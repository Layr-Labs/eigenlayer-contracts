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
            MemberSetInfo storage setInfo =
                memberSetInfo[standbyParams[i].operatorSet.avs][operator][standbyParams[i].operatorSet.id];

            // Update the standby status for the given operator set.
            setInfo.onStandby = standbyParams[i].onStandby;

            emit StandbyParamUpdated(operator, standbyParams[i].operatorSet, standbyParams[i].onStandby);
        }
    }

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
        // Assert `operator` is actually an operator.
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToOperatorSets: operator not registered to EigenLayer yet"
        );
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

        if (operatorSignature.signature.length != 0) {
            // Assert signature provided by `operator` is valid.
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

        MemberInfo storage info = memberInfo[msg.sender][operator];

        // Register `operator` if not already registered.
        if (!info.isOperatorForAVS) {
            info.isOperatorForAVS = true;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, true);
        }

        // Loop over `operatorSetIds` array and register `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            MemberSetInfo storage setInfo = memberSetInfo[msg.sender][operator][operatorSetIds[i]];

            // Assert avs is on standby mode for the given `operator` and `operatorSetIds[i]`.
            if (operatorSignature.signature.length == 0) {
                require(setInfo.onStandby, "AVSDirectory.registerOperatorToOperatorSets: avs not on standby");
            }

            // Assert `operator` has not already been registered to `operatorSetIds[i]`.
            require(
                !setInfo.isMember,
                "AVSDirectory.registerOperatorToOperatorSets: operator already registered to operator set"
            );

            // Mutate `isMember` to `true` for `operatorSetIds[i]`.
            setInfo.isMember = true;

            emit OperatorAddedToOperatorSet(operator, OperatorSet({avs: msg.sender, id: operatorSetIds[i]}));
        }

        // Increase `operatorAVSOperatorSetCount` by `operatorSetIds.length`.
        // You would have to add the operator to 2**256-2 operator sets before overflow is possible here.
        unchecked {
            info.registrationCount += uint248(operatorSetIds.length);
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
            MemberSetInfo storage setInfo = memberSetInfo[msg.sender][operator][operatorSetIds[i]];

            // Assert `operator` is registered for this iterations operator set.
            require(
                setInfo.isMember,
                "AVSDirectory.deregisterOperatorFromOperatorSet: operator not registered for operator set"
            );

            // Mutate `isMember` to `false` for `operatorSetIds[i]`.
            setInfo.isMember = false;

            emit OperatorRemovedFromOperatorSet(operator, OperatorSet({avs: msg.sender, id: operatorSetIds[i]}));
        }

        MemberInfo storage info = memberInfo[msg.sender][operator];

        // The above assertion makes underflow logically impossible here.
        unchecked {
            info.registrationCount -= uint248(operatorSetIds.length);
        }

        // Set the operator as deregistered if no longer registered for any operator sets
        if (info.registrationCount == 0) {
            info.isOperatorForAVS = false;
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
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToAVS: operator signature expired"
        );

        MemberInfo storage info = memberInfo[msg.sender][operator];

        require(!info.isOperatorForAVS, "AVSDirectory.registerOperatorToAVS: operator already registered");
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToAVS: salt already spent"
        );
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet"
        );
        require(
            info.registrationCount == 0,
            "AVSDirectory.registerOperatorToAVS: operator set AVS cannot register operators with legacy method"
        );

        // Calculate the digest hash
        bytes32 operatorRegistrationDigestHash = calculateOperatorAVSRegistrationDigestHash({
            operator: operator,
            avs: msg.sender,
            salt: operatorSignature.salt,
            expiry: operatorSignature.expiry
        });

        // forgefmt: disable-next-item
        // Check that the signature is valid
        EIP1271SignatureUtils.checkSignature_EIP1271({
            signer: operator, 
            digestHash: operatorRegistrationDigestHash, 
            signature: operatorSignature.signature
        });

        // Set the operator as registered
        info.isOperatorForAVS = true;

        // Mark the salt as spent
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

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
        MemberInfo storage info = memberInfo[msg.sender][operator];

        require(info.isOperatorForAVS, "AVSDirectory.deregisterOperatorFromAVS: operator not registered");

        // Set the operator as deregistered
        info.isOperatorForAVS = false;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, false);
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
        return _calculateDigestHash(keccak256(abi.encode(OPERATOR_STANDBY_UPDATE, standbyParams, salt, expiry)));
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
