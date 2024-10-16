// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "./IAVSDirectory.sol";
import "./IPauserRegistry.sol";
import "./IStrategy.sol";
import "./ISignatureUtils.sol";

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
    /// @dev Thrown when an invalid strategy is provided.
    error InvalidStrategy();
    /// @dev Thrown when caller is not the delegation manager.
    error OnlyDelegationManager();
    /// @dev Thrown when an operator attempts to set their allocation for an operatorSet to the same value
    error SameMagnitude();
    /// @dev Thrown when an allocation is attempted for a given operator when they have pending allocations or deallocations.
    error ModificationAlreadyPending();
    /// @dev Thrown when an allocation is attempted that exceeds a given operators total allocatable magnitude.
    error InsufficientAllocatableMagnitude();
    /// @dev Thrown when attempting to use an expired eip-712 signature.
    error SignatureExpired();
    /// @dev Thrown when attempting to spend a spent eip-712 salt.
    error SaltSpent();
    /// @dev Thrown when attempting to slash an operator that has already been slashed at the given timestamp.
    error AlreadySlashedForTimestamp();
    /// @dev Thrown when calling a view function that requires a valid timestamp.
    error InvalidTimestamp();
    /// @dev Thrown when an invalid allocation delay is set
    error InvalidAllocationDelay();
    /// @dev Thrown when a slash is attempted on an operator who has not allocated to the strategy, operatorSet pair
    error OperatorNotAllocated();
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
     * @notice This function takes a list of strategies and adds all completable modifications for each strategy,
     * updating the encumberedMagnitude of the operator as needed.
     *
     * @param operator address to complete modifications for
     * @param strategies a list of strategies to complete modifications for
     * @param numToComplete a list of number of pending modifications to complete for each strategy
     *
     * @dev can be called permissionlessly by anyone
     */
    function clearModificationQueue(
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
}
