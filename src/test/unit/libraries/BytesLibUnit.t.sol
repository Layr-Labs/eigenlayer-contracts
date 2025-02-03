// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/BytesLib.sol";

contract BytesLibUnitTests is Test {
    using BytesLib for bytes;

    function test_Concat_Basic() public pure {
        bytes memory a = hex"1234";
        bytes memory b = hex"5678";
        bytes memory expected = hex"12345678";

        bytes memory result = a.concat(b);
        assertTrue(result.equal(expected));
    }

    function test_Concat_EmptyInputs() public pure {
        bytes memory empty = hex"";
        bytes memory data = hex"1234";
        
        assertTrue(empty.concat(empty).equal(empty));
        assertTrue(data.concat(empty).equal(data));
        assertTrue(empty.concat(data).equal(data));
    }

    function test_Slice_Basic() public pure {
        bytes memory data = hex"0123456789";
        bytes memory expected = hex"234567";
        
        bytes memory result = data.slice(1, 3);
        assertTrue(result.equal(expected));
    }

    function test_Revert_SliceOutOfBounds() public {
        bytes memory data = hex"0123456789";
        
        // Test start beyond length
        vm.expectRevert(BytesLib.OutOfBounds.selector);
        data.slice(10, 1);
        
        // Test length beyond data bounds
        vm.expectRevert(BytesLib.OutOfBounds.selector);
        data.slice(0, 11);
        
        // Test start + length beyond bounds
        vm.expectRevert(BytesLib.OutOfBounds.selector);
        data.slice(5, 6);
    }

    function test_ToAddress() public pure {
        bytes memory data = hex"000000000000000000000000A0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48";
        address expected = 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48;
        
        assertEq(data.toAddress(12), expected);
    }

    function test_ToUint() public pure {
        bytes memory data = hex"000000000000000000000000000000000000000000000000000000000000002A";
        
        assertEq(data.toUint256(0), 42);
        assertEq(data.toUint8(31), 42);
        assertEq(data.toUint16(30), 42);
        assertEq(data.toUint32(28), 42);
        assertEq(data.toUint64(24), 42);
        assertEq(data.toUint96(20), 42);
        assertEq(data.toUint128(16), 42);
    }

    function test_Equal() public pure {
        bytes memory a = hex"1234567890";
        bytes memory b = hex"1234567890";
        bytes memory c = hex"1234567891";
        
        assertTrue(a.equal(b));
        assertFalse(a.equal(c));
    }

    function test_Revert_ToTypesOutOfBounds() public {
        bytes memory tooShort = hex"1234";
        
        vm.expectRevert(BytesLib.OutOfBounds.selector);
        tooShort.toAddress(0);
        
        vm.expectRevert(BytesLib.OutOfBounds.selector);
        tooShort.toUint256(0);
    }
} 