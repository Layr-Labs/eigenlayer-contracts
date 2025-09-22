// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployTaskMailboxImpl} from "./1-deployTaskMailboxImpl.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {MultisigCall, Encode} from "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * @title QueueTaskMailboxUpgrade
 * @notice Queue the TaskMailbox upgrade transaction in the Timelock via the Operations Multisig.
 *         This queues the upgrade to fix the task replay vulnerability.
 */
contract QueueTaskMailboxUpgrade is MultisigBuilder, DeployTaskMailboxImpl {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        if (!(Env.isDestinationChain() && Env._strEq(Env.envVersion(), "1.8.0"))) {
            return;
        }

        bytes memory calldata_to_executor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.schedule({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0,
            delay: timelock.getMinDelay()
        });
    }

    /// @dev Get the calldata to be sent from the timelock to the executor
    function _getCalldataToExecutor() internal returns (bytes memory) {
        /// forgefmt: disable-next-item
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.taskMailbox()),
                impl: address(Env.impl.taskMailbox())
            })
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual override {
        if (!(Env.isDestinationChain() && Env._strEq(Env.envVersion(), "1.8.0"))) {
            return;
        }

        // Deploy the new TaskMailbox implementation first
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

        // Check that the upgrade does not exist in the timelock
        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued yet");

        // Queue the upgrade
        execute();

        // Check that the upgrade has been added to the timelock
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready immediately");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be done");

        // Validate that the TaskMailbox proxy still points to the old implementation
        address currentImpl = Env._getProxyImpl(address(Env.proxy.taskMailbox()));
        address newImpl = address(Env.impl.taskMailbox());
        assertFalse(currentImpl == newImpl, "TaskMailbox proxy should still point to old implementation");
    }
}
