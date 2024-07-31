# Slashing Parameter Configuration

Allocator functionality is explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).

```solidity
interface IAVSDirectory {
    /// STRUCTS

    /**
     * @notice this struct is used in queueAllocation, queueDeallocation in order to specify
     * an operator's slashability for a certain operator set
     *
     * @param operatorSet the operator set to change slashing parameters for
     * @param magnitudeDiff magnitude difference; the difference in proportional parts of the operator's slashable stake
     * that the operator set is getting. This struct is used either in allocating or deallocating
     * slashable stake to an operator set. Slashable stake for an operator set is
     * (slashableMagnitude / sum of all slashableMagnitudes for the strategy/operator + nonSlashableMagnitude) of
     * an operator's delegated stake.
     */
    struct MagnitudeAdjustment {
        OperatorSet operatorSet;
        uint64 magnitudeDiff;
    }

    /**
     * @notice struct used for queued deallocations. Hash of struct is set to true in a mapping
     */
    struct QueuedDeallocation {
        address operator;
        uint16 nonce;
        IStrategy strategy;
        OperatorSet operatorSet;
        uint64 deallocationAmount;
        uint64 queuedTotalMagnitude;
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
    
    function queueAllocation(
        address operator,
        IStrategy strategy,
        MagnitudeAdjustment[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external;

    function queueDeallocation(
        address operator,
        IStrategy strategy,
        MagnitudeAdjustment[] calldata deallocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external returns (QueuedDeallocation[]);

    function completeDeallocation(
        QueuedDeallocation[] calldata deallocations
    ) external;

    function slashOperator(
        address operator,
        OperatorSet operatorSet,
        IStrategy[] calldata strategies,
        uint16 bipsToSlash
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

### queueReallocation

Allocators call this to queue allocation/deallocation to their slashable stake(magnitudes) for a given (operator, IStrategy, operatorSet(avs, operatorSetId)) tuple. For each strategy param given, it queues magnitude updates to the specified operatorSets which will take effect 21 days from the time of calling. This gives stakers 3.5 days to queue withdrawals if they decide

Whether or not an adjustment update is an allocation or deallocation is defined by the `MagnitudeAdjustmentType` enum. For implementation complexity reasons, this is set at the per Strategy level.

All allocations in the call summed with all pending allocations must be less than the operator's current nonslashable magnitude in the strategy in order for all pending allocations to be backed stake that will be slashable when the update takes effect.

A deallocation from a (operator, IStrategy, operatorSet(avs, operatorSetId)) tuple is bounded by the latest pending slashable magnitude defined, as one cannot deallocate more than what is going to be allocated.

The function can be called with an EIP1271 signature from the operator's allocator or by the allocator itself.

Emits

1. `MagnitudeUpdated` for each updated (operator, IStrategy, operatorSet)
2. `TotalAndNonslashableMagnitudeUpdated` for each Strategy

Reverts if

1. The `allocatorSignature` is invalid or `msg.sender` is not the `operator`'s allocator
2. if MagnitudeAdjustmentType == ALLOCATION:
    - The magnitude queued for allocation to operatorSets is greater than the `operator`'s latest pending nonslashable magnitude for the `strategy` (cannot allocate more than nonslashable)
3. if MagnitudeAdjustmentType == DEALLOCATION:
    - The magnitude queued for deallocation from an operatorSet is greater than the `operator`'s latest pending magnitude for the `strategy` (cannot deallocate more than is allocated)
4. There are > 3 pending magnitude updates for any of the operatorSets being updated for the given `operator` and strategy
    1. This is due to gas constraints during slashing

### queueMagnitudeDilution

Allocators call this and for each given strategy decreases all relative proportions of currently allocated magnitudes allowing the allocator to reduce the total slashable proportion. This is done cheaply by increasing their non slashable magnitude to the desired proportion of the new total magnitude. The function can be called with an EIP1271 signature from the operator's allocator or by the allocator itself.

For example, to add 1% to an operator's nonslashable proportion for a given strategy, if the operator's total magnitude for the strategy was 100e18, their allocator would queue a dilution of 1.0101...e18 magnitude. This would result in their nonslashable magnitude being increased by 1% of the total magnitude.

Dilution reduces allocations from all existing operatorSets for the operator and strategy.

Note that, upon slashing, magnitudes for dilution will be decremented to preserve the proportion being diluted because the total magnitude decrements during slashing.

Emits

1. `TotalAndNonslashableMagnitudeUpdated` for each Strategy

Reverts if:

1. The `allocatorSignature` is invalid or `msg.sender` is not the `operator`'s allocator
2. The total magnitude after dilution is greater than `FIXED_TOTAL_MAG_LIMIT`
3. `strategies.length != nonslashableToAdd.length`
4. There are >6 pending updates to the nonslashable magnitude for the given `operator` and strategy
    1. This is due to gas constraints during slashing

### queueMagnitudeConcentration

Allocators call this and for each given strategy increases all relative proportions of currently allocated magnitudes allowing the allocator to increase their total slashable proportion. This is done cheaply by decreasing their non slashable magnitude to the desired proportion of the new total magnitude. The function can be called with an EIP1271 signature from the operator's allocator or by the allocator itself.

For example, to remove 1% from an operator's nonslashable proportion for a given strategy, if the operator's total magnitude for the strategy was 100e18, their allocator would queue a contration of 0.990099e18 magnitude. This would result in their nonslashable magnitude being decreased by 1% of the total magnitude.

Concentration increases allocations from all existing operatorSets for the operator and strategy.

Emits

1. `TotalAndNonSlashableMagnitudeUpdated` for each updated (operator, IStrategy)


Reverts if

1. The `allocatorSignature` is invalid or `msg.sender` is not the `operator`'s allocator
2. The allocator is attempting to remove more nonslashable magnitude than the `operator`'s latest pending nonslashable magnitude for the `strategy`
3. `strategies.length != nonslashableToAdd.length`
4. There are >6 pending updates to the nonslashable magnitude for the given `operator` and strategy
    1. This is due to gas constraints during slashing

### getSlashableBips

For a given timestamp,

$slashableBips_{op,opset,str} = 
\left\lfloor\frac{magnitude_{op,opset,str}}{ nonslashableMagnitude_{op,str} + \sum_{opset_i \in opsets} magnitude_{op,opset_i,str}}*10000\right\rfloor
$

This explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).
