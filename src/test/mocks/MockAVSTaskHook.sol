// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "src/contracts/libraries/OperatorSetLib.sol";
import {IAVSTaskHook} from "src/contracts/interfaces/IAVSTaskHook.sol";
import {ITaskMailboxTypes} from "src/contracts/interfaces/ITaskMailbox.sol";

contract MockAVSTaskHook is IAVSTaskHook {
    uint96 public defaultFee = 1 ether;

    function setDefaultFee(uint96 _fee) external {
        defaultFee = _fee;
    }

    function validatePreTaskCreation(address, /*caller*/ ITaskMailboxTypes.TaskParams memory /*taskParams*/ ) external view {
        //TODO: Implement
    }

    function handlePostTaskCreation(bytes32 /*taskHash*/ ) external {
        //TODO: Implement
    }

    function validatePreTaskResultSubmission(address, /*caller*/ bytes32, /*taskHash*/ bytes memory, /*cert*/ bytes memory /*result*/ )
        external
        view
    {
        //TODO: Implement
    }

    function handlePostTaskResultSubmission(bytes32 /*taskHash*/ ) external {
        //TODO: Implement
    }

    function calculateTaskFee(ITaskMailboxTypes.TaskParams memory /*taskParams*/ ) external view returns (uint96) {
        return defaultFee;
    }
}
