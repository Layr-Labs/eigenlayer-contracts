// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

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
    using SlashingLib for *;
    using SafeERC20 for IERC20;

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
     * @param _eigenPodManager The contract that keeps track of EigenPod stakes for restaking beacon chain ether.
     */
    constructor(
        IDelegationManager _delegation,
        IEigenPodManager _eigenPodManager,
        IAVSDirectory _avsDirectory
    ) StrategyManagerStorage(_delegation, _eigenPodManager, _avsDirectory) {
        _disableInitializers();
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
     * @return shares The amount of new shares in the `strategy` created as part of the action.
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     *
     * WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
     *          where the token balance and corresponding strategy shares are not in sync upon reentrancy.
     */
    function depositIntoStrategy(
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (OwnedShares shares) {
        shares = _depositIntoStrategy(msg.sender, strategy, token, amount);
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
     * @return shares The amount of new shares in the `strategy` created as part of the action.
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     * @dev A signature is required for this function to eliminate the possibility of griefing attacks, specifically those
     * targeting stakers who may be attempting to undelegate.
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
    ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (OwnedShares shares) {
        require(expiry >= block.timestamp, SignatureExpired());
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
        shares = _depositIntoStrategy(staker, strategy, token, amount);
    }

    /// @notice Used by the DelegationManager to remove a Staker's shares from a particular strategy when entering the withdrawal queue
    function removeShares(address staker, IStrategy strategy, Shares shares) external onlyDelegationManager {
        _removeShares(staker, strategy, shares);
    }

    /// @notice Used by the DelegationManager to award a Staker some shares that have passed through the withdrawal queue
    function addOwnedShares(
        address staker,
        IStrategy strategy,
        IERC20 token,
        OwnedShares ownedShares
    ) external onlyDelegationManager {
        _addOwnedShares(staker, token, strategy, ownedShares);
    }

    /// @notice Used by the DelegationManager to convert withdrawn shares to tokens and send them to a recipient
    /// Assumes that shares being passed in have already been accounted for any slashing
    /// and are the `real` shares in the strategy to withdraw
    function withdrawSharesAsTokens(
        address staker,
        IStrategy strategy,
        IERC20 token,
        OwnedShares ownedShares
    ) external onlyDelegationManager {
        strategy.withdraw(staker, token, ownedShares.unwrap());
    }

    /**
     * @notice Owner-only function to change the `strategyWhitelister` address.
     * @param newStrategyWhitelister new address for the `strategyWhitelister`.
     */
    function setStrategyWhitelister(
        address newStrategyWhitelister
    ) external onlyOwner {
        _setStrategyWhitelister(newStrategyWhitelister);
    }

    /**
     * @notice Owner-only function that adds the provided Strategies to the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToWhitelist Strategies that will be added to the `strategyIsWhitelistedForDeposit` mapping (if they aren't in it already)
     */
    function addStrategiesToDepositWhitelist(
        IStrategy[] calldata strategiesToWhitelist
    ) external onlyStrategyWhitelister {
        uint256 strategiesToWhitelistLength = strategiesToWhitelist.length;
        for (uint256 i = 0; i < strategiesToWhitelistLength; ++i) {
            // change storage and emit event only if strategy is not already in whitelist
            if (!strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]]) {
                strategyIsWhitelistedForDeposit[strategiesToWhitelist[i]] = true;
                emit StrategyAddedToDepositWhitelist(strategiesToWhitelist[i]);
            }
        }
    }

    /**
     * @notice Owner-only function that removes the provided Strategies from the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToRemoveFromWhitelist Strategies that will be removed to the `strategyIsWhitelistedForDeposit` mapping (if they are in it)
     */
    function removeStrategiesFromDepositWhitelist(
        IStrategy[] calldata strategiesToRemoveFromWhitelist
    ) external onlyStrategyWhitelister {
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
     * @param token The token that is being deposited (used for indexing)
     * @param strategy The Strategy in which the `staker` is receiving shares
     * @param ownedShares The amount of ownedShares to grant to the `staker`
     * @dev In particular, this function calls `delegation.increaseDelegatedShares(staker, strategy, shares)` to ensure that all
     * delegated shares are tracked, increases the stored share amount in `stakerStrategyShares[staker][strategy]`, and adds `strategy`
     * to the `staker`'s list of strategies, if it is not in the list already.
     */
    function _addOwnedShares(address staker, IERC20 token, IStrategy strategy, OwnedShares ownedShares) internal {
        // sanity checks on inputs
        require(staker != address(0), StakerAddressZero());
        require(ownedShares.unwrap() != 0, SharesAmountZero());

        Shares existingShares = stakerStrategyShares[staker][strategy];

        // if they dont have existing ownedShares of this strategy, add it to their strats
        if (existingShares.unwrap() == 0) {
            require(stakerStrategyList[staker].length < MAX_STAKER_STRATEGY_LIST_LENGTH, MaxStrategiesExceeded());
            stakerStrategyList[staker].push(strategy);
        }

        // add the returned ownedShares to their existing shares for this strategy
        stakerStrategyShares[staker][strategy] = existingShares.add(ownedShares.unwrap()).wrapShares();

        // Increase shares delegated to operator, if needed
        delegation.increaseDelegatedShares({
            staker: staker,
            strategy: strategy,
            existingShares: existingShares,
            addedOwnedShares: ownedShares
        });

        emit Deposit(staker, token, strategy, ownedShares);
    }

    /**
     * @notice Internal function in which `amount` of ERC20 `token` is transferred from `msg.sender` to the Strategy-type contract
     * `strategy`, with the resulting shares credited to `staker`.
     * @param staker The address that will be credited with the new shares.
     * @param strategy The Strategy contract to deposit into.
     * @param token The ERC20 token to deposit.
     * @param amount The amount of `token` to deposit.
     * @return ownedShares The amount of *new* ownedShares in `strategy` that have been credited to the `staker`.
     */
    function _depositIntoStrategy(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) internal onlyStrategiesWhitelistedForDeposit(strategy) returns (OwnedShares ownedShares) {
        // transfer tokens from the sender to the strategy
        token.safeTransferFrom(msg.sender, address(strategy), amount);

        // deposit the assets into the specified strategy and get the equivalent amount of shares in that strategy
        ownedShares = strategy.deposit(token, amount).wrapOwned();

        // add the returned shares to the staker's existing shares for this strategy
        _addOwnedShares(staker, token, strategy, ownedShares);

        return ownedShares;
    }

    /**
     * @notice Decreases the shares that `staker` holds in `strategy` by `shareAmount`.
     * @param staker The address to decrement shares from
     * @param strategy The strategy for which the `staker`'s shares are being decremented
     * @param shareAmount The amount of shares to decrement
     * @dev If the amount of shares represents all of the staker`s shares in said strategy,
     * then the strategy is removed from stakerStrategyList[staker] and 'true' is returned. Otherwise 'false' is returned.
     */
    function _removeShares(address staker, IStrategy strategy, Shares shareAmount) internal returns (bool) {
        // sanity checks on inputs
        require(shareAmount.unwrap() != 0, SharesAmountZero());

        //check that the user has sufficient shares
        Shares userShares = stakerStrategyShares[staker][strategy];

        require(shareAmount.unwrap() <= userShares.unwrap(), SharesAmountTooHigh());

        userShares = userShares.sub(shareAmount.unwrap()).wrapShares();

        // subtract the shares from the staker's existing shares for this strategy
        stakerStrategyShares[staker][strategy] = userShares;

        // if no existing shares, remove the strategy from the staker's dynamic array of strategies
        if (userShares.unwrap() == 0) {
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

    /**
     * @notice Get all details on the staker's deposits and corresponding shares
     * @param staker The staker of interest, whose deposits this function will fetch
     * @return (staker's strategies, shares in these strategies)
     */
    function getDeposits(
        address staker
    ) external view returns (IStrategy[] memory, Shares[] memory) {
        uint256 strategiesLength = stakerStrategyList[staker].length;
        Shares[] memory shares = new Shares[](strategiesLength);

        for (uint256 i = 0; i < strategiesLength; ++i) {
            shares[i] = stakerStrategyShares[staker][stakerStrategyList[staker][i]];
        }
        return (stakerStrategyList[staker], shares);
    }

    function getStakerStrategyList(
        address staker
    ) external view returns (IStrategy[] memory) {
        return stakerStrategyList[staker];
    }

    /// @notice Simple getter function that returns `stakerStrategyList[staker].length`.
    function stakerStrategyListLength(
        address staker
    ) external view returns (uint256) {
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
