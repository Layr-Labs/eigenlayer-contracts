// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-eoa.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

contract Queue is MultisigBuilder, Deploy {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() prank(Env.opsMultisig()) internal virtual override {
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
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls()
            /// core/
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.allocationManager()),
                    impl: address(Env.impl.allocationManager())
                })
            })
            .append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({
                    proxy: address(Env.proxy.delegationManager()),
                    impl: address(Env.impl.delegationManager())
                })
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