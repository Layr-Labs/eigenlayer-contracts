# BLSPubkeyRegistry

This contract is a registry that keeps track of aggregate public key hashes of the quorums of an AVS over time. AVSs that want access to the aggregate public keys of their operator set over time should integrate this registry with their RegistryCoordinator.

## Flows

### registerOperator

The RegistryCoordinator for the AVS makes a call to the BLSPubkeyRegistry to register an operator with a certain public key for a certain set of quorums. The BLSPubkeyRegistry verifies that the operator in fact owns the public key by [making a call to the BLSPublicKeyCompendium](./BLSPublicKeyCompendium.md#integrations). It then, for each quorum the operator is registering for, adds the operator's public key to the aggregate quorum public key, ends the active block range for the previous aggregate public key, and begins the active block range for the new aggregate quorum public key. Updates are stored using the following struct:
```solidity
/// @notice Data structure used to track the history of the Aggregate Public Key of all operators
struct ApkUpdate {
    // first 24 bytes of keccak256(apk_x, apk_y)
    bytes24 apkHash;
    // block number at which the update occurred
    uint32 updateBlockNumber;
    // block number at which the next update occurred
    uint32 nextUpdateBlockNumber;
}
```

The aggregate quorum public key stored in the contract is also overwritten with the new aggregate quorum public key.

The function also returns the hash of the operator's public key as it may be used as is for the operator in the AVS. The [BLSRegistryCoordinator](./BLSRegistryCoordinatorWithIndices.md) uses the hash of the operator's public key as the operator's identifier (operator id) since it lowers gas costs in the [BLSSignatureChecker](./BLSSignatureChecker.md).

### deregisterOperator

The RegistryCoordinator for the AVS makes a call to the BLSPubkeyRegistry to deregister an operator with a certain public key for a certain set of quorums. The BLSPubkeyRegistry verifies that the operator in fact owns the public key by [making a call to the BLSPublicKeyCompendium](./BLSPublicKeyCompendium.md#integrations). It then, for each quorum the operator is registering for, subtracts the operator's public key from the aggregate quorum public key, ends the active block range for the previous aggregate public key, and begins the active block range for the new aggregate quorum public key. 

The aggregate quorum public key stored in the contract is also overwritten with the new aggregate quorum public key.

Note that the contract does not check that the quorums that the operator's public key is being subtracted from are a subset of the quorums the operator is registered for, that logic is expected to be done in the RegistryCoordinator.

## Upstream Dependencies

The main integration with the BLSPublicKeyRegistry is used by the AVSs [BLSSignatureChecker](./BLSSignatureChecker.md). An offchain actor provides a public key, a quorum id, and an index in the array of aggregate quorum public key hashes, and the AVS's signature checker verifies that a certain quorum's aggregate public key hash at a certain block number was in fact the hash of the provided public key. Look at `getApkHashForQuorumAtBlockNumberFromIndex`.
