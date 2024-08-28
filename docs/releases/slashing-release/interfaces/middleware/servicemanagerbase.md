---
description: >-
  POC can be found here
  https://github.com/Layr-Labs/eigenlayer-middleware/pull/263
---

# ServiceManagerBase

```solidity
interface IServiceManager {

    /// EXTERNAL - STATE MODIFYING

    /**
     * @notice Forwards a call to EigenLayer's AVSDirectory contract to register an operator to operator sets
     * @param operator The address of the operator to register.
     * @param operatorSetIds The IDs of the operator sets.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;

    /**
     * @notice Forwards a call to EigenLayer's AVSDirectory contract to deregister an operator from operator sets
     * @param operator The address of the operator to deregister.
     * @param operatorSetIds The IDs of the operator sets.
     */
    function deregisterOperatorFromOperatorSets(address operator, uint32[] calldata operatorSetIds) external;

    /// MIGRATION FUNCTIONS FOR EXISTING OPERATORS PRIOR TO UPGRADE

    /**
     * @notice Migrates the AVS to use operator sets and creates new operator set IDs.
     * @param operatorSetsToCreate An array of operator set IDs to create.
     * @dev This function can only be called by the contract owner.
     */
    function migrateAndCreateOperatorSetIds(uint32[] memory operatorSetsToCreate) external;

    /**
     * @notice Migrates operators to their respective operator sets.
     * @param operatorSetIds A 2D array where each sub-array contains the operator set IDs for a specific operator.
     * @param operators An array of operator addresses to migrate.
     * @dev This function can only be called by the contract owner.
     * @dev Reverts if the migration has already been finalized.
     */
    function migrateToOperatorSets(uint32[][] memory operatorSetIds, address[] memory operators) external;

    /**
     * @notice Finalizes the migration process, preventing further migrations.
     * @dev This function can only be called by the contract owner.
     * @dev Reverts if the migration has already been finalized.
     */
    function finalizeMigration() external;

    /// VIEW

    /**
     * @notice Retrieves the operators to migrate along with their respective operator set IDs.
     * @return operatorSetIdsToCreate An array of operator set IDs to create.
     * @return operatorSetIds A 2D array where each sub-array contains the operator set IDs for a specific operator.
     * @return allOperators An array of all unique operator addresses.
     */
    function getOperatorsToMigrate()
        external
        view
        returns (
            uint32[] memory operatorSetIdsToCreate,
            uint32[][] memory operatorSetIds,
            address[] memory allOperators
        );

    /// @notice Returns the operator set IDs for the given operator address by querying the RegistryCoordinator
    function getOperatorSetIds(address operator) external view returns (uint32[] memory);
}

```

### AVS Migration Process

1. Perform upgrade of AVS middleware contracts

2. `getOperatorsToMigrate()`: Call the view function `getOperatorsToMigrate` to get (operatorSetIdsToCreate, operatorSetIds, allOperators) which will be used as inputs for the state modifying functions below.

3. `migrateAndCreateOperatorSetIds()`: Call `migrateAndCreateOperatorSetIds` with operatorSetIdsToCreate from above. This will create new operatorSetIds based on the existing quorums in the middleware contracts. This will also disable calling registrations/deregistrations from the old M2 legacy functions `registerOperatorFromAVS`, `deregisterOperatorFromAVS`

4. `migrateToOperatorSets()`: Call `migrateToOperatorSets` and for each operatorSetId migrate the respective operators to them.  Note that this could be called multiple times as needed to account for gas limitations.

5. `finalizeMigration()`: Call to prevent further migrations

## Interfaces

#### **getOperatorsToMigrate**

This function retrieves the operators to migrate along with their respective operatorSetIds which should be created. The operatorSetIds to create will return the existing quorums that exist in the middleware contracts.
The values returned are meant to be used in `migrateOperatorToOperatorSets`.

#### **migrateAndCreateOperatorSetIds**

Only callable by ServiceManager contract owner, this will first call `AVSDirectory.becomeOperatorSetAVS` to mark the AVS as an OperatorSet AVS. This will disable calling the old M2 legacy functions `registerOperatorFromAVS`, `deregisterOperatorFromAVS`.
It will also call `AVSDirectory.createOperatorSets` to create the inputted `operatorSetIdsToCreate` in the EigenLayer protocol enabling for slashable allocations to be configured for the respective operatorSets.

Reverts If:

1. msg.sender is not owner
2. ServiceManagerBase is already an operatorSetAVS in the AVSDirectory
3. Any of `operatorSetIdsToCreate` already exist in AVSDirectory

#### **migrateToOperatorSets**

Only callable by ServiceManager contract owner, `migrateToOperatorSets` will migrate old legacy registered operators to their corresponding newly created operatorSetIds (`migrateAndCreateOperatorSetIds` should have already been called). This function will check that the operators are registered in the middleware quorum numbers before calling `AVSDirectory.migrateOperatorsToOperatorSets`. This function can be called as many times as needed to migrate all operators across multiple transactions to account for gas limitations.

Reverts If:

1. msg.sender is not owner
2. `migrationFinalized` is True
3. `operatorSetIds.length != operators.length`
4. For each operatorSetId to register an operator for, the operator is not registered in the corresponding quorum number
5. `AVSDirectory.migrateOperatorsToOperatorSets` reverts \
    1. ServiceManagerBase is not a OperatorSetAVS
    2. operator is not a legacy registered operator

#### **finalizeMigration**

Only callable by ServiceManager contract owner, this will mark the AVS ServiceManager contract as fully migrated and prevents any future calls to `migrateToOperatorSets`.

1. msg.sender is not owner
2. migrationFinalized is True

#### **registerOperatorToOperatorSets**

Forwards a call from the RegistryCoordinator to the AVSDirectory contract to register an operator for the given operatorSetIds.

Reverts if forwarded call reverts:

1. Any of the operatorSets don't exist
2. `operator` is not registered with the DelegationManager
3. `operator` is already registered for any of the operatorSets
4. `operatorSignature` is not from the `operator`


#### **deregisterOperatorFromOperatorSets**

Forwards a call from the RegistryCoordinator to the AVSDirectory contract to deregister an operator for the given operatorSetIds.

Reverts if forwarded call reverts:

1. `operator` is not registered for at least one of the operatorSets

