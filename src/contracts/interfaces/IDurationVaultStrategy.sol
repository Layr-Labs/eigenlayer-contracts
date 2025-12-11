// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./IStrategy.sol";
import "./IDelegationManager.sol";
import "./IAllocationManager.sol";
import "../libraries/OperatorSetLib.sol";

/// @title Interface for time-bound EigenLayer vault strategies.
/// @author Layr Labs, Inc.
/// @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
interface IDurationVaultStrategy is IStrategy {
    enum VaultState {
        UNINITIALIZED,
        DEPOSITS,
        ALLOCATIONS,
        WITHDRAWALS
    }

    struct VaultConfig {
        IERC20 underlyingToken;
        address vaultAdmin;
        uint32 duration;
        uint256 maxPerDeposit;
        uint256 stakeCap;
        string metadataURI;
        OperatorSet operatorSet;
        bytes operatorSetRegistrationData;
        address delegationApprover;
        uint32 operatorAllocationDelay;
        string operatorMetadataURI;
    }

    /// @dev Thrown when attempting to use a zero-address vault admin.
    error InvalidVaultAdmin();
    /// @dev Thrown when attempting to configure a zero duration.
    error InvalidDuration();
    /// @dev Thrown when attempting to mutate configuration from a non-admin.
    error OnlyVaultAdmin();
    /// @dev Thrown when attempting to lock an already locked vault.
    error VaultAlreadyLocked();
    /// @dev Thrown when attempting to deposit after the vault has been locked.
    error DepositsLocked();
    /// @dev Thrown when attempting to withdraw while funds remain locked.
    error WithdrawalsLocked();
    /// @dev Thrown when attempting to remove shares during the allocations period.
    error WithdrawalsLockedDuringAllocations();
    /// @dev Thrown when attempting to add shares when not delegated to the vault operator.
    error MustBeDelegatedToVaultOperator();
    /// @dev Thrown when attempting to mark the vault as matured before duration elapses.
    error DurationNotElapsed();
    /// @dev Thrown when operator integration inputs are missing or invalid.
    error OperatorIntegrationInvalid();

    event VaultInitialized(
        address indexed vaultAdmin,
        IERC20 indexed underlyingToken,
        uint32 duration,
        uint256 maxPerDeposit,
        uint256 stakeCap,
        string metadataURI
    );

    event VaultLocked(uint32 lockedAt, uint32 unlockAt);

    event VaultMatured(uint32 maturedAt);

    event MetadataURIUpdated(string newMetadataURI);

    /// @notice Locks the vault, preventing further deposits / withdrawals until maturity.
    function lock() external;

    /// @notice Marks the vault as matured once the configured duration has elapsed.
    /// @dev After maturation, withdrawals are permitted while deposits remain disabled.
    function markMatured() external;

    /// @notice Updates the vault metadata URI.
    function updateMetadataURI(
        string calldata newMetadataURI
    ) external;

    /// @notice Updates the TVL limits for max deposit per transaction and total stake cap.
    /// @dev Only callable by the vault admin while deposits are open (before lock).
    function updateTVLLimits(
        uint256 newMaxPerDeposit,
        uint256 newStakeCap
    ) external;

    function vaultAdmin() external view returns (address);
    function duration() external view returns (uint32);
    function lockedAt() external view returns (uint32);
    function unlockTimestamp() external view returns (uint32);
    function isLocked() external view returns (bool);
    function isMatured() external view returns (bool);
    function state() external view returns (VaultState);
    function metadataURI() external view returns (string memory);
    function stakeCap() external view returns (uint256);
    function maxPerDeposit() external view returns (uint256);
    function maxTotalDeposits() external view returns (uint256);
    function depositsOpen() external view returns (bool);
    function withdrawalsOpen() external view returns (bool);
    function delegationManager() external view returns (IDelegationManager);
    function allocationManager() external view returns (IAllocationManager);
    function operatorIntegrationConfigured() external view returns (bool);
    function operatorSetRegistered() external view returns (bool);
    function allocationsActive() external view returns (bool);
    function operatorSetInfo() external view returns (address avs, uint32 operatorSetId);

    /// @notice Underlying amount currently queued for withdrawal but not yet completed.
    function queuedUnderlying() external view returns (uint256);
}
