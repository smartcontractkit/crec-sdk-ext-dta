//go:build ignore

package main

// =============================================================================
// GENERATOR CONFIGURATION
// =============================================================================
// Configuration for DTA SDK extension code generation.
// Run with: go run ./gen/main.go ./gen/config.go
//
// The generator will create:
//   - abi_gen.go           : ABI embedding and accessors
//   - extension_gen.go     : Options, Extension struct, New()
//   - operations_gen.go    : Type-safe Prepare* functions
//   - operations_helpers_gen.go : Helper methods per contract
//   - events_gen.go        : Event structs and constants
//   - decode_gen.go        : Event decoders

// -----------------------------------------------------------------------------
// Output Configuration
// -----------------------------------------------------------------------------

// Package name for generated code
const packageName = "v1"

// Module path for imports in generated code
const modulePath = "github.com/smartcontractkit/crec-sdk-ext-dta/v1"

// -----------------------------------------------------------------------------
// Contract Definitions
// -----------------------------------------------------------------------------
// DTA uses two main contracts: DTARequestManagement and DTARequestSettlement.

// ContractConfig defines a contract for code generation.
type ContractConfig struct {
	Name              string            // Contract name (e.g., "DTARequestManagement")
	ABIFile           string            // Path to ABI file (e.g., "abi/DTARequestManagementU.abi.json")
	SkipMethods       map[string]bool   // Methods to skip (e.g., internal, admin-only)
	FuncNameOverrides map[string]string // Method name overrides (method -> custom name)
}

// contracts defines all contracts to generate code for.
var contracts = []ContractConfig{
	{
		Name:    "DTARequestManagement",
		ABIFile: "abi/DTARequestManagementU.abi.json",
		SkipMethods: map[string]bool{
			"ccipReceive":                      true,
			"completeDistributorRequest":       true,
			"directCompleteDistributorRequest": true,
			"initialize":                       true,
			"recoverFunds":                     true,
			"renounceOwnership":                true,
			"transferOwnership":                true,
		},
		FuncNameOverrides: map[string]string{
			// Disambiguate methods that exist on both contracts
			"setCCIPGasLimit": "SetManagementCCIPGasLimitOperation",
			"withdrawTokens":  "WithdrawManagementTokensOperation",
		},
	},
	{
		Name:    "DTARequestSettlement",
		ABIFile: "abi/DTARequestSettlementU.abi.json",
		SkipMethods: map[string]bool{
			"ccipReceive":            true,
			"ccipHandleDTAMessage":   true,
			"directHandleDTAMessage": true,
			"recoverFunds":           true,
			"_executeSettlement":     true,
			"initialize":             true,
		},
		FuncNameOverrides: map[string]string{
			// Disambiguate methods that exist on both contracts
			"setCCIPGasLimit":   "SetSettlementCCIPGasLimitOperation",
			"withdrawTokens":    "WithdrawSettlementTokensOperation",
			"transferOwnership": "TransferDTARequestSettlementOwnershipOperation",
			"renounceOwnership": "RenounceDTARequestSettlementOwnershipOperation",
		},
	},
}

// -----------------------------------------------------------------------------
// Type Overrides
// -----------------------------------------------------------------------------
// Use this to override Solidity types with custom Go types.
// Key format: "methodName_paramName" -> "GoTypeName"
var typeOverrides = map[string]string{
	"allowDTA_mintType": "TokenMintType",
	"allowDTA_burnType": "TokenBurnType",
}

// -----------------------------------------------------------------------------
// Enum Type Mappings
// -----------------------------------------------------------------------------
// Maps Solidity enum internalType strings to Go type names.
// The generator uses this to automatically type event fields as enums.
var enumTypeMapping = map[string]string{
	"enum IDTAMessage.DistributorRequestType":  "DistributorRequestType",
	"enum IDTAMessage.RequestStatus":           "RequestStatus",
	"enum IDTARequestSettlement.TokenMintType": "TokenMintType",
	"enum IDTARequestSettlement.TokenBurnType": "TokenBurnType",
}

// -----------------------------------------------------------------------------
// Events Configuration
// -----------------------------------------------------------------------------
// Events to skip during generation (common/inherited events you don't need)
var skipEvents = map[string]bool{}

// Extra events not in any ABI but needed for decoding
// Use this for events from interfaces or other sources
var extraEvents = []EventDef{
	{
		Name: "AnswerUpdated",
		Fields: []FieldDef{
			{Name: "current", GoName: "Current", GoType: "*big.Int", JSONTag: "current", ParseFunc: `parsing.ScientificNotationToBigInt(params["current"])`, HasError: true, LocalVar: "current", ValueExpr: "current"},
			{Name: "roundId", GoName: "RoundId", GoType: "*big.Int", JSONTag: "round_id", ParseFunc: `parsing.ScientificNotationToBigInt(params["round_id"])`, HasError: true, LocalVar: "roundId", ValueExpr: "roundId"},
			{Name: "updatedAt", GoName: "UpdatedAt", GoType: "*big.Int", JSONTag: "updated_at", ParseFunc: `parsing.ScientificNotationToBigInt(params["updated_at"])`, HasError: true, LocalVar: "updatedAt", ValueExpr: "updatedAt"},
		},
	},
}
