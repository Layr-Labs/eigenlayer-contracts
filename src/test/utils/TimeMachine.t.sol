// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

contract TimeMachine is Test {
    uint[] public snapshots;

    /// -----------------------------------------------------------------------
    /// Setters
    /// -----------------------------------------------------------------------

    function createSnapshot() public returns (uint snapshot) {
        snapshots.push(snapshot = vm.snapshotState());
    }

    function travelToLast() public returns (uint currentSnapshot) {
        // Safety check to make sure createSnapshot is called before attempting
        // to warp so we don't accidentally prevent our own births.
        assertTrue(pastExists(), "Global.warpToPast: invalid usage, past does not exist");
        uint last = lastSnapshot();
        currentSnapshot = createSnapshot();
        vm.revertToState(last);
    }

    function travel(uint snapshot) public {
        vm.revertToState(snapshot);
    }

    /// -----------------------------------------------------------------------
    /// Getters
    /// -----------------------------------------------------------------------

    function lastSnapshot() public view returns (uint) {
        return snapshots[snapshots.length - 1];
    }

    function pastExists() public view returns (bool) {
        return snapshots.length != 0;
    }
}
