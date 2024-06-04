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

        CheckpointProofs memory cpProofs = beaconChain.getCheckpointProofs(validators);

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
                       VERIFY -> START/COMPLETE CHECKPOINT
    *******************************************************************************/

    /// 1. Verify one or more validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares
    function test_VerifyWC_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();
        uint40[] memory subset = _choose(validators);
        uint subsetBalanceWei = beaconChain.totalEffectiveBalanceWei(subset);

        staker.verifyWithdrawalCredentials(subset);
        check_VerifyWC_State(staker, subset, subsetBalanceWei);

        // Advance epoch without generating consensus rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    /// 1. Verify one or more validators' withdrawal credentials
    /// 2. start a checkpoint (in the same block)
    /// 3. complete a checkpoint
    /// => no change in shares
    function test_VerifyWC_SameBlock_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();
        uint40[] memory subset = _choose(validators);
        uint subsetBalanceWei = beaconChain.totalEffectiveBalanceWei(subset);

        staker.verifyWithdrawalCredentials(subset);
        check_VerifyWC_State(staker, subset, subsetBalanceWei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    /// 1. Verify one or more validators' withdrawal credentials
    /// -- fully exit one or more validators to the pod
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares
    function test_VerifyWC_ExitValidators_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();
        uint40[] memory subset = _choose(validators);
        uint subsetBalanceWei = beaconChain.totalEffectiveBalanceWei(subset);

        staker.verifyWithdrawalCredentials(subset);
        check_VerifyWC_State(staker, subset, subsetBalanceWei);

        // Fully exit validators to pod and advance epoch without generating consensus rewards
        staker.exitValidators(subset);
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    /// 1. Verify one or more validators' withdrawal credentials
    /// -- earn rewards on beacon chain (not withdrawn to pod)
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => shares increase by rewards earned
    function test_VerifyWC_Earn_StartCP_CompleteCP(uint24 _rand) public r(_rand) {

    }

    /// 1. Verify one or more validators' withdrawal credentials
    /// -- earn rewards on beacon chain (not withdrawn to pod)
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => shares increase by rewards earned
    function test_VerifyWC_Earn_Withdraw_StartCP_CompleteCP(uint24 _rand) public r(_rand) {

    }

    function test_VerifyAll_Start_CompleteCP_WithRewardsWithdrawn(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        beaconChain.advanceEpoch();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    function test_VerifyAll_Start_CompleteCP_WithRewardsNotWithdrawn(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        beaconChain.advanceEpoch_NoWithdraw();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    function test_VerifyAll_Start_CompleteCP_NoRewards(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        beaconChain.advanceEpoch_NoRewards();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    /*******************************************************************************
                    VERIFY -> EXIT -> START/COMPLETE CHECKPOINT
    *******************************************************************************/

    function test_VerifyAll_ExitAll_Start_CompleteCP_WithRewardsWithdrawn(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        uint64 exitedBalanceGwei = staker.exitValidators(validators);
        beaconChain.advanceEpoch();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithExits_State(staker, validators, exitedBalanceGwei);
    }

    function test_VerifyAll_ExitAll_Start_CompleteCP_WithRewardsNotWithdrawn(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        uint64 exitedBalanceGwei = staker.exitValidators(validators);
        beaconChain.advanceEpoch_NoWithdraw();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithExits_State(staker, validators, exitedBalanceGwei);
    }

    function test_VerifyAll_ExitAll_Start_CompleteCP_NoRewards(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();

        (uint40[] memory validators, uint beaconBalanceWei) = staker.startValidators();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceWei);

        uint64 exitedBalanceGwei = staker.exitValidators(validators);
        beaconChain.advanceEpoch_NoRewards();
        // check pod balances have increased

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithExits_State(staker, validators, exitedBalanceGwei);
    }

    /// @dev Choose a random subset of validators
    /// TODO implement
    function _choose(uint40[] memory validators) internal returns (uint40[] memory) {
        return validators;
    }
}