---
description: >-
  POC can be found here
  https://github.com/Layr-Labs/eigenlayer-middleware/pull/263
---

# ServiceManagerBase

```solidity
interface IServiceManager {

    /// EVENTS
    event OperatorSetStrategiesMigrated(uint32 operatorSetId, IStrategy[] strategies);
    event OperatorMigratedToOperatorSets(address operator, uint32[] indexed operatorSetIds);

    /// EXTERNAL - STATE MODIFYING
    
    /**
     * @notice called by the AVS StakeRegistry whenever a new IStrategy
     * is added to a quorum/operatorSet
     * @dev calls operatorSetManager.addStrategiesToOperatorSet()
     */
    function addStrategiesToOperatorSet(
        uint32 operatorSetId,
        IStrategy[] calldata strategies
    ) external;

    /**
     * @notice called by the AVS StakeRegistry whenever a new IStrategy
     * is removed from a quorum/operatorSet
     * @dev calls operatorSetManager.removeStrategiesFromOperatorSet()
     */
    function removeStrategiesFromOperatorSet(
        uint32 operatorSetId,
        IStrategy[] calldata strategies
    ) external;

    /**
     * @notice One-time call that migrates all existing strategies for each quorum to their respective operator sets
     * Note: a separate migration per operator must be performed to migrate an existing operator to the operator set
     * See migrateOperatorToOperatorSets() below
     * @dev calls operatorSetManager.addStrategiesToOperatorSet()
     */
    function migrateStrategiesToOperatorSets() external;

    /**
     * @notice One-time call to migrate an existing operator to the respective operator sets.
     * The operator needs to provide a signature over the operatorSetIds they are currently registered
     * for. This can be retrieved externally by calling getOperatorSetIds.
     * @param operator the address of the operator to be migrated
     * @param signature the signature of the operator on their intent to migrate
     * @dev calls operatorSetManager.registerOperatorToOperatorSets()
     */
    function migrateOperatorToOperatorSets(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory signature
    ) external;

    /**
     * @notice Once strategies have been migrated and operators have been migrated to operator sets.
     * The ServiceManager owner can eject any operators that have yet to completely migrate fully to operator sets.
     * This final step of the migration process will ensure the full migration of all operators to operator sets.
     * @param operators The list of operators to eject for the given OperatorSet
     * @param operatorSetId This AVS's operatorSetId to eject operators from
     * @dev The RegistryCoordinator MUST set this ServiceManager contract to be the ejector address for this call to succeed
     */
    function ejectNonMigratedOperators(
        address[] calldata operators,
        uint32 operatorSetId
    ) external;

    /**
     * @notice Called by this AVS's RegistryCoordinator to register an operator for its registering operatorSets
     *
     * @param operator the address of the operator to be added to the operator set
     * @param quorumNumbers quorums/operatorSetIds to add the operator to
     * @param signature the signature of the operator on their intent to register
     * @dev msg.sender should be the RegistryCoordinator
     * @dev calls operatorSetManager.registerOperatorToOperatorSets()
     */
    function registerOperatorToOperatorSets(
        address operator,
        bytes calldata quorumNumbers,
        ISignatureUtils.SignatureWithSaltAndExpiry memory signature
    ) external;

    /**
     * @notice Called by this AVS's RegistryCoordinator to deregister an operator for its operatorSets
     *
     * @param operator the address of the operator to be removed from the
     * operator set
     * @param quorumNumbers the quorumNumbers/operatorSetIds to deregister the operator for
     *
     * @dev msg.sender should be the RegistryCoordinator
     * @dev operator must be registered for the given operator sets
     * @dev calls operatorSetManager.deregisterOperatorFromOperatorSets()
     */
    function deregisterOperatorFromOperatorSets(
        address operator,
        bytes calldata quorumNumbers
    ) external;

    /// VIEW

    /// @notice Returns the operator set IDs for the given operator address by querying the RegistryCoordinator
    function getOperatorSetIds(address operator) external view returns (uint32[] memory);
}

```

### AVS Migration Process

1. Perform upgrade of AVS middleware contracts and perform a one-time call of `migrateStrategiesToOperatorSets` to add existing quorum strategies
2. Existing operators are provide a signature which is then passed to `ServiceManager.migrateOperatorToOperatorSets()`  which registers the operator with the OperatorSetManager contract. This can be permissionlessly called by anyone so either the operator can submit the tx themselves or they can provide a signature and EigenLabs or the AVS can migrate the Operator. This is only callable once per nonmigrated Operator that was registered before the upgrade.
3. After a long enough period for Operators to migrate, the ServiceManagerOwner can eject non-migrated operators calling `ejectNonmigratedOperators` . NOTE: the ServiceManager will have to temporarily be set as the `ejector` address for the RegistryCoordinator to perform this step. `ejectNonmigratedOperators` is meant to be called for each Quorum/OperatorSet of the AVS.

## Interfaces

#### **addStrategiesToOperatorSet / removeStrategiesFromOperatorSet**

These two functions are called anytime strategies are added/removed from a quorum in the StakeRegistry. This ensures a callback to the OperatorSetManager to update the strategies in the OperatorSet.

#### **migrateStrategiesToOperatorSets**

This function adds all existing strategies for all existing quorums/operatorSets and pushes it to the OperatorSetManager contract. It can only be called one-time for existing AVS ServiceManagers that existed prior to the OperatorSet upgrade.

#### **migrateOperatorToOperatorSets**

For each existing Operator, to migrate their registration to the OperatorSetManager, they will call this one time with a passed in signature of the operatorSets they are re-registering for.

Reverts If:

* The operator has already migrated
* The call to the OperatorSetManager to register the operator reverts (TODO: link)

#### ejectNonMigratedOperators

Only callable by ServiceManager contract owner, ejectNonmigratedOperators will eject all currently registered operators that haven't migrated their registration with OperatorSets. This requires the RegistryCoordinator of the AVS to set their ejector to temporarily be the ServiceManager contract to allow for this call to succeed.

Reverts If:

* ServiceManager is not the RegistryCoordinator ejector address
* Input param `OperatorSetId >= RegistryCoordinator.quorumCount()`
* The provided operatorSetId is not a valid quorum
* If any of the operators in the input array is registered for the OperatorSet



