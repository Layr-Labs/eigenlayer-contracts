// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-deployContracts.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose:
 *      * enqueue a multisig transaction which;
 *             - upgrades EP
 *  This should be run via the operations multisig.
 */
contract QueueUpgrade is MultisigBuilder, Deploy {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.communityMultisig()) {
        bytes memory calldata_to_executor = _getCalldataToExecutor();

        (bool success,) = address(Env.executorMultisig()).call(calldata_to_executor);
        assertTrue(success, "Upgrade failed");
    }

    /// @dev Get the calldata to be sent from the timelock to the executor
    function _getCalldataToExecutor() internal returns (bytes memory) {
        /// forgefmt: disable-next-item
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: address(Env.beacon.eigenPod()),
            data: Encode.upgradeableBeacon.upgradeTo({
                newImpl: address(Env.impl.eigenPod())
            })
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.communityMultisig()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual override {
        runAsEOA();

        execute();

        _validateNewImplAddresses({areMatching: true});
        _validateProxyAdmins();
        _validateProxyConstructors();
    }

    /// @dev Mirrors the checks done in 1-deployContracts, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal view {
        {
            UpgradeableBeacon eigenPodBeacon = Env.beacon.eigenPod();
            assertTrue(eigenPodBeacon.implementation() == address(Env.impl.eigenPod()), "eigenPodBeacon.impl invalid");

            /// EigenPod
            EigenPod eigenPod = Env.impl.eigenPod();
            assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
            assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");
            assertEq(eigenPod.version(), Env.deployVersion(), "ep.version failed");
        }
    }
}
