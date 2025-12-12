package v1

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/crec-sdk/interfaces/erc20"
	transactTypes "github.com/smartcontractkit/crec-sdk/transact/types"
)

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

// prepareTokenApproveTransaction creates an ERC20 approve transaction.
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
