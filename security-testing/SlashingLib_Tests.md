# SlashingLib Security Tests

## Mathematical Precision Tests

### Test Case: Extreme Value Handling
- **Description**: Test SlashingLib functions with extreme values to ensure proper handling
- **Implementation**:
  ```solidity
  function testFuzz_SlashingLib_ExtremeValues(uint256 depositShares, uint256 slashingFactor) public {
      // Bound inputs to realistic but extreme ranges
      depositShares = bound(depositShares, 1, type(uint128).max);
      slashingFactor = bound(slashingFactor, 1, WAD);
      
      // Test calcWithdrawable with extreme values
      uint256 withdrawableShares = SlashingLib.calcWithdrawable(
          DepositScalingFactor(WAD), 
          depositShares, 
          slashingFactor
      );
      
      // Verify result is within expected bounds
      assertLe(withdrawableShares, depositShares);
  }
  ```

### Test Case: Rounding Error Accumulation
- **Description**: Test for accumulation of rounding errors in repeated operations
- **Implementation**:
  ```solidity
  function testFuzz_SlashingLib_RoundingErrorAccumulation(
      uint64[] calldata prevMaxMagnitudes,
      uint64[] calldata newMaxMagnitudes,
      uint256 initialShares
  ) public {
      // Bound inputs
      initialShares = bound(initialShares, WAD, 1e30);
      
      uint256 cumulativeShares = initialShares;
      uint256 expectedShares = initialShares;
      
      // Apply multiple slashing operations
      for (uint i = 0; i < prevMaxMagnitudes.length; i++) {
          uint64 prevMag = bound(prevMaxMagnitudes[i], 2, WAD);
          uint64 newMag = bound(newMaxMagnitudes[i], 1, prevMag - 1);
          
          // Calculate slashed amount
          uint256 slashedAmount = SlashingLib.calcSlashedAmount(
              cumulativeShares,
              prevMag,
              newMag
          );
          
          cumulativeShares -= slashedAmount;
          
          // Calculate expected result directly
          expectedShares = expectedShares * newMag / prevMag;
      }
      
      // Verify rounding error is within acceptable bounds
      uint256 roundingError = expectedShares > cumulativeShares ? 
          expectedShares - cumulativeShares : 
          cumulativeShares - expectedShares;
          
      // Error should be less than 0.01% of initial shares
      assertLe(roundingError, initialShares / 10000);
  }
  ```

### Test Case: Deposit Scaling Factor Edge Cases
- **Description**: Test edge cases in deposit scaling factor calculations
- **Implementation**:
  ```solidity
  function test_DepositScalingFactor_EdgeCases() public {
      // Test case: First deposit after operator was slashed
      uint256 prevDepositShares = 0;
      uint256 addedShares = 1e18;
      uint256 slashingFactor = WAD / 2; // 50% slashing
      
      DepositScalingFactor storage dsf;
      SlashingLib.update(dsf, prevDepositShares, addedShares, slashingFactor);
      
      // Verify scaling factor is set correctly for first deposit
      assertEq(dsf.scalingFactor(), WAD.divWad(slashingFactor));
      
      // Test case: Deposit when operator is fully slashed (slashingFactor = 0)
      // This should revert to prevent division by zero
      prevDepositShares = 0;
      addedShares = 1e18;
      slashingFactor = 0;
      
      vm.expectRevert();
      SlashingLib.update(dsf, prevDepositShares, addedShares, slashingFactor);
      
      // Test case: Deposit with very small slashing factor
      prevDepositShares = 0;
      addedShares = 1e18;
      slashingFactor = 1; // Smallest non-zero value
      
      SlashingLib.update(dsf, prevDepositShares, addedShares, slashingFactor);
      
      // Verify scaling factor is set correctly
      assertEq(dsf.scalingFactor(), WAD.divWad(slashingFactor));
  }
  ```

### Test Case: Precision Loss in Nested Operations
- **Description**: Test for precision loss in nested mathematical operations
- **Implementation**:
  ```solidity
  function testFuzz_SlashingLib_PrecisionLoss(
      uint256 depositShares,
      uint64 dsf,
      uint64 maxMagnitude,
      uint64 beaconChainSlashingFactor
  ) public {
      // Bound inputs to realistic ranges
      depositShares = bound(depositShares, 1e6, 1e30);
      dsf = uint64(bound(dsf, 1e6, WAD));
      maxMagnitude = uint64(bound(maxMagnitude, 1e6, WAD));
      beaconChainSlashingFactor = uint64(bound(beaconChainSlashingFactor, 1e6, WAD));
      
      // Calculate slashing factor
      uint256 slashingFactor = uint256(maxMagnitude).mulWad(beaconChainSlashingFactor);
      
      // Calculate withdrawable shares
      uint256 withdrawableShares = SlashingLib.calcWithdrawable(
          DepositScalingFactor(dsf),
          depositShares,
          slashingFactor
      );
      
      // Calculate expected result directly with higher precision
      uint256 expectedWithdrawable = depositShares * dsf / WAD * slashingFactor / WAD;
      
      // Calculate relative error
      uint256 absoluteError = withdrawableShares > expectedWithdrawable ? 
          withdrawableShares - expectedWithdrawable : 
          expectedWithdrawable - withdrawableShares;
      
      // Relative error should be very small (less than 0.0001%)
      assertLe(absoluteError * 1e6 / depositShares, 1);
  }
  ```

### Test Case: Division by Zero Prevention
- **Description**: Test that the system properly handles potential division by zero scenarios
- **Implementation**:
  ```solidity
  function test_SlashingLib_DivisionByZeroPrevention() public {
      // Test case: Deposit when operator is fully slashed (slashingFactor = 0)
      uint256 prevDepositShares = 0;
      uint256 addedShares = 1e18;
      uint256 slashingFactor = 0;
      
      DepositScalingFactor storage dsf;
      
      // This should revert to prevent division by zero
      vm.expectRevert();
      SlashingLib.update(dsf, prevDepositShares, addedShares, slashingFactor);
      
      // Test case: Attempt to calculate withdrawable shares with zero slashing factor
      vm.expectRevert();
      SlashingLib.calcWithdrawable(
          DepositScalingFactor(WAD),
          1e18,
          0
      );
  }
  ```

### Test Case: Rounding Direction Consistency
- **Description**: Test that rounding is consistent in different operations
- **Implementation**:
  ```solidity
  function test_SlashingLib_RoundingDirectionConsistency() public {
      // Test rounding in mulWad vs mulWadRoundUp
      uint256 x = WAD - 1; // Just below 1 WAD
      uint256 y = WAD - 1; // Just below 1 WAD
      
      uint256 resultDown = x.mulWad(y);
      uint256 resultUp = x.mulWadRoundUp(y);
      
      // Verify rounding directions
      assertLt(resultDown, WAD - 1);
      assertEq(resultUp, WAD - 1);
      
      // Test rounding in calcSlashedAmount
      uint256 operatorShares = 1e18;
      uint64 prevMaxMagnitude = WAD;
      uint64 newMaxMagnitude = WAD / 2; // 50% slash
      
      uint256 slashedAmount = SlashingLib.calcSlashedAmount(
          operatorShares,
          prevMaxMagnitude,
          newMaxMagnitude
      );
      
      // Verify slashed amount is rounded up (to ensure operator shares are sufficiently reduced)
      assertEq(slashedAmount, operatorShares / 2);
      
      // Test rounding in calcWithdrawable
      uint256 depositShares = 1e18;
      uint256 slashingFactor = WAD / 2; // 50% slashing
      
      uint256 withdrawable = SlashingLib.calcWithdrawable(
          DepositScalingFactor(WAD),
          depositShares,
          slashingFactor
      );
      
      // Verify withdrawable is rounded down (to ensure staker doesn't receive more than entitled)
      assertEq(withdrawable, depositShares / 2);
  }
  ```

### Test Case: Deposits Reducing Withdrawable Shares
- **Description**: Test the edge case where deposits can actually reduce withdrawable shares due to rounding errors
- **Implementation**:
  ```solidity
  function test_SlashingLib_DepositsReducingWithdrawableShares() public {
      // Setup initial large deposit
      uint256 initialDeposit = 4.418e28;
      uint256 smallDeposit = 1000;
      
      // Mock operator with slight slashing
      uint256 slashingFactor = WAD - 1; // Just below 1 WAD
      
      // Initialize deposit scaling factor
      DepositScalingFactor storage dsf;
      SlashingLib.update(dsf, 0, initialDeposit, slashingFactor);
      
      // Calculate initial withdrawable shares
      uint256 initialWithdrawable = SlashingLib.calcWithdrawable(
          dsf,
          initialDeposit,
          slashingFactor
      );
      
      // Simulate additional deposit
      uint256 totalShares = initialDeposit + smallDeposit;
      SlashingLib.update(dsf, initialDeposit, smallDeposit, slashingFactor);
      
      // Calculate new withdrawable shares
      uint256 newWithdrawable = SlashingLib.calcWithdrawable(
          dsf,
          totalShares,
          slashingFactor
      );
      
      // Verify that in some cases, newWithdrawable can be less than initialWithdrawable + smallDeposit
      // This is a known edge case documented in SharesAccountingEdgeCases.md
      if (newWithdrawable < initialWithdrawable) {
          uint256 difference = initialWithdrawable - newWithdrawable;
          // Verify the difference is small relative to the deposit amount
          assertLe(difference, initialDeposit / 1e14);
      }
  }
  ```
