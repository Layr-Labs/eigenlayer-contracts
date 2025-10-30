// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {DeployToken} from "./3-deployToken.s.sol";
import {DeployPauser} from "./2-deployPauser.s.sol";
import {DeployGovernance} from "./1-deployGovernance.s.sol";
import {DeployCore} from "./4-deployCore.s.sol";
import {Queue} from "./5-queueUpgrade.s.sol";
import "../Env.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";

contract Execute is Queue {
    using Env for *;

    function _runAsMultisig() internal virtual override prank(Env.protocolCouncilMultisig()) {
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

    function testScript() public virtual override {
        // 0- Run all the previous steps of the upgrade
        _mode = OperationalMode.EOA;
        DeployGovernance._runAsEOA();
        DeployPauser._runAsEOA();
        DeployToken._runAsEOA();
        DeployCore._runAsEOA();

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
        Queue._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready for execution.");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete.");

        // 2- warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // 3- execute
        execute();
        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");

        // 4- validate the upgrade was successful
        _validateNewImplAddresses({areMatching: true});
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    function _validateNewImplAddresses(
        bool areMatching
    ) internal view {
        /// core/ -- can't check AllocationManager as it didn't exist before this deploy

        function (bool, string memory) internal pure assertion = areMatching ? _assertTrue : _assertFalse;

        assertion(
            Env._getProxyImpl(address(Env.proxy.permissionController())) == address(Env.impl.permissionController()),
            "permissionController impl failed"
        );

        assertion(
            Env._getProxyImpl(address(Env.proxy.delegationManager())) == address(Env.impl.delegationManager()),
            "delegationManager impl failed"
        );

        assertion(
            Env._getProxyImpl(address(Env.proxy.strategyManager())) == address(Env.impl.strategyManager()),
            "strategyManager impl failed"
        );

        assertion(
            Env._getProxyImpl(address(Env.proxy.allocationManager())) == address(Env.impl.allocationManager()),
            "allocationManager impl failed"
        );

        assertion(
            Env._getProxyImpl(address(Env.proxy.rewardsCoordinator())) == address(Env.impl.rewardsCoordinator()),
            "rewardsCoordinator impl failed"
        );

        assertion(
            Env._getProxyImpl(address(Env.proxy.avsDirectory())) == address(Env.impl.avsDirectory()),
            "avsDirectory impl failed"
        );

        assertion(
            Env._getProxyImpl(address(Env.proxy.eigenPodManager())) == address(Env.impl.eigenPodManager()),
            "eigenPodManager impl failed"
        );

        assertion(
            Env._getProxyImpl(address(Env.proxy.strategyFactory())) == address(Env.impl.strategyFactory()),
            "strategyFactory impl failed"
        );

        assertion(
            Env._getProxyImpl(address(Env.proxy.releaseManager())) == address(Env.impl.releaseManager()),
            "releaseManager impl failed"
        );

        /// permissions/
        assertion(
            Env._getProxyImpl(address(Env.proxy.keyRegistrar())) == address(Env.impl.keyRegistrar()),
            "keyRegistrar impl failed"
        );

        /// strategies/
        assertion(
            Env._getProxyImpl(address(Env.proxy.eigenStrategy())) == address(Env.impl.eigenStrategy()),
            "eigenStrategy impl failed"
        );
    }

    /// @dev Mirrors the checks done in 4-deployCore.s.sol, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal view {
        {
            /// core/

            // Permission controller has no constructor

            DelegationManager delegation = Env.proxy.delegationManager();
            assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
            assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
            assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
            assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
            assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
            assertTrue(
                delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid"
            );

            StrategyManager strategyManager = Env.proxy.strategyManager();
            assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
            assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");

            AllocationManager allocationManager = Env.proxy.allocationManager();
            assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
            assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "alm.pR invalid");
            assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "alm.pc invalid");
            assertTrue(allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "alm.deallocDelay invalid");
            assertTrue(
                allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
                "alm.configDelay invalid"
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

            AVSDirectory avsDirectory = Env.proxy.avsDirectory();
            assertTrue(avsDirectory.delegation() == Env.proxy.delegationManager(), "avsD.dm invalid");
            assertTrue(avsDirectory.pauserRegistry() == Env.impl.pauserRegistry(), "avsD.pR invalid");

            ReleaseManager releaseManager = Env.proxy.releaseManager();
            assertTrue(releaseManager.permissionController() == Env.proxy.permissionController(), "rm.pc invalid");
        }

        {
            /// permissions/
            KeyRegistrar keyRegistrar = Env.proxy.keyRegistrar();
            assertTrue(keyRegistrar.permissionController() == Env.proxy.permissionController(), "kr.pc invalid");
            assertTrue(keyRegistrar.allocationManager() == Env.proxy.allocationManager(), "kr.alm invalid");
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

            StrategyFactory strategyFactory = Env.proxy.strategyFactory();
            assertTrue(strategyFactory.strategyManager() == Env.proxy.strategyManager(), "sFact.sm invalid");
            assertTrue(strategyFactory.pauserRegistry() == Env.impl.pauserRegistry(), "sFact.pR invalid");
        }
    }

    /// @dev Call initialize on all proxies to ensure they are initialized
    /// Additionally, validate initialization variables
    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        {
            /// core/
            DelegationManager delegation = Env.proxy.delegationManager();
            vm.expectRevert(errInit);
            delegation.initialize(0);
            assertTrue(delegation.paused() == 0, "dm.paused invalid");

            StrategyManager strategyManager = Env.proxy.strategyManager();
            vm.expectRevert(errInit);
            strategyManager.initialize(address(0), address(0), 0);
            assertTrue(strategyManager.owner() == Env.executorMultisig(), "sm.owner invalid");
            assertTrue(strategyManager.paused() == 0, "sm.paused invalid");
            assertTrue(
                strategyManager.strategyWhitelister() == address(Env.proxy.strategyFactory()), "sm.whitelister invalid"
            );

            AllocationManager allocationManager = Env.proxy.allocationManager();
            vm.expectRevert(errInit);
            allocationManager.initialize(0);
            assertTrue(allocationManager.paused() == 0, "alm.paused invalid");

            RewardsCoordinator rewards = Env.proxy.rewardsCoordinator();
            vm.expectRevert(errInit);
            rewards.initialize(address(0), 0, address(0), 0, 0);
            assertTrue(rewards.owner() == Env.opsMultisig(), "rc.owner invalid");
            assertTrue(rewards.paused() == Env.REWARDS_PAUSE_STATUS(), "rc.paused invalid");
            assertTrue(rewards.rewardsUpdater() == Env.REWARDS_UPDATER(), "rc.updater invalid");
            assertTrue(rewards.activationDelay() == Env.ACTIVATION_DELAY(), "rc.activationDelay invalid");
            assertTrue(rewards.defaultOperatorSplitBips() == Env.DEFAULT_SPLIT_BIPS(), "rc.splitBips invalid");

            AVSDirectory avsDirectory = Env.proxy.avsDirectory();
            vm.expectRevert(errInit);
            avsDirectory.initialize(address(0), 0);
            assertTrue(avsDirectory.owner() == Env.executorMultisig(), "avsD.owner invalid");
            assertTrue(avsDirectory.paused() == 0, "avsD.paused invalid");

            // ReleaseManager is not initializable
        }

        // KeyRegistrar and PermissionController are not initializable

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
            assertEq(address(eigenStrategy.EIGEN()), address(Env.proxy.eigen()), "eigenStrat.EIGEN invalid");
            assertEq(
                address(eigenStrategy.underlyingToken()), address(Env.proxy.beigen()), "eigenStrat.underlying invalid"
            );

            // StrategyBase proxies are initialized when deployed by factory
            StrategyFactory strategyFactory = Env.proxy.strategyFactory();
            vm.expectRevert(errInit);
            strategyFactory.initialize(address(0), 0, UpgradeableBeacon(address(0)));
            assertTrue(strategyFactory.owner() == Env.opsMultisig(), "sFact.owner invalid");
            assertTrue(strategyFactory.paused() == 0, "sFact.paused invalid");
            assertTrue(strategyFactory.strategyBeacon() == Env.beacon.strategyBase(), "sFact.beacon invalid");
        }
    }

    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }
}
