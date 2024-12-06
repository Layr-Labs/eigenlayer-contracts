// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";
import {EncGnosisSafe} from "zeus-templates/utils/EncGnosisSafe.sol";
import {MultisigCallUtils, MultisigCall} from "zeus-templates/utils/MultisigCallUtils.sol";

/**
 * Purpose: Enqueue a transaction which immediately sets `EigenPodManager.PAUSED_START_CHECKPOINT=true` 
 */
contract Pause is MultisigBuilder {
    using MultisigCallUtils for MultisigCall[];
    using EigenLabsUpgrade for *;
    using EncGnosisSafe for *;
    using MultisigCallUtils for *;

    MultisigCall[] private _executorCalls;
    MultisigCall[] private _opsCalls;


    function options() internal override view returns (MultisigOptions memory) {
        return MultisigOptions(
            // TODO: should this run from the pauser multisig
            this._operationsMultisig(),
            Operation.Call
        );
    }

    function runAsMultisig() internal virtual override {
        // TODO: set: 
        //     EigenPodManager.PAUSED_START_CHECKPOINT = true;
    }

    function testDeploy() virtual public {
        execute();

        // TODO: assert side effects
    }
}
