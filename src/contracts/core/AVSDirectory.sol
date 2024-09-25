// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

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
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

    /// @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;
    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS = 1;

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
    constructor(
        IDelegationManager _delegation
    ) AVSDirectoryStorage(_delegation) {
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
     *                    EXTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Called by an AVS to create a list of new operatorSets.
     *
     * @param operatorSetIds The IDs of the operator set to initialize.
     *
     * @dev msg.sender must be the AVS.
     * @dev The AVS may create operator sets before it becomes an operator set AVS.
     */
    function createOperatorSets(
        uint32[] calldata operatorSetIds
    ) external {
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            require(!isOperatorSet[msg.sender][operatorSetIds[i]], InvalidOperatorSet());
            isOperatorSet[msg.sender][operatorSetIds[i]] = true;
            emit OperatorSetCreated(OperatorSet({avs: msg.sender, operatorSetId: operatorSetIds[i]}));
        }
    }

    /**
     * @notice Sets the AVS as an operator set AVS, preventing legacy M2 operator registrations.
     *
     * @dev msg.sender must be the AVS.
     */
    function becomeOperatorSetAVS() external {
        require(!isOperatorSetAVS[msg.sender], InvalidAVS());
        isOperatorSetAVS[msg.sender] = true;
        emit AVSMigratedToOperatorSets(msg.sender);
    }

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
    ) external override onlyWhenNotPaused(PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS) {
        // Assert that the AVS is an operator set AVS.
        require(isOperatorSetAVS[msg.sender], InvalidAVS());

        for (uint256 i = 0; i < operators.length; i++) {
            // Assert that the operator is registered & has not been migrated.
            require(
                avsOperatorStatus[msg.sender][operators[i]] == OperatorAVSRegistrationStatus.REGISTERED,
                InvalidOperator()
            );

            // Migrate operator to operator sets.
            _registerToOperatorSets(operators[i], msg.sender, operatorSetIds[i]);

            // Deregister operator from AVS - this prevents the operator from being migrated again since
            // the AVS can no longer use the legacy M2 registration path
            avsOperatorStatus[msg.sender][operators[i]] = OperatorAVSRegistrationStatus.UNREGISTERED;
            emit OperatorAVSRegistrationStatusUpdated(
                operators[i], msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED
            );
            emit OperatorMigratedToOperatorSets(operators[i], msg.sender, operatorSetIds[i]);
        }
    }

    /**
     *  @notice Called by AVSs to add an operator to a list of operatorSets.
     *
     *  @param operator The address of the operator to be added to the operator set.
     *  @param operatorSetIds The IDs of the operator sets.
     *  @param operatorSignature The signature of the operator on their intent to register.
     *
     *  @dev msg.sender is used as the AVS.
     *  @dev The operator must not have a pending deregistration from the operator set.
     */
    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override onlyWhenNotPaused(PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS) {
        // Assert operator's signature has not expired.
        require(operatorSignature.expiry >= block.timestamp, SignatureExpired());
        // Assert `operator` is actually an operator.
        require(delegation.isOperator(operator), OperatorNotRegistered());
        // Assert that the AVS is an operator set AVS.
        require(isOperatorSetAVS[msg.sender], InvalidAVS());
        // Assert operator's signature `salt` has not already been spent.
        require(!operatorSaltIsSpent[operator][operatorSignature.salt], SaltSpent());

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

        // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        _registerToOperatorSets(operator, msg.sender, operatorSetIds);
    }

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
    ) external override onlyWhenNotPaused(PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS) {
        if (operatorSignature.signature.length == 0) {
            require(msg.sender == operator, InvalidOperator());
        } else {
            // Assert operator's signature has not expired.
            require(operatorSignature.expiry >= block.timestamp, SignatureExpired());
            // Assert operator's signature `salt` has not already been spent.
            require(!operatorSaltIsSpent[operator][operatorSignature.salt], SaltSpent());

            // Assert that `operatorSignature.signature` is a valid signature for operator set deregistrations.
            EIP1271SignatureUtils.checkSignature_EIP1271(
                operator,
                calculateOperatorSetForceDeregistrationTypehash({
                    avs: avs,
                    operatorSetIds: operatorSetIds,
                    salt: operatorSignature.salt,
                    expiry: operatorSignature.expiry
                }),
                operatorSignature.signature
            );

            // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
            operatorSaltIsSpent[operator][operatorSignature.salt] = true;
        }
        _deregisterFromOperatorSets(avs, operator, operatorSetIds);
    }

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
    ) external override onlyWhenNotPaused(PAUSER_OPERATOR_REGISTER_DEREGISTER_TO_OPERATOR_SETS) {
        _deregisterFromOperatorSets(msg.sender, operator, operatorSetIds);
    }

    /**
     *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     *
     *  @param metadataURI The URI for metadata associated with an AVS.
     *
     *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
     */
    function updateAVSMetadataURI(
        string calldata metadataURI
    ) external override {
        emit AVSMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     *
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(
        bytes32 salt
    ) external override {
        // Mutate `operatorSaltIsSpent` to `true` to prevent future spending.
        operatorSaltIsSpent[msg.sender][salt] = true;
    }

    /**
     *
     *        LEGACY EXTERNAL FUNCTIONS - SUPPORT DEPRECATED IN FUTURE RELEASE AFTER SLASHING RELEASE
     *
     */

    /**
     *  @notice Legacy function called by the AVS's service manager contract
     * to register an operator with the AVS. NOTE: this function will be deprecated in a future release
     * after the slashing release. New AVSs should use `registerOperatorToOperatorSets` instead.
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
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert `operatorSignature.expiry` has not elapsed.
        require(operatorSignature.expiry >= block.timestamp, SignatureExpired());

        // Assert that the AVS is not an operator set AVS.
        require(!isOperatorSetAVS[msg.sender], InvalidAVS());

        // Assert that the `operator` is not actively registered to the AVS.
        require(avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED, InvalidOperator());

        // Assert `operator` has not already spent `operatorSignature.salt`.
        require(!operatorSaltIsSpent[operator][operatorSignature.salt], SaltSpent());

        // Assert `operator` is a registered operator.
        require(delegation.isOperator(operator), OperatorDoesNotExist());

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

        // Set the operator as registered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED);
    }

    /**
     *  @notice Legacy function called by an AVS to deregister an operator from the AVS.
     * NOTE: this function will be deprecated in a future release after the slashing release.
     * New AVSs integrating should use `deregisterOperatorFromOperatorSets` instead.
     *
     *  @param operator The address of the operator to deregister.
     *
     *  @dev Only used by legacy M2 AVSs that have not integrated with operator sets.
     */
    function deregisterOperatorFromAVS(
        address operator
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert that operator is registered for the AVS.
        require(avsOperatorStatus[msg.sender][operator] == OperatorAVSRegistrationStatus.REGISTERED, OperatorNotRegistered());
        // Assert that the AVS is not an operator set AVS.
        require(!isOperatorSetAVS[msg.sender], InvalidAVS());

        // Set the operator as deregistered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Helper function used by migration & registration functions to register an operator to operator sets.
     * @param avs The AVS that the operator is registering to.
     * @param operator The operator to register.
     * @param operatorSetIds The IDs of the operator sets.
     */
    function _registerToOperatorSets(address operator, address avs, uint32[] calldata operatorSetIds) internal {
        // Loop over `operatorSetIds` array and register `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            OperatorSet memory operatorSet = OperatorSet(avs, operatorSetIds[i]);

            require(isOperatorSet[avs][operatorSetIds[i]], InvalidOperatorSet());

            bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);

            require(_operatorSetsMemberOf[operator].add(encodedOperatorSet), InvalidOperator());

            _operatorSetMembers[encodedOperatorSet].add(operator);

            OperatorSetRegistrationStatus storage registrationStatus =
                operatorSetStatus[avs][operator][operatorSetIds[i]];

            require(!registrationStatus.registered, InvalidOperator());

            registrationStatus.registered = true;

            emit OperatorAddedToOperatorSet(operator, operatorSet);
        }
    }

    /**
     * @notice Internal function to deregister an operator from an operator set.
     *
     * @param avs The AVS that the operator is deregistering from.
     * @param operator The operator to deregister.
     * @param operatorSetIds The IDs of the operator sets.
     */
    function _deregisterFromOperatorSets(address avs, address operator, uint32[] calldata operatorSetIds) internal {
        // Loop over `operatorSetIds` array and deregister `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            OperatorSet memory operatorSet = OperatorSet(avs, operatorSetIds[i]);

            bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);

            require(_operatorSetsMemberOf[operator].remove(encodedOperatorSet), InvalidOperator());

            _operatorSetMembers[encodedOperatorSet].remove(operator);

            emit OperatorRemovedFromOperatorSet(operator, operatorSet);
        }
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Returns operatorSet an operator is registered to in the order they were registered.
     * @param operator The operator address to query.
     * @param index The index of the enumerated list of operator sets.
     */
    function operatorSetsMemberOfAtIndex(address operator, uint256 index) external view returns (OperatorSet memory) {
        return _decodeOperatorSet(_operatorSetsMemberOf[operator].at(index));
    }

    /**
     * @notice Returns the operator registered to an operatorSet in the order that it was registered.
     * @param operatorSet The operatorSet to query.
     * @param index The index of the enumerated list of operators.
     */
    function operatorSetMemberAtIndex(OperatorSet memory operatorSet, uint256 index) external view returns (address) {
        return _operatorSetMembers[_encodeOperatorSet(operatorSet)].at(index);
    }

    /**
     * @notice Returns an array of operator sets an operator is registered to.
     * @param operator The operator address to query.
     * @param start The starting index of the array to query.
     *  @param length The amount of items of the array to return.
     */
    function getOperatorSetsOfOperator(
        address operator,
        uint256 start,
        uint256 length
    ) public view returns (OperatorSet[] memory operatorSets) {
        uint256 maxLength = _operatorSetsMemberOf[operator].length() - start;
        if (length > maxLength) length = maxLength;
        operatorSets = new OperatorSet[](length);
        for (uint256 i; i < length; ++i) {
            operatorSets[i] = _decodeOperatorSet(_operatorSetsMemberOf[operator].at(start + i));
        }
    }

    /**
     * @notice Returns an array of operators registered to the operatorSet.
     * @param operatorSet The operatorSet to query.
     * @param start The starting index of the array to query.
     * @param length The amount of items of the array to return.
     */
    function getOperatorsInOperatorSet(
        OperatorSet memory operatorSet,
        uint256 start,
        uint256 length
    ) external view returns (address[] memory operators) {
        bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);
        uint256 maxLength = _operatorSetMembers[encodedOperatorSet].length() - start;
        if (length > maxLength) length = maxLength;
        operators = new address[](length);
        for (uint256 i; i < length; ++i) {
            operators[i] = _operatorSetMembers[encodedOperatorSet].at(start + i);
        }
    }

    /**
     * @notice Returns the number of operators registered to an operatorSet.
     * @param operatorSet The operatorSet to get the member count for
     */
    function getNumOperatorsInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetMembers[_encodeOperatorSet(operatorSet)].length();
    }

    /**
     *  @notice Returns the total number of operator sets an operator is registered to.
     *  @param operator The operator address to query.
     */
    function inTotalOperatorSets(
        address operator
    ) external view returns (uint256) {
        return _operatorSetsMemberOf[operator].length();
    }

    /**
     * @notice Returns whether or not an operator is registered to an operator set.
     * @param operator The operator address to query.
     *  @param operatorSet The `OperatorSet` to query.
     */
    function isMember(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        return _operatorSetsMemberOf[operator].contains(_encodeOperatorSet(operatorSet));
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
    ) public view override returns (bytes32) {
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
    ) public view override returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(OPERATOR_SET_REGISTRATION_TYPEHASH, avs, operatorSetIds, salt, expiry))
        );
    }

    /**
     * @notice Calculates the digest hash to be signed by an operator to force deregister from an operator set.
     *
     * @param avs The AVS that operator is deregistering from.
     * @param operatorSetIds An array of operator set IDs the operator is deregistering from.
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid.
     */
    function calculateOperatorSetForceDeregistrationTypehash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH, avs, operatorSetIds, salt, expiry))
        );
    }

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() public view override returns (bytes32) {
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

    /// @notice Returns an EIP-712 encoded hash struct.
    function _calculateDigestHash(
        bytes32 structHash
    ) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _calculateDomainSeparator(), structHash));
    }

    /// @dev Returns an `OperatorSet` encoded into a 32-byte value.
    /// @param operatorSet The `OperatorSet` to encode.
    function _encodeOperatorSet(
        OperatorSet memory operatorSet
    ) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(operatorSet.avs, uint96(operatorSet.operatorSetId)));
    }

    /// @dev Returns an `OperatorSet` decoded from an encoded 32-byte value.
    /// @param encoded The encoded `OperatorSet` to decode.
    /// @dev Assumes `encoded` is encoded via `_encodeOperatorSet(operatorSet)`.
    function _decodeOperatorSet(
        bytes32 encoded
    ) internal pure returns (OperatorSet memory) {
        return OperatorSet({
            avs: address(uint160(uint256(encoded) >> 96)),
            operatorSetId: uint32(uint256(encoded) & type(uint96).max)
        });
    }
}
