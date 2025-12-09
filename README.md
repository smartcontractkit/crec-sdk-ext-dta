# CREC SDK Extension: DTA

A Go SDK extension for Digital Transfer Agent (DTA) operations on CREC (Chainlink Runtime Environment Connect).

## Installation

```bash
go get github.com/smartcontractkit/crec-sdk-ext-dta
```

## Overview

This extension provides utilities for preparing DTA operations for fund subscriptions, redemptions, and management on blockchain networks. It works with DTARequestManagement and DTARequestSettlement smart contracts.

## Usage

```go
import (
    v1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
)

// Create the DTA v1 extension
ext, err := v1.New(&v1.Options{
    DTARequestManagementAddress: "0x...",  // DTARequestManagement contract address
    DTARequestSettlementAddress: "0x...",  // DTARequestSettlement contract address
    AccountAddress:              "0x...",  // Your account address
})
if err != nil {
    log.Fatal(err)
}

// Request subscription
op, err := ext.PrepareRequestSubscriptionOperation(fundAdminAddr, fundTokenId, amount)

// Request subscription with token approval
op, err := ext.PrepareRequestSubscriptionWithTokenApprovalOperation(
    fundAdminAddr, fundTokenId, amount, paymentTokenAddress,
)

// Register fund token with full metadata
op, err := ext.PrepareRegisterFundTokenOperation(fundTokenId, tokenData)

// Complete request processing
op, err := ext.PrepareCompleteRequestProcessingOperation(requestId, true, []byte{})
```

## Available Operations

### DTARequestManagement Operations

| Operation                                              | Description                              |
| ------------------------------------------------------ | ---------------------------------------- |
| `PrepareRequestSubscriptionOperation`                  | Request subscription to a fund           |
| `PrepareRequestRedemptionOperation`                    | Request redemption from a fund           |
| `PrepareRequestSubscriptionWithTokenApprovalOperation` | Request subscription with token approval |
| `PrepareProcessDistributorRequestOperation`            | Process a pending distributor request    |
| `PrepareCancelDistributorRequestOperation`             | Cancel a distributor request             |
| `PrepareRegisterDistributorOperation`                  | Register a new distributor               |
| `PrepareRegisterFundAdminOperation`                    | Register a new fund admin                |
| `PrepareRegisterFundTokenOperation`                    | Register a new fund token                |
| `PrepareAllowDistributorForTokenOperation`             | Allow distributor for a token            |
| `PrepareDisallowDistributorForTokenOperation`          | Disallow distributor for a token         |
| `PrepareForceAllowDistributorForTokenOperation`        | Force allow distributor (admin)          |
| `PrepareVerifyDistributorWalletOperation`              | Verify distributor wallet ownership      |
| `PrepareEnableFundTokenOperation`                      | Enable a fund token                      |
| `PrepareDisableFundTokenOperation`                     | Disable a fund token                     |
| `PrepareSetManagementCCIPGasLimitOperation`            | Set CCIP gas limit                       |
| `PrepareWithdrawManagementTokensOperation`             | Withdraw tokens                          |

### DTARequestSettlement Operations

| Operation                                               | Description                             |
| ------------------------------------------------------- | --------------------------------------- |
| `PrepareAllowDTAOperation`                              | Allow a DTA address for a fund token    |
| `PrepareDisallowDTAOperation`                           | Disallow a DTA address for a fund token |
| `PrepareCompleteRequestProcessingOperation`             | Complete request processing             |
| `PrepareWithdrawSettlementTokensOperation`              | Withdraw tokens from settlement         |
| `PrepareSetSettlementCCIPGasLimitOperation`             | Set CCIP gas limit                      |
| `PrepareTransferDTARequestSettlementOwnershipOperation` | Transfer contract ownership             |
| `PrepareRenounceDTARequestSettlementOwnershipOperation` | Renounce contract ownership             |

### Generic Operations (Power Users)

For methods not covered by type-safe functions:

```go
// Call any DTARequestManagement method
op, err := ext.PrepareManagementOperation("methodName", arg1, arg2, ...)

// Call any DTARequestSettlement method
op, err := ext.PrepareSettlementOperation("methodName", arg1, arg2, ...)
```

## Token Burn Types

| Constant                | Value | Description      |
| ----------------------- | ----- | ---------------- |
| `TokenBurnTypeNone`     | 0     | No burn behavior |
| `TokenBurnTypeBurn`     | 1     | Burn tokens      |
| `TokenBurnTypeTransfer` | 2     | Transfer tokens  |

## Development

### Code Generation

The type-safe `Prepare*` functions are generated from the contract ABIs. To regenerate after ABI changes:

1. Update ABI files in `v1/abi/`:
2. Run the generator:
   ```bash
   make generate
   # or
   go run ./v1/gen/main.go
   ```
3. Run tests:
   ```bash
   go test ./...
   ```

### Project Structure

```
├── v1/                    # DTA v1 SDK
│   ├── abi/               # Embedded ABI JSON files
│   ├── gen/               # Code generator
│   │   └── main.go
│   ├── abi.go             # ABI embedding and parsing
│   ├── dta.go             # Main extension and helpers
│   ├── dta_operations_gen.go  # Generated operations (DO NOT EDIT)
│   ├── decode.go          # Event decoding
│   └── events.go          # Event types and constants
├── types/                 # Shared types for event decoding
└── parsing/               # Numeric parsing utilities
```

## License

See [LICENSE](LICENSE) for details.
