package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	apiClient "github.com/smartcontractkit/crec-api-go/client"
	"github.com/smartcontractkit/crec-sdk-ext-dta/parsing"
	"github.com/smartcontractkit/crec-sdk-ext-dta/types"
)

// Type aliases for shared types
type (
	ConcreteEvent = types.ConcreteEvent
	Event         = types.Event
	Metadata      = types.Metadata
	WorkflowEvent = types.WorkflowEvent
	Transaction   = types.Transaction
	Attribute     = types.Attribute
	Attrs         = types.Attrs
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

// unmarshalVerifiableEvent implements version-specific unmarshaling to populate ConcreteEvent.
func unmarshalVerifiableEvent(data []byte, v *types.VerifiableEvent) error {
	// Use an alias to avoid infinite recursion
	type alias types.VerifiableEvent
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return fmt.Errorf("failed to unmarshal VerifiableEvent envelope: %w", err)
	}

	// Copy envelope to receiver
	*v = types.VerifiableEvent(a)

	// Wrap to get EventName method
	wrapped := VerifiableEvent{VerifiableEvent: *v}
	name := wrapped.EventName()

	// Create the concrete event instance
	var concrete ConcreteEvent
	switch name {
	case EventDistributorRegistered:
		concrete = &DistributorRegistered{
			DistributorAddr: common.HexToAddress(v.Event.Parameters["distributor_addr"]),
		}
	case EventDistributorRequestCanceled:
		concrete = &DistributorRequestCanceled{
			FundTokenId:     common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DistributorAddr: common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			RequestId:       common.HexToHash(v.Event.Parameters["request_id"]),
		}
	case EventDistributorRequestProcessed:
		shares, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["shares"])
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, v.Event.Parameters["shares"])
		}
		status, err := parsing.ScientificNotationToUint8(v.Event.Parameters["status"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse status: %s", name, v.Event.Parameters["status"])
		}
		concrete = &DistributorRequestProcessed{
			RequestId: common.HexToHash(v.Event.Parameters["request_id"]),
			Shares:    shares,
			Status:    status,
			Error:     []byte(v.Event.Parameters["error"]),
		}
	case EventDistributorRequestProcessing:
		shares, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["shares"])
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, v.Event.Parameters["shares"])
		}
		amount, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["amount"])
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, v.Event.Parameters["amount"])
		}
		concrete = &DistributorRequestProcessing{
			FundTokenId:     common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DistributorAddr: common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			RequestId:       common.HexToHash(v.Event.Parameters["request_id"]),
			Shares:          shares,
			Amount:          amount,
		}
	case EventFundAdminRegistered:
		concrete = &FundAdminRegistered{
			FundAdminAddr: common.HexToAddress(v.Event.Parameters["fund_admin_addr"]),
		}
	case EventFundTokenAllowlistUpdated:
		allowed, err := strconv.ParseBool(v.Event.Parameters["allowed"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse allowed: %s", name, v.Event.Parameters["allowed"])
		}
		concrete = &FundTokenAllowlistUpdated{
			FundAdminAddr:   common.HexToAddress(v.Event.Parameters["fund_admin_addr"]),
			FundTokenId:     common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DistributorAddr: common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			Allowed:         allowed,
		}
	case EventFundTokenRegistered:
		tokenChainSelector, err := parsing.ScientificNotationToUint64(v.Event.Parameters["token_chain_selector"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse token_chain_selector: %s", name, v.Event.Parameters["token_chain_selector"])
		}
		concrete = &FundTokenRegistered{
			FundAdminAddr:      common.HexToAddress(v.Event.Parameters["fund_admin_addr"]),
			FundTokenId:        common.HexToHash(v.Event.Parameters["fund_token_id"]),
			FundTokenAddr:      common.HexToAddress(v.Event.Parameters["fund_token_addr"]),
			NavAddr:            common.HexToAddress(v.Event.Parameters["nav_addr"]),
			TokenChainSelector: tokenChainSelector,
		}
	case EventInitialized:
		version, err := parsing.ScientificNotationToUint64(v.Event.Parameters["version"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse version: %s", name, v.Event.Parameters["version"])
		}
		concrete = &Initialized{Version: version}
	case EventInvalidDTARequestSettlement:
		actualChainSelector, err := parsing.ScientificNotationToUint64(v.Event.Parameters["actual_chain_selector"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse actual_chain_selector: %s", name, v.Event.Parameters["actual_chain_selector"])
		}
		concrete = &InvalidDTARequestSettlement{
			FundAdminAddr:            common.HexToAddress(v.Event.Parameters["fund_admin_addr"]),
			FundTokenId:              common.HexToHash(v.Event.Parameters["fund_token_id"]),
			RequestId:                common.HexToHash(v.Event.Parameters["request_id"]),
			ActualChainSelector:      actualChainSelector,
			ActualDTAAdminWalletAddr: common.HexToAddress(v.Event.Parameters["actual_dta_admin_wallet_addr"]),
		}
	case EventMessageFailed:
		concrete = &MessageFailed{
			MessageId: common.HexToHash(v.Event.Parameters["message_id"]),
			Reason:    []byte(v.Event.Parameters["reason"]),
		}
	case EventNativeFundsRecovered:
		amount, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["amount"])
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, v.Event.Parameters["amount"])
		}
		concrete = &NativeFundsRecovered{
			To:     common.HexToAddress(v.Event.Parameters["to"]),
			Amount: amount,
		}
	case EventOwnershipTransferred:
		concrete = &OwnershipTransferred{
			PreviousOwner: common.HexToAddress(v.Event.Parameters["previous_owner"]),
			NewOwner:      common.HexToAddress(v.Event.Parameters["new_owner"]),
		}
	case EventRedemptionRequested:
		shares, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["shares"])
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, v.Event.Parameters["shares"])
		}
		createdAt, err := parsing.ScientificNotationToUint64(v.Event.Parameters["created_at"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse created_at: %s", name, v.Event.Parameters["created_at"])
		}
		concrete = &RedemptionRequested{
			FundTokenId:     common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DistributorAddr: common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			RequestId:       common.HexToHash(v.Event.Parameters["request_id"]),
			Shares:          shares,
			CreatedAt:       createdAt,
		}
	case EventSubscriptionRequested:
		amount, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["amount"])
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, v.Event.Parameters["amount"])
		}
		createdAt, err := parsing.ScientificNotationToUint64(v.Event.Parameters["created_at"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse created_at: %s", name, v.Event.Parameters["created_at"])
		}
		concrete = &SubscriptionRequested{
			FundTokenId:     common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DistributorAddr: common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			RequestId:       common.HexToHash(v.Event.Parameters["request_id"]),
			Amount:          amount,
			CreatedAt:       createdAt,
		}
	case EventAnswerUpdated:
		current, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["current"])
		if !ok {
			return fmt.Errorf("event %s unable to parse current: %s", name, v.Event.Parameters["current"])
		}
		roundId, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["roundId"])
		if !ok {
			return fmt.Errorf("event %s unable to parse roundId: %s", name, v.Event.Parameters["roundId"])
		}
		updatedAt, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["updatedAt"])
		if !ok {
			return fmt.Errorf("event %s unable to parse updatedAt: %s", name, v.Event.Parameters["updatedAt"])
		}
		concrete = &AnswerUpdated{Current: current, RoundId: roundId, UpdatedAt: updatedAt}
	case EventCCIPMessageRecvFailed:
		concrete = &CCIPMessageRecvFailed{
			MessageId: common.HexToHash(v.Event.Parameters["message_id"]),
			Reason:    []byte(v.Event.Parameters["reason"]),
		}
	case EventDTAAdded:
		dtaChainSelector, err := parsing.ScientificNotationToUint64(v.Event.Parameters["dta_chain_selector"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, v.Event.Parameters["dta_chain_selector"])
		}
		concrete = &DTAAdded{
			DtaAddr:          common.HexToAddress(v.Event.Parameters["dta_addr"]),
			DtaChainSelector: dtaChainSelector,
			FundTokenId:      common.HexToHash(v.Event.Parameters["fund_token_id"]),
			FundTokenAddr:    common.HexToAddress(v.Event.Parameters["fund_token_addr"]),
		}
	case EventDTARemoved:
		dtaChainSelector, err := parsing.ScientificNotationToUint64(v.Event.Parameters["dta_chain_selector"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, v.Event.Parameters["dta_chain_selector"])
		}
		concrete = &DTARemoved{
			DtaAddr:          common.HexToAddress(v.Event.Parameters["dta_addr"]),
			DtaChainSelector: dtaChainSelector,
			FundTokenId:      common.HexToHash(v.Event.Parameters["fund_token_id"]),
		}
	case EventDTASettlementClosed:
		requestType, err := parsing.ScientificNotationToUint8(v.Event.Parameters["request_type"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse request_type: %s", name, v.Event.Parameters["request_type"])
		}
		dtaChainSelector, err := parsing.ScientificNotationToUint64(v.Event.Parameters["dta_chain_selector"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, v.Event.Parameters["dta_chain_selector"])
		}
		success, err := strconv.ParseBool(v.Event.Parameters["success"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse success: %s", name, v.Event.Parameters["success"])
		}
		concrete = &DTASettlementClosed{
			DistributorAddr:  common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			RequestType:      uint8(requestType),
			FundTokenId:      common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DtaChainSelector: dtaChainSelector,
			DtaAddr:          common.HexToAddress(v.Event.Parameters["dta_addr"]),
			RequestId:        common.HexToHash(v.Event.Parameters["request_id"]),
			Success:          success,
			Err:              []byte(v.Event.Parameters["err"]),
		}
	case EventDTASettlementOpened:
		requestType, err := parsing.ScientificNotationToUint8(v.Event.Parameters["request_type"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse request_type: %s", name, v.Event.Parameters["request_type"])
		}
		dtaChainSelector, err := parsing.ScientificNotationToUint64(v.Event.Parameters["dta_chain_selector"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, v.Event.Parameters["dta_chain_selector"])
		}
		shares, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["shares"])
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, v.Event.Parameters["shares"])
		}
		amount, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["amount"])
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, v.Event.Parameters["amount"])
		}
		currency, err := parsing.ScientificNotationToUint8(v.Event.Parameters["currency"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse currency: %s", name, v.Event.Parameters["currency"])
		}
		concrete = &DTASettlementOpened{
			DistributorAddr:       common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			RequestType:           requestType,
			FundTokenId:           common.HexToHash(v.Event.Parameters["fund_token_id"]),
			FundAdminAddr:         common.HexToAddress(v.Event.Parameters["fund_admin_addr"]),
			DtaChainSelector:      dtaChainSelector,
			DtaAddr:               common.HexToAddress(v.Event.Parameters["dta_addr"]),
			RequestId:             common.HexToHash(v.Event.Parameters["request_id"]),
			DistributorWalletAddr: common.HexToAddress(v.Event.Parameters["distributor_wallet_addr"]),
			Shares:                shares,
			Amount:                amount,
			Currency:              currency,
		}
	case EventEmptyRequestType:
		concrete = &EmptyRequestType{
			MessageId: common.HexToHash(v.Event.Parameters["message_id"]),
			RequestId: common.HexToHash(v.Event.Parameters["request_id"]),
		}
	case EventInsufficientPaymentTokenBalance:
		amount, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["amount"])
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, v.Event.Parameters["amount"])
		}
		concrete = &InsufficientPaymentTokenBalance{
			FundTokenId:           common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DistributorAddr:       common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			DistributorWalletAddr: common.HexToAddress(v.Event.Parameters["distributor_wallet_addr"]),
			RequestId:             common.HexToHash(v.Event.Parameters["request_id"]),
			Amount:                amount,
		}
	case EventInvalidSubscriptionCrossChainPayment:
		ccipDestTokenAmountsLength, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["ccip_dest_token_amounts_length"])
		if !ok {
			return fmt.Errorf("event %s unable to parse ccip_dest_token_amounts_length: %s", name, v.Event.Parameters["ccip_dest_token_amounts_length"])
		}
		concrete = &InvalidSubscriptionCrossChainPayment{
			FundTokenId:                common.HexToHash(v.Event.Parameters["fund_token_id"]),
			RequestId:                  common.HexToHash(v.Event.Parameters["request_id"]),
			PaymentTokenDestAddr:       common.HexToAddress(v.Event.Parameters["payment_token_dest_addr"]),
			CCIPDestTokenAmountsLength: ccipDestTokenAmountsLength,
			CCIPPaymentTokenAddr:       common.HexToAddress(v.Event.Parameters["ccip_payment_token_addr"]),
		}
	case EventSettlementFailed:
		shares, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["shares"])
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, v.Event.Parameters["shares"])
		}
		amount, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["amount"])
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, v.Event.Parameters["amount"])
		}
		concrete = &SettlementFailed{
			FundTokenId:           common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DistributorAddr:       common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			PaymentTokenAddr:      common.HexToAddress(v.Event.Parameters["payment_token_addr"]),
			DistributorWalletAddr: common.HexToAddress(v.Event.Parameters["distributor_wallet_addr"]),
			RequestId:             common.HexToHash(v.Event.Parameters["request_id"]),
			Shares:                shares,
			Amount:                amount,
			ErrData:               []byte(v.Event.Parameters["err_data"]),
		}
	case EventTokenWithdrawn:
		amount, ok := parsing.ScientificNotationToBigInt(v.Event.Parameters["amount"])
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, v.Event.Parameters["amount"])
		}
		concrete = &TokenWithdrawn{
			Token:     common.HexToAddress(v.Event.Parameters["token"]),
			Recipient: common.HexToAddress(v.Event.Parameters["recipient"]),
			Amount:    amount,
		}
	case EventUnauthorizedSenderDTA:
		reqType, err := parsing.ScientificNotationToUint8(v.Event.Parameters["req_type"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse req_type: %s", name, v.Event.Parameters["req_type"])
		}
		dtaChainSelector, err := parsing.ScientificNotationToUint64(v.Event.Parameters["dta_chain_selector"])
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, v.Event.Parameters["dta_chain_selector"])
		}
		concrete = &UnauthorizedSenderDTA{
			DtaAddr:          common.HexToAddress(v.Event.Parameters["dta_addr"]),
			DtaChainSelector: dtaChainSelector,
			FundTokenId:      common.HexToHash(v.Event.Parameters["fund_token_id"]),
			DistributorAddr:  common.HexToAddress(v.Event.Parameters["distributor_addr"]),
			RequestId:        common.HexToHash(v.Event.Parameters["request_id"]),
			ReqType:          reqType,
		}
	default:
		return fmt.Errorf("unsupported event type: %s", name)
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
