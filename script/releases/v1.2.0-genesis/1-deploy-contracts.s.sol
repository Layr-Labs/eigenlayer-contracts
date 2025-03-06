// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "src/contracts/libraries/BeaconChainProofs.sol";

contract EmptyContract {}

contract DeployFresh is EOADeployer {
    using Env for *;

    // helper fn
    function deployBlankProxy(string memory name) private returns (address) {
        return deployProxy({
            name: name,
            deployedTo: address(new TransparentUpgradeableProxy(
                address(new EmptyContract()),
                Env.proxyAdmin(),
                ""
            ))
        });
    }

    function _runAsEOA() internal override {
        vm.startBroadcast();
        address[] memory pausers = new address[](3);
        pausers[0] = Env.pauserMultisig();
        pausers[1] = Env.opsMultisig();
        pausers[2] = Env.executorMultisig();

        deployBlankProxy({
            admin: 
        });

        deployBlankProxy({
            name: type(PauserRegistry).name,
            implementation: address(new PauserRegistry({
                _pausers: pausers,
                _unpauser: Env.executorMultisig()
            })),
            admin: Env.proxyAdmin(),
            data: ""
        });

        deployProxyAndImpl({
            name: type(PermissionController).name,
            implementation: address(new PermissionController()),
            admin: Env.proxyAdmin(),
            data: ""
        });

        deployProxyAndImpl({
            name: type(DelegationManager).name,
            implementation: address(new DelegationManager({
                _strategyManager: Env.proxy.strategyManager(),
                _eigenPodManager: Env.proxy.eigenPodManager(),
                _allocationManager: Env.proxy.allocationManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _permissionController: Env.proxy.permissionController(),
                _MIN_WITHDRAWAL_DELAY: Env.MIN_WITHDRAWAL_DELAY()
            })),
            admin: Env.proxyAdmin(),
            data: ""
        });

        deployProxyAndImpl({
            name: type(StrategyManager).name,
            implementation: address(new StrategyManager({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            })),
            admin: Env.proxyAdmin(),
            data: ""
        });

        deployProxyAndImpl({
            name: type(AVSDirectory).name,
            implementation: address(new AVSDirectory({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            })),
            admin: Env.proxyAdmin(),
            data: ""
        });

        /// core
        deployProxyAndImpl({
            name: type(AllocationManager).name,
            implementation: address(new AllocationManager({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _permissionController: Env.proxy.permissionController(),
                _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY()
            })),
            admin: Env.proxyAdmin(),
            data: abi.encodeCall(
                AllocationManager.initialize,
                (
                    Env.executorMultisig(), // initialOwner
                    0                       // initialPausedStatus
                )
            )
        });


        deployProxyAndImpl({
            name: type(RewardsCoordinator).name,
            implementation: address(new RewardsCoordinator({
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
            })),
            admin: Env.proxyAdmin(),
            data: ""
        });

        deployProxyAndImpl({
            name: type(EigenPodManager).name,
            implementation: address(new EigenPodManager({
                _ethPOS: Env.ethPOS(),
                _eigenPodBeacon: Env.beacon.eigenPod(),
                _delegationManager: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            })),
            admin: Env.proxyAdmin(),
            data: ""
        });

        deployProxyAndImpl({
            name: type(EigenPod).name,
            implementation: address(new EigenPod({
                _ethPOS: Env.ethPOS(),
                _eigenPodManager: Env.proxy.eigenPodManager(),
                _GENESIS_TIME: Env.EIGENPOD_GENESIS_TIME()
            })),
            admin: Env.proxyAdmin(),
            data: ""
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
    
    function testScript() public virtual {
        _runAsEOA();
    }
}
