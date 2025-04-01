// // SPDX-License-Identifier: BUSL-1.1
// pragma solidity ^0.8.27;

// import "src/test/integration/tests/upgrade/UpgradeTest.t.sol";

// contract Integration_Upgrade_EigenPod_Slashing_Migration is UpgradeTest, EigenPodPausingConstants {
//     uint[] shares;

//     function _init() internal override {
//         _configAssetTypes(HOLDS_ETH);
//         _configUserTypes(DEFAULT);

//         /// 0. Create a staker with underlying assets
//         (staker, strategies, initTokenBalances) = _newRandomStaker();
//         shares = _calculateExpectedShares(strategies, initTokenBalances);

//         ///  1. Deposit into strategies
//         staker.depositIntoEigenlayer(strategies, initTokenBalances);
//         validators = staker.getActiveValidators();
//     }

//     function _completeEigenpodMigration() internal {
//         // Start a checkpoint
//         staker.startCheckpoint();

//         // Pause checkpoint starting
//         cheats.prank(pauserMultisig);
//         eigenPodManager.pause(2 ** PAUSED_START_CHECKPOINT);
//         cheats.expectRevert("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager");
//         staker.startCheckpoint();

//         // Complete in progress checkpoint
//         staker.completeCheckpoint();

//         // Upgrade Contracts for slashing
//         _upgradeEigenLayerContracts();

//         // Unpause EigenPodManager
//         cheats.prank(eigenLayerPauserReg.unpauser());
//         eigenPodManager.unpause(0);
//     }

//     function testFuzz_earnRewards_migrate_exit(uint24) public {
//         // 2. Advance epoch, generating consensus rewards and withdrawing anything over 32 ETH
//         beaconChain.advanceEpoch();

//         // 3. Migrate & Upgrade
//         _completeEigenpodMigration();

//         // 4. Exit validators
//         // Fully exit one or more validators and advance epoch without generating rewards
//         uint40[] memory subset = _choose(validators);
//         uint64 exitedBalanceGwei = staker.exitValidators(subset);
//         beaconChain.advanceEpoch_NoRewards();

//         staker.startCheckpoint();
//         check_StartCheckpoint_WithPodBalance_State(staker, exitedBalanceGwei);
//         staker.completeCheckpoint();
//         check_CompleteCheckpoint_WithExits_State(staker, subset, exitedBalanceGwei);
//     }

//     function testFuzz_slash_migrate(uint24) public {
//         // 2. Slash validators
//         uint40[] memory subset = _choose(validators);
//         uint64 slashedGwei = beaconChain.slashValidators(subset, BeaconChainMock.SlashType(_randUint({min: 0, max: 2})));
//         beaconChain.advanceEpoch_NoRewards();
//         shares[0] = shares[0] - slashedGwei * GWEI_TO_WEI; // Shares should decrease by slashed amount

//         // 3. Migrate & Upgrade
//         _completeEigenpodMigration();

//         // Assertions
//         assert_BCSF_WAD(staker, "BCSF should be WAD");
//         assert_ActiveValidatorCount(staker, validators.length - subset.length, "validator count should decrease");
//         assert_HasExpectedShares(staker, strategies, shares, "shares should decrease by slashed amount");
//     }
// }
