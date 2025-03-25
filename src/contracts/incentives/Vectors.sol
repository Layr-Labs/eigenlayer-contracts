// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

// @notice Packed struct, takes only one slot in storage
struct VectorEntry {
    address key;
    uint96 value;
}

// @notice A list of vector entries paired with a mapping to track which keys the list contains
struct LinearVector {
    VectorEntry[] vector;
    mapping(address => bool) isInVector;
}

// @notice Library for managing entries in a LinearVector. Enforces the behavior of no duplicate keys in the vector.
library LinearVectorOps {

    error CannotAddDuplicateKey();
    error KeyNotInVector();
    error IncorrectIndexInput();

    // TODO: figure out if there is a way to identify the relevant vectorStorage in this event, in the case of multiple vectorStorage structs in one contract
    event EntryAddedToVector(address indexed newKey, uint96 value);
    event KeyRemovedFromVector(address indexed keyRemoved);

    function addEntry(LinearVector storage vectorStorage, VectorEntry memory newEntry) internal {
        require(!vectorStorage.isInVector[newEntry.key], CannotAddDuplicateKey());
        emit EntryAddedToVector(newEntry.key, newEntry.value);
        vectorStorage.isInVector[newEntry.key] = true;
        vectorStorage.vector.push(newEntry);
    }

    function removeKey(LinearVector storage vectorStorage, address keyToRemove, uint256 indexOfKey) internal {
        require(vectorStorage.isInVector[keyToRemove], KeyNotInVector());
        require(vectorStorage.vector[indexOfKey].key == keyToRemove, IncorrectIndexInput());
        // swap and pop
        emit KeyRemovedFromVector(keyToRemove);
        vectorStorage.isInVector[keyToRemove] = false;
        vectorStorage.vector[indexOfKey] = vectorStorage.vector[vectorStorage.vector.length - 1];
        vectorStorage.vector.pop();
    }

    function findKeyIndex(LinearVector storage vectorStorage, address key) internal view returns (uint256) {
        require(vectorStorage.isInVector[key], KeyNotInVector());
        uint256 length = vectorStorage.vector.length;
        uint256 index = 0;
        for (; index < length; ++index) {
            if (key == vectorStorage.vector[index].key) {
                break;
            }
        }
        return index;
    }
}






















