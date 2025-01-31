// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/Snapshots.sol";

contract SnapshotsUnitTests is Test {
    using Snapshots for Snapshots.DefaultWadHistory;

    Snapshots.DefaultWadHistory history;

    function test_Revert_InvalidSnapshotOrdering(uint256 r) public {
        uint32 key = uint32(bound(r, 1, type(uint32).max));
        uint32 smallerKey = uint32(bound(r, 0, key - 1));

        history.push(key, 1);

        vm.expectRevert(Snapshots.InvalidSnapshotOrdering.selector);
        history.push(smallerKey, 2);
    }

    function test_Push_Correctness(uint256 r) public {
        uint32 key = uint32(bound(r, 0, type(uint32).max));
        uint64 value = uint32(bound(r, 0, type(uint64).max));

        history.push(key, value);

        assertEq(history.upperLookup(key), value);
        assertEq(history.latest(), value);
        assertEq(history.length(), 1);
    }

    function test_UpperLookup_InitiallyWad(uint32 r) public view {
        assertEq(history.upperLookup(r), 1e18);
    }

    function test_Latest_InitiallyWad() public view {
        assertEq(history.latest(), 1e18);
    }

    function test_Length_InitiallyZero() public view {
        assertEq(history.length(), 0);
    }
}