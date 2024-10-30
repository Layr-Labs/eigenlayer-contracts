// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.27;

import "./EigenLayerTestHelper.t.sol";

contract PausableTests is EigenLayerTestHelper {

    /// @notice Emitted when the `pauserRegistry` is set to `newPauserRegistry`.
    event PauserRegistrySet(IPauserRegistry pauserRegistry, IPauserRegistry newPauserRegistry);

    ///@dev test that pausing a contract works
    function testPausingWithdrawalsFromStrategyManager(uint256 amountToDeposit, uint256 amountToWithdraw) public {
        cheats.assume(amountToDeposit <= 1e38 - 1);
        cheats.assume(amountToDeposit <= weth.balanceOf(address(this)));
        // if first deposit amount to base strategy is too small, it will revert. ignore that case here.
        cheats.assume(amountToDeposit >= 1);
        cheats.assume(amountToWithdraw <= amountToDeposit);

        address sender = getOperatorAddress(0);
        _testDepositToStrategy(sender, amountToDeposit, weth, wethStrat);

        cheats.startPrank(pauser);
        strategyManager.pause(type(uint256).max);
        cheats.stopPrank();

        // uint256 strategyIndex = 0;

        cheats.prank(sender);

        // TODO: write this to work with completing a queued withdrawal
        // cheats.expectRevert(bytes("Pausable: paused"));
        // strategyManager.withdrawFromStrategy(strategyIndex, wethStrat, weth, amountToWithdraw);
        // cheats.stopPrank();
    }

    function testUnauthorizedPauserStrategyManager(address unauthorizedPauser)
        public
        fuzzedAddress(unauthorizedPauser)
    {
        cheats.assume(!eigenLayerPauserReg.isPauser(unauthorizedPauser));
        cheats.startPrank(unauthorizedPauser);
        cheats.expectRevert(IPausable.OnlyPauser.selector);
        strategyManager.pause(type(uint256).max);
        cheats.stopPrank();
    }

    function testSetPauser(address newPauser) public fuzzedAddress(newPauser) {
        cheats.startPrank(unpauser);
        eigenLayerPauserReg.setIsPauser(newPauser, true);
        cheats.stopPrank();
    }

    function testSetUnpauser(address newUnpauser) public fuzzedAddress(newUnpauser) {
        cheats.startPrank(unpauser);
        eigenLayerPauserReg.setUnpauser(newUnpauser);
        cheats.stopPrank();
    }

    function testSetPauserUnauthorized(address fakePauser, address newPauser)
        public
        fuzzedAddress(newPauser)
        fuzzedAddress(fakePauser)
    {
        cheats.assume(fakePauser != eigenLayerPauserReg.unpauser());
        cheats.startPrank(fakePauser);
        cheats.expectRevert(IPausable.OnlyUnpauser.selector);
        eigenLayerPauserReg.setIsPauser(newPauser, true);
        cheats.stopPrank();
    }
}
