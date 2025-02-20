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
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }
        uint[] memory sharesGotten = _getWithdrawableShares(staker, strategies);
        console.log("Shares after avs slash: ", sharesGotten[0]);

        // 7. Slash Staker on BC
        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators);
        beaconChain.advanceEpoch_NoRewards();
        
        console.log("slashedBalanceGwei: ", slashedBalanceGwei);

        // 8. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedBalanceGwei);
        staker.completeCheckpoint();
        uint64[] memory maxMagnitudes = _getMaxMagnitudes(operator, strategies);
        console.log("maxMagnitudes: ", maxMagnitudes[0]);
        uint64 bcsf = _getBeaconChainSlashingFactor(staker);
        console.log("bcsf: ", bcsf);
        sharesGotten = _getWithdrawableShares(staker, strategies);
        console.log("Shares after bc slash: ", sharesGotten[0]);
        check_CompleteCheckPoint_AVSSLash_BCSLash_HandleRoundDown_State(staker, validators, initDepositShares[0], slashedBalanceGwei);


        // sharesGotten = _getWithdrawableShares(staker, strategies);
        // console.log("Shares after bc slash: ", sharesGotten[0]);
        // console.log("Test test test");

        // Assert state
        // assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
        // assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(staker, BEACONCHAIN_ETH_STRAT, slashedBalanceGwei * GWEI_TO_WEI, "should have decreased withdrawable shares by slashed amount");
    }

}