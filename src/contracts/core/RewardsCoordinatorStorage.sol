// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IRewardsCoordinator.sol";

/**
 * @title Storage variables for the `RewardsCoordinator` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract RewardsCoordinatorStorage is IRewardsCoordinator {
    /**
     *
     *                            CONSTANTS AND IMMUTABLES
     *
     */

    // Constants for all reward types
    /// @notice The interval in seconds at which the calculation for rewards distribution is done.
    /// @dev ALL RewardsSubmission durations must be multiples of this interval. This is going to be configured to 1 week
    uint32 public immutable CALCULATION_INTERVAL_SECONDS;
    /// @notice The maximum amount of time (seconds) that ALL rewards submission can span over
    uint32 public immutable MAX_REWARDS_DURATION;
    /// @notice max amount of time (seconds) that ALL rewards submission can start in the future
    uint32 public immutable MAX_FUTURE_LENGTH;
    /// @notice The cadence at which a snapshot is taken offchain for calculating rewards distributions
    uint32 internal constant SNAPSHOT_CADENCE = 1 days;

    // V1 Rewards Constants
    /// @notice absolute min timestamp (seconds) that a V1 rewards submission can start at
    uint32 public immutable GENESIS_REWARDS_TIMESTAMP;
    /// @notice max amount of time (seconds) that a V1 rewards submission can start in the past
    uint32 public immutable MAX_RETROACTIVE_LENGTH;

    // OperatorSet Rewards Constants
    /// @notice absolute min timestamp (seconds) that an operatorSet rewards submission can start at
    uint32 public immutable OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP;
    /// @notice max of time (seconds) that an operatorSet rewards submission can start in the past
    uint32 public immutable OPERATOR_SET_MAX_RETROACTIVE_LENGTH;

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegationManager;

    /// @notice The StrategyManager contract for EigenLayer
    IStrategyManager public immutable strategyManager;

    /// @notice The AVSDirectory contract for EigenLayer
    IAVSDirectory public immutable avsDirectory;

    /**
     *
     *                                    STORAGE
     *
     */

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

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
    mapping(address => address) public claimerFor;

    /// @notice Mapping: earner => token => total amount claimed
    mapping(address => mapping(IERC20 => uint256)) public cumulativeClaimed;

    /// @notice Used for unique rewardsSubmissionHashes per AVS and for RewardsForAllSubmitters and the tokenHopper
    mapping(address => uint256) public submissionNonce;

    /// @notice Mapping: avs => avsRewardsSubmissionHash => bool to check if rewards submission hash has been submitted
    mapping(address => mapping(bytes32 => bool)) public isAVSRewardsSubmissionHash;

    /// @notice Mapping: rewardsSubmitterForAll => rewardsSubmissionForAllHash => bool to check if rewards submission hash for all has been submitted
    mapping(address => mapping(bytes32 => bool)) public isRewardsSubmissionForAllHash;

    /// @notice Mapping: address => bool to check if the address is permissioned to call createRewardsForAllSubmission
    mapping(address => bool) public isRewardsForAllSubmitter;
    /// @notice Mapping: operator => avs => operatorSetId => OperatorCommissionUpdate history
    mapping(address => mapping(address => mapping(uint32 => mapping(RewardType => OperatorCommissionUpdate[])))) public
        operatorCommissionUpdates;

    /// @notice Mapping: avs => rewardsSubmissionForAllEarnersHash => bool to check
    /// if rewards submission hash for all stakers and operators has been submitted
    mapping(address => mapping(bytes32 => bool)) public isRewardsSubmissionForAllEarnersHash;
    /// @notice Mapping: avs => rewardsSubmissionHash => bool to check if operatorSet rewards submission hash has been submitted
    mapping(address => mapping(bytes32 => bool)) isOperatorSetRewardsSubmissionHash;

    /// @notice Mapping: operator => avs => operatorSetId => OperatorCommissionUpdate history
    mapping(address => mapping(address => mapping(uint32 => mapping(RewardType => OperatorCommissionUpdate[])))) public
        operatorCommissionUpdates;

    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
        IAVSDirectory _avsDirectory,
        uint32 _CALCULATION_INTERVAL_SECONDS,
        uint32 _MAX_REWARDS_DURATION,
        uint32 _MAX_RETROACTIVE_LENGTH,
        uint32 _MAX_FUTURE_LENGTH,
        uint32 _GENESIS_REWARDS_TIMESTAMP,
        uint32 _OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP,
        uint32 _OPERATOR_SET_MAX_RETROACTIVE_LENGTH
    ) {
        require(
            _GENESIS_REWARDS_TIMESTAMP % _CALCULATION_INTERVAL_SECONDS == 0,
            "RewardsCoordinator: GENESIS_REWARDS_TIMESTAMP must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
        require(
            _CALCULATION_INTERVAL_SECONDS % SNAPSHOT_CADENCE == 0,
            "RewardsCoordinator: CALCULATION_INTERVAL_SECONDS must be a multiple of SNAPSHOT_CADENCE"
        );
        require(
            _OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP % _CALCULATION_INTERVAL_SECONDS == 0,
            "RewardsCoordinator: OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP must be a multiple of CALCULATION_INTERVAL_SECONDS"
        );
        delegationManager = _delegationManager;
        strategyManager = _strategyManager;
        avsDirectory = _avsDirectory;
        CALCULATION_INTERVAL_SECONDS = _CALCULATION_INTERVAL_SECONDS;
        MAX_REWARDS_DURATION = _MAX_REWARDS_DURATION;
        MAX_RETROACTIVE_LENGTH = _MAX_RETROACTIVE_LENGTH;
        MAX_FUTURE_LENGTH = _MAX_FUTURE_LENGTH;
        GENESIS_REWARDS_TIMESTAMP = _GENESIS_REWARDS_TIMESTAMP;
        OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP = _OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP;
        OPERATOR_SET_MAX_RETROACTIVE_LENGTH = _OPERATOR_SET_MAX_RETROACTIVE_LENGTH;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[37] private __gap;
}
