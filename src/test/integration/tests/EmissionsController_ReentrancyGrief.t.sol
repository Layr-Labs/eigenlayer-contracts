// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "src/test/integration/IntegrationChecks.t.sol";
import "src/contracts/core/EmissionsController.sol";
import "src/contracts/interfaces/IEmissionsController.sol";

contract MaliciousReentrantToken is ERC20 {
    address immutable emissionsController;
    address immutable rewardsCoordinator;
    bool public attackEnabled;
    bool private attacking;

    constructor(address _ec, address _rc) ERC20("Malicious", "MAL") {
        emissionsController = _ec;
        rewardsCoordinator = _rc;
    }

    function mint(address to, uint amount) external {
        _mint(to, amount);
    }

    function enableAttack() external {
        attackEnabled = true;
    }

    function transferFrom(address from, address to, uint amount) public override returns (bool) {
        if (attackEnabled && msg.sender == rewardsCoordinator && !attacking) {
            attacking = true;
            try EmissionsController(emissionsController).pressButton(1) {} catch {}
            attacking = false;
        }
        return super.transferFrom(from, to, amount);
    }
}

contract Integration_EmissionsController_ReentrancyGrief is IntegrationCheckUtils, IEmissionsControllerTypes {
    MaliciousReentrantToken maliciousToken;
    address attacker = address(0xBAD);
    address incentiveCouncil;

    function setUp() public virtual override {
        super.setUp();

        cheats.prank(rewardsCoordinator.owner());
        rewardsCoordinator.setRewardsForAllSubmitter(address(emissionsController), true);

        maliciousToken = new MaliciousReentrantToken(address(emissionsController), address(rewardsCoordinator));

        incentiveCouncil = emissionsController.incentiveCouncil();
    }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    function _addDistribution(uint64 startEpoch, uint64 totalEpochs, uint64 weight) internal {
        IRewardsCoordinatorTypes.StrategyAndMultiplier[][] memory strategiesAndMultipliers =
            new IRewardsCoordinatorTypes.StrategyAndMultiplier[][](1);
        strategiesAndMultipliers[0] = new IRewardsCoordinatorTypes.StrategyAndMultiplier[](1);
        strategiesAndMultipliers[0][0] = IRewardsCoordinatorTypes.StrategyAndMultiplier({strategy: allStrats[0], multiplier: 1e18});

        cheats.prank(incentiveCouncil);
        emissionsController.addDistribution(
            Distribution({
                distributionType: DistributionType.RewardsForAllEarners,
                weight: weight,
                operatorSet: OperatorSet({avs: address(0), id: 0}),
                strategiesAndMultipliers: strategiesAndMultipliers,
                startEpoch: startEpoch,
                totalEpochs: totalEpochs
            })
        );
    }

    function _pressButton(uint length, uint expectedProcessed) internal {
        vm.recordLogs();
        emissionsController.pressButton(length);

        uint processed;
        Vm.Log[] memory logs = cheats.getRecordedLogs();
        for (uint i = 0; i < logs.length; ++i) {
            if (logs[i].topics.length > 0 && logs[i].topics[0] == IEmissionsControllerEvents.DistributionProcessed.selector) {
                ++processed;
            }
        }
        assertEq(processed, expectedProcessed, "processed != expectedProcessed");
    }

    /// -----------------------------------------------------------------------
    /// Tests
    /// -----------------------------------------------------------------------

    function test_addDistribution_reenter_pressButton_sweep() public {
        // Add simple distribution that runs for 1 epoch.
        _addDistribution({startEpoch: 1, totalEpochs: 1, weight: 10_000});

        // Attempt reentrancy attack via malicious token.
        // Previously, these assertions would revert effectively griefing the system.
        cheats.warp(emissionsController.EMISSIONS_START_TIME() + 1 weeks);
        assertTrue(emissionsController.isButtonPressable(), "button should be pressable before attack");
        _executeReentrancyAttack();
        assertTrue(emissionsController.isButtonPressable(), "button should be pressable after attack");

        // Process the distribution, expecting it to succeed.
        _pressButton({length: 1, expectedProcessed: 1});

        // Attempt to sweep, no funds should be transferred.
        uint councilBalanceBefore = EIGEN.balanceOf(incentiveCouncil);
        emissionsController.sweep();
        assertEq(EIGEN.balanceOf(incentiveCouncil), councilBalanceBefore, "council balance should not have changed");
    }

    function _executeReentrancyAttack() internal {
        maliciousToken.mint(attacker, 1);
        maliciousToken.enableAttack();

        IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategiesAndMultipliers =
            new IRewardsCoordinatorTypes.StrategyAndMultiplier[](1);
        strategiesAndMultipliers[0] = IRewardsCoordinatorTypes.StrategyAndMultiplier({strategy: allStrats[0], multiplier: 1e18});

        IRewardsCoordinatorTypes.RewardsSubmission[] memory submissions = new IRewardsCoordinatorTypes.RewardsSubmission[](1);
        submissions[0] = IRewardsCoordinatorTypes.RewardsSubmission({
            strategiesAndMultipliers: strategiesAndMultipliers,
            token: IERC20(address(maliciousToken)),
            amount: 1,
            startTimestamp: uint32((block.timestamp - 7 days) / 86_400 * 86_400),
            duration: uint32(1 days)
        });

        cheats.startPrank(attacker);
        maliciousToken.approve(address(rewardsCoordinator), type(uint).max);
        rewardsCoordinator.createAVSRewardsSubmission(submissions);
        cheats.stopPrank();
    }
}
