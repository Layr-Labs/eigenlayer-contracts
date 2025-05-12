// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Pectra_Features_Base is IntegrationCheckUtils {
    using ArrayLib for *;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_ETH);
        // (staker, strategies, initTokenBalances) = _newRandomStaker();
        // cheats.assume(initTokenBalances[0] >= 64 ether);

        // // Deposit staker
        // uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        // staker.depositIntoEigenlayer(strategies, initTokenBalances);
        // check_Deposit_State(staker, strategies, shares);
        // initDepositShares = shares;
        // validators = staker.getActiveValidators();

        // // Slash all validators fully
        // slashedGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Full);
        // beaconChain.advanceEpoch_NoRewards(); // Withdraw slashed validators to pod
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

    function _zip(uint40[] memory a1, uint40[] memory a2) internal pure returns (uint40[] memory) {
        uint40[] memory result = new uint40[](a1.length + a2.length);

        uint resultIdx;
        for (uint i = 0; i < a1.length; i++) {
            result[resultIdx] = a1[i];
            resultIdx++;
        }

        for (uint i = 0; i < a2.length; i++) {
            result[resultIdx] = a2[i];
            resultIdx++;
        }

        return result;
    }

    function _printValidators(uint40[] memory validators) internal {
        console.log("Info on %d validators", validators.length);
        console.log("- total effective balance (gwei): %d", beaconChain.totalEffectiveBalanceGwei(validators));

        for (uint i = 0; i < validators.length; i++) {
            uint40 v = validators[i];
            bool hasCompounding = beaconChain.hasCompoundingCreds(v);

            console.log("Validator (%d):", v);

            console.log(hasCompounding ? "- credentials: 0x02" : "- credentials: 0x01");
            console.log("- effective balance (gwei): %d", beaconChain.effectiveBalance(v));
        }
    }
}
