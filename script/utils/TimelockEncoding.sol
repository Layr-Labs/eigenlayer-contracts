// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Script.sol";
import "forge-std/Test.sol";

import "./TxEncodingInterfaces.sol";

contract TimelockEncoding is Test {
    // CALLDATA FOR CALL FROM TIMELOCK TO EXECUTOR MULTISIG
    uint256 safeTxGas = 0;
    uint256 baseGas = 0;
    uint256 gasPrice = 0;
    address gasToken = address(uint160(0));
    address payable refundReceiver = payable(address(uint160(0)));

    // CALDATA FOR CALL TO TIMELOCK
    uint256 timelockValue = 0;
    // empty string, just encode all the data in 'timelockData'
    string timelockSignature;

    // appropriate address on mainnet, Holesky, and many other chains
    address multiSendCallOnly = 0x40A2aCCbd92BCA938b02010E17A5b8929b49130D;

    function encodeForTimelock(
        address to,
        uint256 value,
        bytes memory data,
        uint256 timelockEta
    ) public returns (bytes memory calldata_to_timelock_queuing_action, bytes memory calldata_to_timelock_executing_action) {
        calldata_to_timelock_queuing_action = abi.encodeWithSelector(ITimelock.queueTransaction.selector,
            to,
            value,
            timelockSignature,
            data,
            timelockEta
        );

        emit log_named_bytes("calldata_to_timelock_queuing_action", calldata_to_timelock_queuing_action);

        calldata_to_timelock_executing_action = abi.encodeWithSelector(ITimelock.executeTransaction.selector,
            to,
            value,
            timelockSignature,
            data,
            timelockEta
        );

        emit log_named_bytes("calldata_to_timelock_executing_action", calldata_to_timelock_executing_action);

        return (calldata_to_timelock_queuing_action, calldata_to_timelock_executing_action);
    }

    function encodeForExecutor(
        address from,
        address to,
        uint256 value,
        bytes memory data,
        ISafe.Operation operation
    ) public returns (bytes memory) {
        // encode the "signature" required by the Safe
        bytes1 v = bytes1(uint8(1));
        bytes32 r = bytes32(uint256(uint160(from)));
        bytes32 s;
        bytes memory sig = abi.encodePacked(r,s,v);
        emit log_named_bytes("sig", sig);

        bytes memory final_calldata_to_executor_multisig = abi.encodeWithSelector(ISafe.execTransaction.selector,
            to,
            value,
            data,
            operation,
            safeTxGas,
            baseGas,
            gasPrice,
            gasToken,
            refundReceiver,
            sig
            );

        emit log_named_bytes("final_calldata_to_executor_multisig", final_calldata_to_executor_multisig);

        return final_calldata_to_executor_multisig;
    }

    struct Tx {
        address to;
        uint256 value;
        bytes data;
    }

    function encodeMultisendTxs(Tx[] memory txs) public pure returns (bytes memory) {
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
        return ret;
    }
}
