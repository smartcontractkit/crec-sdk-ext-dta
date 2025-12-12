//go:build ignore

// Generator for DTA v1 extension Prepare* functions and event types.
// Run with: go run ./v1/gen/main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

// ABI file paths relative to the v1 directory
const (
	managementABIPath    = "abi/DTARequestManagementU.abi.json"
	settlementABIPath    = "abi/DTARequestSettlementU.abi.json"
	navAggregatorABIPath = "abi/MockNAVAggregator.abi.json"
)

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
	Name    string
	GoName  string
	GoType  string
	JSONTag string
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
	"allowDTA_mintType": "TokenMintType",
	"allowDTA_burnType": "TokenBurnType",
}

// Enum internal type mappings (Solidity internalType -> Go type)
// Used to map enum types from ABI internalType to Go enum types.
var enumTypeMapping = map[string]string{
	"enum IDTAMessage.DistributorRequestType":  "DistributorRequestType",
	"enum IDTAMessage.RequestStatus":           "RequestStatus",
	"enum IDTARequestSettlement.TokenMintType": "TokenMintType",
	"enum IDTARequestSettlement.TokenBurnType": "TokenBurnType",
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

// Events to skip (common/inherited events we don't need)
var skipEvents = map[string]bool{}

// Extra events not in any ABI but needed for decoding
var extraEvents = []EventDef{}

// Solidity to Go type mapping for functions
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

// Solidity to Go type mapping for events (uses common.Hash for bytes32)
var eventTypeMapping = map[string]string{
	"address": "common.Address",
	"bool":    "bool",
	"bytes":   "[]byte",
	"bytes4":  "[4]byte",
	"bytes32": "common.Hash",
	"string":  "string",
	"uint8":   "uint8",
	"uint24":  "uint32",
	"int24":   "int32",
	"uint40":  "uint64",
	"uint64":  "uint64",
	"uint256": "*big.Int",
	"int256":  "*big.Int",
}

func main() {
	// ABI files are in v1/abi/
	abiDir := "v1"

	managementABI := loadABI(filepath.Join(abiDir, managementABIPath), "DTARequestManagement")
	settlementABI := loadABI(filepath.Join(abiDir, settlementABIPath), "DTARequestSettlement")
	navAggregatorABI := loadABI(filepath.Join(abiDir, navAggregatorABIPath), "MockNAVAggregator")

	// Generate operations
	var funcs []FunctionDef
	mgmtFuncs := parseABI(managementABI, "DTARequestManagement", "prepareManagementOp", managementSkipMethods)
	funcs = append(funcs, mgmtFuncs...)
	settleFuncs := parseABI(settlementABI, "DTARequestSettlement", "prepareSettlementOp", settlementSkipMethods)
	funcs = append(funcs, settleFuncs...)
	sort.Slice(funcs, func(i, j int) bool {
		return funcs[i].GoName < funcs[j].GoName
	})
	generateOperationsFile(funcs)

	// Generate events
	var events []EventDef
	mgmtEvents := parseEvents(managementABI)
	events = append(events, mgmtEvents...)
	settleEvents := parseEvents(settlementABI)
	events = append(events, settleEvents...)
	navEvents := parseEvents(navAggregatorABI)
	events = append(events, navEvents...)

	// Add extra events not in ABIs
	events = append(events, extraEvents...)

	// Deduplicate events by name
	events = deduplicateEvents(events)

	sort.Slice(events, func(i, j int) bool {
		return events[i].Name < events[j].Name
	})
	generateEventsFile(events)
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

func parseEvents(abiJSON string) []EventDef {
	var entries []ABIEntry
	if err := json.Unmarshal([]byte(abiJSON), &entries); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse ABI for events: %v\n", err)
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

		eventDef := EventDef{
			Name: entry.Name,
		}

		for _, input := range entry.Inputs {
			// First check if this is an enum type via internalType
			goType := ""
			if enumType, ok := enumTypeMapping[input.InternalType]; ok {
				goType = enumType
			} else {
				goType = mapEventType(input.Type)
			}

			if goType == "" {
				fmt.Printf("Warning: Unknown event type %s for event %s field %s\n",
					input.Type, entry.Name, input.Name)
				goType = "interface{}"
			}

			goName := toGoFieldName(input.Name)
			jsonTag := toSnakeCase(input.Name)

			eventDef.Fields = append(eventDef.Fields, FieldDef{
				Name:    input.Name,
				GoName:  goName,
				GoType:  goType,
				JSONTag: jsonTag,
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
		return ""
	}

	return ""
}

func mapEventType(solType string) string {
	if t, ok := eventTypeMapping[solType]; ok {
		return t
	}

	// Handle arrays
	if strings.HasSuffix(solType, "[]") {
		baseType := strings.TrimSuffix(solType, "[]")
		if mapped := mapEventType(baseType); mapped != "" {
			return "[]" + mapped
		}
	}

	return ""
}

func toGoFuncName(name, contract string) string {
	key := name + "_" + contract
	if override, ok := funcNameOverrides[key]; ok {
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

func generateOperationsFile(funcs []FunctionDef) {
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

	t := template.Must(template.New("ops").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, funcs); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute operations template: %v\n", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to format operations code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	outputPath := "v1/dta_operations_gen.go"
	if err := os.WriteFile(outputPath, formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write operations file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s with %d functions\n", outputPath, len(funcs))
}

func generateEventsFile(events []EventDef) {
	const tmpl = `// Code generated by gen/main.go. DO NOT EDIT.

package v1

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

// EventName represents an event name from the DTA contracts.
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
		fmt.Fprintf(os.Stderr, "failed to execute events template: %v\n", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to format events code: %v\n", err)
		fmt.Fprintln(os.Stderr, "Raw code:")
		fmt.Fprintln(os.Stderr, buf.String())
		os.Exit(1)
	}

	outputPath := "v1/events_gen.go"
	if err := os.WriteFile(outputPath, formatted, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write events file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s with %d events\n", outputPath, len(events))
}
