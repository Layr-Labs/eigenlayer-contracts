// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationBase.t.sol";

contract IntegrationUtils is IntegrationBase {
    using ArrayLib for *;

    uint numStakers;
    uint numOperators;
    uint numAVSs;

    // Lists of operators created before the m2 (not slashing) upgrade
    //
    // When we call _upgradeEigenLayerContracts, we iterate over
    // these lists and migrate perform the standard migration actions
    // for each user
    User[] operatorsToMigrate;
    User[] stakersToMigrate;

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    /**
     * @dev Create a new user according to configured random variants.
     * This user is ready to deposit into some strategies and has some underlying token balances
     */
    function _newRandomStaker() internal returns (User, IStrategy[] memory, uint[] memory) {
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _randUser(_getStakerName());

        if (!isUpgraded) stakersToMigrate.push(staker);

        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "_newRandomStaker: failed to award token balances");
        return (staker, strategies, tokenBalances);
    }

    /// Given a list of strategies, creates a new user with random token balances in each underlying token
    function _newStaker(IStrategy[] memory strategies) internal returns (User, uint[] memory) {
        (User staker, uint[] memory tokenBalances) = _randUser(_getStakerName(), strategies);

        if (!isUpgraded) stakersToMigrate.push(staker);

        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "_newStaker: failed to award token balances");
        return (staker, tokenBalances);
    }

    /// @dev Creates a new operator with no assets
    function _newRandomOperator() internal returns (User) {
        User operator = _randUser_NoAssets(_getOperatorName());

        /// Registration flow differs for M2 vs Slashing release
        if (!isUpgraded) {
            User_M2(payable(operator)).registerAsOperator_M2();

            operatorsToMigrate.push(operator);
        } else {
            operator.registerAsOperator();

            rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});
        }

        assertTrue(delegationManager().isOperator(address(operator)), "_newRandomOperator: operator should be registered");
        assertEq(delegationManager().delegatedTo(address(operator)), address(operator), "_newRandomOperator: should be self-delegated");
        return operator;
    }

    function _newRandomAVS() internal returns (AVS avs) {
        string memory avsName = string.concat("avs", cheats.toString(numAVSs));
        avs = _genRandAVS(avsName);
        avs.updateAVSMetadataURI("https://example.com");
        ++numAVSs;
    }

    /// @dev Name a newly-created staker ("staker1", "staker2", ...)
    function _getStakerName() private returns (string memory) {
        numStakers++;

        string memory stakerNum = cheats.toString(numStakers);
        string memory namePrefix = isUpgraded ? "staker" : "m2-staker";

        return string.concat(namePrefix, stakerNum);
    }

    /// @dev Name a newly-created operator ("operator1", "operator2", ...)
    function _getOperatorName() private returns (string memory) {
        numOperators++;

        string memory operatorNum = cheats.toString(numOperators);
        string memory namePrefix = isUpgraded ? "operator" : "m2-operator";

        return string.concat(namePrefix, operatorNum);
    }

    /// @dev Send a random amount of ETH (up to 10 gwei) to the destination via `call`,
    /// triggering its fallback function. Sends a gwei-divisible amount as well as a
    /// non-gwei-divisible amount.
    ///
    /// Total sent == `gweiSent + remainderSent`
    function _sendRandomETH(address destination) internal returns (uint64 gweiSent, uint remainderSent) {
        gweiSent = uint64(_randUint({min: 1, max: 10}));
        remainderSent = _randUint({min: 1, max: 100});
        uint totalSent = (gweiSent * GWEI_TO_WEI) + remainderSent;

        cheats.deal(address(this), address(this).balance + totalSent);
        bool r;
        bytes memory d;
        (r, d) = destination.call{value: totalSent}("");

        return (gweiSent, remainderSent);
    }

    /// @dev Choose a random subset of validators (selects AT LEAST ONE)
    function _choose(uint40[] memory validators) internal returns (uint40[] memory) {
        return validators.setLength(_randUint({min: 1, max: validators.length > 1 ? validators.length - 1 : 1}));
    }

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    /// @dev Generate params to allocate all available magnitude to each strategy in the operator set
    function _genAllocation_AllAvailable(User operator, OperatorSet memory operatorSet)
        internal
        view
        returns (AllocateParams memory params)
    {
        return _genAllocation_AllAvailable({
            operator: operator,
            operatorSet: operatorSet,
            strategies: allocationManager().getStrategiesInOperatorSet(operatorSet)
        });
    }

    /// @dev Generate params to allocate all available magnitude to each strategy in the operator set
    function _genAllocation_AllAvailable(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        view
        returns (AllocateParams memory params)
    {
        params.operatorSet = operatorSet;
        params.strategies = strategies;
        params.newMagnitudes = new uint64[](params.strategies.length);

        for (uint i = 0; i < params.strategies.length; i++) {
            IStrategy strategy = params.strategies[i];
            params.newMagnitudes[i] = allocationManager().getMaxMagnitude(address(operator), strategy);
        }
    }

    /// @dev Gen params to allocate half of available magnitude to each strategy in the operator set
    /// returns the params to complete this allocation
    function _genAllocation_HalfAvailable(User operator, OperatorSet memory operatorSet)
        internal
        view
        returns (AllocateParams memory params)
    {
        return _genAllocation_HalfAvailable({
            operator: operator,
            operatorSet: operatorSet,
            strategies: allocationManager().getStrategiesInOperatorSet(operatorSet)
        });
    }

    /// @dev Gen params to allocate half of available magnitude to each strategy in the operator set
    /// returns the params to complete this allocation
    function _genAllocation_HalfAvailable(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        view
        returns (AllocateParams memory params)
    {
        params.operatorSet = operatorSet;
        params.strategies = strategies;
        params.newMagnitudes = new uint64[](params.strategies.length);

        Allocation[] memory allocations = _getAllocations(operator, operatorSet, strategies);
        Magnitudes[] memory magnitudes = _getMagnitudes(operator, strategies);

        for (uint i = 0; i < params.strategies.length; i++) {
            uint64 halfAvailable = uint64(magnitudes[i].allocatable) / 2;
            params.newMagnitudes[i] = allocations[i].currentMagnitude + halfAvailable;
        }
    }

    /// @dev Generate params to allocate a random portion of available magnitude to each strategy
    /// in the operator set. All strategies will have a nonzero allocation, and the minimum allocation
    /// will be 10% of available magnitude
    function _genAllocation_Rand(User operator, OperatorSet memory operatorSet) internal returns (AllocateParams memory params) {
        params.operatorSet = operatorSet;
        params.strategies = allocationManager().getStrategiesInOperatorSet(operatorSet);
        params.newMagnitudes = new uint64[](params.strategies.length);

        Allocation[] memory allocations = _getAllocations(operator, operatorSet, params.strategies);
        Magnitudes[] memory magnitudes = _getMagnitudes(operator, params.strategies);

        for (uint i = 0; i < params.strategies.length; i++) {
            // minimum of 10%, maximum of 100%. increments of 10%.
            uint r = _randUint({min: 1, max: 10});
            uint64 allocation = uint64(magnitudes[i].allocatable) / uint64(r);

            params.newMagnitudes[i] = allocations[i].currentMagnitude + allocation;
        }
    }

    /// @dev Generates params for a half deallocation from all strategies the operator is allocated to in the operator set
    function _genDeallocation_HalfRemaining(User operator, OperatorSet memory operatorSet)
        internal
        view
        returns (AllocateParams memory params)
    {
        return _genDeallocation_HalfRemaining({
            operator: operator,
            operatorSet: operatorSet,
            strategies: allocationManager().getStrategiesInOperatorSet(operatorSet)
        });
    }

    /// @dev Generates params for a half deallocation from all strategies the operator is allocated to in the operator set
    function _genDeallocation_HalfRemaining(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        view
        returns (AllocateParams memory params)
    {
        params.operatorSet = operatorSet;
        params.strategies = strategies;
        params.newMagnitudes = new uint64[](params.strategies.length);

        for (uint i = 0; i < params.strategies.length; i++) {
            IStrategy strategy = params.strategies[i];
            params.newMagnitudes[i] = allocationManager().getEncumberedMagnitude(address(operator), strategy) / 2;
        }
    }

    /// @dev Generates params for a full deallocation from all strategies the operator is allocated to in the operator set
    function _genDeallocation_Full(User operator, OperatorSet memory operatorSet) internal view returns (AllocateParams memory params) {
        return _genDeallocation_Full(operator, operatorSet, allocationManager().getStrategiesInOperatorSet(operatorSet));
    }

    /// @dev Generates params for a full deallocation from all strategies the operator is allocated to in the operator set
    function _genDeallocation_Full(User, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        pure
        returns (AllocateParams memory params)
    {
        params.operatorSet = operatorSet;
        params.strategies = strategies;
        params.newMagnitudes = new uint64[](params.strategies.length);
    }

    /// Generate random slashing between 1 and 99%
    function _genSlashing_Rand(User operator, OperatorSet memory operatorSet) internal returns (SlashingParams memory params) {
        params.operator = address(operator);
        params.operatorSetId = operatorSet.id;
        params.description = "genSlashing_Rand";
        params.strategies = allocationManager().getStrategiesInOperatorSet(operatorSet);
        params.wadsToSlash = new uint[](params.strategies.length);

        /// 1% * rand(1, 99)
        uint slashWad = 1e16 * _randUint({min: 1, max: 99});

        for (uint i = 0; i < params.wadsToSlash.length; i++) {
            params.wadsToSlash[i] = slashWad;
        }
    }

    function _genSlashing_Half(User operator, OperatorSet memory operatorSet) internal view returns (SlashingParams memory params) {
        params.operator = address(operator);
        params.operatorSetId = operatorSet.id;
        params.description = "genSlashing_Half";
        params.strategies = allocationManager().getStrategiesInOperatorSet(operatorSet);
        params.wadsToSlash = new uint[](params.strategies.length);

        // slash 50%
        for (uint i = 0; i < params.wadsToSlash.length; i++) {
            params.wadsToSlash[i] = 5e17;
        }
    }

    function _genSlashing_Full(User operator, OperatorSet memory operatorSet) internal view returns (SlashingParams memory params) {
        params.operator = address(operator);
        params.operatorSetId = operatorSet.id;
        params.description = "_genSlashing_Full";
        params.strategies = allocationManager().getStrategiesInOperatorSet(operatorSet);
        params.wadsToSlash = new uint[](params.strategies.length);

        // slash 100%
        for (uint i = 0; i < params.wadsToSlash.length; i++) {
            params.wadsToSlash[i] = 1e18;
        }
    }

    function _genSlashing_Custom(User operator, OperatorSet memory operatorSet, uint wadsToSlash)
        internal
        view
        returns (SlashingParams memory params)
    {
        params.operator = address(operator);
        params.operatorSetId = operatorSet.id;
        params.description = "_genSlashing_Custom";
        params.strategies = allocationManager().getStrategiesInOperatorSet(operatorSet);
        params.wadsToSlash = new uint[](params.strategies.length);

        for (uint i = 0; i < params.wadsToSlash.length; i++) {
            params.wadsToSlash[i] = wadsToSlash;
        }
    }

    function _randBalanceUpdate(User staker, IStrategy[] memory strategies) internal returns (int[] memory, int[] memory, int[] memory) {
        int[] memory tokenDeltas = new int[](strategies.length);
        int[] memory stakerShareDeltas = new int[](strategies.length);
        int[] memory operatorShareDeltas = new int[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // For native ETH, we're either going to slash the staker's validators,
                // or award them consensus rewards. In either case, the magnitude of
                // the balance update depends on the staker's active validator count
                uint activeValidatorCount = staker.pod().activeValidatorCount();
                int64 deltaGwei;
                if (_randBool()) {
                    uint40[] memory validators = staker.getActiveValidators();
                    emit log_named_uint("slashing validators", validators.length);

                    deltaGwei = -int64(beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor));
                    beaconChain.advanceEpoch_NoRewards();

                    emit log_named_int("slashed amount", deltaGwei);
                } else {
                    emit log("generating consensus rewards for validators");

                    deltaGwei = int64(uint64(activeValidatorCount) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI());
                    beaconChain.advanceEpoch_NoWithdraw();
                }

                tokenDeltas[i] = int(deltaGwei) * int(GWEI_TO_WEI);

                stakerShareDeltas[i] = tokenDeltas[i];
                operatorShareDeltas[i] = _calcNativeETHOperatorShareDelta(staker, stakerShareDeltas[i]);

                emit log_named_int("beacon balance delta (gwei): ", deltaGwei);
                emit log_named_int("staker share delta (gwei): ", stakerShareDeltas[i] / int(GWEI_TO_WEI));
                emit log_named_int("operator share delta (gwei): ", operatorShareDeltas[i] / int(GWEI_TO_WEI));
            } else {
                // For LSTs, mint a random token amount
                uint portion = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});
                StdCheats.deal(address(strat.underlyingToken()), address(staker), portion);

                int delta = int(portion);
                tokenDeltas[i] = delta;
                stakerShareDeltas[i] = int(strat.underlyingToShares(uint(delta)));
                operatorShareDeltas[i] = int(strat.underlyingToShares(uint(delta)));
            }
        }
        return (tokenDeltas, stakerShareDeltas, operatorShareDeltas);
    }

    /// @dev Rolls forward by the minimum withdrawal delay blocks.
    function _rollBlocksForCompleteWithdrawals(Withdrawal[] memory withdrawals) internal {
        uint latest;
        for (uint i = 0; i < withdrawals.length; ++i) {
            if (withdrawals[i].startBlock > latest) latest = withdrawals[i].startBlock;
        }
        cheats.roll(latest + delegationManager().minWithdrawalDelayBlocks() + 1);
    }

    function _rollForward_AllocationDelay(User operator) internal {
        uint32 delay = _getExistingAllocationDelay(operator);
        rollForward(delay);
    }

    function _rollBackward_AllocationDelay(User operator) internal {
        uint32 delay = _getExistingAllocationDelay(operator);
        rollBackward(delay);
    }

    function _rollForward_DeallocationDelay() internal {
        rollForward(allocationManager().DEALLOCATION_DELAY() + 1);
    }

    function _rollBackward_DeallocationDelay() internal {
        rollBackward(allocationManager().DEALLOCATION_DELAY() + 1);
    }

    /// @dev Rolls forward by the default allocation delay blocks.
    function _rollBlocksForCompleteAllocation(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies) internal {
        uint latest;
        for (uint i = 0; i < strategies.length; ++i) {
            uint effectBlock = allocationManager().getAllocation(address(operator), operatorSet, strategies[i]).effectBlock;
            if (effectBlock > latest) latest = effectBlock;
        }
        cheats.roll(latest + 1);
    }

    /// @dev Rolls forward by the default allocation delay blocks.
    function _rollBlocksForCompleteAllocation(User operator, OperatorSet[] memory operatorSets, IStrategy[] memory strategies) internal {
        uint latest;
        for (uint i = 0; i < operatorSets.length; ++i) {
            for (uint j = 0; j < strategies.length; ++j) {
                uint effectBlock = allocationManager().getAllocation(address(operator), operatorSets[i], strategies[j]).effectBlock;
                if (effectBlock > latest) latest = effectBlock;
            }
        }
        cheats.roll(latest + 1);
    }
}
