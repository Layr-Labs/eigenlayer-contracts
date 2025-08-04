// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ITaskMailboxTypes} from "./ITaskMailbox.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";

/**
 * @title IAVSTaskHook
 * @author Layr Labs, Inc.
 * @notice Interface for AVS-specific task lifecycle hooks.
 * @dev This interface allows AVSs to implement custom validation logic for tasks.
 */
interface IAVSTaskHook {
    /**
     * @notice Validates a task before it is created
     * @param caller Address that is creating the task
     * @param taskParams Task parameters
     * @dev This function should revert if the task should not be created
     */
    function validatePreTaskCreation(address caller, ITaskMailboxTypes.TaskParams memory taskParams) external view;

    /**
     * @notice Handles a task after it is created
     * @param taskHash Unique identifier of the task
     * @dev This function can be used to perform additional validation or update AVS-specific state
     */
    function handlePostTaskCreation(
        bytes32 taskHash
    ) external;

    /**
     * @notice Validates a task before it is submitted for verification
     * @param caller Address that is submitting the result
     * @param taskHash Unique identifier of the task
     * @param cert Certificate proving the validity of the result
     * @param result Task execution result data
     * @dev This function should revert if the task should not be verified
     */
    function validatePreTaskResultSubmission(
        address caller,
        bytes32 taskHash,
        bytes memory cert,
        bytes memory result
    ) external view;

    /**
     * @notice Handles a task result submission
     * @param caller Address that submitted the result
     * @param taskHash Unique identifier of the task
     * @dev This function can be used to perform additional validation or update AVS-specific state
     */
    function handlePostTaskResultSubmission(address caller, bytes32 taskHash) external;

    /**
     * @notice Calculates the fee for a task payload against a specific fee market
     * @param taskParams The task parameters
     * @return The fee for the task
     */
    function calculateTaskFee(
        ITaskMailboxTypes.TaskParams memory taskParams
    ) external view returns (uint96);
}
