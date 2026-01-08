// GHOST COPIES:
// For every storage variable we add a ghost field that is kept synchronized by hooks.
// The ghost fields can be accessed by the spec, even inside quantifiers.

// ghost field for the values array
ghost mapping(mathint => bytes32) ghostValues_activeGenerationReservations {
    init_state axiom forall mathint x. ghostValues_activeGenerationReservations[x] == to_bytes32(0);
}
// ghost field for the indexes map
ghost mapping(bytes32 => uint256) ghostIndexes_activeGenerationReservations {
    init_state axiom forall bytes32 x. ghostIndexes_activeGenerationReservations[x] == 0;
}
// ghost field for the length of the values array (stored in offset 0)
ghost uint256 ghostLength_activeGenerationReservations {
    init_state axiom ghostLength_activeGenerationReservations == 0;
    // assumption: it's infeasible to grow the list to these many elements.
    axiom ghostLength_activeGenerationReservations < max_uint256;
}

// HOOKS
// Store hook to synchronize ghostLength_activeGenerationReservations with the length of the _activeGenerationReservations._inner._values array.
hook Sstore currentContract._activeGenerationReservations._inner._values.length uint256 newLength {
    ghostLength_activeGenerationReservations = newLength;
}
// Store hook to synchronize ghostValues_activeGenerationReservations array with _activeGenerationReservations._inner._values.
hook Sstore currentContract._activeGenerationReservations._inner._values[INDEX uint256 index] bytes32 newValue {
    ghostValues_activeGenerationReservations[index] = newValue;
}
// Store hook to synchronize ghostIndexes_activeGenerationReservations array with _activeGenerationReservations._inner._indexes.
hook Sstore currentContract._activeGenerationReservations._inner._indexes[KEY bytes32 value] uint256 newIndex {
    ghostIndexes_activeGenerationReservations[value] = newIndex;
}

// Load hook to synchronize ghostLength_activeGenerationReservations with the length of the _activeGenerationReservations._inner._values array.
hook Sload uint256 length currentContract._activeGenerationReservations._inner._values.length {
    require ghostLength_activeGenerationReservations == length;
}
hook Sload bytes32 value currentContract._activeGenerationReservations._inner._values[INDEX uint256 index] {
    require ghostValues_activeGenerationReservations[index] == value;
}
hook Sload uint256 index currentContract._activeGenerationReservations._inner._indexes[KEY bytes32 value] {
    require ghostIndexes_activeGenerationReservations[value] == index;
}

// INVARIANTS
//  This is the main invariant stating that the indexes and values always match:
//        values[indexes[v] - 1] = v for all values v in the _activeGenerationReservations
//    and indexes[values[i]] = i+1 for all valid indexes i.

invariant _activeGenerationReservationsInvariant()
    (forall uint256 index. 0 <= index && index < ghostLength_activeGenerationReservations => to_mathint(ghostIndexes_activeGenerationReservations[ghostValues_activeGenerationReservations[index]]) == index + 1)
    && (forall bytes32 value. ghostIndexes_activeGenerationReservations[value] == 0 ||
         (ghostValues_activeGenerationReservations[ghostIndexes_activeGenerationReservations[value] - 1] == value && ghostIndexes_activeGenerationReservations[value] >= 1 && ghostIndexes_activeGenerationReservations[value] <= ghostLength_activeGenerationReservations));

