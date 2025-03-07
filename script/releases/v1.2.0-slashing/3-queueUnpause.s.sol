// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {QueueUpgradeAndTimestampSetter} from "./2-queueUpgradeAndTimestampSetter.s.sol";
import "../Env.sol";
import "forge-std/console.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose: Unpauses the EPM.
 * This needs to be done separately from the upgrade, since we must do actions between
 * the upgrade and unpause.
 */
contract QueueUnpause is QueueUpgradeAndTimestampSetter {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override(QueueUpgradeAndTimestampSetter) prank(Env.opsMultisig()) {
        bytes memory calldata_to_executor = _getCalldataToExecutor_queueUnpause();

        TimelockController timelock = Env.timelockController();
        timelock.schedule({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: QueueUpgradeAndTimestampSetter.getTimelockId(),
            salt: 0,
            delay: timelock.getMinDelay()
        });
    }

    /// @dev Get the calldata to be sent from the timelock to the executor
    function _getCalldataToExecutor_queueUnpause() internal virtual returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: address(Env.proxy.eigenPodManager()),
            data: abi.encodeCall(Pausable.unpause, 0)
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: address(Env.multiSendCallOnly()),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual override(QueueUpgradeAndTimestampSetter) {
        runAsEOA();

        // 1. Run queue upgrade logic
        QueueUpgradeAndTimestampSetter._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        // 2. Run queue unpause logic
        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor_queueUnpause();
        console.log("calldata_to_executor");
        console.logBytes(calldata_to_executor);
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: QueueUpgradeAndTimestampSetter.getTimelockId(),
            salt: 0
        });

        // Check that the upgrade does not exist in the timelock
        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        execute();

        // Check that the upgrade has been added to the timelock
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }

    function getTimelockId() public virtual override returns (bytes32) {
        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor_queueUnpause();
        return timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: QueueUpgradeAndTimestampSetter.getTimelockId(),
            salt: 0
        });
    }
}
