// SPDX-License-Identifier: MIT
// forgefmt: disable-start
// Adapted from OpenZeppelin Contracts (last updated v4.9.0) (utils/Checkpoints.sol)

pragma solidity ^0.8.0;

import "@openzeppelin-v4.9.0/contracts/utils/math/Math.sol";
import "@openzeppelin-v4.9.0/contracts/utils/math/SafeCast.sol";

/**
 * @dev This library defines the `History` struct, for checkpointing values as they change at different points in
 * time, and later looking up past values by block timestamp.
 *
 * To create a history of checkpoints define a variable type `TimestampCheckpoints.History` in your contract, and store a new
 * checkpoint for the current transaction timestamp using the {push} function.
 *
 * _Available since v4.5._
 *
 * We disable forge-fmt for this file to simplify diffing against the original OpenZeppelin version:
 * https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.9/contracts/utils/Checkpoints.sol
 */
library TimestampCheckpoints {
    struct History {
        Checkpoint[] _checkpoints;
    }

    struct Checkpoint {
        uint32 _blockTimestamp;
        uint224 _value;
    }

    /**
     * @dev Returns the value at a given block timestamp. If a checkpoint is not available at that timestamp, the closest one
     * before it is returned, or zero otherwise. Because the number returned corresponds to that at the end of the
     * block timestamp, the requested block timestamp must be in the past, excluding the current block timestamp.
     */
    function getAtTimestamp(History storage self, uint256 blockTimestamp) internal view returns (uint256) {
        require(blockTimestamp < block.timestamp, "Checkpoints: timestamp is not in the past");
        uint32 key = SafeCast.toUint32(blockTimestamp);

        uint256 len = self._checkpoints.length;
        uint256 pos = _upperBinaryLookup(self._checkpoints, key, 0, len);
        return pos == 0 ? 0 : _unsafeAccess(self._checkpoints, pos - 1)._value;
    }

    /**
     * @dev Returns the value at a given block timestamp. If a checkpoint is not available at that timestamp, the closest one
     * before it is returned, or zero otherwise. Similar to {upperLookup} but optimized for the case when the searched
     * checkpoint is probably "recent", defined as being among the last sqrt(N) checkpoints where N is the number of
     * checkpoints.
     */
    function getAtProbablyRecentTimestamp(
        History storage self,
        uint256 blockTimestamp
    ) internal view returns (uint256) {
        require(blockTimestamp < block.timestamp, "Checkpoints: timestamp is not in the past");
        uint32 key = SafeCast.toUint32(blockTimestamp);

        uint256 len = self._checkpoints.length;

        uint256 low = 0;
        uint256 high = len;

        if (len > 5) {
            uint256 mid = len - Math.sqrt(len);
            if (key < _unsafeAccess(self._checkpoints, mid)._blockTimestamp) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }

        uint256 pos = _upperBinaryLookup(self._checkpoints, key, low, high);

        return pos == 0 ? 0 : _unsafeAccess(self._checkpoints, pos - 1)._value;
    }

    /**
     * @dev Pushes a value onto a History so that it is stored as the checkpoint for the current block.
     *
     * Returns previous value and new value.
     */
    function push(History storage self, uint256 value) internal returns (uint256, uint256) {
        return _insert(self._checkpoints, SafeCast.toUint32(block.timestamp), SafeCast.toUint224(value));
    }

    /**
     * @dev Pushes a value onto a History, by updating the latest value using binary operation `op`. The new value will
     * be set to `op(latest, delta)`.
     *
     * Returns previous value and new value.
     */
    function push(
        History storage self,
        function(uint256, uint256) view returns (uint256) op,
        uint256 delta
    ) internal returns (uint256, uint256) {
        return push(self, op(latest(self), delta));
    }

    /**
     * @dev Returns the value in the most recent checkpoint, or zero if there are no checkpoints.
     */
    function latest(History storage self) internal view returns (uint224) {
        uint256 pos = self._checkpoints.length;
        return pos == 0 ? 0 : _unsafeAccess(self._checkpoints, pos - 1)._value;
    }

    /**
     * @dev Returns whether there is a checkpoint in the structure (i.e. it is not empty), and if so the key and value
     * in the most recent checkpoint.
     */
    function latestCheckpoint(
        History storage self
    ) internal view returns (bool exists, uint32 _blockTimestamp, uint224 _value) {
        uint256 pos = self._checkpoints.length;
        if (pos == 0) {
            return (false, 0, 0);
        } else {
            Checkpoint memory ckpt = _unsafeAccess(self._checkpoints, pos - 1);
            return (true, ckpt._blockTimestamp, ckpt._value);
        }
    }

    /**
     * @dev Returns the number of checkpoint.
     */
    function length(History storage self) internal view returns (uint256) {
        return self._checkpoints.length;
    }

    /**
     * @dev Pushes a (`key`, `value`) pair into an ordered list of checkpoints, either by inserting a new checkpoint,
     * or by updating the last one.
     */
    function _insert(Checkpoint[] storage self, uint32 key, uint224 value) private returns (uint224, uint224) {
        uint256 pos = self.length;

        if (pos > 0) {
            // Copying to memory is important here.
            Checkpoint memory last = _unsafeAccess(self, pos - 1);

            // Checkpoint keys must be non-decreasing.
            require(last._blockTimestamp <= key, "Checkpoint: decreasing keys");

            // Update or push new checkpoint
            if (last._blockTimestamp == key) {
                _unsafeAccess(self, pos - 1)._value = value;
            } else {
                self.push(Checkpoint({_blockTimestamp: key, _value: value}));
            }
            return (last._value, value);
        } else {
            self.push(Checkpoint({_blockTimestamp: key, _value: value}));
            return (0, value);
        }
    }

    /**
     * @dev Return the index of the last (most recent) checkpoint with key lower or equal than the search key, or `high` if there is none.
     * `low` and `high` define a section where to do the search, with inclusive `low` and exclusive `high`.
     *
     * WARNING: `high` should not be greater than the array's length.
     */
    function _upperBinaryLookup(
        Checkpoint[] storage self,
        uint32 key,
        uint256 low,
        uint256 high
    ) private view returns (uint256) {
        while (low < high) {
            uint256 mid = Math.average(low, high);
            if (_unsafeAccess(self, mid)._blockTimestamp > key) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }
        return high;
    }

    /**
     * @dev Return the index of the first (oldest) checkpoint with key is greater or equal than the search key, or `high` if there is none.
     * `low` and `high` define a section where to do the search, with inclusive `low` and exclusive `high`.
     *
     * WARNING: `high` should not be greater than the array's length.
     */
    function _lowerBinaryLookup(
        Checkpoint[] storage self,
        uint32 key,
        uint256 low,
        uint256 high
    ) private view returns (uint256) {
        while (low < high) {
            uint256 mid = Math.average(low, high);
            if (_unsafeAccess(self, mid)._blockTimestamp < key) {
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
    function _unsafeAccess(Checkpoint[] storage self, uint256 pos) private pure returns (Checkpoint storage result) {
        assembly {
            mstore(0, self.slot)
            result.slot := add(keccak256(0, 0x20), pos)
        }
    }
}
