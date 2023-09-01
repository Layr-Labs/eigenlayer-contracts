# BLSRegistryCoordinatorWithIndices

This contract is deployed for every AVS and serves as the main entrypoint for operators to register and deregister from the AVS. In addition, it is where the AVS defines its operator churn parameters.

## Flows

### registerOperator

When registering the operator must provide 
1. The quorums they are registering for
2. Their BLS public key that they registered with the [BLSPubkeyCompendium](./BLSPublicKeyCompendium.md)
3. The socket (ip:port) at which AVS offchain actors should make requests

The RegistryCoordinator then
1. Registers their BLS public key with the [BLSPubkeyRegistry](BLSPubkeyRegistry.md) and notes the hash of their public key as their operator id
2. Registers them with the StakeRegistry 
3. Registers them with the IndexRegistry
4. Stores the quorum bitmap of the operator using the following struct:
```
/**
 * @notice Data structure for storing info on quorum bitmap updates where the `quorumBitmap` is the bitmap of the 
 * quorums the operator is registered for starting at (inclusive)`updateBlockNumber` and ending at (exclusive) `nextUpdateBlockNumber`
 * @dev nextUpdateBlockNumber is initialized to 0 for the latest update
 */
struct QuorumBitmapUpdate {
    uint32 updateBlockNumber;
    uint32 nextUpdateBlockNumber;
    uint192 quorumBitmap;
}
```

### If quorum full

The following struct is defined for each quorum
```
/**
 * @notice Data structure for storing operator set params for a given quorum. Specifically the 
 * `maxOperatorCount` is the maximum number of operators that can be registered for the quorum,
 * `kickBIPsOfOperatorStake` is the basis points of a new operator needs to have of an operator they are trying to kick from the quorum,
 * and `kickBIPsOfTotalStake` is the basis points of the total stake of the quorum that an operator needs to be below to be kicked.
 */ 
struct OperatorSetParam {
    uint32 maxOperatorCount;
    uint16 kickBIPsOfOperatorStake;
    uint16 kickBIPsOfTotalStake;
}
```

If any of the quorums is full (number of operators in it is `maxOperatorCount`), the operator must provide the public key of an operator which it has more than `kickBIPsOfOperatorStake` than (measured in basis points) and that out has less than `kickBIPsOfTotalStake` than the entire quorum's stake. The provided operators are deregistered from the respective quorums that are full which the registering operator is registering for. Since the operators being kciked may not be the operators with the least stake, the RegistryCoordinator requires that the provided operators are signed off by a permissioned address called a `churnApprover`. Note that the quorum operator caps are due to the cost of BLS signature (dis)aggregation onchain.

Note that the registering operator must also provide the ids of the operators that must swap indexes with the operators being kicked from the quorums it is registering for.

### deregisterOperator

When deregistering, an operator provides
1. The quorums they registered for
2. Their BLS public key
3. The ids of the operators that must swap indices with the deregistering operator in the IndexRegistry

The RegistryCoordinator then deregisters the operator with the BLSPubkeyRegistry, StakeRegistry, and IndexRegistry. It then ends the block range for its stored quorum bitmap for the operator.