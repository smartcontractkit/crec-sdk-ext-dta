// Package v0 provides CREC SDK extension for DTA (Digital Token Asset) v0.1 operations.
//
// This version works with DTAOpenMarketplace and DTAWallet contracts.
//
// # Usage
//
//	ext, err := v0.New(&v0.Options{
//		DTAOpenMarketplaceAddress: "0x...",
//		DTAWalletAddress:          "0x...",
//		AccountAddress:            "0x...",
//	})
package v0

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"

	apiClient "github.com/smartcontractkit/crec-api-go/client"
	"github.com/smartcontractkit/crec-api-go/services/dta/gen/dtaopenmarketplace"
	"github.com/smartcontractkit/crec-api-go/services/dta/gen/dtawallet"

	"github.com/smartcontractkit/crec-sdk/interfaces/erc20"
	transactTypes "github.com/smartcontractkit/crec-sdk/transact/types"
)

// TokenBurnType represents the burn type for DTA operations.
type TokenBurnType uint8

const (
	TokenBurnTypeNone TokenBurnType = iota
	TokenBurnTypeBurn
	TokenBurnTypeTransfer
)

// Options defines the configuration for creating a new CREC DTA v0 extension.
type Options struct {
	// Logger is an optional logger instance. If nil, a default nop logger is used.
	Logger *zerolog.Logger

	// DTAOpenMarketplaceAddress is the address of the DTAOpenMarketplace contract.
	DTAOpenMarketplaceAddress string

	// DTAWalletAddress is the address of the DTAWallet contract.
	DTAWalletAddress string

	// AccountAddress is the address of the account performing the DTA operations.
	AccountAddress string
}

// Extension provides methods for preparing DTA v0 operations.
type Extension struct {
	logger                    *zerolog.Logger
	dtaOpenMarketplaceAddress common.Address
	dtaWalletAddress          common.Address
	accountAddress            common.Address
}

// New creates a new CREC DTA v0 extension with the provided options.
// Returns a pointer to the Extension and an error if any issues occur during initialization.
func New(opts *Options) (*Extension, error) {
	if opts == nil {
		return nil, fmt.Errorf("options is required")
	}

	var logger *zerolog.Logger
	if opts.Logger != nil {
		logger = opts.Logger
	} else {
		nopLogger := zerolog.Nop()
		logger = &nopLogger
	}

	logger.Info().Msg("Creating CREC DTA v0 extension")

	return &Extension{
		logger:                    logger,
		dtaOpenMarketplaceAddress: common.HexToAddress(opts.DTAOpenMarketplaceAddress),
		dtaWalletAddress:          common.HexToAddress(opts.DTAWalletAddress),
		accountAddress:            common.HexToAddress(opts.AccountAddress),
	}, nil
}

// FundTokenData represents the data structure for fund token registration.
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
	DtaWalletAddr                 common.Address
	TimezoneOffsetSecs            *big.Int
	NavTTL                        *big.Int
	PaymentInfo                   DTAPaymentInfo
}

// DTAPaymentInfo represents the payment information for DTA operations.
type DTAPaymentInfo struct {
	OffChainPaymentCurrency uint8
	PaymentTokenSourceAddr  common.Address
	PaymentTokenDestAddr    common.Address
}

// ============================================================================
// DTAOpenMarketplace Operations
// ============================================================================

// PrepareRequestSubscriptionOperation prepares a DTA request subscription operation.
func (e *Extension) PrepareRequestSubscriptionOperation(
	fundAdminAddr common.Address,
	fundTokenId [32]byte,
	amount *big.Int,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("requestSubscription", fundAdminAddr, fundTokenId, amount)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for requestSubscription")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRequestRedemptionOperation prepares a DTA request redemption operation.
func (e *Extension) PrepareRequestRedemptionOperation(
	fundAdminAddr common.Address,
	fundTokenId [32]byte,
	shares *big.Int,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("requestRedemption", fundAdminAddr, fundTokenId, shares)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for requestRedemption")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRequestSubscriptionWithTokenApprovalOperation prepares a subscription operation with token approval.
func (e *Extension) PrepareRequestSubscriptionWithTokenApprovalOperation(
	fundAdminAddr common.Address,
	fundTokenId [32]byte,
	amount *big.Int,
	paymentTokenAddress common.Address,
) (*transactTypes.Operation, error) {
	approveTransaction, err := e.prepareTokenApproveTransaction(&paymentTokenAddress, amount)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to prepare token approve transaction")
		return nil, err
	}

	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("requestSubscription", fundAdminAddr, fundTokenId, amount)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for requestSubscription")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			*approveTransaction,
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareProcessDistributorRequestOperation prepares a process distributor request operation.
func (e *Extension) PrepareProcessDistributorRequestOperation(requestId [32]byte) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("processDistributorRequest", requestId)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for processDistributorRequest")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareCancelDistributorRequestOperation prepares a cancel distributor request operation.
func (e *Extension) PrepareCancelDistributorRequestOperation(requestId [32]byte) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("cancelDistributorRequest", requestId)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for cancelDistributorRequest")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRegisterDistributorOperation prepares a register distributor operation.
func (e *Extension) PrepareRegisterDistributorOperation(
	distributorAddr common.Address,
	distributorWalletAddr common.Address,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("registerDistributor", distributorAddr, distributorWalletAddr)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for registerDistributor")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareUpdateDistributorOperation prepares an update distributor operation.
func (e *Extension) PrepareUpdateDistributorOperation(
	distributorWalletAddr common.Address,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("updateDistributor", distributorWalletAddr)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for updateDistributor")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRegisterFundAdminOperation prepares a register fund admin operation.
func (e *Extension) PrepareRegisterFundAdminOperation(fundAdminAddr common.Address) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("registerFundAdmin", fundAdminAddr)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for registerFundAdmin")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRegisterFundTokenOperation prepares a register fund token operation.
func (e *Extension) PrepareRegisterFundTokenOperation(
	fundTokenId [32]byte,
	tokenData FundTokenData,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	contractTokenData := struct {
		FundTokenAddr                 common.Address
		NavFeedDecimals               uint8
		PurchaseTokenRoundingDecimals uint8
		PurchaseTokenDecimals         uint8
		FundRoundingDecimals          uint8
		FundTokenDecimals             uint8
		RequestsPerDay                uint8
		NavAddr                       common.Address
		TokenChainSelector            uint64
		DtaWalletAddr                 common.Address
		TimezoneOffsetSecs            *big.Int
		NavTTL                        *big.Int
		PaymentInfo                   struct {
			OffChainPaymentCurrency uint8
			PaymentTokenSourceAddr  common.Address
			PaymentTokenDestAddr    common.Address
		}
	}{
		FundTokenAddr:                 tokenData.FundTokenAddr,
		NavFeedDecimals:               tokenData.NavFeedDecimals,
		PurchaseTokenRoundingDecimals: tokenData.PurchaseTokenRoundingDecimals,
		PurchaseTokenDecimals:         tokenData.PurchaseTokenDecimals,
		FundRoundingDecimals:          tokenData.FundRoundingDecimals,
		FundTokenDecimals:             tokenData.FundTokenDecimals,
		RequestsPerDay:                tokenData.RequestsPerDay,
		NavAddr:                       tokenData.NavAddr,
		TokenChainSelector:            tokenData.TokenChainSelector,
		DtaWalletAddr:                 tokenData.DtaWalletAddr,
		TimezoneOffsetSecs:            tokenData.TimezoneOffsetSecs,
		NavTTL:                        tokenData.NavTTL,
		PaymentInfo: struct {
			OffChainPaymentCurrency uint8
			PaymentTokenSourceAddr  common.Address
			PaymentTokenDestAddr    common.Address
		}{
			OffChainPaymentCurrency: tokenData.PaymentInfo.OffChainPaymentCurrency,
			PaymentTokenSourceAddr:  tokenData.PaymentInfo.PaymentTokenSourceAddr,
			PaymentTokenDestAddr:    tokenData.PaymentInfo.PaymentTokenDestAddr,
		},
	}

	calldata, err := abiEncoder.Pack("registerFundToken", fundTokenId, contractTokenData)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for registerFundToken")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareAllowDistributorForTokenOperation prepares an allow distributor for token operation.
func (e *Extension) PrepareAllowDistributorForTokenOperation(
	fundTokenId [32]byte,
	distributorAddr common.Address,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("allowDistributorForToken", fundTokenId, distributorAddr)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for allowDistributorForToken")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareDisallowDistributorForTokenOperation prepares a disallow distributor for token operation.
func (e *Extension) PrepareDisallowDistributorForTokenOperation(
	fundTokenId [32]byte,
	distributorAddr common.Address,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("disallowDistributorForToken", fundTokenId, distributorAddr)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for disallowDistributorForToken")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareEnableFundTokenOperation prepares an enable fund token operation.
func (e *Extension) PrepareEnableFundTokenOperation(fundTokenId [32]byte) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("enableFundToken", fundTokenId)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for enableFundToken")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareDisableFundTokenOperation prepares a disable fund token operation.
func (e *Extension) PrepareDisableFundTokenOperation(fundTokenId [32]byte) (*transactTypes.Operation, error) {
	abiEncoder, err := dtaopenmarketplace.DtaopenmarketplaceMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAOpenMarketplace ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("disableFundToken", fundTokenId)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for disableFundToken")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaOpenMarketplaceAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// ============================================================================
// DTAWallet Operations
// ============================================================================

// PrepareAllowDTAOperation prepares a DTA wallet allow DTA operation.
func (e *Extension) PrepareAllowDTAOperation(
	dtaAddr common.Address,
	dtaChainSelector uint64,
	fundTokenId [32]byte,
	fundTokenAddr common.Address,
	burnType TokenBurnType,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtawallet.DtawalletMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAWallet ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("allowDTA", dtaAddr, dtaChainSelector, fundTokenId, fundTokenAddr, uint8(burnType))
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for allowDTA")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaWalletAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareDisallowDTAOperation prepares a DTA wallet disallow DTA operation.
func (e *Extension) PrepareDisallowDTAOperation(
	dtaAddr common.Address,
	dtaChainSelector uint64,
	fundTokenId [32]byte,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtawallet.DtawalletMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAWallet ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("disallowDTA", dtaAddr, dtaChainSelector, fundTokenId)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for disallowDTA")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaWalletAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareWithdrawTokensOperation prepares a DTA wallet withdraw tokens operation.
func (e *Extension) PrepareWithdrawTokensOperation(
	token common.Address,
	recipient common.Address,
	amount *big.Int,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtawallet.DtawalletMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAWallet ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("withdrawTokens", token, recipient, amount)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for withdrawTokens")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaWalletAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareTransferWalletOwnershipOperation prepares a DTA wallet transfer ownership operation.
func (e *Extension) PrepareTransferWalletOwnershipOperation(newOwner common.Address) (*transactTypes.Operation, error) {
	abiEncoder, err := dtawallet.DtawalletMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAWallet ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("transferOwnership", newOwner)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for transferOwnership")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaWalletAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRenounceWalletOwnershipOperation prepares a DTA wallet renounce ownership operation.
func (e *Extension) PrepareRenounceWalletOwnershipOperation() (*transactTypes.Operation, error) {
	abiEncoder, err := dtawallet.DtawalletMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAWallet ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("renounceOwnership")
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for renounceOwnership")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaWalletAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareCompleteRequestProcessingOperation prepares a DTA wallet complete request processing operation.
func (e *Extension) PrepareCompleteRequestProcessingOperation(
	requestId [32]byte,
	success bool,
	errorData []byte,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtawallet.DtawalletMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTAWallet ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("completeRequestProcessing", requestId, success, errorData)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for completeRequestProcessing")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaWalletAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// ============================================================================
// Event Decoding (Deprecated - needs migration)
// ============================================================================

// ToJson decodes an encoded VerifiableEvent from a CREC event into a JSON byte slice.
// NOTE: This method is temporarily disabled pending migration to new Event structure.
func (e *Extension) ToJson(event *apiClient.Event) ([]byte, error) {
	return nil, fmt.Errorf("ToJson is temporarily disabled - needs migration to new Event structure")
}

// ============================================================================
// Private helpers
// ============================================================================

func (e *Extension) prepareTokenApproveTransaction(
	tokenAddress *common.Address, tokenAmount *big.Int,
) (*transactTypes.Transaction, error) {
	erc20Abi, err := erc20.Erc20MetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to get ERC20 ABI")
		return nil, err
	}

	calldata, err := erc20Abi.Pack("approve", e.dtaOpenMarketplaceAddress, tokenAmount)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to pack calldata for token approve")
		return nil, err
	}

	return &transactTypes.Transaction{
		To:    *tokenAddress,
		Value: big.NewInt(0),
		Data:  calldata,
	}, nil
}

