// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IAllocationManager.sol";
import "../interfaces/IKeyRegistrar.sol";

abstract contract KeyRegistrarStorage is IKeyRegistrar {
    // Constants

    /// @dev Gas limit for pairing operations to prevent DoS
    uint256 internal constant PAIRING_EQUALITY_CHECK_GAS = 400_000;

    /// @dev Hash of zero ECDSA public key (0x04 + 64 zero bytes)
    bytes32 internal constant ZERO_ECDSA_PUBKEY_HASH = keccak256(abi.encodePacked(bytes1(0x04), new bytes(64)));

    // Immutables

    /// @dev Reference to the AllocationManager contract
    IAllocationManager public immutable allocationManager;

    // Mutatables

    /// @dev Maps (operatorSetKey, operator) to their key info
    mapping(bytes32 => mapping(address => KeyInfo)) internal operatorKeyInfo;

    /// @dev Maps operatorSetKey to the key type
    mapping(bytes32 => CurveType) internal operatorSetCurveTypes;

    /// @dev Global mapping of key hash to registration status - enforces global uniqueness
    mapping(bytes32 => bool) internal globalKeyRegistry;

    // Construction

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[47] private __gap;
}
