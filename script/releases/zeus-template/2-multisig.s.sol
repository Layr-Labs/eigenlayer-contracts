// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Addresses, Environment, Params, MultisigCall, MultisigCallUtils, OpsTimelockBuilder} from "zeus-templates/templates/OpsTimelockBuilder.sol";

contract Queue is OpsTimelockBuilder {
    using MultisigCallUtils for MultisigCall[];

    MultisigCall[] internal _executorCalls;

    function queue(Addresses memory addrs, Environment memory env, Params memory params) public override returns (MultisigCall[] memory) {

        //////////////////////////
        // construct executor data here
        //////////////////////////

        return _executorCalls;
    }
}