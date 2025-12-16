//go:build ignore

// Generator for CREC SDK Extension - generates all SDK boilerplate from ABIs.
// Run with: go run ./gen/main.go ./gen/config.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

// =============================================================================
// ABI Types
// =============================================================================

// ABIEntry represents a single entry in the ABI JSON.
type ABIEntry struct {
	Type            string     `json:"type"`
	Name            string     `json:"name"`
	Inputs          []ABIInput `json:"inputs"`
	StateMutability string     `json:"stateMutability"`
	Anonymous       bool       `json:"anonymous"`
}

// ABIInput represents an input parameter in the ABI.
type ABIInput struct {
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	InternalType string     `json:"internalType"`
	Indexed      bool       `json:"indexed"`
	Components   []ABIInput `json:"components"`
}

// =============================================================================
// Code Generation Types
// =============================================================================

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

// EventDef represents a parsed event definition for code generation.
type EventDef struct {
	Name   string
	Fields []FieldDef
}

// FieldDef represents an event field.
type FieldDef struct {
	Name      string
	GoName    string
	GoType    string
	JSONTag   string
	ParseFunc string
	HasError  bool
	IsEnum    bool
	EnumType  string
	LocalVar  string
	ValueExpr string
}

// =============================================================================
// Type Mappings (Solidity -> Go)
// =============================================================================

var typeMapping = map[string]string{
	"address": "common.Address",
	"bool":    "bool",
	"bytes":   "[]byte",
	"bytes4":  "[4]byte",
	"bytes32": "[32]byte",
	"string":  "string",
	"uint8":   "uint8",
	"uint16":  "uint16",
	"uint24":  "uint32",
	"int24":   "int32",
	"uint32":  "uint32",
	"uint40":  "uint64",
	"uint64":  "uint64",
	"uint128": "*big.Int",
	"uint256": "*big.Int",
	"int128":  "*big.Int",
	"int256":  "*big.Int",
}

var eventTypeMapping = map[string]string{
	"address": "common.Address",
	"bool":    "bool",
	"bytes":   "[]byte",
	"bytes4":  "[4]byte",
	"bytes32": "common.Hash",
	"string":  "string",
	"uint8":   "uint8",
	"uint16":  "uint16",
	"uint24":  "uint32",
	"int24":   "int32",
	"uint32":  "uint32",
	"uint40":  "uint64",
	"uint64":  "uint64",
	"uint128": "*big.Int",
	"uint256": "*big.Int",
	"int128":  "*big.Int",
	"int256":  "*big.Int",
}

// =============================================================================
// Main Entry Point
// =============================================================================

func main() {
	fmt.Println("CREC SDK Extension Code Generator")
	fmt.Println("==================================")

	// Collect all functions and events from configured contracts
	var allFuncs []FunctionDef
	var allEvents []EventDef

	for _, contract := range contracts {
		abiJSON := loadABI(contract.ABIFile, contract.Name)
		helperMethod := deriveHelperMethod(contract.Name)

		// Parse functions with per-contract overrides
		funcs := parseABI(abiJSON, contract.Name, helperMethod, contract.SkipMethods, contract.FuncNameOverrides)
		allFuncs = append(allFuncs, funcs...)

		// Parse events
		events := parseEvents(abiJSON)
		allEvents = append(allEvents, events...)
	}

	// Add extra events from config
	allEvents = append(allEvents, extraEvents...)

	// Deduplicate and sort
	allEvents = deduplicateEvents(allEvents)

	sort.Slice(allFuncs, func(i, j int) bool {
		return allFuncs[i].GoName < allFuncs[j].GoName
	})
	sort.Slice(allEvents, func(i, j int) bool {
		return allEvents[i].Name < allEvents[j].Name
	})

	// Generate files
	generateABIFile(contracts)
	generateExtensionFile(contracts)
	generateOperationHelpersFile(contracts)
	generateOperationsFile(allFuncs)
	generateEventsFile(allEvents)
	generateDecodersFile(allEvents)
	generateWatcherValuesFiles(contracts)

	fmt.Println("==================================")
	fmt.Println("Code generation complete!")
}

// =============================================================================
// Name Derivation Helpers
// =============================================================================

// deriveHelperMethod generates helper method name from contract name
// e.g., "MyContract" -> "prepareMyContractOp"
func deriveHelperMethod(contractName string) string {
	return "prepare" + contractName + "Op"
}

// deriveAddressField generates address field name from contract name
// e.g., "MyContract" -> "myContractAddress"
func deriveAddressField(contractName string) string {
	return toLowerFirst(contractName) + "Address"
}

// deriveOptionsField generates options field name from contract name
// e.g., "MyContract" -> "MyContractAddress"
func deriveOptionsField(contractName string) string {
	return contractName + "Address"
}

// deriveABIVarName generates ABI variable name from contract name
// e.g., "MyContract" -> "myContractABI"
func deriveABIVarName(contractName string) string {
	return toLowerFirst(contractName) + "ABI"
}

// deriveABIJSONVarName generates ABI JSON variable name from contract name
// e.g., "MyContract" -> "myContractABIJSON"
func deriveABIJSONVarName(contractName string) string {
	return toLowerFirst(contractName) + "ABIJSON"
}

// deriveABIAccessor generates ABI accessor function name from contract name
// e.g., "MyContract" -> "MyContractABI"
func deriveABIAccessor(contractName string) string {
	return contractName + "ABI"
}

// =============================================================================
// ABI Loading and Parsing
// =============================================================================

func loadABI(path, contractName string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to read %s ABI from %s: %v\n", contractName, path, err)
		fmt.Fprintf(os.Stderr, "       Make sure the ABI file exists in the abi/ directory.\n")
		os.Exit(1)
	}
	fmt.Printf("✓ Loaded %s ABI from %s\n", contractName, path)
	return string(data)
}

func parseABI(abiJSON, contract, helperMethod string, skipMethods map[string]bool, funcOverrides map[string]string) []FunctionDef {
	var entries []ABIEntry
	if err := json.Unmarshal([]byte(abiJSON), &entries); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to parse ABI JSON: %v\n", err)
		os.Exit(1)
	}

	var funcs []FunctionDef
	for _, entry := range entries {
		if entry.Type != "function" {
			continue
		}

		if entry.StateMutability == "view" || entry.StateMutability == "pure" {
			continue
		}

		if skipMethods[entry.Name] {
			fmt.Printf("  ⊘ Skipping %s.%s (in skip list)\n", contract, entry.Name)
			continue
		}

		hasComplexType := false
		for _, input := range entry.Inputs {
			if strings.HasPrefix(input.Type, "tuple") {
				hasComplexType = true
				break
			}
		}

		if hasComplexType {
			fmt.Printf("  ⊘ Skipping %s.%s (has tuple/struct type - add manually to operations.go)\n", contract, entry.Name)
			continue
		}

		// Use per-contract function name override or default
		goName := toGoFuncNameWithOverrides(entry.Name, contract, funcOverrides)

		funcDef := FunctionDef{
			Name:           entry.Name,
			GoName:         goName,
			Contract:       contract,
			HelperMethod:   helperMethod,
			HasComplexType: hasComplexType,
		}

		for _, input := range entry.Inputs {
			goType := mapType(input.Type)
			if goType == "" {
				fmt.Printf("  ⚠ Warning: Unknown type %s for %s.%s parameter %s\n",
					input.Type, contract, entry.Name, input.Name)
				goType = "interface{}"
			}

			typeKey := entry.Name + "_" + input.Name
			needsCast := false
			castToType := ""
			if override, ok := typeOverrides[typeKey]; ok {
				castToType = goType
				goType = override
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

func parseEvents(abiJSON string) []EventDef {
	var entries []ABIEntry
	if err := json.Unmarshal([]byte(abiJSON), &entries); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to parse ABI JSON for events: %v\n", err)
		os.Exit(1)
	}

	var events []EventDef
	for _, entry := range entries {
		if entry.Type != "event" {
			continue
		}

		if skipEvents[entry.Name] {
			continue
		}

		eventDef := EventDef{Name: entry.Name}

		for _, input := range entry.Inputs {
			goType := ""
			isEnum := false
			enumType := ""
			if et, ok := enumTypeMapping[input.InternalType]; ok {
				goType = et
				isEnum = true
				enumType = et
			} else {
				goType = mapEventType(input.Type)
			}

			if goType == "" {
				fmt.Printf("  ⚠ Warning: Unknown event type %s for event %s field %s\n",
					input.Type, entry.Name, input.Name)
				goType = "interface{}"
			}

			goName := toGoFieldName(input.Name)
			jsonTag := toSnakeCase(input.Name)
			localVar := toLowerFirst(goName)
			parseFunc, hasError, valueExpr := getParseInfo(goType, jsonTag, localVar, isEnum, enumType)

			eventDef.Fields = append(eventDef.Fields, FieldDef{
				Name:      input.Name,
				GoName:    goName,
				GoType:    goType,
				JSONTag:   jsonTag,
				ParseFunc: parseFunc,
				HasError:  hasError,
				IsEnum:    isEnum,
				EnumType:  enumType,
				LocalVar:  localVar,
				ValueExpr: valueExpr,
			})
		}

		events = append(events, eventDef)
	}

	return events
}

func deduplicateEvents(events []EventDef) []EventDef {
	seen := make(map[string]bool)
	var result []EventDef
	for _, e := range events {
		if !seen[e.Name] {
			seen[e.Name] = true
			result = append(result, e)
		}
	}
	return result
}

// =============================================================================
// Type Mapping Helpers
// =============================================================================

func mapType(solType string) string {
	if t, ok := typeMapping[solType]; ok {
		return t
	}
	if strings.HasSuffix(solType, "[]") {
		baseType := strings.TrimSuffix(solType, "[]")
		if mapped := mapType(baseType); mapped != "" {
			return "[]" + mapped
		}
	}
	if strings.Contains(solType, "[") && strings.HasSuffix(solType, "]") {
		return ""
	}
	return ""
}

func mapEventType(solType string) string {
	if t, ok := eventTypeMapping[solType]; ok {
		return t
	}
	if strings.HasSuffix(solType, "[]") {
		baseType := strings.TrimSuffix(solType, "[]")
		if mapped := mapEventType(baseType); mapped != "" {
			return "[]" + mapped
		}
	}
	return ""
}

// =============================================================================
// Name Conversion Helpers
// =============================================================================

func toGoFuncNameWithOverrides(name, contract string, overrides map[string]string) string {
	// Check per-contract overrides first (just method name)
	if override, ok := overrides[name]; ok {
		return override
	}
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
	return name
}

func toGoFieldName(name string) string {
	if len(name) == 0 {
		return "Field"
	}
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func toLowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)

	// Count leading uppercase characters
	upperCount := 0
	for i := 0; i < len(runes) && unicode.IsUpper(runes[i]); i++ {
		upperCount++
	}

	if upperCount == 0 {
		return s
	}

	// Determine how many to lowercase:
	// - If only 1 uppercase or entire string is uppercase: lowercase all leading uppers
	// - Otherwise: keep the last uppercase (it starts a camelCase word)
	lowercaseCount := upperCount
	if upperCount > 1 && upperCount < len(runes) {
		lowercaseCount = upperCount - 1
	}

	for i := 0; i < lowercaseCount; i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}

func getParseInfo(goType, jsonTag, localVar string, isEnum bool, enumType string) (parseFunc string, hasError bool, valueExpr string) {
	paramExpr := fmt.Sprintf(`params["%s"]`, jsonTag)

	if isEnum {
		return fmt.Sprintf("parsing.ScientificNotationToUint8(%s)", paramExpr),
			true,
			fmt.Sprintf("%s(%s)", enumType, localVar)
	}

	switch goType {
	case "common.Address":
		return "", false, fmt.Sprintf("common.HexToAddress(%s)", paramExpr)
	case "common.Hash":
		return "", false, fmt.Sprintf("common.HexToHash(%s)", paramExpr)
	case "*big.Int":
		return fmt.Sprintf("parsing.ScientificNotationToBigInt(%s)", paramExpr), true, localVar
	case "uint64":
		return fmt.Sprintf("parsing.ScientificNotationToUint64(%s)", paramExpr), true, localVar
	case "uint32":
		return fmt.Sprintf("parsing.ScientificNotationToUint32(%s)", paramExpr), true, localVar
	case "uint16":
		return fmt.Sprintf("parsing.ScientificNotationToUint16(%s)", paramExpr), true, localVar
	case "uint8":
		return fmt.Sprintf("parsing.ScientificNotationToUint8(%s)", paramExpr), true, localVar
	case "bool":
		return fmt.Sprintf("strconv.ParseBool(%s)", paramExpr), true, localVar
	case "[]byte":
		return "", false, fmt.Sprintf("[]byte(%s)", paramExpr)
	default:
		return "", false, paramExpr
	}
}

// =============================================================================
// Code Generation: ABI File
// =============================================================================

type abiGenData struct {
	Contracts []struct {
		Name        string
		ABIFile     string
		ABIJSONVar  string
		ABIVar      string
		ABIAccessor string
	}
}

func generateABIFile(contractConfigs []ContractConfig) {
	data := abiGenData{}
	for _, c := range contractConfigs {
		data.Contracts = append(data.Contracts, struct {
			Name        string
			ABIFile     string
			ABIJSONVar  string
			ABIVar      string
			ABIAccessor string
		}{
			Name:        c.Name,
			ABIFile:     c.ABIFile,
			ABIJSONVar:  deriveABIJSONVarName(c.Name),
			ABIVar:      deriveABIVarName(c.Name),
			ABIAccessor: deriveABIAccessor(c.Name),
		})
	}

	tmpl := `// Code generated by gen/main.go. DO NOT EDIT.

package ` + packageName + `

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// =============================================================================
// Embedded ABI JSON
// =============================================================================
{{range .Contracts}}
//go:embed {{.ABIFile}}
var {{.ABIJSONVar}} string
{{end}}

// =============================================================================
// Parsed ABIs
// =============================================================================

var (
{{- range .Contracts}}
	{{.ABIVar}} *abi.ABI
{{- end}}
)

func init() {
	var err error
{{range .Contracts}}
	{{.ABIVar}}, err = parseABI({{.ABIJSONVar}})
	if err != nil {
		panic(fmt.Sprintf("failed to parse {{.Name}} ABI: %v", err))
	}
{{end}}
}

func parseABI(jsonABI string) (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(jsonABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// =============================================================================
// ABI Accessors
// =============================================================================
{{range .Contracts}}
// {{.ABIAccessor}} returns the parsed {{.Name}} ABI.
func {{.ABIAccessor}}() *abi.ABI {
	return {{.ABIVar}}
}
{{end}}
`

	t := template.Must(template.New("abi").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute ABI template: %v\n", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to format ABI code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	if err := os.WriteFile("abi_gen.go", formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write ABI file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated abi_gen.go with %d ABIs\n", len(contractConfigs))
}

// =============================================================================
// Code Generation: Extension File
// =============================================================================

type extensionGenData struct {
	PackageName string
	Contracts   []struct {
		Name         string
		OptionsField string
		AddressField string
	}
}

func generateExtensionFile(contractConfigs []ContractConfig) {
	data := extensionGenData{PackageName: packageName}
	for _, c := range contractConfigs {
		data.Contracts = append(data.Contracts, struct {
			Name         string
			OptionsField string
			AddressField string
		}{
			Name:         c.Name,
			OptionsField: deriveOptionsField(c.Name),
			AddressField: deriveAddressField(c.Name),
		})
	}

	tmpl := `// Code generated by gen/main.go. DO NOT EDIT.

package ` + packageName + `

import (
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
)

// Options defines the configuration for creating a new extension.
type Options struct {
	// Logger is an optional logger instance. If nil, slog.Default() is used.
	Logger *slog.Logger

	// AccountAddress is the address of the account performing operations.
	AccountAddress string
{{range .Contracts}}
	// {{.OptionsField}} is the address of the {{.Name}} contract.
	{{.OptionsField}} string
{{end}}
}

// validate checks that all required options are valid.
func (o *Options) validate() error {
	if !common.IsHexAddress(o.AccountAddress) {
		return fmt.Errorf("invalid AccountAddress: %q", o.AccountAddress)
	}
{{range .Contracts}}
	if !common.IsHexAddress(o.{{.OptionsField}}) {
		return fmt.Errorf("invalid {{.OptionsField}}: %q", o.{{.OptionsField}})
	}
{{end}}
	return nil
}

// Extension provides methods for preparing operations.
type Extension struct {
	logger         *slog.Logger
	accountAddress common.Address
{{- range .Contracts}}
	{{.AddressField}} common.Address
{{- end}}
}

// New creates a new extension with the provided options.
func New(opts *Options) (*Extension, error) {
	if opts == nil {
		return nil, fmt.Errorf("options is required")
	}

	if err := opts.validate(); err != nil {
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	logger := opts.Logger
	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("creating CREC SDK extension")

	return &Extension{
		logger:         logger,
		accountAddress: common.HexToAddress(opts.AccountAddress),
{{- range .Contracts}}
		{{.AddressField}}: common.HexToAddress(opts.{{.OptionsField}}),
{{- end}}
	}, nil
}
`

	t := template.Must(template.New("extension").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute extension template: %v\n", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to format extension code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	if err := os.WriteFile("extension_gen.go", formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write extension file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated extension_gen.go with %d contracts\n", len(contractConfigs))
}

// =============================================================================
// Code Generation: Operation Helpers File
// =============================================================================

type helpersGenData struct {
	Contracts []struct {
		Name         string
		HelperMethod string
		AddressField string
		ABIAccessor  string
	}
}

func generateOperationHelpersFile(contractConfigs []ContractConfig) {
	data := helpersGenData{}
	for _, c := range contractConfigs {
		data.Contracts = append(data.Contracts, struct {
			Name         string
			HelperMethod string
			AddressField string
			ABIAccessor  string
		}{
			Name:         c.Name,
			HelperMethod: deriveHelperMethod(c.Name),
			AddressField: deriveAddressField(c.Name),
			ABIAccessor:  deriveABIAccessor(c.Name),
		})
	}

	tmpl := `// Code generated by gen/main.go. DO NOT EDIT.

package ` + packageName + `

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	transactTypes "github.com/smartcontractkit/crec-sdk/transact/types"
)

// Ensure imports are used
var (
	_ = rand.Read
	_ = fmt.Sprintf
	_ = big.NewInt
	_ *abi.ABI
	_ = common.Address{}
	_ *transactTypes.Operation
)

// =============================================================================
// Generic Operation Builders (escape hatches for power users)
// =============================================================================
{{range .Contracts}}
// Prepare{{.Name}}Operation prepares a generic operation on {{.Name}}.
// Use this for methods not covered by the type-safe Prepare* functions.
func (e *Extension) Prepare{{.Name}}Operation(method string, args ...interface{}) (*transactTypes.Operation, error) {
	return e.prepareOperation({{.ABIAccessor}}(), e.{{.AddressField}}, method, args...)
}
{{end}}

// =============================================================================
// Internal Helper Methods (called by generated Prepare* functions)
// =============================================================================
{{range .Contracts}}
func (e *Extension) {{.HelperMethod}}(method string, args ...interface{}) (*transactTypes.Operation, error) {
	return e.prepareOperation({{.ABIAccessor}}(), e.{{.AddressField}}, method, args...)
}
{{end}}

// =============================================================================
// Core Operation Builder
// =============================================================================

func (e *Extension) prepareOperation(
	contractABI *abi.ABI,
	target common.Address,
	method string,
	args ...interface{},
) (*transactTypes.Operation, error) {
	calldata, err := contractABI.Pack(method, args...)
	if err != nil {
		e.logger.Error("failed to pack calldata", "method", method, "error", err)
		return nil, fmt.Errorf("pack %s: %w", method, err)
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
				To:    target,
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}, nil
}

func generateOperationID() (*big.Int, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(b), nil
}
`

	t := template.Must(template.New("helpers").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute helpers template: %v\n", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to format helpers code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	if err := os.WriteFile("operations_helpers_gen.go", formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write helpers file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated operations_helpers_gen.go with %d contracts\n", len(contractConfigs))
}

// =============================================================================
// Code Generation: Operations
// =============================================================================

func generateOperationsFile(funcs []FunctionDef) {
	tmpl := `// Code generated by gen/main.go. DO NOT EDIT.

package ` + packageName + `

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

	t := template.Must(template.New("ops").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, funcs); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute operations template: %v\n", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to format operations code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	if err := os.WriteFile("operations_gen.go", formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write operations file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated operations_gen.go with %d operations\n", len(funcs))
}

// =============================================================================
// Code Generation: Events
// =============================================================================

func generateEventsFile(events []EventDef) {
	tmpl := `// Code generated by gen/main.go. DO NOT EDIT.

package ` + packageName + `

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Ensure imports are used
var (
	_ = big.NewInt
	_ = common.Address{}
	_ = common.Hash{}
)

// EventName represents an event name from the contracts.
type EventName string

const (
{{- range .}}
	Event{{.Name}} EventName = "{{.Name}}"
{{- end}}
	EventUnknown EventName = "Unknown"
)

var allEvents = map[string]EventName{
{{- range .}}
	string(Event{{.Name}}): Event{{.Name}},
{{- end}}
}

func parseEventName(value string) (EventName, bool) {
	ev, ok := allEvents[value]
	return ev, ok
}

func (ev EventName) String() string {
	return string(ev)
}

{{range .}}
// {{.Name}} event.
type {{.Name}} struct {
{{- range .Fields}}
	{{.GoName}} {{.GoType}} ` + "`" + `json:"{{.JSONTag}}"` + "`" + `
{{- end}}
}
{{end}}
`

	t := template.Must(template.New("events").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, events); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute events template: %v\n", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to format events code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	if err := os.WriteFile("events_gen.go", formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write events file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated events_gen.go with %d events\n", len(events))
}

// =============================================================================
// Code Generation: Decoders
// =============================================================================

func generateDecodersFile(events []EventDef) {
	tmpl := `// Code generated by gen/main.go. DO NOT EDIT.

package ` + packageName + `

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/crec-sdk/parsing"
)

// Ensure imports are used
var (
	_ = fmt.Sprintf
	_ = strconv.ParseBool
	_ = common.Address{}
	_ = parsing.ScientificNotationToBigInt
)

// eventDecoder is a function that decodes event parameters into a concrete event type.
type eventDecoder func(params map[string]string, txHash string) (ConcreteEvent, error)

// eventDecoders maps event names to their decoder functions.
var eventDecoders = map[EventName]eventDecoder{
{{- range .}}
	Event{{.Name}}: decode{{.Name}},
{{- end}}
}

{{range .}}
func decode{{.Name}}(params map[string]string, _ string) (ConcreteEvent, error) {
{{- range .Fields}}
{{- if .HasError}}
	{{.LocalVar}}, err := {{.ParseFunc}}
	if err != nil {
		return nil, fmt.Errorf("parse {{.JSONTag}} %q: %w", params["{{.JSONTag}}"], err)
	}
{{- end}}
{{- end}}
	return &{{.Name}}{
{{- range .Fields}}
		{{.GoName}}: {{.ValueExpr}},
{{- end}}
	}, nil
}
{{end}}
`

	t := template.Must(template.New("decoders").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, events); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute decoders template: %v\n", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to format decoders code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	if err := os.WriteFile("decode_gen.go", formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write decoders file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated decode_gen.go with %d decoders\n", len(events))
}

// =============================================================================
// Code Generation: Watcher Values Files
// =============================================================================

func generateWatcherValuesFiles(contractConfigs []ContractConfig) {
	// Ensure watcher/values directory exists
	if err := os.MkdirAll("watcher/values", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to create watcher/values directory: %v\n", err)
		os.Exit(1)
	}

	for _, contract := range contractConfigs {
		generateWatcherValuesFile(contract)
	}
}

// cleanEventInput is a minimal struct for clean JSON output
type cleanEventInput struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Indexed      bool   `json:"indexed"`
	InternalType string `json:"internalType"`
}

// cleanEvent is a minimal struct for clean JSON output
type cleanEvent struct {
	Type      string            `json:"type"`
	Name      string            `json:"name"`
	Inputs    []cleanEventInput `json:"inputs"`
	Anonymous bool              `json:"anonymous"`
}

func generateWatcherValuesFile(contract ContractConfig) {
	// Load and parse ABI to extract events
	abiData, err := os.ReadFile(contract.ABIFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to read ABI for watcher values: %v\n", err)
		os.Exit(1)
	}

	var entries []ABIEntry
	if err := json.Unmarshal(abiData, &entries); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to parse ABI for watcher values: %v\n", err)
		os.Exit(1)
	}

	// Extract only events and convert to clean format
	var events []cleanEvent
	for _, entry := range entries {
		if entry.Type == "event" {
			ce := cleanEvent{
				Type:      entry.Type,
				Name:      entry.Name,
				Anonymous: entry.Anonymous,
			}
			for _, input := range entry.Inputs {
				ce.Inputs = append(ce.Inputs, cleanEventInput{
					Name:         input.Name,
					Type:         input.Type,
					Indexed:      input.Indexed,
					InternalType: input.InternalType,
				})
			}
			events = append(events, ce)
		}
	}

	// Marshal events back to JSON
	eventsJSON, err := json.Marshal(events)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to marshal events JSON: %v\n", err)
		os.Exit(1)
	}

	// Generate lowercase contract name for workflow name
	contractNameLower := strings.ToLower(contract.Name)

	data := struct {
		ContractName      string
		ContractNameLower string
		ABIFile           string
		EventsJSON        string
	}{
		ContractName:      contract.Name,
		ContractNameLower: contractNameLower,
		ABIFile:           contract.ABIFile,
		EventsJSON:        string(eventsJSON),
	}

	tmpl := `# ==========================================================================
# WATCHER VALUES: {{.ContractName}}
# ==========================================================================
# Auto-generated from: {{.ABIFile}}
# Customize and use with: task watcher:config VALUES_FILE=values/values.{{.ContractName}}.yaml

# ==========================================================================
# WORKFLOW.YAML VALUES
# ==========================================================================
# NOTE: workflowName has a maximum of 32 characters
workflowName: "watcher-{{.ContractNameLower}}"

# ==========================================================================
# CONFIG.YAML VALUES
# ==========================================================================

# Network and chain configuration
network: "evm"
chainId: "11155111"
chainSelector: 16015286601757825753

# CREC service configuration
courierUrl: "https://crec.chainlink.com"
service: "myservice"
watcherID: "00000000-0000-0000-0000-000000000000"

# Contract configuration (from ABI)
contractName: "{{.ContractName}}"
contractAddress: "0x0000000000000000000000000000000000000000"

# Event to watch (pick one from the events in contractABI below)
eventName: ""

# Contract ABI (auto-extracted events from {{.ABIFile}})
contractABI: '{{.EventsJSON}}'
`

	t := template.Must(template.New("watcherValues").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute watcher values template: %v\n", err)
		os.Exit(1)
	}

	outputPath := fmt.Sprintf("watcher/values/values.%s.yaml", contract.Name)
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write watcher values file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated %s with %d events\n", outputPath, len(events))
}
