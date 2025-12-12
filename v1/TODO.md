# Event Decoration TODOs

These are ideas for enriching event data with additional fields that could be fetched
via contract reads or other means. The goal is to provide more context around events
for better visibility into the data we care about.

## DistributorRequestCanceled / DistributorRequestProcessed / DistributorRequestProcessing

For requests of any status change, we could call:
```solidity
function getDistributorRequest(bytes32 requestId) external view returns (DistributorRequest memory)
```

Which returns:
```solidity
struct DistributorRequest {
    uint256 shares;
    uint256 amount;
    bytes32 fundTokenId;
    address fundAdminAddr;
    address distributorAddr;
    uint40 createdAt;
    DistributorRequestType requestType;
    RequestStatus status;
}
```

### DistributorRequestCanceled - Additional Fields
```go
DTAAddr         common.Address `json:"dta_addr"`
FundAdminAddr   common.Address `json:"fund_admin_addr"`
Amount          *big.Int       `json:"amount"`
Shares          *big.Int       `json:"shares"`
RequestType     uint8          `json:"request_type"`
Status          uint8          `json:"status"`
CreatedAt       string         `json:"created_at"`
```

### DistributorRequestProcessed - Additional Fields
```go
DTAAddr         common.Address `json:"dta_addr"`
FundAdminAddr   common.Address `json:"fund_admin_addr"`
FundTokenID     common.Hash    `json:"fund_token_id"`
DistributorAddr common.Address `json:"distributor_addr"`
Amount          string         `json:"amount"`
RequestType     uint8          `json:"request_type"`
CreatedAt       string         `json:"created_at"`
```

### DistributorRequestProcessing - Additional Fields
```go
DTAAddr       common.Address `json:"dta_addr"`
FundAdminAddr common.Address `json:"fund_admin_addr"`
RequestType   uint8          `json:"request_type"`
Status        uint8          `json:"status"`
CreatedAt     string         `json:"created_at"`
```

## FundTokenRegistered

These fields would come from an extra read to:
```solidity
function getFundToken(address fundAdminAddr, bytes32 fundTokenId) returns (bool enabled, FundTokenData memory)
```

### FundTokenRegistered - Additional Fields
```go
DtaRequestSettlementAddr      common.Address `json:"dta_request_settlement_addr"`
NavFeedDecimals               uint8          `json:"nav_feed_decimals"`
NavTTL                        uint32         `json:"nav_ttl"`
TimezoneOffsetSecs            int64          `json:"timezone_offset_secs"`
PurchaseTokenDecimals         uint8          `json:"purchase_token_decimals"`
PurchaseTokenRoundingDecimals uint8          `json:"purchase_token_rounding_decimals"`
FundTokenDecimals             uint8          `json:"fund_token_decimals"`
FundRoundingDecimals          uint8          `json:"fund_rounding_decimals"`
RequestsPerDay                uint8          `json:"requests_per_day"`
PaymentTokenSourceAddr        string         `json:"payment_token_source_addr"`
PaymentSourceChainSelector    string         `json:"payment_source_chain_selector"`
PaymentTokenDestAddr          string         `json:"payment_token_dest_addr"`
PaymentDestChainSelector      string         `json:"payment_dest_chain_selector"`
PaymentOffChainCurrency       uint64         `json:"payment_off_chain_currency"`
Enabled                       bool           `json:"enabled"`
```

