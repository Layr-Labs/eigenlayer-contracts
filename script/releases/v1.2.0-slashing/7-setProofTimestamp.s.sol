// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {ExecuteUpgradeAndSetTimestampSubmitter} from "./5-executeUpgradeAndSetTimestampSubmitter.s.sol";
import {QueueUnpause} from "./3-queueUnpause.s.sol";
import {QueueUpgradeAndTimestampSetter} from "./2-queueUpgradeAndTimestampSetter.s.sol";
import {Pause} from "./4-pause.s.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";


/**
 * Purpose: Enqueue a transaction which immediately sets `EigenPodManager.PAUSED_START_CHECKPOINT=true` 
 */
contract SetProofTimestamp is ExecuteUpgradeAndSetTimestampSubmitter {
    using Env for *;

    uint64 proofTimestamp;

    function _runAsMultisig() prank(Env.opsMultisig()) internal virtual override {
        require(proofTimestamp != 0, "proofTimestamp must be set");
        Env.proxy.eigenPodManager().setPectraForkTimestamp(proofTimestamp);
    }

    function setTimestamp(uint64 _proofTimestamp) public {
        proofTimestamp = _proofTimestamp;
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

        // Warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA

        // Execute Upgrade and Set Timestamp Submitter
        ExecuteUpgradeAndSetTimestampSubmitter._runAsMultisig();
        _unsafeResetHasPranked();

        // Set the proof timestamp
        proofTimestamp = 1740434112; // Using holesky pectra fork timestamp for testing
        execute();   

        // Validate that the proof timestamp is set
        assertEq(Env.proxy.eigenPodManager().pectraForkTimestamp(), proofTimestamp, "Proof timestamp is not set");
    }
}