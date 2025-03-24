// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";

/**
 * Purpose: Enqueue a transaction which immediately sets `EigenPodManager.PAUSED_START_CHECKPOINT=true`
 */
contract Pause is MultisigBuilder, EigenPodPausingConstants {
    using Env for *;

    function _runAsMultisig() internal virtual override prank(Env.pauserMultisig()) {
        uint256 mask = 1 << PAUSED_START_CHECKPOINT;

        Env.proxy.eigenPodManager().pause(mask);
    }

    function testScript() public virtual {
        execute();

        assertTrue(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), "Not paused!");

        // Create a new pod and try to start a checkpoint
        EigenPod pod = EigenPod(payable(Env.proxy.eigenPodManager().createPod()));

        // At this point in the upgrade process, we're not using error types yet
        vm.expectRevert("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager");
        pod.startCheckpoint(false);
    }
}
