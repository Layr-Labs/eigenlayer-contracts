// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_FullySlashedEigenpod is IntegrationCheckUtils {
    using ArrayLib for *;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;
    uint64 slashedGwei;

    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initTokenBalances) = _newRandomStaker();

        cheats.assume(initTokenBalances[0] >= 64 ether);

        // Deposit staker
        (uint40[] memory validators,) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(validators);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        initDepositShares = shares;
        check_Deposit_State(staker, strategies, shares);

        // Slash all validators fully
        slashedGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Full);
        beaconChain.advanceEpoch_NoRewards(); // Withdraw slashed validators to pod

        // Start & complete a checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, 0);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_FullySlashed_State(staker, validators, slashedGwei);
    }

    function test_fullSlash_Delegate(uint24 _rand) public rand(_rand) {
        (User operator,,) = _newRandomOperator();

        // Delegate to an operator - should succeed given that delegation only checks the operator's slashing factor
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);
    }

    function test_fullSlash_Revert_Redeposit(uint24 _rand) public rand(_rand) {
        // Start a new validator & verify withdrawal credentials
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        // We should revert on verifyWithdrawalCredentials since the staker's slashing factor is 0
        cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
        staker.verifyWithdrawalCredentials(newValidators);
    }
}