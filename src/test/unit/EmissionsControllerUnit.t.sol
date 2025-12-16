// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/EmissionsController.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract EmissionsControllerUnitTests is EigenLayerUnitTestSetup, IEmissionsControllerErrors, IEmissionsControllerEvents {
    using StdStyle for *;
    using ArrayLib for *;

    uint EMISSIONS_INFLATION_RATE = 50;
    uint EMISSIONS_START_TIME = block.timestamp;
    uint EMISSIONS_EPOCH_LENGTH = 1 weeks;

    address owner = address(0x1);
    address incentiveCouncil = address(0x2);
    EmissionsController emissionsController;

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();
        emissionsController = EmissionsController(
            address(
                new TransparentUpgradeableProxy(
                    address(new EmissionsController(EMISSIONS_INFLATION_RATE, EMISSIONS_START_TIME, EMISSIONS_EPOCH_LENGTH)),
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
    function testFuzz_addDistribution_OnlyIncentiveCouncil(address notIncentiveCouncil, Distribution memory distribution) public {
        cheats.assume(notIncentiveCouncil != incentiveCouncil);
        cheats.prank(notIncentiveCouncil);
        cheats.expectRevert(IEmissionsControllerErrors.CallerIsNotIncentiveCouncil.selector);
        emissionsController.addDistribution(distribution);
    }

    function testFuzz_addDistribution_Correctness(uint weight, DistributionType distributionType) public {
        uint nextDistributionId = emissionsController.getTotalDistributions();
        Distribution memory addedDistribution =
            Distribution({weight: weight, distributionType: distributionType, encodedRewardsSubmission: ""});

        cheats.expectEmit(true, true, true, true);
        emit DistributionAdded(nextDistributionId, 0, addedDistribution);
        cheats.prank(incentiveCouncil);
        uint distributionId = emissionsController.addDistribution(addedDistribution);

        Distribution memory distribution = emissionsController.getDistribution(distributionId);
        assertEq(distributionId, nextDistributionId);
        assertEq(emissionsController.getTotalDistributions(), 1);
        assertEq(distribution.weight, weight);
        assertEq(uint8(distribution.distributionType), uint8(distributionType));
    }
}

contract EmissionsControllerUnitTests_updateDistribution is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_removeDistribution is EmissionsControllerUnitTests {}

/// -----------------------------------------------------------------------
/// View Functions
/// -----------------------------------------------------------------------

contract EmissionsControllerUnitTests_getCurrentEpoch is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_isButtonPressable is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_nextButtonPressTime is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_getTotalDistributions is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_getDistribution is EmissionsControllerUnitTests {}

contract EmissionsControllerUnitTests_getDistributions is EmissionsControllerUnitTests {}
