# Allocator

The `OperatorDetails` is modified to replace the `earningsReciever` field with the new `allocator` role that allocates slashable stake to operatorSets.

```solidity
interface IDelegationManager {
	/// STRUCTS

	/**
     * Struct type used to specify an existing queued withdrawal. Rather than storing the entire struct, only a hash is stored.
     * In functions that operate on existing queued withdrawals -- e.g. completeQueuedWithdrawal`, the data is resubmitted and the hash of the submitted
     * data is computed by `calculateWithdrawalRoot` and checked against the stored hash in order to confirm the integrity of the submitted data.
     */
    struct Withdrawal {
        // The address that originated the Withdrawal
        address staker;
        // The staker's allocator at the time of thhe queing of the withdrawal
        address allocator;
        // The address that can complete the Withdrawal + will receive funds when completing the withdrawal
        address withdrawer;
        // Nonce used to guarantee that otherwise identical withdrawals have unique hashes
        uint256 nonce;
        // Block number when the Withdrawal was created
        uint32 startBlock;
        // Array of strategies that the Withdrawal contains
        IStrategy[] strategies;
        // Array containing the amount of shares in each Strategy in the `strategies` array
        uint256[] shares;
    }

	/// EVENTS

	// @notice Emitted when a new allocator is registered in EigenLayer
    event AllocatorRegistered(address indexed allocator, address selectionApprover);

    /// @notice Emitted when an allocator updates their selection approver
    event SelectionApproverUpdated(address indexed allocator, address selectionApprover);

	/**
     * @notice Emitted when @param allocator indicates that they are updating their MetadataURI string
     * @dev Note that these strings are *never stored in storage* and are instead purely emitted in events for off-chain indexing
     */
    event AllocatorMetadataURIUpdated(address indexed allocator, string metadataURI);

	/// @notice Emitted when an operator queues a handoff of their stake to a target allocator
	event AllocatorHandoffQueued(address operator, address allocator, uint32 effectTimestamp);

	/// @notice Emitted when a handoff from an operator to an allocator is completed for a set of strategies
	event AllocatorHandoffCompleted(address operator, address allocator, IStrategy[] strategies);

	/// @notice Emitted whenever an allocator's shares are increased for a given strategy. Note that shares is the delta in the allocator's shares.
    event AllocatorSharesIncreased(address indexed allocator, address staker, IStrategy strategy, uint256 shares);

    /// @notice Emitted whenever an allocator's shares are decreased for a given strategy. Note that shares is the delta in the allocator's shares.
    event AllocatorSharesIncreased(address indexed allocator, address staker, IStrategy strategy, uint256 shares);

	/// @notice Emitted when @param staker selects @param allocator as their allocator
    event StakerSelected(address indexed staker, address indexed allocator);

    /// @notice Emitted when @param staker unselects their current allocator
    event StakerUnselected(address indexed staker, address indexed allocator);

    /// @notice Emitted when @param staker unselects an allocator via a call not originating from the staker themself
    event StakerForceUnselected(address indexed staker, address indexed allocator);
	
	/// EXTERNAL - STATE MODIFYING

	/**
     * @notice Registers the caller as an allocator in EigenLayer.
     * @param selectionApprover is the address that must approve of the selection of a staker's stake to the allocator. If set to 0, no approval is 
	 * required.
     * @param metadataURI is a URI for the allocator's metadata, i.e. a link providing more details on the allocator.
     *
     * @dev Once an allocator is registered, they cannot 'allocator' as an operator, and they will forever be considered "select themself".
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `AllocatorMetadataURIUpdated` event
	 * @dev revert if the caller is already an operator
     */
    function registerAsAllocator(
        address selectionApprover,
        string calldata metadataURI
    ) external;

	/**
	 * @notice Updates the selection approver for the calling allocator.
	 *
	 * @param selectionApprover is the address that must approve of the selection of a staker's stake to the allocator. If set to 0, no approval is
	 * required.
	 */
    function setSelectionApprover(address selectionApprover) external;

    /**
     * @notice Called by an allocator to emit an `AllocatorMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an allocator
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `AllocatorMetadataURIUpdated` event
     */
    function updateAllocatorMetadataURI(string calldata metadataURI) external;

	/**
	 * @notice Queues a handoff of the calling operator's delegated stake to a target allocator in 14 days
	 *
	 * @param allocator the target allocator to handoff delegated stake to
	 * @param allocatorSignatureAndExpiry the signature of the allocator
	 * @param allocatorSalt A unique single use value tied to an individual signature.
	 *
	 * @dev the handoff must be completed in a separate tx in 14 days. it is permissionless to complete
	 * @dev further delegations and deposits to the operator are prohibited after this function is called
	 */
	function queueAllocatorHandoff(address allocator, SignatureWithExpiry memory allocatorSignatureAndExpiry, bytes32 allocatorSalt) external;

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

	/**
	 * @notice Selects an allocator for the calling staker
	 *
	 * @param allocator the allocator selected by the calling staker
	 * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param approverSalt A unique single use value tied to an individual signature.
     * @dev The approverSignatureAndExpiry is used in the event that:
     *          1) the allocator's `selectionApprover` address is set to a non-zero value.
     *                  AND
     *          2) neither the allocator nor their `selectionApprover` is the `msg.sender`, since in the event that the allocator
     *             or their selectionApprover is the `msg.sender`, then approval is assumed.
	 *
	 * @dev Reverts if the allocatorFor the staker is already set or if the staker's M2 operator has handed off to an allocator
	 */
	function select(address allocator) external;

	/**
     * @notice Caller selects an allocator for a  a staker with valid signatures from both parties.
     * @param staker The account delegating stake to an `operator` account
     * @param allocator The account (`staker`) is selecting for allocation
     * @param stakerSignatureAndExpiry Signed data from the staker authorizing assinging stake to an allocator
     * @param approverSignatureAndExpiry is a parameter that will be used for verifying that the allocator approves of this delegation action in the event that:
     * @param approverSalt Is a salt used to help guarantee signature uniqueness. Each salt can only be used once by a given approver.
     *
     * @dev If `staker` is an EOA, then `stakerSignature` is verified to be a valid ECDSA stakerSignature from `staker`, indicating their intention for this action.
     * @dev If `staker` is a contract, then `stakerSignature` will be checked according to EIP-1271.
     * @dev the allocator's `selectionApprover` address is set to a non-zero value.
     * @dev neither the operator nor their `selectionApprover` is the `msg.sender`, since in the event that the operator or their selectionApprover
     * is the `msg.sender`, then approval is assumed.
     * @dev In the case that `approverSignatureAndExpiry` is not checked, its content is ignored entirely; it's recommended to use an empty input
     * in this case to save on complexity + gas costs
     */
    function selectBySignature(
        address staker,
        address operator,
        SignatureWithExpiry memory stakerSignatureAndExpiry,
        SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) external;
	
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
