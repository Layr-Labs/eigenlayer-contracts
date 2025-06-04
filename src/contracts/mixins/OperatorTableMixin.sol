// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../"
abstract contract OperatorTableMixin {
    function getOperatorTableInfo(bytes calldata tableInfo) internal pure returns (OperatorSet memory operatorSet, KeyType keyType, OperatorSetConfig memory operatorSetInfo) {
        (operatorSet, keyType, operatorSetInfo) = abi.decode(tableInfo, (OperatorSet, KeyType, OperatorSetConfig));
    }
}