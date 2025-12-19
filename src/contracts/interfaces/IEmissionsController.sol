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
}

/// @title IEmissionsControllerTypes
/// @notice Types for the IEmissionsController contract.
interface IEmissionsControllerTypes {
    /// @notice Distribution types as defined in the ELIP.
    /// @dev Ref: "Distribution Submission types may include: createRewardsForAllEarners, createOperatorSetTotalStakeRewardsSubmission, createOperatorSetUniqueStakeRewardsSubmission, EigenDA Distribution, Manual Distribution."
    enum DistributionType {
        Disabled,
        RewardsForAllEarners,
        OperatorSetTotalStake,
        OperatorSetUniqueStake,
        EigenDA,
        Manual
    }

    /// @notice A Distribution structure containing weight, type and strategies.
    /// @dev Ref: "A Distribution consists of N fields: Weight, Distribution-type, Strategies and Multipliers."
    struct Distribution {
        /// The bips denominated weight of the distribution.
        uint256 weight;
        /// The epoch the distribution was last triggered.
        uint256 startEpoch;
        /// The number of epochs to repeat the distribution.
        uint256 stopEpoch;
        /// The type of distribution.
        DistributionType distributionType;
        /// The encoded rewards submission (either `RewardsSubmission` or `OperatorDirectedRewardsSubmission`).
        bytes encodedRewardsSubmission;
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
/// @notice Interface for the EmissionsController contract, which acts as the upgraded ActionGenerator.
/// @dev Ref: "This proposal requires upgrades to the TokenHopper and Action Generator contracts."
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

    /// @notice Triggers the weekly emissions.
    /// @dev Try/catch is used to prevent a single reverting rewards submission from halting emissions.
    /// @dev Pagination is used to prevent out-of-gas errors; multiple calls may be needed to process all submissions.
    /// @dev Ref: "The ActionGenerator today is a contract ... that is triggered by the Hopper. When triggered, it mints new EIGEN tokens..."
    /// @dev Permissionless function that can be called by anyone when `isButtonPressable()` returns true.
    /// @param length The number of distributions to process.
    function pressButton(
        uint256 length
    ) external;

    /// -----------------------------------------------------------------------
    /// Protocol Council Functions
    /// -----------------------------------------------------------------------

    /// @notice Sets the Incentive Council address.
    /// @dev Only the Protocol Council can call this function.
    /// @dev Ref: "Protocol Council Functions: Set Incentive Council multisig address that can interface with the ActionGenerator..."
    /// @param incentiveCouncil The new Incentive Council address.
    function setIncentiveCouncil(
        address incentiveCouncil
    ) external;

    /// -----------------------------------------------------------------------
    /// Incentive Council Functions
    /// -----------------------------------------------------------------------

    /// @notice Adds a new distribution.
    /// @dev Only the Incentive Council can call this function.
    /// @dev Ref: "Incentive Council Functions: addDistribution(weight{int}, distribution-type{see below}, strategiesAndMultipliers())"
    /// @param distribution The distribution to add.
    /// @return distributionId The id of the added distribution.
    function addDistribution(
        Distribution calldata distribution
    ) external returns (uint256 distributionId);

    /// @notice Updates an existing distribution.
    /// @dev Only the Incentive Council can call this function.
    /// @dev Ref: "Incentive Council Functions: updateDistribution(distributionId)"
    /// @param distributionId The id of the distribution to update.
    /// @param distribution The new distribution.
    function updateDistribution(
        uint256 distributionId,
        Distribution calldata distribution
    ) external;

    /// @notice Cancels a distribution.
    /// @dev The distribution remains in storage, but is marked as cancelled.
    /// @dev Only the Incentive Council can call this function.
    /// @dev Ref: Implied by "updateDistribution" and general management of distributions.
    /// @param distributionId The id of the distribution to remove.
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

    /// @notice Returns the next button press time.
    /// @return The next button press time.
    function nextButtonPressTime() external view returns (uint256);

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
