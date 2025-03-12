// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "src/contracts/libraries/BeaconChainProofs.sol";
import {Execute, Queue} from "./3-execute.s.sol";
import {DeployFresh} from "./1-deploy.s.sol";


/**
 * Purpose: optionally enqueue a transaction that pauses all EigenPod functionality.
 */
contract Pause is MultisigBuilder, EigenPodPausingConstants, Execute {
    using Env for *;

    function _runAsMultisig() prank(Env.pauserMultisig()) internal virtual override(Execute, MultisigBuilder) {
        if (!Env.supportsNativeEth()) {
            // disable all eigenpod functionality.
            uint mask = (1 << PAUSED_START_CHECKPOINT) | 
            (1 << PAUSED_NEW_EIGENPODS) | 
            (1 << PAUSED_WITHDRAW_RESTAKED_ETH) | 
            (1 << PAUSED_EIGENPODS_VERIFY_CREDENTIALS) | 
            (1 << PAUSED_NON_PROOF_WITHDRAWALS) | 
            (1 << PAUSED_START_CHECKPOINT) | 
            (1 << PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS) |
            (1 << PAUSED_VERIFY_STALE_BALANCE);

            Env.proxy.eigenPodManager().pause(mask);
        }
    }

    function testScript() public virtual override {
        DeployFresh._runAsEOA();
        _unsafeResetHasPranked();

        Queue._runAsMultisig();
        _unsafeResetHasPranked();

        TimelockController timelock = Env.timelockController();
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA

        Execute._runAsMultisig();
        _unsafeResetHasPranked();

        execute();

        bool shouldBePaused = !Env.supportsNativeEth();

        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), shouldBePaused, "PAUSED_START_CHECKPOINT Not paused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_NEW_EIGENPODS), shouldBePaused, "PAUSED_NEW_EIGENPODS Not paused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_WITHDRAW_RESTAKED_ETH), shouldBePaused, "PAUSED_WITHDRAW_RESTAKED_ETH Not paused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS), shouldBePaused, "PAUSED_EIGENPODS_VERIFY_CREDENTIALS Not paused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_NON_PROOF_WITHDRAWALS), shouldBePaused, "PAUSED_NON_PROOF_WITHDRAWALS Not paused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_START_CHECKPOINT), shouldBePaused, "PAUSED_START_CHECKPOINT Not paused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS), shouldBePaused, "PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS Not paused!");
        assertEq(Env.proxy.eigenPodManager().paused(PAUSED_VERIFY_STALE_BALANCE), shouldBePaused, "PAUSED_VERIFY_STALE_BALANCE Not paused!");
    }
}
