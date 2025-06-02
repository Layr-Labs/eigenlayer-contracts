// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../libraries/OperatorSetLib.sol";

/**
 * @title IOperatorWeightCalculator
 * @notice A contract that calculates the weights for all operators in a given operatorSet
 * @notice This is a base interface to send operator weights to the `IOperatorTableCalculator`
 * @notice We separate this out to be able to plug-in different operator weight calculators for each operatorSet
 */
interface IOperatorWeightCalculator {
    /**
     * @notice Get the weights for all operators in a given operatorSet
     * @param operatorSet The operatorSet to get the weights for
     * @return operators The addresses of the operators in the operatorSet
     * @return weights The weights for each operator in the operatorSet
     */
    function getOperatorWeights(
        OperatorSet calldata operatorSet
    ) external view returns (address[] memory operators, uint256[][] memory weights);
}
