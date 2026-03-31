package certify

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestNewRegisterFundAdminOperationBuilderSupportsV1AndV2(t *testing.T) {
	tests := []struct {
		name    string
		service string
	}{
		{name: "v1", service: dtaServiceV1},
		{name: "v2", service: dtaServiceV2},
	}

	const (
		accountAddress    = "0x3333333333333333333333333333333333333333"
		managementAddress = "0x1111111111111111111111111111111111111111"
		settlementAddress = "0x2222222222222222222222222222222222222222"
	)

	deadline := big.NewInt(123456789)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				DTAService:                  tt.service,
				DTARequestManagementAddress: managementAddress,
				DTARequestSettlementAddress: settlementAddress,
			}

			builder, err := newRegisterFundAdminOperationBuilder(cfg, accountAddress, deadline)
			require.NoError(t, err)

			op, err := builder.PrepareRegisterFundAdminOperation()
			require.NoError(t, err)

			require.Equal(t, common.HexToAddress(accountAddress), op.Account)
			require.NotNil(t, op.Deadline)
			require.Equal(t, deadline.String(), op.Deadline.String())
			require.Len(t, op.Transactions, 1)
			require.Equal(t, common.HexToAddress(managementAddress), op.Transactions[0].To)
		})
	}
}

func TestNewRegisterFundAdminOperationBuilderRejectsUnsupportedService(t *testing.T) {
	cfg := &Config{
		DTAService:                  "dta.v3",
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
	}

	_, err := newRegisterFundAdminOperationBuilder(
		cfg,
		"0x3333333333333333333333333333333333333333",
		big.NewInt(1),
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), `unsupported DTA_SERVICE: "dta.v3"`)
}
