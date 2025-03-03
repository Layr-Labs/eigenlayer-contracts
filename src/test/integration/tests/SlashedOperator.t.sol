// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_SlashedOperator is IntegrationCheckUtils {
    AVS avs;
    User staker;
    User operator;

    OperatorSet operatorSet;

    AllocateParams allocateParams;
    SlashingParams slashParams;

    IStrategy[] strategies;
    IERC20[] tokens;

    uint[] initTokenBalances;
    uint[] initDepositShares;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_LST);

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();
        
        operatorSet = avs.createOperatorSet(strategies);
        tokens = _getUnderlyingTokens(strategies);

        // 1) Register operator for operator set.
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);

        // 2) Operator allocates to operator set.
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_Base_IncrAlloc_State(operator, allocateParams);

        // 3) Roll forward to complete allocation.
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 4) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
    }

    function testFuzz_register_allocate_fullSlash_deposit_delegate(uint24 r) public rand(r) {
        // 5) Staker deposits into strategies.
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 6) Staker delegates to operator who is fully slashed, should fail.
        vm.expectRevert();
        staker.delegateTo(operator);
    }

    function testFuzz_register_allocate_fullSlash_delegate_deposit(uint24 r) public rand(r) {
        // 5) Staker delegates to operator who is fully slashed, should fail.
        staker.delegateTo(operator);
        // NOTE: Leads to division by zero?
        // check_Delegation_State(staker, operator, strategies, new uint256[](strategies.length)); // Initial shares are zero
        
        // 6) Staker deposits into strategies, should fail.
        vm.expectRevert();
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
    }

    function testFuzz_register_allocate_fullSlash_delegate_redelegate_deposit(uint24 r) public rand(r) {
        // 5) Staker delegates to operator who is fully slashed, should fail.
        staker.delegateTo(operator);

        User newOperator = _newRandomOperator_NoAssets();
        newOperator.registerForOperatorSet(operatorSet);

        // 6) Staker redelegates to new operator.
        Withdrawal[] memory withdrawals = staker.redelegate(newOperator);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, newOperator, withdrawals, withdrawalRoots, strategies, new uint256[](strategies.length));

        _rollBlocksForCompleteWithdrawals(withdrawals);

        // 7) Staker deposits into strategies.
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);
    }
}