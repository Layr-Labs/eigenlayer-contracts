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

        // Deal ETH and start 2 validators; both ETH1
        (uint40[] memory validators, uint balanceWei,) = staker.startValidators(2);

        staker.verifyWithdrawalCredentials(validators);
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
}

// contract Integration_FullySlashedEigenpod_Checkpointed is Integration_FullySlashedEigenpod_Base {
//     function _init() internal override {
//         super._init();

//         // // Start & complete a checkpoint
//         // staker.startCheckpoint();
//         // check_StartCheckpoint_WithPodBalance_State(staker, 0);
//         // staker.completeCheckpoint();
//         // check_CompleteCheckpoint_FullySlashed_State(staker, validators, slashedGwei);
//     }

//     function testFuzz_fullSlash_registerStakerAsOperator_Revert_Redeposit(uint24 _rand) public rand(_rand) {
//         // // Register staker as operator
//         // staker.registerAsOperator();

//         // // Start a new validator & verify withdrawal credentials
//         // cheats.deal(address(staker), 32 ether);
//         // (uint40[] memory newValidators,,) = staker.startValidators();
//         // beaconChain.advanceEpoch_NoRewards();

//         // // We should revert on verifyWithdrawalCredentials since the staker's slashing factor is 0
//         // cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
//         // staker.verifyWithdrawalCredentials(newValidators);
//     }

//     function testFuzz_fullSlash_registerStakerAsOperator_delegate_undelegate_completeAsShares(uint24 _rand) public rand(_rand) {
//         // // Register staker as operator
//         // staker.registerAsOperator();
//         // User operator = User(payable(address(staker)));

//         // // Initialize new staker
//         // (User staker2, IStrategy[] memory strategies2, uint[] memory initTokenBalances2) = _newRandomStaker();
//         // uint[] memory shares = _calculateExpectedShares(strategies2, initTokenBalances2);
//         // staker2.depositIntoEigenlayer(strategies2, initTokenBalances2);
//         // check_Deposit_State(staker2, strategies2, shares);

//         // // Delegate to an operator who has now become a staker, this should succeed as slashed operator's BCSF should not affect the staker
//         // staker2.delegateTo(operator);
//         // check_Delegation_State(staker2, operator, strategies2, shares);

//         // // Register as operator and undelegate - the equivalent of redelegating to yourself
//         // Withdrawal[] memory withdrawals = staker2.undelegate();
//         // bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
//         // check_Undelegate_State(staker2, operator, withdrawals, withdrawalRoots, strategies2, shares);

//         // // Complete withdrawals as shares
//         // _rollBlocksForCompleteWithdrawals(withdrawals);
//         // for (uint i = 0; i < withdrawals.length; i++) {
//         //     staker2.completeWithdrawalAsShares(withdrawals[i]);
//         //     check_Withdrawal_AsShares_Undelegated_State(staker2, operator, withdrawals[i], strategies2, shares);
//         // }
//     }
// }