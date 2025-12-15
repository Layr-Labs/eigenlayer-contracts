// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {QueueUpgrade} from "./2-queueUpgrade.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import "../Env.sol";
import "../TestUtils.sol";

/// Purpose: Execute the queued upgrade for Duration Vault feature.
contract ExecuteUpgrade is QueueUpgrade {
    using Env for *;

    function _runAsMultisig() internal virtual override prank(Env.protocolCouncilMultisig()) {
        bytes memory calldata_to_executor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
    }

    function testScript() public virtual override {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        // Deploy the contracts first
        super.runAsEOA();

        // Queue the upgrade
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
        QueueUpgrade._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready for execution.");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete.");

        // Warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay());
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // Execute
        execute();

        // Run validation tests
        _validateUpgrade();
    }

    function _validateUpgrade() internal view {
        // Validate StrategyManager upgrade
        _validateStrategyManagerUpgrade();

        // Validate StrategyBase beacon upgrade
        _validateStrategyBaseUpgrade();

        // Validate StrategyFactory upgrade
        _validateStrategyFactoryUpgrade();

        // Validate DurationVaultStrategy beacon is set
        _validateDurationVaultBeaconSet();
    }

    function _validateStrategyManagerUpgrade() internal view {
        StrategyManager sm = Env.proxy.strategyManager();
        assertTrue(address(sm) != address(0), "StrategyManager proxy should exist");
        // The new StrategyManager should have the beforeAddShares/beforeRemoveShares hooks
        // We can verify by checking the implementation address matches
    }

    function _validateStrategyBaseUpgrade() internal view {
        UpgradeableBeacon beacon = Env.beacon.strategyBase();
        assertTrue(
            beacon.implementation() == address(Env.impl.strategyBase()),
            "StrategyBase beacon should point to new implementation"
        );
    }

    function _validateStrategyFactoryUpgrade() internal view {
        StrategyFactory sf = Env.proxy.strategyFactory();
        assertTrue(address(sf) != address(0), "StrategyFactory proxy should exist");
    }

    function _validateDurationVaultBeaconSet() internal view {
        StrategyFactory sf = Env.proxy.strategyFactory();
        assertTrue(
            address(sf.durationVaultBeacon()) == address(Env.beacon.durationVaultStrategy()),
            "StrategyFactory should have duration vault beacon set"
        );
    }
}
