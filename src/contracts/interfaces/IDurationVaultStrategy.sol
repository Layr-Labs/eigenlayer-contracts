// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./IStrategy.sol";
import "./IDelegationManager.sol";
import "./IAllocationManager.sol";
import "./IRewardsCoordinator.sol";
import "../libraries/OperatorSetLib.sol";

/// @title Errors for IDurationVaultStrategy
interface IDurationVaultStrategyErrors {
    /// @dev Thrown when attempting to use a zero-address vault admin.
    error InvalidVaultAdmin();
    /// @dev Thrown when attempting to configure a zero duration.
    error InvalidDuration();
    /// @dev Thrown when attempting to mutate configuration from a non-admin.
    error OnlyVaultAdmin();
    /// @dev Thrown when attempting to call arbitrator-only functionality from a non-arbitrator.
    error OnlyArbitrator();
    /// @dev Thrown when attempting to configure a zero-address arbitrator.
    error InvalidArbitrator();
    /// @dev Thrown when attempting to lock an already locked vault.
    error VaultAlreadyLocked();
    /// @dev Thrown when attempting to deposit after the vault has been locked.
    error DepositsLocked();
    /// @dev Thrown when attempting to remove shares during the allocations period.
    error WithdrawalsLockedDuringAllocations();
    /// @dev Thrown when attempting to add shares when not delegated to the vault operator.
    error MustBeDelegatedToVaultOperator();
    /// @dev Thrown when attempting to mark the vault as matured before duration elapses.
    error DurationNotElapsed();
    /// @dev Thrown when attempting to use the arbitrator early-advance after the duration has elapsed.
    error DurationAlreadyElapsed();
    /// @dev Thrown when attempting to use the arbitrator early-advance before the vault is locked.
    error VaultNotLocked();
    /// @dev Thrown when operator integration inputs are missing or invalid.
    error OperatorIntegrationInvalid();
    /// @dev Thrown when attempting to deposit into a vault whose underlying token is blacklisted.
    error UnderlyingTokenBlacklisted();
}

/// @title Types for IDurationVaultStrategy
interface IDurationVaultStrategyTypes {
    /// @notice Represents the lifecycle state of a duration vault.
    /// @dev UNINITIALIZED: Vault has not been initialized.
    ///      DEPOSITS: Vault is accepting deposits, withdrawals can be queued.
    ///      ALLOCATIONS: Vault is locked, funds are allocated to the operator set.
    ///      WITHDRAWALS: Duration elapsed, vault is matured and withdrawals are enabled.
    enum VaultState {
        UNINITIALIZED,
        DEPOSITS,
        ALLOCATIONS,
        WITHDRAWALS
    }

    /// @notice Configuration parameters for initializing a duration vault.
    /// @param underlyingToken The ERC20 token that stakers deposit into this vault.
    /// @param vaultAdmin The address authorized to manage vault configuration.
    /// @param duration The lock period in seconds after which the vault matures.
    /// @param maxPerDeposit Maximum amount of underlying tokens accepted per deposit.
    /// @param stakeCap Maximum total underlying tokens the vault will accept.
    /// @param metadataURI URI pointing to vault metadata (description, name, etc.).
    /// @param operatorSet The operator set this vault will register and allocate to.
    /// @param operatorSetRegistrationData Data passed to the AVS registrar during registration.
    /// @param delegationApprover Address that approves staker delegations (0x0 for open delegation).
    /// @param operatorMetadataURI URI pointing to operator metadata.
    struct VaultConfig {
        IERC20 underlyingToken;
        address vaultAdmin;
        address arbitrator;
        uint32 duration;
        uint256 maxPerDeposit;
        uint256 stakeCap;
        string metadataURI;
        OperatorSet operatorSet;
        bytes operatorSetRegistrationData;
        address delegationApprover;
        string operatorMetadataURI;
    }
}

/// @title Events for IDurationVaultStrategy
interface IDurationVaultStrategyEvents {
    /// @notice Emitted when a vault is initialized with its configuration.
    /// @param vaultAdmin The address of the vault administrator.
    /// @param underlyingToken The ERC20 token used for deposits.
    /// @param duration The lock duration in seconds.
    /// @param maxPerDeposit Maximum deposit amount per transaction.
    /// @param stakeCap Maximum total deposits allowed.
    /// @param metadataURI URI pointing to vault metadata.
    event VaultInitialized(
        address indexed vaultAdmin,
        IERC20 indexed underlyingToken,
        uint32 duration,
        uint256 maxPerDeposit,
        uint256 stakeCap,
        string metadataURI
    );

    /// @notice Emitted when the vault is locked, transitioning to ALLOCATIONS state.
    /// @param lockedAt Timestamp when the vault was locked.
    /// @param unlockAt Timestamp when the vault will mature.
    event VaultLocked(uint32 lockedAt, uint32 unlockAt);

    /// @notice Emitted when the vault matures, transitioning to WITHDRAWALS state.
    /// @param maturedAt Timestamp when the vault matured.
    event VaultMatured(uint32 maturedAt);

    /// @notice Emitted when the vault is advanced to WITHDRAWALS early by the arbitrator.
    /// @param arbitrator The arbitrator that performed the early advance.
    /// @param maturedAt Timestamp when the vault transitioned to WITHDRAWALS.
    event VaultAdvancedToWithdrawals(address indexed arbitrator, uint32 maturedAt);

    /// @notice Emitted when the vault metadata URI is updated.
    /// @param newMetadataURI The new metadata URI.
    event MetadataURIUpdated(string newMetadataURI);
}

/// @title Interface for time-bound EigenLayer vault strategies.
/// @author Layr Labs, Inc.
/// @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
/// @dev Duration vault strategies are operator-bound strategies that accept deposits during an
/// open window, lock funds for a configurable duration while allocating to an operator set,
/// and then enable withdrawals after maturity. The vault itself acts as an operator.
interface IDurationVaultStrategy is
    IStrategy,
    IDurationVaultStrategyErrors,
    IDurationVaultStrategyTypes,
    IDurationVaultStrategyEvents
{
    /// @notice Locks the vault, preventing further deposits and withdrawal queuing until maturity.
    /// @dev Transitions state from DEPOSITS to ALLOCATIONS. Allocates full magnitude to the
    /// configured operator set. Only callable by the vault admin.
    function lock() external;

    /// @notice Marks the vault as matured once the configured duration has elapsed.
    /// @dev Transitions state from ALLOCATIONS to WITHDRAWALS.
    ///
    /// Best-effort operator cleanup:
    /// - Attempts to deallocate from the configured operator set and deregister the vault as an operator.
    /// - These external calls are intentionally best-effort so `markMatured()` cannot be bricked and user
    ///   withdrawals cannot be indefinitely locked.
    ///
    /// Common reasons deallocation/deregistration can fail include:
    /// - AllocationManager pausing allocations/deallocations or register/deregister operations.
    /// - AVS-side registrar rejecting deregistration (e.g., operator removed/ejected from an operator set).
    /// - AllocationManager state constraints (e.g., pending allocation modification preventing an update).
    ///
    /// After maturation, withdrawals are permitted while deposits remain disabled. Callable by anyone once
    /// the duration has elapsed.
    function markMatured() external;

    /// @notice Advances the vault to WITHDRAWALS early, after lock but before duration elapses.
    /// @dev Transitions state from ALLOCATIONS to WITHDRAWALS, and triggers the same best-effort operator cleanup
    /// as `markMatured()`. Only callable by the configured arbitrator.
    function advanceToWithdrawals() external;

    /// @notice Updates the vault metadata URI.
    /// @param newMetadataURI The new metadata URI to set.
    /// @dev Only callable by the vault admin.
    function updateMetadataURI(
        string calldata newMetadataURI
    ) external;

    /// @notice Updates the TVL limits for max deposit per transaction and total stake cap.
    /// @param newMaxPerDeposit New maximum deposit amount per transaction.
    /// @param newStakeCap New maximum total deposits allowed.
    /// @dev Only callable by the vault admin while deposits are open (before lock).
    function updateTVLLimits(
        uint256 newMaxPerDeposit,
        uint256 newStakeCap
    ) external;

    /// @notice Returns the vault administrator address.
    function vaultAdmin() external view returns (address);

    /// @notice Returns the arbitrator address.
    function arbitrator() external view returns (address);

    /// @notice Returns the configured lock duration in seconds.
    function duration() external view returns (uint32);

    /// @notice Returns the timestamp when the vault was locked.
    function lockedAt() external view returns (uint32);

    /// @notice Returns the timestamp when the vault will/did mature.
    function unlockTimestamp() external view returns (uint32);

    /// @notice Returns true if the vault has been locked (state != DEPOSITS).
    function isLocked() external view returns (bool);

    /// @notice Returns true if the vault has matured (state == WITHDRAWALS).
    function isMatured() external view returns (bool);

    /// @notice Returns the current vault lifecycle state.
    function state() external view returns (VaultState);

    /// @notice Returns the vault metadata URI.
    function metadataURI() external view returns (string memory);

    /// @notice Returns the maximum total deposits allowed (alias for maxTotalDeposits).
    function stakeCap() external view returns (uint256);

    /// @notice Returns the maximum deposit amount per transaction.
    function maxPerDeposit() external view returns (uint256);

    /// @notice Returns the maximum total deposits allowed.
    function maxTotalDeposits() external view returns (uint256);

    /// @notice Returns true if deposits are currently accepted (state == DEPOSITS).
    function depositsOpen() external view returns (bool);

    /// @notice Returns true if withdrawal queuing is allowed (state != ALLOCATIONS).
    function withdrawalsOpen() external view returns (bool);

    /// @notice Returns the DelegationManager contract reference.
    function delegationManager() external view returns (IDelegationManager);

    /// @notice Returns the AllocationManager contract reference.
    function allocationManager() external view returns (IAllocationManager);

    /// @notice Returns the RewardsCoordinator contract reference.
    function rewardsCoordinator() external view returns (IRewardsCoordinator);

    /// @notice Returns true if operator integration has been configured.
    function operatorIntegrationConfigured() external view returns (bool);

    /// @notice Returns true if the vault is registered to the operator set.
    function operatorSetRegistered() external view returns (bool);

    /// @notice Returns true if allocations are currently active (state == ALLOCATIONS).
    function allocationsActive() external view returns (bool);

    /// @notice Returns the operator set info (AVS address and operator set ID).
    function operatorSetInfo() external view returns (address avs, uint32 operatorSetId);
}
