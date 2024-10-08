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

    /// @notice Delay before deallocations are completable and can be added back into freeMagnitude
    /// In this window, deallocations still remain slashable by the operatorSet they were allocated to.
    uint32 public immutable DEALLOCATION_DELAY;

    /// @dev Delay before alloaction delay modifications take effect.
    uint32 public immutable ALLOCATION_CONFIGURATION_DELAY; // QUESTION: 21 days?

    // Mutatables

    /// @notice Mapping: operator => strategy => snapshotted maxMagnitude
    /// Note that maxMagnitude is monotonically decreasing and is decreased on slashing
    mapping(address => mapping(IStrategy => Snapshots.DefaultWadHistory)) internal _maxMagnitudeHistory;

    /// @notice Mapping: operator => strategy => the amount of magnitude that is not available for allocation
    mapping(address => mapping(IStrategy => uint64)) public encumberedMagnitude;

    /// @notice Mapping: operator => strategy => operatorSet (encoded) => MagnitudeInfo
    mapping(address => mapping(IStrategy => mapping(bytes32 => MagnitudeInfo))) internal _operatorMagnitudeInfo;

    /// @notice Mapping: operator => strategy => operatorSet[] (encoded) to keep track of pending modifications
    mapping(address => mapping(IStrategy => DoubleEndedQueue.Bytes32Deque)) internal modificationQueue;

    /// @notice Mapping: operator => allocation delay (in seconds) for the operator.
    /// This determines how long it takes for allocations to take effect in the future.
    mapping(address => AllocationDelayInfo) internal _allocationDelayInfo;
    
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
    uint256[44] private __gap;
}
