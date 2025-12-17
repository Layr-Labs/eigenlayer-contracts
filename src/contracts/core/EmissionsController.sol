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
        uint256 inflationRate,
        uint256 startTime,
        uint256 cooldownSeconds
    ) EmissionsControllerStorage(inflationRate, startTime, cooldownSeconds) {
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
        // Get the total number of distributions (also next available distribution id).
        distributionId = getTotalDistributions();
        // Append the distribution to the distributions array.
        _distributions.push(distribution);
        // Emit an event for the new distribution.
        emit DistributionAdded(distributionId, getCurrentEpoch(), distribution);
    }

    /// @inheritdoc IEmissionsController
    function updateDistribution(
        uint256 distributionId,
        Distribution calldata distribution
    ) external override onlyIncentiveCouncil {}

    /// @inheritdoc IEmissionsController
    function removeDistribution(
        uint256 distributionId
    ) external override onlyIncentiveCouncil {}

    /// -----------------------------------------------------------------------
    /// View
    /// -----------------------------------------------------------------------

    /// @inheritdoc IEmissionsController
    function getCurrentEpoch() public view returns (uint256) {
        if (block.timestamp < EMISSIONS_START_TIME) return type(uint256).max;
        return (block.timestamp - EMISSIONS_START_TIME) / EMISSIONS_EPOCH_LENGTH;
    }

    /// @inheritdoc IEmissionsController
    function isButtonPressable() external view returns (bool) {
        return _epochTriggered[getCurrentEpoch()];
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
