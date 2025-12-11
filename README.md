# CREC SDK Extension: DTA

A Go SDK extension for Digital Transfer Agent (DTA) operations on CREC (Chainlink Runtime Environment Connect).

## Installation

```bash
go get github.com/smartcontractkit/crec-sdk-ext-dta
```

## Versioning

This SDK uses **contract versioning** (e.g., `v1/`) to match the deployed smart contract ABI versions. This is separate from the Go module's semantic versioning.

| Directory | Contract Version | Description |
|-----------|-----------------|-------------|
| `v1/` | DTA contracts v1 | Current production contracts |

When new contract versions are deployed with breaking ABI changes, a new directory (e.g., `v2/`) will be added. Import the version matching your deployed contracts:

```go
import dtav1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"  // For v1 contracts
// import dtav2 "github.com/smartcontractkit/crec-sdk-ext-dta/v2"  // Future v2 contracts
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

## Project Structure

```
crec-sdk-ext-dta/
├── v1/                           # DTA v1 contract SDK (contract version, not module version)
│   ├── abi/                      # Contract ABIs (canonical source)
│   ├── bindings/                 # Generated Go bindings
│   │   ├── dtarequestmanagement/ # DTARequestManagement contract
│   │   ├── dtarequestsettlement/ # DTARequestSettlement contract
│   │   └── events/               # Generated event types
│   ├── schema/                   # Event validation schemas
│   ├── watcher/                  # CRE watcher workflow for v1 events
│   │   ├── handler/              # Workflow handler logic
│   │   ├── values/               # Configuration values
│   │   ├── config.tmpl           # Config template
│   │   ├── workflow.tmpl         # Workflow template
│   │   └── main.go               # WASM entry point
│   ├── gen/                      # Code generator
│   ├── project.yaml              # CRE project definition
│   ├── abi.go                    # ABI embedding and parsing
│   ├── dta.go                    # Main extension and helpers
│   ├── dta_operations_gen.go     # Generated operations (DO NOT EDIT)
│   ├── decode.go                 # Event decoding
│   └── events.go                 # Event types and constants
├── mocks/                        # Mock server for local testing
├── parsing/                      # Unversioned parsing utilities
├── types/                        # Unversioned shared types
└── Taskfile.yaml                 # Task runner commands
```

## Development

### Prerequisites

Install [Task](https://taskfile.dev/) task runner:

```bash
brew install go-task  # macOS
# or see https://taskfile.dev/installation/
```

### Code Generation

The type-safe `Prepare*` functions are generated from the contract ABIs. To regenerate after ABI changes:

```bash
# Install tools (one-time)
task tools

# Regenerate all code
task generate

# Or run individual generators:
task generate:bindings    # Go bindings from ABIs
task generate:events      # Event types from schema
task generate:operations  # SDK operation functions
```

### Running Tests

```bash
task test
# or
go test ./...
```

### Watcher Workflow

The watcher workflow monitors DTA contract events. All watcher tasks accept a `VERSION` variable (default: `v1`):

```bash
# Configure watcher (generates config.yaml and workflow.yaml from templates)
task watcher:config

# Deploy to CRE
task watcher:deploy

# Config + deploy in one step
task watcher:release

# Simulate with a transaction
task watcher:simulate TX_HASH=0x...

# For future v2 contracts
task watcher:deploy VERSION=v2
```

**Mock server for local simulation:**

```bash
task mock:start   # Start mock server
task mock:stop    # Stop mock server
task mock:logs    # View logs
```

## License

See [LICENSE](LICENSE) for details.