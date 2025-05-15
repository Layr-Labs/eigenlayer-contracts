// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * Purpose: Validate a deployment on a given environment:
 * 1. Check pointers in contracts
 * 2. Validate owners
 * 3. Validate version
 */
contract ValidateDeployment is EOADeployer {
    using Env for *;

    // No-op here, all we are doing is testing the environment
    function _runAsEOA() internal override {}

    function testScript() public virtual {
        _validateNewImplAddresses({areMatching: true});
        _validateProxyAdmins();
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Validate that the `Env.impl` addresses match what the proxy admin reports as the current implementation address.
    function _validateNewImplAddresses(
        bool areMatching
    ) internal view {
        /// core/

        function (bool, string memory) internal pure assertion = areMatching ? _assertTrue : _assertFalse;

        assertion(
            _getProxyImpl(address(Env.proxy.avsDirectory())) == address(Env.impl.avsDirectory()),
            "avsDirectory impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.delegationManager())) == address(Env.impl.delegationManager()),
            "delegationManager impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.rewardsCoordinator())) == address(Env.impl.rewardsCoordinator()),
            "rewardsCoordinator impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.strategyManager())) == address(Env.impl.strategyManager()),
            "strategyManager impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.allocationManager())) == address(Env.impl.allocationManager()),
            "allocationManager impl failed"
        );

        /// permissions/
        assertion(
            _getProxyImpl(address(Env.proxy.permissionController())) == address(Env.impl.permissionController()),
            "permissionController impl failed"
        );

        /// pods/

        assertion(Env.beacon.eigenPod().implementation() == address(Env.impl.eigenPod()), "eigenPod impl failed");

        assertion(
            _getProxyImpl(address(Env.proxy.eigenPodManager())) == address(Env.impl.eigenPodManager()),
            "eigenPodManager impl failed"
        );

        /// strategies/

        assertion(
            _getProxyImpl(address(Env.proxy.eigenStrategy())) == address(Env.impl.eigenStrategy()),
            "eigenStrategy impl failed"
        );

        assertion(
            Env.beacon.strategyBase().implementation() == address(Env.impl.strategyBase()), "strategyBase impl failed"
        );

        // uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        // for (uint256 i = 0; i < count; i++) {
        //     assertion(
        //         _getProxyImpl(address(Env.instance.strategyBaseTVLLimits(i)))
        //             == address(Env.impl.strategyBaseTVLLimits()),
        //         "strategyBaseTVLLimits impl failed"
        //     );
        // }

        assertion(
            _getProxyImpl(address(Env.proxy.strategyFactory())) == address(Env.impl.strategyFactory()),
            "strategyFactory impl failed"
        );
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            _getProxyAdmin(address(Env.proxy.allocationManager())) == pa, "allocationManager proxyAdmin incorrect"
        );

        assertTrue(_getProxyAdmin(address(Env.proxy.avsDirectory())) == pa, "avsDirectory proxyAdmin incorrect");

        assertTrue(
            _getProxyAdmin(address(Env.proxy.delegationManager())) == pa, "delegationManager proxyAdmin incorrect"
        );

        assertTrue(
            _getProxyAdmin(address(Env.proxy.rewardsCoordinator())) == pa, "rewardsCoordinator proxyAdmin incorrect"
        );

        assertTrue(_getProxyAdmin(address(Env.proxy.strategyManager())) == pa, "strategyManager proxyAdmin incorrect");

        assertTrue(
            _getProxyAdmin(address(Env.proxy.allocationManager())) == pa, "allocationManager proxyAdmin incorrect"
        );

        /// permissions
        assertTrue(
            _getProxyAdmin(address(Env.proxy.permissionController())) == pa, "permissionController proxyAdmin incorrect"
        );

        /// pods/

        assertTrue(Env.beacon.eigenPod().owner() == Env.executorMultisig(), "eigenPod beacon owner incorrect");

        assertTrue(_getProxyAdmin(address(Env.proxy.eigenPodManager())) == pa, "eigenPodManager proxyAdmin incorrect");

        /// strategies/

        assertTrue(_getProxyAdmin(address(Env.proxy.eigenStrategy())) == pa, "eigenStrategy proxyAdmin incorrect");

        assertTrue(Env.beacon.strategyBase().owner() == Env.executorMultisig(), "strategyBase beacon owner incorrect");

        // uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        // for (uint256 i = 0; i < count; i++) {
        //     assertTrue(
        //         _getProxyAdmin(address(Env.instance.strategyBaseTVLLimits(i))) == pa,
        //         "strategyBaseTVLLimits proxyAdmin incorrect"
        //     );
        // }

        assertTrue(_getProxyAdmin(address(Env.proxy.strategyFactory())) == pa, "strategyFactory proxyAdmin incorrect");
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
            assertTrue(
                strategyBaseTVLLimits.strategyManager() == Env.proxy.strategyManager(), "stratBaseTVL.sm invalid"
            );
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

    /// @dev Mirrors the checks done in 1-deployContracts, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal view {
        {
            /// permissions/

            // exception: PauserRegistry doesn't have a proxy!
            PauserRegistry registry = Env.impl.pauserRegistry();
            assertTrue(registry.isPauser(Env.pauserMultisig()), "pauser multisig should be pauser");
            assertTrue(registry.isPauser(Env.opsMultisig()), "ops multisig should be pauser");
            assertTrue(registry.isPauser(Env.executorMultisig()), "executor multisig should be pauser");
            assertTrue(registry.unpauser() == Env.executorMultisig(), "executor multisig should be unpauser");

            /// PermissionController has no initial storage
        }

        {
            /// core/

            AllocationManager allocationManager = Env.proxy.allocationManager();
            assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
            assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "alm.pR invalid");
            assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "alm.pc invalid");
            assertTrue(allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "alm.deallocDelay invalid");
            assertTrue(
                allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
                "alm.configDelay invalid"
            );

            AVSDirectory avsDirectory = Env.proxy.avsDirectory();
            assertTrue(avsDirectory.delegation() == Env.proxy.delegationManager(), "avsD.dm invalid");
            assertTrue(avsDirectory.pauserRegistry() == Env.impl.pauserRegistry(), "avsD.pR invalid");

            DelegationManager delegation = Env.proxy.delegationManager();
            assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
            assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
            assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
            assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
            assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
            assertTrue(
                delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid"
            );

            RewardsCoordinator rewards = Env.proxy.rewardsCoordinator();
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

            StrategyManager strategyManager = Env.proxy.strategyManager();
            assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
            assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");
        }

        {
            /// pods/
            UpgradeableBeacon eigenPodBeacon = Env.beacon.eigenPod();
            assertTrue(eigenPodBeacon.implementation() == address(Env.impl.eigenPod()), "eigenPodBeacon.impl invalid");

            EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
            assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
            assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
            assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
            assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");
        }

        {
            /// strategies/
            EigenStrategy eigenStrategy = Env.proxy.eigenStrategy();
            assertTrue(eigenStrategy.strategyManager() == Env.proxy.strategyManager(), "eigStrat.sm invalid");
            assertTrue(eigenStrategy.pauserRegistry() == Env.impl.pauserRegistry(), "eigStrat.pR invalid");

            UpgradeableBeacon strategyBeacon = Env.beacon.strategyBase();
            assertTrue(
                strategyBeacon.implementation() == address(Env.impl.strategyBase()), "strategyBeacon.impl invalid"
            );

            // uint256 count = Env.instance.strategyBaseTVLLimits_Count();
            // for (uint256 i = 0; i < count; i++) {
            //     StrategyBaseTVLLimits strategy = Env.instance.strategyBaseTVLLimits(i);

            //     assertTrue(strategy.strategyManager() == Env.proxy.strategyManager(), "sFact.sm invalid");
            //     assertTrue(strategy.pauserRegistry() == Env.impl.pauserRegistry(), "sFact.pR invalid");
            // }

            StrategyFactory strategyFactory = Env.proxy.strategyFactory();
            assertTrue(strategyFactory.strategyManager() == Env.proxy.strategyManager(), "sFact.sm invalid");
            assertTrue(strategyFactory.pauserRegistry() == Env.impl.pauserRegistry(), "sFact.pR invalid");
        }
    }

    /// @dev Call initialize on all proxies to ensure they are initialized
    /// Additionally, validate initialization variables
    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// permissions/
        // PermissionController is initializable, but does not expose the `initialize` method

        {
            /// core/

            AllocationManager allocationManager = Env.proxy.allocationManager();
            vm.expectRevert(errInit);
            allocationManager.initialize(address(0), 0);
            assertTrue(allocationManager.owner() == Env.executorMultisig(), "alm.owner invalid");
            assertTrue(allocationManager.paused() == 0, "alm.paused invalid");

            AVSDirectory avsDirectory = Env.proxy.avsDirectory();
            vm.expectRevert(errInit);
            avsDirectory.initialize(address(0), 0);
            assertTrue(avsDirectory.owner() == Env.executorMultisig(), "avsD.owner invalid");
            assertTrue(avsDirectory.paused() == 0, "avsD.paused invalid");

            DelegationManager delegation = Env.proxy.delegationManager();
            vm.expectRevert(errInit);
            delegation.initialize(address(0), 0);
            assertTrue(delegation.owner() == Env.executorMultisig(), "dm.owner invalid");
            assertTrue(delegation.paused() == 0, "dm.paused invalid");

            RewardsCoordinator rewards = Env.proxy.rewardsCoordinator();
            vm.expectRevert(errInit);
            rewards.initialize(address(0), 0, address(0), 0, 0);
            assertTrue(rewards.owner() == Env.opsMultisig(), "rc.owner invalid");
            assertTrue(rewards.paused() == 2, "rc.paused invalid");
            assertTrue(rewards.rewardsUpdater() == Env.REWARDS_UPDATER(), "rc.updater invalid");
            assertTrue(rewards.activationDelay() == Env.ACTIVATION_DELAY(), "rc.activationDelay invalid");
            assertTrue(rewards.defaultOperatorSplitBips() == Env.DEFAULT_SPLIT_BIPS(), "rc.splitBips invalid");

            StrategyManager strategyManager = Env.proxy.strategyManager();
            vm.expectRevert(errInit);
            strategyManager.initialize(address(0), address(0), 0);
            assertTrue(strategyManager.owner() == Env.executorMultisig(), "sm.owner invalid");
            assertTrue(strategyManager.paused() == 0, "sm.paused invalid");
            assertTrue(
                strategyManager.strategyWhitelister() == address(Env.proxy.strategyFactory()), "sm.whitelister invalid"
            );
        }

        {
            /// pods/
            // EigenPod proxies are initialized by individual users

            EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
            vm.expectRevert(errInit);
            eigenPodManager.initialize(address(0), 0);
            assertTrue(eigenPodManager.owner() == Env.executorMultisig(), "epm.owner invalid");
            // assertTrue(eigenPodManager.paused() == 0, "epm.paused invalid");
        }

        {
            /// strategies/

            EigenStrategy eigenStrategy = Env.proxy.eigenStrategy();
            vm.expectRevert(errInit);
            eigenStrategy.initialize(IEigen(address(0)), IBackingEigen(address(0)));
            assertTrue(eigenStrategy.paused() == 0, "eigenStrat.paused invalid");
            assertTrue(eigenStrategy.EIGEN() == Env.proxy.eigen(), "eigenStrat.EIGEN invalid");
            assertTrue(eigenStrategy.underlyingToken() == Env.proxy.beigen(), "eigenStrat.underlying invalid");

            // StrategyBase proxies are initialized when deployed by factory

            // uint256 count = Env.instance.strategyBaseTVLLimits_Count();
            // for (uint256 i = 0; i < count; i++) {
            //     StrategyBaseTVLLimits strategy = Env.instance.strategyBaseTVLLimits(i);

            //     emit log_named_address("strat", address(strategy));

            //     vm.expectRevert(errInit);
            //     strategy.initialize(0, 0, IERC20(address(0)));
            //     assertTrue(strategy.maxPerDeposit() == type(uint256).max, "stratTVLLim.maxPerDeposit invalid");
            //     assertTrue(strategy.maxTotalDeposits() == type(uint256).max, "stratTVLLim.maxPerDeposit invalid");
            // }

            StrategyFactory strategyFactory = Env.proxy.strategyFactory();
            vm.expectRevert(errInit);
            strategyFactory.initialize(address(0), 0, UpgradeableBeacon(address(0)));
            assertTrue(strategyFactory.owner() == Env.opsMultisig(), "sFact.owner invalid");
            assertTrue(strategyFactory.paused() == 0, "sFact.paused invalid");
            assertTrue(strategyFactory.strategyBeacon() == Env.beacon.strategyBase(), "sFact.beacon invalid");
        }
    }

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(
        address proxy
    ) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(
        address proxy
    ) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }

    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }
}