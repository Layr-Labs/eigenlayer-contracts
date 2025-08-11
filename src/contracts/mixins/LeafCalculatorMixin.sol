// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IOperatorTableCalculatorTypes} from "../interfaces/IOperatorTableCalculator.sol";

/**
 * @title LeafCalculatorMixin
 * @notice Reusable mixin for calculating operator info and operator table leaf hashes
 * @dev Provides standardized leaf calculation functions for use across multiple contracts and repositories.
 *      This mixin centralizes the leaf hashing logic to ensure consistency across the EigenLayer ecosystem
 *      and maintains proper cryptographic security through salt-based domain separation.
 */
abstract contract LeafCalculatorMixin {
    /// @dev Salt for operator info leaf hash calculation
    /// @dev The salt is used to prevent against second preimage attacks: attacks where an
    /// attacker can create a partial proof using an internal node rather than a leaf to
    /// validate a proof. The salt ensures that leaves cannot be concatenated together to
    /// form a valid proof, as well as reducing the likelihood of an internal node matching
    /// the salt prefix.
    /// @dev Value derived from keccak256("OPERATOR_INFO_LEAF_SALT") = 0x75...
    /// This ensures collision resistance and semantic meaning.
    uint8 public constant OPERATOR_INFO_LEAF_SALT = 0x75;

    /// @dev Salt for operator table leaf hash calculation
    /// @dev The salt is used to prevent against second preimage attacks: attacks where an
    /// attacker can create a partial proof using an internal node rather than a leaf to
    /// validate a proof. The salt ensures that leaves cannot be concatenated together to
    /// form a valid proof, as well as reducing the likelihood of an internal node matching
    /// the salt prefix.
    /// @dev Value derived from keccak256("OPERATOR_TABLE_LEAF_SALT") = 0x8e...
    /// This ensures collision resistance and semantic meaning.
    uint8 public constant OPERATOR_TABLE_LEAF_SALT = 0x8e;

    /**
     * @notice Calculate the leaf hash for an operator info
     * @param operatorInfo The BN254 operator info struct containing the operator's public key and stake weights
     * @return The leaf hash (keccak256 of salt and encoded operator info)
     * @dev The salt is used to prevent against second preimage attacks: attacks where an
     * attacker can create a partial proof using an internal node rather than a leaf to
     * validate a proof. The salt ensures that leaves cannot be concatenated together to
     * form a valid proof, as well as reducing the likelihood of an internal node matching
     * the salt prefix.
     *
     * This is a standard "domain separation" technique in Merkle tree implementations
     * to ensure leaf nodes and internal nodes can never be confused with each other.
     * See Section 2.1 of <https://www.rfc-editor.org/rfc/rfc9162#name-merkle-trees> for more.
     *
     * Uses abi.encodePacked for the salt and abi.encode for the struct to handle complex types
     * (structs with dynamic arrays) while maintaining gas efficiency where possible.
     */
    function calculateOperatorInfoLeaf(
        IOperatorTableCalculatorTypes.BN254OperatorInfo memory operatorInfo
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(OPERATOR_INFO_LEAF_SALT, abi.encode(operatorInfo)));
    }

    /**
     * @notice Calculate the leaf hash for an operator table
     * @param operatorTableBytes The encoded operator table as bytes containing operator set data
     * @return The leaf hash (keccak256 of salt and operator table bytes)
     * @dev The salt is used to prevent against second preimage attacks: attacks where an
     * attacker can create a partial proof using an internal node rather than a leaf to
     * validate a proof. The salt ensures that leaves cannot be concatenated together to
     * form a valid proof, as well as reducing the likelihood of an internal node matching
     * the salt prefix.
     *
     * This is a standard "domain separation" technique in Merkle tree implementations
     * to ensure leaf nodes and internal nodes can never be confused with each other.
     * See Section 2.1 of <https://www.rfc-editor.org/rfc/rfc9162#name-merkle-trees> for more.
     *
     * Uses abi.encodePacked for both salt and bytes for optimal gas efficiency since both
     * are simple byte arrays without complex nested structures.
     */
    function calculateOperatorTableLeaf(
        bytes calldata operatorTableBytes
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(OPERATOR_TABLE_LEAF_SALT, operatorTableBytes));
    }
}
