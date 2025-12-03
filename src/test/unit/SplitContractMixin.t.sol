// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/mixins/SplitContractMixin.sol";

contract SplitContractMixinTest is Test, SplitContractMixin {
    uint value;
    Delegate public delegate;

    constructor() SplitContractMixin(address(0x123)) {
        value = vm.randomUint();
        delegate = new Delegate();
    }

    function getValue() public view returns (uint result) {
        _delegateView(address(delegate));
        result;
    }

    function test_getValue() public {
        (bool success, bytes memory data) = address(this).call(abi.encodeWithSelector(this.getValue.selector));
        assertTrue(success);
        uint result = abi.decode(data, (uint));
        assertEq(result, value);
    }
}

// Mock contract to test delegation
contract Delegate is Test {
    uint value;

    function getValue() public view returns (uint result) {
        return value;
    }
}
