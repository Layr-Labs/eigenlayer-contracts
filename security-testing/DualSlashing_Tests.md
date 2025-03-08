# Dual Slashing Security Tests

## Beacon Chain and AVS Slashing Interaction Tests

### Test Case: Sequential Slashing
- **Description**: Test behavior when an operator is slashed by both beacon chain and AVS
- **Implementation**:
  ```solidity
  function test_DualSlashing_SequentialSlashing() public {
      // Setup operator with delegated stakers and deposits
      // ...
      
      // Record initial withdrawable shares
      uint256 initialWithdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Slash operator on beacon chain
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      
      // Record withdrawable shares after beacon chain slashing
      uint256 afterBeaconSlashingShares = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Verify beacon chain slashing reduced withdrawable shares
      assertLt(afterBeaconSlashingShares, initialWithdrawableShares);
      
      // Slash operator by AVS
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Record withdrawable shares after AVS slashing
      uint256 afterAVSSlashingShares = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Verify AVS slashing further reduced withdrawable shares
      assertLt(afterAVSSlashingShares, afterBeaconSlashingShares);
  }
  ```

### Test Case: Simultaneous Slashing
- **Description**: Test behavior when an operator is slashed by both beacon chain and AVS simultaneously
- **Implementation**:
  ```solidity
  function test_DualSlashing_SimultaneousSlashing() public {
      // Setup operator with delegated stakers and deposits
      // ...
      
      // Record initial withdrawable shares
      uint256 initialWithdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Calculate expected withdrawable shares after dual slashing
      uint256 beaconSlashingFactor = 0.7e18; // 30% slash
      uint256 avsSlashingFactor = 0.5e18; // 50% slash
      uint256 expectedWithdrawableShares = initialWithdrawableShares
          .mulWad(beaconSlashingFactor)
          .mulWad(avsSlashingFactor);
      
      // Perform both slashings
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Record actual withdrawable shares after dual slashing
      uint256 actualWithdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Verify actual shares are close to expected (accounting for rounding)
      uint256 difference = expectedWithdrawableShares > actualWithdrawableShares ?
          expectedWithdrawableShares - actualWithdrawableShares :
          actualWithdrawableShares - expectedWithdrawableShares;
      
      // Difference should be very small relative to initial shares
      assertLe(difference, initialWithdrawableShares / 1e6);
  }
  ```

### Test Case: Full Slashing
- **Description**: Test behavior when an operator is fully slashed (100%)
- **Implementation**:
  ```solidity
  function test_DualSlashing_FullSlashing() public {
      // Setup operator with delegated stakers and deposits
      // ...
      
      // Fully slash the operator
      allocationManager.slashOperator(operator, operatorSetId, 1e18); // 100% slash
      
      // Verify withdrawable shares are zero
      uint256 withdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      assertEq(withdrawableShares, 0);
      
      // Verify staker cannot delegate to fully slashed operator
      address newStaker = address(0x789);
      vm.startPrank(newStaker);
      vm.expectRevert("Cannot delegate to fully slashed operator");
      delegationManager.delegate(operator);
      vm.stopPrank();
  }
  ```

## Complex Share Scaling Mechanism Tests

### Test Case: Share Scaling After Multiple Slashings
- **Description**: Test that share scaling factors are correctly updated after multiple slashing events
- **Implementation**:
  ```solidity
  function test_DualSlashing_ShareScalingAfterMultipleSlashings() public {
      // Setup operator with delegated stakers and deposits
      // ...
      
      // Record initial deposit scaling factor
      DepositScalingFactor initialDSF = delegationManager.getDepositScalingFactor(operator, strategy);
      
      // First slash - beacon chain 30%
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      
      // Record deposit scaling factor after first slash
      DepositScalingFactor dsfAfterFirstSlash = delegationManager.getDepositScalingFactor(operator, strategy);
      
      // Verify deposit scaling factor increased
      assertGt(dsfAfterFirstSlash.scalingFactor(), initialDSF.scalingFactor());
      
      // Second slash - AVS 50%
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Record deposit scaling factor after second slash
      DepositScalingFactor dsfAfterSecondSlash = delegationManager.getDepositScalingFactor(operator, strategy);
      
      // Verify deposit scaling factor increased further
      assertGt(dsfAfterSecondSlash.scalingFactor(), dsfAfterFirstSlash.scalingFactor());
      
      // Deposit after slashing
      uint256 depositAmount = 1e18;
      // ...
      
      // Verify new deposit receives correct shares based on updated scaling factor
      // ...
  }
  ```

### Test Case: Share Scaling with Extreme Values
- **Description**: Test share scaling with extreme values to ensure no overflow/underflow
- **Implementation**:
  ```solidity
  function test_DualSlashing_ShareScalingExtremeValues() public {
      // Setup operator with delegated stakers and deposits
      // ...
      
      // Perform multiple small slashings to increase deposit scaling factor
      for (uint i = 0; i < 10; i++) {
          eigenPodManager.slashBeaconChainETH(operator, 0.1e18);
          allocationManager.slashOperator(operator, operatorSetId, 0.1e18);
      }
      
      // Record deposit scaling factor after multiple slashings
      DepositScalingFactor dsf = delegationManager.getDepositScalingFactor(operator, strategy);
      
      // Verify deposit scaling factor is still within bounds
      assertLt(dsf.scalingFactor(), 1e74); // Upper bound from SharesAccountingEdgeCases.md
      
      // Deposit very large amount
      uint256 largeDeposit = type(uint128).max;
      // ...
      
      // Verify deposit scaling factor calculation doesn't overflow
      // ...
      
      // Verify withdrawable shares calculation doesn't overflow
      uint256 withdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      assertGt(withdrawableShares, 0);
  }
  ```

### Test Case: Share Scaling with Rebasing Tokens
- **Description**: Test interaction between share scaling and rebasing tokens
- **Implementation**:
  ```solidity
  function test_DualSlashing_ShareScalingWithRebasingTokens() public {
      // Setup operator with delegated stakers and rebasing token deposits
      // ...
      
      // Record initial withdrawable shares
      uint256 initialWithdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Trigger rebasing event (increase supply by 10%)
      rebasingToken.rebase(110); // 10% increase
      
      // Record withdrawable shares after rebasing
      uint256 withdrawableSharesAfterRebasing = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Verify rebasing increased withdrawable shares
      assertGt(withdrawableSharesAfterRebasing, initialWithdrawableShares);
      
      // Slash operator
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Record withdrawable shares after slashing
      uint256 withdrawableSharesAfterSlashing = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Verify slashing reduced withdrawable shares
      assertLt(withdrawableSharesAfterSlashing, withdrawableSharesAfterRebasing);
      
      // Trigger another rebasing event
      rebasingToken.rebase(110); // Another 10% increase
      
      // Record withdrawable shares after second rebasing
      uint256 withdrawableSharesAfterSecondRebasing = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Verify rebasing increased withdrawable shares again
      assertGt(withdrawableSharesAfterSecondRebasing, withdrawableSharesAfterSlashing);
  }
  ```

## Edge Cases in Dual Slashing Scenarios

### Test Case: Slashing After Withdrawal Queue
- **Description**: Test behavior when operator is slashed after staker has queued a withdrawal
- **Implementation**:
  ```solidity
  function test_DualSlashing_SlashingAfterWithdrawalQueue() public {
      // Setup operator with delegated stakers and deposits
      // ...
      
      // Queue withdrawal
      uint256 withdrawalIndex = delegationManager.queueWithdrawal(...);
      
      // Record queued withdrawal amount
      uint256 queuedWithdrawalAmount = delegationManager.getQueuedWithdrawalAmount(withdrawalIndex);
      
      // Slash operator
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Complete withdrawal
      delegationManager.completeWithdrawal(withdrawalIndex);
      
      // Verify withdrawn amount is equal to queued withdrawal amount
      // (slashing should not affect already queued withdrawals)
      uint256 withdrawnAmount = token.balanceOf(staker);
      assertEq(withdrawnAmount, queuedWithdrawalAmount);
  }
  ```

### Test Case: Deposit After Dual Slashing
- **Description**: Test behavior when depositing after operator has been slashed by both beacon chain and AVS
- **Implementation**:
  ```solidity
  function test_DualSlashing_DepositAfterDualSlashing() public {
      // Setup operator with delegated stakers and deposits
      // ...
      
      // Slash operator
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Record deposit scaling factor after slashing
      DepositScalingFactor dsf = delegationManager.getDepositScalingFactor(operator, strategy);
      
      // Deposit after slashing
      uint256 depositAmount = 1e18;
      // ...
      
      // Verify deposit received correct shares based on slashed scaling factor
      uint256 expectedShares = depositAmount.divWad(dsf.scalingFactor());
      uint256 actualShares = delegationManager.operatorShares(operator, strategy) - initialOperatorShares;
      
      // Allow for small rounding errors
      uint256 tolerance = depositAmount / 1e6;
      assertApproxEqAbs(actualShares, expectedShares, tolerance);
      
      // Verify withdrawable shares reflect the slashing
      uint256 withdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      uint256 expectedWithdrawableShares = depositAmount.mulWad(0.7e18).mulWad(0.5e18); // 30% and 50% slashing
      
      assertApproxEqAbs(withdrawableShares, expectedWithdrawableShares, tolerance);
  }
  ```

### Test Case: Slashing After Undelegate
- **Description**: Test behavior when operator is slashed after staker has undelegated
- **Implementation**:
  ```solidity
  function test_DualSlashing_SlashingAfterUndelegate() public {
      // Setup operator with delegated stakers and deposits
      // ...
      
      // Queue withdrawal and undelegate
      uint256 withdrawalIndex = delegationManager.queueWithdrawal(...);
      vm.prank(staker);
      delegationManager.undelegate();
      
      // Slash operator
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Complete withdrawal
      delegationManager.completeWithdrawal(withdrawalIndex);
      
      // Verify withdrawn amount is equal to queued withdrawal amount
      // (slashing should not affect already queued withdrawals)
      uint256 withdrawnAmount = token.balanceOf(staker);
      assertEq(withdrawnAmount, queuedWithdrawalAmount);
      
      // Verify staker is no longer delegated
      assertEq(delegationManager.isDelegated(staker), false);
  }
  ```

### Test Case: Dual Slashing with Multiple Strategies
- **Description**: Test behavior when operator is slashed across multiple strategies
- **Implementation**:
  ```solidity
  function test_DualSlashing_MultipleStrategies() public {
      // Setup operator with delegated stakers and deposits in multiple strategies
      // ...
      
      // Record initial withdrawable shares for each strategy
      uint256 initialWithdrawableShares1 = delegationManager.getWithdrawableShares(staker, strategy1);
      uint256 initialWithdrawableShares2 = delegationManager.getWithdrawableShares(staker, strategy2);
      
      // Slash operator on beacon chain (affects all strategies)
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      
      // Slash operator by AVS for specific strategy
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Record withdrawable shares after slashing
      uint256 withdrawableSharesAfterSlashing1 = delegationManager.getWithdrawableShares(staker, strategy1);
      uint256 withdrawableSharesAfterSlashing2 = delegationManager.getWithdrawableShares(staker, strategy2);
      
      // Verify beacon chain slashing affected both strategies
      assertLt(withdrawableSharesAfterSlashing1, initialWithdrawableShares1);
      assertLt(withdrawableSharesAfterSlashing2, initialWithdrawableShares2);
      
      // Verify AVS slashing only affected the specific strategy
      assertEq(
          withdrawableSharesAfterSlashing1,
          initialWithdrawableShares1.mulWad(0.7e18).mulWad(0.5e18) // Both slashings
      );
      assertEq(
          withdrawableSharesAfterSlashing2,
          initialWithdrawableShares2.mulWad(0.7e18) // Only beacon chain slashing
      );
  }
  ```

### Test Case: Dual Slashing with Zero Shares
- **Description**: Test behavior when operator has zero shares in a strategy
- **Implementation**:
  ```solidity
  function test_DualSlashing_ZeroShares() public {
      // Setup operator with no delegated shares
      // ...
      
      // Verify operator has zero shares
      uint256 operatorShares = delegationManager.operatorShares(operator, strategy);
      assertEq(operatorShares, 0);
      
      // Slash operator
      eigenPodManager.slashBeaconChainETH(operator, 0.3e18);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      
      // Verify operator still has zero shares
      operatorShares = delegationManager.operatorShares(operator, strategy);
      assertEq(operatorShares, 0);
      
      // Deposit after slashing
      uint256 depositAmount = 1e18;
      // ...
      
      // Verify deposit received correct shares based on slashed scaling factor
      // ...
  }
  ```
