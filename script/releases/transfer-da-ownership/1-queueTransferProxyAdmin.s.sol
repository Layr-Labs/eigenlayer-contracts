// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {TransferDAOwnershipConstants as C} from "./TransferDAOwnershipConstants.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * Purpose:
 *      * enqueue a multisig transaction which;
 *             - transfers DA Proxy Admin ownership to the EigenDA ops multisig
 *  This should be run via the protocol council multisig.
 */
contract QueueTransferProxyAdmin is MultisigBuilder {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
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
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: C.EIGEN_DA_PROXY_ADMIN,
            data: abi.encodeCall(Ownable.transferOwnership, (C.EIGEN_DA_OPS_MULTISIG))
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual {
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
        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        execute();

        // Check that the upgrade has been added to the timelock
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }
}
