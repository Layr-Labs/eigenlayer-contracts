// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_VerifyWC_StartCP_CompleteCP is IntegrationCheckUtils {
    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        _configUserTypes(DEFAULT);
    }

    function test_GasMetering() public rand(0) {
        (User staker,,) = _newRandomStaker();

        // Deal user 20 full stakers worth of ETH
        emit log_named_string("Dealing 20 * 32 ETH to", staker.NAME());
        cheats.deal(address(staker), 32 ether * 20);
        cheats.pauseGasMetering();

        (uint40[] memory validators,) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        EigenPod pod = staker.pod();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        cheats.startPrank(address(staker));
        cheats.resumeGasMetering();

        uint startGas = gasleft();
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        uint endGas = gasleft();
        cheats.pauseGasMetering();
        uint totalGas = startGas - endGas;
        emit log_named_uint("== num validators", validators.length);
        emit log_named_uint("== verifyWC gas", totalGas);

        // check_VerifyWC_State(staker, validators, beaconBalanceWei);

        beaconChain.advanceEpoch();
        // check pod balances have increased

        staker.startCheckpoint();
        // check_StartCheckpoint_State(staker);

        CheckpointProofs memory cpProofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());

        cheats.resumeGasMetering();
        startGas = gasleft();
        pod.verifyCheckpointProofs({balanceContainerProof: cpProofs.balanceContainerProof, proofs: cpProofs.balanceProofs});
        endGas = gasleft();
        cheats.pauseGasMetering();
        totalGas = startGas - endGas;
        emit log_named_uint("== checkpoint gas", totalGas);

        // check_CompleteCheckpoint_State(staker);
        // revert();
    }

    /**
     *
     *                    VERIFY -> START -> COMPLETE CHECKPOINT
     *                             (TIMING VARIANTS)
     *
     */

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. Verify validators' withdrawal credentials again
    /// => This should fail
    function test_VerifyWC_VerifyWC_Fails(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        cheats.expectRevert(IEigenPodErrors.CredentialsAlreadyVerified.selector);
        staker.verifyWithdrawalCredentials(validators);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. start a checkpoint again
    /// => This should fail
    function test_VerifyWC_StartCP_StartCP_Fails(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        cheats.expectRevert(IEigenPodErrors.CheckpointAlreadyActive.selector);
        staker.startCheckpoint();
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// 4. start a checkpoint without advancing a block
    /// => this should fail
    function test_VerifyWC_StartCP_CompleteCP_StartCP_Fails(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);

        cheats.expectRevert(IEigenPodErrors.CannotCheckpointTwiceInSingleBlock.selector);
        staker.startCheckpoint();
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- move forward 1 or more epochs
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_Advance_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Advance epoch without generating consensus rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- move forward 1 or more epochs
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_StartCP_Advance_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        // Advance epoch without generating consensus rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    /**
     *
     *                    VERIFY -> START -> COMPLETE CHECKPOINT
     *                          (EXIT TO POD VARIANTS)
     *
     */

    /// -- Fully exit validators before verifying withdrawal credentials
    /// 1. Verify validators' withdrawal credentials
    /// => This should fail
    function test_ExitValidators_VerifyWC_Fails(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators,) = staker.startValidators();
        staker.exitValidators(validators);
        beaconChain.advanceEpoch_NoRewards();

        cheats.expectRevert(IEigenPodErrors.ValidatorIsExitingBeaconChain.selector);
        staker.verifyWithdrawalCredentials(validators);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- fully exit validators to pod
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_ExitValidators_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Fully exit one or more validators and advance epoch without generating rewards
        uint40[] memory subset = _choose(validators);
        uint64 exitedBalanceGwei = staker.exitValidators(subset);
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, exitedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithExits_State(staker, subset, exitedBalanceGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- fully exit validators to pod
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    /// -- move forward an epoch
    /// 4. start a checkpoint
    /// 5. complete a checkpoint
    /// => exited balance should be reflected in 4 and 5
    function test_VerifyWC_StartCP_ExitValidators_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        // Fully exit one or more validators and advance epoch without generating rewards
        uint40[] memory subset = _choose(validators);
        uint64 exitedBalanceGwei = staker.exitValidators(subset);
        beaconChain.advanceEpoch_NoRewards();

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, exitedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithExits_State(staker, subset, exitedBalanceGwei);
    }

    /**
     *
     *                    VERIFY -> START -> COMPLETE CHECKPOINT
     *                          (SLASH TO POD VARIANTS)
     *
     */

    /// -- get slashed on beacon chain
    /// 1. Try to verify validators' withdrawal credentials
    /// => this should fail
    function test_SlashToPod_VerifyWC_Fails(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators,) = staker.startValidators();
        beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        // Advance epoch, withdrawing slashed validators to pod
        beaconChain.advanceEpoch_NoRewards();

        cheats.expectRevert(IEigenPodErrors.ValidatorIsExitingBeaconChain.selector);
        staker.verifyWithdrawalCredentials(validators);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- get slashed on beacon chain; exit to pod
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares should decrease by slashed amount
    function test_VerifyWC_SlashToPod_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, validators, slashedBalanceGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- get slashed on beacon chain; exit to pod
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    /// -- move forward an epoch
    /// 4. start a checkpoint
    /// 5. complete a checkpoint
    /// => slashed balance should be reflected in 4 and 5
    function test_VerifyWC_StartCP_SlashToPod_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, validators, slashedBalanceGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. slash validators
    /// 3. start a checkpoint
    /// 4. verify withdrawal credentials for another validator while checkpoint in progress
    /// 5. complete a checkpoint
    /// => Increase in shares between 1 and 4 should reflect the new validator, less the slashed amount
    function test_VerifyWC_Slash_StartCP_VerifyWC_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Slash validators
        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        // Start a checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedBalanceGwei);

        // Start a new validator & verify withdrawal credentials
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(staker, validators, slashedBalanceGwei);
    }

    /**
     *
     *                    VERIFY -> PROVE STALE BALANCE -> COMPLETE CHECKPOINT
     *
     */

    /// 1. Verify validators' withdrawal credentials
    /// -- get slashed on beacon chain; exit to pod
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares should decrease by slashed amount
    function test_VerifyWC_SlashToPod_VerifyStale_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyStaleBalance(validators[0]);
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, validators, slashedBalanceGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- get slashed on beacon chain; do not exit to pod
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares should decrease by slashed amount
    function test_VerifyWC_SlashToCL_VerifyStale_CompleteCP_SlashAgain(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Slash validators but do not process exits to pod
        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoWithdraw();

        staker.verifyStaleBalance(validators[0]);
        check_StartCheckpoint_WithPodBalance_State(staker, 0);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithCLSlashing_State(staker, slashedBalanceGwei);

        // Slash validators again but do not process exits to pod
        uint64 secondSlashedBalanceGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoWithdraw();

        staker.verifyStaleBalance(validators[0]);
        check_StartCheckpoint_WithPodBalance_State(staker, 0);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithCLSlashing_HandleRoundDown_State(staker, secondSlashedBalanceGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- get slashed on beacon chain; exit to pod
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    /// -- move forward an epoch
    /// 4. start a checkpoint
    /// 5. complete a checkpoint
    /// => slashed balance should be reflected in 4 and 5
    function test_VerifyWC_StartCP_SlashToPod_CompleteCP_VerifyStale(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);

        staker.verifyStaleBalance(validators[0]);
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, validators, slashedBalanceGwei);
    }

    /**
     *
     *                    VERIFY -> START -> COMPLETE CHECKPOINT
     *                          (EARN ON CL VARIANTS)
     *
     */

    /// -- earn rewards on beacon chain (not withdrawn to pod)
    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares increase by rewards earned on beacon chain
    function test_EarnOnBeacon_VerifyWC_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch and generate consensus rewards, but don't withdraw to pod
        beaconChain.advanceEpoch_NoWithdraw();
        uint64 beaconBalanceIncreaseGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_EarnOnBeacon_State(staker, beaconBalanceIncreaseGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- earn rewards on beacon chain (not withdrawn to pod)
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares increase by rewards earned on beacon chain
    function test_VerifyWC_EarnOnBeacon_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Advance epoch and generate consensus rewards, but don't withdraw to pod
        beaconChain.advanceEpoch_NoWithdraw();
        uint64 beaconBalanceIncreaseGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_EarnOnBeacon_State(staker, beaconBalanceIncreaseGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- earn rewards on beacon chain (not withdrawn to pod)
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_StartCP_EarnOnBeacon_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        // Advance epoch and generate consensus rewards, but don't withdraw to pod
        beaconChain.advanceEpoch_NoWithdraw();

        staker.completeCheckpoint();
        check_CompleteCheckpoint_EarnOnBeacon_State(staker, 0);
    }

    /**
     *
     *                    VERIFY -> START -> COMPLETE CHECKPOINT
     *                          (EARN TO POD VARIANTS)
     *
     */

    /// -- earn rewards on beacon chain (withdrawn to pod)
    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares increase by rewards withdrawn to pod
    function test_EarnToPod_VerifyWC_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch, generating consensus rewards and withdrawing anything over 32 ETH
        beaconChain.advanceEpoch();
        uint64 expectedWithdrawnGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, expectedWithdrawnGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithPodBalance_State(staker, expectedWithdrawnGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- earn rewards on beacon chain (withdrawn to pod)
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares increase by rewards withdrawn to pod
    function test_VerifyWC_EarnToPod_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Advance epoch, generating consensus rewards and withdrawing anything over 32 ETH
        beaconChain.advanceEpoch();
        uint64 expectedWithdrawnGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, expectedWithdrawnGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithPodBalance_State(staker, expectedWithdrawnGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- earn rewards on beacon chain (withdrawn to pod)
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_StartCP_EarnToPod_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, 0);

        // Advance epoch, generating consensus rewards and withdrawing anything over 32 ETH
        beaconChain.advanceEpoch();
        uint64 expectedWithdrawnGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();

        staker.completeCheckpoint();
        // `pod.balance == gweiSent + remainderSent
        assert_PodBalance_Eq(staker, (expectedWithdrawnGwei * GWEI_TO_WEI), "pod balance should equal expected");
        check_CompleteCheckpoint_WithPodBalance_State(staker, 0);
    }

    /**
     *
     *                    VERIFY -> START -> COMPLETE CHECKPOINT
     *                          (NATIVE ETH VARIANTS)
     *
     */

    /// -- Pod receives native ETH via fallback
    /// 1. start a checkpoint
    /// => checkpoint should auto-complete, awarding shares for ETH in pod
    function test_NativeETH_StartCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        // Send a random amount of ETH to staker's fallback
        (uint64 gweiSent,) = _sendRandomETH(address(staker.pod()));

        // Move forward an epoch so we generate a state root that can be queried in startCheckpoint
        beaconChain.advanceEpoch();

        // should behave identically to partial withdrawals captured by the "earn to pod" variants
        staker.startCheckpoint();
        check_StartCheckpoint_NoValidators_State(staker, gweiSent);
    }

    /// -- Pod receives native ETH via fallback
    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares should account for native ETH
    function test_NativeETH_VerifyWC_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        // Send a random amount of ETH to staker's fallback
        (uint64 gweiSent, uint remainderSent) = _sendRandomETH(address(staker.pod()));

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // should behave identically to partial withdrawals captured by the "earn to pod" variants
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, gweiSent);

        staker.completeCheckpoint();
        // check that `pod.balance == restakedExecutionLayerGwei + remainderSent
        assert_PodBalance_Eq(staker, (gweiSent * GWEI_TO_WEI) + remainderSent, "pod balance should equal expected");
        check_CompleteCheckpoint_WithPodBalance_State(staker, gweiSent);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- Pod receives native ETH via fallback
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares should account for native ETH
    function test_VerifyWC_NativeETH_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Send a random amount of ETH to staker's fallback
        (uint64 gweiSent, uint remainderSent) = _sendRandomETH(address(staker.pod()));

        // should behave identically to partial withdrawals captured by the "earn to pod" variants
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, gweiSent);

        staker.completeCheckpoint();
        // check that `pod.balance == restakedExecutionLayerGwei + remainderSent
        assert_PodBalance_Eq(staker, (gweiSent * GWEI_TO_WEI) + remainderSent, "pod balance should equal expected");
        check_CompleteCheckpoint_WithPodBalance_State(staker, gweiSent);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- Pod receives native ETH via fallback
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_StartCP_NativeETH_CompleteCP(uint24 _rand) public rand(_rand) {
        (User staker,,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // should behave identically to partial withdrawals captured by the "earn to pod" variants
        // ... if we didn't have any partial withdrawals!
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, 0);

        // Send a random amount of ETH to staker's fallback
        (uint64 gweiSent, uint remainderSent) = _sendRandomETH(address(staker.pod()));

        staker.completeCheckpoint();
        // `pod.balance == gweiSent + remainderSent
        assert_PodBalance_Eq(staker, (gweiSent * GWEI_TO_WEI) + remainderSent, "pod balance should equal expected");
        check_CompleteCheckpoint_WithPodBalance_State(staker, 0);
    }
}
