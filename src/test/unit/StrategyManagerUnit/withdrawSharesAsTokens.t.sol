// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/unit/StrategyManagerUnit/StrategyManagerUnit.t.sol";

contract StrategyManagerUnitTests_withdrawSharesAsTokens is StrategyManagerUnitTests {
    function testWithdrawSharesAsTokensRevertsDelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert(bytes("StrategyManager.onlyDelegationManager: not the DelegationManager"));
        invalidDelegationManager.removeShares(strategyManager, address(this), dummyStrat, 1);
    }

    function testWithdrawSharesAsTokensRevertsShareAmountTooHigh(
        address staker,
        uint256 depositAmount,
        uint256 sharesAmount
    ) external {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply() && depositAmount < sharesAmount);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert(bytes("StrategyBase.withdraw: amountShares must be less than or equal to totalShares"));
        delegationManagerMock.withdrawSharesAsTokens(strategyManager, staker, strategy, sharesAmount, token);
    }

    function testWithdrawSharesAsTokensSingleStrategyDeposited(
        address staker,
        uint256 depositAmount,
        uint256 sharesAmount
    ) external {
        cheats.assume(staker != address(0));
        cheats.assume(sharesAmount > 0 && sharesAmount < dummyToken.totalSupply() && depositAmount >= sharesAmount);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        uint256 balanceBefore = token.balanceOf(staker);
        delegationManagerMock.withdrawSharesAsTokens(strategyManager, staker, strategy, sharesAmount, token);
        uint256 balanceAfter = token.balanceOf(staker);
        require(balanceAfter == balanceBefore + sharesAmount, "balanceAfter != balanceBefore + sharesAmount");
    }
}
