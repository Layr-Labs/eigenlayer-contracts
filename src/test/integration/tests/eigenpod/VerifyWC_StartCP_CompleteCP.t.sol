// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_VerifyWC_StartCP_CompleteCP is IntegrationCheckUtils {

    modifier r(uint24 _rand) {
        _configRand({
            _randomSeed: _rand,
            _assetTypes: HOLDS_ETH,
            _userTypes: DEFAULT
        });

        _;
    }

    function test_GasMetering() public r(24) {
        (User staker, ,) = _newRandomStaker();
        cheats.pauseGasMetering();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();
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
        pod.verifyCheckpointProofs({
            balanceContainerProof: cpProofs.balanceContainerProof,
            proofs: cpProofs.balanceProofs
        });
        endGas = gasleft();
        cheats.pauseGasMetering();
        totalGas = startGas - endGas;
        emit log_named_uint("== checkpoint gas", totalGas);

        // check_CompleteCheckpoint_State(staker);
        // revert();
    }

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                                (TIMING VARIANTS)
    *******************************************************************************/

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

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
    /// -- move forward 1 or more epochs
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_Advance_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

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
    function test_VerifyWC_StartCP_Advance_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

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

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (EXIT TO POD VARIANTS)
    *******************************************************************************/

    /// -- Fully exit validators before verifying withdrawal credentials
    /// 1. Verify validators' withdrawal credentials
    /// => This should fail
    function test_ExitValidators_VerifyWC_Fails(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, ) = staker.startValidators();
        staker.exitValidators(validators);
        beaconChain.advanceEpoch_NoRewards();

        cheats.expectRevert("EigenPod._verifyWithdrawalCredentials: validator must not be exiting");
        staker.verifyWithdrawalCredentials(validators);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- fully exit validators to pod
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_ExitValidators_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Fully exit one or more validators and advance epoch without generating rewards
        uint40[] memory subset = _choose(validators);
        uint64 exitedBalanceGwei = staker.exitValidators(subset);
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithExits_State(staker, subset, exitedBalanceGwei);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- fully exit validators to pod
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_StartCP_ExitValidators_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

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
    }

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (SLASH TO POD VARIANTS)
    *******************************************************************************/

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (EARN ON CL VARIANTS)
    *******************************************************************************/

    /// 1. Verify one or more validators' withdrawal credentials
    /// -- earn rewards on beacon chain (not withdrawn to pod)
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => shares increase by rewards earned
    function test_VerifyWC_Earn_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Advance epoch and generate consensus rewards, but don't withdraw excess
        beaconChain.advanceEpoch_NoWithdraw();

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        uint64 beaconBalanceIncreaseGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithRewards_State(staker, beaconBalanceIncreaseGwei, 0);
    }

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (EARN TO POD VARIANTS)
    *******************************************************************************/

    /// 1. Verify validators' withdrawal credentials
    /// -- earn rewards on beacon chain (withdrawn to pod)
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => shares increase by rewards earned
    function test_VerifyWC_EarnToPod_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Advance epoch, generating consensus rewards and withdrawing anything over 32 ETH
        beaconChain.advanceEpoch();
        uint64 beaconBalanceIncreaseGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();

        staker.startCheckpoint();
        check_StartCheckpoint_EarnToPod_State(staker, beaconBalanceIncreaseGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker); // TODO complete_withrewards?
    }

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (NATIVE ETH VARIANTS)
    *******************************************************************************/

    /// 

    /// @dev Choose a random subset of validators
    /// TODO implement
    function _choose(uint40[] memory validators) internal returns (uint40[] memory) {
        return validators;
    }
}