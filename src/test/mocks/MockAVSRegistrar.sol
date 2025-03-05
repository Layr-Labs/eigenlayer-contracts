// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

contract MockAVSRegistrar {
    function supportsAVS(address /*avs*/ ) external pure returns (bool) {
        return true;
    }

    fallback() external {}
}
