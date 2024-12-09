// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/// core/
import "src/contracts/core/AllocationManager.sol";
import "src/contracts/core/AVSDirectory.sol";
import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/RewardsCoordinator.sol";
import "src/contracts/core/StrategyManager.sol";

/// permissions/
import "src/contracts/permissions/PauserRegistry.sol";
import "src/contracts/permissions/PermissionController.sol";

/// pods/
import "src/contracts/pods/EigenPod.sol";
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

        address[] memory pausers = new address[](3);
        pausers[0] = zAddress("pauserMultisig");
        pausers[1] = zAddress("operationsMultisig");
        pausers[2] = zAddress("executorMultisig");

        deployContract({
            name: type(PauserRegistry).name,
            deployedTo: address(new PauserRegistry({
                _pausers: pausers,
                _unpauser: zAddress("executorMultisig")
            }))
        });

        address permissionController_impl = deployImpl({
            name: type(PermissionController).name,
            deployedTo: address(new PermissionController())
        });

        deployProxy({
            name: type(PermissionController).name,
            deployedTo: address(new TransparentUpgradeableProxy({
                _logic: permissionController_impl,
                admin_: zDeployedContract(type(ProxyAdmin).name),
                _data: ""
            }))
        });

        /// core/

        address allocationManager_impl = deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(new AllocationManager({
                _delegation: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name)),
                _permissionController: PermissionController(zDeployedProxy(type(PermissionController).name)),
                _DEALLOCATION_DELAY: zUint32("MIN_WITHDRAWAL_DELAY"),
                _ALLOCATION_CONFIGURATION_DELAY: zUint32("ALLOCATION_CONFIGURATION_DELAY")
            }))
        });

        deployProxy({
            name: type(AllocationManager).name,
            deployedTo: address(new TransparentUpgradeableProxy({
                _logic: allocationManager_impl,
                admin_: zDeployedContract(type(ProxyAdmin).name),
                _data: ""
            }))
        });

        deployImpl({
            name: type(AVSDirectory).name,
            deployedTo: address(new AVSDirectory({
                _delegation: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
            }))
        });

        deployImpl({
            name: type(DelegationManager).name,
            deployedTo: address(new DelegationManager({
                _avsDirectory: AVSDirectory(zDeployedProxy(type(AVSDirectory).name)),
                _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
                _eigenPodManager: EigenPodManager(zDeployedProxy(type(EigenPodManager).name)),
                _allocationManager: AllocationManager(zDeployedProxy(type(AllocationManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name)),
                _permissionController: PermissionController(zDeployedProxy(type(PermissionController).name)),
                _MIN_WITHDRAWAL_DELAY: zUint32("MIN_WITHDRAWAL_DELAY")
            }))
        });

        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(new RewardsCoordinator({
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
            }))
        });

        deployImpl({
            name: type(StrategyManager).name,
            deployedTo: address(new StrategyManager({
                _delegation: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
            }))
        });

        /// pods/

        deployImpl({
            name: type(EigenPodManager).name,
            deployedTo: address(new EigenPodManager({
                _ethPOS: IETHPOSDeposit(zAddress("ethPOS")),
                _eigenPodBeacon: IBeacon(zDeployedProxy(type(EigenPod).name)),
                _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
                _delegationManager: DelegationManager(zDeployedProxy(type(DelegationManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
            }))
        });

        deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod({
                _ethPOS: IETHPOSDeposit(zAddress("ethPOS")),
                _eigenPodManager: EigenPodManager(zDeployedProxy(type(EigenPodManager).name)),
                _GENESIS_TIME: zUint64("EIGENPOD_GENESIS_TIME")
            }))
        });

        /// strategies/

        deployImpl({
            name: type(StrategyBaseTVLLimits).name,
            deployedTo: address(new StrategyBaseTVLLimits({
                _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
            }))
        });

        deployImpl({
            name: type(EigenStrategy).name,
            deployedTo: address(new EigenStrategy({
                _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
            }))
        });

        deployImpl({
            name: type(StrategyFactory).name,
            deployedTo: address(new StrategyFactory({
                _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
            }))
        });

        // for strategies deployed via factory
        deployImpl({
            name: type(StrategyBase).name,
            deployedTo: address(new StrategyBase({
                _strategyManager: StrategyManager(zDeployedProxy(type(StrategyManager).name)),
                _pauserRegistry: PauserRegistry(zDeployedContract(type(PauserRegistry).name))
            }))
        });

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {       
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
