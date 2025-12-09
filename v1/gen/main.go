//go:build ignore

// Generator for DTA v1 extension Prepare* functions.
// Run with: go run ./v1/gen/main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

// ABI file paths relative to the v1 directory
const (
	managementABIPath = "abi/DTARequestManagementU.abi.json"
	settlementABIPath = "abi/DTARequestSettlementU.abi.json"
)

// ABIEntry represents a single entry in the ABI JSON.
type ABIEntry struct {
	Type            string     `json:"type"`
	Name            string     `json:"name"`
	Inputs          []ABIInput `json:"inputs"`
	StateMutability string     `json:"stateMutability"`
}

// ABIInput represents an input parameter in the ABI.
type ABIInput struct {
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	InternalType string     `json:"internalType"`
	Components   []ABIInput `json:"components"`
}

// FunctionDef represents a parsed function definition for code generation.
type FunctionDef struct {
	Name           string
	GoName         string
	Params         []ParamDef
	Contract       string
	HelperMethod   string
	HasComplexType bool
}

// ParamDef represents a function parameter.
type ParamDef struct {
	Name       string
	GoType     string
	NeedsCast  bool
	CastToType string
}

// Custom function name overrides for better API design
// Used to disambiguate functions that exist on both contracts or need custom naming
var funcNameOverrides = map[string]string{
	"transferOwnership_DTARequestSettlement": "TransferDTARequestSettlementOwnershipOperation",
	"renounceOwnership_DTARequestSettlement": "RenounceDTARequestSettlementOwnershipOperation",
	// Functions that exist on both contracts need unique names
	"setCCIPGasLimit_DTARequestManagement": "SetManagementCCIPGasLimitOperation",
	"setCCIPGasLimit_DTARequestSettlement": "SetSettlementCCIPGasLimitOperation",
	"withdrawTokens_DTARequestManagement":  "WithdrawManagementTokensOperation",
	"withdrawTokens_DTARequestSettlement":  "WithdrawSettlementTokensOperation",
}

// Custom type overrides (method_param -> GoType)
var typeOverrides = map[string]string{
	"allowDTA_burnType": "TokenBurnType",
}

var managementSkipMethods = map[string]bool{
	"ccipReceive":                      true,
	"completeDistributorRequest":       true,
	"directCompleteDistributorRequest": true,
	"initialize":                       true,
	"recoverFunds":                     true,
	"renounceOwnership":                true,
	"transferOwnership":                true,
}

var settlementSkipMethods = map[string]bool{
	"ccipReceive":            true,
	"ccipHandleDTAMessage":   true,
	"directHandleDTAMessage": true,
	"recoverFunds":           true,
	"_executeSettlement":     true,
	"initialize":             true,
}

// Solidity to Go type mapping
var typeMapping = map[string]string{
	"address": "common.Address",
	"bool":    "bool",
	"bytes":   "[]byte",
	"bytes4":  "[4]byte",
	"bytes32": "[32]byte",
	"string":  "string",
	"uint8":   "uint8",
	"uint24":  "uint32", // Go doesn't have uint24
	"int24":   "int32",  // Go doesn't have int24
	"uint40":  "uint64", // Go doesn't have uint40
	"uint64":  "uint64",
	"uint256": "*big.Int",
	"int256":  "*big.Int",
}

func main() {
	// ABI files are in v1/abi/
	abiDir := "v1"

	managementABI := loadABI(filepath.Join(abiDir, managementABIPath), "DTARequestManagement")
	settlementABI := loadABI(filepath.Join(abiDir, settlementABIPath), "DTARequestSettlement")

	var funcs []FunctionDef

	// Parse DTARequestManagement functions
	mgmtFuncs := parseABI(managementABI, "DTARequestManagement", "prepareManagementOp", managementSkipMethods)
	funcs = append(funcs, mgmtFuncs...)

	// Parse DTARequestSettlement functions
	settleFuncs := parseABI(settlementABI, "DTARequestSettlement", "prepareSettlementOp", settlementSkipMethods)
	funcs = append(funcs, settleFuncs...)

	// Sort by name for consistent output
	sort.Slice(funcs, func(i, j int) bool {
		return funcs[i].GoName < funcs[j].GoName
	})

	// Generate the code
	code := generateCode(funcs)

	// Format the code
	formatted, err := format.Source([]byte(code))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to format generated code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, code)
		os.Exit(1)
	}

	// Write to file
	outputPath := "v1/dta_operations_gen.go"
	if err := os.WriteFile(outputPath, formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s with %d functions\n", outputPath, len(funcs))
}

// loadABI reads an ABI JSON file and returns its contents.
func loadABI(path, contractName string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read %s ABI from %s: %v\n", contractName, path, err)
		os.Exit(1)
	}
	fmt.Printf("Loaded %s ABI from %s\n", contractName, path)
	return string(data)
}

func parseABI(abiJSON, contract, helperMethod string, skipMethods map[string]bool) []FunctionDef {
	var entries []ABIEntry
	if err := json.Unmarshal([]byte(abiJSON), &entries); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse ABI: %v\n", err)
		os.Exit(1)
	}

	var funcs []FunctionDef
	for _, entry := range entries {
		// Only process functions
		if entry.Type != "function" {
			continue
		}

		// Skip view/pure functions (read-only)
		if entry.StateMutability == "view" || entry.StateMutability == "pure" {
			continue
		}

		// Skip explicitly excluded methods
		if skipMethods[entry.Name] {
			continue
		}

		// Check for complex types (tuples)
		hasComplexType := false
		for _, input := range entry.Inputs {
			if strings.HasPrefix(input.Type, "tuple") {
				hasComplexType = true
				break
			}
		}

		// Skip functions with complex types - they need manual handling
		if hasComplexType {
			fmt.Printf("Skipping %s.%s (has complex type)\n", contract, entry.Name)
			continue
		}

		funcDef := FunctionDef{
			Name:           entry.Name,
			GoName:         toGoFuncName(entry.Name, contract),
			Contract:       contract,
			HelperMethod:   helperMethod,
			HasComplexType: hasComplexType,
		}

		for _, input := range entry.Inputs {
			goType := mapType(input.Type)
			if goType == "" {
				fmt.Printf("Warning: Unknown type %s for %s.%s parameter %s\n",
					input.Type, contract, entry.Name, input.Name)
				goType = "interface{}"
			}

			// Check for type overrides
			typeKey := entry.Name + "_" + input.Name
			needsCast := false
			castToType := ""
			if override, ok := typeOverrides[typeKey]; ok {
				castToType = goType // Store original type for cast
				goType = override   // Use override type in signature
				needsCast = true
			}

			funcDef.Params = append(funcDef.Params, ParamDef{
				Name:       toGoParamName(input.Name),
				GoType:     goType,
				NeedsCast:  needsCast,
				CastToType: castToType,
			})
		}

		funcs = append(funcs, funcDef)
	}

	return funcs
}

func mapType(solType string) string {
	if t, ok := typeMapping[solType]; ok {
		return t
	}

	// Handle arrays
	if strings.HasSuffix(solType, "[]") {
		baseType := strings.TrimSuffix(solType, "[]")
		if mapped := mapType(baseType); mapped != "" {
			return "[]" + mapped
		}
	}

	// Handle fixed-size arrays like bytes32[]
	if strings.Contains(solType, "[") && strings.HasSuffix(solType, "]") {
		// This is a fixed-size array like address[10]
		return ""
	}

	return ""
}

func toGoFuncName(name, contract string) string {
	// Check for overrides first
	key := name + "_" + contract
	if override, ok := funcNameOverrides[key]; ok {
		return override
	}

	// Capitalize first letter
	if len(name) == 0 {
		return name
	}
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes) + "Operation"
}

func toGoParamName(name string) string {
	if len(name) == 0 {
		return "arg"
	}
	// Keep first letter lowercase for Go convention
	return name
}

func generateCode(funcs []FunctionDef) string {
	const tmpl = `// Code generated by gen/main.go. DO NOT EDIT.

package v1

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	transactTypes "github.com/smartcontractkit/crec-sdk/transact/types"
)

// Ensure imports are used
var (
	_ = big.NewInt
	_ = common.Address{}
	_ *transactTypes.Operation
)

{{range .}}
// Prepare{{.GoName}} prepares a {{.Name}} operation on {{.Contract}}.
func (e *Extension) Prepare{{.GoName}}({{range $i, $p := .Params}}{{if $i}}, {{end}}{{$p.Name}} {{$p.GoType}}{{end}}) (*transactTypes.Operation, error) {
	return e.{{.HelperMethod}}("{{.Name}}"{{range .Params}}, {{if .NeedsCast}}{{.CastToType}}({{.Name}}){{else}}{{.Name}}{{end}}{{end}})
}
{{end}}
`

	t := template.Must(template.New("code").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, funcs); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute template: %v\n", err)
		os.Exit(1)
	}

	return buf.String()
}
