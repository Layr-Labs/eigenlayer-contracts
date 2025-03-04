// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_CompleteCP_WithNoAddedShares is IntegrationCheckUtils {
    using ArrayLib for *;
    
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    IAllocationManagerTypes.AllocateParams allocateParams;

    User staker;
    uint40[] validators;
    uint64 beaconBalanceGwei;
    IStrategy[] strategies;
    uint[] initDepositShares;
    uint64 slashedGwei;

    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initDepositShares) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        cheats.assume(initDepositShares[0] >= 64 ether);

        // 1. start validators, advance epoch
        (validators, beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        
        // 2. deposit: verify wc
        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        uint[] memory shares = _calculateExpectedShares(strategies, initDepositShares);
        check_Deposit_State(staker, strategies, shares);

        // 3. slash validators (half)
        uint40[] memory slashedValidators = _choose(validators);
        slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Half);
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        
        // 4. start/complete cp to reduce BCSF
        staker.startCheckpoint();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(staker, slashedValidators, slashedGwei);
    }

    /**
     * @notice Test sets up an EigenPod which has a non-WAD BCSF. After queue withdrawing all depositShares
     * which sets it to 0, they can then complete checkpoints repeatedly with 0 shares increase to increase the staker DSF each time
     */
    function test_completeCP_withNoAddedShares(uint24 _rand) public rand(_rand) {
        // 5. queue withdraw all depositShares having it set to 0
        uint withdrawableSharesBefore = _getStakerWithdrawableShares(staker, strategies)[0];
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);

        // 6. advance epoch no rewards
        // start/complete cp repeatedly with 0 shares increase to increase the staker DSF each time
        for (uint256 i = 0; i < 10; i++) {
            beaconChain.advanceEpoch_NoWithdrawNoRewards();
            staker.startCheckpoint();
            staker.completeCheckpoint();
        }

        // Staker deposit shares should be 0 from queue withdrawing all depositShares
        // therefore the depositScalingFactor should also be reset WAD
        assertEq(eigenPodManager.podOwnerDepositShares(address(staker)), 0);
        assertEq(delegationManager.depositScalingFactor(address(staker), beaconChainETHStrategy), WAD);

        // 7. deposit: can either verify wc or start/complete cp or complete the withdrawals as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        staker.completeWithdrawalsAsShares(withdrawals);

        // 8. delegateTo an operator
        staker.delegateTo(operator);
        // End state: staker and operator have much higher inflated withdrawable and delegated shares respectively
        // The staker's withdrawable shares should be <= from withdrawable shares before (should be equal but could be less due to rounding)
        uint withdrawableSharesAfter = _getStakerWithdrawableShares(staker, strategies)[0];
        uint operatorShares = delegationManager.operatorShares(address(operator), strategies[0]);
        assertLe(
            withdrawableSharesAfter,
            withdrawableSharesBefore,
            "staker withdrawable shares should be <= from withdrawable shares before"
        );
        assertLe(
            operatorShares,
            withdrawableSharesBefore,
            "operatorShares should be <= from withdrawable shares before"
        );
    }
}
