// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./IStrategy.sol";

/**
 * @title Interface for time-bound EigenLayer vault strategies.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IDurationVaultStrategy is IStrategy {
    struct VaultConfig {
        IERC20 underlyingToken;
        address vaultAdmin;
        uint64 depositWindowStart;
        uint64 depositWindowEnd;
        uint64 duration;
        uint256 maxPerDeposit;
        uint256 stakeCap;
        string metadataURI;
    }

    /// @dev Thrown when attempting to use a zero-address vault admin.
    error InvalidVaultAdmin();
    /// @dev Thrown when attempting to configure a zero duration.
    error InvalidDuration();
    /// @dev Thrown when attempting to configure an invalid deposit window.
    error InvalidDepositWindow();
    /// @dev Thrown when attempting to mutate configuration from a non-admin.
    error OnlyVaultAdmin();
    /// @dev Thrown when attempting to lock an already locked vault.
    error VaultAlreadyLocked();
    /// @dev Thrown when attempting to unlock or update a vault that has not been locked yet.
    error VaultNotLocked();
    /// @dev Thrown when attempting to deposit before the window opens.
    error DepositWindowNotStarted();
    /// @dev Thrown when attempting to deposit after the window closes (either manually or via lock).
    error DepositWindowClosed();
    /// @dev Thrown when attempting to withdraw while funds remain locked.
    error WithdrawalsLocked();
    /// @dev Thrown when attempting to mark the vault as matured before duration elapses.
    error DurationNotElapsed();

    event VaultInitialized(
        address indexed vaultAdmin,
        IERC20 indexed underlyingToken,
        uint64 depositWindowStart,
        uint64 depositWindowEnd,
        uint64 duration,
        uint256 maxPerDeposit,
        uint256 stakeCap,
        string metadataURI
    );

    event VaultLocked(uint64 lockedAt, uint64 unlockAt);

    event VaultMatured(uint64 maturedAt);

    event VaultAdminUpdated(address indexed previousAdmin, address indexed newAdmin);

    event DepositWindowUpdated(uint64 newStart, uint64 newEnd);

    event MetadataURIUpdated(string newMetadataURI);

    /**
     * @notice Locks the vault, preventing further deposits / withdrawals until maturity.
     */
    function lock() external;

    /**
     * @notice Marks the vault as matured once the configured duration has elapsed.
     * @dev After maturation, withdrawals are permitted while deposits remain disabled.
     */
    function markMatured() external;

    /**
     * @notice Updates the vault metadata URI.
     */
    function updateMetadataURI(
        string calldata newMetadataURI
    ) external;

    /**
     * @notice Updates the deposit window bounds. Only callable before the vault is locked.
     */
    function updateDepositWindow(uint64 newStart, uint64 newEnd) external;

    /**
     * @notice Transfers vault admin privileges to a new address.
     */
    function transferVaultAdmin(
        address newVaultAdmin
    ) external;

    function vaultAdmin() external view returns (address);
    function depositWindowStart() external view returns (uint64);
    function depositWindowEnd() external view returns (uint64);
    function duration() external view returns (uint64);
    function lockedAt() external view returns (uint64);
    function unlockTimestamp() external view returns (uint64);
    function isLocked() external view returns (bool);
    function isMatured() external view returns (bool);
    function metadataURI() external view returns (string memory);
    function stakeCap() external view returns (uint256);
    function depositsOpen() external view returns (bool);
    function withdrawalsOpen() external view returns (bool);
}
