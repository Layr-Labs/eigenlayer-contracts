// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";

import "../interfaces/IAllocationManager.sol";
import "../interfaces/IDelegationManager.sol";

import {Snapshots} from "../libraries/Snapshots.sol";

abstract contract AllocationManagerStorage is IAllocationManager {
    using Snapshots for Snapshots.DefaultWadHistory;

    // Constants

    /// @dev Index for flag that pauses operator allocations/deallocations when set.
    uint8 internal constant PAUSED_MODIFY_ALLOCATIONS = 0;

    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;

    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION = 2;

    // Immutables

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegation;

    /// @notice Delay before deallocations are clearable and can be added back into freeMagnitude
    /// In this window, deallocations still remain slashable by the operatorSet they were allocated to.
    uint32 public immutable DEALLOCATION_DELAY;

    /// @notice Delay before alloaction delay modifications take effect.
    uint32 public immutable ALLOCATION_CONFIGURATION_DELAY;

    // Mutatables

    /// AVS => OPERATOR SET

    /// @dev Contains the AVS's configured registrar contract that handles registration/deregistration
    /// Note: if set to 0, defaults to the AVS's address
    mapping(address avs => IAVSRegistrar) internal _avsRegistrar;

    /// @dev Lists the operator set ids an AVS has created
    mapping(address avs => EnumerableSet.UintSet) internal _operatorSets;

    /// @dev Lists the strategies an AVS supports for an operator set
    mapping(bytes32 operatorSetKey => EnumerableSet.AddressSet) internal _operatorSetStrategies;

    /// @dev Lists the members of an AVS's operator set
    mapping(bytes32 operatorSetKey => EnumerableSet.AddressSet) internal _operatorSetMembers;

    /// OPERATOR => OPERATOR SET (REGISTRATION/DEREGISTRATION)

    /// @notice Returns the allocation delay info for each `operator`; the delay and whether or not it's previously been set.
    mapping(address operator => AllocationDelayInfo) internal _allocationDelayInfo;

    /// @dev Lists the operator sets the operator is registered for. Note that an operator
    /// can be registered without allocated stake. Likewise, an operator can allocate
    /// without being registered.
    mapping(address operator => EnumerableSet.Bytes32Set) internal registeredSets;

    /// @dev Lists the operator sets the operator has outstanding allocations in.
    mapping(address operator => EnumerableSet.Bytes32Set) internal allocatedSets;

    /// @dev Contains the operator's registration status for an operator set.
    mapping(address operator => mapping(bytes32 operatorSetKey => RegistrationStatus)) internal registrationStatus;

    /// @dev For an operator set, lists all strategies an operator has outstanding allocations from.
    mapping(address operator => mapping(bytes32 operatorSetKey => EnumerableSet.AddressSet)) internal
        allocatedStrategies;

    /// @dev For an operator set and strategy, the current allocated magnitude and any pending modification
    mapping(address operator => mapping(bytes32 operatorSetKey => mapping(IStrategy strategy => Allocation))) internal
        allocations;

    /// OPERATOR => STRATEGY (MAX/USED AND DEALLOCATIONS)

    /// @dev Contains a history of the operator's maximum magnitude for a given strategy
    mapping(address operator => mapping(IStrategy strategy => Snapshots.DefaultWadHistory)) internal
        _maxMagnitudeHistory;

    /// @dev For a strategy, contains the amount of magnitude an operator has allocated to operator sets
    mapping(address operator => mapping(IStrategy strategy => uint64)) internal encumberedMagnitude;

    /// @dev For a strategy, keeps an ordered queue of operator sets that have pending deallocations
    /// These must be completed in order to free up magnitude for future allocation
    mapping(address operator => mapping(IStrategy strategy => DoubleEndedQueue.Bytes32Deque)) internal deallocationQueue;

    /// @dev Lists the AVSs who has registered metadata and claimed itself as an AVS
    /// @notice bool is not used and is always true if the avs has registered metadata
    mapping(address avs => bool) internal _avsRegisteredMetadata;

    // Construction

    constructor(IDelegationManager _delegation, uint32 _DEALLOCATION_DELAY, uint32 _ALLOCATION_CONFIGURATION_DELAY) {
        delegation = _delegation;
        DEALLOCATION_DELAY = _DEALLOCATION_DELAY;
        ALLOCATION_CONFIGURATION_DELAY = _ALLOCATION_CONFIGURATION_DELAY;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[36] private __gap;
}
