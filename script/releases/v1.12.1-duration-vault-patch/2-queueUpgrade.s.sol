// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {CoreUpgradeQueueBuilder} from "../CoreUpgradeQueueBuilder.sol";
import "../Env.sol";
import "./1-deployImplementations.s.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/// Purpose: queue the Sepolia-only DurationVaultStrategy beacon upgrade.
contract QueueUpgrade is DeployImplementations, MultisigBuilder {
    using Env for *;
    using Encode for *;
    using CoreUpgradeQueueBuilder for *;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        _requireSepoliaPatchEnv();

        bytes memory calldataToExecutor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.schedule({
            target: Env.executorMultisig(),
            value: 0,
            data: calldataToExecutor,
            predecessor: 0,
            salt: 0,
            delay: timelock.getMinDelay()
        });
    }

    function _getCalldataToExecutor() internal returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls();
        executorCalls.upgradeDurationVaultStrategy();

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        _requireSepoliaPatchEnv();
        runAsEOA();

        TimelockController timelock = Env.timelockController();
        bytes memory calldataToExecutor = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldataToExecutor,
            predecessor: 0,
            salt: 0
        });

        assertFalse(timelock.isOperationPending(txHash), "transaction should not be queued");
        _validatePreUpgradeState();

        execute();

        assertTrue(timelock.isOperationPending(txHash), "transaction should be queued");
    }

    function _validatePreUpgradeState() internal view {
        require(
            address(Env.impl.durationVaultStrategy()).code.length > 0,
            "DurationVaultStrategy implementation not deployed"
        );
        require(
            Env.beacon.durationVaultStrategy().implementation() != address(Env.impl.durationVaultStrategy()),
            "DurationVaultStrategy beacon already upgraded"
        );
    }
}
