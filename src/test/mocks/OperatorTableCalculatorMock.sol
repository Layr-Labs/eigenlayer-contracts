// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IOperatorTableCalculator.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract OperatorTableCalculatorMock is IOperatorTableCalculator {
    using OperatorSetLib for OperatorSet;

    mapping(bytes32 => bytes) internal _operatorTableBytes;
    mapping(bytes32 => address[]) internal _operators;
    mapping(bytes32 => mapping(address => uint)) internal _operatorWeights;

    function calculateOperatorTableBytes(OperatorSet memory operatorSet) external view returns (bytes memory) {
        return _operatorTableBytes[operatorSet.key()];
    }

    function setOperatorTableBytes(OperatorSet memory operatorSet, bytes memory operatorTableBytes) external {
        _operatorTableBytes[operatorSet.key()] = operatorTableBytes;
    }

    function getOperatorSetWeights(OperatorSet calldata operatorSet)
        external
        view
        returns (address[] memory operators, uint[][] memory weights)
    {
        bytes32 key = operatorSet.key();
        operators = _operators[key];

        weights = new uint[][](operators.length);
        for (uint i = 0; i < operators.length; i++) {
            weights[i] = new uint[](1);
            weights[i][0] = _operatorWeights[key][operators[i]];
        }

        return (operators, weights);
    }

    function getOperatorWeights(OperatorSet calldata operatorSet, address operator) external view returns (uint[] memory weights) {
        weights = new uint[](1);
        weights[0] = _operatorWeights[operatorSet.key()][operator];
        return weights;
    }

    // Helper functions for testing
    function setOperators(OperatorSet memory operatorSet, address[] memory operators) external {
        _operators[operatorSet.key()] = operators;
    }

    function setOperatorWeight(OperatorSet memory operatorSet, address operator, uint weight) external {
        _operatorWeights[operatorSet.key()][operator] = weight;
    }
}
