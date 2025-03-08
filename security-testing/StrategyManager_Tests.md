# StrategyManager Security Tests

## Deposit and Withdrawal Security Tests

### Test Case: Deposit Validation
- **Description**: Test that deposits properly validate inputs and handle edge cases
- **Implementation**:
  ```solidity
  function test_StrategyManager_DepositValidation() public {
      // Test with zero amount
      vm.expectRevert("StrategyManager.depositIntoStrategy: cannot deposit zero tokens");
      strategyManager.depositIntoStrategy(strategy, token, 0);
      
      // Test with invalid strategy
      vm.expectRevert("StrategyManager.onlyStrategiesWhitelistedForDeposit: strategy not whitelisted");
      strategyManager.depositIntoStrategy(invalidStrategy, token, 1e18);
      
      // Test with non-existent token
      vm.expectRevert();
      strategyManager.depositIntoStrategy(strategy, address(0), 1e18);
      
      // Test with insufficient allowance
      vm.expectRevert("ERC20: insufficient allowance");
      strategyManager.depositIntoStrategy(strategy, token, 1e18);
  }
  ```

### Test Case: Withdrawal Security
- **Description**: Test that withdrawals properly validate inputs and handle edge cases
- **Implementation**:
  ```solidity
  function test_StrategyManager_WithdrawalSecurity() public {
      // Setup deposits
      // ...
      
      // Test withdrawal with invalid shares amount
      vm.expectRevert("StrategyManager.queueWithdrawal: withdrawing more shares than staker has");
      strategyManager.queueWithdrawal(
          [strategy],
          [token],
          [1e18 + 1], // More than deposited
          address(this)
      );
      
      // Test withdrawal with mismatched arrays
      vm.expectRevert("StrategyManager.queueWithdrawal: input length mismatch");
      strategyManager.queueWithdrawal(
          [strategy, strategy],
          [token],
          [1e18],
          address(this)
      );
      
      // Test withdrawal to zero address
      vm.expectRevert("StrategyManager.queueWithdrawal: cannot withdraw to zero address");
      strategyManager.queueWithdrawal(
          [strategy],
          [token],
          [1e18],
          address(0)
      );
  }
  ```

### Test Case: Withdrawal Queue Manipulation
- **Description**: Test that withdrawal queue cannot be manipulated
- **Implementation**:
  ```solidity
  function test_StrategyManager_WithdrawalQueueManipulation() public {
      // Setup deposits and queue withdrawal
      // ...
      uint256 withdrawalIndex = strategyManager.queueWithdrawal(
          [strategy],
          [token],
          [1e18],
          address(this)
      );
      
      // Attempt to complete withdrawal with invalid index
      vm.expectRevert("StrategyManager.completeQueuedWithdrawal: withdrawal does not exist");
      strategyManager.completeQueuedWithdrawal(
          withdrawalIndex + 1,
          [strategy],
          [token],
          [1e18],
          address(this)
      );
      
      // Attempt to complete withdrawal with incorrect parameters
      vm.expectRevert("StrategyManager.completeQueuedWithdrawal: withdrawal params do not match");
      strategyManager.completeQueuedWithdrawal(
          withdrawalIndex,
          [strategy, strategy], // Different array
          [token],
          [1e18],
          address(this)
      );
      
      // Attempt to complete withdrawal as non-withdrawer
      address attacker = address(0x789);
      vm.prank(attacker);
      vm.expectRevert("StrategyManager.completeQueuedWithdrawal: only withdrawer can complete");
      strategyManager.completeQueuedWithdrawal(
          withdrawalIndex,
          [strategy],
          [token],
          [1e18],
          address(this)
      );
  }
  ```

### Test Case: Reentrancy in Deposit and Withdrawal
- **Description**: Test that deposit and withdrawal functions are protected against reentrancy attacks
- **Implementation**:
  ```solidity
  function test_StrategyManager_ReentrancyProtection() public {
      // Setup a malicious contract that attempts reentrancy
      MaliciousReentrancyContract attacker = new MaliciousReentrancyContract(strategyManager);
      
      // Fund the attacker
      deal(token, address(attacker), 10e18);
      
      // Attempt reentrancy attack on deposit function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackDeposit(strategy, token, 1e18);
      
      // Setup for withdrawal reentrancy
      // ...
      
      // Attempt reentrancy attack on withdrawal function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackWithdrawal();
  }
  ```

## Share Calculation Tests

### Test Case: Share Calculation Precision
- **Description**: Test precision of share calculations for different token types
- **Implementation**:
  ```solidity
  function testFuzz_StrategyManager_ShareCalculationPrecision(uint256 depositAmount) public {
      // Bound input to realistic range
      depositAmount = bound(depositAmount, 1, 1e30);
      
      // Deposit tokens
      uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
      deal(token, staker, depositAmount);
      
      vm.startPrank(staker);
      token.approve(address(strategyManager), depositAmount);
      uint256 sharesReceived = strategyManager.depositIntoStrategy(strategy, token, depositAmount);
      vm.stopPrank();
      
      // Verify shares calculation
      uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
      assertEq(sharesAfter, sharesBefore + sharesReceived);
      
      // Verify shares can be withdrawn
      vm.startPrank(staker);
      strategyManager.queueWithdrawal(
          [strategy],
          [token],
          [sharesReceived],
          staker
      );
      vm.stopPrank();
  }
  ```

### Test Case: Share Calculation for Different Token Types
- **Description**: Test share calculations for different token types (standard, rebasing, fee-on-transfer)
- **Implementation**:
  ```solidity
  function test_StrategyManager_ShareCalculationDifferentTokens() public {
      // Setup different token types
      address standardToken = address(new StandardERC20());
      address rebasingToken = address(new RebasingERC20());
      address feeToken = address(new FeeOnTransferERC20());
      
      // Create strategies for each token
      address standardStrategy = address(new MockStrategy(standardToken));
      address rebasingStrategy = address(new MockStrategy(rebasingToken));
      address feeStrategy = address(new MockStrategy(feeToken));
      
      // Whitelist strategies
      // ...
      
      // Test standard token
      deal(standardToken, staker, 1e18);
      vm.startPrank(staker);
      IERC20(standardToken).approve(address(strategyManager), 1e18);
      uint256 standardShares = strategyManager.depositIntoStrategy(standardStrategy, standardToken, 1e18);
      vm.stopPrank();
      
      // Test rebasing token
      deal(rebasingToken, staker, 1e18);
      vm.startPrank(staker);
      IERC20(rebasingToken).approve(address(strategyManager), 1e18);
      uint256 rebasingShares = strategyManager.depositIntoStrategy(rebasingStrategy, rebasingToken, 1e18);
      vm.stopPrank();
      
      // Test fee-on-transfer token
      deal(feeToken, staker, 1e18);
      vm.startPrank(staker);
      IERC20(feeToken).approve(address(strategyManager), 1e18);
      uint256 feeShares = strategyManager.depositIntoStrategy(feeStrategy, feeToken, 1e18);
      vm.stopPrank();
      
      // Verify shares are calculated correctly for each token type
      // ...
  }
  ```

### Test Case: Share Calculation with Extreme Values
- **Description**: Test share calculations with extreme values to ensure no overflow/underflow
- **Implementation**:
  ```solidity
  function test_StrategyManager_ShareCalculationExtremeValues() public {
      // Test with very small deposit
      uint256 smallDeposit = 1; // 1 wei
      deal(token, staker, smallDeposit);
      
      vm.startPrank(staker);
      token.approve(address(strategyManager), smallDeposit);
      uint256 smallShares = strategyManager.depositIntoStrategy(strategy, token, smallDeposit);
      vm.stopPrank();
      
      // Verify small deposit results in non-zero shares
      assertGt(smallShares, 0);
      
      // Test with very large deposit
      uint256 largeDeposit = type(uint128).max;
      deal(token, staker, largeDeposit);
      
      vm.startPrank(staker);
      token.approve(address(strategyManager), largeDeposit);
      uint256 largeShares = strategyManager.depositIntoStrategy(strategy, token, largeDeposit);
      vm.stopPrank();
      
      // Verify large deposit doesn't overflow
      assertGt(largeShares, 0);
      
      // Verify shares can be withdrawn
      vm.startPrank(staker);
      strategyManager.queueWithdrawal(
          [strategy],
          [token],
          [largeShares],
          staker
      );
      vm.stopPrank();
  }
  ```

## Strategy Upgrade Security Tests

### Test Case: Strategy Upgrade Authorization
- **Description**: Test that only authorized users can upgrade strategies
- **Implementation**:
  ```solidity
  function test_StrategyManager_StrategyUpgradeAuthorization() public {
      // Setup initial strategy
      // ...
      
      // Create new strategy implementation
      address newStrategyImpl = address(new MockStrategy(token));
      
      // Attempt upgrade as non-admin
      address nonAdmin = address(0x789);
      vm.startPrank(nonAdmin);
      vm.expectRevert("PermissionController.onlyPermitted: not permitted");
      strategyManager.upgradeStrategy(strategy, newStrategyImpl);
      vm.stopPrank();
      
      // Upgrade as admin
      address admin = address(0x123);
      vm.startPrank(admin);
      strategyManager.upgradeStrategy(strategy, newStrategyImpl);
      vm.stopPrank();
      
      // Verify upgrade was successful
      // ...
  }
  ```

### Test Case: Strategy Upgrade State Preservation
- **Description**: Test that strategy upgrades preserve state
- **Implementation**:
  ```solidity
  function test_StrategyManager_StrategyUpgradeStatePreservation() public {
      // Setup initial strategy and deposit
      // ...
      
      // Record initial state
      uint256 initialTotalShares = IStrategy(strategy).totalShares();
      uint256 initialUnderlyingTokenBalance = IERC20(token).balanceOf(strategy);
      
      // Create new strategy implementation
      address newStrategyImpl = address(new MockStrategy(token));
      
      // Upgrade strategy
      vm.prank(admin);
      strategyManager.upgradeStrategy(strategy, newStrategyImpl);
      
      // Verify state is preserved
      uint256 newTotalShares = IStrategy(strategy).totalShares();
      uint256 newUnderlyingTokenBalance = IERC20(token).balanceOf(strategy);
      
      assertEq(newTotalShares, initialTotalShares);
      assertEq(newUnderlyingTokenBalance, initialUnderlyingTokenBalance);
      
      // Verify functionality still works after upgrade
      // Deposit
      deal(token, staker, 1e18);
      vm.startPrank(staker);
      token.approve(address(strategyManager), 1e18);
      strategyManager.depositIntoStrategy(strategy, token, 1e18);
      vm.stopPrank();
      
      // Verify deposit increased total shares
      assertGt(IStrategy(strategy).totalShares(), newTotalShares);
  }
  ```

### Test Case: Strategy Upgrade Security Checks
- **Description**: Test that strategy upgrades include proper security checks
- **Implementation**:
  ```solidity
  function test_StrategyManager_StrategyUpgradeSecurityChecks() public {
      // Setup initial strategy
      // ...
      
      // Attempt upgrade to invalid implementation (zero address)
      vm.startPrank(admin);
      vm.expectRevert("StrategyManager.upgradeStrategy: implementation cannot be zero address");
      strategyManager.upgradeStrategy(strategy, address(0));
      vm.stopPrank();
      
      // Attempt upgrade to implementation with different underlying token
      address differentToken = address(new StandardERC20());
      address invalidImpl = address(new MockStrategy(differentToken));
      
      vm.startPrank(admin);
      vm.expectRevert("StrategyManager.upgradeStrategy: new implementation has different underlying token");
      strategyManager.upgradeStrategy(strategy, invalidImpl);
      vm.stopPrank();
      
      // Attempt upgrade to non-contract address
      address nonContract = address(0x123);
      
      vm.startPrank(admin);
      vm.expectRevert("StrategyManager.upgradeStrategy: implementation must be a contract");
      strategyManager.upgradeStrategy(strategy, nonContract);
      vm.stopPrank();
  }
  ```

## Rebasing Token Handling Tests

### Test Case: Rebasing Token Deposit and Withdrawal
- **Description**: Test deposit and withdrawal with rebasing tokens
- **Implementation**:
  ```solidity
  function test_StrategyManager_RebasingTokenHandling() public {
      // Setup rebasing token and strategy
      RebasingERC20 rebasingToken = new RebasingERC20();
      address rebasingStrategy = address(new MockStrategy(address(rebasingToken)));
      
      // Whitelist strategy
      // ...
      
      // Initial deposit
      uint256 initialDeposit = 1e18;
      deal(address(rebasingToken), staker, initialDeposit);
      
      vm.startPrank(staker);
      rebasingToken.approve(address(strategyManager), initialDeposit);
      uint256 sharesReceived = strategyManager.depositIntoStrategy(rebasingStrategy, address(rebasingToken), initialDeposit);
      vm.stopPrank();
      
      // Trigger rebasing event (increase supply by 10%)
      rebasingToken.rebase(110); // 10% increase
      
      // Verify strategy's underlying token balance increased
      uint256 newBalance = rebasingToken.balanceOf(rebasingStrategy);
      assertEq(newBalance, initialDeposit * 110 / 100);
      
      // Withdraw shares
      vm.startPrank(staker);
      uint256 withdrawalIndex = strategyManager.queueWithdrawal(
          [rebasingStrategy],
          [address(rebasingToken)],
          [sharesReceived],
          staker
      );
      vm.stopPrank();
      
      // Complete withdrawal
      // ...
      
      // Verify withdrawn amount includes rebasing gains
      uint256 withdrawnAmount = rebasingToken.balanceOf(staker);
      assertEq(withdrawnAmount, initialDeposit * 110 / 100);
  }
  ```

### Test Case: Rebasing Token Share Value
- **Description**: Test that share value is correctly calculated for rebasing tokens
- **Implementation**:
  ```solidity
  function test_StrategyManager_RebasingTokenShareValue() public {
      // Setup rebasing token and strategy
      RebasingERC20 rebasingToken = new RebasingERC20();
      MockStrategy rebasingStrategy = new MockStrategy(address(rebasingToken));
      
      // Whitelist strategy
      // ...
      
      // Initial deposit
      uint256 initialDeposit = 1e18;
      deal(address(rebasingToken), staker, initialDeposit);
      
      vm.startPrank(staker);
      rebasingToken.approve(address(strategyManager), initialDeposit);
      uint256 sharesReceived = strategyManager.depositIntoStrategy(address(rebasingStrategy), address(rebasingToken), initialDeposit);
      vm.stopPrank();
      
      // Record initial share value
      uint256 initialShareValue = rebasingStrategy.sharesToUnderlyingView(sharesReceived);
      assertEq(initialShareValue, initialDeposit);
      
      // Trigger rebasing event (increase supply by 10%)
      rebasingToken.rebase(110); // 10% increase
      
      // Verify share value increased
      uint256 newShareValue = rebasingStrategy.sharesToUnderlyingView(sharesReceived);
      assertEq(newShareValue, initialDeposit * 110 / 100);
      
      // Deposit again after rebasing
      uint256 secondDeposit = 1e18;
      deal(address(rebasingToken), staker, secondDeposit);
      
      vm.startPrank(staker);
      rebasingToken.approve(address(strategyManager), secondDeposit);
      uint256 sharesReceived2 = strategyManager.depositIntoStrategy(address(rebasingStrategy), address(rebasingToken), secondDeposit);
      vm.stopPrank();
      
      // Verify fewer shares received for same deposit amount due to increased share value
      assertLt(sharesReceived2, sharesReceived);
  }
  ```

### Test Case: Negative Rebasing
- **Description**: Test behavior with negative rebasing (token supply decreases)
- **Implementation**:
  ```solidity
  function test_StrategyManager_NegativeRebasing() public {
      // Setup rebasing token and strategy
      RebasingERC20 rebasingToken = new RebasingERC20();
      address rebasingStrategy = address(new MockStrategy(address(rebasingToken)));
      
      // Whitelist strategy
      // ...
      
      // Initial deposit
      uint256 initialDeposit = 1e18;
      deal(address(rebasingToken), staker, initialDeposit);
      
      vm.startPrank(staker);
      rebasingToken.approve(address(strategyManager), initialDeposit);
      uint256 sharesReceived = strategyManager.depositIntoStrategy(rebasingStrategy, address(rebasingToken), initialDeposit);
      vm.stopPrank();
      
      // Trigger negative rebasing event (decrease supply by 10%)
      rebasingToken.rebase(90); // 10% decrease
      
      // Verify strategy's underlying token balance decreased
      uint256 newBalance = rebasingToken.balanceOf(rebasingStrategy);
      assertEq(newBalance, initialDeposit * 90 / 100);
      
      // Withdraw shares
      vm.startPrank(staker);
      uint256 withdrawalIndex = strategyManager.queueWithdrawal(
          [rebasingStrategy],
          [address(rebasingToken)],
          [sharesReceived],
          staker
      );
      vm.stopPrank();
      
      // Complete withdrawal
      // ...
      
      // Verify withdrawn amount reflects rebasing loss
      uint256 withdrawnAmount = rebasingToken.balanceOf(staker);
      assertEq(withdrawnAmount, initialDeposit * 90 / 100);
  }
  ```
