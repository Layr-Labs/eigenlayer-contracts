// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {QueueUnpause} from "./3-queueUnpause.s.sol";
import {QueueUpgradeAndTimestampSetter} from "./2-queueUpgradeAndTimestampSetter.s.sol";
import {Pause} from "./4-pause.s.sol";
import "../Env.sol";
import "forge-std/console.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose: Executes the upgrade from step 2, which
 * upgrades the EPM/EP and sets the timestamp submitter to the ops multisig.
 */
contract ExecuteUpgradeAndSetTimestampSubmitter is QueueUnpause, Pause {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override(QueueUnpause, Pause) prank(Env.protocolCouncilMultisig()) {
        bytes memory calldata_to_executor = QueueUpgradeAndTimestampSetter._getCalldataToExecutor_queueUpgrade();

        TimelockController timelock = Env.timelockController();
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
    }

    function testScript() public virtual override(QueueUnpause, Pause) {
        // 1-4 are completed in _completeSteps1_4()
        _completeSteps1_4();

        // Warp past delay
        TimelockController timelock = Env.timelockController();
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(
            timelock.isOperationReady(QueueUpgradeAndTimestampSetter.getTimelockId()),
            true,
            "Transaction should be executable."
        );

        // 5. Execute
        execute();
        assertTrue(
            timelock.isOperationDone(QueueUpgradeAndTimestampSetter.getTimelockId()), "Transaction should be complete."
        );

        // Validate that the operations multisig is the timestamp submitter
        assertEq(
            Env.proxy.eigenPodManager().proofTimestampSetter(),
            Env.opsMultisig(),
            "Timestamp submitter is not the operations multisig"
        );

        // Check that the unpause is not complete
        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), "Not paused!");
        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS), "Not paused!");
        assertFalse(timelock.isOperationDone(QueueUnpause.getTimelockId()), "Transaction should NOT be complete.");

        // Validations
        _validateNewImplAddresses(true);
        _validateProxyAdmins();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Mirrors the checks done in 1-deployContracts, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal virtual {
        UpgradeableBeacon eigenPodBeacon = Env.beacon.eigenPod();
        assertTrue(eigenPodBeacon.implementation() == address(Env.impl.eigenPod()), "eigenPodBeacon.impl invalid");

        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
        assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
        assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
        assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");
    }

    /// @dev Call initialize on all proxies to ensure they are initialized
    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";
        // EigenPod proxies are initialized by individual users

        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        vm.expectRevert(errInit);
        eigenPodManager.initialize(address(0), 0);
        assertTrue(eigenPodManager.owner() == Env.executorMultisig(), "epm.owner invalid");
    }

    function _completeSteps1_4() internal {
        // 0. Get Queue Transactions
        TimelockController timelock = Env.timelockController();
        assertFalse(
            timelock.isOperationPending(QueueUpgradeAndTimestampSetter.getTimelockId()),
            "Transaction should not be queued."
        );
        assertFalse(
            timelock.isOperationReady(QueueUnpause.getTimelockId()), "Transaction should not be ready for execution."
        );

        // 1. Deploy Impls
        runAsEOA();

        // 2. Queue Upgrade and Set Timestamp Submitter
        QueueUpgradeAndTimestampSetter._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(
            timelock.isOperationPending(QueueUpgradeAndTimestampSetter.getTimelockId()), "Transaction should be queued."
        );
        assertFalse(
            timelock.isOperationReady(QueueUpgradeAndTimestampSetter.getTimelockId()),
            "Transaction should NOT be ready for execution."
        );
        assertFalse(
            timelock.isOperationDone(QueueUpgradeAndTimestampSetter.getTimelockId()),
            "Transaction should NOT be complete."
        );

        // 3. Queue Unpause
        QueueUnpause._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(QueueUnpause.getTimelockId()), "Transaction should be queued.");
        assertFalse(
            timelock.isOperationReady(QueueUnpause.getTimelockId()), "Transaction should NOT be ready for execution."
        );
        assertFalse(timelock.isOperationDone(QueueUnpause.getTimelockId()), "Transaction should NOT be complete.");

        // 4. Run Pausing Logic
        Pause._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), "Not paused!");
        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS), "Not paused!");
    }
}
