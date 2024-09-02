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

    bytes execUpgradeCalldata = hex"0825f38f000000000000000000000000369e6f597e22eab55ffb173c6d9cd234bd699111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000066d867e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003646a76120200000000000000000000000040a2accbd92bca938b02010e17a5b8929b49130d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000001648d80ff0a00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000112005a2a4f2f3c18f09179b6703e63d9edd165909073000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000243659cfe60000000000000000000000006d225e974fa404d25ffb84ed6e242ffa18ef6430008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000091e677b07f7af907ec9a428aafa9fc14a0d3a338000000000000000000000000731a0ad160e407393ff662231add6dd145ad3fea0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041000000000000000000000000a6db1a8c5a981d1536266d2a393c5f8ddb210eaf0000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";
    uint timelockETA = 1725458400;    
    address podOwner = 0xefF584E8336dA7A23EE32ea19a937b016D69d589;
    address testPod = 0xA6f93249580EC3F08016cD3d4154AADD70aC3C96;

    function test_PEPE_Upgrade() public r(0) {
        EigenPod pod = EigenPod(payable(testPod));

        emit log_named_uint("pod balance", testPod.balance);

        cheats.startPrank(podOwner);

        cheats.expectRevert();
        pod.activeValidatorCount();
        cheats.expectRevert();
        pod.startCheckpoint(false);

        cheats.stopPrank();

        _doUpgrade();
        beaconChain.advanceEpoch_NoWithdraw();

        cheats.startPrank(podOwner);
        require(pod.activeValidatorCount() == 1, "incorrect active validator count");
        pod.startCheckpoint(true);
        IEigenPod.Checkpoint memory cp = pod.currentCheckpoint();
        emit log_named_bytes32("cp.blockRoot", cp.beaconBlockRoot);
        emit log_named_uint("cp.proofsRemaining", cp.proofsRemaining);
        emit log_named_uint("cp.podBalanceGwei", cp.podBalanceGwei);
        emit log_named_int("cp.balanceDeltasGwei", cp.balanceDeltasGwei);
    }

    function _doUpgrade() internal {
        cheats.warp(timelockETA);
        cheats.prank(operationsMultisig);
        timelock.call(execUpgradeCalldata);
    }

    function test_GasMetering() public r(0) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();
        // Deal user 20 full stakers worth of ETH
        emit log_named_string("Dealing 20 * 32 ETH to", staker.NAME());
        cheats.deal(address(staker), 32 ether * 20);
        cheats.pauseGasMetering();

        (uint40[] memory validators, ) = staker.startValidators();
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
        _upgradeEigenLayerContracts();

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
    function test_VerifyWC_VerifyWC_Fails(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        cheats.expectRevert("EigenPod._verifyWithdrawalCredentials: validator must be inactive to prove withdrawal credentials");
        staker.verifyWithdrawalCredentials(validators);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. start a checkpoint again
    /// => This should fail
    function test_VerifyWC_StartCP_StartCP_Fails(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        cheats.expectRevert("EigenPod._startCheckpoint: must finish previous checkpoint before starting another");
        staker.startCheckpoint();
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// 4. start a checkpoint without advancing a block
    /// => this should fail
    function test_VerifyWC_StartCP_CompleteCP_StartCP_Fails(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);

        cheats.expectRevert("EigenPod._startCheckpoint: cannot checkpoint twice in one block");
        staker.startCheckpoint();
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- move forward 1 or more epochs
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_Advance_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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
        _upgradeEigenLayerContracts();

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
        _upgradeEigenLayerContracts();

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
        _upgradeEigenLayerContracts();

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
    function test_VerifyWC_StartCP_ExitValidators_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (SLASH TO POD VARIANTS)
    *******************************************************************************/

    /// -- get slashed on beacon chain
    /// 1. Try to verify validators' withdrawal credentials
    /// => this should fail
    function test_SlashToPod_VerifyWC_Fails(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, ) = staker.startValidators();
        beaconChain.slashValidators(validators);
        // Advance epoch, withdrawing slashed validators to pod
        beaconChain.advanceEpoch_NoRewards();
        
        cheats.expectRevert("EigenPod._verifyWithdrawalCredentials: validator must not be exiting");
        staker.verifyWithdrawalCredentials(validators);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- get slashed on beacon chain; exit to pod
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares should decrease by slashed amount
    function test_VerifyWC_SlashToPod_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators);
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
    function test_VerifyWC_StartCP_SlashToPod_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators);
        beaconChain.advanceEpoch_NoRewards();

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, validators, slashedBalanceGwei);
    }

    /*******************************************************************************
                       VERIFY -> PROVE STALE BALANCE -> COMPLETE CHECKPOINT
    *******************************************************************************/

    /// 1. Verify validators' withdrawal credentials
    /// -- get slashed on beacon chain; exit to pod
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares should decrease by slashed amount
    function test_VerifyWC_SlashToPod_VerifyStale_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators);
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
    function test_VerifyWC_SlashToCL_VerifyStale_CompleteCP_SlashAgain(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Slash validators but do not process exits to pod
        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators);
        beaconChain.advanceEpoch_NoWithdraw();

        staker.verifyStaleBalance(validators[0]);
        check_StartCheckpoint_WithPodBalance_State(staker, 0);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithCLSlashing_State(staker, slashedBalanceGwei);

        // Slash validators again but do not process exits to pod
        uint64 secondSlashedBalanceGwei = beaconChain.slashValidators(validators);
        beaconChain.advanceEpoch_NoWithdraw();

        staker.verifyStaleBalance(validators[0]);
        check_StartCheckpoint_WithPodBalance_State(staker, 0);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithCLSlashing_State(staker, secondSlashedBalanceGwei);
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
    function test_VerifyWC_StartCP_SlashToPod_CompleteCP_VerifyStale(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        // Advance epoch without generating rewards
        beaconChain.advanceEpoch_NoRewards();

        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        uint64 slashedBalanceGwei = beaconChain.slashValidators(validators);
        beaconChain.advanceEpoch_NoRewards();

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);

        staker.verifyStaleBalance(validators[0]);
        check_StartCheckpoint_WithPodBalance_State(staker, beaconBalanceGwei - slashedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, validators, slashedBalanceGwei);
    }

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (EARN ON CL VARIANTS)
    *******************************************************************************/

    /// -- earn rewards on beacon chain (not withdrawn to pod)
    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares increase by rewards earned on beacon chain
    function test_EarnOnBeacon_VerifyWC_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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
    function test_VerifyWC_EarnOnBeacon_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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
    function test_VerifyWC_StartCP_EarnOnBeacon_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (EARN TO POD VARIANTS)
    *******************************************************************************/

    /// -- earn rewards on beacon chain (withdrawn to pod)
    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares increase by rewards withdrawn to pod
    function test_EarnToPod_VerifyWC_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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
    function test_VerifyWC_EarnToPod_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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
    function test_VerifyWC_StartCP_EarnToPod_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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

    /*******************************************************************************
                       VERIFY -> START -> COMPLETE CHECKPOINT
                             (NATIVE ETH VARIANTS)
    *******************************************************************************/

    /// -- Pod receives native ETH via fallback
    /// 1. start a checkpoint
    /// => checkpoint should auto-complete, awarding shares for ETH in pod
    function test_NativeETH_StartCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();
        // Send a random amount of ETH to staker's fallback
        (uint64 gweiSent, ) = _sendRandomETH(address(staker.pod()));

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
    function test_NativeETH_VerifyWC_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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
        // check that `pod.balance == withdrawableRestakedExecutionLayerGwei + remainderSent
        assert_PodBalance_Eq(staker, (gweiSent * GWEI_TO_WEI) + remainderSent, "pod balance should equal expected");
        check_CompleteCheckpoint_WithPodBalance_State(staker, gweiSent);
    }

    /// 1. Verify validators' withdrawal credentials
    /// -- Pod receives native ETH via fallback
    /// 2. start a checkpoint
    /// 3. complete a checkpoint
    /// => after 3, shares should account for native ETH
    function test_VerifyWC_NativeETH_StartCP_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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
        // check that `pod.balance == withdrawableRestakedExecutionLayerGwei + remainderSent
        assert_PodBalance_Eq(staker, (gweiSent * GWEI_TO_WEI) + remainderSent, "pod balance should equal expected");
        check_CompleteCheckpoint_WithPodBalance_State(staker, gweiSent);
    }

    /// 1. Verify validators' withdrawal credentials
    /// 2. start a checkpoint
    /// -- Pod receives native ETH via fallback
    /// 3. complete a checkpoint
    /// => no change in shares between 1 and 3
    function test_VerifyWC_StartCP_NativeETH_CompleteCP(uint24 _rand) public r(_rand) {
        (User staker, ,) = _newRandomStaker();
        _upgradeEigenLayerContracts();

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