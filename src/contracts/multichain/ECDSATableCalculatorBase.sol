// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IECDSATableCalculator.sol";
import "../interfaces/IKeyRegistrar.sol";
import "../libraries/Merkle.sol";

/**
 * @title ECDSATableCalculatorBase
 * @notice Abstract contract that provides base functionality for calculating ECDSA operator tables
 * @dev This contract contains all the core logic for operator table calculations,
 *      with weight calculation left to be implemented by derived contracts
 */
abstract contract ECDSATableCalculatorBase is IECDSATableCalculator {
    using Merkle for bytes32[];

    // Immutables
    /// @notice KeyRegistrar contract for managing operator keys
    IKeyRegistrar public immutable keyRegistrar;

    constructor(
        IKeyRegistrar _keyRegistrar
    ) {
        keyRegistrar = _keyRegistrar;
    }

    /// @inheritdoc IECDSATableCalculator
    function calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) external view virtual returns (ECDSAOperatorInfo[] memory operatorInfos) {
        return _calculateOperatorTable(operatorSet);
    }

    /// @inheritdoc IOperatorTableCalculator
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view virtual returns (bytes memory operatorTableBytes) {
        return abi.encode(_calculateOperatorTable(operatorSet));
    }

    /// @inheritdoc IOperatorTableCalculator
    function getOperatorWeights(
        OperatorSet calldata operatorSet
    ) external view virtual returns (address[] memory operators, uint256[][] memory weights) {
        return _getOperatorWeights(operatorSet);
    }

    /// @inheritdoc IOperatorTableCalculator
    function getOperatorWeight(
        OperatorSet calldata operatorSet,
        address operator
    ) external view virtual returns (uint256 weight) {
        (address[] memory operators, uint256[][] memory weights) = _getOperatorWeights(operatorSet);

        // Find the index of the operator in the operators array
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                return weights[i][0];
            }
        }

        return 0;
    }

    /**
     * @notice Abstract function to get the operator weights for a given operatorSet
     * @param operatorSet The operatorSet to get the weights for
     * @return operators The addresses of the operators in the operatorSet
     * @return weights The weights for each operator in the operatorSet, this is a 2D array where the first index is the operator
     * and the second index is the type of weight
     * @dev Must be implemented by derived contracts to define specific weight calculation logic
     */
    function _getOperatorWeights(
        OperatorSet calldata operatorSet
    ) internal view virtual returns (address[] memory operators, uint256[][] memory weights);

    /**
     * @notice Calculates the operator table for a given operatorSet
     * @param operatorSet The operatorSet to calculate the operator table for
     * @return operatorInfos The operator table for the given operatorSet
     * @dev This function:
     * 1. Gets operator weights from the weight calculator
     * 2. Creates ECDSAOperatorInfo structs for each operator with registered ECDSA keys
     */
    function _calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) internal view returns (ECDSAOperatorInfo[] memory operatorInfos) {
        // Get the weights for all operators in the operatorSet
        (address[] memory operators, uint256[][] memory weights) = _getOperatorWeights(operatorSet);

        // If there are no weights, return an empty array
        if (weights.length == 0) {
            return new ECDSAOperatorInfo[](0);
        }

        // Create the operator infos array with maximum possible size
        operatorInfos = new ECDSAOperatorInfo[](operators.length);
        uint256 operatorCount = 0;

        for (uint256 i = 0; i < operators.length; i++) {
            // Skip if the operator has not registered their ECDSA key
            if (!keyRegistrar.isRegistered(operatorSet, operators[i])) {
                continue;
            }

            // Get the ECDSA address (public key) for the operator
            address ecdsaAddress = keyRegistrar.getECDSAAddress(operatorSet, operators[i]);

            // Create the ECDSAOperatorInfo struct
            operatorInfos[operatorCount] = ECDSAOperatorInfo({pubkey: ecdsaAddress, weights: weights[i]});

            operatorCount++;
        }

        // If no operators have registered keys, return empty array
        if (operatorCount == 0) {
            return new ECDSAOperatorInfo[](0);
        }

        // Resize the array to the actual number of operators with registered keys
        assembly {
            mstore(operatorInfos, operatorCount)
        }

        return operatorInfos;
    }
}
