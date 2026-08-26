// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_DisableEigenPod is IntegrationCheckUtils {
    using ArrayLib for *;

    function _init() internal virtual override {
        _configUserTypes(DEFAULT);
    }

    function testFuzz_deposit_queue_disable_cleanup(uint24 _random) public rand(_random) {
        _configAssetTypes(HOLDS_ETH);
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        // 1. Deposit native ETH into EigenLayer.
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, depositShares);
        _checkpointPod(staker);

        // 2. Queue all native ETH shares.
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(
            staker, User(payable(address(0))), strategies, depositShares, withdrawableShares, withdrawals, withdrawalRoots
        );

        // 3. Disable the pod once the beacon withdrawal is no longer slashable.
        _rollBlocksForCompleteWithdrawals(withdrawals);
        staker.disablePod();
        assertTrue(staker.pod().restakingDisabled(), "pod should be disabled");
        assertEq(eigenPodManager.podOwnerDepositShares(address(staker)), 0, "native shares should remain queued out");
        assertEq(delegationManager.getQueuedWithdrawalRoots(address(staker)).length, 0, "disable should clear beacon withdrawal");
    }

    function testFuzz_depositQueueMixedSeparate_disable_completeLST(uint24 _random) public rand(_random) {
        IStrategy[] memory strategies = new IStrategy[](2);
        strategies[0] = beaconChainETHStrategy;
        strategies[1] = lstStrats[0];
        (User staker, uint[] memory tokenBalances) = _newStaker(strategies);

        // 1. Deposit native ETH and one LST.
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, depositShares);
        _checkpointPod(staker);

        // 2. Queue beacon-chain ETH and LST as separate withdrawals.
        IStrategy[] memory beaconStrategy = beaconChainETHStrategy.toArray();
        uint[] memory beaconShares = depositShares[0].toArrayU256();
        uint[] memory beaconWithdrawableShares = _getStakerWithdrawableShares(staker, beaconStrategy);
        Withdrawal[] memory beaconWithdrawals = staker.queueWithdrawals(beaconStrategy, beaconShares);
        check_QueuedWithdrawal_State(
            staker,
            User(payable(address(0))),
            beaconStrategy,
            beaconShares,
            beaconWithdrawableShares,
            beaconWithdrawals,
            _getWithdrawalHashes(beaconWithdrawals)
        );

        IStrategy[] memory lstStrategy = strategies[1].toArray();
        uint[] memory lstShares = depositShares[1].toArrayU256();
        uint[] memory lstWithdrawableShares = _getStakerWithdrawableShares(staker, lstStrategy);
        Withdrawal[] memory lstWithdrawals = staker.queueWithdrawals(lstStrategy, lstShares);
        check_QueuedWithdrawal_State(
            staker,
            User(payable(address(0))),
            lstStrategy,
            lstShares,
            lstWithdrawableShares,
            lstWithdrawals,
            _getWithdrawalHashes(lstWithdrawals)
        );

        // 3. Disable the pod.
        _rollBlocksForCompleteWithdrawals(lstWithdrawals);
        staker.disablePod();
        assertTrue(staker.pod().restakingDisabled(), "pod should be disabled");
        assertEq(delegationManager.getQueuedWithdrawalRoots(address(staker)).length, 1, "only LST withdrawal should remain");

        // 4. Pure LST withdrawal remains completable after native restaking is disabled.
        uint[] memory expectedTokens = _calculateExpectedTokens(lstStrategy, lstWithdrawableShares);
        IERC20[] memory tokens = staker.completeWithdrawalAsTokens(lstWithdrawals[0]);
        check_Withdrawal_AsTokens_State(
            staker, User(payable(address(0))), lstWithdrawals[0], lstStrategy, lstWithdrawableShares, tokens, expectedTokens
        );
        assertEq(delegationManager.getQueuedWithdrawalRoots(address(staker)).length, 0, "all withdrawals should be complete or cleared");
    }

    function testFuzz_deposit_beaconSlash_queue_disable(uint24 _random) public rand(_random) {
        _configAssetTypes(HOLDS_ETH);
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        cheats.assume(tokenBalances[0] >= 64 ether);

        // 1. Deposit native ETH into EigenLayer.
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, depositShares);

        // 2. Slash on the beacon chain and checkpoint the lower balance.
        uint40[] memory slashedValidators = _choose(staker.getActiveValidators());
        uint64 slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.startCheckpoint();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithCLSlashing_HandleRoundDown_State(staker, slashedGwei);
        assertLt(eigenPodManager.beaconChainSlashingFactor(address(staker)), WAD, "BCSF should be slashed");

        // 3. Queue all remaining native ETH shares.
        depositShares = _getStakerDepositShares(staker, strategies);
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositShares);
        check_QueuedWithdrawal_State(
            staker, User(payable(address(0))), strategies, depositShares, withdrawableShares, withdrawals, _getWithdrawalHashes(withdrawals)
        );

        // 4. A beacon-chain-slashed pod can still disable once the conservation check passes.
        _rollBlocksForCompleteWithdrawals(withdrawals);
        staker.disablePod();
        assertTrue(staker.pod().restakingDisabled(), "slashed pod should be disabled");
    }

    function testFuzz_delegateToSlashedOperator_beaconSlash_queue_disableWithinRoundingTolerance(uint24 _random) public rand(_random) {
        _configAssetTypes(HOLDS_ETH);
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();
        cheats.assume(tokenBalances[0] >= 64 ether);

        // 1. Slash an operator before the staker delegates, creating a non-trivial DSF on later native deposits.
        {
            OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
            operator.registerForOperatorSet(operatorSet);
            check_Registration_State_NoAllocation(operator, operatorSet, strategies);
            AllocateParams memory allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
            operator.modifyAllocations(allocateParams);
            check_IncrAlloc_State_Slashable_NoDelegatedStake(operator, allocateParams);
            _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

            SlashingParams memory slashParams = _genSlashing_Custom(operator, operatorSet, 333_333_333_333_333_333);
            (uint slashId,) = avs.slashOperator(slashParams);
            check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);
        }

        // 2. Delegate to the already-slashed operator, then deposit native ETH.
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, new IStrategy[](0), new uint[](0));
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, depositShares);
        assertGt(delegationManager.depositScalingFactor(address(staker), beaconChainETHStrategy), WAD, "DSF should forgive prior slash");

        // 3. Beacon slash the pod to combine operator magnitude rounding with BCSF rounding.
        {
            uint40[] memory slashedValidators = _choose(staker.getActiveValidators());
            uint64 slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Minor);
            beaconChain.advanceEpoch_NoWithdrawNoRewards();
            staker.startCheckpoint();
            staker.completeCheckpoint();
            check_CompleteCheckpoint_WithCLSlashing_HandleRoundDown_State(staker, slashedGwei);
        }

        // 4. Queue all native ETH shares and verify the disable conservation difference is bounded by 1 gwei.
        depositShares = _getStakerDepositShares(staker, strategies);
        Withdrawal[] memory withdrawals;
        {
            uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
            withdrawals = staker.queueWithdrawals(strategies, depositShares);
            check_QueuedWithdrawal_State(
                staker, operator, strategies, depositShares, withdrawableShares, withdrawals, _getWithdrawalHashes(withdrawals)
            );
        }

        _rollBlocksForCompleteWithdrawals(withdrawals);
        (uint podBalanceWei, uint queuedBeaconWei) = _disableConservationValues(staker);
        assertLe(podBalanceWei, queuedBeaconWei + GWEI_TO_WEI, "valid rounding gap should fit in 1 gwei");

        // 5. Disable succeeds with the same real rounded values.
        staker.disablePod();
        assertTrue(staker.pod().restakingDisabled(), "pod should be disabled");
    }

    function testFuzz_delegate_queue_slashBeforeDelay_revert_disable(uint24 _random) public rand(_random) {
        _configAssetTypes(HOLDS_ETH);
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        // 1. Deposit native ETH and delegate to an operator.
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, depositShares);
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, depositShares);
        _checkpointPod(staker);

        // 2. Allocate the delegated native stake to an AVS.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);
        AllocateParams memory allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 3. Queue all native shares.
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositShares);
        check_QueuedWithdrawal_State(
            staker, operator, strategies, depositShares, withdrawableShares, withdrawals, _getWithdrawalHashes(withdrawals)
        );

        // 4. Slash while the queued withdrawal is still slashable.
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        (uint slashId,) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        // 5. Disable must not let the staker escape the queue slash.
        _rollBlocksForCompleteWithdrawals(withdrawals);
        (uint podBalanceWei, uint queuedBeaconWei) = _disableConservationValues(staker);
        assertGt(podBalanceWei, queuedBeaconWei + GWEI_TO_WEI, "queue slash should reduce entitlement beyond tolerance");

        cheats.expectRevert(IEigenPodManagerErrors.PodValueExceedsQueuedWithdrawals.selector);
        staker.disablePod();
    }

    function testFuzz_deposit_delegate_undelegate_disable_cleanup(uint24 _random) public rand(_random) {
        _configAssetTypes(HOLDS_ETH);
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        User operator = _newRandomOperator_NoAssets();

        // 1. Deposit native ETH and delegate.
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, depositShares);
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, depositShares);
        _checkpointPod(staker);

        // 2. Undelegate, creating a real delegation-lifecycle queued beacon withdrawal.
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 3. Disable after the queue is no longer slashable.
        _rollBlocksForCompleteWithdrawals(withdrawals);
        staker.disablePod();
        assertTrue(staker.pod().restakingDisabled(), "pod should be disabled");

        // 4. Disable clears the undelegate-created beacon withdrawal.
        assertEq(delegationManager.getQueuedWithdrawalRoots(address(staker)).length, 0, "disable should clear beacon withdrawal");
    }

    function testFuzz_deposit_queue_disable_exitValidators_sweep(uint24 _random) public rand(_random) {
        User staker = _newEmptyStaker();

        // 1. Start 0x01 validators and verify them into EigenLayer.
        (uint40[] memory validators, uint64 totalBalanceGwei) = staker.startETH1Validators(uint8(_randUint({min: 2, max: 8})));
        staker.verifyWithdrawalCredentials(validators);
        IStrategy[] memory strategies = beaconChainETHStrategy.toArray();
        uint[] memory depositShares = (totalBalanceGwei * GWEI_TO_WEI).toArrayU256();
        check_Deposit_State(staker, strategies, depositShares);
        _checkpointPod(staker);

        // 2. Queue all native ETH shares and disable once the queue is no longer slashable.
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositShares);
        check_QueuedWithdrawal_State(
            staker, User(payable(address(0))), strategies, depositShares, withdrawableShares, withdrawals, _getWithdrawalHashes(withdrawals)
        );
        _rollBlocksForCompleteWithdrawals(withdrawals);
        staker.disablePod();
        assertTrue(staker.pod().restakingDisabled(), "pod should be disabled");

        // 3. After disable, the owner can request full validator exits without re-entering restaking.
        uint64 exitedBalanceGwei = staker.exitValidators(validators);
        beaconChain.advanceEpoch_NoRewards();
        assertEq(address(staker.pod()).balance, uint(exitedBalanceGwei) * GWEI_TO_WEI, "exited ETH should land in disabled pod");

        // 4. The owner can sweep ETH that arrived in the disabled pod.
        address recipient = address(0xBEEF);
        uint recipientBalanceBefore = recipient.balance;
        uint podBalanceBefore = address(staker.pod()).balance;
        staker.withdrawDisabledPodETH(recipient);
        assertEq(address(staker.pod()).balance, 0, "pod should be swept");
        assertEq(recipient.balance, recipientBalanceBefore + podBalanceBefore, "recipient should receive disabled pod ETH");
    }

    function _disableConservationValues(User staker) internal view returns (uint podBalanceWei, uint queuedBeaconWei) {
        (Withdrawal[] memory withdrawals, uint[][] memory shares) = delegationManager.getQueuedWithdrawals(address(staker));

        for (uint i; i < withdrawals.length; ++i) {
            for (uint j; j < withdrawals[i].strategies.length; ++j) {
                if (withdrawals[i].strategies[j] == beaconChainETHStrategy) queuedBeaconWei += shares[i][j];
            }
        }

        IEigenPodTypes.Checkpoint memory checkpoint = staker.pod().currentCheckpoint();
        int podBalanceGwei = int(uint(staker.pod().withdrawableRestakedExecutionLayerGwei())) + int(uint(checkpoint.prevBeaconBalanceGwei))
            + int(checkpoint.balanceDeltasGwei);
        podBalanceWei = podBalanceGwei <= 0 ? 0 : uint(podBalanceGwei) * GWEI_TO_WEI;
    }

    function _checkpointPod(User staker) internal {
        staker.startCheckpoint();
        staker.completeCheckpoint();
    }
}
