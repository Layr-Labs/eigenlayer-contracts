// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/strategies/StrategyBase.sol";
import "../../contracts/permissions/PauserRegistry.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/ERC20_SetTransferReverting_Mock.sol";

import "forge-std/Test.sol";

contract StrategyBaseUnitTests is Test {
    Vm cheats = Vm(VM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;
    IStrategyManager public strategyManager;
    IERC20 public underlyingToken;
    StrategyBase public strategyImplementation;
    StrategyBase public strategy;

    address public pauser = address(555);
    address public unpauser = address(999);

    uint initialSupply = 1e36;
    address initialOwner = address(this);

    /**
     * @notice virtual shares used as part of the mitigation of the common 'share inflation' attack vector.
     * Constant value chosen to reasonably reduce attempted share inflation by the first depositor, while still
     * incurring reasonably small losses to depositors
     */
    uint internal constant SHARES_OFFSET = 1e3;
    /**
     * @notice virtual balance used as part of the mitigation of the common 'share inflation' attack vector
     * Constant value chosen to reasonably reduce attempted share inflation by the first depositor, while still
     * incurring reasonably small losses to depositors
     */
    uint internal constant BALANCE_OFFSET = 1e3;

    event ExchangeRateEmitted(uint rate);

    function setUp() public virtual {
        proxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        strategyManager = IStrategyManager(address(new StrategyManagerMock(IDelegationManager(address(0)))));

        underlyingToken = new ERC20PresetFixedSupply("Test Token", "TEST", initialSupply, initialOwner);

        strategyImplementation = new StrategyBase(strategyManager, pauserRegistry, "v9.9.9");

        strategy = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken)
                )
            )
        );
    }

    function testCannotReinitialize() public {
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        strategy.initialize(underlyingToken);
    }

    function testCannotReceiveZeroShares() public {
        uint amountToDeposit = 0;

        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(IStrategyErrors.NewSharesZero.selector);
        strategy.deposit(underlyingToken, amountToDeposit);
        cheats.stopPrank();
    }

    function testDepositWithZeroPriorBalanceAndZeroPriorShares(uint amountToDeposit) public virtual {
        // sanity check / filter
        cheats.assume(amountToDeposit <= underlyingToken.balanceOf(address(this)));
        cheats.assume(amountToDeposit >= 1);

        uint totalSharesBefore = strategy.totalShares();

        underlyingToken.transfer(address(strategy), amountToDeposit);

        cheats.prank(address(strategyManager));
        cheats.expectEmit(true, true, true, true, address(strategy));
        emit ExchangeRateEmitted(1e18);
        uint newShares = strategy.deposit(underlyingToken, amountToDeposit);

        require(newShares == amountToDeposit, "newShares != amountToDeposit");
        uint totalSharesAfter = strategy.totalShares();
        require(totalSharesAfter - totalSharesBefore == newShares, "totalSharesAfter - totalSharesBefore != newShares");
        require(strategy.sharesToUnderlying(1e18) == 1e18);
    }

    function testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares(uint priorTotalShares, uint amountToDeposit) public virtual {
        cheats.assume(priorTotalShares >= 1 && amountToDeposit > 0);

        testDepositWithZeroPriorBalanceAndZeroPriorShares(priorTotalShares);

        // sanity check / filter
        cheats.assume(amountToDeposit <= underlyingToken.balanceOf(address(this)));

        uint totalSharesBefore = strategy.totalShares();

        underlyingToken.transfer(address(strategy), amountToDeposit);

        cheats.prank(address(strategyManager));
        uint newShares = strategy.deposit(underlyingToken, amountToDeposit);

        require(newShares == amountToDeposit, "newShares != amountToDeposit");
        uint totalSharesAfter = strategy.totalShares();
        require(totalSharesAfter - totalSharesBefore == newShares, "totalSharesAfter - totalSharesBefore != newShares");
    }

    function testDepositFailsWhenDepositsPaused() public {
        // pause deposits
        cheats.prank(pauser);
        strategy.pause(1);

        uint amountToDeposit = 1e18;
        underlyingToken.transfer(address(strategy), amountToDeposit);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        cheats.prank(address(strategyManager));
        strategy.deposit(underlyingToken, amountToDeposit);
    }

    function testDepositFailsWhenCallingFromNotStrategyManager(address caller) public {
        cheats.assume(caller != address(strategy.strategyManager()) && caller != address(proxyAdmin));

        uint amountToDeposit = 1e18;
        underlyingToken.transfer(address(strategy), amountToDeposit);

        cheats.expectRevert(IStrategyErrors.OnlyStrategyManager.selector);
        cheats.prank(caller);
        strategy.deposit(underlyingToken, amountToDeposit);
    }

    function testDepositFailsWhenNotUsingUnderlyingToken(address notUnderlyingToken) public {
        cheats.assume(notUnderlyingToken != address(underlyingToken));

        uint amountToDeposit = 1e18;

        cheats.expectRevert(IStrategyErrors.OnlyUnderlyingToken.selector);
        cheats.prank(address(strategyManager));
        strategy.deposit(IERC20(notUnderlyingToken), amountToDeposit);
    }

    function testDepositFailForTooManyShares() public {
        // Deploy token with 1e39 total supply
        underlyingToken = new ERC20PresetFixedSupply("Test Token", "TEST", 1e39, initialOwner);

        strategyImplementation = new StrategyBase(strategyManager, pauserRegistry, "v9.9.9");

        strategy = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken)
                )
            )
        );

        // Transfer underlying token to strategy
        uint amountToDeposit = 1e39;
        underlyingToken.transfer(address(strategy), amountToDeposit);

        // Deposit
        cheats.prank(address(strategyManager));
        cheats.expectRevert(IStrategyErrors.TotalSharesExceedsMax.selector);
        strategy.deposit(underlyingToken, amountToDeposit);
    }

    function testWithdrawWithPriorTotalSharesAndAmountSharesEqual(uint amountToDeposit) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint sharesToWithdraw = strategy.totalShares();
        uint strategyBalanceBefore = underlyingToken.balanceOf(address(strategy));

        uint tokenBalanceBefore = underlyingToken.balanceOf(address(this));
        cheats.prank(address(strategyManager));
        cheats.expectEmit(true, true, true, true, address(strategy));
        emit ExchangeRateEmitted(1e18);
        strategy.withdraw(address(this), underlyingToken, sharesToWithdraw);

        uint tokenBalanceAfter = underlyingToken.balanceOf(address(this));
        uint totalSharesAfter = strategy.totalShares();

        require(totalSharesAfter == 0, "shares did not decrease appropriately");
        require(
            tokenBalanceAfter - tokenBalanceBefore == strategyBalanceBefore,
            "tokenBalanceAfter - tokenBalanceBefore != strategyBalanceBefore"
        );
    }

    function testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual(uint96 amountToDeposit, uint96 sharesToWithdraw) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint totalSharesBefore = strategy.totalShares();
        cheats.assume(sharesToWithdraw <= totalSharesBefore);
        uint strategyBalanceBefore = underlyingToken.balanceOf(address(strategy));
        uint tokenBalanceBefore = underlyingToken.balanceOf(address(this));

        cheats.prank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, sharesToWithdraw);

        uint tokenBalanceAfter = underlyingToken.balanceOf(address(this));
        uint totalSharesAfter = strategy.totalShares();

        require(totalSharesBefore - totalSharesAfter == sharesToWithdraw, "shares did not decrease appropriately");
        require(
            tokenBalanceAfter - tokenBalanceBefore == (strategyBalanceBefore * sharesToWithdraw) / totalSharesBefore,
            "token balance did not increase appropriately"
        );
    }

    function testWithdrawZeroAmount(uint amountToDeposit) public {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint amountToWithdraw = 0;

        uint sharesBefore = strategy.totalShares();
        uint tokenBalanceBefore = underlyingToken.balanceOf(address(this));
        cheats.prank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, amountToWithdraw);

        require(sharesBefore == strategy.totalShares(), "shares changed");
        require(tokenBalanceBefore == underlyingToken.balanceOf(address(this)), "token balance changed");
    }

    function testWithdrawFailsWhenWithdrawalsPaused(uint amountToDeposit) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        // pause withdrawals
        cheats.prank(pauser);
        strategy.pause(2);

        uint amountToWithdraw = 1e18;

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        cheats.prank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, amountToWithdraw);
    }

    function testWithdrawalFailsWhenCallingFromNotStrategyManager(address caller) public {
        cheats.assume(caller != address(strategy.strategyManager()) && caller != address(proxyAdmin));

        uint amountToDeposit = 1e18;
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint amountToWithdraw = 1e18;

        cheats.expectRevert(IStrategyErrors.OnlyStrategyManager.selector);
        cheats.prank(caller);
        strategy.withdraw(address(this), underlyingToken, amountToWithdraw);
    }

    function testWithdrawalFailsWhenNotUsingUnderlyingToken(address notUnderlyingToken) public {
        cheats.assume(notUnderlyingToken != address(underlyingToken));

        uint amountToWithdraw = 1e18;

        cheats.expectRevert(IStrategyErrors.OnlyUnderlyingToken.selector);
        cheats.prank(address(strategyManager));
        strategy.withdraw(address(this), IERC20(notUnderlyingToken), amountToWithdraw);
    }

    function testWithdrawFailsWhenSharesGreaterThanTotalShares(uint amountToDeposit, uint sharesToWithdraw) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint totalSharesBefore = strategy.totalShares();

        // since we are checking strictly greater than in this test
        cheats.assume(sharesToWithdraw > totalSharesBefore);

        cheats.expectRevert(IStrategyErrors.WithdrawalAmountExceedsTotalDeposits.selector);
        cheats.prank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, sharesToWithdraw);
    }

    function testWithdrawalFailsWhenTokenTransferFails() public {
        underlyingToken = new ERC20_SetTransferReverting_Mock(initialSupply, initialOwner);

        strategy = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken)
                )
            )
        );

        uint amountToDeposit = 1e18;
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint amountToWithdraw = 1e18;
        ERC20_SetTransferReverting_Mock(address(underlyingToken)).setTransfersRevert(true);

        cheats.expectRevert();
        cheats.prank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, amountToWithdraw);
    }

    // uint240 input to prevent overflow
    function testIntegrityOfSharesToUnderlyingWithZeroTotalShares(uint240 amountSharesToQuery) public view {
        uint underlyingFromShares = strategy.sharesToUnderlying(amountSharesToQuery);
        require(underlyingFromShares == amountSharesToQuery, "underlyingFromShares != amountSharesToQuery");

        uint underlyingFromSharesView = strategy.sharesToUnderlyingView(amountSharesToQuery);
        require(underlyingFromSharesView == amountSharesToQuery, "underlyingFromSharesView != amountSharesToQuery");
    }

    function testDeposit_ZeroAmount() public {
        cheats.prank(address(strategyManager));
        cheats.expectRevert(IStrategyErrors.NewSharesZero.selector);
        strategy.deposit(underlyingToken, 0);
    }

    // amountSharesToQuery input is uint96 to prevent overflow
    function testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares(
        uint amountToDeposit,
        uint amountToTransfer,
        uint96 amountSharesToQuery
    ) public virtual {
        // sanity check / filter
        cheats.assume(amountToDeposit <= underlyingToken.balanceOf(address(this)));
        cheats.assume(amountToDeposit >= 1);

        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        cheats.assume(amountToTransfer <= underlyingToken.balanceOf(address(this)));
        underlyingToken.transfer(address(strategy), amountToTransfer);
        uint strategyBalance = underlyingToken.balanceOf(address(strategy));
        uint virtualBalance = strategyBalance + BALANCE_OFFSET;

        uint expectedValueOut = (virtualBalance * amountSharesToQuery) / (strategy.totalShares() + SHARES_OFFSET);

        uint underlyingFromShares = strategy.sharesToUnderlying(amountSharesToQuery);
        require(underlyingFromShares == expectedValueOut, "underlyingFromShares != expectedValueOut");

        uint underlyingFromSharesView = strategy.sharesToUnderlyingView(amountSharesToQuery);
        require(underlyingFromSharesView == expectedValueOut, "underlyingFromSharesView != expectedValueOut");
    }

    // amountUnderlyingToQuery input is uint96 to prevent overflow
    function testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares(
        uint amountToDeposit,
        uint amountToTransfer,
        uint96 amountUnderlyingToQuery
    ) public virtual {
        // sanity check / filter
        cheats.assume(amountToDeposit <= underlyingToken.balanceOf(address(this)));
        cheats.assume(amountToDeposit >= 1);

        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        cheats.assume(amountToTransfer <= underlyingToken.balanceOf(address(this)));
        underlyingToken.transfer(address(strategy), amountToTransfer);
        uint strategyBalance = underlyingToken.balanceOf(address(strategy));

        uint expectedValueOut = ((strategy.totalShares() + SHARES_OFFSET) * amountUnderlyingToQuery) / (strategyBalance + BALANCE_OFFSET);

        uint sharesFromUnderlying = strategy.underlyingToShares(amountUnderlyingToQuery);
        require(sharesFromUnderlying == expectedValueOut, "sharesFromUnderlying != expectedValueOut");

        uint sharesFromUnderlyingView = strategy.underlyingToSharesView(amountUnderlyingToQuery);
        require(sharesFromUnderlyingView == expectedValueOut, "sharesFromUnderlyingView != expectedValueOut");
    }

    // verify that small remaining share amounts (nonzero in particular) are allowed
    function testCanWithdrawDownToSmallShares(uint amountToDeposit, uint32 sharesToLeave) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint totalSharesBefore = strategy.totalShares();
        // filter out underflow
        cheats.assume(sharesToLeave <= totalSharesBefore);

        uint sharesToWithdraw = totalSharesBefore - sharesToLeave;

        cheats.prank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, sharesToWithdraw);
    }
}
