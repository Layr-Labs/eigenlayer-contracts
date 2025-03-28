// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

/// @notice Testing the rounding behavior when the DSF is high and there are multiple deposits
contract Integration_HighDSF_Multiple_Deposits is IntegrationCheckUtils {
    using ArrayLib for *;

    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;
    SlashingParams slashParams;

    User staker;
    IStrategy[] strategies;
    IERC20[] tokens; // underlying token for each strategy
    uint[] initTokenBalances;
    uint[] initDepositShares;

    /**
     * Shared setup:
     * 1. create a new staker, operator, and avs
     * 2. create an operator set and register an operator, allocate all magnitude to the operator set
     * 3. slash operator to 1 magnitude remaining
     * 4. delegate to operator
     */
    function _init() internal override {
        // 1. create a new staker, operator, and avs
        _configAssetTypes(HOLDS_LST);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        // 2. Create an operator set and register an operator, allocate all magnitude to the operator set
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable_NoDelegatedStake(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 3. slash operator to 1 magnitude remaining
        slashParams = _genSlashing_Custom(operator, operatorSet, WAD - 1);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // 4. delegate to operator
        staker.delegateTo(operator);

        uint slashingFactor = staker.getSlashingFactor(strategies[0]);
        assertEq(slashingFactor, 1, "slashing factor should be 1");
    }

    /// @notice Test setup with a staker with slashingFactor of 1 (maxMagnitude = 1)
    /// with repeat deposits to increase the DSF. Limiting number of fuzzed runs to speed up tests since this
    /// for loops several times.
    /// forge-config: default.fuzz.runs = 10
    function test_multiple_deposits(uint24 _r) public rand(_r) {
        // deposit initial assets into strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // Repeat the deposit 50 times
        // Gas intensive so we pause gas metering for this loop
        cheats.pauseGasMetering();
        for (uint i = 0; i < 50; i++) {
            _dealAmounts(staker, strategies, initTokenBalances);
            staker.depositIntoEigenlayer(strategies, initTokenBalances);
            initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
            check_Deposit_State(staker, strategies, initDepositShares);
        }
        cheats.resumeGasMetering();

        // Check that the DSF is still bounded without overflow
        for (uint i = 0; i < strategies.length; i++) {
            assertGe(delegationManager.depositScalingFactor(address(staker), strategies[i]), WAD, "DSF should be >= WAD");
            // theoretical upper bound on DSF is 1e74
            assertLt(delegationManager.depositScalingFactor(address(staker), strategies[i]), 1e74, "DSF should be < 1e74");
        }
    }
}
