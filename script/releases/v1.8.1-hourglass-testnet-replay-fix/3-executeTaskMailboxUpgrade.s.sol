// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueTaskMailboxUpgrade} from "./2-queueTaskMailboxUpgrade.s.sol";
import {Encode} from "zeus-templates/utils/Encode.sol";

/**
 * @title ExecuteTaskMailboxUpgrade
 * @notice Execute the queued TaskMailbox upgrade after the timelock delay.
 *         This completes the upgrade to fix the task replay vulnerability where certificates
 *         could be replayed across different tasks directed at the same operator set.
 */
contract ExecuteTaskMailboxUpgrade is QueueTaskMailboxUpgrade {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal override prank(Env.protocolCouncilMultisig()) {
        if (!(Env.isDestinationChain() && Env._strEq(Env.envVersion(), "1.8.0"))) {
            return;
        }

        bytes memory calldata_to_executor = _getCalldataToExecutor();
        TimelockController timelock = Env.timelockController();

        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
    }

    function testScript() public virtual override {
        if (!(Env.isDestinationChain() && Env._strEq(Env.envVersion(), "1.8.0"))) {
            return;
        }

        // 1 - Deploy. The new TaskMailbox implementation has been deployed
        runAsEOA();

        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });

        // 2 - Queue. Check that the operation IS ready
        QueueTaskMailboxUpgrade._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready immediately");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete");

        // 3 - Warp past the timelock delay
        vm.warp(block.timestamp + timelock.getMinDelay());
        assertTrue(timelock.isOperationReady(txHash), "Transaction should be ready for execution");

        // 4 - Execute the upgrade
        execute();
        assertTrue(timelock.isOperationDone(txHash), "v1.8.1 TaskMailbox upgrade should be complete");

        // 5 - Validate the upgrade was successful
        _validateUpgradeComplete();
        _validateProxyAdmin();
        _validateProxyConstructor();
        _validateProxyInitialized();
        _validateGetMessageHash();
    }

    /// @dev Validate that the TaskMailbox proxy now points to the new implementation
    function _validateUpgradeComplete() internal view {
        address currentImpl = Env._getProxyImpl(address(Env.proxy.taskMailbox()));
        address expectedImpl = address(Env.impl.taskMailbox());

        assertTrue(currentImpl == expectedImpl, "TaskMailbox proxy should point to new implementation");
    }

    /// @dev Validate the proxy's constructor values through the proxy
    function _validateProxyConstructor() internal view {
        TaskMailbox taskMailbox = Env.proxy.taskMailbox();

        // Validate version
        assertEq(
            keccak256(bytes(taskMailbox.version())),
            keccak256(bytes(Env.deployVersion())),
            "TaskMailbox version mismatch"
        );

        // Validate certificate verifiers
        assertTrue(
            taskMailbox.BN254_CERTIFICATE_VERIFIER() == address(Env.proxy.bn254CertificateVerifier()),
            "TaskMailbox BN254_CERTIFICATE_VERIFIER mismatch"
        );
        assertTrue(
            taskMailbox.ECDSA_CERTIFICATE_VERIFIER() == address(Env.proxy.ecdsaCertificateVerifier()),
            "TaskMailbox ECDSA_CERTIFICATE_VERIFIER mismatch"
        );

        // Validate MAX_TASK_SLA
        assertEq(taskMailbox.MAX_TASK_SLA(), Env.MAX_TASK_SLA(), "TaskMailbox MAX_TASK_SLA mismatch");
    }

    /// @dev Validate that the proxy cannot be re-initialized
    function _validateProxyInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        TaskMailbox taskMailbox = Env.proxy.taskMailbox();

        vm.expectRevert(errInit);
        taskMailbox.initialize(
            address(0), // owner
            0, // feeSplit
            address(0) // feeSplitCollector
        );
    }

    /// @dev Validate that the new getMessageHash function works correctly
    function _validateGetMessageHash() internal view {
        TaskMailbox taskMailbox = Env.proxy.taskMailbox();

        // Test the new getMessageHash function with sample data
        bytes32 taskHash = keccak256("test_task");
        bytes memory result = abi.encode("test_result");

        // The new implementation should compute messageHash as keccak256(abi.encode(taskHash, result))
        bytes32 expectedMessageHash = keccak256(abi.encode(taskHash, result));
        bytes32 actualMessageHash = taskMailbox.getMessageHash(taskHash, result);

        assertTrue(
            expectedMessageHash == actualMessageHash,
            "getMessageHash should compute correct message hash with taskHash included"
        );

        // Verify that different tasks with same result produce different message hashes
        bytes32 differentTaskHash = keccak256("different_task");
        bytes32 differentMessageHash = taskMailbox.getMessageHash(differentTaskHash, result);

        assertFalse(
            actualMessageHash == differentMessageHash,
            "Different tasks with same result should produce different message hashes"
        );
    }
}
