# Rewards

Notably we don't include a foundation usecase due to changes in the foundation rewards flow.

These functions are additional to the existing functions in the Rewards protocol.

```solidity
interface IRewardsCoordinator {

    /// STRUCTS & ENUMS
    
    enum RewardType {
        UNIFORM,
        DELEGATED_STAKE
    }
    
    struct StrategyAndMultiplier {
        IStrategy strategy;
        uint96 multiplier;
    }
    
    struct OperatorSetRewardsSubmission {
        RewardType rewardType;
        uint32 operatorSetId;
        StrategyAndMultiplier[] strategiesAndMultipliers;
        IERC20 token;
        uint256 amount;
        uint32 startTimestamp;
        uint32 duration;
    }
    
    /// EVENTS

    event OperatorCommissionUpdated(
        address indexed operator,
        IAVSDirectory.OperatorSet indexed operatorSet
        RewardType rewardType,
        uint16 indexed newCommissionBips,
        uint32 effectTimestamp,
    );
    
    event OperatorSetRewardsCreated(
        address indexed avs,
        uint256 indexed submissionNonce,
        bytes32 indexed rewardsSubmissionHash,
        OperatorSetRewardsSubmission rewardsSubmission
    );
    
    /// EXTERNAL - STATE MODIFYING

    /**
     * @notice Sets the commission an operator takes in bips for a given reward type and operatorSet
     *
     * @param operatorSet The operatorSet to update commission for
     * @param rewardType The associated rewardType to update commission for 
     * @param commissionBips The commision, in bips
     *
     * @dev The commission can range from 1 to 10000
     * @dev The commission update takes effect after 7 days
     * @dev Default commission for all operators is 1000 bips (10%)
     */
    function setOperatorCommission(
        IAVSDirectory.OperatorSet calldata operatorSet,
        RewardType rewardType,
        uint16 commissionBips
    ) external returns (uint32 effectTimestamp);
    
    /**
     * @notice Creates a new rewards submission on behalf of an AVS for a given operatorSet, 
     * to be split amongst the set of stakers delegated to operators who are 
     * registered to the msg.sender AVS and the given operatorSetId
     *
     * @param rewardsSubmissions The operatorSet rewards submission being created for
     *
     * @dev msg.sender is the AVS in the operatorSet the rewards submission is being made to
     * @dev AVSs set their reward type depending on what metric they want rewards
     * distributed proportional to
     * @dev The tokens in the rewards submissions are sent to the `RewardsCoordinator` contract
     * @dev Strategies of each rewards submission must be in ascending order of addresses to check for duplicates
     */
    function rewardOperatorSetForRange(
        OperatorSetRewardsSubmission[] calldata rewardsSubmissions
    ) external;
}
```

### setOperatorCommission

Sets the commission an operator takes in bips for a given reward type and operatorSet. Commission is active after 7 days. Notably, this gives stakers the guarantee that the commission updates they see when delegating will remain for 7 days, but does not give them the ability to withdraw before new updates occur.

Reverts if:

1. The commission bips are not between \[1,10000]
   1. The minimum of 1 is used to interpret 0 as the default 1000 bip commission
2. The OperatorSet does not exist in the AVSDirectory

### rewardOperatorSetForRange

Called AVSs to create a set of `RewardsSubmission`s that rewards all stakers delegated to operators in one of their operatorSets over a certain time range. The payment specifies a certain amount of a certain token to distribute to the stakers (and their operators via commissions) evenly over the time range.

1. `UNIFORM` rewards award each registered operator the same amount of tokens
2. `DELEGATED_STAKE` rewards award each registered operator an amount of tokens proportional to the weight of their delegated stake according to the supplied strategies and multipliers

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
