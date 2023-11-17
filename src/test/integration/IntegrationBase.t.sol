// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import "src/test/integration/IntegrationDeployer.t.sol";
import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/User.t.sol";

abstract contract IntegrationBase is IntegrationDeployer {

    /**
     * Gen/Init methods:
     */

    /**
     * @dev Create a new user according to configured random variants.
     * This user is ready to deposit into some strategies and has some underlying token balances
     */
    function _newRandomStaker() internal returns (User, IStrategy[] memory, uint[] memory) {
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _randUser();
        
        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "_newStaker: failed to award token balances");

        return (staker, strategies, tokenBalances);
    }

    function _newRandomOperator() internal returns (User, IStrategy[] memory, uint[] memory) {
        (User operator, IStrategy[] memory strategies, uint[] memory tokenBalances) = _randUser();
        
        operator.registerAsOperator();
        operator.depositIntoEigenlayer(strategies, tokenBalances);

        assert_Snap_AddedStakerShares(operator, strategies, tokenBalances, "_newOperator: failed to add delegatable shares");
        assert_Snap_AddedOperatorShares(operator, strategies, tokenBalances, "_newOperator: failed to award shares to operator");
        assertTrue(delegationManager.isOperator(address(operator)), "_newOperator: operator should be registered");

        return (operator, strategies, tokenBalances);
    }

    /** 
     * Common assertions:
     */

    function assert_HasNoDelegatableShares(User user, string memory err) internal {
        (IStrategy[] memory strategies, uint[] memory shares) = 
            delegationManager.getDelegatableShares(address(user));
        
        assertEq(strategies.length, 0, err);
        assertEq(strategies.length, shares.length, "assert_HasNoDelegatableShares: return length mismatch");
    }

    function assert_HasUnderlyingTokenBalances(
        User user, 
        IStrategy[] memory strategies, 
        uint[] memory expectedBalances, 
        string memory err
    ) internal {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            
            uint expectedBalance = expectedBalances[i];
            uint tokenBalance;
            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokenBalance = address(user).balance;
            } else {
                tokenBalance = strat.underlyingToken().balanceOf(address(user));
            }

            assertEq(expectedBalance, tokenBalance, err);
        }
    }

    function assert_HasNoUnderlyingTokenBalance(User user, IStrategy[] memory strategies, string memory err) internal {
        assert_HasUnderlyingTokenBalances(user, strategies, new uint[](strategies.length), err);
    }

    function assert_HasExpectedShares(
        User user, 
        IStrategy[] memory strategies, 
        uint[] memory expectedShares, 
        string memory err
    ) internal {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint actualShares;
            if (strat == BEACONCHAIN_ETH_STRAT) {
                // TODO
                // actualShares = eigenPodManager.podOwnerShares(address(user));
                revert("unimplemented");
            } else {
                actualShares = strategyManager.stakerStrategyShares(address(user), strat);
            }

            assertEq(expectedShares[i], actualShares, err);
        }
    }

    function assert_HasOperatorShares(
        User user, 
        IStrategy[] memory strategies, 
        uint[] memory expectedShares, 
        string memory err
    ) internal {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint actualShares = delegationManager.operatorShares(address(user), strat);

            assertEq(expectedShares[i], actualShares, err);
        }
    }

    function assert_AllWithdrawalsPending(bytes32[] memory withdrawalRoots, string memory err) internal {
        for (uint i = 0; i < withdrawalRoots.length; i++) {
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoots[i]), err);
        }
    }

    function assert_ValidWithdrawalHashes(
        IDelegationManager.Withdrawal[] memory withdrawals,
        bytes32[] memory withdrawalRoots,
        string memory err
    ) internal {
        for (uint i = 0; i < withdrawals.length; i++) {
            assertEq(withdrawalRoots[i], delegationManager.calculateWithdrawalRoot(withdrawals[i]), err);
        }
    }
    
    /**
     * Snapshot assertions combine Timemachine's snapshots with assertions
     * that allow easy comparisons between prev/cur values
     */

    /// @dev Check that the operator has `addedShares` additional shares for each
    /// strategy since the last snapshot
    function assert_Snap_AddedOperatorShares(
        User operator, 
        IStrategy[] memory strategies, 
        uint[] memory addedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + addedShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the operator has `removedShares` prior shares for each
    /// strategy since the last snapshot
    function assert_Snap_RemovedOperatorShares(
        User operator, 
        IStrategy[] memory strategies, 
        uint[] memory removedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);

        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the staker has `addedShares` additional shares for each
    /// strategy since the last snapshot
    function assert_Snap_AddedStakerShares(
        User staker, 
        IStrategy[] memory strategies, 
        uint[] memory addedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerShares(staker, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + addedShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the staker has `removedShares` prior shares for each
    /// strategy since the last snapshot
    function assert_Snap_RemovedStakerShares(
        User staker, 
        IStrategy[] memory strategies, 
        uint[] memory removedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerShares(staker, strategies);

        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    function assert_Snap_IncreasedQueuedWithdrawals(
        User staker, 
        IDelegationManager.Withdrawal[] memory withdrawals,
        string memory err
    ) internal {
        uint curQueuedWithdrawals = _getCumulativeWithdrawals(staker);
        // Use timewarp to get previous cumulative withdrawals
        uint prevQueuedWithdrawals = _getPrevCumulativeWithdrawals(staker);

        assertEq(prevQueuedWithdrawals + withdrawals.length, curQueuedWithdrawals, err);
    }

    function assert_Snap_IncreasedTokenBalances(
        User staker,
        IERC20[] memory tokens,
        uint[] memory addedTokens,
        string memory err
    ) internal {
        uint[] memory curTokenBalances = _getTokenBalances(staker, tokens);
        // Use timewarp to get previous token balances
        uint[] memory prevTokenBalances = _getPrevTokenBalances(staker, tokens);

        for (uint i = 0; i < tokens.length; i++) {
            uint prevBalance = prevTokenBalances[i];
            uint curBalance = curTokenBalances[i];

            assertEq(prevBalance + addedTokens[i], curBalance, err);
        }
    }

    /**
     * Helpful getters:
     */

    /// @dev For some strategies/underlying token balances, calculate the expected shares received
    /// from depositing all tokens
    function _calculateExpectedShares(IStrategy[] memory strategies, uint[] memory tokenBalances) internal returns (uint[] memory) {
        uint[] memory expectedShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint tokenBalance = tokenBalances[i];
            if (strat == BEACONCHAIN_ETH_STRAT) {
                // TODO - need to calculate this
                // expectedShares[i] = eigenPodManager.underlyingToShares(tokenBalance);
                revert("_calculateExpectedShares: unimplemented for native eth");
            } else {
                expectedShares[i] = strat.underlyingToShares(tokenBalance);
            }
        }

        return expectedShares;
    }

    /// @dev For some strategies/underlying token balances, calculate the expected shares received
    /// from depositing all tokens
    function _calculateExpectedTokens(IStrategy[] memory strategies, uint[] memory shares) internal returns (uint[] memory) {
        uint[] memory expectedTokens = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // TODO - need to calculate this
                // expectedTokens[i] = eigenPodManager.underlyingToShares(tokenBalance);
                revert("_calculateExpectedShares: unimplemented for native eth");
            } else {
                expectedTokens[i] = strat.sharesToUnderlying(shares[i]);
            }
        }

        return expectedTokens;
    }

    modifier timewarp() {
        uint curState = timeMachine.warpToLast();
        _;
        timeMachine.warpToPresent(curState);
    }

    /// @dev Uses timewarp modifier to get operator shares at the last snapshot
    function _getPrevOperatorShares(
        User operator, 
        IStrategy[] memory strategies
    ) internal timewarp() returns (uint[] memory) {
        return _getOperatorShares(operator, strategies);
    }

    /// @dev Looks up each strategy and returns a list of the operator's shares
    function _getOperatorShares(User operator, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory curShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            curShares[i] = delegationManager.operatorShares(address(operator), strategies[i]);
        }

        return curShares;
    }

    /// @dev Uses timewarp modifier to get staker shares at the last snapshot
    function _getPrevStakerShares(
        User staker, 
        IStrategy[] memory strategies
    ) internal timewarp() returns (uint[] memory) {
        return _getStakerShares(staker, strategies);
    }

    /// @dev Looks up each strategy and returns a list of the staker's shares
    function _getStakerShares(User staker, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory curShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // curShares[i] = eigenPodManager.podOwnerShares(address(staker));
                revert("TODO: unimplemented");
            } else {
                curShares[i] = strategyManager.stakerStrategyShares(address(staker), strat);
            }
        }

        return curShares;
    }

    function _getPrevCumulativeWithdrawals(User staker) internal timewarp() returns (uint) {
        return _getCumulativeWithdrawals(staker);
    }

    function _getCumulativeWithdrawals(User staker) internal view returns (uint) {
        return delegationManager.cumulativeWithdrawalsQueued(address(staker));
    }

    function _getPrevTokenBalances(User staker, IERC20[] memory tokens) internal timewarp() returns (uint[] memory) {
        return _getTokenBalances(staker, tokens);
    }

    function _getTokenBalances(User staker, IERC20[] memory tokens) internal view returns (uint[] memory) {
        uint[] memory balances = new uint[](tokens.length);

        for (uint i = 0; i < tokens.length; i++) {
            balances[i] = tokens[i].balanceOf(address(staker));
        }

        return balances;
    }
}