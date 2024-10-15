// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../mixins/SignatureUtils.sol";
import "../permissions/Pausable.sol";
import "./AVSDirectoryStorage.sol";

contract AVSDirectory is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AVSDirectoryStorage,
    ReentrancyGuardUpgradeable,
    SignatureUtils
{
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManager,
     * and eigenpodManager contracts
     */
    constructor(
        IDelegationManager _delegation,
        uint32 _DEALLOCATION_DELAY
    ) AVSDirectoryStorage(_delegation, _DEALLOCATION_DELAY) {
        _disableInitializers();
    }

    /// @inheritdoc IAVSDirectory
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
    }

    /**
     *
     *                    EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc IAVSDirectory
    function createOperatorSets(
        uint32[] calldata operatorSetIds
    ) external {
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            require(!isOperatorSet[msg.sender][operatorSetIds[i]], InvalidOperatorSet());
            isOperatorSet[msg.sender][operatorSetIds[i]] = true;
            emit OperatorSetCreated(OperatorSet({avs: msg.sender, operatorSetId: operatorSetIds[i]}));
        }
    }

    /// @inheritdoc IAVSDirectory
    function becomeOperatorSetAVS() external {
        require(!isOperatorSetAVS[msg.sender], InvalidAVS());
        isOperatorSetAVS[msg.sender] = true;
        emit AVSMigratedToOperatorSets(msg.sender);
    }

    /// @inheritdoc IAVSDirectory
    function migrateOperatorsToOperatorSets(
        address[] calldata operators,
        uint32[][] calldata operatorSetIds
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
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

    /// @inheritdoc IAVSDirectory
    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
        // Assert operator's signature has not expired.
        require(operatorSignature.expiry >= block.timestamp, SignatureExpired());
        // Assert `operator` is actually an operator.
        require(delegation.isOperator(operator), OperatorNotRegisteredToEigenLayer());
        // Assert that the AVS is an operator set AVS.
        require(isOperatorSetAVS[msg.sender], InvalidAVS());
        // Assert operator's signature `salt` has not already been spent.
        require(!operatorSaltIsSpent[operator][operatorSignature.salt], SaltSpent());

        // Assert that `operatorSignature.signature` is a valid signature for operator set registrations.
        _checkIsValidSignatureNow({
            signer: operator,
            signableDigest: calculateOperatorSetRegistrationDigestHash({
                avs: msg.sender,
                operatorSetIds: operatorSetIds,
                salt: operatorSignature.salt,
                expiry: operatorSignature.expiry
            }),
            signature: operatorSignature.signature
        });

        // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        _registerToOperatorSets(operator, msg.sender, operatorSetIds);
    }

    /// @inheritdoc IAVSDirectory
    function forceDeregisterFromOperatorSets(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
        if (operatorSignature.signature.length == 0) {
            require(msg.sender == operator, InvalidOperator());
        } else {
            // Assert operator's signature has not expired.
            require(operatorSignature.expiry >= block.timestamp, SignatureExpired());
            // Assert operator's signature `salt` has not already been spent.
            require(!operatorSaltIsSpent[operator][operatorSignature.salt], SaltSpent());

            // Assert that `operatorSignature.signature` is a valid signature for operator set deregistrations.
            _checkIsValidSignatureNow({
                signer: operator,
                signableDigest: calculateOperatorSetForceDeregistrationTypehash({
                    avs: avs,
                    operatorSetIds: operatorSetIds,
                    salt: operatorSignature.salt,
                    expiry: operatorSignature.expiry
                }),
                signature: operatorSignature.signature
            });

            // Mutate `operatorSaltIsSpent` to `true` to prevent future respending.
            operatorSaltIsSpent[operator][operatorSignature.salt] = true;
        }
        _deregisterFromOperatorSets(avs, operator, operatorSetIds);
    }

    /// @inheritdoc IAVSDirectory
    function deregisterOperatorFromOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
        _deregisterFromOperatorSets(msg.sender, operator, operatorSetIds);
    }

    /// @inheritdoc IAVSDirectory
    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] calldata strategies) external override {
        OperatorSet memory operatorSet = OperatorSet(msg.sender, operatorSetId);
        require(isOperatorSet[msg.sender][operatorSetId], InvalidOperatorSet());
        bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);
        for (uint256 i = 0; i < strategies.length; i++) {
            require(
                _operatorSetStrategies[encodedOperatorSet].add(address(strategies[i])), StrategyAlreadyInOperatorSet()
            );
            emit StrategyAddedToOperatorSet(operatorSet, strategies[i]);
        }
    }

    /// @inheritdoc IAVSDirectory
    function removeStrategiesFromOperatorSet(uint32 operatorSetId, IStrategy[] calldata strategies) external override {
        OperatorSet memory operatorSet = OperatorSet(msg.sender, operatorSetId);
        require(isOperatorSet[msg.sender][operatorSetId], InvalidOperatorSet());
        bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);
        for (uint256 i = 0; i < strategies.length; i++) {
            require(
                _operatorSetStrategies[encodedOperatorSet].remove(address(strategies[i])), StrategyNotInOperatorSet()
            );
            emit StrategyRemovedFromOperatorSet(operatorSet, strategies[i]);
        }
    }

    /// @inheritdoc IAVSDirectory
    function updateAVSMetadataURI(
        string calldata metadataURI
    ) external override {
        emit AVSMetadataURIUpdated(msg.sender, metadataURI);
    }

    /// @inheritdoc IAVSDirectory
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

    /// @inheritdoc IAVSDirectory
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert `operatorSignature.expiry` has not elapsed.
        require(operatorSignature.expiry >= block.timestamp, SignatureExpired());

        // Assert that the AVS is not an operator set AVS.
        require(!isOperatorSetAVS[msg.sender], InvalidAVS());

        // Assert that the `operator` is not actively registered to the AVS.
        require(
            avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED,
            OperatorAlreadyRegisteredToAVS()
        );

        // Assert `operator` has not already spent `operatorSignature.salt`.
        require(!operatorSaltIsSpent[operator][operatorSignature.salt], SaltSpent());

        // Assert `operator` is a registered operator.
        require(delegation.isOperator(operator), OperatorNotRegisteredToEigenLayer());

        // Assert that `operatorSignature.signature` is a valid signature for the operator AVS registration.
        _checkIsValidSignatureNow({
            signer: operator,
            signableDigest: calculateOperatorAVSRegistrationDigestHash({
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

    /// @inheritdoc IAVSDirectory
    function deregisterOperatorFromAVS(
        address operator
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        // Assert that operator is registered for the AVS.
        require(
            avsOperatorStatus[msg.sender][operator] == OperatorAVSRegistrationStatus.REGISTERED,
            OperatorNotRegisteredToAVS()
        );
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

            _operatorSetsMemberOf[operator].add(encodedOperatorSet);

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

            _operatorSetsMemberOf[operator].remove(encodedOperatorSet);

            _operatorSetMembers[encodedOperatorSet].remove(operator);

            OperatorSetRegistrationStatus storage registrationStatus =
                operatorSetStatus[avs][operator][operatorSetIds[i]];

            require(registrationStatus.registered, InvalidOperator());

            registrationStatus.registered = false;
            registrationStatus.lastDeregisteredTimestamp = uint32(block.timestamp);

            emit OperatorRemovedFromOperatorSet(operator, operatorSet);
        }
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IAVSDirectory
    function operatorSetsMemberOfAtIndex(address operator, uint256 index) external view returns (OperatorSet memory) {
        return _decodeOperatorSet(_operatorSetsMemberOf[operator].at(index));
    }

    /// @inheritdoc IAVSDirectory
    function operatorSetMemberAtIndex(OperatorSet memory operatorSet, uint256 index) external view returns (address) {
        return _operatorSetMembers[_encodeOperatorSet(operatorSet)].at(index);
    }

    /// @inheritdoc IAVSDirectory
    function getNumOperatorSetsOfOperator(
        address operator
    ) external view returns (uint256) {
        return _operatorSetsMemberOf[operator].length();
    }

    /// @inheritdoc IAVSDirectory
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

    /// @inheritdoc IAVSDirectory
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

    /// @inheritdoc IAVSDirectory
    function getStrategiesInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (IStrategy[] memory strategies) {
        bytes32 encodedOperatorSet = _encodeOperatorSet(operatorSet);
        uint256 length = _operatorSetStrategies[encodedOperatorSet].length();

        strategies = new IStrategy[](length);
        for (uint256 i; i < length; ++i) {
            strategies[i] = IStrategy(_operatorSetStrategies[encodedOperatorSet].at(i));
        }
    }

    /// @inheritdoc IAVSDirectory
    function getNumOperatorsInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetMembers[_encodeOperatorSet(operatorSet)].length();
    }

    /// @inheritdoc IAVSDirectory
    function inTotalOperatorSets(
        address operator
    ) external view returns (uint256) {
        return _operatorSetsMemberOf[operator].length();
    }

    /// @inheritdoc IAVSDirectory
    function isMember(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        return _operatorSetsMemberOf[operator].contains(_encodeOperatorSet(operatorSet));
    }

    /// @inheritdoc IAVSDirectory
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        if (isMember(operator, operatorSet)) return true;

        OperatorSetRegistrationStatus memory status =
            operatorSetStatus[operatorSet.avs][operator][operatorSet.operatorSetId];

        return block.timestamp < status.lastDeregisteredTimestamp + DEALLOCATION_DELAY;
    }

    /// @inheritdoc IAVSDirectory
    function isOperatorSetBatch(
        OperatorSet[] calldata operatorSets
    ) public view returns (bool) {
        for (uint256 i = 0; i < operatorSets.length; ++i) {
            if (!isOperatorSet[operatorSets[i].avs][operatorSets[i].operatorSetId]) return false;
        }
        return true;
    }

    /// @inheritdoc IAVSDirectory
    function calculateOperatorAVSRegistrationDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) public view override returns (bytes32) {
        return _calculateSignableDigest(
            keccak256(abi.encode(OPERATOR_AVS_REGISTRATION_TYPEHASH, operator, avs, salt, expiry))
        );
    }

    /// @inheritdoc IAVSDirectory
    function calculateOperatorSetRegistrationDigestHash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) public view override returns (bytes32) {
        return _calculateSignableDigest(
            keccak256(abi.encode(OPERATOR_SET_REGISTRATION_TYPEHASH, avs, operatorSetIds, salt, expiry))
        );
    }

    /// @inheritdoc IAVSDirectory
    function calculateOperatorSetForceDeregistrationTypehash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateSignableDigest(
            keccak256(abi.encode(OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH, avs, operatorSetIds, salt, expiry))
        );
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
