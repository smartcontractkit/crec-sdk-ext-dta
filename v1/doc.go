// Package v1 provides CREC SDK extension for DTA (Digital Transfer Agent) v1.0 operations.
//
// This package works with DTARequestManagement and DTARequestSettlement contracts
// to facilitate tokenized fund subscription and redemption workflows.
//
// # Usage
//
//	ext, err := v1.New(&v1.Options{
//		DTARequestManagementAddress: "0x...",
//		DTARequestSettlementAddress: "0x...",
//		AccountAddress:              "0x...",
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Prepare a subscription operation
//	op, err := ext.PrepareRequestSubscriptionOperation(fundAdmin, fundTokenId, amount)
//
// # Operations
//
// The extension provides type-safe Prepare* functions for all contract operations.
// For methods not covered by the generated functions, use the generic helpers:
//
//	op, err := ext.PrepareDTARequestManagementOperation("methodName", arg1, arg2)
//	op, err := ext.PrepareDTARequestSettlementOperation("methodName", arg1, arg2)
//
// # Events
//
// Event types are generated from the contract ABIs for use in decoding:
//
//	event, err := v1.DecodeEvent(eventParams)
//	switch e := event.Data.(type) {
//	case *v1.DistributorRequestCreated:
//		// Handle subscription/redemption request
//	case *v1.DistributorRequestCompleted:
//		// Handle completion
//	}
package v1

