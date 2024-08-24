# OperatorSets

```solidity
interface IAVSDirectory {

	/// STRUCTS & ENUMS
	
	/// @notice Enum representing the status of an operator's registration with an AVS
	enum OperatorAVSRegistrationStatus {
 		UNREGISTERED,       // Operator not registered to AVS
 		REGISTERED          // Operator registered to AVS
 	}

	/// @notice Enum representing the registration status of an operator with an AVS.
    /// @notice Only used by legacy M2 AVSs that have not integrated with operatorSets.
    enum OperatorAVSRegistrationStatus {
        UNREGISTERED, // Operator not registered to AVS
        REGISTERED // Operator registered to AVS
    }
    
	/// @notice Struct representing an operator set
	struct OperatorSet {
		address avs;
		uint32 operatorSetId;
	}
	
	/// EVENTS

	/// @notice Emitted when an operator set is created by an AVS.
    event OperatorSetCreated(OperatorSet operatorSet);

	/**
     *  @notice Emitted when an operator's registration status with an AVS id udpated
     *  @notice Only used by legacy M2 AVSs that have not integrated with operatorSets.
     */
    event OperatorAVSRegistrationStatusUpdated(
        address indexed operator, address indexed avs, OperatorAVSRegistrationStatus status
    );

	/// @notice Emitted when an operator is added to an operator set.
    event OperatorAddedToOperatorSet(address indexed operator, OperatorSet operatorSet);

	/// @notice Emitted when an operator is removed from an operator set.
    event OperatorRemovedFromOperatorSet(address indexed operator, OperatorSet operatorSet);

	/// @notice Emitted when an AVS updates their metadata URI (Uniform Resource Identifier).
    /// @dev The URI is never stored; it is simply emitted through an event for off-chain indexing.
    event AVSMetadataURIUpdated(address indexed avs, string metadataURI);

    /// @notice Emitted when an AVS migrates to using operator sets.
    event AVSMigratedToOperatorSets(address indexed avs);

    /// @notice Emitted when an operator is migrated from M2 registration to operator sets.
    event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds);
	
	/// EXTERNAL - STATE MODIFYING

	/**
     * @notice Called by an AVS to create a list of new operatorSets.
     *
     * @param operatorSetIds The IDs of the operator set to initialize.
     *
     * @dev msg.sender must be the AVS.
     * @dev The AVS may create operator sets before it becomes an operator set AVS.
     */
    function createOperatorSets(uint32[] calldata operatorSetIds) external;
	
	/**
     * @notice Sets the AVS as an operator set AVS, preventing legacy M2 operator registrations.
     *
     * @dev msg.sender must be the AVS.
     */
    function becomeOperatorSetAVS() external;

	/**
	 * @notice Called by an AVS to register an operator with the AVS.
	 * 
	 * @param operator The address of the operator to register.
	 * @param operatorSignature The signature, salt, and expiry of the operator's signature.
	 * 
	 * @dev msg.sender is the AVS
	 * @dev only used by legacy M2 AVSs that haven't integrated with operatorSets
	 */
	function registerOperatorToAVS(
		address operator,
		ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
	) external;
	
	/**
	 * @notice Called by an AVS to deregister an operator with the AVS.
	 * 
	 * @param operator The address of the operator to deregister.
	 * 
	 * @dev only used by legacy M2 AVSs that haven't integrated with operatorSets
	 */
	function deregisterOperatorFromAVS(address operator) external;

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

	/**
     * @notice Called by AVSs to add an operator to list of operatorSets.
     *
     * @param operator The address of the operator to be added to the operator set.
     * @param operatorSetIds The IDs of the operator sets.
     * @param operatorSignature The signature of the operator on their intent to register.
     *
     * @dev msg.sender is used as the AVS.
     * @dev The operator must not have a pending deregistration from the operator set.
     */
    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;
	
	/**
     *  @notice Called by AVSs to remove an operator from an operator set.
     *
     *  @param operator The address of the operator to be removed from the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *
     *  @dev msg.sender is used as the AVS.
     */
    function deregisterOperatorFromOperatorSets(
		address operator,
		uint32[] calldata operatorSetIds
	) external;

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
	
	// VIEW
	
	/**
	 * @param avs to check operator set existence for
	 * @param operatorSetId of the operatorSet to check existence for
	 *
	 * @return whether the operatorSet exists
	 * 
	 * @dev This is a mapping in storage
	 */
	function isOperatorSet(address avs, uint32 operatorSetId) returns (bool);
	
	/**
	 * @param operator the operator to get the status of
	 * @param avs to check operatorSet membership in
	 * @param operatorSetId of the operatorSet to check whether the operator was in
	 * 
	 * @return whether the operator was in a given operatorSet
	 *
	 * @dev This is a mapping in storage
	 */
	function isMember(
		address operator,
		OperatorSet memory operatorSet
	) external view returns (bool);



	/**
	 * @param avs to check operatorSet status of
	 * 
	 * @return whether the avs is an operatorSet AVS
	 *
	 * @dev This is a mapping in storage
	 */
	function isOperatorSetAVS(address avs) external view returns(bool);
}
```
### createOperatorSets

This is called by an AVS to create a group of operatorSets.

Reverts if:

1. OperatorSets with the any of the given `operatorSetIds` already exist

### registerOperatorToAVS - M2 Legacy

This is called by AVSs when operators initially register for their AVS. It requires a [EIP 1271](https://eips.ethereum.org/EIPS/eip-1271) signature from the operator. This is used in legacy M2 AVSs. Once the AVS begins migrating to operatorSet based registration, it can no longer use the legacy functions. 

Reverts if:

1. The `operatorSignature` is not from `operator`
2. If `operator` is not registered with the DelegationManager
3. If `operator` is already registered with the AVS
4. The AVS is an operatorSet AVS.

### deregisterOperatorFromAVS - M2 Legacy

This is called by AVSs when operators deregister from their AVS. This is used in legacy M2 AVSs.

Reverts if:

1. the operator is not registered for the AVS

### registerOperatorToOperatorSets

This is called by AVSs that have integrated with operatorSets in order to add an operator to a specified list of operatorSets. A signature from the operator is required.

Emits a `OperatorAddedToOperatorSet` event for each operatorSet the operator is added to.

Reverts if:

1. Any of the operatorSets don't exist
2. `operator` is not registered with the DelegationManager
3. `operator` is already registered for any of the operatorSets
4. `operatorSignature` is not from the `operator`

### migrateOperatorsToOperatorSets

This is called by an AVS to migrate its M2 legacy registered operators to a list of operator sets. This function can only be called once for a given operator that is already registered to the AVS. No consent is required by the operator. If the operator disagrees with the migration, it can unilaterally remove itself from any operatorSet. The operator is deregistered from legacy M2 registration once this function is called - the `OperatorAVSRegistrationStatus` with the UNREGISTERED parameter is emitted. 

Emits a `OperatorAddedToOperatorSet` event for each operatorSet the operator is added to. Also emits an `OperatorMigratedToOperatorSets` for all operatorSets the operator is migrated to. 

Reverts If:
1. The operator does not have a legacy M2 registration with the AVS
2. The operator has already been migrated
3. Any of the target operatorSets have not been created
4. The operator has already been registered for any of the operatorSets
5. The AVS is not a operatorSet AVS, i.e isOperatorSetAVS(avs) returns false


### deregisterOperatorFromOperatorSets

This is called by AVSs that have integrated with operatorSets in order to remove an operator from a list of operatorSets. Note that deregistering doesn't immediately void your slashable stake from the operatorSet as there is a slashable window that extends past your deregistration timestamp. 

Emits a `OperatorRemovedFromOperatorSet` event for each operatorSet the operator is removed from.

Reverts if:

1. `operator` is not registered for at least one of the operatorSets

### forceDeregisterFromOperatorSets

This is called by operator or on their behalf (using a signature) in order to deregister from a list of operatorSets of a given AVS. This is intended to be used in cases where AVS contracts are compromised and don't allow operators to deregister. As noted in `deregisterOperatorFromOperatorSets`, operators still remain slashable for a slashable time window after deregistering. This is an accepted risk from malicious AVSs that must be considered by operators when they register.

AVSs must update to force deregistrations asynchronously through through their own means (e.g. AVSSync).

Reverts if:

1. `msg.sender` is not the `operator` or `operatorSignature` is not from `operator`
2. `operator` is not registered for at least one of the operatorSets

### isOperatorSet

Returns whether an operatorSet has been created.

### isMember

Returns whether an operator is in an operatorSet.
