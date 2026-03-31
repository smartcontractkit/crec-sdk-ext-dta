package operations

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestPreparedOperationsDefaultDeadlineToZero(t *testing.T) {
	ext, err := New(&Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
	})
	require.NoError(t, err)

	op, err := ext.PrepareRegisterFundAdminOperation()
	require.NoError(t, err)
	require.NotNil(t, op.Deadline)
	require.Zero(t, op.Deadline.Sign())
}

func TestPreparedOperationsCloneConfiguredDeadline(t *testing.T) {
	configuredDeadline := big.NewInt(123456789)

	ext, err := New(&Options{
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		AccountAddress:              "0x3333333333333333333333333333333333333333",
		Deadline:                    configuredDeadline,
	})
	require.NoError(t, err)

	configuredDeadline.SetInt64(1)

	op1, err := ext.PrepareRegisterFundAdminOperation()
	require.NoError(t, err)
	require.NotNil(t, op1.Deadline)
	require.Equal(t, "123456789", op1.Deadline.String())

	op1.Deadline.SetInt64(42)

	op2, err := ext.PrepareRegisterDistributorOperation(common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
	require.NoError(t, err)
	require.NotNil(t, op2.Deadline)
	require.Equal(t, "123456789", op2.Deadline.String())
}
