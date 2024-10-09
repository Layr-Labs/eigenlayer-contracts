// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./MultisigBuilder.sol";

import "script/utils/SafeTxUtils.sol";

/// @notice template for an OpsMultisig script that goes through the timelock
/// @dev writing a script is done from the perspective of the OpsMultisig
abstract contract OpsTimelockBuilder is MultisigBuilder {

    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;

    MultisigCall[] _executorCalls;
    MultisigCall[] _opsCalls;

    /// @return a SafeTx object for a Gnosis Safe to ingest
    function queue(string memory envPath) public returns (SafeTx memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        MultisigCall[] memory calls = _queue(addrs, env, params);

        // encode calls for executor
        bytes memory executorCalldata = makeExecutorCalldata(calls, params.multiSendCallOnly, addrs.timelock);

        // encode executor data for timelock
        bytes memory timelockCalldata = abi.encodeWithSelector(
            ITimelock.queueTransaction.selector,
            addrs.executorMultisig,
            0,
            "",
            executorCalldata,
            type(uint256).max
        );

        // encode timelock data for ops multisig
        return SafeTx({
            to: addrs.timelock,
            value: 0,
            data: timelockCalldata,
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }

    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (MultisigCall[] memory);

    /// @notice helper function to create calldata for executor
    /// can be used for queue or execute
    function makeExecutorCalldata(MultisigCall[] memory calls, address multiSendCallOnly, address timelock) internal pure returns (bytes memory) {
        bytes memory data = calls.encodeMultisendTxs();

        bytes memory executorCalldata = SafeTx({
            to: multiSendCallOnly,
            value: 0,
            data: data,
            op: EncGnosisSafe.Operation.DelegateCall
        }).encodeForExecutor(timelock);

        return executorCalldata;
    }
}