// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_BeaconChain_AVS_Ordering is IntegrationCheckUtils {
    using ArrayLib for *;

    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;

    User staker;
    uint64 beaconBalanceGwei;
    uint40[] validators;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);

        // Create staker, operator, and avs
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        // 1. Deposit into strategies
        (validators, beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(validators);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Delegate staker to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create an operator set containing the strategies held by the staker
        operatorSet = avs.createOperatorSet(strategies);

        // 4. Register operator to AVS
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);
        
        // 5. Allocate to the operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);    

        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }

    function testFuzz_avsSlash_bcSlash_checkpoint(uint24 _random) public rand(_random){ 
        // 6. Slash Operator by AVS
        SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            console.log("Strategies to slash: %s", strategiesToSlash.length);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should be slashed");
        }

        // 7. Slash Staker on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators);
        beaconChain.advanceEpoch_NoRewards();
        
        // 8. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckPoint_AfterAVSAndBCSlash(staker, validators, initDepositShares[0], slashedAmountGwei, allocateParams, slashingParams);
    }

    function testFuzz_bcSlash_checkpoint_avsSlash(uint24 _random) public rand(_random) {
        // 6. Slash staker on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators);
        beaconChain.advanceEpoch_NoRewards();

        // 7. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, validators, slashedAmountGwei);

        // 8. Slash operator by AVS
        SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            console.log("Strategies to slash: %s", strategiesToSlash.length);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterBCAndAVSSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should be slashed");
        }
    }

    function testFuzz_avsSlash_verifyValidator_bcSlash_checkpoint(uint24 _random) public rand(_random) {
        // 6. Slash operator by AVS
        SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            console.log("Strategies to slash: %s", strategiesToSlash.length);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should be slashed");
        }

        // 7. Verify Validator
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);
        assert_Snap_Added_Staker_WithdrawableSharesAtLeast(staker, BEACONCHAIN_ETH_STRAT.toArray(), uint256(addedBeaconBalanceGwei * GWEI_TO_WEI).toArrayU256(), "staker withdrawable shares should increase by the added beacon balance");

        // 8. Slash first validators on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators);
        beaconChain.advanceEpoch_NoRewards();

        // 9. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithAVSAndBCSlashing_HandleRoundDown_State(staker, validators, slashedAmountGwei);
    }

}