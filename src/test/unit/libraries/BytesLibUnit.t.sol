// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/BytesLib.sol";

contract BytesLibHarness {
    function concat(bytes memory a, bytes memory b) public pure returns (bytes memory) {
        return BytesLib.concat(a, b);
    }

    function slice(bytes memory data, uint start, uint length) public pure returns (bytes memory) {
        return BytesLib.slice(data, start, length);
    }

    function toAddress(bytes memory data, uint offset) public pure returns (address) {
        return BytesLib.toAddress(data, offset);
    }

    function toUint256(bytes memory data, uint offset) public pure returns (uint) {
        return BytesLib.toUint256(data, offset);
    }

    function toUint128(bytes memory data, uint offset) public pure returns (uint128) {
        return BytesLib.toUint128(data, offset);
    }

    function toUint96(bytes memory data, uint offset) public pure returns (uint96) {
        return BytesLib.toUint96(data, offset);
    }

    function toUint64(bytes memory data, uint offset) public pure returns (uint64) {
        return BytesLib.toUint64(data, offset);
    }

    function toUint32(bytes memory data, uint offset) public pure returns (uint32) {
        return BytesLib.toUint32(data, offset);
    }

    function toUint16(bytes memory data, uint offset) public pure returns (uint16) {
        return BytesLib.toUint16(data, offset);
    }

    function toUint8(bytes memory data, uint offset) public pure returns (uint8) {
        return BytesLib.toUint8(data, offset);
    }

    function equal(bytes memory a, bytes memory b) public pure returns (bool) {
        return BytesLib.equal(a, b);
    }
}

contract BytesLibUnitTests is Test {
    BytesLibHarness harness;

    function setUp() public {
        harness = new BytesLibHarness();
    }

    function test_Concat_Basic() public view {
        bytes memory a = hex"1234";
        bytes memory b = hex"5678";
        bytes memory expected = hex"12345678";

        bytes memory result = harness.concat(a, b);
        assertTrue(harness.equal(result, expected));
    }

    function test_Concat_EmptyInputs() public view {
        bytes memory empty = hex"";
        bytes memory data = hex"1234";

        assertTrue(harness.equal(harness.concat(empty, empty), empty));
        assertTrue(harness.equal(harness.concat(data, empty), data));
        assertTrue(harness.equal(harness.concat(empty, data), data));
    }

    function test_Slice_Basic() public view {
        bytes memory data = hex"0123456789";
        bytes memory expected = hex"234567";

        bytes memory result = harness.slice(data, 1, 3);
        assertTrue(harness.equal(result, expected));
    }

    function test_Revert_SliceOutOfBounds() public {
        bytes memory data = hex"0123456789";

        // Test start beyond length
        vm.expectRevert(BytesLib.OutOfBounds.selector);
        harness.slice(data, 10, 1);

        // Test length beyond data bounds
        vm.expectRevert(BytesLib.OutOfBounds.selector);
        harness.slice(data, 0, 11);

        // Test start + length beyond bounds
        vm.expectRevert(BytesLib.OutOfBounds.selector);
        harness.slice(data, 5, 6);
    }

    function test_ToAddress() public view {
        bytes memory data = hex"000000000000000000000000A0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48";
        address expected = 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48;

        assertEq(harness.toAddress(data, 12), expected);
    }

    function test_ToUint() public view {
        bytes memory data = hex"000000000000000000000000000000000000000000000000000000000000002A";

        assertEq(harness.toUint256(data, 0), 42);
        assertEq(harness.toUint8(data, 31), 42);
        assertEq(harness.toUint16(data, 30), 42);
        assertEq(harness.toUint32(data, 28), 42);
        assertEq(harness.toUint64(data, 24), 42);
        assertEq(harness.toUint96(data, 20), 42);
        assertEq(harness.toUint128(data, 16), 42);
    }

    function test_Equal() public view {
        bytes memory a = hex"1234567890";
        bytes memory b = hex"1234567890";
        bytes memory c = hex"1234567891";

        assertTrue(harness.equal(a, b));
        assertFalse(harness.equal(a, c));
    }

    function test_Revert_ToTypesOutOfBounds() public {
        bytes memory tooShort = hex"1234";

        vm.expectRevert(BytesLib.OutOfBounds.selector);
        harness.toAddress(tooShort, 0);

        vm.expectRevert(BytesLib.OutOfBounds.selector);
        harness.toUint256(tooShort, 0);
    }
}
