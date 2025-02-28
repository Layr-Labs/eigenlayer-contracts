// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

/// @notice Tests where we slash native eth on the Beacon Chain and by an OperatorSet 
contract Integration_DualSlashing_Base is IntegrationCheckUtils {
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

    function _init() internal virtual override {
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
}

contract Integration_DualSlashing_Base_BeaconChainFirst is Integration_DualSlashing_Base {

    function testFuzz_bcSlash_checkpoint_avsSlash(uint24 _random) public rand(_random) {
        // 6. Slash staker on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        // 7. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, validators, slashedAmountGwei);

        // 8. Slash operator by AVS
        SlashingParams memory slashingParams;
        {
            slashingParams = _genSlashing_Rand(operator, operatorSet);
            avs.slashOperator(slashingParams);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterBCSlash_AVSSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should be slashed");
        }
    }
}

contract Integration_DualSlashing_Base_AVSFirst is Integration_DualSlashing_Base {
    using ArrayLib for *;

    SlashingParams slashingParams;

    function _init() internal virtual override {
        super._init();

        // 6. Slash operator by AVS
        slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashingParams);
        check_Base_Slashing_State(operator, allocateParams, slashingParams);
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
        assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should be slashed");
    }

    /// @dev Validates behavior of "restaking", ie. that the funds can be slashed twice
    function testFuzz_avsSlash_bcSlash_checkpoint(uint24 _random) public rand(_random){ 
        // 7. Slash Staker on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();
        
        // 8. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_AfterAVSSlash_BCSlash(staker, validators, initDepositShares[0], slashedAmountGwei, allocateParams, slashingParams);
    }

    /// @notice Because the validator is proven prior to the BC slash, the system applies the new balance 
    ///         to the BC and AVS slash combined
    function testFuzz_avsSlash_verifyValidator_bcSlash_checkpoint(uint24 _random) public rand(_random) {
        // 7. Verify Validator
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        uint beaconSharesAdded = uint256(addedBeaconBalanceGwei * GWEI_TO_WEI);
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);
        assert_Snap_Added_Staker_WithdrawableShares_AtLeast(staker, BEACONCHAIN_ETH_STRAT.toArray(), beaconSharesAdded.toArrayU256(), "staker withdrawable shares should increase by the added beacon balance");

        // 8. Slash first validators on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        // 9. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_AfterAVSSlash_ValidatorProven_BCSlash(staker, validators, initDepositShares[0], beaconSharesAdded, allocateParams, slashingParams);
    }

    /// @dev Same as above, but validator is proven after BC slash (this ordering doesn't matter to EL)
    function testFuzz_avsSlash_bcSlash_verifyValidator_checkpoint(uint24 _random) public rand(_random) {
        // 7. Slash Staker on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        // 8. Verify Validator
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        uint beaconSharesAdded = uint256(addedBeaconBalanceGwei * GWEI_TO_WEI);
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);
        assert_Snap_Added_Staker_WithdrawableShares_AtLeast(staker, BEACONCHAIN_ETH_STRAT.toArray(), beaconSharesAdded.toArrayU256(), "staker withdrawable shares should increase by the added beacon balance");

        // 9. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_AfterAVSSlash_ValidatorProven_BCSlash(staker, validators, initDepositShares[0], beaconSharesAdded, allocateParams, slashingParams);
    }

    /// @notice The validator proven should not be affected by the BC or AVS slashes
    function testFuzz_avsSlash_bcSlash_checkpoint_verifyValidator(uint24 _rand) public rand(_rand) {
        // 7. Slash Staker on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();
        
        // 8. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_AfterAVSSlash_BCSlash(staker, validators, initDepositShares[0], slashedAmountGwei, allocateParams, slashingParams);

        // 9. Verify Validator
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        uint beaconSharesAdded = uint256(addedBeaconBalanceGwei * GWEI_TO_WEI);
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);
        assert_Snap_Added_Staker_WithdrawableShares_AtLeast(staker, BEACONCHAIN_ETH_STRAT.toArray(), beaconSharesAdded.toArrayU256(), "staker withdrawable shares should increase by the added beacon balance");
    }

    /// @notice The balance increase results in the pods not processing the beacon slash as a slash, given
    ///         that the checkpoint had a positive delta
    function testFuzz_avsSlash_bcSlash_balanceIncrease_checkpoint(uint24 _rand) public rand(_rand) {
        // 7. Slash Staker on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        // 8. Send 32 ETH to pod, some random amount of ETH, greater than the amount slashed
        uint ethToDeposit = 32 ether;
        cheats.deal(address(staker), ethToDeposit);
        cheats.prank(address(staker));
        address(staker.pod()).call{value: ethToDeposit}("");
        uint64 beaconSharesAddedGwei = uint64(ethToDeposit / GWEI_TO_WEI);

        // 9. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei + beaconSharesAddedGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_AfterAVSSlash_ETHDeposit_BCSlash(staker, validators, slashedAmountGwei, beaconSharesAddedGwei);
    }

    /// @notice The balance increase occurs after the slashings are processed, so it should be unaffected by the slashings
    function testFuzz_avsSlash_bcSlash_checkpoint_balanceIncrease(uint24 _rand) public rand(_rand) {
        // 7. Slash Staker on BC
        uint64 slashedAmountGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();
        
        // 8. Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedAmountGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_AfterAVSSlash_BCSlash(staker, validators, initDepositShares[0], slashedAmountGwei, allocateParams, slashingParams);

        // 9. Send 32 ETH to pod, some random amount of ETH, greater than the amount slashed
        uint ethToDeposit = 32 ether;
        cheats.deal(address(staker), ethToDeposit);
        cheats.prank(address(staker));
        address(staker.pod()).call{value: ethToDeposit}("");
        uint64 beaconSharesAddedGwei = uint64(ethToDeposit / GWEI_TO_WEI);

        // 10. Checkpoint. This should immediately complete as there are no more active validators
        beaconChain.advanceEpoch_NoRewards();
        staker.startCheckpoint();
        check_StartCheckpoint_NoValidators_State(staker, beaconSharesAddedGwei);
    }
}