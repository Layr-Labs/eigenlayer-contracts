// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployToken} from "script/releases/v0.0.0-source-genesis/3-deployToken.s.sol";
import {DeployPauser} from "script/releases/v0.0.0-source-genesis/2-deployPauser.s.sol";
import {DeployGovernance} from "script/releases/v0.0.0-source-genesis/1-deployGovernance.s.sol";
import "../Env.sol";

/// This script deploys the following contracts/proxies:
/// - DelegationManager
/// - StrategyManager
/// - AllocationManager
/// - RewardsCoordinator
/// - AVSDirectory
/// - EigenPodManager
/// - StrategyFactory
/// - ReleaseManager
/// - PermissionController
/// - KeyRegistrar
/// - CrossChainRegistry
/// - EigenStrategy
/// - EigenPod - Implementation & Beacon
/// - StrategyBase - Implementation & Beacon
contract DeployCore is DeployToken {
    using Env for *;

    EmptyContract emptyContract;

    function deployBlankProxy(
        string memory name
    ) private returns (address) {
        return deployProxy({
            name: name,
            deployedTo: address(new TransparentUpgradeableProxy(address(emptyContract), Env.proxyAdmin(), ""))
        });
    }

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        // 0. Deploy the empty contract
        emptyContract = new EmptyContract();

        // 1. Deploy empty proxies for the core contracts
        deployBlankProxy({name: type(DelegationManager).name});

        deployBlankProxy({name: type(StrategyManager).name});

        deployBlankProxy({name: type(AllocationManager).name});

        deployBlankProxy({name: type(RewardsCoordinator).name});

        deployBlankProxy({name: type(AVSDirectory).name});

        deployBlankProxy({name: type(EigenPodManager).name});

        deployBlankProxy({name: type(StrategyFactory).name});

        deployBlankProxy({name: type(EigenStrategy).name});

        deployBlankProxy({name: type(PermissionController).name});

        deployBlankProxy({name: type(KeyRegistrar).name});

        deployBlankProxy({name: type(CrossChainRegistry).name});

        deployBlankProxy({name: type(ReleaseManager).name});

        // 2. Deploy the beacons and their associated implementations
        deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod(Env.ethPOS(), Env.proxy.eigenPodManager(), Env.deployVersion()))
        });

        deployBeacon({
            name: type(EigenPod).name,
            deployedTo: address(new UpgradeableBeacon(address(Env.impl.eigenPod())))
        });

        deployImpl({
            name: type(StrategyBase).name,
            deployedTo: address(
                new StrategyBase({
                    _strategyManager: Env.proxy.strategyManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployBeacon({
            name: type(StrategyBase).name,
            deployedTo: address(new UpgradeableBeacon(address(Env.impl.strategyBase())))
        });

        // 3. Deploy the non-beacon implementation contracts
        deployImpl({
            name: type(DelegationManager).name,
            deployedTo: address(
                new DelegationManager(
                    Env.proxy.strategyManager(),
                    Env.proxy.eigenPodManager(),
                    Env.proxy.allocationManager(),
                    Env.impl.pauserRegistry(),
                    Env.proxy.permissionController(),
                    Env.MIN_WITHDRAWAL_DELAY(),
                    Env.deployVersion()
                )
            )
        });

        deployImpl({
            name: type(StrategyManager).name,
            deployedTo: address(
                new StrategyManager(
                    Env.proxy.allocationManager(),
                    Env.proxy.delegationManager(),
                    Env.impl.pauserRegistry(),
                    Env.deployVersion()
                )
            )
        });

        deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(
                new AllocationManager(
                    Env.proxy.delegationManager(),
                    Env.proxy.eigenStrategy(),
                    Env.impl.pauserRegistry(),
                    Env.proxy.permissionController(),
                    Env.MIN_WITHDRAWAL_DELAY(),
                    Env.ALLOCATION_CONFIGURATION_DELAY(),
                    Env.deployVersion()
                )
            )
        });

        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(
                new RewardsCoordinator(
                    IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams({
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
                        version: Env.deployVersion()
                    })
                )
            )
        });

        deployImpl({
            name: type(AVSDirectory).name,
            deployedTo: address(
                new AVSDirectory(Env.proxy.delegationManager(), Env.impl.pauserRegistry(), Env.deployVersion())
            )
        });

        deployImpl({
            name: type(EigenPodManager).name,
            deployedTo: address(
                new EigenPodManager(
                    Env.ethPOS(),
                    Env.beacon.eigenPod(),
                    Env.proxy.delegationManager(),
                    Env.impl.pauserRegistry(),
                    Env.deployVersion()
                )
            )
        });

        deployImpl({
            name: type(StrategyFactory).name,
            deployedTo: address(
                new StrategyFactory(Env.proxy.strategyManager(), Env.impl.pauserRegistry(), Env.deployVersion())
            )
        });

        deployImpl({
            name: type(EigenStrategy).name,
            deployedTo: address(
                new EigenStrategy({
                    _strategyManager: Env.proxy.strategyManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(PermissionController).name,
            deployedTo: address(new PermissionController(Env.deployVersion()))
        });

        deployImpl({
            name: type(KeyRegistrar).name,
            deployedTo: address(
                new KeyRegistrar(Env.proxy.permissionController(), Env.proxy.allocationManager(), Env.deployVersion())
            )
        });

        deployImpl({
            name: type(CrossChainRegistry).name,
            deployedTo: address(
                new CrossChainRegistry(
                    Env.proxy.allocationManager(),
                    Env.proxy.keyRegistrar(),
                    Env.proxy.permissionController(),
                    Env.impl.pauserRegistry(),
                    Env.deployVersion()
                )
            )
        });

        deployImpl({
            name: type(ReleaseManager).name,
            deployedTo: address(new ReleaseManager(Env.proxy.permissionController(), Env.deployVersion()))
        });

        // 4. Transfer ownership of beacons to the executor multisig
        Env.beacon.eigenPod().transferOwnership(Env.executorMultisig());
        Env.beacon.strategyBase().transferOwnership(Env.executorMultisig());

        vm.stopBroadcast();
    }

    function testScript() public virtual override {
        // Deploy older contracts. We have to manually set the EOA mode so we don't revert
        _mode = OperationalMode.EOA;
        DeployGovernance._runAsEOA();
        DeployPauser._runAsEOA();
        DeployToken._runAsEOA();

        // Run the deploy proxies script
        runAsEOA();

        _validateProxyAdmins();
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateBeacons();
        _validateVersion();
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.delegationManager())) == pa, "delegationManager proxyAdmin incorrect"
        );

        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.strategyManager())) == pa, "strategyManager proxyAdmin incorrect"
        );

        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.allocationManager())) == pa, "allocationManager proxyAdmin incorrect"
        );

        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.rewardsCoordinator())) == pa, "rewardsCoordinator proxyAdmin incorrect"
        );

        assertTrue(Env._getProxyAdmin(address(Env.proxy.avsDirectory())) == pa, "avsDirectory proxyAdmin incorrect");

        assertTrue(Env._getProxyAdmin(address(Env.proxy.releaseManager())) == pa, "releaseManager proxyAdmin incorrect");

        /// permissions/
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.permissionController())) == pa,
            "permissionController proxyAdmin incorrect"
        );
        assertTrue(Env._getProxyAdmin(address(Env.proxy.keyRegistrar())) == pa, "keyRegistrar proxyAdmin incorrect");

        /// multichain/
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.crossChainRegistry())) == pa, "crossChainRegistry proxyAdmin incorrect"
        );

        /// strategies/
        assertTrue(Env._getProxyAdmin(address(Env.proxy.eigenStrategy())) == pa, "eigenStrategy proxyAdmin incorrect");

        assertTrue(Env.beacon.strategyBase().owner() == Env.executorMultisig(), "strategyBase beacon owner incorrect");

        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.strategyFactory())) == pa, "strategyFactory proxyAdmin incorrect"
        );

        /// pods/
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.eigenPodManager())) == pa, "eigenPodManager proxyAdmin incorrect"
        );

        assertTrue(Env.beacon.eigenPod().owner() == Env.executorMultisig(), "eigenPod beacon owner incorrect");
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        {
            /// core/

            AllocationManager allocationManager = Env.impl.allocationManager();
            assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
            assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "alm.pR invalid");
            assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "alm.pc invalid");
            assertTrue(allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "alm.deallocDelay invalid");
            assertTrue(
                allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
                "alm.configDelay invalid"
            );

            AVSDirectory avsDirectory = Env.impl.avsDirectory();
            assertTrue(avsDirectory.delegation() == Env.proxy.delegationManager(), "avsD.dm invalid");
            assertTrue(avsDirectory.pauserRegistry() == Env.impl.pauserRegistry(), "avsD.pR invalid");

            DelegationManager delegation = Env.impl.delegationManager();
            assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
            assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
            assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
            assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
            assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
            assertTrue(
                delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid"
            );

            RewardsCoordinator rewards = Env.impl.rewardsCoordinator();
            assertTrue(rewards.delegationManager() == Env.proxy.delegationManager(), "rc.dm invalid");
            assertTrue(rewards.strategyManager() == Env.proxy.strategyManager(), "rc.sm invalid");
            assertTrue(rewards.allocationManager() == Env.proxy.allocationManager(), "rc.alm invalid");
            assertTrue(rewards.pauserRegistry() == Env.impl.pauserRegistry(), "rc.pR invalid");
            assertTrue(rewards.permissionController() == Env.proxy.permissionController(), "rc.pc invalid");
            assertTrue(
                rewards.CALCULATION_INTERVAL_SECONDS() == Env.CALCULATION_INTERVAL_SECONDS(), "rc.calcInterval invalid"
            );
            assertTrue(rewards.MAX_REWARDS_DURATION() == Env.MAX_REWARDS_DURATION(), "rc.rewardsDuration invalid");
            assertTrue(rewards.MAX_RETROACTIVE_LENGTH() == Env.MAX_RETROACTIVE_LENGTH(), "rc.retroLength invalid");
            assertTrue(rewards.MAX_FUTURE_LENGTH() == Env.MAX_FUTURE_LENGTH(), "rc.futureLength invalid");
            assertTrue(rewards.GENESIS_REWARDS_TIMESTAMP() == Env.GENESIS_REWARDS_TIMESTAMP(), "rc.genesis invalid");

            StrategyManager strategyManager = Env.impl.strategyManager();
            assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
            assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");

            ReleaseManager releaseManager = Env.impl.releaseManager();
            assertTrue(releaseManager.permissionController() == Env.proxy.permissionController(), "rm.pc invalid");
        }

        {
            // PermissionController has no constructor

            /// permissions/
            KeyRegistrar keyRegistrar = Env.impl.keyRegistrar();
            assertTrue(keyRegistrar.permissionController() == Env.proxy.permissionController(), "kr.pc invalid");
            assertTrue(keyRegistrar.allocationManager() == Env.proxy.allocationManager(), "kr.alm invalid");
        }

        {
            /// multichain/
            CrossChainRegistry crossChainRegistry = Env.impl.crossChainRegistry();
            assertTrue(crossChainRegistry.allocationManager() == Env.proxy.allocationManager(), "ccr.alm invalid");
            assertTrue(crossChainRegistry.keyRegistrar() == Env.proxy.keyRegistrar(), "ccr.kr invalid");
            assertTrue(crossChainRegistry.permissionController() == Env.proxy.permissionController(), "ccr.pc invalid");
            assertTrue(crossChainRegistry.pauserRegistry() == Env.impl.pauserRegistry(), "ccr.pR invalid");
        }

        {
            /// pods/
            EigenPod eigenPod = Env.impl.eigenPod();
            assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
            assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");

            EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
            assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
            assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
            assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
            assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");
        }

        {
            /// strategies/
            EigenStrategy eigenStrategy = Env.impl.eigenStrategy();
            assertTrue(eigenStrategy.strategyManager() == Env.proxy.strategyManager(), "eigStrat.sm invalid");
            assertTrue(eigenStrategy.pauserRegistry() == Env.impl.pauserRegistry(), "eigStrat.pR invalid");

            StrategyBase strategyBase = Env.impl.strategyBase();
            assertTrue(strategyBase.strategyManager() == Env.proxy.strategyManager(), "stratBase.sm invalid");
            assertTrue(strategyBase.pauserRegistry() == Env.impl.pauserRegistry(), "stratBase.pR invalid");

            StrategyFactory strategyFactory = Env.impl.strategyFactory();
            assertTrue(strategyFactory.strategyManager() == Env.proxy.strategyManager(), "sFact.sm invalid");
            assertTrue(strategyFactory.pauserRegistry() == Env.impl.pauserRegistry(), "sFact.pR invalid");
        }
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// permissions/
        // PermissionController is initializable, but does not expose the `initialize` method

        {
            /// core/

            AllocationManager allocationManager = Env.impl.allocationManager();
            vm.expectRevert(errInit);
            allocationManager.initialize(0);

            AVSDirectory avsDirectory = Env.impl.avsDirectory();
            vm.expectRevert(errInit);
            avsDirectory.initialize(address(0), 0);

            DelegationManager delegation = Env.impl.delegationManager();
            vm.expectRevert(errInit);
            delegation.initialize(0);

            RewardsCoordinator rewards = Env.impl.rewardsCoordinator();
            vm.expectRevert(errInit);
            rewards.initialize(address(0), 0, address(0), 0, 0);

            StrategyManager strategyManager = Env.impl.strategyManager();
            vm.expectRevert(errInit);
            strategyManager.initialize(address(0), address(0), 0);

            // ReleaseManager is not initializable
        }

        // KeyRegistrar and PermissionController are not initializable

        {
            /// multichain/
            CrossChainRegistry crossChainRegistry = Env.impl.crossChainRegistry();
            vm.expectRevert(errInit);
            crossChainRegistry.initialize(address(0), 0, 0);
        }

        {
            /// pods/
            EigenPod eigenPod = Env.impl.eigenPod();
            vm.expectRevert(errInit);
            eigenPod.initialize(address(0));

            EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
            vm.expectRevert(errInit);
            eigenPodManager.initialize(address(0), 0);
        }

        {
            /// strategies/
            EigenStrategy eigenStrategy = Env.impl.eigenStrategy();
            vm.expectRevert(errInit);
            eigenStrategy.initialize(IEigen(address(0)), IBackingEigen(address(0)));

            StrategyBase strategyBase = Env.impl.strategyBase();
            vm.expectRevert(errInit);
            strategyBase.initialize(IERC20(address(0)));

            StrategyFactory strategyFactory = Env.impl.strategyFactory();
            vm.expectRevert(errInit);
            strategyFactory.initialize(address(0), 0, UpgradeableBeacon(address(0)));
        }
    }

    /// @dev Validate that the beacons are set to the correct implementation contracts
    function _validateBeacons() internal view {
        assertEq(
            address(Env.beacon.eigenPod().implementation()),
            address(Env.impl.eigenPod()),
            "eigenPod beacon implementation incorrect"
        );
        assertEq(
            address(Env.beacon.strategyBase().implementation()),
            address(Env.impl.strategyBase()),
            "strategyBase beacon implementation incorrect"
        );
    }

    function _validateVersion() internal view {
        // On future upgrades, just tick the major/minor/patch to validate
        string memory expected = Env.deployVersion();

        {
            /// core/
            assertEq(Env.impl.allocationManager().version(), expected, "allocationManager version mismatch");
            assertEq(Env.impl.avsDirectory().version(), expected, "avsDirectory version mismatch");
            assertEq(Env.impl.delegationManager().version(), expected, "delegationManager version mismatch");
            assertEq(Env.impl.rewardsCoordinator().version(), expected, "rewardsCoordinator version mismatch");
            assertEq(Env.impl.strategyManager().version(), expected, "strategyManager version mismatch");
            assertEq(Env.impl.releaseManager().version(), expected, "releaseManager version mismatch");
        }

        {
            /// permissions/
            assertEq(Env.impl.permissionController().version(), expected, "permissionController version mismatch");
            assertEq(Env.impl.keyRegistrar().version(), expected, "keyRegistrar version mismatch");
        }

        {
            /// multichain/
            assertEq(Env.impl.crossChainRegistry().version(), expected, "crossChainRegistry version mismatch");
        }

        {
            /// pods/
            assertEq(Env.impl.eigenPod().version(), expected, "eigenPod version mismatch");
            assertEq(Env.impl.eigenPodManager().version(), expected, "eigenPodManager version mismatch");
        }

        {
            /// strategies/
            assertEq(Env.impl.eigenStrategy().version(), expected, "eigenStrategy version mismatch");
            assertEq(Env.impl.strategyBase().version(), expected, "strategyBase version mismatch");
            assertEq(Env.impl.strategyFactory().version(), expected, "strategyFactory version mismatch");
        }
    }
}
