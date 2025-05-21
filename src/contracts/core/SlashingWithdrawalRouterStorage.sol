// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/utils/structs/EnumerableMapUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/utils/structs/EnumerableSetUpgradeable.sol";
import "../interfaces/ISlashingWithdrawalRouter.sol";
import "../interfaces/IAllocationManager.sol";
import "../interfaces/IStrategyManager.sol";
import "../interfaces/IStrategy.sol";

abstract contract SlashingWithdrawalRouterStorage is ISlashingWithdrawalRouter {
    // Constants

    /// @dev The default burn address for slashed funds.
    address internal constant DEFAULT_BURN_ADDRESS = 0x00000000000000000000000000000000000E16E4;

    /// @notice The pause status for the `burnOrRedistributeShares` function.
    /// @dev Allows all burn or redistribution outflows to be temporarily halted.
    uint8 public constant PAUSED_BURN_OR_REDISTRIBUTE_SHARES = 0;

    // Immutable Storage

    /// @notice Returns the EigenLayer `AllocationManager` address.
    IAllocationManager public immutable allocationManager;

    /// @notice Returns the EigenLayer `StrategyManager` address.
    IStrategyManager public immutable strategyManager;

    // Mutable Storage

    /// @dev Returns a list of operator sets that have pending slash IDs.
    EnumerableSetUpgradeable.Bytes32Set internal _pendingOperatorSets;

    /// @dev Returns a list of pending slash IDs for a given operator set.
    mapping(bytes32 operatorSetKey => EnumerableSetUpgradeable.UintSet) internal _pendingSlashIds;

    /// @dev Returns an enumerable mapping of strategies to their underlying amounts for a given slash ID.
    mapping(bytes32 operatorSetKey => mapping(uint256 slashId => EnumerableMapUpgradeable.AddressToUintMap)) internal
        _pendingBurnOrRedistributions;

    /// @dev Returns the start block for a given slash ID.
    mapping(bytes32 operatorSetKey => mapping(uint256 slashId => uint32 startBlock)) internal _slashIdToStartBlock;

    /// @notice Returns the paused status for a given operator set and slash ID.
    mapping(bytes32 operatorSetKey => mapping(uint256 slashId => bool paused)) internal _paused;

    /// @dev Returns the burn or redistribution delay for a given strategy.
    uint32 internal _globalBurnOrRedistributionDelayBlocks;

    /// @dev Returns the operator set delay for a given strategy.
    mapping(address strategy => uint32 delay) internal _strategyBurnOrRedistributionDelayBlocks;

    // Constructor

    constructor(IAllocationManager _allocationManager, IStrategyManager _strategyManager) {
        allocationManager = _allocationManager;
        strategyManager = _strategyManager;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[42] private __gap;
}
