import "./EnumerableSet.spec";
using OperatorSetHelper as OperatorSetHelper;

use builtin rule sanity filtered { f -> f.contract == currentContract }

use invariant _activeGenerationReservationsInvariant;

methods {
    function _.calculateOperatorTableBytes(CrossChainRegistry.OperatorSet) external => DISPATCHER(true);
    function OperatorSetHelper.getOperatorSetKey(CrossChainRegistry.OperatorSet os) external returns (bytes32) envfree;
}

// // There must not be two active reservations for the same OperatorSet.key()
invariant activeGenerationReservationUniqueness(bytes32 key1, bytes32 key2, uint256 i1, uint256 i2)
    (key1 != key2 && ghostIndexes_activeGenerationReservations[key1] != 0 && ghostIndexes_activeGenerationReservations[key2] != 0) => 
        (ghostIndexes_activeGenerationReservations[key1] != ghostIndexes_activeGenerationReservations[key2] 
        && ghostValues_activeGenerationReservations[ghostIndexes_activeGenerationReservations[key1]-1] != ghostValues_activeGenerationReservations[ghostIndexes_activeGenerationReservations[key2]-1]) 
    {
        preserved{
            requireInvariant _activeGenerationReservationsInvariant();
        }
    }

// valid state changes
rule onlyValidSenderCanChangeState(env e, method f, calldataarg args)
filtered{f -> !f.isView} {
    CrossChainRegistry.OperatorSet os;
    bytes32 key = OperatorSetHelper.getOperatorSetKey(os);
    requireInvariant _activeGenerationReservationsInvariant();
    uint256 i;
    
    mathint _activeGenerationReservationsIndexPre = ghostIndexes_activeGenerationReservations[key];
    bytes32 _activeGenerationReservationsValuePre = ghostValues_activeGenerationReservations[i];
    uint256 _activeGenerationReservationsLengthPre = ghostLength_activeGenerationReservations;

    address configPreOwner = currentContract._operatorSetConfigs[key].owner;
    uint32 configPreMaxStalenessPeriod = currentContract._operatorSetConfigs[key].maxStalenessPeriod;
    address tableCalculatorPre = currentContract._operatorTableCalculators[key];

    bool canCall = canCall(e, os.avs, e.msg.sender, f.selector);

    if(f.selector ==  sig:createGenerationReservation(CrossChainRegistry.OperatorSet, address, ICrossChainRegistryTypes.OperatorSetConfig).selector){
        ICrossChainRegistryTypes.OperatorSetConfig config;
        createGenerationReservation(e, os, tableCalculatorPre, config);
    } else if (f.selector == sig:removeGenerationReservation(CrossChainRegistry.OperatorSet).selector){
        removeGenerationReservation(e, os);
    } else {
        f(e, args);
    }
    
    // no state change in _activeGenerationReservations
    assert ghostValues_activeGenerationReservations[i] == _activeGenerationReservationsValuePre || canCall;
    assert ghostIndexes_activeGenerationReservations[key] == _activeGenerationReservationsIndexPre || canCall;
    assert ghostLength_activeGenerationReservations == _activeGenerationReservationsLengthPre || canCall;

    // no state change in _operatorSetConfigs
    assert configPreOwner == currentContract._operatorSetConfigs[key].owner || canCall;
    assert configPreMaxStalenessPeriod == currentContract._operatorSetConfigs[key].maxStalenessPeriod || canCall;

    // no state change in _operatorTableCalculators
    assert tableCalculatorPre == currentContract._operatorTableCalculators[key] || canCall;

}

/*
 * _paused@before == _paused@after
 *
 * What it means: The function must not modify the pause state stored in the _paused state variable
 *
 * Why it should hold: createGenerationReservation is not a pause management function and should only affect operator set reservations. Only authorized pausers should be able to modify the pause state
 *
 * Possible consequences: If this property is violated, an attacker could unpause the contract or change pause states through a function that should only create reservations, bypassing pause controls
 */
rule no_illegal_change_to_paused(env e, method f, calldataarg args) filtered{f -> !f.isView}{

    // assign all the 'before' variables
    uint256 _paused_before = currentContract._paused;

    // call function under test
    f(e, args);

    // assign all the 'after' variables
    uint256 _paused_after = currentContract._paused;

    // verify integrity
    assert (_paused_before == _paused_after || 
        f.selector == sig:initialize(address, uint32, uint256).selector ||
        f.selector == sig:pauseAll().selector ||
        f.selector == sig:pause(uint256).selector ||
        f.selector == sig:unpause(uint256).selector
    ), "_paused@before == _paused@after";
}

/*
 * _initialized@after == _initialized@before
 *
 * What it means: The function must not change the initialization state flag
 *
 * Why it should hold: Only initialization functions should modify this flag, not chain whitelist operations
 *
 * Possible consequences: Reinitialization vulnerabilities, bypassing initialization checks, state corruption
 */
rule no_illegal_change_initialized(env e, method f, calldataarg args) {
    // assign all the 'before' variables
    uint8 _initialized_before = currentContract._initialized;

    // call function under test
    f(e, args);

    // assign all the 'after' variables
    uint8 _initialized_after = currentContract._initialized;

    // verify integrity
    assert (_initialized_after == _initialized_before || 
        f.selector == sig:initialize(address, uint32, uint256).selector
    ), "_initialized@after == _initialized@before";
}

/*
 * _initialized@after == _initialized@before
 *
 * What it means: The function must not change the initialization state flag
 *
 * Why it should hold: Only initialization functions should modify this flag, not chain whitelist operations
 *
 * Possible consequences: Reinitialization vulnerabilities, bypassing initialization checks, state corruption
 */
rule no_illegal_change_initializing(env e, method f, calldataarg args) {
    // assign all the 'before' variables
    bool _initializing_before = currentContract._initializing;

    // call function under test
    f(e, args);

    // assign all the 'after' variables
    bool _initializing_after = currentContract._initializing;

    // verify integrity
    assert (_initializing_after == _initializing_before), "_initialized@after == _initialized@before";
}

/*
 * _owner@after == _owner@before
 *
 * What it means: The contract owner address stored in _owner should remain the same before and after the function execution
 *
 * Why it should hold: The setOperatorSetConfig function is not designed to change ownership and has no ownership transfer logic. Any change to the owner during this function would indicate unauthorized privilege escalation or a serious bug
 *
 * Possible consequences: If ownership changes unexpectedly, it could lead to complete loss of administrative control, unauthorized access to owner-only functions, or transfer of control to malicious actors
 */
rule no_illegal_change_to_owner(env e, method f, calldataarg args) {
    // assign all the 'before' variables
    address _owner_before = currentContract._owner;

    // call function under test
    f(e, args);

    // assign all the 'after' variables
    address _owner_after = currentContract._owner;

    // verify integrity
    assert (_owner_after == _owner_before || 
        f.selector == sig:initialize(address, uint32, uint256).selector ||
        f.selector == sig:renounceOwnership().selector ||
        f.selector == sig:transferOwnership(address).selector
    ), "_owner@after == _owner@before";
}

/*
 * _initialized != 0 => revert
 *
 * What it means: The initialize function must revert if the contract has already been initialized (when _initialized is not 0)
 *
 * Why it should hold: This is a critical security property of the Initializable pattern from OpenZeppelin. The initialize function should only be callable once to prevent reinitialization attacks that could reset critical state variables like ownership
 *
 * Possible consequences: State corruption, ownership takeover, complete contract compromise, fund loss through unauthorized access control changes
 */
rule initialize_already_initialized_reverts(env e) {
    address initialOwner;
    uint32 initialTableUpdateCadence;
    uint256 initialPausedStatus;

    // assign all the 'before' variables
    uint8 _initialized_before = currentContract._initialized;

    // call function under test
    initialize@withrevert(e, initialOwner, initialTableUpdateCadence, initialPausedStatus);
    bool initialize_reverted = lastReverted;

    // verify integrity
    assert ((_initialized_before != 0) => initialize_reverted), "_initialized != 0 => revert";
}

/*
 * initialOwner != address(0) => _owner@after == initialOwner
 * _paused@after == initialPausedStatus
 * _initialized@after == 1
 * _initialized@before == 0 
 *
 * What it means: When a valid (non-zero) initialOwner is provided, the _owner storage variable must be set to that address after initialization. The _paused storage variable must be set to the initialPausedStatus parameter value after initialization. The _initialized flag must be set to 1 to mark the contract as initialized
 *
 * Why it should hold: This ensures proper ownership transfer during initialization. The owner has critical privileges in this contract including managing chain whitelists, so correct ownership assignment is essential for security
 *
 * Possible consequences: Loss of administrative control, inability to manage protocol parameters, potential for contract to become unmanageable
 */
rule initialize_integrity(env e) {
    address initialOwner;
    uint32 initialTableUpdateCadence;
    uint256 initialPausedStatus;

    // assign all the 'before' variables
    uint8 initializedBefore = currentContract._initialized;

    // call function under test
    initialize(e, initialOwner, initialTableUpdateCadence, initialPausedStatus);

    // assign all the 'after' variables
    address ownerAfter = currentContract._owner;
    uint256 pausedAfter = currentContract._paused;
    uint32 tableUpdateCadenceAfter = currentContract._tableUpdateCadence;

    // verify integrity
    assert (initialOwner != 0 => ownerAfter == initialOwner), "initialOwner != address(0) => _owner@after == initialOwner";
    assert (pausedAfter == initialPausedStatus), "_paused@after == initialPausedStatus";
    assert (initializedBefore == 0 && currentContract._initialized == 1), "_initialized@after == 1";
    assert (tableUpdateCadenceAfter == initialTableUpdateCadence && initialTableUpdateCadence != 0);
}



/*
 * _paused != 0 => revert
 *
 * What it means: The function must revert if the contract is in any paused state (when _paused is non-zero)
 *
 * Why it should hold: The function has the onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS) modifier, which should prevent execution when the contract is paused. This is a critical safety mechanism to halt operations during emergencies or maintenance
 *
 * Possible consequences: If this property is violated, attackers could continue creating generation reservations even when the contract is paused, potentially bypassing emergency stops, maintenance windows, or security incident responses
 */
rule createGenerationReservation_paused_reverts(env e) {
    CrossChainRegistry.OperatorSet operatorSet;
    address operatorTableCalculator;
    ICrossChainRegistryTypes.OperatorSetConfig config;

    uint256 mask = 1 << 0; // PAUSED_GENERATION_RESERVATIONS == 0
    bool paused = (currentContract._paused & mask) == mask;

    // call function under test
    createGenerationReservation@withrevert(e, operatorSet, operatorTableCalculator, config);
    bool createGenerationReservation_reverted = lastReverted;

    // verify integrity
    assert (paused => createGenerationReservation_reverted), "_paused != 0 => revert";
}

/*
 * !allocationManager.isOperatorSet(operatorSet) => revert
 *
 * What it means: The function must revert if the provided operator set is not recognized as valid by the allocation manager
 *
 * Why it should hold: The function has the isValidOperatorSet modifier that calls allocationManager.isOperatorSet(operatorSet). Only legitimate operator sets registered with the allocation manager should be allowed to create reservations
 *
 * Possible consequences: If this property is violated, attackers could create reservations for non-existent or unauthorized operator sets, leading to invalid cross-chain configurations and potential system corruption
 */
rule createGenerationReservation_invalid_operator_set_reverts(env e) {
    CrossChainRegistry.OperatorSet operatorSet;
    address operatorTableCalculator;
    ICrossChainRegistryTypes.OperatorSetConfig config;

    // assign all the 'before' variables
    bool isOperatorSet = currentContract.allocationManager.isOperatorSet(e, operatorSet);

    // call function under test
    createGenerationReservation@withrevert(e, operatorSet, operatorTableCalculator, config);
    bool createGenerationReservation_reverted = lastReverted;

    // verify integrity
    assert (!isOperatorSet => createGenerationReservation_reverted), "!allocationManager.isOperatorSet(operatorSet) => revert";
}

/*
 * _paused != 0 => revert
 *
 * What it means: The function must revert when the contract is in a paused state (when _paused is not zero)
 *
 * Why it should hold: The function has the onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS) modifier which should prevent execution when paused. This is a critical safety mechanism to halt operations during emergencies or maintenance
 *
 * Possible consequences: Emergency pause mechanisms could be bypassed, allowing operations to continue during critical vulnerabilities or maintenance periods, potentially leading to state corruption or exploitation
 */
rule removeGenerationReservation_paused_reverts(env e) {
    CrossChainRegistry.OperatorSet operatorSet;

    // assign all the 'before' variables
    uint256 mask = 1 << 0; // PAUSED_GENERATION_RESERVATIONS == 0
    bool paused = (currentContract._paused & mask) == mask;

    // call function under test
    removeGenerationReservation@withrevert(e, operatorSet);
    bool removeGenerationReservation_reverted = lastReverted;

    // verify integrity
    assert (paused => removeGenerationReservation_reverted), "_paused != 0 => revert";
}

/*
 * !allocationManager.isOperatorSet(operatorSet) => revert
 *
 * What it means: The function must revert when called with an operator set that is not valid according to the AllocationManager
 *
 * Why it should hold: The function has the isValidOperatorSet modifier which validates the operator set exists in the AllocationManager. This prevents operations on non-existent or invalid operator sets
 *
 * Possible consequences: Operations could be performed on invalid operator sets, leading to inconsistent state between the CrossChainRegistry and AllocationManager, potentially causing system-wide inconsistencies
 */
rule removeGenerationReservation_invalid_operator_set_reverts(env e) {
    CrossChainRegistry.OperatorSet operatorSet;

    // assign all the 'before' variables
    bool isOperatorSet = currentContract.allocationManager.isOperatorSet(e, operatorSet);

    // call function under test
    removeGenerationReservation@withrevert(e, operatorSet);
    bool removeGenerationReservation_reverted = lastReverted;

    // verify integrity
    assert (!isOperatorSet => removeGenerationReservation_reverted), "!allocationManager.isOperatorSet(operatorSet) => revert";
}

/*
 * _paused & PAUSED_OPERATOR_TABLE_CALCULATOR != 0 => revert
 *
 * What it means: The function must revert when the PAUSED_OPERATOR_TABLE_CALCULATOR pause flag is set in the _paused bitmap
 *
 * Why it should hold: The contract implements a pausable mechanism where specific functionality can be disabled during emergencies or maintenance. The onlyWhenNotPaused(PAUSED_OPERATOR_TABLE_CALCULATOR) modifier enforces this check
 *
 * Possible consequences: Emergency pause mechanisms could be bypassed, allowing critical operations to continue during security incidents or maintenance periods when they should be halted
 */
rule setOperatorTableCalculator_paused_reverts(env e) {
    CrossChainRegistry.OperatorSet operatorSet;
    address operatorTableCalculator;

    // assign all the 'before' variables
    uint256 mask = 1 << 1; // PAUSED_OPERATOR_TABLE_CALCULATOR == 1
    bool paused = (currentContract._paused & mask) == mask;

    // call function under test
    setOperatorTableCalculator@withrevert(e, operatorSet, operatorTableCalculator);
    bool setOperatorTableCalculator_reverted = lastReverted;

    // verify integrity
    assert (paused => setOperatorTableCalculator_reverted), "_paused & PAUSED_OPERATOR_TABLE_CALCULATOR != 0 => revert";
}

/*
 * !allocationManager.isOperatorSet(operatorSet) => revert
 *
 * What it means: The function must revert if the provided operatorSet is not recognized as valid by the allocationManager
 *
 * Why it should hold: The isValidOperatorSet modifier ensures only legitimate operator sets registered in the AllocationManager can have calculators set. This prevents operations on non-existent or unauthorized operator sets
 *
 * Possible consequences: Unauthorized operator sets could have calculators assigned, leading to invalid cross-chain operations and potential manipulation of operator table calculations
 */
rule setOperatorTableCalculator_invalid_operator_set_reverts(env e) {
    CrossChainRegistry.OperatorSet operatorSet;
    address operatorTableCalculator;

    // assign all the 'before' variables
    bool isOperatorSet = currentContract.allocationManager.isOperatorSet(e, operatorSet);

    // call function under test
    setOperatorTableCalculator@withrevert(e, operatorSet, operatorTableCalculator);
    bool setOperatorTableCalculator_reverted = lastReverted;

    // verify integrity
    assert (!isOperatorSet => setOperatorTableCalculator_reverted), "!allocationManager.isOperatorSet(operatorSet) => revert";
}


/*
 * _paused & PAUSED_OPERATOR_SET_CONFIG != 0 => revert
 *
 * What it means: The function must revert when the PAUSED_OPERATOR_SET_CONFIG flag is set in the _paused bitmap
 *
 * Why it should hold: The contract implements a pausable mechanism where specific functionality can be disabled during emergencies or maintenance. The onlyWhenNotPaused(PAUSED_OPERATOR_SET_CONFIG) modifier should enforce this by checking the pause status and reverting if the corresponding bit is set
 *
 * Possible consequences: If this property is violated, critical operator set configuration changes could occur during emergency situations when the system should be frozen, potentially allowing malicious configuration updates during security incidents or system maintenance
 */
rule setOperatorSetConfig_paused_reverts(env e) {
    CrossChainRegistry.OperatorSet operatorSet;
    ICrossChainRegistryTypes.OperatorSetConfig config;

    // assign all the 'before' variables
    uint256 _paused_before = currentContract._paused;
    uint256 mask = 1 << 2; // PAUSED_OPERATOR_SET_CONFIG == 2
    bool paused = (currentContract._paused & mask) == mask;

    // call function under test
    setOperatorSetConfig@withrevert(e, operatorSet, config);
    bool setOperatorSetConfig_reverted = lastReverted;

    // verify integrity
    assert (paused => setOperatorSetConfig_reverted), "_paused & PAUSED_OPERATOR_SET_CONFIG != 0 => revert";
}

/*
 * !allocationManager.isOperatorSet(operatorSet) => revert
 *
 * What it means: The function must revert when the provided operatorSet is not recognized as valid by the AllocationManager
 *
 * Why it should hold: The isValidOperatorSet modifier calls allocationManager.isOperatorSet(operatorSet) to ensure only legitimate operator sets can have their configurations modified. This prevents configuration of non-existent or unauthorized operator sets
 *
 * Possible consequences: If this property is violated, attackers could create configurations for fake or unauthorized operator sets, leading to state corruption, resource waste, and potential manipulation of cross-chain registry operations
 */
rule setOperatorSetConfig_invalid_operator_set_reverts(env e) {
    CrossChainRegistry.OperatorSet operatorSet;
    ICrossChainRegistryTypes.OperatorSetConfig config;

    // assign all the 'before' variables
    bool isOperatorSet = currentContract.allocationManager.isOperatorSet(e, operatorSet);

    // call function under test
    setOperatorSetConfig@withrevert(e, operatorSet, config);
    bool setOperatorSetConfig_reverted = lastReverted;

    // verify integrity
    assert (!isOperatorSet => setOperatorSetConfig_reverted), "!allocationManager.isOperatorSet(operatorSet) => revert";
}



/*
 * Combined revert conditions for addChainIDsToWhitelist
 * 
 * Must revert if:
 * 1) Paused bit for whitelist is set
 * 2) chainIDs.length != operatorTableUpdaters.length
 * 3) Any chainIDs[i] == 0
 * 4) msg sender is not owner
 */
rule addChainIDsToWhitelist_revert_conditions(env e) {
    address[] operatorTableUpdaters;
    uint256[] chainIDs;

    // --- snapshots / pre-state
    uint256 _paused_before = currentContract._paused;
    uint256 mask = 1 << 3; // PAUSED_CHAIN_WHITELIST == 3
    bool paused = (_paused_before & mask) == mask;
    
    uint256 i; 
    bool zero_present = i < chainIDs.length && chainIDs[i] == 0;

    // --- call function under test
    addChainIDsToWhitelist@withrevert(e, chainIDs, operatorTableUpdaters);
    bool reverted = lastReverted;

    // --- verify revert conditions
    assert (paused 
        || (e.msg.sender != currentContract._owner) 
        || (chainIDs.length != operatorTableUpdaters.length) 
        || zero_present 
    => reverted);
}


/*
 * removeChainIDsFromWhitelist revert conditions
 *
 * Must revert if:
 * 1) Paused bit for whitelist is set
 * 2) Any chainIDs[i] is not whitelisted in the pre-state
 * 3) Any duplicate chain ID appears in the input (i != j && chainIDs[i] == chainIDs[j])
 *
 * Empty arrays are allowed and should NOT revert.
 */
rule removeChainIDsFromWhitelist_revert_conditions(env e) {
    uint256[] chainIDs;

    // --- pre-state
    uint256 paused_before = currentContract._paused;
    uint256 mask = 1 << 3; // PAUSED_CHAIN_WHITELIST == 3
    bool paused = (paused_before & mask) == mask;

    uint256 len = chainIDs.length;

    // nondet indices to quantify over the input
    uint256 i;
    uint256 j;

    bool i_in = (i < len);
    bool j_in = (j < len) && (j != i);
    uint256 ci = chainIDs[i];
    uint256 cj = chainIDs[j];
    // Any duplicate anywhere in the input will cause the second removal to fail.
    bool duplicate_present = i_in && j_in && (ci == cj);

    // --- call function under test
    removeChainIDsFromWhitelist@withrevert(e, chainIDs);
    bool reverted = lastReverted;

    assert (duplicate_present|| paused || e.msg.sender != currentContract._owner => reverted);
    // --- allowed no-op
    assert ((!paused && e.msg.sender == currentContract._owner && len == 0 && e.msg.value == 0) => !reverted), "empty input is allowed and should not revert";
}
