// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IStrategy.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAVSDirectory.sol";

/**
 * @title Storage variables for the `StrategyManager` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract StrategyManagerStorage is IStrategyManager {
    // Constants

    /// @notice The EIP-712 typehash for the deposit struct used by the contract
    bytes32 public constant DEPOSIT_TYPEHASH =
        keccak256("Deposit(address staker,address strategy,address token,uint256 amount,uint256 nonce,uint256 expiry)");

    // maximum length of dynamic arrays in `stakerStrategyList` mapping, for sanity's sake
    uint8 internal constant MAX_STAKER_STRATEGY_LIST_LENGTH = 32;

    // index for flag that pauses deposits when set
    uint8 internal constant PAUSED_DEPOSITS = 0;

    /// @notice default address for burning slashed shares and transferring underlying tokens
    address public constant DEFAULT_BURN_ADDRESS = 0x00000000000000000000000000000000000E16E4;

    // Immutables

    IDelegationManager public immutable delegation;

    // Mutatables

    /// @dev Do not remove, deprecated storage.
    bytes32 internal __deprecated_DOMAIN_SEPARATOR;

    /// @notice Returns the signature `nonce` for each `signer`.
    mapping(address signer => uint256 nonce) public nonces;

    /// @notice Returns the permissioned address that can whitelist strategies.
    address public strategyWhitelister;

    /// @dev Do not remove, deprecated storage.
    uint256 private __deprecated_withdrawalDelayBlocks;

    /// @notice Returns the number of deposited `shares` for a `staker` for a given `strategy`.
    /// @dev All of these shares may not be withdrawable if the staker has delegated to an operator that has been slashed.
    mapping(address staker => mapping(IStrategy strategy => uint256 shares)) public stakerDepositShares;

    /// @notice Returns a list of the `strategies` that a `staker` is currently staking in.
    mapping(address staker => IStrategy[] strategies) public stakerStrategyList;

    /// @dev Do not remove, deprecated storage.
    mapping(bytes32 withdrawalRoot => bool pending) private __deprecated_withdrawalRootPending;

    /// @dev Do not remove, deprecated storage.
    mapping(address staker => uint256 totalQueued) private __deprecated_numWithdrawalsQueued;

    /// @notice Returns whether a `strategy` is `whitelisted` for deposits.
    mapping(IStrategy strategy => bool whitelisted) public strategyIsWhitelistedForDeposit;

    /// @dev Do not remove, deprecated storage.
    mapping(address avs => uint256 shares) private __deprecated_beaconChainETHSharesToDecrementOnWithdrawal;

    /// @dev Do not remove, deprecated storage.
    mapping(IStrategy strategy => bool) private __deprecated_thirdPartyTransfersForbidden;

    /// @notice Returns the amount of `shares` that have been slashed on EigenLayer but not burned yet. Takes 3 storage slots.
    EnumerableMap.AddressToUintMap internal burnableShares;

    // Construction

    /**
     * @param _delegation The delegation contract of EigenLayer.
     */
    constructor(
        IDelegationManager _delegation
    ) {
        delegation = _delegation;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[36] private __gap;
}
