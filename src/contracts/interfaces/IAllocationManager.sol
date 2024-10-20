// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IPauserRegistry.sol";
import "./IStrategy.sol";
import "./ISignatureUtils.sol";

/// @notice Struct representing an operator set
struct OperatorSet {
    address avs;
    uint32 operatorSetId;
}

interface IAllocationManagerErrors {
    /// @dev Thrown when `wadToSlash` is zero or greater than 1e18
    error InvalidWadToSlash();
    /// @dev Thrown when `operator` is not a registered operator.
    error OperatorNotRegistered();
    /// @dev Thrown when two array parameters have mismatching lengths.
    error InputArrayLengthMismatch();
    /// @dev Thrown when an operator's allocation delay has yet to be set.
    error UninitializedAllocationDelay();
    /// @dev Thrown when provided `expectedTotalMagnitude` for a given allocation does not match `currentTotalMagnitude`.
    error InvalidExpectedTotalMagnitude();
    /// @dev Thrown when an invalid operator set is provided.
    error InvalidOperatorSet();
    /// @dev Thrown when an invalid operator is provided.
    error InvalidOperator();
    /// @dev Thrown when caller is not the delegation manager.
    error OnlyDelegationManager();
    /// @dev Thrown when an operator attempts to set their allocation for an operatorSet to the same value
    error SameMagnitude();
    /// @dev Thrown when an allocation is attempted for a given operator when they have pending allocations or deallocations.
    error ModificationAlreadyPending();
    /// @dev Thrown when an allocation is attempted that exceeds a given operators total allocatable magnitude.
    error InsufficientAllocatableMagnitude();
    /// @dev Thrown when attempting to slash an operator that has already been slashed at the given timestamp.
    error AlreadySlashedForTimestamp();
    /// @dev Thrown when calling a view function that requires a valid timestamp.
    error InvalidTimestamp();
    /// @dev Thrown when an invalid allocation delay is set
    error InvalidAllocationDelay();
    /// @dev Thrown when a slash is attempted on an operator who has not allocated to the strategy, operatorSet pair
    error OperatorNotAllocated();
    /// @dev Thrown when a strategy is already added to an operator set.
    error StrategyAlreadyInOperatorSet();
    /// @dev Thrown when a strategy is not in an operator set.
    error StrategyNotInOperatorSet();

    /// @dev Thrown when an operator does not exist in the DelegationManager
    error OperatorNotRegisteredToEigenLayer();
    /// @dev Thrown when an invalid AVS is provided.
    error InvalidAVS();

    /// @dev Thrown when attempting to spend a spent eip-712 salt.
    error SaltSpent();
    /// @dev Thrown when attempting to use an expired eip-712 signature.
    error SignatureExpired();
}

interface IAllocationManagerTypes {
    /**
     * @notice struct used to modify the allocation of slashable magnitude to list of operatorSets
     * @param strategy the strategy to allocate magnitude for
     * @param expectedMaxMagnitude the expected max magnitude of the operator (used to combat against race conditions with slashing)
     * @param operatorSets the operatorSets to allocate magnitude for
     * @param magnitudes the magnitudes to allocate for each operatorSet
     */
    struct MagnitudeAllocation {
        IStrategy strategy;
        uint64 expectedMaxMagnitude;
        OperatorSet[] operatorSets;
        uint64[] magnitudes;
    }

    /**
     * @notice struct used for operator magnitude updates. Stored in _operatorMagnitudeInfo mapping
     * @param currentMagnitude the current magnitude of the operator
     * @param pendingDiff the pending magnitude difference of the operator
     * @param effectTimestamp the timestamp at which the pending magnitude will take effect
     */
    struct MagnitudeInfo {
        uint64 currentMagnitude;
        int128 pendingDiff;
        uint32 effectTimestamp;
    }

    /**
     * @notice Struct containing allocation delay metadata for a given operator.
     * @param delay Current allocation delay if `pendingDelay` is non-zero and `pendingDelayEffectTimestamp` has elapsed.
     * @param pendingDelay Current allocation delay if it's non-zero and `pendingDelayEffectTimestamp` has elapsed.
     * @param effectTimestamp The timestamp for which `pendingDelay` becomes the curren allocation delay.
     */
    struct AllocationDelayInfo {
        uint32 delay;
        uint32 pendingDelay;
        uint32 effectTimestamp;
    }

    /**
     * @notice Struct containing parameters to slashing
     * @param operator the address to slash
     * @param operatorSetId the ID of the operatorSet the operator is being slashed on behalf of
     * @param strategies the set of strategies to slash
     * @param wadToSlash the parts in 1e18 to slash, this will be proportional to the operator's
     * slashable stake allocation for the operatorSet
     * @param description the description of the slashing provided by the AVS for legibility
     */
    struct SlashingParams {
        address operator;
        uint32 operatorSetId;
        IStrategy[] strategies;
        uint256 wadToSlash;
        string description;
    }

    /**
     * @param encumberedMagnitude the effective magnitude allocated to all operator sets
     * for the strategy
     * @param currentMagnitude the effective current magnitude allocated to a single operator set
     * for the strategy
     * @param pendingDiff the pending change in magnitude, if one exists
     * @param effectTimestamp the time after which `pendingDiff` will take effect
     */
    struct PendingMagnitudeInfo {
        uint64 encumberedMagnitude;
        uint64 currentMagnitude;
        int128 pendingDiff;
        uint32 effectTimestamp;
    }

    /**
     * @notice Struct representing the registration status of an operator with an operator set.
     * Keeps track of last deregistered timestamp for slashability concerns.
     * @param registered whether the operator is registered with the operator set
     * @param lastDeregisteredTimestamp the timestamp at which the operator was last deregistered
     */
    struct OperatorSetRegistrationStatus {
        bool registered;
        uint32 lastDeregisteredTimestamp;
    }
}

interface IAllocationManagerEvents is IAllocationManagerTypes {
    /// @notice Emitted when operator updates their allocation delay.
    event AllocationDelaySet(address operator, uint32 delay, uint32 effectTimestamp);

    /// @notice Emitted when an operator's magnitude is updated for a given operatorSet and strategy
    event OperatorSetMagnitudeUpdated(
        address operator, OperatorSet operatorSet, IStrategy strategy, uint64 magnitude, uint32 effectTimestamp
    );

    /// @notice Emitted when operator's encumbered magnitude is updated for a given strategy
    event EncumberedMagnitudeUpdated(address operator, IStrategy strategy, uint64 encumberedMagnitude);

    /// @notice Emitted when an operator's total magnitude is updated for a given strategy
    event MaxMagnitudeUpdated(address operator, IStrategy strategy, uint64 totalMagnitude);

    /// @notice Emitted when an operator is slashed by an operator set for a strategy
    /// `wadSlashed` is the proportion of the operator's total delegated stake that was slashed
    event OperatorSlashed(
        address operator, OperatorSet operatorSet, IStrategy[] strategies, uint256[] wadSlashed, string description
    );

    /// @notice Emitted when an operator set is created by an AVS.
    event OperatorSetCreated(OperatorSet operatorSet);

    /// @notice Emitted when an operator is added to an operator set.
    event OperatorAddedToOperatorSet(address indexed operator, OperatorSet operatorSet);

    /// @notice Emitted when an operator is removed from an operator set.
    event OperatorRemovedFromOperatorSet(address indexed operator, OperatorSet operatorSet);

    /// @notice Emitted when a strategy is added to an operator set.
    event StrategyAddedToOperatorSet(OperatorSet operatorSet, IStrategy strategy);

    /// @notice Emitted when a strategy is removed from an operator set.
    event StrategyRemovedFromOperatorSet(OperatorSet operatorSet, IStrategy strategy);

    /// @notice Emitted when an AVS migrates to using operator sets.
    event AVSMigratedToOperatorSets(address indexed avs);

    /// @notice Emitted when an operator is migrated from M2 registration to operator sets.
    event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds);
}

interface IAllocationManager is ISignatureUtils, IAllocationManagerErrors, IAllocationManagerEvents {
    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     */
    function initialize(address initialOwner, IPauserRegistry _pauserRegistry, uint256 initialPausedStatus) external;

    /**
     * @notice Called by an AVS to slash an operator in a given operator set
     */
    function slashOperator(
        SlashingParams calldata params
    ) external;

    /**
     * @notice Modifies the propotions of slashable stake allocated to a list of operatorSets for a set of strategies
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @dev Updates encumberedMagnitude for the updated strategies
     * @dev msg.sender is used as operator
     */
    function modifyAllocations(
        MagnitudeAllocation[] calldata allocations
    ) external;

    /**
     * @notice This function takes a list of strategies and adds all completable deallocations for each strategy,
     * updating the encumberedMagnitude of the operator as needed.
     *
     * @param operator address to complete deallocations for
     * @param strategies a list of strategies to complete deallocations for
     * @param numToComplete a list of number of pending deallocations to complete for each strategy
     *
     * @dev can be called permissionlessly by anyone
     */
    function clearDeallocationQueue(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToComplete
    ) external;

    /**
     * @notice Called by the delegation manager to set an operator's allocation delay.
     * This is set when the operator first registers, and is the time between an operator
     * allocating magnitude to an operator set, and the magnitude becoming slashable.
     * @dev Note that if an operator's allocation delay is 0, it has not been set yet,
     * and the operator will be unable to allocate magnitude to any operator set.
     * @param operator The operator to set the delay on behalf of.
     * @param delay the allocation delay in seconds
     */
    function setAllocationDelay(address operator, uint32 delay) external;

    /**
     * @notice Called by an operator to set their allocation delay. This is the time between an operator
     * allocating magnitude to an operator set, and the magnitude becoming slashable.
     * @dev Note that if an operator's allocation delay is 0, it has not been set yet,
     * and the operator will be unable to allocate magnitude to any operator set.
     * @param delay the allocation delay in seconds
     */
    function setAllocationDelay(
        uint32 delay
    ) external;

    /**
     * @notice Called by an AVS to create a list of new operatorSets.
     *
     * @param operatorSetIds The IDs of the operator set to initialize.
     *
     * @dev msg.sender must be the AVS.
     * @dev The AVS may create operator sets before it becomes an operator set AVS.
     */
    function createOperatorSets(
        uint32[] calldata operatorSetIds
    ) external;

    /**
     * @notice Sets the AVS as an operator set AVS, preventing legacy M2 operator registrations.
     *
     * @dev msg.sender must be the AVS.
     */
    function becomeOperatorSetAVS() external;

    // /**
    //  * @notice Called by an AVS to migrate operators that have a legacy M2 registration to operator sets.
    //  *
    //  * @param operators The list of operators to migrate
    //  * @param operatorSetIds The list of operatorSets to migrate the operators to
    //  *
    //  * @dev The msg.sender used is the AVS
    //  * @dev The operator can only be migrated at most once per AVS
    //  * @dev The AVS can no longer register operators via the legacy M2 registration path once it begins migration
    //  * @dev The operator is deregistered from the M2 legacy AVS once migrated
    //  */
    // function migrateOperatorsToOperatorSets(
    //     address[] calldata operators,
    //     uint32[][] calldata operatorSetIds
    // ) external;

    /**
     *  @notice Called by AVSs to add an operator to a list of operatorSets.
     *
     *  @param operator The address of the operator to be added to the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *  @param operatorSignature The signature of the operator on their intent to register.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev The operator must not have a pending deregistration from the operator set.
     */
    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;

    /**
     * @notice Called by an operator to deregister from an operator set
     *
     * @param operator The operator to deregister from the operatorSets.
     * @param avs The address of the AVS to deregister the operator from.
     * @param operatorSetIds The IDs of the operator sets.
     * @param operatorSignature the signature of the operator on their intent to deregister or empty if the operator itself is calling
     *
     * @dev if the operatorSignature is empty, the caller must be the operator
     * @dev this will likely only be called in case the AVS contracts are in a state that prevents operators from deregistering
     */
    function forceDeregisterFromOperatorSets(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;

    /**
     *  @notice Called by AVSs to remove an operator from an operator set.
     *
     *  @param operator The address of the operator to be removed from the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *
     *  @dev msg.sender is used as the AVS.
     */
    function deregisterOperatorFromOperatorSets(address operator, uint32[] calldata operatorSetIds) external;

    /**
     *  @notice Called by AVSs to add a set of strategies to an operator set.
     *
     *  @param operatorSetId The ID of the operator set.
     *  @param strategies The addresses of the strategies to be added to the operator set.
     *
     *  @dev msg.sender is used as the AVS.
     */
    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] calldata strategies) external;

    /**
     *  @notice Called by AVSs to remove a set of strategies from an operator set.
     *
     *  @param operatorSetId The ID of the operator set.
     *  @param strategies The addresses of the strategies to be removed from the operator set.
     *
     *  @dev msg.sender is used as the AVS.
     */
    function removeStrategiesFromOperatorSet(uint32 operatorSetId, IStrategy[] calldata strategies) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Returns the effective magnitude info for each of an operator's operator sets.
     * This method fetches the complete list of an operator's operator sets, then applies any
     * completable allocation modifications to return the effective, up-to-date current and
     * pending magnitude allocations for each operator set.
     * @param operator the operator to query
     * @param strategy the strategy to get allocation info for
     * @return the list of the operator's operator sets
     * @return the corresponding allocation details for each operator set
     */
    function getAllocationInfo(
        address operator,
        IStrategy strategy
    ) external view returns (OperatorSet[] memory, MagnitudeInfo[] memory);

    /**
     * @notice Returns the effective magnitude info for each operator set. This method
     * automatically applies any completable modifications, returning the effective
     * current and pending allocations for each operator set.
     * @param operator the operator to query
     * @param strategy the strategy to get allocation info for
     * @param operatorSets the operatorSets to get allocation info for
     * @return The magnitude info for each operator set
     */
    function getAllocationInfo(
        address operator,
        IStrategy strategy,
        OperatorSet[] calldata operatorSets
    ) external view returns (MagnitudeInfo[] memory);

    /**
     * @notice Returns the effective magnitude info for each operator for each strategy for the operatorSet This method
     * automatically applies any completable modifications, returning the effective
     * current and pending allocations for each operator set.
     * @param operatorSet the operator set to query
     * @param strategies the strategies to get allocation info for
     * @param operators the operators to get allocation info for
     * @return The magnitude info for each operator for each strategy
     */
    function getAllocationInfo(
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies,
        address[] calldata operators
    ) external view returns (MagnitudeInfo[][] memory);

    /**
     * @notice For a strategy, get the amount of magnitude not currently allocated to any operator set
     * @param operator the operator to query
     * @param strategy the strategy to get allocatable magnitude for
     * @return magnitude available to be allocated to an operator set
     */
    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64);

    /**
     * @notice Returns the maximum magnitude an operator can allocate for the given strategies
     * @dev The max magnitude of an operator starts at WAD (1e18), and is decreased anytime
     * the operator is slashed. This value acts as a cap on the total magnitude of the operator.
     * @param operator the operator to query
     * @param strategies the strategies to get the max magnitudes for
     * @return the max magnitudes for each strategy
     */
    function getMaxMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory);

    /**
     * @notice Returns the maximum magnitude an operator can allocate for the given strategies
     * at a given timestamp
     * @dev The max magnitude of an operator starts at WAD (1e18), and is decreased anytime
     * the operator is slashed. This value acts as a cap on the total magnitude of the operator.
     * @param operator the operator to query
     * @param strategies the strategies to get the max magnitudes for
     * @param timestamp the timestamp at which to check the max magnitudes
     * @return the max magnitudes for each strategy
     */
    function getMaxMagnitudesAtTimestamp(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint64[] memory);

    /**
     * @notice Returns the time in seconds between an operator allocating slashable magnitude
     * and the magnitude becoming slashable. If the delay has not been set, `isSet` will be false.
     * @dev The operator must have a configured delay before allocating magnitude
     * @param operator The operator to query
     * @return isSet Whether the operator has configured a delay
     * @return delay The time in seconds between allocating magnitude and magnitude becoming slashable
     */
    function getAllocationDelay(
        address operator
    ) external view returns (bool isSet, uint32 delay);

    /**
     * @notice returns the minimum operatorShares and the slashableOperatorShares for an operator, list of strategies,
     * and an operatorSet before a given timestamp. This is used to get the shares to weight operators by given ones slashing window.
     * @param operatorSet the operatorSet to get the shares for
     * @param operators the operators to get the shares for
     * @param strategies the strategies to get the shares for
     * @param beforeTimestamp the timestamp to get the shares at
     */
    function getMinDelegatedAndSlashableOperatorShares(
        OperatorSet calldata operatorSet,
        address[] calldata operators,
        IStrategy[] calldata strategies,
        uint32 beforeTimestamp
    ) external view returns (uint256[][] memory, uint256[][] memory);

    function isOperatorSetAVS(
        address avs
    ) external view returns (bool);

    /// @notice Returns true if the operator set is valid.
    function isOperatorSet(address avs, uint32 operatorSetId) external view returns (bool);

    /**
     * @notice Returns operator set an operator is registered to in the order they were registered.
     * @param operator The operator address to query.
     * @param index The index of the enumerated list of operator sets.
     */
    function operatorSetsMemberOfAtIndex(address operator, uint256 index) external view returns (OperatorSet memory);

    /**
     * @notice Retursn the operator registered to an operatorSet in the order that it was registered.
     *  @param operatorSet The operatorSet to query.
     *  @param index The index of the enumerated list of operators.
     */
    function operatorSetMemberAtIndex(OperatorSet memory operatorSet, uint256 index) external view returns (address);

    /**
     * @notice Returns the number of operator sets an operator is registered to.
     * @param operator the operator address to query
     */
    function getNumOperatorSetsOfOperator(
        address operator
    ) external view returns (uint256);

    /**
     * @notice Returns an array of operator sets an operator is registered to.
     * @param operator The operator address to query.
     * @param start The starting index of the array to query.
     *  @param length The amount of items of the array to return.
     */
    function getOperatorSetsOfOperator(
        address operator,
        uint256 start,
        uint256 length
    ) external view returns (OperatorSet[] memory operatorSets);

    /**
     * @notice Returns an array of operators registered to the operatorSet.
     * @param operatorSet The operatorSet to query.
     * @param start The starting index of the array to query.
     * @param length The amount of items of the array to return.
     */
    function getOperatorsInOperatorSet(
        OperatorSet memory operatorSet,
        uint256 start,
        uint256 length
    ) external view returns (address[] memory operators);

    /**
     * @notice Returns an array of strategies in the operatorSet.
     * @param operatorSet The operatorSet to query.
     */
    function getStrategiesInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (IStrategy[] memory strategies);

    /**
     * @notice Returns the number of operators registered to an operatorSet.
     * @param operatorSet The operatorSet to get the member count for
     */
    function getNumOperatorsInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (uint256);

    /**
     *  @notice Returns the total number of operator sets an operator is registered to.
     *  @param operator The operator address to query.
     */
    function inTotalOperatorSets(
        address operator
    ) external view returns (uint256);

    /**
     * @notice Returns whether or not an operator is registered to an operator set.
     * @param operator The operator address to query.
     * @param operatorSet The `OperatorSet` to query.
     */
    function isMember(address operator, OperatorSet memory operatorSet) external view returns (bool);

    /**
     * @notice Returns whether or not an operator is slashable for an operator set.
     * @param operator The operator address to query.
     * @param operatorSet The `OperatorSet` to query.ß
     */
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) external view returns (bool);

    /**
     * @notice Returns whether or not an operator is registered to all provided operator sets.
     * @param operatorSets The list of operator sets to check.
     */
    function isOperatorSetBatch(
        OperatorSet[] calldata operatorSets
    ) external view returns (bool);

    function operatorSetStatus(
        address avs,
        address operator,
        uint32 operatorSetId
    ) external view returns (bool registered, uint32 lastDeregisteredTimestamp);

    /**
     * @notice Calculates the digest hash to be signed by an operator to register with an operator set.
     *
     * @param avs The AVS that operator is registering to operator sets for.
     * @param operatorSetIds An array of operator set IDs the operator is registering to.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateOperatorSetRegistrationDigestHash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32);

    /**
     * @notice Calculates the digest hash to be signed by an operator to force deregister from an operator set.
     *
     * @param avs The AVS that operator is deregistering from.
     * @param operatorSetIds An array of operator set IDs the operator is deregistering from.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateOperatorSetForceDeregistrationTypehash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32);

    /// @notice The EIP-712 typehash for the OperatorSetRegistration struct used by the contract.
    function OPERATOR_SET_REGISTRATION_TYPEHASH() external view returns (bytes32);
}
