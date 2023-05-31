// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./StrategyBase.sol";
import "forge-std/Test.sol";

/**
 * @title A Strategy implementation inheriting from `StrategyBase` that limits the total amount of deposits it will accept.
 * @dev Note that this implementation still converts between any amount of shares or underlying tokens in its view functions;
 * these functions purposefully do not take the TVL limit into account.
 * @author Layr Labs, Inc.
 */
contract StrategyBaseTVLLimits is StrategyBase, Test {
    /// The maximum deposit (in underlyingToken) that this strategy will accept per deposit
    uint256 public maxDepositPerAddress;

    /// The maximum deposits (in underlyingToken) that this strategy will hold
    uint256 public maxTotalDeposits;

    /// @notice Emitted when `maxDepositPerAddress` value is updated from `previousValue` to `newValue`
    event MaxDepositPerAddressUpdated(uint256 previousValue, uint256 newValue);

    /// @notice Emitted when `maxTotalDeposits` value is updated from `previousValue` to `newValue`
    event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue);

    mapping(address => uint256) public depositorToDepositAmount;


    constructor(IStrategyManager _strategyManager) StrategyBase(_strategyManager) {}

    function initialize(uint256 _maxDepositPerAddress, uint256 _maxTotalDeposits, IERC20 _underlyingToken, IPauserRegistry _pauserRegistry)
        public virtual initializer
    {
        _setTVLLimits(_maxDepositPerAddress, _maxTotalDeposits);
        _initializeStrategyBase(_underlyingToken, _pauserRegistry);
    }

    /**
     * @notice Sets the maximum deposits (in underlyingToken) that this strategy will hold and accept per deposit
     * @param newMaxTotalDeposits The new maximum deposits
     * @dev Callable only by the pauser of this contract
     * @dev We note that there is a potential race condition between a call to this function that lowers either or both of these limits and call(s)
     * to `deposit`, that may result in some calls to `deposit` reverting.
     */
    function setTVLLimits(uint256 newMaxDepositPerAddress, uint256 newMaxTotalDeposits) external onlyPauser {
        _setTVLLimits(newMaxDepositPerAddress, newMaxTotalDeposits);
    }

    /// @notice Simple getter function that returns the current values of `maxPerDeposit` and `maxTotalDeposits`.
    function getTVLLimits() external view returns (uint256, uint256) {
        return (maxDepositPerAddress, maxTotalDeposits);
    }

    /// @notice Internal setter for TVL limits
    function _setTVLLimits(uint256 newMaxDepositPerAddress, uint256 newMaxTotalDeposits) internal {
        emit MaxDepositPerAddressUpdated(maxDepositPerAddress, newMaxDepositPerAddress);
        emit MaxTotalDepositsUpdated(maxTotalDeposits, newMaxTotalDeposits);
        require(newMaxDepositPerAddress <= newMaxTotalDeposits, "StrategyBaseTVLLimits._setTVLLimits: maxPerDeposit exceeds maxTotalDeposits");
        maxDepositPerAddress = newMaxDepositPerAddress;
        maxTotalDeposits = newMaxTotalDeposits;
    }

    /**
     * @notice Called in the external `deposit` function, before any logic is executed. Makes sure that deposits don't exceed configured maximum.
     * @dev Unused token param is the token being deposited. This is already checked in the `deposit` function.
     * @dev Note that the `maxTotalDeposits` is purely checked against the current `_tokenBalance()`, since by this point in the deposit flow, the
     * tokens should have already been transferred to this Strategy by the StrategyManager
     * @dev We note as well that this makes it possible for various race conditions to occur:
     * a) multiple simultaneous calls to `deposit` may result in some of these calls reverting due to `maxTotalDeposits` being reached.
     * b) transferring funds directly to this Strategy (although not generally in someone's economic self interest) in order to reach `maxTotalDeposits`
     * is a route by which someone can cause calls to `deposit` to revert.
     * c) increases in the token balance of this contract through other effects – including token rebasing – may cause similar issues to (a) and (b).
     * @param amount The amount of `token` being deposited
     */
    function _beforeDeposit(address depositor, IERC20 /*token*/, uint256 amount) internal virtual override {
        require(depositorToDepositAmount[depositor] + amount <= maxDepositPerAddress, "StrategyBaseTVLLimits: max deposit per address exceeded");
        require(_tokenBalance() <= maxTotalDeposits, "StrategyBaseTVLLimits: max deposits exceeded");
        depositorToDepositAmount[depositor] += amount;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}