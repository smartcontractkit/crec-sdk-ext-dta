package v1_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	v1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
)

func TestDtaV1_New(t *testing.T) {
	tests := []struct {
		name       string
		opts       *v1.Options
		wantErr    bool
		errContain string
	}{
		{
			name:       "nil options returns error",
			opts:       nil,
			wantErr:    true,
			errContain: "options is required",
		},
		{
			name: "valid options creates extension",
			opts: &v1.Options{
				DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
				DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
				AccountAddress:              "0x3333333333333333333333333333333333333333",
			},
			wantErr: false,
		},
		{
			name: "invalid management address",
			opts: &v1.Options{
				DTARequestManagementAddress: "invalid",
				DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
				AccountAddress:              "0x3333333333333333333333333333333333333333",
			},
			wantErr:    true,
			errContain: "invalid DTARequestManagementAddress",
		},
		{
			name: "invalid settlement address",
			opts: &v1.Options{
				DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
				DTARequestSettlementAddress: "not-an-address",
				AccountAddress:              "0x3333333333333333333333333333333333333333",
			},
			wantErr:    true,
			errContain: "invalid DTARequestSettlementAddress",
		},
		{
			name: "invalid account address",
			opts: &v1.Options{
				DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
				DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
				AccountAddress:              "",
			},
			wantErr:    true,
			errContain: "invalid AccountAddress",
		},
		{
			name: "empty management address",
			opts: &v1.Options{
				DTARequestManagementAddress: "",
				DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
				AccountAddress:              "0x3333333333333333333333333333333333333333",
			},
			wantErr:    true,
			errContain: "invalid DTARequestManagementAddress",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ext, err := v1.New(tt.opts)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, ext)
				if tt.errContain != "" {
					require.Contains(t, err.Error(), tt.errContain)
				}
			} else {
				require.NoError(t, err)
				require.NotNil(t, ext)
			}
		})
	}
}

func TestDtaV1_PrepareRequestSubscriptionOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	fundAdminAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))
	amount := big.NewInt(1000000)

	op, err := ext.PrepareRequestSubscriptionOperation(fundAdminAddr, fundTokenId, amount)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[0].To)
}

func TestDtaV1_PrepareRequestSubscriptionWithTokenApprovalOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	fundAdminAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))
	amount := big.NewInt(1000000)
	paymentToken := common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")

	op, err := ext.PrepareRequestSubscriptionWithTokenApprovalOperation(fundAdminAddr, fundTokenId, amount, paymentToken)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 2)
	// First tx is approval to payment token
	require.Equal(t, paymentToken, op.Transactions[0].To)
	// Second tx is subscription to request management
	require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[1].To)
}

func TestDtaV1_PrepareAllowDTAOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	dtaAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fundAdminAddr := common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))
	fundTokenAddr := common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")

	op, err := ext.PrepareAllowDTAOperation(dtaAddr, 1234, fundAdminAddr, fundTokenId, fundTokenAddr, v1.TokenMintTypeIssueTokens, v1.TokenBurnTypeBurn)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	// Transaction goes to settlement address
	require.Equal(t, common.HexToAddress("0x2222222222222222222222222222222222222222"), op.Transactions[0].To)
}

func TestDtaV1_PrepareCompleteRequestProcessingOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	var requestId [32]byte
	copy(requestId[:], []byte("testrequest"))

	op, err := ext.PrepareCompleteRequestProcessingOperation(requestId, true, []byte{})
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	// Transaction goes to settlement address
	require.Equal(t, common.HexToAddress("0x2222222222222222222222222222222222222222"), op.Transactions[0].To)
}

func TestDtaV1_PrepareRegisterFundAdminOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	op, err := ext.PrepareRegisterFundAdminOperation()
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[0].To)
}

func TestDtaV1_PrepareRegisterDistributorOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	distributorWalletAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	op, err := ext.PrepareRegisterDistributorOperation(distributorWalletAddr)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[0].To)
}

func TestDtaV1_PrepareRequestRedemptionOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	fundAdminAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))
	shares := big.NewInt(1000000)

	op, err := ext.PrepareRequestRedemptionOperation(fundAdminAddr, fundTokenId, shares)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[0].To)
}

func TestDtaV1_PrepareDisallowDTAOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	dtaAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fundAdminAddr := common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))

	op, err := ext.PrepareDisallowDTAOperation(dtaAddr, 1234, fundAdminAddr, fundTokenId)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	require.Equal(t, common.HexToAddress("0x2222222222222222222222222222222222222222"), op.Transactions[0].To)
}

func TestDtaV1_PrepareSetManagementCCIPGasLimitOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	gasLimit := big.NewInt(500000)

	op, err := ext.PrepareSetManagementCCIPGasLimitOperation(gasLimit)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[0].To)
}

func TestDtaV1_PrepareSetSettlementCCIPGasLimitOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	gasLimit := big.NewInt(500000)

	op, err := ext.PrepareSetSettlementCCIPGasLimitOperation(gasLimit)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	require.Equal(t, common.HexToAddress("0x2222222222222222222222222222222222222222"), op.Transactions[0].To)
}

func TestDtaV1_PrepareGenericOperations(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	t.Run("management operation", func(t *testing.T) {
		op, err := ext.PrepareManagementOperation("registerFundAdmin")
		require.NoError(t, err)
		require.NotNil(t, op)
		require.Len(t, op.Transactions, 1)
		require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[0].To)
	})

	t.Run("settlement operation", func(t *testing.T) {
		op, err := ext.PrepareSettlementOperation("renounceOwnership")
		require.NoError(t, err)
		require.NotNil(t, op)
		require.Len(t, op.Transactions, 1)
		require.Equal(t, common.HexToAddress("0x2222222222222222222222222222222222222222"), op.Transactions[0].To)
	})

	t.Run("invalid method returns error", func(t *testing.T) {
		_, err := ext.PrepareManagementOperation("nonExistentMethod")
		require.Error(t, err)
		require.Contains(t, err.Error(), "pack nonExistentMethod")
	})
}

func TestDtaV1_OperationIDsAreUnique(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	// Create multiple operations and ensure IDs are unique
	seenIDs := make(map[string]bool)
	for i := 0; i < 100; i++ {
		op, err := ext.PrepareRegisterFundAdminOperation()
		require.NoError(t, err)
		idStr := op.ID.String()
		require.False(t, seenIDs[idStr], "duplicate operation ID: %s", idStr)
		seenIDs[idStr] = true
	}
}
