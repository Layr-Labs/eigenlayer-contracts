// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Slashing_Combined is IntegrationCheckUtils {
    // Helper struct to reduce stack variables
    struct TestContext {
        User staker;
        User operator;
        AVS avs;
        IStrategy[] strategies;
        uint256[] tokenBalances;
        uint40[] validators;
        OperatorSet operatorSet;
        IAllocationManagerTypes.AllocateParams allocateParams;
        IAllocationManagerTypes.SlashingParams slashingParams;
    }

    function testFuzz_deposit_slashBeacon_delegate_slashEigen_withdraw(uint24 _random) public {
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
        (ctx.avs,) = _newRandomAVS();

        // Handle validator setup and slashing
        _handleValidatorSetupAndSlashing(ctx);
        
        // Handle delegation and operator setup
        _handleDelegationAndOperatorSetup(ctx);
        
        // Handle EigenLayer slashing
        _handleEigenLayerSlashing(ctx);
        
        // Handle withdrawal
        _handleWithdrawal(ctx);
    }

    function _handleValidatorSetupAndSlashing(TestContext memory ctx) internal {
        // Create and verify validators
        (ctx.validators,) = ctx.staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        ctx.staker.verifyWithdrawalCredentials(ctx.validators);
        
        // Slash validators and record via checkpoint
        uint40[] memory slashedValidators = _chooseSubset(ctx.validators);
        uint64 slashedGwei = beaconChain.slashValidators(slashedValidators);
        beaconChain.advanceEpoch_NoRewards();
        
        ctx.staker.startCheckpoint();
        console.log("Active validator count after starting checkpoint:", ctx.staker.pod().activeValidatorCount());
        ctx.staker.completeCheckpoint();
        console.log("Active validator count after completing checkpoint:", ctx.staker.pod().activeValidatorCount());
        check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(ctx.staker, slashedValidators, slashedGwei);
    }

    function _handleDelegationAndOperatorSetup(TestContext memory ctx) internal {
        // Handle delegation
        ctx.staker.delegateTo(ctx.operator);
        check_Delegation_State(
            ctx.staker, 
            ctx.operator, 
            ctx.strategies, 
            _getStakerDepositShares(ctx.staker, ctx.strategies)
        );

        // Setup operator set
        ctx.operatorSet = ctx.avs.createOperatorSet(ctx.strategies);
        ctx.operator.registerForOperatorSet(ctx.operatorSet);
        
        ctx.allocateParams = ctx.operator.modifyAllocations(
            ctx.operatorSet,
            _randMagnitudes({sum: 1 ether, len: ctx.strategies.length})
        );
        _rollBlocksForCompleteAllocation(ctx.operator, ctx.operatorSet, ctx.strategies);
    }

    function _handleEigenLayerSlashing(TestContext memory ctx) internal {
        (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) = 
            _randStrategiesAndWadsToSlash(ctx.operatorSet);
            
        ctx.slashingParams = ctx.avs.slashOperator(
            ctx.operator, 
            ctx.operatorSet.id, 
            strategiesToSlash, 
            wadsToSlash
        );
        
        assert_Snap_Allocations_Slashed(
            ctx.slashingParams, 
            ctx.operatorSet, 
            true, 
            "operator allocations should be slashed"
        );
        assert_Snap_Unchanged_StakerDepositShares(
            ctx.staker, 
            "staker deposit shares should be unchanged after slashing"
        );
        assert_Snap_StakerWithdrawableShares_AfterSlash(
            ctx.staker, 
            ctx.allocateParams, 
            ctx.slashingParams, 
            "staker deposit shares should be slashed"
        );
    }

    function _handleWithdrawal(TestContext memory ctx) internal {
        EigenPod pod = ctx.staker.pod();
        console.log("Before checkpoint:");
        console.log("Restaked execution layer gwei:", pod.withdrawableRestakedExecutionLayerGwei());
        console.log("Total shares:", eigenPodManager.stakerDepositShares(address(ctx.staker), beaconChainETHStrategy));
        
        // Advance beacon chain epoch
        beaconChain.advanceEpoch_NoRewards();
        
        // checkpoint 
        //ctx.staker.startCheckpoint();
        //ctx.staker.completeCheckpoint();

        console.log("After checkpoint:");
        console.log("Restaked execution layer gwei:", pod.withdrawableRestakedExecutionLayerGwei());
        console.log("Total shares:", eigenPodManager.stakerDepositShares(address(ctx.staker), beaconChainETHStrategy));

        // Queue withdrawal
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = ctx.staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        console.log("After undelegate:");
        console.log("withdrawals length:", withdrawals.length);
        for (uint i = 0; i < withdrawals.length; i++) {
            console.log("Is pending?", delegationManager.pendingWithdrawals(withdrawalRoots[i]));
        }

        // Complete withdrawals after delay
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint256 i = 0; i < withdrawals.length; i++) {
            console.log("Attempting to complete withdrawal", i);
            console.log("Is pending before completion?", delegationManager.pendingWithdrawals(withdrawalRoots[i]));
            
            uint[] memory expectedTokens = _calculateExpectedTokens(
                withdrawals[i].strategies,
                withdrawals[i].scaledShares
            );
            
            for (uint256 i = 0; i < expectedTokens.length; i++) {
                console.log(expectedTokens[i]);
            }
            IERC20[] memory tokens = ctx.staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State_AfterSlash(
                ctx.staker,
                ctx.operator,
                withdrawals[i],
                ctx.allocateParams,
                ctx.slashingParams,
                expectedTokens
            );
        }

        // Final state checks
        assert_HasNoDelegatableShares(ctx.staker, "staker should have withdrawn all shares");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be completed");
    }

}