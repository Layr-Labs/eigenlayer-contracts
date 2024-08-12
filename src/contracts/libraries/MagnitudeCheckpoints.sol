// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin-upgrades-v4.9.0/contracts/utils/math/MathUpgradeable.sol";
import "@openzeppelin-upgrades-v4.9.0/contracts/utils/math/SafeCastUpgradeable.sol";

/**
 * @title Library for handling checkpointed magnitudes as part of allocating and slashing.
 * @notice This library is using OpenZeppelin's CheckpointsUpgradeable library (v4.9.0)
 * and removes structs and functions that are unessential.
 * Interfaces and structs are renamed for clarity and usage (magnitudes, timestamps, etc).
 * Some additional functions have also been added for convenience.
 * @dev This library defines the `History` struct, for checkpointing values as they change at different points in
 * time, and later looking up past values by block number. See {Votes} as an example.
 *
 * To create a history of checkpoints define a variable type `Checkpoints.History` in your contract, and store a new
 * checkpoint for the current transaction block using the {push} function.
 *
 * _Available since v4.5._
 */
library MagnitudeCheckpoints {
    struct History {
        MagnitudeCheckpoint[] _magnitudes;
    }

    struct MagnitudeCheckpoint {
        uint32 _timestamp;
        uint224 _value;
    }

    /**
     * @dev Pushes a (`timestamp`, `value`) pair into a History so that it is stored as the checkpoint.
     *
     * Returns previous value and new value.
     */
    function push(History storage self, uint32 timestamp, uint224 value) internal returns (uint224, uint224) {
        return _insert(self._magnitudes, timestamp, value);
    }

    /**
     * @dev Returns the value in the first (oldest) checkpoint with timestamp greater or equal than the search timestamp, or zero if there is none.
     */
    function lowerLookup(History storage self, uint32 timestamp) internal view returns (uint224) {
        uint256 len = self._magnitudes.length;
        uint256 pos = _lowerBinaryLookup(self._magnitudes, timestamp, 0, len);
        return pos == len ? 0 : _unsafeAccess(self._magnitudes, pos)._value;
    }

    /**
     * @dev Returns the value in the last (most recent) checkpoint with timestamp lower or equal than the search timestamp, or zero if there is none.
     */
    function upperLookup(History storage self, uint32 timestamp) internal view returns (uint224) {
        uint256 len = self._magnitudes.length;
        uint256 pos = _upperBinaryLookup(self._magnitudes, timestamp, 0, len);
        return pos == 0 ? 0 : _unsafeAccess(self._magnitudes, pos - 1)._value;
    }

    /**
     * @dev Returns the value in the last (most recent) checkpoint with timestamp lower or equal than the search timestamp, or zero if there is none.
     *
     * NOTE: This is a variant of {upperLookup} that is optimised to find "recent" checkpoint (checkpoints with high timestamps).
     */
    function upperLookupRecent(History storage self, uint32 timestamp) internal view returns (uint224) {
        uint256 len = self._magnitudes.length;

        uint256 low = 0;
        uint256 high = len;

        if (len > 5) {
            uint256 mid = len - MathUpgradeable.sqrt(len);
            if (timestamp < _unsafeAccess(self._magnitudes, mid)._timestamp) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }

        uint256 pos = _upperBinaryLookup(self._magnitudes, timestamp, low, high);

        return pos == 0 ? 0 : _unsafeAccess(self._magnitudes, pos - 1)._value;
    }

    /**
     * @dev Returns the value in the most recent checkpoint, or zero if there are no checkpoints.
     */
    function latest(History storage self) internal view returns (uint224) {
        uint256 pos = self._magnitudes.length;
        return pos == 0 ? 0 : _unsafeAccess(self._magnitudes, pos - 1)._value;
    }

    /**
     * @dev Returns whether there is a checkpoint in the structure (i.e. it is not empty), and if so the timestamp and value
     * in the most recent checkpoint.
     */
    function latestCheckpoint(History storage self) internal view returns (bool exists, uint32 _timestamp, uint224 _value) {
        uint256 pos = self._magnitudes.length;
        if (pos == 0) {
            return (false, 0, 0);
        } else {
            MagnitudeCheckpoint memory ckpt = _unsafeAccess(self._magnitudes, pos - 1);
            return (true, ckpt._timestamp, ckpt._value);
        }
    }

    /**
     * @dev Returns the number of checkpoint.
     */
    function length(History storage self) internal view returns (uint256) {
        return self._magnitudes.length;
    }

    /**
     * @dev Pushes a (`timestamp`, `value`) pair into an ordered list of checkpoints, either by inserting a new checkpoint,
     * or by updating the last one.
     */
    function _insert(MagnitudeCheckpoint[] storage self, uint32 timestamp, uint224 value) private returns (uint224, uint224) {
        uint256 pos = self.length;

        if (pos > 0) {
            // Copying to memory is important here.
            MagnitudeCheckpoint memory last = _unsafeAccess(self, pos - 1);

            // Checkpoint timestamps must be non-decreasing.
            require(last._timestamp <= timestamp, "Checkpoint: decreasing timestamps");

            // Update or push new checkpoint
            if (last._timestamp == timestamp) {
                _unsafeAccess(self, pos - 1)._value = value;
            } else {
                self.push(MagnitudeCheckpoint({_timestamp: timestamp, _value: value}));
            }
            return (last._value, value);
        } else {
            self.push(MagnitudeCheckpoint({_timestamp: timestamp, _value: value}));
            return (0, value);
        }
    }

    /**
     * @dev Return the index of the last (most recent) checkpoint with timestamp lower or equal than the search timestamp, or `high` if there is none.
     * `low` and `high` define a section where to do the search, with inclusive `low` and exclusive `high`.
     *
     * WARNING: `high` should not be greater than the array's length.
     */
    function _upperBinaryLookup(
        MagnitudeCheckpoint[] storage self,
        uint32 timestamp,
        uint256 low,
        uint256 high
    ) private view returns (uint256) {
        while (low < high) {
            uint256 mid = MathUpgradeable.average(low, high);
            if (_unsafeAccess(self, mid)._timestamp > timestamp) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }
        return high;
    }

    /**
     * @dev Return the index of the first (oldest) checkpoint with timestamp is greater or equal than the search timestamp, or `high` if there is none.
     * `low` and `high` define a section where to do the search, with inclusive `low` and exclusive `high`.
     *
     * WARNING: `high` should not be greater than the array's length.
     */
    function _lowerBinaryLookup(
        MagnitudeCheckpoint[] storage self,
        uint32 timestamp,
        uint256 low,
        uint256 high
    ) private view returns (uint256) {
        while (low < high) {
            uint256 mid = MathUpgradeable.average(low, high);
            if (_unsafeAccess(self, mid)._timestamp < timestamp) {
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
    function _unsafeAccess(
        MagnitudeCheckpoint[] storage self,
        uint256 pos
    ) private pure returns (MagnitudeCheckpoint storage result) {
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
     * @dev Returns the value in the last (most recent) checkpoint with timestamp lower or equal than the search timestamp, or zero if there is none.
     * In addition, returns the position of the checkpoint in the array.
     *
     * NOTE: That if value != 0 && pos == 0, then that means the value is the first checkpoint and actually exists
     * a checkpoint DNE iff value == 0 && pos == 0
     */
    function upperLookupWithPos(History storage self, uint32 timestamp) internal view returns (uint224, uint256) {
        uint256 len = self._magnitudes.length;
        uint256 pos = _upperBinaryLookup(self._magnitudes, timestamp, 0, len);
        return pos == 0 ? (0, 0) : (_unsafeAccess(self._magnitudes, pos - 1)._value, pos - 1);
    }

    /**
     * @dev Returns the value in the last (most recent) checkpoint with timestamp lower or equal than the search timestamp, or zero if there is none.
     * In addition, returns the position of the checkpoint in the array.
     *
     * NOTE: This is a variant of {upperLookup} that is optimised to find "recent" checkpoint (checkpoints with high timestamps).
     * NOTE: That if value != 0 && pos == 0, then that means the value is the first checkpoint and actually exists
     * a checkpoint DNE iff value == 0 && pos == 0 => value == 0
     */
    function upperLookupRecentWithPos(History storage self, uint32 timestamp) internal view returns (uint224, uint256) {
        uint256 len = self._magnitudes.length;

        uint256 low = 0;
        uint256 high = len;

        if (len > 5) {
            uint256 mid = len - MathUpgradeable.sqrt(len);
            if (timestamp < _unsafeAccess(self._magnitudes, mid)._timestamp) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }

        uint256 pos = _upperBinaryLookup(self._magnitudes, timestamp, low, high);

        return pos == 0 ? (0, 0) : (_unsafeAccess(self._magnitudes, pos - 1)._value, pos - 1);
    }

    /// @notice WARNING: this function is only used because of the invariant property
    /// that from the current timestamp, all future checkpointed magnitude values are strictly > current value.
    /// Use function with extreme care for other situations.
    function decrementAtAndFutureCheckpoints(
        History storage self,
        uint32 timestamp,
        uint224 decrementValue
    ) internal {
        (uint224 value, uint256 pos) = upperLookupRecentWithPos(self, timestamp);

        // if there is no checkpoint, return
        if (value == 0 && pos == 0) {
            pos = type(uint256).max;
        }

        uint256 len = self._magnitudes.length;
        while (pos < len) {
            MagnitudeCheckpoint storage current = _unsafeAccess(self._magnitudes, pos);

            // reverts from underflow. Expected to never happen in our usage
            current._value -= decrementValue;

            unchecked {
                ++pos;
            }
        }
    }
}
