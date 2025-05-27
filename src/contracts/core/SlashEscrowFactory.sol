// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/ClonesUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../mixins/SemVerMixin.sol";
import "./SlashEscrowFactoryStorage.sol";

contract SlashEscrowFactory is Initializable, SlashEscrowFactoryStorage, OwnableUpgradeable, Pausable, SemVerMixin {
    using SafeERC20 for IERC20;
    using OperatorSetLib for *;
    using EnumerableSet for *;
    using ClonesUpgradeable for address;
    /**
     *
     *                         INITIALIZATION
     *
     */

    constructor(
        IAllocationManager _allocationManager,
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        ISlashEscrow _slashEscrowImplementation,
        string memory _version
    )
        SlashEscrowFactoryStorage(_allocationManager, _strategyManager, _slashEscrowImplementation)
        Pausable(_pauserRegistry)
        SemVerMixin(_version)
    {
        _disableInitializers();
    }

    /// @inheritdoc ISlashEscrowFactory
    function initialize(
        address initialOwner,
        uint256 initialPausedStatus,
        uint32 initialGlobalDelayBlocks
    ) external initializer {
        _transferOwnership(initialOwner);
        _setPausedStatus(initialPausedStatus);
        _setGlobalBurnOrRedistributionDelay(initialGlobalDelayBlocks);
    }

    /**
     *
     *                         ACTIONS
     *
     */

    /// @inheritdoc ISlashEscrowFactory
    function initiateSlashEscrow(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external virtual {
        require(msg.sender == address(strategyManager), OnlyStrategyManager());

        // Create storage pointers for readability.
        EnumerableSet.Bytes32Set storage pendingOperatorSets = _pendingOperatorSets;
        EnumerableSet.UintSet storage pendingSlashIds = _pendingSlashIds[operatorSet.key()];
        EnumerableSet.AddressSet storage pendingStrategiesForSlashId =
            _pendingStrategiesForSlashId[operatorSet.key()][slashId];

        // Add the slash ID, operator set, and strategy to their respective pending sets.
        pendingSlashIds.add(slashId);
        pendingOperatorSets.add(operatorSet.key());
        pendingStrategiesForSlashId.add(address(strategy));

        // Set the start block for the slash ID.
        _slashIdToStartBlock[operatorSet.key()][slashId] = uint32(block.number);
        emit StartBurnOrRedistribution(operatorSet, slashId, strategy, uint32(block.number));
    }

    /// @inheritdoc ISlashEscrowFactory
    function releaseSlashEscrow(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external virtual onlyWhenNotPaused(PAUSED_BURN_OR_REDISTRIBUTE_SHARES) {
        address redistributionRecipient = allocationManager.getRedistributionRecipient(operatorSet);

        // If the redistribution recipient is not the default burn address...
        if (redistributionRecipient != DEFAULT_BURN_ADDRESS) {
            require(msg.sender == redistributionRecipient, OnlyRedistributionRecipient());
        }

        // Assert that the slash ID is not paused
        require(!isBurnOrRedistributionPaused(operatorSet, slashId), IPausable.CurrentlyPaused());

        // Assert that the escrow delay has elapsed
        require(block.number >= getBurnOrRedistributionCompleteBlock(operatorSet, slashId), EscrowDelayNotElapsed());

        // Calling `decreaseBurnOrRedistributableShares` will transfer the underlying tokens to the `SlashEscrow`.
        // NOTE: While `decreaseBurnOrRedistributableShares` may have already been called, we call it again to ensure that the
        // underlying tokens are actually in escrow before processing and removing storage (which would otherwise prevent
        // the tokens from being released).
        strategyManager.decreaseBurnOrRedistributableShares(operatorSet, slashId);

        // Deploy the counterfactual `SlashEscrow`.
        ISlashEscrow slashEscrow = deploySlashEscrow(operatorSet, slashId);

        // Release the slashEscrow
        _processSlashEscrow(operatorSet, slashId, slashEscrow, redistributionRecipient);
    }

    /// @inheritdoc ISlashEscrowFactory
    function deploySlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) public returns (ISlashEscrow) {
        ISlashEscrow slashEscrow = getSlashEscrow(operatorSet, slashId);

        // If the slash escrow is not deployed...
        if (!isDeployedSlashEscrow(slashEscrow)) {
            return ISlashEscrow(
                address(slashEscrowImplementation).cloneDeterministic(computeSlashEscrowSalt(operatorSet, slashId))
            );
        }

        return slashEscrow;
    }

    /**
     *
     *                         PAUSABLE ACTIONS
     *
     */

    /// @inheritdoc ISlashEscrowFactory
    function pauseRedistribution(OperatorSet calldata operatorSet, uint256 slashId) external virtual onlyPauser {
        _checkNewPausedStatus(operatorSet, slashId, true);
        _paused[operatorSet.key()][slashId] = true;
        emit RedistributionPaused(operatorSet, slashId);
    }

    /// @inheritdoc ISlashEscrowFactory
    function unpauseRedistribution(OperatorSet calldata operatorSet, uint256 slashId) external virtual onlyUnpauser {
        _checkNewPausedStatus(operatorSet, slashId, false);
        _paused[operatorSet.key()][slashId] = false;
        emit RedistributionUnpaused(operatorSet, slashId);
    }

    /**
     *
     *                         OWNER ACTIONS
     *
     */

    /// @inheritdoc ISlashEscrowFactory
    function setGlobalBurnOrRedistributionDelay(
        uint32 delay
    ) external onlyOwner {
        _setGlobalBurnOrRedistributionDelay(delay);
    }

    /// @inheritdoc ISlashEscrowFactory
    function setStrategyBurnOrRedistributionDelay(IStrategy strategy, uint32 delay) external onlyOwner {
        _strategyBurnOrRedistributionDelayBlocks[address(strategy)] = delay;
        emit StrategyBurnOrRedistributionDelaySet(strategy, delay);
    }

    /**
     *
     *                         HELPERS
     *
     */

    /// @notice Processes the slash escrow.
    function _processSlashEscrow(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        ISlashEscrow slashEscrow,
        address redistributionRecipient
    ) internal {
        // Create storage pointers for readability.
        EnumerableSet.Bytes32Set storage pendingOperatorSets = _pendingOperatorSets;
        EnumerableSet.UintSet storage pendingSlashIds = _pendingSlashIds[operatorSet.key()];
        EnumerableSet.AddressSet storage pendingStrategiesForSlashId =
            _pendingStrategiesForSlashId[operatorSet.key()][slashId];

        // Iterate over the escrow array in reverse order and pop the processed entries from storage.
        uint256 totalPendingForSlashId = pendingStrategiesForSlashId.length();
        for (uint256 i = totalPendingForSlashId; i > 0; --i) {
            address strategy = pendingStrategiesForSlashId.at(i - 1);

            // Burn or redistribute the underlying tokens for the strategy.
            slashEscrow.burnOrRedistributeUnderlyingTokens(
                ISlashEscrowFactory(address(this)),
                slashEscrowImplementation,
                operatorSet,
                slashId,
                redistributionRecipient,
                IStrategy(strategy)
            );

            // Remove the strategy and underlying amount from the pending burn or redistributions map.
            pendingStrategiesForSlashId.remove(strategy);
            emit BurnOrRedistributionComplete(operatorSet, slashId, IStrategy(strategy), redistributionRecipient);
        }

        // Remove the slash ID from the pending slash IDs set.
        pendingSlashIds.remove(slashId);

        // Delete the start block for the slash ID.
        delete _slashIdToStartBlock[operatorSet.key()][slashId];

        // Remove the operator set from the pending operator sets set if there are no more pending slash IDs.
        if (pendingSlashIds.length() == 0) {
            pendingOperatorSets.remove(operatorSet.key());
        }
    }

    /// @notice Sets the global burn or redistribution delay.
    function _setGlobalBurnOrRedistributionDelay(
        uint32 delay
    ) internal {
        _globalBurnOrRedistributionDelayBlocks = delay;
        emit GlobalBurnOrRedistributionDelaySet(delay);
    }

    /// @notice Checks that the new paused status is not the same as the current paused status.
    /// @dev This is needed for event sanitization.
    function _checkNewPausedStatus(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        bool newPauseStatus
    ) internal view {
        require(_paused[operatorSet.key()][slashId] != newPauseStatus, IPausable.InvalidNewPausedStatus());
    }

    /**
     *
     *                         GETTERS
     *
     */

    /// @inheritdoc ISlashEscrowFactory
    function getPendingOperatorSets() public view returns (OperatorSet[] memory operatorSets) {
        bytes32[] memory operatorSetKeys = _pendingOperatorSets.values();

        operatorSets = new OperatorSet[](operatorSetKeys.length);

        for (uint256 i = 0; i < operatorSetKeys.length; ++i) {
            operatorSets[i] = operatorSetKeys[i].decode();
        }

        return operatorSets;
    }

    /// @inheritdoc ISlashEscrowFactory
    function getTotalPendingOperatorSets() external view returns (uint256) {
        return _pendingOperatorSets.length();
    }

    /// @inheritdoc ISlashEscrowFactory
    function isPendingOperatorSet(
        OperatorSet calldata operatorSet
    ) external view returns (bool) {
        return _pendingOperatorSets.contains(operatorSet.key());
    }

    /// @inheritdoc ISlashEscrowFactory
    function getPendingSlashIds(
        OperatorSet memory operatorSet
    ) public view returns (uint256[] memory) {
        return _pendingSlashIds[operatorSet.key()].values();
    }

    /// @inheritdoc ISlashEscrowFactory
    function getTotalPendingSlashIds(
        OperatorSet calldata operatorSet
    ) external view returns (uint256) {
        return _pendingSlashIds[operatorSet.key()].length();
    }

    /// @inheritdoc ISlashEscrowFactory
    function getPendingEscrows()
        external
        view
        returns (
            OperatorSet[] memory operatorSets,
            bool[] memory isRedistributing,
            uint256[][] memory slashIds,
            uint32[][] memory completeBlocks
        )
    {
        operatorSets = getPendingOperatorSets();
        isRedistributing = new bool[](operatorSets.length);
        slashIds = new uint256[][](operatorSets.length);
        completeBlocks = new uint32[][](operatorSets.length);

        // Populate all arrays.
        for (uint256 i = 0; i < operatorSets.length; i++) {
            // Get whether the operator set is redistributing.
            isRedistributing[i] = allocationManager.isRedistributingOperatorSet(operatorSets[i]);

            // Get the pending slash IDs for the operator set.
            slashIds[i] = getPendingSlashIds(operatorSets[i]);

            // For each slashId, get the complete block.
            completeBlocks[i] = new uint32[](slashIds[i].length);
            for (uint256 j = 0; j < slashIds[i].length; j++) {
                completeBlocks[i][j] = getBurnOrRedistributionCompleteBlock(operatorSets[i], slashIds[i][j]);
            }
        }
    }

    /// @inheritdoc ISlashEscrowFactory
    function isPendingSlashId(OperatorSet calldata operatorSet, uint256 slashId) external view returns (bool) {
        return _pendingSlashIds[operatorSet.key()].contains(slashId);
    }

    /// @inheritdoc ISlashEscrowFactory
    function getPendingStrategiesForSlashId(
        OperatorSet memory operatorSet,
        uint256 slashId
    ) public view returns (IStrategy[] memory strategies) {
        EnumerableSet.AddressSet storage pendingStrategiesForSlashId =
            _pendingStrategiesForSlashId[operatorSet.key()][slashId];

        uint256 length = pendingStrategiesForSlashId.length();

        strategies = new IStrategy[](length);

        for (uint256 i = 0; i < length; ++i) {
            address strategy = pendingStrategiesForSlashId.at(i);

            strategies[i] = IStrategy(strategy);
        }
    }

    /// @inheritdoc ISlashEscrowFactory
    function getPendingStrategiesForSlashIds(
        OperatorSet memory operatorSet
    ) public view returns (IStrategy[][] memory strategies) {
        EnumerableSet.UintSet storage pendingSlashIds = _pendingSlashIds[operatorSet.key()];

        uint256 length = pendingSlashIds.length();

        strategies = new IStrategy[][](length);

        for (uint256 i = 0; i < length; ++i) {
            strategies[i] = getPendingStrategiesForSlashId(operatorSet, pendingSlashIds.at(i));
        }
    }

    /// @inheritdoc ISlashEscrowFactory
    function getTotalPendingStrategiesForSlashId(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (uint256) {
        return _pendingStrategiesForSlashId[operatorSet.key()][slashId].length();
    }

    /// @inheritdoc ISlashEscrowFactory
    function getPendingUnderlyingAmountForStrategy(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external view returns (uint256) {
        return strategy.underlyingToken().balanceOf(address(getSlashEscrow(operatorSet, slashId)));
    }

    /// @inheritdoc ISlashEscrowFactory
    function isBurnOrRedistributionPaused(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) public view returns (bool) {
        return _paused[operatorSet.key()][slashId];
    }

    /// @inheritdoc ISlashEscrowFactory
    function getBurnOrRedistributionStartBlock(
        OperatorSet memory operatorSet,
        uint256 slashId
    ) public view returns (uint256) {
        return _slashIdToStartBlock[operatorSet.key()][slashId];
    }

    /// @inheritdoc ISlashEscrowFactory
    function getBurnOrRedistributionCompleteBlock(
        OperatorSet memory operatorSet,
        uint256 slashId
    ) public view returns (uint32) {
        IStrategy[] memory strategies = getPendingStrategiesForSlashId(operatorSet, slashId);

        // Loop through all strategies and return the max delay
        uint32 maxStrategyDelay;
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint32 delay = getStrategyBurnOrRedistributionDelay(IStrategy(address(strategies[i])));
            if (delay > maxStrategyDelay) {
                maxStrategyDelay = delay;
            }
        }

        // The escrow can be released once the max strategy delay has elapsed.
        return uint32(getBurnOrRedistributionStartBlock(operatorSet, slashId) + maxStrategyDelay + 1);
    }

    /// @inheritdoc ISlashEscrowFactory
    function getStrategyBurnOrRedistributionDelay(
        IStrategy strategy
    ) public view returns (uint32) {
        uint32 globalDelay = _globalBurnOrRedistributionDelayBlocks;
        uint32 strategyDelay = _strategyBurnOrRedistributionDelayBlocks[address(strategy)];

        // Return whichever delay is greater.
        return strategyDelay > globalDelay ? strategyDelay : globalDelay;
    }

    /// @inheritdoc ISlashEscrowFactory
    function getGlobalBurnOrRedistributionDelay() external view returns (uint32) {
        return _globalBurnOrRedistributionDelayBlocks;
    }

    /// @inheritdoc ISlashEscrowFactory
    function computeSlashEscrowSalt(OperatorSet calldata operatorSet, uint256 slashId) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(operatorSet.key(), slashId));
    }

    /// @inheritdoc ISlashEscrowFactory
    function isDeployedSlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) public view returns (bool) {
        return isDeployedSlashEscrow(getSlashEscrow(operatorSet, slashId));
    }

    /// @inheritdoc ISlashEscrowFactory
    function isDeployedSlashEscrow(
        ISlashEscrow slashEscrow
    ) public view returns (bool) {
        return address(slashEscrow).code.length != 0;
    }

    /// @inheritdoc ISlashEscrowFactory
    function getSlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) public view returns (ISlashEscrow) {
        return ISlashEscrow(
            address(slashEscrowImplementation).predictDeterministicAddress({
                salt: computeSlashEscrowSalt(operatorSet, slashId),
                deployer: address(this)
            })
        );
    }
}
