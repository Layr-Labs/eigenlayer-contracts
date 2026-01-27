// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IEmissionsController.sol";
import "src/test/utils/Logger.t.sol";
import "src/test/integration/TimeMachine.t.sol";

contract IncentiveCouncil is Logger, IEmissionsControllerTypes {
    using print for *;

    modifier createSnapshot() {
        timeMachine.createSnapshot();
        _;
    }

    IEmissionsController immutable emissionsController;
    address immutable incentiveCouncil;
    TimeMachine immutable timeMachine;

    constructor(IEmissionsController _emissionsController, address _incentiveCouncil, TimeMachine _timeMachine) {
        emissionsController = _emissionsController;
        incentiveCouncil = _incentiveCouncil;
        timeMachine = _timeMachine;
    }

    function NAME() public view override returns (string memory) {
        return "IncentiveCouncil";
    }

    function _randomStrategiesAndMultipliers(IStrategy[] memory strategies)
        internal
        returns (IRewardsCoordinatorTypes.StrategyAndMultiplier[][] memory strategiesAndMultipliers)
    {
        uint numStrategies = vm.randomUint({min: 1, max: strategies.length});
        strategiesAndMultipliers = new IRewardsCoordinatorTypes.StrategyAndMultiplier[][](numStrategies);
        uint multiplierLeft = 1e18;

        for (uint i = 0; i < numStrategies; i++) {
            strategiesAndMultipliers[i] = new IRewardsCoordinatorTypes.StrategyAndMultiplier[](1);
            uint multiplier;

            // For all strategies except the last, randomly assign multiplier
            // ensuring we leave enough for remaining strategies (at least 1 per strategy)
            if (i < numStrategies - 1) {
                uint remainingStrategies = numStrategies - i - 1;
                uint minMultiplierNeeded = remainingStrategies; // Reserve at least 1 for each remaining strategy
                uint maxAvailable = multiplierLeft > minMultiplierNeeded ? multiplierLeft - minMultiplierNeeded : 1;

                // Ensure max is at least min (1) to avoid underflow
                if (maxAvailable < 1) maxAvailable = 1;

                multiplier = vm.randomUint({min: 1, max: maxAvailable});
                multiplierLeft -= multiplier;
            } else {
                // Last strategy gets all remaining multiplier
                multiplier = multiplierLeft;
                multiplierLeft = 0;
            }

            strategiesAndMultipliers[i][0] =
                IRewardsCoordinatorTypes.StrategyAndMultiplier({strategy: strategies[i], multiplier: uint96(multiplier)});
        }
        return strategiesAndMultipliers;
    }

    function _randomOperatorSet(OperatorSet[] memory operatorSets) internal returns (OperatorSet memory operatorSet) {
        operatorSet = operatorSets[vm.randomUint({min: 0, max: operatorSets.length - 1})];
        return operatorSet;
    }

    function _randomDistributionType(bool allowDisabled) internal returns (DistributionType distributionType) {
        return DistributionType(uint8(vm.randomUint({min: allowDisabled ? 0 : 1, max: 5})));
    }

    function addDistributions(
        OperatorSet[] memory operatorSets,
        IStrategy[] memory strategies,
        uint numDistributions,
        uint64 startEpoch,
        uint16 totalWeight
    ) public createSnapshot returns (uint[] memory distributionIds, Distribution[] memory distributions) {
        print.method("addDistributions");

        // Generate random distribution types
        DistributionType[] memory distributionTypes = new DistributionType[](numDistributions);
        for (uint i = 0; i < numDistributions; ++i) {
            distributionTypes[i] = _randomDistributionType({allowDisabled: false});
        }

        (distributionIds, distributions) = _addDistributionsInternal(operatorSets, strategies, distributionTypes, startEpoch, totalWeight);

        print.gasUsed();
        return (distributionIds, distributions);
    }

    function addDistributionsWithTypes(
        OperatorSet[] memory operatorSets,
        IStrategy[] memory strategies,
        DistributionType[] memory distributionTypes,
        uint64 startEpoch,
        uint16 totalWeight
    ) public createSnapshot returns (uint[] memory distributionIds, Distribution[] memory distributions) {
        print.method("addDistributionsWithTypes");

        (distributionIds, distributions) = _addDistributionsInternal(operatorSets, strategies, distributionTypes, startEpoch, totalWeight);

        print.gasUsed();
        return (distributionIds, distributions);
    }

    function _addDistributionsInternal(
        OperatorSet[] memory operatorSets,
        IStrategy[] memory strategies,
        DistributionType[] memory distributionTypes,
        uint64 startEpoch,
        uint16 totalWeight
    ) internal returns (uint[] memory distributionIds, Distribution[] memory distributions) {
        uint numDistributions = distributionTypes.length;
        distributionIds = new uint[](numDistributions);
        distributions = new Distribution[](numDistributions);

        uint16 weightLeft = totalWeight;

        for (uint i = 0; i < numDistributions; ++i) {
            uint16 weight;

            // For all distributions except the last, randomly assign weight
            // ensuring we leave enough for remaining distributions (at least 1 per distribution)
            if (i < numDistributions - 1) {
                uint remainingDistributions = numDistributions - i - 1;
                uint minWeightNeeded = remainingDistributions; // Reserve at least 1 for each remaining dist
                uint maxAvailable = weightLeft > minWeightNeeded ? weightLeft - minWeightNeeded : 1;

                // Ensure max is at least min (1) to avoid underflow
                if (maxAvailable < 1) maxAvailable = 1;

                weight = uint16(vm.randomUint({min: 1, max: maxAvailable}));
                weightLeft -= weight;
            } else {
                // Last distribution gets all remaining weight
                weight = weightLeft;
                weightLeft = 0;
            }

            DistributionType distributionType = distributionTypes[i];

            distributions[i] = Distribution({
                weight: weight,
                startEpoch: startEpoch,
                totalEpochs: 1,
                distributionType: distributionType,
                operatorSet: (distributionType == DistributionType.OperatorSetTotalStake
                        || distributionType == DistributionType.OperatorSetUniqueStake || distributionType == DistributionType.EigenDA)
                    ? _randomOperatorSet(operatorSets)
                    : OperatorSet({avs: address(0), id: 0}),
                strategiesAndMultipliers: _randomStrategiesAndMultipliers(strategies)
            });

            vm.prank(incentiveCouncil);
            distributionIds[i] = emissionsController.addDistribution(distributions[i]);
        }

        return (distributionIds, distributions);
    }

    function updateDistributions(uint[] memory distributionIds, Distribution[] memory distributions, uint16[] memory weights)
        public
        createSnapshot
        returns (Distribution[] memory updated)
    {
        print.method("updateDistributions");
        // To avoid TotalWeightExceedsMax errors during sequential updates,
        // we need to update in an order that doesn't cause intermediate totals to exceed MAX_TOTAL_WEIGHT.
        // Strategy: First decrease weights, then increase/keep same weights.

        // Store original weights for comparison
        uint16[] memory originalWeights = new uint16[](distributions.length);
        for (uint i = 0; i < distributions.length; ++i) {
            originalWeights[i] = uint16(distributions[i].weight);
        }

        // First pass: decrease weights
        for (uint i = 0; i < distributions.length; ++i) {
            if (weights[i] < originalWeights[i]) {
                distributions[i].weight = weights[i];
                vm.prank(incentiveCouncil);
                emissionsController.updateDistribution(distributionIds[i], distributions[i]);
            }
        }

        // Second pass: increase or keep same weights
        for (uint i = 0; i < distributions.length; ++i) {
            if (weights[i] >= originalWeights[i]) {
                distributions[i].weight = weights[i];
                vm.prank(incentiveCouncil);
                emissionsController.updateDistribution(distributionIds[i], distributions[i]);
            }
        }

        print.gasUsed();
        return distributions;
    }

    /// @dev Calls emissionsController.pressButton with a specified length.
    /// @param length The number of distributions to process.
    /// @return processed The number of distributions processed.
    function pressButton(uint length) public createSnapshot returns (uint processed) {
        print.method("pressButton");
        vm.recordLogs();
        emissionsController.pressButton(length);
        Vm.Log[] memory logs = cheats.getRecordedLogs();
        for (uint i = 0; i < logs.length; ++i) {
            if (logs[i].topics.length > 0 && logs[i].topics[0] == IEmissionsControllerEvents.DistributionProcessed.selector) {
                ++processed;
            }
        }
        print.gasUsed();
        return processed;
    }

    /// @dev Calls emissionsController.sweep.
    function sweep() public createSnapshot returns (bool swept) {
        print.method("sweep");
        vm.recordLogs();
        emissionsController.sweep();
        Vm.Log[] memory logs = cheats.getRecordedLogs();
        for (uint i = 0; i < logs.length; ++i) {
            if (logs[i].topics.length > 0 && logs[i].topics[0] == IEmissionsControllerEvents.Swept.selector) return true;
        }
        print.gasUsed();
        return false;
    }

    /// @dev Warps the chain to a random timestamp within an epoch.
    /// @param epoch The epoch to warp to.
    function warpToEpoch(uint epoch) public createSnapshot {
        print.method("warpToEpoch");
        uint epochLength = emissionsController.EMISSIONS_EPOCH_LENGTH();
        uint randomExcess = vm.randomUint({min: 0, max: epochLength - 1}); // random time wtihin the epoch
        cheats.warp(emissionsController.EMISSIONS_START_TIME() + (epoch * epochLength) + randomExcess);
    }
}
