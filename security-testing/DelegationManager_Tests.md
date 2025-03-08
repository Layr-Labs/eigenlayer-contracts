# DelegationManager Security Tests

## Reentrancy Protection Tests

### Test Case: Reentrancy in Delegation
- **Description**: Test that delegation functions are protected against reentrancy attacks
- **Implementation**:
  ```solidity
  function test_DelegationManager_ReentrancyProtection_Delegate() public {
      // Setup a malicious contract that attempts reentrancy
      MaliciousReentrancyContract attacker = new MaliciousReentrancyContract(delegationManager);
      
      // Fund the attacker
      deal(address(attacker), 10 ether);
      
      // Attempt reentrancy attack on delegate function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackDelegate(operator);
  }
  ```

### Test Case: Reentrancy in Withdrawal
- **Description**: Test that withdrawal functions are protected against reentrancy attacks
- **Implementation**:
  ```solidity
  function test_DelegationManager_ReentrancyProtection_QueueWithdrawal() public {
      // Setup a malicious contract that attempts reentrancy
      MaliciousReentrancyContract attacker = new MaliciousReentrancyContract(delegationManager);
      
      // Setup delegation and deposits
      // ...
      
      // Attempt reentrancy attack on queueWithdrawal function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackQueueWithdrawal();
  }
  ```

### Test Case: Reentrancy in Undelegate
- **Description**: Test that undelegate function is protected against reentrancy attacks
- **Implementation**:
  ```solidity
  function test_DelegationManager_ReentrancyProtection_Undelegate() public {
      // Setup a malicious contract that attempts reentrancy
      MaliciousReentrancyContract attacker = new MaliciousReentrancyContract(delegationManager);
      
      // Setup delegation
      vm.prank(address(attacker));
      delegationManager.delegate(operator);
      
      // Attempt reentrancy attack on undelegate function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackUndelegate();
  }
  ```

## Delegation State Transition Tests

### Test Case: Delegation State Consistency
- **Description**: Test that delegation state remains consistent through various operations
- **Implementation**:
  ```solidity
  function test_DelegationManager_StateConsistency() public {
      // Setup initial state
      address staker = address(0x123);
      address operator = address(0x456);
      
      // Delegate
      vm.startPrank(staker);
      delegationManager.delegate(operator);
      
      // Verify delegation state
      assertEq(delegationManager.isDelegated(staker), true);
      assertEq(delegationManager.delegatedTo(staker), operator);
      
      // Undelegate
      delegationManager.undelegate();
      
      // Verify state after undelegation
      assertEq(delegationManager.isDelegated(staker), false);
      assertEq(delegationManager.delegatedTo(staker), address(0));
      vm.stopPrank();
  }
  ```

### Test Case: Delegation State After Operator Slashing
- **Description**: Test that delegation state remains consistent after operator is slashed
- **Implementation**:
  ```solidity
  function test_DelegationManager_StateAfterSlashing() public {
      // Setup initial state
      address staker = address(0x123);
      address operator = address(0x456);
      
      // Delegate
      vm.prank(staker);
      delegationManager.delegate(operator);
      
      // Slash operator
      // ...
      
      // Verify delegation state is still consistent
      assertEq(delegationManager.isDelegated(staker), true);
      assertEq(delegationManager.delegatedTo(staker), operator);
      
      // Verify withdrawable shares are reduced
      // ...
  }
  ```

### Test Case: Multiple Delegations and Undelegations
- **Description**: Test that multiple delegations and undelegations maintain consistent state
- **Implementation**:
  ```solidity
  function test_DelegationManager_MultipleDelegations() public {
      // Setup initial state
      address staker = address(0x123);
      address operator1 = address(0x456);
      address operator2 = address(0x789);
      
      // First delegation
      vm.startPrank(staker);
      delegationManager.delegate(operator1);
      
      // Verify delegation state
      assertEq(delegationManager.isDelegated(staker), true);
      assertEq(delegationManager.delegatedTo(staker), operator1);
      
      // Undelegate
      delegationManager.undelegate();
      
      // Verify state after undelegation
      assertEq(delegationManager.isDelegated(staker), false);
      assertEq(delegationManager.delegatedTo(staker), address(0));
      
      // Second delegation
      delegationManager.delegate(operator2);
      
      // Verify delegation state
      assertEq(delegationManager.isDelegated(staker), true);
      assertEq(delegationManager.delegatedTo(staker), operator2);
      vm.stopPrank();
  }
  ```

## Withdrawal Queue Processing Tests

### Test Case: Withdrawal Queue Ordering
- **Description**: Test that withdrawals are processed in the correct order
- **Implementation**:
  ```solidity
  function test_DelegationManager_WithdrawalQueueOrdering() public {
      // Setup initial state with multiple stakers and deposits
      // ...
      
      // Queue multiple withdrawals
      uint256[] memory withdrawalIndices = new uint256[](3);
      withdrawalIndices[0] = delegationManager.queueWithdrawal(...);
      withdrawalIndices[1] = delegationManager.queueWithdrawal(...);
      withdrawalIndices[2] = delegationManager.queueWithdrawal(...);
      
      // Verify withdrawals are processed in order
      for (uint i = 0; i < withdrawalIndices.length; i++) {
          // Complete withdrawal
          delegationManager.completeWithdrawal(withdrawalIndices[i]);
          
          // Verify withdrawal was processed correctly
          // ...
      }
  }
  ```

### Test Case: Withdrawal After Slashing
- **Description**: Test that withdrawals after slashing return the correct amount
- **Implementation**:
  ```solidity
  function test_DelegationManager_WithdrawalAfterSlashing() public {
      // Setup initial state with staker, operator, and deposits
      // ...
      
      // Record initial withdrawable shares
      uint256 initialWithdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Slash operator
      // ...
      
      // Queue withdrawal
      uint256 withdrawalIndex = delegationManager.queueWithdrawal(...);
      
      // Record withdrawable shares after slashing
      uint256 withdrawableSharesAfterSlashing = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Verify withdrawable shares are reduced
      assertLt(withdrawableSharesAfterSlashing, initialWithdrawableShares);
      
      // Complete withdrawal
      delegationManager.completeWithdrawal(withdrawalIndex);
      
      // Verify withdrawal amount matches withdrawable shares after slashing
      // ...
  }
  ```

### Test Case: Withdrawal Queue Manipulation
- **Description**: Test that withdrawal queue cannot be manipulated
- **Implementation**:
  ```solidity
  function test_DelegationManager_WithdrawalQueueManipulation() public {
      // Setup initial state
      // ...
      
      // Queue withdrawal
      uint256 withdrawalIndex = delegationManager.queueWithdrawal(...);
      
      // Attempt to complete withdrawal with invalid index
      vm.expectRevert("DelegationManager.completeWithdrawal: withdrawal does not exist");
      delegationManager.completeWithdrawal(withdrawalIndex + 1);
      
      // Attempt to complete withdrawal with someone else's index
      address otherStaker = address(0x789);
      vm.prank(otherStaker);
      vm.expectRevert("DelegationManager.completeWithdrawal: not withdrawer");
      delegationManager.completeWithdrawal(withdrawalIndex);
  }
  ```

## Slashed Operator Handling Tests

### Test Case: Delegation to Fully Slashed Operator
- **Description**: Test behavior when attempting to delegate to a fully slashed operator
- **Implementation**:
  ```solidity
  function test_DelegationManager_DelegateToFullySlashedOperator() public {
      // Setup a fully slashed operator
      address operator = address(0x456);
      
      // Slash the operator fully for a strategy
      // ...
      
      // Attempt to delegate to the fully slashed operator
      address staker = address(0x123);
      vm.startPrank(staker);
      
      // This should revert with appropriate error
      vm.expectRevert("Cannot delegate to fully slashed operator");
      delegationManager.delegate(operator);
      vm.stopPrank();
  }
  ```

### Test Case: Withdrawable Shares After Slashing
- **Description**: Test that withdrawable shares are correctly calculated after slashing
- **Implementation**:
  ```solidity
  function test_DelegationManager_WithdrawableSharesAfterSlashing() public {
      // Setup initial state with staker, operator, and deposits
      // ...
      
      // Record initial withdrawable shares
      uint256 initialWithdrawableShares = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Slash operator by 50%
      uint256 slashingPercentage = 0.5e18; // 50% in WAD
      // ...
      
      // Record withdrawable shares after slashing
      uint256 withdrawableSharesAfterSlashing = delegationManager.getWithdrawableShares(staker, strategy);
      
      // Verify withdrawable shares are reduced by approximately 50%
      uint256 expectedWithdrawableShares = initialWithdrawableShares.mulWad(1e18 - slashingPercentage);
      uint256 tolerance = initialWithdrawableShares / 1000; // 0.1% tolerance for rounding
      
      assertApproxEqAbs(
          withdrawableSharesAfterSlashing,
          expectedWithdrawableShares,
          tolerance
      );
  }
  ```

### Test Case: Deposit After Operator Slashing
- **Description**: Test behavior when depositing to a slashed operator
- **Implementation**:
  ```solidity
  function test_DelegationManager_DepositAfterOperatorSlashing() public {
      // Setup initial state with staker, operator, and deposits
      // ...
      
      // Slash operator
      // ...
      
      // Deposit additional assets
      uint256 additionalDeposit = 1e18;
      // ...
      
      // Verify deposit scaling factor is adjusted correctly
      // ...
      
      // Verify withdrawable shares reflect the slashing
      // ...
  }
  ```

### Test Case: Operator Shares Invariant After Slashing
- **Description**: Test that operator shares invariant is maintained after slashing
- **Implementation**:
  ```solidity
  function test_DelegationManager_OperatorSharesInvariantAfterSlashing() public {
      // Setup multiple stakers delegated to the same operator
      // ...
      
      // Record total operator shares
      uint256 operatorShares = delegationManager.operatorShares(operator, strategy);
      
      // Record sum of withdrawable shares for all stakers
      uint256 totalWithdrawableShares = 0;
      for (uint i = 0; i < stakers.length; i++) {
          totalWithdrawableShares += delegationManager.getWithdrawableShares(stakers[i], strategy);
      }
      
      // Slash operator
      // ...
      
      // Record new operator shares
      uint256 newOperatorShares = delegationManager.operatorShares(operator, strategy);
      
      // Record new sum of withdrawable shares
      uint256 newTotalWithdrawableShares = 0;
      for (uint i = 0; i < stakers.length; i++) {
          newTotalWithdrawableShares += delegationManager.getWithdrawableShares(stakers[i], strategy);
      }
      
      // Verify invariant: operatorShares >= sum(withdrawableShares)
      assertGe(newOperatorShares, newTotalWithdrawableShares);
      
      // Verify both values decreased after slashing
      assertLt(newOperatorShares, operatorShares);
      assertLt(newTotalWithdrawableShares, totalWithdrawableShares);
  }
  ```
