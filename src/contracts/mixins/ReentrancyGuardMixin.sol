// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/**
 * @dev Variant of {ReentrancyGuard} that uses transient storage.
 *
 * NOTE: This variant only works on networks where EIP-1153 is available.
 */
abstract contract ReentrancyGuardMixin {
    // keccak256(abi.encode(uint256(keccak256("eigenlayer.storage.ReentrancyGuardMixin")) - 1)) & ~bytes32(uint256(0xff))
    uint256 private constant _REENTRANCY_GUARD_SLOT =
        0x61bb794ad7a504b3613420bc192fca11ecb0ea36bf99527d17aa6bd66a5db500;

    /// @dev Unauthorized reentrant call.
    error Reentrancy();

    /// @dev Guards a function from reentrancy.
    modifier nonReentrant() virtual {
        uint256 s = _REENTRANCY_GUARD_SLOT;
        /// @solidity memory-safe-assembly
        assembly {
            if tload(s) {
                mstore(0x00, 0xab143c06) // `Reentrancy()`.
                revert(0x1c, 0x04)
            }
            tstore(s, address())
        }
        _;
        /// @solidity memory-safe-assembly
        assembly {
            tstore(s, 0)
        }
    }
}