# AVSDirectory

## Overview

The AVSDirectory contract is where registration relationships are defined between AVSs, operatorSets, and operators. Registration and deregistration are used in the protocol to activate and deactivate slashable stake allocations. They're also used to make the protocol more legible to external integrations.

The slashing release introduces the concept of operatorSets, which are simply an (address, uint32) pair that the define an AVS and an operator set ID. OperatorSets are used to group operators by different tasks and sets of tokens. For example, EigenDA has an ETH/LST operatorSet and an Eigen operatorSet. A bridge may have on operatorSet for all operators that serve a particular chain. Overall, operatorSets are mainly used for protocol legibility.

Functionality is provided for AVSs to migrate from an pre-operatorSet registration model to an operatorSet model. Direct to AVS registration is still supported for AVSs that have not migrated to the operatorSet model, but is slated to be deprecated soon in the future.

## `becomeOperatorSetAVS`
```solidity
/**
 * @notice Sets the AVS as an operator set AVS, preventing legacy M2 operator registrations.
 *
 * @dev msg.sender must be the AVS.
 */
function becomeOperatorSetAVS() external;
```

AVSs call this to become an operator set AVS. Once an AVS becomes an operator set AVS, they can no longer register operators via the legacy M2 registration path. This is a seperate function to help avoid accidental migrations to the operator set AVS model.

## `createOperatorSets`
```solidity
/**
 * @notice Called by an AVS to create a list of new operatorSets.
 *
 * @param operatorSetIds The IDs of the operator set to initialize.
 *
 * @dev msg.sender must be the AVS.
 */
function createOperatorSets(
    uint32[] calldata operatorSetIds
) external;
```

AVSs use this function to create a list of new operator sets.They must call this function before they add any operators to the operator sets. The operator set IDs must be not already exist.

This can be called before the AVS becomes an operator set AVS. (TODO: we should make this so that it can only be called after the AVS becomes an operator set AVS?)

## `migrateOperatorsToOperatorSets`
```solidity
/**
 * @notice Called by an AVS to migrate operators that have a legacy M2 registration to operator sets.
 *
 * @param operators The list of operators to migrate
 * @param operatorSetIds The list of operatorSets to migrate the operators to
 *
 * @dev The msg.sender used is the AVS
 * @dev The operator can only be migrated at most once per AVS
 * @dev The AVS can no longer register operators via the legacy M2 registration path once it begins migration
 * @dev The operator is deregistered from the M2 legacy AVS once migrated
 */
function migrateOperatorsToOperatorSets(
    address[] calldata operators,
    uint32[][] calldata operatorSetIds
) external;
```

AVSs that launched before the slashing release can use this function to migrate operators that have a legacy M2 registration to operator sets. Each operator can only be migrated once for the AVS and the AVS can no longer register operators via the legacy M2 registration path once it begins migration.

## `registerOperatorToOperatorSets`
```solidity
/**
 *  @notice Called by AVSs to add an operator to list of operatorSets.
 *
 *  @param operator The address of the operator to be added to the operator set.
 *  @param operatorSetIds The IDs of the operator sets.
 *  @param operatorSignature The signature of the operator on their intent to register.
 *
 *  @dev msg.sender is used as the AVS.
 */
function registerOperatorToOperatorSets(
    address operator,
    uint32[] calldata operatorSetIds,
    ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
) external;
```

AVSs use this function to add an operator to a list of operator sets. The operator's signature is required to confirm their intent to register.  If the operator has a slashable stake allocation to the AVS, it takes effect when the operator is registered (and up to `DEALLOCATION_DELAY` seconds after the operator is deregistered).

The operator set must exist before the operator can be added to it and the AVS must be an operator set AVS.

## `deregisterOperatorFromOperatorSets`
```solidity
/**
 *  @notice Called by AVSs to remove an operator from an operator set.
 *
 *  @param operator The address of the operator to be removed from the operator set.
 *  @param operatorSetIds The IDs of the operator sets.
 *
 *  @dev msg.sender is used as the AVS.
 */
function deregisterOperatorFromOperatorSets(address operator, uint32[] calldata operatorSetIds) external;
```

AVSs use this function to remove an operator from an operator set. The operator is still slashable for its slashable stake allocation to the AVS until `DEALLOCATION_DELAY` seconds after the operator is deregistered.

The operator must be registered to the operator set before they can be deregistered from it.


## `forceDeregisterFromOperatorSets`
```solidity
/**
 * @notice Called by an operator to deregister from an operator set
 *
 * @param operator The operator to deregister from the operatorSets.
 * @param avs The address of the AVS to deregister the operator from.
 * @param operatorSetIds The IDs of the operator sets.
 * @param operatorSignature the signature of the operator on their intent to deregister or empty if the operator itself is calling
 *
 * @dev if the operatorSignature is empty, the caller must be the operator
 * @dev this will likely only be called in case the AVS contracts are in a state that prevents operators from deregistering
 */
function forceDeregisterFromOperatorSets(
    address operator,
    address avs,
    uint32[] calldata operatorSetIds,
    ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
) external;
```

Operators can use this function to deregister from an operator set without requiring the AVS to sign off on the deregistration. This function is intended to be used in cases where the AVS contracts are in a state that prevents operators from deregistering (either malicious or unintentional).

Operators can also deallocate their slashable stake allocation seperately to avoid slashing risk, so this function is mainly for external integrations to interpret the correct state of the protocol.

## `updateAVSMetadataURI`
```solidity
/**
 *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
 *
 *  @param metadataURI The URI for metadata associated with an AVS.
 *
 *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
 */
function updateAVSMetadataURI(
    string calldata metadataURI
) external;
```

This function allows an AVS to update the metadata URI associated with the AVS. The metadata URI is never stored on-chain and is only emitted in the `AVSMetadataURIUpdated` event.

## View Functions

See the [AVS Directory Inteface](../../../src/contracts/interfaces/IAVSDirectory.sol) for view functions. 