// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./StrategyBaseTVLLimits.sol";
import "./DurationVaultStrategyStorage.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAllocationManager.sol";
import "../libraries/OperatorSetLib.sol";

/**
 * @title Duration-bound EigenLayer vault strategy with configurable deposit caps and windows.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
contract DurationVaultStrategy is DurationVaultStrategyStorage, StrategyBaseTVLLimits {
    using OperatorSetLib for OperatorSet;

    /// @notice Constant representing the full allocation magnitude (1 WAD) for allocation manager calls.
    uint64 internal constant FULL_ALLOCATION = 1e18;

    /// @notice Maximum allowable duration (approximately 2 years).
    uint32 internal constant MAX_DURATION = uint32(2 * 365 days);

    /// @notice Delegation manager reference used to register the vault as an operator.
    IDelegationManager public immutable override delegationManager;

    /// @notice Allocation manager reference used to register/allocate operator sets.
    IAllocationManager public immutable override allocationManager;

    error OperatorIntegrationInvalid();

    modifier onlyVaultAdmin() {
        if (msg.sender != vaultAdmin) revert OnlyVaultAdmin();
        _;
    }

    constructor(
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        string memory _version,
        IDelegationManager _delegationManager,
        IAllocationManager _allocationManager
    ) StrategyBaseTVLLimits(_strategyManager, _pauserRegistry, _version) {
        if (address(_delegationManager) == address(0) || address(_allocationManager) == address(0)) {
            revert OperatorIntegrationInvalid();
        }
        delegationManager = _delegationManager;
        allocationManager = _allocationManager;
    }

    /**
     * @notice Initializes the vault configuration.
     */
    function initialize(
        VaultConfig memory config
    ) public initializer {
        if (config.vaultAdmin == address(0)) revert InvalidVaultAdmin();
        if (config.duration == 0 || config.duration > MAX_DURATION) revert InvalidDuration();
        _setTVLLimits(config.maxPerDeposit, config.stakeCap);
        _initializeStrategyBase(config.underlyingToken);

        vaultAdmin = config.vaultAdmin;
        duration = config.duration;
        metadataURI = config.metadataURI;

        _configureOperatorIntegration(config);
        _state = VaultState.Deposits;

        emit VaultInitialized(
            vaultAdmin, config.underlyingToken, duration, config.maxPerDeposit, config.stakeCap, metadataURI
        );
    }

    /**
     * @notice Locks the vault, preventing new deposits and withdrawals until maturity.
     */
    function lock() external override onlyVaultAdmin {
        if (_state != VaultState.Deposits) revert VaultAlreadyLocked();

        uint32 currentTimestamp = _currentTimestamp();
        lockedAt = currentTimestamp;
        uint32 newUnlockAt = currentTimestamp + duration;
        if (newUnlockAt < currentTimestamp) revert InvalidDuration();
        unlockAt = newUnlockAt;

        _state = VaultState.Allocations;

        emit VaultLocked(lockedAt, unlockAt);

        _allocateFullMagnitude();
    }

    /**
     * @notice Marks the vault as matured once the configured duration elapses. Callable by anyone.
     */
    function markMatured() external override {
        if (_state == VaultState.Withdrawals) {
            // already recorded; noop
            return;
        }
        if (_state != VaultState.Allocations || block.timestamp < unlockAt) revert DurationNotElapsed();

        _state = VaultState.Withdrawals;
        maturedAt = _currentTimestamp();
        emit VaultMatured(maturedAt);

        _deallocateAll();
        _deregisterFromOperatorSet();
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
    function unlockTimestamp() public view override returns (uint32) {
        return unlockAt;
    }

    /// @inheritdoc IDurationVaultStrategy
    function isLocked() public view override returns (bool) {
        return _state != VaultState.Deposits;
    }

    /// @inheritdoc IDurationVaultStrategy
    function isMatured() public view override returns (bool) {
        return _state == VaultState.Withdrawals;
    }

    /// @inheritdoc IDurationVaultStrategy
    function state() public view override returns (VaultState) {
        return _state;
    }

    /// @inheritdoc IDurationVaultStrategy
    function stakeCap() external view override returns (uint256) {
        return maxTotalDeposits;
    }

    /// @inheritdoc IDurationVaultStrategy
    function depositsOpen() public view override returns (bool) {
        return _state == VaultState.Deposits;
    }

    /// @inheritdoc IDurationVaultStrategy
    function withdrawalsOpen() public view override returns (bool) {
        return _state != VaultState.Allocations;
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
        if (!depositsOpen()) {
            revert DepositsLocked();
        }
        super._beforeDeposit(token, amount);
    }

    function _beforeWithdrawal(address recipient, IERC20 token, uint256 amountShares) internal virtual override {
        if (!withdrawalsOpen()) {
            revert WithdrawalsLocked();
        }
        super._beforeWithdrawal(recipient, token, amountShares);
    }

    function _configureOperatorIntegration(
        VaultConfig memory config
    ) internal {
        if (config.operatorSetAVS == address(0) || config.operatorSetId == 0) {
            revert OperatorIntegrationInvalid();
        }
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
        allocationsActive = false;
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
    function _currentTimestamp() internal view returns (uint32) {
        uint256 ts = block.timestamp;
        if (ts > type(uint32).max) revert TimestampOverflow();
        return uint32(ts);
    }
}
