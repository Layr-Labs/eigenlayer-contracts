// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../mixins/SignatureUtilsMixin.sol";
import "../permissions/Pausable.sol";
import "./AVSDirectoryStorage.sol";

contract AVSDirectory is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AVSDirectoryStorage,
    ReentrancyGuardUpgradeable,
    SignatureUtilsMixin
{
    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy manager, delegationManager,
     * and eigenpodManager contracts
     */
    constructor(
        IDelegationManager _delegation,
        IPauserRegistry _pauserRegistry,
        string memory _version
    ) AVSDirectoryStorage(_delegation) Pausable(_pauserRegistry) SignatureUtilsMixin(_version) {
        _disableInitializers();
    }

    /// @inheritdoc IAVSDirectory
    function initialize(address initialOwner, uint256 initialPausedStatus) external initializer {
        _setPausedStatus(initialPausedStatus);
        _transferOwnership(initialOwner);
    }

    /**
     *
     *                    EXTERNAL FUNCTIONS
     *
     */

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
        ISignatureUtilsMixinTypes.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
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
            signature: operatorSignature.signature,
            expiry: operatorSignature.expiry
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

        // Set the operator as deregistered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

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
}
