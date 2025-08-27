// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueSourceChain} from "./2-queueSourceChain.s.sol";
import {Encode} from "zeus-templates/utils/Encode.sol";
import {console} from "forge-std/console.sol";

/// @notice Executes the upgrade for v1.7.0 multichain testnet final
contract ExecuteSourceChain is QueueSourceChain {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal override prank(Env.protocolCouncilMultisig()) {
        if (!Env.isSourceChain()) {
            return;
        }

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
        if (!Env.isSourceChain()) {
            return;
        }

        // 1 - Deploy. The contracts have been deployed in the deploy upgrade script
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

        /// 2 - Queue. Check that the operation IS ready
        QueueSourceChain._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready for execution.");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete.");

        // 3 - warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertTrue(timelock.isOperationReady(txHash), "Transaction should be executable.");

        // 4 - execute
        execute();
        assertTrue(timelock.isOperationDone(txHash), "v1.7.0 multichain testnet final txn should be complete.");

        // 5 - Validate
        _validateNewImplAddresses({areMatching: true});
        _validateProxyAdmins();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Mirrors the checks done in 1-deploySourceChain, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal view {
        KeyRegistrar keyRegistrar = Env.proxy.keyRegistrar();
        assertEq(keyRegistrar.version(), Env.deployVersion(), "keyRegistrar version mismatch");
        assertTrue(
            keyRegistrar.permissionController() == Env.proxy.permissionController(),
            "keyRegistrar permissionController mismatch"
        );
        assertTrue(
            keyRegistrar.allocationManager() == Env.proxy.allocationManager(), "keyRegistrar allocationManager mismatch"
        );

        CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
        assertEq(crossChainRegistry.version(), Env.deployVersion(), "crossChainRegistry version mismatch");
        assertTrue(
            crossChainRegistry.allocationManager() == Env.proxy.allocationManager(),
            "crossChainRegistry allocationManager mismatch"
        );
        assertTrue(
            crossChainRegistry.keyRegistrar() == Env.proxy.keyRegistrar(), "crossChainRegistry keyRegistrar mismatch"
        );
        assertTrue(
            crossChainRegistry.permissionController() == Env.proxy.permissionController(),
            "crossChainRegistry permissionController mismatch"
        );
        assertTrue(
            crossChainRegistry.pauserRegistry() == Env.impl.pauserRegistry(),
            "crossChainRegistry pauserRegistry mismatch"
        );

        ReleaseManager releaseManager = Env.proxy.releaseManager();
        assertEq(releaseManager.version(), Env.deployVersion(), "releaseManager version mismatch");
        assertTrue(
            releaseManager.permissionController() == Env.proxy.permissionController(),
            "releaseManager permissionController mismatch"
        );
    }

    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// CrossChainRegistry
        CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
        vm.expectRevert(errInit);
        crossChainRegistry.initialize(address(0), 1 days, 0);

        // ReleaseManager and KeyRegistrar don't have initialize functions
    }
}
