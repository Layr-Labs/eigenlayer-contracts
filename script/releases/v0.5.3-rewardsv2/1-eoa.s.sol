// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {IDelegationManager} from "src/contracts/interfaces/IDelegationManager.sol";
import {DelegationManager} from "src/contracts/core/DelegationManager.sol";
import {StrategyManager} from "src/contracts/core/StrategyManager.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";
import {Test, console} from "forge-std/Test.sol";
import {IPauserRegistry} from "src/contracts/interfaces/IPauserRegistry.sol";

contract Deploy is EOADeployer {
    using EigenLabsUpgrade for *;

    function _runAsEOA() internal override {
        zUpdateUint16(string("REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS"), uint16(1000));
        zUpdateUint32(string("REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"), uint32(1 days));

        vm.startBroadcast();
        deploySingleton(
            address(
                new RewardsCoordinator(
                    IDelegationManager(zDeployedProxy(type(DelegationManager).name)),
                    StrategyManager(zDeployedProxy(type(StrategyManager).name)),
                    zUint32("REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"),
                    zUint32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION"),
                    zUint32("REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH"),
                    zUint32("REWARDS_COORDINATOR_MAX_FUTURE_LENGTH"),
                    zUint32("REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP")
                )
            ),
            this.impl(type(RewardsCoordinator).name)
        );

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        // Deploy RewardsCoordinator Implementation
        address oldImpl = zDeployedImpl(type(RewardsCoordinator).name);
        runAsEOA();
        address newImpl = zDeployedImpl(type(RewardsCoordinator).name);
        assertTrue(oldImpl != newImpl, "impl should be different");

        Deployment[] memory deploys = deploys();

        // sanity check that zDeployedImpl is returning our deployment.
        assertEq(deploys[0].deployedTo, zDeployedImpl(type(RewardsCoordinator).name));

        RewardsCoordinator rewardsCoordinatorImpl = RewardsCoordinator(zDeployedImpl(type(RewardsCoordinator).name));

        address owner = this._operationsMultisig();
        IPauserRegistry pauserRegistry = IPauserRegistry(this._pauserRegistry());
        uint64 initPausedStatus = zUint64("REWARDS_COORDINATOR_INIT_PAUSED_STATUS");
        address rewardsUpdater = zAddress("REWARDS_COORDINATOR_UPDATER");
        uint32 activationDelay = zUint32("REWARDS_COORDINATOR_ACTIVATION_DELAY");
        uint16 defaultOperatorSplitBips = zUint16("REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS");

        // Ensure that the implementation contract cannot be initialized.
        vm.expectRevert("Initializable: contract is already initialized");
        rewardsCoordinatorImpl.initialize(
            owner,
            pauserRegistry,
            initPausedStatus,
            rewardsUpdater,
            activationDelay,
            defaultOperatorSplitBips
        );

        // Assert Immutables and State Variables set through initialize
        assertEq(rewardsCoordinatorImpl.owner(), address(0), "expected owner");
        assertEq(address(rewardsCoordinatorImpl.pauserRegistry()), address(0), "expected pauserRegistry");
        assertEq(address(rewardsCoordinatorImpl.rewardsUpdater()), address(0), "expected rewardsUpdater");
        assertEq(rewardsCoordinatorImpl.activationDelay(), 0, "expected activationDelay");
        assertEq(rewardsCoordinatorImpl.defaultOperatorSplitBips(), 0, "expected defaultOperatorSplitBips");

        assertEq(
            address(rewardsCoordinatorImpl.delegationManager()),
            zDeployedProxy(type(DelegationManager).name),
            "expected delegationManager"
        );
        assertEq(
            address(rewardsCoordinatorImpl.strategyManager()),
            zDeployedProxy(type(StrategyManager).name),
            "expected strategyManager"
        );

        assertEq(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            zUint32("REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"),
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertEq(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            1 days,
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertGt(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            0,
            "expected non-zero REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );

        assertEq(rewardsCoordinatorImpl.MAX_REWARDS_DURATION(), zUint32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION"));
        assertGt(rewardsCoordinatorImpl.MAX_REWARDS_DURATION(), 0);

        assertEq(
            rewardsCoordinatorImpl.MAX_RETROACTIVE_LENGTH(),
            zUint32("REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH")
        );
        assertGt(rewardsCoordinatorImpl.MAX_RETROACTIVE_LENGTH(), 0);

        assertEq(rewardsCoordinatorImpl.MAX_FUTURE_LENGTH(), zUint32("REWARDS_COORDINATOR_MAX_FUTURE_LENGTH"));
        assertGt(rewardsCoordinatorImpl.MAX_FUTURE_LENGTH(), 0);

        assertEq(
            rewardsCoordinatorImpl.GENESIS_REWARDS_TIMESTAMP(),
            zUint32("REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP")
        );
        assertGt(rewardsCoordinatorImpl.GENESIS_REWARDS_TIMESTAMP(), 0);

        assertEq(deploys.length, 1, "expected exactly 1 deployment");
        assertEq(
            keccak256(bytes(deploys[0].name)),
            keccak256(bytes(this.impl(type(RewardsCoordinator).name))),
            "zeusTest: Deployment name is not RewardsCoordinator"
        );
        assertTrue(deploys[0].singleton == true, "zeusTest: RewardsCoordinator should be a singleton.");
        assertNotEq(deploys[0].deployedTo, address(0), "zeusTest: Should deploy to non-zero address.");
    }
}
