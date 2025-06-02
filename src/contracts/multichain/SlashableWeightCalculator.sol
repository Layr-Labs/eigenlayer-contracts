// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";
import "../interfaces/IStrategy.sol";
import {
    IOperatorWeightCalculator,
    OperatorSet,
    OperatorSetLib
} from "../interfaces/multichain/IOperatorWeightCalculator.sol";
import "./SlashableWeightCalculatorStorage.sol";

/**
 * @title SlashableWeightCalculator
 * @notice A contract that calculates the operator weights for a given operatorSet
 * @dev This contract assumes that all operator stakes are stored in the `AllocationManager`
 */
contract SlashableWeightCalculator is Initializable, OwnableUpgradeable, SlashableWeightCalculatorStorage {
    using OperatorSetLib for OperatorSet;
    using Math for uint256;

    error LookaheadBlocksTooHigh();
    error ArrayLengthMismatch();

    /// @notice Emitted when a strategy multiplier is set
    event StrategyMultiplierSet(IStrategy indexed strategy, uint256 multiplier);
    /// @notice Emitted when the lookahead blocks are set
    event LookaheadBlocksSet(uint256 lookaheadBlocks);

    constructor(
        IAllocationManager _allocationManager
    ) SlashableWeightCalculatorStorage(_allocationManager) {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract
     * @param initialOwner The owner of the contract
     */
    function initialize(
        address initialOwner
    ) external initializer {
        _transferOwnership(initialOwner);
    }

    /**
     * @notice Get the operator weights for a given operatorSet based on the slashable stake.
     * @param operatorSet The operatorSet to get the weights for
     * @return operators The addresses of the operators in the operatorSet
     * @return weights The weights for each operator in the operatorSet, this is a 2D array where the first index is the operator
     * and the second index is the type of weight. In this case its of length 1 and returns the slashable stake for the operatorSet.
     */
    function getOperatorWeights(
        OperatorSet calldata operatorSet
    ) public view virtual returns (address[] memory operators, uint256[][] memory weights) {
        // Get all operators & strategies in the operatorSet
        operators = allocationManager.getMembers(operatorSet);
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);

        // Get the minimum slashable stake for each operator
        uint256[][] memory minSlashableStake = allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: operators,
            strategies: strategies,
            futureBlock: uint32(block.number + lookaheadBlocks)
        });

        // Cache the operatorSetKey for weights
        bytes32 operatorSetKey = operatorSet.key();

        weights = new uint256[][](operators.length);
        for (uint256 operatorIndex = 0; operatorIndex < operators.length; ++operatorIndex) {
            // Initialize operator weights array of length 1 just for slashable stake
            weights[operatorIndex] = new uint256[](1);
            // 1. For the given operator, loop through the strategies and calculate the operator's weight for the operatorSet
            for (uint256 stratIndex = 0; stratIndex < strategies.length; ++stratIndex) {
                // Update the weight for the operator and strategy, only if there's a nonzero minimum slashable stake
                if (minSlashableStake[operatorIndex][stratIndex] > 0) {
                    // Get the multiplier for the strategy
                    uint256 multiplier = _multipliers[operatorSetKey][strategies[stratIndex]];

                    // We're only returning the weights of slashable stake in this calculator, hence the 0 index
                    weights[operatorIndex][0] +=
                        minSlashableStake[operatorIndex][stratIndex].mulDiv(multiplier, WEIGHTING_DIVISOR);
                }
            }
        }

        return (operators, weights);
    }

    /**
     * @notice Get the weight for a given operator in a given operatorSet based on the slashable stake.
     * @param operatorSet The operatorSet to get the weight for
     * @param operator The operator to get the weight for
     * @return weight The weight for the operator in the operatorSet
     */
    function getOperatorWeight(
        OperatorSet calldata operatorSet,
        address operator
    ) public view virtual returns (uint256 weight) {
        (address[] memory operators, uint256[][] memory weights) = getOperatorWeights(operatorSet);

        // Find the index of the operator in the operators array
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                return weights[i][0];
            }
        }

        return 0;
    }

    /**
     * @notice Set the multipliers for each strategy. Note that the strategy doesn't necessarily need to be in the operatorSet
     * to be able to be set here.
     * @param operatorSet The operatorSet to set the multipliers for
     * @param strategies The strategies to set the multipliers for
     * @param multipliers The multipliers to set for the strategies
     */
    function setStrategyMultipliers(
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies,
        uint256[] calldata multipliers
    ) external onlyOwner {
        // Validate the lengths of strategies and multipliers
        require(strategies.length == multipliers.length, ArrayLengthMismatch());

        // Get the key for the operatorSet
        bytes32 operatorSetKey = operatorSet.key();

        for (uint256 i = 0; i < strategies.length; i++) {
            _multipliers[operatorSetKey][strategies[i]] = multipliers[i];
            emit StrategyMultiplierSet(strategies[i], multipliers[i]);
        }
    }

    /**
     * @notice Set the lookahead blocks for the slashable stake calculation
     * @param _lookaheadBlocks The lookahead blocks to set
     */
    function setLookaheadBlocks(
        uint256 _lookaheadBlocks
    ) external onlyOwner {
        require(_lookaheadBlocks < allocationManager.DEALLOCATION_DELAY(), LookaheadBlocksTooHigh());
        lookaheadBlocks = _lookaheadBlocks;
        emit LookaheadBlocksSet(_lookaheadBlocks);
    }
}
