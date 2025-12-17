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
}

/// -----------------------------------------------------------------------
/// Permissionless Trigger
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_pressButton is EmissionsControllerUnitTests {}

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
                encodedRewardsSubmission: ""
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
                encodedRewardsSubmission: ""
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
                encodedRewardsSubmission: ""
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
                encodedRewardsSubmission: ""
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
            Distribution({weight: weight, startEpoch: 0, stopEpoch: 0, distributionType: distributionType, encodedRewardsSubmission: ""});

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
                encodedRewardsSubmission: ""
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
                encodedRewardsSubmission: ""
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
                encodedRewardsSubmission: ""
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
                encodedRewardsSubmission: ""
            })
        );

        cheats.prank(incentiveCouncil);
        emissionsController.removeDistribution(distributionId);

        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.DistributionIsDisabled.selector);
        emissionsController.updateDistribution(
            distributionId,
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                encodedRewardsSubmission: ""
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
                encodedRewardsSubmission: ""
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
                encodedRewardsSubmission: ""
            })
        );
    }

    function testFuzz_updateDistribution_TotalWeightExceedsMax(uint weight) public {
        weight = bound(weight, 10_001, type(uint).max);

        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                encodedRewardsSubmission: ""
            })
        );

        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.TotalWeightExceedsMax.selector);
        emissionsController.updateDistribution(
            0,
            Distribution({
                weight: weight,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                encodedRewardsSubmission: ""
            })
        );
    }
}

// TODO: update distirbution test reverts if disabled via removeDistribution

contract EmissionsControllerUnitTests_removeDistribution is EmissionsControllerUnitTests {
    function test_revert_removeDistribution_OnlyIncentiveCouncil() public {
        address notIncentiveCouncil = address(0x3);
        cheats.assume(notIncentiveCouncil != incentiveCouncil);
        cheats.prank(notIncentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.CallerIsNotIncentiveCouncil.selector);
        emissionsController.removeDistribution(0);
    }

    function test_revert_removeDistribution_NonExistentDistribution() public {
        cheats.prank(incentiveCouncil);
        cheats.expectRevert(stdError.indexOOBError); // team may want an explicit check for this
        emissionsController.removeDistribution(0);
    }

    function test_revert_removeDistribution_DistributionIsDisabled() public {
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                encodedRewardsSubmission: ""
            })
        );
        cheats.prank(incentiveCouncil);
        emissionsController.removeDistribution(0);

        cheats.prank(incentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.DistributionIsDisabled.selector);
        emissionsController.removeDistribution(0);
    }

    function test_removeDistribution_Correctness() public {
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(
            Distribution({
                weight: 10_000,
                startEpoch: 0,
                stopEpoch: 0,
                distributionType: DistributionType.RewardsForAllEarners,
                encodedRewardsSubmission: ""
            })
        );

        cheats.prank(incentiveCouncil);
        cheats.expectEmit(true, true, true, true);
        emit DistributionRemoved(distributionId, type(uint).max);
        emissionsController.removeDistribution(distributionId);

        Distribution memory distribution = emissionsController.getDistribution(distributionId);
        assertEq(distribution.distributionType, DistributionType.Disabled);
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

contract EmissionsControllerUnitTests_isButtonPressable is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_nextButtonPressTime is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_getTotalDistributions is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_getDistribution is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_getDistributions is EmissionsControllerUnitTests {}
