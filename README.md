# CREC SDK Extension: DTA

A Go SDK extension for Digital Token Asset (DTA) operations on CREC.

## Installation

```bash
go get github.com/smartcontractkit/crec-sdk-ext-dta
```

## Overview

This extension provides utilities for preparing DTA (Digital Token Asset) operations for fund subscriptions, redemptions, and management on blockchain networks.

The extension supports two versions:
- **v0**: Works with DTAOpenMarketplace and DTAWallet contracts
- **v1**: Works with DTARequestManagement and DTARequestSettlement contracts

## Usage

### v0 (DTAOpenMarketplace / DTAWallet)

```go
import (
    v0 "github.com/smartcontractkit/crec-sdk-ext-dta/v0"
)

// Create the DTA v0 extension
ext, err := v0.New(&v0.Options{
    DTAOpenMarketplaceAddress: "0x...",  // DTAOpenMarketplace contract address
    DTAWalletAddress:          "0x...",  // DTAWallet contract address
    AccountAddress:            "0x...",  // Your account address
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

// Register distributor
op, err := ext.PrepareRegisterDistributorOperation(distributorAddr, distributorWalletAddr)

// Allow DTA for fund token
op, err := ext.PrepareAllowDTAOperation(dtaAddr, chainSelector, fundTokenId, fundTokenAddr, v0.TokenBurnTypeBurn)
```

### v1 (DTARequestManagement / DTARequestSettlement)

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

### DTAOpenMarketplace / DTARequestManagement Operations

| Operation | Description |
|-----------|-------------|
| `PrepareRequestSubscriptionOperation` | Request subscription to a fund |
| `PrepareRequestRedemptionOperation` | Request redemption from a fund |
| `PrepareRequestSubscriptionWithTokenApprovalOperation` | Request subscription with token approval |
| `PrepareProcessDistributorRequestOperation` | Process a pending distributor request |
| `PrepareCancelDistributorRequestOperation` | Cancel a distributor request |
| `PrepareRegisterDistributorOperation` | Register a new distributor |
| `PrepareRegisterFundAdminOperation` | Register a new fund admin |
| `PrepareRegisterFundTokenOperation` | Register a new fund token |
| `PrepareAllowDistributorForTokenOperation` | Allow distributor for a token |
| `PrepareDisallowDistributorForTokenOperation` | Disallow distributor for a token |
| `PrepareEnableFundTokenOperation` | Enable a fund token |
| `PrepareDisableFundTokenOperation` | Disable a fund token |

### DTAWallet / DTARequestSettlement Operations

| Operation | Description |
|-----------|-------------|
| `PrepareAllowDTAOperation` | Allow a DTA address for a fund token |
| `PrepareDisallowDTAOperation` | Disallow a DTA address for a fund token |
| `PrepareWithdrawTokensOperation` | Withdraw tokens from the wallet |
| `PrepareTransferOwnershipOperation` | Transfer contract ownership |
| `PrepareRenounceOwnershipOperation` | Renounce contract ownership |
| `PrepareCompleteRequestProcessingOperation` | Complete request processing |

### v1-Specific Operations

| Operation | Description |
|-----------|-------------|
| `PrepareVerifyDistributorWalletOperation` | Verify distributor wallet ownership |
| `PrepareForceAllowDistributorForTokenOperation` | Force allow distributor (admin) |

## Token Burn Types

| Constant | Value | Description |
|----------|-------|-------------|
| `TokenBurnTypeNone` | 0 | No burn behavior |
| `TokenBurnTypeBurn` | 1 | Burn tokens |
| `TokenBurnTypeTransfer` | 2 | Transfer tokens |

## License

See [LICENSE](LICENSE) for details.
