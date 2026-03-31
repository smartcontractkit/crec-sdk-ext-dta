// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title MockDTARequestSettlement
/// @notice Shared settlement mock for testing the DTA v1/v2 extension watchers
///         and deployed certification flows.
/// @dev    Emits DTASettlementOpened and provides a stub completeRequestProcessing.
contract MockDTARequestSettlement {

    // --- Enums (match IDTAMessage) ---
    enum DistributorRequestType { None, Subscription, Redemption }
    enum Currency { None, USD, EUR, GBP }

    // --- Events (match DTARequestSettlement ABI) ---
    event DTASettlementOpened(
        address indexed fundAdminAddr,
        bytes32 indexed fundTokenId,
        uint8   indexed requestType,            // DistributorRequestType
        address distributorAddr,
        uint64  dtaChainSelector,
        address dtaAddr,                         // DTARequestManagement address
        bytes32 requestId,
        address distributorWalletAddr,
        uint256 shares,
        uint256 amount,
        uint8   currency                         // Currency enum
    );

    address public owner;
    address public managementAddr;

    bytes32 public constant MOCK_FUND_TOKEN_ID = keccak256("mock-fund-token");
    address public constant MOCK_FUND_ADMIN = address(0xAAAA);
    address public constant MOCK_DISTRIBUTOR = address(0xD15C);
    address public constant MOCK_DISTRIBUTOR_WALLET = address(0xF111);

    constructor(address _managementAddr) {
        owner = msg.sender;
        managementAddr = _managementAddr;
    }

    function updateManagementAddr(address _managementAddr) external {
        require(msg.sender == owner, "only owner");
        managementAddr = _managementAddr;
    }

    /// @notice Emit a DTASettlementOpened event with deterministic data.
    ///         dtaAddr is set to the managementAddr so the handler can call
    ///         getDistributorRequest and getFundToken on it.
    function emitDTASettlementOpened() external returns (bytes32 requestId) {
        requestId = keccak256(abi.encodePacked(block.number, msg.sender, "settlement"));
        emit DTASettlementOpened(
            MOCK_FUND_ADMIN,                        // fundAdminAddr (indexed)
            MOCK_FUND_TOKEN_ID,                     // fundTokenId (indexed)
            uint8(DistributorRequestType.Subscription), // requestType (indexed)
            MOCK_DISTRIBUTOR,                       // distributorAddr
            16015286601757825753,                    // dtaChainSelector (Sepolia)
            managementAddr,                         // dtaAddr → points to management mock
            requestId,                              // requestId
            MOCK_DISTRIBUTOR_WALLET,                // distributorWalletAddr
            1000e18,                                // shares
            5000e6,                                 // amount
            uint8(Currency.USD)                     // currency
        );
    }

    /// @notice Stub for completeRequestProcessing — the handler only reads its ABI signature
    ///         to build a PaymentCallback, it does not actually call this function.
    function completeRequestProcessing(
        bytes32 /* requestId */,
        uint256 /* shares */,
        uint256 /* amount */,
        uint256 /* nav */,
        bool    /* success */,
        bytes calldata /* err */
    ) external pure {
        // no-op: mock stub
    }
}
