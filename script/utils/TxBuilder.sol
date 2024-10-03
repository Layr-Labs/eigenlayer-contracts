// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

struct Tx {
    address to;
    uint256 value;
    bytes data;
}

struct Txs {
    uint txCount;
    mapping(uint => Tx) txns;
}

library TxBuilder {

    function append(Txs storage txs, address to, bytes memory data) internal returns (Txs storage) {
        txs.txns[txs.txCount] = Tx({
            to: to,
            value: 0,
            data: data
        });

        txs.txCount++;
        return txs;
    }

    function append(Txs storage txs, address to, uint value, bytes memory data) internal returns (Txs storage) {
        txs.txns[txs.txCount] = Tx({
            to: to,
            value: value,
            data: data
        });

        txs.txCount++;
        return txs;
    }

    function toArray(Txs storage txs) internal view returns (Tx[] memory) {
        Tx[] memory arr = new Tx[](txs.txCount);

        for (uint i = 0; i < arr.length; i++) {
            arr[i] = txs.txns[i];
        }

        return arr;
    }
}