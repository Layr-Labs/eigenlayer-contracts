// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";

contract UpgradeCounter is MultisigBuilder {

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (Transaction memory) {

        bytes memory calldata_to_executor;
        return Transaction({
            to: addrs.timelock,
            value: 0,
            data: calldata_to_executor,
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }
}