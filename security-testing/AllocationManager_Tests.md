# AllocationManager Security Tests

## Allocation Security Tests

### Test Case: Allocation Validation
- **Description**: Test that allocations properly validate inputs and handle edge cases
- **Implementation**:
  ```solidity
  function test_AllocationManager_AllocationValidation() public {
      // Test allocation with invalid operator
      vm.expectRevert("AllocationManager.allocate: operator not registered");
      allocationManager.allocate(invalidOperator, operatorSetId, 1e18);
      
      // Test allocation with invalid operatorSetId
      vm.expectRevert("AllocationManager.allocate: operatorSet not registered");
      allocationManager.allocate(operator, invalidOperatorSetId, 1e18);
      
      // Test allocation with zero shares
      vm.expectRevert("AllocationManager.allocate: cannot allocate zero shares");
      allocationManager.allocate(operator, operatorSetId, 0);
      
      // Test allocation with more shares than delegated
      vm.expectRevert("AllocationManager.allocate: cannot allocate more shares than delegated");
      allocationManager.allocate(operator, operatorSetId, 1e30); // Very large amount
  }
  ```

### Test Case: Allocation Limits
- **Description**: Test that allocations respect limits
- **Implementation**:
  ```solidity
  function test_AllocationManager_AllocationLimits() public {
      // Setup operator with delegated shares
      // ...
      
      // Record initial delegated shares
      uint256 delegatedShares = delegationManager.operatorShares(operator, strategy);
      
      // Allocate maximum allowed shares
      allocationManager.allocate(operator, operatorSetId, delegatedShares);
      
      // Verify allocation was successful
      uint256 allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(allocatedShares, delegatedShares);
      
      // Attempt to allocate more shares
      vm.expectRevert("AllocationManager.allocate: cannot allocate more shares than delegated");
      allocationManager.allocate(operator, operatorSetId, 1);
      
      // Deallocate some shares
      allocationManager.deallocate(operator, operatorSetId, delegatedShares / 2);
      
      // Verify deallocation was successful
      allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(allocatedShares, delegatedShares / 2);
      
      // Allocate remaining shares
      allocationManager.allocate(operator, operatorSetId, delegatedShares / 2);
      
      // Verify allocation was successful
      allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(allocatedShares, delegatedShares);
  }
  ```

### Test Case: Allocation Authorization
- **Description**: Test that only authorized users can allocate
- **Implementation**:
  ```solidity
  function test_AllocationManager_AllocationAuthorization() public {
      // Setup operator with delegated shares
      // ...
      
      // Attempt allocation as non-operator
      address nonOperator = address(0x789);
      vm.startPrank(nonOperator);
      vm.expectRevert("AllocationManager.onlyOperator: not operator");
      allocationManager.allocate(operator, operatorSetId, 1e18);
      vm.stopPrank();
      
      // Allocate as operator
      vm.startPrank(operator);
      allocationManager.allocate(operator, operatorSetId, 1e18);
      vm.stopPrank();
      
      // Verify allocation was successful
      uint256 allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(allocatedShares, 1e18);
  }
  ```

### Test Case: Allocation After Slashing
- **Description**: Test allocation behavior after operator is slashed
- **Implementation**:
  ```solidity
  function test_AllocationManager_AllocationAfterSlashing() public {
      // Setup operator with delegated shares and initial allocation
      // ...
      
      // Record initial allocation
      uint256 initialAllocation = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      
      // Slash operator
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18); // 50% slash
      
      // Record allocation after slashing
      uint256 allocationAfterSlashing = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      
      // Verify allocation was reduced by 50%
      assertEq(allocationAfterSlashing, initialAllocation / 2);
      
      // Attempt to allocate more shares
      uint256 additionalShares = 1e18;
      vm.startPrank(operator);
      allocationManager.allocate(operator, operatorSetId, additionalShares);
      vm.stopPrank();
      
      // Verify new allocation
      uint256 newAllocation = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(newAllocation, allocationAfterSlashing + additionalShares);
  }
  ```

## Deallocation Security Tests

### Test Case: Deallocation Validation
- **Description**: Test that deallocations properly validate inputs and handle edge cases
- **Implementation**:
  ```solidity
  function test_AllocationManager_DeallocationValidation() public {
      // Setup operator with delegated shares and allocation
      // ...
      
      // Test deallocation with invalid operator
      vm.expectRevert("AllocationManager.deallocate: operator not registered");
      allocationManager.deallocate(invalidOperator, operatorSetId, 1e18);
      
      // Test deallocation with invalid operatorSetId
      vm.expectRevert("AllocationManager.deallocate: operatorSet not registered");
      allocationManager.deallocate(operator, invalidOperatorSetId, 1e18);
      
      // Test deallocation with zero shares
      vm.expectRevert("AllocationManager.deallocate: cannot deallocate zero shares");
      allocationManager.deallocate(operator, operatorSetId, 0);
      
      // Test deallocation with more shares than allocated
      uint256 allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      vm.expectRevert("AllocationManager.deallocate: cannot deallocate more shares than allocated");
      allocationManager.deallocate(operator, operatorSetId, allocatedShares + 1);
  }
  ```

### Test Case: Deallocation Authorization
- **Description**: Test that only authorized users can deallocate
- **Implementation**:
  ```solidity
  function test_AllocationManager_DeallocationAuthorization() public {
      // Setup operator with delegated shares and allocation
      // ...
      
      // Attempt deallocation as non-operator
      address nonOperator = address(0x789);
      vm.startPrank(nonOperator);
      vm.expectRevert("AllocationManager.onlyOperator: not operator");
      allocationManager.deallocate(operator, operatorSetId, 1e18);
      vm.stopPrank();
      
      // Deallocate as operator
      vm.startPrank(operator);
      allocationManager.deallocate(operator, operatorSetId, 1e18);
      vm.stopPrank();
      
      // Verify deallocation was successful
      uint256 allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(allocatedShares, initialAllocation - 1e18);
  }
  ```

### Test Case: Complete Deallocation
- **Description**: Test behavior when all shares are deallocated
- **Implementation**:
  ```solidity
  function test_AllocationManager_CompleteDeallocation() public {
      // Setup operator with delegated shares and allocation
      // ...
      
      // Record initial allocation
      uint256 initialAllocation = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      
      // Deallocate all shares
      vm.startPrank(operator);
      allocationManager.deallocate(operator, operatorSetId, initialAllocation);
      vm.stopPrank();
      
      // Verify allocation is zero
      uint256 allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(allocatedShares, 0);
      
      // Verify operator is still registered
      bool isRegistered = allocationManager.isOperator(operator);
      assertTrue(isRegistered);
  }
  ```

## Operator Set Management Tests

### Test Case: Operator Registration
- **Description**: Test operator registration process
- **Implementation**:
  ```solidity
  function test_AllocationManager_OperatorRegistration() public {
      // Setup new operator
      address newOperator = address(0x123);
      
      // Register operator
      vm.startPrank(newOperator);
      allocationManager.registerOperator();
      vm.stopPrank();
      
      // Verify operator is registered
      bool isRegistered = allocationManager.isOperator(newOperator);
      assertTrue(isRegistered);
      
      // Attempt to register again
      vm.startPrank(newOperator);
      vm.expectRevert("AllocationManager.registerOperator: already registered");
      allocationManager.registerOperator();
      vm.stopPrank();
  }
  ```

### Test Case: Operator Deregistration
- **Description**: Test operator deregistration process
- **Implementation**:
  ```solidity
  function test_AllocationManager_OperatorDeregistration() public {
      // Setup registered operator
      // ...
      
      // Verify operator is registered
      bool isRegistered = allocationManager.isOperator(operator);
      assertTrue(isRegistered);
      
      // Deregister operator
      vm.startPrank(operator);
      allocationManager.deregisterOperator();
      vm.stopPrank();
      
      // Verify operator is no longer registered
      isRegistered = allocationManager.isOperator(operator);
      assertFalse(isRegistered);
      
      // Attempt to deregister again
      vm.startPrank(operator);
      vm.expectRevert("AllocationManager.deregisterOperator: not registered");
      allocationManager.deregisterOperator();
      vm.stopPrank();
  }
  ```

### Test Case: Operator Set Registration
- **Description**: Test operator set registration process
- **Implementation**:
  ```solidity
  function test_AllocationManager_OperatorSetRegistration() public {
      // Setup new operator set
      address operatorSetOwner = address(0x123);
      
      // Register operator set
      vm.startPrank(operatorSetOwner);
      uint256 newOperatorSetId = allocationManager.registerOperatorSet();
      vm.stopPrank();
      
      // Verify operator set is registered
      bool isRegistered = allocationManager.isOperatorSetRegistered(newOperatorSetId);
      assertTrue(isRegistered);
      
      // Verify operator set owner
      address owner = allocationManager.getOperatorSetOwner(newOperatorSetId);
      assertEq(owner, operatorSetOwner);
  }
  ```

### Test Case: Operator Set Deregistration
- **Description**: Test operator set deregistration process
- **Implementation**:
  ```solidity
  function test_AllocationManager_OperatorSetDeregistration() public {
      // Setup registered operator set
      // ...
      
      // Verify operator set is registered
      bool isRegistered = allocationManager.isOperatorSetRegistered(operatorSetId);
      assertTrue(isRegistered);
      
      // Deregister operator set
      vm.startPrank(operatorSetOwner);
      allocationManager.deregisterOperatorSet(operatorSetId);
      vm.stopPrank();
      
      // Verify operator set is no longer registered
      isRegistered = allocationManager.isOperatorSetRegistered(operatorSetId);
      assertFalse(isRegistered);
      
      // Attempt to deregister again
      vm.startPrank(operatorSetOwner);
      vm.expectRevert("AllocationManager.deregisterOperatorSet: not registered");
      allocationManager.deregisterOperatorSet(operatorSetId);
      vm.stopPrank();
  }
  ```

## Slashing Mechanism Tests

### Test Case: Slashing Mechanism Security
- **Description**: Test that slashing properly validates inputs and handles edge cases
- **Implementation**:
  ```solidity
  function test_AllocationManager_SlashingMechanism() public {
      // Setup allocations
      // ...
      
      // Test slashing with invalid operatorSetId
      vm.expectRevert("AllocationManager.slashOperator: operatorSet not registered");
      allocationManager.slashOperator(operator, invalidOperatorSetId, 0.5e18);
      
      // Test slashing with invalid operator
      vm.expectRevert("AllocationManager.slashOperator: operator not registered");
      allocationManager.slashOperator(invalidOperator, operatorSetId, 0.5e18);
      
      // Test slashing with zero amount
      vm.expectRevert("AllocationManager.slashOperator: cannot slash zero amount");
      allocationManager.slashOperator(operator, operatorSetId, 0);
      
      // Test slashing with amount > WAD
      vm.expectRevert("AllocationManager.slashOperator: cannot slash more than 100%");
      allocationManager.slashOperator(operator, operatorSetId, 1.1e18);
  }
  ```

### Test Case: Slashing Authorization
- **Description**: Test that only authorized users can slash operators
- **Implementation**:
  ```solidity
  function test_AllocationManager_SlashingAuthorization() public {
      // Setup allocations
      // ...
      
      // Attempt slashing as non-owner
      address nonOwner = address(0x789);
      vm.startPrank(nonOwner);
      vm.expectRevert("AllocationManager.onlyOperatorSetOwner: not owner");
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      vm.stopPrank();
      
      // Slash as operator set owner
      vm.startPrank(operatorSetOwner);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      vm.stopPrank();
      
      // Verify slashing was successful
      uint256 newMagnitude = allocationManager.getOperatorAllocation(operator, operatorSetId).currentMagnitude;
      assertEq(newMagnitude, WAD / 2); // 50% of WAD
  }
  ```

### Test Case: Rounding Up on Slashing
- **Description**: Test that slashing rounds up to prevent DOS attacks
- **Implementation**:
  ```solidity
  function test_AllocationManager_RoundingUpOnSlashing() public {
      // Setup allocation with very small amount
      // ...
      
      // Record initial magnitude
      uint256 initialMagnitude = allocationManager.getOperatorAllocation(operator, operatorSetId).currentMagnitude;
      assertEq(initialMagnitude, WAD); // Should be WAD initially
      
      // Slash with very small percentage
      uint256 slashAmount = 1; // Smallest non-zero value
      
      // This should still result in some slashing
      vm.startPrank(operatorSetOwner);
      allocationManager.slashOperator(operator, operatorSetId, slashAmount);
      vm.stopPrank();
      
      // Verify operator was slashed by at least 1 wei
      uint256 newMagnitude = allocationManager.getOperatorAllocation(operator, operatorSetId).currentMagnitude;
      assertLt(newMagnitude, initialMagnitude);
  }
  ```

### Test Case: Multiple Slashing
- **Description**: Test behavior when operator is slashed multiple times
- **Implementation**:
  ```solidity
  function test_AllocationManager_MultipleSlashing() public {
      // Setup allocations
      // ...
      
      // Record initial magnitude
      uint256 initialMagnitude = allocationManager.getOperatorAllocation(operator, operatorSetId).currentMagnitude;
      assertEq(initialMagnitude, WAD); // Should be WAD initially
      
      // First slash - 50%
      vm.startPrank(operatorSetOwner);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      vm.stopPrank();
      
      // Verify first slash
      uint256 magnitudeAfterFirstSlash = allocationManager.getOperatorAllocation(operator, operatorSetId).currentMagnitude;
      assertEq(magnitudeAfterFirstSlash, WAD / 2); // 50% of WAD
      
      // Second slash - 50% of remaining
      vm.startPrank(operatorSetOwner);
      allocationManager.slashOperator(operator, operatorSetId, 0.5e18);
      vm.stopPrank();
      
      // Verify second slash
      uint256 magnitudeAfterSecondSlash = allocationManager.getOperatorAllocation(operator, operatorSetId).currentMagnitude;
      assertEq(magnitudeAfterSecondSlash, WAD / 4); // 25% of original WAD
  }
  ```

### Test Case: Full Slashing
- **Description**: Test behavior when operator is fully slashed
- **Implementation**:
  ```solidity
  function test_AllocationManager_FullSlashing() public {
      // Setup allocations
      // ...
      
      // Fully slash operator
      vm.startPrank(operatorSetOwner);
      allocationManager.slashOperator(operator, operatorSetId, 1e18); // 100% slash
      vm.stopPrank();
      
      // Verify operator is fully slashed
      uint256 newMagnitude = allocationManager.getOperatorAllocation(operator, operatorSetId).currentMagnitude;
      assertEq(newMagnitude, 0);
      
      // Verify operator can still deallocate
      vm.startPrank(operator);
      allocationManager.deallocate(operator, operatorSetId, 1e18);
      vm.stopPrank();
      
      // Verify deallocation was successful
      uint256 allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(allocatedShares, initialAllocation - 1e18);
  }
  ```

## Registration/Deregistration Security Tests

### Test Case: Registration with Existing Delegation
- **Description**: Test registration process for an operator with existing delegation
- **Implementation**:
  ```solidity
  function test_AllocationManager_RegistrationWithExistingDelegation() public {
      // Setup operator with delegation but not registered in AllocationManager
      // ...
      
      // Register operator
      vm.startPrank(operator);
      allocationManager.registerOperator();
      vm.stopPrank();
      
      // Verify operator is registered
      bool isRegistered = allocationManager.isOperator(operator);
      assertTrue(isRegistered);
      
      // Verify operator can allocate shares
      vm.startPrank(operator);
      allocationManager.allocate(operator, operatorSetId, 1e18);
      vm.stopPrank();
      
      // Verify allocation was successful
      uint256 allocatedShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(allocatedShares, 1e18);
  }
  ```

### Test Case: Deregistration with Existing Allocations
- **Description**: Test deregistration process for an operator with existing allocations
- **Implementation**:
  ```solidity
  function test_AllocationManager_DeregistrationWithExistingAllocations() public {
      // Setup operator with allocations
      // ...
      
      // Attempt to deregister operator with allocations
      vm.startPrank(operator);
      vm.expectRevert("AllocationManager.deregisterOperator: has allocations");
      allocationManager.deregisterOperator();
      vm.stopPrank();
      
      // Deallocate all shares
      vm.startPrank(operator);
      allocationManager.deallocate(operator, operatorSetId, allocatedShares);
      vm.stopPrank();
      
      // Verify deallocation was successful
      uint256 remainingShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(remainingShares, 0);
      
      // Now deregister operator
      vm.startPrank(operator);
      allocationManager.deregisterOperator();
      vm.stopPrank();
      
      // Verify operator is no longer registered
      bool isRegistered = allocationManager.isOperator(operator);
      assertFalse(isRegistered);
  }
  ```

### Test Case: Operator Set Deregistration with Existing Allocations
- **Description**: Test deregistration process for an operator set with existing allocations
- **Implementation**:
  ```solidity
  function test_AllocationManager_OperatorSetDeregistrationWithExistingAllocations() public {
      // Setup operator set with allocations
      // ...
      
      // Attempt to deregister operator set with allocations
      vm.startPrank(operatorSetOwner);
      vm.expectRevert("AllocationManager.deregisterOperatorSet: has allocations");
      allocationManager.deregisterOperatorSet(operatorSetId);
      vm.stopPrank();
      
      // Deallocate all shares
      vm.startPrank(operator);
      allocationManager.deallocate(operator, operatorSetId, allocatedShares);
      vm.stopPrank();
      
      // Verify deallocation was successful
      uint256 remainingShares = allocationManager.getOperatorAllocation(operator, operatorSetId).currentShares;
      assertEq(remainingShares, 0);
      
      // Now deregister operator set
      vm.startPrank(operatorSetOwner);
      allocationManager.deregisterOperatorSet(operatorSetId);
      vm.stopPrank();
      
      // Verify operator set is no longer registered
      bool isRegistered = allocationManager.isOperatorSetRegistered(operatorSetId);
      assertFalse(isRegistered);
  }
  ```

### Test Case: Registration/Deregistration Reentrancy
- **Description**: Test that registration and deregistration functions are protected against reentrancy attacks
- **Implementation**:
  ```solidity
  function test_AllocationManager_RegistrationReentrancy() public {
      // Setup a malicious contract that attempts reentrancy
      MaliciousReentrancyContract attacker = new MaliciousReentrancyContract(allocationManager);
      
      // Attempt reentrancy attack on registerOperator function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackRegisterOperator();
      
      // Attempt reentrancy attack on deregisterOperator function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackDeregisterOperator();
      
      // Attempt reentrancy attack on registerOperatorSet function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackRegisterOperatorSet();
      
      // Attempt reentrancy attack on deregisterOperatorSet function
      vm.expectRevert("ReentrancyGuard: reentrant call");
      attacker.attackDeregisterOperatorSet();
  }
  ```
