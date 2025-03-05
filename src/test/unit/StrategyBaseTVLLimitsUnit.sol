// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./StrategyBaseUnit.t.sol";

import "../../contracts/strategies/StrategyBaseTVLLimits.sol";

contract StrategyBaseTVLLimitsUnitTests is StrategyBaseUnitTests {
    StrategyBaseTVLLimits public strategyBaseTVLLimitsImplementation;
    StrategyBaseTVLLimits public strategyWithTVLLimits;

    // defaults for tests, used in setup
    uint maxTotalDeposits = 3200e18;
    uint maxPerDeposit = 32e18;

    /// @notice Emitted when `maxPerDeposit` value is updated from `previousValue` to `newValue`
    event MaxPerDepositUpdated(uint previousValue, uint newValue);

    /// @notice Emitted when `maxTotalDeposits` value is updated from `previousValue` to `newValue`
    event MaxTotalDepositsUpdated(uint previousValue, uint newValue);

    function setUp() public virtual override {
        // copy setup for StrategyBaseUnitTests
        StrategyBaseUnitTests.setUp();

        // depoloy the TVL-limited strategy
        strategyBaseTVLLimitsImplementation = new StrategyBaseTVLLimits(strategyManager, pauserRegistry, "v9.9.9");
        strategyWithTVLLimits = StrategyBaseTVLLimits(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyBaseTVLLimitsImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        StrategyBaseTVLLimits.initialize.selector, maxPerDeposit, maxTotalDeposits, underlyingToken, pauserRegistry
                    )
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

    function testSetTVLLimits(uint maxPerDepositFuzzedInput, uint maxTotalDepositsFuzzedInput) public {
        _setTVLLimits(maxPerDepositFuzzedInput, maxTotalDepositsFuzzedInput);
        (uint _maxPerDeposit, uint _maxDeposits) = strategyWithTVLLimits.getTVLLimits();

        assertEq(_maxPerDeposit, maxPerDepositFuzzedInput);
        assertEq(_maxDeposits, maxTotalDepositsFuzzedInput);
    }

    function testSetTVLLimitsFailsWhenNotCalledByUnpauser(
        uint maxPerDepositFuzzedInput,
        uint maxTotalDepositsFuzzedInput,
        address notUnpauser
    ) public {
        cheats.assume(notUnpauser != address(proxyAdmin));
        cheats.assume(notUnpauser != unpauser);
        cheats.prank(notUnpauser);
        cheats.expectRevert(IPausable.OnlyUnpauser.selector);
        strategyWithTVLLimits.setTVLLimits(maxPerDepositFuzzedInput, maxTotalDepositsFuzzedInput);
    }

    function testSetInvalidMaxPerDepositAndMaxDeposits(uint maxPerDepositFuzzedInput, uint maxTotalDepositsFuzzedInput) public {
        cheats.assume(maxTotalDepositsFuzzedInput < maxPerDepositFuzzedInput);
        cheats.prank(unpauser);
        cheats.expectRevert(IStrategyErrors.MaxPerDepositExceedsMax.selector);
        strategyWithTVLLimits.setTVLLimits(maxPerDepositFuzzedInput, maxTotalDepositsFuzzedInput);
    }

    function testDepositMoreThanMaxPerDeposit(uint maxPerDepositFuzzedInput, uint maxTotalDepositsFuzzedInput, uint amount) public {
        cheats.assume(amount > maxPerDepositFuzzedInput);
        _setTVLLimits(maxPerDepositFuzzedInput, maxTotalDepositsFuzzedInput);

        cheats.prank(address(strategyManager));
        cheats.expectRevert(IStrategyErrors.MaxPerDepositExceedsMax.selector);
        strategyWithTVLLimits.deposit(underlyingToken, amount);
    }

    function testDepositMorethanMaxDeposits() public {
        maxTotalDeposits = 1e12;
        maxPerDeposit = 3e11;
        uint numDeposits = maxTotalDeposits / maxPerDeposit;

        _setTVLLimits(maxPerDeposit, maxTotalDeposits);

        underlyingToken.transfer(address(strategyManager), maxTotalDeposits);
        cheats.startPrank(address(strategyManager));
        for (uint i = 0; i < numDeposits; i++) {
            underlyingToken.transfer(address(strategyWithTVLLimits), maxPerDeposit);
            strategyWithTVLLimits.deposit(underlyingToken, maxPerDeposit);
        }
        cheats.stopPrank();

        underlyingToken.transfer(address(strategyWithTVLLimits), maxPerDeposit);
        require(underlyingToken.balanceOf(address(strategyWithTVLLimits)) > maxTotalDeposits, "bad test setup");

        cheats.prank(address(strategyManager));
        cheats.expectRevert(IStrategyErrors.BalanceExceedsMaxTotalDeposits.selector);
        strategyWithTVLLimits.deposit(underlyingToken, maxPerDeposit);
    }

    function testDepositValidAmount(uint depositAmount) public {
        maxTotalDeposits = 1e12;
        maxPerDeposit = 3e11;
        cheats.assume(depositAmount > 0);
        cheats.assume(depositAmount < maxPerDeposit);

        _setTVLLimits(maxPerDeposit, maxTotalDeposits);

        // we need to actually transfer the tokens to the strategy to avoid underflow in the `deposit` calculation
        underlyingToken.transfer(address(strategyWithTVLLimits), depositAmount);

        uint sharesBefore = strategyWithTVLLimits.totalShares();
        cheats.prank(address(strategyManager));
        strategyWithTVLLimits.deposit(underlyingToken, depositAmount);

        require(strategyWithTVLLimits.totalShares() == depositAmount + sharesBefore, "total shares not updated correctly");
    }

    function testDepositTVLLimit_ThenChangeTVLLimit(uint maxTotalDepositsFuzzedInput, uint newMaxTotalDepositsFuzzedInput) public {
        cheats.assume(maxTotalDepositsFuzzedInput > 0);
        cheats.assume(newMaxTotalDepositsFuzzedInput > maxTotalDepositsFuzzedInput);
        cheats.assume(newMaxTotalDepositsFuzzedInput < initialSupply);
        cheats.prank(unpauser);
        strategyWithTVLLimits.setTVLLimits(maxTotalDepositsFuzzedInput, maxTotalDepositsFuzzedInput);

        underlyingToken.transfer(address(strategyWithTVLLimits), maxTotalDepositsFuzzedInput);

        uint sharesBefore = strategyWithTVLLimits.totalShares();

        cheats.prank(address(strategyManager));
        strategyWithTVLLimits.deposit(underlyingToken, maxTotalDepositsFuzzedInput);

        require(strategyWithTVLLimits.totalShares() == maxTotalDepositsFuzzedInput + sharesBefore, "total shares not updated correctly");

        cheats.prank(unpauser);
        strategyWithTVLLimits.setTVLLimits(newMaxTotalDepositsFuzzedInput, newMaxTotalDepositsFuzzedInput);

        underlyingToken.transfer(address(strategyWithTVLLimits), newMaxTotalDepositsFuzzedInput - maxTotalDepositsFuzzedInput);

        sharesBefore = strategyWithTVLLimits.totalShares();
        cheats.prank(address(strategyManager));
        strategyWithTVLLimits.deposit(underlyingToken, newMaxTotalDepositsFuzzedInput - maxTotalDepositsFuzzedInput);

        require(strategyWithTVLLimits.totalShares() == newMaxTotalDepositsFuzzedInput, "total shares not updated correctly");
    }

    /// @notice General-purpose test, re-useable, handles whether the deposit should revert or not and returns 'true' if it did revert.
    function testDeposit_WithTVLLimits(uint maxPerDepositFuzzedInput, uint maxTotalDepositsFuzzedInput, uint depositAmount)
        public
        returns (bool depositReverted)
    {
        cheats.assume(maxPerDepositFuzzedInput < maxTotalDepositsFuzzedInput);
        // need to filter this to make sure the deposit amounts can actually be transferred
        cheats.assume(depositAmount <= initialSupply);
        // set TVL limits
        _setTVLLimits(maxPerDepositFuzzedInput, maxTotalDepositsFuzzedInput);

        // we need to calculate this before transferring tokens to the strategy
        uint expectedSharesOut = strategyWithTVLLimits.underlyingToShares(depositAmount);

        // we need to actually transfer the tokens to the strategy to avoid underflow in the `deposit` calculation
        underlyingToken.transfer(address(strategyWithTVLLimits), depositAmount);

        if (depositAmount > maxPerDepositFuzzedInput) {
            cheats.prank(address(strategyManager));
            cheats.expectRevert(IStrategyErrors.MaxPerDepositExceedsMax.selector);
            strategyWithTVLLimits.deposit(underlyingToken, depositAmount);

            // transfer the tokens back from the strategy to not mess up the state
            cheats.prank(address(strategyWithTVLLimits));
            underlyingToken.transfer(address(this), depositAmount);

            // return 'true' since the call to `deposit` reverted
            return true;
        } else if (underlyingToken.balanceOf(address(strategyWithTVLLimits)) > maxTotalDepositsFuzzedInput) {
            cheats.prank(address(strategyManager));
            cheats.expectRevert(IStrategyErrors.MaxPerDepositExceedsMax.selector);
            strategyWithTVLLimits.deposit(underlyingToken, depositAmount);

            // transfer the tokens back from the strategy to not mess up the state
            cheats.prank(address(strategyWithTVLLimits));
            underlyingToken.transfer(address(this), depositAmount);

            // return 'true' since the call to `deposit` reverted
            return true;
        } else {
            uint totalSharesBefore = strategyWithTVLLimits.totalShares();
            if (expectedSharesOut == 0) {
                cheats.prank(address(strategyManager));
                cheats.expectRevert(IStrategyErrors.NewSharesZero.selector);
                strategyWithTVLLimits.deposit(underlyingToken, depositAmount);

                // transfer the tokens back from the strategy to not mess up the state
                cheats.prank(address(strategyWithTVLLimits));
                underlyingToken.transfer(address(this), depositAmount);

                // return 'true' since the call to `deposit` reverted
                return true;
            } else {
                cheats.prank(address(strategyManager));
                strategyWithTVLLimits.deposit(underlyingToken, depositAmount);

                require(strategyWithTVLLimits.totalShares() == expectedSharesOut + totalSharesBefore, "total shares not updated correctly");

                // return 'false' since the call to `deposit` succeeded
                return false;
            }
        }
    }

    // sets the TVL Limits and checks that events were emitted correctly
    function _setTVLLimits(uint _maxPerDeposit, uint _maxTotalDeposits) internal {
        cheats.assume(_maxPerDeposit < _maxTotalDeposits);
        (uint _maxPerDepositBefore, uint _maxTotalDepositsBefore) = strategyWithTVLLimits.getTVLLimits();
        cheats.expectEmit(true, true, true, true, address(strategyWithTVLLimits));
        cheats.prank(unpauser);
        emit MaxPerDepositUpdated(_maxPerDepositBefore, _maxPerDeposit);
        emit MaxTotalDepositsUpdated(_maxTotalDepositsBefore, _maxTotalDeposits);
        strategyWithTVLLimits.setTVLLimits(_maxPerDeposit, _maxTotalDeposits);
    }

    /// OVERRIDING EXISTING TESTS TO FILTER INPUTS THAT WOULD FAIL DUE TO DEPOSIT-LIMITING
    modifier filterToValidDepositAmounts(uint amountToDeposit) {
        (uint _maxPerDeposit, uint _maxTotalDeposits) = strategyWithTVLLimits.getTVLLimits();
        cheats.assume(amountToDeposit <= _maxPerDeposit && amountToDeposit <= _maxTotalDeposits);
        _;
    }

    function testCanWithdrawDownToSmallShares(uint amountToDeposit, uint32 sharesToLeave)
        public
        virtual
        override
        filterToValidDepositAmounts(amountToDeposit)
    {
        StrategyBaseUnitTests.testCanWithdrawDownToSmallShares(amountToDeposit, sharesToLeave);
    }

    function testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares(uint priorTotalShares, uint amountToDeposit)
        public
        virtual
        override
        filterToValidDepositAmounts(priorTotalShares)
        filterToValidDepositAmounts(amountToDeposit)
    {
        StrategyBaseUnitTests.testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares(priorTotalShares, amountToDeposit);
    }

    function testDepositWithZeroPriorBalanceAndZeroPriorShares(uint amountToDeposit)
        public
        virtual
        override
        filterToValidDepositAmounts(amountToDeposit)
    {
        StrategyBaseUnitTests.testDepositWithZeroPriorBalanceAndZeroPriorShares(amountToDeposit);
    }

    function testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares(
        uint amountToDeposit,
        uint amountToTransfer,
        uint96 amountSharesToQuery
    ) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares(
            amountToDeposit, amountToTransfer, amountSharesToQuery
        );
    }

    function testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares(
        uint amountToDeposit,
        uint amountToTransfer,
        uint96 amountUnderlyingToQuery
    ) public virtual override filterToValidDepositAmounts(amountToDeposit) {
        StrategyBaseUnitTests.testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares(
            amountToDeposit, amountToTransfer, amountUnderlyingToQuery
        );
    }

    function testWithdrawFailsWhenSharesGreaterThanTotalShares(uint amountToDeposit, uint sharesToWithdraw)
        public
        virtual
        override
        filterToValidDepositAmounts(amountToDeposit)
    {
        StrategyBaseUnitTests.testWithdrawFailsWhenSharesGreaterThanTotalShares(amountToDeposit, sharesToWithdraw);
    }

    function testWithdrawFailsWhenWithdrawalsPaused(uint amountToDeposit)
        public
        virtual
        override
        filterToValidDepositAmounts(amountToDeposit)
    {
        StrategyBaseUnitTests.testWithdrawFailsWhenWithdrawalsPaused(amountToDeposit);
    }

    function testWithdrawWithPriorTotalSharesAndAmountSharesEqual(uint amountToDeposit)
        public
        virtual
        override
        filterToValidDepositAmounts(amountToDeposit)
    {
        StrategyBaseUnitTests.testWithdrawWithPriorTotalSharesAndAmountSharesEqual(amountToDeposit);
    }

    function testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual(uint96 amountToDeposit, uint96 sharesToWithdraw)
        public
        virtual
        override
        filterToValidDepositAmounts(amountToDeposit)
    {
        StrategyBaseUnitTests.testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual(amountToDeposit, sharesToWithdraw);
    }
}
