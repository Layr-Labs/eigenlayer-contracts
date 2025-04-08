// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_EigenPod_Base is UpgradeTest {
    User staker;
    IStrategy[] strategies;
    uint[] tokenBalances;
    uint[] shares;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_ETH);
        _configUserTypes(DEFAULT);

        /// 0. Create a staker with underlying assets
        (staker, strategies, tokenBalances) = _newRandomStaker();
        shares = _calculateExpectedShares(strategies, tokenBalances);

        ///  1. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
    }
}

contract Integration_Upgrade_EigenPod_SlashAfterUpgrade is Integration_Upgrade_EigenPod_Base {
    uint40[] validators;
    uint64 slashedGwei;

    function testFuzz_deposit_upgrade_slash_completeCheckpoint(uint24 _rand) public rand(_rand) {
        uint64 initBeaconBalanceGwei = uint64(tokenBalances[0] / GWEI_TO_WEI);

        /// 2. Upgrade contracts
        _upgradeEigenLayerContracts();

        /// 3. Slash the staker partially
        validators = staker.getActiveValidators();
        slashedGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards(); // Withdraw slashed validators to pod

        // 4. Start and complete a checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_WithPodBalance_State(staker, initBeaconBalanceGwei - slashedGwei);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(staker, validators, slashedGwei);
    }
}

contract Integration_Upgrade_EigenPod_FullSlash is Integration_Upgrade_EigenPod_Base {
    uint40[] validators;
    uint64 slashedGwei;

    function _init() internal override {
        super._init();

        /// 2. Fully slash the staker
        validators = staker.getActiveValidators();
        slashedGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Full);
        beaconChain.advanceEpoch_NoRewards(); // Withdraw slashed validators to pod

        /// 3. Upgrade contracts
        _upgradeEigenLayerContracts();
    }

    function testFuzz_deposit_fullSlash_upgrade_delegate(uint24 _rand) public rand(_rand) {
        /// 4. Delegate to operator
        (User operator,,) = _newRandomOperator();
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);
    }

    function testFuzz_deposit_fullSlash_upgrade_deposit_delegate(uint24 _rand) public rand(_rand) {
        // 5. Start a new validator & verify withdrawal credentials
        cheats.deal(address(staker), 32 ether);
        tokenBalances[0] = tokenBalances[0] + 32 ether;
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);
        shares = _calculateExpectedShares(strategies, tokenBalances);

        // 6. Delegate to operator
        (User operator,,) = _newRandomOperator();
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);
    }
}

contract Integration_Upgrade_EigenPod_NegativeShares is Integration_Upgrade_EigenPod_Base {
    User operator;
    Withdrawal withdrawal;
    int[] tokenDeltas;
    int[] balanceUpdateShareDelta;

    function _init() internal override {
        super._init();

        // 3. Delegate to operator
        (operator,,) = _newRandomOperator();
        staker.delegateTo(operator);

        /// 4. Queue a withdrawal for all shares
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawal = withdrawals[0];

        /// 5. Update balance randomly (can be positive or negative)
        (tokenDeltas, balanceUpdateShareDelta,) = _randBalanceUpdate(staker, strategies);
        staker.updateBalances(strategies, tokenDeltas);

        /// 6. Upgrade contracts
        _upgradeEigenLayerContracts();
    }

    function testFuzz_deposit_delegate_updateBalance_upgrade_completeAsShares(uint24 _rand) public rand(_rand) {
        /// 7. Complete the withdrawal as shares
        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = withdrawal;
        _rollBlocksForCompleteWithdrawals(withdrawals);
        staker.completeWithdrawalAsShares(withdrawal);

        // Manually complete checks since we could still negative shares prior to the upgrade, causing a revert in the share check
        (uint[] memory expectedOperatorShareDelta, int[] memory expectedStakerShareDelta) =
            _getPostWithdrawalExpectedShareDeltas(balanceUpdateShareDelta[0], withdrawal);
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
        assert_Snap_Added_OperatorShares(
            operator, withdrawal.strategies, expectedOperatorShareDelta, "operator should have received shares"
        );
        assert_Snap_Delta_StakerShares(staker, strategies, expectedStakerShareDelta, "staker should have received expected shares");
    }

    function testFuzz_deposit_delegate_updateBalance_upgrade_completeAsTokens(uint24 _rand) public rand(_rand) {
        /// 7. Complete the withdrawal as tokens
        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = withdrawal;
        _rollBlocksForCompleteWithdrawals(withdrawals);
        IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawal);
        uint[] memory expectedTokens = _getPostWithdrawalExpectedTokenDeltas(balanceUpdateShareDelta[0], withdrawal);

        // Manually complete checks since we could still negative shares prior to the upgrade, causing a revert in the share check
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Added_TokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
        assert_Snap_Unchanged_OperatorShares(operator, "operator shares should not have changed");

        // If we had a positive balance update, then the staker shares should not have changed
        if (balanceUpdateShareDelta[0] > 0) {
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should not have changed");
        }
        // Else, the staker shares should have increased by the magnitude of the negative share delta
        else {
            int[] memory expectedStakerShareDelta = new int[](1);
            expectedStakerShareDelta[0] = -balanceUpdateShareDelta[0];
            assert_Snap_Delta_StakerShares(staker, strategies, expectedStakerShareDelta, "staker should have received expected shares");
        }
    }

    function _getPostWithdrawalExpectedShareDeltas(int _balanceUpdateShareDelta, Withdrawal memory _withdrawal)
        internal
        pure
        returns (uint[] memory, int[] memory)
    {
        uint[] memory operatorShareDelta = new uint[](1);
        int[] memory stakerShareDelta = new int[](1);
        // The staker share delta is the withdrawal scaled shares since it can go from negative to positive
        stakerShareDelta[0] = int(_withdrawal.scaledShares[0]);

        if (_balanceUpdateShareDelta > 0) {
            // If balanceUpdateShareDelta is positive, then the operator delta is the withdrawal scaled shares
            operatorShareDelta[0] = _withdrawal.scaledShares[0];
        } else {
            // Operator shares never went negative, so we can just add the withdrawal scaled shares and the negative share delta
            operatorShareDelta[0] = uint(int(_withdrawal.scaledShares[0]) + _balanceUpdateShareDelta);
        }

        return (operatorShareDelta, stakerShareDelta);
    }

    function _getPostWithdrawalExpectedTokenDeltas(int _balanceUpdateShareDelta, Withdrawal memory _withdrawal)
        internal
        pure
        returns (uint[] memory)
    {
        uint[] memory expectedTokenDeltas = new uint[](1);
        if (_balanceUpdateShareDelta > 0) {
            // If we had a positive balance update, then the expected token delta is the withdrawal scaled shares
            expectedTokenDeltas[0] = _withdrawal.scaledShares[0];
        } else {
            // If we had a negative balance update, then the expected token delta is the withdrawal scaled shares plus the negative share delta
            expectedTokenDeltas[0] = uint(int(_withdrawal.scaledShares[0]) + _balanceUpdateShareDelta);
        }
        return expectedTokenDeltas;
    }
}
