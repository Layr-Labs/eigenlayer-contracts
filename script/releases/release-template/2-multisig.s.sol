// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {ITimelock} from "zeus-templates/interfaces/ITimelock.sol";

contract Queue is MultisigBuilder {
    using MultisigCallUtils for MultisigCall[];

    MultisigCall[] private _executorCalls;
    MultisigCall[] private _opsCalls;

    function _queue() internal returns (MultisigCall[] memory) {
        //////////////////////////
        // construct executor data here
        //////////////////////////

        return _executorCalls;
    }

    /**
     * @dev Can be overridden so that executing script can inherit this contract.
     */
    function _execute() internal virtual override returns (MultisigCall[] memory) {
        // get the queue data
        MultisigCall[] memory calls = _queue();

        address multiSendCallOnly = zeusAddress("MultiSendCallOnly");
        address timelock = zeusAddress("Timelock");

        // encode calls for executor
        bytes memory executorCalldata = calls.makeExecutorCalldata(multiSendCallOnly, timelock);

        address executorMultisig = zeusAddress("ExecutorMultisig");

        // encode executor data for timelock
        bytes memory timelockCalldata = abi.encodeWithSelector(
            ITimelock.queueTransaction.selector, executorMultisig, 0, "", executorCalldata, type(uint256).max
        );

        _opsCalls.append(timelock, timelockCalldata);

        // encode timelock data for ops multisig
        return _opsCalls;
    }

    function zeusTest() public override {
        // Test function implementation
    }
}
