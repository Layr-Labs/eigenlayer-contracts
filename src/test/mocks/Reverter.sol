// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

contract Reverter {
    fallback() external {
        revert("Reverter: I am a contract that always reverts");
    }
}

contract ReverterWithDecimals is Reverter {
    function decimals() external pure returns (uint8) {
        return 18;
    }
}
