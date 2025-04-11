// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-eoa.s.sol";
import "../Env.sol";
import "forge-std/console.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose: Queues an upgrade to EP/EPM
 * and sets the timestamp submitter to the ops multisig
 */
contract QueueUpgradeAndTimestampSetter is MultisigBuilder, Deploy {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        bytes memory calldata_to_executor = _getCalldataToExecutor_queueUpgrade();

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
    function _getCalldataToExecutor_queueUpgrade() internal virtual returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: address(Env.beacon.eigenPod()),
            data: Encode.upgradeableBeacon.upgradeTo({newImpl: address(Env.impl.eigenPod())})
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.eigenPodManager()),
                impl: address(Env.impl.eigenPodManager())
            })
        });

        // Set the timestamp submitter to the ops multisig
        executorCalls.append({
            to: address(Env.proxy.eigenPodManager()),
            data: abi.encodeCall(EigenPodManager.setProofTimestampSetter, (address(Env.opsMultisig())))
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: address(Env.multiSendCallOnly()),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual {
        runAsEOA();

        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor_queueUpgrade();
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });

        // Check that the upgrade does not exist in the timelock
        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        execute();

        // Check that the upgrade has been added to the timelock
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }

    function getTimelockId() public virtual returns (bytes32) {
        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor_queueUpgrade();
        return timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
    }
}
