// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import "../interfaces/IAllocationManager.sol";
import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";

import {Snapshots} from "../libraries/Snapshots.sol";

abstract contract AllocationManagerStorage is IAllocationManager {
    // Constants
    
    /// @dev Index for flag that pauses operator allocations/deallocations when set.
    uint8 internal constant PAUSED_MODIFY_ALLOCATIONS = 0;

    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;

    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION = 2;

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

    /// @notice Mapping: operator => strategy => operatorSet[] (encoded) to keep track of pending deallocations
    mapping(address => mapping(IStrategy => DoubleEndedQueue.Bytes32Deque)) internal deallocationQueue;

    /// @notice Mapping: operator => allocation delay (in seconds) for the operator.
    /// This determines how long it takes for allocations to take effect in the future.
    mapping(address => AllocationDelayInfo) internal _allocationDelayInfo;

    /// @notice Mapping: avs => operatorSetId => Whether or not an operator set is valid.
    mapping(address => mapping(uint32 => bool)) public isOperatorSet;

    /// @notice Mapping: operator => List of operator sets that operator is registered to.
    /// @dev Each item is formatted as such: bytes32(abi.encodePacked(avs, uint96(operatorSetId)))
    mapping(address => EnumerableSet.Bytes32Set) internal _operatorSetsMemberOf;

    /// @notice Mapping: operatorSet => List of operators that are registered to the operatorSet
    /// @dev Each key is formatted as such: bytes32(abi.encodePacked(avs, uint96(operatorSetId)))
    mapping(bytes32 => EnumerableSet.AddressSet) internal _operatorSetMembers;

    /// @notice Mapping: operatorSet => List of strategies that the operatorSet contains
    /// @dev Each key is formatted as such: bytes32(abi.encodePacked(avs, uint96(operatorSetId)))
    mapping(bytes32 => EnumerableSet.AddressSet) internal _operatorSetStrategies;

    /// @notice Mapping: operator => avs => operatorSetId => operator registration status
    mapping(address => mapping(address => mapping(uint32 => OperatorSetRegistrationStatus))) public operatorSetStatus;

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
    uint256[47] private __gap;
}
