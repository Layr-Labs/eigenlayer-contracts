// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./StrategyBase.sol";
import "./DurationVaultStrategyStorage.sol";
import "../interfaces/IDurationVaultStrategy.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAllocationManager.sol";
import "../libraries/OperatorSetLib.sol";

/// @title Duration-bound EigenLayer vault strategy with configurable deposit caps and windows.
/// @author Layr Labs, Inc.
/// @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
contract DurationVaultStrategy is DurationVaultStrategyStorage, StrategyBase {
    using OperatorSetLib for OperatorSet;

    /// @dev Thrown when attempting to allocate while a pending allocation modification already exists.
    error PendingAllocation();

    /// @notice Emitted when `maxPerDeposit` value is updated from `previousValue` to `newValue`.
    event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue);

    /// @notice Emitted when `maxTotalDeposits` value is updated from `previousValue` to `newValue`.
    event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue);

    /// @notice Delegation manager reference used to register the vault as an operator.
    IDelegationManager public immutable override delegationManager;

    /// @notice Allocation manager reference used to register/allocate operator sets.
    IAllocationManager public immutable override allocationManager;

    modifier onlyVaultAdmin() {
        require(msg.sender == vaultAdmin, OnlyVaultAdmin());
        _;
    }

    constructor(
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        IDelegationManager _delegationManager,
        IAllocationManager _allocationManager
    ) StrategyBase(_strategyManager, _pauserRegistry) {
        require(
            address(_delegationManager) != address(0) && address(_allocationManager) != address(0),
            OperatorIntegrationInvalid()
        );
        delegationManager = _delegationManager;
        allocationManager = _allocationManager;
    }

    /// @notice Initializes the vault configuration.
    function initialize(
        VaultConfig memory config
    ) public initializer {
        require(config.vaultAdmin != address(0), InvalidVaultAdmin());
        require(config.duration != 0 && config.duration <= MAX_DURATION, InvalidDuration());
        _setTVLLimits(config.maxPerDeposit, config.stakeCap);
        _initializeStrategyBase(config.underlyingToken);

        vaultAdmin = config.vaultAdmin;
        duration = config.duration;
        metadataURI = config.metadataURI;

        _configureOperatorIntegration(config);
        _state = VaultState.DEPOSITS;

        emit VaultInitialized(
            vaultAdmin,
            config.underlyingToken,
            duration,
            config.maxPerDeposit,
            config.stakeCap,
            metadataURI
        );
    }

    /// @notice Locks the vault, preventing new deposits and withdrawals until maturity.
    function lock() external override onlyVaultAdmin {
        require(_state == VaultState.DEPOSITS, VaultAlreadyLocked());

        uint32 currentTimestamp = uint32(block.timestamp);
        lockedAt = currentTimestamp;
        unlockAt = currentTimestamp + duration;

        _state = VaultState.ALLOCATIONS;

        emit VaultLocked(lockedAt, unlockAt);

        _allocateFullMagnitude();
    }

    /// @notice Marks the vault as matured once the configured duration elapses. Callable by anyone.
    function markMatured() external override {
        if (_state == VaultState.WITHDRAWALS) {
            // already recorded; noop
            return;
        }
        require(_state == VaultState.ALLOCATIONS, DurationNotElapsed());
        require(block.timestamp >= unlockAt, DurationNotElapsed());
        _state = VaultState.WITHDRAWALS;
        maturedAt = uint32(block.timestamp);
        emit VaultMatured(maturedAt);

        _deallocateAll();
        _deregisterFromOperatorSet();
    }

    /// @notice Updates the metadata URI describing the vault.
    function updateMetadataURI(
        string calldata newMetadataURI
    ) external override onlyVaultAdmin {
        metadataURI = newMetadataURI;
        emit MetadataURIUpdated(newMetadataURI);
    }

    /// @notice Updates the TVL limits for max deposit per transaction and total stake cap.
    /// @dev Only callable by the vault admin while deposits are open (before lock).
    function updateTVLLimits(
        uint256 newMaxPerDeposit,
        uint256 newStakeCap
    ) external override onlyVaultAdmin {
        require(depositsOpen(), DepositsLocked());
        _setTVLLimits(newMaxPerDeposit, newStakeCap);
    }

    /// @notice Allows the unpauser to update TVL limits, mirroring `StrategyBaseTVLLimits`.
    function setTVLLimits(
        uint256 newMaxPerDeposit,
        uint256 newMaxTotalDeposits
    ) external onlyUnpauser {
        // Keep vault config changes constrained to the deposits window.
        require(depositsOpen(), DepositsLocked());
        _setTVLLimits(newMaxPerDeposit, newMaxTotalDeposits);
    }

    /// @notice Returns the current TVL limits (per-deposit and total stake cap).
    /// @dev Helper for tests and parity with `StrategyBaseTVLLimits`.
    function getTVLLimits() external view returns (uint256, uint256) {
        return (maxPerDeposit, maxTotalDeposits);
    }

    /// @inheritdoc StrategyBase
    /// @dev State and TVL limit checks are performed in `beforeAddShares` which is called
    /// after deposit by the StrategyManager. This hook only validates the token type.
    function _beforeDeposit(
        IERC20 token,
        uint256 /*amount*/
    ) internal virtual override {
        require(_state == VaultState.DEPOSITS, DepositsLocked());
        super._beforeDeposit(token, 0);
    }

    /// @inheritdoc IDurationVaultStrategy
    function unlockTimestamp() public view override returns (uint32) {
        return unlockAt;
    }

    /// @inheritdoc IDurationVaultStrategy
    function isLocked() public view override returns (bool) {
        return _state != VaultState.DEPOSITS;
    }

    /// @inheritdoc IDurationVaultStrategy
    function isMatured() public view override returns (bool) {
        return _state == VaultState.WITHDRAWALS;
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
        return _state == VaultState.DEPOSITS;
    }

    /// @inheritdoc IDurationVaultStrategy
    function withdrawalsOpen() public view override returns (bool) {
        return _state != VaultState.ALLOCATIONS;
    }

    /// @inheritdoc IStrategy
    function beforeAddShares(
        address staker,
        uint256 shares
    ) external view override(IStrategy, StrategyBase) onlyStrategyManager {
        require(_state == VaultState.DEPOSITS, DepositsLocked());
        require(delegationManager.delegatedTo(staker) == address(this), MustBeDelegatedToVaultOperator());

        // Enforce per-deposit cap using the minted shares as proxy for underlying.
        uint256 amountUnderlying = sharesToUnderlyingView(shares);
        require(amountUnderlying <= maxPerDeposit, MaxPerDepositExceedsMax());

        // Enforce total cap using operatorShares (active, non-queued shares).
        // At this point, operatorShares hasn't been updated yet, so we add the new shares.
        uint256 currentOperatorShares = _getOperatorShares();
        uint256 postDepositUnderlying = sharesToUnderlyingView(currentOperatorShares + shares);
        require(postDepositUnderlying <= maxTotalDeposits, BalanceExceedsMaxTotalDeposits());
    }

    /// @notice Returns the current operator shares for this vault/strategy.
    function _getOperatorShares() internal view returns (uint256) {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(address(this));
        uint256[] memory sharesArray = delegationManager.getOperatorShares(address(this), strategies);
        return sharesArray[0];
    }

    /// @inheritdoc IStrategy
    function beforeRemoveShares(
        address,
        uint256
    ) external view override(IStrategy, StrategyBase) onlyStrategyManager {
        // Queuing withdrawals is blocked during ALLOCATIONS. Withdrawals queued during
        // DEPOSITS can complete during ALLOCATIONS since they were queued before lock.
        require(_state != VaultState.ALLOCATIONS, WithdrawalsLockedDuringAllocations());
    }

    /// @notice Sets the maximum deposits (in underlyingToken) that this strategy will hold and accept per deposit.
    function _setTVLLimits(
        uint256 newMaxPerDeposit,
        uint256 newMaxTotalDeposits
    ) internal {
        emit MaxPerDepositUpdated(maxPerDeposit, newMaxPerDeposit);
        emit MaxTotalDepositsUpdated(maxTotalDeposits, newMaxTotalDeposits);
        require(newMaxPerDeposit <= newMaxTotalDeposits, MaxPerDepositExceedsMax());
        maxPerDeposit = newMaxPerDeposit;
        maxTotalDeposits = newMaxTotalDeposits;
    }

    /// @inheritdoc IDurationVaultStrategy
    function operatorIntegrationConfigured() public pure override returns (bool) {
        return true;
    }

    /// @inheritdoc IDurationVaultStrategy
    function operatorSetInfo() external view override returns (address avs, uint32 operatorSetId) {
        return (_operatorSet.avs, _operatorSet.id);
    }

    function operatorSetRegistered() public view override returns (bool) {
        return _state == VaultState.DEPOSITS || _state == VaultState.ALLOCATIONS;
    }

    function allocationsActive() public view override returns (bool) {
        return _state == VaultState.ALLOCATIONS;
    }

    function _beforeWithdrawal(
        address, /*recipient*/
        IERC20 token,
        uint256 /*amountShares*/
    ) internal view virtual override {
        // Withdrawals are only blocked at the queue stage (beforeRemoveShares).
        // Once queued, withdrawals can complete in any state. The DelegationManager
        // validates that only properly-queued withdrawals reach this point.
        require(token == underlyingToken, OnlyUnderlyingToken());
    }

    function _configureOperatorIntegration(
        VaultConfig memory config
    ) internal {
        require(config.operatorSet.avs != address(0), OperatorIntegrationInvalid());
        _operatorSet = config.operatorSet;

        // Set allocation delay strictly greater than withdrawal delay to protect pre-lock queued withdrawals.
        uint32 minWithdrawal = delegationManager.minWithdrawalDelayBlocks();
        uint32 allocationDelay = minWithdrawal + 1;

        // apply allocation delay at registration
        delegationManager.registerAsOperator(config.delegationApprover, allocationDelay, config.operatorMetadataURI);

        IAllocationManager.RegisterParams memory params;
        params.avs = config.operatorSet.avs;
        params.operatorSetIds = new uint32[](1);
        params.operatorSetIds[0] = config.operatorSet.id;
        params.data = config.operatorSetRegistrationData;
        allocationManager.registerForOperatorSets(address(this), params);
    }

    function _allocateFullMagnitude() internal {
        // Ensure no pending allocation modification exists for this operator/operatorSet/strategy.
        // Pending modifications would cause ModificationAlreadyPending() in AllocationManager.modifyAllocations.
        IAllocationManager.Allocation memory alloc =
            allocationManager.getAllocation(address(this), _operatorSet, IStrategy(address(this)));
        require(alloc.effectBlock == 0, PendingAllocation());

        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = _operatorSet;
        params[0].strategies = new IStrategy[](1);
        params[0].strategies[0] = IStrategy(address(this));
        params[0].newMagnitudes = new uint64[](1);
        params[0].newMagnitudes[0] = FULL_ALLOCATION;
        allocationManager.modifyAllocations(address(this), params);
    }

    function _deallocateAll() internal {
        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = _operatorSet;
        params[0].strategies = new IStrategy[](1);
        params[0].strategies[0] = IStrategy(address(this));
        params[0].newMagnitudes = new uint64[](1);
        params[0].newMagnitudes[0] = 0;
        // This call is best-effort: failures should not brick `markMatured()` and lock user funds.
        try allocationManager.modifyAllocations(address(this), params) {} catch {}
    }

    function _deregisterFromOperatorSet() internal {
        IAllocationManager.DeregisterParams memory params;
        params.operator = address(this);
        params.avs = _operatorSet.avs;
        params.operatorSetIds = new uint32[](1);
        params.operatorSetIds[0] = _operatorSet.id;
        // This call is best-effort: failures should not brick `markMatured()` and lock user funds.
        try allocationManager.deregisterFromOperatorSets(params) {} catch {}
    }
}
