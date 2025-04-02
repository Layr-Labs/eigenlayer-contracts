// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-deployContracts.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose:
 *      * enqueue a multisig transaction which;
 *             - upgrades DM, EPM, EP
 *  This should be run via the protocol council multisig.
 */
contract QueueUpgrade is MultisigBuilder, Deploy {
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
        // Core - Pods - Strategies - Permissions
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.delegationManager()),
                impl: address(Env.impl.delegationManager())
            })
        }).append({
            to: address(Env.beacon.eigenPod()),
            data: Encode.upgradeableBeacon.upgradeTo({newImpl: address(Env.impl.eigenPod())})
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.eigenPodManager()),
                impl: address(Env.impl.eigenPodManager())
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
