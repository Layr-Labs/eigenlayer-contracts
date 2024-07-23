# Allocator

TODO: Handle elimination of operator delegation and deposits

```solidity
interface IDelegationManager {

	/// @notice Emitted when an operator queues a handoff of their stake to a target allocator
	event AllocatorHandoffQueued(address operator, address allocator, uint32 effectTimestamp);

	/**
	 * @notice Queues a handoff of the calling operator's delegated stake to a target allocator in 14 days
	 *
	 * @param allocator the target allocator to handoff delegated stake to
	 * @param approverSignatureAndExpiry the signature of the allocator's delegation approver on their intent to accept the handoff
	 * @param allocatorSalt A unique single use value tied to an individual signature.
	 *
	 * @dev the handoff can be completed in a separate tx in 14 days. it is permissionless to complete
	 * @dev further delegations and deposits to the operator are prohibited after this function is called
	 */
	function queueAllocatorHandoff(address allocator, SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 allocatorSalt) external;
}

interface IAllocatorManager {
	/// STRUCTS
	
	/**
	 * Struct type used to specify an existing queued withdrawal. Rather than storing the entire struct, only a hash is stored.
	 * In functions that operate on existing queued withdrawals -- e.g. completeQueuedWithdrawal`, the data is resubmitted and the hash of the submitted
	 * data is computed by `calculateWithdrawalRoot` and checked against the stored hash in order to confirm the integrity of the submitted data.
	 */
	struct Withdrawal {
    	// The address that originated the Withdrawal
    	address staker;
	    // NEW: The staker's allocator at the time of thhe queing of the withdrawal
    	address allocator;
	    // The address that can complete the Withdrawal + will receive funds when completing the withdrawal
    	address withdrawer;
	    // Nonce used to guarantee that otherwise identical withdrawals have unique hashes
	    uint256 nonce;
    	// Block number when the Withdrawal was created
	    uint32 startBlock;
	    // Array of strategies that the Withdrawal contains
	    IStrategy strategy;
	    // Array containing the amount of shares in each Strategy in the `strategies` array
	    uint256 scaledShares;
	}

	/// EVENTS

	// @notice Emitted when a new allocator is registered in EigenLayer
	event AllocatorRegistered(address indexed allocator, address delegationApprovers);

	/// @notice Emitted when an allocator updates their delegation approver
	event DelegationApproverUpdated(address indexed allocator, address delegationApprover);

	/**
	 * @notice Emitted when @param allocator indicates that they are updating their MetadataURI string
	 * @dev Note that these strings are *never stored in storage* and are instead purely emitted in events for off-chain indexing
	 */
	event AllocatorMetadataURIUpdated(address indexed allocator, string metadataURI);

	/// @notice Emitted whenever an allocator's shares are increased for a given strategy. Note that shares is the delta in the allocator's shares.
	event AllocatorSharesIncreased(address indexed allocator, address staker, IStrategy strategy, uint256 shares);

	/// @notice Emitted whenever an allocator's shares are decreased for a given strategy. Note that shares is the delta in the allocator's shares.
	event AllocatorSharesIncreased(address indexed allocator, address staker, IStrategy strategy, uint256 shares);

	/// @notice Emitted when @param staker delegated to @param allocator as their allocator
	event StakerDelegated(address indexed staker, address indexed allocator);

	/// @notice Emitted when @param staker undelegates from their current allocator
	event StakerUndelegated(address indexed staker, address indexed allocator);

	/// @notice Emitted when @param staker undelegates from an allocator via a call not originating from the staker themself
	event StakerForceUndelegated(address indexed staker, address indexed allocator);

	/// @notice Emitted when a handoff from an operator to an allocator is completed for a set of strategies
	event AllocatorHandoffCompleted(address operator, address allocator, IStrategy[] strategies);
	
	/// EXTERNAL - STATE MODIFYING

	/**
     * @notice Registers the caller as an allocator in EigenLayer.
     * @param delegationApprover is the address that must approve of the delegation of a staker's stake to the allocator. If set to 0, no approval is 
	 * required.
     * @param metadataURI is a URI for the allocator's metadata, i.e. a link providing more details on the allocator.
     *
     * @dev Once an allocator is registered, they cannot 'allocator' as an operator, and they will forever be considered "delegated themself".
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `AllocatorMetadataURIUpdated` event
	 * @dev reverts if the caller is delegated to a pre-SDA operator
     */
    function registerAsAllocator(
        address delegationApprover,
        string calldata metadataURI
    ) external;

	/**
	 * @notice Updates the delegation approver for the calling allocator.
	 *
	 * @param delegationApprover is the address that must approve of the delegations of a staker's stake to the allocator. If set to 0, no approval is
	 * required.
	 */
    function setDelegationApprover(address delegationApprover) external;

    /**
     * @notice Called by an allocator to emit an `AllocatorMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an allocator
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `AllocatorMetadataURIUpdated` event
     */
    function updateAllocatorMetadataURI(string calldata metadataURI) external;

	/**
	 * @notice Delegates to an allocator for the calling staker
	 *
	 * @param allocator the allocator delegated to by the calling staker
	 * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param approverSalt A unique single use value tied to an individual signature.
     * @dev The approverSignatureAndExpiry is used in the event that:
     *          1) the allocator's `delegationApprover` address is set to a non-zero value.
     *                  AND
     *          2) neither the allocator nor their `delegationApprover` is the `msg.sender`, since in the event that the allocator
     *             or their delegationApprover is the `msg.sender`, then approval is assumed.
	 *
	 * @dev Reverts if the allocatorFor the staker is already set or if the staker's M2 operator has handed off to an allocator
	 */
	function delegateTo(
		address allocator,
		SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
	) external;

	/**
     * @notice Caller delegates a staker to an allocator with valid signatures from both parties.
     * @param staker The account delegating stake to an `operator` account
     * @param allocator The account (`staker`) is delegating to
     * @param stakerSignatureAndExpiry Signed data from the staker authorizing assinging stake to an allocator
     * @param approverSignatureAndExpiry is a parameter that will be used for verifying that the allocator approves of this delegation action in the event that:
     * @param approverSalt Is a salt used to help guarantee signature uniqueness. Each salt can only be used once by a given approver.
     *
     * @dev If `staker` is an EOA, then `stakerSignature` is verified to be a valid ECDSA stakerSignature from `staker`, indicating their intention for this action.
     * @dev If `staker` is a contract, then `stakerSignature` will be checked according to EIP-1271.
     * @dev the allocator's `delegationApprover` address is set to a non-zero value.
     * @dev neither the operator nor their `delegationApprover` is the `msg.sender`, since in the event that the operator or their delegationApprover
     * is the `msg.sender`, then approval is assumed.
     * @dev In the case that `approverSignatureAndExpiry` is not checked, its content is ignored entirely; it's recommended to use an empty input
     * in this case to save on complexity + gas costs
     */
    function delegateToBySignature(
        address staker,
        address allocator,
        SignatureWithExpiry memory stakerSignatureAndExpiry,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) external;

	/**
     * @notice Undelegates the staker's currently chosed allocator and queues a withdrawal of all of the staker's shares
     * @param staker The account to be undelegated.
     * @return withdrawalRoot The root of the newly queued withdrawal, if a withdrawal was queued. Otherwise just bytes32(0).
     *
     * @dev Reverts if the `staker` is also an allocator, since allocator are not allowed to undelegate from themselves.
     * @dev Reverts if the caller is not the staker, nor the allocator who the staker is delegated to, nor the allocator's specified "delegationApprover"
     * @dev Reverts if the `staker` is already undelegated.
     */
    function undelegate(address staker) external returns (bytes32[] memory withdrawalRoot);

	/**
	 * @notice Completes a handoff queued via queueHandoff.
	 *
	 * @param operator the operator in the queued handoff
	 * @param allocator the allocator in the queued handoff
	 * @param strategies the strategies to be handed off
	 * 
	 * @dev must be called 14 days after the handoff was queued
	 * @dev the allocator's shares are incremented by the operator's shares for each strategy and the operator's shares are decremented by the operator's
	 * shares for each strategy.
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
	function delegatedTo(address staker) external view returns(address allocator);
}
```

### registerAsAllocator

This is called by accounts that wish to become allocators in EigenLayer. Allocators are responsible for allocating slashable stake to operatorSets. The `delegationApprover` is the address that must approve of the delegation of a staker's stake to the allocator. If set to 0, no approval is required.

Emits an `AllocatorRegistered` event.

Reverts if:
1. The caller is delegated to an pre-SDA operator. This includes pre-SDA operators themselves.

### setDelegationApprover

This is called by an allocator to update their delegation approver. The `delegationApprover` is the address that must approve of the delegation of a staker's stake to the allocator. If set to 0, no approval is required.

Emits a `DelegationApproverUpdated` event.

Reverts if:
1. The caller is not an allocator

### updateAllocatorMetadataURI

This is called by an allocator to update their metadata URI. The `metadataURI` is a URI for the allocator's metadata, i.e. a link providing more details on the allocator.

Emits a `AllocatorMetadataURIUpdated` event.

Reverts if:
1. The caller is not an allocator

### delegateTo

This is called by stakers to delegate to an allocator. The `allocator` is the allocator being delegated to by the calling staker. An EIP-1271 signature from the allocator's delegation approver is required.

Stakers that are delegated to an M2 operator that has queued a handoff to an allocator are not allowed to delegate to an allocator. However, stakers that are delegated to an M2 operator that has not queued a handoff to an allocator are allowed to delegate to an allocator and instanaeously move their stake to the allocator.

Emits a `StakerDelegated` event.

Reverts if:
1. `approverSignatureAndExpiry` is not from the allocator's `delegationApprover` and the allocator's `delegationApprover` is set to a non-zero value
2. the allocator for the staker is already set
3. the staker is delegated to an M2 operator that has queued a handoff to an allocator

### delegateToBySignature

Similar to `delegateTo`, but uses a signature from the staker. The staker's signature is verified to be a valid ECDSA signature from the staker, indicating their intention for this action. If the staker is a contract, the signature will be checked according to EIP-1271.

Emits a `StakerDelegated` event.

Reverts if:
1. `stakerSignatureAndExpiry` is not from the staker
2. `approverSignatureAndExpiry` is not from the allocator's `delegationApprover` and the allocator's `delegationApprover` is set to a non-zero value
3. the allocator for the staker is already set
4. the staker is delegated to an M2 operator that has queued a handoff to an allocator

### undelegate

This is called by stakers to undelegate from their current allocator and queue a withdrawal of all of the staker's shares.

The staker's allocator and the allocator's delegation approver are allowed to undelegate the staker. This is intended to help for compliance reasons.

Emits a `StakerUndelegated` event.

Reverts if:
1. the staker is also an allocator
2. the caller is not the staker, nor the allocator who the staker is delegated to, nor the allocator's specified delegation approver
3. the staker is already undelegation

### queueAllocatorHandoff

This is called by operators to queue a handoff of their delegated stake to a target allocator in 14 days. This allows stake to flow from staker-determined-operators to staker-determined-allocators initially without any staker involvement. A EIP-1271 signature from the allocator's delegation approver is required.

The handoff must be completed in a separate transaction.

Emits a `AllocatorHandoffQueued` event.

Reverts if:
1. the operator is not registered with the DelegationManager
2. the allocator is not registered with the DelegationManager
3. `approverSignatureAndExpiry` is not from the allocator's delegation approver and the allocator's delegation approver is set to a non-zero value
2. the operator has already queued a handoff to an allocator

### completeAllocatorHandoff

This is called permissionlessly by anyone to complete a handoff queued via `queueAllocatorHandoff`. The handoff can be completed 14 days after it was queued. The call accepts an array of strategies to be handed off. If all strategies are not handed off, this function can be called by anyone else to complete the handoff for different strategies.

The allocator's shares are incremented by the operator's shares for each strategy and the operator's shares are decremented by the operator's shares for each strategy.

Emits a `AllocatorHandoffCompleted` event.

Reverts if:
1. the handoff doesn't exist
2. the handoff is not 14 days old
3. at least one of the provided strategies has already been handed off