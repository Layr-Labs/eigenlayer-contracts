// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

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

        deployImpl({
            name: type(PermissionController).name,
            deployedTo: address(new PermissionController())
        });

        deployProxy({
            name: type(PermissionController).name,
            deployedTo: address(new TransparentUpgradeableProxy({
                _logic: address(Env.impl.permissionController()),
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
                _data: abi.encodeCall(
                    AllocationManager.initialize,
                    (
                        Env.executorMultisig(), // initialOwner
                        0                       // initialPausedStatus
                    )
                )
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
                _eigenPodBeacon: Env.beacon.eigenPod(),
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

    function testScript() public virtual {       
        _runAsEOA();

        _validateNewImplAddresses({ areMatching: false });
        _validateProxyAdmins();
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateStrategiesAreWhitelisted();
    }

    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    function _validateNewImplAddresses(bool areMatching) internal view {
        /// core/ -- can't check AllocationManager as it didn't exist before this deploy

        function (address, address, string memory) internal pure assertion =
            areMatching ? _assertMatch : _assertNotMatch;

        assertion(
            _getProxyImpl(address(Env.proxy.avsDirectory())),
            address(Env.impl.avsDirectory()),
            "avsDirectory impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.delegationManager())),
            address(Env.impl.delegationManager()),
            "delegationManager impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.rewardsCoordinator())),
            address(Env.impl.rewardsCoordinator()),
            "rewardsCoordinator impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.strategyManager())),
            address(Env.impl.strategyManager()),
            "strategyManager impl failed"
        );
        
        /// permissions/ -- can't check these because PauserRegistry has no proxy, and
        /// PermissionController proxy didn't exist before this deploy

        /// pods/

        assertion(
            Env.beacon.eigenPod().implementation(),
            address(Env.impl.eigenPod()),
            "eigenPod impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.eigenPodManager())),
            address(Env.impl.eigenPodManager()),
            "eigenPodManager impl failed"
        );

        /// strategies/

        assertion(
            _getProxyImpl(address(Env.proxy.eigenStrategy())),
            address(Env.impl.eigenStrategy()),
            "eigenStrategy impl failed"
        );

        assertion(
            Env.beacon.strategyBase().implementation(),
            address(Env.impl.strategyBase()),
            "strategyBase impl failed"
        );

        uint count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint i = 0; i < count; i++) {
            assertion(
                _getProxyImpl(address(Env.instance.strategyBaseTVLLimits(i))),
                address(Env.impl.strategyBaseTVLLimits()),
                "strategyBaseTVLLimits impl failed"
            );
        }

        assertion(
            _getProxyImpl(address(Env.proxy.strategyFactory())),
            address(Env.impl.strategyFactory()),
            "strategyFactory impl failed"
        );
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            _getProxyAdmin(address(Env.proxy.allocationManager())) == pa,
            "allocationManager proxyAdmin incorrect"
        );

        assertTrue(
            _getProxyAdmin(address(Env.proxy.avsDirectory())) == pa,
            "avsDirectory proxyAdmin incorrect"
        );

        assertTrue(
            _getProxyAdmin(address(Env.proxy.delegationManager())) == pa,
            "delegationManager proxyAdmin incorrect"
        );

        assertTrue(
            _getProxyAdmin(address(Env.proxy.rewardsCoordinator())) == pa,
            "rewardsCoordinator proxyAdmin incorrect"
        );

        assertTrue(
            _getProxyAdmin(address(Env.proxy.strategyManager())) == pa,
            "strategyManager proxyAdmin incorrect"
        );
        
        /// permissions/ -- can't check these because PauserRegistry has no proxy, and
        /// PermissionController proxy didn't exist before this deploy

        /// pods/

        assertTrue(
            Env.beacon.eigenPod().owner() == Env.executorMultisig(),
            "eigenPod beacon owner incorrect"
        );

        assertTrue(
            _getProxyAdmin(address(Env.proxy.eigenPodManager())) == pa,
            "eigenPodManager proxyAdmin incorrect"
        );

        /// strategies/

        assertTrue(
            _getProxyAdmin(address(Env.proxy.eigenStrategy())) == pa,
            "eigenStrategy proxyAdmin incorrect"
        );

        assertTrue(
            Env.beacon.strategyBase().owner() == Env.executorMultisig(),
            "strategyBase beacon owner incorrect"
        );

        uint count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint i = 0; i < count; i++) {
            assertTrue(
                _getProxyAdmin(address(Env.instance.strategyBaseTVLLimits(i))) == pa,
                "strategyBaseTVLLimits proxyAdmin incorrect"
            );
        }

        assertTrue(
            _getProxyAdmin(address(Env.proxy.strategyFactory())) == pa,
            "strategyFactory proxyAdmin incorrect"
        );
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        {
            /// permissions/

            PauserRegistry registry = Env.impl.pauserRegistry();
            assertTrue(registry.isPauser(Env.pauserMultisig()), "pauser multisig should be pauser");
            assertTrue(registry.isPauser(Env.opsMultisig()), "ops multisig should be pauser");
            assertTrue(registry.isPauser(Env.executorMultisig()), "executor multisig should be pauser");
            assertTrue(registry.unpauser() == Env.executorMultisig(), "executor multisig should be unpauser");

            /// PermissionController has no initial storage
        }
        
        {
            /// core/

            AllocationManager allocationManager = Env.impl.allocationManager();
            assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
            assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "alm.pR invalid");
            assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "alm.pc invalid");
            assertTrue(allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "alm.deallocDelay invalid");
            assertTrue(allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(), "alm.configDelay invalid");
            
            AVSDirectory avsDirectory = Env.impl.avsDirectory();
            assertTrue(avsDirectory.delegation() == Env.proxy.delegationManager(), "avsD.dm invalid");
            assertTrue(avsDirectory.pauserRegistry() == Env.impl.pauserRegistry(), "avsD.pR invalid");

            DelegationManager delegation = Env.impl.delegationManager();
            assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
            assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
            assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
            assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
            assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
            assertTrue(delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid");

            RewardsCoordinator rewards = Env.impl.rewardsCoordinator();
            assertTrue(rewards.delegationManager() == Env.proxy.delegationManager(), "rc.dm invalid");
            assertTrue(rewards.strategyManager() == Env.proxy.strategyManager(), "rc.sm invalid");
            assertTrue(rewards.allocationManager() == Env.proxy.allocationManager(), "rc.alm invalid");
            assertTrue(rewards.pauserRegistry() == Env.impl.pauserRegistry(), "rc.pR invalid");
            assertTrue(rewards.permissionController() == Env.proxy.permissionController(), "rc.pc invalid");
            assertTrue(rewards.CALCULATION_INTERVAL_SECONDS() == Env.CALCULATION_INTERVAL_SECONDS(), "rc.calcInterval invalid");
            assertTrue(rewards.MAX_REWARDS_DURATION() == Env.MAX_REWARDS_DURATION(), "rc.rewardsDuration invalid");
            assertTrue(rewards.MAX_RETROACTIVE_LENGTH() == Env.MAX_RETROACTIVE_LENGTH(), "rc.retroLength invalid");
            assertTrue(rewards.MAX_FUTURE_LENGTH() == Env.MAX_FUTURE_LENGTH(), "rc.futureLength invalid");
            assertTrue(rewards.GENESIS_REWARDS_TIMESTAMP() == Env.GENESIS_REWARDS_TIMESTAMP(), "rc.genesis invalid");
            
            StrategyManager strategyManager = Env.impl.strategyManager();
            assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
            assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");
        }

        {
            /// pods/
            EigenPod eigenPod = Env.impl.eigenPod();
            assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
            assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");
            assertTrue(eigenPod.GENESIS_TIME() == Env.EIGENPOD_GENESIS_TIME(), "ep.genesis invalid");

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

            StrategyBaseTVLLimits strategyBaseTVLLimits = Env.impl.strategyBaseTVLLimits();
            assertTrue(strategyBaseTVLLimits.strategyManager() == Env.proxy.strategyManager(), "stratBaseTVL.sm invalid");
            assertTrue(strategyBaseTVLLimits.pauserRegistry() == Env.impl.pauserRegistry(), "stratBaseTVL.pR invalid");

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
            allocationManager.initialize(address(0), 0);
            
            AVSDirectory avsDirectory = Env.impl.avsDirectory();
            vm.expectRevert(errInit);
            avsDirectory.initialize(address(0), 0);

            DelegationManager delegation = Env.impl.delegationManager();
            vm.expectRevert(errInit);
            delegation.initialize(address(0), 0);

            RewardsCoordinator rewards = Env.impl.rewardsCoordinator();
            vm.expectRevert(errInit);
            rewards.initialize(address(0), 0, address(0), 0, 0);

            StrategyManager strategyManager = Env.impl.strategyManager();
            vm.expectRevert(errInit);
            strategyManager.initialize(address(0), address(0), 0);
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

            StrategyBaseTVLLimits strategyBaseTVLLimits = Env.impl.strategyBaseTVLLimits();
            vm.expectRevert(errInit);
            strategyBaseTVLLimits.initialize(0, 0, IERC20(address(0)));

            StrategyFactory strategyFactory = Env.impl.strategyFactory();
            vm.expectRevert(errInit);
            strategyFactory.initialize(address(0), 0, UpgradeableBeacon(address(0)));
        }
    }

    /// @dev Iterate over StrategyBaseTVLLimits instances and validate that each is
    /// whitelisted for deposit
    function _validateStrategiesAreWhitelisted() internal view {
        uint count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint i = 0; i < count; i++) {
            StrategyBaseTVLLimits strategy = Env.instance.strategyBaseTVLLimits(i);
            
            // emit log_named_uint("strategy", i);
            // IERC20Metadata underlying = IERC20Metadata(address(strategy.underlyingToken()));
            // emit log_named_string("- name", underlying.name());
            // emit log_named_string("- symbol", underlying.symbol());
            // emit log_named_uint("- totalShares", strategy.totalShares());

            bool isWhitelisted = Env.proxy.strategyManager().strategyIsWhitelistedForDeposit(strategy);
            // emit log_named_string("- is whitelisted", isWhitelisted ? "true" : "false");
            assertTrue(isWhitelisted, "not whitelisted!!");
        }
    }

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }

    function _assertMatch(address a, address b, string memory err) private pure {
        assertEq(a, b, err);
    }

    function _assertNotMatch(address a, address b, string memory err) private pure {
        assertNotEq(a, b, err);
    }
}
