package bundle

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"

	"github.com/andybalholm/brotli"
	crecbundle "github.com/smartcontractkit/crec-sdk/extension/bundle"
)

//go:embed binary.wasm.br.b64
var wasmBinaryBrB64 string

var wasmBinary = decodeCompressedBinary(wasmBinaryBrB64)

//go:embed DTARequestManagementU.abi.json
var dtaRequestManagementABI string

//go:embed DTARequestSettlementU.abi.json
var dtaRequestSettlementABI string

// Get returns the DTA extension watcher bundle.
func Get() *crecbundle.Bundle {
	return &crecbundle.Bundle{
		Service:    "dta.v1",
		WasmBinary: wasmBinary,
		Contracts:  contracts,
		Events:         events,
	}
}

var contracts = []crecbundle.Contract{
	{Name: "DTARequestManagement", ABI: dtaRequestManagementABI},
	{Name: "DTARequestSettlement", ABI: dtaRequestSettlementABI},
}

// DataSchema for events enriched with distributor_request + fund_token_data.
var onChainEnrichmentDataSchema = json.RawMessage(`{
	"type": "object",
	"properties": {
		"on_chain": {
			"type": "array",
			"description": "On-chain reference data fetched by the handler",
			"items": {
				"type": "object",
				"properties": {
					"source": {
						"type": "object",
						"properties": {
							"contract_address": {"type": "string"},
							"contract_function_signature": {"type": "string"},
							"call_data": {"type": "string"},
							"block": {"type": "string"}
						}
					},
					"data": {"type": "object"}
				}
			}
		}
	}
}`)

// DataSchema for settlement events: on-chain enrichment + off-chain currency + payment request (Opened only).
var settlementDataSchema = json.RawMessage(`{
	"type": "object",
	"properties": {
		"on_chain": {
			"type": "array",
			"description": "On-chain reference data (distributor_request, fund_token_data)",
			"items": {
				"type": "object",
				"properties": {
					"source": {"type": "object"},
					"data": {"type": "object"}
				}
			}
		},
		"off_chain": {
			"type": "array",
			"description": "Off-chain reference data (currency code)",
			"items": {
				"type": "object",
				"properties": {
					"source": {"type": "object"},
					"data": {"type": "object"}
				}
			}
		},
		"requests": {
			"type": "array",
			"description": "Generated requests (payment_request for DTASettlementOpened)",
			"items": {
				"type": "object",
				"properties": {
					"type": {"type": "string"},
					"value": {"type": "object"}
				}
			}
		}
	}
}`)

var events = []crecbundle.Event{
	// --- DTARequestManagement events ---
	{Name: "DistributorRegistered", TriggerContract: "DTARequestManagement", Description: "New distributor registered", ParamsSchema: ParamsSchemas["DistributorRegistered"]},
	{Name: "DistributorRequestCanceled", TriggerContract: "DTARequestManagement", Description: "Distributor request canceled", ParamsSchema: ParamsSchemas["DistributorRequestCanceled"], DataSchema: onChainEnrichmentDataSchema},
	{Name: "DistributorRequestProcessed", TriggerContract: "DTARequestManagement", Description: "Distributor request completed", ParamsSchema: ParamsSchemas["DistributorRequestProcessed"], DataSchema: onChainEnrichmentDataSchema},
	{Name: "DistributorRequestProcessing", TriggerContract: "DTARequestManagement", Description: "Distributor request being processed", ParamsSchema: ParamsSchemas["DistributorRequestProcessing"], DataSchema: onChainEnrichmentDataSchema},
	{Name: "FundAdminRegistered", TriggerContract: "DTARequestManagement", Description: "New fund admin registered", ParamsSchema: ParamsSchemas["FundAdminRegistered"]},
	{Name: "FundTokenAllowlistUpdated", TriggerContract: "DTARequestManagement", Description: "Fund token allowlist updated", ParamsSchema: ParamsSchemas["FundTokenAllowlistUpdated"], DataSchema: onChainEnrichmentDataSchema},
	{Name: "FundTokenRegistered", TriggerContract: "DTARequestManagement", Description: "New fund token created", ParamsSchema: ParamsSchemas["FundTokenRegistered"], DataSchema: onChainEnrichmentDataSchema},
	{Name: "RedemptionRequested", TriggerContract: "DTARequestManagement", Description: "Investor redemption request", ParamsSchema: ParamsSchemas["RedemptionRequested"], DataSchema: onChainEnrichmentDataSchema},
	{Name: "SubscriptionRequested", TriggerContract: "DTARequestManagement", Description: "Investor subscription request", ParamsSchema: ParamsSchemas["SubscriptionRequested"], DataSchema: onChainEnrichmentDataSchema},
	// --- DTARequestSettlement events ---
	{Name: "DTASettlementOpened", TriggerContract: "DTARequestSettlement", Description: "DTA settlement initiated", ParamsSchema: ParamsSchemas["DTASettlementOpened"], DataSchema: settlementDataSchema},
	{Name: "DTASettlementClosed", TriggerContract: "DTARequestSettlement", Description: "DTA settlement completed or failed", ParamsSchema: ParamsSchemas["DTASettlementClosed"], DataSchema: settlementDataSchema},
}

func decodeCompressedBinary(encoded string) []byte {
	data := strings.TrimSpace(encoded)
	if data == "" {
		return nil
	}
	compressed, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic("crec bundle: invalid base64 wasm: " + err.Error())
	}
	raw, err := io.ReadAll(brotli.NewReader(bytes.NewReader(compressed)))
	if err != nil {
		panic("crec bundle: decompress wasm: " + err.Error())
	}
	return raw
}
