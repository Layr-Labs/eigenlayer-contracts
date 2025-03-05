// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/mocks/BeaconChainMock.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

/// @notice Testing the rounding behavior when operator magnitude is initially 1
contract Integration_Register_Allocate_Slash_VerifyWC is IntegrationCheckUtils {
    using ArrayLib for *;

    AVS avs;
    OperatorSet operatorSet;

    User operator;
    IAllocationManagerTypes.AllocateParams allocateParams;

    User staker;
    IStrategy[] strategies;
    // uint[] initTokenBalances;
    uint[] initDepositShares;
    uint40[] validators;
    uint64 beaconBalanceGwei;
    uint64 slashedGwei;

    /**
     * 1. Create an operatorSet and register the operator allocating all magnitude
     * 2. slash operator to 1 magnitude remaining
     * 3. delegate staker to operator
     * 4. deposit (verify withdrawal credentials)
     */
    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initDepositShares) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        cheats.assume(initDepositShares[0] >= 64 ether);

        // 1. Create an operator set and register an operator
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);
        // Allocate all magnitude to the operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable_NoDelegatedStake(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 2. Fully slash operator
        SlashingParams memory slashParams = _genSlashing_Custom(operator, operatorSet, WAD - 1);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // 3. Delegate to an operator
        staker.delegateTo(operator);
        // delegate staker without any depositShares in beaconChainETHStrategy yet
        IStrategy[] memory emptyStrategies;
        uint[] memory emptyTokenBalances;
        check_Delegation_State(staker, operator, emptyStrategies, emptyTokenBalances);

        // 4. deposit/verify withdrawal credentials
        (validators, beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);
    }

    /**
     * Test sequence following _init()
     * 4. slash validators on beacon chain (start/complete checkpoint)
     * 5. queue withdrawal
     * 6. complete withdrawal
     * Given the operator has 1 magnitude remaining, the being slashed on the beacon chain
     * and calculating a non-WAD BCSF, their slashing factor should be rounded down to 0. Resulting in
     * the staker having 0 withdrawable shares.
     */
    function test_slashBC_startCompleteCP_queue_complete(uint24 _r) public rand(_r) {
        // 4. slash validators on beacon chain (start/complete checkpoint)
        uint40[] memory slashedValidators = _choose(validators);
        slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Minor);
        // ensure non zero amount slashed gwei so that we can test rounding down behavior
        cheats.assume(slashedGwei > 0);
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        // start and complete checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckPoint_WithSlashing_LowMagnitude_State(staker, slashedGwei);

        // 5. queue withdrawal
        (, uint[] memory withdrawShares) = _randWithdrawal(strategies, initDepositShares);

        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, withdrawShares);
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State({
            staker: staker,
            operator: operator,
            strategies: strategies,
            depositShares: withdrawShares, // amount of deposit shares to withdraw
            withdrawableShares: 0.toArrayU256(), // amount of withdrawable shares is 0 due to slashing factor being 0
            withdrawals: withdrawals,
            withdrawalRoots: roots
        });

        // Operator's maxMagnitude is 1 and staker was slashed to non-WAD BCSF. Therefore
        // staker should have been rounded down to 0
        uint slashingFactor = staker.getSlashingFactor(beaconChainETHStrategy);
        assertEq(slashingFactor, 0, "slashing factor should be rounded down to 0");

        // 6. complete withdrawal
        // only strategy is beaconChainETHStrategy
        _rollBlocksForCompleteWithdrawals(withdrawals);

        staker.completeWithdrawalAsShares(withdrawals[0]);
        check_Withdrawal_AsShares_State(staker, operator, withdrawals[0], strategies, 0.toArrayU256());
    }

    /**
     * Test sequence following _init()
     * 4. fully slash operator on opSet again to 0 magnitude
     * 5. undelegate
     * 6. redeposit (start/complete checkpoint)
     * This is testing when an operator is fully slashed with 0 magnitude, the staker can still undelegate
     * and "redeposit" to Eigenlayer.
     */
    function test_slash_undelegate_redeposit(uint24 _r) public rand(_r) {
        // 4. AVS slashes operator again to 0 magnitude and fully slashed
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        slashParams.wadsToSlash[0] = WAD;
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // 5. undelegate results in 0 delegated shares removed since operator has 0 magnitude and staker is fully slashed too
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 6. redeposit (start/complete checkpoint or verifyWC)
        if (_randBool()) {
            // Verify WC
            (validators, beaconBalanceGwei) = staker.startValidators(uint8(_randUint(3, 10)));
            beaconChain.advanceEpoch();

            staker.verifyWithdrawalCredentials(validators);
            check_VerifyWC_State(staker, validators, beaconBalanceGwei);
        } else {
            // Start/complete CP
            beaconChain.advanceEpoch();
            staker.startCheckpoint();
            check_StartCheckpoint_State(staker);

            staker.completeCheckpoint();
            check_CompleteCheckpoint_State(staker);
        }
    }

    /**
     * Test sequence following _init()
     * 4. fully slash operator on opSet again to 0 magnitude
     * 5. undelegate
     * 6. complete withdrawals as shares(although amount 0 from fully slashed operator)
     * 7. redeposit (start/complete checkpoint)
     * This is testing when an operator is fully slashed with 0 magnitude, the staker can still undelegate,
     * complete withdrawals as shares(0 shares though), and redeposit to Eigenlayer.
     */
    function test_slash_undelegate_completeAsShares_startCompleteCP(uint24 _r) public rand(_r) {
        // 4. AVS slashes operator again to 0 magnitude and fully slashed
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        slashParams.wadsToSlash[0] = WAD;
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // 5. undelegate results in 0 delegated shares removed since operator has 0 magnitude and staker is fully slashed too
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 6. complete withdrawals as shares(although amount 0 from fully slashed operator)
        _rollBlocksForCompleteWithdrawals(withdrawals);

        staker.completeWithdrawalAsShares(withdrawals[0]);
        check_Withdrawal_AsShares_State(staker, operator, withdrawals[0], strategies, 0.toArrayU256());

        // 7. redeposit (start/complete checkpoint)
        beaconChain.advanceEpoch();
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);
    }

    /**
     * Test sequence following _init()
     * 4. fully slash operator on opSet again to 0 magnitude
     * 5. undelegate
     * 6. complete withdrawals as tokens(although amount 0 from fully slashed operator)
     * 7. redeposit (start new validators and verifywc)
     * This is testing when an operator is fully slashed with 0 magnitude, the staker can still undelegate,
     * complete withdrawals as tokens(0 tokens though), and redeposit to Eigenlayer.
     */
    function test_slash_undelegate_completeAsTokens_verifyWC(uint24 _r) public rand(_r) {
        // 4. AVS slashes operator again to 0 magnitude and fully slashed
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        slashParams.wadsToSlash[0] = WAD;
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // 5. undelegate results in 0 delegated shares removed since operator has 0 magnitude and staker is fully slashed too
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 6. complete withdrawals as tokens(although amount 0 from fully slashed operator)
        // This also exits validators on the beacon chain
        _rollBlocksForCompleteWithdrawals(withdrawals);
        IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[0]);
        check_Withdrawal_AsTokens_State(staker, operator, withdrawals[0], strategies, 0.toArrayU256(), tokens, 0.toArrayU256());

        // 7. deposit/verify withdrawal credentials
        // randomly startup 1-10 validators
        (validators, beaconBalanceGwei) = staker.startValidators(uint8(_randUint(1, 10)));
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);
    }

    /**
     * Test sequence following _init()
     * 4. queue withdrawal
     * 5. complete withdrawal as tokens
     * This is testing a staker can queue a withdrawal and complete as tokens even
     * though the operator has 1 maxMagnitude
     */
    function test_queueAllShares_completeAsTokens(uint24 _r) public rand(_r) {
        // 4. queue withdrawal
        // ( , uint[] memory withdrawShares) = _randWithdrawal(strategies, initDepositShares);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State({
            staker: staker,
            operator: operator,
            strategies: strategies,
            depositShares: initDepositShares,
            withdrawableShares: initDepositShares,
            withdrawals: withdrawals,
            withdrawalRoots: roots
        });

        // 5. complete withdrawal as tokens
        // - exits validators
        // - advances epoch
        // - starts/completes checkpoint
        _rollBlocksForCompleteWithdrawals(withdrawals);
        IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[0]);
        check_Withdrawal_AsTokens_State(staker, operator, withdrawals[0], strategies, initDepositShares, tokens, initDepositShares);
    }
}
