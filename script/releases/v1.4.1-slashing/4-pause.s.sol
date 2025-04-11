// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";

/**
 * Purpose: Enqueue a transaction which immediately sets
 *  `EigenPodManager.PAUSED_START_CHECKPOINT=true` and
 *  `EigenPodManager.PAUSED_EIGENPODS_VERIFY_CREDENTIALS=true`
 */
contract Pause is MultisigBuilder, EigenPodPausingConstants {
    using Env for *;

    function _runAsMultisig() internal virtual override prank(Env.pauserMultisig()) {
        uint256 mask = 2 ** PAUSED_START_CHECKPOINT | 2 ** PAUSED_EIGENPODS_VERIFY_CREDENTIALS
            | 2 ** PAUSED_VERIFY_STALE_BALANCE | 2 ** PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS;

        Env.proxy.eigenPodManager().pause(mask);
    }

    function testScript() public virtual {
        execute();

        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), "Not paused!");
        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS), "Not paused!");
        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_VERIFY_STALE_BALANCE), "Not paused!");
        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS), "Not paused!");

        // Create a new pod and try to verify credentials + start checkpoint
        EigenPod pod = EigenPod(payable(Env.proxy.eigenPodManager().createPod()));

        // Revert verifying withdrawal credentials
        vm.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        BeaconChainProofs.StateRootProof memory emptyProof;
        pod.verifyWithdrawalCredentials(0, emptyProof, new uint40[](0), new bytes[](0), new bytes32[][](0));

        // Revert starting checkpoint
        vm.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        pod.startCheckpoint(false);

        // Revert verifying stale balance
        BeaconChainProofs.ValidatorProof memory validatorProof;
        vm.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        pod.verifyStaleBalance(0, emptyProof, validatorProof);

        // Revert completing checkpoint
        BeaconChainProofs.BalanceContainerProof memory balanceContainerProof;
        BeaconChainProofs.BalanceProof[] memory proofs;
        vm.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        pod.verifyCheckpointProofs(balanceContainerProof, proofs);
    }
}
