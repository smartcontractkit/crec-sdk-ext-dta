// Package v2 provides CREC SDK extension for DTA (Digital Transfer Agent) v2.0.
//
// This package is organized into sub-packages:
//
//   - v2/events:     Lightweight event types, decoders, and constants.
//   - v2/operations: Extension client for preparing on-chain operations.
//
// The root v2 package provides [DecodeFromEvent] for SDK consumers to decode
// watcher event payloads into typed Go structs with enrichment data.
//
// # Decoding Events
//
//	decoded, err := v2.DecodeFromEvent(ctx, event)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(decoded.EventName())
//
// # Preparing Operations
//
//	import "github.com/smartcontractkit/crec-sdk-ext-dta/v2/operations"
//
//	ext, err := operations.New(&operations.Options{
//		DTARequestManagementAddress: "0x...",
//		DTARequestSettlementAddress: "0x...",
//		AccountAddress:              "0x...",
//	})
//	op, err := ext.PrepareRequestSubscriptionOperation(fundAdmin, fundTokenId, amount, referenceID)
package v2
