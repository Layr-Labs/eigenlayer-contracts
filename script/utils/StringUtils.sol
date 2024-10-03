// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

library StringUtils {

    function eq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }

    function concat(string memory a, string memory b) internal pure returns (string memory) {
        return string.concat(a, b);
    }
}