package operations

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/crec-sdk-ext-dta/v1/events"
	"github.com/smartcontractkit/crec-sdk/interfaces/erc20"
	transactTypes "github.com/smartcontractkit/crec-sdk/transact/types"
)

// ============================================================================
// Special Operations (multi-transaction or complex types)
// ============================================================================
// These operations cannot be auto-generated because they either:
// - Involve multiple transactions (e.g., approve + call)
// - Have complex struct types that require manual handling

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

	calldata, err := DTARequestManagementABI().Pack("requestSubscription", fundAdminAddr, fundTokenId, amount)
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
	tokenData events.FundTokenData,
) (*transactTypes.Operation, error) {
	calldata, err := DTARequestManagementABI().Pack("registerFundToken", fundTokenId, tokenData)
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
// Internal Helpers (not generated)
// ============================================================================

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
