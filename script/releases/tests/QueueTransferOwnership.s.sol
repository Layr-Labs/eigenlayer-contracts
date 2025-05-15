// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import "zeus-templates/utils/Encode.sol";

import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * Purpose:
 *      * enqueue a multisig transaction which;
 *             - transfers DA contract ownership to the EigenDA ops multisig
 *  This should be run via the core ops multisig.
 */
contract TransferOwnership is MultisigBuilder {
    using Encode for *;
    using Env for *;

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

    function _getCalldataToExecutor() internal returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: address(Env.proxy.rewardsCoordinator()),
            data: abi.encodeCall(Ownable.transferOwnership, (Env.opsMultisig()))
        }).append({
            to: address(Env.proxy.strategyFactory()),
            data: abi.encodeCall(Ownable.transferOwnership, (Env.opsMultisig()))
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual {
        // _runAsMultisig();

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