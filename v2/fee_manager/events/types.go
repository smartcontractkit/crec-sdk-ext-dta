package events

// FeeTokenType represents the type of fee token.
// Maps to IFeeManager.FeeTokenType enum in Solidity.
type FeeTokenType uint8

const (
	FeeTokenTypeERC20  FeeTokenType = iota // Standard ERC-20 token
	FeeTokenTypeNative                     // Native chain token (ETH, etc.)
	FeeTokenTypeLink                       // LINK token (ERC-677, uses onTokenTransfer)
)
