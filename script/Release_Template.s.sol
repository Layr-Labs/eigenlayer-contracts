// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ConfigParser.sol";
import "script/utils/EncodeSafeMultisendMainnet.sol";
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

/// TODO: break all abstract contracts out into their own file in a `template` directory

/// @notice template for an EOA script
abstract contract EOABuilder is ConfigParser {

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
abstract contract MultisigBuilder is ConfigParser, EncodeSafeTransactionMainnet {

    /// @return a Transaction object for a Gnosis Safe to ingest
    function execute(string memory envPath) public returns (bytes memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        Tx[] memory txs = _execute(addrs, env, params);

        return encodeMultisendTxs(txs);
    }

    function test_Execute(string memory envPath) public {
        execute(envPath);
    }

    /// @notice to be implemented by inheriting contract
    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Tx[] memory);
}

abstract contract NestedMultisigBuilder is ConfigParser {

    /// @return a Transaction object for a Gnosis Safe to ingest
    /// @dev this object is intended to hold calldata to be sent to *yet another* Safe
    /// which will contain the actual relevant calldata
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
abstract contract OpsTimelockBuilder is NestedMultisigBuilder {

    /// @return a Transaction object for a Gnosis Safe to ingest
    function queue(string memory envPath) public returns (bytes memory) {
        // TODO

        // get response from _queue()
        // encode for Timelock
        // return encoded call for Ops Multisig

        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        TimelockTx[] memory ttx = _queue(addrs, env, params);

        return encodeTimelockTxn(ttx);
    }
    
    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory);

    function _makeTimelockTxns(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (Transaction memory);
}
