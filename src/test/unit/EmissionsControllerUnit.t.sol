// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/EmissionsController.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract EmissionsControllerUnitTests is EigenLayerUnitTestSetup, IEmissionsControllerErrors, IEmissionsControllerEvents {
    using StdStyle for *;
    using ArrayLib for *;

    uint EMISSIONS_INFLATION_RATE = 50;
    // Use a fixed start time that's aligned with CALCULATION_INTERVAL_SECONDS (1 day = 86400)
    // 7 days from an aligned base ensures proper alignment
    uint EMISSIONS_START_TIME = 7 days;
    uint EMISSIONS_EPOCH_LENGTH = 1 weeks;
    uint REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS = 86_400;

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
                            IAllocationManager(address(allocationManagerMock)),
                            IRewardsCoordinator(address(rewardsCoordinatorMock)),
                            IPauserRegistry(address(pauserRegistry)),
                            EMISSIONS_INFLATION_RATE,
                            EMISSIONS_START_TIME,
                            EMISSIONS_EPOCH_LENGTH,
                            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS
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

    function _getTotalDistributionsProcessed() public returns (uint) {
        Vm.Log[] memory logs = cheats.getRecordedLogs();
        uint processed = 0;
        for (uint i = 0; i < logs.length; i++) {
            if (logs[i].topics.length > 0 && logs[i].topics[0] == IEmissionsControllerEvents.DistributionProcessed.selector) {
                ++processed;
            }
        }
        return processed;
    }

    function _pressButton(uint epoch, uint length, uint expectedProcessed, bool expectedPressable) public {
        cheats.warp(EMISSIONS_START_TIME + epoch * EMISSIONS_EPOCH_LENGTH);
        assertEq(emissionsController.getCurrentEpoch(), epoch);
        cheats.recordLogs();
        emissionsController.pressButton(length);
        assertEq(_getTotalDistributionsProcessed(), expectedProcessed, "processed != expectedProcessed");
        assertEq(emissionsController.isButtonPressable(), expectedPressable, "isButtonPressable != expectedPressable");
    }

    function _addDefaultDistribution(uint64 startEpoch, uint64 totalEpochs, uint64 weight) public {
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: weight,
                startEpoch: startEpoch,
                totalEpochs: totalEpochs,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }
}

/// -----------------------------------------------------------------------
/// Initialization
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_Initialization_Setters is EmissionsControllerUnitTests {
    function test_constructor_setters() public {
        assertEq(address(emissionsController.EIGEN()), address(eigenMock));
        assertEq(address(emissionsController.BACKING_EIGEN()), address(backingEigenMock));
        assertEq(address(emissionsController.ALLOCATION_MANAGER()), address(allocationManagerMock));
        assertEq(address(emissionsController.REWARDS_COORDINATOR()), address(rewardsCoordinatorMock));
        assertEq(address(emissionsController.pauserRegistry()), address(pauserRegistry));
        assertEq(emissionsController.EMISSIONS_INFLATION_RATE(), EMISSIONS_INFLATION_RATE);
        assertEq(emissionsController.EMISSIONS_START_TIME(), EMISSIONS_START_TIME);
        assertEq(emissionsController.EMISSIONS_EPOCH_LENGTH(), EMISSIONS_EPOCH_LENGTH);
    }

    function test_initialize_setters() public {
        assertEq(emissionsController.owner(), owner);
        assertEq(emissionsController.incentiveCouncil(), incentiveCouncil);
        assertEq(emissionsController.paused(), 0);
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
    /// -----------------------------------------------------------------------
    /// Revert Tests
    /// -----------------------------------------------------------------------
    /// @notice Assert the function reverts when paused.
    function test_revert_pressButton_WhenPaused() public {
        _addDefaultDistribution({startEpoch: 0, totalEpochs: 1, weight: 10_000});
        cheats.prank(pauser);
        emissionsController.pauseAll();
        cheats.warp(EMISSIONS_START_TIME);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        emissionsController.pressButton(1);
    }

    /// @notice Assert the function reverts when emissions have not started yet.
    function test_revert_pressButton_EmissionsNotStarted() public {
        _addDefaultDistribution({startEpoch: 0, totalEpochs: 1, weight: 10_000});
        cheats.expectRevert(IEmissionsControllerErrors.EmissionsNotStarted.selector);
        emissionsController.pressButton(1);
    }

    /// @notice Assert the function reverts when no distributions are left to be processed (none to start with).
    function test_revert_pressButton_AllDistributionsProcessed_NoDistributions() public {
        cheats.warp(EMISSIONS_START_TIME);
        cheats.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        emissionsController.pressButton(0);
    }

    /// @notice Assert the function reverts when no distributions are left to be processed (some to start with).
    function test_revert_pressButton_AllDistributionsProcessed() public {
        _addDefaultDistribution({startEpoch: 0, totalEpochs: 1, weight: 10_000});
        cheats.warp(EMISSIONS_START_TIME);
        emissionsController.pressButton(1);
        cheats.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        emissionsController.pressButton(1);
    }

    /// -----------------------------------------------------------------------
    /// Timing Fuzz Tests
    /// -----------------------------------------------------------------------

    /// @notice Assert the function processes a single epoch distribution correctly.
    function testFuzz_pressButton_SingleEpochDistribution(uint64 startEpoch) public {
        startEpoch = uint64(bound(startEpoch, 0, type(uint8).max));
        _addDefaultDistribution({startEpoch: startEpoch, totalEpochs: 1, weight: 10_000});
        _pressButton({epoch: startEpoch, length: 1, expectedProcessed: 1, expectedPressable: false});
        _pressButton({epoch: startEpoch + 1, length: 1, expectedProcessed: 0, expectedPressable: false});
    }

    /// @notice Assert the function processes a multiple epoch distribution correctly.
    function testFuzz_pressButton_MultipleEpochDistribution(uint64 startEpoch, uint64 totalEpochs) public {
        startEpoch = uint64(bound(startEpoch, 0, type(uint8).max));
        totalEpochs = uint64(bound(totalEpochs, 2, type(uint8).max));
        _addDefaultDistribution({startEpoch: startEpoch, totalEpochs: totalEpochs, weight: 10_000});
        for (uint i = 0; i < totalEpochs; i++) {
            _pressButton({epoch: startEpoch + i, length: 1, expectedProcessed: 1, expectedPressable: false});
        }
        _pressButton({epoch: startEpoch + totalEpochs, length: 1, expectedProcessed: 0, expectedPressable: false});
    }

    /// @notice Assert the function processes an infinite distribution correctly.
    function testFuzz_pressButton_InfiniteDistribution(uint64 startEpoch) public {
        startEpoch = uint64(bound(startEpoch, 0, type(uint8).max));
        _addDefaultDistribution({startEpoch: startEpoch, totalEpochs: 0, weight: 10_000});
        for (uint64 epoch = 0; epoch < 100; ++epoch) {
            _pressButton({epoch: startEpoch + epoch, length: 1, expectedProcessed: 1, expectedPressable: false});
        }
    }

    /// -----------------------------------------------------------------------
    /// Minting Behavior Tests
    /// -----------------------------------------------------------------------

    /// @notice Assert the function mints only once per epoch.
    function test_pressButton_MintsOnlyOncePerEpoch() public {
        _addDefaultDistribution({startEpoch: 0, totalEpochs: 1, weight: 5000});
        _addDefaultDistribution({startEpoch: 0, totalEpochs: 1, weight: 5000});

        uint balanceBefore = eigenMock.balanceOf(address(emissionsController));
        _pressButton({epoch: 0, length: 1, expectedProcessed: 1, expectedPressable: true});
        uint balanceAfter = eigenMock.balanceOf(address(emissionsController));
        assertEq(balanceAfter, balanceBefore + EMISSIONS_INFLATION_RATE);

        balanceBefore = eigenMock.balanceOf(address(emissionsController));
        _pressButton({epoch: 0, length: 1, expectedProcessed: 1, expectedPressable: false});
        balanceAfter = eigenMock.balanceOf(address(emissionsController));
        assertEq(balanceAfter, balanceBefore);
    }

    /// -----------------------------------------------------------------------
    /// Distribution Skipping Tests
    /// -----------------------------------------------------------------------

    function testFuzz_pressButton_SkipsDisabledDistributions(uint64 startEpoch) public {} // TODO: implement

    /// @notice Assert the function skips distributions that have not started yet.
    function testFuzz_pressButton_SkipsNotYetStartedDistributions(uint64 startEpoch) public {
        startEpoch = uint64(bound(startEpoch, 1, type(uint8).max));
        _addDefaultDistribution({startEpoch: startEpoch, totalEpochs: 1, weight: 10_000});
        for (uint64 epoch = 0; epoch < startEpoch; ++epoch) {
            _pressButton({epoch: epoch, length: 1, expectedProcessed: 0, expectedPressable: false});
        }
        _pressButton({epoch: startEpoch, length: 1, expectedProcessed: 1, expectedPressable: false});
    }

    /// @notice Assert the function skips distributions that have ended.
    function testFuzz_pressButton_SkipsEndedDistributions(uint64 startEpoch, uint64 totalEpochs) public {
        startEpoch = uint64(bound(startEpoch, 1, type(uint8).max));
        _addDefaultDistribution({startEpoch: startEpoch, totalEpochs: 1, weight: 10_000});
        for (uint i = 1; i < 5; ++i) {
            _pressButton({epoch: startEpoch + i, length: 1, expectedProcessed: 0, expectedPressable: false});
        }
    }

    /// @notice Assert the function skips distributions with zero weight.
    function testFuzz_pressButton_SkipsZeroWeightDistribution(uint64 startEpoch) public {
        startEpoch = uint64(bound(startEpoch, 1, type(uint8).max));
        _addDefaultDistribution({startEpoch: startEpoch, totalEpochs: 1, weight: 0});
        _pressButton({epoch: startEpoch, length: 1, expectedProcessed: 0, expectedPressable: false});
    }

    /// -----------------------------------------------------------------------
    /// Length Edge Cases
    /// -----------------------------------------------------------------------

    function test_pressButton_LengthZeroProcessesNone() public {
        _addDefaultDistribution({startEpoch: 0, totalEpochs: 1, weight: 10_000});
        _pressButton({epoch: 0, length: 0, expectedProcessed: 0, expectedPressable: true});
    }
}

contract EmissionsControllerUnitTests_sweep is EmissionsControllerUnitTests {
    /// @notice Assert the function reverts when paused.
    function test_revert_sweep_WhenPaused() public {
        cheats.warp(EMISSIONS_START_TIME);
        cheats.prank(pauser);
        emissionsController.pause(1);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        emissionsController.sweep();
    }

    /// @notice Assert the function transfers tokens to the incentive council.
    function test_sweep_TransfersTokensToIncentiveCouncil() public {
        _addDefaultDistribution({startEpoch: 0, totalEpochs: 1, weight: 10_000});
        cheats.warp(EMISSIONS_START_TIME);
        emissionsController.pressButton(1);

        uint sweepAmount = 100 ether;
        deal(address(eigenMock), address(emissionsController), sweepAmount);
        assertEq(eigenMock.balanceOf(address(emissionsController)), sweepAmount);

        uint councilBalanceBefore = eigenMock.balanceOf(incentiveCouncil);

        cheats.expectEmit(true, true, true, true);
        emit Swept(incentiveCouncil, sweepAmount);
        emissionsController.sweep();

        assertEq(eigenMock.balanceOf(address(emissionsController)), 0);
        assertEq(eigenMock.balanceOf(incentiveCouncil), councilBalanceBefore + sweepAmount);
    }

    function test_sweep_DoesNothingWhenButtonPressable() public {
        // Add distribution
        _addDefaultDistribution({startEpoch: 0, totalEpochs: 1, weight: 10_000});

        // Warp to epoch 0 - button is pressable
        cheats.warp(EMISSIONS_START_TIME);

        // Give some EIGEN tokens directly to the controller
        uint sweepAmount = 100 ether;
        deal(address(eigenMock), address(emissionsController), sweepAmount);

        // Sweep should do nothing when button is pressable
        emissionsController.sweep();

        // Tokens should still be in the controller
        assertEq(eigenMock.balanceOf(address(emissionsController)), sweepAmount);
    }

    function test_sweep_DoesNothingWhenBalanceZeroAfterButtonNotPressable() public {
        // Don't add any distributions - button won't be pressable after start

        // Warp to epoch 0
        cheats.warp(EMISSIONS_START_TIME);

        // Can't press button with no distributions
        assertFalse(emissionsController.isButtonPressable());

        // Balance is 0, sweep should do nothing (no transfer, no event)
        uint councilBalanceBefore = eigenMock.balanceOf(incentiveCouncil);
        assertEq(eigenMock.balanceOf(address(emissionsController)), 0);
        emissionsController.sweep();
        assertEq(eigenMock.balanceOf(address(emissionsController)), 0);
        assertEq(eigenMock.balanceOf(incentiveCouncil), councilBalanceBefore);
    }
}

/// -----------------------------------------------------------------------
/// Owner Functions
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_setIncentiveCouncil is EmissionsControllerUnitTests {
    function test_revert_setIncentiveCouncil_ZeroAddress() public {
        cheats.prank(owner);
        cheats.expectRevert(IPausable.InputAddressZero.selector);
        emissionsController.setIncentiveCouncil(address(0));
    }

    function test_revert_setIncentiveCouncil_OnlyOwner(address notOwner) public {
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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

        uint nextDistributionId = emissionsController.getTotalProcessableDistributions();

        // Use defaultStrategiesAndMultipliers for non-Manual types, empty for Manual
        IRewardsCoordinatorTypes.StrategyAndMultiplier[][] memory strategiesAndMultipliers =
            distributionType == DistributionType.Manual ? emptyStrategiesAndMultipliers() : defaultStrategiesAndMultipliers();

        Distribution memory addedDistribution = Distribution({
            weight: uint64(weight),
            startEpoch: 0,
            totalEpochs: 0,
            distributionType: distributionType,
            operatorSet: emptyOperatorSet(),
            strategiesAndMultipliers: strategiesAndMultipliers
        });

        allocationManagerMock.setIsOperatorSet(addedDistribution.operatorSet, true);

        cheats.expectEmit(true, true, true, true);
        emit DistributionAdded(nextDistributionId, type(uint).max, addedDistribution);
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(addedDistribution);

        Distribution memory distribution = emissionsController.getDistribution(distributionId);
        assertEq(distributionId, nextDistributionId);
        assertEq(emissionsController.getTotalProcessableDistributions(), 1);
        assertEq(distribution.weight, weight);
        assertEq(uint8(distribution.distributionType), uint8(distributionType));
    }

    function test_revert_addDistribution_AllDistributionsMustBeProcessed() public {
        // Add first distribution before emissions start
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                totalEpochs: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Warp to after emissions start (epoch 0)
        cheats.warp(EMISSIONS_START_TIME);

        // Now there's 1 distribution but 0 processed, button is pressable
        assertTrue(emissionsController.isButtonPressable());

        // Attempt to add another distribution before processing the first one should revert
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.AllDistributionsMustBeProcessed.selector);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 1,
                totalEpochs: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
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
                totalEpochs: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
    }

    function test_revert_updateDistribution_AllDistributionsMustBeProcessed() public {
        // Add first distribution before emissions start
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                totalEpochs: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Warp to after emissions start (epoch 0)
        cheats.warp(EMISSIONS_START_TIME);

        // Now there's 1 distribution but 0 processed, button is pressable
        assertTrue(emissionsController.isButtonPressable());

        // Attempt to update the distribution before processing it should revert
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.AllDistributionsMustBeProcessed.selector);
        emissionsController.updateDistribution(
            distributionId,
            Distribution({
                weight: 3000,
                startEpoch: 1,
                totalEpochs: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
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
                totalEpochs: 1,
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
                totalEpochs: 1,
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

    function test_isButtonPressable_BeforeEmissionsStart_WithDistributions() public {
        // Add a distribution before emissions start
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                totalEpochs: 0, // infinite duration
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Before emissions start, getCurrentEpoch() returns type(uint256).max
        // This accesses _epochs[type(uint256).max] which is uninitialized
        // Button should NOT be pressable before emissions start
        assertEq(emissionsController.getCurrentEpoch(), type(uint).max, "Current epoch should be max before start");
        assertEq(emissionsController.getTotalProcessableDistributions(), 1, "Should have 1 distribution");
        assertFalse(emissionsController.isButtonPressable(), "Button should not be pressable before emissions start");
    }
}

contract EmissionsControllerUnitTests_nextTimeButtonPressable is EmissionsControllerUnitTests {
    function test_nextTimeButtonPressable_BeforeStart() public {
        assertEq(emissionsController.nextTimeButtonPressable(), EMISSIONS_START_TIME);
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

contract EmissionsControllerUnitTests_lastTimeButtonPressable is EmissionsControllerUnitTests {
    function test_lastTimeButtonPressable_BeforeStart() public {
        assertEq(emissionsController.lastTimeButtonPressable(), type(uint).max);
    }

    function test_lastTimeButtonPressable_AtStart() public {
        // At emissions start (epoch 0)
        cheats.warp(EMISSIONS_START_TIME);
        assertEq(emissionsController.lastTimeButtonPressable(), EMISSIONS_START_TIME);
    }

    function test_lastTimeButtonPressable_AfterMultipleEpochs() public {
        // Warp to middle of epoch 5
        cheats.warp(EMISSIONS_START_TIME + 5 * EMISSIONS_EPOCH_LENGTH + EMISSIONS_EPOCH_LENGTH / 2);
        assertEq(emissionsController.getCurrentEpoch(), 5);
        // Last time pressable should be start of epoch 5
        assertEq(emissionsController.lastTimeButtonPressable(), EMISSIONS_START_TIME + 5 * EMISSIONS_EPOCH_LENGTH);
    }

    function testFuzz_lastTimeButtonPressable_Correctness(uint numEpochs) public {
        numEpochs = bound(numEpochs, 0, 1000);

        // Warp to arbitrary epoch
        cheats.warp(EMISSIONS_START_TIME + numEpochs * EMISSIONS_EPOCH_LENGTH);
        uint currentEpoch = emissionsController.getCurrentEpoch();
        assertEq(currentEpoch, numEpochs);

        // Last button press time should be start of current epoch
        assertEq(emissionsController.lastTimeButtonPressable(), EMISSIONS_START_TIME + currentEpoch * EMISSIONS_EPOCH_LENGTH);
    }
}

contract EmissionsControllerUnitTests_getTotalProcessableDistributions is EmissionsControllerUnitTests {
    function test_getTotalProcessableDistributions_InitiallyZero() public {
        assertEq(emissionsController.getTotalProcessableDistributions(), 0);
    }

    function test_getTotalProcessableDistributions_AfterAdding() public {
        // Add first distribution
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                totalEpochs: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
        assertEq(emissionsController.getTotalProcessableDistributions(), 1);

        // Add second distribution
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 2000,
                startEpoch: 0,
                totalEpochs: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );
        assertEq(emissionsController.getTotalProcessableDistributions(), 2);
    }

    function testFuzz_getTotalProcessableDistributions_Correctness(uint8 count) public {
        count = uint8(bound(count, 0, 50)); // Reasonable upper bound for gas

        for (uint i = 0; i < count; i++) {
            cheats.prank(incentiveCouncil);
            emissionsController.addDistribution(
                Distribution({
                    weight: 100,
                    startEpoch: 0,
                    totalEpochs: 1,
                    distributionType: DistributionType.RewardsForAllEarners,
                    operatorSet: emptyOperatorSet(),
                    strategiesAndMultipliers: defaultStrategiesAndMultipliers()
                })
            );
        }

        assertEq(emissionsController.getTotalProcessableDistributions(), count);
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
                totalEpochs: 10,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Retrieve and verify
        Distribution memory retrieved = emissionsController.getDistribution(distributionId);
        assertEq(retrieved.weight, 5000);
        assertEq(retrieved.startEpoch, 0);
        assertEq(retrieved.totalEpochs, 10);
        assertEq(uint8(retrieved.distributionType), uint8(DistributionType.RewardsForAllEarners));
    }

    function test_getDistribution_MultipleDistributions() public {
        allocationManagerMock.setIsOperatorSet(emptyOperatorSet(), true);

        // Add multiple distributions with different parameters
        cheats.prank(incentiveCouncil);
        uint distributionId0 = emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                totalEpochs: 5,
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
                totalEpochs: 7,
                distributionType: DistributionType.OperatorSetTotalStake,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Verify first distribution
        Distribution memory retrieved0 = emissionsController.getDistribution(distributionId0);
        assertEq(retrieved0.weight, 3000);
        assertEq(retrieved0.startEpoch, 0);
        assertEq(retrieved0.totalEpochs, 5);
        assertEq(uint8(retrieved0.distributionType), uint8(DistributionType.RewardsForAllEarners));

        // Verify second distribution
        Distribution memory retrieved1 = emissionsController.getDistribution(distributionId1);
        assertEq(retrieved1.weight, 4000);
        assertEq(retrieved1.startEpoch, 1);
        assertEq(retrieved1.totalEpochs, 7);
        assertEq(uint8(retrieved1.distributionType), uint8(DistributionType.OperatorSetTotalStake));
    }

    function test_getDistribution_AfterUpdate() public {
        allocationManagerMock.setIsOperatorSet(emptyOperatorSet(), true);

        // Add a distribution
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                totalEpochs: 5,
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
                totalEpochs: 13,
                distributionType: DistributionType.OperatorSetUniqueStake,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Verify updated values
        Distribution memory retrieved = emissionsController.getDistribution(distributionId);
        assertEq(retrieved.weight, 4000);
        assertEq(retrieved.startEpoch, 2);
        assertEq(retrieved.totalEpochs, 13);
        assertEq(uint8(retrieved.distributionType), uint8(DistributionType.OperatorSetUniqueStake));
    }
}

contract EmissionsControllerUnitTests_getDistributions is EmissionsControllerUnitTests {
    /// @notice Test that getDistributions bounds the length parameter to avoid out-of-bounds errors.
    /// @dev This is a regression test for the length bounding feature added to getDistributions.
    function test_getDistributions_LengthBounding() public {
        // Add 2 distributions
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                totalEpochs: 5,
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
                totalEpochs: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                operatorSet: emptyOperatorSet(),
                strategiesAndMultipliers: defaultStrategiesAndMultipliers()
            })
        );

        // Request length that exceeds available distributions from start=0
        // Should return 2 distributions instead of reverting
        Distribution[] memory distributions = emissionsController.getDistributions(0, 5);
        assertEq(distributions.length, 2, "Returned array should have requested length");

        // Request length that exceeds available distributions from start=1
        // Should return 1 distribution instead of reverting
        distributions = emissionsController.getDistributions(1, 5);
        assertEq(distributions.length, 1, "Returned array should have requested length");
    }

    function test_getDistributions_All() public {
        allocationManagerMock.setIsOperatorSet(emptyOperatorSet(), true);

        // Add multiple distributions
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                totalEpochs: 5,
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
                totalEpochs: 7,
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
                totalEpochs: 8,
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
                    totalEpochs: 5,
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
                totalEpochs: 5,
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
        allocationManagerMock.setIsOperatorSet(emptyOperatorSet(), true);

        // Add multiple distributions
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                totalEpochs: 5,
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
                totalEpochs: 7,
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
                    totalEpochs: 5,
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
