// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_EigenPod_Slashing_Migration is UpgradeTest, EigenPodPausingConstants {
    
    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        _configUserTypes(DEFAULT);
    }

    /**
     * 1. Verify validators' withdrawal credentials
     *    -- earn rewards on beacon chain (withdrawn to pod)
     * 2. Start a checkpoint
     * 3. Pause starting checkpoints
     * 4. Complete in progress checkpoint
     * 5. Upgrade EigenPod contracts
     * 6. Exit subset of Validators 
     */
    function test_upgrade_eigenpod_migration(uint24 _rand) public rand(_rand) {
        // Initialize state
        (User staker, ,) = _newRandomStaker();    

        (uint40[] memory validators, ) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards(); 

        // 1. Verify validators' withdrawal credentials
        staker.verifyWithdrawalCredentials(validators);

        // Advance epoch, generating consensus rewards and withdrawing anything over 32 ETH
        beaconChain.advanceEpoch();

        // 2. Start a checkpoint
        staker.startCheckpoint();

        // 3. Pause checkpoint starting
        cheats.prank(pauserMultisig);
        eigenPodManager.pause(2 ** PAUSED_START_CHECKPOINT);
        cheats.expectRevert("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager");
        staker.startCheckpoint();

        // 4. Complete in progress checkpoint
        staker.completeCheckpoint();

        // 5. Upgrade Contracts for slashing
        _upgradeEigenLayerContracts();

        // Unpause EigenPodManager
        cheats.prank(eigenLayerPauserReg.unpauser());
        eigenPodManager.unpause(0);

        // 6. Exit validators
        // Fully exit one or more validators and advance epoch without generating rewards
        uint40[] memory subset = _choose(validators);
        uint64 exitedBalanceGwei = staker.exitValidators(subset);
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, exitedBalanceGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithExits_State(staker, subset, exitedBalanceGwei);
    }
}