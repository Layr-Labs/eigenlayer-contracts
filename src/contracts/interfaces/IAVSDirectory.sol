// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./ISignatureUtils.sol";
import "./IStrategy.sol";

interface IAVSDirectory is ISignatureUtils {
    /// @notice Enum representing the status of an operator's registration with an AVS
    enum OperatorAVSRegistrationStatus {
        UNREGISTERED,       // Operator not registered to AVS
        REGISTERED          // Operator registered to AVS
    }

    /// @notice Enum represnting an operator set for an AVS
    struct OperatorSet {
        address avs;
        uint32 id;
    }

    /**
     * @notice Emitted when @param avs indicates that they are updating their MetadataURI string
     * @dev Note that these strings are *never stored in storage* and are instead purely emitted in events for off-chain indexing
     */
    event AVSMetadataURIUpdated(address indexed avs, string metadataURI);

    /// @notice Emitted when an operator's registration status for an AVS is updated
    event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, OperatorAVSRegistrationStatus status);

    /**
     * @notice Emitted when an operator is added to an operator set
     */
    event OperatorAddedToOperatorSet(address operator, OperatorSet operatorSet);
    
    /**
     * @notice Emitted when an operator is removed from an operator set
     */
	event OperatorRemovedFromOperatorSet(address operator, OperatorSet operatorSet);

    /// @notice Emitted when a strategy is added to an operator set
	event OperatorSetStrategyAdded(OperatorSet operatorSet, IStrategy strategy);
	
	/// @notice Emitted when a strategy is removed from an operator set
	event OperatorSetStrategyRemoved(OperatorSet operatorSet, IStrategy strategy);

    /**
     * @notice Called by an avs to register an operator with the avs.
     * @param operator The address of the operator to register.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;

    /**
     * @notice Called by an avs to deregister an operator with the avs.
     * @param operator The address of the operator to deregister.
     */
    function deregisterOperatorFromAVS(address operator) external;

	/**
	 * @notice Called by AVSs to add an operator to an operator set
	 * 
	 * @param operator the address of the operator to be added to the operator set
	 * @param operatorSetID the ID of the operator set
	 * @param operatorSignature the signature of the operator on their intent to register
	 * @dev msg.sender is used as the AVS
	 * @dev operator must not have a deregistration from the operator set
	 * @dev if this is the first operator set in the AVS that the operator is 
	 * registering for, a OperatorAVSRegistrationStatusUpdated event is emitted with 
	 * a REGISTERED status
	 */
	function registerOperatorToOperatorSet(address operator, uint32 operatorSetID, ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) external;

	/**
	 * @notice Called by AVSs or operators to remove an operator to from operator set
	 * 
	 * @param operator the address of the operator to be removed from the 
	 * operator set
	 * @param operatorSetID the ID of the operator set
	 * 
	 * @dev msg.sender is used as the AVS
	 * @dev operator must be registered for msg.sender AVS and the given 
	 * operator set
     * @dev if this removes operator from all operator sets for the msg.sender AVS
     * then an OperatorAVSRegistrationStatusUpdated event is emitted with a DEREGISTERED
     * status
	 */
	function deregisterOperatorFromOperatorSet(
		address operator, 
		uint32 operatorSetID
	) external;

    /**
     * @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an AVS
     * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `AVSMetadataURIUpdated` event
     */
    function updateAVSMetadataURI(string calldata metadataURI) external;

    /**
     * @notice Returns whether or not the salt has already been used by the operator.
     * @dev Salts is used in the `registerOperatorToAVS` function.
     */
    function operatorSaltIsSpent(address operator, bytes32 salt) external view returns (bool);

    /**
     * @notice Calculates the digest hash to be signed by an operator to register with an AVS
     * @param operator The account registering as an operator
     * @param avs The AVS the operator is registering to
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid
     */
    function calculateOperatorAVSRegistrationDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32);

    /**
     * @notice Calculates the digest hash to be signed by an operator to register with an operator set
     * @param operator The operator set that the operator is registering to
     * @param avs The AVS the operator is registering to
     * @param operatorSetID The ID of the operator set
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid
     */
    function calculateOperatorSetRegistrationDigestHash(
        address operator,
        address avs,
        uint32 operatorSetID,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32);

    /// @notice The EIP-712 typehash for the Registration struct used by the contract
    function OPERATOR_AVS_REGISTRATION_TYPEHASH() external view returns (bytes32);

    /// @notice The EIP-712 typehash for the OperatorSetRegistration struct used by the contract
    function OPERATOR_SET_REGISTRATION_TYPEHASH() external view returns (bytes32);
}
