// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployContracts} from "./1-deployContracts.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {CoreUpgradeQueueBuilder} from "../CoreUpgradeQueueBuilder.sol";
import "../Env.sol";
import "../TestUtils.sol";

/// Purpose: Queue the upgrade for Duration Vault feature.
/// This script queues upgrades to:
/// - StrategyManager proxy
/// - StrategyBase beacon
/// - StrategyFactory proxy
/// - Creates new DurationVaultStrategy beacon
/// - Sets the duration vault beacon on StrategyFactory
contract QueueUpgrade is DeployContracts {
    using Env for *;
    using Encode for *;
    using CoreUpgradeQueueBuilder for MultisigCall[];

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

        /// core/
        executorCalls.upgradeStrategyManager();

        /// strategies/
        executorCalls.upgradeStrategyBase();
        executorCalls.upgradeStrategyFactory();

        // Set the duration vault beacon on the strategy factory
        // The beacon was deployed in step 1 (deployContracts)
        executorCalls.setDurationVaultBeacon();

        return Encode.gnosisSafe
            .execTransaction({
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

        // Deploy the contracts first
        super.runAsEOA();

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
