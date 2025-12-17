// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "./storage/EmissionsControllerStorage.sol";

contract EmissionsController is Initializable, OwnableUpgradeable, EmissionsControllerStorage {
    /// @dev Modifier that checks if the caller is the incentive council.
    modifier onlyIncentiveCouncil() {
        // Check if the caller is the incentive council.
        require(msg.sender == incentiveCouncil, CallerIsNotIncentiveCouncil());
        _;
    }

    /// -----------------------------------------------------------------------
    /// Initialization
    /// -----------------------------------------------------------------------
    constructor(
        IEigen eigen,
        IRewardsCoordinator rewardsCoordinator,
        uint256 inflationRate,
        uint256 startTime,
        uint256 cooldownSeconds
    ) EmissionsControllerStorage(eigen, rewardsCoordinator, inflationRate, startTime, cooldownSeconds) {
        _disableInitializers();
    }

    /// @inheritdoc IEmissionsController
    function initialize(
        address initialOwner,
        address initialIncentiveCouncil
    ) external override initializer {
        // Set the initial owner.
        _transferOwnership(initialOwner);
        // Set the initial incentive council.
        _setIncentiveCouncil(initialIncentiveCouncil);
    }

    /// -----------------------------------------------------------------------
    /// Permissionless Trigger
    /// -----------------------------------------------------------------------

    /// @inheritdoc IEmissionsController
    function pressButton() external override {}

    /// -----------------------------------------------------------------------
    /// Owner Functions
    /// -----------------------------------------------------------------------

    /// @inheritdoc IEmissionsController
    function setIncentiveCouncil(
        address newIncentiveCouncil
    ) external override onlyOwner {
        _setIncentiveCouncil(newIncentiveCouncil);
    }

    /// @dev Internal helper to set the incentive council.
    function _setIncentiveCouncil(
        address newIncentiveCouncil
    ) internal {
        // Set the new incentive council.
        incentiveCouncil = newIncentiveCouncil;
        // Emit an event for the updated incentive council.
        emit IncentiveCouncilUpdated(newIncentiveCouncil);
    }

    /// -----------------------------------------------------------------------
    /// Incentive Council Functions
    /// -----------------------------------------------------------------------

    /// @inheritdoc IEmissionsController
    function addDistribution(
        Distribution calldata distribution
    ) external override onlyIncentiveCouncil returns (uint256 distributionId) {
        uint256 currentEpoch = getCurrentEpoch();

        // Question: How should be handle added distributions before the first epoch?

        // Check if the distribution is disabled.
        if (distribution.distributionType == DistributionType.Disabled) {
            revert CannotAddDisabledDistribution();
        }
        // Asserts the following:
        // - The start epoch is in the future.
        // - The total weight of all distributions does not exceed the max total weight.
        _checkDistribution(distribution, currentEpoch);
        // Get the total number of distributions (also next available distribution id).
        distributionId = getTotalDistributions();
        // Append the distribution to the distributions array.
        _distributions.push(distribution);
        // Emit an event for the new distribution.
        emit DistributionAdded(distributionId, currentEpoch, distribution);
    }

    /// @inheritdoc IEmissionsController
    function updateDistribution(
        uint256 distributionId,
        Distribution calldata distribution
    ) external override onlyIncentiveCouncil {
        uint256 currentEpoch = getCurrentEpoch();

        // Check if the distribution is disabled.
        if (distribution.distributionType == DistributionType.Disabled) {
            revert CannotUpdateDisabledDistribution();
        }

        // Check if the distribution is disabled.
        _checkDisabled(_distributions[distributionId]);
        // Asserts the following:
        // - The start epoch is in the future.
        // - The total weight of all distributions does not exceed the max total weight.
        _checkDistribution(distribution, currentEpoch);
        // Update the distribution in the distributions array.
        _distributions[distributionId] = distribution;
        // Emit an event for the updated distribution.
        emit DistributionUpdated(distributionId, currentEpoch, distribution);
    }

    /// @inheritdoc IEmissionsController
    function removeDistribution(
        uint256 distributionId
    ) external override onlyIncentiveCouncil {
        // Check if the distribution is already disabled.
        _checkDisabled(_distributions[distributionId]);
        // Set the distribution type to disabled.
        // Prevents further updates to the distribution.
        _distributions[distributionId].distributionType = DistributionType.Disabled;
        // Emit an event for the removed distribution.
        emit DistributionRemoved(distributionId, getCurrentEpoch());
    }

    /// -----------------------------------------------------------------------
    /// Internal Helpers
    /// -----------------------------------------------------------------------

    function _checkDisabled(
        Distribution storage distribution
    ) internal view {
        // Check if the distribution is disabled.
        if (distribution.distributionType == DistributionType.Disabled) {
            revert DistributionIsDisabled();
        }
    }

    function _checkDistribution(
        Distribution calldata distribution,
        uint256 currentEpoch
    ) internal view {
        // Check if the start epoch is in the future.
        // Prevents updating a distribution to a past or current epoch.
        if (distribution.startEpoch <= currentEpoch) {
            revert StartEpochMustBeInTheFuture();
        }

        // Check if the new total weight of all distributions exceeds max total weight.
        // Prevents distributing more supply than inflation rate allows.
        if (distribution.weight + totalWeight > MAX_TOTAL_WEIGHT) {
            revert TotalWeightExceedsMax();
        }
    }

    /// -----------------------------------------------------------------------
    /// View
    /// -----------------------------------------------------------------------

    /// @inheritdoc IEmissionsController
    function getCurrentEpoch() public view returns (uint256) {
        // If the start time has not elapsed, default to max uint256.
        if (block.timestamp < EMISSIONS_START_TIME) return type(uint256).max;
        // Calculate the current epoch by dividing the time since start by the epoch length.
        return (block.timestamp - EMISSIONS_START_TIME) / EMISSIONS_EPOCH_LENGTH;
    }

    /// @inheritdoc IEmissionsController
    function isButtonPressable() external view returns (bool) {
        return !_epochTriggered[getCurrentEpoch()];
    }

    /// @inheritdoc IEmissionsController
    function nextButtonPressTime() external view returns (uint256) {
        return EMISSIONS_START_TIME + EMISSIONS_EPOCH_LENGTH * (getCurrentEpoch() + 1);
    }

    /// @inheritdoc IEmissionsController
    function getTotalDistributions() public view returns (uint256) {
        return _distributions.length;
    }

    /// @inheritdoc IEmissionsController
    function getDistribution(
        uint256 distributionId
    ) external view returns (Distribution memory) {
        return _distributions[distributionId];
    }

    /// @inheritdoc IEmissionsController
    function getDistributions(
        uint256 start,
        uint256 length
    ) external view returns (Distribution[] memory distributions) {
        // Create a new array in memory with the given length.
        distributions = new Distribution[](length);
        // Copy the specified subset of distributions from the storage array to the memory array.
        for (uint256 i = 0; i < length; ++i) {
            distributions[i] = _distributions[start + i];
        }
    }
}
