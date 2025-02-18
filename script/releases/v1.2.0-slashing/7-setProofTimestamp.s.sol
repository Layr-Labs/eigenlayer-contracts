// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";

/**
 * Purpose: Enqueue a transaction which immediately sets `EigenPodManager.PAUSED_START_CHECKPOINT=true` 
 */
contract SetProofTimestamp is MultisigBuilder {
    using Env for *;

    function _runAsMultisig() prank(Env.opsMultisig()) internal virtual override {
        
    }

    function testScript() public virtual {
        
    }
}