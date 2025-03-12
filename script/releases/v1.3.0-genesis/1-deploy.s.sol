// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "src/contracts/libraries/BeaconChainProofs.sol";

import {IRewardsCoordinatorTypes} from "src/contracts/interfaces/IRewardsCoordinator.sol";

contract EmptyContract {}

contract DeployFresh is EOADeployer {
    using Env for *;

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

        deployBlankProxy({
            name: type(DelegationManager).name
        });

        deployBlankProxy({
            name: type(StrategyManager).name
        });
        deployBlankProxy({
            name: type(StrategyFactory).name
        });

        deployBlankProxy({
            name: type(AVSDirectory).name
        });

        deployBlankProxy({
            name: type(AllocationManager).name
        });

        deployBlankProxy({
            name: type(RewardsCoordinator).name
        });

        deployBlankProxy({
            name: type(EigenPodManager).name
        });

        deployBlankProxy({
            name: type(EigenPod).name
        });

        deployBlankProxy({
            name: type(PermissionController).name
        });

        address eigenPodImplementation = deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod(
                Env.ethPOS(),
                Env.proxy.eigenPodManager(),
                Env.EIGENPOD_GENESIS_TIME(),
                Env.deployToVersion()
            ))
        });

        // TODO: rm this in favor of a method on `EOADeployer`.
        deployContract({
            name: string.concat(type(EigenPod).name, "_Beacon"),
            deployedTo: address(new UpgradeableBeacon(eigenPodImplementation))
        });

        address[] memory pausers = new address[](3);
        pausers[0] = Env.pauserMultisig();
        pausers[1] = Env.opsMultisig();
        pausers[2] = Env.executorMultisig();

        // implementations
        deployImpl({
            name: type(PauserRegistry).name,
            deployedTo: address(new PauserRegistry(pausers, Env.executorMultisig()))
        });

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        deployImpl({
            name: type(DelegationManager).name,
            deployedTo: address(new DelegationManager(Env.proxy.strategyManager(), Env.proxy.eigenPodManager(), Env.proxy.allocationManager(), Env.impl.pauserRegistry(), Env.proxy.permissionController(), Env.MIN_WITHDRAWAL_DELAY(), Env.deployToVersion()))
        });

        deployImpl({
            name: type(StrategyManager).name,
            deployedTo: address(new StrategyManager(Env.proxy.delegationManager(), Env.impl.pauserRegistry(), Env.deployToVersion()))
        });

        deployImpl({
            name: type(AVSDirectory).name,
            deployedTo: address(new AVSDirectory(Env.proxy.delegationManager(), Env.impl.pauserRegistry(), Env.deployToVersion()))
        });

        deployImpl({
            name: type(EigenPodManager).name,
            deployedTo: address(new EigenPodManager(
                Env.ethPOS(),
                Env.beacon.eigenPod(),
                Env.proxy.delegationManager(),
                Env.impl.pauserRegistry(),
                Env.deployToVersion()
            ))
        });

        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(new RewardsCoordinator(IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams({
                delegationManager: Env.proxy.delegationManager(),
                strategyManager: Env.proxy.strategyManager(),
                allocationManager: Env.proxy.allocationManager(),
                pauserRegistry: Env.impl.pauserRegistry(),
                permissionController: Env.proxy.permissionController(),
                CALCULATION_INTERVAL_SECONDS: Env.CALCULATION_INTERVAL_SECONDS(),
                MAX_REWARDS_DURATION: Env.MAX_REWARDS_DURATION(),
                MAX_RETROACTIVE_LENGTH: Env.MAX_RETROACTIVE_LENGTH(),
                MAX_FUTURE_LENGTH: Env.MAX_FUTURE_LENGTH(),
                GENESIS_REWARDS_TIMESTAMP: Env.GENESIS_REWARDS_TIMESTAMP(),
                version: Env.deployToVersion()
            })))
        });

        deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(
                new AllocationManager(Env.proxy.delegationManager(), Env.impl.pauserRegistry(), Env.proxy.permissionController(), Env.MIN_WITHDRAWAL_DELAY(), Env.ALLOCATION_CONFIGURATION_DELAY(), Env.deployToVersion())
            )
        });

        deployImpl({
            name: type(PermissionController).name,
            deployedTo: address(
                new PermissionController(
                    Env.deployToVersion()
                )
            )
        });

        deployImpl({
            name: type(StrategyFactory).name,
            deployedTo: address(new StrategyFactory(Env.impl.strategyManager(), Env.impl.pauserRegistry(), Env.deployToVersion()))
        });
        
        /// strategies/
        deployImpl({
            name: type(StrategyBaseTVLLimits).name,
            deployedTo: address(new StrategyBaseTVLLimits({
                _strategyManager: Env.proxy.strategyManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _version: Env.deployToVersion()
            }))
        });

        deployImpl({
            name: type(EigenStrategy).name,
            deployedTo: address(new EigenStrategy({
                _strategyManager: Env.proxy.strategyManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _version: Env.deployToVersion()
            }))
        });

        deployImpl({
            name: type(StrategyFactory).name,
            deployedTo: address(new StrategyFactory({
                _strategyManager: Env.proxy.strategyManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _version: Env.deployToVersion()
            }))
        });

        // for strategies deployed via factory
        deployImpl({
            name: type(StrategyBase).name,
            deployedTo: address(new StrategyBase({
                _strategyManager: Env.proxy.strategyManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _version: Env.deployToVersion()
            }))
        });

        vm.stopBroadcast();
    }
    
    function testDeploy() public virtual {
        _runAsEOA();
    }
}
