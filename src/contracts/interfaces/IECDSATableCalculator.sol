// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/libraries/OperatorSetLib.sol";

interface IECDSATableCalculatorTypes {
    /**
     * @notice A struct that contains information about a single operator
     * @param pubkey The address of the operator
     * @param weights The weights of the operator for a single operatorSet
     * @dev The `weights` array can be defined as a list of arbitrary groupings. For example,
     * it can be [slashable_stake, delegated_stake, strategy_i_stake, ...]
     */
    struct ECDSAOperatorInfo {
        address pubkey;
        uint96[] weights;
    }
}

interface IECDSATableCalculator is IECDSATableCalculatorTypes {
    /**
     * @notice calculates the operatorInfos for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return operatorInfos the list of operatorInfos for the given operatorSet
     * @dev The output of this function is converted to bytes via the `calculateOperatorTableBytes` function
     */
    function calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) external view returns (ECDSAOperatorInfo[] memory operatorInfos);

    /**
     * @notice calculates the operatorInfos for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return operatorTableBytes the operatorTableBytes for the given operatorSet
     */
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view returns (bytes memory operatorTableBytes);
}
