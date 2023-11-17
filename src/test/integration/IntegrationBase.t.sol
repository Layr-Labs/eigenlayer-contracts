// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import "src/test/integration/IntegrationDeployer.t.sol";
import "src/test/integration/Global.t.sol";
import "src/test/integration/users/User.sol";

abstract contract Flags is Test {

    uint constant FLAG = 1;

    /// @dev Asset flags
    /// These are used with _configRand to determine what assets are given
    /// to a user when they are created.
    uint constant NO_ASSETS = (FLAG << 0); // will have no assets
    uint constant HOLDS_LST = (FLAG << 1); // will hold some random amount of LSTs
    uint constant HOLDS_ETH = (FLAG << 2); // will hold some random amount of ETH
    uint constant HOLDS_MIX = (FLAG << 3); // will hold a mix of LSTs and ETH

    /// @dev Withdrawal flags
    /// These are used with _configRand to determine how a user conducts a withdrawal
    uint constant FULL_WITHDRAW_SINGLE = (FLAG << 0); // stakers will withdraw all assets using a single queued withdrawal
    uint constant FULL_WITHDRAW_MULTI  = (FLAG << 1); // stakers will withdraw all assets using multiple queued withdrawals
    uint constant PART_WITHDRAW_SINGLE = (FLAG << 2); // stakers will withdraw some, but not all assets

    /// Note: Thought about the following flags (but did not implement) -
    ///
    /// WithdrawerType (SELF_WITHDRAWER, OTHER_WITHDRAWER)
    ///     - especially with EPM share handling, this felt like it deserved its own test rather than a fuzzy state
    /// CompletionType (AS_TOKENS, AS_SHARES)
    ///     - same reason as above
    ///
    /// DepositMethod (DEPOSIT_STD, DEPOSIT_WITH_SIG)
    ///     - could still do this!
    /// WithdrawalMethod (QUEUE_WITHDRAWAL, UNDELEGATE, REDELEGATE)
    ///     - could still do this! 
    ///     - This would trigger staker.queueWithdrawals to use either `queueWithdrawals` or `undelegate` under the hood
    ///     - "redelegate" would be like the above, but adding a new `delegateTo` step after undelegating

    bytes32 random;

    bytes stakerAssets;
    bytes operatorAssets;
    bytes withdrawTypes;

    function _configRand(
        uint16 _randomSeed, 
        uint _stakerAssets,
        uint _operatorAssets,
        uint _withdrawType
    ) internal {
        // Using uint16 for the seed type so that if a test fails, it's easier
        // to manually use the seed to replay the same test.
        // TODO - should we expand the type?
        emit log_named_uint("_configRand: set random seed to ", _randomSeed);
        random = keccak256(abi.encodePacked(_randomSeed));

        emit log_named_uint("staker assets: ", _stakerAssets);
        emit log_named_uint("operator assets: ", _operatorAssets);

        emit log_named_uint("HOLDS_LST: ", HOLDS_LST);

        // Convert flag bitmaps to bytes of set bits for easy use with _randUint
        stakerAssets = _bitmapToBytes(_stakerAssets);
        operatorAssets = _bitmapToBytes(_operatorAssets);
        withdrawTypes = _bitmapToBytes(_withdrawType);

        emit log("staker assets");
        for (uint i = 0; i < stakerAssets.length; i++) {
            emit log_named_uint("- ", uint(uint8(stakerAssets[i])));
        }

        emit log("operator assets");
        for (uint i = 0; i < operatorAssets.length; i++) {
            emit log_named_uint("- ", uint(uint8(operatorAssets[i])));
        }

        assertTrue(stakerAssets.length != 0, "_configRand: no staker assets selected");
        assertTrue(operatorAssets.length != 0, "_configRand: no operator assets selected");
        assertTrue(withdrawTypes.length != 0, "_configRand: no withdrawal type selected");
    }

    /**
     * @dev Converts a bitmap into an array of bytes
     * @dev Each byte in the input is processed as indicating a single bit to flip in the bitmap
     */
    function _bitmapToBytes(uint bitmap) internal returns (bytes memory bytesArray) {
        for (uint i = 0; i < 256; ++i) {
            // Mask for i-th bit
            uint mask = uint(1 << i);

            // emit log_named_uint("mask: ", mask);

            // If the i-th bit is flipped, add a byte to the return array
            if (bitmap & mask != 0) {
                emit log_named_uint("i: ", i);
                bytesArray = bytes.concat(bytesArray, bytes1(uint8(1 << i)));
            }
        }
        return bytesArray;
    }

    /// @dev Uses `random` to return a random uint, with a range given by `min` and `max` (inclusive)
    /// @return `min` <= result <= `max`
    function _randUint(uint min, uint max) internal returns (uint) {        
        uint range = max - min + 1;

        // calculate the number of bits needed for the range
        uint bitsNeeded = 0;
        uint tempRange = range;
        while (tempRange > 0) {
            bitsNeeded++;
            tempRange >>= 1;
        }

        // create a mask for the required number of bits
        // and extract the value from the hash
        uint mask = (1 << bitsNeeded) - 1;
        uint value = uint(random) & mask;

        // in case value is out of range, wrap around or retry
        while (value >= range) {
            value = (value - range) & mask;
        }

        // Hash `random` with itself so the next value we generate is different
        random = keccak256(abi.encodePacked(random));
        return min + value;
    }
}

abstract contract IntegrationBase is IntegrationDeployer, Flags {

    Global global;

    function setUp() public virtual override {
        super.setUp();

        global = new Global();
    }

    /**
     * Gen/Init methods:
     */

    /// @dev Uses configured randomness to create a Staker-type user
    /// @return The new User, the strategies the user has token balances for, and the
    /// token balances themselves.
    function _newStaker() internal returns (User, IStrategy[] memory, uint[] memory) {
        // Select assets for the staker
        uint assetIdx = _randUint({ min: 0, max: stakerAssets.length - 1 });
        emit log_named_uint("_newStaker: rand idx: ", assetIdx);

        uint flag = uint(uint8(stakerAssets[assetIdx]));

        emit log_named_uint("_newStaker: asset flag is ", flag);

        User staker = new User(delegationManager, strategyManager, eigenPodManager, global);
        IStrategy[] memory strategies;
        uint[] memory tokenBalances;

        if (flag == NO_ASSETS) {
            emit log("_newStaker: creating NO_ASSETS user");
        } else if (flag == HOLDS_LST) {
            emit log("_newStaker: creating HOLDS_LST user");

            strategies = new IStrategy[](1);
            tokenBalances = new uint[](1);

            // TODO rand more than 1 strategy
            uint idx = _randUint({ min: 0, max: _strategies.length - 1 });

            IStrategy strategy = _strategies[idx];
            IERC20 underlyingToken = strategy.underlyingToken();
        
            // TODO better rand token balance
            uint balance = _randUint({ min: 1e6, max: 5e6 });

            // award tokens
            StdCheats.deal(address(underlyingToken), address(staker), balance);

            strategies[0] = strategy;
            tokenBalances[0] = balance;

            emit log_named_string("_newStaker: user receiving token: ", IERC20Metadata(address(underlyingToken)).name());
            emit log_named_uint("_newStaker: awarded tokens to user: ", balance);
        } else if (flag == HOLDS_ETH) {
            emit log("_newStaker: creating HOLDS_ETH user");
            // Need to support native eth initialization
            revert("unimplemented");
        } else if (flag == HOLDS_MIX) {
            emit log("_newStaker: creating HOLDS_MIX user");
            // Need to support native eth initialization
            revert("unimplemented");
        } else {
            revert("_newStaker: invalid flag");
        }

        return (staker, strategies, tokenBalances);
    }

    function _newOperator() internal returns (User, IStrategy[] memory, uint[] memory) {
        // Select assets for the operator
        uint assetIdx = _randUint({ min: 0, max: operatorAssets.length - 1 });
        emit log_named_uint("_newOperator: rand idx: ", assetIdx);
        
        uint flag = uint(uint8(operatorAssets[assetIdx]));

        emit log_named_uint("_newOperator: asset flag is ", flag);

        User operator = new User(delegationManager, strategyManager, eigenPodManager, global);
        IStrategy[] memory strategies;
        uint[] memory tokenBalances;

        if (flag == NO_ASSETS) {
            emit log("_newOperator: creating NO_ASSETS user");
        } else if (flag == HOLDS_LST) {
            emit log("_newOperator: creating HOLDS_LST user");

            strategies = new IStrategy[](1);
            tokenBalances = new uint[](1);

            // TODO rand more than 1 strategy
            uint idx = _randUint({ min: 0, max: _strategies.length - 1 });

            IStrategy strategy = _strategies[idx];
            IERC20 underlyingToken = strategy.underlyingToken();
        
            // TODO better rand token balance
            uint balance = _randUint({ min: 1e6, max: 5e6 });

            // award tokens
            StdCheats.deal(address(underlyingToken), address(operator), balance);

            strategies[0] = strategy;
            tokenBalances[0] = balance;

            emit log_named_string("_newOperator: user receiving token: ", IERC20Metadata(address(underlyingToken)).name());
            emit log_named_uint("_newOperator: awarded tokens to user: ", balance);
        } else if (flag == HOLDS_ETH) {
            emit log("_newOperator: creating HOLDS_ETH user");
            // Need to support native eth initialization
            revert("unimplemented");
        } else if (flag == HOLDS_MIX) {
            emit log("_newOperator: creating HOLDS_MIX user");
            // Need to support native eth initialization
            revert("unimplemented");
        } else {
            revert("_newOperator: invalid flag");
        }

        operator.registerAsOperator();
        assertTrue(delegationManager.isOperator(address(operator)), "_newOperator: operator should be registered");
        // assert_HasNoDelegatableShares(operator, "_newOperator: new operator should not have delegatable shares");

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

    function assert_HasNoUnderlyingTokenBalance(User user, IStrategy[] memory strategies, string memory err) internal {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint tokenBalance;
            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokenBalance = address(user).balance;
            } else {
                tokenBalance = strat.underlyingToken().balanceOf(address(user));
            }

            assertEq(0, tokenBalance, err);
        }
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
     * Snapshot assertions combine Global's snapshots with assertions
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
        uint curState = global.warpToLast();
        _;
        global.warpToPresent(curState);
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