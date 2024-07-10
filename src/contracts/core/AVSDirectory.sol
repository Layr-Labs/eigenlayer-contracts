// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "./AVSDirectoryStorage.sol";

contract AVSDirectory is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AVSDirectoryStorage,
    ReentrancyGuardUpgradeable
{
    /// @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    /// @dev Returns the chain ID from the time the contract was deployed.
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManager, slasher,
     * and eigenpodManager contracts
     */
    constructor(IDelegationManager _delegation) AVSDirectoryStorage(_delegation) {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     * minWithdrawalDelayBlocks is set only once here
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _transferOwnership(initialOwner);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /**
     *  @notice Called by AVSs to add an operator to an operator set.
     *
     *  @param operator The address of the operator to be added to the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *  @param operatorSignature The signature of the operator on their intent to register.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev The operator must not have a pending deregistration from the operator set.
     *  @dev If this is the first operator set in the AVS that the operator is
     *  registering for, a OperatorAVSRegistrationStatusUpdated event is emitted with
     *  a REGISTERED status.
     */
    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert operator's signature has not expired.
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToOperatorSets: operator signature expired"
        );
        // Assert operator's signature `salt` has not already been spent.
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToOperatorSets: salt already spent"
        );
        // Assert `operator` is actually an operator.
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToOperatorSets: operator not registered to EigenLayer yet"
        );
        
        // If a signature is not provided...
        if (operatorSignature.signature.length != 0) {
            // Assert that `operatorSignature.signature` is a valid signature for operator set registrations.
            EIP1271SignatureUtils.checkSignature_EIP1271(
                operator,
                calculateOperatorSetRegistrationDigestHash({
                    avs: msg.sender,
                    operatorSetIds: operatorSetIds,
                    salt: operatorSignature.salt,
                    expiry: operatorSignature.expiry
                }),
                operatorSignature.signature
            );
        }

        // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        MemberInfo storage member = memberInfo[msg.sender][operator];

        // Register `operator` if not already registered.
        if (!member.isLegacyOperator) {
            member.isLegacyOperator = true;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, true);
        }

        // Loop over `operatorSetIds` array and register `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            MemberSetInfo storage setMember = memberSetInfo[msg.sender][operator][operatorSetIds[i]];

            // Assert avs is on standby mode for the given `operator` and `operatorSetIds[i]`.
            if (operatorSignature.signature.length == 0) {
                require(setMember.onStandby, "AVSDirectory.registerOperatorToOperatorSets: avs not on standby");
            }

            // Assert `operator` has not already been registered to `operatorSetIds[i]`.
            require(
                !setMember.isSetOperator,
                "AVSDirectory.registerOperatorToOperatorSets: operator already registered to operator set"
            );

            // Mutate `setMember.isSetOperator` to `true`.
            setMember.isSetOperator = true;

            emit OperatorAddedToOperatorSet(operator, OperatorSet({avs: msg.sender, id: operatorSetIds[i]}));
        }

        // Increase `member.inTotalSets` by `operatorSetIds.length`.
        // You would have to add the operator to 2**256-2 operator sets before overflow is possible here.
        unchecked {
            member.inTotalSets += uint248(operatorSetIds.length);
        }
    }

    /**
     *  @notice Called by AVSs or operators to remove an operator from an operator set.
     *
     *  @param operator The address of the operator to be removed from the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev The operator must be registered for the msg.sender AVS and the given operator set.
     *  @dev If this call removes the operator from all operator sets for the msg.sender AVS,
     *  then an OperatorAVSRegistrationStatusUpdated event is emitted with a DEREGISTERED status.
     */
    function deregisterOperatorFromOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Loop over `operatorSetIds` array and deregister `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            MemberSetInfo storage setMember = memberSetInfo[msg.sender][operator][operatorSetIds[i]];

            // Assert `operator` is registered for this iterations operator set.
            require(
                setMember.isSetOperator,
                "AVSDirectory.deregisterOperatorFromOperatorSet: operator not registered for operator set"
            );

            // Mutate `setMember.isSetOperator` to `false`.
            setMember.isSetOperator = false;

            emit OperatorRemovedFromOperatorSet(operator, OperatorSet({avs: msg.sender, id: operatorSetIds[i]}));
        }

        MemberInfo storage member = memberInfo[msg.sender][operator];

        // The above assertion makes underflow logically impossible here.
        unchecked {
            member.inTotalSets -= uint248(operatorSetIds.length);
        }

        // Set the operator as deregistered if no longer registered for any operator sets
        if (member.inTotalSets == 0) {
            member.isLegacyOperator = false;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, false);
        }
    }

    /**
     *  @notice Called by the AVS's service manager contract to register an operator with the AVS.
     *
     *  @param operator The address of the operator to register.
     *  @param operatorSignature The signature, salt, and expiry of the operator's signature.
     *
     *  @dev msg.sender must be the AVS.
     *  @dev Only used by legacy M2 AVSs that have not integrated with operator sets.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert `operatorSignature.expiry` has not elapsed.
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToAVS: operator signature expired"
        );

        MemberInfo storage member = memberInfo[msg.sender][operator];

        // Assert `operator` is not already registered to the (caller) avs.
        require(!member.isLegacyOperator, "AVSDirectory.registerOperatorToAVS: operator already registered");

        // Assert `operator` is not actively registered to any operator sets.
        require(
            member.inTotalSets == 0,
            "AVSDirectory.registerOperatorToAVS: operator set AVS cannot register operators with legacy method"
        );

        // Assert `operator` has not already spent `operatorSignature.salt`.
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToAVS: salt already spent"
        );

        // Assert `operator` is a registered operator.
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet"
        );

        // Assert that `operatorSignature.signature` is a valid signature for the operator AVS registration.
        EIP1271SignatureUtils.checkSignature_EIP1271({
            signer: operator,
            digestHash: calculateOperatorAVSRegistrationDigestHash({
                operator: operator,
                avs: msg.sender,
                salt: operatorSignature.salt,
                expiry: operatorSignature.expiry
            }),
            signature: operatorSignature.signature
        });

        // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;
        
        // Mutate `member.isLegacyOperator` to `true`.
        member.isLegacyOperator = true;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, true);
    }

    /**
     *  @notice Called by an AVS to deregister an operator from the AVS.
     *
     *  @param operator The address of the operator to deregister.
     *
     *  @dev Only used by legacy M2 AVSs that have not integrated with operator sets.
     */
    function deregisterOperatorFromAVS(address operator)
        external
        onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS)
    {
        MemberInfo storage member = memberInfo[msg.sender][operator];

        // Assert `member.isLegacyOperator` is actively registered to the (caller) avs.
        require(member.isLegacyOperator, "AVSDirectory.deregisterOperatorFromAVS: operator not registered");

        // Mutate `member.isLegacyOperator` to `false`.
        member.isLegacyOperator = false;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, false);
    }

    /**
     * @notice Updates the standby parameters for an operator across multiple operator sets.
     * Allows the AVS to add the operator to a given operator set if they are not already registered.
     *
     * @param operator The address of the operator for which the standby parameters are being updated.
     * @param standbyParams The new standby parameters for the operator.
     * @param operatorSignature If non-empty, the signature of the operator authorizing the modification.
     *                  If empty, the `msg.sender` must be the operator.
     */
    function updateStandbyParams(
        address operator,
        StandbyParam[] calldata standbyParams,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {
        // Check if the operator's signature is provided or not.
        if (operatorSignature.signature.length == 0) {
            // If the signature length is zero, assert that the caller is the operator.
            require(msg.sender == operator, "AVSDirectory.updateStandbyParams: invalid signature");
        } else {
            // If a signature is provided, perform the following checks:

            // Assert that the operator's signature has not expired.
            require(
                operatorSignature.expiry >= block.timestamp,
                "AVSDirectory.updateStandbyParams: operator signature expired"
            );

            // Assert that the operator's signature cannot be replayed.
            require(
                !operatorSaltIsSpent[operator][operatorSignature.salt], "AVSDirectory.updateStandbyParams: salt spent"
            );

            // Assert that the operator's signature is valid.
            EIP1271SignatureUtils.checkSignature_EIP1271({
                signer: operator,
                digestHash: calculateUpdateStandbyDigestHash(
                    standbyParams, operatorSignature.salt, operatorSignature.expiry
                ),
                signature: operatorSignature.signature
            });

            // Mark the salt as spent to prevent replay attacks.
            operatorSaltIsSpent[operator][operatorSignature.salt] = true;
        }

        // Loop through each standby parameter and update the state accordingly.
        for (uint256 i; i < standbyParams.length; ++i) {
            MemberSetInfo storage setMember =
                memberSetInfo[standbyParams[i].operatorSet.avs][operator][standbyParams[i].operatorSet.id];

            // Update the standby status for the given operator set.
            setMember.onStandby = standbyParams[i].onStandby;

            emit StandbyParamUpdated(operator, standbyParams[i].operatorSet, standbyParams[i].onStandby);
        }
    }

    /**
     *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     *
     *  @param metadataURI The URI for metadata associated with an AVS.
     *
     *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
     */
    function updateAVSMetadataURI(string calldata metadataURI) external {
        emit AVSMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(bytes32 salt) external {
        // Mutate `operatorSaltIsSpent` to `true` to prevent future spending.
        operatorSaltIsSpent[msg.sender][salt] = true;
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     *  @notice Calculates the digest hash to be signed by an operator to register with an AVS.
     *
     *  @param standbyParams The newly updated standby mode parameters.
     *  @param salt A unique and single-use value associated with the approver's signature.
     *  @param expiry The time after which the approver's signature becomes invalid.
     */
    function calculateUpdateStandbyDigestHash(
        StandbyParam[] calldata standbyParams,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return
            _calculateDigestHash(keccak256(abi.encode(OPERATOR_STANDBY_UPDATE_TYPEHASH, standbyParams, salt, expiry)));
    }

    /**
     *  @notice Calculates the digest hash to be signed by an operator to register with an AVS.
     *
     *  @param operator The account registering as an operator.
     *  @param avs The AVS the operator is registering with.
     *  @param salt A unique and single-use value associated with the approver's signature.
     *  @param expiry The time after which the approver's signature becomes invalid.
     */
    function calculateOperatorAVSRegistrationDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return
            _calculateDigestHash(keccak256(abi.encode(OPERATOR_AVS_REGISTRATION_TYPEHASH, operator, avs, salt, expiry)));
    }

    /**
     * @notice Calculates the digest hash to be signed by an operator to register with an operator set.
     *
     * @param avs The AVS that operator is registering to operator sets for.
     * @param operatorSetIds An array of operator set IDs the operator is registering to.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateOperatorSetRegistrationDigestHash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(OPERATOR_SET_REGISTRATION_TYPEHASH, avs, operatorSetIds, salt, expiry))
        );
    }

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() public view returns (bytes32) {
        return _calculateDomainSeparator();
    }

    /// @notice Internal function for calculating the current domain separator of this contract
    function _calculateDomainSeparator() internal view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
        }
    }

    function _calculateDigestHash(bytes32 structHash) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _calculateDomainSeparator(), structHash));
    }
}
