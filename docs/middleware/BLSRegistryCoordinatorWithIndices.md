# BLSRegistryCoordinatorWithIndices

This contract is deployed for every AVS and serves as the main entrypoint for operators to register and deregister from the AVS. In addition, it is where the AVS defines its operator churn parameters.

## Flows

### registerOperator

When registering the operator must provide 
1. The quorums they are registering for
2. Their BLS public key that they registered with the [BLSPubkeyCompendium](./BLSPublicKeyCompendium.md)
3. The socket (ip:port) at which AVS offchain actors should make requests

The RegistryCoordinator then
1. Registers the operator's BLS public key with the [BLSPubkeyRegistry](BLSPubkeyRegistry.md) and notes the hash of their public key as their operator id
2. Registers the operator with the [StakeRegistry](./StakeRegistry.md) 
3. Registers the operator with the [IndexRegistry](./IndexRegistry.md)
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

Operators can be registered for certain quorums and later register for other (non-overlapping) quorums.

### If quorum full

The following struct is defined for each quorum
```
/**
 * @notice Data structure for storing operator set params for a given quorum. Specifically the 
 * `maxOperatorCount` is the maximum number of operators that can be registered for the quorum
 * `kickBIPsOfOperatorStake` is the multiple (in basis points) of stake that a new operator must have, as compared the operator that they are kicking out of the quorum 
 * `kickBIPsOfTotalStake` is the fraction (in basis points) of the total stake of the quorum that an operator needs to be below to be kicked.
 */ 
struct OperatorSetParam {
    uint32 maxOperatorCount;
    uint16 kickBIPsOfOperatorStake;
    uint16 kickBIPsOfTotalStake;
}
```

If any of the quorums is full (number of operators in it is `maxOperatorCount`), the newly registering operator must provide the public key of another operator to be kicked. The new and kicked operator must satisfy the two conditions:
1. the new operator has an amount of stake that is at least `kickBIPsOfOperatorStake` multiple of the kicked operator's stake
2. the kicked operator has less than `kickBIPsOfTotalStake` fraction of the quorum's total stake

The provided operators are deregistered from the respective quorums that are full which the registering operator is registering for. Since the operators being kicked may not be the operators with the least stake, the RegistryCoordinator requires that the provided operators are signed off by a permissioned address called a `churnApprover`. Note that the quorum operator caps are due to the cost of BLS signature (dis)aggregation onchain.

Operators register with a list of 
```
/**
* @notice Data structure for the parameters needed to kick an operator from a quorum with number `quorumNumber`, used during registration churn.
* Specifically the `operator` is the address of the operator to kick, `pubkey` is the BLS public key of the operator,
*/
struct OperatorKickParam {
    uint8 quorumNumber;
    address operator;
    BN254.G1Point pubkey; 
}
```
For each quorum they need to kick operators from. This list, along with the id of the registering operator needs to be signed (along with a salt and expiry) by an actor known as the *churnApprover*. Operators will make a request to the churnApprover offchain before registering for their signature, if needed.

### deregisterOperator

When deregistering, an operator provides
1. The quorums they registered for
2. Their BLS public key
3. The ids of the operators that must swap indices with the [deregistering operator in the IndexRegistry](./IndexRegistry.md#deregisteroperator).

The RegistryCoordinator then deregisters the operator with the BLSPubkeyRegistry, StakeRegistry, and IndexRegistry. It then ends the block range for its stored quorum bitmap for the operator.

Operators can deregister from a subset of quorums that they are registered for.

## Upstream Dependencies

Operators register and deregister with the AVS for certain quorums through this contract.

EigenLabs intends to run the EigenDA churnApprover.