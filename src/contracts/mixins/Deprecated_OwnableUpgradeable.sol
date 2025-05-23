// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

/**
 * @title Deprecated_OwnableUpgradeable
 * @dev This contract can be inherited in place of OpenZeppelin's OwnableUpgradeable
 * to maintain the same storage layout while effectively removing the Ownable functionality.
 *
 * This is useful in cases where a contract previously used OwnableUpgradeable but no longer
 * needs ownership functionality, yet must maintain the same storage slots to ensure
 * compatibility with existing deployed proxies.
 *
 * The contract preserves the same storage layout as OwnableUpgradeable:
 * - It keeps the `_owner` storage variable in the same slot
 * - It maintains the same storage gap for future upgrades
 */
abstract contract Deprecated_OwnableUpgradeable is Initializable {
    address private _owner;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
