// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueDestinationChain} from "./5-queueDestinationChain.s.sol";
import {Encode} from "zeus-templates/utils/Encode.sol";

/// @notice Executes the upgrade for v1.7.0 multichain testnet final destination chain
contract ExecuteDestinationChain is QueueDestinationChain {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal override prank(Env.protocolCouncilMultisig()) {
        if (!Env.isDestinationChain()) {
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
        if (!Env.isDestinationChain()) {
            return;
        }

        // 1 - Deploy. The contracts have been deployed in the deploy upgrade script
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

        /// 2 - Queue. Check that the operation IS ready
        QueueDestinationChain._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready for execution.");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete.");

        // 3 - warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertTrue(timelock.isOperationReady(txHash), "Transaction should be executable.");

        // 4 - execute
        execute();
        assertTrue(
            timelock.isOperationDone(txHash),
            "v1.8.0 multichain testnet final destination chain txn should be complete."
        );

        // 5 - Validate
        _validateNewImplAddresses({areMatching: true});
        _validateProxyAdmins();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Mirrors the checks done in 4-deployDestinationChain, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal view {
        OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
        assertEq(operatorTableUpdater.version(), Env.deployVersion(), "operatorTableUpdater version mismatch");
        assertTrue(
            operatorTableUpdater.bn254CertificateVerifier() == Env.proxy.bn254CertificateVerifier(),
            "operatorTableUpdater bn254CertificateVerifier mismatch"
        );
        assertTrue(
            operatorTableUpdater.ecdsaCertificateVerifier() == Env.proxy.ecdsaCertificateVerifier(),
            "operatorTableUpdater ecdsaCertificateVerifier mismatch"
        );

        ECDSACertificateVerifier ecdsaCertificateVerifier = Env.proxy.ecdsaCertificateVerifier();
        assertEq(ecdsaCertificateVerifier.version(), Env.deployVersion(), "ecdsaCertificateVerifier version mismatch");
        assertTrue(
            ecdsaCertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
            "ecdsaCertificateVerifier operatorTableUpdater mismatch"
        );

        BN254CertificateVerifier bn254CertificateVerifier = Env.proxy.bn254CertificateVerifier();
        assertEq(bn254CertificateVerifier.version(), Env.deployVersion(), "bn254CertificateVerifier version mismatch");
        assertTrue(
            bn254CertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
            "bn254CertificateVerifier operatorTableUpdater mismatch"
        );

        TaskMailbox taskMailbox = Env.proxy.taskMailbox();
        assertEq(taskMailbox.version(), Env.deployVersion(), "taskMailbox version mismatch");
        assertTrue(
            taskMailbox.BN254_CERTIFICATE_VERIFIER() == address(Env.proxy.bn254CertificateVerifier()),
            "taskMailbox BN254_CERTIFICATE_VERIFIER mismatch"
        );
        assertTrue(
            taskMailbox.ECDSA_CERTIFICATE_VERIFIER() == address(Env.proxy.ecdsaCertificateVerifier()),
            "taskMailbox ECDSA_CERTIFICATE_VERIFIER mismatch"
        );
        assertEq(taskMailbox.MAX_TASK_SLA(), Env.MAX_TASK_SLA(), "taskMailbox MAX_TASK_SLA mismatch");
    }

    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// OperatorTableUpdater - dummy parameters
        OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
        OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;

        vm.expectRevert(errInit);
        operatorTableUpdater.initialize(
            address(0), // owner
            0, // initial paused status
            dummyOperatorSet, // globalRootConfirmerSet
            0, // globalRootConfirmationThreshold
            dummyBN254Info // globalRootConfirmerSetInfo
        );

        /// TaskMailbox
        TaskMailbox taskMailbox = Env.proxy.taskMailbox();
        vm.expectRevert(errInit);
        taskMailbox.initialize(
            address(0), // owner
            0, // feeSplit
            address(0) // feeSplitCollector
        );

        // ECDSACertificateVerifier and BN254CertificateVerifier don't have initialize functions
    }
}
