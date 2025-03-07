// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_Pectra is UpgradeTest, EigenPodPausingConstants {
    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        _configUserTypes(DEFAULT);
    }

    function test_Upgrade_VerifyWC_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        // 1. Pause, Fork, and Upgrade
        _pauseForkAndUpgrade();

        // 2. Set Pectra Fork Timestamp & unpause
        _setTimestampAndUnpause();

        // 3. Initialize Staker
        (User staker,,) = _newRandomStaker();
        (uint40[] memory validators, uint64 beaconBalanceGwei,) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        // 4. Verify Withdrawal Credentials
        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // 4. Start Checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        // 5. Complete Checkpoint
        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    function test_VerifyWC_StartCP_Fork_CompleteCP(uint24 _rand) public rand(_rand) {
        // Initialize state
        (User staker,,) = _newRandomStaker();
        (uint40[] memory validators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        // 1. Verify validators' withdrawal credentials
        staker.verifyWithdrawalCredentials(validators);

        // 2. Start a checkpoint
        staker.startCheckpoint();

        // 3. Pause, Fork, and Upgrade
        _pauseForkAndUpgrade();

        // 4. Set Pectra Fork Timestamp & unpause
        _setTimestampAndUnpause();

        // 5. Complete in progress checkpoint
        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    function test_VerifyWC_Fork_EarnToPod_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        // Initialize state
        (User staker,,) = _newRandomStaker();
        (uint40[] memory validators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        // 1. Verify validators' withdrawal credentials
        staker.verifyWithdrawalCredentials(validators);

        // 2. Pause, Fork, and Upgrade
        _pauseForkAndUpgrade();

        // 3. Set timestamp and unpause
        _setTimestampAndUnpause();

        // 4. Advance epoch, generating consensus rewards and withdrawing anything over Max EB
        // Not: Nothing is withdrawn because all validators were created with 32 ETH
        beaconChain.advanceEpoch();
        uint64 expectedEarnedGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();

        // 5. Start a checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        // 6. Complete in progress checkpoint
        staker.completeCheckpoint();
        check_CompleteCheckpoint_EarnOnBeacon_State(staker, expectedEarnedGwei);
    }

    function _pauseForkAndUpgrade() internal {
        // 1. Pause starting checkpoint, completing, and credential proofs
        cheats.prank(pauserMultisig);
        eigenPodManager.pause(
            2 ** PAUSED_START_CHECKPOINT | 2 ** PAUSED_EIGENPODS_VERIFY_CREDENTIALS | 2 ** PAUSED_VERIFY_STALE_BALANCE
                | 2 ** PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS
        );

        // 2. Fork to Pectra
        uint64 pectraForkTimestamp = uint64(block.timestamp) + 12;
        BeaconChainMock_DenebForkable(address(beaconChain)).forkToPectra(pectraForkTimestamp);

        // 3. Upgrade EigenPodManager & EigenPod
        _upgradeEigenLayerContracts();

        // 4. Set proof timestamp setter to operations multisig
        cheats.prank(eigenPodManager.owner());
        eigenPodManager.setProofTimestampSetter(address(operationsMultisig));
    }

    function _setTimestampAndUnpause() internal {
        // 1. Set Timestamp
        cheats.startPrank(eigenPodManager.proofTimestampSetter());
        eigenPodManager.setPectraForkTimestamp(BeaconChainMock_DenebForkable(address(beaconChain)).pectraForkTimestamp());
        cheats.stopPrank();

        // 2. Randomly warp to just after the fork timestamp
        // If we do not warp, proofs will be against deneb state
        if (_randBool()) {
            // If we warp, proofs will be against electra state
            cheats.warp(block.timestamp + 1);
        }

        // 3. Unpause
        cheats.prank(eigenLayerPauserReg.unpauser());
        eigenPodManager.unpause(0);
    }
}
