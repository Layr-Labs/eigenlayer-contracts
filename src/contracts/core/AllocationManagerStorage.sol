// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IAllocationManager.sol";
import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";
import {Snapshots} from "../libraries/Snapshots.sol";

abstract contract AllocationManagerStorage is IAllocationManager {
    using Snapshots for Snapshots.DefaultWadHistory;

    // Constants

    /// @dev Index for flag that pauses operator allocations/deallocations when set.
    uint8 internal constant PAUSED_MODIFY_ALLOCATIONS = 0;

    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;

    /// @dev BIPS factor for slashable bips
    uint256 internal constant BIPS_FACTOR = 10_000;

    /// @dev Maximum number of pending updates that can be queued for allocations/deallocations
    uint256 internal constant MAX_PENDING_UPDATES = 1;

    // Immutables

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegation;

    /// @notice The AVSDirectory contract for EigenLayer
    IAVSDirectory public immutable avsDirectory;

    /// @notice Delay before deallocations are clearable and can be added back into freeMagnitude
    /// In this window, deallocations still remain slashable by the operatorSet they were allocated to.
    uint32 public immutable DEALLOCATION_DELAY;

    /// @notice Delay before alloaction delay modifications take effect.
    uint32 public immutable ALLOCATION_CONFIGURATION_DELAY;

    // Mutatables

    /// @notice Returns snapshots of max magnitude for each `operator` for a given `strategy`.
    /// @dev This value starts at 100% (1e18) and decreases with slashing.
    mapping(address operator => mapping(IStrategy strategy => Snapshots.DefaultWadHistory)) internal
        _maxMagnitudeHistory;

    /// @notice Returns the amount of magnitude that is not available for allocation for each `operator` for a given `strategy`.
    /// @dev This value increases with allocations and slashing, and decreases with deallocations; should never exceed 100% (1e18).
    mapping(address operator => mapping(IStrategy strategy => uint64)) public encumberedMagnitude;

    /// @notice Returns the magnitude info for each `operator` for a given `strategy` and operator set (`operatorSetKey`).
    mapping(address operator => mapping(IStrategy strategy => mapping(bytes32 operatorSetKey => MagnitudeInfo)))
        internal _operatorMagnitudeInfo;

    /// @notice Returns pending deallocations for each `operator` for a given `strategy`.
    mapping(address operator => mapping(IStrategy strategy => DoubleEndedQueue.Bytes32Deque)) internal deallocationQueue;

    /// @notice Returns the allocation delay info for each `operator`; the delay and whether or not it's previously been set.
    mapping(address operator => AllocationDelayInfo) internal _allocationDelayInfo;

    // Construction

    constructor(
        IDelegationManager _delegation,
        IAVSDirectory _avsDirectory,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    ) {
        delegation = _delegation;
        avsDirectory = _avsDirectory;
        DEALLOCATION_DELAY = _DEALLOCATION_DELAY;
        ALLOCATION_CONFIGURATION_DELAY = _ALLOCATION_CONFIGURATION_DELAY;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[45] private __gap;
}
