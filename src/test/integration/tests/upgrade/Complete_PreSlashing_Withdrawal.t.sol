// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_Complete_PreSlashing_Withdrawal is UpgradeTest {
    struct TestState {
        User staker;
        User operator;
        IStrategy[] strategies;
        uint256[] tokenBalances;
        uint256[] shares;
        uint256[] withdrawalShares;
        uint256[] expectedTokens;
        Withdrawal[] withdrawals;
        bool isPartial;
    }

    function _init_(uint24 _random, bool withOperator, bool withDelegation) internal returns (TestState memory state) {
        (state.staker, state.strategies, state.tokenBalances) = _newRandomStaker();
        state.shares = _calculateExpectedShares(state.strategies, state.tokenBalances);
        
        if (withOperator) {
            (state.operator, ,) = _newRandomOperator();
        } else {
            state.operator = User(payable(0));
        }

        if (withDelegation) {
            state.staker.delegateTo(state.operator);
        }

        state.staker.depositIntoEigenlayer(state.strategies, state.tokenBalances);
        
        // Setup withdrawal shares (full or partial)
        state.isPartial = _randBool();
        state.withdrawalShares = new uint256[](state.shares.length);
        for (uint256 i = 0; i < state.shares.length; i++) {
            state.withdrawalShares[i] = state.isPartial ? state.shares[i] / 2 : state.shares[i];
        }
        
        return state;
    }

    function testFuzz_deposit_queue_upgrade_completeAsShares(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, false, false);
        
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _upgradeEigenLayerContracts();
    
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        for (uint i = 0; i < state.withdrawals.length; i++) {
            state.staker.completeWithdrawalAsShares(state.withdrawals[i]);
            check_Withdrawal_AsShares_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares);
        }
    }

    function testFuzz_deposit_queue_upgrade_completeAsTokens(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, false, false);
        state.expectedTokens = _calculateExpectedTokens(state.strategies, state.withdrawalShares);
        
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _upgradeEigenLayerContracts();
    
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        for (uint i = 0; i < state.withdrawals.length; i++) {
            IERC20[] memory tokens = state.staker.completeWithdrawalAsTokens(state.withdrawals[i]);
            check_Withdrawal_AsTokens_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares, tokens, state.expectedTokens);
        }
    }

    function testFuzz_delegate_deposit_queue_upgrade_completeAsShares(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, true, true);
        
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _upgradeEigenLayerContracts();

        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        for (uint i = 0; i < state.withdrawals.length; i++) {
            state.staker.completeWithdrawalAsShares(state.withdrawals[i]);
            check_Withdrawal_AsShares_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares);
        }
    }

    function testFuzz_delegate_deposit_queue_upgrade_completeAsTokens(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, true, true);
        state.expectedTokens = _calculateExpectedTokens(state.strategies, state.withdrawalShares);
        
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _upgradeEigenLayerContracts();

        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        for (uint i = 0; i < state.withdrawals.length; i++) {
            IERC20[] memory tokens = state.staker.completeWithdrawalAsTokens(state.withdrawals[i]);
            check_Withdrawal_AsTokens_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares, tokens, state.expectedTokens);
        }
    }

    function testFuzz_upgrade_delegate_queuePartial_completeAsShares(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, true, false);
        
        _upgradeEigenLayerContracts();
        state.staker.delegateTo(state.operator);
        
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        
        for (uint i = 0; i < state.withdrawals.length; i++) {
            state.staker.completeWithdrawalAsShares(state.withdrawals[i]);
            check_Withdrawal_AsShares_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares);
        }
    }

    function testFuzz_upgrade_delegate_queuePartial_completeAsTokens(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, true, false);
        state.expectedTokens = _calculateExpectedTokens(state.strategies, state.withdrawalShares);
        
        _upgradeEigenLayerContracts();
        state.staker.delegateTo(state.operator);
        
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        
        for (uint i = 0; i < state.withdrawals.length; i++) {
            IERC20[] memory tokens = state.staker.completeWithdrawalAsTokens(state.withdrawals[i]);
            check_Withdrawal_AsTokens_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares, tokens, state.expectedTokens);
        }
    }

    function testFuzz_delegate_deposit_queue_completeBeforeUpgrade_asShares(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, true, true);
        
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        
        // We must roll forward by the delay twice since pre-upgrade delay is half as long as post-upgrade delay.
        rollForward(delegationManager.minWithdrawalDelayBlocks() + 1);
        _upgradeEigenLayerContracts();

        for (uint256 i = 0; i < state.withdrawals.length; i++) {
            state.staker.completeWithdrawalAsShares(state.withdrawals[i]);
            check_Withdrawal_AsShares_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares);
        }
    }

    function testFuzz_delegate_deposit_queue_completeBeforeUpgrade_asTokens(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, true, true);
        state.expectedTokens = _calculateExpectedTokens(state.strategies, state.withdrawalShares);
        
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        
        // We must roll forward by the delay twice since pre-upgrade delay is half as long as post-upgrade delay.
        rollForward(delegationManager.minWithdrawalDelayBlocks() + 1);
        _upgradeEigenLayerContracts();

        for (uint i = 0; i < state.withdrawals.length; i++) {
            IERC20[] memory tokens = state.staker.completeWithdrawalAsTokens(state.withdrawals[i]);
            check_Withdrawal_AsTokens_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares, tokens, state.expectedTokens);
        }
    }

    function testFuzz_upgrade_allocate_correctSlashableStake(uint24 _random) public rand(_random) {
        TestState memory state = _init_(_random, true, false);
        state.staker.delegateTo(state.operator);

        _upgradeEigenLayerContracts();
        (AVS avs,) = _newRandomAVS();

        state.operator.setAllocationDelay(1);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});

        OperatorSet memory operatorSet = avs.createOperatorSet(state.strategies);
        state.operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(state.operator, operatorSet, allStrats);

        AllocateParams memory allocateParams = AllocateParams({
            operatorSet: operatorSet,
            strategies: state.strategies,
            newMagnitudes: _randMagnitudes({sum: 1 ether, len: state.strategies.length})
        });
        state.operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(state.operator, allocateParams);
    }
}