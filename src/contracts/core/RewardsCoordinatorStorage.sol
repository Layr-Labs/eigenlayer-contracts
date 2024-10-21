// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IRewardsCoordinator.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IStrategyManager.sol";

/**
 * @title Storage variables for the `RewardsCoordinator` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract RewardsCoordinatorStorage is IRewardsCoordinator {
    // Constants

    /// @notice The maximum rewards token amount for a single rewards submission, constrained by off-chain calculation
    uint256 internal constant MAX_REWARDS_AMOUNT = 1e38 - 1;

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

    /// @dev Salt for the earner leaf, meant to distinguish from tokenLeaf since they have the same sized data
    uint8 internal constant EARNER_LEAF_SALT = 0;
    /// @dev Salt for the token leaf, meant to distinguish from earnerLeaf since they have the same sized data
    uint8 internal constant TOKEN_LEAF_SALT = 1;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    // Immtuables

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegationManager;

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /// @notice The interval in seconds at which the calculation for rewards distribution is done.
    /// @dev RewardsSubmission durations must be multiples of this interval. This is going to be configured to 1 week
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

    bytes32 internal __deprecated_DOMAIN_SEPARATOR;

    /**
     * @notice List of roots submited by the rewardsUpdater
     * @dev Array is internal with an external getter so we can return a `DistributionRoot[] memory` object
     */
    DistributionRoot[] internal _distributionRoots;

    /// Slot 3
    /// @notice The address of the entity that can update the contract with new merkle roots
    address public rewardsUpdater;
    /// @notice Delay in timestamp (seconds) before a posted root can be claimed against
    uint32 public activationDelay;
    /// @notice Timestamp for last submitted DistributionRoot
    uint32 public currRewardsCalculationEndTimestamp;
    /// @notice the commission for all operators across all avss
    uint16 public globalOperatorCommissionBips;

    /// @notice Mapping: earner => the address of the entity who can call `processClaim` on behalf of the earner
    mapping(address earner => address) public claimerFor;

    /// @notice Mapping: earner => token => total amount claimed
    mapping(address earner => mapping(IERC20 token => uint256)) public cumulativeClaimed;

    /// @notice Used for unique rewardsSubmissionHashes per AVS and for RewardsForAllSubmitters and the tokenHopper
    mapping(address avs => uint256) public submissionNonce;

    /// @notice Mapping: avs => avsRewardsSubmissionHash => bool to check if rewards submission hash has been submitted
    mapping(address avs => mapping(bytes32 avsRewardsSubmissionHash => bool)) public isAVSRewardsSubmissionHash;

    /// @notice Mapping: avs => rewardsSubmissionForAllHash => bool to check if rewards submission hash for all has been submitted
    mapping(address avs => mapping(bytes32 rewardsSubmissionForAllHash => bool)) public isRewardsSubmissionForAllHash;

    /// @notice Mapping: address => bool to check if the address is permissioned to call createRewardsForAllSubmission
    mapping(address rewardsForAllSubmitter => bool) public isRewardsForAllSubmitter;

    /// @notice Mapping: avs => rewardsSubmissionForAllEarnersHash => bool to check
    /// if rewards submission hash for all stakers and operators has been submitted
    mapping(address avs => mapping(bytes32 rewardsSubmissionForAllEarnersHash => bool)) public
        isRewardsSubmissionForAllEarnersHash;

    // Construction

    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
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
    uint256[40] private __gap;
}
