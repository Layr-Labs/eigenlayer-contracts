# Slashing Parameter Configuration

Allocation functionality and magnitudes are explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).

```solidity
interface IAVSDirectory {
    /// STRUCTS

    /**
     * @notice this struct is used in allocate and queueDeallocation in order to specify an operator's slashability for a certain operator set
     *
     * @param strategy the strategy to adjust slashable stake for
     * @param operatorSets the operator sets to adjust slashable stake for
     * @param magnitudeDiff magnitude difference; the difference in proportional parts of the operator's slashable stake
     * that is slashable by the operatorSet.
     * Slashable stake for an operator set is (magnitude / sum of all magnitudes for the strategy/operator + nonSlashableMagnitude) of
     * an operator's delegated stake.
     */
    struct MagnitudeAdjustment {
        IStrategy strategy;
        OperatorSet[] operatorSets;
        uint64[] magnitudeDiffs;
    }

    /// EVENTS

    event MagnitudeAllocation(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeToAllocate,
        uint32 effectTimestamp,
    );

    event MagnitudeDeallocationQueued(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeToDeallocate,
        uint32 effectTimestamp,
    );

    event MagnitudeDeallocationCompleted(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint64 magnitudeCompletedDeallocation
    );

    /// EXTERNAL - STATE MODIFYING
    
    /**
     * @notice Queues a set of magnitude adjustments to increase the slashable stake of an operator set for the given operator for the given strategy.
     * Nonslashable magnitude for each strategy will decrement by the sum of all 
     * allocations for that strategy and the allocations will take effect 21 days from calling.
     *
     * @param operator address to increase allocations for
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     */
    function allocate(
        address operator,
        MagnitudeAdjustment[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external;

    /**
     * @notice Queues a set of magnitude adjustments to decrease the slashable stake of an operator set for the given operator for the given strategy.
     * The deallocations will take effect 21 days from calling. In order for the operator to have their nonslashable magnitude increased, 
     * they must call the contract again to complete the deallocation. Stake deallocations are still subject to slashing until 21 days have passed since queuing.
     *
     * @param operator address to decrease allocations for
     * @param deallocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     */
    function queueDeallocation(
        address operator,
        MagnitudeAdjustment[] calldata deallocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external;

    /**
     * @notice Complete queued deallocations of slashable stake for an operator.
     * Increments the nonslashable magnitude of the operator by the sum of all deallocation amounts for each strategy. 
     * If the operator was slashed, this will be a smaller amount than during queuing.
     *
     * @param operator address to complete deallocations for
     * @param operatorSets the sets of operator sets to complete deallocations for
     * @param strategies a 2d list of strategies to complete deallocations for, one list for each operatorSet
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     *
     * @dev can be called permissionlessly by anyone
     */
    function completeDeallocations(
        address operator,
        OperatorSet[] calldata operatorSets,
        IStrategy[][] calldata strategies,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external;

    /// VIEW

    /**
     * @param operator the operator to get the slashable bips for
     * @param operatorSet the operatorSet to get the slashable bips for
     * @param strategy the strategy to get the slashable bips for
     * @param timestamp the timestamp to get the slashable bips for for
     *
     * @return slashableBips the slashable bips of the given strategy owned by
     * the given OperatorSet for the given operator and timestamp
     */
    function getSlashableBips(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy strategy,
        uint32 timestamp
    ) external returns (uint16 slashableBips);
}
```

### allocate

Operators call this to allocate to their slashable stake (magnitudes) for a list of (operator, IStrategy, operatorSet(avs, operatorSetId)) tuples. For each adjustment param given, it queues magnitude updates to the specified operatorSets which will take effect 21 days from the time of calling. This gives the operator's stakers 3.5 days to queue withdrawals if they disagree with the changes to their staking portfolio.

All allocations in the call are summed and checked to be less than the nonslashable magnitude available for allocation for the `strategy`. This is in order for all allocations to be backed stake that will be slashable when the update takes effect.

The function can be called with an EIP1271 signature from the operator or by the operator itself.

Emits

1. `MagnitudeAllocation` for each updated (operator, IStrategy, operatorSet)

Reverts if

1. The `operatorSignature` is invalid or `msg.sender` is not the `operator`
2. The magnitude queued for allocation to operatorSets summed with all pending allocations for the `(operator, strategy)` pair is greater less than the nonslashable magnitude for the `strategy` (cannot allocate more than nonslashable)

### queueDeallocation

Operators call this to deallocate from their slashable stake (magnitudes) for a list of (operator, IStrategy, operatorSet(avs, operatorSetId)) tuples. Queued deallocations are no longer slashable after 21 days from the time of queueing.
These deallocations must be completed in a 2-tx step process by calling `completeDeallocations` after 21 days have passed which increments the nonslashable magnitude of the operator, to be allocated to different operatorSets.

The sum of all pending deallocations from a (operator, IStrategy, operatorSet(avs, operatorSetId)) tuple are verified to be less than the currently allocated magnitude. This is in order to make sure that the operator cannot deallocate more than they have allocated.

The function can be called with an EIP1271 signature from the operator or by the operator itself.

Emits

1. `MagnitudeDeallocationQueued` for each updated (operator, IStrategy, operatorSet)

Reverts if

1. The `operatorSignature` is invalid or `msg.sender` is not the `operator`
2. The sum of all pending deallocations from an operatorSet is greater than the `operator`'s current magnitude for the operatorSet and `strategy` (cannot deallocate more than is allocated)

### completeDeallocations

Operators call this to complete their queued deallocations for a list of (operator, IStrategy, operatorSet(avs, operatorSetId)) tuples. This increments the nonslashable magnitude of the operator by the sum of all completable deallocations.

Deallocations may increment nonslashable magnitude upon completion less than expected due to slashing events while they were queued. See [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912) for more information.

This function can be called permissionlessly by anyone.

Emits

1. `MagnitudeDeallocationCompleted` for each updated (operator, IStrategy, operatorSet)

Reverts if:

1. The `operatorSignature` is invalid or `msg.sender` is not the `operator`
2. The number of `operatorSets` is not the same as the number of `strategies` lists
3. There are duplicate strategies in any of the `strategies` lists

### getSlashableBips

For a given timestamp,

$slashableBips_{op,opset,str} = 
\left\lfloor\frac{magnitude_{op,opset,str}}{ nonslashableMagnitude_{op,str} + \sum_{opset_i \in opsets} magnitude_{op,opset_i,str}}*10000\right\rfloor
$

This explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).
