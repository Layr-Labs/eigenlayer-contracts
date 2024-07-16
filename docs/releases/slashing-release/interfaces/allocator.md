# Allocator

The `OperatorDetails` is modified to replace the `earningsReciever` field with the new `allocator` role that allocates slashable stake to operatorSets.

```solidity
interface IDelegationManager {
	/// STRUCTS
	
	// @notice Struct used for storing information about a single operator who has registered with EigenLayer
	struct OperatorDetails {
		// @notice address that can allocate slashable stake and register/deregister the operator to AVS operatorSets
		address allocator;
		/**
		 * @notice Address to verify signatures when a staker wishes to delegate to the operator, as well as controlling "forced undelegations".
		 * @dev Signature verification follows these rules:
		 * 1) If this address is left as address(0), then any staker will be free to delegate to the operator, i.e. no signature verification will be performed.
		 * 2) If this address is an EOA (i.e. it has no code), then we follow standard ECDSA signature verification for delegations to the operator.
		 * 3) If this address is a contract (i.e. it has code) then we forward a call to the contract and verify that it returns the correct EIP-1271 "magic value".
		 */
		address delegationApprover;
		// @notice noop
		uint32 stakerOptOutWindowBlocks;
	}

	/// EVENTS

	/// @notice Emitted when an operator updates their OperatorDetails to @param newOperatorDetails
	event OperatorDetailsModified(address indexed operator, OperatorDetails newOperatorDetails);
	
	/// EXTERNAL - STATE MODIFYING
	
	/**
	* @notice Updates an operator's stored `OperatorDetails`.
	* @param newOperatorDetails is the updated `OperatorDetails` for the operator, to replace their current OperatorDetails`.
	*
	* @dev The caller must have previously registered as an operator in EigenLayer.
	*/
	function modifyOperatorDetails(OperatorDetails calldata newOperatorDetails) external
	
	/// VIEW
	
	/**
	 * @param operator the operator to get the allocator of
	 *
	 * @returns allocator the operator's allocator
	 * @dev an operator's default allocator is themselves if they don't set it
	 */
	function getAllocator(address operator) 
		external view returns(address stakeAllocator);
}
```

### modifyOperatorDetails

Updates the operator details for the operator (`msg.sender`) immediately. Emits a `OperatorDetailsModified` event.

Never reverts.

### getAllocator

Returns the current allocator for the given operator. If the allocator has never been set, it returns the operator.
