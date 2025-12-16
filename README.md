# CREC SDK Extension: DTA

A Go SDK extension for Digital Transfer Agent (DTA) operations on CREC (Chainlink Runtime Environment Connect).

## Installation

```bash
go get github.com/smartcontractkit/crec-sdk-ext-dta
```

## Versioning

This SDK uses **contract versioning** (e.g., `v1/`) to match deployed smart contract ABI versions. This is separate from the Go module's semantic versioning.

| Directory | Contract Version | Description                  |
| --------- | ---------------- | ---------------------------- |
| `v1/`     | DTA contracts v1 | Current production contracts |

When new contract versions are deployed with breaking ABI changes, a new directory (e.g., `v2/`) will be added:

```go
import dtav1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
```

## Overview

This extension provides utilities for preparing DTA operations for fund subscriptions, redemptions, and management on blockchain networks. It works with DTARequestManagement and DTARequestSettlement smart contracts.

## Usage

```go
import (
    dtav1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
)

// Create the DTA v1 extension
ext, err := dtav1.New(&dtav1.Options{
    DTARequestManagementAddress: "0x...",
    DTARequestSettlementAddress: "0x...",
    AccountAddress:              "0x...",
})
if err != nil {
    log.Fatal(err)
}

// Request subscription
op, err := ext.PrepareRequestSubscriptionOperation(fundAdminAddr, fundTokenId, amount)

// Request subscription with token approval (multi-transaction)
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

### Generic Operations

For contract methods not covered by the type-safe functions:

```go
op, err := ext.PrepareDTARequestManagementOperation("methodName", arg1, arg2, ...)
op, err := ext.PrepareDTARequestSettlementOperation("methodName", arg1, arg2, ...)
```

## Types

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

## Project Structure

```
crec-sdk-ext-dta/
├── v1/                              # DTA v1 contract SDK
│   ├── abi/                         # Contract ABIs
│   ├── bindings/                    # Generated Go bindings (abigen)
│   ├── gen/                         # Code generator
│   │   ├── main.go                  # Generator (from template)
│   │   └── config.go                # DTA-specific configuration
│   ├── watcher/                     # CRE watcher workflow
│   │   ├── handler/                 # Workflow handler
│   │   ├── values/                  # Configuration values
│   │   └── *.tmpl                   # Templates
│   ├── doc.go                       # Package documentation
│   ├── types.go                     # Enum and struct types
│   ├── operations.go                # Custom operations (multi-tx, complex types)
│   ├── decode.go                    # Event decoding logic
│   ├── abi_gen.go                   # Generated: ABI embedding
│   ├── extension_gen.go             # Generated: Options, Extension, New()
│   ├── operations_gen.go            # Generated: Type-safe Prepare* functions
│   ├── operations_helpers_gen.go    # Generated: Helper methods
│   ├── events_gen.go                # Generated: Event structs
│   └── decode_gen.go                # Generated: Event decoders
├── mocks/                           # Mock server for local testing
└── Taskfile.yaml                    # Task runner commands
```

## Development

### Prerequisites

Install [Task](https://taskfile.dev/):

```bash
brew install go-task  # macOS
```

### Code Generation

Operations and event types are generated from contract ABIs. To regenerate after ABI changes:

```bash
# Install tools (one-time)
task tools

# Regenerate all code
task generate
```

### Running Tests

```bash
task test
```

### Watcher Workflow

The watcher workflow monitors DTA contract events:

```bash
# Configure watcher
task watcher:config

# Deploy to CRE
task watcher:deploy

# Config + deploy
task watcher:release

# Simulate with a transaction
task watcher:simulate TX_HASH=0x...
```

**Mock server for local simulation:**

```bash
task mock:start
task mock:stop
task mock:logs
```

## License

See [LICENSE](LICENSE) for details.
