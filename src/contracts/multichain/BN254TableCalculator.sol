// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IBN254TableCalculator.sol";
import "../interfaces/IKeyRegistrar.sol";
import "../interfaces/IAllocationManager.sol";
import "../libraries/Merkle.sol";
import "../libraries/BN254.sol";

/**
 * @title BN254TableCalculator
 * @notice Contract that calculates BN254 operator tables for a given operatorSet
 * @dev This contract calculates operator weights and formats them into the required table structure
 */
contract BN254TableCalculator is IBN254TableCalculator {
    using Merkle for bytes32[];
    using BN254 for BN254.G1Point;

    // Immutables & Constants
    /// @notice KeyRegistrar contract for managing operator keys
    IKeyRegistrar public immutable keyRegistrar;
    /// @notice AllocationManager contract for managing operator allocations
    IAllocationManager public immutable allocationManager;
    /// @notice The default lookahead blocks for the slashable stake lookup
    uint256 public immutable LOOKAHEAD_BLOCKS;

    constructor(IKeyRegistrar _keyRegistrar, IAllocationManager _allocationManager, uint256 _LOOKAHEAD_BLOCKS) {
        keyRegistrar = _keyRegistrar;
        allocationManager = _allocationManager;
        LOOKAHEAD_BLOCKS = _LOOKAHEAD_BLOCKS;
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
     * @notice Get the operator weights for a given operatorSet based on the slashable stake.
     * @param operatorSet The operatorSet to get the weights for
     * @return operators The addresses of the operators in the operatorSet
     * @return weights The weights for each operator in the operatorSet, this is a 2D array where the first index is the operator
     * and the second index is the type of weight. In this case its of length 1 and returns the slashable stake for the operatorSet.
     */
    function _getOperatorWeights(
        OperatorSet calldata operatorSet
    ) internal view virtual returns (address[] memory operators, uint256[][] memory weights) {
        // Get all operators & strategies in the operatorSet
        operators = allocationManager.getMembers(operatorSet);
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);

        // Get the minimum slashable stake for each operator
        uint256[][] memory minSlashableStake = allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: operators,
            strategies: strategies,
            futureBlock: uint32(block.number + LOOKAHEAD_BLOCKS)
        });

        weights = new uint256[][](operators.length);
        for (uint256 operatorIndex = 0; operatorIndex < operators.length; ++operatorIndex) {
            // Initialize operator weights array of length 1 just for slashable stake
            weights[operatorIndex] = new uint256[](1);
            // 1. For the given operator, loop through the strategies and sum together to calculate the operator's weight for the operatorSet
            for (uint256 stratIndex = 0; stratIndex < strategies.length; ++stratIndex) {
                weights[operatorIndex][0] += minSlashableStake[operatorIndex][stratIndex];
            }
        }

        return (operators, weights);
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
        (address[] memory operators, uint256[][] memory weights) = _getOperatorWeights(operatorSet);

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
