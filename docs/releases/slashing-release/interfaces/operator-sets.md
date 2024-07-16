# OperatorSets

Registration/Deregistration functions are updated to require the operator's allocator's signature instead of the operator's signature.

```solidity
interface IAVSDirectory {
	
	/// EXTERNAL - STATE MODIFYING

	/**
	 * @notice Called by an AVS to add an operator to a list of operatorSets.
	 * 
	 * @param operator the address of the operator to be added to the operatorSets
	 * @param operatorSetIds the IDs of the operatorSets
	 * @param allocatorSignature the signature of the allocator on their intent to register the operator
	 *
	 * @dev msg.sender is used as the AVS
	 */
	function registerOperatorToOperatorSets(
		address operator,
		uint32[] calldata operatorSetIds,
		ISignatureUtils.SignatureWithSaltAndExpiry memory allocatorSignature
	) external;

	/**
	 * @notice Called by an allocator or on their behalf to deregister an operator from a list of operatorSets for given AVS without AVS intervention.
	 * 
	 * @param operator the operator to deregister from the operatorSets
	 * @param avs the AVS to deregister from
	 * @param operatorSetIDs the list of operatorSets to deregister from
	 * @param allocatorSignature the signature of the allocator on their intent to deregister the operator or empty if the allocator itself is calling
	 *
	 * @dev if the allocatorSignature is empty, the caller must be the allocator
	 * @dev this will likely only be called in case the AVS contracts are in a state that prevents operator from deregistering
	 */
	 function forceDeregisterFromOperatorSets(
		address operator,
		address avs, 
		uint32[] operatorSetIds,
		ISignatureUtils.SignatureWithSaltAndExpiry memory allocatorSignature
	);
}
```

### registerOperatorToOperatorSets

This is called by AVSs that have integrated with operatorSets in order to add an operator to a specified list of operatorSets. A signature from the operator's allocator is required.

Emits a `OperatorAddedToOperatorSet` event for each operatorSet the operator is added to.

Reverts if:

1. Any of the operatorSets don't exist
2. `operator` is not registered with the DelegationManager
3. `operator` is already registered for any of the operatorSets
4. `allocatorSignature` is not from the `operator`'s allocator

### forceDeregisterFromOperatorSets

This is called by an allocator or on their behalf (using a signature) in order to deregister an operator from a list of operatorSets of a given AVS. This is intended to be used in cases where AVS contracts are compromised and don't allow operators to deregister. There is no risk from staying registered, but the functionality is included for aesthetic reasons.

AVSs must update to force deregistrations asynchronously through through their own means (e.g. AVSSync).

Reverts if:

1. `msg.sender` is not the `operator` or `allocatorSignature` is not from `operator`'s allocator
2. `operator` is not registered for at least one of the operatorSets