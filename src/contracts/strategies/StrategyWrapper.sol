// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IStrategyManager.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title Extremely simple implementation of `IStrategy` interface.
 * @author Layr Labs, Inc.
 * @notice Simple, basic, "do-nothing" Strategy that holds a single underlying token and returns it on withdrawals.
 * Assumes shares are always 1-to-1 with the underlyingToken.
 * @dev Unlike `StrategyBase`, this contract is *not* designed to be inherited from.
 * @dev This contract is expressly *not* intended for use with 'fee-on-transfer'-type tokens.
 * Setting the `underlyingToken` to be a fee-on-transfer token may result in improper accounting.
 */
contract StrategyWrapper is IStrategy {
    using SafeERC20 for IERC20;

    /// @notice EigenLayer's StrategyManager contract
    IStrategyManager public immutable strategyManager;

    /// @notice The underyling token for shares in this Strategy
    IERC20 public immutable underlyingToken;

    /// @notice The total number of extant shares in thie Strategy
    uint256 public totalShares;

    modifier onlyStrategyManager() {
        require(msg.sender == address(strategyManager), "StrategyWrapper.onlyStrategyManager");
        _;
    }

    constructor(IStrategyManager _strategyManager, IERC20 _underlyingToken) {
        strategyManager = _strategyManager;
        underlyingToken = _underlyingToken;
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
     * @return newShares is the number of new shares issued at the current exchange ratio.
     */
    function deposit(IERC20 token, uint256 amount) external override onlyStrategyManager returns (uint256) {
        // call hook to allow for any pre-deposit logic
        _beforeDeposit(token, amount);

        require(token == underlyingToken, "StrategyWrapper.deposit: Can only deposit underlyingToken");
        totalShares += amount;
        return amount;
    }

    /**
     * @notice Used to withdraw tokens from this Strategy, to the `depositor`'s address
     * @param token is the ERC20 token being transferred out
     * @param amountShares is the amount of shares being withdrawn
     * @dev This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
     * other functions, and individual share balances are recorded in the strategyManager as well.
     */
    function withdraw(address depositor, IERC20 token, uint256 amountShares)
        external
        override
        onlyStrategyManager
    {
        // call hook to allow for any pre-withdrawal logic
        _beforeWithdrawal(depositor, token, amountShares);

        require(token == underlyingToken, "StrategyWrapper.withdraw: Can only withdraw the strategy token");
        require(
            amountShares <= totalShares,
            "StrategyWrapper.withdraw: amountShares must be less than or equal to totalShares"
        );
        // Decrease `totalShares` to reflect withdrawal. Unchecked arithmetic since we just checked this above.
        unchecked {
            totalShares -= amountShares;
        }
        underlyingToken.safeTransfer(depositor, amountShares);
    }

    /**
     * @notice Called in the external `deposit` function, before any logic is executed. Expected to be overridden if strategies want such logic.
     * @param token The token being deposited
     * @param amount The amount of `token` being deposited
     */
    function _beforeDeposit(IERC20 token, uint256 amount)  internal virtual {}

    /**
     * @notice Called in the external `withdraw` function, before any logic is executed.  Expected to be overridden if strategies want such logic.
     * @param depositor The address of the depositor
     * @param token The token being withdrawn
     * @param amountShares The amount of shares being withdrawn
     */
    function _beforeWithdrawal(address depositor, IERC20 token, uint256 amountShares) internal virtual {}

    /**
     * @notice Currently returns a brief string explaining the strategy's goal & purpose, but for more complex
     * strategies, may be a link to metadata that explains in more detail.
     */
    function explanation() external pure override returns (string memory) {
        return "Wrapper Strategy to simply store tokens. Assumes fixed 1-to-1 share-underlying exchange.";
    }

    /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
     * @notice In contrast to `sharesToUnderlying`, this function guarantees no state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function sharesToUnderlyingView(uint256 amountShares) public pure override returns (uint256) {
        return amountShares;
    }

    /**
     * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
     * @notice In contrast to `sharesToUnderlyingView`, this function **may** make state modifications
     * @param amountShares is the amount of shares to calculate its conversion into the underlying token
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function sharesToUnderlying(uint256 amountShares) public pure override returns (uint256) {
        return amountShares;
    }

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
     * @notice In contrast to `underlyingToShares`, this function guarantees no state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into strategy shares
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function underlyingToSharesView(uint256 amountUnderlying) external pure returns (uint256) {
        return amountUnderlying;
    }

    /**
     * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
     * @notice In contrast to `underlyingToSharesView`, this function **may** make state modifications
     * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into strategy shares
     * @dev Implementation for these functions in particular may vary signifcantly for different strategies
     */
    function underlyingToShares(uint256 amountUnderlying) external pure returns (uint256) {
        return amountUnderlying;
    }

    /**
     * @notice convenience function for fetching the current underlying value of all of the `user`'s shares in
     * this strategy. In contrast to `userUnderlying`, this function guarantees no state modifications
     */
    function userUnderlyingView(address user) external view returns (uint256) {
        return sharesToUnderlyingView(shares(user));
    }

    /**
     * @notice convenience function for fetching the current underlying value of all of the `user`'s shares in
     * this strategy. In contrast to `userUnderlyingView`, this function **may** make state modifications
     */
    function userUnderlying(address user) external view returns (uint256) {
        return sharesToUnderlying(shares(user));
    }

    /**
     * @notice convenience function for fetching the current total shares of `user` in this strategy, by
     * querying the `strategyManager` contract
     */
    function shares(address user) public view returns (uint256) {
        return IStrategyManager(strategyManager).stakerStrategyShares(user, IStrategy(address(this)));
    }
}
