// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/templates/QueueBuilder.sol";
import "script/utils/MultisigCallUtils.sol";

contract Queue is QueueBuilder {
    using MultisigCallUtils for MultisigCall[];

    function _queue(Addresses memory addrs, Environment memory env, Params memory params) public override returns (MultisigCall[] memory) {

        //////////////////////////
        // construct executor data here
        //////////////////////////

        return _executorCalls;
    }
}