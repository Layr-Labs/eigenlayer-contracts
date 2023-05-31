// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./StrategyBaseUnit.t.sol";

import "../../contracts/strategies/StrategyBaseTVLLimits.sol";

contract StrategyBaseTVLLimitsUnitTests is StrategyBaseUnitTests {
    StrategyBaseTVLLimits public strategyBaseTVLLimitsImplementation;
    StrategyBaseTVLLimits public strategyWithTVLLimits;

    // defaults for tests, used in setup
    uint256 maxTotalDeposits = 3200e18;
    uint256 maxDepositPerAddress = 32e18;

    /// @notice Emitted when `maxDepositPerAddress` value is updated from `previousValue` to `newValue`
    event MaxDepositPerAddressUpdated(uint256 previousValue, uint256 newValue);

    /// @notice Emitted when `maxTotalDeposits` value is updated from `previousValue` to `newValue`
    event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue);

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
                    abi.encodeWithSelector(StrategyBaseTVLLimits.initialize.selector, maxDepositPerAddress, maxTotalDeposits, underlyingToken, pauserRegistry)
                )
            )
        );

        // verify that limits were set up correctly
        require(strategyWithTVLLimits.maxTotalDeposits() == maxTotalDeposits, "bad test setup");
        require(strategyWithTVLLimits.maxDepositPerAddress() == maxDepositPerAddress, "bad test setup");

        // replace the strategy from the non-TVL-limited tests with the TVL-limited strategy
        // this should result in all the StrategyBaseUnitTests also running against the `StrategyBaseTVLLimits` contract
        strategy = strategyWithTVLLimits;
    }

    function testSetTVLLimits(uint256 maxDepositPerAddressFuzzedInput, uint256 maxTotalDepositsFuzzedInput) public {
        _setTVLLimits(maxDepositPerAddressFuzzedInput, maxTotalDepositsFuzzedInput);
        (uint256 _maxDepositPerAddress, uint256 _maxDeposits) = strategyWithTVLLimits.getTVLLimits();

        assertEq(_maxDepositPerAddress, maxDepositPerAddressFuzzedInput);
        assertEq(_maxDeposits, maxTotalDepositsFuzzedInput);
    }

    function testSetTVLLimitsFailsWhenNotCalledByPauser(uint256 maxDepositPerAddressFuzzedInput, uint256 maxTotalDepositsFuzzedInput, address notPauser) public {
        cheats.assume(notPauser != address(proxyAdmin));
        cheats.assume(notPauser != pauser);
        cheats.startPrank(notPauser);
        cheats.expectRevert(bytes("msg.sender is not permissioned as pauser"));
        strategyWithTVLLimits.setTVLLimits(maxDepositPerAddressFuzzedInput, maxTotalDepositsFuzzedInput);
        cheats.stopPrank();
    }

    function testSetInvalidMaxDepositPerAddressAndMaxDeposits(uint256 maxDepositPerAddressFuzzedInput, uint256 maxTotalDepositsFuzzedInput) public {
        cheats.assume(maxTotalDepositsFuzzedInput < maxDepositPerAddressFuzzedInput);
        cheats.startPrank(pauser);
        cheats.expectRevert(bytes("StrategyBaseTVLLimits._setTVLLimits: maxDepositPerAddress exceeds maxTotalDeposits"));
        strategyWithTVLLimits.setTVLLimits(maxDepositPerAddressFuzzedInput, maxTotalDepositsFuzzedInput);
        cheats.stopPrank();
    }

    function testDepositMoreThanMaxDepositPerAddress(uint256 maxDepositPerAddressFuzzedInput, uint256 maxTotalDepositsFuzzedInput, uint256 amount) public {
        cheats.assume(amount > maxDepositPerAddressFuzzedInput);
        _setTVLLimits(maxDepositPerAddressFuzzedInput, maxTotalDepositsFuzzedInput);

        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBaseTVLLimits: max deposit per address exceeded"));
        strategyWithTVLLimits.deposit(depositor, underlyingToken, amount);
        cheats.stopPrank();
    }

    function testDepositMorethanMaxDeposits() public {
        maxTotalDeposits = 1e12;
        maxDepositPerAddress = 3e11;
        uint256 numDeposits = maxTotalDeposits / maxDepositPerAddress;

        _setTVLLimits(maxDepositPerAddress, maxTotalDeposits);

        underlyingToken.transfer(address(strategyManager), maxTotalDeposits);
        cheats.startPrank(address(strategyManager));
        for (uint256 i = 0; i < numDeposits; i++) {
            underlyingToken.transfer(address(strategyWithTVLLimits), maxDepositPerAddress);
            strategyWithTVLLimits.deposit(depositor, underlyingToken, maxDepositPerAddress);
        }
        cheats.stopPrank();

        underlyingToken.transfer(address(strategyWithTVLLimits), maxDepositPerAddress);
        require(underlyingToken.balanceOf(address(strategyWithTVLLimits)) > maxTotalDeposits, "bad test setup");

        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(bytes("StrategyBaseTVLLimits: max deposits exceeded"));
        strategyWithTVLLimits.deposit(depositor, underlyingToken, maxDepositPerAddress);
        cheats.stopPrank();
    }

    function testDepositValidAmount(uint256 depositAmount) public {
        maxTotalDeposits = 1e12;
        maxDepositPerAddress = 3e11;
        cheats.assume(depositAmount > 0);
        cheats.assume(depositAmount < maxDepositPerAddress);

        _setTVLLimits(maxDepositPerAddress, maxTotalDeposits);

        // we need to actually transfer the tokens to the strategy to avoid underflow in the `deposit` calculation
        underlyingToken.transfer(address(strategyWithTVLLimits), depositAmount);

        uint256 sharesBefore = strategyWithTVLLimits.totalShares();
        cheats.startPrank(address(strategyManager));
        strategyWithTVLLimits.deposit(depositor, underlyingToken, depositAmount);
        cheats.stopPrank();

        require(strategyWithTVLLimits.totalShares() == depositAmount + sharesBefore, "total shares not updated correctly");
    }

    function testDepositTVLLimit_ThenChangeTVLLimit(uint256 maxTotalDepositsFuzzedInput, uint256 newMaxTotalDepositsFuzzedInput) public {
        cheats.assume(maxTotalDepositsFuzzedInput > 0);
        cheats.assume(newMaxTotalDepositsFuzzedInput > maxTotalDepositsFuzzedInput);
        cheats.assume(newMaxTotalDepositsFuzzedInput < initialSupply);
        cheats.startPrank(pauser);
        strategyWithTVLLimits.setTVLLimits(maxTotalDepositsFuzzedInput, maxTotalDepositsFuzzedInput);
        cheats.stopPrank();

        underlyingToken.transfer(address(strategyWithTVLLimits), maxTotalDepositsFuzzedInput);

        uint256 sharesBefore = strategyWithTVLLimits.totalShares();

        cheats.startPrank(address(strategyManager));
        strategyWithTVLLimits.deposit(depositor, underlyingToken, maxTotalDepositsFuzzedInput);
        cheats.stopPrank();

        require(strategyWithTVLLimits.totalShares() == maxTotalDepositsFuzzedInput + sharesBefore, "total shares not updated correctly");

        cheats.startPrank(pauser);
        strategyWithTVLLimits.setTVLLimits(newMaxTotalDepositsFuzzedInput, newMaxTotalDepositsFuzzedInput);
        cheats.stopPrank();

        underlyingToken.transfer(address(strategyWithTVLLimits), newMaxTotalDepositsFuzzedInput - maxTotalDepositsFuzzedInput);

        sharesBefore = strategyWithTVLLimits.totalShares();
        cheats.startPrank(address(strategyManager));    
        strategyWithTVLLimits.deposit(depositor, underlyingToken, newMaxTotalDepositsFuzzedInput - maxTotalDepositsFuzzedInput);
        cheats.stopPrank();

        require(strategyWithTVLLimits.totalShares() == newMaxTotalDepositsFuzzedInput, "total shares not updated correctly");
    }

    /// @notice General-purpose test, re-useable, handles whether the deposit should revert or not and returns 'true' if it did revert.
    function testDeposit_WithTVLLimits(uint256 maxDepositPerAddressFuzzedInput, uint256 maxTotalDepositsFuzzedInput, uint256 depositAmount)
        public returns (bool depositReverted)
    {
        cheats.assume(maxDepositPerAddressFuzzedInput < maxTotalDepositsFuzzedInput);
        // need to filter this to make sure the deposit amounts can actually be transferred
        cheats.assume(depositAmount <= initialSupply);
        // set TVL limits
        _setTVLLimits(maxDepositPerAddressFuzzedInput, maxTotalDepositsFuzzedInput);

        // we need to calculate this before transferring tokens to the strategy
        uint256 expectedSharesOut = strategyWithTVLLimits.underlyingToShares(depositAmount);

        // we need to actually transfer the tokens to the strategy to avoid underflow in the `deposit` calculation
        underlyingToken.transfer(address(strategyWithTVLLimits), depositAmount);

        if (depositAmount > maxDepositPerAddressFuzzedInput) {
            cheats.startPrank(address(strategyManager));
            cheats.expectRevert("StrategyBaseTVLLimits: max per deposit exceeded");
            strategyWithTVLLimits.deposit(depositor, underlyingToken, depositAmount);
            cheats.stopPrank();

            // transfer the tokens back from the strategy to not mess up the state
            cheats.startPrank(address(strategyWithTVLLimits));
            underlyingToken.transfer(address(this), depositAmount);
            cheats.stopPrank();

            // return 'true' since the call to `deposit` reverted
            return true;
        } else if (underlyingToken.balanceOf(address(strategyWithTVLLimits)) > maxTotalDepositsFuzzedInput) {
            cheats.startPrank(address(strategyManager));
            cheats.expectRevert("StrategyBaseTVLLimits: max deposits exceeded");
            strategyWithTVLLimits.deposit(depositor, underlyingToken, depositAmount);
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
                strategyWithTVLLimits.deposit(depositor, underlyingToken, depositAmount);
                cheats.stopPrank();

                // transfer the tokens back from the strategy to not mess up the state
                cheats.startPrank(address(strategyWithTVLLimits));
                underlyingToken.transfer(address(this), depositAmount);
                cheats.stopPrank();

                // return 'true' since the call to `deposit` reverted
                return true;
            } else {
                cheats.startPrank(address(strategyManager));
                strategyWithTVLLimits.deposit(depositor, underlyingToken, depositAmount);
                cheats.stopPrank();

                require(strategyWithTVLLimits.totalShares() == expectedSharesOut + totalSharesBefore, "total shares not updated correctly");

                // return 'false' since the call to `deposit` succeeded
                return false;
            }
        }
    }

    // sets the TVL Limits and checks that events were emitted correctly
    function _setTVLLimits(uint256 _maxDepositPerAddress, uint256 _maxTotalDeposits) internal {
        cheats.assume(_maxDepositPerAddress < _maxTotalDeposits);
        (uint256 _maxDepositPerAddressBefore, uint256 _maxTotalDepositsBefore) = strategyWithTVLLimits.getTVLLimits();
        cheats.startPrank(pauser);
        cheats.expectEmit(true, true, true, true, address(strategyWithTVLLimits));
        emit MaxDepositPerAddressUpdated(_maxDepositPerAddressBefore, _maxDepositPerAddress);
        emit MaxTotalDepositsUpdated(_maxTotalDepositsBefore, _maxTotalDeposits);
        strategyWithTVLLimits.setTVLLimits(_maxDepositPerAddress, _maxTotalDeposits);
        cheats.stopPrank();
    }

    /// OVERRIDING EXISTING TESTS TO FILTER INPUTS THAT WOULD FAIL DUE TO DEPOSIT-LIMITING
    modifier filterToValidDepositAmounts(uint256 amountToDeposit) {
        (uint256 _maxDepositPerAddress, uint256 _maxTotalDeposits) = strategyWithTVLLimits.getTVLLimits();
        cheats.assume(amountToDeposit <= _maxDepositPerAddress && amountToDeposit <= _maxTotalDeposits);
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