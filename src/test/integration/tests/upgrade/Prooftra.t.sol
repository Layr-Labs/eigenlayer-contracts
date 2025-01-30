// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_Pectra is UpgradeTest {
    
    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        _configUserTypes(DEFAULT);
    }        

    function test_VerifyWC_StartCP_Fork_CompleteCP(uint24 _rand) public rand(_rand) {
        // Initialize state
        (User staker, ,) = _newRandomStaker();    

        (uint40[] memory validators, ) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards(); 

        // 1. Verify validators' withdrawal credentials
        staker.verifyWithdrawalCredentials(validators);

        // 2. Start a checkpoint
        staker.startCheckpoint();

        // 3. Upgrade Contracts for Pectra
        _upgradeEigenLayerContracts();
        _hardForkToPectra();

        // 4. Complete in progress checkpoint
        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    function test_VerifyWC_Fork_EarnToPod_StartCP_CompleteCP(uint24 _rand) public rand(_rand) {
        // Initialize state
        (User staker, ,) = _newRandomStaker();    

        (uint40[] memory validators, ) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards(); 

        // 1. Verify validators' withdrawal credentials
        staker.verifyWithdrawalCredentials(validators);

        // 2. Upgrade Contracts for Pectra
        _upgradeEigenLayerContracts();
        _hardForkToPectra();

        // Randomly advance timestamp to be just after the hard fork timestamp
        if (_randBool()) {
            cheats.warp(block.timestamp + 1);
        }

        // 3. Advance epoch, generating consensus rewards and withdrawing anything over 32 ETH
        beaconChain.advanceEpoch();
        uint64 expectedWithdrawnGwei = uint64(validators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI();

        // 4. Start a checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, expectedWithdrawnGwei);

        // 5. Complete in progress checkpoint
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithPodBalance_State(staker, expectedWithdrawnGwei);
    }

    function _hardForkToPectra() internal {
        beaconChain.forkToPectra();
    }

}