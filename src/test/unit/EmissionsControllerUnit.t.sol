// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/EmissionsController.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract EmissionsControllerUnitTests is EigenLayerUnitTestSetup, IEmissionsControllerErrors, IEmissionsControllerEvents {
    using StdStyle for *;
    using ArrayLib for *;

    uint EMISSIONS_INFLATION_RATE = 50;
    uint EMISSIONS_START_TIME = block.timestamp + 1 weeks;
    uint EMISSIONS_EPOCH_LENGTH = 1 weeks;

    address owner = address(0x1);
    address incentiveCouncil = address(0x2);
    EmissionsController emissionsController;

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();
        emissionsController = EmissionsController(
            address(
                new TransparentUpgradeableProxy(
                    address(
                        new EmissionsController(
                            IEigen(address(eigenMock)),
                            IBackingEigen(address(backingEigenMock)),
                            IRewardsCoordinator(address(rewardsCoordinatorMock)),
                            IPauserRegistry(address(pauserRegistry)),
                            EMISSIONS_INFLATION_RATE,
                            EMISSIONS_START_TIME,
                            EMISSIONS_EPOCH_LENGTH
                        )
                    ),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(EmissionsController.initialize.selector, owner, incentiveCouncil, 0)
                )
            )
        );
    }

    function emptyOperatorSet() public pure returns (OperatorSet memory) {
        return OperatorSet({avs: address(0), id: 0});
    }

    function emptyStrategiesAndMultipliers() public pure returns (IRewardsCoordinatorTypes.StrategyAndMultiplier[][] memory) {
        return new IRewardsCoordinatorTypes.StrategyAndMultiplier[][](0);
    }

    function defaultStrategiesAndMultipliers() public pure returns (IRewardsCoordinatorTypes.StrategyAndMultiplier[][] memory) {
        IRewardsCoordinatorTypes.StrategyAndMultiplier[][] memory submissions = new IRewardsCoordinatorTypes.StrategyAndMultiplier[][](1);
        submissions[0] = new IRewardsCoordinatorTypes.StrategyAndMultiplier[](0);
        return submissions;
    }
}

/// -----------------------------------------------------------------------
/// Initialization
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_Initialization_Setters is EmissionsControllerUnitTests {
    function test_constructor_setters() public {
        assertEq(emissionsController.EMISSIONS_INFLATION_RATE(), EMISSIONS_INFLATION_RATE);
        assertEq(emissionsController.EMISSIONS_START_TIME(), EMISSIONS_START_TIME);
        assertEq(emissionsController.EMISSIONS_EPOCH_LENGTH(), EMISSIONS_EPOCH_LENGTH);
    }

    function test_initialize_setters() public {
        assertEq(emissionsController.EMISSIONS_INFLATION_RATE(), EMISSIONS_INFLATION_RATE);
        assertEq(emissionsController.EMISSIONS_START_TIME(), EMISSIONS_START_TIME);
        assertEq(emissionsController.EMISSIONS_EPOCH_LENGTH(), EMISSIONS_EPOCH_LENGTH);
    }

    function test_revert_initialize_AlreadyInitialized() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        emissionsController.initialize(owner, incentiveCouncil, 0);
    }
}

/// -----------------------------------------------------------------------
/// Permissionless Trigger
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_pressButton is EmissionsControllerUnitTests {
    function test_revert_pressButton_AllDistributionsProcessed_NoDistributions() public {
        // Warp to after emissions start
        cheats.warp(EMISSIONS_START_TIME);

        // Attempt to press button with no distributions should revert
        cheats.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        emissionsController.pressButton(0);
    }

    function test_revert_pressButton_AllDistributionsProcessed() public {
        // Add a distribution first (before emissions start, so startEpoch 0 is in the future)
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Warp to after emissions start (now at epoch 0)
        cheats.warp(EMISSIONS_START_TIME);

        // Process all distributions
        emissionsController.pressButton(1);

        // Attempt to press button again should revert
        cheats.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        emissionsController.pressButton(1);
    }
}

/// -----------------------------------------------------------------------
/// Owner Functions
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_setIncentiveCouncil is EmissionsControllerUnitTests {
    function testFuzz_setIncentiveCouncil_OnlyOwner(address notOwner) public {
        cheats.assume(notOwner != owner);
        cheats.assume(notOwner != address(eigenLayerProxyAdmin));
        cheats.prank(notOwner);
        cheats.expectRevert("Ownable: caller is not the owner");
        emissionsController.setIncentiveCouncil(incentiveCouncil);
    }

    function testFuzz_setIncentiveCouncil_Correctness(address newIncentiveCouncil) public {
        cheats.expectEmit(true, true, true, true);
        emit IncentiveCouncilUpdated(newIncentiveCouncil);
        cheats.prank(owner);
        emissionsController.setIncentiveCouncil(newIncentiveCouncil);
        assertEq(emissionsController.incentiveCouncil(), newIncentiveCouncil);
    }
}

/// -----------------------------------------------------------------------
/// Incentive Council Functions
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_addDistribution is EmissionsControllerUnitTests {
    function test_revert_addDistribution_OnlyIncentiveCouncil() public {
        address notIncentiveCouncil = address(0x3);
        cheats.expectRevert(IEmissionsControllerErrors.CallerIsNotIncentiveCouncil.selector);
        emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function test_revert_addDistribution_DisabledDistribution() public {
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.CannotAddDisabledDistribution.selector);
        emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.Disabled,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function test_revert_addDistribution_RewardsSubmissionsCannotBeEmpty() public {
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.RewardsSubmissionsCannotBeEmpty.selector);
        emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: emptyStrategiesAndMultipliers()
            })
        );
    }

    function test_revert_addDistribution_StartEpochMustBeInTheFuture() public {
        cheats.warp(EMISSIONS_START_TIME);
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.StartEpochMustBeInTheFuture.selector);
        emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function testFuzz_revert_addDistribution_TotalWeightExceedsMax(uint weight) public {
        weight = bound(weight, 10_001, type(uint64).max);
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.TotalWeightExceedsMax.selector);
        emissionsController.addDistribution(
            Distribution({
                weight: uint64(weight),
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function test_revert_addDistribution_TotalWeightExceedsMax_MultipleDistributions() public {
        // Add first distribution with weight 6000
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 6000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Attempt to add second distribution with weight 5000, total would be 11000 > 10000
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.TotalWeightExceedsMax.selector);
        emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function testFuzz_addDistribution_Correctness(uint weight, uint8 distributionTypeUint8) public {
        weight = bound(weight, 0, 10_000);
        DistributionType distributionType = DistributionType(
            bound(uint8(distributionTypeUint8), uint8(DistributionType.RewardsForAllEarners), uint8(type(DistributionType).max))
        );

        uint nextDistributionId = emissionsController.getTotalDistributions();

        // Use defaultStrategiesAndMultipliers for non-Manual types, empty for Manual
        IRewardsCoordinatorTypes.StrategyAndMultiplier[][] memory strategiesAndMultipliers =
            distributionType == DistributionType.Manual ? emptyStrategiesAndMultipliers() : defaultStrategiesAndMultipliers();

        Distribution memory addedDistribution = Distribution({
            weight: uint64(weight),
            startEpoch: 0,
            stopEpoch: 0,
            distributionType: distributionType,
            operatorSet: emptyOperatorSet(),
            strategiesAndMultipliers: strategiesAndMultipliers
        });

        cheats.expectEmit(true, true, true, true);
        emit DistributionAdded(nextDistributionId, type(uint).max, addedDistribution);
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(addedDistribution);

        Distribution memory distribution = emissionsController.getDistribution(distributionId);
        assertEq(distributionId, nextDistributionId);
        assertEq(emissionsController.getTotalDistributions(), 1);
        assertEq(distribution.weight, weight);
        assertEq(uint8(distribution.distributionType), uint8(distributionType));
    }
}

contract EmissionsControllerUnitTests_updateDistribution is EmissionsControllerUnitTests {
    function test_revert_updateDistribution_OnlyIncentiveCouncil() public {
        address notIncentiveCouncil = address(0x3);
        cheats.assume(notIncentiveCouncil != incentiveCouncil);
        cheats.prank(notIncentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.CallerIsNotIncentiveCouncil.selector);
        emissionsController.updateDistribution(
            0,
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function test_revert_updateDistribution_NonExistentDistribution() public {
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(stdError.indexOOBError); // team may want an explicit check for this
        emissionsController.updateDistribution(
            0,
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function test_revert_updateDistribution_StartEpochMustBeInTheFuture() public {
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        cheats.warp(EMISSIONS_START_TIME);
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.StartEpochMustBeInTheFuture.selector);
        emissionsController.updateDistribution(
            distributionId,
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function test_revert_updateDistribution_RewardsSubmissionsCannotBeEmpty() public {
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.RewardsSubmissionsCannotBeEmpty.selector);
        emissionsController.updateDistribution(
            distributionId,
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: emptyStrategiesAndMultipliers()
            })
        );
    }

    // NOTE: Fuzz test removed - covered by test_revert_updateDistribution_TotalWeightExceedsMax_MultipleDistributions

    function test_revert_updateDistribution_TotalWeightExceedsMax_MultipleDistributions() public {
        // Add first distribution with weight 6000
        cheats.prank(incentiveCouncil);
        uint distributionId1 = emissionsController.addDistribution(
            Distribution({
                weight: 6000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Add second distribution with weight 3000
        cheats.prank(incentiveCouncil);
        uint distributionId2 = emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Attempt to update second distribution to weight 5000, total would be 11000 > 10000
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.TotalWeightExceedsMax.selector);
        emissionsController.updateDistribution(
            distributionId2,
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    // Regression test for bug where _checkDistribution was called after updating totalWeight,
    // causing valid weight updates to incorrectly revert due to double-counting the new weight
    function test_updateDistribution_DoesNotDoubleCountNewWeight() public {
        // Add first distribution with weight 6000
        cheats.prank(incentiveCouncil);
        uint distributionId1 = emissionsController.addDistribution(
            Distribution({
                weight: 6000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Add second distribution with weight 1000 (total = 7000)
        cheats.prank(incentiveCouncil);
        uint distributionId2 = emissionsController.addDistribution(
            Distribution({
                weight: 1000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Update second distribution to weight 4000 (total should be 10000, which is valid)
        // Without the fix, this would incorrectly check: 4000 + (7000 - 1000 + 4000) = 14000 > 10000
        // With the fix, this correctly checks: 4000 + (7000 - 1000) = 10000 <= 10000
        cheats.prank(incentiveCouncil);
        emissionsController.updateDistribution(
            distributionId2,
            Distribution({
                weight: 4000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Verify the update succeeded
        Distribution memory updated = emissionsController.getDistribution(distributionId2);
        assertEq(updated.weight, 4000);
        assertEq(emissionsController.totalWeight(), 10_000);
    }
}

/// -----------------------------------------------------------------------
/// View Functions
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_getCurrentEpoch is EmissionsControllerUnitTests {
    function test_getCurrentEpoch_MaxBeforeStart() public {
        assertEq(emissionsController.getCurrentEpoch(), type(uint).max);
    }

    function test_getCurrentEpoch_ZeroAtStart() public {
        vm.warp(EMISSIONS_START_TIME);
        assertEq(emissionsController.getCurrentEpoch(), 0);
    }

    function test_getCurrentEpoch_MonotonicallyIncreasingFromZero() public {
        vm.warp(EMISSIONS_START_TIME);
        uint n = 10;
        for (uint i = 1; i < n; i++) {
            assertEq(emissionsController.getCurrentEpoch(), i - 1);
            cheats.warp(block.timestamp + EMISSIONS_EPOCH_LENGTH);
            assertEq(emissionsController.getCurrentEpoch(), i);
        }
    }
}

contract EmissionsControllerUnitTests_isButtonPressable is EmissionsControllerUnitTests {
    function test_isButtonPressable_NoDistributions() public {
        // Before emissions start, no distributions
        assertFalse(emissionsController.isButtonPressable());

        // After emissions start, still no distributions
        cheats.warp(EMISSIONS_START_TIME);
        assertFalse(emissionsController.isButtonPressable());
    }

    function test_isButtonPressable_WithDistributions() public {
        // Add a distribution
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Should be pressable after emissions start
        cheats.warp(EMISSIONS_START_TIME);
        assertTrue(emissionsController.isButtonPressable());
    }

    function test_isButtonPressable_AfterProcessing() public {
        // Add a distribution
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Warp to emissions start
        cheats.warp(EMISSIONS_START_TIME);
        assertTrue(emissionsController.isButtonPressable());

        // Process all distributions
        emissionsController.pressButton(1);

        // Should not be pressable after processing all
        assertFalse(emissionsController.isButtonPressable());
    }
}

contract EmissionsControllerUnitTests_nextTimeButtonPressable is EmissionsControllerUnitTests {
    function test_revert_nextTimeButtonPressable_BeforeStart() public {
        // Before emissions start, getCurrentEpoch() returns type(uint).max
        // The calculation (type(uint).max + 1) will overflow and revert in Solidity 0.8.x
        cheats.expectRevert(stdError.arithmeticError);
        emissionsController.nextTimeButtonPressable();
    }

    function test_nextTimeButtonPressable_AtStart() public {
        // At emissions start (epoch 0)
        cheats.warp(EMISSIONS_START_TIME);
        assertEq(emissionsController.nextTimeButtonPressable(), EMISSIONS_START_TIME + EMISSIONS_EPOCH_LENGTH);
    }

    function test_nextTimeButtonPressable_AfterMultipleEpochs() public {
        // Warp to epoch 5
        cheats.warp(EMISSIONS_START_TIME + 5 * EMISSIONS_EPOCH_LENGTH);
        assertEq(emissionsController.getCurrentEpoch(), 5);
        assertEq(emissionsController.nextTimeButtonPressable(), EMISSIONS_START_TIME + 6 * EMISSIONS_EPOCH_LENGTH);
    }

    function testFuzz_nextTimeButtonPressable_Correctness(uint numEpochs) public {
        numEpochs = bound(numEpochs, 0, 1000);

        // Warp to arbitrary epoch
        cheats.warp(EMISSIONS_START_TIME + numEpochs * EMISSIONS_EPOCH_LENGTH);
        uint currentEpoch = emissionsController.getCurrentEpoch();
        assertEq(currentEpoch, numEpochs);

        // Next button press time should be start of next epoch
        assertEq(emissionsController.nextTimeButtonPressable(), EMISSIONS_START_TIME + (currentEpoch + 1) * EMISSIONS_EPOCH_LENGTH);
    }
}

contract EmissionsControllerUnitTests_getTotalDistributions is EmissionsControllerUnitTests {
    function test_getTotalDistributions_InitiallyZero() public {
        assertEq(emissionsController.getTotalDistributions(), 0);
    }

    function test_getTotalDistributions_AfterAdding() public {
        // Add first distribution
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
        assertEq(emissionsController.getTotalDistributions(), 1);

        // Add second distribution
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 2000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
        assertEq(emissionsController.getTotalDistributions(), 2);
    }

    function testFuzz_getTotalDistributions_Correctness(uint8 count) public {
        count = uint8(bound(count, 0, 50)); // Reasonable upper bound for gas

        for (uint i = 0; i < count; i++) {
            cheats.prank(incentiveCouncil);
            emissionsController.addDistribution(
                Distribution({
                    weight: 100,
                    startEpoch: 0,
                    stopEpoch: 1,
                    distributionType: DistributionType.RewardsForAllEarners,
                    operatorSet: emptyOperatorSet(),
                    strategiesAndMultipliers: defaultStrategiesAndMultipliers()
                })
            );
        }

        assertEq(emissionsController.getTotalDistributions(), count);
    }
}

contract EmissionsControllerUnitTests_getDistribution is EmissionsControllerUnitTests {
    function test_revert_getDistribution_NonExistent() public {
        cheats.expectRevert(stdError.indexOOBError);
        emissionsController.getDistribution(0);
    }

    function test_getDistribution_SingleDistribution() public {
        // Add a distribution
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 10,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Retrieve and verify
        Distribution memory retrieved = emissionsController.getDistribution(distributionId);
        assertEq(retrieved.weight, 5000);
        assertEq(retrieved.startEpoch, 0);
        assertEq(retrieved.stopEpoch, 10);
        assertEq(uint8(retrieved.distributionType), uint8(DistributionType.RewardsForAllEarners));
    }

    function test_getDistribution_MultipleDistributions() public {
        // Add multiple distributions with different parameters
        cheats.prank(incentiveCouncil);
        uint distributionId0 = emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        cheats.prank(incentiveCouncil);
        uint distributionId1 = emissionsController.addDistribution(
            Distribution({
                weight: 4000,
                startEpoch: 1,
                stopEpoch: 8,
                distributionType: DistributionType.OperatorSetTotalStake,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Verify first distribution
        Distribution memory retrieved0 = emissionsController.getDistribution(distributionId0);
        assertEq(retrieved0.weight, 3000);
        assertEq(retrieved0.startEpoch, 0);
        assertEq(retrieved0.stopEpoch, 5);
        assertEq(uint8(retrieved0.distributionType), uint8(DistributionType.RewardsForAllEarners));

        // Verify second distribution
        Distribution memory retrieved1 = emissionsController.getDistribution(distributionId1);
        assertEq(retrieved1.weight, 4000);
        assertEq(retrieved1.startEpoch, 1);
        assertEq(retrieved1.stopEpoch, 8);
        assertEq(uint8(retrieved1.distributionType), uint8(DistributionType.OperatorSetTotalStake));
    }

    function test_getDistribution_AfterUpdate() public {
        // Add a distribution
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Update the distribution (new weight 4000 is still within limits)
        cheats.prank(incentiveCouncil);
        emissionsController.updateDistribution(
            distributionId,
            Distribution({
                weight: 4000,
                startEpoch: 2,
                stopEpoch: 15,
                distributionType: DistributionType.OperatorSetUniqueStake,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Verify updated values
        Distribution memory retrieved = emissionsController.getDistribution(distributionId);
        assertEq(retrieved.weight, 4000);
        assertEq(retrieved.startEpoch, 2);
        assertEq(retrieved.stopEpoch, 15);
        assertEq(uint8(retrieved.distributionType), uint8(DistributionType.OperatorSetUniqueStake));
    }
}

contract EmissionsControllerUnitTests_getDistributions is EmissionsControllerUnitTests {
    function test_revert_getDistributions_OutOfBounds() public {
        // Add 2 distributions
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 2000,
                startEpoch: 0,
                stopEpoch: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Try to get distributions beyond available length
        cheats.expectRevert(stdError.indexOOBError);
        emissionsController.getDistributions(0, 3);

        cheats.expectRevert(stdError.indexOOBError);
        emissionsController.getDistributions(1, 2);
    }

    function test_getDistributions_All() public {
        // Add multiple distributions
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 4000,
                startEpoch: 1,
                stopEpoch: 8,
                distributionType: DistributionType.OperatorSetTotalStake,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 2000,
                startEpoch: 2,
                stopEpoch: 10,
                distributionType: DistributionType.OperatorSetUniqueStake,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Get all distributions
        Distribution[] memory distributions = emissionsController.getDistributions(0, 3);

        assertEq(distributions.length, 3);
        assertEq(distributions[0].weight, 3000);
        assertEq(distributions[1].weight, 4000);
        assertEq(distributions[2].weight, 2000);
    }

    function test_getDistributions_Subset() public {
        // Add multiple distributions (total weight: 100 * 5 = 500, well under 10000 limit)
        for (uint i = 0; i < 5; i++) {
            cheats.prank(incentiveCouncil);
            emissionsController.addDistribution(
                Distribution({
                    weight: uint64(100 * (i + 1)),
                    startEpoch: 0,
                    stopEpoch: 5,
                    distributionType: DistributionType.RewardsForAllEarners,
                    operatorSet: emptyOperatorSet(),
                    strategiesAndMultipliers: defaultStrategiesAndMultipliers()
                })
            );
        }

        // Get subset starting at index 1, length 3
        Distribution[] memory distributions = emissionsController.getDistributions(1, 3);

        assertEq(distributions.length, 3);
        assertEq(distributions[0].weight, 200); // Index 1
        assertEq(distributions[1].weight, 300); // Index 2
        assertEq(distributions[2].weight, 400); // Index 3
    }

    function test_getDistributions_EmptyArray() public {
        // Add some distributions
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Get 0 length
        Distribution[] memory distributions = emissionsController.getDistributions(0, 0);
        assertEq(distributions.length, 0);
    }

    function test_getDistributions_SingleElement() public {
        // Add multiple distributions
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 4000,
                startEpoch: 1,
                stopEpoch: 8,
                distributionType: DistributionType.OperatorSetTotalStake,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Get single element at index 1
        Distribution[] memory distributions = emissionsController.getDistributions(1, 1);

        assertEq(distributions.length, 1);
        assertEq(distributions[0].weight, 4000);
        assertEq(uint8(distributions[0].distributionType), uint8(DistributionType.OperatorSetTotalStake));
    }

    function testFuzz_getDistributions_Correctness(uint8 totalCount, uint8 start, uint8 length) public {
        // Limit totalCount to 10 to avoid exceeding MAX_TOTAL_WEIGHT (10000)
        // Each distribution has weight 100 * (i+1), so max weight per distribution is 100 * 10 = 1000
        // Total max weight = 100 + 200 + ... + 1000 = 5500, well under 10000
        totalCount = uint8(bound(totalCount, 1, 10));
        start = uint8(bound(start, 0, totalCount - 1));
        // Ensure we don't go out of bounds
        uint8 maxLength = totalCount - start;
        length = uint8(bound(length, 0, maxLength));

        // Add distributions
        for (uint i = 0; i < totalCount; i++) {
            cheats.prank(incentiveCouncil);
            emissionsController.addDistribution(
                Distribution({
                    weight: uint64(100 * (i + 1)),
                    startEpoch: 0,
                    stopEpoch: 5,
                    distributionType: DistributionType.RewardsForAllEarners,
                    operatorSet: emptyOperatorSet(),
                    strategiesAndMultipliers: defaultStrategiesAndMultipliers()
                })
            );
        }

        // Get subset
        Distribution[] memory distributions = emissionsController.getDistributions(start, length);

        assertEq(distributions.length, length);

        // Verify each element
        for (uint i = 0; i < length; i++) {
            assertEq(distributions[i].weight, 100 * (start + i + 1));
        }
    }
}

/// -----------------------------------------------------------------------
/// Pausable Token Flows
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_Pausable is EmissionsControllerUnitTests {
    function test_revert_pressButton_WhenPaused() public {
        // Add a distribution before emissions start
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 10,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Warp to after emissions start
        cheats.warp(EMISSIONS_START_TIME);

        // Pause token flows (PAUSED_TOKEN_FLOWS = 0, so bit flag is 2^0 = 1)
        cheats.prank(pauser);
        emissionsController.pause(1);

        // Attempt to press button should revert
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        emissionsController.pressButton(1);
    }

    function test_revert_sweep_WhenPaused() public {
        // Warp to after emissions start
        cheats.warp(EMISSIONS_START_TIME);

        // Pause token flows (PAUSED_TOKEN_FLOWS = 0, so bit flag is 2^0 = 1)
        cheats.prank(pauser);
        emissionsController.pause(1);

        // Attempt to sweep should revert
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        emissionsController.sweep();
    }
}
