// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade. 
 */
contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        /// permissions/

        address[] memory pausers = new address[](3);
        pausers[0] = Env.pauserMultisig();
        pausers[1] = Env.opsMultisig();
        pausers[2] = Env.executorMultisig();

        deployImpl({
            name: type(PauserRegistry).name,
            deployedTo: address(new PauserRegistry({
                _pausers: pausers,
                _unpauser: Env.executorMultisig()
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
                admin_: Env.proxyAdmin(),
                _data: ""
            }))
        });

        /// core/

        deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(new AllocationManager({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _permissionController: Env.proxy.permissionController(),
                _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY()
            }))
        });

        deployProxy({
            name: type(AllocationManager).name,
            deployedTo: address(new TransparentUpgradeableProxy({
                _logic: address(Env.impl.allocationManager()),
                admin_: Env.proxyAdmin(),
                _data: ""
            }))
        });

        deployImpl({
            name: type(AVSDirectory).name,
            deployedTo: address(new AVSDirectory({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
        });

        deployImpl({
            name: type(DelegationManager).name,
            deployedTo: address(new DelegationManager({
                _avsDirectory: Env.proxy.avsDirectory(),
                _strategyManager: Env.proxy.strategyManager(),
                _eigenPodManager: Env.proxy.eigenPodManager(),
                _allocationManager: Env.proxy.allocationManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _permissionController: Env.proxy.permissionController(),
                _MIN_WITHDRAWAL_DELAY: Env.MIN_WITHDRAWAL_DELAY()
            }))
        });

        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(new RewardsCoordinator({
                _delegationManager: Env.proxy.delegationManager(),
                _strategyManager: Env.proxy.strategyManager(),
                _allocationManager: Env.proxy.allocationManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _permissionController: Env.proxy.permissionController(),
                _CALCULATION_INTERVAL_SECONDS: Env.CALCULATION_INTERVAL_SECONDS(),
                _MAX_REWARDS_DURATION: Env.MAX_REWARDS_DURATION(),
                _MAX_RETROACTIVE_LENGTH: Env.MAX_RETROACTIVE_LENGTH(),
                _MAX_FUTURE_LENGTH: Env.MAX_FUTURE_LENGTH(),
                _GENESIS_REWARDS_TIMESTAMP: Env.GENESIS_REWARDS_TIMESTAMP()
            }))
        });

        deployImpl({
            name: type(StrategyManager).name,
            deployedTo: address(new StrategyManager({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
        });

        /// pods/

        deployImpl({
            name: type(EigenPodManager).name,
            deployedTo: address(new EigenPodManager({
                _ethPOS: Env.ethPOS(),
                _eigenPodBeacon: IBeacon(address(Env.proxy.eigenPod())),
                _strategyManager: Env.proxy.strategyManager(),
                _delegationManager: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
        });

        deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod({
                _ethPOS: Env.ethPOS(),
                _eigenPodManager: Env.proxy.eigenPodManager(),
                _GENESIS_TIME: Env.EIGENPOD_GENESIS_TIME()
            }))
        });

        /// strategies/

        deployImpl({
            name: type(StrategyBaseTVLLimits).name,
            deployedTo: address(new StrategyBaseTVLLimits({
                _strategyManager: Env.proxy.strategyManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
        });

        deployImpl({
            name: type(EigenStrategy).name,
            deployedTo: address(new EigenStrategy({
                _strategyManager: Env.proxy.strategyManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
        });

        deployImpl({
            name: type(StrategyFactory).name,
            deployedTo: address(new StrategyFactory({
                _strategyManager: Env.proxy.strategyManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
        });

        // for strategies deployed via factory
        deployImpl({
            name: type(StrategyBase).name,
            deployedTo: address(new StrategyBase({
                _strategyManager: Env.proxy.strategyManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
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
