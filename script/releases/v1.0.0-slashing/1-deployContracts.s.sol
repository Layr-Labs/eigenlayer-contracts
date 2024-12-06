// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/// core/
import {AllocationManager} from "src/contracts/core/AllocationManager.sol";
import {AVSDirectory} from "src/contracts/core/AVSDirectory.sol";
import {DelegationManager} from "src/contracts/core/DelegationManager.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {StrategyManager} from "src/contracts/core/StrategyManager.sol";

/// permissions/
import {PauserRegistry} from "src/contracts/permissions/PauserRegistry.sol";
import {PermissionController} from "src/contracts/permissions/PermissionController.sol";

/// pods/
import {EigenPod} from "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodManager.sol";

/// strategies/
import "src/contracts/strategies/EigenStrategy.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/strategies/StrategyBaseTVLLimits.sol";
import "src/contracts/strategies/StrategyFactory.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade. 
 */
contract Deploy is EOADeployer {
    using EigenLabsUpgrade for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        /// permissions/

        // TODO: pauser registry?

        PermissionController permissionController_impl = new PermissionController();

        deploySingleton({
            deployedTo: address(permissionController_impl),
            name: this.impl(type(PermissionController).name)
        });

        TransparentUpgradeableProxy permissionController_proxy = new TransparentUpgradeableProxy({
            _logic: address(permissionController_impl),
            admin_: zDeployedContract(type(ProxyAdmin).name),
            _data: ""
        });

        deploySingleton({
            deployedTo: address(permissionController_proxy),
            name: this.proxy(type(PermissionController).name)
        });

        /// core/

        AllocationManager allocationManager_impl = new AllocationManager({
            _delegation: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name)),
            _permissionController: PermissionController(zDeployedProxy(type(PermissionController).name)),
            _DEALLOCATION_DELAY: zUint32("MIN_WITHDRAWAL_DELAY"),
            _ALLOCATION_CONFIGURATION_DELAY: zUint32("ALLOCATION_CONFIGURATION_DELAY")
        });

        deploySingleton({
            deployedTo: address(allocationManager_impl),
            name: this.impl(type(AllocationManager).name)
        });

        AVSDirectory avsDirectory_impl = new AVSDirectory({
            _delegation: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
        });

        deploySingleton({
            deployedTo: address(avsDirectory_impl),
            name: this.impl(type(AVSDirectory).name)
        });

        DelegationManager delegationManager_impl = new DelegationManager({
            _avsDirectory: AVSDirectory(zDeployedProxy(type(AVSDirectory).name)),
            _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
            _eigenPodManager: EigenPodManager(zDeployedProxy(type(EigenPodManager).name)),
            _allocationManager: AllocationManager(zDeployedProxy(type(AllocationManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name)),
            _permissionController: PermissionController(zDeployedProxy(type(PermissionController).name)),
            _MIN_WITHDRAWAL_DELAY: zUint32("MIN_WITHDRAWAL_DELAY")
        });

        deploySingleton({
            deployedTo: address(delegationManager_impl),
            name: this.impl(type(DelegationManager).name)
        });

        RewardsCoordinator rewardsCoordinator_impl = new RewardsCoordinator({
            _delegationManager: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
            _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
            _allocationManager: AllocationManager(zDeployedProxy(type(AllocationManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name)),
            _permissionController: PermissionController(zDeployedProxy(type(PermissionController).name)),
            _CALCULATION_INTERVAL_SECONDS: zUint32("CALCULATION_INTERVAL_SECONDS"),
            _MAX_REWARDS_DURATION: zUint32("MAX_REWARDS_DURATION"),
            _MAX_RETROACTIVE_LENGTH: zUint32("MAX_RETROACTIVE_LENGTH"),
            _MAX_FUTURE_LENGTH: zUint32("MAX_FUTURE_LENGTH"),
            _GENESIS_REWARDS_TIMESTAMP: zUint32("GENESIS_REWARDS_TIMESTAMP")
        });

        deploySingleton({
            deployedTo: address(rewardsCoordinator_impl),
            name: this.impl(type(RewardsCoordinator).name)
        });

        StrategyManager strategyManager_impl = new StrategyManager({
            _delegation: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
        });

        deploySingleton({
            deployedTo: address(strategyManager_impl),
            name: this.impl(type(StrategyManager).name)
        });

        /// pods/

        EigenPodManager eigenPodManager_impl = new EigenPodManager({
            _ethPOS: IETHPOSDeposit(zAddress("ethPOS")),
            _eigenPodBeacon: IBeacon(zDeployedProxy(type(EigenPod).name)),
            _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
            _delegationManager: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
        });

        deploySingleton({
            deployedTo: address(eigenPodManager_impl),
            name: this.impl(type(EigenPodManager).name)
        });

        EigenPod eigenPod_impl = new EigenPod({
            _ethPOS: IETHPOSDeposit(zAddress("ethPOS")),
            _eigenPodManager: EigenPodManager(zDeployedProxy(type(EigenPodManager).name)),
            _GENESIS_TIME: zUint64("EIGENPOD_GENESIS_TIME")
        });

        deploySingleton({
            deployedTo: address(eigenPod_impl),
            name: this.impl(type(EigenPod).name)
        });

        /// strategies/

        StrategyBaseTVLLimits preLongtailStrategy_impl = new StrategyBaseTVLLimits({
            _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
        });

        deploySingleton({
            deployedTo: address(preLongtailStrategy_impl),
            name: this.impl(type(StrategyBaseTVLLimits).name)
        });

        EigenStrategy eigenStrategy_impl = new EigenStrategy({
            _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
        });

        deploySingleton({
            deployedTo: address(eigenStrategy_impl),
            name: this.impl(type(EigenStrategy).name)
        });

        StrategyFactory strategyFactory_impl = new StrategyFactory({
            _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
        });

        deploySingleton({
            deployedTo: address(strategyFactory_impl),
            name: this.impl(type(StrategyFactory).name)
        });

        // for strategies deployed via factory
        StrategyBase strategyBase_impl = new StrategyBase({
            _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
            _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
        });

        deploySingleton({
            deployedTo: address(strategyBase_impl),
            name: this.impl(type(StrategyBase).name)
        });

        vm.stopBroadcast();
    }

    function testDeploy() public {
        _runAsEOA();
        Deployment[] memory deploys = deploys();

        emit log_named_uint("num deploys", deploys.length);

        for (uint i = 0; i < deploys.length; i++) {
            emit log_named_uint("deploy ", i);
            emit log_named_string(" - name", deploys[i].name);
            emit log_named_address(" - deployedTo", deploys[i].deployedTo);
        }
    }
}
