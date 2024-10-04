// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ConfigParser.sol";
import "script/utils/Encoders.sol";
import "script/utils/Interfaces.sol";

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

/// @dev from EncodeSafeMultisendMainnet
struct MultisigCall {
    address to;
    uint256 value;
    bytes data;
}

library MultisigCallHelper {

    function append(
        MultisigCall[] storage multisigCalls,
        address to,
        uint256 value,
        bytes memory data
    ) internal returns (MultisigCall[] storage) {
        multisigCalls.push(MultisigCall({
            to: to,
            value: value,
            data: data
        }));

        return multisigCalls;
    }

    /// @notice appends a multisig call with a value of 0
    function append(
        MultisigCall[] storage multisigCalls,
        address to,
        bytes memory data
    ) internal returns (MultisigCall[] storage) {
        multisigCalls.push(MultisigCall({
            to: to,
            value: 0,
            data: data
        }));

        return multisigCalls;
    }

    function encodeMultisendTxs(MultisigCall[] memory txs) public pure returns (bytes memory) {
        bytes memory ret = new bytes(0);
        for (uint256 i = 0; i < txs.length; i++) {
            ret = abi.encodePacked(
                ret,
                abi.encodePacked(
                    uint8(0),
                    txs[i].to,
                    txs[i].value,
                    uint256(txs[i].data.length),
                    txs[i].data
                )
            );
        }


        return abi.encodeWithSelector(IMultiSend.multiSend.selector, ret);
    }
}

library TransactionHelper {
    function encodeForExecutor(
        Transaction memory t
    ) public returns (bytes memory) {
        bytes1 v = bytes1(uint8(1));
        bytes32 r = bytes32(uint256(uint160(0xA6Db1A8C5a981d1536266D2a393c5F8dDb210EAF))); // HARDCODED MAINNET TIMELOCK
        bytes32 s;
        bytes memory sig = abi.encodePacked(r,s,v);

        bytes memory final_calldata_to_executor_multisig = abi.encodeWithSelector(
            ISafe.execTransaction.selector,
            t.to,
            t.value,
            t.data,
            t.op,
            0, // safeTxGas
            0, // baseGas
            0, // gasPrice
            address(uint160(0)), // gasToken
            payable(address(uint160(0))), // refundReceiver
            sig
        );

        return final_calldata_to_executor_multisig;
    }
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
abstract contract MultisigBuilder is ConfigParser {

    using MultisigCallHelper for *;

    /// @return a Transaction object for a Gnosis Safe to ingest
    function execute(string memory envPath) public returns (Transaction memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        MultisigCall[] memory calls = _execute(addrs, env, params);

        bytes memory data = calls.encodeMultisendTxs();

        return Transaction({
            to: params.multiSendCallOnly,
            value: 0,
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

        execute(envPath);
        _testExecute(
            addrs,
            env,
            params
        );
    }

    /// @notice to be implemented by inheriting contract
    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (MultisigCall[] memory);

    /// @notice to be implemented for testing
    function _testExecute(Addresses memory addrs, Environment memory env, Params memory params) internal virtual;
}

/// @notice template for an OpsMultisig script that goes through the timelock
/// @dev writing a script is done from the perspective of the OpsMultisig
abstract contract OpsTimelockBuilder is MultisigBuilder {

    using MultisigCallHelper for *;
    using TransactionHelper for *;

    /// @return a Transaction object for a Gnosis Safe to ingest
    function queue(string memory envPath) public returns (Transaction memory) {
        (
            Addresses memory addrs,
            Environment memory env,
            Params memory params
        ) = _readConfigFile(envPath);

        MultisigCall[] memory calls = _queue(addrs, env, params);

        // ENCODE CALLS FOR EXECUTOR
        bytes memory data = calls.encodeMultisendTxs();

        bytes memory executorCalldata = Transaction({
            to: params.multiSendCallOnly,
            value: 0,
            data: data,
            op: EncGnosisSafe.Operation.DelegateCall
        }).encodeForExecutor();

        // ENCODE EXECUTOR CALLDATA FOR TIMELOCK
        bytes memory timelockCalldata = abi.encodeWithSelector(
            ITimelock.queueTransaction.selector,
            addrs.executorMultisig,
            0,
            "",
            executorCalldata,
            type(uint256).max
        );

        // ENCODE TIMELOCK DATA FOR OPS MULTISIG
        return Transaction({
            to: addrs.timelock,
            value: 0,
            data: timelockCalldata,
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }

    function makeTimelockTx(Addresses memory addrs, Environment memory env, Params memory params) public returns (Transaction memory) {
        MultisigCall[] memory calls = _makeTimelockTx(addrs, env, params);

        bytes memory data = calls.encodeMultisendTxs();

        return Transaction({
            to: params.multiSendCallOnly,
            value: 0,
            data: data,
            op: EncGnosisSafe.Operation.DelegateCall
        });
    }

    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (MultisigCall[] memory);

    function _makeTimelockTx(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (MultisigCall[] memory);
}

/// @notice to be used for CommunityMultisig
/// NOTE: WIP
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