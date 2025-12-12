package v1_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	v1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
	transactTypes "github.com/smartcontractkit/crec-sdk/transact/types"
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

func TestDtaV1_PrepareRegisterFundTokenOperation(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))

	tokenData := v1.FundTokenData{
		FundTokenAddr:                 common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
		NavFeedDecimals:               8,
		PurchaseTokenRoundingDecimals: 2,
		PurchaseTokenDecimals:         6,
		FundRoundingDecimals:          2,
		FundTokenDecimals:             18,
		RequestsPerDay:                1,
		NavAddr:                       common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
		TokenChainSelector:            1234,
		DtaRequestSettlementAddr:      common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"),
		TimezoneOffsetSecs:            big.NewInt(0),
		NavTTL:                        big.NewInt(3600),
		PaymentInfo: v1.DTAPayment{
			OffChainPaymentCurrency: 0,
			PaymentTokenSourceAddr:  common.HexToAddress("0xDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD"),
			PaymentTokenDestAddr:    common.HexToAddress("0xEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE"),
		},
	}

	op, err := ext.PrepareRegisterFundTokenOperation(fundTokenId, tokenData)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[0].To)
	require.NotEmpty(t, op.Transactions[0].Data)
}

func TestDtaV1_PrepareDistributorOperations(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))
	distributorAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	tests := []struct {
		name           string
		prepareOp      func() (interface{}, error)
		expectedTarget common.Address
	}{
		{
			name: "AllowDistributorForToken",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareAllowDistributorForTokenOperation(fundTokenId, distributorAddr)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
		{
			name: "DisallowDistributorForToken",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareDisallowDistributorForTokenOperation(fundTokenId, distributorAddr)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
		{
			name: "ForceAllowDistributorForToken",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareForceAllowDistributorForTokenOperation(fundTokenId, distributorAddr)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
		{
			name: "VerifyDistributorWallet",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareVerifyDistributorWalletOperation(distributorAddr)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.prepareOp()
			require.NoError(t, err)
			require.NotNil(t, result)

			op, ok := result.(*transactTypes.Operation)
			require.True(t, ok)
			require.Len(t, op.Transactions, 1)
			require.Equal(t, tt.expectedTarget, op.Transactions[0].To)
		})
	}
}

func TestDtaV1_PrepareFundTokenOperations(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))

	tests := []struct {
		name           string
		prepareOp      func() (interface{}, error)
		expectedTarget common.Address
	}{
		{
			name: "EnableFundToken",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareEnableFundTokenOperation(fundTokenId)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
		{
			name: "DisableFundToken",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareDisableFundTokenOperation(fundTokenId)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.prepareOp()
			require.NoError(t, err)
			require.NotNil(t, result)

			op, ok := result.(*transactTypes.Operation)
			require.True(t, ok)
			require.Len(t, op.Transactions, 1)
			require.Equal(t, tt.expectedTarget, op.Transactions[0].To)
		})
	}
}

func TestDtaV1_PrepareDistributorRequestOperations(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	var requestId [32]byte
	copy(requestId[:], []byte("testrequest"))

	tests := []struct {
		name           string
		prepareOp      func() (interface{}, error)
		expectedTarget common.Address
	}{
		{
			name: "CancelDistributorRequest",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareCancelDistributorRequestOperation(requestId)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
		{
			name: "ProcessDistributorRequest",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareProcessDistributorRequestOperation(requestId)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.prepareOp()
			require.NoError(t, err)
			require.NotNil(t, result)

			op, ok := result.(*transactTypes.Operation)
			require.True(t, ok)
			require.Len(t, op.Transactions, 1)
			require.Equal(t, tt.expectedTarget, op.Transactions[0].To)
		})
	}
}

func TestDtaV1_PrepareOwnershipOperations(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	newOwner := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	tests := []struct {
		name           string
		prepareOp      func() (interface{}, error)
		expectedTarget common.Address
	}{
		{
			name: "RenounceDTARequestSettlementOwnership",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareRenounceDTARequestSettlementOwnershipOperation()
			},
			expectedTarget: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		},
		{
			name: "TransferDTARequestSettlementOwnership",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareTransferDTARequestSettlementOwnershipOperation(newOwner)
			},
			expectedTarget: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.prepareOp()
			require.NoError(t, err)
			require.NotNil(t, result)

			op, ok := result.(*transactTypes.Operation)
			require.True(t, ok)
			require.Len(t, op.Transactions, 1)
			require.Equal(t, tt.expectedTarget, op.Transactions[0].To)
		})
	}
}

func TestDtaV1_PrepareWithdrawTokensOperations(t *testing.T) {
	ext, err := v1.New(&v1.Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	token := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	recipient := common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	amount := big.NewInt(1000000)

	tests := []struct {
		name           string
		prepareOp      func() (interface{}, error)
		expectedTarget common.Address
	}{
		{
			name: "WithdrawManagementTokens",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareWithdrawManagementTokensOperation(token, recipient, amount)
			},
			expectedTarget: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		},
		{
			name: "WithdrawSettlementTokens",
			prepareOp: func() (interface{}, error) {
				return ext.PrepareWithdrawSettlementTokensOperation(token, recipient, amount)
			},
			expectedTarget: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.prepareOp()
			require.NoError(t, err)
			require.NotNil(t, result)

			op, ok := result.(*transactTypes.Operation)
			require.True(t, ok)
			require.Len(t, op.Transactions, 1)
			require.Equal(t, tt.expectedTarget, op.Transactions[0].To)
		})
	}
}
