// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ConfigParser.sol";
import "script/utils/Encoders.sol";
import "script/utils/Interfaces.sol";
import "script/utils/MultisigCallUtils.s.sol";
import "script/utils/SafeTxUtils.sol";

/// @notice Deployment data struct
struct Deployment {
    string name;
    address deployedTo;
}

/// TODO: break all abstract contracts out into their own file in a `template` directory

/// @notice template for an EOA script
abstract contract EOADeployer is ConfigParser {
    Deployment[] internal _deployments;

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

abstract contract QueueBuilder is ConfigParser {

        using MultisigCallUtils for MultisigCall[];
        using SafeTxUtils for SafeTx;

        MultisigCall[] _executorCalls;

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

        /// @notice to be implemented by inheriting contract
        function _queue(Addresses memory addrs, Environment memory env, Params memory params) public virtual returns (MultisigCall[] memory);

        /// @notice helper function to create calldata for executor
        /// can be used for queue or execute
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
}

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