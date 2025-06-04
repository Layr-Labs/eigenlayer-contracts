// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../libraries/OperatorSetLib.sol";
import "./IOperatorTableCalculator.sol";

interface IECDSATableCalculatorTypes {
    /**
     * @notice A struct that contains information about a single operator
     * @param pubkey The address of the signing ECDSA key of the operator and not the operator address itself.
     * This is read from the KeyRegistrar contract.
     * @param weights The weights of the operator for a single operatorSet
     * @dev The `weights` array can be defined as a list of arbitrary groupings. For example,
     * it can be [slashable_stake, delegated_stake, strategy_i_stake, ...]
     */
    struct ECDSAOperatorInfo {
        address pubkey;
        uint256[] weights;
    }
}

interface IECDSATableCalculatorEvents {
    /// @notice Emitted when the lookahead blocks are set
    event LookaheadBlocksSet(uint256 lookaheadBlocks);
}

interface IECDSATableCalculatorErrors {
    /// @notice Emitted when the lookahead blocks are too high
    error LookaheadBlocksTooHigh();
}

interface IECDSATableCalculator is
    IOperatorTableCalculator,
    IECDSATableCalculatorTypes,
    IECDSATableCalculatorEvents,
    IECDSATableCalculatorErrors
{
    /**
     * @notice calculates the operatorInfos for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return operatorInfos the list of operatorInfos for the given operatorSet
     * @dev The output of this function is converted to bytes via the `calculateOperatorTableBytes` function
     */
    function calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) external view returns (ECDSAOperatorInfo[] memory operatorInfos);
}
