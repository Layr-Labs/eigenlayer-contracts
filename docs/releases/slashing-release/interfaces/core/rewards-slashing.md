# Rewards

Notably we don't include a foundation usecase due to changes in the foundation rewards flow.

The only difference is the new `SLASHABLE_STAKE` `RewardType` added.

```solidity
interface IRewardsCoordinator {

    /// STRUCTS & ENUMS
    
    enum RewardType {
        UNIFORM,
        DELEGATED_STAKE,
        SLASHABLE_STAKE
    }
}
```

### rewardOperatorSetForRange

Called AVSs to create a `RewardsSubmission` that rewards all stakers delegated to operators in one of their operatorSets over a certain time range. The payment specifies a certain amount of a certain token to distribute to the stakers (and their operators via commissions) evenly over the time range.

1. `UNIFORM` rewards award each registered operator the same amount of tokens
2. `DELEGATED_STAKE` rewards awards each registered operator an amount of tokens proportional to the weight of their delegated stake according to the supplied strategies and multipliers
3. `SLASHABLE_STAKE` rewards awards each registered operator an amount of tokens proportional to the weight of their slasahble stake according to the supplied strategies and multipliers

Rewards to operators are further distributed to each of their stakers proportional to their weight given by the supplied strategies and multipliers. The operator retains their commission on the total award given to them.

The set of `RewardType`s is intended to extent over the next releases.

Reverts if:

1. The reward's time range starts before `block.timestamp - MAX_RETROACTIVE_LENGTH`
2. The reward's time range starts after `block.timestamp + MAX_FUTURE_LENGTH`
3. The reward's duration of the reward submission is not a multiple of `CALCULATION_INTERVAL_SECONDS`
4. The reward's duration is greater than `MAX_REWARDS_DURATION`
5. The reward's amount is > `1e38-1`
6. `msg.sender` doesn't have a sufficient approval or allowance to transfer the tokens in the rewards submissions
6. Any of the strategies being rewarded for are not whitelisted for deposit
7. The operatorSetId is not stored in the AVSDirectory