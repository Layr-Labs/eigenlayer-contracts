// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ConfigParser.sol";
import {EncGnosisSafe} from "script/utils/Encoders.sol";

struct Deployment {
    string name;
    address deployedTo;
}

struct Transaction {
    address to;
    uint256 value;
    bytes data;
    EncGnosisSafe.Operation op;
}

abstract contract EOABuilder {
    function deploy(string memory envPath) public returns (Deployment[] memory) {
        // TODO
    }
    function _deploy(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Deployment[] memory);
}

abstract contract MultisigBuilder is ConfigParser {

    function execute(string memory envPath) public returns (Transaction memory) {
        // TODO - pull from Releasoor.run
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        return _execute(addrs, env, params);
    }

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory);
}

abstract contract OpsTimelockBuilder is MultisigBuilder {
    function queue(string memory envPath) public returns (Transaction memory) {
        // TODO
    }
    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory);
    function _makeTimelockTxns(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory);
}

contract ExampleScript is MultisigBuilder {

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