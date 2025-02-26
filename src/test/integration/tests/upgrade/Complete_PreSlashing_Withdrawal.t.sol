// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_PreSlashing_Withdrawal_Base is UpgradeTest {
    struct TestState {
        User staker;
        User operator;
        IStrategy[] strategies;
        uint256[] tokenBalances;
        uint256[] shares;
        uint256[] withdrawalShares;
        uint256[] expectedTokens;
        Withdrawal[] withdrawals;
        bool isPartial;
        bool completeAsTokens;
    }

    function _init_(uint24 randomness, bool withOperator, bool withDelegation) internal virtual returns (TestState memory state) {
        // Create staker
        (state.staker, state.strategies, state.tokenBalances) = _newRandomStaker();
        state.shares = _calculateExpectedShares(state.strategies, state.tokenBalances);

        // Delegate staker to operator
        state.operator = withOperator ? _newRandomOperator_NoAssets() : User(payable(0));
        if (withDelegation) state.staker.delegateTo(state.operator);

        // Deposit into EigenLayer
        state.staker.depositIntoEigenlayer(state.strategies, state.tokenBalances);

        // Setup withdrawal shares (full or partial)
        state.isPartial = _randBool();
        state.withdrawalShares = new uint256[](state.shares.length);
        for (uint256 i = 0; i < state.shares.length; i++) {
            state.withdrawalShares[i] = state.isPartial ? state.shares[i] / 2 : state.shares[i];
        }

        state.expectedTokens = _calculateExpectedTokens(state.strategies, state.withdrawalShares);
        state.completeAsTokens = _randBool();
    }

    function _completeWithdrawal(TestState memory state) internal virtual {
        for (uint256 i = 0; i < state.withdrawals.length; i++) {
            if (state.completeAsTokens) {
                IERC20[] memory tokens = state.staker.completeWithdrawalAsTokens(state.withdrawals[i]);
                check_Withdrawal_AsTokens_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares, tokens, state.expectedTokens);
            } else {
                state.staker.completeWithdrawalAsShares(state.withdrawals[i]);
                check_Withdrawal_AsShares_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares);
            }
        }
    }
}

contract Integration_Upgrade_Complete_PreSlashing_Withdrawal is Integration_Upgrade_PreSlashing_Withdrawal_Base {

    function testFuzz_deposit_queue_upgrade_complete(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: false, withDelegation: false});
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _upgradeEigenLayerContracts();
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        _completeWithdrawal(state);
    }

    function testFuzz_delegate_deposit_queue_upgrade_complete(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: true, withDelegation: true});
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _upgradeEigenLayerContracts();
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        _completeWithdrawal(state);
    }

    function testFuzz_upgrade_delegate_queuePartial_complete(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: true, withDelegation: false});
        _upgradeEigenLayerContracts();
        state.staker.delegateTo(state.operator);
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        _completeWithdrawal(state);
    }

    function testFuzz_delegate_deposit_queue_completeBeforeUpgrade(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: true, withDelegation: true});
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        // We must roll forward by the delay twice since pre-upgrade delay is half as long as post-upgrade delay.
        rollForward(delegationManager.minWithdrawalDelayBlocks() + 1);
        _upgradeEigenLayerContracts();
        _completeWithdrawal(state);
    }
}

contract Integration_Upgrade_CompletePreSlashingWithdrawal_AfterSlashing is Integration_Upgrade_PreSlashing_Withdrawal_Base {
    using ArrayLib for *;

    AVS avs;
    OperatorSet operatorSet;
    AllocateParams allocateParams;

    function _init_(uint24 randomness, bool withOperator, bool withDelegation) internal override returns (TestState memory state) {
        _configAssetTypes(HOLDS_ETH);
        state = super._init_({randomness: randomness, withOperator: withOperator, withDelegation: withDelegation});

        // Queue Full Withdrawals
        state.isPartial = false;
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);

        // Upgrade contracts
        _upgradeEigenLayerContracts();

        // Set allocation delay, register for operatorSet & allocate fully
        state.operator.setAllocationDelay(1);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});
        (avs,) = _newRandomAVS();
        operatorSet = avs.createOperatorSet(state.strategies);
        allocateParams = _genAllocation_AllAvailable(state.operator, operatorSet);

        state.operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(state.operator, operatorSet, allStrats);

        state.operator.modifyAllocations(allocateParams);
        check_Base_IncrAlloc_State(state.operator, allocateParams);
        _rollBlocksForCompleteAllocation(state.operator, operatorSet, state.strategies);
    }

    function testFuzz_delegate_deposit_queue_upgrade_slash_complete(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: true, withDelegation: true});

        // Slash operator by AVS
        SlashingParams memory slashParams = _genSlashing_Half(state.operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(state.operator, allocateParams, slashParams);

        // Complete withdrawals - staker shouldn't be affected by slashing, except for BC
        // TODO: handle this properly
        state.completeAsTokens = true;
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        _completeWithdrawal(state);

        console.log("withdrawal completed");
        console.log("dsf: ", _getDepositScalingFactor(state.staker, BEACONCHAIN_ETH_STRAT));

        // Start a new validator & verify withdrawal credentials
        cheats.deal(address(state.staker), 32 ether);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = state.staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        state.staker.verifyWithdrawalCredentials(newValidators);
        IStrategy[] memory strats = new IStrategy[](1);
        strats[0] = BEACONCHAIN_ETH_STRAT;
        uint withdrawableShares = _getWithdrawableShares(state.staker, strats)[0];
        console.log("withdrawable shares: ", withdrawableShares);
        require(true == false);
        // check_VerifyWC_State(state.staker, newValidators, addedBeaconBalanceGwei);
    }
}
