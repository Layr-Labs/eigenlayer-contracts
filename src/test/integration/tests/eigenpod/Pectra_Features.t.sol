// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Pectra_Features_Base is IntegrationCheckUtils {
    using ArrayLib for *;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_ETH);
    }

    function testFuzz_consolidate(uint24 _r) public rand(_r) {
        User staker = _newEmptyStaker();

        // Deal ETH and start 0x01 validators
        (uint40[] memory validators, uint64 totalBalanceGwei) = staker.startETH1Validators(uint8(_randUint({min: 2, max: 65})));

        staker.verifyWithdrawalCredentials(validators);
        check_Deposit_State(staker, beaconChainETHStrategy.toArray(), (totalBalanceGwei * GWEI_TO_WEI).toArrayU256());

        (uint40[] memory newValidators, uint40[] memory consolidated) = staker.maxConsolidation(validators);

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithConsolidations_State(staker, consolidated);

        beaconChain.advanceEpoch_NoWithdraw();

        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_EarnOnBeacon_State(staker, uint64(newValidators.length) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI());
    }

    function testFuzz_partialWithdrawals(uint24 _r) public rand(_r) {
        User staker = _newEmptyStaker();

        // Deal ETH and start 0x02 validators
        (uint40[] memory validators, uint64 totalBalanceGwei) = staker.startCompoundingValidators(uint8(_randUint({min: 2, max: 10})));

        staker.verifyWithdrawalCredentials(validators);
        check_Deposit_State(staker, beaconChainETHStrategy.toArray(), (totalBalanceGwei * GWEI_TO_WEI).toArrayU256());

        // Create partial withdrawal requests for each validator, down to 32 ETH
        uint64 amountWithdrawnGwei = staker.withdrawExcess(validators);

        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, amountWithdrawnGwei);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithPodBalance_State(staker, amountWithdrawnGwei);
    }
}
