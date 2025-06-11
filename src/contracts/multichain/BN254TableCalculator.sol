// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./BN254TableCalculatorBase.sol";
import "../interfaces/IAllocationManager.sol";

/**
 * @title BN254TableCalculator
 * @notice Implementation that calculates BN254 operator tables using the sum of the minimum slashable stake weights
 */
contract BN254TableCalculator is BN254TableCalculatorBase {
    // Immutables
    /// @notice AllocationManager contract for managing operator allocations
    IAllocationManager public immutable allocationManager;
    /// @notice The default lookahead blocks for the slashable stake lookup
    uint256 public immutable LOOKAHEAD_BLOCKS;

    constructor(
        IKeyRegistrar _keyRegistrar,
        IAllocationManager _allocationManager,
        uint256 _LOOKAHEAD_BLOCKS
    ) BN254TableCalculatorBase(_keyRegistrar) {
        allocationManager = _allocationManager;
        LOOKAHEAD_BLOCKS = _LOOKAHEAD_BLOCKS;
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
    ) internal view override returns (address[] memory operators, uint256[][] memory weights) {
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
}
