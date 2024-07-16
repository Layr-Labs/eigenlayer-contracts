# Slashing Parameter Configuration

Allocator functionality is explained [here](https://www.notion.so/eigen-labs/Allocator-Functionality-282a008ab7a14c79a25ec2954f8f5912).

```solidity
interface IAVSDirectory {
    /// STRUCTS

    enum MagnitudeAdjustmentType {
        ALLOCATION,
        DEALLOCATION
    }

    /**
     * @notice this struct is used in MagnitudeAdjustmentsParam in order to specify
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
     * @notice Input param struct used for functions queueAllocation and queueDeallocation.
     * A structure defining a set of operator-based slashing configurations
     * to manage slashable stake.
     * @param strategy the strategy to update magnitudes for
     * @param magnitudeAdjustmentType the type of adjustment to make to the operator's slashable stake
     * This is set at the per Strategy level vs per OperatorSet to make TotalAndNonslashableUpdate(s)
     * easier to update. Allocations and deallocations for the same strategy could be performed by having
     * an additional array index when calling `queueReallocation`
     * @param magnitudeAdjustments the input magnitude parameters defining the adjustments to the proportion
     * the operator set is able to slash an operator. This is passed as an array of
     * operator sets, magnitudes diffs, MagnitudeAdjustmentType enum for a given strategy
     */
    struct MagnitudeAdjustmentsParam {
        IStrategy strategy;
        MagnitudeAdjustmentType magnitudeAdjustmentType;
        MagnitudeAdjustment[] magnitudeAdjustments;
    }

    /**
     * @notice Used for historical magnitude updates in mapping
     * operator => IStrategy => avs => operatorSetId => MagnitudeUpdate[]
     * New updates are pushed whenever queueDeallocations or queueAllocations is called
     * @param timestamp timestamp of MagnitudeUpdate, if timestamp > block.timestamp then it is currently pending
     * @param magnitude the magnitude/proportion value of the (operator, Strategy, operatorSet) at timestamp
     */
    struct MagnitudeUpdate {
        uint32 timestamp;
        uint64 magnitude;
    }

    /**
     * @notice Used for historical magnitude updates in mapping
     * operator => IStrategy => TotalMagnitudeUpdate[]
     * New total magnitude updates are pushed whenever magnitude changing functions are called
     * @param timestamp timestamp of TotalAndNonslashableUpdate, if timestamp > block.timestamp then it is currently pending
     * @param totalMagnitude total magnitude amount for a strategy which equals
     * nonslashableMagnitude amount + sum of each operatorSet's allocated magnitudes
     * @param nonslashableMagnitude nonslashable magnitude that CANNOT be slashed if timestamp <= block.timestamp
     * Note if timestamp > block.timestamp, this is an upper bound on the slashable amount
     * as it may still be a pending deallocation (still slashable)
     * @param cumulativeAllocationSum monotonically increasing sum of all magnitudes when allocating
     * required to ensure all allocations are backed by nonslashable magnitude
     */
    struct TotalAndNonslashableUpdate {
        uint32 timestamp;
        uint64 totalMagnitude;
        uint64 nonslashableMagnitude;
        uint64 cumulativeAllocationSum;
    }

    /// EVENTS

    event MagnitudeUpdated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint32 effectTimestamp,
        uint64 slashableMagnitude
    );

    event TotalAndNonslashableMagnitudeUpdated(
        address operator,
        IStrategy strategy,
        uint32 effectTimestamp,
        uint64 nonslashableMagnitude,
        uint64 totalSlashableMagnitude
    );

    /// EXTERNAL - STATE MODIFYING

    /**
     * @notice Queues magnitude adjustment updates of type ALLOCATION or DEALLOCATION
     * The magnitude allocation takes 21 days from time when it is queued to take effect.
     *
     * @param operator the operator whom the magnitude parameters are being adjusted
     * @param allocationParams allocation adjustment params with differences to add/subtract to current magnitudes
     * @param allocatorSignature if non-empty is the signature of the allocator on
     * the modification. if empty, the msg.sender must be the allocator for the
     * operator
     *
     * @dev pushes a MagnitudeUpdate and TotalAndNonslashableUpdate to take effect in 21 days
     * @dev If ALLOCATION adjustment type:
     * reverts if sum of magnitudeDiffs > nonslashable magnitude for the latest pending update - sum of all pending allocations
     * since one cannot allocate more than is nonslahsable
     * @dev if DEALLOCATION adjustment type:
     * reverts if magnitudeDiff > allocated magnitude for the latest pending update
     * since one cannot deallocate more than is already allocated
     * @dev reverts if there are more than 3 pending allocations/deallocations for the given (operator, strategy, operatorSet) tuple
     * @dev emits events MagnitudeUpdated, TotalAndNonslashableMagnitudeUpdated
     */
    function queueReallocation(
        address operator,
        MagnitudeAdjustmentsParam[] calldata allocationParams,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) external returns (uint32 effectTimestamp);

    /**
     * @notice For each strategy, add to nonslashable magnitude to dilute all operatorSet slashable stake allocations
     * proportionally. This is an efficient way of adjusting total magnitude and slashing risk.
     *
     * @param operator the address that will have their nonslashable magnitude added to, diluting
     * relative proportions of all existing operatorSet magnitudes
     * @param strategies array of IStrategy's to dilute
     * @param nonslashableToAdd array values of the magnitude amount added to nonslashable magnitude
     * @param allocatorSignature if non-empty is the signature of the allocator on
     * the modification. if empty, the msg.sender must be the allocator for the
     * operator
     *
     * @dev reverts if nonslashableToAdd > FIXED_TOTAL_MAG_LIMIT
     * @dev reverts if there are more than 6 pending changes to changes to the nonslashable magnitude for the strategy
     * @dev reverts if strategies.length != nonslashableToAdd.length
     * @dev emits event TotalAndNonSlashableMagnitudeUpdated
     */
    function queueMagnitudeDilution(
        address operator,
        IStrategy[] calldata strategies,
        uint64[] calldata nonslashableToAdd,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) external returns (uint64 newNonslashableMagnitude, uint64 newTotalMagnitude);

    /**
     * @notice Decrement the nonslashable magnitude to increase all operatorSet slashable stake allocations
     * proportionally. This is an efficient way of adjusting total magnitude and slashing risk.
     *
     * @param operator the address that will have their nonslashable magnitude decremented, increasing
     * relative proportions of all existing operatorSet magnitudes
     * @param strategies array of IStrategy's to dilute
     * @param nonslashableToDecrement array values of the magnitude amount decremented from the nonslashable magnitude
     * @param allocatorSignature if non-empty is the signature of the allocator on
     * the modification. if empty, the msg.sender must be the allocator for the
     * operator
     *
     * @dev reverts if nonslashableToDecrement > latest pending nonslashableMagnitude
     * @dev reverts if there are more than 6 pending changes to changes to the nonslashable magnitude for the strategy
     * @dev reverts if strategies.length != nonslashableToDecrement.length
     * @dev emits event TotalAndNonSlashableMagnitudeUpdated
     */
    function queueMagnitudeConcentration(
        address operator,
        IStrategy[] calldata strategies,
        uint64[] calldata nonslashableToDecrement,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) external returns (uint64 newNonslashableMagnitude, uint64 newTotalMagnitude);

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
