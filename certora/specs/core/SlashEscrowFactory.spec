import "../erc20cvl.spec";
import "../libraries/EnumerableSet.spec";
using SlashEscrowFactoryHarness as factory;
using SlashEscrow as escrow;
using StrategyManager as StrategyManager;
using OperatorSetLib as OperatorSetLib;
using AllocationManager as AllocationManager;

use invariant _pendingSlashIdsInvariant;
use invariant _pendingStrategiesForSlashIdInvariant;
use invariant _pendingOperatorSetsInvariant;

methods {
    // Strategy methods
    function _.underlyingToken() external => DISPATCHER(true);
    function _.withdraw(address,address,uint256) external => DISPATCHER(true);

    // SlashEscrow
    function _.releaseTokens(address, address, SlashEscrowFactoryHarness.OperatorSet,uint256, address, address) external => DISPATCHER(true);

    unresolved external in StrategyManager._ => DISPATCH [
        getSlashEscrow(SlashEscrowFactoryHarness.OperatorSet, uint256),
        StrategyBase.withdraw(address,address,uint256)
    ] default HAVOC_ECF;

    // SlashEscrowFactory methods
    function initiateSlashEscrow(SlashEscrowFactoryHarness.OperatorSet, uint256, address) external;
    function releaseSlashEscrow(SlashEscrowFactoryHarness.OperatorSet, uint256) external;
    function releaseSlashEscrowByStrategy(SlashEscrowFactoryHarness.OperatorSet, uint256, address) external;
    function getSlashEscrow(SlashEscrowFactoryHarness.OperatorSet, uint256) external returns (address) envfree;
    function isDeployedSlashEscrow(SlashEscrowFactoryHarness.OperatorSet, uint256) external returns (bool) envfree;
    function getEscrowStartBlock(SlashEscrowFactoryHarness.OperatorSet, uint256) external returns (uint256) envfree;
    function getEscrowCompleteBlock(SlashEscrowFactoryHarness.OperatorSet, uint256) external returns (uint32) envfree;
    function getStrategyEscrowDelay(address) external returns (uint32) envfree;
    function getGlobalEscrowDelay() external returns (uint32) envfree;
    function isEscrowPaused(SlashEscrowFactoryHarness.OperatorSet, uint256) external returns (bool) envfree;
    function isPendingSlashId(SlashEscrowFactoryHarness.OperatorSet, uint256) external returns (bool) envfree;
    function getTotalPendingSlashIds(SlashEscrowFactoryHarness.OperatorSet) external returns (uint256) envfree;
    function getTotalPendingStrategiesForSlashId(SlashEscrowFactoryHarness.OperatorSet, uint256 slashId) external returns (uint256) envfree;
    function getPendingUnderlyingAmountForStrategy(SlashEscrowFactoryHarness.OperatorSet, uint256 slashId, address strategy) external returns (uint256) envfree;
    function computeSlashEscrowSalt(SlashEscrowFactoryHarness.OperatorSet, uint256 slashId) external returns (bytes32) envfree;
    function getTotalPendingOperatorSets() external returns (uint256) envfree;

    //harness methods
    function isStrategyPending(SlashEscrowFactoryHarness.OperatorSet operatorSet, uint256 slashId, address strategy) external returns (bool) envfree;
    function getOperatorSetKey(SlashEscrowFactoryHarness.OperatorSet) external returns (bytes32) envfree;

    function getPendingStrategiesForSlashId(SlashEscrowFactoryHarness.OperatorSet, uint256) external returns (address[]) envfree;
    
    // SlashEscrow methods
    function releaseTokens(address, address, SlashEscrowFactoryHarness.OperatorSet, uint256, address, address) external;
    
    // StrategyManager methods
    function clearBurnOrRedistributableShares(SlashEscrowFactoryHarness.OperatorSet, uint256) external;
    function clearBurnOrRedistributableSharesByStrategy(SlashEscrowFactoryHarness.OperatorSet, uint256, address) external;

}

// Rule: Escrow deployment determinism
rule escrowDeploymentDeterminism(env e1, env e2) {
    SlashEscrowFactoryHarness.OperatorSet operatorSet;
    uint256 slashId;
    
    bytes32 salt1 = computeSlashEscrowSalt(operatorSet, slashId);
    bytes32 salt2 = computeSlashEscrowSalt(operatorSet, slashId);
    
    assert salt1 == salt2;
    
    address escrow1 = getSlashEscrow(operatorSet, slashId);
    address escrow2 = getSlashEscrow(operatorSet, slashId);
    
    assert escrow1 == escrow2;
}


// @title initiateSlashEscrow starts a countdown (sets start block)
rule initiateSlashEscrowStartsCountdown(
    SlashEscrowFactoryHarness.OperatorSet operatorSet,
    uint256 slashId,
    address strategy
) {
    env e;
    
    bool isAlreadyDeployed = isDeployedSlashEscrow(operatorSet, slashId);
    // Call initiateSlashEscrow
    initiateSlashEscrow(e, operatorSet, slashId, strategy);
    
    // Postcondition: Start block is set to current block
    uint256 startBlockAfter = getEscrowStartBlock(operatorSet, slashId);
    
    assert !isAlreadyDeployed => startBlockAfter == require_uint32(e.block.number);
}

// @title Countdown is identical for all strategies in same slash
rule countdownIdenticalForSameSlash(
    SlashEscrowFactoryHarness.OperatorSet operatorSet,
    uint256 slashId,
    address strategy1,
    address strategy2
) {
    env e1;
    env e2;
    require e1.block.number == e2.block.number; // Same block
    
    // First strategy initiation
    initiateSlashEscrow(e1, operatorSet, slashId, strategy1);
    uint256 startBlock1 = getEscrowStartBlock(operatorSet, slashId);
    
    // Second strategy initiation in same block
    initiateSlashEscrow(e2, operatorSet, slashId, strategy2);
    uint256 startBlock2 = getEscrowStartBlock(operatorSet, slashId);
    
    // Both strategies must have the same start block
    assert startBlock1 == startBlock2;
}

// @title releaseSlashEscrow can only be called after delays have passed
rule releaseSlashEscrowOnlyAfterDelay(
    SlashEscrowFactoryHarness.OperatorSet operatorSet,
    uint256 slashId
) {
    env e;
    
    // Preconditions
    require isPendingSlashId(operatorSet, slashId);
    require !isEscrowPaused(operatorSet, slashId);
    
    uint32 completeBlock = getEscrowCompleteBlock(operatorSet, slashId);
    
    // If current block is before complete block, release should revert
    releaseSlashEscrow@withrevert(e, operatorSet, slashId);
    assert e.block.number < completeBlock => lastReverted;
}

// @title releaseSlashEscrowByStrategy can only be called after delays have passed
rule releaseSlashEscrowByStrategyOnlyAfterDelay(
    SlashEscrowFactoryHarness.OperatorSet operatorSet,
    uint256 slashId,
    address strategy
) {
    env e;
    
    // Preconditions
    require isPendingSlashId(operatorSet, slashId);
    require !isEscrowPaused(operatorSet, slashId);
    
    uint32 completeBlock = getEscrowCompleteBlock(operatorSet, slashId);
    
    // If current block is before complete block, release should revert
    releaseSlashEscrowByStrategy@withrevert(e, operatorSet, slashId, strategy);
    assert e.block.number < completeBlock => lastReverted; 
}

// @title After successful release, slash ID is no longer pending
rule releaseSlashEscrowRemovesPendingStatus(
    SlashEscrowFactoryHarness.OperatorSet operatorSet,
    uint256 slashId
) {
    env e;
    
    // Release the escrow
    releaseSlashEscrow(e, operatorSet, slashId);
    
    // Slash ID should no longer be pending (if all strategies were processed)
    assert getPendingStrategiesForSlashId(operatorSet, slashId).length == 0 => !isPendingSlashId(operatorSet, slashId);
}

// Parametric rule: State consistency across operations
rule stateConsistencyAcrossOperations(env e, method f) filtered { f-> !f.isView }{
    SlashEscrowFactoryHarness.OperatorSet operatorSet;
    
    requireInvariant _pendingSlashIdsInvariant();
    requireInvariant _pendingStrategiesForSlashIdInvariant();
    requireInvariant _pendingOperatorSetsInvariant();

    uint256 pendingSetsBefore = getTotalPendingOperatorSets();
    uint256 pendingSlashIdsBefore = getTotalPendingSlashIds(operatorSet);
    
    calldataarg args;
    f(e, args);
    
    uint256 pendingSetsAfter = getTotalPendingOperatorSets();
    uint256 pendingSlashIdsAfter = getTotalPendingSlashIds(operatorSet);
    
    // Pending counts should only increase with initiateSlashEscrow
    assert f.selector == sig:initiateSlashEscrow(SlashEscrowFactoryHarness.OperatorSet,uint256,address).selector =>
        (pendingSetsAfter >= pendingSetsBefore && pendingSlashIdsAfter >= pendingSlashIdsBefore);
    
    // Pending counts should only decrease with release functions
    assert (f.selector == sig:releaseSlashEscrow(SlashEscrowFactoryHarness.OperatorSet,uint256).selector ||
            f.selector == sig:releaseSlashEscrowByStrategy(SlashEscrowFactoryHarness.OperatorSet,uint256,address).selector) =>
        (pendingSetsAfter <= pendingSetsBefore && pendingSlashIdsAfter <= pendingSlashIdsBefore);
}


// @title Paused escrows cannot be released
rule pausedEscrowsCannotBeReleased(
    SlashEscrowFactoryHarness.OperatorSet operatorSet,
    uint256 slashId
) {
    env e;
    
    // Precondition: Escrow is paused
    require isEscrowPaused(operatorSet, slashId);
    
    // Attempt to release should revert
    releaseSlashEscrow@withrevert(e, operatorSet, slashId);
    assert lastReverted;
    
    // Also test strategy-specific release
    address strategy;
    releaseSlashEscrowByStrategy@withrevert(e, operatorSet, slashId, strategy);
    assert lastReverted;
}

// @title Slash escrow address is deterministic
rule slashEscrowAddressDeterministic(
    env e, 
    OperatorSetLib.OperatorSet operatorSet1,
    OperatorSetLib.OperatorSet operatorSet2,
    uint256 slashId1,
    uint256 slashId2, 
    address strategy
) {

    initiateSlashEscrow(e, operatorSet1, slashId1, strategy);
    initiateSlashEscrow(e, operatorSet2, slashId2, strategy);

    address escrow1 = getSlashEscrow(operatorSet1, slashId1);
    address escrow2 = getSlashEscrow(operatorSet2, slashId2);
    
    // Same operator set and slash ID should produce same escrow address
    assert (getOperatorSetKey(operatorSet1) == getOperatorSetKey(operatorSet2) && slashId1 == slashId2) => (escrow1 == escrow2);
}

// @title Multiple initiations of same slash don't change start block
rule multipleInitiationsSameStartBlock(
    OperatorSetLib.OperatorSet operatorSet,
    uint256 slashId,
    address strategy1,
    address strategy2
) {
    env e1;
    env e2;
    require e1.block.number == e2.block.number;
    
    // First initiation
    initiateSlashEscrow(e1, operatorSet, slashId, strategy1);
    uint256 startBlock1 = getEscrowStartBlock(operatorSet, slashId);
    
    // Second initiation with different strategy
    initiateSlashEscrow(e2, operatorSet, slashId, strategy2);
    uint256 startBlock2 = getEscrowStartBlock(operatorSet, slashId);
    
    // Start block should remain the same
    assert startBlock1 == startBlock2;
}

// @title Only strategy manager can call initiateSlashEscrow
rule onlyStrategyManagerCanInitiate(
    SlashEscrowFactoryHarness.OperatorSet operatorSet,
    uint256 slashId,
    address strategy
) {
    env e;
    
    // If caller is not strategy manager, call should revert
    initiateSlashEscrow@withrevert(e, operatorSet, slashId, strategy);
    assert e.msg.sender != StrategyManager => lastReverted;
}

// @title Start block is always set for pending slash IDs
invariant startBlockSetForPendingSlashIds(SlashEscrowFactoryHarness.OperatorSet operatorSet, uint256 slashId)
    isPendingSlashId(operatorSet, slashId) => getEscrowStartBlock(operatorSet, slashId) > 0 {
        preserved with (env e) {
            require (require_uint32(e.block.number) != 0, "uint32(block.number) used as a starting block for escrow, would overflow in approx 1.6k years");
            requireInvariant _pendingSlashIdsInvariant();
            requireInvariant _pendingStrategiesForSlashIdInvariant();
            requireInvariant _pendingOperatorSetsInvariant();
        }
    }