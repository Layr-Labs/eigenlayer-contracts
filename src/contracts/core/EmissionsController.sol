// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../libraries/OperatorSetLib.sol";
import "../permissions/Pausable.sol";
import "./storage/EmissionsControllerStorage.sol";

contract EmissionsController is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    ReentrancyGuardUpgradeable,
    EmissionsControllerStorage
{
    using SafeERC20 for IERC20;

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
        IPauserRegistry pauserRegistry,
        uint256 inflationRate,
        uint256 startTime,
        uint256 cooldownSeconds
    )
        EmissionsControllerStorage(eigen, backingEigen, rewardsCoordinator, inflationRate, startTime, cooldownSeconds)
        Pausable(pauserRegistry)
    {
        _disableInitializers();
    }

    /// @inheritdoc IEmissionsController
    function initialize(
        address initialOwner,
        address initialIncentiveCouncil,
        uint256 initialPausedStatus
    ) external override initializer {
        // Set the initial owner.
        _transferOwnership(initialOwner);
        // Set the initial incentive council.
        _setIncentiveCouncil(initialIncentiveCouncil);
        // Set the initial paused status.
        _setPausedStatus(initialPausedStatus);
    }

    /// -----------------------------------------------------------------------
    /// Permissionless Trigger
    /// -----------------------------------------------------------------------

    /// @inheritdoc IEmissionsController
    function sweep() external override nonReentrant onlyWhenNotPaused(PAUSED_TOKEN_FLOWS) {
        uint256 amount = EIGEN.balanceOf(address(this));
        if (!isButtonPressable() && amount != 0) {
            IERC20(EIGEN).safeTransfer(incentiveCouncil, amount);
            emit Swept(incentiveCouncil, amount);
        }
    }

    /// @inheritdoc IEmissionsController
    function pressButton(
        uint256 length
    ) external override nonReentrant onlyWhenNotPaused(PAUSED_TOKEN_FLOWS) {
        uint256 currentEpoch = getCurrentEpoch();
        uint256 totalDistributions = getTotalDistributions();
        uint256 nextDistributionId = _epochs[currentEpoch].totalProcessed;

        // Check if all distributions have already been processed.
        if (nextDistributionId >= totalDistributions) {
            revert AllDistributionsProcessed();
        }

        // Mint the total amount of bEIGEN/EIGEN needed for all distributions.
        if (!_epochs[currentEpoch].minted) {
            // NOTE: Approvals may not be entirely spent.

            // Max approve EIGEN for spending bEIGEN.
            BACKING_EIGEN.approve(address(EIGEN), EMISSIONS_INFLATION_RATE);
            // Max approve RewardsCoordinator for spending EIGEN.
            EIGEN.approve(address(REWARDS_COORDINATOR), EMISSIONS_INFLATION_RATE);

            // First mint the bEIGEN in order to wrap it into EIGEN.
            BACKING_EIGEN.mint(address(this), EMISSIONS_INFLATION_RATE);
            // Then wrap it into EIGEN.
            EIGEN.wrap(EMISSIONS_INFLATION_RATE);

            // Mark the epoch as minted.
            _epochs[currentEpoch].minted = true;
        }

        // Calculate the start timestamp for the distribution (equivalent to `lastTimeButtonPressable()`).
        uint256 startTimestamp = EMISSIONS_START_TIME + EMISSIONS_EPOCH_LENGTH * currentEpoch;
        // Calculate the last index to process.
        uint256 lastIndex = nextDistributionId + length;

        // If length exceeds total distributions, set last index to total distributions (exclusive upper bound).
        if (lastIndex > totalDistributions) lastIndex = totalDistributions;

        // Process distributions starting from the next one to process...
        for (uint256 i = nextDistributionId; i < lastIndex; ++i) {
            Distribution memory distribution = _distributions[i];

            // Skip disabled distributions...
            if (distribution.distributionType == DistributionType.Disabled) continue;
            // Skip distributions that haven't started yet...
            if (distribution.startEpoch > currentEpoch) continue;
            // Skip distributions that have ended...
            if (distribution.stopEpoch <= currentEpoch) continue;

            _processDistribution(i, currentEpoch, startTimestamp, distribution);
        }

        // Update total processed count for this epoch.
        _epochs[currentEpoch].totalProcessed = uint248(lastIndex);
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
        // Success flag for the distribution.
        bool success;

        if (distribution.distributionType != DistributionType.Manual) {
            // Calculate the amount per submission.
            uint256 amountPerSubmission = totalAmount / distribution.strategiesAndMultipliers.length;
            // Update the rewards submissions start timestamp, duration, and amount.
            IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions =
                new IRewardsCoordinator.RewardsSubmission[](distribution.strategiesAndMultipliers.length);
            for (uint256 i = 0; i < rewardsSubmissions.length; ++i) {
                rewardsSubmissions[i] = IRewardsCoordinatorTypes.RewardsSubmission({
                    strategiesAndMultipliers: distribution.strategiesAndMultipliers[i],
                    token: EIGEN,
                    amount: amountPerSubmission,
                    startTimestamp: uint32(startTimestamp),
                    duration: uint32(EMISSIONS_EPOCH_LENGTH)
                });
            }

            // Dispatch the `RewardsCoordinator` call based on the distribution type.
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
            } else {
                revert InvalidDistributionType(); // Only reachable if the distribution type is `Disabled`.
            }
        } else {
            (success,) =
                address(EIGEN).call(abi.encodeWithSelector(IERC20.transfer.selector, incentiveCouncil, totalAmount));
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
        // Checks

        uint256 currentEpoch = getCurrentEpoch();

        // Check if the distribution is disabled.
        if (distribution.distributionType == DistributionType.Disabled) {
            revert CannotAddDisabledDistribution();
        }

        uint256 totalWeightBefore = totalWeight;

        // Asserts the following:
        // - The start epoch is in the future.
        // - The total weight of all distributions does not exceed the max total weight.
        _checkDistribution(distribution, currentEpoch, totalWeightBefore);

        // Effects

        // Update the total weight.
        totalWeight = totalWeightBefore + distribution.weight;
        // Append the distribution to the distributions array.
        _distributions.push(distribution);
        // Emit an event for the new distribution.
        emit DistributionAdded(distributionId, currentEpoch, distribution);

        // Return distribution id (also next available distribution id).
        return getTotalDistributions();
    }

    /// @inheritdoc IEmissionsController
    function updateDistribution(
        uint256 distributionId,
        Distribution calldata distribution
    ) external override onlyIncentiveCouncil {
        // Checks

        uint256 currentEpoch = getCurrentEpoch();

        // Check if the distribution (from calldata) is disabled.
        if (distribution.distributionType == DistributionType.Disabled) {
            revert CannotDisableDistributionViaUpdate();
        }

        // Check if the distribution (in storage) is disabled.
        _checkDisabled(_distributions[distributionId]);

        uint256 totalWeightBefore = totalWeight;
        uint256 weight = _distributions[distributionId].weight;

        // Asserts the following:
        // - The start epoch is in the future.
        // - The total weight of all distributions does not exceed the max total weight.
        _checkDistribution(distribution, currentEpoch, totalWeightBefore - weight);

        // Effects

        // Add the new weight.
        totalWeight = totalWeightBefore - weight + distribution.weight;
        // Update the distribution in the distributions array.
        _distributions[distributionId] = distribution;
        // Emit an event for the updated distribution.
        emit DistributionUpdated(distributionId, currentEpoch, distribution);
    }

    /// @inheritdoc IEmissionsController
    function disableDistribution(
        uint256 distributionId
    ) external override onlyIncentiveCouncil {
        // Checks

        // Check if the distribution is already disabled.
        _checkDisabled(_distributions[distributionId]);

        // Effects

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
        uint256 currentEpoch,
        uint256 totalWeightBefore
    ) internal pure {
        // Check if the start epoch is in the future.
        // Prevents updating a distribution to a past or current epoch.
        if (currentEpoch != type(uint256).max) {
            // After emissions start - require future epochs only
            if (distribution.startEpoch <= currentEpoch) {
                revert StartEpochMustBeInTheFuture();
            }
        }

        // Check if the new total weight of all distributions exceeds max total weight.
        // Prevents distributing more supply than inflation rate allows.
        if (distribution.weight + totalWeightBefore > MAX_TOTAL_WEIGHT) {
            revert TotalWeightExceedsMax();
        }

        // Check if rewards submissions array is empty for non-Manual distributions.
        // Manual distributions handle rewards differently and don't require submissions.
        if (
            distribution.distributionType != DistributionType.Manual
                && distribution.strategiesAndMultipliers.length == 0
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
    function isButtonPressable() public view returns (bool) {
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
