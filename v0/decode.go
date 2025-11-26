package v0

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
// v0 uses Metadata.WorkflowEvent.Attributes to extract event data.
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

	// Helper to get attribute value
	getAttr := func(key string) string {
		if attr, ok := v.Metadata.WorkflowEvent.Attributes[key]; ok {
			return attr.Value
		}
		return ""
	}

	// Create the concrete event instance
	var concrete ConcreteEvent
	switch name {
	case EventDistributorRegistered:
		concrete = &DistributorRegistered{
			DistributorAddr: common.HexToAddress(getAttr("distributor_addr")),
		}
	case EventDistributorRequestCanceled:
		concrete = &DistributorRequestCanceled{
			FundTokenId:     common.HexToHash(getAttr("fund_token_id")),
			DistributorAddr: common.HexToAddress(getAttr("distributor_addr")),
			RequestId:       common.HexToHash(getAttr("request_id")),
		}
	case EventDistributorRequestProcessed:
		shares, ok := parsing.ScientificNotationToBigInt(getAttr("shares"))
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, getAttr("shares"))
		}
		status, err := strconv.ParseUint(getAttr("status"), 10, 8)
		if err != nil {
			return fmt.Errorf("event %s unable to parse status: %s", name, getAttr("status"))
		}
		concrete = &DistributorRequestProcessed{
			RequestId: common.HexToHash(getAttr("request_id")),
			Shares:    shares,
			Status:    uint8(status),
			Error:     []byte(getAttr("error")),
		}
	case EventDistributorRequestProcessing:
		shares, ok := parsing.ScientificNotationToBigInt(getAttr("shares"))
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, getAttr("shares"))
		}
		amount, ok := parsing.ScientificNotationToBigInt(getAttr("amount"))
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, getAttr("amount"))
		}
		concrete = &DistributorRequestProcessing{
			FundTokenId:     common.HexToHash(getAttr("fund_token_id")),
			DistributorAddr: common.HexToAddress(getAttr("distributor_addr")),
			RequestId:       common.HexToHash(getAttr("request_id")),
			Shares:          shares,
			Amount:          amount,
		}
	case EventFundAdminRegistered:
		concrete = &FundAdminRegistered{
			FundAdminAddr: common.HexToAddress(getAttr("fund_admin_addr")),
		}
	case EventFundTokenAllowlistUpdated:
		allowed, err := strconv.ParseBool(getAttr("allowed"))
		if err != nil {
			return fmt.Errorf("event %s unable to parse allowed: %s", name, getAttr("allowed"))
		}
		concrete = &FundTokenAllowlistUpdated{
			FundAdminAddr:   common.HexToAddress(getAttr("fund_admin_addr")),
			FundTokenId:     common.HexToHash(getAttr("fund_token_id")),
			DistributorAddr: common.HexToAddress(getAttr("distributor_addr")),
			Allowed:         allowed,
		}
	case EventFundTokenRegistered:
		tokenChainSelector, err := strconv.ParseUint(getAttr("token_chain_selector"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse token_chain_selector: %s", name, getAttr("token_chain_selector"))
		}
		concrete = &FundTokenRegistered{
			FundAdminAddr:      common.HexToAddress(getAttr("fund_admin_addr")),
			FundTokenId:        common.HexToHash(getAttr("fund_token_id")),
			FundTokenAddr:      common.HexToAddress(getAttr("fund_token_addr")),
			NavAddr:            common.HexToAddress(getAttr("nav_addr")),
			TokenChainSelector: tokenChainSelector,
		}
	case EventInitialized:
		version, err := strconv.ParseUint(getAttr("version"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse version: %s", name, getAttr("version"))
		}
		concrete = &Initialized{Version: version}
	case EventInvalidDTAWallet:
		actualChainSelector, err := strconv.ParseUint(getAttr("actual_chain_selector"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse actual_chain_selector: %s", name, getAttr("actual_chain_selector"))
		}
		concrete = &InvalidDTAWallet{
			FundAdminAddr:            common.HexToAddress(getAttr("fund_admin_addr")),
			FundTokenId:              common.HexToHash(getAttr("fund_token_id")),
			RequestId:                common.HexToHash(getAttr("request_id")),
			ActualChainSelector:      actualChainSelector,
			ActualDTAAdminWalletAddr: common.HexToAddress(getAttr("actual_dta_admin_wallet_addr")),
		}
	case EventMessageFailed:
		concrete = &MessageFailed{
			MessageId: common.HexToHash(getAttr("message_id")),
			Reason:    []byte(getAttr("reason")),
		}
	case EventNativeFundsRecovered:
		amount, ok := parsing.ScientificNotationToBigInt(getAttr("amount"))
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, getAttr("amount"))
		}
		concrete = &NativeFundsRecovered{
			To:     common.HexToAddress(getAttr("to")),
			Amount: amount,
		}
	case EventOwnershipTransferred:
		concrete = &OwnershipTransferred{
			PreviousOwner: common.HexToAddress(getAttr("previous_owner")),
			NewOwner:      common.HexToAddress(getAttr("new_owner")),
		}
	case EventRedemptionRequested:
		shares, ok := parsing.ScientificNotationToBigInt(getAttr("shares"))
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, getAttr("shares"))
		}
		createdAt, err := strconv.ParseUint(getAttr("created_at"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse created_at: %s", name, getAttr("created_at"))
		}
		concrete = &RedemptionRequested{
			FundTokenId:     common.HexToHash(getAttr("fund_token_id")),
			DistributorAddr: common.HexToAddress(getAttr("distributor_addr")),
			RequestId:       common.HexToHash(getAttr("request_id")),
			Shares:          shares,
			CreatedAt:       createdAt,
		}
	case EventSubscriptionRequested:
		amount, ok := parsing.ScientificNotationToBigInt(getAttr("amount"))
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, getAttr("amount"))
		}
		createdAt, err := strconv.ParseUint(getAttr("created_at"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse created_at: %s", name, getAttr("created_at"))
		}
		concrete = &SubscriptionRequested{
			FundTokenId:     common.HexToHash(getAttr("fund_token_id")),
			DistributorAddr: common.HexToAddress(getAttr("distributor_addr")),
			RequestId:       common.HexToHash(getAttr("request_id")),
			Amount:          amount,
			CreatedAt:       createdAt,
		}
	case EventAnswerUpdated:
		current, ok := parsing.ScientificNotationToBigInt(getAttr("current"))
		if !ok {
			return fmt.Errorf("event %s unable to parse current: %s", name, getAttr("current"))
		}
		roundId, ok := parsing.ScientificNotationToBigInt(getAttr("roundId"))
		if !ok {
			return fmt.Errorf("event %s unable to parse roundId: %s", name, getAttr("roundId"))
		}
		updatedAt, ok := parsing.ScientificNotationToBigInt(getAttr("updatedAt"))
		if !ok {
			return fmt.Errorf("event %s unable to parse updatedAt: %s", name, getAttr("updatedAt"))
		}
		concrete = &AnswerUpdated{Current: current, RoundId: roundId, UpdatedAt: updatedAt}
	case EventCCIPMessageRecvFailed:
		concrete = &CCIPMessageRecvFailed{
			MessageId: common.HexToHash(getAttr("message_id")),
			Reason:    []byte(getAttr("reason")),
		}
	case EventDTAAdded:
		dtaChainSelector, err := strconv.ParseUint(getAttr("dta_chain_selector"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, getAttr("dta_chain_selector"))
		}
		concrete = &DTAAdded{
			DtaAddr:          common.HexToAddress(getAttr("dta_addr")),
			DtaChainSelector: dtaChainSelector,
			FundTokenId:      common.HexToHash(getAttr("fund_token_id")),
			FundTokenAddr:    common.HexToAddress(getAttr("fund_token_addr")),
		}
	case EventDTARemoved:
		dtaChainSelector, err := strconv.ParseUint(getAttr("dta_chain_selector"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, getAttr("dta_chain_selector"))
		}
		concrete = &DTARemoved{
			DtaAddr:          common.HexToAddress(getAttr("dta_addr")),
			DtaChainSelector: dtaChainSelector,
			FundTokenId:      common.HexToHash(getAttr("fund_token_id")),
		}
	case EventDTASettlementClosed:
		requestType, err := strconv.ParseUint(getAttr("request_type"), 10, 8)
		if err != nil {
			return fmt.Errorf("event %s unable to parse request_type: %s", name, getAttr("request_type"))
		}
		dtaChainSelector, err := strconv.ParseUint(getAttr("dta_chain_selector"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, getAttr("dta_chain_selector"))
		}
		success, err := strconv.ParseBool(getAttr("success"))
		if err != nil {
			return fmt.Errorf("event %s unable to parse success: %s", name, getAttr("success"))
		}
		concrete = &DTASettlementClosed{
			DistributorAddr:  common.HexToAddress(getAttr("distributor_addr")),
			RequestType:      uint8(requestType),
			FundTokenId:      common.HexToHash(getAttr("fund_token_id")),
			DtaChainSelector: dtaChainSelector,
			DtaAddr:          common.HexToAddress(getAttr("dta_addr")),
			RequestId:        common.HexToHash(getAttr("request_id")),
			Success:          success,
			Err:              []byte(getAttr("err")),
		}
	case EventDTASettlementOpened:
		requestType, err := strconv.ParseUint(getAttr("request_type"), 10, 8)
		if err != nil {
			return fmt.Errorf("event %s unable to parse request_type: %s", name, getAttr("request_type"))
		}
		dtaChainSelector, err := strconv.ParseUint(getAttr("dta_chain_selector"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, getAttr("dta_chain_selector"))
		}
		shares, ok := parsing.ScientificNotationToBigInt(getAttr("shares"))
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, getAttr("shares"))
		}
		amount, ok := parsing.ScientificNotationToBigInt(getAttr("amount"))
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, getAttr("amount"))
		}
		currency, err := strconv.ParseUint(getAttr("currency"), 10, 8)
		if err != nil {
			return fmt.Errorf("event %s unable to parse currency: %s", name, getAttr("currency"))
		}
		concrete = &DTASettlementOpened{
			DistributorAddr:       common.HexToAddress(getAttr("distributor_addr")),
			RequestType:           uint8(requestType),
			FundTokenId:           common.HexToHash(getAttr("fund_token_id")),
			FundAdminAddr:         common.HexToAddress(getAttr("fund_admin_addr")),
			DtaChainSelector:      dtaChainSelector,
			DtaAddr:               common.HexToAddress(getAttr("dta_addr")),
			RequestId:             common.HexToHash(getAttr("request_id")),
			DistributorWalletAddr: common.HexToAddress(getAttr("distributor_wallet_addr")),
			Shares:                shares,
			Amount:                amount,
			Currency:              uint8(currency),
		}
	case EventEmptyRequestType:
		concrete = &EmptyRequestType{
			MessageId: common.HexToHash(getAttr("message_id")),
			RequestId: common.HexToHash(getAttr("request_id")),
		}
	case EventInsufficientPaymentTokenBalance:
		amount, ok := parsing.ScientificNotationToBigInt(getAttr("amount"))
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, getAttr("amount"))
		}
		concrete = &InsufficientPaymentTokenBalance{
			FundTokenId:           common.HexToHash(getAttr("fund_token_id")),
			DistributorAddr:       common.HexToAddress(getAttr("distributor_addr")),
			DistributorWalletAddr: common.HexToAddress(getAttr("distributor_wallet_addr")),
			RequestId:             common.HexToHash(getAttr("request_id")),
			Amount:                amount,
		}
	case EventSettlementFailed:
		shares, ok := parsing.ScientificNotationToBigInt(getAttr("shares"))
		if !ok {
			return fmt.Errorf("event %s unable to parse shares: %s", name, getAttr("shares"))
		}
		amount, ok := parsing.ScientificNotationToBigInt(getAttr("amount"))
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, getAttr("amount"))
		}
		concrete = &SettlementFailed{
			FundTokenId:           common.HexToHash(getAttr("fund_token_id")),
			DistributorAddr:       common.HexToAddress(getAttr("distributor_addr")),
			PaymentTokenAddr:      common.HexToAddress(getAttr("payment_token_addr")),
			DistributorWalletAddr: common.HexToAddress(getAttr("distributor_wallet_addr")),
			RequestId:             common.HexToHash(getAttr("request_id")),
			Shares:                shares,
			Amount:                amount,
			ErrData:               []byte(getAttr("err_data")),
		}
	case EventTokenWithdrawn:
		amount, ok := parsing.ScientificNotationToBigInt(getAttr("amount"))
		if !ok {
			return fmt.Errorf("event %s unable to parse amount: %s", name, getAttr("amount"))
		}
		concrete = &TokenWithdrawn{
			Token:     common.HexToAddress(getAttr("token")),
			Recipient: common.HexToAddress(getAttr("recipient")),
			Amount:    amount,
		}
	case EventUnauthorizedSenderDTA:
		reqType, err := strconv.ParseUint(getAttr("req_type"), 10, 8)
		if err != nil {
			return fmt.Errorf("event %s unable to parse req_type: %s", name, getAttr("req_type"))
		}
		dtaChainSelector, err := strconv.ParseUint(getAttr("dta_chain_selector"), 10, 64)
		if err != nil {
			return fmt.Errorf("event %s unable to parse dta_chain_selector: %s", name, getAttr("dta_chain_selector"))
		}
		concrete = &UnauthorizedSenderDTA{
			DtaAddr:          common.HexToAddress(getAttr("dta_addr")),
			DtaChainSelector: dtaChainSelector,
			FundTokenId:      common.HexToHash(getAttr("fund_token_id")),
			DistributorAddr:  common.HexToAddress(getAttr("distributor_addr")),
			RequestId:        common.HexToHash(getAttr("request_id")),
			ReqType:          uint8(reqType),
		}
	default:
		return fmt.Errorf("unsupported event type: %s", getAttr("event_type"))
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
