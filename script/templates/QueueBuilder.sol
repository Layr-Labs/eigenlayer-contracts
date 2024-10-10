// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Addresses, Environment, Params, ConfigParser} from "script/utils/ConfigParser.sol";
import {MultisigCall, MultisigCallUtils} from "script/utils/MultisigCallUtils.sol";
import {SafeTx, SafeTxUtils} from "script/utils/SafeTxUtils.sol";

import {ITimelock} from "script/utils/Interfaces.sol";
import {EncGnosisSafe} from "script/utils/Encoders.sol";

/**
 * @title QueueBuilder
 * @dev Abstract contract for queueing transactions for the Executor Multisig via the Ops Multisig and Timelock.
 * @dev See <https://docs.eigenlayer.xyz/eigenlayer/security/multisig-governance> for more details on the relationship between the Executor Multisig, Operations Multisig, and Timelock.
 */
abstract contract QueueBuilder is ConfigParser {

    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;

    /**
     * @dev To be used in _queue() to craft executor calls.
     */
    MultisigCall[] internal _executorCalls;

    /**
     * @notice Prepares a SafeTx for the Ops Multisig to queue on the Timelock.
     * @param envPath The path to the environment configuration file.
     * @return A SafeTx object for a Gnosis Safe to ingest.
     */
    function queue(string memory envPath) public returns (SafeTx memory) {
        // read config file for relevant environment
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        // get calls for Executor from inheriting script
        MultisigCall[] memory calls = _queue(addrs, env, params);

        // encode calls as MultiSend data
        bytes memory executorCalldata = makeExecutorCalldata(calls, params.multiSendCallOnly, addrs.timelock);

        // encode Executor data for Timelock
        bytes memory timelockCalldata = abi.encodeWithSelector(
            ITimelock.queueTransaction.selector,
            addrs.executorMultisig,
            0,
            "",
            executorCalldata,
            type(uint256).max
        );

        // encode Timelock data for Ops Multisig
        return SafeTx({
            to: addrs.timelock,
            value: 0,
            data: timelockCalldata,
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }

    /**
     * @notice Helper function to create calldata to be run by the Executor.
     * @param calls The array of MultisigCall objects.
     * @param multiSendCallOnly The address for multi-send call only.
     * @param timelock The address of the timelock.
     * @return The encoded calldata for the executor.
     */
    function makeExecutorCalldata(MultisigCall[] memory calls, address multiSendCallOnly, address timelock) public pure returns (bytes memory) {
        bytes memory data = calls.encodeMultisendTxs();

        bytes memory executorCalldata = SafeTx({
            to: multiSendCallOnly,
            value: 0,
            data: data,
            op: EncGnosisSafe.Operation.DelegateCall
        }).encodeForExecutor(timelock);

        return executorCalldata;
    }

    /**
     * @notice To be implemented by inheriting contract.
     * @param addrs The relevant addresses for the target environment.
     * @param env The environment configuration.
     * @param params The relevant parameters for the target environment.
     * @return An array of MultisigCall objects scoped to the Executor.
     * @dev This function should be implemented from the context of the
     * Executor Multisig. The queue() function will then prepare the SafeTx
     * for the Ops Multisig to queue on the Timelock.
     * @dev This function is inteded to be called by the "Execute" script after
     * the queue delay period has passed.
     */
    function _queue(Addresses memory addrs, Environment memory env, Params memory params) public virtual returns (MultisigCall[] memory);
}
