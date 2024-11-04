// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

type Randomness is uint256;

using Random for Randomness global;

library Random {
    /// -----------------------------------------------------------------------
    /// Constants
    /// -----------------------------------------------------------------------

    // Equivalent to: `uint256(keccak256("RANDOMNESS.SEED"))`.
    uint256 constant SEED = 0x93bfe7cafd9427243dc4fe8c6e706851eb6696ba8e48960dd74ecc96544938ce;
    
    /// Equivalent to: `uint256(keccak256("RANDOMNESS.SEED"))`.
    uint256 constant SLOT = 0xd0660badbab446a974e6a19901c78a2ad88d7e4f1710b85e1cfc0878477344fd;

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    function set(Randomness r) internal returns (Randomness) {
        /// @solidity memory-safe-assembly
        assembly {
            sstore(SLOT, r)
        }
        return r;
    }

    function shuffle(Randomness r) internal returns (Randomness) {
        /// @solidity memory-safe-assembly
        assembly {
            mstore(0x00, sload(SLOT))
            mstore(0x20, r)
            r := keccak256(0x00, 0x20)
        }
        return r.set();
    }

    function Uint256(Randomness r, uint256 min, uint256 max) internal returns (uint256) {
        return max <= min ? min : r.Uint256() % (max - min) + min;
    }

    function Uint256(Randomness r) internal returns (uint256) {
        return r.shuffle().unwrap();
    }

    function Uint64(Randomness r, uint64 min, uint64 max) internal returns (uint64) {
        return uint64(Uint256(r, min, max));
    }

    function Uint64(Randomness r) internal returns (uint64) {
        return uint64(Uint256(r));
    }

    function Uint32(Randomness r, uint32 min, uint32 max) internal returns (uint32) {
        return uint32(Uint256(r, min, max));
    }

    function Uint32(Randomness r) internal returns (uint32) {
        return uint32(Uint256(r));
    }

    function Bytes32(Randomness r) internal returns (bytes32) {
        return bytes32(r.Uint256());
    }

    function Address(Randomness r) internal returns (address) {
        return address(uint160(r.Uint256(1, type(uint160).max)));
    }

    function wrap(uint256 r) internal pure returns (Randomness) {
        return Randomness.wrap(r);
    }

    function unwrap(Randomness r) internal pure returns (uint256) {
        return Randomness.unwrap(r);
    }
}
