// Package v1 provides CREC SDK extension for DTA (Digital Token Asset) v1.0 operations.
//
// This version works with DTARequestManagement and DTARequestSettlement contracts.
//
// # Usage
//
//	ext, err := v1.New(&v1.Options{
//		DTARequestManagementAddress: "0x...",
//		DTARequestSettlementAddress: "0x...",
//		AccountAddress:              "0x...",
//	})
package v1

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"

	"github.com/smartcontractkit/crec-api-go/services/dta/gen/dtarequestmanagement"
	"github.com/smartcontractkit/crec-api-go/services/dta/gen/dtarequestsettlement"

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

// Options defines the configuration for creating a new CREC DTA v1 extension.
type Options struct {
	// Logger is an optional logger instance. If nil, a default nop logger is used.
	Logger *zerolog.Logger

	// DTARequestManagementAddress is the address of the DTARequestManagement contract.
	DTARequestManagementAddress string

	// DTARequestSettlementAddress is the address of the DTARequestSettlement contract.
	DTARequestSettlementAddress string

	// AccountAddress is the address of the account performing the DTA operations.
	AccountAddress string
}

// Extension provides methods for preparing DTA v1 operations.
type Extension struct {
	logger                      *zerolog.Logger
	dtaRequestManagementAddress common.Address
	dtaRequestSettlementAddress common.Address
	accountAddress              common.Address
}

// New creates a new CREC DTA v1 extension with the provided options.
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

	logger.Info().Msg("Creating CREC DTA v1 extension")

	return &Extension{
		logger:                      logger,
		dtaRequestManagementAddress: common.HexToAddress(opts.DTARequestManagementAddress),
		dtaRequestSettlementAddress: common.HexToAddress(opts.DTARequestSettlementAddress),
		accountAddress:              common.HexToAddress(opts.AccountAddress),
	}, nil
}

// ============================================================================
// DTARequestManagement Operations
// ============================================================================

// PrepareRequestSubscriptionOperation prepares a DTA request subscription operation.
func (e *Extension) PrepareRequestSubscriptionOperation(
	fundAdminAddr common.Address,
	fundTokenId [32]byte,
	amount *big.Int,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
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
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
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

	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareProcessDistributorRequestOperation prepares a process distributor request operation.
func (e *Extension) PrepareProcessDistributorRequestOperation(requestId [32]byte) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareCancelDistributorRequestOperation prepares a cancel distributor request operation.
func (e *Extension) PrepareCancelDistributorRequestOperation(requestId [32]byte) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRegisterDistributorOperation prepares a register distributor operation.
func (e *Extension) PrepareRegisterDistributorOperation(
	distributorWalletAddr common.Address,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("registerDistributor", distributorWalletAddr)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for registerDistributor")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRegisterFundAdminOperation prepares a register fund admin operation.
func (e *Extension) PrepareRegisterFundAdminOperation(fundAdminAddr common.Address) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRegisterFundTokenOperation prepares a register fund token operation.
func (e *Extension) PrepareRegisterFundTokenOperation(
	fundTokenId [32]byte,
	tokenData dtarequestmanagement.IFundTokenRegistryFundTokenData,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
		return nil, err
	}

	contractTokenData := dtarequestmanagement.IFundTokenRegistryFundTokenData{
		FundTokenAddr:                 tokenData.FundTokenAddr,
		NavFeedDecimals:               tokenData.NavFeedDecimals,
		PurchaseTokenRoundingDecimals: tokenData.PurchaseTokenRoundingDecimals,
		PurchaseTokenDecimals:         tokenData.PurchaseTokenDecimals,
		FundRoundingDecimals:          tokenData.FundRoundingDecimals,
		FundTokenDecimals:             tokenData.FundTokenDecimals,
		RequestsPerDay:                tokenData.RequestsPerDay,
		NavAddr:                       tokenData.NavAddr,
		TokenChainSelector:            tokenData.TokenChainSelector,
		DtaRequestSettlementAddr:      tokenData.DtaRequestSettlementAddr,
		TimezoneOffsetSecs:            tokenData.TimezoneOffsetSecs,
		NavTTL:                        tokenData.NavTTL,
		PaymentInfo: dtarequestmanagement.IDTAMessageDTAPayment{
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
				To:    e.dtaRequestManagementAddress,
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
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
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
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareEnableFundTokenOperation prepares an enable fund token operation.
func (e *Extension) PrepareEnableFundTokenOperation(fundTokenId [32]byte) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareDisableFundTokenOperation prepares a disable fund token operation.
func (e *Extension) PrepareDisableFundTokenOperation(fundTokenId [32]byte) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
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
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareVerifyDistributorWalletOperation prepares a verify distributor wallet operation.
func (e *Extension) PrepareVerifyDistributorWalletOperation(distributorAddr common.Address) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("verifyDistributorWallet", distributorAddr)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for verifyDistributorWallet")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareForceAllowDistributorForTokenOperation prepares a force allow distributor for token operation (admin function).
func (e *Extension) PrepareForceAllowDistributorForTokenOperation(
	fundTokenId [32]byte,
	distributorAddr common.Address,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestmanagement.DtarequestmanagementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestManagement ABI")
		return nil, err
	}

	calldata, err := abiEncoder.Pack("forceAllowDistributorForToken", fundTokenId, distributorAddr)
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to pack calldata for forceAllowDistributorForToken")
		return nil, err
	}

	return &transactTypes.Operation{
		ID:      big.NewInt(time.Now().Unix()),
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    e.dtaRequestManagementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// ============================================================================
// DTARequestSettlement Operations
// ============================================================================

// PrepareAllowDTAOperation prepares a DTARequestSettlement allow DTA operation.
func (e *Extension) PrepareAllowDTAOperation(
	dtaAddr common.Address,
	dtaChainSelector uint64,
	fundTokenId [32]byte,
	fundTokenAddr common.Address,
	burnType TokenBurnType,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestsettlement.DtarequestsettlementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestSettlement ABI")
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
				To:    e.dtaRequestSettlementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareDisallowDTAOperation prepares a DTARequestSettlement disallow DTA operation.
func (e *Extension) PrepareDisallowDTAOperation(
	dtaAddr common.Address,
	dtaChainSelector uint64,
	fundTokenId [32]byte,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestsettlement.DtarequestsettlementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestSettlement ABI")
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
				To:    e.dtaRequestSettlementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareWithdrawTokensOperation prepares a DTARequestSettlement withdraw tokens operation.
func (e *Extension) PrepareWithdrawTokensOperation(
	token common.Address,
	recipient common.Address,
	amount *big.Int,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestsettlement.DtarequestsettlementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestSettlement ABI")
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
				To:    e.dtaRequestSettlementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareTransferDTARequestSettlementOwnershipOperation prepares a DTARequestSettlement transfer ownership operation.
func (e *Extension) PrepareTransferDTARequestSettlementOwnershipOperation(newOwner common.Address) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestsettlement.DtarequestsettlementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestSettlement ABI")
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
				To:    e.dtaRequestSettlementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareRenounceDTARequestSettlementOwnershipOperation prepares a DTARequestSettlement renounce ownership operation.
func (e *Extension) PrepareRenounceDTARequestSettlementOwnershipOperation() (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestsettlement.DtarequestsettlementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestSettlement ABI")
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
				To:    e.dtaRequestSettlementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// PrepareCompleteRequestProcessingOperation prepares a DTARequestSettlement complete request processing operation.
func (e *Extension) PrepareCompleteRequestProcessingOperation(
	requestId [32]byte,
	success bool,
	errorData []byte,
) (*transactTypes.Operation, error) {
	abiEncoder, err := dtarequestsettlement.DtarequestsettlementMetaData.GetAbi()
	if err != nil {
		e.logger.Error().Err(err).Msg("Failed to get DTARequestSettlement ABI")
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
				To:    e.dtaRequestSettlementAddress,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
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

	calldata, err := erc20Abi.Pack("approve", e.dtaRequestManagementAddress, tokenAmount)
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

