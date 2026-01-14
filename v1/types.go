package v1

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
	TokenMintTypeMint        TokenMintType = iota // mint(address account, uint256 amount) - ERC3643, CMTAT
	TokenMintTypeIssueTokens                      // issueTokens(address _to, uint256 _value) - DSToken (BUIDL)
)

// TokenBurnType represents the burn type for DTA operations.
// Maps to IDTARequestSettlement.TokenBurnType enum in Solidity.
type TokenBurnType uint8

const (
	TokenBurnTypeBurn           TokenBurnType = iota // burn(address account, uint256 value) - ERC3643
	TokenBurnTypeBurnFrom                            // burnFrom(address account, uint256 value) - CMTAT
	TokenBurnTypeBurnWithReason                      // burn(address _who, uint256 _value, string _reason) - DSToken (BUIDL)
	TokenBurnTypeForceBurn                           // forceBurn(address account, uint256 amount, string reason) - CMTAT v2.3.0
)

// DistributorRequestType represents the type of distributor request.
// Maps to IDTAMessage.DistributorRequestType enum in Solidity.
type DistributorRequestType uint8

const (
	DistributorRequestTypeNone         DistributorRequestType = iota // Zero value
	DistributorRequestTypeSubscription                               // Subscription request
	DistributorRequestTypeRedemption                                 // Redemption request
)

// RequestStatus represents the status of a distributor request.
// Maps to IDTAMessage.RequestStatus enum in Solidity.
type RequestStatus uint8

const (
	RequestStatusNone       RequestStatus = iota // Zero value
	RequestStatusPending                         // Request is pending
	RequestStatusProcessing                      // Request is being processed
	RequestStatusProcessed                       // Request has been processed
	RequestStatusCanceled                        // Request was canceled
	RequestStatusFailed                          // Request failed
)

// =============================================================================
// Solidity Struct Types
// =============================================================================
// These types map to Solidity structs used in contract function parameters.

// DTAPayment represents payment information for DTA operations.
// Maps to payment-related struct fields in Solidity.
type DTAPayment struct {
	OffChainPaymentCurrency uint8
	PaymentTokenSourceAddr  common.Address
	PaymentTokenDestAddr    common.Address
}

// FundTokenData represents the data for registering a fund token.
// Maps to the FundTokenData struct in DTARequestManagement.
type FundTokenData struct {
	FundTokenAddr                 common.Address
	NavFeedDecimals               uint8
	PurchaseTokenRoundingDecimals uint8
	PurchaseTokenDecimals         uint8
	FundRoundingDecimals          uint8
	FundTokenDecimals             uint8
	RequestsPerDay                uint8
	NavAddr                       common.Address
	TokenChainSelector            uint64
	DtaRequestSettlementAddr      common.Address
	TimezoneOffsetSecs            *big.Int
	NavTTL                        *big.Int
	PaymentInfo                   DTAPayment
}

// DistributorRequest represents the data for a distributor request.
// Maps to the IDTADistributor.DistributorRequest struct in DTARequestManagement.
type DistributorRequest struct {
	Shares          *big.Int
	Amount          *big.Int
	FundTokenId     [32]byte
	FundAdminAddr   common.Address
	DistributorAddr common.Address
	CreatedAt       *big.Int
	RequestType     uint8
	Status          uint8
}
