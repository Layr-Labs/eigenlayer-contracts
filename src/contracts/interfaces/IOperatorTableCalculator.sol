// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";

/// @notice A base operator table calculator that all operator table calculators (ECDSA, BN254) must implement
interface IOperatorTableCalculator {
    /**
     * @notice Calculates the operatorTableBytes for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return operatorTableBytes the operatorTableBytes for the given operatorSet
     */
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view returns (bytes memory operatorTableBytes);
}
