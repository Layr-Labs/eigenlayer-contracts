// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/ClonesUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../mixins/SemVerMixin.sol";
import "./SlashEscrowFactoryStorage.sol";

contract SlashEscrowFactory is
    Initializable,
    SlashEscrowFactoryStorage,
    OwnableUpgradeable,
    Pausable,
    SemVerMixin,
    ReentrancyGuardUpgradeable
{
    using SafeERC20 for IERC20;
    using OperatorSetLib for *;
    using EnumerableSet for *;
    using ClonesUpgradeable for address;

    modifier onlyStrategyManager() {
        require(msg.sender == address(strategyManager), OnlyStrategyManager());
        _;
    }

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
        _setGlobalEscrowDelay(initialGlobalDelayBlocks);
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
    ) external onlyStrategyManager {
        // Create storage pointers for readability.
        EnumerableSet.UintSet storage pendingSlashIds = _pendingSlashIds[operatorSet.key()];
        EnumerableSet.AddressSet storage pendingStrategiesForSlashId =
            _pendingStrategiesForSlashId[operatorSet.key()][slashId];

        // Note: Since this function can be called multiple times for the same operatorSet/slashId, we check
        // if the slash escrow is already deployed. If it is not, we deploy it and update the pending mappings.
        if (!isDeployedSlashEscrow(operatorSet, slashId)) {
            // Deploy the `SlashEscrow`.
            _deploySlashEscrow(operatorSet, slashId);

            // Update the pending mappings.
            _pendingOperatorSets.add(operatorSet.key());
            pendingSlashIds.add(slashId);
        }

        // Add the strategy to the pending strategies for the slash ID.
        pendingStrategiesForSlashId.add(address(strategy));

        // Calculate the complete block for the strategy, and fetch the current complete block.
        // The escrow delay for a strategy is determined by taking the maximum of:
        // 1. The global escrow delay (applies to all strategies).
        // 2. The strategy-specific delay (if set).
        uint32 completeBlock = uint32(block.number + getStrategyEscrowDelay(strategy) + 1);
        uint32 currentCompleteBlock = _slashIdToCompleteBlock[operatorSet.key()][slashId];

        // Only update the maturity block if the new calculated maturity block (with the strategy delay)
        // is further in the future than the currently stored value. This ensures the escrow cannot be released
        // before the longest required delay among all strategies for this slashId.
        if (completeBlock > currentCompleteBlock) {
            _slashIdToCompleteBlock[operatorSet.key()][slashId] = completeBlock;
        }

        // Emit the start escrow event. We can use the block.number here because all strategies
        // in a given operatorSet/slashId will have their escrow initiated in the same transaction.
        emit StartEscrow(operatorSet, slashId, strategy, uint32(block.number));
    }

    /// @inheritdoc ISlashEscrowFactory
    function releaseSlashEscrow(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external onlyWhenNotPaused(PAUSED_RELEASE_ESCROW) nonReentrant {
        address redistributionRecipient = allocationManager.getRedistributionRecipient(operatorSet);

        _checkReleaseSlashEscrow(operatorSet, slashId, redistributionRecipient);

        // Calling `clearBurnOrRedistributableShares` will transfer the underlying tokens to the `SlashEscrow`.
        // NOTE: While `clearBurnOrRedistributableShares` may have already been called, we call it again to ensure that the
        // underlying tokens are actually in escrow before processing and removing storage (which would otherwise prevent
        // the tokens from being released).
        strategyManager.clearBurnOrRedistributableShares(operatorSet, slashId);

        // Process the slash escrow for each strategy.
        address[] memory strategies = _pendingStrategiesForSlashId[operatorSet.key()][slashId].values();
        for (uint256 i = 0; i < strategies.length; ++i) {
            _processSlashEscrowByStrategy({
                operatorSet: operatorSet,
                slashId: slashId,
                slashEscrow: getSlashEscrow(operatorSet, slashId),
                redistributionRecipient: redistributionRecipient,
                strategy: IStrategy(strategies[i])
            });
        }

        // Update the slash escrow storage.
        _updateSlashEscrowStorage(operatorSet, slashId);
    }

    /// @inheritdoc ISlashEscrowFactory
    function releaseSlashEscrowByStrategy(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external virtual onlyWhenNotPaused(PAUSED_RELEASE_ESCROW) nonReentrant {
        address redistributionRecipient = allocationManager.getRedistributionRecipient(operatorSet);

        _checkReleaseSlashEscrow(operatorSet, slashId, redistributionRecipient);

        // Calling `clearBurnOrRedistributableSharesByStrategy` will transfer the underlying tokens to the `SlashEscrow`.
        // NOTE: While the strategy may have already been cleared, we call it again to ensure that the
        // underlying tokens are actually in escrow before processing and removing storage (which would otherwise prevent
        // the tokens from being released).
        strategyManager.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategy);

        // Release the slashEscrow.
        _processSlashEscrowByStrategy({
            operatorSet: operatorSet,
            slashId: slashId,
            slashEscrow: getSlashEscrow(operatorSet, slashId),
            redistributionRecipient: redistributionRecipient,
            strategy: strategy
        });

        // Update the slash escrow storage.
        _updateSlashEscrowStorage(operatorSet, slashId);
    }

    /**
     *
     *                         PAUSER/UNPAUSER ACTIONS
     *
     */

    /// @inheritdoc ISlashEscrowFactory
    function pauseEscrow(OperatorSet calldata operatorSet, uint256 slashId) external virtual onlyPauser {
        require(!_paused[operatorSet.key()][slashId], IPausable.InvalidNewPausedStatus());
        _paused[operatorSet.key()][slashId] = true;
        emit EscrowPaused(operatorSet, slashId);
    }

    /// @inheritdoc ISlashEscrowFactory
    function unpauseEscrow(OperatorSet calldata operatorSet, uint256 slashId) external virtual onlyUnpauser {
        require(_paused[operatorSet.key()][slashId], IPausable.InvalidNewPausedStatus());
        _paused[operatorSet.key()][slashId] = false;
        emit EscrowUnpaused(operatorSet, slashId);
    }

    /**
     *
     *                         OWNER ACTIONS
     *
     */

    /// @inheritdoc ISlashEscrowFactory
    function setGlobalEscrowDelay(
        uint32 delay
    ) external onlyOwner {
        _setGlobalEscrowDelay(delay);
    }

    /// @inheritdoc ISlashEscrowFactory
    function setStrategyEscrowDelay(IStrategy strategy, uint32 delay) external onlyOwner {
        _strategyEscrowDelayBlocks[address(strategy)] = delay;
        emit StrategyEscrowDelaySet(strategy, delay);
    }

    /**
     *
     *                         HELPERS
     *
     */

    /// @notice Checks that the slash escrow can be released.
    function _checkReleaseSlashEscrow(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        address redistributionRecipient
    ) internal view {
        // If the redistribution recipient is not the default burn address...
        if (redistributionRecipient != DEFAULT_BURN_ADDRESS) {
            require(msg.sender == redistributionRecipient, OnlyRedistributionRecipient());
        }

        // Assert that the slash ID is not paused
        require(!isEscrowPaused(operatorSet, slashId), IPausable.CurrentlyPaused());

        // Assert that the escrow delay has elapsed
        // `getEscrowCompleteBlock` returns the block number at which the escrow can be released, so
        // we require that the current block number is greater than OR equal to the complete block.
        require(block.number >= getEscrowCompleteBlock(operatorSet, slashId), EscrowDelayNotElapsed());
    }

    /// @notice Processes the slash escrow for a single strategy.
    function _processSlashEscrowByStrategy(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        ISlashEscrow slashEscrow,
        address redistributionRecipient,
        IStrategy strategy
    ) internal {
        // Create storage pointer for readability.
        EnumerableSet.AddressSet storage pendingStrategiesForSlashId =
            _pendingStrategiesForSlashId[operatorSet.key()][slashId];

        // Burn or redistribute the underlying tokens for the strategy.
        slashEscrow.releaseTokens({
            slashEscrowFactory: ISlashEscrowFactory(address(this)),
            slashEscrowImplementation: slashEscrowImplementation,
            operatorSet: operatorSet,
            slashId: slashId,
            recipient: redistributionRecipient,
            strategy: strategy
        });

        // Remove the strategy and underlying amount from the pending strategies escrow map.
        pendingStrategiesForSlashId.remove(address(strategy));
        emit EscrowComplete(operatorSet, slashId, strategy, redistributionRecipient);
    }

    function _updateSlashEscrowStorage(OperatorSet calldata operatorSet, uint256 slashId) internal {
        // Create storage pointers for readability.
        EnumerableSet.Bytes32Set storage pendingOperatorSets = _pendingOperatorSets;
        EnumerableSet.UintSet storage pendingSlashIds = _pendingSlashIds[operatorSet.key()];
        uint256 totalPendingForSlashId = _pendingStrategiesForSlashId[operatorSet.key()][slashId].length();

        // If there are no more strategies to process, remove the slash ID from the pending slash IDs set.
        if (totalPendingForSlashId == 0) {
            pendingSlashIds.remove(slashId);

            // Delete the start block for the slash ID.
            delete _slashIdToCompleteBlock[operatorSet.key()][slashId];

            // If there are no more slash IDs for the operator set, remove the operator set from the pending operator sets set.
            if (pendingSlashIds.length() == 0) {
                pendingOperatorSets.remove(operatorSet.key());
            }
        }
    }

    /// @notice Sets the global escrow delay.
    function _setGlobalEscrowDelay(
        uint32 delay
    ) internal {
        _globalEscrowDelayBlocks = delay;
        emit GlobalEscrowDelaySet(delay);
    }

    /**
     * @notice Deploys a `SlashEscrow`
     * @param operatorSet The operator set whose slash escrow is being deployed.
     * @param slashId The slash ID of the slash escrow that is being deployed.
     * @dev The slash escrow is deployed in `initiateSlashEscrow`
     */
    function _deploySlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) internal {
        address(slashEscrowImplementation).cloneDeterministic(computeSlashEscrowSalt(operatorSet, slashId));
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
                completeBlocks[i][j] = getEscrowCompleteBlock(operatorSets[i], slashIds[i][j]);
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
    function isEscrowPaused(OperatorSet calldata operatorSet, uint256 slashId) public view returns (bool) {
        return _paused[operatorSet.key()][slashId] || paused(PAUSED_RELEASE_ESCROW);
    }

    /// @inheritdoc ISlashEscrowFactory
    function getEscrowCompleteBlock(OperatorSet memory operatorSet, uint256 slashId) public view returns (uint32) {
        return _slashIdToCompleteBlock[operatorSet.key()][slashId];
    }

    /// @inheritdoc ISlashEscrowFactory
    function getStrategyEscrowDelay(
        IStrategy strategy
    ) public view returns (uint32) {
        uint32 globalDelay = _globalEscrowDelayBlocks;
        uint32 strategyDelay = _strategyEscrowDelayBlocks[address(strategy)];

        // Return whichever delay is greater.
        return strategyDelay > globalDelay ? strategyDelay : globalDelay;
    }

    /// @inheritdoc ISlashEscrowFactory
    function getGlobalEscrowDelay() external view returns (uint32) {
        return _globalEscrowDelayBlocks;
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
