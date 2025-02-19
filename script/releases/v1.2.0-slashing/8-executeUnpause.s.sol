// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {ExecuteUpgradeAndSetTimestampSubmitter} from "./5-executeUpgradeAndSetTimestampSubmitter.s.sol";
import {QueueUnpause} from "./3-queueUnpause.s.sol";
import {QueueUpgradeAndTimestampSetter} from "./2-queueUpgradeAndTimestampSetter.s.sol";
import {Pause} from "./4-pause.s.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {SetProofTimestamp} from "./7-setProofTimestamp.s.sol";
/**
 * Purpose: Executes the unpause transaction from step 3
 *
*/
contract ExecuteUnpause is SetProofTimestamp {
    using Env for *;

    function _runAsMultisig() prank(Env.protocolCouncilMultisig()) internal virtual override {
        bytes memory calldata_to_executor = QueueUnpause._getCalldataToExecutor_queueUnpause();

        TimelockController timelock = Env.timelockController();
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: QueueUpgradeAndTimestampSetter.getTimelockId(),
            salt: 0
        });
    }

    function testScript() public virtual override {
        TimelockController timelock = Env.timelockController();
        // Deploy Impls
        _runAsEOA();

        // Queue Upgrade and Set Timestamp Submitter
        QueueUpgradeAndTimestampSetter._runAsMultisig();
        _unsafeResetHasPranked();

        // Queue Unpause
        QueueUnpause._runAsMultisig();
        _unsafeResetHasPranked();

        // Run Pausing Logic
        Pause._runAsMultisig();
        _unsafeResetHasPranked();

        // Warp past delay - this works for both the upgrade & the unpause
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA

        // Execute Upgrade and Set Timestamp Submitter
        ExecuteUpgradeAndSetTimestampSubmitter._runAsMultisig();
        _unsafeResetHasPranked();

        // Set the proof timestamp
        SetProofTimestamp.setTimestamp(1740434112); // Using holesky pectra fork timestamp for testing
        SetProofTimestamp._runAsMultisig();
        _unsafeResetHasPranked();

        // Check that unpausing is ready (predecessor is done & we've warped past delay)
        assertTrue(timelock.isOperationDone(QueueUpgradeAndTimestampSetter.getTimelockId()), "Predecessor should be done.");
        assertTrue(timelock.isOperationReady(QueueUnpause.getTimelockId()), "Transaction should be executable.");


        // Execute Unpause
        execute();

        // Validate that the unpause is successful
        assertEq(Env.proxy.eigenPodManager().paused(), 0, "Not unpaused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), false, "Not unpaused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS), false, "Not unpaused!");
    }
}