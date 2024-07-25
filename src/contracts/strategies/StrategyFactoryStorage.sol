// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IStrategyFactory.sol";

/**
 * @title Storage for the StrategyFactory contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
abstract contract StrategyFactoryStorage is IStrategyFactory {
    // @notice Upgradeable beacon which new Strategies deployed by this contract point to
    IBeacon public strategyBeacon;

    // @notice Mapping token => Strategy contract for the token
    mapping(IERC20 => IStrategy) public tokenStrategy;

    // @notice Mapping token => Whether or not a strategy can be deployed for the token
    mapping(IERC20 => bool) public isBlacklisted;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
