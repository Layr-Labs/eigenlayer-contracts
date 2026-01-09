// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";
import "../Env.sol";
import "../TestUtils.sol";

import {EmissionsController} from "src/contracts/core/EmissionsController.sol";

/// Purpose: Deploy new implementations for RewardsCoordinator and EmissionsController
contract DeployImplementations is EOADeployer {
    using Env for *;

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        // Deploy new RewardsCoordinator implementation to account for new fee changes
        deployRewardsCoordinator();

        // Deploy new EmissionsController implementation
        deployEmissionsController();

        vm.stopBroadcast();
    }

    function deployRewardsCoordinator() internal onlyEOA returns (RewardsCoordinator deployed) {
        deployed = new RewardsCoordinator({
            params: IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams({
                delegationManager: Env.proxy.delegationManager(),
                strategyManager: Env.proxy.strategyManager(),
                allocationManager: Env.proxy.allocationManager(),
                pauserRegistry: Env.impl.pauserRegistry(),
                permissionController: Env.proxy.permissionController(),
                CALCULATION_INTERVAL_SECONDS: Env.CALCULATION_INTERVAL_SECONDS(),
                MAX_REWARDS_DURATION: Env.MAX_REWARDS_DURATION(),
                MAX_RETROACTIVE_LENGTH: Env.MAX_RETROACTIVE_LENGTH(),
                MAX_FUTURE_LENGTH: Env.MAX_FUTURE_LENGTH(),
                GENESIS_REWARDS_TIMESTAMP: Env.GENESIS_REWARDS_TIMESTAMP()
            })
        });
        deployImpl({name: type(RewardsCoordinator).name, deployedTo: address(deployed)});
    }

    function deployEmissionsController() internal onlyEOA returns (EmissionsController deployed) {
        deployed = new EmissionsController({
            eigen: IEigen(address(Env.proxy.eigen())),
            backingEigen: Env.proxy.beigen(),
            rewardsCoordinator: Env.proxy.rewardsCoordinator(),
            pauserRegistry: Env.impl.pauserRegistry(),
            inflationRate: Env.EMISSIONS_INFLATION_RATE(),
            startTime: Env.EMISSIONS_START_TIME(),
            cooldownSeconds: Env.EMISSIONS_COOLDOWN_SECONDS()
        });
        deployImpl({name: type(EmissionsController).name, deployedTo: address(deployed)});
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        // Deploy implementations as EOA
        runAsEOA();

        // Validate implementations
        TestUtils.validateEmissionsControllerInitialized(Env.impl.emissionsController());
        TestUtils.validateRewardsCoordinatorConstructor(Env.impl.rewardsCoordinator());
    }
}

