// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/Snapshots.sol";

contract SnapshotsHarness {
    using Snapshots for Snapshots.DefaultWadHistory;
    using Snapshots for Snapshots.DefaultZeroHistory;

    Snapshots.DefaultWadHistory internal wadHistory;
    Snapshots.DefaultZeroHistory internal zeroHistory;

    function pushWad(uint32 key, uint64 value) public {
        wadHistory.push(key, value);
    }

    function upperLookupWad(uint32 key) public view returns (uint64) {
        return wadHistory.upperLookup(key);
    }

    function latestWad() public view returns (uint64) {
        return wadHistory.latest();
    }

    function lengthWad() public view returns (uint) {
        return wadHistory.length();
    }

    function pushZero(uint32 key, uint value) public {
        zeroHistory.push(key, value);
    }

    function upperLookupZero(uint32 key) public view returns (uint) {
        return zeroHistory.upperLookup(key);
    }

    function latestZero() public view returns (uint) {
        return zeroHistory.latest();
    }

    function lengthZero() public view returns (uint) {
        return zeroHistory.length();
    }
}

contract SnapshotsUnitTests is Test {
    SnapshotsHarness harness;

    function setUp() public {
        harness = new SnapshotsHarness();
    }

    /// forge-config: default.allow_internal_expect_revert = true
    function test_Revert_InvalidSnapshotOrdering(uint r) public {
        uint32 key = uint32(bound(r, 1, type(uint32).max));
        uint32 smallerKey = uint32(bound(r, 0, key - 1));

        harness.pushWad(key, 1);

        vm.expectRevert(Snapshots.InvalidSnapshotOrdering.selector);
        harness.pushWad(smallerKey, 2);
    }

    function test_Push_Correctness(uint r) public {
        uint32 key = uint32(bound(r, 0, type(uint32).max));
        uint64 value = uint32(bound(r, 0, type(uint64).max));

        harness.pushWad(key, value);

        assertEq(harness.upperLookupWad(key), value);
        assertEq(harness.latestWad(), value);
        assertEq(harness.lengthWad(), 1);
    }

    function test_UpperLookup_InitiallyWad(uint32 r) public view {
        assertEq(harness.upperLookupWad(r), 1e18);
    }

    function test_Latest_InitiallyWad() public view {
        assertEq(harness.latestWad(), 1e18);
    }

    function test_Length_InitiallyZero() public view {
        assertEq(harness.lengthWad(), 0);
    }

    function test_Revert_InvalidSnapshotOrdering_ZeroHistory(uint r) public {
        uint32 key = uint32(bound(r, 1, type(uint32).max));
        uint32 smallerKey = uint32(bound(r, 0, key - 1));

        harness.pushZero(key, 1);

        vm.expectRevert(Snapshots.InvalidSnapshotOrdering.selector);
        harness.pushZero(smallerKey, 2);
    }

    function test_Push_Correctness_ZeroHistory(uint r) public {
        uint32 key = uint32(bound(r, 0, type(uint32).max));
        uint value = bound(r, 0, type(uint224).max);

        harness.pushZero(key, value);

        assertEq(harness.upperLookupZero(key), value);
        assertEq(harness.latestZero(), value);
        assertEq(harness.lengthZero(), 1);
    }

    function test_UpperLookup_InitiallyZero(uint32 r) public view {
        assertEq(harness.upperLookupZero(r), 0);
    }

    function test_Latest_InitiallyZero() public view {
        assertEq(harness.latestZero(), 0);
    }

    function test_Length_InitiallyZero_ZeroHistory() public view {
        assertEq(harness.lengthZero(), 0);
    }
}
