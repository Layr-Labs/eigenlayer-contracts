# Slashing Parameter Configuration

Allocation functionality is explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).

```solidity
interface IAVSDirectory {
    /// STRUCTS

    /**
     * @notice this struct is used in queueAllocation, queueDeallocation in order to specify
     * an operator's slashability for a certain operator set
     *
     * @param strategy strategy to adjust allocations/deallocations for
     * @param operatorSets the operator sets to adjust magnitudes for
     * @param magnitudeDiff magnitude difference; the difference in proportional parts of the operator's slashable stake
     * that the operator set is getting. This struct is used either in allocating or deallocating
     * slashable stake to an operator set. Slashable stake for an operator set is
     * (slashableMagnitude / sum of all slashableMagnitudes for the strategy/operator + nonSlashableMagnitude) of
     * an operator's delegated stake.
     */
    struct MagnitudeAdjustment {
        IStrategy strategy;
        OperatorSet[] operatorSets;
        uint64[] magnitudeDiffs;
    }

    /**
     * @notice struct used for queued deallocations. Hash of struct is set in storage to be referenced later
     * when completing deallocations.
     * @param operator address for performing deallocations for
     * @param nonce
     */
    struct QueuedDeallocation {
        address operator;
        uint16 nonce;
        uint64 queuedTotalMagnitude;
        IStrategy strategy;
        OperatorSet[] operatorSets;
        uint64[] deallocationAmounts;
    }

    /// EVENTS

    event MagnitudeUpdated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint32 effectTimestamp,
        uint64 slashableMagnitude
    );

    event TotalMagnitudeUpdated(
        address operator,
        IStrategy strategy,
        uint64 nonslashableMagnitude
    );

    event NonslashableMagnitudeUpdated(
        address operator,
        IStrategy strategy,
        uint32 effectTimestamp,
        uint64 nonslashableMagnitude
    )

    /// EXTERNAL - STATE MODIFYING
    
    /**
     * @notice Allocate slashable magnitude for an operator.
     * For each strategy, takes an array of magnitude adjustments(increments)
     * and allocates them to their respective operatorSets. 
     * Nonslashable magnitude for the strategy will decrement by the sum of all 
     * allocations and these allocation increases will take effect 21 days(in blocks) from now.
     *
     * @param operator address to increase allocations for
     * @param allocations array of magnitude adjustments for multiple strategies and
     * corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     */
    function allocate(
        address operator,
        MagnitudeAdjustment[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external;

    /**
     * @notice Queue deallocate slashable magnitude for an operator.
     * For each strategy, takes an array of magnitude adjustments(increments)
     * and deallocates them to their respective operatorSets. These deallocations are queued
     * and can be completable 21 days(in blocks) from now.
     * NOTE: queued deallocations are still subject to slashing until completed.
     *
     * @param operator address to increase allocations for
     * @param deallocations array of magnitude adjustments for multiple strategies and
     * corresponding operator sets
     * @param operatorSignature signature of the operator if msg.sender is not the operator
     * @return queuedDeallocations the queued deallocation structs used as inputs to completeDeallocation
     */
    function queueDeallocation(
        address operator,
        MagnitudeAdjustment[] calldata deallocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external returns (QueuedDeallocation[]);

    /**
     * @notice Complete queued deallocations of magnitude for an operator.
     * Deallocates magnitude from an operatorSet that currently has allocations.
     * Decrementing the operatorSet magnitude and incrementing the nonslashable magnitude as a result.
     * NOTE: that the nonslashable amounts from deallocationAmounts may actually be less than expected due to slashing
     *
     * @param queuedDeallocations deallocations that were queued and are to be completed
     */
    function completeDeallocation(
        QueuedDeallocation[] calldata queuedDeallocations
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

Operators call this to allocate to their slashable stake(magnitudes) for a given (operator, IStrategy, operatorSet(avs, operatorSetId)) tuple. For each strategy param given, it queues magnitude updates to the specified operatorSets which will take effect 21 days from the time of calling. This gives stakers 3.5 days to queue withdrawals if they decide

All allocations in the call summed with all pending allocations must be less than the operator's current nonslashable magnitude in the strategy in order for all pending allocations to be backed stake that will be slashable when the update takes effect.

The function can be called with an EIP1271 signature from the operator or by the operator itself.

Emits

1. `MagnitudeUpdated` for each updated (operator, IStrategy, operatorSet)
2. `NonslashableMagnitudeUpdated` for each Strategy

Reverts if

1. The `operatorSignature` is invalid or `msg.sender` is not the `operator`
2. The magnitude queued for allocation to operatorSets summed with all pending allocations is greater than the `operator`'s latest pending nonslashable magnitude for the `strategy` (cannot allocate more than nonslashable)

### queueDeallocation


Operators call this to deallocate from their slashable stake(magnitudes) for a given (operator, IStrategy, operatorSet(avs, operatorSetId)) tuple.
A queued deallocation from a (operator, IStrategy, operatorSet(avs, operatorSetId)) tuple is bounded by the latest pending slashable magnitude defined, as one cannot deallocate more than what is going to be allocated.
These deallocations must be completed in a 2-tx step process by `queueDeallocation` and `comepleteDeallocation`. Queued deallocations are only completable after 21 days and are still subject to slashing from the operatorSet
while queued.

The function can be called with an EIP1271 signature from the operator or by the operator itself.

Emits

1. `MagnitudeUpdated` for each updated (operator, IStrategy, operatorSet)

Reverts if

1. The `operatorSignature` is invalid or `msg.sender` is not the `operator`
2. The magnitude queued for deallocation from an operatorSet is greater than the `operator`'s latest pending magnitude for the `strategy` (cannot deallocate more than is allocated)

### completeDeallocation

Operators call this to complete their queued deallocations after they have been queued for at least 21 days. Deallocations must all be completed in the order they were queued (tracked by an incrementing nonce).
Deallocation amounts may increment nonslashable magnitude upon completion less than expected due to slashing events while they were queued. This is done by keeping track of the totalMagnitude
at time of queuing and the totalMagnitude at the time of deallocation completion.

This function can be called permissionlessly by anyone once it is completable.

Emits

1. `NonslashableMagnitudeUpdated` for each updated (operator, IStrategy)

Reverts if

1. Any of the queued deallocations have not been queued for long enough
2. Queued deallocations are not provided sequentially and in order of the last completed nonce.
3. Any of the hashes of the queued deallocation inputs were not set in storage i.e were not actually previously queued


### getSlashableBips

For a given timestamp,

$slashableBips_{op,opset,str} = 
\left\lfloor\frac{magnitude_{op,opset,str}}{ nonslashableMagnitude_{op,str} + \sum_{opset_i \in opsets} magnitude_{op,opset_i,str}}*10000\right\rfloor
$

This explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).
