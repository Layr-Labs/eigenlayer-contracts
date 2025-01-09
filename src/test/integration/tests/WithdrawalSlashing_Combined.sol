// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_WithdrawalSlashing_Combined is IntegrationCheckUtils {
    // Helper struct to reduce stack variables
    struct TestContext {
        User staker;
        User operator;
        IStrategy[] strategies;
        uint256[] tokenBalances;
        uint40[] validators;
        bytes32[] withdrawalRoots; 
        IDelegationManagerTypes.Withdrawal[] withdrawals;
        uint64 slashedGwei;
    }
    function testFuzz_deposit_delegate_queueWithdrawal_slashBeacon_checkpoint(uint24 _random) public {
        TestContext memory ctx;
        
        // Initial setup and configuration
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_ETH,
            _userTypes: DEFAULT
        });
        
        _upgradeEigenLayerContracts();

        // Initialize actors and store in context
        (ctx.staker, ctx.strategies, ctx.tokenBalances) = _newRandomStaker();
        (ctx.operator,,) = _newRandomOperator();

        // Handle validator setup and delegation
        _handleValidatorSetupAndDelegation(ctx);

        // Queue withdrawal before slashing
        _handleQueueWithdrawal(ctx);
        
        // Execute slashing while withdrawal is in queue
        _handleBeaconChainSlashing(ctx);

        // Start a checkpoint to reflect slashing
        _handlePostSlashingCheckpoint(ctx);

        // Complete the withdrawal and verify slashing was applied
        _handleWithdrawalCompletion(ctx);
    }

    function _handleValidatorSetupAndDelegation(TestContext memory ctx) internal {
        // Create and verify validators
        (ctx.validators,) = ctx.staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        ctx.staker.verifyWithdrawalCredentials(ctx.validators);

        // Delegate to operator
        ctx.staker.delegateTo(ctx.operator);
        check_Delegation_State(
            ctx.staker,
            ctx.operator,
            ctx.strategies,
            _getStakerDepositShares(ctx.staker, ctx.strategies)
        );
    }

    function _handleQueueWithdrawal(TestContext memory ctx) internal {
        // Queue withdrawal by undelegating
        ctx.withdrawals = ctx.staker.undelegate();
        ctx.withdrawalRoots = _getWithdrawalHashes(ctx.withdrawals);
        
        // Verify withdrawal state
        assert_AllWithdrawalsPending(ctx.withdrawalRoots, "withdrawals should be pending");
        assert_ValidWithdrawalHashes(ctx.withdrawals, ctx.withdrawalRoots, "withdrawal hashes should be valid");
    }

    function _handleBeaconChainSlashing(TestContext memory ctx) internal {
        // Choose subset of validators to slash
        uint40[] memory slashedValidators = _chooseSubset(ctx.validators);
        
        // Execute slashing on beacon chain 
        ctx.slashedGwei = beaconChain.slashValidators(slashedValidators);
        beaconChain.advanceEpoch_NoRewards();

        console.log("Slashed amount (gwei)", ctx.slashedGwei);
    }

    function _handlePostSlashingCheckpoint(TestContext memory ctx) internal {
        // Start and complete checkpoint to reflect slashing
        ctx.staker.startCheckpoint();
        ctx.staker.completeCheckpoint();
        console.log("Active validator count after completing checkpoint:", ctx.staker.pod().activeValidatorCount());

    }

    function _handleWithdrawalCompletion(TestContext memory ctx) internal {
        // Advance blocks to complete withdrawal
        _rollBlocksForCompleteWithdrawals(ctx.withdrawals);

        // Complete each withdrawal and verify state
        for (uint i = 0; i < ctx.withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(
                ctx.withdrawals[i].strategies,
                ctx.withdrawals[i].scaledShares
            );

            IERC20[] memory tokens = ctx.staker.completeWithdrawalAsTokens(ctx.withdrawals[i]);
            
            check_Withdrawal_AsTokens_State(
                ctx.staker,
                ctx.operator,
                ctx.withdrawals[i],
                ctx.withdrawals[i].strategies,
                ctx.withdrawals[i].scaledShares,
                tokens,
                expectedTokens
            );
        }

        // Final checks
        assert_HasNoDelegatableShares(ctx.staker, "staker should have no shares after withdrawal");
        assert_NoWithdrawalsPending(ctx.withdrawalRoots, "all withdrawals should be completed");
    }
}