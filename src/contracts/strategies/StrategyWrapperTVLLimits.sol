// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./StrategyBase.sol";

/**
 * @title A Strategy implementation inheriting from `StrategyWrapper` that limits the total amount of deposits it will accept.
 * @author Layr Labs, Inc.
 */
contract StrategyWrapperTVLLimits is StrategyBase {
    /// The maximum deposits (in underlyingToken) that this strategy will hold
    uint256 public maxDeposits;

    constructor(IStrategyManager _strategyManager) StrategyBase(_strategyManager) {}

    function initialize(IERC20 _underlyingToken, IPauserRegistry _pauserRegistry, uint256 _maxDeposits) public virtual initializer {
        _initializeStrategyBase(_underlyingToken, _pauserRegistry);
        maxDeposits = _maxDeposits;
    }

    /**
     * @notice Sets the maximum deposits (in underlyingToken) that this strategy will hold
     * @param newMaxDeposits The new maximum deposits
     */
    function setMaxDeposits(uint256 newMaxDeposits) external onlyPauser {
        maxDeposits = newMaxDeposits;
    }

    /**
     * Called in the external `deposit` function, before any logic is executed. Makes sure that deposits don't exceed configured maximum.
     * @param token The token being deposited
     * @param amount The amount of `token` being deposited
     */
    function _beforeDeposit(IERC20 token, uint256 amount) internal virtual override {
        require(_tokenBalance() + amount <= maxDeposits, "StrategyWrapperTVLLimits: max deposits exceeded");
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}