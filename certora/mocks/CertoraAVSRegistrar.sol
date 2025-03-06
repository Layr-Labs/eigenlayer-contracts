// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IAVSRegistrar.sol";

contract CertoraAVSRegistrar is IAVSRegistrar {
    function registerOperator(address, address, uint32[] calldata, bytes calldata) external {}
    function deregisterOperator(address, address, uint32[] calldata) external {}

    function supportsAVS(
        address
    ) external view returns (bool) {
        return true;
    }
}
