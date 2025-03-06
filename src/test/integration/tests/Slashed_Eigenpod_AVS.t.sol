// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_SlashedEigenpod_AVS_Base is IntegrationCheckUtils {
    using ArrayLib for *;
    using SlashingLib for *;
    using Math for uint;

    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;
    SlashingParams slashParams;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        cheats.assume(initTokenBalances[0] >= 64 ether);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);

        // 4. Register for operator set
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        // 5. Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }
}

contract Integration_SlashedEigenpod_AVS_Checkpoint is Integration_SlashedEigenpod_AVS_Base {
    function _init() internal override {
        super._init();

        // 6. Slash
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        beaconChain.advanceEpoch_NoRewards();
    }

    /// @dev Asserts that the DSF isn't updated after a slash & checkpoint with 0 balance
    function testFuzz_deposit_delegate_allocate_slash_checkpointZeroBalance(uint24 _rand) public rand(_rand) {
        // 7. Start & complete checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_ZeroBalanceDelta_State(staker);
    }
}

contract Integration_SlashedEigenpod_AVS_Withdraw is Integration_SlashedEigenpod_AVS_Base {
    using Math for uint;
    using SlashingLib for uint;

    uint[] withdrawableSharesAfterSlash;

    function _init() internal override {
        super._init();

        // Slash or queue a withdrawal in a random order
        if (_randBool()) {
            // Slash -> Queue Withdrawal
            // 7. Slash
            slashParams = _genSlashing_Half(operator, operatorSet);
            avs.slashOperator(slashParams);
            check_Base_Slashing_State(operator, allocateParams, slashParams);

            // 8. Queue Withdrawal for all shares.
            Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
            bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
            withdrawableSharesAfterSlash = _calcWithdrawable(staker, strategies, initDepositShares);
            check_QueuedWithdrawal_State(
                staker, operator, strategies, initDepositShares, withdrawableSharesAfterSlash, withdrawals, withdrawalRoots
            );
        } else {
            // Queue Withdrawal -> Slash
            // 7. Queue Withdrawal for all shares
            Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
            bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
            withdrawableSharesAfterSlash = _calcWithdrawable(staker, strategies, initDepositShares);
            check_QueuedWithdrawal_State(
                staker, operator, strategies, initDepositShares, withdrawableSharesAfterSlash, withdrawals, withdrawalRoots
            );

            // 8. Slash
            slashParams = _genSlashing_Half(operator, operatorSet);
            avs.slashOperator(slashParams);
            check_Base_Slashing_State(operator, allocateParams, slashParams);
        }
        withdrawableSharesAfterSlash = _calcWithdrawable(staker, strategies, initDepositShares);

        beaconChain.advanceEpoch_NoRewards();
    }

    /// @dev Asserts that the DSF isn't updated after a slash/queue and a checkpoint with 0 balance.
    /// @dev The staker should subsequently not be able to inflate their withdrawable shares
    function testFuzz_deposit_delegate_allocate_slashAndQueue_checkpoint_redeposit(uint24 _rand) public rand(_rand) {
        // 9.  Start & complete checkpoint.
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_ZeroBalanceDelta_State(staker);

        // 10. Redeposit
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);
    }

    /// @dev Asserts that the staker cannot inflate withdrawable shares after redepositing
    function testFuzz_deposit_delegate_allocate_slashAndQueue_completeAsTokens_redeposit(uint24 _rand) public rand(_rand) {
        Withdrawal[] memory withdrawals = _getQueuedWithdrawals(staker);
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // 9. Complete withdrawal as tokens
        for (uint i = 0; i < withdrawals.length; ++i) {
            IERC20[] memory tokens = _getUnderlyingTokens(withdrawals[i].strategies);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawableSharesAfterSlash);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], strategies, withdrawals[i].scaledShares, tokens, expectedTokens
            );
        }

        // 10. Redeposit
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);
    }

    /// @dev Asserts that the staker cannot inflate withdrawable shares after checkpointing & completing as shares
    function testFuzz_deposit_delegate_allocate_slashAndQueue_checkPoint_completeAsShares(uint24 _rand) public rand(_rand) {
        Withdrawal[] memory withdrawals = _getQueuedWithdrawals(staker);
        _rollBlocksForCompleteWithdrawals(withdrawals);
        uint[] memory withdrawableShares = _calcWithdrawable(staker, strategies, initDepositShares);

        // 9.  Start & complete checkpoint, since the next step does not.
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_ZeroBalanceDelta_State(staker);

        // 10. Complete withdrawal as shares. Deposit scaling factor is doubled because operator was slashed by half.
        staker.completeWithdrawalAsShares(withdrawals[0]);
        check_Withdrawal_AsShares_State(staker, operator, withdrawals[0], strategies, withdrawableShares);
    }
}
