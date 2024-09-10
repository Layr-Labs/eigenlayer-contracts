// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "./IAVSDirectory.sol";
import "./IStrategy.sol";
import "./ISignatureUtils.sol";

interface IAllocationManager is ISignatureUtils {
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
    /// @dev Thrown when provided operator sets are not in ascending order.
    error OperatorSetsNotInAscendingOrder();
    /// @dev Thrown when an allocation is attempted for a given operator when they have pending allocations or deallocations.
    error PendingAllocationOrDeallocation();
    /// @dev Thrown when an allocation is attempted that exceeds a given operators total allocatable magnitude.
    error InsufficientAllocatableMagnitude();
    /// @dev Thrown when attempting to use an expired eip-712 signature.
    error SignatureExpired();
    /// @dev Thrown when attempting to spend a spent eip-712 salt.
    error SaltSpent();

    /**
     * @notice struct used to modify the allocation of slashable magnitude to list of operatorSets
     * @param strategy the strategy to allocate magnitude for
     * @param expectedTotalMagnitude the expected total magnitude of the operator used to combat against race conditions with slashing
     * @param operatorSets the operatorSets to allocate magnitude for
     * @param magnitudes the magnitudes to allocate for each operatorSet
     */
    struct MagnitudeAllocation {
        IStrategy strategy;
        uint64 expectedTotalMagnitude;
        OperatorSet[] operatorSets;
        uint64[] magnitudes;
    }

    /**
     * @notice struct used for pending free magnitude. Stored in (operator, strategy, operatorSet) mapping
     * to be used in completeDeallocations.
     * @param magnitudeDiff the amount of magnitude to deallocate
     * @param completableTimestamp the timestamp at which the deallocation can be completed, 21 days from when queued
     */
    struct PendingFreeMagnitude {
        uint64 magnitudeDiff;
        uint32 completableTimestamp;
    }

    /**
     * @notice Struct containing info regarding free allocatable magnitude.
     * @param nextPendingFreeMagnitudeIndex The next available update index.
     * @param freeMagnitude The total amount of free allocatable magnitude.
     */
    struct OperatorMagnitudeInfo {
        uint192 nextPendingFreeMagnitudeIndex;
        uint64 freeMagnitude;
    }

    /// @notice Emitted when an operator set is created by an AVS.
    event OperatorSetCreated(OperatorSet operatorSet);

    /// @notice Emitted when an operator allocates slashable magnitude to an operator set
    event MagnitudeAllocated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeToAllocate,
        uint32 effectTimestamp
    );

    /// @notice Emitted when an operator queues deallocations of slashable magnitude from an operator set
    event MagnitudeQueueDeallocated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeToDeallocate,
        uint32 completableTimestamp
    );

    /// @notice Emitted when an operator completes deallocations of slashable magnitude from an operator set
    /// and adds back magnitude to free allocatable magnitude
    event MagnitudeDeallocationCompleted(
        address operator, IStrategy strategy, OperatorSet operatorSet, uint64 freeMagnitudeAdded
    );

    /// @notice Emitted when an operator is slashed by an operator set for a strategy
    event OperatorSlashed(address operator, uint32 operatorSetId, IStrategy strategy, uint16 bipsToSlash);

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Modifies the propotions of slashable stake allocated to a list of operatorSets for a set of strategies
     * @param operator address to modify allocations for
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     * @dev updates freeMagnitude for the updated strategies
     * @dev must be called by the operator
     */
    function modifyAllocations(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external;

    /**
     * @notice For all pending deallocations that have become completable, their pending free magnitude can be
     * added back to the free magnitude of the (operator, strategy) amount. This function takes a list of strategies
     * and adds all completable deallocations for each strategy, updating the freeMagnitudes of the operator
     *
     * @param operator address to complete deallocations for
     * @param strategies a list of strategies to complete deallocations for
     *
     * @dev can be called permissionlessly by anyone
     */
    function updateFreeMagnitude(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToComplete
    ) external;

    /**
     * @notice Called by an AVS to slash an operator for given operatorSetId, list of strategies, and bipsToSlash.
     * For each given (operator, operatorSetId, strategy) tuple, bipsToSlash
     * bips of the operatorSet's slashable stake allocation will be slashed
     *
     * @param operator the address to slash
     * @param operatorSetId the ID of the operatorSet the operator is being slashed on behalf of
     * @param strategies the set of strategies to slash
     * @param bipsToSlash the number of bips to slash, this will be proportional to the
     * operator's slashable stake allocation for the operatorSet
     */
    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint16 bipsToSlash
    ) external;

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(
        bytes32 salt
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Get the allocatable magnitude for an operator and strategy based on number of pending deallocations
     * that could be completed at the same time. This is the sum of freeMagnitude and the sum of all pending completable deallocations.
     * @param operator the operator to get the allocatable magnitude for
     * @param strategy the strategy to get the allocatable magnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete, 0 to complete all (uint8 max 256)
     */
    function getAllocatableMagnitude(
        address operator,
        IStrategy strategy,
        uint16 numToComplete
    ) external view returns (uint64);

    /**
     * @notice operator is slashable by operatorSet if currently registered OR last deregistered within 21 days
     * @param operator the operator to check slashability for
     * @param operatorSet the operatorSet to check slashability for
     * @return bool if the operator is slashable by the operatorSet
     */
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) external view returns (bool);

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     *
     * @return operatorSets the operator sets the operator is a member of and the current slashable magnitudes for each strategy
     */
    function getCurrentSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (OperatorSet[] memory, uint64[][] memory);

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     * @param timestamp the timestamp to get the slashable magnitude for
     *
     * @return operatorSets the operator sets the operator is a member of and the slashable magnitudes for each strategy
     */
    function getSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (OperatorSet[] memory, uint64[][] memory);

    /**
     * @param operator the operator to get the slashable ppm for
     * @param operatorSet the operatorSet to get the slashable ppm for
     * @param strategies the strategies to get the slashable ppm for
     * @param timestamp the timestamp to get the slashable ppm for for
     * @param linear whether the search should be linear (from the most recent) or binary
     *
     * @return slashablePPM the slashable ppm of the given list of strategies allocated to
     * the given OperatorSet for the given operator and timestamp
     */
    function getSlashablePPM(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies,
        uint32 timestamp,
        bool linear
    ) external view returns (uint24[] memory);

    /**
     * @notice Returns the current total magnitudes of an operator for a given set of strategies
     * @param operator the operator to get the total magnitude for
     * @param strategies the strategies to get the total magnitudes for
     * @return totalMagnitudes the total magnitudes for each strategy
     */
    function getTotalMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory);

    /**
     * @notice Returns the total magnitudes of an operator for a given set of strategies at a given timestamp
     * @param operator the operator to get the total magnitude for
     * @param strategies the strategies to get the total magnitudes for
     * @param timestamp the timestamp to get the total magnitudes at
     * @return totalMagnitudes the total magnitudes for each strategy
     */
    function getTotalMagnitudesAtTimestamp(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint64[] memory);

    /**
     * @notice Returns the current total magnitude of an operator for a given strategy
     * @param operator the operator to get the total magnitude for
     * @param strategy the strategy to get the total magnitude for
     * @return totalMagnitude the total magnitude for the strategy
     */
    function getTotalMagnitude(address operator, IStrategy strategy) external view returns (uint64);

    /**
     * @notice Returns the total magnitude of an operator for a given strategy at a given timestamp
     * @param operator the operator to get the total magnitude for
     * @param strategy the strategy to get the total magnitude for
     * @param timestamp the timestamp to get the total magnitude at
     * @return totalMagnitude the total magnitude for the strategy
     */
    function getTotalMagnitudeAtTimestamp(
        address operator,
        IStrategy strategy,
        uint32 timestamp
    ) external view returns (uint64);

    /**
     * @notice Calculates the digest hash to be signed by an operator to modify magnitude allocations
     * @param operator The operator to allocate or deallocate magnitude for.
     * @param allocations The magnitude allocations/deallocations to be made.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateMagnitudeAllocationDigestHash(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32);

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() external view returns (bytes32);
}
