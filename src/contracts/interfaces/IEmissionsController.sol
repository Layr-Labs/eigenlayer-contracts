// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IRewardsCoordinator.sol";
import "./IEigen.sol";
import "./IBackingEigen.sol";

/// @title IEmissionsControllerErrors
/// @notice Errors for the IEmissionsController contract.
interface IEmissionsControllerErrors {
    /// @dev Thrown when caller is not the incentive council.
    error CallerIsNotIncentiveCouncil();
    /// @dev Thrown when the start epoch is the current or past epoch.
    error StartEpochMustBeInTheFuture();
    /// @dev Thrown when the total weight of all distributions exceeds the max total weight (100%).
    error TotalWeightExceedsMax();
    /// @dev Thrown when the distribution is disabled.
    error DistributionIsDisabled();
    /// @dev Thrown when attempting to add a disabled distribution.
    error CannotAddDisabledDistribution();
    /// @dev Thrown when attempting to update a disabled distribution.
    error CannotDisableDistributionViaUpdate();
    /// @dev Thrown when all distributions have been processed for the current epoch.
    error AllDistributionsProcessed();
    /// @dev Thrown when the distribution type is invalid. Should be unreachable.
    error InvalidDistributionType();
    /// @dev Thrown when rewards submissions array is empty for a distribution that requires it.
    error RewardsSubmissionsCannotBeEmpty();
}

/// @title IEmissionsControllerTypes
/// @notice Types for the IEmissionsController contract.
interface IEmissionsControllerTypes {
    /// @notice Distribution types that determine how emissions are routed to the RewardsCoordinator.
    /// @dev Each type maps to a specific RewardsCoordinator function or minting mechanism.
    ///      - Disabled: Distribution is inactive and skipped during processing
    ///      - RewardsForAllEarners: Calls `createRewardsForAllEarners` for protocol-wide rewards
    ///      - OperatorSetTotalStake: Calls `createTotalStakeRewardsSubmission` for operator set rewards weighted by total stake
    ///      - OperatorSetUniqueStake: Calls `createUniqueStakeRewardsSubmission` for operator set rewards weighted by unique stake
    ///      - EigenDA: Calls `createAVSRewardsSubmission` for EigenDA-specific rewards
    ///      - Manual: Directly mints bEIGEN to specified recipients without RewardsCoordinator interaction
    enum DistributionType {
        Disabled,
        RewardsForAllEarners,
        OperatorSetTotalStake,
        OperatorSetUniqueStake,
        EigenDA,
        Manual
    }

    /// @notice A struct containing the total minted and processed amounts for an epoch.
    struct Epoch {
        /// Whether the epoch has been minted.
        bool minted;
        /// The total number of distributions processed for the epoch.
        uint248 totalProcessed;
    }

    /// @notice A Distribution structure defining how a portion of emissions should be allocated.
    /// @dev Distributions are stored in an append-only array and processed each epoch by `pressButton`.
    ///      The weight determines the proportion of total emissions allocated to this distribution.
    ///      Active distributions are processed sequentially.
    struct Distribution {
        /// The bips denominated weight of the distribution.
        uint64 weight;
        /// The epoch the distribution was last triggered ().
        uint64 startEpoch;
        /// The number of epochs to repeat the distribution (1256 years supported).
        uint64 stopEpoch;
        /// The type of distribution.
        DistributionType distributionType;
        /// The operator set (Required only for OperatorSetTotalStake and OperatorSetUniqueStake distribution types).
        OperatorSet operatorSet;
        /// The strategies and their respective multipliers for distributing rewards.
        IRewardsCoordinatorTypes.StrategyAndMultiplier[][] strategiesAndMultipliers;
    }
}

/// @title IEmissionsControllerEvents
/// @notice Events for the IEmissionsController contract.
interface IEmissionsControllerEvents is IEmissionsControllerTypes {
    /// @notice Emitted when a distribution is updated.
    /// @param distributionId The id of the distribution.
    /// @param epoch The epoch the distribution was updated.
    /// @param distribution The distribution.
    event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, Distribution distribution);

    /// @notice Emitted when a distribution is added.
    /// @param distributionId The id of the distribution.
    /// @param epoch The epoch the distribution was added.
    /// @param distribution The distribution.
    event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, Distribution distribution);

    /// @notice Emitted when a distribution is removed.
    /// @param distributionId The id of the distribution.
    /// @param epoch The epoch the distribution was removed.
    event DistributionRemoved(uint256 indexed distributionId, uint256 indexed epoch);

    /// @notice Emitted when a distribution is processed.
    /// @param distributionId The id of the distribution.
    /// @param epoch The epoch the distribution was processed.
    /// @param distribution The distribution.
    event DistributionProcessed(
        uint256 indexed distributionId,
        uint256 indexed epoch,
        Distribution distribution,
        bool success
    );

    /// @notice Emitted when the Incentive Council address is updated.
    /// @param incentiveCouncil The new Incentive Council address.
    event IncentiveCouncilUpdated(address indexed incentiveCouncil);
}

/// @title IEmissionsController
/// @notice Interface for the EmissionsController contract that manages programmatic EIGEN emissions.
/// @dev The EmissionsController mints EIGEN at a fixed inflation rate per epoch and distributes it
///      according to configured distributions. It replaces the legacy ActionGenerator pattern with
///      a more flexible distribution system controlled by the Incentive Council.
interface IEmissionsController is IEmissionsControllerErrors, IEmissionsControllerEvents {
    /// -----------------------------------------------------------------------
    /// Constants
    /// -----------------------------------------------------------------------
    /// @notice The EIGEN token address.
    /// @dev Immutable/constant variable that requires an upgrade to modify.
    function EIGEN() external view returns (IEigen);

    /// @notice The BACKING_EIGEN token address.
    /// @dev Immutable/constant variable that requires an upgrade to modify.
    function BACKING_EIGEN() external view returns (IBackingEigen);

    /// @notice The RewardsCoordinator address.
    /// @dev Immutable/constant variable that requires an upgrade to modify.
    function REWARDS_COORDINATOR() external view returns (IRewardsCoordinator);

    /// @notice The max total weight of all distributions.
    /// @dev Constant variable that requires an upgrade to modify.
    function MAX_TOTAL_WEIGHT() external view returns (uint256);

    /// @notice The rate of inflation for emissions.
    /// @dev Immutable/constant variable that requires an upgrade to modify.
    function EMISSIONS_INFLATION_RATE() external view returns (uint256);

    /// @notice The start time of the emissions.
    /// @dev Immutable/constant variable that requires an upgrade to modify.
    function EMISSIONS_START_TIME() external view returns (uint256);

    /// @notice The cooldown seconds of the emissions.
    /// @dev Immutable/constant variable that requires an upgrade to modify.
    function EMISSIONS_EPOCH_LENGTH() external view returns (uint256);

    /// -----------------------------------------------------------------------
    /// Initialization Functions
    /// -----------------------------------------------------------------------

    /// @notice Initializes the contract.
    /// @param initialOwner The initial owner address.
    /// @param incentiveCouncil The initial Incentive Council address.
    function initialize(
        address initialOwner,
        address incentiveCouncil
    ) external;

    /// -----------------------------------------------------------------------
    /// Permissionless Trigger
    /// -----------------------------------------------------------------------

    /// @notice Triggers emissions for the current epoch and processes distributions.
    /// @dev This function mints EMISSIONS_INFLATION_RATE of bEIGEN, wraps it to EIGEN, then processes
    ///      distributions from _totalProcessed[currentEpoch] up to the specified length.
    ///      Each distribution receives a proportional amount based on its weight and submits to RewardsCoordinator.
    ///      Uses low-level calls to prevent reverts in one distribution from blocking others.
    ///      Can be called multiple times per epoch to paginate through all distributions if needed.
    /// @dev Permissionless function that can be called by anyone when `isButtonPressable()` returns true.
    /// @param length The number of distributions to process (upper bound for the loop).
    function pressButton(
        uint256 length
    ) external;

    /// -----------------------------------------------------------------------
    /// Protocol Council Functions
    /// -----------------------------------------------------------------------

    /// @notice Sets the Incentive Council address.
    /// @dev Only the contract owner (Protocol Council) can call this function.
    ///      The Incentive Council has exclusive rights to add, update, and disable distributions.
    ///      This separation of powers allows the Protocol Council to delegate distribution management
    ///      without giving up ownership of the contract.
    /// @param incentiveCouncil The new Incentive Council address.
    function setIncentiveCouncil(
        address incentiveCouncil
    ) external;

    /// -----------------------------------------------------------------------
    /// Incentive Council Functions
    /// -----------------------------------------------------------------------

    /// @notice Adds a new distribution to the emissions schedule.
    /// @dev Only the Incentive Council can call this function.
    ///      The distribution is appended to the _distributions array and assigned the next available ID.
    ///      The distribution's weight is added to totalWeight, which must not exceed MAX_TOTAL_WEIGHT.
    ///      The startEpoch must be in the future to prevent immediate processing.
    ///      Cannot add distributions with DistributionType.Disabled.
    /// @param distribution The distribution to add.
    /// @return distributionId The id of the added distribution.
    function addDistribution(
        Distribution calldata distribution
    ) external returns (uint256 distributionId);

    /// @notice Updates an existing distribution's parameters.
    /// @dev Only the Incentive Council can call this function.
    ///      Replaces the entire distribution at the given ID with the new distribution.
    ///      Adjusts totalWeight by subtracting the old weight and adding the new weight.
    ///      Cannot update distributions that are already disabled.
    ///      Cannot use this function to disable a distribution (must use disableDistribution).
    /// @param distributionId The id of the distribution to update.
    /// @param distribution The new distribution parameters.
    function updateDistribution(
        uint256 distributionId,
        Distribution calldata distribution
    ) external;

    /// @notice Disables a distribution permanently.
    /// @dev The distribution remains in storage but is marked as Disabled and skipped during processing.
    ///      Only the Incentive Council can call this function.
    ///      The distribution's weight is subtracted from totalWeight, freeing up allocation capacity.
    ///      This action is permanent - disabled distributions cannot be re-enabled or updated.
    ///      The distribution will be skipped in all future epochs regardless of startEpoch/stopEpoch.
    /// @param distributionId The id of the distribution to disable.
    function disableDistribution(
        uint256 distributionId
    ) external;

    /// -----------------------------------------------------------------------
    /// View
    /// -----------------------------------------------------------------------

    /// @notice Returns the current Incentive Council address.
    /// @return The Incentive Council address.
    function incentiveCouncil() external view returns (address);

    /// @notice Returns the total weight of all distributions.
    /// @return The total weight of all distributions.
    function totalWeight() external view returns (uint256);

    /// @notice Returns the current epoch.
    /// @return The current epoch.
    function getCurrentEpoch() external view returns (uint256);

    /// @notice Checks if the emissions can be triggered.
    /// @return True if the cooldown has passed and the system is ready.
    function isButtonPressable() external view returns (bool);

    /// @notice Returns the next time the button will be pressable.
    /// @return The next time the button will be pressable.
    function nextTimeButtonPressable() external view returns (uint256);

    /// @notice Returns the last time the button was pressable.
    /// @return The last time the button was pressable.
    function lastTimeButtonPressable() external view returns (uint256);

    /// @notice Returns the total number of distributions.
    /// @return The total number of distributions.
    function getTotalDistributions() external view returns (uint256);

    /// @notice Returns a distribution by index.
    /// @param distributionId The id of the distribution.
    /// @return The Distribution struct at the given index.
    function getDistribution(
        uint256 distributionId
    ) external view returns (Distribution memory);

    /// @notice Returns a subset of distributions.
    /// @param start The start index of the distributions.
    /// @param length The length of the distributions.
    /// @return An append-only array of Distribution structs.
    function getDistributions(
        uint256 start,
        uint256 length
    ) external view returns (Distribution[] memory);
}
