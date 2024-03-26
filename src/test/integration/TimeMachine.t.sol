// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";

contract TimeMachine is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    bool pastExists = false;
    uint lastSnapshot;

    uint64 public proofGenStartTime;

    function createSnapshot() public returns (uint) {
        uint snapshot = cheats.snapshot();
        lastSnapshot = snapshot;
        pastExists = true;
        return snapshot;
    }

    function warpToLast() public returns (uint curState) {
        // Safety check to make sure createSnapshot is called before attempting to warp
        // so we don't accidentally prevent our own births
        assertTrue(pastExists, "Global.warpToPast: invalid usage, past does not exist");

        curState = cheats.snapshot();
        cheats.revertTo(lastSnapshot);
        return curState;
    }

    function warpToPresent(uint curState) public {
        cheats.revertTo(curState);
    }

    /// @dev Sets the timestamp we use for proof gen to now,
    /// then sets block timestamp to now + secondsAgo.
    ///
    /// This means we can create mock proofs using an oracle time
    /// of `proofGenStartTime`.
    function setProofGenStartTime(uint secondsAgo) public {
        proofGenStartTime = uint64(block.timestamp);
        cheats.warp(block.timestamp + secondsAgo);
    }
}
