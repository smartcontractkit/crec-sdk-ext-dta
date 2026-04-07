# CREC SDK Extension: DTA

A Go SDK extension for Digital Transfer Agent (DTA) operations on CREC (Chainlink Runtime Environment Connect).

## Installation

```bash
go get github.com/smartcontractkit/crec-sdk-ext-dta
```

## Versioning

This SDK uses **contract versioning** (e.g., `v1/`, `v2/`) to match deployed smart contract ABI versions. This is separate from the Go module's semantic versioning.

| Directory | Contract Version | Description                                        |
| --------- | ---------------- | -------------------------------------------------- |
| `v1/`     | DTA contracts v1 | Production contracts                               |
| `v2/`     | DTA contracts v2 | Adds `referenceID`, distributor authorization model |

```go
import dtav1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
import dtav2 "github.com/smartcontractkit/crec-sdk-ext-dta/v2"
```

Both versions share the same package layout (`events/`, `operations/`, `watcher/`, `decode.go`) and the same enum/struct types (`TokenMintType`, `TokenBurnType`, `RequestStatus`, `FundTokenData`, `DistributorRequest`). The differences are strictly ABI-driven.

## Overview

This extension provides utilities for preparing DTA operations for fund subscriptions, redemptions, and management on blockchain networks. It works with DTARequestManagement and DTARequestSettlement smart contracts.

## Usage (v1)

```go
import (
    dtav1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
)

ext, err := dtav1.New(&dtav1.Options{
    DTARequestManagementAddress: "0x...",
    DTARequestSettlementAddress: "0x...",
    AccountAddress:              "0x...",
})

op, err := ext.PrepareRequestSubscriptionOperation(fundAdminAddr, fundTokenId, amount)
```

## Usage (v2)

```go
import (
    dtav2 "github.com/smartcontractkit/crec-sdk-ext-dta/v2"
)

ext, err := dtav2.New(&dtav2.Options{
    DTARequestManagementAddress: "0x...",
    DTARequestSettlementAddress: "0x...",
    AccountAddress:              "0x...",
})

// v2 operations accept a referenceID parameter
op, err := ext.PrepareRequestSubscriptionOperation(fundAdminAddr, fundTokenId, amount, referenceID)
```

## v2 Changes from v1

### New `referenceID` parameter

Subscription and redemption operations now accept a `referenceID [32]byte` parameter for correlating on-chain requests with off-chain records:

| Operation (v2)                          | Signature change                                                   |
| --------------------------------------- | ------------------------------------------------------------------ |
| `PrepareRequestSubscriptionOperation`   | Added `referenceID [32]byte` as last parameter                     |
| `PrepareRequestRedemptionOperation`     | Added `referenceID [32]byte` as last parameter                     |

The corresponding events (`SubscriptionRequested`, `RedemptionRequested`) now include a `ReferenceID common.Hash` field.

### New event: `DistributorAuthorizationUpdated`

Emitted when a distributor's authorization status changes for a fund token:

| Field              | Type             |
| ------------------ | ---------------- |
| `DistributorAddr`  | `common.Address` |
| `FundAdminAddr`    | `common.Address` |
| `FundTokenId`      | `common.Hash`    |
| `Authorized`       | `bool`           |

### New operations

| Operation                                     | Description                                  |
| --------------------------------------------- | -------------------------------------------- |
| `PrepareAuthorizeDistributorForTokenOperation` | Authorize a distributor for a fund token     |
| `PrepareRevokeDistributorForTokenOperation`    | Revoke a distributor's token authorization   |

### Removed operations (v1 only)

| Operation                                      | Replacement                                    |
| ---------------------------------------------- | ---------------------------------------------- |
| `PrepareForceAllowDistributorForTokenOperation` | Use `PrepareAuthorizeDistributorForTokenOperation` |
| `PrepareVerifyDistributorWalletOperation`       | Removed from v2 contracts                      |

## Available Operations (v1)

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

### Generic Operations

For contract methods not covered by the type-safe functions:

```go
op, err := ext.PrepareDTARequestManagementOperation("methodName", arg1, arg2, ...)
op, err := ext.PrepareDTARequestSettlementOperation("methodName", arg1, arg2, ...)
```

## Types

Types are shared across both v1 and v2 (defined in `<version>/events/types.go`).

### Token Mint Types

| Constant                   | Value | Description                                       |
| -------------------------- | ----- | ------------------------------------------------- |
| `TokenMintTypeMint`        | 0     | `mint(address, uint256)` - ERC3643, CMTAT         |
| `TokenMintTypeIssueTokens` | 1     | `issueTokens(address, uint256)` - DSToken (BUIDL) |

### Token Burn Types

| Constant                      | Value | Description                                          |
| ----------------------------- | ----- | ---------------------------------------------------- |
| `TokenBurnTypeBurn`           | 0     | `burn(address, uint256)` - ERC3643                   |
| `TokenBurnTypeBurnFrom`       | 1     | `burnFrom(address, uint256)` - CMTAT                 |
| `TokenBurnTypeBurnWithReason` | 2     | `burn(address, uint256, string)` - DSToken (BUIDL)   |
| `TokenBurnTypeForceBurn`      | 3     | `forceBurn(address, uint256, string)` - CMTAT v2.3.0 |

### Request Status

| Constant                  | Value | Description                |
| ------------------------- | ----- | -------------------------- |
| `RequestStatusNone`       | 0     | Zero value                 |
| `RequestStatusPending`    | 1     | Request is pending         |
| `RequestStatusProcessing` | 2     | Request is being processed |
| `RequestStatusProcessed`  | 3     | Request has been processed |
| `RequestStatusCanceled`   | 4     | Request was canceled       |
| `RequestStatusFailed`     | 5     | Request failed             |

## Decoding Events

Both v1 and v2 provide a `DecodeFromEvent` function that extracts typed events and enrichment data from raw CREC events:

```go
decoded, err := dtav2.DecodeFromEvent(ctx, event)
fmt.Println(decoded.EventName())        // e.g. EventSubscriptionRequested
fmt.Println(decoded.ConcreteEvent)      // typed event struct
fmt.Println(decoded.FundTokenData)      // enrichment: fund token metadata (if available)
fmt.Println(decoded.DistributorRequest) // enrichment: request details (if available)
```

## Project Structure

```
crec-sdk-ext-dta/
├── v1/                              # DTA v1 contract SDK
│   ├── events/                      # Event types and decoders
│   │   ├── types.go                 # Enum and struct types
│   │   ├── events_gen.go            # Generated: Event structs
│   │   └── decode_gen.go            # Generated: Event decoders
│   ├── operations/                  # Operations SDK
│   │   ├── operations.go            # Custom operations (multi-tx, complex types)
│   │   ├── abi_gen.go               # Generated: ABI embedding
│   │   ├── extension_gen.go         # Generated: Options, Extension, New()
│   │   ├── operations_gen.go        # Generated: Type-safe Prepare* functions
│   │   └── operations_helpers_gen.go # Generated: Helper methods
│   ├── watcher/                     # CRE watcher workflow
│   │   ├── handler/                 # Workflow handler (enrichment logic)
│   │   ├── bundle/                  # Bundle definition and schemas
│   │   └── values/                  # Configuration values
│   ├── doc.go                       # Package documentation
│   └── decode.go                    # Event decoding + enrichment
├── v2/                              # DTA v2 contract SDK (same layout as v1)
│   ├── events/
│   ├── operations/
│   ├── watcher/
│   ├── doc.go
│   └── decode.go
├── abi/                             # Contract ABI files
│   ├── v1/
│   └── v2/
├── mocks/                           # Mock server for local testing
└── Taskfile.yaml                    # Task runner commands
```

## Development

### Prerequisites

Install [Task](https://taskfile.dev/):

```bash
brew install go-task  # macOS
```

### Running Tests

```bash
task test
```

### Watcher Workflow

The watcher workflow monitors DTA contract events:

```bash
# Configure watcher (specify VERSION)
task watcher:config VERSION=v2

# Deploy to CRE
task watcher:deploy VERSION=v2

# Config + deploy
task watcher:release VERSION=v2

# Simulate with a transaction
task watcher:simulate VERSION=v2 TX_HASH=0x...
```

**Mock server for local simulation:**

```bash
task mock:start
task mock:stop
task mock:logs
```

## Related

- [crec-sdk](https://github.com/smartcontractkit/crec-sdk) — CREC client for channels, events, watchers, and operations
- [crec-workflow-utils](https://github.com/smartcontractkit/crec-workflow-utils) — Shared utilities for event-listener workflows

## License

[BUSL](LICENSE)
