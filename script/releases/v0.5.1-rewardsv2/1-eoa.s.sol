// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {IDelegationManager} from "src/contracts/interfaces/IDelegationManager.sol";
import {DelegationManager} from "src/contracts/core/DelegationManager.sol";
import {StrategyManager} from "src/contracts/core/StrategyManager.sol";
import {Env} from "../Env.sol";
import {Test, console} from "forge-std/Test.sol";
import {IPauserRegistry} from "src/contracts/interfaces/IPauserRegistry.sol";
import {ZEnvHelpers, State} from "zeus-templates/utils/ZEnvHelpers.sol";  

contract Deploy is EOADeployer {
    using Env for *;
    using ZEnvHelpers for *;

    function _runAsEOA() internal override {
        zUpdateUint16(string("REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS"), uint16(1000));

        State storage state = ZEnvHelpers.state();

        vm.startBroadcast();

        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(
                new RewardsCoordinator(
                    Env.proxy.delegationManager(),
                    Env.proxy.strategyManager(),
                    state.envU32("REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"),
                    state.envU32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION"),
                    state.envU32("REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH"),
                    state.envU32("REWARDS_COORDINATOR_MAX_FUTURE_LENGTH"),
                    state.envU32("REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP")
                )
            )
        });

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        // Deploy RewardsCoordinator Implementation
        address oldImpl = address(Env.impl.rewardsCoordinator());
        runAsEOA();
        address newImpl = address(Env.impl.rewardsCoordinator());
        assertTrue(oldImpl != newImpl, "impl should be different");

        Deployment[] memory deploys = deploys();

        // sanity check that deployed impl is returning our deployment.
        assertEq(deploys[0].deployedTo, address(Env.impl.rewardsCoordinator()));

        RewardsCoordinator rewardsCoordinatorImpl = Env.impl.rewardsCoordinator();

        State storage state = ZEnvHelpers.state();

        address owner = address(Env.opsMultisig());
        IPauserRegistry pauserRegistry = Env.impl.pauserRegistry();
        uint64 initPausedStatus = state.envU64("REWARDS_COORDINATOR_INIT_PAUSED_STATUS");
        address rewardsUpdater = state.envAddress("REWARDS_COORDINATOR_UPDATER");
        uint32 activationDelay = state.envU32("REWARDS_COORDINATOR_ACTIVATION_DELAY");
        uint16 defaultOperatorSplitBips = state.envU16("REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS");

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
            address(Env.proxy.delegationManager()),
            "expected delegationManager"
        );
        assertEq(
            address(rewardsCoordinatorImpl.strategyManager()),
            address(Env.proxy.strategyManager()),
            "expected strategyManager"
        );

        assertEq(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            state.envU32("REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"),
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertGt(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            0,
            "expected non-zero REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );

        assertEq(rewardsCoordinatorImpl.MAX_REWARDS_DURATION(), state.envU32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION"));
        assertGt(rewardsCoordinatorImpl.MAX_REWARDS_DURATION(), 0);

        assertEq(
            rewardsCoordinatorImpl.MAX_RETROACTIVE_LENGTH(),
            state.envU32("REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH")
        );
        assertGt(rewardsCoordinatorImpl.MAX_RETROACTIVE_LENGTH(), 0);

        assertEq(rewardsCoordinatorImpl.MAX_FUTURE_LENGTH(), state.envU32("REWARDS_COORDINATOR_MAX_FUTURE_LENGTH"));
        assertGt(rewardsCoordinatorImpl.MAX_FUTURE_LENGTH(), 0);

        assertEq(
            rewardsCoordinatorImpl.GENESIS_REWARDS_TIMESTAMP(),
            state.envU32("REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP")
        );
        assertGt(rewardsCoordinatorImpl.GENESIS_REWARDS_TIMESTAMP(), 0);

        assertEq(deploys.length, 1, "expected exactly 1 deployment");
        assertEq(
            keccak256(bytes(deploys[0].name)),
            keccak256(bytes("RewardsCoordinator_Impl")),
            "zeusTest: Deployment name is not RewardsCoordinator"
        );
        assertTrue(deploys[0].singleton == true, "zeusTest: RewardsCoordinator should be a singleton.");
        assertNotEq(deploys[0].deployedTo, address(0), "zeusTest: Should deploy to non-zero address.");
    }
}
