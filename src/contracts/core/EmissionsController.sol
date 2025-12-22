// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "../libraries/OperatorSetLib.sol";
import "./storage/EmissionsControllerStorage.sol";

// Add fee, it's enabled by default, with settable fee recipient.
// AVS's must also have a way to opt-out of the fee.

contract EmissionsController is Initializable, OwnableUpgradeable, EmissionsControllerStorage {
    /// @dev Modifier that checks if the caller is the incentive council.
    modifier onlyIncentiveCouncil() {
        _checkOnlyIncentiveCouncil();
        _;
    }

    /// -----------------------------------------------------------------------
    /// Initialization
    /// -----------------------------------------------------------------------
    constructor(
        IEigen eigen,
        IBackingEigen backingEigen,
        IRewardsCoordinator rewardsCoordinator,
        uint256 inflationRate,
        uint256 startTime,
        uint256 cooldownSeconds
    ) EmissionsControllerStorage(eigen, backingEigen, rewardsCoordinator, inflationRate, startTime, cooldownSeconds) {
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

        // NOTE: Reviewers, is this okay? Avoids the need for approvals on each button press.

        // Max approve EIGEN for spending bEIGEN.
        BACKING_EIGEN.approve(address(EIGEN), type(uint256).max);
        // Max approve RewardsCoordinator for spending EIGEN.
        EIGEN.approve(address(REWARDS_COORDINATOR), type(uint256).max);
    }

    /// -----------------------------------------------------------------------
    /// Permissionless Trigger
    /// -----------------------------------------------------------------------

    /// @inheritdoc IEmissionsController
    function pressButton(
        uint256 length
    ) external override {
        uint256 currentEpoch = getCurrentEpoch();
        uint256 totalDistributions = getTotalDistributions();
        uint256 nextDistributionId = _epochs[currentEpoch].totalProcessed;

        // Check if all distributions have already been processed.
        if (nextDistributionId >= totalDistributions) {
            revert AllDistributionsProcessed();
        }

        // Mint the total amount of bEIGEN/EIGEN needed for all distributions.
        if (!_epochs[currentEpoch].minted) {
            // First mint the bEIGEN in order to wrap it into EIGEN.
            BACKING_EIGEN.mint(address(this), EMISSIONS_INFLATION_RATE);
            // Then wrap it into EIGEN.
            EIGEN.wrap(EMISSIONS_INFLATION_RATE);

            // Mark the epoch as minted.
            _epochs[currentEpoch].minted = true;
        }

        // Calculate the start timestamp for the distribution (equivalent to `lastTimeButtonPressable()`).
        uint256 startTimestamp = EMISSIONS_START_TIME + EMISSIONS_EPOCH_LENGTH * currentEpoch;

        // Process distributions starting from the next one to process
        for (uint256 i = nextDistributionId; i < length; ++i) {
            Distribution memory distribution = _distributions[i];

            // Skip disabled distributions...
            if (distribution.distributionType == DistributionType.Disabled) continue;
            // Skip distributions that haven't started yet...
            if (distribution.startEpoch > currentEpoch) continue;
            // Skip distributions that have ended...
            if (distribution.stopEpoch <= currentEpoch) continue;

            _processDistribution(i, currentEpoch, startTimestamp, distribution);
        }

        // QUESTION: Should be burn excess supply that wasn't distributed (e.g. if the distribution reverts)?

        // Update total processed count for this epoch (equals next index to process)
        _epochs[currentEpoch].totalProcessed = uint248(length);
    }

    /// @dev Internal helper that processes a distribution.
    /// @param distributionId The id of the distribution to process.
    /// @param currentEpoch The current epoch.
    /// @param distribution The distribution (from storage).
    function _processDistribution(
        uint256 distributionId,
        uint256 currentEpoch,
        uint256 startTimestamp,
        Distribution memory distribution
    ) internal {
        // Calculate the total amount of emissions for the distribution.
        uint256 totalAmount = EMISSIONS_INFLATION_RATE * distribution.weight / MAX_TOTAL_WEIGHT;
        // Calculate the amount per submission (total amount divided by the number of submissions).
        uint256 amountPerSubmission = totalAmount / distribution.partialRewardsSubmissions.length;

        // Update the rewards submissions start timestamp, duration, and amount.
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions =
            new IRewardsCoordinator.RewardsSubmission[](distribution.partialRewardsSubmissions.length);
        for (uint256 i = 0; i < rewardsSubmissions.length; ++i) {
            rewardsSubmissions[i] = IRewardsCoordinatorTypes.RewardsSubmission({
                strategiesAndMultipliers: distribution.partialRewardsSubmissions[i].strategiesAndMultipliers,
                token: distribution.partialRewardsSubmissions[i].token,
                amount: amountPerSubmission,
                startTimestamp: uint32(startTimestamp),
                duration: uint32(EMISSIONS_EPOCH_LENGTH) // QUESTION: Should this be configurable?
            });
        }

        // Dispatch the `RewardsCoordinator` call based on the distribution type.
        bool success;
        if (distribution.distributionType == DistributionType.RewardsForAllEarners) {
            success = _tryCallRewardsCoordinator(
                abi.encodeCall(IRewardsCoordinator.createRewardsForAllEarners, (rewardsSubmissions))
            );
        } else if (distribution.distributionType == DistributionType.OperatorSetUniqueStake) {
            success = _tryCallRewardsCoordinator(
                abi.encodeCall(
                    IRewardsCoordinator.createUniqueStakeRewardsSubmission,
                    (distribution.operatorSet, rewardsSubmissions)
                )
            );
        } else if (distribution.distributionType == DistributionType.OperatorSetTotalStake) {
            success = _tryCallRewardsCoordinator(
                abi.encodeCall(
                    IRewardsCoordinator.createTotalStakeRewardsSubmission,
                    (distribution.operatorSet, rewardsSubmissions)
                )
            );
        } else if (distribution.distributionType == DistributionType.EigenDA) {
            success = _tryCallRewardsCoordinator(
                abi.encodeCall(IRewardsCoordinator.createAVSRewardsSubmission, (rewardsSubmissions))
            );
        } else if (distribution.distributionType == DistributionType.Manual) {
            // We use call to prevent further distributions from being processed if the mint fails.
            (success,) =
                address(BACKING_EIGEN).call(abi.encodeWithSelector(IBackingEigen.mint.selector, rewardsSubmissions));
        } else {
            revert InvalidDistributionType(); // Only reachable if the distribution type is `Disabled`.
        }

        // Emit an event for the processed distribution.
        emit DistributionProcessed(distributionId, currentEpoch, distribution, success);
    }

    /// @dev Internal helper that try/calls the RewardsCoordinator returning success or failure.
    /// This is needed as using try/catch requires decoding the calldata, which can revert preventing further distributions.
    /// @param abiEncodedCall The ABI encoded call to the RewardsCoordinator.
    /// @return success True if the function call was successful, false otherwise.
    function _tryCallRewardsCoordinator(
        bytes memory abiEncodedCall
    ) internal returns (bool success) {
        (success,) = address(REWARDS_COORDINATOR).call(abiEncodedCall);
    }

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

        // operatorSet only non-zero if distribution type is OperatorSetTotalStake or OperatorSetUniqueStake.
        // Asserts the following:
        // - The start epoch is in the future.
        // - The total weight of all distributions does not exceed the max total weight.
        _checkDistribution(distribution, currentEpoch);
        // Get the total number of distributions (also next available distribution id).
        distributionId = getTotalDistributions();

        // Append the distribution to the distributions array.
        _distributions.push(distribution);
        // Update the total weight.
        totalWeight += distribution.weight;
        // Emit an event for the new distribution.
        emit DistributionAdded(distributionId, currentEpoch, distribution);
    }

    /// @inheritdoc IEmissionsController
    function updateDistribution(
        uint256 distributionId,
        Distribution calldata distribution
    ) external override onlyIncentiveCouncil {
        uint256 currentEpoch = getCurrentEpoch();

        // Check if the distribution (from calldata) is disabled.
        if (distribution.distributionType == DistributionType.Disabled) {
            revert CannotDisableDistributionViaUpdate();
        }

        // Check if the distribution (in storage) is disabled.
        _checkDisabled(_distributions[distributionId]);

        uint256 weight = _distributions[distributionId].weight;
        // Subtract the old weight.
        totalWeight -= weight;
        // Add the new weight.
        totalWeight += distribution.weight;

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
    function disableDistribution(
        uint256 distributionId
    ) external override onlyIncentiveCouncil {
        // Check if the distribution is already disabled.
        _checkDisabled(_distributions[distributionId]);
        // Subtract the weight from total weight.
        totalWeight -= _distributions[distributionId].weight;
        // Set the distribution type to disabled.
        // Prevents further updates to the distribution.
        _distributions[distributionId].distributionType = DistributionType.Disabled;
        // Emit an event for the removed distribution.
        emit DistributionRemoved(distributionId, getCurrentEpoch());
    }

    /// -----------------------------------------------------------------------
    /// Internal Helpers
    /// -----------------------------------------------------------------------

    function _checkOnlyIncentiveCouncil() internal view {
        // Check if the caller is the incentive council.
        if (msg.sender != incentiveCouncil) revert CallerIsNotIncentiveCouncil();
    }

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
        if (distribution.startEpoch == currentEpoch) {
            revert StartEpochMustBeInTheFuture();
        }

        // Check if the new total weight of all distributions exceeds max total weight.
        // Prevents distributing more supply than inflation rate allows.
        if (distribution.weight + totalWeight > MAX_TOTAL_WEIGHT) {
            revert TotalWeightExceedsMax();
        }

        // Check if rewards submissions array is empty for non-Manual distributions.
        // Manual distributions handle rewards differently and don't require submissions.
        if (
            distribution.distributionType != DistributionType.Manual
                && distribution.partialRewardsSubmissions.length == 0
        ) {
            revert RewardsSubmissionsCannotBeEmpty();
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
        return _epochs[getCurrentEpoch()].totalProcessed < getTotalDistributions();
    }

    /// @inheritdoc IEmissionsController
    function nextTimeButtonPressable() external view returns (uint256) {
        return EMISSIONS_START_TIME + EMISSIONS_EPOCH_LENGTH * (getCurrentEpoch() + 1);
    }

    /// @inheritdoc IEmissionsController
    function lastTimeButtonPressable() public view returns (uint256) {
        return EMISSIONS_START_TIME + EMISSIONS_EPOCH_LENGTH * getCurrentEpoch();
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
