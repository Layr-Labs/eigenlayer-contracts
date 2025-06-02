// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/multichain/IECDSATableCalculator.sol";
import "../interfaces/multichain/IOperatorWeightCalculator.sol";
import "../interfaces/IKeyRegistrar.sol";

/**
 * @title Storage variables for the `ECDSATableCalculator` contract.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract ECDSATableCalculatorStorage is IECDSATableCalculator {
    // Immutables
    /// @notice OperatorWeightCalculator contract for calculating operator weights
    IOperatorWeightCalculator public immutable operatorWeightCalculator;

    /// @notice KeyRegistrar contract for managing operator keys
    IKeyRegistrar public immutable keyRegistrar;

    constructor(IOperatorWeightCalculator _operatorWeightCalculator, IKeyRegistrar _keyRegistrar) {
        operatorWeightCalculator = _operatorWeightCalculator;
        keyRegistrar = _keyRegistrar;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
