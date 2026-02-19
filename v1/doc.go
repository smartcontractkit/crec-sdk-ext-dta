// Package v1 provides CREC SDK extension for DTA (Digital Transfer Agent) v1.0.
//
// This package is organized into sub-packages:
//
//   - v1/events:     Lightweight event types, decoders, and constants.
//   - v1/contracts:  Embedded ABI JSON and parsed ABI accessors.
//   - v1/operations: Extension client for preparing on-chain operations.
//
// The root v1 package provides [DecodeFromEvent] for SDK consumers to decode
// watcher event payloads into typed Go structs.
//
// # Decoding Events
//
//	decoded, err := v1.DecodeFromEvent(ctx, event)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(decoded.EventName())
//
// # Preparing Operations
//
//	import "github.com/smartcontractkit/crec-sdk-ext-dta/v1/operations"
//
//	ext, err := operations.New(&operations.Options{
//		DTARequestManagementAddress: "0x...",
//		DTARequestSettlementAddress: "0x...",
//		AccountAddress:              "0x...",
//	})
//	op, err := ext.PrepareRequestSubscriptionOperation(fundAdmin, fundTokenId, amount)
package v1
