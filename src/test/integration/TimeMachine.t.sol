// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/test/utils/Logger.t.sol";

contract TimeMachine is Test, Logger {
    uint256[] public snapshots;

    function NAME() public view virtual override returns (string memory) {
        return "TimeMachine";
    }

    /// -----------------------------------------------------------------------
    /// Setters
    /// -----------------------------------------------------------------------

    function createSnapshot() public returns (uint256 snapshot) {
        snapshots.push(snapshot = cheats.snapshotState());
        print.method("createSnapshot", cheats.toString(snapshot));
    }

    function travelToLast() public returns (uint256 currentSnapshot) {
        // Safety check to make sure createSnapshot is called before attempting
        // to warp so we don't accidentally prevent our own births.
        assertTrue(pastExists(), "Global.warpToPast: invalid usage, past does not exist");
        uint256 last = lastSnapshot();
        print.method("travelToLast", cheats.toString(last));
        currentSnapshot = createSnapshot();
        cheats.revertToState(last);
    }

    function travel(
        uint256 snapshot
    ) public {
        print.method("travel", cheats.toString(snapshot));
        cheats.revertToState(snapshot);
    }

    /// -----------------------------------------------------------------------
    /// Getters
    /// -----------------------------------------------------------------------

    function lastSnapshot() public view returns (uint256) {
        return snapshots[snapshots.length - 1];
    }

    function pastExists() public view returns (bool) {
        return snapshots.length != 0;
    }
}
