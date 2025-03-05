// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IRewardsCoordinator.sol";

/**
 * @title Storage variables for the `RewardsCoordinator` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract RewardsCoordinatorStorage is IRewardsCoordinator {
    // Constants

    /// @dev Index for flag that pauses calling createAVSRewardsSubmission
    uint8 internal constant PAUSED_AVS_REWARDS_SUBMISSION = 0;
    /// @dev Index for flag that pauses calling createRewardsForAllSubmission
    uint8 internal constant PAUSED_REWARDS_FOR_ALL_SUBMISSION = 1;
    /// @dev Index for flag that pauses calling processClaim
    uint8 internal constant PAUSED_PROCESS_CLAIM = 2;
    /// @dev Index for flag that pauses submitRoots and disableRoot
    uint8 internal constant PAUSED_SUBMIT_DISABLE_ROOTS = 3;
    /// @dev Index for flag that pauses calling rewardAllStakersAndOperators
    uint8 internal constant PAUSED_REWARD_ALL_STAKERS_AND_OPERATORS = 4;
    /// @dev Index for flag that pauses calling createOperatorDirectedAVSRewardsSubmission
    uint8 internal constant PAUSED_OPERATOR_DIRECTED_AVS_REWARDS_SUBMISSION = 5;
    /// @dev Index for flag that pauses calling setOperatorAVSSplit
    uint8 internal constant PAUSED_OPERATOR_AVS_SPLIT = 6;
    /// @dev Index for flag that pauses calling setOperatorPISplit
    uint8 internal constant PAUSED_OPERATOR_PI_SPLIT = 7;
    /// @dev Index for flag that pauses calling setOperatorSetSplit
    uint8 internal constant PAUSED_OPERATOR_SET_SPLIT = 8;
    /// @dev Index for flag that pauses calling setOperatorSetPerformanceRewardsSubmission
    uint8 internal constant PAUSED_OPERATOR_DIRECTED_OPERATOR_SET_REWARDS_SUBMISSION = 9;

    /// @dev Salt for the earner leaf, meant to distinguish from tokenLeaf since they have the same sized data
    uint8 internal constant EARNER_LEAF_SALT = 0;
    /// @dev Salt for the token leaf, meant to distinguish from earnerLeaf since they have the same sized data
    uint8 internal constant TOKEN_LEAF_SALT = 1;

    /// @notice The maximum rewards token amount for a single rewards submission, constrained by off-chain calculation
    uint256 internal constant MAX_REWARDS_AMOUNT = 1e38 - 1;
    /// @notice Equivalent to 100%, but in basis points.
    uint16 internal constant ONE_HUNDRED_IN_BIPS = 10_000;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    // Immutables

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegationManager;

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /// @notice The AllocationManager contract for EigenLayer
    IAllocationManager public immutable allocationManager;

    /// @notice The interval in seconds at which the calculation for rewards distribution is done.
    /// @dev RewardsSubmission durations must be multiples of this interval. This is going to be configured to 1 day
    uint32 public immutable CALCULATION_INTERVAL_SECONDS;
    /// @notice The maximum amount of time (seconds) that a rewards submission can span over
    uint32 public immutable MAX_REWARDS_DURATION;
    /// @notice max amount of time (seconds) that a rewards submission can start in the past
    uint32 public immutable MAX_RETROACTIVE_LENGTH;
    /// @notice max amount of time (seconds) that a rewards submission can start in the future
    uint32 public immutable MAX_FUTURE_LENGTH;
    /// @notice absolute min timestamp (seconds) that a rewards submission can start at
    uint32 public immutable GENESIS_REWARDS_TIMESTAMP;
    /// @notice The cadence at which a snapshot is taken offchain for calculating rewards distributions
    uint32 internal constant SNAPSHOT_CADENCE = 1 days;

    // Mutatables

    /// @dev Do not remove, deprecated storage.
    bytes32 internal __deprecated_DOMAIN_SEPARATOR;

    /**
     * @notice List of roots submitted by the rewardsUpdater
     * @dev Array is internal with an external getter so we can return a `DistributionRoot[] memory` object
     */
    DistributionRoot[] internal _distributionRoots;

    /// Slot 2
    /// @notice The address of the entity that can update the contract with new merkle roots
    address public rewardsUpdater;
    /// @notice Delay in timestamp (seconds) before a posted root can be claimed against
    uint32 public activationDelay;
    /// @notice Timestamp for last submitted DistributionRoot
    uint32 public currRewardsCalculationEndTimestamp;
    /// @notice the default split for all operators across all avss in bips.
    uint16 public defaultOperatorSplitBips;

    /// @notice Returns the `claimer` for a given `earner`.
    /// @dev The claimer is able to call `processClaim` on behalf of the `earner`.
    mapping(address earner => address claimer) public claimerFor;

    /// @notice Returns the total claimed amount for an `earner` for a given `token`.
    mapping(address earner => mapping(IERC20 token => uint256 totalClaimed)) public cumulativeClaimed;

    /// @notice Returns the submission `nonce` for an `avs`.
    mapping(address avs => uint256 nonce) public submissionNonce;

    /// @notice Returns whether a `hash` is a `valid` rewards submission hash for a given `avs`.
    mapping(address avs => mapping(bytes32 hash => bool valid)) public isAVSRewardsSubmissionHash;

    /// @notice Returns whether a `hash` is a `valid` rewards submission for all hash for a given `avs`.
    mapping(address avs => mapping(bytes32 hash => bool valid)) public isRewardsSubmissionForAllHash;

    /// @notice Returns whether a `submitter` is a `valid` rewards for all submitter.
    mapping(address submitter => bool valid) public isRewardsForAllSubmitter;

    /// @notice Returns whether a `hash` is a `valid` rewards submission for all earners hash for a given `avs`.
    mapping(address avs => mapping(bytes32 hash => bool valid)) public isRewardsSubmissionForAllEarnersHash;

    /// @notice Returns whether a `hash` is a `valid` operator set performance rewards submission hash for a given `avs`.
    mapping(address avs => mapping(bytes32 hash => bool valid)) public isOperatorDirectedAVSRewardsSubmissionHash;

    /// @notice Returns the `split` an `operator` takes for an `avs`.
    mapping(address operator => mapping(address avs => OperatorSplit split)) internal _operatorAVSSplitBips;

    /// @notice Returns the `split` an `operator` takes for Programmatic Incentives.
    mapping(address operator => OperatorSplit split) internal _operatorPISplitBips;

    /// @notice Returns the `split` an `operator` takes for a given operator set.
    mapping(address operator => mapping(bytes32 operatorSetKey => OperatorSplit split)) internal _operatorSetSplitBips;

    /// @notice Returns whether a `hash` is a `valid` operator set performance rewards submission hash for a given `avs`.
    mapping(address avs => mapping(bytes32 hash => bool valid)) public
        isOperatorDirectedOperatorSetRewardsSubmissionHash;

    // Construction

    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
        IAllocationManager _allocationManager,
        uint32 _CALCULATION_INTERVAL_SECONDS,
        uint32 _MAX_REWARDS_DURATION,
        uint32 _MAX_RETROACTIVE_LENGTH,
        uint32 _MAX_FUTURE_LENGTH,
        uint32 _GENESIS_REWARDS_TIMESTAMP
    ) {
        require(
            _GENESIS_REWARDS_TIMESTAMP % _CALCULATION_INTERVAL_SECONDS == 0, InvalidGenesisRewardsTimestampRemainder()
        );
        require(_CALCULATION_INTERVAL_SECONDS % SNAPSHOT_CADENCE == 0, InvalidCalculationIntervalSecondsRemainder());
        delegationManager = _delegationManager;
        strategyManager = _strategyManager;
        allocationManager = _allocationManager;
        CALCULATION_INTERVAL_SECONDS = _CALCULATION_INTERVAL_SECONDS;
        MAX_REWARDS_DURATION = _MAX_REWARDS_DURATION;
        MAX_RETROACTIVE_LENGTH = _MAX_RETROACTIVE_LENGTH;
        MAX_FUTURE_LENGTH = _MAX_FUTURE_LENGTH;
        GENESIS_REWARDS_TIMESTAMP = _GENESIS_REWARDS_TIMESTAMP;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[35] private __gap;
}
