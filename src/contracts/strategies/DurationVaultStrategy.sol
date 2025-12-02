// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./StrategyBaseTVLLimits.sol";
import "../interfaces/IDurationVaultStrategy.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAllocationManager.sol";
import "../libraries/OperatorSetLib.sol";

/**
 * @title Duration-bound EigenLayer vault strategy with configurable deposit caps and windows.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
contract DurationVaultStrategy is StrategyBaseTVLLimits, IDurationVaultStrategy {
    using OperatorSetLib for OperatorSet;
    /// @notice Address empowered to configure and lock the vault.

    address public vaultAdmin;

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

    /// @notice Delegation manager reference used to register the vault as an operator.
    IDelegationManager public delegationManager;

    /// @notice Allocation manager reference used to register/allocate operator sets.
    IAllocationManager public allocationManager;

    /// @notice Stored operator set metadata for integration with the allocation manager.
    OperatorSet internal _operatorSet;

    /// @notice True when allocations are currently active (i.e. slashable) for the configured operator set.
    bool public allocationsActive;

    /// @notice True when the vault remains registered for the operator set.
    bool public operatorSetRegistered;

    /// @notice Constant representing the full allocation magnitude (1 WAD) for allocation manager calls.
    uint64 internal constant FULL_ALLOCATION = 1e18;

    error OperatorIntegrationInvalid();

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
        _setTVLLimits(config.maxPerDeposit, config.stakeCap);
        _initializeStrategyBase(config.underlyingToken);

        vaultAdmin = config.vaultAdmin;
        duration = config.duration;
        metadataURI = config.metadataURI;

        _configureOperatorIntegration(config);

        emit VaultInitialized(
            vaultAdmin, config.underlyingToken, duration, config.maxPerDeposit, config.stakeCap, metadataURI
        );
    }

    /**
     * @notice Locks the vault, preventing new deposits and withdrawals until maturity.
     */
    function lock() external override onlyVaultAdmin {
        if (lockedAt != 0) revert VaultAlreadyLocked();

        lockedAt = uint64(block.timestamp);
        uint256 rawUnlockTimestamp = uint256(lockedAt) + uint256(duration);
        require(rawUnlockTimestamp <= type(uint64).max, InvalidDuration());
        unlockAt = uint64(rawUnlockTimestamp);

        emit VaultLocked(lockedAt, unlockAt);

        if (!allocationsActive) {
            _allocateFullMagnitude();
        }
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

        if (allocationsActive) {
            _deallocateAll();
            allocationsActive = false;
        }
        if (operatorSetRegistered) {
            _deregisterFromOperatorSet();
        }
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
        return !isLocked();
    }

    /// @inheritdoc IDurationVaultStrategy
    function withdrawalsOpen() public view override returns (bool) {
        if (!isLocked()) {
            return true;
        }
        return isMatured();
    }

    /// @inheritdoc IDurationVaultStrategy
    function operatorIntegrationConfigured() public pure override returns (bool) {
        return true;
    }

    /// @inheritdoc IDurationVaultStrategy
    function operatorSetInfo() external view override returns (address avs, uint32 operatorSetId) {
        return (_operatorSet.avs, _operatorSet.id);
    }

    function _beforeDeposit(IERC20 token, uint256 amount) internal virtual override {
        if (isLocked()) {
            revert DepositsLocked();
        }
        super._beforeDeposit(token, amount);
    }

    function _beforeWithdrawal(address recipient, IERC20 token, uint256 amountShares) internal virtual override {
        if (isLocked() && !isMatured()) {
            revert WithdrawalsLocked();
        }
        super._beforeWithdrawal(recipient, token, amountShares);
    }

    function _configureOperatorIntegration(
        VaultConfig memory config
    ) internal {
        bool hasDelegation = address(config.delegationManager) != address(0);
        bool hasAllocation = address(config.allocationManager) != address(0);
        bool hasAVS = config.operatorSetAVS != address(0);
        bool hasOperatorSetId = config.operatorSetId != 0;

        if (!(hasDelegation && hasAllocation && hasAVS && hasOperatorSetId)) {
            revert OperatorIntegrationInvalid();
        }

        delegationManager = config.delegationManager;
        allocationManager = config.allocationManager;
        _operatorSet = OperatorSet({avs: config.operatorSetAVS, id: config.operatorSetId});

        delegationManager.registerAsOperator(
            config.delegationApprover, config.operatorAllocationDelay, config.operatorMetadataURI
        );

        IAllocationManager.RegisterParams memory params;
        params.avs = config.operatorSetAVS;
        params.operatorSetIds = new uint32[](1);
        params.operatorSetIds[0] = config.operatorSetId;
        params.data = config.operatorSetRegistrationData;
        allocationManager.registerForOperatorSets(address(this), params);

        operatorSetRegistered = true;
    }

    function _allocateFullMagnitude() internal {
        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = _operatorSet;
        params[0].strategies = new IStrategy[](1);
        params[0].strategies[0] = IStrategy(address(this));
        params[0].newMagnitudes = new uint64[](1);
        params[0].newMagnitudes[0] = FULL_ALLOCATION;
        allocationManager.modifyAllocations(address(this), params);
        allocationsActive = true;
    }

    function _deallocateAll() internal {
        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = _operatorSet;
        params[0].strategies = new IStrategy[](1);
        params[0].strategies[0] = IStrategy(address(this));
        params[0].newMagnitudes = new uint64[](1);
        params[0].newMagnitudes[0] = 0;
        allocationManager.modifyAllocations(address(this), params);
    }

    function _deregisterFromOperatorSet() internal {
        IAllocationManager.DeregisterParams memory params;
        params.operator = address(this);
        params.avs = _operatorSet.avs;
        params.operatorSetIds = new uint32[](1);
        params.operatorSetIds[0] = _operatorSet.id;
        allocationManager.deregisterFromOperatorSets(params);
        operatorSetRegistered = false;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     */
    uint256[36] private __gap;
}
