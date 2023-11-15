// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./utils/User.sol";

contract ExampleIntegrationTest is EigenLayerDeployer {

    function testDepositAndWithdrawWethAsShares() public {
        // create a new user
        User staker = new User();

        uint256 amountToDeposit = 10e18;
        uint256 amountToWithdraw = 5e18;

        // give staker funds to deposit
        weth.transfer(staker.user(), amountToDeposit);

        staker.depositToStrategy(amountToDeposit, weth, wethStrat);

        uint256 userSharesBeforeWithdrawal = strategyManager.stakerStrategyShares(staker.user(), wethStrat);

        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrat;
        uint256[] memory shares = new uint256[](1);
        shares[0] = amountToWithdraw;
        address withdrawer = staker.user();

        staker.queueWithdrawal(strategies, shares, withdrawer);

        staker.completeOldestQueuedWithdrawalAsShares();

        uint256 userSharesAfterWithdrawal = strategyManager.stakerStrategyShares(staker.user(), wethStrat);

        assertEq(userSharesBeforeWithdrawal, userSharesAfterWithdrawal, "user shares changed inappropriately");
    }
}
