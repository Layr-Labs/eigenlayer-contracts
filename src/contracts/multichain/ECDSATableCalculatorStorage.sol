// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/multichain/IECDSATableCalculator.sol";
import "../interfaces/IKeyRegistrar.sol";
import "../interfaces/IAllocationManager.sol";

/**
 * @title Storage variables for the `ECDSATableCalculator` contract.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract ECDSATableCalculatorStorage is IECDSATableCalculator {
    // Immutables
    /// @notice KeyRegistrar contract for managing operator keys
    IKeyRegistrar public immutable keyRegistrar;
    /// @notice AllocationManager contract for managing operator allocations
    IAllocationManager public immutable allocationManager;

    // Storage variables
    /// @notice The lookahead blocks for the slashable stake calculation
    uint256 public lookaheadBlocks;

    constructor(IKeyRegistrar _keyRegistrar, IAllocationManager _allocationManager) {
        keyRegistrar = _keyRegistrar;
        allocationManager = _allocationManager;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
