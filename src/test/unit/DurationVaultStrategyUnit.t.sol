// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./StrategyBaseTVLLimitsUnit.sol";
import "../../contracts/strategies/DurationVaultStrategy.sol";
import "../../contracts/interfaces/IDurationVaultStrategy.sol";

contract DurationVaultStrategyUnitTests is StrategyBaseTVLLimitsUnitTests {
    DurationVaultStrategy public durationVaultImplementation;
    DurationVaultStrategy public durationVault;

    uint64 internal defaultDuration = 30 days;
    uint64 internal defaultDepositWindowLength = 7 days;

    function setUp() public virtual override {
        StrategyBaseUnitTests.setUp();

        durationVaultImplementation = new DurationVaultStrategy(strategyManager, pauserRegistry, "9.9.9");

        IDurationVaultStrategy.VaultConfig memory config = IDurationVaultStrategy.VaultConfig({
            underlyingToken: underlyingToken,
            vaultAdmin: address(this),
            depositWindowStart: 0,
            depositWindowEnd: uint64(block.timestamp + defaultDepositWindowLength),
            duration: defaultDuration,
            maxPerDeposit: maxPerDeposit,
            stakeCap: maxTotalDeposits,
            metadataURI: "ipfs://duration-vault"
        });

        durationVault = DurationVaultStrategy(
            address(
                new TransparentUpgradeableProxy(
                    address(durationVaultImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(DurationVaultStrategy.initialize.selector, config)
                )
            )
        );

        strategy = StrategyBase(address(durationVault));
        strategyWithTVLLimits = StrategyBaseTVLLimits(address(durationVault));
    }

    function testDepositWindowNotStarted() public {
        // reconfigure deposit window to start in future
        durationVault.updateDepositWindow(uint64(block.timestamp + 1 hours), uint64(block.timestamp + 2 hours));

        uint depositAmount = 1e18;
        underlyingToken.transfer(address(durationVault), depositAmount);

        cheats.prank(address(strategyManager));
        cheats.expectRevert(IDurationVaultStrategy.DepositWindowNotStarted.selector);
        durationVault.deposit(underlyingToken, depositAmount);
    }

    function testDepositWindowClosedAfterEnd() public {
        cheats.warp(block.timestamp + defaultDepositWindowLength + 1);

        uint depositAmount = 1e18;
        underlyingToken.transfer(address(durationVault), depositAmount);

        cheats.prank(address(strategyManager));
        cheats.expectRevert(IDurationVaultStrategy.DepositWindowClosed.selector);
        durationVault.deposit(underlyingToken, depositAmount);
    }

    function testDepositsBlockedAfterManualLock() public {
        durationVault.lock();

        uint depositAmount = 1e18;

        underlyingToken.transfer(address(durationVault), depositAmount);
        cheats.prank(address(strategyManager));
        cheats.expectRevert(IDurationVaultStrategy.DepositWindowClosed.selector);
        durationVault.deposit(underlyingToken, depositAmount);
    }

    function testWithdrawalsBlockedUntilMaturity() public {
        // prepare deposit
        uint depositAmount = 10 ether;
        underlyingToken.transfer(address(durationVault), depositAmount);
        cheats.prank(address(strategyManager));
        durationVault.deposit(underlyingToken, depositAmount);

        durationVault.lock();

        assertTrue(durationVault.isLocked(), "vault should be locked");
        assertFalse(durationVault.withdrawalsOpen(), "withdrawals should be closed before maturity");

        uint shares = durationVault.totalShares();

        // Attempt withdrawal before maturity
        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(IDurationVaultStrategy.WithdrawalsLocked.selector);
        durationVault.withdraw(address(this), underlyingToken, shares);
        cheats.stopPrank();

        cheats.warp(block.timestamp + defaultDuration + 1);

        cheats.startPrank(address(strategyManager));
        durationVault.withdraw(address(this), underlyingToken, durationVault.totalShares());
        cheats.stopPrank();
    }
}
