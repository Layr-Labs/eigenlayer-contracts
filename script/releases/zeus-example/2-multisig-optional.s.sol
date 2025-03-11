// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Deploy} from "./1-eoa.s.sol";
import "../Env.sol";
import "forge-std/console.sol";
import "zeus-templates/utils/ZEnvHelpers.sol";  

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {EmptyContract} from "./EmptyContract.sol";

/**
 * Purpose: Tests that zeus can handle multisig steps without transactions.
 */
contract MultisigOptional is MultisigBuilder, Deploy {
    using Env for *;
    using Encode for *;
    using ZEnvHelpers for *;

    function _runAsMultisig()
        internal
        virtual
        override
        prank(Env.opsMultisig())
    {
        // this script does not output any transactions.
        EmptyContract c = EmptyContract(ZEnvHelpers.state().deployedImpl(type(EmptyContract).name));
        if ((c.stateVar() % 2) == 1) {
            // we only increment if it's odd! otherwise, this step does nothing.
            c.increment();
        }
    }  

    function testScript() public virtual {
        runAsEOA();
        
        EmptyContract c = EmptyContract(ZEnvHelpers.state().deployedImpl(type(EmptyContract).name));
        uint256 prev = c.stateVar();

        execute();

        if (prev % 2 == 0) {
            assertEq(prev, c.stateVar());
        } else {
            assertEq(prev + 1, c.stateVar());
        }
    }
}
