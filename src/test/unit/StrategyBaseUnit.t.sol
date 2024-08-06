// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/strategies/StrategyBase.sol";
import "../../contracts/permissions/PauserRegistry.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/ERC20_SetTransferReverting_Mock.sol";

import "forge-std/Test.sol";

contract StrategyBaseUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;
    IStrategyManager public strategyManager;
    IERC20 public underlyingToken;
    StrategyBase public strategyImplementation;
    StrategyBase public strategy;

    address public pauser = address(555);
    address public unpauser = address(999);

    uint256 initialSupply = 1e36;
    address initialOwner = address(this);

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

    event ExchangeRateEmitted(uint256 rate);

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);
        
        strategyManager = new StrategyManagerMock();

        underlyingToken = new ERC20PresetFixedSupply("Test Token", "TEST", initialSupply, initialOwner);

        strategyImplementation = new StrategyBase(strategyManager);

        strategy = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, pauserRegistry)
                )
            )
        );
    }

    function testCannotReinitialize() public {
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        strategy.initialize(underlyingToken, pauserRegistry);
    }

    function testCannotReceiveZeroShares() public {
        uint256 amountToDeposit = 0;

        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBase.deposit: newShares cannot be zero"));
        strategy.deposit(underlyingToken, amountToDeposit);
        cheats.stopPrank();
    }

    function testDepositWithZeroPriorBalanceAndZeroPriorShares(uint256 amountToDeposit) public virtual {
        // sanity check / filter
        cheats.assume(amountToDeposit <= underlyingToken.balanceOf(address(this)));
        cheats.assume(amountToDeposit >= 1);

        uint256 totalSharesBefore = strategy.totalShares();

        underlyingToken.transfer(address(strategy), amountToDeposit);

        cheats.startPrank(address(strategyManager));
        cheats.expectEmit(true, true, true, true, address(strategy));
        emit ExchangeRateEmitted(1e18);
        uint256 newShares = strategy.deposit(underlyingToken, amountToDeposit);
        cheats.stopPrank();

        require(newShares == amountToDeposit, "newShares != amountToDeposit");
        uint256 totalSharesAfter = strategy.totalShares();
        require(totalSharesAfter - totalSharesBefore == newShares, "totalSharesAfter - totalSharesBefore != newShares");
        require(strategy.sharesToUnderlying(1e18) == 1e18);
    }

    function testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares(uint256 priorTotalShares, uint256 amountToDeposit) public virtual {
        cheats.assume(priorTotalShares >= 1 && amountToDeposit > 0);

        testDepositWithZeroPriorBalanceAndZeroPriorShares(priorTotalShares);

        // sanity check / filter
        cheats.assume(amountToDeposit <= underlyingToken.balanceOf(address(this)));

        uint256 totalSharesBefore = strategy.totalShares();

        underlyingToken.transfer(address(strategy), amountToDeposit);

        cheats.startPrank(address(strategyManager));

        uint256 newShares = strategy.deposit(underlyingToken, amountToDeposit);
        cheats.stopPrank();

        require(newShares == amountToDeposit, "newShares != amountToDeposit");
        uint256 totalSharesAfter = strategy.totalShares();
        require(totalSharesAfter - totalSharesBefore == newShares, "totalSharesAfter - totalSharesBefore != newShares");
    }

    function testDepositFailsWhenDepositsPaused() public {
        // pause deposits
        cheats.startPrank(pauser);
        strategy.pause(1);
        cheats.stopPrank();

        uint256 amountToDeposit = 1e18;
        underlyingToken.transfer(address(strategy), amountToDeposit);

        cheats.expectRevert(bytes("Pausable: index is paused"));
        cheats.startPrank(address(strategyManager));
        strategy.deposit(underlyingToken, amountToDeposit);
        cheats.stopPrank();
    }


    function testDepositFailsWhenCallingFromNotStrategyManager(address caller) public {
        cheats.assume(caller != address(strategy.strategyManager()) && caller != address(proxyAdmin));

        uint256 amountToDeposit = 1e18;
        underlyingToken.transfer(address(strategy), amountToDeposit);

        cheats.expectRevert(bytes("StrategyBase.onlyStrategyManager"));
        cheats.startPrank(caller);
        strategy.deposit(underlyingToken, amountToDeposit);
        cheats.stopPrank();
    }

    function testDepositFailsWhenNotUsingUnderlyingToken(address notUnderlyingToken) public {
        cheats.assume(notUnderlyingToken != address(underlyingToken));

        uint256 amountToDeposit = 1e18;

        cheats.expectRevert(bytes("StrategyBase.deposit: Can only deposit underlyingToken"));
        cheats.startPrank(address(strategyManager));
        strategy.deposit(IERC20(notUnderlyingToken), amountToDeposit);
        cheats.stopPrank();
    }

    function testDepositFailForTooManyShares() public {
        // Deploy token with 1e39 total supply
        underlyingToken = new ERC20PresetFixedSupply("Test Token", "TEST", 1e39, initialOwner);

        strategyImplementation = new StrategyBase(strategyManager);

        strategy = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, pauserRegistry)
                )
            )
        );

        // Transfer underlying token to strategy
        uint256 amountToDeposit = 1e39;
        underlyingToken.transfer(address(strategy), amountToDeposit);


        // Deposit
        cheats.prank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBase.deposit: totalShares exceeds `MAX_TOTAL_SHARES`"));
        strategy.deposit(underlyingToken, amountToDeposit);
    }

    function testWithdrawWithPriorTotalSharesAndAmountSharesEqual(uint256 amountToDeposit) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint256 sharesToWithdraw = strategy.totalShares();
        uint256 strategyBalanceBefore = underlyingToken.balanceOf(address(strategy));

        uint256 tokenBalanceBefore = underlyingToken.balanceOf(address(this));
        cheats.startPrank(address(strategyManager));
        cheats.expectEmit(true, true, true, true, address(strategy));
        emit ExchangeRateEmitted(1e18);
        strategy.withdraw(address(this), underlyingToken, sharesToWithdraw);

        cheats.stopPrank();

        uint256 tokenBalanceAfter = underlyingToken.balanceOf(address(this));
        uint256 totalSharesAfter = strategy.totalShares();

        require(totalSharesAfter == 0, "shares did not decrease appropriately");
        require(tokenBalanceAfter - tokenBalanceBefore == strategyBalanceBefore, "tokenBalanceAfter - tokenBalanceBefore != strategyBalanceBefore");
    }

    function testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual(uint96 amountToDeposit, uint96 sharesToWithdraw) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint256 totalSharesBefore = strategy.totalShares();
        cheats.assume(sharesToWithdraw <= totalSharesBefore);
        uint256 strategyBalanceBefore = underlyingToken.balanceOf(address(strategy));
        uint256 tokenBalanceBefore = underlyingToken.balanceOf(address(this));


        cheats.startPrank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, sharesToWithdraw);
        cheats.stopPrank();

        uint256 tokenBalanceAfter = underlyingToken.balanceOf(address(this));
        uint256 totalSharesAfter = strategy.totalShares();

        require(totalSharesBefore - totalSharesAfter == sharesToWithdraw, "shares did not decrease appropriately");
        require(tokenBalanceAfter - tokenBalanceBefore == (strategyBalanceBefore * sharesToWithdraw) / totalSharesBefore,
            "token balance did not increase appropriately");
    }

    function testWithdrawZeroAmount(uint256 amountToDeposit) public {
         cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint256 amountToWithdraw = 0;

        uint256 sharesBefore = strategy.totalShares();
        uint256 tokenBalanceBefore = underlyingToken.balanceOf(address(this));
        cheats.startPrank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, amountToWithdraw);
        cheats.stopPrank();

        require(sharesBefore == strategy.totalShares(), "shares changed");
        require(tokenBalanceBefore == underlyingToken.balanceOf(address(this)), "token balance changed");
    }

    function testWithdrawFailsWhenWithdrawalsPaused(uint256 amountToDeposit) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        // pause withdrawals
        cheats.startPrank(pauser);
        strategy.pause(2);
        cheats.stopPrank();

        uint256 amountToWithdraw = 1e18;

        cheats.expectRevert(bytes("Pausable: index is paused"));
        cheats.startPrank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, amountToWithdraw);
        cheats.stopPrank();
    }

    function testWithdrawalFailsWhenCallingFromNotStrategyManager(address caller) public {
        cheats.assume(caller != address(strategy.strategyManager()) && caller != address(proxyAdmin));

        uint256 amountToDeposit = 1e18;
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint256 amountToWithdraw = 1e18;

        cheats.expectRevert(bytes("StrategyBase.onlyStrategyManager"));
        cheats.startPrank(caller);
        strategy.withdraw(address(this), underlyingToken, amountToWithdraw);
        cheats.stopPrank();
    }

    function testWithdrawalFailsWhenNotUsingUnderlyingToken(address notUnderlyingToken) public {
        cheats.assume(notUnderlyingToken != address(underlyingToken));

        uint256 amountToWithdraw = 1e18;

        cheats.expectRevert(bytes("StrategyBase.withdraw: Can only withdraw the strategy token"));
        cheats.startPrank(address(strategyManager));
        strategy.withdraw(address(this), IERC20(notUnderlyingToken), amountToWithdraw);
        cheats.stopPrank();
    }

    function testWithdrawFailsWhenSharesGreaterThanTotalShares(uint256 amountToDeposit, uint256 sharesToWithdraw) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint256 totalSharesBefore = strategy.totalShares();

        // since we are checking strictly greater than in this test
        cheats.assume(sharesToWithdraw > totalSharesBefore);

        cheats.expectRevert(bytes("StrategyBase.withdraw: amountShares must be less than or equal to totalShares"));
        cheats.startPrank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, sharesToWithdraw);
        cheats.stopPrank();
    }

    function testWithdrawalFailsWhenTokenTransferFails() public {
        underlyingToken = new ERC20_SetTransferReverting_Mock(initialSupply, initialOwner);

        strategy = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, pauserRegistry)
                )
            )
        );

        uint256 amountToDeposit = 1e18;
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint256 amountToWithdraw = 1e18;
        ERC20_SetTransferReverting_Mock(address(underlyingToken)).setTransfersRevert(true);

        cheats.expectRevert();
        cheats.startPrank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, amountToWithdraw);
        cheats.stopPrank();
    }

    // uint240 input to prevent overflow
    function testIntegrityOfSharesToUnderlyingWithZeroTotalShares(uint240 amountSharesToQuery) public {
        uint256 underlyingFromShares = strategy.sharesToUnderlying(amountSharesToQuery);
        require(underlyingFromShares == amountSharesToQuery, "underlyingFromShares != amountSharesToQuery");

        uint256 underlyingFromSharesView = strategy.sharesToUnderlyingView(amountSharesToQuery);
        require(underlyingFromSharesView == amountSharesToQuery, "underlyingFromSharesView != amountSharesToQuery");
    }

    function testDeposit_ZeroAmount() public {
        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBase.deposit: newShares cannot be zero"));
        strategy.deposit(underlyingToken, 0);
        cheats.stopPrank();
    }


    // amountSharesToQuery input is uint96 to prevent overflow
    function testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares(uint256 amountToDeposit, uint256 amountToTransfer, uint96 amountSharesToQuery
    ) public virtual {
        // sanity check / filter
        cheats.assume(amountToDeposit <= underlyingToken.balanceOf(address(this)));
        cheats.assume(amountToDeposit >= 1);

        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        cheats.assume(amountToTransfer <= underlyingToken.balanceOf(address(this)));
        underlyingToken.transfer(address(strategy), amountToTransfer);
        uint256 strategyBalance = underlyingToken.balanceOf(address(strategy));
        uint256 virtualBalance = strategyBalance + BALANCE_OFFSET;

        uint256 expectedValueOut = (virtualBalance * amountSharesToQuery) / (strategy.totalShares() + SHARES_OFFSET);

        uint256 underlyingFromShares = strategy.sharesToUnderlying(amountSharesToQuery);
        require(underlyingFromShares == expectedValueOut, "underlyingFromShares != expectedValueOut");

        uint256 underlyingFromSharesView = strategy.sharesToUnderlyingView(amountSharesToQuery);
        require(underlyingFromSharesView == expectedValueOut, "underlyingFromSharesView != expectedValueOut");
    }

    // amountUnderlyingToQuery input is uint96 to prevent overflow
    function testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares(uint256 amountToDeposit, uint256 amountToTransfer, uint96 amountUnderlyingToQuery
    ) public virtual {
        // sanity check / filter
        cheats.assume(amountToDeposit <= underlyingToken.balanceOf(address(this)));
        cheats.assume(amountToDeposit >= 1);

        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        cheats.assume(amountToTransfer <= underlyingToken.balanceOf(address(this)));
        underlyingToken.transfer(address(strategy), amountToTransfer);
        uint256 strategyBalance = underlyingToken.balanceOf(address(strategy));

        uint256 expectedValueOut = ((strategy.totalShares() + SHARES_OFFSET) * amountUnderlyingToQuery) / (strategyBalance + BALANCE_OFFSET);

        uint256 sharesFromUnderlying = strategy.underlyingToShares(amountUnderlyingToQuery);
        require(sharesFromUnderlying == expectedValueOut, "sharesFromUnderlying != expectedValueOut");

        uint256 sharesFromUnderlyingView = strategy.underlyingToSharesView(amountUnderlyingToQuery);
        require(sharesFromUnderlyingView == expectedValueOut, "sharesFromUnderlyingView != expectedValueOut");
    }

    // verify that small remaining share amounts (nonzero in particular) are allowed
    function testCanWithdrawDownToSmallShares(uint256 amountToDeposit, uint32 sharesToLeave) public virtual {
        cheats.assume(amountToDeposit >= 1);
        testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);

        uint256 totalSharesBefore = strategy.totalShares();
        // filter out underflow
        cheats.assume(sharesToLeave <= totalSharesBefore);

        uint256 sharesToWithdraw = totalSharesBefore - sharesToLeave;

        cheats.startPrank(address(strategyManager));
        strategy.withdraw(address(this), underlyingToken, sharesToWithdraw);
        cheats.stopPrank();
    }
}
