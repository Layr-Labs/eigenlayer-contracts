# StakeRegistry

This contract is deployed for every AVS and keeps track of the AVS's operators' stakes over time and the total stakes for each quorum. In addition, this contract also handles the adding and modification of quorum.

# Definitions

A **quorum** is defined by list of the following structs:
```
struct StrategyAndWeightingMultiplier {
    IStrategy strategy;
    uint96 multiplier;
}
```

## Flows

### createQuorum

The owner of the StakeRegistry can create a quorum by providing the list of `StrategyAndWeightingMultiplier`s. Quorums cannot be removed.

### modifyQuorum

The owner of the StakeRegistry can modify the set of strategies and they multipliers for a certain quorum.

### registerOperator

The RegistryCoordinator for the AVS makes a call to the StakeRegistry to register an operator for a certain set of quorums. For each of the quorums being registered for, the StakeRegistry calculates a linear combination of the operator's delegated shares of each `strategy` in the quorum and their corresponding `multiplier` to get a `stake`. The contract then stores the stake in the following struct:
```
/// @notice struct used to store the stakes of an individual operator or the sum of all operators' stakes, for storage
struct OperatorStakeUpdate {
    // the block number at which the stake amounts were updated and stored
    uint32 updateBlockNumber;
    // the block number at which the *next update* occurred.
    /// @notice This entry has the value **0** until another update takes place.
    uint32 nextUpdateBlockNumber;
    // stake weight for the quorum
    uint96 stake;
}
```
For each quorum the operator is a part of.

### deregisterOperator

The RegistryCoordinator for the AVS calls the StakeRegistry to deregister an operator for a certain set of quorums. For each of the quorums being registered for, the StakeRegistry ends the block range of the current `OperatorStakeUpdate` for the operator for the quorum.

Note that the contract does not check that the quorums that the operator is being deregistered from are a subset of the quorums the operator is registered for, that logic is expected to be done in the RegistryCoordinator.

### updateStakes

An offchain actor can provide a list of operator ids, their corresponding addresses, and a few other witnesses in order to recalculate the stakes of the provided operators for all of the quorums each operator is registered for. This ends block range of the current `OperatorStakeUpdate`s for each of the quorums for each of the provided operators and pushes a new update for each of them.

This has more implications after slashing is enabled... TODO

## Upstream Dependencies

The main integration with the StakeRegistry is used by the AVSs [BLSSignatureChecker](./BLSSignatureChecker.md). An offchain actor provides an operator id, a quorum id, and an index in the array of the operator's stake updates to verify the stake of an operator at a particular block number. They also provide a quorum id and an index in the array of total stake updates to verify the stake of the entire quorum at a particular block number.