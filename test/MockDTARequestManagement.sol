// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title MockDTARequestManagement
/// @notice Minimal v1-compatible mock for testing the DTA v1 extension watcher
///         and deployed v1 certification flow.
/// @dev    This contract intentionally uses the v1 DistributorRequest tuple shape
///         without `referenceID`. Use MockDTARequestManagementV2 for dta.v2.
contract MockDTARequestManagement {

    // --- Enums (match IDTAMessage) ---
    enum DistributorRequestType { None, Subscription, Redemption }
    enum RequestStatus { None, Pending, Processing, Processed, Canceled, Failed }
    enum Currency { None, USD, EUR, GBP }

    // --- Structs (match Solidity ABI) ---
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

    struct DTAPayment {
        Currency offChainPaymentCurrency;
        address paymentTokenSourceAddr;
        address paymentTokenDestAddr;
    }

    struct FundTokenData {
        address fundTokenAddr;
        uint8 navFeedDecimals;
        uint8 purchaseTokenRoundingDecimals;
        uint8 purchaseTokenDecimals;
        uint8 fundRoundingDecimals;
        uint8 fundTokenDecimals;
        uint8 requestsPerDay;
        address navAddr;
        uint64 tokenChainSelector;
        address dtaRequestSettlementAddr;
        int24 timezoneOffsetSecs;
        uint24 navTTL;
        DTAPayment paymentInfo;
    }

    // --- Events (match DTARequestManagement ABI) ---
    event DistributorRequestProcessing(
        address indexed fundAdminAddr,
        bytes32 indexed fundTokenId,
        address indexed distributorAddr,
        bytes32 requestId,
        uint256 shares,
        uint256 amount
    );

    // Fixed fund token ID, fund admin and distributor for deterministic mock data
    bytes32 public constant MOCK_FUND_TOKEN_ID = keccak256("mock-fund-token");
    address public constant MOCK_FUND_ADMIN = address(0xAAAA);
    address public constant MOCK_DISTRIBUTOR = address(0xD15C);

    address public owner;
    address public settlementAddr;

    constructor(address _settlementAddr) {
        owner = msg.sender;
        settlementAddr = _settlementAddr;
    }

    function updateSettlementAddr(address _settlementAddr) external {
        require(msg.sender == owner, "only owner");
        settlementAddr = _settlementAddr;
    }

    /// @notice Emit a DistributorRequestProcessing event with deterministic data.
    function emitDistributorRequestProcessing() external returns (bytes32 requestId) {
        requestId = keccak256(abi.encodePacked(block.number, msg.sender));
        emit DistributorRequestProcessing(
            MOCK_FUND_ADMIN,        // fundAdminAddr (indexed)
            MOCK_FUND_TOKEN_ID,     // fundTokenId (indexed)
            MOCK_DISTRIBUTOR,       // distributorAddr (indexed)
            requestId,              // requestId
            1000e18,                // shares
            5000e6                  // amount (e.g. 5000 USDC)
        );
    }

    /// @notice Returns hardcoded DistributorRequest for any requestId.
    function getDistributorRequest(bytes32 /* requestId */) external view returns (DistributorRequest memory) {
        return DistributorRequest({
            shares: 1000e18,
            amount: 5000e6,
            fundTokenId: MOCK_FUND_TOKEN_ID,
            fundAdminAddr: MOCK_FUND_ADMIN,
            distributorAddr: MOCK_DISTRIBUTOR,
            createdAt: uint40(block.timestamp),
            requestType: DistributorRequestType.Subscription,
            status: RequestStatus.Processing
        });
    }

    /// @notice Returns hardcoded FundTokenData for any fundAdmin + fundTokenId.
    function getFundToken(address /* fundAdminAddr */, bytes32 /* fundTokenId */)
        external
        view
        returns (bool enabled, FundTokenData memory)
    {
        enabled = true;
        return (enabled, FundTokenData({
            fundTokenAddr: address(0xBBBB),
            navFeedDecimals: 8,
            purchaseTokenRoundingDecimals: 2,
            purchaseTokenDecimals: 6,
            fundRoundingDecimals: 2,
            fundTokenDecimals: 18,
            requestsPerDay: 1,
            navAddr: address(0xCCCC),
            tokenChainSelector: 16015286601757825753,
            dtaRequestSettlementAddr: settlementAddr,
            timezoneOffsetSecs: 0,
            navTTL: 3600,
            paymentInfo: DTAPayment({
                offChainPaymentCurrency: Currency.USD,
                paymentTokenSourceAddr: address(0xDDDD),
                paymentTokenDestAddr: address(0xEEEE)
            })
        }));
    }
}
