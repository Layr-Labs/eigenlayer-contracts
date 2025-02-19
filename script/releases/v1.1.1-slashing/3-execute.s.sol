// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {Queue} from "./2-multisig.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is Queue {
    using Env for *;

    function _runAsMultisig() internal override(Queue) prank(Env.protocolCouncilMultisig()) {
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

    function testScript() public virtual override(Queue) {
        // 0. Deploy Impls
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

        // 1. Queue Upgrade
        Queue._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        // 2. Warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // 3- execute
        execute();

        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");

        // 4. Validate
        _validateNewImplAddresses(true);
        _validateProxyConstructors();
        _validateProxyInitialized();
    }

    function _validateProxyConstructors() internal view {
        AllocationManager allocationManager = Env.proxy.allocationManager();
        assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
        assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "alm.pR invalid");
        assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "alm.pc invalid");
        assertTrue(
            allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "alm.deallocationDelay invalid"
        );
        assertTrue(
            allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
            "alm.allocationConfigurationDelay invalid"
        );
    }

    /// @dev Call initialize on all deployed proxies to ensure initializers are disabled
    function _validateProxyInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        AllocationManager allocationManager = Env.proxy.allocationManager();
        vm.expectRevert(errInit);
        allocationManager.initialize(address(0), 0);
    }
}
