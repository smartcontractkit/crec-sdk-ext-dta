package operations

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestPreparedOperationsDefaultDeadlineToZero(t *testing.T) {
	ext, err := New(&Options{
		AccountAddress:    "0x1111111111111111111111111111111111111111",
		FeeManagerAddress: "0x2222222222222222222222222222222222222222",
	})
	require.NoError(t, err)

	op, err := ext.PrepareCreateUserTokenAccountOperation(common.HexToAddress("0x3333333333333333333333333333333333333333"))
	require.NoError(t, err)
	require.NotNil(t, op.Deadline)
	require.Zero(t, op.Deadline.Sign())
}

func TestPreparedOperationsCloneConfiguredDeadline(t *testing.T) {
	configuredDeadline := big.NewInt(555666777)

	ext, err := New(&Options{
		AccountAddress:    "0x1111111111111111111111111111111111111111",
		FeeManagerAddress: "0x2222222222222222222222222222222222222222",
		Deadline:          configuredDeadline,
	})
	require.NoError(t, err)

	configuredDeadline.SetInt64(1)

	op1, err := ext.PrepareCreateUserTokenAccountOperation(common.HexToAddress("0x3333333333333333333333333333333333333333"))
	require.NoError(t, err)
	require.NotNil(t, op1.Deadline)
	require.Equal(t, "555666777", op1.Deadline.String())

	op1.Deadline.SetInt64(42)

	var actionID [32]byte
	actionID[31] = 1

	op2, err := ext.PrepareAllowProductForActionOperation(
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
		actionID,
		common.HexToAddress("0x5555555555555555555555555555555555555555"),
	)
	require.NoError(t, err)
	require.NotNil(t, op2.Deadline)
	require.Equal(t, "555666777", op2.Deadline.String())
}
