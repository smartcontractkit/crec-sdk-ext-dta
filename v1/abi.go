package v1

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

//go:embed abi/DTARequestManagementU.abi.json
var dtaRequestManagementABIJSON string

//go:embed abi/DTARequestSettlementU.abi.json
var dtaRequestSettlementABIJSON string

var (
	dtaRequestManagementABI *abi.ABI
	dtaRequestSettlementABI *abi.ABI
)

func init() {
	var err error

	dtaRequestManagementABI, err = parseABI(dtaRequestManagementABIJSON)
	if err != nil {
		panic(fmt.Sprintf("failed to parse DTARequestManagement ABI: %v", err))
	}

	dtaRequestSettlementABI, err = parseABI(dtaRequestSettlementABIJSON)
	if err != nil {
		panic(fmt.Sprintf("failed to parse DTARequestSettlement ABI: %v", err))
	}
}

func parseABI(jsonABI string) (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(jsonABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// ManagementABI returns the parsed DTARequestManagement ABI.
func ManagementABI() *abi.ABI {
	return dtaRequestManagementABI
}

// SettlementABI returns the parsed DTARequestSettlement ABI.
func SettlementABI() *abi.ABI {
	return dtaRequestSettlementABI
}
