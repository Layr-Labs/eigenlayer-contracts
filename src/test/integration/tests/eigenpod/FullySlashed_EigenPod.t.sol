// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_FullySlashedEigenpod_Base is IntegrationCheckUtils {
    using ArrayLib for *;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;
    uint64 slashedGwei;
    uint40[] validators;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        cheats.assume(initTokenBalances[0] >= 64 ether);

        // Deposit staker
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);
        initDepositShares = shares;
        validators = staker.getActiveValidators();

        // Slash all validators fully
        slashedGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Full);
        beaconChain.advanceEpoch_NoRewards(); // Withdraw slashed validators to pod
    }
}

contract Integration_FullySlashedEigenpod_Checkpointed is Integration_FullySlashedEigenpod_Base {
    function _init() internal override {
        super._init();

        // Start & complete a checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, 0);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_FullySlashed_State(staker, validators, slashedGwei);
    }

    function testFuzz_fullSlash_Delegate(uint24 _rand) public rand(_rand) {
        (User operator,,) = _newRandomOperator();

        // Delegate to an operator - should succeed given that delegation only checks the operator's slashing factor
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);
    }

    function testFuzz_fullSlash_Revert_Redeposit(uint24 _rand) public rand(_rand) {
        // Start a new validator & verify withdrawal credentials
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        // We should revert on verifyWithdrawalCredentials since the staker's slashing factor is 0
        cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
        staker.verifyWithdrawalCredentials(newValidators);
    }

    function testFuzz_fullSlash_registerStakerAsOperator_Revert_Redeposit(uint24 _rand) public rand(_rand) {
        // Register staker as operator
        staker.registerAsOperator();

        // Start a new validator & verify withdrawal credentials
        cheats.deal(address(staker), 32 ether);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();

        // We should revert on verifyWithdrawalCredentials since the staker's slashing factor is 0
        cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
        staker.verifyWithdrawalCredentials(newValidators);
    }

    function testFuzz_fullSlash_registerStakerAsOperator_delegate_undelegate_completeAsShares(uint24 _rand) public rand(_rand) {
        // Register staker as operator
        staker.registerAsOperator();
        User operator = User(payable(address(staker)));

        // Initialize new staker
        (User staker2, IStrategy[] memory strategies2, uint[] memory initTokenBalances2) = _newRandomStaker();
        uint[] memory shares = _calculateExpectedShares(strategies2, initTokenBalances2);
        staker2.depositIntoEigenlayer(strategies2, initTokenBalances2);
        check_Deposit_State(staker2, strategies2, shares);

        // Delegate to an operator who has now become a staker, this should succeed as slashed operator's BCSF should not affect the staker
        staker2.delegateTo(operator);
        check_Delegation_State(staker2, operator, strategies2, shares);

        // Register as operator and undelegate - the equivalent of redelegating to yourself
        Withdrawal[] memory withdrawals = staker2.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker2, operator, withdrawals, withdrawalRoots, strategies2, shares);

        // Complete withdrawals as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker2.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker2, operator, withdrawals[i], strategies2, shares);
        }
    }
}

contract Integration_FullySlashedEigenpod_NotCheckpointed is Integration_FullySlashedEigenpod_Base {
    /// @dev Adding funds prior to checkpointing allows the pod to not be "bricked"
    function testFuzz_proveValidator_checkpoint_queue_completeAsTokens(uint24 _rand) public rand(_rand) {
        // Deal ETH to staker
        uint amount = 32 ether;
        cheats.deal(address(staker), amount);
        uint[] memory initTokenBalances2 = new uint[](1);
        initTokenBalances2[0] = amount;

        // Deposit staker
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances2);
        staker.depositIntoEigenlayer(strategies, initTokenBalances2);
        check_Deposit_State(staker, strategies, shares);

        // Checkpoint slashed EigenPod
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, 0);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(staker, validators, slashedGwei);

        // Queue Full Withdrawal
        uint[] memory depositShares = _getStakerDepositShares(staker, strategies);
        uint[] memory withdrawableShares = _getWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(
            staker, User(payable(address(0))), strategies, depositShares, withdrawableShares, withdrawals, withdrawalRoots
        );

        // Complete withdrawal as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = _getUnderlyingTokens(withdrawals[i].strategies);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawableShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                staker, User(payable(address(0))), withdrawals[i], withdrawals[i].strategies, withdrawableShares, tokens, expectedTokens
            );
        }
    }

    function testFuzz_depositMinimumAmount_checkpoint(uint24 _rand) public rand(_rand) {
        // Deal ETH to staker, minimum amount to be checkpointed
        uint64 podBalanceGwei = 1;
        uint amountToDeal = 1 * GWEI_TO_WEI;
        bool isBricked;

        // Randomly deal 1 less than minimum amount to be checkpointed such that the pod is bricked
        if (_randBool()) {
            amountToDeal -= 1;
            podBalanceGwei -= 1;
            isBricked = true;
        }

        // Send ETH to pod
        cheats.prank(address(staker));
        (bool success,) = address(staker.pod()).call{value: amountToDeal}("");
        require(success, "pod call failed");

        // Checkpoint slashed EigenPod
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, podBalanceGwei);
        staker.completeCheckpoint();
        if (isBricked) {
            // BCSF is asserted to be zero here
            check_CompleteCheckpoint_FullySlashed_State(staker, validators, slashedGwei);
        } else {
            check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(staker, validators, slashedGwei - podBalanceGwei);
        }
    }
}
