// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import "forge-std/console.sol";

import {DeployGovernance} from "./1-deployGovernance.s.sol";
import {DeployPauser} from "./2-deployPauser.s.sol";
import {DeployToken} from "./3-deployToken.s.sol";
import {DeployCore} from "./4-deployCore.s.sol";
import {Queue} from "./5-queueUpgrade.s.sol";
import {Execute} from "./6-executeUpgrade.s.sol";

/// @notice Proposes a transaction to set the proof timestamp
/// @dev This is needed for Eigenpods due to Pectra fork timestamp being set after the fork
contract SetForkTimestamp is Execute {
    using Env for *;
    using ZEnvHelpers for *;

    uint64 proofTimestamp;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        proofTimestamp = ZEnvHelpers.state().envU64("PECTRA_FORK_TIMESTAMP");

        // Sanity check that the timestamp is nonzero
        require(proofTimestamp > 0, "proofTimestamp is zero");

        Env.proxy.eigenPodManager().setPectraForkTimestamp(proofTimestamp);
    }

    function setTimestamp(
        uint64 _proofTimestamp
    ) public {
        proofTimestamp = _proofTimestamp;
    }

    function testScript() public virtual override {
        /// Complete steps 1-6 of the upgrade script
        _mode = OperationalMode.EOA;
        DeployGovernance._runAsEOA();
        DeployPauser._runAsEOA();
        DeployToken._runAsEOA();
        DeployCore._runAsEOA();
        Queue._runAsMultisig();
        _unsafeResetHasPranked();
        vm.warp(block.timestamp + Env.timelockController().getMinDelay()); // 1 tick after ETA
        Execute._runAsMultisig();
        _unsafeResetHasPranked();

        // Execute the set fork timestamp
        execute();

        // Validate that the proof timestamp is set
        assertEq(Env.proxy.eigenPodManager().pectraForkTimestamp(), proofTimestamp, "Proof timestamp is not set");
    }

    function _getProofTimestamp() internal view returns (uint64) {
        string memory timestampPath =
            string.concat(vm.projectRoot(), "/script/releases/v1.4.1-slashing/forkTimestamp.txt");
        string memory timestamp = vm.readFile(timestampPath);
        return uint64(vm.parseUint(timestamp));
    }
}
