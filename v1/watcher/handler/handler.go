package handler

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	gethAbi "github.com/ethereum/go-ethereum/accounts/abi"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	apiModels "github.com/smartcontractkit/crec-api-go/models"
	workflows "github.com/smartcontractkit/crec-workflow-utils"

	dtaevents "github.com/smartcontractkit/crec-sdk-ext-dta/v1/events"
)

var (
	CompleteRequestProcessing         string = "completeRequestProcessing"
	DTARequestManagement              string = "DTARequestManagement"
	DTARequestSettlement              string = "DTARequestSettlement"
	DTASettlementOpenedEventSignature string = "DTASettlementOpened(address,bytes32,uint8,address,uint64,address,bytes32,address,uint256,uint256,uint8)"
	DTASettlementClosedEventSignature string = "DTASettlementClosed(address,bytes32,uint8,address,uint64,address,bytes32,bool,bytes)"
	WorkflowService                   string = "dta.v1"
)

type GetDistributorRequestInput struct {
	GetDistributorRequestMethod gethAbi.Method
	DtaChainSelector            string
	DtaAddr                     string
	RequestId                   []byte
}

type GetFundTokenInput struct {
	GetFundTokenMethod gethAbi.Method
	DtaChainSelector   string
	DtaAddr            string
	FundAdminAddr      string
	FundTokenId        []byte
}

// OnLog processes EVM log events from DTA contracts.
// It decodes event parameters, composes workflow metadata, and posts signed events.
func OnLog(cfg *workflows.Config, rt cre.Runtime, payload *evm.Log, confidence apiModels.ConfidenceLevel) (string, error) {

	event, err := workflows.BuildEVMEventFromLog(rt, cfg, payload, confidence)
	if err != nil {
		return "", err
	}
	if event == nil {
		return "", fmt.Errorf("event is nil")
	}

	var referenceData *workflows.ReferenceData
	if cfg.DetectEventTriggerConfig.ContractName == DTARequestManagement {
		referenceData, err = buildReferenceDataFromDTARequestManagementEvent(rt, cfg, *event)
		if err != nil {
			return "", err
		}
	} else if cfg.DetectEventTriggerConfig.ContractName == DTARequestSettlement &&
		(event.EventSignature == DTASettlementOpenedEventSignature || event.EventSignature == DTASettlementClosedEventSignature) {
		referenceData, err = buildReferenceDataFromDTASettlementEvent(rt, cfg, *event)
		if err != nil {
			return "", err
		}
	}

	referenceDataBytes, err := json.Marshal(referenceData)
	if err != nil {
		return "", err
	}
	typeAndValue := workflows.TypeAndValue{
		Type:  workflows.RawMessageTypeReferenceData,
		Value: json.RawMessage(referenceDataBytes),
	}
	typeAndValueBytes, err := json.Marshal(typeAndValue)
	if err != nil {
		return "", err
	}
	var referenceDataMap map[string]interface{}
	if err := json.Unmarshal(typeAndValueBytes, &referenceDataMap); err != nil {
		return "", err
	}

	abiJSON, err := workflows.GetContractABI(cfg, cfg.DetectEventTriggerConfig.ContractName)
	if err != nil {
		return "", err
	}
	eventName, err := workflows.GetEventNameFromLog(cfg, payload, abiJSON)
	if err != nil {
		return "", err
	}

	verifiableEvent, err := workflows.BuildVerifiableEventForEVMEvent(rt, cfg, event, cfg.Service, eventName, &referenceDataMap)
	if err != nil {
		return "", err
	}

	encodedVerifiableEvent, err := workflows.EncodeVerifiableEvent(verifiableEvent)
	if err != nil {
		return "", err
	}

	rt.Logger().Info("verifiableEvent", "encodedVerifiableEvent", encodedVerifiableEvent)

	return workflows.SignAndPostVerifiableEvent(rt, cfg, verifiableEvent)
}

func buildReferenceDataFromDTARequestManagementEvent(rt cre.Runtime, cfg *workflows.Config, event apiModels.EVMEvent) (*workflows.ReferenceData, error) {
	var onChainReferenceData []workflows.OnChainReferenceData
	getDistributorRequestInput, err := buildGetDistributorRequestInputFromDTARequestManagementEvent(cfg, event)
	if err != nil {
		rt.Logger().Warn("failed to build getDistributorRequestInputFromDTARequestManagementEvent, skipping distributor request", "error", err)
	}
	var distributorRequestRef *workflows.OnChainReferenceData
	if getDistributorRequestInput != nil {
		distributorRequestRef, err = fetchAndDecodeDistributorRequest(rt, *getDistributorRequestInput)
		if err != nil {
			rt.Logger().Warn("failed to fetch and decode distributor request, skipping distributor request", "error", err)
		}
		if distributorRequestRef != nil {
			onChainReferenceData = append(onChainReferenceData, *distributorRequestRef)
		}
	}

	var distributorRequest *dtaevents.DistributorRequest
	if distributorRequestRef != nil {
		distributorRequestBytes, err := json.Marshal(distributorRequestRef.Data["distributor_request"])
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(distributorRequestBytes, &distributorRequest); err != nil {
			return nil, err
		}
	}

	fundTokenDataRequest, err := buildFundTokenDataRequestFromDTARequestManagementEvent(rt, cfg, event, distributorRequest)
	if err != nil {
		return nil, err
	}
	if fundTokenDataRequest != nil {
		fundTokenDataRef, err := fetchAndDecodeFundToken(rt, *fundTokenDataRequest)
		if err != nil {
			return nil, err
		}
		onChainReferenceData = append(onChainReferenceData, fundTokenDataRef)
	}
	return &workflows.ReferenceData{
		OnChain: onChainReferenceData,
	}, nil
}

func buildReferenceDataFromDTASettlementEvent(rt cre.Runtime, cfg *workflows.Config, event apiModels.EVMEvent) (*workflows.ReferenceData, error) {
	if event.Params == nil {
		return nil, fmt.Errorf("event params are nil")
	}
	var referenceData *workflows.ReferenceData
	var onChainReferenceData []workflows.OnChainReferenceData
	getDistributorRequestInput, err := buildGetDistributorRequestInputFromDTASettlementEvent(cfg, *event.Params)
	if err != nil {
		return nil, err
	}
	if getDistributorRequestInput != nil {
		distributorRequestRef, err := fetchAndDecodeDistributorRequest(rt, *getDistributorRequestInput)
		if err != nil {
			return nil, err
		}
		if distributorRequestRef != nil {
			onChainReferenceData = append(onChainReferenceData, *distributorRequestRef)
		}
	}

	fundTokenDataRequest, err := buildFundTokenDataRequestFromDTASettlementEvent(cfg, *event.Params)
	if err != nil {
		return nil, err
	}
	fundTokenDataRef, err := fetchAndDecodeFundToken(rt, fundTokenDataRequest)
	if err != nil {
		return nil, err
	}
	onChainReferenceData = append(onChainReferenceData, fundTokenDataRef)
	fundTokenData, ok := fundTokenDataRef.Data["fund_token_data"].(dtaevents.FundTokenData)
	if !ok {
		return nil, fmt.Errorf("fund_token_data not found or not a FundTokenData")
	}
	referenceData = &workflows.ReferenceData{
		OnChain: onChainReferenceData,
		OffChain: []workflows.OffChainReferenceData{
			workflows.GetCurrencyCodeAsOffChainReferenceData(fundTokenData.PaymentInfo.OffChainPaymentCurrency),
		},
	}
	if event.EventSignature == DTASettlementOpenedEventSignature {
		paymentRequest, err := buildPaymentRequest(cfg, event, fundTokenData)
		if err != nil {
			return nil, err
		}
		paymentRequestBytes, err := json.Marshal(paymentRequest)
		if err != nil {
			return nil, err
		}
		referenceData.Requests = []workflows.TypeAndValue{
			{
				Type:  workflows.RawMessageTypePaymentRequest,
				Value: json.RawMessage(paymentRequestBytes),
			},
		}
	}
	return referenceData, nil
}

func buildGetDistributorRequestInputFromDTASettlementEvent(cfg *workflows.Config, params map[string]any) (*GetDistributorRequestInput, error) {
	abiJSON, err := workflows.GetContractABI(cfg, DTARequestManagement)
	if err != nil {
		return nil, err
	}
	parsedABI, err := gethAbi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	getDistributorRequestMethod, ok := parsedABI.Methods["getDistributorRequest"]
	if !ok {
		return nil, fmt.Errorf("getDistributorRequest method not found")
	}

	dtaChainSelectorUint64, ok := params["dta_chain_selector"].(uint64)
	if !ok {
		return nil, fmt.Errorf("dta_chain_selector not found or not a uint64")
	}
	dtaChainSelectorStr := strconv.FormatUint(dtaChainSelectorUint64, 10)
	dtaAddr, ok := params["dta_addr"].(string)
	if !ok {
		return nil, fmt.Errorf("dta_addr not found or not a string")
	}

	var requestIdBytes []byte
	switch v := params["request_id"].(type) {
	case []byte:
		requestIdBytes = v
	case string:
		hexStr := strings.TrimPrefix(v, "0x")
		requestIdBytes, err = hex.DecodeString(hexStr)
		if err != nil {
			return nil, fmt.Errorf("failed to decode request_id hex string: %w", err)
		}
	default:
		return nil, fmt.Errorf("request_id has unsupported type: %T", v)
	}
	if len(requestIdBytes) != 32 {
		return nil, fmt.Errorf("request_id must be 32 bytes, got %d bytes", len(requestIdBytes))
	}

	return &GetDistributorRequestInput{
		GetDistributorRequestMethod: getDistributorRequestMethod,
		DtaChainSelector:            dtaChainSelectorStr,
		DtaAddr:                     dtaAddr,
		RequestId:                   requestIdBytes,
	}, nil
}

func buildGetDistributorRequestInputFromDTARequestManagementEvent(cfg *workflows.Config, event apiModels.EVMEvent) (*GetDistributorRequestInput, error) {
	abiJSON, err := workflows.GetContractABI(cfg, DTARequestManagement)
	if err != nil {
		return nil, err
	}
	parsedABI, err := gethAbi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	getDistributorRequestMethod, ok := parsedABI.Methods["getDistributorRequest"]
	if !ok {
		return nil, fmt.Errorf("getDistributorRequest method not found")
	}

	if event.Params == nil {
		return nil, fmt.Errorf("event params are nil")
	}
	params := *event.Params

	var requestIdBytes []byte
	switch v := params["request_id"].(type) {
	case []byte:
		requestIdBytes = v
	case string:
		hexStr := strings.TrimPrefix(v, "0x")
		requestIdBytes, err = hex.DecodeString(hexStr)
		if err != nil {
			return nil, fmt.Errorf("failed to decode request_id hex string: %w", err)
		}
	default:
		return nil, fmt.Errorf("request_id has unsupported type: %T", v)
	}
	if len(requestIdBytes) != 32 {
		return nil, fmt.Errorf("request_id must be 32 bytes, got %d bytes", len(requestIdBytes))
	}

	return &GetDistributorRequestInput{
		GetDistributorRequestMethod: getDistributorRequestMethod,
		DtaChainSelector:            cfg.ChainSelector,
		DtaAddr:                     event.Address,
		RequestId:                   requestIdBytes,
	}, nil
}

func fetchAndDecodeDistributorRequest(rt cre.Runtime, request GetDistributorRequestInput) (*workflows.OnChainReferenceData, error) {
	to := gethCommon.HexToAddress(request.DtaAddr)
	methodID := crypto.Keccak256([]byte("getDistributorRequest(bytes32)"))[:4]
	callData := make([]byte, 4+32)
	copy(callData[:4], methodID)
	copy(callData[4:4+32], request.RequestId)

	chainSelector, err := strconv.ParseUint(request.DtaChainSelector, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid chain selector: %w", err)
	}

	cli := &evm.Client{ChainSelector: chainSelector}
	callContractReply, err := cli.CallContract(rt, &evm.CallContractRequest{
		Call: &evm.CallMsg{
			To:   to.Bytes(),
			Data: callData,
		},
	}).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to call getDistributorRequest: %w", err)
	}
	if len(callContractReply.Data) == 0 {
		return nil, fmt.Errorf("empty getDistributorRequest response")
	}

	vals, err := request.GetDistributorRequestMethod.Outputs.UnpackValues(callContractReply.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getDistributorRequest response: %w", err)
	}
	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 output, got %d", len(vals))
	}

	distributorRequestStruct := vals[0]
	distributorRequestBytes, err := json.Marshal(distributorRequestStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal distributorRequest struct: %w", err)
	}

	var distributorRequest dtaevents.DistributorRequest
	if err := json.Unmarshal(distributorRequestBytes, &distributorRequest); err != nil {
		return nil, fmt.Errorf("failed to unmarshal distributorRequest to struct: %w", err)
	}

	return &workflows.OnChainReferenceData{
		Source: workflows.OnChainReferenceDataSource{
			ContractAddress:           request.DtaAddr,
			ContractFunctionSignature: request.GetDistributorRequestMethod.Sig,
			CallData:                  "0x" + hex.EncodeToString(callData),
			Block:                     "latest",
		},
		Data: map[string]any{
			"distributor_request": distributorRequest,
		},
	}, nil
}

func buildFundTokenDataRequestFromDTASettlementEvent(cfg *workflows.Config, params map[string]any) (GetFundTokenInput, error) {
	abiJSON, err := workflows.GetContractABI(cfg, DTARequestManagement)
	if err != nil {
		return GetFundTokenInput{}, err
	}
	parsedABI, err := gethAbi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return GetFundTokenInput{}, err
	}
	getFundTokenMethod, ok := parsedABI.Methods["getFundToken"]
	if !ok {
		return GetFundTokenInput{}, fmt.Errorf("getFundToken method not found")
	}

	dtaChainSelectorUint64, ok := params["dta_chain_selector"].(uint64)
	if !ok {
		return GetFundTokenInput{}, fmt.Errorf("dta_chain_selector not found or not a uint64")
	}
	dtaChainSelectorStr := strconv.FormatUint(dtaChainSelectorUint64, 10)
	dtaAddr, ok := params["dta_addr"].(string)
	if !ok {
		return GetFundTokenInput{}, fmt.Errorf("dta_addr not found or not a string")
	}

	fundAdminAddr, ok := params["fund_admin_addr"].(string)
	if !ok {
		return GetFundTokenInput{}, fmt.Errorf("fund_admin_addr not found or not a string")
	}

	var fundTokenId []byte

	switch v := params["fund_token_id"].(type) {
	case []byte:
		fundTokenId = v
	case string:
		hexStr := strings.TrimPrefix(v, "0x")
		var err error
		fundTokenId, err = hex.DecodeString(hexStr)
		if err != nil {
			return GetFundTokenInput{}, fmt.Errorf("failed to decode fund_token_id hex string %q: %w", v, err)
		}
	default:
		return GetFundTokenInput{}, fmt.Errorf("fund_token_id has unsupported type: %T, value: %v", v, v)
	}

	if len(fundTokenId) != 32 {
		return GetFundTokenInput{}, fmt.Errorf("fund_token_id must be 32 bytes, got %d bytes", len(fundTokenId))
	}

	return GetFundTokenInput{
		GetFundTokenMethod: getFundTokenMethod,
		DtaChainSelector:   dtaChainSelectorStr,
		DtaAddr:            dtaAddr,
		FundAdminAddr:      fundAdminAddr,
		FundTokenId:        fundTokenId,
	}, nil
}

func buildFundTokenDataRequestFromDTARequestManagementEvent(rt cre.Runtime, cfg *workflows.Config, event apiModels.EVMEvent, distributorRequest *dtaevents.DistributorRequest) (*GetFundTokenInput, error) {
	abiJSON, err := workflows.GetContractABI(cfg, DTARequestManagement)
	if err != nil {
		return nil, err
	}
	parsedABI, err := gethAbi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}

	getFundTokenMethod, ok := parsedABI.Methods["getFundToken"]
	if !ok {
		return nil, fmt.Errorf("getFundToken method not found")
	}

	if event.Params == nil {
		return nil, fmt.Errorf("event params are nil")
	}
	params := *event.Params

	var fundAdminAddr string
	if distributorRequest != nil {
		fundAdminAddr = distributorRequest.FundAdminAddr.Hex()
	} else {
		fundAdminAddr, ok = params["fund_admin_addr"].(string)
		if !ok {
			rt.Logger().Warn("fund_admin_addr not found or not a string, skipping getFundToken call")
			return nil, nil
		}
	}

	var fundTokenId []byte
	if distributorRequest != nil {
		fundTokenId = distributorRequest.FundTokenId[:]
	} else {
		switch v := params["fund_token_id"].(type) {
		case []byte:
			fundTokenId = v
		case string:
			hexStr := strings.TrimPrefix(v, "0x")
			var err error
			fundTokenId, err = hex.DecodeString(hexStr)
			if err != nil {
				rt.Logger().Warn("failed to decode fund_token_id hex string, skipping getFundToken call", "fund_token_id", v, "error", err)
				return nil, nil
			}
		default:
			rt.Logger().Warn("fund_token_id has unsupported type, skipping getFundToken call", "type", fmt.Sprintf("%T", v), "value", v)
			return nil, nil
		}
	}

	if len(fundTokenId) != 32 {
		rt.Logger().Warn("fund_token_id must be 32 bytes, skipping getFundToken call", "byteCount", len(fundTokenId))
		return nil, nil
	}

	return &GetFundTokenInput{
		GetFundTokenMethod: getFundTokenMethod,
		DtaChainSelector:   cfg.ChainSelector,
		DtaAddr:            event.Address,
		FundAdminAddr:      fundAdminAddr,
		FundTokenId:        fundTokenId,
	}, nil
}

func fetchAndDecodeFundToken(rt cre.Runtime, request GetFundTokenInput) (workflows.OnChainReferenceData, error) {
	to := gethCommon.HexToAddress(request.DtaAddr)
	methodID := crypto.Keccak256([]byte("getFundToken(address,bytes32)"))[:4]
	callData := make([]byte, 4+32+32)
	copy(callData[:4], methodID)

	// Properly left-pad the address to 32 bytes
	fundAdminAddrBytes := gethCommon.HexToAddress(request.FundAdminAddr).Bytes() // 20 bytes
	fundAdminAddrPadded := make([]byte, 32)
	copy(fundAdminAddrPadded[32-len(fundAdminAddrBytes):], fundAdminAddrBytes)
	copy(callData[4:4+32], fundAdminAddrPadded)

	// Ensure fundTokenId is 32 bytes (already assumed, but for safety)
	fundTokenIdPadded := make([]byte, 32)
	copy(fundTokenIdPadded, request.FundTokenId)
	copy(callData[4+32:4+32+32], fundTokenIdPadded)

	chainSelector, err := strconv.ParseUint(request.DtaChainSelector, 10, 64)
	if err != nil {
		return workflows.OnChainReferenceData{}, fmt.Errorf("invalid chain selector: %w", err)
	}

	cli := &evm.Client{ChainSelector: chainSelector}
	callContractReply, err := cli.CallContract(rt, &evm.CallContractRequest{
		Call: &evm.CallMsg{
			To:   to.Bytes(),
			Data: callData,
		},
	}).Await()
	if err != nil {
		rt.Logger().Error("failed to call getFundToken", "error", err)
		return workflows.OnChainReferenceData{}, err
	}
	if len(callContractReply.Data) == 0 {
		return workflows.OnChainReferenceData{}, fmt.Errorf("empty getFundToken response")
	}

	vals, err := request.GetFundTokenMethod.Outputs.UnpackValues(callContractReply.Data)
	if err != nil {
		return workflows.OnChainReferenceData{}, fmt.Errorf("failed to unpack getFundToken response: %w", err)
	}
	if len(vals) != 2 {
		return workflows.OnChainReferenceData{}, fmt.Errorf("expected 2 outputs, got %d", len(vals))
	}
	enabled, ok := vals[0].(bool)
	if !ok {
		return workflows.OnChainReferenceData{}, fmt.Errorf("enabled is not a bool")
	}

	// fundTokenData is a struct, not a map - convert it to map[string]any
	// First try to marshal the struct to JSON, then unmarshal to dtaevents.FundTokenData
	fundTokenDataStruct := vals[1]
	fundTokenDataBytes, err := json.Marshal(fundTokenDataStruct)
	if err != nil {
		return workflows.OnChainReferenceData{}, fmt.Errorf("failed to marshal fundTokenData struct: %w", err)
	}

	var fundTokenData dtaevents.FundTokenData
	if err := json.Unmarshal(fundTokenDataBytes, &fundTokenData); err != nil {
		return workflows.OnChainReferenceData{}, fmt.Errorf("failed to unmarshal fundTokenData to struct: %w", err)
	}

	return workflows.OnChainReferenceData{
		Source: workflows.OnChainReferenceDataSource{
			ContractAddress:           request.DtaAddr,
			ContractFunctionSignature: request.GetFundTokenMethod.Sig,
			CallData:                  "0x" + hex.EncodeToString(callData),
			Block:                     "latest",
		},
		Data: map[string]any{
			"enabled":         enabled,
			"fund_token_data": fundTokenData,
		},
	}, nil
}

func buildPaymentRequest(cfg *workflows.Config, event apiModels.EVMEvent, fundTokenData dtaevents.FundTokenData) (workflows.PaymentRequest, error) {
	abiJSON, _ := workflows.GetContractABI(cfg, DTARequestSettlement)
	parsedABI, err := gethAbi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return workflows.PaymentRequest{}, err
	}

	callbackMethod, ok := parsedABI.Methods[CompleteRequestProcessing]
	if !ok {
		return workflows.PaymentRequest{}, fmt.Errorf("%s method not found", CompleteRequestProcessing)
	}

	expiration := int64(event.BlockTimestamp) + int64(time.Hour)

	if event.Params == nil {
		return workflows.PaymentRequest{}, fmt.Errorf("event params are nil")
	}

	decodedEvent, err := decodeDTASettlementOpened(*event.Params)
	if err != nil {
		return workflows.PaymentRequest{}, fmt.Errorf("failed to decode DTA settlement opened event: %w", err)
	}

	currencyCode := workflows.GetCurrencyCode(fundTokenData.PaymentInfo.OffChainPaymentCurrency)

	var sender string
	var receiver string
	switch decodedEvent.RequestType {
	case dtaevents.DistributorRequestTypeSubscription:
		sender = decodedEvent.DistributorAddr.Hex()
		receiver = decodedEvent.FundAdminAddr.Hex()
	case dtaevents.DistributorRequestTypeRedemption:
		sender = decodedEvent.FundAdminAddr.Hex()
		receiver = decodedEvent.DistributorAddr.Hex()
	default:
		return workflows.PaymentRequest{}, fmt.Errorf("unknown request type: %d", decodedEvent.RequestType)
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(fundTokenData.NavFeedDecimals)), nil)
	quotient := new(big.Int).Div(decodedEvent.Amount, divisor)
	if !quotient.IsInt64() {
		return workflows.PaymentRequest{}, fmt.Errorf("amount overflow: %s", quotient.String())
	}
	amount := workflows.Fixed2(quotient.Int64())

	return workflows.PaymentRequest{
		ApplicationType: WorkflowService,
		ApplicationAddr: event.Address,
		E2EID:           decodedEvent.RequestId.Hex(),
		Sender:          sender,
		Receiver:        receiver,
		Currency:        currencyCode,
		Amount:          amount,
		Expiration:      &expiration,
		CustomCallback: &workflows.PaymentCallback{
			ContractAddress:   event.Address,
			FunctionName:      CompleteRequestProcessing,
			FunctionSignature: callbackMethod.Sig,
		},
	}, nil
}

func decodeDTASettlementOpened(params map[string]any) (dtaevents.DTASettlementOpened, error) {
	// Convert map[string]any to map[string]string (same conversion used in decode.go)
	stringParams := make(map[string]string, len(params))
	for k, v := range params {
		stringParams[k] = fmt.Sprintf("%v", v)
	}

	// Use the generated decoder from decode_gen.go via eventDecoders
	decoder, ok := dtaevents.EventDecoders()[dtaevents.EventDTASettlementOpened]
	if !ok {
		return dtaevents.DTASettlementOpened{}, fmt.Errorf("decoder not found for DTASettlementOpened")
	}

	// Call the decoder (txHash not needed for this event)
	concrete, err := decoder(stringParams, "")
	if err != nil {
		return dtaevents.DTASettlementOpened{}, fmt.Errorf("decode DTASettlementOpened: %w", err)
	}

	// Type assert to the concrete type
	event, ok := concrete.(*dtaevents.DTASettlementOpened)
	if !ok {
		return dtaevents.DTASettlementOpened{}, fmt.Errorf("decoded event is not DTASettlementOpened")
	}

	return *event, nil
}
