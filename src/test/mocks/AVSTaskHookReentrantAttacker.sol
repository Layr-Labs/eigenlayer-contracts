// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IBN254CertificateVerifierTypes} from "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import {OperatorSet} from "src/contracts/libraries/OperatorSetLib.sol";
import {BN254} from "src/contracts/libraries/BN254.sol";

import {TaskMailbox} from "src/contracts/avs/task/TaskMailbox.sol";
import {ITaskMailboxTypes} from "src/contracts/interfaces/ITaskMailbox.sol";
import {IAVSTaskHook} from "src/contracts/interfaces/IAVSTaskHook.sol";

/**
 * @title AVSTaskHookReentrantAttacker
 * @notice Mock contract for testing reentrancy protection in AVSTaskHook
 */
contract AVSTaskHookReentrantAttacker is IAVSTaskHook, ITaskMailboxTypes {
    TaskMailbox public taskMailbox;

    // Store attack parameters individually to avoid memory/storage issues
    address public refundCollector;
    address public executorOperatorSetAvs;
    uint32 public executorOperatorSetId;
    bytes public payload;

    bytes32 public attackTaskHash;
    bytes public result;
    bool public attackOnPost;
    bool public attackCreateTask;

    constructor(address _taskMailbox) {
        taskMailbox = TaskMailbox(_taskMailbox);
    }

    function setAttackParams(
        TaskParams memory _taskParams,
        bytes32 _attackTaskHash,
        IBN254CertificateVerifierTypes.BN254Certificate memory, // unused but kept for interface compatibility
        bytes memory _result,
        bool _attackOnPost,
        bool _attackCreateTask
    ) external {
        // Store TaskParams fields individually
        refundCollector = _taskParams.refundCollector;
        executorOperatorSetAvs = _taskParams.executorOperatorSet.avs;
        executorOperatorSetId = _taskParams.executorOperatorSet.id;
        payload = _taskParams.payload;

        attackTaskHash = _attackTaskHash;
        result = _result;
        attackOnPost = _attackOnPost;
        attackCreateTask = _attackCreateTask;
    }

    function validatePreTaskCreation(address, TaskParams memory) external view {}

    function handlePostTaskCreation(bytes32) external override {
        if (attackOnPost) {
            if (attackCreateTask) {
                // Reconstruct TaskParams for the attack
                TaskParams memory taskParams = TaskParams({
                    refundCollector: refundCollector,
                    executorOperatorSet: OperatorSet(executorOperatorSetAvs, executorOperatorSetId),
                    payload: payload
                });
                // Try to reenter createTask
                taskMailbox.createTask(taskParams);
            } else {
                // Reconstruct certificate for the attack
                IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
                    referenceTimestamp: uint32(block.timestamp),
                    messageHash: attackTaskHash,
                    signature: BN254.G1Point(0, 0),
                    apk: BN254.G2Point([uint(0), uint(0)], [uint(0), uint(0)]),
                    nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
                });
                // Try to reenter submitResult
                taskMailbox.submitResult(attackTaskHash, abi.encode(cert), result);
            }
        }
    }

    function validatePreTaskResultSubmission(address, bytes32, bytes memory, bytes memory) external view {}

    function handlePostTaskResultSubmission(address, bytes32) external {
        if (!attackOnPost) {
            if (attackCreateTask) {
                // Reconstruct TaskParams for the attack
                TaskParams memory taskParams = TaskParams({
                    refundCollector: refundCollector,
                    executorOperatorSet: OperatorSet(executorOperatorSetAvs, executorOperatorSetId),
                    payload: payload
                });
                // Try to reenter createTask
                taskMailbox.createTask(taskParams);
            } else {
                // Reconstruct certificate for the attack
                IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
                    referenceTimestamp: uint32(block.timestamp),
                    messageHash: attackTaskHash,
                    signature: BN254.G1Point(0, 0),
                    apk: BN254.G2Point([uint(0), uint(0)], [uint(0), uint(0)]),
                    nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
                });
                // Try to reenter submitResult
                taskMailbox.submitResult(attackTaskHash, abi.encode(cert), result);
            }
        }
    }

    function calculateTaskFee(ITaskMailboxTypes.TaskParams memory) external pure returns (uint96) {
        return 0;
    }
}
