// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import "src/test/integration/TypeImporter.t.sol";
import "src/test/integration/IntegrationDeployer.t.sol";

contract IntegrationGetters is IntegrationDeployer, TypeImporter {
    using ArrayLib for *;

    struct Magnitudes {
        uint encumbered;
        uint allocatable;
        uint max;
    }

    modifier timewarp() {
        uint curState = timeMachine.travelToLast();
        _;
        timeMachine.travel(curState);
    }

    function _getPrevAllocations(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        timewarp
        returns (Allocation[] memory)
    {
        return _getAllocations(operator, operatorSet, strategies);
    }

    function _getAllocations(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        view
        returns (Allocation[] memory allocations)
    {
        allocations = new Allocation[](strategies.length);
        for (uint i = 0; i < strategies.length; ++i) {
            allocations[i] = allocationManager.getAllocation(address(operator), operatorSet, strategies[i]);
        }
    }

    function _getPrevAllocatedStrats(User operator, OperatorSet memory operatorSet) internal timewarp returns (IStrategy[] memory) {
        return _getAllocatedStrats(operator, operatorSet);
    }

    function _getAllocatedStrats(User operator, OperatorSet memory operatorSet) internal view returns (IStrategy[] memory) {
        return allocationManager.getAllocatedStrategies(address(operator), operatorSet);
    }

    function _getPrevAllocatedSets(User operator) internal timewarp returns (OperatorSet[] memory) {
        return _getAllocatedSets(operator);
    }

    function _getAllocatedSets(User operator) internal view returns (OperatorSet[] memory) {
        return allocationManager.getAllocatedSets(address(operator));
    }

    function _getPrevRegisteredSets(User operator) internal timewarp returns (OperatorSet[] memory) {
        return _getRegisteredSets(operator);
    }

    function _getRegisteredSets(User operator) internal view returns (OperatorSet[] memory) {
        return allocationManager.getRegisteredSets(address(operator));
    }

    function _getPrevMembers(OperatorSet memory operatorSet) internal timewarp returns (address[] memory) {
        return _getMembers(operatorSet);
    }

    function _getMembers(OperatorSet memory operatorSet) internal view returns (address[] memory) {
        return allocationManager.getMembers(operatorSet);
    }

    function _getPrevMagnitudes(User operator, IStrategy[] memory strategies) internal timewarp returns (Magnitudes[] memory) {
        return _getMagnitudes(operator, strategies);
    }

    function _getMagnitudes(User operator, IStrategy[] memory strategies) internal view returns (Magnitudes[] memory magnitudes) {
        magnitudes = new Magnitudes[](strategies.length);
        for (uint i = 0; i < strategies.length; ++i) {
            magnitudes[i] = Magnitudes({
                encumbered: allocationManager.getEncumberedMagnitude(address(operator), strategies[i]),
                allocatable: allocationManager.getAllocatableMagnitude(address(operator), strategies[i]),
                max: allocationManager.getMaxMagnitude(address(operator), strategies[i])
            });
        }
    }

    function _getPrevMinSlashableStake(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        timewarp
        returns (uint[] memory)
    {
        return _getMinSlashableStake(operator, operatorSet, strategies);
    }

    function _getPrevMinSlashableStake(address operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        timewarp
        returns (uint[] memory)
    {
        return _getMinSlashableStake(operator, operatorSet, strategies);
    }

    function _getMinSlashableStake(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        view
        returns (uint[] memory)
    {
        return allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: address(operator).toArray(),
            strategies: strategies,
            futureBlock: uint32(block.number)
        })[0];
    }

    function _getMinSlashableStake(address operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        view
        returns (uint[] memory)
    {
        return allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: address(operator).toArray(),
            strategies: strategies,
            futureBlock: uint32(block.number)
        })[0];
    }

    function _getPrevAllocatedStake(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        timewarp
        returns (uint[] memory)
    {
        return _getAllocatedStake(operator, operatorSet, strategies);
    }

    function _getAllocatedStake(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies)
        internal
        view
        returns (uint[] memory)
    {
        return allocationManager.getAllocatedStake({
            operatorSet: operatorSet,
            operators: address(operator).toArray(),
            strategies: strategies
        })[0];
    }

    function _getPrevIsSlashable(User operator, OperatorSet memory operatorSet) internal timewarp returns (bool) {
        return _getIsSlashable(operator, operatorSet);
    }

    function _getIsSlashable(User operator, OperatorSet memory operatorSet) internal view returns (bool) {
        return allocationManager.isOperatorSlashable(address(operator), operatorSet);
    }

    function _getPrevIsMemberOfSet(User operator, OperatorSet memory operatorSet) internal timewarp returns (bool) {
        return _getIsMemberOfSet(operator, operatorSet);
    }

    function _getIsMemberOfSet(User operator, OperatorSet memory operatorSet) internal view returns (bool) {
        return allocationManager.isMemberOfOperatorSet(address(operator), operatorSet);
    }

    function _getPrevBurnableShares(IStrategy[] memory strategies) internal timewarp returns (uint[] memory) {
        return _getBurnableShares(strategies);
    }

    function _getBurnableShares(IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory burnableShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            if (strategies[i] == beaconChainETHStrategy) burnableShares[i] = eigenPodManager.burnableETHShares();
            else burnableShares[i] = strategyManager.getBurnableShares(strategies[i]);
        }

        return burnableShares;
    }

    function _getPrevSlashableSharesInQueue(User operator, IStrategy[] memory strategies) internal timewarp returns (uint[] memory) {
        return _getSlashableSharesInQueue(operator, strategies);
    }

    function _getSlashableSharesInQueue(User operator, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory slashableShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            slashableShares[i] = delegationManager.getSlashableSharesInQueue(address(operator), strategies[i]);
        }

        return slashableShares;
    }

    function _getPrevOperatorShares(User operator, IStrategy[] memory strategies) internal timewarp returns (uint[] memory) {
        return _getOperatorShares(operator, strategies);
    }

    function _getOperatorShares(User operator, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory curShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            curShares[i] = delegationManager.operatorShares(address(operator), strategies[i]);
        }

        return curShares;
    }

    function _getPrevStakerDepositShares(User staker, IStrategy[] memory strategies) internal timewarp returns (uint[] memory) {
        return _getStakerDepositShares(staker, strategies);
    }

    function _getStakerDepositShares(User staker, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory curShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                int shares = eigenPodManager.podOwnerDepositShares(address(staker));

                if (shares < 0) revert("_getStakerDepositShares: negative shares");

                curShares[i] = uint(shares);
            } else {
                curShares[i] = strategyManager.stakerDepositShares(address(staker), strat);
            }
        }

        return curShares;
    }

    function _getPrevStakerDepositSharesInt(User staker, IStrategy[] memory strategies) internal timewarp returns (int[] memory) {
        return _getStakerDepositSharesInt(staker, strategies);
    }

    function _getStakerDepositSharesInt(User staker, IStrategy[] memory strategies) internal view returns (int[] memory) {
        int[] memory curShares = new int[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) curShares[i] = eigenPodManager.podOwnerDepositShares(address(staker));
            else curShares[i] = int(strategyManager.stakerDepositShares(address(staker), strat));
        }

        return curShares;
    }

    function _getPrevStakerWithdrawableShares(User staker, IStrategy[] memory strategies) internal timewarp returns (uint[] memory) {
        return _getStakerWithdrawableShares(staker, strategies);
    }

    function _getStakerWithdrawableShares(User staker, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(address(staker), strategies);
        return withdrawableShares;
    }

    function _getPrevBeaconChainSlashingFactor(User staker) internal timewarp returns (uint64) {
        return _getBeaconChainSlashingFactor(staker);
    }

    function _getBeaconChainSlashingFactor(User staker) internal view returns (uint64) {
        return eigenPodManager.beaconChainSlashingFactor(address(staker));
    }

    function _getPrevCumulativeWithdrawals(User staker) internal timewarp returns (uint) {
        return _getCumulativeWithdrawals(staker);
    }

    function _getCumulativeWithdrawals(User staker) internal view returns (uint) {
        return delegationManager.cumulativeWithdrawalsQueued(address(staker));
    }

    function _getPrevTokenBalances(User staker, IERC20[] memory tokens) internal timewarp returns (uint[] memory) {
        return _getTokenBalances(staker, tokens);
    }

    function _getTokenBalances(User staker, IERC20[] memory tokens) internal view returns (uint[] memory) {
        uint[] memory balances = new uint[](tokens.length);

        for (uint i = 0; i < tokens.length; i++) {
            if (tokens[i] == NATIVE_ETH) balances[i] = address(staker).balance;
            else balances[i] = tokens[i].balanceOf(address(staker));
        }

        return balances;
    }

    function _getPrevTotalStrategyShares(IStrategy[] memory strategies) internal timewarp returns (uint[] memory) {
        return _getTotalStrategyShares(strategies);
    }

    function _getTotalStrategyShares(IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory shares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            if (strategies[i] != BEACONCHAIN_ETH_STRAT) shares[i] = strategies[i].totalShares();
        }

        return shares;
    }

    function _getDepositScalingFactors(User staker, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory depositScalingFactors = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            depositScalingFactors[i] = _getDepositScalingFactor(staker, strategies[i]);
        }
        return depositScalingFactors;
    }

    function _getDepositScalingFactor(User staker, IStrategy strategy) internal view returns (uint) {
        return delegationManager.depositScalingFactor(address(staker), strategy);
    }

    function _getPrevDepositScalingFactors(User staker, IStrategy[] memory strategies) internal timewarp returns (uint[] memory) {
        return _getDepositScalingFactors(staker, strategies);
    }

    function _getPrevWithdrawableShares(User staker, IStrategy[] memory strategies) internal timewarp returns (uint[] memory) {
        return _getWithdrawableShares(staker, strategies);
    }

    function _getWithdrawableShares(User staker, IStrategy[] memory strategies) internal view returns (uint[] memory withdrawableShares) {
        (withdrawableShares,) = delegationManager.getWithdrawableShares(address(staker), strategies);
    }

    function _getWithdrawableShares(User staker, IStrategy strategy) internal view returns (uint withdrawableShares) {
        (uint[] memory _withdrawableShares,) = delegationManager.getWithdrawableShares(address(staker), strategy.toArray());
        return _withdrawableShares[0];
    }

    function _getActiveValidatorCount(User staker) internal view returns (uint) {
        EigenPod pod = staker.pod();
        return pod.activeValidatorCount();
    }

    function _getPrevActiveValidatorCount(User staker) internal timewarp returns (uint) {
        return _getActiveValidatorCount(staker);
    }

    function _getValidatorStatuses(User staker, bytes32[] memory pubkeyHashes) internal view returns (VALIDATOR_STATUS[] memory) {
        EigenPod pod = staker.pod();
        VALIDATOR_STATUS[] memory statuses = new VALIDATOR_STATUS[](pubkeyHashes.length);

        for (uint i = 0; i < statuses.length; i++) {
            statuses[i] = pod.validatorStatus(pubkeyHashes[i]);
        }

        return statuses;
    }

    function _getPrevValidatorStatuses(User staker, bytes32[] memory pubkeyHashes) internal timewarp returns (VALIDATOR_STATUS[] memory) {
        return _getValidatorStatuses(staker, pubkeyHashes);
    }

    function _getCheckpointTimestamp(User staker) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.currentCheckpointTimestamp();
    }

    function _getPrevCheckpointTimestamp(User staker) internal timewarp returns (uint64) {
        return _getCheckpointTimestamp(staker);
    }

    function _getLastCheckpointTimestamp(User staker) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.lastCheckpointTimestamp();
    }

    function _getPrevLastCheckpointTimestamp(User staker) internal timewarp returns (uint64) {
        return _getLastCheckpointTimestamp(staker);
    }

    function _getWithdrawableRestakedGwei(User staker) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.withdrawableRestakedExecutionLayerGwei();
    }

    function _getPrevWithdrawableRestakedGwei(User staker) internal timewarp returns (uint64) {
        return _getWithdrawableRestakedGwei(staker);
    }

    function _getCheckpointPodBalanceGwei(User staker) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return uint64(pod.currentCheckpoint().podBalanceGwei);
    }

    function _getPrevCheckpointPodBalanceGwei(User staker) internal timewarp returns (uint64) {
        return _getCheckpointPodBalanceGwei(staker);
    }

    function _getCheckpointBalanceExited(User staker, uint64 checkpointTimestamp) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.checkpointBalanceExitedGwei(checkpointTimestamp);
    }

    function _getPrevCheckpointBalanceExited(User staker, uint64 checkpointTimestamp) internal timewarp returns (uint64) {
        return _getCheckpointBalanceExited(staker, checkpointTimestamp);
    }
}
