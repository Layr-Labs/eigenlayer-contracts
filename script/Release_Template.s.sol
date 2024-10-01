// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./utils/Releasoor.s.sol";

abstract contract EOABuilder {
    function deploy(string memory envPath) public returns (Deployment[]);
    function _deploy(Addresses memory addrs, Env, Params) internal virtual returns (Deployment[]);
}

abstract contract MultisigBuilder {
    function execute(string memory envPath) public returns (Transaction memory) {
        // TODO - pull from Releasoor.run
    }

    function _execute(Addresses memory addrs, Env, Params) internal virtual returns (Calls[] memory);
}

abstract contract OpsTimelockBuilder is MultisigBuilder {
    function queue(string memory envPath) public returns (Calls[]);
    function _queue(Addresses memory addrs, Env, Params) internal virtual returns (Calls[] memory);
    function _makeTimelockTxns(Addresses memory addrs, Env, Params) internal virtual returns (Calls[]);
}

contract ExampleScript is MultisigBuilder {

    struct Transaction {
        address to;
        uint value;
        bytes data;
        Operation op;
    }

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory) {


        return Transaction({
            to: addrs.admin.timelock,
            data: calldata_to_timelock_queueing_action
        });
    }
}