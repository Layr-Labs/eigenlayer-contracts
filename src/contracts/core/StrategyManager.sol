// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../mixins/SignatureUtilsMixin.sol";
import "../interfaces/IEigenPodManager.sol";
import "../permissions/Pausable.sol";
import "./StrategyManagerStorage.sol";

/**
 * @title The primary entry- and exit-point for funds into and out of EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This contract is for managing deposits in different strategies. The main
 * functionalities are:
 * - adding and removing strategies that any delegator can deposit into
 * - enabling deposit of assets into specified strategy(s)
 */
contract StrategyManager is
    Initializable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    Pausable,
    StrategyManagerStorage,
    SignatureUtilsMixin
{
    using OperatorSetLib for *;
    using SlashingLib for *;
    using SafeERC20 for IERC20;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.UintSet;

    modifier onlyStrategyWhitelister() {
        require(msg.sender == strategyWhitelister, OnlyStrategyWhitelister());
        _;
    }

    modifier onlyStrategiesWhitelistedForDeposit(
        IStrategy strategy
    ) {
        require(strategyIsWhitelistedForDeposit[strategy], StrategyNotWhitelisted());
        _;
    }

    modifier onlyDelegationManager() {
        require(msg.sender == address(delegation), OnlyDelegationManager());
        _;
    }

    /**
     * @param _delegation The delegation contract of EigenLayer.
     */
    constructor(
        IAllocationManager _allocationManager,
        IDelegationManager _delegation,
        IPauserRegistry _pauserRegistry,
        string memory _version
    ) StrategyManagerStorage(_allocationManager, _delegation) Pausable(_pauserRegistry) SignatureUtilsMixin(_version) {
        _disableInitializers();
    }

    // EXTERNAL FUNCTIONS

    /**
     * @notice Initializes the strategy manager contract. Sets the `pauserRegistry` (currently **not** modifiable after being set),
     * and transfers contract ownership to the specified `initialOwner`.
     * @param initialOwner Ownership of this contract is transferred to this address.
     * @param initialStrategyWhitelister The initial value of `strategyWhitelister` to set.
     */
    function initialize(
        address initialOwner,
        address initialStrategyWhitelister,
        uint256 initialPausedStatus
    ) external initializer {
        _setPausedStatus(initialPausedStatus);
        _transferOwnership(initialOwner);
        _setStrategyWhitelister(initialStrategyWhitelister);
    }

    /// @inheritdoc IStrategyManager
    function depositIntoStrategy(
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (uint256 depositShares) {
        depositShares = _depositIntoStrategy(msg.sender, strategy, token, amount);
    }

    /// @inheritdoc IStrategyManager
    function depositIntoStrategyWithSignature(
        IStrategy strategy,
        IERC20 token,
        uint256 amount,
        address staker,
        uint256 expiry,
        bytes memory signature
    ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (uint256 depositShares) {
        // Cache staker's nonce to avoid sloads.
        uint256 nonce = nonces[staker];
        // Assert that the signature is valid.
        _checkIsValidSignatureNow({
            signer: staker,
            signableDigest: calculateStrategyDepositDigestHash(staker, strategy, token, amount, nonce, expiry),
            signature: signature,
            expiry: expiry
        });
        // Increment the nonce for the staker.
        unchecked {
            nonces[staker] = nonce + 1;
        }
        // deposit the tokens (from the `msg.sender`) and credit the new shares to the `staker`
        depositShares = _depositIntoStrategy(staker, strategy, token, amount);
    }

    /// @inheritdoc IShareManager
    function removeDepositShares(
        address staker,
        IStrategy strategy,
        uint256 depositSharesToRemove
    ) external onlyDelegationManager nonReentrant returns (uint256) {
        return _removeDepositShares(staker, strategy, depositSharesToRemove);
    }

    /// @inheritdoc IShareManager
    function addShares(
        address staker,
        IStrategy strategy,
        uint256 shares
    ) external onlyDelegationManager nonReentrant returns (uint256, uint256) {
        return _addShares(staker, strategy, shares);
    }

    /// @inheritdoc IShareManager
    function withdrawSharesAsTokens(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 shares
    ) external onlyDelegationManager nonReentrant {
        strategy.withdraw(staker, token, shares);
    }

    /// @inheritdoc IShareManager
    function increaseBurnOrRedistributableShares(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy,
        uint256 sharesToBurn
    ) external onlyDelegationManager nonReentrant {
        // Create storage pointer for readability.
        EnumerableMap.AddressToUintMap storage burnOrRedistributableShares =
            _burnOrRedistributableShares[operatorSet.key()][slashId];

        // Sanity check that the strategy is not already in the slash's burn or redistributable shares.
        // This should never happen because the `AllocationManager` ensures that strategies for a given slash are unique.
        // Add the shares to the operator set's burn or redistributable shares.
        require(burnOrRedistributableShares.set(address(strategy), sharesToBurn), StrategyAlreadyInSlash());

        // NOTE: Duplicate operator sets and slash ids will not revert, but will not be added.
        _pendingOperatorSets.add(operatorSet.key());
        _pendingSlashIds[operatorSet.key()].add(slashId);

        emit BurnOrRedistributableSharesIncreased(operatorSet, slashId, strategy, sharesToBurn);
    }

    /// @inheritdoc IStrategyManager
    function clearBurnOrRedistributableShares(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external nonReentrant returns (uint256[] memory) {
        // Get the strategies to clear.
        address[] memory strategies = _burnOrRedistributableShares[operatorSet.key()][slashId].keys();
        uint256 length = strategies.length;
        uint256[] memory amounts = new uint256[](length);

        // Note: We don't need to iterate backwards since we're indexing into the `EnumerableMap` directly.
        for (uint256 i = 0; i < length; ++i) {
            amounts[i] = _clearBurnOrRedistributableShares({
                operatorSet: operatorSet,
                slashId: slashId,
                strategy: IStrategy(strategies[i]),
                recipient: allocationManager.getRedistributionRecipient(operatorSet)
            });
        }

        return amounts;
    }

    /// @inheritdoc IStrategyManager
    function clearBurnOrRedistributableSharesByStrategy(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external nonReentrant returns (uint256) {
        return _clearBurnOrRedistributableShares({
            operatorSet: operatorSet,
            slashId: slashId,
            strategy: strategy,
            recipient: allocationManager.getRedistributionRecipient(operatorSet)
        });
    }

    /// @inheritdoc IStrategyManager
    function burnShares(
        IStrategy strategy
    ) external nonReentrant {
        (, uint256 sharesToBurn) = EnumerableMap.tryGet(burnableShares, address(strategy));
        EnumerableMap.remove(burnableShares, address(strategy));
        emit BurnableSharesDecreased(strategy, sharesToBurn);

        // Burning acts like withdrawing, except that the destination is to the burn address.
        // If we have no shares to burn, we don't need to call the strategy.
        if (sharesToBurn != 0) {
            strategy.withdraw(DEFAULT_BURN_ADDRESS, strategy.underlyingToken(), sharesToBurn);
        }
    }

    /// @inheritdoc IStrategyManager
    function setStrategyWhitelister(
        address newStrategyWhitelister
    ) external onlyOwner nonReentrant {
        _setStrategyWhitelister(newStrategyWhitelister);
    }

    /// @inheritdoc IStrategyManager
    function addStrategiesToDepositWhitelist(
        IStrategy[] calldata strategiesToWhitelist
    ) external onlyStrategyWhitelister nonReentrant {
        uint256 strategiesToWhitelistLength = strategiesToWhitelist.length;
        for (uint256 i = 0; i < strategiesToWhitelistLength; ++i) {
            // change storage and emit event only if strategy is not already in whitelist
            if (!strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]]) {
                strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]] = true;
                emit StrategyAddedToDepositWhitelist(strategiesToWhitelist[i]);
            }
        }
    }

    /// @inheritdoc IStrategyManager
    function removeStrategiesFromDepositWhitelist(
        IStrategy[] calldata strategiesToRemoveFromWhitelist
    ) external onlyStrategyWhitelister nonReentrant {
        uint256 strategiesToRemoveFromWhitelistLength = strategiesToRemoveFromWhitelist.length;
        for (uint256 i = 0; i < strategiesToRemoveFromWhitelistLength; ++i) {
            // change storage and emit event only if strategy is already in whitelist
            if (strategyIsWhitelistedForDeposit[strategiesToRemoveFromWhitelist[i]]) {
                strategyIsWhitelistedForDeposit[strategiesToRemoveFromWhitelist[i]] = false;
                emit StrategyRemovedFromDepositWhitelist(strategiesToRemoveFromWhitelist[i]);
            }
        }
    }

    // INTERNAL FUNCTIONS

    /**
     * @notice This function adds `shares` for a given `strategy` to the `staker` and runs through the necessary update logic.
     * @param staker The address to add shares to
     * @param strategy The Strategy in which the `staker` is receiving shares
     * @param shares The amount of shares to grant to the `staker`
     * @dev In particular, this function calls `delegation.increaseDelegatedShares(staker, strategy, shares)` to ensure that all
     * delegated shares are tracked, increases the stored share amount in `stakerDepositShares[staker][strategy]`, and adds `strategy`
     * to the `staker`'s list of strategies, if it is not in the list already.
     */
    function _addShares(address staker, IStrategy strategy, uint256 shares) internal returns (uint256, uint256) {
        // sanity checks on inputs
        require(staker != address(0), StakerAddressZero());
        require(shares != 0, SharesAmountZero());

        uint256 prevDepositShares = stakerDepositShares[staker][strategy];

        // if they dont have prevDepositShares of this strategy, add it to their strats
        if (prevDepositShares == 0) {
            require(stakerStrategyList[staker].length < MAX_STAKER_STRATEGY_LIST_LENGTH, MaxStrategiesExceeded());
            stakerStrategyList[staker].push(strategy);
        }

        // add the returned depositedShares to their existing shares for this strategy
        stakerDepositShares[staker][strategy] = prevDepositShares + shares;

        emit Deposit(staker, strategy, shares);
        return (prevDepositShares, shares);
    }

    /**
     * @notice Internal function in which `amount` of ERC20 `token` is transferred from `msg.sender` to the Strategy-type contract
     * `strategy`, with the resulting shares credited to `staker`.
     * @param staker The address that will be credited with the new shares.
     * @param strategy The Strategy contract to deposit into.
     * @param token The ERC20 token to deposit.
     * @param amount The amount of `token` to deposit.
     * @return shares The amount of *new* shares in `strategy` that have been credited to the `staker`.
     */
    function _depositIntoStrategy(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) internal onlyStrategiesWhitelistedForDeposit(strategy) returns (uint256 shares) {
        // transfer tokens from the sender to the strategy
        token.safeTransferFrom(msg.sender, address(strategy), amount);

        // deposit the assets into the specified strategy and get the equivalent amount of shares in that strategy
        shares = strategy.deposit(token, amount);

        // add the returned shares to the staker's existing shares for this strategy
        (uint256 prevDepositShares, uint256 addedShares) = _addShares(staker, strategy, shares);

        // Increase shares delegated to operator
        delegation.increaseDelegatedShares({
            staker: staker,
            strategy: strategy,
            prevDepositShares: prevDepositShares,
            addedShares: addedShares
        });

        return shares;
    }

    /**
     * @notice Decreases the shares that `staker` holds in `strategy` by `depositSharesToRemove`.
     * @param staker The address to decrement shares from
     * @param strategy The strategy for which the `staker`'s shares are being decremented
     * @param depositSharesToRemove The amount of deposit shares to decrement
     * @dev If the amount of shares represents all of the staker`s shares in said strategy,
     * then the strategy is removed from stakerStrategyList[staker] and 'true' is returned. Otherwise 'false' is returned.
     * Also returns the user's updated deposit shares after decrement.
     */
    function _removeDepositShares(
        address staker,
        IStrategy strategy,
        uint256 depositSharesToRemove
    ) internal returns (uint256) {
        // sanity checks on inputs
        require(depositSharesToRemove != 0, SharesAmountZero());

        //check that the user has sufficient shares
        uint256 userDepositShares = stakerDepositShares[staker][strategy];

        // This check technically shouldn't actually ever revert because depositSharesToRemove is already
        // checked to not exceed max amount of shares when the withdrawal was queued in the DelegationManager
        require(depositSharesToRemove <= userDepositShares, SharesAmountTooHigh());
        userDepositShares = userDepositShares - depositSharesToRemove;

        // subtract the shares from the staker's existing shares for this strategy
        stakerDepositShares[staker][strategy] = userDepositShares;

        // if no existing shares, remove the strategy from the staker's dynamic array of strategies
        if (userDepositShares == 0) {
            _removeStrategyFromStakerStrategyList(staker, strategy);
        }

        return userDepositShares;
    }

    /**
     * @notice Removes `strategy` from `staker`'s dynamic array of strategies, i.e. from `stakerStrategyList[staker]`
     * @param staker The user whose array will have an entry removed
     * @param strategy The Strategy to remove from `stakerStrategyList[staker]`
     */
    function _removeStrategyFromStakerStrategyList(address staker, IStrategy strategy) internal {
        //loop through all of the strategies, find the right one, then replace
        uint256 stratsLength = stakerStrategyList[staker].length;
        uint256 j = 0;
        for (; j < stratsLength; ++j) {
            if (stakerStrategyList[staker][j] == strategy) {
                //replace the strategy with the last strategy in the list
                stakerStrategyList[staker][j] = stakerStrategyList[staker][stakerStrategyList[staker].length - 1];
                break;
            }
        }
        // if we didn't find the strategy, revert
        require(j != stratsLength, StrategyNotFound());
        // pop off the last entry in the list of strategies
        stakerStrategyList[staker].pop();
    }

    /**
     * @notice Clears burn/redistributable shares and sends underlying tokens to recipient.
     * @param operatorSet The operator set to clear the shares for.
     * @param slashId The slash id to clear the shares for.
     * @param strategy The strategy to clear the shares for.
     * @param recipient The recipient to withdraw the shares to.
     */
    function _clearBurnOrRedistributableShares(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy,
        address recipient
    ) internal returns (uint256) {
        EnumerableMap.AddressToUintMap storage burnOrRedistributableShares =
            _burnOrRedistributableShares[operatorSet.key()][slashId];

        (, uint256 sharesToRemove) = burnOrRedistributableShares.tryGet(address(strategy));
        burnOrRedistributableShares.remove(address(strategy));

        uint256 amountOut;
        if (sharesToRemove != 0) {
            // Withdraw the shares to the burn address.
            amountOut = IStrategy(strategy).withdraw({
                recipient: recipient,
                token: IStrategy(strategy).underlyingToken(),
                amountShares: sharesToRemove
            });

            // Emit an event to notify the that burnable shares have been decreased.
            emit BurnOrRedistributableSharesDecreased(operatorSet, slashId, strategy, sharesToRemove);
        }

        uint256 remainingStrategies = burnOrRedistributableShares.keys().length;

        // If there are no more strategies to burn or redistribute...
        if (remainingStrategies == 0) {
            // Remove the slash id from the pending slash ids.
            _pendingSlashIds[operatorSet.key()].remove(slashId);

            // If there are no more pending slash ids for this operator set, remove the operator set from the pending operator sets.
            if (_pendingSlashIds[operatorSet.key()].length() == 0) {
                _pendingOperatorSets.remove(operatorSet.key());
            }
        }

        return amountOut;
    }

    /**
     * @notice Internal function for modifying the `strategyWhitelister`. Used inside of the `setStrategyWhitelister` and `initialize` functions.
     * @param newStrategyWhitelister The new address for the `strategyWhitelister` to take.
     */
    function _setStrategyWhitelister(
        address newStrategyWhitelister
    ) internal {
        emit StrategyWhitelisterChanged(strategyWhitelister, newStrategyWhitelister);
        strategyWhitelister = newStrategyWhitelister;
    }

    // VIEW FUNCTIONS

    /// @inheritdoc IStrategyManager
    function getDeposits(
        address staker
    ) external view returns (IStrategy[] memory, uint256[] memory) {
        uint256 strategiesLength = stakerStrategyList[staker].length;
        uint256[] memory depositedShares = new uint256[](strategiesLength);

        for (uint256 i = 0; i < strategiesLength; ++i) {
            depositedShares[i] = stakerDepositShares[staker][stakerStrategyList[staker][i]];
        }
        return (stakerStrategyList[staker], depositedShares);
    }

    function getStakerStrategyList(
        address staker
    ) external view returns (IStrategy[] memory) {
        return stakerStrategyList[staker];
    }

    /// @inheritdoc IStrategyManager
    function stakerStrategyListLength(
        address staker
    ) external view returns (uint256) {
        return stakerStrategyList[staker].length;
    }

    /// @inheritdoc IStrategyManager
    function calculateStrategyDepositDigestHash(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount,
        uint256 nonce,
        uint256 expiry
    ) public view returns (bytes32) {
        /// forgefmt: disable-next-item
        return _calculateSignableDigest(
            keccak256(
                abi.encode(
                    DEPOSIT_TYPEHASH, 
                    staker, 
                    strategy, 
                    token, 
                    amount, 
                    nonce, 
                    expiry
                )
            )
        );
    }

    /// @inheritdoc IStrategyManager
    function getBurnOrRedistributableShares(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (IStrategy[] memory, uint256[] memory) {
        // Store the data structure for readability
        EnumerableMap.AddressToUintMap storage burnOrRedistributableShares =
            _burnOrRedistributableShares[operatorSet.key()][slashId];

        address[] memory keys = burnOrRedistributableShares.keys();
        IStrategy[] memory strategies = new IStrategy[](keys.length);
        uint256[] memory shares = new uint256[](keys.length);

        for (uint256 i = 0; i < keys.length; ++i) {
            strategies[i] = IStrategy(keys[i]);
            (, shares[i]) = burnOrRedistributableShares.tryGet(keys[i]);
        }

        return (strategies, shares);
    }

    /// @inheritdoc IStrategyManager
    function getBurnOrRedistributableShares(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external view returns (uint256) {
        (, uint256 shares) = _burnOrRedistributableShares[operatorSet.key()][slashId].tryGet(address(strategy));
        return shares;
    }

    /// @inheritdoc IStrategyManager
    function getBurnOrRedistributableCount(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (uint256) {
        return _burnOrRedistributableShares[operatorSet.key()][slashId].length();
    }

    /// @inheritdoc IStrategyManager
    function getBurnableShares(
        IStrategy strategy
    ) external view returns (uint256) {
        (, uint256 shares) = EnumerableMap.tryGet(burnableShares, address(strategy));
        return shares;
    }

    /// @inheritdoc IStrategyManager
    function getStrategiesWithBurnableShares() external view returns (address[] memory, uint256[] memory) {
        uint256 totalEntries = EnumerableMap.length(burnableShares);

        address[] memory strategies = new address[](totalEntries);
        uint256[] memory shares = new uint256[](totalEntries);

        for (uint256 i = 0; i < totalEntries; i++) {
            (address strategy, uint256 shareAmount) = EnumerableMap.at(burnableShares, i);
            strategies[i] = strategy;
            shares[i] = shareAmount;
        }

        return (strategies, shares);
    }

    /// @inheritdoc IStrategyManager
    function getPendingOperatorSets() external view returns (OperatorSet[] memory) {
        uint256 totalEntries = _pendingOperatorSets.length();
        OperatorSet[] memory operatorSets = new OperatorSet[](totalEntries);
        for (uint256 i = 0; i < totalEntries; i++) {
            operatorSets[i] = _pendingOperatorSets.at(i).decode();
        }
        return operatorSets;
    }

    /// @inheritdoc IStrategyManager
    function getPendingSlashIds(
        OperatorSet calldata operatorSet
    ) external view returns (uint256[] memory) {
        return _pendingSlashIds[operatorSet.key()].values();
    }
}
