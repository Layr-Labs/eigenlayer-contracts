// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./StrategyBaseTVLLimits.sol";
import "../interfaces/IDurationVaultStrategy.sol";

/**
 * @title Duration-bound EigenLayer vault strategy with configurable deposit caps and windows.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
contract DurationVaultStrategy is StrategyBaseTVLLimits, IDurationVaultStrategy {
    /// @notice Address empowered to configure and lock the vault.
    address public vaultAdmin;

    /// @notice Timestamp when the deposit window opens. Zero means "immediately open".
    uint64 public depositWindowStart;

    /// @notice Timestamp when the deposit window closes. Zero means "no enforced close" until lock().
    uint64 public depositWindowEnd;

    /// @notice The enforced lock duration once `lock` is called.
    uint64 public duration;

    /// @notice Timestamp when the vault was locked. Zero indicates the vault is not yet locked.
    uint64 public lockedAt;

    /// @notice Timestamp when the vault unlocks (set at lock time).
    uint64 public unlockAt;

    /// @notice Timestamp when the vault was marked as matured (purely informational).
    uint64 public maturedAt;

    /// @notice Optional metadata URI describing the vault configuration.
    string public metadataURI;

    modifier onlyVaultAdmin() {
        if (msg.sender != vaultAdmin) revert OnlyVaultAdmin();
        _;
    }

    constructor(
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        string memory _version
    ) StrategyBaseTVLLimits(_strategyManager, _pauserRegistry, _version) {}

    /**
     * @notice Initializes the vault configuration.
     */
    function initialize(
        VaultConfig memory config
    ) public initializer {
        if (config.vaultAdmin == address(0)) revert InvalidVaultAdmin();
        if (config.duration == 0) revert InvalidDuration();
        if (config.depositWindowEnd != 0 && config.depositWindowEnd <= config.depositWindowStart) {
            revert InvalidDepositWindow();
        }

        _setTVLLimits(config.maxPerDeposit, config.stakeCap);
        _initializeStrategyBase(config.underlyingToken);

        vaultAdmin = config.vaultAdmin;
        depositWindowStart = config.depositWindowStart;
        depositWindowEnd = config.depositWindowEnd;
        duration = config.duration;
        metadataURI = config.metadataURI;

        emit VaultInitialized(
            vaultAdmin,
            config.underlyingToken,
            depositWindowStart,
            depositWindowEnd,
            duration,
            config.maxPerDeposit,
            config.stakeCap,
            metadataURI
        );
    }

    /**
     * @notice Locks the vault, preventing new deposits and withdrawals until maturity.
     */
    function lock() external override onlyVaultAdmin {
        if (lockedAt != 0) revert VaultAlreadyLocked();
        if (depositWindowStart != 0 && block.timestamp < depositWindowStart) revert DepositWindowNotStarted();

        lockedAt = uint64(block.timestamp);
        uint256 rawUnlockTimestamp = uint256(lockedAt) + uint256(duration);
        require(rawUnlockTimestamp <= type(uint64).max, InvalidDuration());
        unlockAt = uint64(rawUnlockTimestamp);

        // Closing the deposit window at the lock time ensures future deposits revert even if no explicit end was set.
        depositWindowEnd = lockedAt;

        emit VaultLocked(lockedAt, unlockAt);
    }

    /**
     * @notice Marks the vault as matured once the configured duration elapses. Callable by anyone.
     */
    function markMatured() external override {
        if (!isMatured()) revert DurationNotElapsed();
        if (maturedAt != 0) {
            // already recorded; noop
            return;
        }
        maturedAt = uint64(block.timestamp);
        emit VaultMatured(maturedAt);
    }

    /**
     * @notice Updates the metadata URI describing the vault.
     */
    function updateMetadataURI(
        string calldata newMetadataURI
    ) external override onlyVaultAdmin {
        metadataURI = newMetadataURI;
        emit MetadataURIUpdated(newMetadataURI);
    }

    /**
     * @notice Updates the deposit window bounds. Cannot be called after locking.
     */
    function updateDepositWindow(uint64 newStart, uint64 newEnd) external override onlyVaultAdmin {
        if (lockedAt != 0) revert VaultAlreadyLocked();
        if (newEnd != 0 && newEnd <= newStart) revert InvalidDepositWindow();
        depositWindowStart = newStart;
        depositWindowEnd = newEnd;
        emit DepositWindowUpdated(newStart, newEnd);
    }

    /**
     * @notice Transfers admin privileges to a new address.
     */
    function transferVaultAdmin(
        address newVaultAdmin
    ) external override onlyVaultAdmin {
        if (newVaultAdmin == address(0)) revert InvalidVaultAdmin();
        emit VaultAdminUpdated(vaultAdmin, newVaultAdmin);
        vaultAdmin = newVaultAdmin;
    }

    /// @inheritdoc IDurationVaultStrategy
    function unlockTimestamp() public view override returns (uint64) {
        return unlockAt;
    }

    /// @inheritdoc IDurationVaultStrategy
    function isLocked() public view override returns (bool) {
        return lockedAt != 0;
    }

    /// @inheritdoc IDurationVaultStrategy
    function isMatured() public view override returns (bool) {
        uint64 lockTimestamp = lockedAt;
        if (lockTimestamp == 0) {
            return false;
        }
        return block.timestamp >= unlockAt && unlockAt != 0;
    }

    /// @inheritdoc IDurationVaultStrategy
    function stakeCap() external view override returns (uint256) {
        return maxTotalDeposits;
    }

    /// @inheritdoc IDurationVaultStrategy
    function depositsOpen() public view override returns (bool) {
        return _depositsWithinWindow() && !isLocked();
    }

    /// @inheritdoc IDurationVaultStrategy
    function withdrawalsOpen() public view override returns (bool) {
        if (!isLocked()) {
            return true;
        }
        return isMatured();
    }

    /**
     * @notice Internal helper verifying deposit timing constraints.
     */
    function _depositsWithinWindow() internal view returns (bool) {
        uint64 start = depositWindowStart;
        uint64 end = depositWindowEnd;
        if (start != 0 && block.timestamp < start) {
            return false;
        }
        if (end != 0 && block.timestamp > end) {
            return false;
        }
        return true;
    }

    function _beforeDeposit(IERC20 token, uint256 amount) internal virtual override {
        if (!isLocked()) {
            uint64 start = depositWindowStart;
            if (start != 0 && block.timestamp < start) revert DepositWindowNotStarted();
            uint64 end = depositWindowEnd;
            if (end != 0 && block.timestamp > end) revert DepositWindowClosed();
        } else {
            revert DepositWindowClosed();
        }

        super._beforeDeposit(token, amount);
    }

    function _beforeWithdrawal(address recipient, IERC20 token, uint256 amountShares) internal virtual override {
        if (isLocked() && !isMatured()) {
            revert WithdrawalsLocked();
        }
        super._beforeWithdrawal(recipient, token, amountShares);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     */
    uint256[40] private __gap;
}
