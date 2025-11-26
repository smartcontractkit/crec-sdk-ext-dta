package v0_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	v0 "github.com/smartcontractkit/crec-sdk-ext-dta/v0.1"
)

func TestDtaV0_New(t *testing.T) {
	tests := []struct {
		name    string
		opts    *v0.Options
		wantErr bool
	}{
		{
			name:    "nil options returns error",
			opts:    nil,
			wantErr: true,
		},
		{
			name: "valid options creates extension",
			opts: &v0.Options{
				DTAOpenMarketplaceAddress: "0x1111111111111111111111111111111111111111",
				DTAWalletAddress:          "0x2222222222222222222222222222222222222222",
				AccountAddress:            "0x3333333333333333333333333333333333333333",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ext, err := v0.New(tt.opts)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, ext)
			} else {
				require.NoError(t, err)
				require.NotNil(t, ext)
			}
		})
	}
}

func TestDtaV0_PrepareRequestSubscriptionOperation(t *testing.T) {
	ext, err := v0.New(&v0.Options{
		DTAOpenMarketplaceAddress: "0x1111111111111111111111111111111111111111",
		DTAWalletAddress:          "0x2222222222222222222222222222222222222222",
		AccountAddress:            "0x3333333333333333333333333333333333333333",
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

func TestDtaV0_PrepareRequestSubscriptionWithTokenApprovalOperation(t *testing.T) {
	ext, err := v0.New(&v0.Options{
		DTAOpenMarketplaceAddress: "0x1111111111111111111111111111111111111111",
		DTAWalletAddress:          "0x2222222222222222222222222222222222222222",
		AccountAddress:            "0x3333333333333333333333333333333333333333",
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
	// Second tx is subscription to marketplace
	require.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), op.Transactions[1].To)
}

func TestDtaV0_PrepareAllowDTAOperation(t *testing.T) {
	ext, err := v0.New(&v0.Options{
		DTAOpenMarketplaceAddress: "0x1111111111111111111111111111111111111111",
		DTAWalletAddress:          "0x2222222222222222222222222222222222222222",
		AccountAddress:            "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	dtaAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	var fundTokenId [32]byte
	copy(fundTokenId[:], []byte("testtoken"))
	fundTokenAddr := common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")

	op, err := ext.PrepareAllowDTAOperation(dtaAddr, 1234, fundTokenId, fundTokenAddr, v0.TokenBurnTypeBurn)
	require.NoError(t, err)
	require.NotNil(t, op)
	require.Len(t, op.Transactions, 1)
	// Transaction goes to wallet address
	require.Equal(t, common.HexToAddress("0x2222222222222222222222222222222222222222"), op.Transactions[0].To)
}

