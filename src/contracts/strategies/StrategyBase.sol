// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IStrategyManager.sol";
import "../permissions/Pausable.sol";
import "../mixins/SemVerMixin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

/**
 * @title Base implementation of `IStrategy` interface, designed to be inherited from by more complex strategies.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Simple, basic, "do-nothing" Strategy that holds a single underlying token and returns it on withdrawals.
 * Implements minimal versions of the IStrategy functions, this contract is designed to be inherited by
 * more complex strategies, which can then override its functions as necessary.
 * @dev Note that some functions have their mutability restricted; developers inheriting from this contract cannot broaden
 * the mutability without modifying this contract itself.
 * @dev This contract is expressly *not* intended for use with 'fee-on-transfer'-type tokens.
 * Setting the `underlyingToken` to be a fee-on-transfer token may result in improper accounting.
 * @notice This contract functions similarly to an ERC4626 vault, only without issuing a token.
 * To mitigate against the common "inflation attack" vector, we have chosen to use the 'virtual shares' mitigation route,
 * similar to [OpenZeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/extensions/ERC4626.sol).
 * We acknowledge that this mitigation has the known downside of the virtual shares causing some losses to users, which are pronounced
 * particularly in the case of the share exchange rate changing significantly, either positively or negatively.
 * For a fairly thorough discussion of this issue and our chosen mitigation strategy, we recommend reading through
 * [this thread](https://github.com/OpenZeppelin/openzeppelin-contracts/issues/3706) on the OpenZeppelin repo.
 * We specifically use a share offset of `SHARES_OFFSET` and a balance offset of `BALANCE_OFFSET`.
 */
contract StrategyBase is Initializable, Pausable, IStrategy, SemVerMixin {
    using SafeERC20 for IERC20;

    uint8 internal constant PAUSED_DEPOSITS = 0;
    uint8 internal constant PAUSED_WITHDRAWALS = 1;

    /**
     * @notice virtual shares used as part of the mitigation of the common 'share inflation' attack vector.
     * Constant value chosen to reasonably reduce attempted share inflation by the first depositor, while still
     * incurring reasonably small losses to depositors
     */
    uint256 internal constant SHARES_OFFSET = 1e3;
    /**
     * @notice virtual balance used as part of the mitigation of the common 'share inflation' attack vector
     * Constant value chosen to reasonably reduce attempted share inflation by the first depositor, while still
     * incurring reasonably small losses to depositors
     */
    uint256 internal constant BALANCE_OFFSET = 1e3;

    /**
     * @notice The maximum total shares for a given strategy
     * @dev This constant prevents overflow in offchain services for rewards
     */
    uint256 internal constant MAX_TOTAL_SHARES = 1e38 - 1;

    /// @notice EigenLayer's StrategyManager contract
    IStrategyManager public immutable strategyManager;

    /// @notice The underlying token for shares in this Strategy
    IERC20 public underlyingToken;

    /// @notice The total number of extant shares in this Strategy
    uint256 public totalShares;

    /// @notice Simply checks that the `msg.sender` is the `strategyManager`, which is an address stored immutably at construction.
    modifier onlyStrategyManager() {
        require(msg.sender == address(strategyManager), OnlyStrategyManager());
        _;
    }

    /// @notice Since this contract is designed to be initializable, the constructor simply sets `strategyManager`, the only immutable variable.
    constructor(
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        string memory _version
    ) Pausable(_pauserRegistry) SemVerMixin(_version) {
        strategyManager = _strategyManager;
        _disableInitializers();
    }

    function initialize(
        IERC20 _underlyingToken
    ) public virtual initializer {
        _initializeStrategyBase(_underlyingToken);
    }

    /// @notice Sets the `underlyingToken` and `pauserRegistry` for the strategy.
    function _initializeStrategyBase(
        IERC20 _underlyingToken
    ) internal onlyInitializing {
        underlyingToken = _underlyingToken;
        _setPausedStatus(_UNPAUSE_ALL);
        emit StrategyTokenSet(underlyingToken, IERC20Metadata(address(_underlyingToken)).decimals());
    }

    /**
     * @notice Used to deposit tokens into this Strategy
     * @param token is the ERC20 token being deposited
     * @param amount is the amount of token being deposited
     * @dev This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
     * `depositIntoStrategy` function, and individual share balances are recorded in the strategyManager as well.
     * @dev Note that the assumption is made that `amount` of `token` has already been transferred directly to this contract
     * (as performed in the StrategyManager's deposit functions). In particular, setting the `underlyingToken` of this contract
     * to be a fee-on-transfer token will break the assumption that the amount this contract *received* of the token is equal to
     * the amount that was input when the transfer was performed (i.e. the amount transferred 'out' of the depositor's balance).
     * @dev Note that any validation of `token` is done inside `_beforeDeposit`. This can be overridden if needed.
     * @return newShares is the number of new shares issued at the current exchange ratio.
     */
    function deposit(
        IERC20 token,
        uint256 amount
    ) external virtual override onlyWhenNotPaused(PAUSED_DEPOSITS) onlyStrategyManager returns (uint256 newShares) {
        // call hook to allow for any pre-deposit logic
        _beforeDeposit(token, amount);

        // copy `totalShares` value to memory, prior to any change
        uint256 priorTotalShares = totalShares;

        /**
         * @notice calculation of newShares *mirrors* `underlyingToShares(amount)`, but is different since the balance of `underlyingToken`
         * has already been increased due to the `strategyManager` transferring tokens to this strategy prior to calling this function
         */
        // account for virtual shares and balance
        uint256 virtualShareAmount = priorTotalShares + SHARES_OFFSET;
        uint256 virtualTokenBalance = _tokenBalance() + BALANCE_OFFSET;
        // calculate the prior virtual balance to account for the tokens that were already transferred to this contract
        uint256 virtualPriorTokenBalance = virtualTokenBalance - amount;
        newShares = (amount * virtualShareAmount) / virtualPriorTokenBalance;

        // extra check for correctness / against edge case where share rate can be massively inflated as a 'griefing' sort of attack
        require(newShares != 0, NewSharesZero());

        // update total share amount to account for deposit
        totalShares = (priorTotalShares + newShares);
        require(totalShares <= MAX_TOTAL_SHARES, TotalSharesExceedsMax());

        // emit exchange rate
        _emitExchangeRate(virtualTokenBalance, totalShares + SHARES_OFFSET);

        return newShares;
    }

    /**
     * @notice Used to withdraw tokens from this Strategy, to the `recipient`'s address
     * @param recipient is the address to receive the withdrawn funds
     * @param token is the ERC20 token being transferred out
     * @param amountShares is the amount of shares being withdrawn
     * @dev This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
     * other functions, and individual share balances are recorded in the strategyManager as well.
     * @dev Note that any validation of `token` is done inside `_beforeWithdrawal`. This can be overridden if needed.
     */
    function withdraw(
        address recipient,
        IERC20 token,
        uint256 amountShares
    ) external virtual override onlyWhenNotPaused(PAUSED_WITHDRAWALS) onlyStrategyManager {
        // call hook to allow for any pre-withdrawal logic
        _beforeWithdrawal(recipient, token, amountShares);

        // copy `totalShares` value to memory, prior to any change
        uint256 priorTotalShares = totalShares;
        require(amountShares <= priorTotalShares, WithdrawalAmountExceedsTotalDeposits());

        /**
         * @notice calculation of amountToSend *mirrors* `sharesToUnderlying(amountShares)`, but is different since the `totalShares` has already
         * been decremented. Specifically, notice how we use `priorTotalShares` here instead of `totalShares`.
         */
        // account for virtual shares and balance
        uint256 virtualPriorTotalShares = priorTotalShares + SHARES_OFFSET;
        uint256 virtualTokenBalance = _tokenBalance() + BALANCE_OFFSET;
        // calculate ratio based on virtual shares and balance, being careful to multiply before dividing
        uint256 amountToSend = (virtualTokenBalance * amountShares) / virtualPriorTotalShares;

        // Decrease the `totalShares` value to reflect the withdrawal
        totalShares = priorTotalShares - amountShares;

        // emit exchange rate
        _emitExchangeRate(virtualTokenBalance - amountToSend, totalShares + SHARES_OFFSET);

        _afterWithdrawal(recipient, token, amountToSend);
    }

    /**
     * @notice Called in the external `deposit` function, before any logic is executed. Expected to be overridden if strategies want such logic.
     * @param token The token being deposited
     */
    function _beforeDeposit(
        IERC20 token,
        uint256 // amount
    ) internal virtual {
        require(token == underlyingToken, OnlyUnderlyingToken());
    }

    /**
     * @notice Called in the external `withdraw` function, before any logic is executed.  Expected to be overridden if strategies want such logic.
     * @param token The token being withdrawn
     */
    function _beforeWithdrawal(
        address, // recipient
        IERC20 token,
        uint256 // amountShares
    ) internal virtual {
        require(token == underlyingToken, OnlyUnderlyingToken());
    }

    /**
     * @notice Transfers tokens to the recipient after a withdrawal is processed
     * @dev Called in the external `withdraw` function after all logic is executed
     * @param recipient The destination of the tokens
     * @param token The ERC20 being transferred
     * @param amountToSend The amount of `token` to transfer
     */
    function _afterWithdrawal(address recipient, IERC20 token, uint256 amountToSend) internal virtual {
        token.safeTransfer(recipient, amountToSend);
    }

    /// @inheritdoc IStrategy
    function explanation() external pure virtual override returns (string memory) {
        return "Base Strategy implementation to inherit from for more complex implementations";
    }

    /// @inheritdoc IStrategy
    function sharesToUnderlyingView(
        uint256 amountShares
    ) public view virtual override returns (uint256) {
        // account for virtual shares and balance
        uint256 virtualTotalShares = totalShares + SHARES_OFFSET;
        uint256 virtualTokenBalance = _tokenBalance() + BALANCE_OFFSET;
        // calculate ratio based on virtual shares and balance, being careful to multiply before dividing
        return (virtualTokenBalance * amountShares) / virtualTotalShares;
    }

    /// @inheritdoc IStrategy
    function sharesToUnderlying(
        uint256 amountShares
    ) public view virtual override returns (uint256) {
        return sharesToUnderlyingView(amountShares);
    }

    /// @inheritdoc IStrategy
    function underlyingToSharesView(
        uint256 amountUnderlying
    ) public view virtual returns (uint256) {
        // account for virtual shares and balance
        uint256 virtualTotalShares = totalShares + SHARES_OFFSET;
        uint256 virtualTokenBalance = _tokenBalance() + BALANCE_OFFSET;
        // calculate ratio based on virtual shares and balance, being careful to multiply before dividing
        return (amountUnderlying * virtualTotalShares) / virtualTokenBalance;
    }

    /// @inheritdoc IStrategy
    function underlyingToShares(
        uint256 amountUnderlying
    ) external view virtual returns (uint256) {
        return underlyingToSharesView(amountUnderlying);
    }

    /// @inheritdoc IStrategy
    function userUnderlyingView(
        address user
    ) external view virtual returns (uint256) {
        return sharesToUnderlyingView(shares(user));
    }

    /// @inheritdoc IStrategy
    function userUnderlying(
        address user
    ) external virtual returns (uint256) {
        return sharesToUnderlying(shares(user));
    }

    /// @inheritdoc IStrategy
    function shares(
        address user
    ) public view virtual returns (uint256) {
        return strategyManager.stakerDepositShares(user, IStrategy(address(this)));
    }

    /// @notice Internal function used to fetch this contract's current balance of `underlyingToken`.
    // slither-disable-next-line dead-code
    function _tokenBalance() internal view virtual returns (uint256) {
        return underlyingToken.balanceOf(address(this));
    }

    /// @notice Internal function used to emit the exchange rate of the strategy in wad (18 decimals)
    /// @dev Tokens that do not have 18 decimals must have offchain services scale the exchange rate down to proper magnitude
    function _emitExchangeRate(uint256 virtualTokenBalance, uint256 virtualTotalShares) internal {
        // Emit asset over shares ratio.
        emit ExchangeRateEmitted((1e18 * virtualTokenBalance) / virtualTotalShares);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
