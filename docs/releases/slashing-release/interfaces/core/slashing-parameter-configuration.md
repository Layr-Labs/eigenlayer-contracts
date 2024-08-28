# Slashing Parameter Configuration

Allocation functionality and magnitudes are explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).

```solidity
interface IAVSDirectory {
    /// STRUCTS

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
     * @param completableTimestamp the timestamp at which the deallocation can be completed, 17.5 days from when queued
     */
    struct PendingFreeMagnitude {
        uint64 magnitudeDiff;
        uint32 completableTimestamp;
    }

    /// EVENTS

    event MagnitudeAllocated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeToAllocate,
        uint32 effectTimestamp,
    );

    event MagnitudeDeallocated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeToDeallocate,
        uint32 completeableTimestamp,
    );


    /// EXTERNAL - STATE MODIFYING
    
    /**
     * @notice Modifies the propotions of slashable stake allocated to a list of operatorSets for a set of strategies
     * This will revert if the operator has never set their allocationDelay before.
     * @param operator address to modify allocations for
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     * @dev Updates freeMagnitude for the updated strategies
     * @dev Must be called by the operator or with a valid operator signature
     * @dev For each allocation, allocation.operatorSets MUST be ordered in ascending order according to the
     * encoding of the operatorSet. This is to prevent duplicate operatorSets being passed in. The easiest way to ensure
     * ordering is to sort allocated operatorSets by address first, and then sort for each avs by ascending operatorSetIds.
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
        uint8[] calldata numToComplete
    ) external;

    /**
     * @notice Called by operators to set their allocation delay. Can only have 1 pending update
     * to their allocationDelay. If operator's have never set their allocationDelay before then they cannot
     * call modifyAllocations.
     * @param delay the allocation delay in seconds
     */
    function setAllocationDelay(uint32 delay) external;

    /// VIEW

    /**
     * @param operator the operator to get the slashableProportion for
     * @param operatorSet the operatorSet to get the slashableProportion for
     * @param strategy the strategy to get the slashableProportion for
     * @param timestamp the timestamp to get the slashableProportion for for
     *
     * @return slashableProportion the slashable proportion in parts per 1e18 of the given strategy owned by
     * the given OperatorSet for the given operator and timestamp
     */
    function getSlashableProportion(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy strategy,
        uint32 timestamp
    ) external returns (uint64 slashableProportion);

    /**
     * @notice Get the allocatable magnitude for an operator and strategy based on number of pending deallocations
     * that could be completed at the same time. This is the sum of freeMagnitude and the sum of all pending completable deallocations.
     * @param operator the operator to get the allocatable magnitude for
     * @param strategy the strategy to get the allocatable magnitude for
     * @param numToComplete the number of pending free magnitudes deallocations to complete
     */
    function getAllocatableMagnitude(
        address operator,
        IStrategy strategy,
        uint16 numToComplete
    ) external view returns (uint64);

    /**
     * @notice Get the allocation delay (in seconds) for an operator. Can only be configured one-time
     * from calling initializeAllocationDelay.
     * @param operator the operator to get the allocation delay for
     * @param timestamp the timestamp to read the operator's allocationDelay
     */
    function getAllocationDelay(address operator, uint32 timestamp) external view returns (uint32);

    /**
     * @notice operator is slashable by operatorSet if currently registered OR last deregistered within 17.5 days
     * @param operator the operator to check slashability for
     * @param operatorSet the operatorSet to check slashability for
     * @return bool if the operator is slashable by the operatorSet
     */
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) external view returns (bool);

}
```

### modifyAllocations

Operators call this to set the proportions of slashable stake allocated to a list of operatorSets for a set of strategies. Depending on what the new magnitude/proportion and current magnitude value is configured to, this will either create a pending allocation or deallocation. Allocation delays can be configured by operators but setting a new value itself has a delay of 21 days. Deallocation delays all have a built-in delay of 17.5 days until they take effect. Deallocations are still susceptible to slashing by the operatorSet up until they take effect. Pending allocations on the other hand are not slashable (referring to the added increase in magnitude).

Note that allocation delays are more of a concern to stakers as it is increasing the risk to an operator's delegated stakers. Thus they may need enough time to undelegate and withdraw their stake from an operator if they do not agree with a pending allocation. However, some operators are also their own staker (ex. LRTs) and may not care about this allocation delay so they could configure it to be 0.

Deallocation delays are important for AVSs and their operatorSets to ensure they can have a minimum future forecast of slashable stake so that all of their economic security cannot leave their system instantly. 17.5 days gives AVSs enough time to read updates made to slashable stakes and for the existing tasks they have created to have their slashing windows expire.

Operators call this to allocate to their slashable stake (magnitudes) for a list of (operator, IStrategy, operatorSet(avs, operatorSetId)) tuples. For each MagnitudeAllocation param given, it has a given IStrategy and it queues magnitude updates to the specified operatorSets which will take effect 21 days from the time of calling. This gives the operator's stakers 3.5 days to queue withdrawals if they disagree with the changes to their staking portfolio.

All allocations in the call are summed and checked to be less than the free magnitude available for allocation for the `strategy`. This free magnitude amount is nonslashable by an operatorSet. This is in order for all allocations to be backed stake that will be slashable when the update takes effect.

The function can be called with an EIP1271 signature from the operator or by the operator itself.

Emits

1. For each allocation input, \
        `MagnitudeAllocated` is emitted for the given (operator, IStrategy, operatorSet)
2. For each deallocation input, \
`MagnitudeDeallocated` is emitted for the given (operator, IStrategy, operatorSet)

Reverts if

1. The operator is not a registered operator in the DelegationManager
2. The `operatorSignature` is invalid AND `msg.sender` is not the `operator`
3. The operator has never set their allocationDelay before
4. For each IStrategy, The operator's totalMagnitude doesn't match `expectedTotalMagnitude`
5. For each `MagnitudeAllocation`: \
    1. allocation.operatorSets.length != allocation.magnitudes.length
    2. One of the provided operatorSets does not exist
    3. There exists a pending allocation or deallocation. Can only have 1 pending change
6. If the sum of all magnitude allocations for a IStrategy is greater than the free magnitude that is available to allocate.

### updateFreeMagnitude

For all pending deallocations that have become completable, their pending free magnitude can be
added back to the free magnitude of the (operator, IStrategy) amount. This is by default done whenever `modifyAllocations` is called but this a separate interface in case of gas limitations in a tx because there is no bound on the number of pending completable deallocations for a given IStrategy.

### setAllocationDelay

Operators can call this to set their allocationDelay. Operators will not be able to call `modifyAllocations` if they
have never set their allocation delay before.

Reverts if

1. msg.sender is not a operator
2. msg.sender already has a pending update to their delay

### getSlashableProportion

For a given timestamp,

$slashableProportion_{op,opset,str} = 
\left\lfloor\frac{magnitude_{op,opset,str}}{ nonslashableMagnitude_{op,str} + \sum_{opset_i \in opsets} magnitude_{op,opset_i,str}}*1e18\right\rfloor
$

This explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).

### getAllocatableMagnitude

Returns the available free magnitude that can be allocated for a given Strategy. In the underlying implementation, this will calculate the actual freeMagnitude amount in storage and add all completable deallocation amounts as well. 

### getAllocationDelay

Returns the current allocation delay for a operator at a given timestamp

### isOperatorSlashable

Checks that a operator is slashable for an operatorSet which is defined as the following being true: operator is currently registered OR operator deregistered within last 17.5 days