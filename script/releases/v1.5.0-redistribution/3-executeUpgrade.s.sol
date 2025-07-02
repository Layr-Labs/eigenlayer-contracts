// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueUpgrade} from "./2-queueUpgrade.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is QueueUpgrade {
    using Env for *;

    function _runAsMultisig() internal override prank(Env.protocolCouncilMultisig()) {
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
        QueueUpgrade._runAsMultisig();
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

        _validateNewImplAddresses({areMatching: true});
        _validateProxyAdmins();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Mirrors the checks done in 1-deployContracts, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal view {
        {
            /// AllocationManager
            AllocationManager allocationManager = Env.impl.allocationManager();
            assertTrue(
                allocationManager.delegation() == Env.proxy.delegationManager(), "allocationManager.delegation invalid"
            );
            assertTrue(
                allocationManager.eigenStrategy() == Env.proxy.eigenStrategy(),
                "allocationManager.eigenStrategy invalid"
            );
            assertTrue(
                allocationManager.pauserRegistry() == Env.impl.pauserRegistry(),
                "allocationManager.pauserRegistry invalid"
            );
            assertTrue(
                allocationManager.permissionController() == Env.proxy.permissionController(),
                "allocationManager.permissionController invalid"
            );
            assertTrue(allocationManager.DEALLOCATION_DELAY() == 1 days, "allocationManager.DEALLOCATION_DELAY invalid");
            assertTrue(
                allocationManager.ALLOCATION_CONFIGURATION_DELAY() == 1 days,
                "allocationManager.ALLOCATION_CONFIGURATION_DELAY invalid"
            );
            assertEq(allocationManager.version(), Env.deployVersion(), "allocationManager.version failed");
        }

        {
            /// StrategyManager
            StrategyManager strategyManager = Env.impl.strategyManager();
            assertTrue(
                strategyManager.allocationManager() == Env.proxy.allocationManager(),
                "strategyManager.allocationManager invalid"
            );
            assertTrue(
                strategyManager.delegation() == Env.proxy.delegationManager(), "strategyManager.delegation invalid"
            );
            assertTrue(
                strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "strategyManager.pauserRegistry invalid"
            );
            assertEq(strategyManager.version(), Env.deployVersion(), "strategyManager.version failed");
        }
    }

    /// @dev Call initialize on all proxies to ensure they are initialized
    /// Additionally, validate initialization variables
    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// AllocationManager
        AllocationManager allocationManager = Env.impl.allocationManager();
        vm.expectRevert(errInit);
        allocationManager.initialize(0);

        /// StrategyManager
        StrategyManager strategyManager = Env.impl.strategyManager();
        vm.expectRevert(errInit);
        strategyManager.initialize(address(0), address(0), 0);
    }
}
