// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ConfigParser.sol";
import {EncGnosisSafe} from "script/utils/Encoders.sol";

/// @notice Deployment data struct
struct Deployment {
    string name;
    address deployedTo;
}

/// @notice Tx data struct
/// @dev based on <https://docs.safe.global/sdk/api-kit/guides/propose-and-confirm-transactions#propose-a-transaction-to-the-service>
struct Transaction {
    address to;
    uint256 value;
    bytes data;
    EncGnosisSafe.Operation op;
}

/// @notice template for an EOA script
abstract contract EOABuilder is ConfigParser {
    Deployment[] internal deployments;

    function deploy(string memory envPath) public returns (Deployment[] memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        return _deploy(addrs, env, params);
    }

    function _deploy(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Deployment[] memory);
}

/// @notice template for a Multisig script
abstract contract MultisigBuilder is ConfigParser {

    function execute(string memory envPath) public returns (Transaction memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        return _execute(addrs, env, params);
    }

    /// @notice to be implemented by inheriting contract
    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory);
}


/// @notice template for an OpsMultisig script that goes through the timelock
abstract contract OpsTimelockBuilder is MultisigBuilder {
    function queue(string memory envPath) public returns (Transaction memory) {
        // TODO
    }
    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory);

    function _makeTimelockTxns(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory);
}
