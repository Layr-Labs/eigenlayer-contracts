// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin-upgrades-v4.9.0/contracts/utils/math/MathUpgradeable.sol";
import "@openzeppelin-upgrades-v4.9.0/contracts/utils/math/SafeCastUpgradeable.sol";

/**
 * @title Library for handling snapshots as part of allocating and slashing.
 * @notice This library is using OpenZeppelin's CheckpointsUpgradeable library (v4.9.0)
 * and removes structs and functions that are unessential.
 * Interfaces and structs are renamed for clarity and usage (timestamps, etc).
 * Some additional functions have also been added for convenience.
 * @dev This library defines the `History` struct, for snapshotting values as they change at different points in
 * time, and later looking up past values by block number. See {Votes} as an example.
 *
 * To create a history of snapshots define a variable type `Snapshots.History` in your contract, and store a new
 * snapshot for the current transaction block using the {push} function.
 *
 * _Available since v4.5._
 */
library Snapshots {
    struct History {
        Snapshot[] _snapshots;
    }

    struct Snapshot {
        uint32 _key;
        uint224 _value;
    }

    /**
     * @dev Pushes a (`key`, `value`) pair into a History so that it is stored as the snapshot.
     *
     * Returns previous value and new value.
     */
    function push(History storage self, uint32 key, uint224 value) internal returns (uint224, uint224) {
        return _insert(self._snapshots, key, value);
    }

    /**
     * @dev Returns the value in the first (oldest) snapshot with key greater or equal than the search key, or zero if there is none.
     */
    function lowerLookup(History storage self, uint32 key) internal view returns (uint224) {
        uint256 len = self._snapshots.length;
        uint256 pos = _lowerBinaryLookup(self._snapshots, key, 0, len);
        return pos == len ? 0 : _unsafeAccess(self._snapshots, pos)._value;
    }

    /**
     * @dev Returns the value in the last (most recent) snapshot with key lower or equal than the search key, or zero if there is none.
     */
    function upperLookup(History storage self, uint32 key) internal view returns (uint224) {
        uint256 len = self._snapshots.length;
        uint256 pos = _upperBinaryLookup(self._snapshots, key, 0, len);
        return pos == 0 ? 0 : _unsafeAccess(self._snapshots, pos - 1)._value;
    }

    /**
     * @dev Returns the value in the last (most recent) snapshot with key lower or equal than the search key, or zero if there is none.
     *
     * NOTE: This is a variant of {upperLookup} that is optimised to find "recent" snapshot (snapshots with high keys).
     */
    function upperLookupRecent(History storage self, uint32 key) internal view returns (uint224) {
        uint256 len = self._snapshots.length;

        uint256 low = 0;
        uint256 high = len;

        if (len > 5) {
            uint256 mid = len - MathUpgradeable.sqrt(len);
            if (key < _unsafeAccess(self._snapshots, mid)._key) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }

        uint256 pos = _upperBinaryLookup(self._snapshots, key, low, high);

        return pos == 0 ? 0 : _unsafeAccess(self._snapshots, pos - 1)._value;
    }

    /**
     * @dev Returns the value in the most recent snapshot, or zero if there are no snapshots.
     */
    function latest(
        History storage self
    ) internal view returns (uint224) {
        uint256 pos = self._snapshots.length;
        return pos == 0 ? 0 : _unsafeAccess(self._snapshots, pos - 1)._value;
    }

    /**
     * @dev Returns whether there is a snapshot in the structure (i.e. it is not empty), and if so the key and value
     * in the most recent snapshot.
     */
    function latestSnapshot(
        History storage self
    ) internal view returns (bool exists, uint32 _key, uint224 _value) {
        uint256 pos = self._snapshots.length;
        if (pos == 0) {
            return (false, 0, 0);
        } else {
            Snapshot memory ckpt = _unsafeAccess(self._snapshots, pos - 1);
            return (true, ckpt._key, ckpt._value);
        }
    }

    /**
     * @dev Returns the number of snapshot.
     */
    function length(
        History storage self
    ) internal view returns (uint256) {
        return self._snapshots.length;
    }

    /**
     * @dev Pushes a (`key`, `value`) pair into an ordered list of snapshots, either by inserting a new snapshot,
     * or by updating the last one.
     */
    function _insert(Snapshot[] storage self, uint32 key, uint224 value) private returns (uint224, uint224) {
        uint256 pos = self.length;

        if (pos > 0) {
            // Copying to memory is important here.
            Snapshot memory last = _unsafeAccess(self, pos - 1);

            // Snapshot keys must be non-decreasing.
            require(last._key <= key, "Snapshot: decreasing keys");

            // Update or push new snapshot
            if (last._key == key) {
                _unsafeAccess(self, pos - 1)._value = value;
            } else {
                self.push(Snapshot({_key: key, _value: value}));
            }
            return (last._value, value);
        } else {
            self.push(Snapshot({_key: key, _value: value}));
            return (0, value);
        }
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
     * @dev Return the index of the first (oldest) snapshot with key is greater or equal than the search key, or `high` if there is none.
     * `low` and `high` define a section where to do the search, with inclusive `low` and exclusive `high`.
     *
     * WARNING: `high` should not be greater than the array's length.
     */
    function _lowerBinaryLookup(
        Snapshot[] storage self,
        uint32 key,
        uint256 low,
        uint256 high
    ) private view returns (uint256) {
        while (low < high) {
            uint256 mid = MathUpgradeable.average(low, high);
            if (_unsafeAccess(self, mid)._key < key) {
                low = mid + 1;
            } else {
                high = mid;
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

    /**
     *
     *                         ADDITIONAL FUNCTIONS FROM EIGEN-LABS
     *
     */

    /**
     * @dev Returns the value in the last (most recent) snapshot with key lower or equal than the search key, or zero if there is none.
     * This function is a linear search for keys that are close to the end of the array.
     */
    function upperLookupLinear(History storage self, uint32 key) internal view returns (uint224) {
        uint256 len = self._snapshots.length;
        for (uint256 i = len; i > 0; --i) {
            Snapshot storage current = _unsafeAccess(self._snapshots, i - 1);
            if (current._key <= key) {
                return current._value;
            }
        }
        return 0;
    }

    /**
     * @dev Returns the value in the last (most recent) snapshot with key lower or equal than the search key, or zero if there is none.
     * In addition, returns the position of the snapshot in the array.
     *
     * NOTE: That if value != 0 && pos == 0, then that means the value is the first snapshot and actually exists
     * a snapshot DNE iff value == 0 && pos == 0
     */
    function upperLookupWithPos(History storage self, uint32 key) internal view returns (uint224, uint256) {
        uint256 len = self._snapshots.length;
        uint256 pos = _upperBinaryLookup(self._snapshots, key, 0, len);
        return pos == 0 ? (0, 0) : (_unsafeAccess(self._snapshots, pos - 1)._value, pos - 1);
    }

    /**
     * @dev Returns the value in the last (most recent) snapshot with key lower or equal than the search key, or zero if there is none.
     * In addition, returns the position of the snapshot in the array.
     *
     * NOTE: This is a variant of {upperLookup} that is optimised to find "recent" snapshot (snapshots with high keys).
     * NOTE: That if value != 0 && pos == 0, then that means the value is the first snapshot and actually exists
     * a snapshot DNE iff value == 0 && pos == 0 => value == 0
     */
    function upperLookupRecentWithPos(
        History storage self,
        uint32 key
    ) internal view returns (uint224, uint256, uint256) {
        uint256 len = self._snapshots.length;

        uint256 low = 0;
        uint256 high = len;

        if (len > 5) {
            uint256 mid = len - MathUpgradeable.sqrt(len);
            if (key < _unsafeAccess(self._snapshots, mid)._key) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }

        uint256 pos = _upperBinaryLookup(self._snapshots, key, low, high);

        return pos == 0 ? (0, 0, len) : (_unsafeAccess(self._snapshots, pos - 1)._value, pos - 1, len);
    }

    /// @notice WARNING: this function is only used because of the invariant property
    /// that from the current key, all future snapshotted magnitude values are strictly > current value.
    /// Use function with extreme care for other situations.
    function decrementAtAndFutureSnapshots(History storage self, uint32 key, uint224 decrementValue) internal {
        (uint224 value, uint256 pos, uint256 len) = upperLookupRecentWithPos(self, key);

        // if there is no snapshot, return
        if (value == 0 && pos == 0) {
            pos = type(uint256).max;
        }

        while (pos < len) {
            Snapshot storage current = _unsafeAccess(self._snapshots, pos);

            // reverts from underflow. Expected to never happen in our usage
            current._value -= decrementValue;

            unchecked {
                ++pos;
            }
        }
    }
}