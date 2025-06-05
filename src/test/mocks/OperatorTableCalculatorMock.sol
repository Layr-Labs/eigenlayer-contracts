// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IOperatorTableCalculator.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract OperatorTableCalculatorMock is IOperatorTableCalculator {
    using OperatorSetLib for OperatorSet;

    mapping(bytes32 => bytes) internal _operatorTableBytes;

    function calculateOperatorTableBytes(OperatorSet memory operatorSet) external view returns (bytes memory) {
        return _operatorTableBytes[operatorSet.key()];
    }

    function setOperatorTableBytes(OperatorSet memory operatorSet, bytes memory operatorTableBytes) external {
        _operatorTableBytes[operatorSet.key()] = operatorTableBytes;
    }
} 