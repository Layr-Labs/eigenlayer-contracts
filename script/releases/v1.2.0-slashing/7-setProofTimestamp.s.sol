// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {ExecuteUpgradeAndSetTimestampSubmitter} from "./5-executeUpgradeAndSetTimestampSubmitter.s.sol";
import {QueueUnpause} from "./3-queueUnpause.s.sol";
import {QueueUpgradeAndTimestampSetter} from "./2-queueUpgradeAndTimestampSetter.s.sol";
import {Pause} from "./4-pause.s.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";


/**
 * Purpose: Propose a transaction to set the proof timestamp
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
        // 1-4 are completed in _completeSteps1_4()
        _completeSteps1_4();

        // Warp past delay
        TimelockController timelock = Env.timelockController();
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA

        // 5. Execute Upgrade and Set Timestamp Submitter
        ExecuteUpgradeAndSetTimestampSubmitter._runAsMultisig();
        _unsafeResetHasPranked();

        // 6. Set the proof timestamp
        proofTimestamp = 1740434112; // Using holesky pectra fork timestamp for testing
        execute();   

        // Validate that the proof timestamp is set
        assertEq(Env.proxy.eigenPodManager().pectraForkTimestamp(), proofTimestamp, "Proof timestamp is not set");
    }
}