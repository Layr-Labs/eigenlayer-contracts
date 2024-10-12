// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Addresses, Environment, Params, ConfigParser} from "script/utils/ConfigParser.sol";
import {MultisigCall, MultisigCallUtils} from "script/utils/MultisigCallUtils.sol";
import {SafeTx, EncGnosisSafe} from "script/utils/SafeTxUtils.sol";

/**
 * @title MultisigBuilder
 * @dev Abstract contract for building arbitrary multisig scripts.
 */
abstract contract MultisigBuilder is ConfigParser {

    using MultisigCallUtils for MultisigCall[];

    /**
     * @notice Constructs a SafeTx object for a Gnosis Safe to ingest.
     * @param envPath The path to the relevant environment configuration file.
     * @return A SafeTx struct containing the transaction data to post to the Safe API.
     */
    function execute(string memory envPath) public returns (SafeTx memory) {
        // read in config file for relevant environment
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        // get calls for Multisig from inheriting script
        MultisigCall[] memory calls = _execute(addrs, env, params);

        // encode calls as MultiSend data
        bytes memory data = calls.encodeMultisendTxs();

        // creates and return SafeTx object
        // assumes 0 value (ETH) being sent to multisig
        return SafeTx({
            to: params.multiSendCallOnly,
            value: 0,
            data: data,
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }

    /**
     * @notice To be implemented by inheriting contract.
     * @param addrs A struct containing the addresses involved in the multisig call.
     * @param env A struct containing the environment settings for the multisig call.
     * @param params A struct containing the parameters for the multisig call.
     * @return An array of MultisigCall objects.
     */
    function _execute(Addresses memory addrs, Environment memory env, Params memory params) public virtual returns (MultisigCall[] memory);
}