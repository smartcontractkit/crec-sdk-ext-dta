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
	"crypto/rand"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

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

// DTAPayment represents payment information for DTA operations.
type DTAPayment struct {
	OffChainPaymentCurrency uint8
	PaymentTokenSourceAddr  common.Address
	PaymentTokenDestAddr    common.Address
}

// FundTokenData represents the data for registering a fund token.
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

// Options defines the configuration for creating a new CREC DTA v1 extension.
type Options struct {
	// Logger is an optional logger instance. If nil, a default nop logger is used.
	Logger *slog.Logger

	// DTARequestManagementAddress is the address of the DTARequestManagement contract.
	DTARequestManagementAddress string

	// DTARequestSettlementAddress is the address of the DTARequestSettlement contract.
	DTARequestSettlementAddress string

	// AccountAddress is the address of the account performing the DTA operations.
	AccountAddress string
}

// validate checks that all required options are valid.
func (o *Options) validate() error {
	if !common.IsHexAddress(o.DTARequestManagementAddress) {
		return fmt.Errorf("invalid DTARequestManagementAddress: %q", o.DTARequestManagementAddress)
	}
	if !common.IsHexAddress(o.DTARequestSettlementAddress) {
		return fmt.Errorf("invalid DTARequestSettlementAddress: %q", o.DTARequestSettlementAddress)
	}
	if !common.IsHexAddress(o.AccountAddress) {
		return fmt.Errorf("invalid AccountAddress: %q", o.AccountAddress)
	}
	return nil
}

// Extension provides methods for preparing DTA v1 operations.
type Extension struct {
	logger                      *slog.Logger
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

	if err := opts.validate(); err != nil {
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	logger := opts.Logger
	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("creating CREC DTA v1 extension")

	return &Extension{
		logger:                      logger,
		dtaRequestManagementAddress: common.HexToAddress(opts.DTARequestManagementAddress),
		dtaRequestSettlementAddress: common.HexToAddress(opts.DTARequestSettlementAddress),
		accountAddress:              common.HexToAddress(opts.AccountAddress),
	}, nil
}

// ============================================================================
// Generic Operation Builders (for power users)
// ============================================================================

// PrepareManagementOperation prepares a generic DTARequestManagement operation.
// This is an escape hatch for power users who need to call contract methods
// not covered by the type-safe Prepare* functions.
func (e *Extension) PrepareManagementOperation(method string, args ...interface{}) (*transactTypes.Operation, error) {
	return e.prepareOperation(
		ManagementABI(),
		e.dtaRequestManagementAddress,
		method,
		args...,
	)
}

// PrepareSettlementOperation prepares a generic DTARequestSettlement operation.
// This is an escape hatch for power users who need to call contract methods
// not covered by the type-safe Prepare* functions.
func (e *Extension) PrepareSettlementOperation(method string, args ...interface{}) (*transactTypes.Operation, error) {
	return e.prepareOperation(
		SettlementABI(),
		e.dtaRequestSettlementAddress,
		method,
		args...,
	)
}

// ============================================================================
// Internal Helpers
// ============================================================================

// prepareOperation is the internal helper that all operation builders use.
func (e *Extension) prepareOperation(
	contractABI *abi.ABI,
	target common.Address,
	method string,
	args ...interface{},
) (*transactTypes.Operation, error) {
	calldata, err := contractABI.Pack(method, args...)
	if err != nil {
		e.logger.Error("failed to pack calldata", "method", method, "error", err)
		return nil, fmt.Errorf("pack %s: %w", method, err)
	}

	opID, err := generateOperationID()
	if err != nil {
		e.logger.Error("failed to generate operation ID", "error", err)
		return nil, fmt.Errorf("generate operation ID: %w", err)
	}

	return &transactTypes.Operation{
		ID:      opID,
		Account: e.accountAddress,
		Transactions: []transactTypes.Transaction{
			{
				To:    target,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

// generateOperationID creates a cryptographically random operation ID.
func generateOperationID() (*big.Int, error) {
	// Generate 16 bytes of random data (128 bits)
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(b), nil
}

// prepareManagementOp is a convenience wrapper for DTARequestManagement operations.
func (e *Extension) prepareManagementOp(method string, args ...interface{}) (*transactTypes.Operation, error) {
	return e.prepareOperation(
		ManagementABI(),
		e.dtaRequestManagementAddress,
		method,
		args...,
	)
}

// prepareSettlementOp is a convenience wrapper for DTARequestSettlement operations.
func (e *Extension) prepareSettlementOp(method string, args ...interface{}) (*transactTypes.Operation, error) {
	return e.prepareOperation(
		SettlementABI(),
		e.dtaRequestSettlementAddress,
		method,
		args...,
	)
}

// ============================================================================
// Special Operations (multi-transaction or complex types)
// ============================================================================

// PrepareRequestSubscriptionWithTokenApprovalOperation prepares a subscription operation with token approval.
func (e *Extension) PrepareRequestSubscriptionWithTokenApprovalOperation(
	fundAdminAddr common.Address,
	fundTokenId [32]byte,
	amount *big.Int,
	paymentTokenAddress common.Address,
) (*transactTypes.Operation, error) {
	approveTransaction, err := e.prepareTokenApproveTransaction(&paymentTokenAddress, amount)
	if err != nil {
		e.logger.Error("failed to prepare token approve transaction", "error", err)
		return nil, fmt.Errorf("prepare token approve: %w", err)
	}

	calldata, err := ManagementABI().Pack("requestSubscription", fundAdminAddr, fundTokenId, amount)
	if err != nil {
		e.logger.Error("failed to pack calldata for requestSubscription", "error", err)
		return nil, fmt.Errorf("pack requestSubscription: %w", err)
	}

	opID, err := generateOperationID()
	if err != nil {
		e.logger.Error("failed to generate operation ID", "error", err)
		return nil, fmt.Errorf("generate operation ID: %w", err)
	}

	return &transactTypes.Operation{
		ID:      opID,
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

// PrepareRegisterFundTokenOperation prepares a register fund token operation.
// This function has a complex struct type that cannot be auto-generated.
func (e *Extension) PrepareRegisterFundTokenOperation(
	fundTokenId [32]byte,
	tokenData FundTokenData,
) (*transactTypes.Operation, error) {
	calldata, err := ManagementABI().Pack("registerFundToken", fundTokenId, tokenData)
	if err != nil {
		e.logger.Error("failed to pack calldata for registerFundToken", "error", err)
		return nil, fmt.Errorf("pack registerFundToken: %w", err)
	}

	opID, err := generateOperationID()
	if err != nil {
		e.logger.Error("failed to generate operation ID", "error", err)
		return nil, fmt.Errorf("generate operation ID: %w", err)
	}

	return &transactTypes.Operation{
		ID:      opID,
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
// Private helpers
// ============================================================================

func (e *Extension) prepareTokenApproveTransaction(
	tokenAddress *common.Address, tokenAmount *big.Int,
) (*transactTypes.Transaction, error) {
	erc20Abi, err := erc20.Erc20MetaData.GetAbi()
	if err != nil {
		e.logger.Error("failed to get ERC20 ABI", "error", err)
		return nil, fmt.Errorf("get ERC20 ABI: %w", err)
	}

	calldata, err := erc20Abi.Pack("approve", e.dtaRequestManagementAddress, tokenAmount)
	if err != nil {
		e.logger.Error("failed to pack calldata for token approve", "error", err)
		return nil, fmt.Errorf("pack approve: %w", err)
	}

	return &transactTypes.Transaction{
		To:    *tokenAddress,
		Value: big.NewInt(0),
		Data:  calldata,
	}, nil
}
