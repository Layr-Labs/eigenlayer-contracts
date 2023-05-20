// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./StrategyBaseUnit.t.sol";

import "../../contracts/strategies/StrategyBaseTVLLimits.sol";

contract StrategyBaseTVLLimitsUnitTests is StrategyBaseUnitTests {
    StrategyBaseTVLLimits public strategyBaseTVLLimitsImplementation;
    StrategyBaseTVLLimits public strategyWithTVLLimits;

    // defaults for tests, used in setup
    uint256 maxTotalDeposits = 3200e18;
    uint256 maxPerDeposit = 32e18;

    function setUp() virtual public override {
        // copy setup for StrategyBaseUnitTests
        StrategyBaseUnitTests.setUp();

        // depoloy the TVL-limited strategy
        strategyBaseTVLLimitsImplementation = new StrategyBaseTVLLimits(strategyManager);
        strategyWithTVLLimits = StrategyBaseTVLLimits(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyBaseTVLLimitsImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBaseTVLLimits.initialize.selector, maxPerDeposit, maxTotalDeposits, underlyingToken, pauserRegistry)
                )
            )
        );

        // verify that limits were set up correctly
        require(strategyWithTVLLimits.maxTotalDeposits() == maxTotalDeposits, "bad test setup");
        require(strategyWithTVLLimits.maxPerDeposit() == maxPerDeposit, "bad test setup");

        // replace the strategy from the non-TVL-limited tests with the TVL-limited strategy
        // this should result in all the StrategyBaseUnitTests also running against the `StrategyBaseTVLLimits` contract
        strategy = strategyWithTVLLimits;
    }

    function testSetTVLLimits(uint256 maxDepositsFuzzedInput, uint256 maxPerDepositFuzzedInput) public {
        cheats.startPrank(pauser);
        strategyWithTVLLimits.setTVLLimits(maxPerDepositFuzzedInput, maxDepositsFuzzedInput);
        cheats.stopPrank();
        (uint256 _maxPerDeposit, uint256 _maxDeposits) = strategyWithTVLLimits.getTVLLimits();

        assertEq(_maxPerDeposit, maxPerDepositFuzzedInput);
        assertEq(_maxDeposits, maxDepositsFuzzedInput);
    }

    function testSetTVLLimitsFailsWhenNotCalledByPauser(uint256 maxDepositsFuzzedInput, uint256 maxPerDepositFuzzedInput, address notPauser) public {
        cheats.assume(notPauser != address(proxyAdmin));
        cheats.assume(notPauser != pauser);
        cheats.startPrank(notPauser);
        cheats.expectRevert(bytes("msg.sender is not permissioned as pauser"));
        strategyWithTVLLimits.setTVLLimits(maxPerDepositFuzzedInput, maxDepositsFuzzedInput);
        cheats.stopPrank();
    }

    function testDepositMoreThanMaxPerDeposit(uint256 maxDepositsFuzzedInput, uint256 maxPerDepositFuzzedInput, uint256 amount) public {
        cheats.assume(amount > maxPerDepositFuzzedInput);
        cheats.startPrank(pauser);
        strategyWithTVLLimits.setTVLLimits(maxPerDepositFuzzedInput, maxDepositsFuzzedInput);
        cheats.stopPrank();

        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBaseTVLLimits: max per deposit exceeded"));
        strategyWithTVLLimits.deposit(underlyingToken, amount);
        cheats.stopPrank();
    }

    function testDepositMorethanMaxDeposits() public {
        maxTotalDeposits = 1e12;
        maxPerDeposit = 3e11;
        uint256 numDeposits = maxTotalDeposits / maxPerDeposit;

        underlyingToken.transfer(address(strategyManager), maxTotalDeposits);
        cheats.startPrank(pauser);
        strategyWithTVLLimits.setTVLLimits(maxPerDeposit, maxTotalDeposits);
        cheats.stopPrank();

        cheats.startPrank(address(strategyManager));
        for (uint256 i = 0; i < numDeposits; i++) {
            underlyingToken.transfer(address(strategyWithTVLLimits), maxPerDeposit);
            strategyWithTVLLimits.deposit(underlyingToken, maxPerDeposit);
        }
        cheats.stopPrank();

        underlyingToken.transfer(address(strategyWithTVLLimits), maxPerDeposit);
        require(underlyingToken.balanceOf(address(strategyWithTVLLimits)) > maxTotalDeposits, "bad test setup");

        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBaseTVLLimits: max deposits exceeded"));
        strategyWithTVLLimits.deposit(underlyingToken, maxPerDeposit);
        cheats.stopPrank();
    }

    function testDepositValidAmount(uint256 depositAmount) public {
        maxTotalDeposits = 1e12;
        maxPerDeposit = 3e11;
        cheats.assume(depositAmount > 0);
        cheats.assume(depositAmount < maxPerDeposit);

        cheats.startPrank(pauser);
        strategyWithTVLLimits.setTVLLimits(maxPerDeposit, maxTotalDeposits);
        cheats.stopPrank();

        // we need to actually transfer the tokens to the strategy to avoid underflow in the `deposit` calculation
        underlyingToken.transfer(address(strategyWithTVLLimits), depositAmount);

        uint256 sharesBefore = strategyWithTVLLimits.totalShares();
        cheats.startPrank(address(strategyManager));
        strategyWithTVLLimits.deposit(underlyingToken, depositAmount);
        cheats.stopPrank();

        require(strategyWithTVLLimits.totalShares() == depositAmount + sharesBefore, "total shares not updated correctly");
    }

    /// @notice General-purpose test, re-useable, handles whether the deposit should revert or not and returns 'true' if it did revert.
    function testDeposit_WithTVLLimits(uint256 maxDepositsFuzzedInput, uint256 maxPerDepositFuzzedInput, uint256 depositAmount)
        public returns (bool depositReverted)
    {
        // need to filter this to make sure the deposit amounts can actually be transferred
        cheats.assume(depositAmount <= initialSupply);
        // set TVL limits
        cheats.startPrank(pauser);
        strategyWithTVLLimits.setTVLLimits(maxPerDepositFuzzedInput, maxDepositsFuzzedInput);
        cheats.stopPrank();

        // we need to calculate this before transferring tokens to the strategy
        uint256 expectedSharesOut = strategyWithTVLLimits.underlyingToShares(depositAmount);

        // we need to actually transfer the tokens to the strategy to avoid underflow in the `deposit` calculation
        underlyingToken.transfer(address(strategyWithTVLLimits), depositAmount);

        if (depositAmount > maxPerDepositFuzzedInput) {
            cheats.startPrank(address(strategyManager));
            cheats.expectRevert("StrategyBaseTVLLimits: max per deposit exceeded");
            strategyWithTVLLimits.deposit(underlyingToken, depositAmount);
            cheats.stopPrank();

            // transfer the tokens back from the strategy to not mess up the state
            cheats.startPrank(address(strategyWithTVLLimits));
            underlyingToken.transfer(address(this), depositAmount);
            cheats.stopPrank();

            // return 'true' since the call to `deposit` reverted
            return true;
        } else if (underlyingToken.balanceOf(address(strategyWithTVLLimits)) > maxDepositsFuzzedInput) {
            cheats.startPrank(address(strategyManager));
            cheats.expectRevert("StrategyBaseTVLLimits: max deposits exceeded");
            strategyWithTVLLimits.deposit(underlyingToken, depositAmount);
            cheats.stopPrank();

            // transfer the tokens back from the strategy to not mess up the state
            cheats.startPrank(address(strategyWithTVLLimits));
            underlyingToken.transfer(address(this), depositAmount);
            cheats.stopPrank();

            // return 'true' since the call to `deposit` reverted
            return true;
        } else {
            uint256 totalSharesBefore = strategyWithTVLLimits.totalShares();
            if (expectedSharesOut == 0) {
                cheats.startPrank(address(strategyManager));
                cheats.expectRevert("StrategyBase.deposit: newShares cannot be zero");
                strategyWithTVLLimits.deposit(underlyingToken, depositAmount);
                cheats.stopPrank();

                // transfer the tokens back from the strategy to not mess up the state
                cheats.startPrank(address(strategyWithTVLLimits));
                underlyingToken.transfer(address(this), depositAmount);
                cheats.stopPrank();

                // return 'true' since the call to `deposit` reverted
                return true;
            } else {
                cheats.startPrank(address(strategyManager));
                strategyWithTVLLimits.deposit(underlyingToken, depositAmount);
                cheats.stopPrank();

                require(strategyWithTVLLimits.totalShares() == expectedSharesOut + totalSharesBefore, "total shares not updated correctly");

                // return 'false' since the call to `deposit` succeeded
                return false;
            }
        }
    }

    /// OVERRIDING EXISTING TESTS TO FILTER INPUTS THAT WOULD FAIL DUE TO DEPOSIT-LIMITING
    modifier filterToValidDepositAmounts(uint256 amountToDeposit) {
        (uint256 _maxPerDeposit, uint256 _maxTotalDeposits) = strategyWithTVLLimits.getTVLLimits();
        cheats.assume(amountToDeposit <= _maxPerDeposit && amountToDeposit <= _maxTotalDeposits);
        _;
    }

    function testCanWithdrawDownToSmallShares(uint256 amountToDeposit, uint32 sharesToLeave) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testCanWithdrawDownToSmallShares(amountToDeposit, sharesToLeave);
    }

    function testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares(uint256 priorTotalShares, uint256 amountToDeposit
    ) public virtual override filterToValidDepositAmounts(priorTotalShares) filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares(priorTotalShares, amountToDeposit);
    }

    function testDepositWithZeroPriorBalanceAndZeroPriorShares(uint256 amountToDeposit) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);
    }

    function testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares(uint256 amountToDeposit, uint256 amountToTransfer, uint96 amountSharesToQuery
    ) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares(amountToDeposit, amountToTransfer, amountSharesToQuery);
    }

    function testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares(uint256 amountToDeposit, uint256 amountToTransfer, uint96 amountUnderlyingToQuery
    ) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares(amountToDeposit, amountToTransfer, amountUnderlyingToQuery);
    }

    function testWithdrawFailsWhenSharesGreaterThanTotalShares(uint256 amountToDeposit, uint256 sharesToWithdraw
    ) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testWithdrawFailsWhenSharesGreaterThanTotalShares(amountToDeposit, sharesToWithdraw);
    }

    function testWithdrawFailsWhenWithdrawalsPaused(uint256 amountToDeposit) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testWithdrawFailsWhenWithdrawalsPaused(amountToDeposit);
    }

    function testWithdrawWithPriorTotalSharesAndAmountSharesEqual(uint256 amountToDeposit) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testWithdrawWithPriorTotalSharesAndAmountSharesEqual(amountToDeposit);
    }

    function testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual(uint96 amountToDeposit, uint96 sharesToWithdraw
    ) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual(amountToDeposit, sharesToWithdraw);
    }
}