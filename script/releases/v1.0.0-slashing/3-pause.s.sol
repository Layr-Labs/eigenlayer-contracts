// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {EncGnosisSafe} from "zeus-templates/utils/EncGnosisSafe.sol";
import {MultisigCallUtils, MultisigCall} from "zeus-templates/utils/MultisigCallUtils.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * Purpose: Enqueue a transaction which immediately sets `EigenPodManager.PAUSED_START_CHECKPOINT=true` 
 */
contract Pause is MultisigBuilder {
    using MultisigCallUtils for MultisigCall[];
    using EncGnosisSafe for *;
    using MultisigCallUtils for *;
    using Env for *;

    MultisigCall[] private _executorCalls;
    MultisigCall[] private _opsCalls;


    function options() internal override view returns (MultisigOptions memory) {
        return MultisigOptions(
            Env.pauserMultisig(),
            Operation.Call
        );
    }

    function runAsMultisig() internal virtual override {
        // TODO: set: 
        //    pause EigenPodManager
    }

    function testDeploy() virtual public {
        execute();

        // TODO: assert side effects
    }
}
