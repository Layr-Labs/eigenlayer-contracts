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
                            EMISSIONS_INFLATION_RATE,
                            EMISSIONS_START_TIME,
                            EMISSIONS_EPOCH_LENGTH
                        )
                    ),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(EmissionsController.initialize.selector, owner, incentiveCouncil)
                )
            )
        );
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
        emissionsController.initialize(owner, incentiveCouncil);
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
        // Encode an empty array of RewardsSubmissions
        IRewardsCoordinatorTypes.RewardsSubmission[] memory emptySubmissions = new IRewardsCoordinatorTypes.RewardsSubmission[](0);
        bytes memory encodedSubmission = abi.encode(emptySubmissions);

        // Add a distribution first (before emissions start, so startEpoch 0 is in the future)
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: encodedSubmission
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );
    }

    function testFuzz_revert_addDistribution_TotalWeightExceedsMax(uint weight) public {
        weight = bound(weight, 10_001, type(uint).max);
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.TotalWeightExceedsMax.selector);
        emissionsController.addDistribution(
            Distribution({
                weight: weight,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );
    }

    function testFuzz_addDistribution_Correctness(uint weight, uint8 distributionTypeUint8) public {
        weight = bound(weight, 0, 10_000);
        DistributionType distributionType = DistributionType(
            bound(uint8(distributionTypeUint8), uint8(DistributionType.RewardsForAllEarners), uint8(type(DistributionType).max))
        );

        uint nextDistributionId = emissionsController.getTotalDistributions();
        Distribution memory addedDistribution =
            Distribution({weight: weight, startEpoch: 0, stopEpoch: 0, distributionType: distributionType, rewardsCoordinatorCalldata: ""});

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
                rewardsCoordinatorCalldata: ""
            })
        );
    }

    function test_revert_updateDistribution_CannotDisableDistributionViaUpdate() public {
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.CannotDisableDistributionViaUpdate.selector);
        emissionsController.updateDistribution(
            0,
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.Disabled,
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );
    }

    function test_revert_updateDistribution_DistributionIsDisabled() public {
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.disableDistribution(distributionId);

        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.DistributionIsDisabled.selector);
        emissionsController.updateDistribution(
            distributionId,
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );
    }
}

contract EmissionsControllerUnitTests_disableDistribution is EmissionsControllerUnitTests {
    function test_revert_disableDistribution_OnlyIncentiveCouncil() public {
        address notIncentiveCouncil = address(0x3);
        cheats.assume(notIncentiveCouncil != incentiveCouncil);
        cheats.prank(notIncentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.CallerIsNotIncentiveCouncil.selector);
        emissionsController.disableDistribution(0);
    }

    function test_revert_disableDistribution_NonExistentDistribution() public {
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(stdError.indexOOBError); // team may want an explicit check for this
        emissionsController.disableDistribution(0);
    }

    function test_revert_disableDistribution_DistributionIsDisabled() public {
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
            })
        );
        cheats.prank(incentiveCouncil);
        emissionsController.disableDistribution(0);

        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.DistributionIsDisabled.selector);
        emissionsController.disableDistribution(0);
    }

    function test_disableDistribution_Correctness() public {
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
            })
        );

        cheats.prank(incentiveCouncil);
        cheats.expectEmit(true, true, true, true);
        emit DistributionRemoved(distributionId, type(uint).max);
        emissionsController.disableDistribution(distributionId);

        Distribution memory distribution = emissionsController.getDistribution(distributionId);
        assertEq(uint8(distribution.distributionType), uint8(DistributionType.Disabled));
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
                rewardsCoordinatorCalldata: ""
            })
        );

        // Should be pressable after emissions start
        cheats.warp(EMISSIONS_START_TIME);
        assertTrue(emissionsController.isButtonPressable());
    }

    function test_isButtonPressable_AfterProcessing() public {
        // Encode an empty array of RewardsSubmissions
        IRewardsCoordinatorTypes.RewardsSubmission[] memory emptySubmissions = new IRewardsCoordinatorTypes.RewardsSubmission[](0);
        bytes memory encodedSubmission = abi.encode(emptySubmissions);

        // Add a distribution
        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: encodedSubmission
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

contract EmissionsControllerUnitTests_nextButtonPressTime is EmissionsControllerUnitTests {
    function test_revert_nextButtonPressTime_BeforeStart() public {
        // Before emissions start, getCurrentEpoch() returns type(uint).max
        // The calculation (type(uint).max + 1) will overflow and revert in Solidity 0.8.x
        cheats.expectRevert(stdError.arithmeticError);
        emissionsController.nextButtonPressTime();
    }

    function test_nextButtonPressTime_AtStart() public {
        // At emissions start (epoch 0)
        cheats.warp(EMISSIONS_START_TIME);
        assertEq(emissionsController.nextButtonPressTime(), EMISSIONS_START_TIME + EMISSIONS_EPOCH_LENGTH);
    }

    function test_nextButtonPressTime_AfterMultipleEpochs() public {
        // Warp to epoch 5
        cheats.warp(EMISSIONS_START_TIME + 5 * EMISSIONS_EPOCH_LENGTH);
        assertEq(emissionsController.getCurrentEpoch(), 5);
        assertEq(emissionsController.nextButtonPressTime(), EMISSIONS_START_TIME + 6 * EMISSIONS_EPOCH_LENGTH);
    }

    function testFuzz_nextButtonPressTime_Correctness(uint numEpochs) public {
        numEpochs = bound(numEpochs, 0, 1000);

        // Warp to arbitrary epoch
        cheats.warp(EMISSIONS_START_TIME + numEpochs * EMISSIONS_EPOCH_LENGTH);
        uint currentEpoch = emissionsController.getCurrentEpoch();
        assertEq(currentEpoch, numEpochs);

        // Next button press time should be start of next epoch
        assertEq(emissionsController.nextButtonPressTime(), EMISSIONS_START_TIME + (currentEpoch + 1) * EMISSIONS_EPOCH_LENGTH);
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );
        assertEq(emissionsController.getTotalDistributions(), 2);
    }

    function test_getTotalDistributions_AfterDisabling() public {
        // Add distributions
        cheats.prank(incentiveCouncil);
        uint distributionId1 = emissionsController.addDistribution(
            Distribution({
                weight: 3000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 2000,
                startEpoch: 0,
                stopEpoch: 1,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
            })
        );

        assertEq(emissionsController.getTotalDistributions(), 2);

        // Disable one distribution - total count should not change
        cheats.prank(incentiveCouncil);
        emissionsController.disableDistribution(distributionId1);
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
                    rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );

        cheats.prank(incentiveCouncil);
        uint distributionId1 = emissionsController.addDistribution(
            Distribution({
                weight: 4000,
                startEpoch: 1,
                stopEpoch: 8,
                distributionType: DistributionType.OperatorSetTotalStake,
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );

        // Verify updated values
        Distribution memory retrieved = emissionsController.getDistribution(distributionId);
        assertEq(retrieved.weight, 4000);
        assertEq(retrieved.startEpoch, 2);
        assertEq(retrieved.stopEpoch, 15);
        assertEq(uint8(retrieved.distributionType), uint8(DistributionType.OperatorSetUniqueStake));
    }

    function test_getDistribution_AfterDisable() public {
        // Add a distribution
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 5000,
                startEpoch: 0,
                stopEpoch: 10,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
            })
        );

        // Disable the distribution
        cheats.prank(incentiveCouncil);
        emissionsController.disableDistribution(distributionId);

        // Verify it's disabled
        Distribution memory retrieved = emissionsController.getDistribution(distributionId);
        assertEq(uint8(retrieved.distributionType), uint8(DistributionType.Disabled));
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
                rewardsCoordinatorCalldata: ""
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 2000,
                startEpoch: 0,
                stopEpoch: 5,
                distributionType: DistributionType.RewardsForAllEarners,
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 4000,
                startEpoch: 1,
                stopEpoch: 8,
                distributionType: DistributionType.OperatorSetTotalStake,
                rewardsCoordinatorCalldata: ""
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 2000,
                startEpoch: 2,
                stopEpoch: 10,
                distributionType: DistributionType.OperatorSetUniqueStake,
                rewardsCoordinatorCalldata: ""
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
                    weight: 100 * (i + 1),
                    startEpoch: 0,
                    stopEpoch: 5,
                    distributionType: DistributionType.RewardsForAllEarners,
                    rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
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
                rewardsCoordinatorCalldata: ""
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                weight: 4000,
                startEpoch: 1,
                stopEpoch: 8,
                distributionType: DistributionType.OperatorSetTotalStake,
                rewardsCoordinatorCalldata: ""
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
                    weight: 100 * (i + 1),
                    startEpoch: 0,
                    stopEpoch: 5,
                    distributionType: DistributionType.RewardsForAllEarners,
                    rewardsCoordinatorCalldata: ""
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
