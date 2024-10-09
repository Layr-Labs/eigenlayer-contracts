// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ConfigParser.sol";

import "script/utils/SafeTxUtils.sol";

/// @notice to be used for CommunityMultisig
/// NOTE: WIP
abstract contract NestedMultisigBuilder is ConfigParser {

    /// @return a SafeTx object for a Gnosis Safe to ingest
    /// @dev this object is intended to hold calldata to be sent to *yet another* Safe
    /// which will contain the actual relevant calldata
    function execute(string memory envPath) public returns (SafeTx memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        return _execute(addrs, env, params);
    }

    /// @notice to be implemented by inheriting contract
    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (SafeTx memory);
}