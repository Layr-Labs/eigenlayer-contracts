// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployDestinationChain} from "./4-deployDestinationChain.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {MultisigCall, Encode} from "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose:
 *      * enqueue a multisig transaction which;
 *             - upgrades OperatorTableUpdater, ECDSACertificateVerifier, BN254CertificateVerifier
 *  This should be run via the protocol council multisig.
 */
contract QueueDestinationChain is MultisigBuilder, DeployDestinationChain {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        if (!Env.isDestinationChain()) {
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
                proxy: address(Env.proxy.operatorTableUpdater()),
                impl: address(Env.impl.operatorTableUpdater())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.ecdsaCertificateVerifier()),
                impl: address(Env.impl.ecdsaCertificateVerifier())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.bn254CertificateVerifier()),
                impl: address(Env.impl.bn254CertificateVerifier())
            })
        }).append({
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
        if (!Env.isDestinationChain()) {
            return;
        }

        // Deploy contracts first
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
        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        execute();

        // Check that the upgrade has been added to the timelock
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }
}
