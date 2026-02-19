package bundle

import (
	_ "embed"
	"encoding/json"

	crecbundle "github.com/smartcontractkit/crec-sdk/extension/bundle"
)

//go:embed binary.wasm
var wasmBinary []byte

//go:embed config.tmpl
var configTemplate []byte

//go:embed DTARequestManagementU.abi.json
var dtaRequestManagementUABI string

//go:embed DTARequestSettlementU.abi.json
var dtaRequestSettlementUABI string

// Get returns the DTA v2 extension watcher bundle.
func Get() *crecbundle.Bundle {
	return &crecbundle.Bundle{
		Service:        "dta.v2",
		WasmBinary:     wasmBinary,
		ConfigTemplate: configTemplate,
		Contracts:      contracts,
		Events:         events,
	}
}

var contracts = []crecbundle.Contract{
	{Name: "DTARequestManagement", ABI: dtaRequestManagementUABI},
	{Name: "DTARequestSettlement", ABI: dtaRequestSettlementUABI},
	// TODO: Add MockNAVAggregator contract when ABI is available.
	// Used by the handler for NAV price feed enrichment on settlement events.
	// {Name: "MockNAVAggregator", ABI: mockNAVAggregatorABI},
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
	{Name: "RedemptionRequested", TriggerContract: "DTARequestManagement", Description: "Redemption request", ParamsSchema: ParamsSchemas["RedemptionRequested"], DataSchema: onChainEnrichmentDataSchema},
	{Name: "SubscriptionRequested", TriggerContract: "DTARequestManagement", Description: "Subscription request", ParamsSchema: ParamsSchemas["SubscriptionRequested"], DataSchema: onChainEnrichmentDataSchema},
	// --- DTARequestSettlement events ---
	{Name: "DTASettlementOpened", TriggerContract: "DTARequestSettlement", Description: "DTA settlement initiated", ParamsSchema: ParamsSchemas["DTASettlementOpened"], DataSchema: settlementDataSchema},
	{Name: "DTASettlementClosed", TriggerContract: "DTARequestSettlement", Description: "DTA settlement completed or failed", ParamsSchema: ParamsSchemas["DTASettlementClosed"], DataSchema: settlementDataSchema},
}
