# DTA Watcher

The DTA watcher is a CRE (Chainlink Runtime Environment) workflow compiled to WASM that monitors DTA smart contract events on-chain. When an event is detected, the handler decodes it, enriches it with on-chain reference data, and posts a signed `VerifiableEvent` to the CREC API for delivery to users.

## Processing Flow

```
EVM Log (on-chain event)
  │
  ▼
OnLog()                              ─── Main handler entry point
  │
  ├─ BuildEVMEventFromLog()          ─── Decode raw log using contract ABI
  │
  ├─ DTARequestManagement event?
  │    └─ buildReferenceDataFromDTARequestManagementEvent()
  │         ├─ getDistributorRequest()    ─── On-chain read (if request_id present)
  │         └─ getFundToken()             ─── On-chain read (if fund_admin_addr + fund_token_id present)
  │
  ├─ DTASettlementOpened or DTASettlementClosed event?
  │    └─ buildReferenceDataFromDTASettlementEvent()
  │         ├─ getDistributorRequest()    ─── On-chain read
  │         ├─ getFundToken()             ─── On-chain read
  │         ├─ currency code lookup       ─── Off-chain reference data
  │         └─ buildPaymentRequest()      ─── Payment request for settlement
  │
  ├─ BuildVerifiableEventForEVMEvent()  ─── Package into VerifiableEvent
  │
  ▼
SignAndPostVerifiableEvent()            ─── Sign and POST to CREC API
```

## VerifiableEvent Structure

The output `VerifiableEvent` has this shape:

```json
{
  "name": "SubscriptionRequested",
  "service": "dta.v1",
  "chain_family": "evm",
  "chain_selector": "16015286601757825753",
  "timestamp": "2025-01-15T12:00:00Z",
  "chain_event": {
    "address": "0x...",
    "chain_id": "11155111",
    "tx_hash": "0x...",
    "block_number": 12345678,
    "block_timestamp": 1705305600,
    "event_signature": "SubscriptionRequested(address,bytes32,address,bytes32,uint256,uint40)",
    "topic_hash": "0x...",
    "log_index": 0,
    "params": { ... }
  },
  "data": {
    "type": "reference_data",
    "value": {
      "on_chain": [ ... ],
      "off_chain": [ ... ],
      "requests": [ ... ]
    }
  }
}
```

The `params` field contains the decoded Solidity event parameters. The `data` field contains enrichment data fetched by the handler.

## Bundle Schema Fields

Each event in the bundle exposes two JSON Schema fields:

- **`ParamsSchema`** — describes `chain_event.params` (the decoded Solidity event parameters). Auto-generated from the contract ABI by `gen/main.go` into `params_schema_gen.go`.
- **`DataSchema`** — describes the `data` enrichment payload (reference data, off-chain lookups, generated requests). Hand-written in `bundle.go` since it depends on handler logic, not just the ABI. `nil` for events with no enrichment.

Consumers can use `ParamsSchema` to decode the event parameters and `DataSchema` to understand what enrichment data (if any) is attached.

> **Tip:** If you're receiving events from the CREC API as `apiClient.Event`, the easiest path is to call `v1.DecodeFromEvent(ctx, event)` which handles all decoding for you — parsing the `VerifiableEvent`, decoding chain event params into typed Go structs, and extracting enrichment data (`FundTokenData`, `DistributorRequest`, `PaymentRequests`). The schemas here are primarily useful for non-Go consumers or UI rendering.

## Supported Events

The watcher monitors events across two DTA contracts. Each event's `params` schema is documented below.

---

### DTARequestManagement Events

#### DistributorRegistered

New distributor registered in the DTA system.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `distributor_addr` | `address` | string | Distributor address |

**Enrichment:** None (no `request_id` or `fund_token_id` to look up).

#### DistributorRequestCanceled

A distributor's subscription or redemption request was canceled.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address (indexed) |
| `fund_token_id` | `bytes32` | string | Fund token identifier (indexed) |
| `distributor_addr` | `address` | string | Distributor address (indexed) |
| `request_id` | `bytes32` | string | Request identifier |

**Enrichment:** `distributor_request` (on-chain), `fund_token_data` (on-chain).

#### DistributorRequestProcessed

A distributor's request has been fully processed with a result.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `request_id` | `bytes32` | string | Request identifier |
| `shares` | `uint256` | string | Number of shares |
| `status` | `uint8` | integer | Request status enum (0=None, 1=Pending, 2=Processing, 3=Processed, 4=Canceled, 5=Failed) |
| `error` | `bytes` | string | Error data (hex) |

**Enrichment:** `distributor_request` (on-chain), `fund_token_data` (on-chain).

#### DistributorRequestProcessing

A distributor's request is being actively processed.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address (indexed) |
| `fund_token_id` | `bytes32` | string | Fund token identifier (indexed) |
| `distributor_addr` | `address` | string | Distributor address (indexed) |
| `request_id` | `bytes32` | string | Request identifier |
| `shares` | `uint256` | string | Number of shares |
| `amount` | `uint256` | string | Amount |

**Enrichment:** `distributor_request` (on-chain), `fund_token_data` (on-chain).

#### FundAdminRegistered

A new fund administrator was registered.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address |

**Enrichment:** None.

#### FundTokenAllowlistUpdated

A distributor's allowlist status for a fund token was changed.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address (indexed) |
| `fund_token_id` | `bytes32` | string | Fund token identifier (indexed) |
| `distributor_addr` | `address` | string | Distributor address (indexed) |
| `allowed` | `bool` | boolean | Whether the distributor is allowed |

**Enrichment:** `fund_token_data` (on-chain).

#### FundTokenRegistered

A new fund token was created and registered.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address (indexed) |
| `fund_token_id` | `bytes32` | string | Fund token identifier (indexed) |
| `fund_token_addr` | `address` | string | Fund token contract address (indexed) |
| `nav_addr` | `address` | string | NAV feed address |
| `token_chain_selector` | `uint64` | string | Token chain selector |

**Enrichment:** `fund_token_data` (on-chain).

#### RedemptionRequested

An investor requested to redeem shares.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address (indexed) |
| `fund_token_id` | `bytes32` | string | Fund token identifier (indexed) |
| `distributor_addr` | `address` | string | Distributor address (indexed) |
| `request_id` | `bytes32` | string | Request identifier |
| `shares` | `uint256` | string | Number of shares to redeem |
| `created_at` | `uint40` | integer | Creation timestamp |

**Enrichment:** `distributor_request` (on-chain), `fund_token_data` (on-chain).

#### SubscriptionRequested

An investor requested to subscribe to a fund.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address (indexed) |
| `fund_token_id` | `bytes32` | string | Fund token identifier (indexed) |
| `distributor_addr` | `address` | string | Distributor address (indexed) |
| `request_id` | `bytes32` | string | Request identifier |
| `amount` | `uint256` | string | Subscription amount |
| `created_at` | `uint40` | integer | Creation timestamp |

**Enrichment:** `distributor_request` (on-chain), `fund_token_data` (on-chain).

---

### DTARequestSettlement Events

#### DTASettlementOpened

A settlement was initiated for a subscription or redemption.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address (indexed) |
| `fund_token_id` | `bytes32` | string | Fund token identifier (indexed) |
| `request_type` | `uint8` | integer | 1 = Subscription, 2 = Redemption (indexed) |
| `distributor_addr` | `address` | string | Distributor address |
| `dta_chain_selector` | `uint64` | string | DTA chain selector |
| `dta_addr` | `address` | string | DTA contract address |
| `request_id` | `bytes32` | string | Request identifier |
| `distributor_wallet_addr` | `address` | string | Distributor wallet address |
| `shares` | `uint256` | string | Number of shares |
| `amount` | `uint256` | string | Amount |
| `currency` | `uint8` | integer | Currency enum |

**Enrichment:** `distributor_request` (on-chain), `fund_token_data` (on-chain), currency code (off-chain), `payment_request` (generated).

#### DTASettlementClosed

A settlement completed or failed.

| Parameter | Solidity Type | JSON Type | Description |
|-----------|--------------|-----------|-------------|
| `fund_admin_addr` | `address` | string | Fund admin address (indexed) |
| `fund_token_id` | `bytes32` | string | Fund token identifier (indexed) |
| `request_type` | `uint8` | integer | 1 = Subscription, 2 = Redemption (indexed) |
| `distributor_addr` | `address` | string | Distributor address |
| `dta_chain_selector` | `uint64` | string | DTA chain selector |
| `dta_addr` | `address` | string | DTA contract address |
| `request_id` | `bytes32` | string | Request identifier |
| `success` | `bool` | boolean | Whether the settlement succeeded |
| `err` | `bytes` | string | Error data (hex bytes, empty on success) |

**Enrichment:** `distributor_request` (on-chain), `fund_token_data` (on-chain), currency code (off-chain). Reuses the same settlement enrichment pipeline as DTASettlementOpened, minus the payment request.

---

## Reference Data (Enrichment)

For events that have a `request_id`, `fund_admin_addr`, or `fund_token_id`, the handler enriches the event with on-chain data by calling view functions on the DTA contracts.

### DistributorRequest

Fetched via `DTARequestManagement.getDistributorRequest(bytes32 requestId)`.

| Field | Type | Description |
|-------|------|-------------|
| `shares` | string | Number of shares |
| `amount` | string | Amount |
| `fund_token_id` | string | Fund token ID (bytes32 hex) |
| `fund_admin_addr` | string | Fund admin address |
| `distributor_addr` | string | Distributor address |
| `created_at` | string | Creation timestamp |
| `request_type` | integer | 1=Subscription, 2=Redemption |
| `status` | integer | 0=None, 1=Pending, 2=Processing, 3=Processed, 4=Canceled, 5=Failed |

### FundTokenData

Fetched via `DTARequestManagement.getFundToken(address fundAdmin, bytes32 fundTokenId)`.

Returns an `enabled` boolean and a `FundTokenData` struct containing token configuration, NAV feed details, payment info, and chain selectors.

### PaymentRequest (DTASettlementOpened only)

Generated for settlement events to facilitate off-chain payment processing.

| Field | Type | Description |
|-------|------|-------------|
| `application_type` | string | Always "dta.v1" |
| `application_addr` | string | Settlement contract address |
| `e2e_id` | string | Request ID (end-to-end identifier) |
| `sender` | string | Payment sender address |
| `receiver` | string | Payment receiver address |
| `currency` | string | ISO currency code |
| `chain_id` | string | Chain ID |
| `amount` | number | Payment amount (scaled from on-chain decimals) |
| `expiration` | integer | Expiration timestamp |
| `custom_callback` | object | Callback to `completeRequestProcessing` on settlement contract |

## Solidity → JSON Type Mapping

The handler decodes Solidity types to JSON values using these conventions:

| Solidity Type | JSON Type | Notes |
|---------------|-----------|-------|
| `address` | string | 0x-prefixed hex, 42 chars |
| `bytes32` | string | 0x-prefixed hex, 66 chars |
| `bytes` | string | 0x-prefixed variable-length hex |
| `uint256` | string | Decimal string (too large for JSON number) |
| `uint128` | string | Decimal string |
| `uint64` | string | Decimal string (can exceed JS safe integer) |
| `uint40` | integer | Fits in JSON number |
| `uint8` | integer | Fits in JSON number |
| `bool` | boolean | |
