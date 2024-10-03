// SPDX-License-Identifier: UNLICENSED
pragma solidity =0.8.12;

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// script that calculates calldata input to provide to timelock to call the executor multisig and execute an action.
// in this particular case, it encodes an action where the executor multisig will call the community multisig with no data and a value of 1e9 (i.e. one gwei)

interface ISafe {
    /// @dev Allows to execute a Safe transaction confirmed by required number of owners and then pays the account that submitted the transaction.
    ///      Note: The fees are always transferred, even if the user transaction fails.
    /// @param to Destination address of Safe transaction.
    /// @param value Ether value of Safe transaction.
    /// @param data Data payload of Safe transaction.
    /// @param operation Operation type of Safe transaction.
    /// @param safeTxGas Gas that should be used for the Safe transaction.
    /// @param baseGas Gas costs that are independent of the transaction execution(e.g. base transaction fee, signature check, payment of the refund)
    /// @param gasPrice Gas price that should be used for the payment calculation.
    /// @param gasToken Token address (or 0 if ETH) that is used for the payment.
    /// @param refundReceiver Address of receiver of gas payment (or 0 if tx.origin).
    /// @param signatures Packed signature data ({bytes32 r}{bytes32 s}{uint8 v})
    function execTransaction(
        address to,
        uint256 value,
        bytes calldata data,
        uint8 operation,
        uint256 safeTxGas,
        uint256 baseGas,
        uint256 gasPrice,
        address gasToken,
        address payable refundReceiver,
        bytes memory signatures
    ) external;
}

interface ITimelock {
    function queueTransaction(address target, uint value, string memory signature, bytes memory data, uint eta) external returns (bytes32);
        /*
        {
        require(msg.sender == admin, "Timelock::queueTransaction: Call must come from admin.");
        require(eta >= getBlockTimestamp().add(delay), "Timelock::queueTransaction: Estimated execution block must satisfy delay.");

        bytes32 txHash = keccak256(abi.encode(target, value, signature, data, eta));
        queuedTransactions[txHash] = true;

        emit QueueTransaction(txHash, target, value, signature, data, eta);
        return txHash;
        }
        */

    function executeTransaction(address target, uint value, string memory signature, bytes memory data, uint eta) external payable returns (bytes memory);
}

contract Enum {
    enum Operation {Call, DelegateCall}
}

contract EncodeSafeTransactionMainnet is Test {
    // CALLDATA FOR CALL FROM TIMELOCK TO EXECUTOR MULTISIG
    // regular call
    uint256 safeTxGas = 0;
    uint256 baseGas = 0;
    uint256 gasPrice = 0;
    address gasToken = address(uint160(0));
    address payable refundReceiver = payable(address(uint160(0)));

    // CALDATA FOR CALL TO TIMELOCK
    address timelockTarget = address(0x1); // NOTE: dummy value
    uint256 timelockValue = 0;
    // empty string, just encode all the data in 'timelockData'
    string timelockSignature;

    function encodeForTimelock(
        address to,
        uint256 value,
        bytes memory data,
        Enum.Operation operation,
        uint256 timelockEta
    ) public {

        address timelock = address(0x0); // NOTE: dummy value

        bytes memory final_calldata_timelock_to_executor_multisig = encodeForExecutor(timelock, to, value, data, operation);

        emit log_named_bytes("final_calldata_timelock_to_executor_multisig", final_calldata_timelock_to_executor_multisig);

        bytes memory calldata_to_timelock = abi.encodeWithSelector(ITimelock.queueTransaction.selector,
            timelockTarget,
            timelockValue,
            timelockSignature,
            final_calldata_timelock_to_executor_multisig,
            timelockEta
            );

        emit log_named_bytes("calldata_to_timelock", calldata_to_timelock);
    }

    function encodeForExecutor(
        address from,
        address to,
        uint256 value,
        bytes memory data,
        Enum.Operation operation
    ) public returns (bytes memory) {
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