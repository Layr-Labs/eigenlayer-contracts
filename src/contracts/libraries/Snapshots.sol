// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin-upgrades/contracts/utils/math/MathUpgradeable.sol";

import "./SlashingLib.sol";

/**
 * @title Library for handling snapshots as part of allocating and slashing.
 * @notice This library is using OpenZeppelin's CheckpointsUpgradeable library (v4.9.0)
 * and removes structs and functions that are unessential.
 * Interfaces and structs are renamed for clarity and usage.
 * Some additional functions have also been added for convenience.
 * @dev This library defines the `DefaultWadHistory` and `DefaultZeroHistory` struct, for snapshotting values as they change at different points in
 * time, and later looking up past values by block number. See {Votes} as an example.
 *
 * To create a history of snapshots define a variable type `Snapshots.DefaultWadHistory` or `Snapshots.DefaultZeroHistory` in your contract,
 * and store a new snapshot for the current transaction block using the {push} function. If there is no history yet, the value is either WAD or 0,
 * depending on the type of History struct used. This is implemented because for the AllocationManager we want the
 * the default value to be WAD(1e18) but when used in the DelegationManager we want the default value to be 0.
 *
 * _Available since v4.5._
 */
library Snapshots {
    struct DefaultWadHistory {
        Snapshot[] _snapshots;
    }

    struct DefaultZeroHistory {
        Snapshot[] _snapshots;
    }

    struct Snapshot {
        uint32 _key;
        uint224 _value;
    }

    error InvalidSnapshotOrdering();

    /**
     * @dev Pushes a (`key`, `value`) pair into a DefaultWadHistory so that it is stored as the snapshot.
     */
    function push(DefaultWadHistory storage self, uint32 key, uint64 value) internal {
        _insert(self._snapshots, key, value);
    }

    /**
     * @dev Pushes a (`key`, `value`) pair into a DefaultZeroHistory so that it is stored as the snapshot.
     * `value` is cast to uint224. Responsibility for the safety of this operation falls outside of this library.
     */
    function push(DefaultZeroHistory storage self, uint32 key, uint256 value) internal {
        _insert(self._snapshots, key, uint224(value));
    }

    /**
     * @dev Return default value of WAD if there are no snapshots for DefaultWadHistory.
     * This is used for looking up maxMagnitudes in the AllocationManager.
     */
    function upperLookup(DefaultWadHistory storage self, uint32 key) internal view returns (uint64) {
        return uint64(_upperLookup(self._snapshots, key, WAD));
    }

    /**
     * @dev Return default value of 0 if there are no snapshots for DefaultZeroHistory.
     * This is used for looking up cumulative scaled shares in the DelegationManager.
     */
    function upperLookup(DefaultZeroHistory storage self, uint32 key) internal view returns (uint256) {
        return _upperLookup(self._snapshots, key, 0);
    }

    /**
     * @dev Returns the value in the most recent snapshot, or WAD if there are no snapshots.
     */
    function latest(
        DefaultWadHistory storage self
    ) internal view returns (uint64) {
        return uint64(_latest(self._snapshots, WAD));
    }

    /**
     * @dev Returns the value in the most recent snapshot, or 0 if there are no snapshots.
     */
    function latest(
        DefaultZeroHistory storage self
    ) internal view returns (uint256) {
        return uint256(_latest(self._snapshots, 0));
    }

    /**
     * @dev Returns the number of snapshots.
     */
    function length(
        DefaultWadHistory storage self
    ) internal view returns (uint256) {
        return self._snapshots.length;
    }

    /**
     * @dev Returns the number of snapshots.
     */
    function length(
        DefaultZeroHistory storage self
    ) internal view returns (uint256) {
        return self._snapshots.length;
    }

    /**
     * @dev Pushes a (`key`, `value`) pair into an ordered list of snapshots, either by inserting a new snapshot,
     * or by updating the last one.
     */
    function _insert(Snapshot[] storage self, uint32 key, uint224 value) private {
        uint256 pos = self.length;

        if (pos > 0) {
            // Validate that inserted keys are always >= the previous key
            Snapshot memory last = _unsafeAccess(self, pos - 1);
            require(last._key <= key, InvalidSnapshotOrdering());

            // Update existing snapshot if `key` matches
            if (last._key == key) {
                _unsafeAccess(self, pos - 1)._value = value;
                return;
            }
        }

        // `key` was not in the list; push as a new entry
        self.push(Snapshot({_key: key, _value: value}));
    }

    /**
     * @dev Returns the value in the last (most recent) snapshot with key lower or equal than the search key, or `defaultValue` if there is none.
     */
    function _upperLookup(
        Snapshot[] storage snapshots,
        uint32 key,
        uint224 defaultValue
    ) private view returns (uint224) {
        uint256 len = snapshots.length;
        uint256 pos = _upperBinaryLookup(snapshots, key, 0, len);
        return pos == 0 ? defaultValue : _unsafeAccess(snapshots, pos - 1)._value;
    }

    /**
     * @dev Returns the value in the most recent snapshot, or `defaultValue` if there are no snapshots.
     */
    function _latest(Snapshot[] storage snapshots, uint224 defaultValue) private view returns (uint224) {
        uint256 pos = snapshots.length;
        return pos == 0 ? defaultValue : _unsafeAccess(snapshots, pos - 1)._value;
    }

    /**
     * @dev Return the index of the last (most recent) snapshot with key lower or equal than the search key, or `high` if there is none.
     * `low` and `high` define a section where to do the search, with inclusive `low` and exclusive `high`.
     *
     * WARNING: `high` should not be greater than the array's length.
     */
    function _upperBinaryLookup(
        Snapshot[] storage self,
        uint32 key,
        uint256 low,
        uint256 high
    ) private view returns (uint256) {
        while (low < high) {
            uint256 mid = MathUpgradeable.average(low, high);
            if (_unsafeAccess(self, mid)._key > key) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }
        return high;
    }

    /**
     * @dev Access an element of the array without performing bounds check. The position is assumed to be within bounds.
     */
    function _unsafeAccess(Snapshot[] storage self, uint256 pos) private pure returns (Snapshot storage result) {
        assembly {
            mstore(0, self.slot)
            result.slot := add(keccak256(0, 0x20), pos)
        }
    }
}
