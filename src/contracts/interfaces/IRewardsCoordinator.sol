// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../libraries/OperatorSetLib.sol";

import "./IAllocationManager.sol";
import "./IDelegationManager.sol";
import "./IStrategyManager.sol";
import "./IPauserRegistry.sol";
import "./IPermissionController.sol";
import "./IStrategy.sol";
import "./ISemVerMixin.sol";

interface IRewardsCoordinatorErrors {
    /// @dev Thrown when msg.sender is not allowed to call a function
    error UnauthorizedCaller();
    /// @dev Thrown when a earner not an AVS or Operator
    error InvalidEarner();

    /// Invalid Inputs

    /// @dev Thrown when an input address is zero
    error InvalidAddressZero();
    /// @dev Thrown when an invalid root is provided.
    error InvalidRoot();
    /// @dev Thrown when an invalid root index is provided.
    error InvalidRootIndex();
    /// @dev Thrown when input arrays length is zero.
    error InputArrayLengthZero();
    /// @dev Thrown when two array parameters have mismatching lengths.
    error InputArrayLengthMismatch();
    /// @dev Thrown when provided root is not for new calculated period.
    error NewRootMustBeForNewCalculatedPeriod();
    /// @dev Thrown when rewards end timestamp has not elapsed.
    error RewardsEndTimestampNotElapsed();
    /// @dev Thrown when an invalid operator set is provided.
    error InvalidOperatorSet();

    /// Rewards Submissions

    /// @dev Thrown when input `amount` is zero.
    error AmountIsZero();
    /// @dev Thrown when input `amount` exceeds maximum.
    error AmountExceedsMax();
    /// @dev Thrown when input `split` exceeds `ONE_HUNDRED_IN_BIPS`
    error SplitExceedsMax();
    /// @dev Thrown when an operator attempts to set a split before the previous one becomes active
    error PreviousSplitPending();
    /// @dev Thrown when input `duration` exceeds maximum.
    error DurationExceedsMax();
    /// @dev Thrown when input `duration` is zero.
    error DurationIsZero();
    /// @dev Thrown when input `duration` is not evenly divisble by CALCULATION_INTERVAL_SECONDS.
    error InvalidDurationRemainder();
    /// @dev Thrown when GENESIS_REWARDS_TIMESTAMP is not evenly divisble by CALCULATION_INTERVAL_SECONDS.
    error InvalidGenesisRewardsTimestampRemainder();
    /// @dev Thrown when CALCULATION_INTERVAL_SECONDS is not evenly divisble by SNAPSHOT_CADENCE.
    error InvalidCalculationIntervalSecondsRemainder();
    /// @dev Thrown when `startTimestamp` is not evenly divisble by CALCULATION_INTERVAL_SECONDS.
    error InvalidStartTimestampRemainder();
    /// @dev Thrown when `startTimestamp` is too far in the future.
    error StartTimestampTooFarInFuture();
    /// @dev Thrown when `startTimestamp` is too far in the past.
    error StartTimestampTooFarInPast();
    /// @dev Thrown when an attempt to use a non-whitelisted strategy is made.
    error StrategyNotWhitelisted();
    /// @dev Thrown when `strategies` is not sorted in ascending order.
    error StrategiesNotInAscendingOrder();
    /// @dev Thrown when `operators` are not sorted in ascending order
    error OperatorsNotInAscendingOrder();
    /// @dev Thrown when an operator-directed rewards submission is not retroactive
    error SubmissionNotRetroactive();

    /// Claims

    /// @dev Thrown when an invalid earner claim proof is provided.
    error InvalidClaimProof();
    /// @dev Thrown when an invalid token leaf index is provided.
    error InvalidTokenLeafIndex();
    /// @dev Thrown when an invalid earner leaf index is provided.
    error InvalidEarnerLeafIndex();
    /// @dev Thrown when cumulative earnings are not greater than cumulative claimed.
    error EarningsNotGreaterThanClaimed();

    /// Reward Root Checks

    /// @dev Thrown if a root has already been disabled.
    error RootDisabled();
    /// @dev Thrown if a root has not been activated yet.
    error RootNotActivated();
    /// @dev Thrown if a root has already been activated.
    error RootAlreadyActivated();
}

interface IRewardsCoordinatorTypes {
    /**
     * @notice A linear combination of strategies and multipliers for AVSs to weigh
     * EigenLayer strategies.
     * @param strategy The EigenLayer strategy to be used for the rewards submission
     * @param multiplier The weight of the strategy in the rewards submission
     */
    struct StrategyAndMultiplier {
        IStrategy strategy;
        uint96 multiplier;
    }

    /**
     * @notice A reward struct for an operator
     * @param operator The operator to be rewarded
     * @param amount The reward amount for the operator
     */
    struct OperatorReward {
        address operator;
        uint256 amount;
    }

    /**
     * @notice A split struct for an Operator
     * @param oldSplitBips The old split in basis points. This is the split that is active if `block.timestamp < activatedAt`
     * @param newSplitBips The new split in basis points. This is the split that is active if `block.timestamp >= activatedAt`
     * @param activatedAt The timestamp at which the split will be activated
     */
    struct OperatorSplit {
        uint16 oldSplitBips;
        uint16 newSplitBips;
        uint32 activatedAt;
    }

    /**
     * Sliding Window for valid RewardsSubmission startTimestamp
     *
     * Scenario A: GENESIS_REWARDS_TIMESTAMP IS WITHIN RANGE
     *         <-----MAX_RETROACTIVE_LENGTH-----> t (block.timestamp) <---MAX_FUTURE_LENGTH--->
     *             <--------------------valid range for startTimestamp------------------------>
     *             ^
     *         GENESIS_REWARDS_TIMESTAMP
     *
     *
     * Scenario B: GENESIS_REWARDS_TIMESTAMP IS OUT OF RANGE
     *         <-----MAX_RETROACTIVE_LENGTH-----> t (block.timestamp) <---MAX_FUTURE_LENGTH--->
     *         <------------------------valid range for startTimestamp------------------------>
     *     ^
     * GENESIS_REWARDS_TIMESTAMP
     * @notice RewardsSubmission struct submitted by AVSs when making rewards for their operators and stakers
     * RewardsSubmission can be for a time range within the valid window for startTimestamp and must be within max duration.
     * See `createAVSRewardsSubmission()` for more details.
     * @param strategiesAndMultipliers The strategies and their relative weights
     * cannot have duplicate strategies and need to be sorted in ascending address order
     * @param token The rewards token to be distributed
     * @param amount The total amount of tokens to be distributed
     * @param startTimestamp The timestamp (seconds) at which the submission range is considered for distribution
     * could start in the past or in the future but within a valid range. See the diagram above.
     * @param duration The duration of the submission range in seconds. Must be <= MAX_REWARDS_DURATION
     */
    struct RewardsSubmission {
        StrategyAndMultiplier[] strategiesAndMultipliers;
        IERC20 token;
        uint256 amount;
        uint32 startTimestamp;
        uint32 duration;
    }

    /**
     * @notice OperatorDirectedRewardsSubmission struct submitted by AVSs when making operator-directed rewards for their operators and stakers.
     * @param strategiesAndMultipliers The strategies and their relative weights.
     * @param token The rewards token to be distributed.
     * @param operatorRewards The rewards for the operators.
     * @param startTimestamp The timestamp (seconds) at which the submission range is considered for distribution.
     * @param duration The duration of the submission range in seconds.
     * @param description Describes what the rewards submission is for.
     */
    struct OperatorDirectedRewardsSubmission {
        StrategyAndMultiplier[] strategiesAndMultipliers;
        IERC20 token;
        OperatorReward[] operatorRewards;
        uint32 startTimestamp;
        uint32 duration;
        string description;
    }

    /**
     * @notice A distribution root is a merkle root of the distribution of earnings for a given period.
     * The RewardsCoordinator stores all historical distribution roots so that earners can claim their earnings against older roots
     * if they wish but the merkle tree contains the cumulative earnings of all earners and tokens for a given period so earners (or their claimers if set)
     * only need to claim against the latest root to claim all available earnings.
     * @param root The merkle root of the distribution
     * @param rewardsCalculationEndTimestamp The timestamp (seconds) until which rewards have been calculated
     * @param activatedAt The timestamp (seconds) at which the root can be claimed against
     */
    struct DistributionRoot {
        bytes32 root;
        uint32 rewardsCalculationEndTimestamp;
        uint32 activatedAt;
        bool disabled;
    }

    /**
     * @notice Internal leaf in the merkle tree for the earner's account leaf
     * @param earner The address of the earner
     * @param earnerTokenRoot The merkle root of the earner's token subtree
     * Each leaf in the earner's token subtree is a TokenTreeMerkleLeaf
     */
    struct EarnerTreeMerkleLeaf {
        address earner;
        bytes32 earnerTokenRoot;
    }

    /**
     * @notice The actual leaves in the distribution merkle tree specifying the token earnings
     * for the respective earner's subtree. Each leaf is a claimable amount of a token for an earner.
     * @param token The token for which the earnings are being claimed
     * @param cumulativeEarnings The cumulative earnings of the earner for the token
     */
    struct TokenTreeMerkleLeaf {
        IERC20 token;
        uint256 cumulativeEarnings;
    }

    /**
     * @notice A claim against a distribution root called by an
     * earners claimer (could be the earner themselves). Each token claim will claim the difference
     * between the cumulativeEarnings of the earner and the cumulativeClaimed of the claimer.
     * Each claim can specify which of the earner's earned tokens they want to claim.
     * See `processClaim()` for more details.
     * @param rootIndex The index of the root in the list of DistributionRoots
     * @param earnerIndex The index of the earner's account root in the merkle tree
     * @param earnerTreeProof The proof of the earner's EarnerTreeMerkleLeaf against the merkle root
     * @param earnerLeaf The earner's EarnerTreeMerkleLeaf struct, providing the earner address and earnerTokenRoot
     * @param tokenIndices The indices of the token leaves in the earner's subtree
     * @param tokenTreeProofs The proofs of the token leaves against the earner's earnerTokenRoot
     * @param tokenLeaves The token leaves to be claimed
     * @dev The merkle tree is structured with the merkle root at the top and EarnerTreeMerkleLeaf as internal leaves
     * in the tree. Each earner leaf has its own subtree with TokenTreeMerkleLeaf as leaves in the subtree.
     * To prove a claim against a specified rootIndex(which specifies the distributionRoot being used),
     * the claim will first verify inclusion of the earner leaf in the tree against _distributionRoots[rootIndex].root.
     * Then for each token, it will verify inclusion of the token leaf in the earner's subtree against the earner's earnerTokenRoot.
     */
    struct RewardsMerkleClaim {
        uint32 rootIndex;
        uint32 earnerIndex;
        bytes earnerTreeProof;
        EarnerTreeMerkleLeaf earnerLeaf;
        uint32[] tokenIndices;
        bytes[] tokenTreeProofs;
        TokenTreeMerkleLeaf[] tokenLeaves;
    }

    /**
     * @notice Parameters for the RewardsCoordinator constructor
     * @param delegationManager The address of the DelegationManager contract
     * @param strategyManager The address of the StrategyManager contract
     * @param allocationManager The address of the AllocationManager contract
     * @param pauserRegistry The address of the PauserRegistry contract
     * @param permissionController The address of the PermissionController contract
     * @param CALCULATION_INTERVAL_SECONDS The interval at which rewards are calculated
     * @param MAX_REWARDS_DURATION The maximum duration of a rewards submission
     * @param MAX_RETROACTIVE_LENGTH The maximum retroactive length of a rewards submission
     * @param MAX_FUTURE_LENGTH The maximum future length of a rewards submission
     * @param GENESIS_REWARDS_TIMESTAMP The timestamp at which rewards are first calculated
     * @param version The semantic version of the contract (e.g. "v1.2.3")
     * @dev Needed to avoid stack-too-deep errors
     */
    struct RewardsCoordinatorConstructorParams {
        IDelegationManager delegationManager;
        IStrategyManager strategyManager;
        IAllocationManager allocationManager;
        IPauserRegistry pauserRegistry;
        IPermissionController permissionController;
        uint32 CALCULATION_INTERVAL_SECONDS;
        uint32 MAX_REWARDS_DURATION;
        uint32 MAX_RETROACTIVE_LENGTH;
        uint32 MAX_FUTURE_LENGTH;
        uint32 GENESIS_REWARDS_TIMESTAMP;
        string version;
    }
}

interface IRewardsCoordinatorEvents is IRewardsCoordinatorTypes {
    /// @notice emitted when an AVS creates a valid RewardsSubmission
    event AVSRewardsSubmissionCreated(
        address indexed avs,
        uint256 indexed submissionNonce,
        bytes32 indexed rewardsSubmissionHash,
        RewardsSubmission rewardsSubmission
    );

    /// @notice emitted when a valid RewardsSubmission is created for all stakers by a valid submitter
    event RewardsSubmissionForAllCreated(
        address indexed submitter,
        uint256 indexed submissionNonce,
        bytes32 indexed rewardsSubmissionHash,
        RewardsSubmission rewardsSubmission
    );

    /// @notice emitted when a valid RewardsSubmission is created when rewardAllStakersAndOperators is called
    event RewardsSubmissionForAllEarnersCreated(
        address indexed tokenHopper,
        uint256 indexed submissionNonce,
        bytes32 indexed rewardsSubmissionHash,
        RewardsSubmission rewardsSubmission
    );

    /**
     * @notice Emitted when an AVS creates a valid `OperatorDirectedRewardsSubmission`
     * @param caller The address calling `createOperatorDirectedAVSRewardsSubmission`.
     * @param avs The avs on behalf of which the operator-directed rewards are being submitted.
     * @param operatorDirectedRewardsSubmissionHash Keccak256 hash of (`avs`, `submissionNonce` and `operatorDirectedRewardsSubmission`).
     * @param submissionNonce Current nonce of the avs. Used to generate a unique submission hash.
     * @param operatorDirectedRewardsSubmission The Operator-Directed Rewards Submission. Contains the token, start timestamp, duration, operator rewards, description and, strategy and multipliers.
     */
    event OperatorDirectedAVSRewardsSubmissionCreated(
        address indexed caller,
        address indexed avs,
        bytes32 indexed operatorDirectedRewardsSubmissionHash,
        uint256 submissionNonce,
        OperatorDirectedRewardsSubmission operatorDirectedRewardsSubmission
    );

    /**
     * @notice Emitted when an AVS creates a valid `OperatorDirectedRewardsSubmission` for an operator set.
     * @param caller The address calling `createOperatorDirectedOperatorSetRewardsSubmission`.
     * @param operatorDirectedRewardsSubmissionHash Keccak256 hash of (`avs`, `submissionNonce` and `operatorDirectedRewardsSubmission`).
     * @param operatorSet The operatorSet on behalf of which the operator-directed rewards are being submitted.
     * @param submissionNonce Current nonce of the avs. Used to generate a unique submission hash.
     * @param operatorDirectedRewardsSubmission The Operator-Directed Rewards Submission. Contains the token, start timestamp, duration, operator rewards, description and, strategy and multipliers.
     */
    event OperatorDirectedOperatorSetRewardsSubmissionCreated(
        address indexed caller,
        bytes32 indexed operatorDirectedRewardsSubmissionHash,
        OperatorSet operatorSet,
        uint256 submissionNonce,
        OperatorDirectedRewardsSubmission operatorDirectedRewardsSubmission
    );

    /// @notice rewardsUpdater is responsible for submitting DistributionRoots, only owner can set rewardsUpdater
    event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater);

    event RewardsForAllSubmitterSet(
        address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue
    );

    event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay);
    event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips);

    /**
     * @notice Emitted when the operator split for an AVS is set.
     * @param caller The address calling `setOperatorAVSSplit`.
     * @param operator The operator on behalf of which the split is being set.
     * @param avs The avs for which the split is being set by the operator.
     * @param activatedAt The timestamp at which the split will be activated.
     * @param oldOperatorAVSSplitBips The old split for the operator for the AVS.
     * @param newOperatorAVSSplitBips The new split for the operator for the AVS.
     */
    event OperatorAVSSplitBipsSet(
        address indexed caller,
        address indexed operator,
        address indexed avs,
        uint32 activatedAt,
        uint16 oldOperatorAVSSplitBips,
        uint16 newOperatorAVSSplitBips
    );

    /**
     * @notice Emitted when the operator split for Programmatic Incentives is set.
     * @param caller The address calling `setOperatorPISplit`.
     * @param operator The operator on behalf of which the split is being set.
     * @param activatedAt The timestamp at which the split will be activated.
     * @param oldOperatorPISplitBips The old split for the operator for Programmatic Incentives.
     * @param newOperatorPISplitBips The new split for the operator for Programmatic Incentives.
     */
    event OperatorPISplitBipsSet(
        address indexed caller,
        address indexed operator,
        uint32 activatedAt,
        uint16 oldOperatorPISplitBips,
        uint16 newOperatorPISplitBips
    );

    /**
     * @notice Emitted when the operator split for a given operatorSet is set.
     * @param caller The address calling `setOperatorSetSplit`.
     * @param operator The operator on behalf of which the split is being set.
     * @param operatorSet The operatorSet for which the split is being set.
     * @param activatedAt The timestamp at which the split will be activated.
     * @param oldOperatorSetSplitBips The old split for the operator for the operatorSet.
     * @param newOperatorSetSplitBips The new split for the operator for the operatorSet.
     */
    event OperatorSetSplitBipsSet(
        address indexed caller,
        address indexed operator,
        OperatorSet operatorSet,
        uint32 activatedAt,
        uint16 oldOperatorSetSplitBips,
        uint16 newOperatorSetSplitBips
    );

    event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer);

    /// @notice rootIndex is the specific array index of the newly created root in the storage array
    event DistributionRootSubmitted(
        uint32 indexed rootIndex,
        bytes32 indexed root,
        uint32 indexed rewardsCalculationEndTimestamp,
        uint32 activatedAt
    );

    event DistributionRootDisabled(uint32 indexed rootIndex);

    /// @notice root is one of the submitted distribution roots that was claimed against
    event RewardsClaimed(
        bytes32 root,
        address indexed earner,
        address indexed claimer,
        address indexed recipient,
        IERC20 token,
        uint256 claimedAmount
    );
}

/**
 * @title Interface for the `IRewardsCoordinator` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Allows AVSs to make "Rewards Submissions", which get distributed amongst the AVSs' confirmed
 * Operators and the Stakers delegated to those Operators.
 * Calculations are performed based on the completed RewardsSubmission, with the results posted in
 * a Merkle root against which Stakers & Operators can make claims.
 */
interface IRewardsCoordinator is IRewardsCoordinatorErrors, IRewardsCoordinatorEvents, ISemVerMixin {
    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, rewardsUpdater and
     * configures the initial paused status, activationDelay, and defaultOperatorSplitBips.
     */
    function initialize(
        address initialOwner,
        uint256 initialPausedStatus,
        address _rewardsUpdater,
        uint32 _activationDelay,
        uint16 _defaultSplitBips
    ) external;

    /**
     * @notice Creates a new rewards submission on behalf of an AVS, to be split amongst the
     * set of stakers delegated to operators who are registered to the `avs`
     * @param rewardsSubmissions The rewards submissions being created
     * @dev Expected to be called by the ServiceManager of the AVS on behalf of which the submission is being made
     * @dev The duration of the `rewardsSubmission` cannot exceed `MAX_REWARDS_DURATION`
     * @dev The duration of the `rewardsSubmission` cannot be 0 and must be a multiple of `CALCULATION_INTERVAL_SECONDS`
     * @dev The tokens are sent to the `RewardsCoordinator` contract
     * @dev Strategies must be in ascending order of addresses to check for duplicates
     * @dev This function will revert if the `rewardsSubmission` is malformed,
     * e.g. if the `strategies` and `weights` arrays are of non-equal lengths
     */
    function createAVSRewardsSubmission(
        RewardsSubmission[] calldata rewardsSubmissions
    ) external;

    /**
     * @notice similar to `createAVSRewardsSubmission` except the rewards are split amongst *all* stakers
     * rather than just those delegated to operators who are registered to a single avs and is
     * a permissioned call based on isRewardsForAllSubmitter mapping.
     * @param rewardsSubmissions The rewards submissions being created
     */
    function createRewardsForAllSubmission(
        RewardsSubmission[] calldata rewardsSubmissions
    ) external;

    /**
     * @notice Creates a new rewards submission for all earners across all AVSs.
     * Earners in this case indicating all operators and their delegated stakers. Undelegated stake
     * is not rewarded from this RewardsSubmission. This interface is only callable
     * by the token hopper contract from the Eigen Foundation
     * @param rewardsSubmissions The rewards submissions being created
     */
    function createRewardsForAllEarners(
        RewardsSubmission[] calldata rewardsSubmissions
    ) external;

    /**
     * @notice Creates a new operator-directed rewards submission on behalf of an AVS, to be split amongst the operators and
     * set of stakers delegated to operators who are registered to the `avs`.
     * @param avs The AVS on behalf of which the reward is being submitted
     * @param operatorDirectedRewardsSubmissions The operator-directed rewards submissions being created
     * @dev Expected to be called by the ServiceManager of the AVS on behalf of which the submission is being made
     * @dev The duration of the `rewardsSubmission` cannot exceed `MAX_REWARDS_DURATION`
     * @dev The duration of the `rewardsSubmission` cannot be 0 and must be a multiple of `CALCULATION_INTERVAL_SECONDS`
     * @dev The tokens are sent to the `RewardsCoordinator` contract
     * @dev The `RewardsCoordinator` contract needs a token approval of sum of all `operatorRewards` in the `operatorDirectedRewardsSubmissions`, before calling this function.
     * @dev Strategies must be in ascending order of addresses to check for duplicates
     * @dev Operators must be in ascending order of addresses to check for duplicates.
     * @dev This function will revert if the `operatorDirectedRewardsSubmissions` is malformed.
     */
    function createOperatorDirectedAVSRewardsSubmission(
        address avs,
        OperatorDirectedRewardsSubmission[] calldata operatorDirectedRewardsSubmissions
    ) external;

    /**
     * @notice Creates a new operator-directed rewards submission for an operator set, to be split amongst the operators and
     * set of stakers delegated to operators who are part of the operator set.
     * @param operatorSet The operator set for which the rewards are being submitted
     * @param operatorDirectedRewardsSubmissions The operator-directed rewards submissions being created
     * @dev Expected to be called by the AVS that created the operator set
     * @dev The duration of the `rewardsSubmission` cannot exceed `MAX_REWARDS_DURATION`
     * @dev The duration of the `rewardsSubmission` cannot be 0 and must be a multiple of `CALCULATION_INTERVAL_SECONDS`
     * @dev The tokens are sent to the `RewardsCoordinator` contract
     * @dev The `RewardsCoordinator` contract needs a token approval of sum of all `operatorRewards` in the `operatorDirectedRewardsSubmissions`, before calling this function
     * @dev Strategies must be in ascending order of addresses to check for duplicates
     * @dev Operators must be in ascending order of addresses to check for duplicates
     * @dev This function will revert if the `operatorDirectedRewardsSubmissions` is malformed
     */
    function createOperatorDirectedOperatorSetRewardsSubmission(
        OperatorSet calldata operatorSet,
        OperatorDirectedRewardsSubmission[] calldata operatorDirectedRewardsSubmissions
    ) external;

    /**
     * @notice Claim rewards against a given root (read from _distributionRoots[claim.rootIndex]).
     * Earnings are cumulative so earners don't have to claim against all distribution roots they have earnings for,
     * they can simply claim against the latest root and the contract will calculate the difference between
     * their cumulativeEarnings and cumulativeClaimed. This difference is then transferred to recipient address.
     * @param claim The RewardsMerkleClaim to be processed.
     * Contains the root index, earner, token leaves, and required proofs
     * @param recipient The address recipient that receives the ERC20 rewards
     * @dev only callable by the valid claimer, that is
     * if claimerFor[claim.earner] is address(0) then only the earner can claim, otherwise only
     * claimerFor[claim.earner] can claim the rewards.
     */
    function processClaim(RewardsMerkleClaim calldata claim, address recipient) external;

    /**
     * @notice Batch claim rewards against a given root (read from _distributionRoots[claim.rootIndex]).
     * Earnings are cumulative so earners don't have to claim against all distribution roots they have earnings for,
     * they can simply claim against the latest root and the contract will calculate the difference between
     * their cumulativeEarnings and cumulativeClaimed. This difference is then transferred to recipient address.
     * @param claims The RewardsMerkleClaims to be processed.
     * Contains the root index, earner, token leaves, and required proofs
     * @param recipient The address recipient that receives the ERC20 rewards
     * @dev only callable by the valid claimer, that is
     * if claimerFor[claim.earner] is address(0) then only the earner can claim, otherwise only
     * claimerFor[claim.earner] can claim the rewards.
     * @dev This function may fail to execute with a large number of claims due to gas limits. Use a smaller array of claims if necessary.
     */
    function processClaims(RewardsMerkleClaim[] calldata claims, address recipient) external;

    /**
     * @notice Creates a new distribution root. activatedAt is set to block.timestamp + activationDelay
     * @param root The merkle root of the distribution
     * @param rewardsCalculationEndTimestamp The timestamp until which rewards have been calculated
     * @dev Only callable by the rewardsUpdater
     */
    function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) external;

    /**
     * @notice allow the rewardsUpdater to disable/cancel a pending root submission in case of an error
     * @param rootIndex The index of the root to be disabled
     */
    function disableRoot(
        uint32 rootIndex
    ) external;

    /**
     * @notice Sets the address of the entity that can call `processClaim` on ehalf of an earner
     * @param claimer The address of the entity that can call `processClaim` on behalf of the earner
     * @dev Assumes msg.sender is the earner
     */
    function setClaimerFor(
        address claimer
    ) external;

    /**
     * @notice Sets the address of the entity that can call `processClaim` on behalf of an earner
     * @param earner The address to set the claimer for
     * @param claimer The address of the entity that can call `processClaim` on behalf of the earner
     * @dev Only callable by operators or AVSs. We define an AVS that has created at least one
     *      operatorSet in the `AllocationManager`
     */
    function setClaimerFor(address earner, address claimer) external;

    /**
     * @notice Sets the delay in timestamp before a posted root can be claimed against
     * @dev Only callable by the contract owner
     * @param _activationDelay The new value for activationDelay
     */
    function setActivationDelay(
        uint32 _activationDelay
    ) external;

    /**
     * @notice Sets the default split for all operators across all avss.
     * @param split The default split for all operators across all avss in bips.
     * @dev Only callable by the contract owner.
     */
    function setDefaultOperatorSplit(
        uint16 split
    ) external;

    /**
     * @notice Sets the split for a specific operator for a specific avs
     * @param operator The operator who is setting the split
     * @param avs The avs for which the split is being set by the operator
     * @param split The split for the operator for the specific avs in bips.
     * @dev Only callable by the operator
     * @dev Split has to be between 0 and 10000 bips (inclusive)
     * @dev The split will be activated after the activation delay
     */
    function setOperatorAVSSplit(address operator, address avs, uint16 split) external;

    /**
     * @notice Sets the split for a specific operator for Programmatic Incentives.
     * @param operator The operator on behalf of which the split is being set.
     * @param split The split for the operator for Programmatic Incentives in bips.
     * @dev Only callable by the operator
     * @dev Split has to be between 0 and 10000 bips (inclusive)
     * @dev The split will be activated after the activation delay
     */
    function setOperatorPISplit(address operator, uint16 split) external;

    /**
     * @notice Sets the split for a specific operator for a specific operatorSet.
     * @param operator The operator who is setting the split.
     * @param operatorSet The operatorSet for which the split is being set by the operator.
     * @param split The split for the operator for the specific operatorSet in bips.
     * @dev Only callable by the operator
     * @dev Split has to be between 0 and 10000 bips (inclusive)
     * @dev The split will be activated after the activation delay
     */
    function setOperatorSetSplit(address operator, OperatorSet calldata operatorSet, uint16 split) external;

    /**
     * @notice Sets the permissioned `rewardsUpdater` address which can post new roots
     * @dev Only callable by the contract owner
     * @param _rewardsUpdater The address of the new rewardsUpdater
     */
    function setRewardsUpdater(
        address _rewardsUpdater
    ) external;

    /**
     * @notice Sets the permissioned `rewardsForAllSubmitter` address which can submit createRewardsForAllSubmission
     * @dev Only callable by the contract owner
     * @param _submitter The address of the rewardsForAllSubmitter
     * @param _newValue The new value for isRewardsForAllSubmitter
     */
    function setRewardsForAllSubmitter(address _submitter, bool _newValue) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @notice Delay in timestamp (seconds) before a posted root can be claimed against
    function activationDelay() external view returns (uint32);

    /// @notice The timestamp until which RewardsSubmissions have been calculated
    function currRewardsCalculationEndTimestamp() external view returns (uint32);

    /// @notice Mapping: earner => the address of the entity who can call `processClaim` on behalf of the earner
    function claimerFor(
        address earner
    ) external view returns (address);

    /// @notice Mapping: claimer => token => total amount claimed
    function cumulativeClaimed(address claimer, IERC20 token) external view returns (uint256);

    /// @notice the default split for all operators across all avss
    function defaultOperatorSplitBips() external view returns (uint16);

    /// @notice the split for a specific `operator` for a specific `avs`
    function getOperatorAVSSplit(address operator, address avs) external view returns (uint16);

    /// @notice the split for a specific `operator` for Programmatic Incentives
    function getOperatorPISplit(
        address operator
    ) external view returns (uint16);

    /// @notice Returns the split for a specific `operator` for a given `operatorSet`
    function getOperatorSetSplit(address operator, OperatorSet calldata operatorSet) external view returns (uint16);

    /// @notice return the hash of the earner's leaf
    function calculateEarnerLeafHash(
        EarnerTreeMerkleLeaf calldata leaf
    ) external pure returns (bytes32);

    /// @notice returns the hash of the earner's token leaf
    function calculateTokenLeafHash(
        TokenTreeMerkleLeaf calldata leaf
    ) external pure returns (bytes32);

    /// @notice returns 'true' if the claim would currently pass the check in `processClaims`
    /// but will revert if not valid
    function checkClaim(
        RewardsMerkleClaim calldata claim
    ) external view returns (bool);

    /// @notice returns the number of distribution roots posted
    function getDistributionRootsLength() external view returns (uint256);

    /// @notice returns the distributionRoot at the specified index
    function getDistributionRootAtIndex(
        uint256 index
    ) external view returns (DistributionRoot memory);

    /// @notice returns the current distributionRoot
    function getCurrentDistributionRoot() external view returns (DistributionRoot memory);

    /// @notice loop through the distribution roots from reverse and get latest root that is not disabled and activated
    /// i.e. a root that can be claimed against
    function getCurrentClaimableDistributionRoot() external view returns (DistributionRoot memory);

    /// @notice loop through distribution roots from reverse and return index from hash
    function getRootIndexFromHash(
        bytes32 rootHash
    ) external view returns (uint32);

    /// @notice The address of the entity that can update the contract with new merkle roots
    function rewardsUpdater() external view returns (address);

    /**
     * @notice The interval in seconds at which the calculation for a RewardsSubmission distribution is done.
     * @dev Rewards Submission durations must be multiples of this interval.
     */
    function CALCULATION_INTERVAL_SECONDS() external view returns (uint32);

    /// @notice The maximum amount of time (seconds) that a RewardsSubmission can span over
    function MAX_REWARDS_DURATION() external view returns (uint32);

    /// @notice max amount of time (seconds) that a submission can start in the past
    function MAX_RETROACTIVE_LENGTH() external view returns (uint32);

    /// @notice max amount of time (seconds) that a submission can start in the future
    function MAX_FUTURE_LENGTH() external view returns (uint32);

    /// @notice absolute min timestamp (seconds) that a submission can start at
    function GENESIS_REWARDS_TIMESTAMP() external view returns (uint32);
}
