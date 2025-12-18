// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../releases/Env.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";

// Types
import "src/contracts/interfaces/IRewardsCoordinator.sol";

/// Purpose: Update the generator on a PREPROD/TESTNET environment
contract SetSubmitter is MultisigBuilder {
    using Env for *;

    address submitter = 0x1E4c2A1F731F9AEeaBCB5ad8d5EdAbE6b4a04492;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        Env.proxy.rewardsCoordinator().setRewardsForAllSubmitter(submitter, true);
    }

    function testScript() public virtual {
        _runAsMultisig();
        assertEq(Env.proxy.rewardsCoordinator().isRewardsForAllSubmitter(submitter), true);
    }
}

