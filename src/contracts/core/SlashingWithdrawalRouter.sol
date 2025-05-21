// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../mixins/SemVerMixin.sol";
import "./SlashingWithdrawalRouterStorage.sol";

contract SlashingWithdrawalRouter is
    Initializable,
    SlashingWithdrawalRouterStorage,
    OwnableUpgradeable,
    Pausable,
    SemVerMixin
{
    using SafeERC20 for IERC20;
    using OperatorSetLib for *;
    using EnumerableSetUpgradeable for *;
    using EnumerableMapUpgradeable for EnumerableMapUpgradeable.AddressToUintMap;

    /**
     *
     *                         INITIALIZATION
     *
     */
    constructor(
        IAllocationManager _allocationManager,
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        string memory _version
    )
        SlashingWithdrawalRouterStorage(_allocationManager, _strategyManager)
        Pausable(_pauserRegistry)
        SemVerMixin(_version)
    {
        _disableInitializers();
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function initialize(address initialOwner, uint256 initialPausedStatus) external initializer {
        _transferOwnership(initialOwner);
        _setPausedStatus(initialPausedStatus);

        // Set the global burn or redistribution delay to 3.5 days in blocks assuming 12 second blocks.
        _globalBurnOrRedistributionDelayBlocks = 3.5 days / 12 seconds;
    }

    /**
     *
     *                         ACTIONS
     *
     */

    /// @inheritdoc ISlashingWithdrawalRouter
    function startBurnOrRedistributeShares(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy,
        uint256 underlyingAmount
    ) external virtual {
        // Assert that the caller is the `StrategyManager`.
        require(msg.sender == address(strategyManager), OnlyStrategyManager());

        // Create a storage pointer to `_pendingOperatorSets`.
        EnumerableSetUpgradeable.Bytes32Set storage pendingOperatorSets = _pendingOperatorSets;

        // Create a storage pointer to `_pendingSlashIds`.
        EnumerableSetUpgradeable.UintSet storage pendingSlashIds = _pendingSlashIds[operatorSet.key()];

        // Create a storage pointer to `_pendingBurnOrRedistributions`.
        EnumerableMapUpgradeable.AddressToUintMap storage pendingBurnOrRedistributions =
            _pendingBurnOrRedistributions[operatorSet.key()][slashId];

        // Add the slash ID to the pending slash IDs set.
        pendingSlashIds.add(slashId);

        // Add the operator set to the pending operator sets set.
        if (!pendingOperatorSets.contains(operatorSet.key())) {
            pendingOperatorSets.add(operatorSet.key());
        }

        // Add the strategy and underlying amount to the pending burn or redistributions map.
        pendingBurnOrRedistributions.set(address(strategy), underlyingAmount);

        // Set the start block for the slash ID.
        _slashIdToStartBlock[operatorSet.key()][slashId] = uint32(block.number);

        // Emit an event to notify that a burn or redistribution has been started.
        emit StartBurnOrRedistribution(operatorSet, slashId, strategy, underlyingAmount, uint32(block.number));
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function burnOrRedistributeShares(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external virtual onlyWhenNotPaused(PAUSED_BURN_OR_REDISTRIBUTE_SHARES) {
        // Fetch the redistribution recipient for the operator set from the AllocationManager.
        address redistributionRecipient = allocationManager.getRedistributionRecipient(operatorSet);

        // If the redistribution recipient is not the default burn address...
        if (redistributionRecipient != DEFAULT_BURN_ADDRESS) {
            // Assert that the caller is the redistribution recipient.
            require(msg.sender == redistributionRecipient, OnlyRedistributionRecipient());
        }

        // Assert that the escrow is not paused.
        require(!_paused[operatorSet.key()][slashId], IPausable.CurrentlyPaused());

        // Create a storage pointer to `_pendingOperatorSets`.
        EnumerableSetUpgradeable.Bytes32Set storage pendingOperatorSets = _pendingOperatorSets;

        // Create a storage pointer to `_pendingSlashIds`.
        EnumerableSetUpgradeable.UintSet storage pendingSlashIds = _pendingSlashIds[operatorSet.key()];

        // Create a storage pointer to `_pendingBurnOrRedistributions`.
        EnumerableMapUpgradeable.AddressToUintMap storage pendingBurnOrRedistributions =
            _pendingBurnOrRedistributions[operatorSet.key()][slashId];

        // Process the burn or redistribution.
        _processBurnOrRedistribution(pendingBurnOrRedistributions, operatorSet, slashId, redistributionRecipient);

        // Remove the slash ID and operator set from their respective pending lists if no more strategies remain to be processed.
        _tryClearPendingOperatorSetsAndSlashIds(
            pendingOperatorSets, pendingSlashIds, pendingBurnOrRedistributions, operatorSet, slashId
        );
    }

    /**
     *
     *                         PAUSABLE ACTIONS
     *
     */

    /// @inheritdoc ISlashingWithdrawalRouter
    function pauseRedistribution(OperatorSet calldata operatorSet, uint256 slashId) external virtual onlyPauser {
        // Check that the new paused status is not the same as the current paused status.
        _checkNewPausedStatus(operatorSet, slashId, true);

        // Set the paused flag to true.
        _paused[operatorSet.key()][slashId] = true;

        // Emit an event to notify that a redistribution has been paused.
        emit RedistributionPaused(operatorSet, slashId);
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function unpauseRedistribution(OperatorSet calldata operatorSet, uint256 slashId) external virtual onlyUnpauser {
        // Check that the new paused status is not the same as the current paused status.
        _checkNewPausedStatus(operatorSet, slashId, false);

        // Set the paused flag to false.
        _paused[operatorSet.key()][slashId] = false;

        // Emit an event to notify that a redistribution has been unpaused.
        emit RedistributionUnpaused(operatorSet, slashId);
    }

    /**
     *
     *                         OWNER ACTIONS
     *
     */

    /// @inheritdoc ISlashingWithdrawalRouter
    function setGlobalBurnOrRedistributionDelay(
        uint256 delay
    ) external onlyOwner {
        // Set the global burn or redistribution delay.
        _globalBurnOrRedistributionDelayBlocks = uint32(delay);

        // Emit an event to notify that a global burn or redistribution delay has been set.
        emit GlobalBurnOrRedistributionDelaySet(delay);
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function setStrategyBurnOrRedistributionDelay(IStrategy strategy, uint256 delay) external onlyOwner {
        // Set the burn or redistribution delay.
        _strategyBurnOrRedistributionDelayBlocks[address(strategy)] = uint32(delay);

        // Emit an event to notify that a burn or redistribution delay has been set.
        emit StrategyBurnOrRedistributionDelaySet(strategy, delay);
    }

    /**
     *
     *                         HELPERS
     *
     */

    /// @notice Processes burn or redistribution of escrowed tokens for a given operator set and slash ID.
    /// @dev Iterates through pending burn/redistributions in reverse order, checking if each strategy's delay period has elapsed.
    /// If delay has passed, transfers underlying tokens to redistribution recipient and removes from pending map.
    function _processBurnOrRedistribution(
        EnumerableMapUpgradeable.AddressToUintMap storage pendingBurnOrRedistributions,
        OperatorSet calldata operatorSet,
        uint256 slashId,
        address redistributionRecipient
    ) internal {
        // Fetch the total number of pending burn or redistributions for the slash ID before processing.
        uint256 totalPendingForSlashId = pendingBurnOrRedistributions.length();

        // Fetch the start block for the slash ID.
        uint256 startBlock = getBurnOrRedistributionStartBlock(operatorSet, slashId);

        // Iterate over the escrow array in reverse order and pop the processed entries from storage.
        for (uint256 i = totalPendingForSlashId; i > 0; --i) {
            (address strategy, uint256 underlyingAmount) = pendingBurnOrRedistributions.at(i - 1);

            // Fetch the burn or redistribution delay for the strategy.
            uint256 delay = getStrategyBurnOrRedistributionDelay(IStrategy(strategy));

            // Skip this element if the delay has not passed.
            if (startBlock + delay >= block.number) {
                continue;
            }

            // Remove the strategy and underlying amount from the pending burn or redistributions map.
            pendingBurnOrRedistributions.remove(strategy);

            // Transfer the escrowed tokens to the caller.
            IStrategy(strategy).underlyingToken().safeTransfer(redistributionRecipient, underlyingAmount);

            // Emit an event to notify that a burn or redistribution has occurred.
            emit BurnOrRedistribution(
                operatorSet, slashId, IStrategy(strategy), underlyingAmount, redistributionRecipient
            );
        }
    }

    /// @notice Attempts to clear the pending operator sets and slash IDs if no more strategies remain to be processed.
    /// @dev Removes slash ID from pending slash IDs set and deletes start block for slash ID if no more strategies remain.
    /// Also removes operator set from pending operator sets set if no more slash IDs exist for it.
    function _tryClearPendingOperatorSetsAndSlashIds(
        EnumerableSetUpgradeable.Bytes32Set storage pendingOperatorSets,
        EnumerableSetUpgradeable.UintSet storage pendingSlashIds,
        EnumerableMapUpgradeable.AddressToUintMap storage pendingBurnOrRedistributions,
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) internal {
        // Fetch the total number of pending burn or redistributions for the slash ID after processing.
        uint256 totalPendingForSlashId = pendingBurnOrRedistributions.length();

        // If there are no more strategies to process, remove the slash ID from the pending slash IDs set.
        if (totalPendingForSlashId == 0) {
            // Remove the slash ID from the pending slash IDs set.
            pendingSlashIds.remove(slashId);

            // Delete the start block for the slash ID.
            delete _slashIdToStartBlock[operatorSet.key()][slashId];

            // If there are no more slash IDs for the operator set, remove the operator set from the pending operator sets set.
            if (pendingSlashIds.length() == 0) {
                // Remove the operator set from the pending operator sets set.
                pendingOperatorSets.remove(operatorSet.key());
            }
        }
    }

    /// @notice Checks that the new paused status is not the same as the current paused status.
    /// @dev This is needed for event sanitization.
    function _checkNewPausedStatus(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        bool newPauseStatus
    ) internal view {
        // Assert that the new paused status is not the same as the current paused status.
        require(_paused[operatorSet.key()][slashId] != newPauseStatus, IPausable.InvalidNewPausedStatus());
    }

    /**
     *
     *                         GETTERS
     *
     */

    /// @inheritdoc ISlashingWithdrawalRouter
    function getPendingOperatorSets() external view returns (OperatorSet[] memory operatorSets) {
        bytes32[] memory operatorSetKeys = _pendingOperatorSets.values();

        operatorSets = new OperatorSet[](operatorSetKeys.length);

        for (uint256 i = 0; i < operatorSetKeys.length; ++i) {
            operatorSets[i] = operatorSetKeys[i].decode();
        }

        return operatorSets;
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getPendingSlashIds(
        OperatorSet calldata operatorSet
    ) external view returns (uint256[] memory) {
        return _pendingSlashIds[operatorSet.key()].values();
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getPendingBurnOrRedistributions(
        OperatorSet memory operatorSet,
        uint256 slashId
    ) public view returns (IStrategy[] memory strategies, uint256[] memory underlyingAmounts) {
        EnumerableMapUpgradeable.AddressToUintMap storage pendingBurnOrRedistributions =
            _pendingBurnOrRedistributions[operatorSet.key()][slashId];

        uint256 length = pendingBurnOrRedistributions.length();

        strategies = new IStrategy[](length);
        underlyingAmounts = new uint256[](length);

        for (uint256 i = 0; i < length; ++i) {
            (address strategy, uint256 underlyingAmount) = pendingBurnOrRedistributions.at(i);

            strategies[i] = IStrategy(strategy);
            underlyingAmounts[i] = underlyingAmount;
        }
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getPendingBurnOrRedistributions(
        OperatorSet memory operatorSet
    ) public view returns (IStrategy[][] memory strategies, uint256[][] memory underlyingAmounts) {
        EnumerableSetUpgradeable.UintSet storage pendingSlashIds = _pendingSlashIds[operatorSet.key()];

        uint256 length = pendingSlashIds.length();

        strategies = new IStrategy[][](length);
        underlyingAmounts = new uint256[][](length);

        for (uint256 i = 0; i < length; ++i) {
            (strategies[i], underlyingAmounts[i]) = getPendingBurnOrRedistributions(operatorSet, pendingSlashIds.at(i));
        }
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getPendingBurnOrRedistributions()
        public
        view
        returns (IStrategy[][][] memory strategies, uint256[][][] memory underlyingAmounts)
    {
        bytes32[] memory operatorSetKeys = _pendingOperatorSets.values();

        uint256 length = operatorSetKeys.length;

        strategies = new IStrategy[][][](length);
        underlyingAmounts = new uint256[][][](length);

        for (uint256 i = 0; i < length; ++i) {
            (strategies[i], underlyingAmounts[i]) = getPendingBurnOrRedistributions(operatorSetKeys[i].decode());
        }
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getPendingBurnOrRedistributionsCount(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (uint256) {
        return _pendingBurnOrRedistributions[operatorSet.key()][slashId].length();
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getPendingUnderlyingAmountForStrategy(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external view returns (uint256) {
        (, uint256 underlyingAmount) =
            _pendingBurnOrRedistributions[operatorSet.key()][slashId].tryGet(address(strategy));

        return underlyingAmount;
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function isBurnOrRedistributionPaused(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (bool) {
        return _paused[operatorSet.key()][slashId];
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getBurnOrRedistributionStartBlock(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) public view returns (uint256) {
        return _slashIdToStartBlock[operatorSet.key()][slashId];
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getStrategyBurnOrRedistributionDelay(
        IStrategy strategy
    ) public view returns (uint256) {
        // Fetch the global and strategy burn or redistribution delay.
        uint256 globalDelay = _globalBurnOrRedistributionDelayBlocks;
        uint256 strategyDelay = _strategyBurnOrRedistributionDelayBlocks[address(strategy)];

        // If the strategy delay is less than the global delay, return the strategy delay.
        // Otherwise, return the global delay.
        return strategyDelay < globalDelay ? strategyDelay : globalDelay;
    }

    /// @inheritdoc ISlashingWithdrawalRouter
    function getGlobalBurnOrRedistributionDelay() external view returns (uint256) {
        return _globalBurnOrRedistributionDelayBlocks;
    }
}
