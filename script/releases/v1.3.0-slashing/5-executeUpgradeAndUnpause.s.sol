// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueAndUnpause} from "./2-queueUpgradeAndUnpause.s.sol";
import {Pause} from "./3-pause.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is QueueAndUnpause, Pause {
    using Env for *;

    function _runAsMultisig() internal override(Pause, QueueAndUnpause) prank(Env.protocolCouncilMultisig()) {
        bytes memory calldata_to_executor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
    }

    function testScript() public virtual override(Pause, QueueAndUnpause) {
        runAsEOA();

        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });

        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        // 1- run queueing logic
        QueueAndUnpause._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready for execution.");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete.");

        // 2- run pausing logic
        Pause._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), "EPM is not paused!");

        // 2- warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // 3- execute
        execute();

        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");

        _validateNewImplAddresses({areMatching: true});
        _validateProxyAdmins();
        _validateProxyConstructors();
        _validateProxiesInitialized();
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

            uint256 count = Env.instance.strategyBaseTVLLimits_Count();
            for (uint256 i = 0; i < count; i++) {
                StrategyBaseTVLLimits strategy = Env.instance.strategyBaseTVLLimits(i);

                assertTrue(strategy.strategyManager() == Env.proxy.strategyManager(), "sFact.sm invalid");
                assertTrue(strategy.pauserRegistry() == Env.impl.pauserRegistry(), "sFact.pR invalid");
            }

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
            assertTrue(rewards.paused() == Env.REWARDS_PAUSE_STATUS(), "rc.paused invalid");
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
            assertTrue(eigenPodManager.paused() == 0, "epm.paused invalid");
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

            uint256 count = Env.instance.strategyBaseTVLLimits_Count();
            for (uint256 i = 0; i < count; i++) {
                StrategyBaseTVLLimits strategy = Env.instance.strategyBaseTVLLimits(i);

                emit log_named_address("strat", address(strategy));

                vm.expectRevert(errInit);
                strategy.initialize(0, 0, IERC20(address(0)));
                assertTrue(strategy.maxPerDeposit() == type(uint256).max, "stratTVLLim.maxPerDeposit invalid");
                assertTrue(strategy.maxTotalDeposits() == type(uint256).max, "stratTVLLim.maxPerDeposit invalid");
            }

            StrategyFactory strategyFactory = Env.proxy.strategyFactory();
            vm.expectRevert(errInit);
            strategyFactory.initialize(address(0), 0, UpgradeableBeacon(address(0)));
            assertTrue(strategyFactory.owner() == Env.opsMultisig(), "sFact.owner invalid");
            assertTrue(strategyFactory.paused() == 0, "sFact.paused invalid");
            assertTrue(strategyFactory.strategyBeacon() == Env.beacon.strategyBase(), "sFact.beacon invalid");
        }
    }
}
