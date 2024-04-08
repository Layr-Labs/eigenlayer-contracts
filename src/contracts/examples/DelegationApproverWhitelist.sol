// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../libraries/Merkle.sol";

/**
 * @title DelegationApproverWhitelist
 * @author Layr Labs, Inc.
 * @notice Allows an operator to efficiently maintain a "whitelist" of stakers who can are approved to delegate to them.
 * @dev Usage:
 * - An EigenLayer operator sets this contract as their `delegationApprover` in EigenLayer's DelegationManager
 * - This contract's owner (who could be the operator or could be another party) sets a merkle root of approved hashes
 * - Any time a callback is made to this contract from EigenLayer's DelegationManager contract,
 * (either by calling `DelegationManager.delegateTo` or `DelegationManager.delegateToBySignature`),
 * the caller must provide a valid `merkleProofAndIndex` input to this contract's `isValidSignature` function, proving inclusion
 * of the relevant `delegationHash` in the Merkle tree corresponding to the present `merkleRoot`.
 * - In practice, the contract owner will need to construct a Merkle tree of approved hashes -- which could be derived from
 * a "whitelist" of approved stakers and referencing DelegationManager.calculateDelegationApprovalDigestHash -- and then update
 * the `merkleRoot` stored in this contract, and either
 * (a) publish the Merkle tree for these stakers to pull proofs from, or
 * (b) provide the proofs to stakers upon request.
 */
contract DelegationApproverWhitelist is Ownable {
    // @notice EigenLayer's DelegationManager contract
    address immutable public eigenLayerDelegationManager;

    // @notice Merkle root for the tree of approved delegation hashes
    bytes32 public merkleRoot;

    event MerkleRootUpdated(bytes32 previousRoot, bytes32 newRoot);

    constructor(address _eigenLayerDelegationManager, address initialOwner) {
        eigenLayerDelegationManager = _eigenLayerDelegationManager;
        _transferOwnership(initialOwner);
    }

    /**
     * @notice Allows the contract owner to update the stored Merkle root
     * @param _merkleRoot New Merkle root to represent the updated whitelist.
     */
    function updateMerkleRoot(bytes32 _merkleRoot) external onlyOwner {
        emit MerkleRootUpdated(merkleRoot, _merkleRoot);
        merkleRoot = _merkleRoot;
    }

    /**
     * @dev Verifies whether or not a staker is included in the operator's whitelist using a Merkle proof.
     * @param delegationHash Hash of the data associated with the staker. See `DelegationManager.calculateDelegationApprovalDigestHash`
     * @param merkleProofAndIndex ABI-encoded Merkle proof and index to verify if the staker is in the whitelist.
     * @return Returns the EIP-1271 magic value indicating if the signature (Merkle proof) is valid or not.
     * @dev merkleProofAndIndex is simply `abi.encode(merkleProof, index)`
     */
    function isValidSignature(bytes32 delegationHash, bytes calldata merkleProofAndIndex) public view returns (bytes4) {
        if (msg.sender != eigenLayerDelegationManager) {
            // "magic bytes" for invalid signature
            return 0xffffffff;
        } else {
            (bytes memory proof, uint256 index) = abi.decode(merkleProofAndIndex, (bytes, uint256));
            if (
                Merkle.verifyInclusionKeccak({
                    proof: proof,
                    root: merkleRoot,
                    leaf: delegationHash,
                    index: index
                })
            ) {
                // "magic bytes" for valid signature
                return 0x1626ba7e;
            } else {
                return 0xffffffff;
            }
        }
    }
}