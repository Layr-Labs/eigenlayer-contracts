# IndexRegistry

This contract assigns each operator an index (0 indexed) within each of its quorums. If a quorum has $n$ operators, each operator will be assigned an index $0$ through $n-1$. This contract is used for AVSs that need a common ordering among all operators in a quorum that is accessible onchain. For example, this will be used in proofs of custody by EigenDA. This contract also keeps a list of all operators that have ever joined the AVS for convenience purposes in offchain software that are out of scope for this document.

## Flows

### registerOperator

The RegistryCoordinator for the AVS makes call to the IndexRegistry to register an operator for a certain set of quorums. The IndexRegistry will assign the next index in each of the quorums the operator is registering for to the operator storing the following struct:
```solidity
// struct used to give definitive ordering to operators at each blockNumber. 
// NOTE: this struct is slightly abused for also storing the total number of operators for each quorum over time
struct OperatorIndexUpdate {
    // blockNumber number from which `index` was the operators index
    // the operator's index or the total number of operators at a `blockNumber` is the first entry such that `blockNumber >= entry.fromBlockNumber`
    uint32 fromBlockNumber;
    // index of the operator in array of operators, or the total number of operators if in the 'totalOperatorsHistory'
    // index = type(uint32).max = OPERATOR_DEREGISTERED_INDEX implies the operator was deregistered
    uint32 index;
}
```

The IndexRegistry also adds the operator's id to the append only list of operators that have registered for the middleware and it stores the total number of operators after the registering operator has registered for each of the quorums the operator is registering for by pushing the above struct to a growing array.

### deregisterOperator

The RegistryCoordinator for the AVS makes call to the IndexRegistry to deregister an operator for a certain set of quorums. The RegistryCoordinator provides a witness of the ids of the operators that have the greatest index in each of the quorums that the operator is deregistering from. The IndexRegistry then, for each quorum the operator is deregistering from, 
1. Decrements the total number of operators in the quorum
2. Makes sure the provided "greatest index operator id" in fact has the greatest index in the quorum by checking it against the total number of operators in the quorum
3. Sets the index of the "greatest index operator" to the index of the deregistering operator
4. Sets the index of the deregistering operator to `OPERATOR_DEREGISTERED_INDEX = type(uint32).max`

Steps 3 and 4 are done via pushing the above struct to a growing array that is kept track of for each operator.

Note that the contract does not check that the quorums that the operator is being deregistered from are a subset of the quorums the operator is registered for, that logic is expected to be done in the RegistryCoordinator.

## Upstream Dependencies

The [BLSOperatorStateRetriever](./BLSOperatorStateRetriever.md) uses the globally ordered list of all operators every registered for the AVS to serve information about the active operator set for the AVS to offchain nodes.