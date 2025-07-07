// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-deployContracts.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {MultisigCall, Encode} from "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose:
 *      * enqueue a multisig transaction which;
 *             - upgrades DM, ALM, SM, EPM
 *             - upgrades strategies (EigenStrategy, StrategyBase, StrategyBaseTVLLimits)
 *  This should be run via the protocol council multisig.
 */
contract QueueUpgrade is MultisigBuilder, Deploy {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        if (!Env.isSourceChain()) {
            return;
        }

        IERC20[] memory tokensToBlacklist = new IERC20[](2);
        tokensToBlacklist[0] = IERC20(address(Env.proxy.eigen()));
        tokensToBlacklist[1] = IERC20(address(Env.proxy.beigen()));

        Env.proxy.strategyFactory().blacklistTokens(tokensToBlacklist);
    }

    function testScript() public virtual override {
        if (!Env.isSourceChain()) {
            return;
        }

        execute();


        assertTrue(
            Env.proxy.strategyFactory().isBlacklisted(IERC20(address(Env.proxy.eigen()))), "eigen should be blacklisted"
        );
        assertTrue(
            Env.proxy.strategyFactory().isBlacklisted(IERC20(address(Env.proxy.beigen()))),
            "beigen should be blacklisted"
        );
    }
}
