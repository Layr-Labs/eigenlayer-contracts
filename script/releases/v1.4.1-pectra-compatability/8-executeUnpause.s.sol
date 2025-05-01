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
 */

contract ExecuteUnpause is SetProofTimestamp {
    using Env for *;

    function _runAsMultisig() internal virtual override prank(Env.protocolCouncilMultisig()) {
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
        // 1-4 are completed in _completeSteps1_4()
        _completeSteps1_4();

        // Warp past delay - this works for both the upgrade & the unpause
        TimelockController timelock = Env.timelockController();
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA

        // 5. Execute Upgrade and Set Timestamp Submitter
        ExecuteUpgradeAndSetTimestampSubmitter._runAsMultisig();
        _unsafeResetHasPranked();

        // 7. Set the proof timestamp
        SetProofTimestamp._runAsMultisig();
        _unsafeResetHasPranked();

        // Check that unpausing is ready (predecessor is done & we've warped past delay)
        assertTrue(
            timelock.isOperationDone(QueueUpgradeAndTimestampSetter.getTimelockId()), "Predecessor should be done."
        );
        assertTrue(timelock.isOperationReady(QueueUnpause.getTimelockId()), "Transaction should be executable.");

        // 8. Execute Unpause
        execute();

        // Validate that the unpause is successful
        assertEq(Env.proxy.eigenPodManager().paused(), 0, "Not unpaused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), false, "Not unpaused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS), false, "Not unpaused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_VERIFY_STALE_BALANCE), false, "Not unpaused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS), false, "Not unpaused!");
    }
}
