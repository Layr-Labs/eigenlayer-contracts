// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../interfaces/IEigenPodManager.sol";
import "../permissions/Pausable.sol";
import "./StrategyManagerStorage.sol";
import "../libraries/EIP1271SignatureUtils.sol";

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
    StrategyManagerStorage
{
    using SafeERC20 for IERC20;

    // index for flag that pauses deposits when set
    uint8 internal constant PAUSED_DEPOSITS = 0;

    // chain id at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    modifier onlyStrategyWhitelister() {
        require(
            msg.sender == strategyWhitelister, "StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister"
        );
        _;
    }

    modifier onlyStrategiesWhitelistedForDeposit(IStrategy strategy) {
        require(
            strategyIsWhitelistedForDeposit[strategy],
            "StrategyManager.onlyStrategiesWhitelistedForDeposit: strategy not whitelisted"
        );
        _;
    }

    modifier onlyDelegationManager() {
        require(msg.sender == address(delegation), "StrategyManager.onlyDelegationManager: not the DelegationManager");
        _;
    }

    /**
     * @param _delegation The delegation contract of EigenLayer.
     * @param _slasher The primary slashing contract of EigenLayer.
     * @param _eigenPodManager The contract that keeps track of EigenPod stakes for restaking beacon chain ether.
     */
    constructor(
        IDelegationManager _delegation,
        IEigenPodManager _eigenPodManager,
        ISlasher _slasher,
        IAVSDirectory _avsDirectory
    ) StrategyManagerStorage(_delegation, _eigenPodManager, _slasher, _avsDirectory) {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    // EXTERNAL FUNCTIONS

    /**
     * @notice Initializes the strategy manager contract. Sets the `pauserRegistry` (currently **not** modifiable after being set),
     * and transfers contract ownership to the specified `initialOwner`.
     * @param _pauserRegistry Used for access control of pausing.
     * @param initialOwner Ownership of this contract is transferred to this address.
     * @param initialStrategyWhitelister The initial value of `strategyWhitelister` to set.
     * @param  initialPausedStatus The initial value of `_paused` to set.
     */
    function initialize(
        address initialOwner,
        address initialStrategyWhitelister,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
        _setStrategyWhitelister(initialStrategyWhitelister);
    }

    /**
     * @notice Deposits `amount` of `token` into the specified `strategy`, with the resultant shares credited to `msg.sender`
     * @param strategy is the specified strategy where deposit is to be made,
     * @param token is the denomination in which the deposit is to be made,
     * @param amount is the amount of token to be deposited in the strategy by the staker
     * @return scaledShares The amount of new scaled shares in the `strategy` created as part of the action.
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     *
     * WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
     *          where the token balance and corresponding strategy shares are not in sync upon reentrancy.
     */
    function depositIntoStrategy(
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (uint256 scaledShares) {
        scaledShares = _depositIntoStrategy(msg.sender, strategy, token, amount);
    }

    /**
     * @notice Used for depositing an asset into the specified strategy with the resultant shares credited to `staker`,
     * who must sign off on the action.
     * Note that the assets are transferred out/from the `msg.sender`, not from the `staker`; this function is explicitly designed
     * purely to help one address deposit 'for' another.
     * @param strategy is the specified strategy where deposit is to be made,
     * @param token is the denomination in which the deposit is to be made,
     * @param amount is the amount of token to be deposited in the strategy by the staker
     * @param staker the staker that the deposited assets will be credited to
     * @param expiry the timestamp at which the signature expires
     * @param signature is a valid signature from the `staker`. either an ECDSA signature if the `staker` is an EOA, or data to forward
     * following EIP-1271 if the `staker` is a contract
     * @return scaledShares The amount of new scaled shares in the `strategy` created as part of the action.
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     * @dev A signature is required for this function to eliminate the possibility of griefing attacks, specifically those
     * targeting stakers who may be attempting to undelegate.
     * @dev Cannot be called if thirdPartyTransfersForbidden is set to true for this strategy
     *
     *  WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
     *          where the token balance and corresponding strategy shares are not in sync upon reentrancy
     */
    function depositIntoStrategyWithSignature(
        IStrategy strategy,
        IERC20 token,
        uint256 amount,
        address staker,
        uint256 expiry,
        bytes memory signature
    ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (uint256 scaledShares) {
        require(
            !thirdPartyTransfersForbidden[strategy],
            "StrategyManager.depositIntoStrategyWithSignature: third transfers disabled"
        );
        require(expiry >= block.timestamp, "StrategyManager.depositIntoStrategyWithSignature: signature expired");
        // calculate struct hash, then increment `staker`'s nonce
        uint256 nonce = nonces[staker];
        bytes32 structHash = keccak256(abi.encode(DEPOSIT_TYPEHASH, staker, strategy, token, amount, nonce, expiry));
        unchecked {
            nonces[staker] = nonce + 1;
        }

        // calculate the digest hash
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));

        /**
         * check validity of signature:
         * 1) if `staker` is an EOA, then `signature` must be a valid ECDSA signature from `staker`,
         * indicating their intention for this action
         * 2) if `staker` is a contract, then `signature` will be checked according to EIP-1271
         */
        EIP1271SignatureUtils.checkSignature_EIP1271(staker, digestHash, signature);

        // deposit the tokens (from the `msg.sender`) and credit the new shares to the `staker`
        scaledShares = _depositIntoStrategy(staker, strategy, token, amount);
    }

    /// @notice Used by the DelegationManager to remove a Staker's scaled shares from a particular strategy when entering the withdrawal queue
    function removeScaledShares(address staker, IStrategy strategy, uint256 scaledShares) external onlyDelegationManager {
        _removeScaledShares(staker, strategy, scaledShares);
    }

    /// @notice Used by the DelegationManager to award a Staker some scaled shares that have passed through the withdrawal queue
    function addScaledShares(
        address staker,
        IERC20 token,
        IStrategy strategy,
        uint256 scaledShares
    ) external onlyDelegationManager {
        _addScaledShares(staker, token, strategy, scaledShares);
    }

    /// @notice Used by the DelegationManager to convert withdrawn shares to tokens and send them to a recipient
    /// Assumes that shares being passed in have already been descaled accordingly to account for any slashing
    /// and are the `real` shares in the strategy to withdraw
    function withdrawSharesAsTokens(
        address recipient,
        IStrategy strategy,
        uint256 strategyShares,
        IERC20 token
    ) external onlyDelegationManager {
        strategy.withdraw(recipient, token, strategyShares);
    }

    /**
     * If true for a strategy, a user cannot depositIntoStrategyWithSignature into that strategy for another staker
     * and also when performing DelegationManager.queueWithdrawals, a staker can only withdraw to themselves.
     * Defaulted to false for all existing strategies.
     * @param strategy The strategy to set `thirdPartyTransfersForbidden` value to
     * @param value bool value to set `thirdPartyTransfersForbidden` to
     */
    function setThirdPartyTransfersForbidden(IStrategy strategy, bool value) external onlyStrategyWhitelister {
        _setThirdPartyTransfersForbidden(strategy, value);
    }

    /**
     * @notice Owner-only function to change the `strategyWhitelister` address.
     * @param newStrategyWhitelister new address for the `strategyWhitelister`.
     */
    function setStrategyWhitelister(address newStrategyWhitelister) external onlyOwner {
        _setStrategyWhitelister(newStrategyWhitelister);
    }

    /**
     * @notice Owner-only function that adds the provided Strategies to the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToWhitelist Strategies that will be added to the `strategyIsWhitelistedForDeposit` mapping (if they aren't in it already)
     * @param thirdPartyTransfersForbiddenValues bool values to set `thirdPartyTransfersForbidden` to for each strategy
     */
    function addStrategiesToDepositWhitelist(
        IStrategy[] calldata strategiesToWhitelist,
        bool[] calldata thirdPartyTransfersForbiddenValues
    ) external onlyStrategyWhitelister {
        require(
            strategiesToWhitelist.length == thirdPartyTransfersForbiddenValues.length,
            "StrategyManager.addStrategiesToDepositWhitelist: array lengths do not match"
        );
        uint256 strategiesToWhitelistLength = strategiesToWhitelist.length;
        for (uint256 i = 0; i < strategiesToWhitelistLength;) {
            // change storage and emit event only if strategy is not already in whitelist
            if (!strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]]) {
                strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]] = true;
                emit StrategyAddedToDepositWhitelist(strategiesToWhitelist[i]);
                _setThirdPartyTransfersForbidden(strategiesToWhitelist[i], thirdPartyTransfersForbiddenValues[i]);
            }
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Owner-only function that removes the provided Strategies from the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToRemoveFromWhitelist Strategies that will be removed to the `strategyIsWhitelistedForDeposit` mapping (if they are in it)
     */
    function removeStrategiesFromDepositWhitelist(IStrategy[] calldata strategiesToRemoveFromWhitelist)
        external
        onlyStrategyWhitelister
    {
        uint256 strategiesToRemoveFromWhitelistLength = strategiesToRemoveFromWhitelist.length;
        for (uint256 i = 0; i < strategiesToRemoveFromWhitelistLength;) {
            // change storage and emit event only if strategy is already in whitelist
            if (strategyIsWhitelistedForDeposit[strategiesToRemoveFromWhitelist[i]]) {
                strategyIsWhitelistedForDeposit[strategiesToRemoveFromWhitelist[i]] = false;
                emit StrategyRemovedFromDepositWhitelist(strategiesToRemoveFromWhitelist[i]);
                // Set mapping value to default false value
                _setThirdPartyTransfersForbidden(strategiesToRemoveFromWhitelist[i], false);
            }
            unchecked {
                ++i;
            }
        }
    }

    // INTERNAL FUNCTIONS

    /**
     * @notice This function adds `scaledShares` for a given `strategy` to the `staker` and runs through the necessary update logic.
     * @param staker The address to add scaledShares to
     * @param token The token that is being deposited (used for indexing)
     * @param strategy The Strategy in which the `staker` is receiving scaledShares
     * @param scaledShares The amount of scaled shares to grant to the `staker`
     * @dev In particular, this function calls `delegation.increaseDelegatedScaledShares(staker, strategy, scaledShares)` to ensure that all
     * delegated scaledShares are tracked, increases the stored share amount in `stakerStrategyScaledShares[staker][strategy]`, and adds `strategy`
     * to the `staker`'s list of strategies, if it is not in the list already.
     */
    function _addScaledShares(address staker, IERC20 token, IStrategy strategy, uint256 scaledShares) internal {
        // sanity checks on inputs
        require(staker != address(0), "StrategyManager._addScaledShares: staker cannot be zero address");
        require(scaledShares != 0, "StrategyManager._addScaledShares: shares should not be zero!");

        // if they dont have existing scaled shares of this strategy, add it to their strats
        if (stakerStrategyScaledShares[staker][strategy] == 0) {
            require(
                stakerStrategyList[staker].length < MAX_STAKER_STRATEGY_LIST_LENGTH,
                "StrategyManager._addScaledShares: deposit would exceed MAX_STAKER_STRATEGY_LIST_LENGTH"
            );
            stakerStrategyList[staker].push(strategy);
        }

        // add the returned scaled shares to their existing scaled shares for this strategy
        stakerStrategyScaledShares[staker][strategy] += scaledShares;

        emit Deposit(staker, token, strategy, scaledShares);
    }

    /**
     * @notice Internal function in which `amount` of ERC20 `token` is transferred from `msg.sender` to the Strategy-type contract
     * `strategy`, with the resulting scaledShares credited to `staker`.
     * @param staker The address that will be credited with the new scaledShares.
     * @param strategy The Strategy contract to deposit into.
     * @param token The ERC20 token to deposit.
     * @param amount The amount of `token` to deposit.
     * @return scaledShares The amount of *new* scaled shares in `strategy` that have been credited to the `staker`.
     */
    function _depositIntoStrategy(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) internal onlyStrategiesWhitelistedForDeposit(strategy) returns (uint256 scaledShares) {
        // transfer tokens from the sender to the strategy
        token.safeTransferFrom(msg.sender, address(strategy), amount);

        // deposit the assets into the specified strategy and get the equivalent amount of shares in that strategy
        uint256 shares = strategy.deposit(token, amount);

        // scale strategy shares before storing
        scaledShares = delegation.getStakerScaledShares(staker, strategy, shares);

        // add the returned scaled shares to the staker's existing scaled shares for this strategy
        _addScaledShares(staker, token, strategy, scaledShares);

        // Increase scaled shares delegated to operator, if needed
        delegation.increaseDelegatedScaledShares(staker, strategy, scaledShares);

        return scaledShares;
    }

    /**
     * @notice Decreases the scaled shares that `staker` holds in `strategy` by `scaledSharesAmount`.
     * @param staker The address to decrement scaled shares from
     * @param strategy The strategy for which the `staker`'s scaled shares are being decremented
     * @param scaledSharesAmount The amount of scaled shares to decrement
     * @dev If the amount of scaled shares represents all of the staker`s shares in said strategy,
     * then the strategy is removed from stakerStrategyList[staker] and 'true' is returned. Otherwise 'false' is returned.
     */
    function _removeScaledShares(address staker, IStrategy strategy, uint256 scaledSharesAmount) internal returns (bool) {
        // sanity checks on inputs
        require(scaledSharesAmount != 0, "StrategyManager._removeScaledShares: scaledSharesAmount should not be zero!");

        //check that the user has sufficient scaled shares
        uint256 userShares = stakerStrategyScaledShares[staker][strategy];

        require(scaledSharesAmount <= userShares, "StrategyManager._removeScaledShares: scaledSharesAmount too high");
        //unchecked arithmetic since we just checked this above
        unchecked {
            userShares = userShares - scaledSharesAmount;
        }

        // subtract the shares from the staker's existing shares for this strategy
        stakerStrategyScaledShares[staker][strategy] = userShares;

        // if no existing scaled shares, remove the strategy from the staker's dynamic array of strategies
        if (userShares == 0) {
            _removeStrategyFromStakerStrategyList(staker, strategy);

            // return true in the event that the strategy was removed from stakerStrategyList[staker]
            return true;
        }
        // return false in the event that the strategy was *not* removed from stakerStrategyList[staker]
        return false;
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
        for (; j < stratsLength;) {
            if (stakerStrategyList[staker][j] == strategy) {
                //replace the strategy with the last strategy in the list
                stakerStrategyList[staker][j] = stakerStrategyList[staker][stakerStrategyList[staker].length - 1];
                break;
            }
            unchecked {
                ++j;
            }
        }
        // if we didn't find the strategy, revert
        require(j != stratsLength, "StrategyManager._removeStrategyFromStakerStrategyList: strategy not found");
        // pop off the last entry in the list of strategies
        stakerStrategyList[staker].pop();
    }

    /**
     * @notice Internal function for modifying `thirdPartyTransfersForbidden`.
     * Used inside of the `setThirdPartyTransfersForbidden` and `addStrategiesToDepositWhitelist` functions.
     * @param strategy The strategy to set `thirdPartyTransfersForbidden` value to
     * @param value bool value to set `thirdPartyTransfersForbidden` to
     */
    function _setThirdPartyTransfersForbidden(IStrategy strategy, bool value) internal {
        emit UpdatedThirdPartyTransfersForbidden(strategy, value);
        thirdPartyTransfersForbidden[strategy] = value;
    }

    /**
     * @notice Internal function for modifying the `strategyWhitelister`. Used inside of the `setStrategyWhitelister` and `initialize` functions.
     * @param newStrategyWhitelister The new address for the `strategyWhitelister` to take.
     */
    function _setStrategyWhitelister(address newStrategyWhitelister) internal {
        emit StrategyWhitelisterChanged(strategyWhitelister, newStrategyWhitelister);
        strategyWhitelister = newStrategyWhitelister;
    }

    // VIEW FUNCTIONS
    /**
     * @notice Get all details on the staker's deposits and corresponding scaled shares
     * @param staker The staker of interest, whose deposits this function will fetch
     * @return (staker's strategies, scaled shares in these strategies)
     */
    function getDeposits(address staker) external view returns (IStrategy[] memory, uint256[] memory) {
        uint256 strategiesLength = stakerStrategyList[staker].length;
        uint256[] memory scaledShares = new uint256[](strategiesLength);

        for (uint256 i = 0; i < strategiesLength;) {
            scaledShares[i] = stakerStrategyScaledShares[staker][stakerStrategyList[staker][i]];
            unchecked {
                ++i;
            }
        }
        return (stakerStrategyList[staker], scaledShares);
    }

    /// @notice Returns the current shares of `user` in `strategy`
    function stakerStrategyShares(
        address user,
        IStrategy strategy
    ) public view returns (uint256 shares) {
        uint256 scaledShares = stakerStrategyScaledShares[user][strategy];
        return delegation.getStakerShares(user, strategy, scaledShares);
    }

    /// @notice Simple getter function that returns `stakerStrategyList[staker].length`.
    function stakerStrategyListLength(address staker) external view returns (uint256) {
        return stakerStrategyList[staker].length;
    }

    /**
     * @notice Getter function for the current EIP-712 domain separator for this contract.
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     */
    function domainSeparator() public view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return _calculateDomainSeparator();
        }
    }

    // @notice Internal function for calculating the current domain separator of this contract
    function _calculateDomainSeparator() internal view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
    }
}
