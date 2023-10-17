// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

interface IFunctionGateway {
    function request(
        bytes32 functionId,
        bytes memory inputs,
        bytes4 callbackSelector,
        bytes memory context
    ) external payable;
}