// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

contract MockDecimals {
    function decimals() public pure returns (uint8) {
        return 18;
    }
}
