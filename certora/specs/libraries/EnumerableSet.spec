methods {
}

// GHOST COPIES:
// For every storage variable we add a ghost field that is kept synchronized by hooks.
// The ghost fields can be accessed by the spec, even inside quantifiers.

// ghost field for the values array
ghost mapping(bytes32 => mapping (mathint => bytes32)) ghostValues_pendingSlashIds {
    init_state axiom forall bytes32 key. forall mathint x. ghostValues_pendingSlashIds[key][x] == to_bytes32(0);
}
// ghost field for the indexes map
ghost mapping(bytes32 => mapping (bytes32 => uint256)) ghostIndexes_pendingSlashIds {
    init_state axiom forall bytes32 key. forall bytes32 x. ghostIndexes_pendingSlashIds[key][x] == 0;
}
// ghost field for the length of the values array (stored in offset 0)
ghost mapping(bytes32 => uint256) ghostLength_pendingSlashIds {
    init_state axiom forall bytes32 key. ghostLength_pendingSlashIds[key] == 0;
    // assumption: it's infeasible to grow the list to these many elements.
    axiom forall bytes32 key. ghostLength_pendingSlashIds[key] < max_uint256;
}

// HOOKS
// Store hook to synchronize ghostLength_pendingSlashIds with the length of the set._inner._values array.
hook Sstore currentContract._pendingSlashIds[KEY bytes32 key]._inner._values.length uint256 newLength {
    ghostLength_pendingSlashIds[key] = newLength;
}
// Store hook to synchronize ghostValues_pendingSlashIds array with _pendingSlashIds._inner._values.
hook Sstore currentContract._pendingSlashIds[KEY bytes32 key]._inner._values[INDEX uint256 index] bytes32 newValue {
    ghostValues_pendingSlashIds[key][index] = newValue;
}
// Store hook to synchronize ghostIndexes array with _pendingSlashIds._inner._indexes.
hook Sstore currentContract._pendingSlashIds[KEY bytes32 key]._inner._indexes[KEY bytes32 value] uint256 newIndex {
    ghostIndexes_pendingSlashIds[key][value] = newIndex;
}

// The load hooks can use require to ensure that the ghost field has the same information as the storage.
// The require is sound, since the store hooks ensure the contents are always the same.  However we cannot
// prove that with invariants, since this would require the invariant to read the storage for all elements
// and neither storage access nor function calls are allowed in quantifiers.
//
// By following this simple pattern it is ensured that the ghost state and the storage are always the same
// and that the solver can use this knowledge in the proofs.

// Load hook to synchronize ghostLength_pendingSlashIds with the length of the _pendingSlashIds._inner._values array.
hook Sload uint256 length currentContract._pendingSlashIds[KEY bytes32 key]._inner._values.length {
    require ghostLength_pendingSlashIds[key] == length;
}
hook Sload bytes32 value currentContract._pendingSlashIds[KEY bytes32 key]._inner._values[INDEX uint256 index] {
    require ghostValues_pendingSlashIds[key][index] == value;
}
hook Sload uint256 index currentContract._pendingSlashIds[KEY bytes32 key]._inner._indexes[KEY bytes32 value] {
    require ghostIndexes_pendingSlashIds[key][value] == index;
}

// INVARIANTS

//  This is the main invariant stating that the indexes and values always match:
//        values[indexes[v] - 1] = v for all values v in the _pendingSlashIds
//    and indexes[values[i]] = i+1 for all valid indexes i.

invariant _pendingSlashIdsInvariant()
    (forall bytes32 key. forall uint256 index. 0 <= index && index < ghostLength_pendingSlashIds[key] => to_mathint(ghostIndexes_pendingSlashIds[key][ghostValues_pendingSlashIds[key][index]]) == index + 1)
    && (forall bytes32 key. forall bytes32 value. ghostIndexes_pendingSlashIds[key][value] == 0 ||
         (ghostValues_pendingSlashIds[key][ghostIndexes_pendingSlashIds[key][value] - 1] == value && ghostIndexes_pendingSlashIds[key][value] >= 1 && ghostIndexes_pendingSlashIds[key][value] <= ghostLength_pendingSlashIds[key]));






// ghost field for the values array
ghost mapping(bytes32 => mapping (uint256 => mapping (mathint => bytes32))) ghostValues_pendingStrategiesForSlashId {
    init_state axiom forall bytes32 key. forall uint256 slashId. forall mathint x. ghostValues_pendingStrategiesForSlashId[key][slashId][x] == to_bytes32(0);
}
// ghost field for the indexes map
ghost mapping(bytes32 => mapping (uint256 => mapping (bytes32 => uint256))) ghostIndexes_pendingStrategiesForSlashId {
    init_state axiom forall bytes32 key. forall uint256 slashId. forall bytes32 x. ghostIndexes_pendingStrategiesForSlashId[key][slashId][x] == 0;
}
// ghost field for the length of the values array (stored in offset 0)
ghost mapping(bytes32 => mapping (uint256 => uint256))ghostLength_pendingStrategiesForSlashId {
    init_state axiom forall bytes32 key. forall uint256 slashId. ghostLength_pendingStrategiesForSlashId[key][slashId] == 0;
    // assumption: it's infeasible to grow the list to these many elements.
    axiom forall bytes32 key. forall uint256 slashId.  ghostLength_pendingStrategiesForSlashId[key][slashId] < max_uint256;
}

// HOOKS
// Store hook to synchronize ghostLength_pendingStrategiesForSlashId with the length of the set._inner._values array.
hook Sstore currentContract._pendingStrategiesForSlashId[KEY bytes32 key][KEY uint256 slashId]._inner._values.length uint256 newLength {
    ghostLength_pendingStrategiesForSlashId[key][slashId] = newLength;
}
// Store hook to synchronize ghostValues_pendingStrategiesForSlashId array with _pendingStrategiesForSlashId._inner._values.
hook Sstore currentContract._pendingStrategiesForSlashId[KEY bytes32 key][KEY uint256 slashId]._inner._values[INDEX uint256 index] bytes32 newValue {
    ghostValues_pendingStrategiesForSlashId[key][slashId][index] = newValue;
}
// Store hook to synchronize ghostIndexes array with _pendingStrategiesForSlashId._inner._indexes.
hook Sstore currentContract._pendingStrategiesForSlashId[KEY bytes32 key][KEY uint256 slashId]._inner._indexes[KEY bytes32 value] uint256 newIndex {
    ghostIndexes_pendingStrategiesForSlashId[key][slashId][value] = newIndex;
}

// Load hook to synchronize ghostLength_pendingStrategiesForSlashId with the length of the _pendingStrategiesForSlashId._inner._values array.
hook Sload uint256 length currentContract._pendingStrategiesForSlashId[KEY bytes32 key][KEY uint256 slashId]._inner._values.length {
    require ghostLength_pendingStrategiesForSlashId[key][slashId] == length;
}
hook Sload bytes32 value currentContract._pendingStrategiesForSlashId[KEY bytes32 key][KEY uint256 slashId]._inner._values[INDEX uint256 index] {
    require ghostValues_pendingStrategiesForSlashId[key][slashId][index] == value;
}
hook Sload uint256 index currentContract._pendingStrategiesForSlashId[KEY bytes32 key][KEY uint256 slashId]._inner._indexes[KEY bytes32 value] {
    require ghostIndexes_pendingStrategiesForSlashId[key][slashId][value] == index;
}

// INVARIANTS
//  This is the main invariant stating that the indexes and values always match:
//        values[indexes[v] - 1] = v for all values v in the _pendingStrategiesForSlashId
//    and indexes[values[i]] = i+1 for all valid indexes i.

invariant _pendingStrategiesForSlashIdInvariant()
    (forall bytes32 key. forall uint256 slashId. forall uint256 index. 0 <= index && index < ghostLength_pendingStrategiesForSlashId[key][slashId] => to_mathint(ghostIndexes_pendingStrategiesForSlashId[key][slashId][ghostValues_pendingStrategiesForSlashId[key][slashId][index]]) == index + 1)
    && (forall bytes32 key. forall uint256 slashId. forall bytes32 value. ghostIndexes_pendingStrategiesForSlashId[key][slashId][value] == 0 ||
         (ghostValues_pendingStrategiesForSlashId[key][slashId][ghostIndexes_pendingStrategiesForSlashId[key][slashId][value] - 1] == value && ghostIndexes_pendingStrategiesForSlashId[key][slashId][value] >= 1 && ghostIndexes_pendingStrategiesForSlashId[key][slashId][value] <= ghostLength_pendingStrategiesForSlashId[key][slashId]));





// GHOST COPIES:
// For every storage variable we add a ghost field that is kept synchronized by hooks.
// The ghost fields can be accessed by the spec, even inside quantifiers.

// ghost field for the values array
ghost mapping(mathint => bytes32) ghostValues_pendingOperatorSets {
    init_state axiom forall mathint x. ghostValues_pendingOperatorSets[x] == to_bytes32(0);
}
// ghost field for the indexes map
ghost mapping(bytes32 => uint256) ghostIndexes_pendingOperatorSets {
    init_state axiom forall bytes32 x. ghostIndexes_pendingOperatorSets[x] == 0;
}
// ghost field for the length of the values array (stored in offset 0)
ghost uint256 ghostLength_pendingOperatorSets {
    init_state axiom ghostLength_pendingOperatorSets == 0;
    // assumption: it's infeasible to grow the list to these many elements.
    axiom ghostLength_pendingOperatorSets < max_uint256;
}

// HOOKS
// Store hook to synchronize ghostLength_pendingOperatorSets with the length of the _pendingOperatorSets._inner._values array.
hook Sstore currentContract._pendingOperatorSets._inner._values.length uint256 newLength {
    ghostLength_pendingOperatorSets = newLength;
}
// Store hook to synchronize ghostValues_pendingOperatorSets array with _pendingOperatorSets._inner._values.
hook Sstore currentContract._pendingOperatorSets._inner._values[INDEX uint256 index] bytes32 newValue {
    ghostValues_pendingOperatorSets[index] = newValue;
}
// Store hook to synchronize ghostIndexes_pendingOperatorSets array with _pendingOperatorSets._inner._indexes.
hook Sstore currentContract._pendingOperatorSets._inner._indexes[KEY bytes32 value] uint256 newIndex {
    ghostIndexes_pendingOperatorSets[value] = newIndex;
}

// Load hook to synchronize ghostLength_pendingOperatorSets with the length of the _pendingOperatorSets._inner._values array.
hook Sload uint256 length currentContract._pendingOperatorSets._inner._values.length {
    require ghostLength_pendingOperatorSets == length;
}
hook Sload bytes32 value currentContract._pendingOperatorSets._inner._values[INDEX uint256 index] {
    require ghostValues_pendingOperatorSets[index] == value;
}
hook Sload uint256 index currentContract._pendingOperatorSets._inner._indexes[KEY bytes32 value] {
    require ghostIndexes_pendingOperatorSets[value] == index;
}

// INVARIANTS
//  This is the main invariant stating that the indexes and values always match:
//        values[indexes[v] - 1] = v for all values v in the _pendingOperatorSets
//    and indexes[values[i]] = i+1 for all valid indexes i.

invariant _pendingOperatorSetsInvariant()
    (forall uint256 index. 0 <= index && index < ghostLength_pendingOperatorSets => to_mathint(ghostIndexes_pendingOperatorSets[ghostValues_pendingOperatorSets[index]]) == index + 1)
    && (forall bytes32 value. ghostIndexes_pendingOperatorSets[value] == 0 ||
         (ghostValues_pendingOperatorSets[ghostIndexes_pendingOperatorSets[value] - 1] == value && ghostIndexes_pendingOperatorSets[value] >= 1 && ghostIndexes_pendingOperatorSets[value] <= ghostLength_pendingOperatorSets));
