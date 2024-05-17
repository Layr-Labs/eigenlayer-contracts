// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";

import "src/test/integration/IntegrationDeployer.t.sol";
import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";

abstract contract IntegrationBase is IntegrationDeployer {

    using Strings for *;

    uint numStakers = 0;
    uint numOperators = 0;

    // Lists of stakers/operators created before the m2 upgrade
    //
    // When we call _upgradeEigenLayerContracts, we iterate over
    // these lists and migrate perform the standard migration actions
    // for each user
    User[] stakersToMigrate;
    User[] operatorsToMigrate;

    /**
     * Gen/Init methods:
     */

    /**
     * @dev Create a new user according to configured random variants.
     * This user is ready to deposit into some strategies and has some underlying token balances
     */
    function _newRandomStaker() internal returns (User, IStrategy[] memory, uint[] memory) {
        string memory stakerName;

        User staker;
        IStrategy[] memory strategies;
        uint[] memory tokenBalances;

        if (forkType == MAINNET && !isUpgraded) {
            stakerName = string.concat("- M1_Staker", numStakers.toString());

            (staker, strategies, tokenBalances) = _randUser(stakerName);

            stakersToMigrate.push(staker);
        } else {
            stakerName = string.concat("- Staker", numStakers.toString());

            (staker, strategies, tokenBalances) = _randUser(stakerName);
        }

        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "_newRandomStaker: failed to award token balances");

        numStakers++;
        return (staker, strategies, tokenBalances);
    }

    /**
     * @dev Create a new operator according to configured random variants.
     * This user will immediately deposit their randomized assets into eigenlayer.
     * @notice If forktype is mainnet and not upgraded, then the operator will only randomize LSTs assets and deposit them
     * as ETH podowner shares are not available yet. 
     */
    function _newRandomOperator() internal returns (User, IStrategy[] memory, uint[] memory) {
        User operator;
        IStrategy[] memory strategies;
        uint[] memory tokenBalances;

        if (forkType == MAINNET && !isUpgraded) {
            string memory operatorName = string.concat("- M1_Operator", numOperators.toString());

            // Create an operator for M1. We omit native ETH because we want to
            // check staker/operator shares, and we don't award native ETH shares in M1
            (operator, strategies, tokenBalances) = _randUser_NoETH(operatorName);

            User_M1(payable(address(operator))).depositIntoEigenlayer_M1(strategies, tokenBalances);
            uint[] memory addedShares = _calculateExpectedShares(strategies, tokenBalances);

            assert_Snap_Added_StakerShares(operator, strategies, addedShares, "_newRandomOperator: failed to add delegatable shares");

            operatorsToMigrate.push(operator);
        } else {
            string memory operatorName = string.concat("- Operator", numOperators.toString());

            (operator, strategies, tokenBalances) = _randUser(operatorName);

            uint[] memory addedShares = _calculateExpectedShares(strategies, tokenBalances);

            operator.registerAsOperator();
            operator.depositIntoEigenlayer(strategies, tokenBalances);

            assert_Snap_Added_StakerShares(operator, strategies, addedShares, "_newRandomOperator: failed to add delegatable shares");
            assert_Snap_Added_OperatorShares(operator, strategies, addedShares, "_newRandomOperator: failed to award shares to operator");
            assertTrue(delegationManager.isOperator(address(operator)), "_newRandomOperator: operator should be registered");
        }

        numOperators++;
        return (operator, strategies, tokenBalances);
    }

    /// @dev If we're on mainnet, upgrade contracts to M2 and migrate stakers/operators
    function _upgradeEigenLayerContracts() internal {
        if (forkType == MAINNET) {
            require(!isUpgraded, "_upgradeEigenLayerContracts: already performed m2 upgrade");

            emit log("_upgradeEigenLayerContracts: upgrading mainnet to m2");
            _upgradeMainnetContracts();

            emit log("===Migrating Stakers/Operators===");

            // Enable restaking for stakers' pods
            for (uint i = 0; i < stakersToMigrate.length; i++) {
                stakersToMigrate[i].activateRestaking();
            }

            // Register operators with DelegationManager
            for (uint i = 0; i < operatorsToMigrate.length; i++) {
                operatorsToMigrate[i].registerAsOperator();
            }

            emit log("======");

            // Bump block.timestamp forward to allow verifyWC proofs for migrated pods
            emit log("advancing block time to start of next epoch:");

            uint64 curEpoch = _timeStampToEpoch(uint64(block.timestamp));
            uint64 nextEpochStartTime = _nextEpochStartTimestamp(curEpoch);

            emit log_named_uint("- current time", block.timestamp);
            emit log_named_uint("- current epoch", curEpoch);
            emit log_named_uint("- next epoch starts at", nextEpochStartTime);

            cheats.warp(nextEpochStartTime);
            beaconChain.setNextTimestamp(nextEpochStartTime);

            emit log("======");

            isUpgraded = true;
            emit log("_upgradeEigenLayerContracts: m2 upgrade complete");
        } else if (forkType == HOLESKY) {
            require(!isUpgraded, "_upgradeEigenLayerContracts: already performed m2 upgrade");

            emit log("_upgradeEigenLayerContracts: upgrading holesky to m2");
            _upgradeHoleskyContracts();

            isUpgraded = true;
            emit log("_upgradeEigenLayerContracts: m2 upgrade complete");
        }
    }

    /// From EigenPod.sol, using genesis time configured in ExistingDeploymentParser
    function _timeStampToEpoch(uint64 timestamp) internal view returns (uint64) {
        require(timestamp >= EIGENPOD_GENESIS_TIME, "EigenPod._timestampToEpoch: timestamp is before genesis");
        return (timestamp - EIGENPOD_GENESIS_TIME) / BeaconChainProofs.SECONDS_PER_EPOCH;
    }

    /// From EigenPod.sol, using genesis time configured in ExistingDeploymentParser
    function _nextEpochStartTimestamp(uint64 epoch) internal view returns (uint64) {
        return  
            EIGENPOD_GENESIS_TIME + ((1 + epoch) * BeaconChainProofs.SECONDS_PER_EPOCH);
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

            assertApproxEqAbs(expectedBalance, tokenBalance, 1, err);
            // assertEq(expectedBalance, tokenBalance, err);
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
                // This method should only be used for tests that handle positive
                // balances. Negative balances are an edge case that require
                // the own tests and helper methods.
                int shares = eigenPodManager.podOwnerShares(address(user));
                if (shares < 0) {
                    revert("assert_HasExpectedShares: negative shares");
                }

                actualShares = uint(shares);
            } else {
                actualShares = strategyManager.stakerStrategyShares(address(user), strat);
            }

            assertApproxEqAbs(expectedShares[i], actualShares, 1, err);
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

            assertApproxEqAbs(expectedShares[i], actualShares, 1, err);
        }
    }

    /// @dev Asserts that ALL of the `withdrawalRoots` is in `delegationManager.pendingWithdrawals`
    function assert_AllWithdrawalsPending(bytes32[] memory withdrawalRoots, string memory err) internal {
        for (uint i = 0; i < withdrawalRoots.length; i++) {
            assert_WithdrawalPending(withdrawalRoots[i], err);
        }
    }

    /// @dev Asserts that NONE of the `withdrawalRoots` is in `delegationManager.pendingWithdrawals`
    function assert_NoWithdrawalsPending(bytes32[] memory withdrawalRoots, string memory err) internal {
        for (uint i = 0; i < withdrawalRoots.length; i++) {
            assert_WithdrawalNotPending(withdrawalRoots[i], err);
        }
    }

    /// @dev Asserts that the hash of each withdrawal corresponds to the provided withdrawal root
    function assert_WithdrawalPending(bytes32 withdrawalRoot, string memory err) internal {
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), err);
    }

    function assert_WithdrawalNotPending(bytes32 withdrawalRoot, string memory err) internal {
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), err);
    }

    function assert_ValidWithdrawalHashes(
        IDelegationManager.Withdrawal[] memory withdrawals,
        bytes32[] memory withdrawalRoots,
        string memory err
    ) internal {
        for (uint i = 0; i < withdrawals.length; i++) {
            assert_ValidWithdrawalHash(withdrawals[i], withdrawalRoots[i], err);
        }
    }

    function assert_ValidWithdrawalHash(
        IDelegationManager.Withdrawal memory withdrawal,
        bytes32 withdrawalRoot,
        string memory err
    ) internal {
        assertEq(withdrawalRoot, delegationManager.calculateWithdrawalRoot(withdrawal), err);
    }
    
    /*******************************************************************************
                                SNAPSHOT ASSERTIONS
                       TIME TRAVELERS ONLY BEYOND THIS POINT
    *******************************************************************************/

    /// Snapshot assertions for delegationManager.operatorShares:

    /// @dev Check that the operator has `addedShares` additional operator shares 
    // for each strategy since the last snapshot
    function assert_Snap_Added_OperatorShares(
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
            assertApproxEqAbs(prevShares[i] + addedShares[i], curShares[i], 1, err);
        }
    }

    /// @dev Check that the operator has `removedShares` fewer operator shares
    /// for each strategy since the last snapshot
    function assert_Snap_Removed_OperatorShares(
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

    /// @dev Check that the operator's shares in ALL strategies have not changed
    /// since the last snapshot
    function assert_Snap_Unchanged_OperatorShares(
        User operator,
        string memory err
    ) internal {
        IStrategy[] memory strategies = allStrats;

        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);

        // For each strategy, check (prev == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Delta_OperatorShares(
        User operator, 
        IStrategy[] memory strategies, 
        int[] memory shareDeltas,
        string memory err
    ) internal {
        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            uint expectedShares;
            if (shareDeltas[i] < 0) {
                expectedShares = prevShares[i] - uint(-shareDeltas[i]);
            } else {
                expectedShares = prevShares[i] + uint(shareDeltas[i]);
            }
            assertEq(expectedShares, curShares[i], err);
        }
    }

    /// Snapshot assertions for strategyMgr.stakerStrategyShares and eigenPodMgr.podOwnerShares:

    /// @dev Check that the staker has `addedShares` additional delegatable shares
    /// for each strategy since the last snapshot
    function assert_Snap_Added_StakerShares(
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
            assertApproxEqAbs(prevShares[i] + addedShares[i], curShares[i], 1, err);            
        }
    }

    /// @dev Check that the staker has `removedShares` fewer delegatable shares
    /// for each strategy since the last snapshot
    function assert_Snap_Removed_StakerShares(
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

    /// @dev Check that the staker's delegatable shares in ALL strategies have not changed
    /// since the last snapshot
    function assert_Snap_Unchanged_StakerShares(
        User staker,
        string memory err
    ) internal {
        IStrategy[] memory strategies = allStrats;

        uint[] memory curShares = _getStakerShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerShares(staker, strategies);

        // For each strategy, check (prev == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Removed_StrategyShares(
        IStrategy[] memory strategies,
        uint[] memory removedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getTotalStrategyShares(strategies);

        // Use timewarp to get previous strategy shares
        uint[] memory prevShares = _getPrevTotalStrategyShares(strategies);

        for (uint i = 0; i < strategies.length; i++) {
            // Ignore BeaconChainETH strategy since it doesn't keep track of global strategy shares
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                continue;
            }
            uint prevShare = prevShares[i];
            uint curShare = curShares[i];

            assertEq(prevShare - removedShares[i], curShare, err);
        }
    }

    function assert_Snap_Unchanged_StrategyShares(
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        uint[] memory curShares = _getTotalStrategyShares(strategies);

        // Use timewarp to get previous strategy shares
        uint[] memory prevShares = _getPrevTotalStrategyShares(strategies);

        for (uint i = 0; i < strategies.length; i++) {
            uint prevShare = prevShares[i];
            uint curShare = curShares[i];

            assertEq(prevShare, curShare, err);
        }
    }

    function assert_Snap_Delta_StakerShares(
        User staker, 
        IStrategy[] memory strategies, 
        int[] memory shareDeltas,
        string memory err
    ) internal {
        int[] memory curShares = _getStakerSharesInt(staker, strategies);
        // Use timewarp to get previous staker shares
        int[] memory prevShares = _getPrevStakerSharesInt(staker, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + shareDeltas[i], curShares[i], err);
        }
    }

    /// Snapshot assertions for underlying token balances:

    /// @dev Check that the staker has `addedTokens` additional underlying tokens 
    // since the last snapshot
    function assert_Snap_Added_TokenBalances(
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

    /// @dev Check that the staker has `removedTokens` fewer underlying tokens 
    // since the last snapshot
    function assert_Snap_Removed_TokenBalances(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory removedTokens,
        string memory err
    ) internal {
        IERC20[] memory tokens = _getUnderlyingTokens(strategies);

        uint[] memory curTokenBalances = _getTokenBalances(staker, tokens);
        // Use timewarp to get previous token balances
        uint[] memory prevTokenBalances = _getPrevTokenBalances(staker, tokens);

        for (uint i = 0; i < tokens.length; i++) {
            uint prevBalance = prevTokenBalances[i];
            uint curBalance = curTokenBalances[i];

            assertEq(prevBalance - removedTokens[i], curBalance, err);
        }
    }

    /// @dev Check that the staker's underlying token balance for ALL tokens have
    /// not changed since the last snapshot
    function assert_Snap_Unchanged_TokenBalances(
        User staker,
        string memory err
    ) internal {
        IERC20[] memory tokens = allTokens;

        uint[] memory curTokenBalances = _getTokenBalances(staker, tokens);
        // Use timewarp to get previous token balances
        uint[] memory prevTokenBalances = _getPrevTokenBalances(staker, tokens);

        for (uint i = 0; i < tokens.length; i++) {
            assertEq(prevTokenBalances[i], curTokenBalances[i], err);
        }
    }

    /// Other snapshot assertions:

    function assert_Snap_Added_QueuedWithdrawals(
        User staker, 
        IDelegationManager.Withdrawal[] memory withdrawals,
        string memory err
    ) internal {
        uint curQueuedWithdrawals = _getCumulativeWithdrawals(staker);
        // Use timewarp to get previous cumulative withdrawals
        uint prevQueuedWithdrawals = _getPrevCumulativeWithdrawals(staker);

        assertEq(prevQueuedWithdrawals + withdrawals.length, curQueuedWithdrawals, err);
    }

    function assert_Snap_Added_QueuedWithdrawal(
        User staker, 
        IDelegationManager.Withdrawal memory /*withdrawal*/,
        string memory err
    ) internal {
        uint curQueuedWithdrawal = _getCumulativeWithdrawals(staker);
        // Use timewarp to get previous cumulative withdrawals
        uint prevQueuedWithdrawal = _getPrevCumulativeWithdrawals(staker);

        assertEq(prevQueuedWithdrawal + 1, curQueuedWithdrawal, err);
    }

    /*******************************************************************************
                                UTILITY METHODS
    *******************************************************************************/

    function _randWithdrawal(
        IStrategy[] memory strategies, 
        uint[] memory shares
    ) internal returns (IStrategy[] memory, uint[] memory) {
        uint stratsToWithdraw = _randUint({ min: 1, max: strategies.length });

        IStrategy[] memory withdrawStrats = new IStrategy[](stratsToWithdraw);
        uint[] memory withdrawShares = new uint[](stratsToWithdraw);

        for (uint i = 0; i < stratsToWithdraw; i++) {
            uint sharesToWithdraw;

            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                // For native eth, withdraw a random amount of gwei (at least 1)
                uint portion = _randUint({ min: 1, max: shares[i] / GWEI_TO_WEI });
                portion *= GWEI_TO_WEI;

                sharesToWithdraw = shares[i] - portion;
            } else {
                // For LSTs, withdraw a random amount of shares (at least 1)
                uint portion = _randUint({ min: 1, max: shares[i] });

                sharesToWithdraw = shares[i] - portion;
            }

            withdrawStrats[i] = strategies[i];
            withdrawShares[i] = sharesToWithdraw;
        }

        return (withdrawStrats, withdrawShares);
    }

    /**
     * Helpful getters:
     */
    function _randBalanceUpdate(
        User staker,
        IStrategy[] memory strategies
    ) internal returns (int[] memory, int[] memory, int[] memory) {

        int[] memory tokenDeltas = new int[](strategies.length);
        int[] memory stakerShareDeltas = new int[](strategies.length);
        int[] memory operatorShareDeltas = new int[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // TODO - could choose and set a "next updatable validator" at random here
                uint40 validator = staker.getUpdatableValidator();
                uint64 beaconBalanceGwei = beaconChain.balanceOfGwei(validator);

                // For native eth, add or remove a random amount of Gwei - minimum 1
                // and max of the current beacon chain balance
                int64 deltaGwei = int64(int(_randUint({ min: 1, max: beaconBalanceGwei })));
                bool addTokens = _randBool();
                deltaGwei = addTokens ? deltaGwei : -deltaGwei;

                tokenDeltas[i] = int(deltaGwei) * int(GWEI_TO_WEI);

                // stakerShareDeltas[i] = _calculateSharesDelta(newPodBalanceGwei, oldPodBalanceGwei);
                stakerShareDeltas[i] = _calcNativeETHStakerShareDelta(staker, validator, beaconBalanceGwei, deltaGwei);
                operatorShareDeltas[i] = _calcNativeETHOperatorShareDelta(staker, stakerShareDeltas[i]);

                emit log_named_uint("current beacon balance (gwei): ", beaconBalanceGwei);
                // emit log_named_uint("current validator pod balance (gwei): ", oldPodBalanceGwei);
                emit log_named_int("beacon balance delta (gwei): ", deltaGwei);
                emit log_named_int("staker share delta (gwei): ", stakerShareDeltas[i] / int(GWEI_TO_WEI));
                emit log_named_int("operator share delta (gwei): ", operatorShareDeltas[i] / int(GWEI_TO_WEI));
            } else {
                // For LSTs, mint a random token amount
                uint portion = _randUint({ min: MIN_BALANCE, max: MAX_BALANCE });
                StdCheats.deal(address(strat.underlyingToken()), address(staker), portion);

                int delta = int(portion);
                tokenDeltas[i] = delta;
                stakerShareDeltas[i] = int(strat.underlyingToShares(uint(delta)));
                operatorShareDeltas[i] = int(strat.underlyingToShares(uint(delta)));
            }
        }
        return (tokenDeltas, stakerShareDeltas, operatorShareDeltas);
    }

    function _calcNativeETHStakerShareDelta(
        User staker, 
        uint40 validatorIndex, 
        uint64 beaconBalanceGwei, 
        int64 deltaGwei
    ) internal view returns (int) {
        uint64 oldPodBalanceGwei = 
            staker
                .pod()
                .validatorPubkeyHashToInfo(beaconChain.pubkeyHash(validatorIndex))
                .restakedBalanceGwei;
        
        uint64 newPodBalanceGwei = _calcPodBalance(beaconBalanceGwei, deltaGwei);

        return (int(uint(newPodBalanceGwei)) - int(uint(oldPodBalanceGwei))) * int(GWEI_TO_WEI);
    }
    
    function _calcPodBalance(uint64 beaconBalanceGwei, int64 deltaGwei) internal pure returns (uint64) {
        uint64 podBalanceGwei;
        if (deltaGwei < 0) {
            podBalanceGwei = beaconBalanceGwei - uint64(uint(int(-deltaGwei)));
        } else {
            podBalanceGwei = beaconBalanceGwei + uint64(uint(int(deltaGwei)));
        }

        if (podBalanceGwei > MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR) {
            podBalanceGwei = MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR;
        }

        return podBalanceGwei;
    }

    function _calcNativeETHOperatorShareDelta(User staker, int shareDelta) internal view returns (int) {
        int curPodOwnerShares = eigenPodManager.podOwnerShares(address(staker));
        int newPodOwnerShares = curPodOwnerShares + shareDelta;

        if (curPodOwnerShares <= 0) {
            // if the shares started negative and stayed negative, then there cannot have been an increase in delegateable shares
            if (newPodOwnerShares <= 0) {
                return 0;
            // if the shares started negative and became positive, then the increase in delegateable shares is the ending share amount
            } else {
                return newPodOwnerShares;
            }
        } else {
            // if the shares started positive and became negative, then the decrease in delegateable shares is the starting share amount
            if (newPodOwnerShares <= 0) {
                return (-curPodOwnerShares);
            // if the shares started positive and stayed positive, then the change in delegateable shares
            // is the difference between starting and ending amounts
            } else {
                return (newPodOwnerShares - curPodOwnerShares);
            }
        }
    }

    /// @dev For some strategies/underlying token balances, calculate the expected shares received
    /// from depositing all tokens
    function _calculateExpectedShares(IStrategy[] memory strategies, uint[] memory tokenBalances) internal returns (uint[] memory) {
        uint[] memory expectedShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint tokenBalance = tokenBalances[i];
            if (strat == BEACONCHAIN_ETH_STRAT) {
                expectedShares[i] = tokenBalance;
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
                expectedTokens[i] = shares[i];
            } else {
                expectedTokens[i] = strat.sharesToUnderlying(shares[i]);
            }
        }

        return expectedTokens;
    }

    function _getWithdrawalHashes(
        IDelegationManager.Withdrawal[] memory withdrawals
    ) internal view returns (bytes32[] memory) {
        bytes32[] memory withdrawalRoots = new bytes32[](withdrawals.length);

        for (uint i = 0; i < withdrawals.length; i++) {
            withdrawalRoots[i] = delegationManager.calculateWithdrawalRoot(withdrawals[i]);
        }

        return withdrawalRoots;
    }

    /// @dev Converts a list of strategies to underlying tokens
    function _getUnderlyingTokens(IStrategy[] memory strategies) internal view returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](strategies.length);

        for (uint i = 0; i < tokens.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokens[i] = NATIVE_ETH;
            } else {
                tokens[i] = strat.underlyingToken();
            }
        }

        return tokens;
    }

    modifier timewarp() {
        uint curState = timeMachine.warpToLast();
        _;
        timeMachine.warpToPresent(curState);
    }

    /// @dev Given a list of strategies, roll the block number forward to the
    /// a valid blocknumber to completeWithdrawals
    function _rollBlocksForCompleteWithdrawals(IStrategy[] memory strategies) internal {
        // uint256 blocksToRoll = delegationManager.minWithdrawalDelayBlocks();
        // for (uint i = 0; i < strategies.length; i++) {
        //     uint256 withdrawalDelayBlocks = delegationManager.strategyWithdrawalDelayBlocks(strategies[i]);
        //     if (withdrawalDelayBlocks > blocksToRoll) {
        //         blocksToRoll = withdrawalDelayBlocks;
        //     }
        // }
        cheats.roll(block.number + delegationManager.getWithdrawalDelay(strategies));
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
                // This method should only be used for tests that handle positive
                // balances. Negative balances are an edge case that require
                // the own tests and helper methods.
                int shares = eigenPodManager.podOwnerShares(address(staker));
                if (shares < 0) {
                    revert("_getStakerShares: negative shares");
                }

                curShares[i] = uint(shares);
            } else {
                curShares[i] = strategyManager.stakerStrategyShares(address(staker), strat);
            }
        }

        return curShares;
    }

    /// @dev Uses timewarp modifier to get staker shares at the last snapshot
    function _getPrevStakerSharesInt(
        User staker, 
        IStrategy[] memory strategies
    ) internal timewarp() returns (int[] memory) {
        return _getStakerSharesInt(staker, strategies);
    }

    /// @dev Looks up each strategy and returns a list of the staker's shares
    function _getStakerSharesInt(User staker, IStrategy[] memory strategies) internal view returns (int[] memory) {
        int[] memory curShares = new int[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                curShares[i] = eigenPodManager.podOwnerShares(address(staker));
            } else {
                curShares[i] = int(strategyManager.stakerStrategyShares(address(staker), strat));
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
            if (tokens[i] == NATIVE_ETH) {
                balances[i] = address(staker).balance;
            } else {
                balances[i] = tokens[i].balanceOf(address(staker));
            }
        }

        return balances;
    }

    function _getPrevTotalStrategyShares(IStrategy[] memory strategies) internal timewarp() returns (uint[] memory) {
        return _getTotalStrategyShares(strategies);
    }

    function _getTotalStrategyShares(IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory shares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            if (strategies[i] != BEACONCHAIN_ETH_STRAT) {
                shares[i] = strategies[i].totalShares();
            }
            // BeaconChainETH strategy doesn't keep track of global strategy shares, so we ignore
        }

        return shares;
    }
}