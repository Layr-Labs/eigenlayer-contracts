// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IDurationVaultStrategy.sol";
import "../libraries/OperatorSetLib.sol";

/// @title Storage layout for DurationVaultStrategy.
/// @author Layr Labs, Inc.
/// @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
abstract contract DurationVaultStrategyStorage is IDurationVaultStrategy {
    /// @notice Constant representing the full allocation magnitude (1 WAD) for allocation manager calls.
    uint64 internal constant FULL_ALLOCATION = 1e18;

    /// @notice Maximum allowable duration (approximately 2 years).
    uint32 internal constant MAX_DURATION = uint32(2 * 365 days);

    /// @notice Address empowered to configure and lock the vault.
    address public vaultAdmin;

    /// @notice Address empowered to advance the vault to withdrawals early (after lock, before duration elapses).
    address public arbitrator;

    /// @notice The enforced lock duration once `lock` is called.
    uint32 public duration;

    /// @notice Timestamp when the vault was locked. Zero indicates the vault is not yet locked.
    uint32 public lockedAt;

    /// @notice Timestamp when the vault unlocks (set at lock time).
    uint32 public unlockAt;

    /// @notice Timestamp when the vault was marked as matured (purely informational).
    uint32 public maturedAt;

    /// @notice Tracks the lifecycle of the vault (deposits -> allocations -> withdrawals).
    VaultState internal _state;

    /// @notice Optional metadata URI describing the vault configuration.
    string public metadataURI;

    /// @notice Stored operator set metadata for integration with the allocation manager.
    OperatorSet internal _operatorSet;

    /// @notice The maximum deposit (in underlyingToken) that this strategy will accept per deposit.
    uint256 public maxPerDeposit;

    /// @notice The maximum deposits (in underlyingToken) that this strategy will hold.
    uint256 public maxTotalDeposits;

    /// @dev This empty reserved space is put in place to allow future versions to add new
    /// variables without shifting down storage in the inheritance chain.
    /// Storage slots used: vaultAdmin (1) + arbitrator (1) + duration/lockedAt/unlockAt/maturedAt/_state (packed, 1) +
    /// metadataURI (1) + _operatorSet (1) + maxPerDeposit (1) + maxTotalDeposits (1) = 6.
    /// Gap: 50 - 7 = 43.
    uint256[43] private __gap;
}
