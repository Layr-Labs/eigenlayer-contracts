// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

abstract contract IntegrationBase {

    Vm cheats = Vm(HEVM_ADDRESS);

    IDelegationManager delegationManager;
    IEigenPodManager eigenPodManager;
    IStrategyManager strategyManager;

    Global global;

    function setUp() virtual public {
        // Deploy contracts
        
        // Deploy strategies and assign to strategymgr

        // initialize contracts
    }
    
    /**
     * Gen/Init methods:
     */

    /// Creates a blank user object and sanity checks setup
    function _newRegisteredOperator() internal returns (IUser) {
        IUser operator = IUser(new User());

        operator.registerAsOperator();

        assertTrue(delegationManager.isOperator(operator), "_gen_Operator_Registered: operator should be registered");
        assert_HasNoDelegatableShares(operator, "new operator should not have delegatable shares");

        return IUser(operator);
    }

    /// This method is very unfinished, but this should combine with _configRand in Test.t.sol
    /// to generate a User object with the specific assets allowed in the test.
    function _newRandomStaker() internal returns (IUser) {
        UserType uType = UserType(_randUint({ min: type(UserType).min, max: type(UserType).max }));

        User user;
        if (uType == UserType.MIXED_ASSETS) {
            user = new 
        }
        
        IStrategy[] memory strategies = _selectRandomStrategies();
        uint[] memory tokenBalances = _awardRandomUnderlyingBalances(staker, strategies);

        assertEq(strategies.length, tokenBalances.length, "todo");
        assertTrue(strategies.length != 0, "todo");

        // Validate that staker has correct balance in each selected strategy's underlying token
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint balance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                assertEq(tokenBalance, address(staker).balance, "staker should hold correct amount of native ETH");
            } else {
                assertEq(
                    tokenBalance, strat.underlyingToken.balanceOf(address(staker)), 
                    "staker should hold correct amount of underlying token"
                );
            }
        }

        return staker;
    }

    /** 
     * Common assertions:
     */

    function assert_HasNoDelegatableShares(IUser user, string memory err) internal {
        (IStrategy[] memory strategies, uint[] memory shares) = 
            delegationManager.getDelegatableShares(user);
        
        assertEq(strategies.length, 0, err);
        assertEq(strategies.length, shares.length, "assert_HasNoDelegatableShares: return length mismatch");
    }

    function assert_HasNoUnderlyingTokenBalance(IUser user, IStrategy[] memory strategies, string memory err) internal {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint tokenBalance;
            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokenBalance = address(user).balance;
            } else {
                tokenBalance = strat.underlyingToken.balanceOf(user);
            }

            assertEq(0, tokenBalance, err);
        }
    }

    function assert_HasExpectedShares(IUser user, IStrategy[] memory strategies, uint[] memory expectedShares) internal {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint actualShares;
            if (strat == BEACONCHAIN_ETH_STRAT) {
                actualShares = eigenPodManager.podOwnerShares(user);
            } else {
                actualShares = strategyManager.stakerStrategyShares(user);
            }

            assertEq(expectedShares[i], actualShares, err);
        }
    }

    function assert_HasOperatorShares(IUser user, IStrategy[] memory strategies, uint[] memory expectedShares) internal {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint actualShares = delegationManager.operatorShares(user, strat);

            assertEq(expectedShares[i], actualShares, err);
        }
    }

    /**
     * Snapshot assertions combine Global's snapshots with assertions
     * that allow easy comparisons between prev/cur values
     */

    /// @dev Check that the operator has `addedShares` additional shares for each
    /// strategy since the last snapshot
    function assert_Snap_AddedOperatorShares(
        IUser operator, 
        IStrategy[] memory strategies, 
        uint[] memory addedShares,
        string memory err
    ) internal {
        // Put current state on stack and revert to prior state to fetch
        // previous operator shares
        uint curState = global.warpToLast();
        uint[] memory prevShares = _getOperatorShares(operator, strategies);
        global.warpToPresent(curState); // Snap back to reality

        uint[] memory curShares = _getOperatorShares(operator, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + addedShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the operator has `removedShares` prior shares for each
    /// strategy since the last snapshot
    function assert_Snap_RemovedOperatorShares(
        IUser operator, 
        IStrategy[] memory strategies, 
        uint[] memory removedShares,
        string memory err
    ) internal {
        // Put current state on stack and revert to prior state to fetch
        // previous operator shares
        uint curState = global.warpToLast();
        uint[] memory prevShares = _getOperatorShares(operator, strategies);
        global.warpToPresent(curState); // Snap back to reality

        uint[] memory curShares = _getOperatorShares(operator, strategies);

        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the staker has `addedShares` additional shares for each
    /// strategy since the last snapshot
    function assert_Snap_AddedStakerShares(
        IUser staker, 
        IStrategy[] memory strategies, 
        uint[] memory addedShares,
        string memory err
    ) internal {
        // Put current state on stack and revert to prior state to fetch
        // previous staker shares
        uint curState = global.warpToLast();
        uint[] memory prevShares = _getStakerShares(staker, strategies);
        global.warpToPresent(curState); // Snap back to reality

        uint[] memory curShares = _getStakerShares(staker, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + addedShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the staker has `removedShares` prior shares for each
    /// strategy since the last snapshot
    function assert_Snap_RemovedStakerShares(
        IUser staker, 
        IStrategy[] memory strategies, 
        uint[] memory removedShares,
        string memory err
    ) internal {
        // Put current state on stack and revert to prior state to fetch
        // previous staker shares
        uint curState = global.warpToLast();
        uint[] memory prevShares = _getStakerShares(staker, strategies);
        global.warpToPresent(curState); // Snap back to reality

        uint[] memory curShares = _getStakerShares(staker, strategies);

        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    /**
     * Helpful getters:
     */

    /// @dev For some strategies/underlying token balances, calculate the expected shares received
    /// from depositing all tokens
    function _calculateExpectedShares(IStrategy[] memory strategies, uint[] memory tokenBalances) internal view returns (uint[] memory) {
        uint[] memory expectedShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint tokenBalance = tokenBalances[i];
            if (strat == BEACONCHAIN_ETH_STRAT) {
                // TODO - need to calculate this
                // expectedShares[i] = eigenPodManager.underlyingToShares(tokenBalance);
            } else {
                expectedShares[i] = strat.underlyingToShares(tokenBalance);
            }
        }

        return expectedShares;
    }

    /// @dev Looks up each strategy and returns a list of the operator's shares
    function _getOperatorShares(IUser operator, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory curShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            curShares[i] = delegationManager.operatorShares(operator, strategies[i]);
        }

        return curShares;
    }

    /// @dev Looks up each strategy and returns a list of the staker's shares
    function _getStakerShares(IUser staker, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory curShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                curShares[i] = eigenPodManager.podOwnerShares(staker);
            } else {
                curShares[i] = strategyManager.stakerStrategyShares(staker, strat);
            }
        }

        return curShares;
    }
}