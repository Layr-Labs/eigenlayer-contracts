// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ConfigParser.sol";
import "script/utils/MultisigCallUtils.sol";
import {SafeTx, EncGnosisSafe} from "script/utils/SafeTxUtils.sol";

/// @notice template for a Multisig script
abstract contract MultisigBuilder is ConfigParser {

    using MultisigCallUtils for MultisigCall[];
    MultisigCall[] internal _multisigCalls;

    /// @return a SafeTx object for a Gnosis Safe to ingest
    function execute(string memory envPath) public returns (SafeTx memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        // populate _multisigCalls
        _execute(addrs, env, params);

        bytes memory data = _multisigCalls.encodeMultisendTxs();

        return SafeTx({
            to: params.multiSendCallOnly,
            value: 0, // TODO: determine if this should be user-controlled
            data: data,
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }

    function testExecute(string memory envPath) public {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        _execute(addrs, env, params);
        _testExecute(addrs, env, params);
    }

    /// @notice to be implemented by inheriting contract
    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (MultisigCall[] memory);

    /// @notice to be implemented for testing
    function _testExecute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual;
}