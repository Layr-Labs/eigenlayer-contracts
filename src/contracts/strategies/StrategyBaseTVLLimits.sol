// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./StrategyBase.sol";

/**
 * @title A Strategy implementation inheriting from `StrategyBase` that limits the total amount of deposits it will accept.
 * @dev Note that this implementation still converts between any amount of shares or underlying tokens in its view functions;
 * these functions purposefully do not take the TVL limit into account.
 * @author Layr Labs, Inc.
 */
contract StrategyBaseTVLLimits is StrategyBase {
    /// The maximum deposit (in underlyingToken) that this strategy will accept per deposit
    uint256 public maxPerDeposit;

    /// The maximum deposits (in underlyingToken) that this strategy will hold
    uint256 public maxTotalDeposits;

    /// @notice Emitted when `maxPerDeposit` value is updated from `previousValue` to `newValue`
    event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue);

    /// @notice Emitted when `maxTotalDeposits` value is updated from `previousValue` to `newValue`
    event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue);


    constructor(IStrategyManager _strategyManager) StrategyBase(_strategyManager) {}

    function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, IERC20 _underlyingToken, IPauserRegistry _pauserRegistry)
        public virtual initializer
    {
        _setTVLLimits(_maxPerDeposit, _maxTotalDeposits);
        _initializeStrategyBase(_underlyingToken, _pauserRegistry);
    }

    /**
     * @notice Sets the maximum deposits (in underlyingToken) that this strategy will hold and accept per deposit
     * @param newMaxTotalDeposits The new maximum deposits
     * @dev Callable only by the pauser of this contract
     */
    function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) external onlyPauser {
        _setTVLLimits(newMaxPerDeposit, newMaxTotalDeposits);
    }

    /// @notice Simple getter function that returns the current values of `maxPerDeposit` and `maxTotalDeposits`.
    function getTVLLimits() external view returns (uint256, uint256) {
        return (maxPerDeposit, maxTotalDeposits);
    }

    /// @notice Internal setter for TVL limits
    function _setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) internal {
        emit MaxPerDepositUpdated(maxPerDeposit, newMaxPerDeposit);
        emit MaxPerDepositUpdated(maxTotalDeposits, newMaxTotalDeposits);
        maxPerDeposit = newMaxPerDeposit;
        maxTotalDeposits = newMaxTotalDeposits;
    }

    /**
     * @notice Called in the external `deposit` function, before any logic is executed. Makes sure that deposits don't exceed configured maximum.
     * @param token The token being deposited
     * @param amount The amount of `token` being deposited
     */
    function _beforeDeposit(IERC20 token, uint256 amount) internal virtual override {
        require(amount <= maxPerDeposit, "StrategyBaseTVLLimits: max per deposit exceeded");
        require(_tokenBalance() + amount <= maxTotalDeposits, "StrategyBaseTVLLimits: max deposits exceeded");
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}