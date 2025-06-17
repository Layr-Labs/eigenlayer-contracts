// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IBN254TableCalculator.sol";
import "../interfaces/IKeyRegistrar.sol";
import "../libraries/Merkle.sol";
import "../libraries/BN254.sol";

/**
 * @title BN254TableCalculatorBase
 * @notice Abstract contract that provides base functionality for calculating BN254 operator tables
 * @dev This contract contains all the core logic for operator table calculations,
 *      with weight calculation left to be implemented by derived contracts
 */
abstract contract BN254TableCalculatorBase is IBN254TableCalculator {
    using Merkle for bytes32[];
    using BN254 for BN254.G1Point;

    // Immutables
    /// @notice KeyRegistrar contract for managing operator keys
    IKeyRegistrar public immutable keyRegistrar;

    constructor(
        IKeyRegistrar _keyRegistrar
    ) {
        keyRegistrar = _keyRegistrar;
    }

    /// @inheritdoc IBN254TableCalculator
    function calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) external view virtual returns (BN254OperatorSetInfo memory operatorSetInfo) {
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

    /// @inheritdoc IBN254TableCalculator
    function getOperatorInfos(
        OperatorSet calldata operatorSet
    ) external view virtual returns (BN254OperatorInfo[] memory) {
        // Get the weights for all operators
        (address[] memory operators, uint256[][] memory weights) = _getOperatorWeights(operatorSet);

        BN254OperatorInfo[] memory operatorInfos = new BN254OperatorInfo[](operators.length);

        for (uint256 i = 0; i < operators.length; i++) {
            // Skip if the operator has not registered their key
            if (!keyRegistrar.isRegistered(operatorSet, operators[i])) {
                continue;
            }

            (BN254.G1Point memory g1Point,) = keyRegistrar.getBN254Key(operatorSet, operators[i]);
            operatorInfos[i] = BN254OperatorInfo({pubkey: g1Point, weights: weights[i]});
        }

        return operatorInfos;
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
     * @notice Calculates the operator table for a given operatorSet, also calculates the aggregate pubkey for the operatorSet
     * @param operatorSet The operatorSet to calculate the operator table for
     * @return operatorSetInfo The operator table for the given operatorSet
     * @dev This function:
     * 1. Gets operator weights from the weight calculator
     * 2. Collates weights into total weights
     * 3. Creates a merkle tree of operator info
     *    - assumes that the operator has a registered BN254 key
     * 4. Calculates the aggregate public key
     */
    function _calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) internal view returns (BN254OperatorSetInfo memory operatorSetInfo) {
        // Get the weights for all operators in the operatorSet
        (address[] memory operators, uint256[][] memory weights) = _getOperatorWeights(operatorSet);

        // If there are no weights, return an empty operator set info
        if (weights.length == 0) {
            return BN254OperatorSetInfo({
                operatorInfoTreeRoot: bytes32(0),
                numOperators: 0,
                aggregatePubkey: BN254.G1Point(0, 0),
                totalWeights: new uint256[](0)
            });
        }

        // Initialize arrays
        uint256 subArrayLength = weights[0].length;
        uint256[] memory totalWeights = new uint256[](subArrayLength);
        bytes32[] memory operatorInfoLeaves = new bytes32[](operators.length);
        BN254.G1Point memory aggregatePubkey;
        uint256 operatorCount = 0;

        for (uint256 i = 0; i < operators.length; i++) {
            // Skip if the operator has not registered their key
            if (!keyRegistrar.isRegistered(operatorSet, operators[i])) {
                continue;
            }

            // Read the weights for the operator and encode them into the operatorInfoLeaves
            // for all weights, add them to the total weights. The ith index returns the weights array for the ith operator
            for (uint256 j = 0; j < subArrayLength; j++) {
                totalWeights[j] += weights[i][j];
            }
            (BN254.G1Point memory g1Point,) = keyRegistrar.getBN254Key(operatorSet, operators[i]);
            operatorInfoLeaves[i] = keccak256(abi.encode(BN254OperatorInfo({pubkey: g1Point, weights: weights[i]})));

            // Add the operator's G1 point to the aggregate pubkey
            aggregatePubkey = aggregatePubkey.plus(g1Point);

            // Increment the operator count
            operatorCount++;
        }

        // If there are no operators, return an empty operator set info
        if (operatorCount == 0) {
            return BN254OperatorSetInfo({
                operatorInfoTreeRoot: bytes32(0),
                numOperators: 0,
                aggregatePubkey: BN254.G1Point(0, 0),
                totalWeights: new uint256[](0)
            });
        }

        // Resize the operatorInfoLeaves array to the number of operators and merkleize
        assembly {
            mstore(operatorInfoLeaves, operatorCount)
        }

        bytes32 operatorInfoTreeRoot = operatorInfoLeaves.merkleizeKeccak();

        return BN254OperatorSetInfo({
            operatorInfoTreeRoot: operatorInfoTreeRoot,
            numOperators: operatorCount,
            aggregatePubkey: aggregatePubkey,
            totalWeights: totalWeights
        });
    }
}
