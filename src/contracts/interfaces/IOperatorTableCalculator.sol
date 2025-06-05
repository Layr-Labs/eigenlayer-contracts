// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../libraries/OperatorSetLib.sol";

/// @notice A base operator table calculator that all operator table calculators (ECDSA, BN254) must implement
interface IOperatorTableCalculator {
    /**
     * @notice Calculates the operator table for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return operatorTableBytes the operatorTableBytes for the given operatorSet
     */
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view returns (bytes memory operatorTableBytes);

    /**
     * @notice Get the operator weights for a given operatorSet based on the slashable stake.
     * @param operatorSet The operatorSet to get the weights for
     * @return operators The addresses of the operators in the operatorSet
     * @return weights The weights for each operator in the operatorSet, this is a 2D array where the first index is the operator
     * and the second index is the type of weight. In this case its of length 1 and returns the slashable stake for the operatorSet.
     */
    function getOperatorWeights(
        OperatorSet calldata operatorSet
    ) external view returns (address[] memory operators, uint256[][] memory weights);

    /**
     * @notice Get the weight for a given operator in a given operatorSet based on the slashable stake.
     * @param operatorSet The operatorSet to get the weight for
     * @param operator The operator to get the weight for
     * @return weight The weight for the operator in the operatorSet
     */
    function getOperatorWeight(
        OperatorSet calldata operatorSet,
        address operator
    ) external view returns (uint256 weight);
}
