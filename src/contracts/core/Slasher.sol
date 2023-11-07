// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../libraries/EIP1271SignatureUtils.sol";
import "./SlasherStorage.sol";

/**
 * @title The primary 'slashing' contract for EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This contract allows AVSs and operators to emit events to be viewed by offchain services before slashing is enabled.
 */
contract Slasher is SlasherStorage {

    constructor(IStrategyManager _strategyManager, IDelegationManager _delegation) SlasherStorage(_strategyManager, _delegation, "v0.0.1") {
    }

    // EXTERNAL FUNCTIONS
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
    }

    /**
     * @notice External function called by AVSs to register an operator without the ability to slash it
     * @param operator The address of the registering operator
     * @param signatureWithSaltAndExpiry The signature, salt, and expiry of the operator's signature
     */
    function registerOperatorNoSlashingWithAVS(address operator, ISignatureUtils.SignatureWithSaltAndExpiry memory signatureWithSaltAndExpiry) external {
        require(signatureWithSaltAndExpiry.expiry > block.timestamp, "Slasher.registerOperatorNoSlashingWithAVS: signature expired");
        require(isOperatorSaltUsed[msg.sender][signatureWithSaltAndExpiry.salt] == false, "Slasher.registerOperatorNoSlashingWithAVS: salt already used");
        require(!registeredWithoutSlashing[msg.sender][operator], "Slasher.registerOperatorNoSlashingWithAVS: operator already registered without slashing");
        // check operator signature
        EIP1271SignatureUtils.checkSignature_EIP1271(
            operator, 
            calculateRegistrationNoSlashingDigestHash({
                operator: operator, 
                avs: msg.sender, 
                salt: signatureWithSaltAndExpiry.salt, 
                expiry: signatureWithSaltAndExpiry.expiry
            }), 
            signatureWithSaltAndExpiry.signature
        );
        isOperatorSaltUsed[operator][signatureWithSaltAndExpiry.salt] = true;
        // register operator
        registeredWithoutSlashing[msg.sender][operator] = true;

        emit OperatorRegistrationStatusUpdate(operator, msg.sender, RegistrationStatus.REGISTERED);
    }

    /**
     * @notice Public function called by AVSs to deregister an operator once the operator is already registered without slashing
     * @param operator The address of the registering operator
     */
    function deregisterOperatorNoSlashingFromAVS(address operator) external {
        require(registeredWithoutSlashing[msg.sender][operator], "Slasher.deregisterOperatorNoSlashingFromAVS: operator is not registered without slashing");
        // register operator
        registeredWithoutSlashing[msg.sender][operator] = false;

        emit OperatorRegistrationStatusUpdate(operator, msg.sender, RegistrationStatus.DEREGISTERED);
    }

    
    /**
     * @notice Public function for the the operator message digest to sign for registration without slashing
     * @param operator The address of the registering operator 
     * @param avs The avs the operator is registering with
     * @param salt The salt to use for the operator's signature
     * @param expiry The desired expiry time of the operator's signature
     */
    function calculateRegistrationNoSlashingDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the digest hash
        return _hashTypedDataV4(keccak256(abi.encode(REGISTRATION_WITHOUT_SLASHING_TYPEHASH, operator, avs, salt, expiry)));
    }
}
