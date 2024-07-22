# Allocator

The `OperatorDetails` is modified to replace the `earningsReciever` field with the new `allocator` role that allocates slashable stake to operatorSets.

```solidity
interface IDelegationManager {
	/// EVENTS

	/// @notice Emitted when an operator updates their OperatorDetails to @param newOperatorDetails
	event AllocatorHandoffQueued(address operator, address allocator, uint32 effectTimestamp);

	
	event AllocatorHandoffCompleted(address operator, address allocator, IStrategy[] strategies);
	
	/// EXTERNAL - STATE MODIFYING

	/**
	 * @notice Queues a handoff of the calling operator's delegated stake to a target allocator in 14 days
	 *
	 * @param allocator the target allocator to handoff delegated stake to
	 * @param allocatorSignatureAndExpiry the signature of the allocator
	 *
	 * @dev the handoff must be completed in a separate tx in 14 days. it is permissionless to complete
	 * @dev further delegations and deposits to the operator are prohibited after this function is called
	 */
	function queueAllocatorHandoff(address allocator, SignatureWithExpiry memory allocatorSignatureAndExpiry) external;

	/**
	 * @notice Completes a handoff queued via queueHandoff.
	 *
	 * @param operator the operator in the queued handoff
	 * @param allocator the allocator in the queued handoff
	 * @param strategies the strategies to be handed off
	 * 
	 * @dev must be called 14 days after the handoff was queued
	 * @dev the allocator's shares are incremented by the operators shares for each strategy
	 * @dev if all strategies are not handed off, this function can be called by anyone else to 
	 * complete the handoff for different strategies
	 */
	function completeAllocatorHandoff(address operator, address allocator, IStrategy[] calldata strategies) external;
	
	/// VIEW
	
	/**
	 * @param staker the staker to get the allocator of
	 *
	 * @returns allocator the staker's allocator
	 */
	function allocatorFor(address staker) external view returns(address stakeAllocator);
}
```

### modifyOperatorDetails

Updates the operator details for the operator (`msg.sender`) immediately. Emits a `OperatorDetailsModified` event.

Never reverts.

### getAllocator

Returns the current allocator for the given operator. If the allocator has never been set, it returns the operator.
