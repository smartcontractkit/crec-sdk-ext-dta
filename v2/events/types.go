package events

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// =============================================================================
// Solidity Enum Types
// =============================================================================
// These types map to Solidity enums defined in the DTA contracts.
// Since Solidity ABIs encode enums as uint8 without preserving names,
// these must be manually maintained to match the contract definitions.

// TokenMintType represents the mint type for DTA operations.
// Maps to IDTARequestSettlement.TokenMintType enum in Solidity.
type TokenMintType uint8

const (
	// TokenMintTypeMint uses mint(address account, uint256 amount) - ERC3643, CMTAT.
	TokenMintTypeMint TokenMintType = iota
	// TokenMintTypeIssueTokens uses issueTokens(address _to, uint256 _value) - DSToken (BUIDL).
	TokenMintTypeIssueTokens
)

// TokenBurnType represents the burn type for DTA operations.
// Maps to IDTARequestSettlement.TokenBurnType enum in Solidity.
type TokenBurnType uint8

const (
	// TokenBurnTypeBurn uses burn(address account, uint256 value) - ERC3643.
	TokenBurnTypeBurn TokenBurnType = iota
	// TokenBurnTypeBurnFrom uses burnFrom(address account, uint256 value) - CMTAT.
	TokenBurnTypeBurnFrom
	// TokenBurnTypeBurnWithReason uses burn(address _who, uint256 _value, string _reason) - DSToken (BUIDL).
	TokenBurnTypeBurnWithReason
	// TokenBurnTypeForceBurn uses forceBurn(address account, uint256 amount, string reason) - CMTAT v2.3.0.
	TokenBurnTypeForceBurn
)

// DistributorRequestType represents the type of distributor request.
// Maps to IDTAMessage.DistributorRequestType enum in Solidity.
type DistributorRequestType uint8

const (
	// DistributorRequestTypeNone is the zero value for request type.
	DistributorRequestTypeNone DistributorRequestType = iota
	// DistributorRequestTypeSubscription indicates a subscription request.
	DistributorRequestTypeSubscription
	// DistributorRequestTypeRedemption indicates a redemption request.
	DistributorRequestTypeRedemption
)

// RequestStatus represents the status of a distributor request.
// Maps to IDTAMessage.RequestStatus enum in Solidity.
type RequestStatus uint8

const (
	// RequestStatusNone is the zero value for request status.
	RequestStatusNone RequestStatus = iota
	// RequestStatusPending indicates the request is pending.
	RequestStatusPending
	// RequestStatusProcessing indicates the request is being processed.
	RequestStatusProcessing
	// RequestStatusProcessed indicates the request has been processed.
	RequestStatusProcessed
	// RequestStatusCanceled indicates the request was canceled.
	RequestStatusCanceled
	// RequestStatusFailed indicates the request failed.
	RequestStatusFailed
)

// =============================================================================
// Solidity Struct Types
// =============================================================================
// These types map to Solidity structs used in contract function parameters.

// DTAPayment represents payment information for DTA operations.
// Maps to payment-related struct fields in Solidity.
type DTAPayment struct {
	// OffChainPaymentCurrency indicates whether payment is off-chain.
	OffChainPaymentCurrency uint8
	// PaymentTokenSourceAddr is the source address for the payment token.
	PaymentTokenSourceAddr common.Address
	// PaymentTokenDestAddr is the destination address for the payment token.
	PaymentTokenDestAddr common.Address
}

// FundTokenData represents the data for registering a fund token.
// Maps to the FundTokenData struct in DTARequestManagement.
type FundTokenData struct {
	// FundTokenAddr is the address of the fund token contract.
	FundTokenAddr common.Address
	// NavFeedDecimals is the decimal precision of the NAV feed.
	NavFeedDecimals uint8
	// PurchaseTokenRoundingDecimals controls rounding for purchase token amounts.
	PurchaseTokenRoundingDecimals uint8
	// PurchaseTokenDecimals is the decimal precision of the purchase token.
	PurchaseTokenDecimals uint8
	// FundRoundingDecimals controls rounding for fund share amounts.
	FundRoundingDecimals uint8
	// FundTokenDecimals is the decimal precision of the fund token.
	FundTokenDecimals uint8
	// RequestsPerDay limits the number of requests per day per investor.
	RequestsPerDay uint8
	// NavAddr is the address of the NAV feed contract.
	NavAddr common.Address
	// TokenChainSelector is the CCIP chain selector for the fund token.
	TokenChainSelector uint64
	// DtaRequestSettlementAddr is the address of the DTA request settlement contract.
	DtaRequestSettlementAddr common.Address
	// TimezoneOffsetSecs is the timezone offset in seconds for NAV updates.
	TimezoneOffsetSecs *big.Int
	// NavTTL is the time-to-live for NAV values.
	NavTTL *big.Int
	// PaymentInfo holds payment token and currency configuration.
	PaymentInfo DTAPayment
}

// DistributorRequest represents the data for a distributor request.
// Maps to the IDTADistributor.DistributorRequest struct in DTARequestManagement.
type DistributorRequest struct {
	// Shares is the number of fund shares for the request.
	Shares *big.Int
	// Amount is the payment amount for the request.
	Amount *big.Int
	// FundTokenId identifies the fund token.
	FundTokenId [32]byte
	// ReferenceID identifies the external reference associated with the request.
	ReferenceID [32]byte
	// FundAdminAddr is the address of the fund administrator.
	FundAdminAddr common.Address
	// DistributorAddr is the address of the distributor.
	DistributorAddr common.Address
	// CreatedAt is the timestamp when the request was created.
	CreatedAt *big.Int
	// RequestType indicates subscription or redemption (see DistributorRequestType).
	RequestType uint8
	// Status indicates the request status (see RequestStatus).
	Status uint8
}
