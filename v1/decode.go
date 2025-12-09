package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	apiClient "github.com/smartcontractkit/crec-api-go/client"
	"github.com/smartcontractkit/crec-sdk-ext-dta/parsing"
	"github.com/smartcontractkit/crec-sdk-ext-dta/types"
)

// VerifiableEvent wraps the shared VerifiableEvent type to add version-specific methods.
type VerifiableEvent struct {
	types.VerifiableEvent
}

// EventName determines and returns the event name from the workflow attributes or outer event name;
// defaults to EventUnknown if not resolvable.
func (v VerifiableEvent) EventName() EventName {
	return types.GetEventName(v.VerifiableEvent, EventUnknown, parseEvent)
}

// eventDecoder is a function that decodes event parameters into a concrete event type.
type eventDecoder func(params map[string]string, txHash string) (types.ConcreteEvent, error)

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
	EventInsufficientPaymentTokenBalance:      decodeInsufficientPaymentTokenBalance,
	EventInvalidSubscriptionCrossChainPayment: decodeInvalidSubscriptionCrossChainPayment,
	EventSettlementFailed:                     decodeSettlementFailed,
	EventTokenWithdrawn:                       decodeTokenWithdrawn,
	EventUnauthorizedSenderDTA:                decodeUnauthorizedSenderDTA,
}

// unmarshalVerifiableEvent implements version-specific unmarshaling to populate ConcreteEvent.
func unmarshalVerifiableEvent(data []byte, v *types.VerifiableEvent) error {
	// Use an alias to avoid infinite recursion
	type alias types.VerifiableEvent
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return fmt.Errorf("unmarshal envelope: %w", err)
	}

	// Copy envelope to receiver
	*v = types.VerifiableEvent(a)

	// Wrap to get EventName method
	wrapped := VerifiableEvent{VerifiableEvent: *v}
	name := wrapped.EventName()
	txHash := v.Transaction.Hash

	// Look up decoder in registry
	decoder, ok := eventDecoders[name]
	if !ok {
		return fmt.Errorf("unsupported event type %q (tx=%s)", name, txHash)
	}

	// Decode the event
	concrete, err := decoder(v.Event.Parameters, txHash)
	if err != nil {
		return fmt.Errorf("decode %s (tx=%s): %w", name, txHash, err)
	}

	v.ConcreteEvent = concrete
	return nil
}

// DecodeFromEvent extracts the WatcherEventPayload from an apiClient.Event and converts it
// to a VerifiableEvent with the ConcreteEvent populated based on the event type.
func DecodeFromEvent(ctx context.Context, event apiClient.Event) (VerifiableEvent, error) {
	ve, err := types.DecodeFromEvent(ctx, event, unmarshalVerifiableEvent)
	if err != nil {
		return VerifiableEvent{}, err
	}
	return VerifiableEvent{VerifiableEvent: ve}, nil
}

// ============================================================================
// Event Decoders
// ============================================================================

func decodeDistributorRegistered(params map[string]string, _ string) (types.ConcreteEvent, error) {
	return &DistributorRegistered{
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
	}, nil
}

func decodeDistributorRequestCanceled(params map[string]string, _ string) (types.ConcreteEvent, error) {
	return &DistributorRequestCanceled{
		FundTokenId:     common.HexToHash(params["fund_token_id"]),
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
		RequestId:       common.HexToHash(params["request_id"]),
	}, nil
}

func decodeDistributorRequestProcessed(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	shares, ok := parsing.ScientificNotationToBigInt(params["shares"])
	if !ok {
		return nil, fmt.Errorf("parse shares %q: invalid format", params["shares"])
	}
	status, err := parsing.ScientificNotationToUint8(params["status"])
	if err != nil {
		return nil, fmt.Errorf("parse status %q: %w", params["status"], err)
	}
	return &DistributorRequestProcessed{
		RequestId: common.HexToHash(params["request_id"]),
		Shares:    shares,
		Status:    status,
		Error:     []byte(params["error"]),
	}, nil
}

func decodeDistributorRequestProcessing(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	shares, ok := parsing.ScientificNotationToBigInt(params["shares"])
	if !ok {
		return nil, fmt.Errorf("parse shares %q: invalid format", params["shares"])
	}
	amount, ok := parsing.ScientificNotationToBigInt(params["amount"])
	if !ok {
		return nil, fmt.Errorf("parse amount %q: invalid format", params["amount"])
	}
	return &DistributorRequestProcessing{
		FundTokenId:     common.HexToHash(params["fund_token_id"]),
		DistributorAddr: common.HexToAddress(params["distributor_addr"]),
		RequestId:       common.HexToHash(params["request_id"]),
		Shares:          shares,
		Amount:          amount,
	}, nil
}

func decodeFundAdminRegistered(params map[string]string, _ string) (types.ConcreteEvent, error) {
	return &FundAdminRegistered{
		FundAdminAddr: common.HexToAddress(params["fund_admin_addr"]),
	}, nil
}

func decodeFundTokenAllowlistUpdated(params map[string]string, txHash string) (types.ConcreteEvent, error) {
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

func decodeFundTokenRegistered(params map[string]string, txHash string) (types.ConcreteEvent, error) {
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

func decodeInitialized(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	version, err := parsing.ScientificNotationToUint64(params["version"])
	if err != nil {
		return nil, fmt.Errorf("parse version %q: %w", params["version"], err)
	}
	return &Initialized{Version: version}, nil
}

func decodeInvalidDTARequestSettlement(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	actualChainSelector, err := parsing.ScientificNotationToUint64(params["actual_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse actual_chain_selector %q: %w", params["actual_chain_selector"], err)
	}
	return &InvalidDTARequestSettlement{
		FundAdminAddr:            common.HexToAddress(params["fund_admin_addr"]),
		FundTokenId:              common.HexToHash(params["fund_token_id"]),
		RequestId:                common.HexToHash(params["request_id"]),
		ActualChainSelector:      actualChainSelector,
		ActualDTAAdminWalletAddr: common.HexToAddress(params["actual_dta_admin_wallet_addr"]),
	}, nil
}

func decodeMessageFailed(params map[string]string, _ string) (types.ConcreteEvent, error) {
	return &MessageFailed{
		MessageId: common.HexToHash(params["message_id"]),
		Reason:    []byte(params["reason"]),
	}, nil
}

func decodeNativeFundsRecovered(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	amount, ok := parsing.ScientificNotationToBigInt(params["amount"])
	if !ok {
		return nil, fmt.Errorf("parse amount %q: invalid format", params["amount"])
	}
	return &NativeFundsRecovered{
		To:     common.HexToAddress(params["to"]),
		Amount: amount,
	}, nil
}

func decodeOwnershipTransferred(params map[string]string, _ string) (types.ConcreteEvent, error) {
	return &OwnershipTransferred{
		PreviousOwner: common.HexToAddress(params["previous_owner"]),
		NewOwner:      common.HexToAddress(params["new_owner"]),
	}, nil
}

func decodeRedemptionRequested(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	shares, ok := parsing.ScientificNotationToBigInt(params["shares"])
	if !ok {
		return nil, fmt.Errorf("parse shares %q: invalid format", params["shares"])
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

func decodeSubscriptionRequested(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	amount, ok := parsing.ScientificNotationToBigInt(params["amount"])
	if !ok {
		return nil, fmt.Errorf("parse amount %q: invalid format", params["amount"])
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

func decodeAnswerUpdated(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	current, ok := parsing.ScientificNotationToBigInt(params["current"])
	if !ok {
		return nil, fmt.Errorf("parse current %q: invalid format", params["current"])
	}
	roundId, ok := parsing.ScientificNotationToBigInt(params["roundId"])
	if !ok {
		return nil, fmt.Errorf("parse roundId %q: invalid format", params["roundId"])
	}
	updatedAt, ok := parsing.ScientificNotationToBigInt(params["updatedAt"])
	if !ok {
		return nil, fmt.Errorf("parse updatedAt %q: invalid format", params["updatedAt"])
	}
	return &AnswerUpdated{Current: current, RoundId: roundId, UpdatedAt: updatedAt}, nil
}

func decodeCCIPMessageRecvFailed(params map[string]string, _ string) (types.ConcreteEvent, error) {
	return &CCIPMessageRecvFailed{
		MessageId: common.HexToHash(params["message_id"]),
		Reason:    []byte(params["reason"]),
	}, nil
}

func decodeDTAAdded(params map[string]string, txHash string) (types.ConcreteEvent, error) {
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

func decodeDTARemoved(params map[string]string, txHash string) (types.ConcreteEvent, error) {
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

func decodeDTASettlementClosed(params map[string]string, txHash string) (types.ConcreteEvent, error) {
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
		RequestType:      requestType,
		FundTokenId:      common.HexToHash(params["fund_token_id"]),
		DtaChainSelector: dtaChainSelector,
		DtaAddr:          common.HexToAddress(params["dta_addr"]),
		RequestId:        common.HexToHash(params["request_id"]),
		Success:          success,
		Err:              []byte(params["err"]),
	}, nil
}

func decodeDTASettlementOpened(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	requestType, err := parsing.ScientificNotationToUint8(params["request_type"])
	if err != nil {
		return nil, fmt.Errorf("parse request_type %q: %w", params["request_type"], err)
	}
	dtaChainSelector, err := parsing.ScientificNotationToUint64(params["dta_chain_selector"])
	if err != nil {
		return nil, fmt.Errorf("parse dta_chain_selector %q: %w", params["dta_chain_selector"], err)
	}
	shares, ok := parsing.ScientificNotationToBigInt(params["shares"])
	if !ok {
		return nil, fmt.Errorf("parse shares %q: invalid format", params["shares"])
	}
	amount, ok := parsing.ScientificNotationToBigInt(params["amount"])
	if !ok {
		return nil, fmt.Errorf("parse amount %q: invalid format", params["amount"])
	}
	currency, err := parsing.ScientificNotationToUint8(params["currency"])
	if err != nil {
		return nil, fmt.Errorf("parse currency %q: %w", params["currency"], err)
	}
	return &DTASettlementOpened{
		DistributorAddr:       common.HexToAddress(params["distributor_addr"]),
		RequestType:           requestType,
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

func decodeEmptyRequestType(params map[string]string, _ string) (types.ConcreteEvent, error) {
	return &EmptyRequestType{
		MessageId: common.HexToHash(params["message_id"]),
		RequestId: common.HexToHash(params["request_id"]),
	}, nil
}

func decodeInsufficientPaymentTokenBalance(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	amount, ok := parsing.ScientificNotationToBigInt(params["amount"])
	if !ok {
		return nil, fmt.Errorf("parse amount %q: invalid format", params["amount"])
	}
	return &InsufficientPaymentTokenBalance{
		FundTokenId:           common.HexToHash(params["fund_token_id"]),
		DistributorAddr:       common.HexToAddress(params["distributor_addr"]),
		DistributorWalletAddr: common.HexToAddress(params["distributor_wallet_addr"]),
		RequestId:             common.HexToHash(params["request_id"]),
		Amount:                amount,
	}, nil
}

func decodeInvalidSubscriptionCrossChainPayment(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	ccipDestTokenAmountsLength, ok := parsing.ScientificNotationToBigInt(params["ccip_dest_token_amounts_length"])
	if !ok {
		return nil, fmt.Errorf("parse ccip_dest_token_amounts_length %q: invalid format", params["ccip_dest_token_amounts_length"])
	}
	return &InvalidSubscriptionCrossChainPayment{
		FundTokenId:                common.HexToHash(params["fund_token_id"]),
		RequestId:                  common.HexToHash(params["request_id"]),
		PaymentTokenDestAddr:       common.HexToAddress(params["payment_token_dest_addr"]),
		CCIPDestTokenAmountsLength: ccipDestTokenAmountsLength,
		CCIPPaymentTokenAddr:       common.HexToAddress(params["ccip_payment_token_addr"]),
	}, nil
}

func decodeSettlementFailed(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	shares, ok := parsing.ScientificNotationToBigInt(params["shares"])
	if !ok {
		return nil, fmt.Errorf("parse shares %q: invalid format", params["shares"])
	}
	amount, ok := parsing.ScientificNotationToBigInt(params["amount"])
	if !ok {
		return nil, fmt.Errorf("parse amount %q: invalid format", params["amount"])
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

func decodeTokenWithdrawn(params map[string]string, txHash string) (types.ConcreteEvent, error) {
	amount, ok := parsing.ScientificNotationToBigInt(params["amount"])
	if !ok {
		return nil, fmt.Errorf("parse amount %q: invalid format", params["amount"])
	}
	return &TokenWithdrawn{
		Token:     common.HexToAddress(params["token"]),
		Recipient: common.HexToAddress(params["recipient"]),
		Amount:    amount,
	}, nil
}

func decodeUnauthorizedSenderDTA(params map[string]string, txHash string) (types.ConcreteEvent, error) {
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
		ReqType:          reqType,
	}, nil
}
