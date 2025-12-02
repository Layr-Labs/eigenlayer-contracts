// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IDurationVaultStrategy.sol";
import "../libraries/OperatorSetLib.sol";

/**
 * @title Storage layout for DurationVaultStrategy.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
abstract contract DurationVaultStrategyStorage is IDurationVaultStrategy {
    /// @notice Address empowered to configure and lock the vault.
    address public vaultAdmin;

    /// @notice The enforced lock duration once `lock` is called.
    uint32 public duration;

    /// @notice Timestamp when the vault was locked. Zero indicates the vault is not yet locked.
    uint32 public lockedAt;

    /// @notice Timestamp when the vault unlocks (set at lock time).
    uint32 public unlockAt;

    /// @notice Timestamp when the vault was marked as matured (purely informational).
    uint32 public maturedAt;

    /// @notice Optional metadata URI describing the vault configuration.
    string public metadataURI;

    /// @notice Stored operator set metadata for integration with the allocation manager.
    OperatorSet internal _operatorSet;

    /// @notice Tracks the lifecycle of the vault (deposits -> allocations -> withdrawals).
    VaultState internal _state;

    /// @notice True when allocations are currently active (i.e. slashable) for the configured operator set.
    bool public allocationsActive;

    /// @notice True when the vault remains registered for the operator set.
    bool public operatorSetRegistered;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     */
    uint256[41] private __gap;
}


