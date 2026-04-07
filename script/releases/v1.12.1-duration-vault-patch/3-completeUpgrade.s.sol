// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {QueueUpgrade} from "./2-queueUpgrade.s.sol";
import "../Env.sol";
import "../TestUtils.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/// Purpose: execute the Sepolia-only DurationVaultStrategy beacon upgrade.
contract ExecuteUpgrade is QueueUpgrade {
    using Env for *;

    function _runAsMultisig() internal virtual override prank(Env.protocolCouncilMultisig()) {
        _requireSepoliaPatchEnv();

        bytes memory calldataToExecutor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldataToExecutor,
            predecessor: 0,
            salt: 0
        });
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        _requireSepoliaPatchEnv();

        super.runAsEOA();

        TimelockController timelock = Env.timelockController();
        bytes memory calldataToExecutor = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldataToExecutor,
            predecessor: 0,
            salt: 0
        });

        assertFalse(timelock.isOperationPending(txHash), "transaction should not be queued");
        QueueUpgrade._runAsMultisig();
        _unsafeResetHasPranked();

        assertTrue(timelock.isOperationPending(txHash), "transaction should be queued");
        assertFalse(timelock.isOperationReady(txHash), "transaction should not be ready");
        assertFalse(timelock.isOperationDone(txHash), "transaction should not be complete");

        vm.warp(block.timestamp + timelock.getMinDelay());
        assertTrue(timelock.isOperationReady(txHash), "transaction should be executable");

        execute();

        TestUtils.validateDurationVaultStrategyProxyAdmin();
        TestUtils.validateDurationVaultStrategyStorage();
        TestUtils.validateDurationVaultStrategyImplConstructors();
        TestUtils.validateDurationVaultStrategyImplAddressesMatchProxy();
        TestUtils.validateDurationVaultStrategyProtocolRegistry();
    }
}
