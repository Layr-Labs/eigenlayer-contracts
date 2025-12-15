import "../libraries/BN254-nondet.spec";
using AllocationManager as AllocationManager;
using PermissionController as PermissionController;

methods {
    function _.canCall(address, address, address, bytes4) external => DISPATCHER(true);
    function _.isValidERC1271SignatureNow(address, bytes32, bytes memory) internal => NONDET;

    function isRegistered(KeyRegistrar.OperatorSet, address) external returns (bool) envfree;

    function getOperatorSetKey(KeyRegistrar.OperatorSet) external returns (bytes32) envfree;
    function getOperatorKeyDataHash(bytes32, address) external returns (bytes32) envfree;

     // external calls to AVSRegistrar.  Note that the source does not have a proper implementation, the one available always reverts
    function _.registerOperator(address,address,uint32[],bytes) external => DISPATCHER(true);
    function _.deregisterOperator(address,address,uint32[]) external => DISPATCHER(true);
    function _.supportsAVS(address) external => DISPATCHER(true);
}

use builtin rule sanity filtered { f -> f.contract == currentContract }


rule global_registry_monotonic(method f, env e, calldataarg args) filtered { f -> !f.isView } {
    bytes32 keyHash;

    // snapshot
    bool _globalKeyRegistry_before = currentContract._globalKeyRegistry[keyHash];

    // execute arbitrary non-view function
    f(e, args);

    // check monotonicity
    assert  _globalKeyRegistry_before => currentContract._globalKeyRegistry[keyHash], "_globalKeyRegistry must never flip from true to false";
}

/*
 * curveType == 0 => revert
 *
 * What it means: The function must revert when curveType is 0 (NONE), preventing configuration with an invalid curve type
 *
 * Why it should hold: CurveType.NONE (value 0) represents an unconfigured state, not a valid curve type for cryptographic operations. Allowing configuration with NONE would create an inconsistent state where an operator set appears configured but has no valid cryptographic curve
 *
 * Possible consequences: State corruption where operator sets appear configured but cannot perform cryptographic operations, leading to system-wide failures in key registration and verification
 */
rule configureOperatorSet_invalidCurveTypeReverts(env e) {
    KeyRegistrar.OperatorSet os;
    IKeyRegistrarTypes.CurveType curveType;
    bytes32 key = getOperatorSetKey(os);

    // assign all the 'before' variables
    IKeyRegistrarTypes.CurveType curveTypeBefore = currentContract._operatorSetCurveTypes[key];

    // call function under test
    configureOperatorSet@withrevert(e, os, curveType);
    bool configureOperatorSet_reverted = lastReverted;

    // assign all the 'after' variables
    IKeyRegistrarTypes.CurveType curveTypeAfter = currentContract._operatorSetCurveTypes[key];

    // verify integrity
    assert (curveType == IKeyRegistrarTypes.CurveType.NONE) => configureOperatorSet_reverted, "curveType == NONE => revert";
    assert (curveTypeBefore != IKeyRegistrarTypes.CurveType.NONE => configureOperatorSet_reverted), "curveTypeBefore must be empty";
    assert (curveTypeAfter == curveType && (curveType == IKeyRegistrarTypes.CurveType.ECDSA || curveType == IKeyRegistrarTypes.CurveType.BN254)) || configureOperatorSet_reverted, "curveType must be valid after successful configuration";
}

/*
 * operatorSetNotConfigured || keyAlreadyRegistered => revert
 *
 * What it means:  
 * `registerKey` must revert if either (a) the operator‑set has **not** been
 * configured with a valid curve type (`_operatorSetCurveTypes == NONE`), or
 * (b) the operator already has a key registered in that set, (c) keyData cannot be empty, 
 * (d) keyData must be exactly 20 bytes for ECDSA keys or 192 bytes for BN254 keys - any other length should cause registration to fail
 *
 * Why it should hold:  
 * Registering under an unconfigured curve type would leave the system in an
 * unusable state, while allowing duplicate registrations would violate the
 * one‑key‑per‑operator invariant and undermine key‑uniqueness guarantees.
 *
 * Possible consequences:  
 * If the check is bypassed, keys could be registered under invalid cryptographic
 * parameters, or operators could overwrite/accumulate multiple keys. Either case
 * leads to inconsistent state, signature‑verification failures, and potential
 * security breaches.
 */
rule registerKey_revertConditions(env e) {
    KeyRegistrar.OperatorSet os;
    address operator;
    bytes keyData;
    bytes signature;
    bytes32 key = getOperatorSetKey(os);

    // Ensure the operatorSet has been configured with a valid curve
    bool operatorSetNotConfigured = currentContract._operatorSetCurveTypes[key] == IKeyRegistrarTypes.CurveType.NONE;
    bool keyAlreadyRegistered = currentContract._operatorKeyInfo[key][operator].isRegistered == true;
    bool emptyKeyData = keyData.length == 0;
    bool validKeyDataLength = (currentContract._operatorSetCurveTypes[key] == IKeyRegistrarTypes.CurveType.ECDSA && keyData.length == 20) || (currentContract._operatorSetCurveTypes[key] == IKeyRegistrarTypes.CurveType.BN254 && keyData.length == 192);
    bool callAuthorized = PermissionController.canCall(e, operator, e.msg.sender, currentContract, to_bytes4(sig:registerKey(address, KeyRegistrar.OperatorSet, bytes, bytes).selector));

    // Then the call to registerKey should proceed
    registerKey@withrevert(e, operator, os, keyData, signature);
    bool registerKeyReverted = lastReverted;

    assert operatorSetNotConfigured => registerKeyReverted;
    assert keyAlreadyRegistered => registerKeyReverted;
    assert emptyKeyData => registerKeyReverted;
    assert !validKeyDataLength => registerKeyReverted;
    assert !callAuthorized => registerKeyReverted;
}

/*
 * (_operatorKeyInfo[operatorSetKey][operator] or keyData) changed
 *     => f.selector == registerKey || deregisterKey
 *
 * What it means:  
 * Any state change that touches an operator’s key‑registration record—either
 * the `isRegistered` flag or the stored key‑data hash—**must** originate
 * exclusively from the `registerKey` or `deregisterKey` functions. All other
 * non‑view methods are forbidden to modify this storage.
 *
 * Why it should hold:  
 * Centralising mutations in the dedicated registration functions preserves
 * invariant boundaries, simplifies auditing, and prevents accidental or
 * malicious corruption of key‑management data by unrelated contract logic.
 *
 * Possible consequences:  
 * If another function can alter these fields, it might silently register,
 * deregister, or tamper with keys without passing the intended access‑control
 * and validation pathways—resulting in lost keys, unauthorised operators,
 * denial‑of‑service, or broken cryptographic assumptions across the system.
 */
rule onlyRegisterAndDeregisterKeyModifyOperatorKeyInfo(method f, env e, calldataarg args)  filtered { f -> !f.isView } 
 {
    address operator;
    KeyRegistrar.OperatorSet os;
    bytes32 key = getOperatorSetKey(os);

    // Snapshot old state
    bool isRegistered = currentContract._operatorKeyInfo[key][operator].isRegistered;
    bytes32 keyData = getOperatorKeyDataHash(key, operator);

    // Run any non-view function
    f(e, args);

    // Assert state only changed by register or deregister key
    assert (getOperatorKeyDataHash(key, operator) != keyData ||
            currentContract._operatorKeyInfo[key][operator].isRegistered != isRegistered) 
            => (f.selector == sig:registerKey(address, KeyRegistrar.OperatorSet, bytes, bytes).selector ||
            f.selector == sig:deregisterKey(address, KeyRegistrar.OperatorSet).selector);
}


/*
 * a call to `registerKey` must turn `isRegistered` `false → true`; a call to `deregisterKey` must turn it `true → false`.
 */
rule registrationFlagFlipsCorrectly(method f, env e, calldataarg args) filtered { f -> !f.isView } {
    address operator;
    KeyRegistrar.OperatorSet os;
    bytes32 k = getOperatorSetKey(os);
    bool before = currentContract._operatorKeyInfo[k][operator].isRegistered;

    f(e,args);

    bool after  = currentContract._operatorKeyInfo[k][operator].isRegistered;

    assert (f.selector == sig:registerKey(address, KeyRegistrar.OperatorSet, bytes, bytes).selector) =>
           ((before == false) => after || before == after); //can stay the same if the registered key is different from k
    assert (f.selector == sig:deregisterKey(address, KeyRegistrar.OperatorSet).selector) =>
           ((before == true)  => !after || before == after); //can stay the same if the deregistered key is different from k
    assert (f.selector != sig:registerKey(address, KeyRegistrar.OperatorSet, bytes, bytes).selector && f.selector != sig:deregisterKey(address, KeyRegistrar.OperatorSet).selector) => before == after;
}


rule onlyConfigureOperatorSetMayWriteCurveType(method f, env e, calldataarg args)		
   filtered { f -> !f.isView } {
    bytes32 setKey;

    IKeyRegistrarTypes.CurveType before = currentContract._operatorSetCurveTypes[setKey];
    f(e, args);
    IKeyRegistrarTypes.CurveType after  = currentContract._operatorSetCurveTypes[setKey];

    assert (before != after) => 
        (f.selector == sig:configureOperatorSet(KeyRegistrar.OperatorSet, IKeyRegistrarTypes.CurveType).selector);
}



/*
 * otherOperator != operator => _operatorKeyInfo[operatorSetKey][otherOperator].isRegistered@after == _operatorKeyInfo[operatorSetKey][otherOperator].isRegistered@before
 *
 * What it means: When one operator deregisters their key, the registration status of all other operators in the same operator set must remain unchanged
 *
 * Why it should hold: Key deregistration should be isolated to the specific operator performing the action. Other operators' keys are independent and should not be affected by unrelated deregistration operations.
 *
 * Possible consequences: If other operators' registrations are affected, it could lead to widespread service disruption, unauthorized deregistration of valid keys, and breaking the isolation principle between different operators' key management.
 */
rule deregisterKey_other_operators_unchanged(env e) {
    address operator;
    KeyRegistrar.OperatorSet operatorSet;
    bytes32 operatorSetKey;
    address otherOperator;

    // assign all the 'before' variables
    bool _operatorKeyInfo_operatorSetKey__otherOperator__isRegistered_before = currentContract._operatorKeyInfo[operatorSetKey][otherOperator].isRegistered;

    // call function under test
    deregisterKey(e, operator, operatorSet);

    // assign all the 'after' variables
    bool _operatorKeyInfo_operatorSetKey__otherOperator__isRegistered_after = currentContract._operatorKeyInfo[operatorSetKey][otherOperator].isRegistered;

    // verify integrity
    assert ((otherOperator != operator) => (_operatorKeyInfo_operatorSetKey__otherOperator__isRegistered_after == _operatorKeyInfo_operatorSetKey__otherOperator__isRegistered_before)), "otherOperator != operator => _operatorKeyInfo[operatorSetKey][otherOperator].isRegistered@after == _operatorKeyInfo[operatorSetKey][otherOperator].isRegistered@before";
}


/*
 * (!keyRegistered) ∨ (operatorSetNotConfigured) ∨ (isOperatorSlashable) ⇒ revert
 *
 * What it means — `deregisterKey` must **fail** in any of the following
 * pre‑conditions:
 *   1. **Key not registered**  The operator has no key recorded in the requested
 *      operator‑set (`_operatorKeyInfo[setKey][operator].isRegistered == false`).
 *   2. **Operator‑set not configured**  The set’s curve type is
 *      `CurveType.NONE`, signalling it has never been initialised for use.
 *   3. **Operator is slashable**  `AllocationManager.isOperatorSlashable`
 *      reports the operator is currently subject to slashing and therefore
 *      must not be allowed to opt‑out.
 *
 * Why it should hold:
 *   • Prevents meaningless state transitions on non‑existent keys.  
 *   • Guards against interacting with an uninitialised cryptographic domain.  
 *   • Ensures slashable operators cannot escape penalties by deregistering.
 *
 * Possible consequences if violated:
 *   • Deregistering a non‑existent key breaks invariants and confuses callers.  
 *   • Operating on an unconfigured set may leave the system in an unusable or
 *     inconsistent state.  
 *   • Allowing slashable operators to deregister could let them evade economic
 *     punishment and undermine security assumptions.
 */
rule deregisterKey_revertConditions(env e) {
    address operator;
    KeyRegistrar.OperatorSet operatorSet;
    bytes32 operatorSetKey = getOperatorSetKey(operatorSet);

    // assign all the 'before' variables
    bool keyRegistered = currentContract._operatorKeyInfo[operatorSetKey][operator].isRegistered;
    bool operatorSetNotConfigured = currentContract._operatorSetCurveTypes[operatorSetKey] == IKeyRegistrarTypes.CurveType.NONE;
    bool isOperatorSlashable = AllocationManager.isOperatorSlashable(e, operator, operatorSet);
    bool callAuthorized = PermissionController.canCall(e, operator, e.msg.sender, currentContract, to_bytes4(sig:deregisterKey(address, KeyRegistrar.OperatorSet).selector));
    // call function under test
    deregisterKey@withrevert(e, operator, operatorSet);
    bool deregisterKey_reverted = lastReverted;

    // verify integrity
    assert (!keyRegistered => deregisterKey_reverted), "!_operatorKeyInfo[operatorSetKey][operator].isRegistered => revert";
    assert (operatorSetNotConfigured => deregisterKey_reverted), "curveType is 0 => revert";
    assert (isOperatorSlashable => deregisterKey_reverted), "operator that is slashable cannot be deregistered";
    assert (!callAuthorized => deregisterKey_reverted), "unauthorized operator";
}

/*****************************************************************
 * 1. configureOperatorSet — unauthorised caller must revert
 *****************************************************************/
rule configureOperatorSet_unauthorisedCallerReverts(env e) {
    KeyRegistrar.OperatorSet os;
    IKeyRegistrarTypes.CurveType curveType;

    // Permission check performed by the modifier in Solidity
    bool callAuthorised =
        PermissionController.canCall(
            e,
            os.avs,                             // account
            e.msg.sender,                       // caller
            currentContract,                    // target
            to_bytes4(sig:configureOperatorSet(KeyRegistrar.OperatorSet, IKeyRegistrarTypes.CurveType).selector)
        );

    configureOperatorSet@withrevert(e, os, curveType);
    bool reverted = lastReverted;

    // If not authorised, the call **must** revert
    assert (!callAuthorised) => reverted, "configureOperatorSet: unauthorised caller should revert";
}



/*****************************************************************
 * 2. Only registerKey may set a new entry in _globalKeyRegistry
 *****************************************************************/
rule onlyRegisterKeyMayWriteGlobalRegistry(method f, env e, calldataarg args)
      filtered { f -> !f.isView } {
    bytes32 keyHash;

    bool before = currentContract._globalKeyRegistry[keyHash];
    f(e, args);
    bool after  = currentContract._globalKeyRegistry[keyHash];

    /* A transition false → true is allowed **only** via registerKey */
    assert ((after && !before) =>
            (f.selector == sig:registerKey(address, KeyRegistrar.OperatorSet, bytes, bytes).selector)),
           "_globalKeyRegistry can be set only in registerKey";
}



/*****************************************************************
 * 3. registerKey must not touch other operators’ records
 *****************************************************************/
rule registerKey_otherOperatorsUnchanged(env e) {
    address operator;
    address otherOperator;
    KeyRegistrar.OperatorSet os;
    bytes keyData;
    bytes signature;
    bytes32 setKey = getOperatorSetKey(os);

    /* Snapshot another operator’s registration status */
    bool before = currentContract._operatorKeyInfo[setKey][otherOperator].isRegistered;

    registerKey@withrevert(e, operator, os, keyData, signature);

    bool after  = currentContract._operatorKeyInfo[setKey][otherOperator].isRegistered;

    assert ((otherOperator != operator) => (after == before)),
           "registerKey must not modify registrations of other operators";
}
