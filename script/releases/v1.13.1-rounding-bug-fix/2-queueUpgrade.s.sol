// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {DeployImplementations} from "./1-deployImplementations.s.sol";
import {CoreUpgradeQueueBuilder} from "../CoreUpgradeQueueBuilder.sol";
import "../Env.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {IProtocolRegistry, IProtocolRegistryTypes} from "src/contracts/interfaces/IProtocolRegistry.sol";

/// Purpose: Queue the DelegationManager upgrade for the queued slash accounting fix.
/// This script queues an upgrade to:
/// - DelegationManager (fix queued slashable share accounting under rounding)
/// - ProtocolRegistry semantic version bump to 1.13.1
contract QueueUpgrade is DeployImplementations, MultisigBuilder {
    using Env for *;
    using Encode for *;
    using CoreUpgradeQueueBuilder for *;

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
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls();

        // 1. Upgrade DelegationManager proxy to the new implementation.
        executorCalls.upgradeDelegationManager();

        // 2. Bump ProtocolRegistry semantic version to 1.13.1.
        executorCalls.append({
            to: address(Env.proxy.protocolRegistry()),
            data: abi.encodeCall(
                IProtocolRegistry.ship,
                (
                    new address[](0),
                    new IProtocolRegistryTypes.DeploymentConfig[](0),
                    new string[](0),
                    Env.deployVersion()
                )
            )
        });

        return Encode.gnosisSafe
            .execTransaction({
                from: address(Env.timelockController()),
                to: Env.multiSendCallOnly(),
                op: Encode.Operation.DelegateCall,
                data: Encode.multiSend(executorCalls)
            });
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed() || Env._versionGte(Env.envVersion(), Env.deployVersion())) {
            return;
        }

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

        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        assertEq(
            address(Env.impl.delegationManager()).code.length, 0, "implementation must remain undeployed while queued"
        );

        execute();

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
    }
}
