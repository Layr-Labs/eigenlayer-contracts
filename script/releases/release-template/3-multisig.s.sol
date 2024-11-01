// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {SafeTx, SafeTxUtils} from "zeus-templates/utils/SafeTxUtils.sol";
import {ITimelock} from "zeus-templates/interfaces/ITimelock.sol";
import {Queue} from "./2-multisig.s.sol";

contract Execute is Queue {
    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;

    MultisigCall[] private _opsCalls;

    /**
     * @dev Overrides the previous _execute function to execute the queued transactions.
     */
    function _execute() internal override returns (MultisigCall[] memory) {
        MultisigCall[] memory _executorCalls = _queue();

        address multiSendCallOnly = zeusAddress("MultiSendCallOnly");
        address timelock = zeusAddress("Timelock");

        bytes memory executorCalldata = _executorCalls.makeExecutorCalldata(multiSendCallOnly, timelock);

        // execute queued transaction
        _opsCalls.append({
            to: timelock,
            value: 0,
            data: abi.encodeWithSelector(ITimelock.executeTransaction.selector, executorCalldata)
        });

        //////////////////////////
        // add more opsCalls here
        //////////////////////////

        return _opsCalls;
    }
}
