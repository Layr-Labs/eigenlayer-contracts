// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

contract MockAVSRegistrar {
    address _avs;

    function avs() external view returns (address) {
        return _avs;
    }

    function setAvs(
        address newAvs
    ) external {
        _avs = newAvs;
    }

    fallback() external {}
}
