package v1

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	apiClient "github.com/smartcontractkit/crec-api-go/client"
	"github.com/smartcontractkit/crec-sdk-ext-dta/parsing"
)

// ConcreteEvent represents any decoded concrete event payload.
type ConcreteEvent interface{}

// DecodedEvent wraps WatcherEventPayload with a decoded ConcreteEvent.
type DecodedEvent struct {
	apiClient.WatcherEventPayload
	ConcreteEvent ConcreteEvent
}

// EventName returns the parsed event name from the payload.
func (e DecodedEvent) EventName() EventName {
	name, ok := parseEventName(e.Event.EventName)
	if !ok {
		return EventUnknown
	}
	return name
}

// eventDecoder is a function that decodes event parameters into a concrete event type.
type eventDecoder func(params map[string]string, txHash string) (ConcreteEvent, error)

// eventDecoders maps event names to their decoder functions.
var eventDecoders = map[EventName]eventDecoder{
	// DTARequestManagementU events
	EventDistributorRegistered:        decodeDistributorRegistered,
	EventDistributorRequestCanceled:   decodeDistributorRequestCanceled,
	EventDistributorRequestProcessed:  decodeDistributorRequestProcessed,
	EventDistributorRequestProcessing: decodeDistributorRequestProcessing,
	EventFundAdminRegistered:          decodeFundAdminRegistered,
	EventFundTokenAllowlistUpdated:    decodeFundTokenAllowlistUpdated,
	EventFundTokenRegistered:          decodeFundTokenRegistered,
	EventInitialized:                  decodeInitialized,
	EventInvalidDTARequestSettlement:  decodeInvalidDTARequestSettlement,
	EventMessageFailed:                decodeMessageFailed,
	EventNativeFundsRecovered:         decodeNativeFundsRecovered,
	EventOwnershipTransferred:         decodeOwnershipTransferred,
	EventRedemptionRequested:          decodeRedemptionRequested,
	EventSubscriptionRequested:        decodeSubscriptionRequested,

	// NAV events
	EventAnswerUpdated: decodeAnswerUpdated,

	// DTARequestSettlementU events
	EventCCIPMessageRecvFailed:                decodeCCIPMessageRecvFailed,
	EventDTAAdded:                             decodeDTAAdded,
	EventDTARemoved:                           decodeDTARemoved,
	EventDTASettlementClosed:                  decodeDTASettlementClosed,
	EventDTASettlementOpened:                  decodeDTASettlementOpened,
	EventEmptyRequestType:                     decodeEmptyRequestType,
	EventInvalidSubscriptionCrossChainPayment: decodeInvalidSubscriptionCrossChainPayment,
	EventSettlementFailed:                     decodeSettlementFailed,
	EventTokenWithdrawn:                       decodeTokenWithdrawn,
	EventUnauthorizedSenderDTA:                decodeUnauthorizedSenderDTA,
}

// DecodeFromEvent extracts the WatcherEventPayload from an apiClient.Event and decodes
// the ConcreteEvent based on the event type.
func DecodeFromEvent(_ context.Context, event apiClient.Event) (DecodedEvent, error) {
	payload, err := event.Payload.AsWatcherEventPayload()
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("extract watcher payload: %w", err)
	}

	// Convert Data map to string params for decoders
	params := make(map[string]string, len(payload.Event.Data))
	for k, v := range payload.Event.Data {
		params[k] = fmt.Sprintf("%v", v)
	}

	name, ok := parseEventName(payload.Event.EventName)
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event name: %s", payload.Event.EventName)
	}
	decoder, ok := eventDecoders[name]
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event decoder for event name: %s", name)
	}

	concrete, err := decoder(params, payload.Transaction.Hash)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("decode %s: %w", name, err)
	}

	return DecodedEvent{
		WatcherEventPayload: payload,
		ConcreteEvent:       concrete,
	}, nil
}

// ============================================================================
// Event Decoders
// ============================================================================

func decodeDistributorRegistered(params map[string]string, _ string) (ConcreteEvent, error) {
	return &DistributorRegistered{
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
	}, nil
}

func decodeDistributorRequestCanceled(params map[string]string, _ string) (ConcreteEvent, error) {
	return &DistributorRequestCanceled{
		FundTokenId:     common.HexToHash(params["fund_token_id"]),
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
		RequestId:       common.HexToHash(params["request_id"]),
	}, nil
}

func decodeDistributorRequestProcessed(params map[string]string, _ string) (ConcreteEvent, error) {
	shares, err := parsing.ScientificNotationToBigInt(params["shares"])
	if err != nil {
		return nil, fmt.Errorf("parse shares %q: %w", params["shares"], err)
	}
	status, err := parsing.ScientificNotationToUint8(params["status"])
	if err != nil {
		return nil, fmt.Errorf("parse status %q: %w", params["status"], err)
	}
	return &DistributorRequestProcessed{
		RequestId: common.HexToHash(params["request_id"]),
		Shares:    shares,
		Status:    RequestStatus(status),
		Error:     []byte(params["error"]),
	}, nil
}

func decodeDistributorRequestProcessing(params map[string]string, _ string) (ConcreteEvent, error) {
	shares, err := parsing.ScientificNotationToBigInt(params["shares"])
	if err != nil {
		return nil, fmt.Errorf("parse shares %q: %w", params["shares"], err)
	}
	amount, err := parsing.ScientificNotationToBigInt(params["amount"])
	if err != nil {
		return nil, fmt.Errorf("parse amount %q: %w", params["amount"], err)
	}
	return &DistributorRequestProcessing{
		FundTokenId:     common.HexToHash(params["fund_token_id"]),
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
		RequestId:       common.HexToHash(params["request_id"]),
		Shares:          shares,
		Amount:          amount,
	}, nil
}

func decodeFundAdminRegistered(params map[string]string, _ string) (ConcreteEvent, error) {
	return &FundAdminRegistered{
		FundAdminAddr: common.HexToAddress(params["fund_admin_addr"]),
	}, nil
}

func decodeFundTokenAllowlistUpdated(params map[string]string, _ string) (ConcreteEvent, error) {
	allowed, err := strconv.ParseBool(params["allowed"])
	if err != nil {
		return nil, fmt.Errorf("parse allowed %q: %w", params["allowed"], err)
	}
	return &FundTokenAllowlistUpdated{
		FundAdminAddr:   common.HexToAddress(params["fund_admin_addr"]),
		FundTokenId:     common.HexToHash(params["fund_token_id"]),
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
		Allowed:         allowed,
	}, nil
}

func decodeFundTokenRegistered(params map[string]string, _ string) (ConcreteEvent, error) {
	tokenChainSelector, err := parsing.ScientificNotationToUint64(params["token_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse token_chain_selector %q: %w", params["token_chain_selector"], err)
	}
	return &FundTokenRegistered{
		FundAdminAddr:      common.HexToAddress(params["fund_admin_addr"]),
		FundTokenId:        common.HexToHash(params["fund_token_id"]),
		FundTokenAddr:      common.HexToAddress(params["fund_token_addr"]),
		NavAddr:            common.HexToAddress(params["nav_addr"]),
		TokenChainSelector: tokenChainSelector,
	}, nil
}

func decodeInitialized(params map[string]string, _ string) (ConcreteEvent, error) {
	version, err := parsing.ScientificNotationToUint64(params["version"])
	if err != nil {
		return nil, fmt.Errorf("parse version %q: %w", params["version"], err)
	}
	return &Initialized{Version: version}, nil
}

func decodeInvalidDTARequestSettlement(params map[string]string, _ string) (ConcreteEvent, error) {
	actualChainSelector, err := parsing.ScientificNotationToUint64(params["actual_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse actual_chain_selector %q: %w", params["actual_chain_selector"], err)
	}
	return &InvalidDTARequestSettlement{
		FundAdminAddr:                  common.HexToAddress(params["fund_admin_addr"]),
		FundTokenId:                    common.HexToHash(params["fund_token_id"]),
		RequestId:                      common.HexToHash(params["request_id"]),
		ActualChainSelector:            actualChainSelector,
		ActualDTARequestSettlementAddr: common.HexToAddress(params["actual_dta_request_settlement_addr"]),
	}, nil
}

func decodeMessageFailed(params map[string]string, _ string) (ConcreteEvent, error) {
	return &MessageFailed{
		MessageId: common.HexToHash(params["message_id"]),
		Reason:    []byte(params["reason"]),
	}, nil
}

func decodeNativeFundsRecovered(params map[string]string, _ string) (ConcreteEvent, error) {
	amount, err := parsing.ScientificNotationToBigInt(params["amount"])
	if err != nil {
		return nil, fmt.Errorf("parse amount %q: %w", params["amount"], err)
	}
	return &NativeFundsRecovered{
		To:     common.HexToAddress(params["to"]),
		Amount: amount,
	}, nil
}

func decodeOwnershipTransferred(params map[string]string, _ string) (ConcreteEvent, error) {
	return &OwnershipTransferred{
		PreviousOwner: common.HexToAddress(params["previous_owner"]),
		NewOwner:      common.HexToAddress(params["new_owner"]),
	}, nil
}

func decodeRedemptionRequested(params map[string]string, _ string) (ConcreteEvent, error) {
	shares, err := parsing.ScientificNotationToBigInt(params["shares"])
	if err != nil {
		return nil, fmt.Errorf("parse shares %q: %w", params["shares"], err)
	}
	createdAt, err := parsing.ScientificNotationToUint64(params["created_at"])
	if err != nil {
		return nil, fmt.Errorf("parse created_at %q: %w", params["created_at"], err)
	}
	return &RedemptionRequested{
		FundTokenId:     common.HexToHash(params["fund_token_id"]),
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
		RequestId:       common.HexToHash(params["request_id"]),
		Shares:          shares,
		CreatedAt:       createdAt,
	}, nil
}

func decodeSubscriptionRequested(params map[string]string, _ string) (ConcreteEvent, error) {
	amount, err := parsing.ScientificNotationToBigInt(params["amount"])
	if err != nil {
		return nil, fmt.Errorf("parse amount %q: %w", params["amount"], err)
	}
	createdAt, err := parsing.ScientificNotationToUint64(params["created_at"])
	if err != nil {
		return nil, fmt.Errorf("parse created_at %q: %w", params["created_at"], err)
	}
	return &SubscriptionRequested{
		FundTokenId:     common.HexToHash(params["fund_token_id"]),
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
		RequestId:       common.HexToHash(params["request_id"]),
		Amount:          amount,
		CreatedAt:       createdAt,
	}, nil
}

func decodeAnswerUpdated(params map[string]string, _ string) (ConcreteEvent, error) {
	current, err := parsing.ScientificNotationToBigInt(params["current"])
	if err != nil {
		return nil, fmt.Errorf("parse current %q: %w", params["current"], err)
	}
	roundId, err := parsing.ScientificNotationToBigInt(params["roundId"])
	if err != nil {
		return nil, fmt.Errorf("parse roundId %q: %w", params["roundId"], err)
	}
	updatedAt, err := parsing.ScientificNotationToBigInt(params["updatedAt"])
	if err != nil {
		return nil, fmt.Errorf("parse updatedAt %q: %w", params["updatedAt"], err)
	}
	return &AnswerUpdated{Current: current, RoundId: roundId, UpdatedAt: updatedAt}, nil
}

func decodeCCIPMessageRecvFailed(params map[string]string, _ string) (ConcreteEvent, error) {
	return &CCIPMessageRecvFailed{
		MessageId: common.HexToHash(params["message_id"]),
		Reason:    []byte(params["reason"]),
	}, nil
}

func decodeDTAAdded(params map[string]string, _ string) (ConcreteEvent, error) {
	dtaChainSelector, err := parsing.ScientificNotationToUint64(params["dta_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse dta_chain_selector %q: %w", params["dta_chain_selector"], err)
	}
	return &DTAAdded{
		DtaAddr:          common.HexToAddress(params["dta_addr"]),
		DtaChainSelector: dtaChainSelector,
		FundTokenId:      common.HexToHash(params["fund_token_id"]),
		FundTokenAddr:    common.HexToAddress(params["fund_token_addr"]),
	}, nil
}

func decodeDTARemoved(params map[string]string, _ string) (ConcreteEvent, error) {
	dtaChainSelector, err := parsing.ScientificNotationToUint64(params["dta_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse dta_chain_selector %q: %w", params["dta_chain_selector"], err)
	}
	return &DTARemoved{
		DtaAddr:          common.HexToAddress(params["dta_addr"]),
		DtaChainSelector: dtaChainSelector,
		FundTokenId:      common.HexToHash(params["fund_token_id"]),
	}, nil
}

func decodeDTASettlementClosed(params map[string]string, _ string) (ConcreteEvent, error) {
	requestType, err := parsing.ScientificNotationToUint8(params["request_type"])
	if err != nil {
		return nil, fmt.Errorf("parse request_type %q: %w", params["request_type"], err)
	}
	dtaChainSelector, err := parsing.ScientificNotationToUint64(params["dta_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse dta_chain_selector %q: %w", params["dta_chain_selector"], err)
	}
	success, err := strconv.ParseBool(params["success"])
	if err != nil {
		return nil, fmt.Errorf("parse success %q: %w", params["success"], err)
	}
	return &DTASettlementClosed{
		DistributorAddr:  common.HexToAddress(params["distributor_addr"]),
		RequestType:      DistributorRequestType(requestType),
		FundTokenId:      common.HexToHash(params["fund_token_id"]),
		DtaChainSelector: dtaChainSelector,
		DtaAddr:          common.HexToAddress(params["dta_addr"]),
		RequestId:        common.HexToHash(params["request_id"]),
		Success:          success,
		Err:              []byte(params["err"]),
	}, nil
}

func decodeDTASettlementOpened(params map[string]string, _ string) (ConcreteEvent, error) {
	requestType, err := parsing.ScientificNotationToUint8(params["request_type"])
	if err != nil {
		return nil, fmt.Errorf("parse request_type %q: %w", params["request_type"], err)
	}
	dtaChainSelector, err := parsing.ScientificNotationToUint64(params["dta_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse dta_chain_selector %q: %w", params["dta_chain_selector"], err)
	}
	shares, err := parsing.ScientificNotationToBigInt(params["shares"])
	if err != nil {
		return nil, fmt.Errorf("parse shares %q: %w", params["shares"], err)
	}
	amount, err := parsing.ScientificNotationToBigInt(params["amount"])
	if err != nil {
		return nil, fmt.Errorf("parse amount %q: %w", params["amount"], err)
	}
	currency, err := parsing.ScientificNotationToUint8(params["currency"])
	if err != nil {
		return nil, fmt.Errorf("parse currency %q: %w", params["currency"], err)
	}
	return &DTASettlementOpened{
		DistributorAddr:       common.HexToAddress(params["distributor_addr"]),
		RequestType:           DistributorRequestType(requestType),
		FundTokenId:           common.HexToHash(params["fund_token_id"]),
		FundAdminAddr:         common.HexToAddress(params["fund_admin_addr"]),
		DtaChainSelector:      dtaChainSelector,
		DtaAddr:               common.HexToAddress(params["dta_addr"]),
		RequestId:             common.HexToHash(params["request_id"]),
		DistributorWalletAddr: common.HexToAddress(params["distributor_wallet_addr"]),
		Shares:                shares,
		Amount:                amount,
		Currency:              currency,
	}, nil
}

func decodeEmptyRequestType(params map[string]string, _ string) (ConcreteEvent, error) {
	return &EmptyRequestType{
		MessageId: common.HexToHash(params["message_id"]),
		RequestId: common.HexToHash(params["request_id"]),
	}, nil
}

func decodeInvalidSubscriptionCrossChainPayment(params map[string]string, _ string) (ConcreteEvent, error) {
	ccipDestTokenAmountsLength, err := parsing.ScientificNotationToBigInt(params["ccip_dest_token_amounts_length"])
	if err != nil {
		return nil, fmt.Errorf("parse ccip_dest_token_amounts_length %q: %w", params["ccip_dest_token_amounts_length"], err)
	}
	return &InvalidSubscriptionCrossChainPayment{
		FundAdminAddr:              common.HexToAddress(params["fund_admin_addr"]),
		FundTokenId:                common.HexToHash(params["fund_token_id"]),
		RequestId:                  common.HexToHash(params["request_id"]),
		PaymentTokenDestAddr:       common.HexToAddress(params["payment_token_dest_addr"]),
		CcipDestTokenAmountsLength: ccipDestTokenAmountsLength,
		CcipPaymentTokenAddr:       common.HexToAddress(params["ccip_payment_token_addr"]),
	}, nil
}

func decodeSettlementFailed(params map[string]string, _ string) (ConcreteEvent, error) {
	shares, err := parsing.ScientificNotationToBigInt(params["shares"])
	if err != nil {
		return nil, fmt.Errorf("parse shares %q: %w", params["shares"], err)
	}
	amount, err := parsing.ScientificNotationToBigInt(params["amount"])
	if err != nil {
		return nil, fmt.Errorf("parse amount %q: %w", params["amount"], err)
	}
	return &SettlementFailed{
		FundTokenId:           common.HexToHash(params["fund_token_id"]),
		DistributorAddr:       common.HexToAddress(params["distributor_addr"]),
		PaymentTokenAddr:      common.HexToAddress(params["payment_token_addr"]),
		DistributorWalletAddr: common.HexToAddress(params["distributor_wallet_addr"]),
		RequestId:             common.HexToHash(params["request_id"]),
		Shares:                shares,
		Amount:                amount,
		ErrData:               []byte(params["err_data"]),
	}, nil
}

func decodeTokenWithdrawn(params map[string]string, _ string) (ConcreteEvent, error) {
	amount, err := parsing.ScientificNotationToBigInt(params["amount"])
	if err != nil {
		return nil, fmt.Errorf("parse amount %q: %w", params["amount"], err)
	}
	return &TokenWithdrawn{
		Token:     common.HexToAddress(params["token"]),
		Recipient: common.HexToAddress(params["recipient"]),
		Amount:    amount,
	}, nil
}

func decodeUnauthorizedSenderDTA(params map[string]string, _ string) (ConcreteEvent, error) {
	reqType, err := parsing.ScientificNotationToUint8(params["req_type"])
	if err != nil {
		return nil, fmt.Errorf("parse req_type %q: %w", params["req_type"], err)
	}
	dtaChainSelector, err := parsing.ScientificNotationToUint64(params["dta_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse dta_chain_selector %q: %w", params["dta_chain_selector"], err)
	}
	return &UnauthorizedSenderDTA{
		DtaAddr:          common.HexToAddress(params["dta_addr"]),
		DtaChainSelector: dtaChainSelector,
		FundTokenId:      common.HexToHash(params["fund_token_id"]),
		DistributorAddr:  common.HexToAddress(params["distributor_addr"]),
		RequestId:        common.HexToHash(params["request_id"]),
		ReqType:          DistributorRequestType(reqType),
	}, nil
}
