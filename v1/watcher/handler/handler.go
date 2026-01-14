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
	workflows "github.com/smartcontractkit/cre-workflow-utils"
	dtav1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
)

var (
	CompleteRequestProcessing    string = "completeRequestProcessing"
	DTARequestManagement         string = "DTARequestManagement"
	DTARequestSettlement         string = "DTARequestSettlement"
	DTASettlementOpenedEventName string = "DTASettlementOpened"
	WorkflowDomain               string = "dta"
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
func OnLog(cfg *workflows.Config, rt cre.Runtime, payload *evm.Log) (string, error) {

	trigger := workflows.BuildTrigger(cfg.ChainID, payload)
	event, err := workflows.BuildEvent(rt, cfg, payload)
	if err != nil {
		return "", err
	}

	var referenceData *workflows.ReferenceData
	if cfg.DetectEventTriggerConfig.ContractName == DTARequestManagement {
		referenceData, err = buildReferenceDataFromDTARequestManagementEvent(rt, cfg, event)
		if err != nil {
			return "", err
		}
	} else if cfg.DetectEventTriggerConfig.ContractName == DTARequestSettlement && cfg.DetectEventTriggerConfig.ContractEventName == DTASettlementOpenedEventName {
		referenceData, err = buildReferenceDataFromDTASettlementOpenedEvent(rt, cfg, trigger, event)
		if err != nil {
			return "", err
		}
	}

	verifiableEvent, err := workflows.BuildAndHashVerifiableEvent(&WorkflowDomain, trigger, event, referenceData)
	if err != nil {
		return "", err
	}

	return workflows.GenerateAndPostReport(cfg, rt, verifiableEvent)
}

func buildReferenceDataFromDTARequestManagementEvent(rt cre.Runtime, cfg *workflows.Config, event workflows.Event) (*workflows.ReferenceData, error) {
	var onChainReferenceData []workflows.OnChainReferenceData
	getDistributorRequestInput, err := buildGetDistributorRequestInputFromDTARequestManagementEvent(rt, cfg, event)
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

	var distributorRequest *dtav1.DistributorRequest
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

func buildReferenceDataFromDTASettlementOpenedEvent(rt cre.Runtime, cfg *workflows.Config, trigger workflows.Trigger, event workflows.Event) (*workflows.ReferenceData, error) {
	var referenceData *workflows.ReferenceData
	var onChainReferenceData []workflows.OnChainReferenceData
	getDistributorRequestInput, err := buildGetDistributorRequestInputFromDTASettlementOpenedEvent(rt, cfg, event.Args)
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

	fundTokenDataRequest, err := buildFundTokenDataRequestFromDTASettlementOpenedEvent(rt, cfg, event.Args)
	if err != nil {
		return nil, err
	}
	fundTokenDataRef, err := fetchAndDecodeFundToken(rt, fundTokenDataRequest)
	if err != nil {
		return nil, err
	}
	onChainReferenceData = append(onChainReferenceData, fundTokenDataRef)
	fundTokenData, ok := fundTokenDataRef.Data["fund_token_data"].(dtav1.FundTokenData)
	if !ok {
		return nil, fmt.Errorf("fund_token_data not found or not a FundTokenData")
	}
	referenceData = &workflows.ReferenceData{
		OnChain: onChainReferenceData,
		OffChain: []workflows.OffChainReferenceData{
			workflows.GetCurrencyCodeAsOffChainReferenceData(fundTokenData.PaymentInfo.OffChainPaymentCurrency),
		},
	}
	if cfg.DetectEventTriggerConfig.ContractEventName == DTASettlementOpenedEventName {
		paymentRequest, err := buildPaymentRequest(rt, cfg, trigger, event, fundTokenData)
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

func buildGetDistributorRequestInputFromDTASettlementOpenedEvent(rt cre.Runtime, cfg *workflows.Config, params map[string]any) (*GetDistributorRequestInput, error) {
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

func buildGetDistributorRequestInputFromDTARequestManagementEvent(rt cre.Runtime, cfg *workflows.Config, event workflows.Event) (*GetDistributorRequestInput, error) {
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

	var requestIdBytes []byte
	switch v := event.Args["request_id"].(type) {
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
		DtaAddr:                     event.ContractAddress,
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

	var distributorRequest dtav1.DistributorRequest
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

func buildFundTokenDataRequestFromDTASettlementOpenedEvent(rt cre.Runtime, cfg *workflows.Config, params map[string]any) (GetFundTokenInput, error) {
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

func buildFundTokenDataRequestFromDTARequestManagementEvent(rt cre.Runtime, cfg *workflows.Config, event workflows.Event, distributorRequest *dtav1.DistributorRequest) (*GetFundTokenInput, error) {
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

	var fundAdminAddr string
	if distributorRequest != nil {
		fundAdminAddr = distributorRequest.FundAdminAddr.Hex()
	} else {
		fundAdminAddr, ok = event.Args["fund_admin_addr"].(string)
		if !ok {
			rt.Logger().Warn("fund_admin_addr not found or not a string, skipping getFundToken call")
			return nil, nil
		}
	}

	var fundTokenId []byte
	if distributorRequest != nil {
		fundTokenId = distributorRequest.FundTokenId[:]
	} else {
		switch v := event.Args["fund_token_id"].(type) {
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
		DtaAddr:            event.ContractAddress,
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
	// First try to marshal the struct to JSON, then unmarshal to dtav1.FundTokenData
	fundTokenDataStruct := vals[1]
	fundTokenDataBytes, err := json.Marshal(fundTokenDataStruct)
	if err != nil {
		return workflows.OnChainReferenceData{}, fmt.Errorf("failed to marshal fundTokenData struct: %w", err)
	}

	var fundTokenData dtav1.FundTokenData
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

func buildPaymentRequest(rt cre.Runtime, cfg *workflows.Config, trigger workflows.Trigger, event workflows.Event, fundTokenData dtav1.FundTokenData) (workflows.PaymentRequest, error) {
	abiJSON, _ := workflows.GetContractABI(cfg, DTARequestSettlement)
	parsedABI, err := gethAbi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return workflows.PaymentRequest{}, err
	}

	callbackMethod, ok := parsedABI.Methods[CompleteRequestProcessing]
	if !ok {
		return workflows.PaymentRequest{}, fmt.Errorf("%s method not found", CompleteRequestProcessing)
	}

	expiration := event.BlockTimestamp.Unix() + int64(time.Hour)

	decodedEvent, err := decodeDTASettlementOpened(event.Args)
	if err != nil {
		return workflows.PaymentRequest{}, fmt.Errorf("failed to decode DTA settlement opened event: %w", err)
	}

	currencyCode := workflows.GetCurrencyCode(fundTokenData.PaymentInfo.OffChainPaymentCurrency)

	var sender string
	var receiver string
	switch decodedEvent.RequestType {
	case dtav1.DistributorRequestTypeSubscription:
		sender = decodedEvent.DistributorWalletAddr.Hex()
		receiver = decodedEvent.FundAdminAddr.Hex()
	case dtav1.DistributorRequestTypeRedemption:
		sender = decodedEvent.FundAdminAddr.Hex()
		receiver = decodedEvent.DistributorWalletAddr.Hex()
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
		ApplicationType: WorkflowDomain,
		ApplicationAddr: event.ContractAddress,
		E2EID:           decodedEvent.RequestId.Hex(),
		Sender:          sender,
		Receiver:        receiver,
		Currency:        currencyCode,
		ChainID:         trigger.ChainID,
		Amount:          amount,
		Expiration:      &expiration,
		CustomCallback: &workflows.PaymentCallback{
			ContractAddress:   event.ContractAddress,
			FunctionName:      CompleteRequestProcessing,
			FunctionSignature: callbackMethod.Sig,
		},
	}, nil
}

func decodeDTASettlementOpened(params map[string]any) (dtav1.DTASettlementOpened, error) {
	// Convert map[string]any to map[string]string (same conversion used in decode.go)
	stringParams := make(map[string]string, len(params))
	for k, v := range params {
		stringParams[k] = fmt.Sprintf("%v", v)
	}

	// Use the generated decoder from decode_gen.go via eventDecoders
	decoder, ok := dtav1.EventDecoders()[dtav1.EventDTASettlementOpened]
	if !ok {
		return dtav1.DTASettlementOpened{}, fmt.Errorf("decoder not found for DTASettlementOpened")
	}

	// Call the decoder (txHash not needed for this event)
	concrete, err := decoder(stringParams, "")
	if err != nil {
		return dtav1.DTASettlementOpened{}, fmt.Errorf("decode DTASettlementOpened: %w", err)
	}

	// Type assert to the concrete type
	event, ok := concrete.(*dtav1.DTASettlementOpened)
	if !ok {
		return dtav1.DTASettlementOpened{}, fmt.Errorf("decoded event is not DTASettlementOpened")
	}

	return *event, nil
}
