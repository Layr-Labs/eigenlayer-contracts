// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {DeployProtocolRegistryProxy} from "./1-deployProtocolRegistryProxy.s.sol";
import {DeployProtocolRegistryImpl} from "./2-deployProtocolRegistryImpl.s.sol";
import {UpgradeProtocolRegistry} from "./3-upgradeProtocolRegistry.s.sol";
import "../Env.sol";
import "../TestUtils.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade.
 * Contracts deployed:
 * /// Permissions
 * - PermissionController
 * - KeyRegistrar
 * /// Core
 * - AllocationManager
 * - AVSDirectory
 * - DelegationManager
 * - ReleaseManager
 * - RewardsCoordinator
 * - StrategyManager
 * /// Pods
 * - EigenPod
 * - EigenPodManager
 * /// Strategies
 * - EigenStrategy
 * - StrategyBase
 * - StrategyBaseTVLLimits
 * - StrategyFactory
 * /// Multichain
 * - CrossChainRegistry
 * - OperatorTableUpdater
 * - ECDSACertificateVerifier
 * - BN254CertificateVerifier
 * /// AVS
 * - TaskMailbox
 */
contract Deploy is UpgradeProtocolRegistry {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        /// permissions/

        deployImpl({name: type(PermissionController).name, deployedTo: address(new PermissionController())});

        deployImpl({
            name: type(KeyRegistrar).name,
            deployedTo: address(
                new KeyRegistrar({
                    _permissionController: Env.proxy.permissionController(),
                    _allocationManager: Env.proxy.allocationManager(),
                    _version: Env.deployVersion()
                })
            )
        });

        /// core/

        deployImpl({
            name: type(AllocationManagerView).name,
            deployedTo: address(
                new AllocationManagerView({
                    _delegation: Env.proxy.delegationManager(),
                    _eigenStrategy: Env.proxy.eigenStrategy(),
                    _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                    _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY()
                })
            )
        });

        deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(
                new AllocationManager({
                    _allocationManagerView: Env.impl.allocationManagerView(),
                    _delegation: Env.proxy.delegationManager(),
                    _eigenStrategy: Env.proxy.eigenStrategy(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _permissionController: Env.proxy.permissionController(),
                    _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                    _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY()
                })
            )
        });

        deployImpl({
            name: type(AVSDirectory).name,
            deployedTo: address(
                new AVSDirectory({
                    _delegation: Env.proxy.delegationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(DelegationManager).name,
            deployedTo: address(
                new DelegationManager({
                    _strategyManager: Env.proxy.strategyManager(),
                    _eigenPodManager: Env.proxy.eigenPodManager(),
                    _allocationManager: Env.proxy.allocationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _permissionController: Env.proxy.permissionController(),
                    _MIN_WITHDRAWAL_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                    _version: Env.deployVersion()
                })
            )
        });

        // Deploy impl AND proxy for protocol registry
        deployImpl({name: type(ProtocolRegistry).name, deployedTo: address(new ProtocolRegistry())});

        deployImpl({
            name: type(ReleaseManager).name,
            deployedTo: address(new ReleaseManager({_permissionController: Env.proxy.permissionController()}))
        });

        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(
                new RewardsCoordinator({
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
                })
            )
        });

        deployImpl({
            name: type(StrategyManager).name,
            deployedTo: address(
                new StrategyManager({
                    _allocationManager: Env.proxy.allocationManager(),
                    _delegation: Env.proxy.delegationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        /// pods/

        deployImpl({
            name: type(EigenPodManager).name,
            deployedTo: address(
                new EigenPodManager({
                    _ethPOS: Env.ethPOS(),
                    _eigenPodBeacon: Env.beacon.eigenPod(),
                    _delegationManager: Env.proxy.delegationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry()
                })
            )
        });

        deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod({_ethPOS: Env.ethPOS(), _eigenPodManager: Env.proxy.eigenPodManager()}))
        });

        /// strategies/

        deployImpl({
            name: type(EigenStrategy).name,
            deployedTo: address(
                new EigenStrategy({
                    _strategyManager: Env.proxy.strategyManager(),
                    _pauserRegistry: Env.impl.pauserRegistry()
                })
            )
        });

        // for strategies deployed via factory
        deployImpl({
            name: type(StrategyBase).name,
            deployedTo: address(
                new StrategyBase({_strategyManager: Env.proxy.strategyManager(), _pauserRegistry: Env.impl.pauserRegistry()})
            )
        });

        deployImpl({
            name: type(StrategyBaseTVLLimits).name,
            deployedTo: address(
                new StrategyBaseTVLLimits({
                    _strategyManager: Env.proxy.strategyManager(),
                    _pauserRegistry: Env.impl.pauserRegistry()
                })
            )
        });

        deployImpl({
            name: type(StrategyFactory).name,
            deployedTo: address(
                new StrategyFactory({
                    _strategyManager: Env.proxy.strategyManager(),
                    _pauserRegistry: Env.impl.pauserRegistry()
                })
            )
        });

        /// multichain/

        deployImpl({
            name: type(CrossChainRegistry).name,
            deployedTo: address(
                new CrossChainRegistry({
                    _allocationManager: Env.proxy.allocationManager(),
                    _keyRegistrar: Env.proxy.keyRegistrar(),
                    _permissionController: Env.proxy.permissionController(),
                    _pauserRegistry: Env.impl.pauserRegistry()
                })
            )
        });

        deployImpl({
            name: type(OperatorTableUpdater).name,
            deployedTo: address(
                new OperatorTableUpdater({
                    _bn254CertificateVerifier: Env.proxy.bn254CertificateVerifier(),
                    _ecdsaCertificateVerifier: Env.proxy.ecdsaCertificateVerifier(),
                    _pauserRegistry: Env.impl.pauserRegistry()
                })
            )
        });

        deployImpl({
            name: type(ECDSACertificateVerifier).name,
            deployedTo: address(
                new ECDSACertificateVerifier({
                    _operatorTableUpdater: Env.proxy.operatorTableUpdater(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(BN254CertificateVerifier).name,
            deployedTo: address(new BN254CertificateVerifier({_operatorTableUpdater: Env.proxy.operatorTableUpdater()}))
        });

        /// avs/

        deployImpl({
            name: type(TaskMailbox).name,
            deployedTo: address(
                new TaskMailbox({
                    _bn254CertificateVerifier: address(Env.proxy.bn254CertificateVerifier()),
                    _ecdsaCertificateVerifier: address(Env.proxy.ecdsaCertificateVerifier()),
                    _maxTaskSLA: Env.MAX_TASK_SLA()
                })
            )
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual override {
        // Deploy protocol registry and initialize it
        UpgradeProtocolRegistry.testScript();

        // Deploy the core contracts
        runAsEOA();

        // Run tests
        TestUtils.validateProxyAdmins();
        TestUtils.validateImplConstructors();
        TestUtils.validateImplsNotInitializable();

        // Check individual version addresses
        TestUtils.validateKeyRegistrarVersion();
        TestUtils.validateAVSDirectoryVersion();
        TestUtils.validateDelegationManagerVersion();
        TestUtils.validateStrategyManagerVersion();
        TestUtils.validateECDSACertificateVerifierVersion();
    }
}
