// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IStrategyFactory.sol";

/// @title Storage for the StrategyFactory contract.
/// @author Layr Labs, Inc.
/// @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
abstract contract StrategyFactoryStorage is IStrategyFactory {
    /// @dev Deprecated: strategyBeacon is now immutable in StrategyFactory.
    /// Slot preserved for storage layout compatibility with existing proxies.
    uint256 private __deprecated_strategyBeacon;

    /// @notice Mapping token => Strategy contract for the token
    /// The strategies in this mapping are deployed by the StrategyFactory.
    /// The factory can only deploy a single strategy per token address
    /// These strategies MIGHT not be whitelisted in the StrategyManager,
    /// though deployNewStrategy does whitelist by default.
    /// These strategies MIGHT not be the only strategy for the underlying token
    /// as additional strategies can be whitelisted by the owner of the factory.
    mapping(IERC20 => IStrategy) public deployedStrategies;

    /// @notice Mapping token => Whether or not a strategy can be deployed for the token
    mapping(IERC20 => bool) public isBlacklisted;

    /// @dev Deprecated: durationVaultBeacon is now immutable in StrategyFactory.
    /// Slot preserved for storage layout compatibility with existing proxies.
    uint256 private __deprecated_durationVaultBeacon;

    /// @notice Mapping token => all duration vault strategies deployed for the token.
    mapping(IERC20 => IDurationVaultStrategy[]) public durationVaultsByToken;

    /// @dev This empty reserved space is put in place to allow future versions to add new
    /// variables without shifting down storage in the inheritance chain.
    /// See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
    /// Storage slots used: __deprecated_strategyBeacon (1) + deployedStrategies (1) + isBlacklisted (1) +
    /// __deprecated_durationVaultBeacon (1) + durationVaultsByToken (1) = 5 slots. Gap: 51 - 5 = 46.
    uint256[46] private __gap;
}
