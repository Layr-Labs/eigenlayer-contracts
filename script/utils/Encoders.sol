// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "./Interfaces.sol";
import "./TxBuilder.sol";

library EncProxyAdmin {

    function upgrade(address proxy, address implementation) internal pure returns (bytes memory) {
        return abi.encodeWithSelector(IProxyAdmin.upgrade.selector, proxy, implementation);
    }
}

library EncUpgradeableBeacon {

    function upgradeTo(address implementation) internal pure returns (bytes memory) {
        return abi.encodeWithSelector(IUpgradeableBeacon.upgradeTo.selector, implementation);
    }
}

library EncTimelock {

    uint constant TIMELOCK_VALUE = 0;
    string constant TIMELOCK_SIGNATURE = "";

    function queueTransaction(address target, uint eta, bytes memory data) internal pure returns (bytes memory) {
        return abi.encodeWithSelector(ITimelock.queueTransaction.selector,
            target,
            TIMELOCK_VALUE,
            TIMELOCK_SIGNATURE,
            data,
            eta
        );
    }

    function executeTransaction(address target, uint eta, bytes memory data) internal pure returns (bytes memory) {
        return abi.encodeWithSelector(ITimelock.executeTransaction.selector,
            target,
            TIMELOCK_VALUE,
            TIMELOCK_SIGNATURE,
            data,
            eta
        );
    }

    function cancelTransaction(address target, uint eta, bytes memory data) internal pure returns (bytes memory) {
        return abi.encodeWithSelector(ITimelock.cancelTransaction.selector,
            target,
            TIMELOCK_VALUE,
            TIMELOCK_SIGNATURE,
            data,
            eta
        );
    }
}

library EncGnosisSafe {

    enum Operation {
        Call,
        DelegateCall
    }

    uint constant SAFE_TX_GAS = 0;
    uint constant BASE_GAS = 0;
    uint constant GAS_PRICE = 0;
    address constant GAS_TOKEN = address(uint160(0));
    address constant REFUND_RECEIVER = payable(address(uint160(0)));

    function execTransaction(
        address from,
        address to,
        bytes memory data,
        Operation op
    ) internal pure returns (bytes memory) {
        return encodeForExecutor(from, to, 0, data, op);
    }

    function execTransaction(
        address from,
        address to,
        uint value,
        bytes memory data,
        Operation op
    ) internal pure returns (bytes memory) {
        return encodeForExecutor(from, to, value, data, op);
    }

    function encodeForExecutor(
        address from,
        address to,
        uint value,
        bytes memory data,
        Operation op
    ) internal pure returns (bytes memory) {
        bytes1 v = bytes1(uint8(1));
        bytes32 r = bytes32(uint256(uint160(from)));
        bytes32 s;
        bytes memory sig = abi.encodePacked(r,s,v);

        bytes memory final_calldata_to_executor_multisig = abi.encodeWithSelector(ISafe.execTransaction.selector,
            to,
            value,
            data,
            op,
            SAFE_TX_GAS,
            BASE_GAS,
            GAS_PRICE,
            GAS_TOKEN,
            REFUND_RECEIVER,
            sig
        );

        return final_calldata_to_executor_multisig;
    }
}

library EncMultiSendCallOnly {

    using TxBuilder for *;
    
    function multiSend(Txs storage txs) internal view returns (bytes memory) {
        Tx[] memory arr = txs.toArray();

        bytes memory ret = new bytes(0);
        for (uint i = 0; i < arr.length; i++) {
            ret = abi.encodePacked(
                ret,
                abi.encodePacked(
                    uint8(0),
                    arr[i].to,
                    arr[i].value,
                    uint(arr[i].data.length),
                    arr[i].data
                )
            );
        }

        return abi.encodeWithSelector(IMultiSend.multiSend.selector, ret);
    }
}