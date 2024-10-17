// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Addresses, Environment, Params, MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {SafeTx, SafeTxUtils} from "zeus-templates/utils/SafeTxUtils.sol";
import {ITimelock} from "zeus-templates/interfaces/ITimelock.sol";
import {Queue} from "./2-multisig.s.sol";

contract Execute is MultisigBuilder {

    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;

    MultisigCall[] internal _opsCalls;

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {

        Queue queue = new Queue();

        MultisigCall[] memory _executorCalls = queue.queue(addrs, env, params);

        bytes memory executorCalldata = queue.makeExecutorCalldata(
            _executorCalls,
            params.multiSendCallOnly,
            addrs.timelock
        );

        // execute queued transaction
        _opsCalls.append({
            to: addrs.timelock,
            value: 0,
            data: abi.encodeWithSelector(
                ITimelock.executeTransaction.selector,
                executorCalldata
            )
        });

        //////////////////////////
        // add more opsCalls here
        //////////////////////////

        return _opsCalls;
    }
}
