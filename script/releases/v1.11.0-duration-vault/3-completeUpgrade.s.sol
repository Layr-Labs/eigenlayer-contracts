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
        // Only execute on version 1.10.0
        if (!Env._strEq(Env.envVersion(), "1.10.0")) {
            return;
        }

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
        if (!Env.isCoreProtocolDeployed() || !Env.isSource() || !Env._strEq(Env.envVersion(), "1.10.0")) {
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

        // Run validation tests using TestUtils
        // Note: We skip validateProxyStorage() because it validates all core contracts
        // including RewardsCoordinator.activationDelay(), which this upgrade doesn't touch.
        // The activationDelay may have been changed on-chain and won't match the env config.
        TestUtils.validateProxyAdmins();
        TestUtils.validateProxyConstructors();
        TestUtils.validateProxiesAlreadyInitialized();
        TestUtils.validateImplAddressesMatchProxy();
        TestUtils.validateProtocolRegistry();

        // v1.11.0+ DurationVaultStrategy-specific validations
        TestUtils.validateDurationVaultStrategyProxyAdmin();
        TestUtils.validateDurationVaultStrategyStorage();
        TestUtils.validateDurationVaultStrategyImplConstructors();
        TestUtils.validateDurationVaultStrategyImplAddressesMatchProxy();
        TestUtils.validateDurationVaultStrategyProtocolRegistry();
    }

    /// @notice ensure that the gas used
    function testGas() public virtual {
        if (!Env.isCoreProtocolDeployed() || !Env.isSource() || !Env._strEq(Env.envVersion(), "1.10.0")) {
            return;
        }

        uint256 MAX_GAS_PER_TX = 16_777_216;

        // Deploy the contracts first
        super.runAsEOA();

        QueueUpgrade._runAsMultisig();
        TimelockController timelock = Env.timelockController();

        _unsafeResetHasPranked(); // reset hasPranked so we can use it again
        // Warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay());
        _startPrank(Env.protocolCouncilMultisig());
        bytes memory calldata_to_executor = _getCalldataToExecutor();

        vm.startSnapshotGas("execute");
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
        uint256 gasUsed = vm.stopSnapshotGas("execute");
        assertTrue(
            gasUsed <= MAX_GAS_PER_TX, "TX Gas should be less than or equal to the EIP-7825 transaction gas limit cap"
        );
        _stopPrank();
    }
}
