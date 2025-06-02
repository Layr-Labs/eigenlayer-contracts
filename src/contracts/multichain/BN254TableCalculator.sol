// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "../interfaces/multichain/IBN254TableCalculator.sol";
import "../interfaces/multichain/IOperatorWeightCalculator.sol";
import "../interfaces/IKeyRegistrar.sol";
import "../libraries/Merkle.sol";
import "../libraries/BN254.sol";
import "./BN254TableCalculatorStorage.sol";

/**
 * @title BN254TableCalculator
 * @notice Contract that calculates BN254 operator tables for a given operatorSet
 * @dev This contract uses an IOperatorWeightCalculator to get operator weights and formats them into the required table structure
 */
contract BN254TableCalculator is Initializable, OwnableUpgradeable, BN254TableCalculatorStorage {
    using Merkle for bytes32[];
    using BN254 for BN254.G1Point;

    constructor(
        IOperatorWeightCalculator _operatorWeightCalculator,
        IKeyRegistrar _keyRegistrar
    ) BN254TableCalculatorStorage(_operatorWeightCalculator, _keyRegistrar) {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract.
     */
    function initialize(
        address initialOwner
    ) external initializer {
        _transferOwnership(initialOwner);
    }

    /// @inheritdoc IBN254TableCalculator
    function calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) public view virtual returns (BN254OperatorSetInfo memory operatorSetInfo) {
        return _calculateOperatorTable(operatorSet);
    }

    /// @inheritdoc IBN254TableCalculator
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view virtual returns (bytes memory operatorTableBytes) {
        return abi.encode(_calculateOperatorTable(operatorSet));
    }

    /// @inheritdoc IBN254TableCalculator
    function getOperatorInfos(
        OperatorSet calldata operatorSet
    ) external view virtual returns (BN254OperatorInfo[] memory) {
        // Get the weights for all operators
        (address[] memory operators, uint256[][] memory weights) =
            operatorWeightCalculator.getOperatorWeights(operatorSet);

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
        (address[] memory operators, uint256[][] memory weights) =
            operatorWeightCalculator.getOperatorWeights(operatorSet);

        // Collate weights into a single array of total weights
        uint256 subArrayLength = weights[0].length;
        uint256[] memory totalWeights = new uint256[](subArrayLength);
        bytes32[] memory operatorInfoLeaves = new bytes32[](operators.length);
        BN254.G1Point memory aggregatePubkey;

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
        }

        bytes32 operatorInfoTreeRoot = operatorInfoLeaves.merkleizeKeccak();

        return BN254OperatorSetInfo({
            operatorInfoTreeRoot: operatorInfoTreeRoot,
            numOperators: operators.length,
            aggregatePubkey: aggregatePubkey,
            totalWeights: totalWeights
        });
    }
}
