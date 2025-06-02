// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IStrategy} from "../interfaces/IStrategy.sol";
import {IAllocationManager} from "../interfaces/IAllocationManager.sol";
import {IOperatorWeightCalculator} from "../interfaces/multichain/IOperatorWeightCalculator.sol";

/**
 * @title Storage variables for the `SlashableWeightCalculator` contract.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract SlashableWeightCalculatorStorage is IOperatorWeightCalculator {
    // Constants
    uint256 public constant WEIGHTING_DIVISOR = 1e18;

    // Immutables
    IAllocationManager public immutable allocationManager;

    // Storage variables
    /// @notice The multipliers for each strategy, per operatorSet
    mapping(bytes32 operatorSetKey => mapping(IStrategy => uint256)) internal _multipliers;

    /// @notice The lookahead blocks for the slashable stake calculation
    uint256 public lookaheadBlocks = 100_800;

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
