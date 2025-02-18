// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

contract MockAVSRegistrar {
    function supportsAVS(
        address /*avs*/
    ) external pure returns (bool) {
        return true;
    }

    fallback () external {}
}

contract MockAVSRegistrarAlt {
    bool public isSupported;
    
    function supportsAVS(
        address /*avs*/
    ) external view returns (bool) {
        return isSupported;
    }

    function registerOperator(
        address /*operator*/,
        address /*avs*/,
        uint32[] calldata /*operatorSetIds*/,
        bytes calldata /*data*/
    ) external {}

    function deregisterOperator(
        address /*operator*/,
        address /*avs*/,
        uint32[] calldata /*operatorSetIds*/
    ) external {}

    function setSupportsAVS(bool value) external {
        isSupported = value;
    }
}