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
     * @param expiry The desired expiry time of the operator's signature
     * @param signature The signature of the operator
     */
    function registerOperatorWithoutSlashingWithAVS(address operator, uint256 expiry, bytes memory signature) external {
        require(expiry > block.timestamp, "Slasher.registerOperatorWithoutSlashingWithAVS: signature expired");
        require(!registeredWithoutSlashing[msg.sender][operator], "Slasher.registerOperatorWithoutSlashingWithAVS: operator already registered without slashing");
        // check operator signature
        EIP1271SignatureUtils.checkSignature_EIP1271(operator, calculateRegistrationWithoutSlashingDigestHash(operator, msg.sender, nonces[msg.sender]++, expiry), signature);
        // register operator
        registeredWithoutSlashing[msg.sender][operator] = true;
    }

    /**
     * @notice Public function called by AVSs to deregister an operator once the operator is already registered without slashing
     * @param operator The address of the registering operator
     */
    function deregisterOperatorWithoutSlashingWithAVS(address operator) external {
        require(registeredWithoutSlashing[msg.sender][operator], "Slasher.deregisterOperatorWithoutSlashingWithAVS: operator is not registered without slashing");
        // register operator
        registeredWithoutSlashing[msg.sender][operator] = false;
    }

    
    /**
     * @notice Public function for the the operator message digest to sign for registration without slashing
     * @param operator The address of the registering operator 
     * @param avs The avs the operator is registering with
     * @param nonce The nonce to use for the operator's signature
     * @param expiry The desired expiry time of the operator's signature
     */
    function calculateRegistrationWithoutSlashingDigestHash(
        address operator,
        address avs,
        uint256 nonce,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the digest hash
        return _hashTypedDataV4(keccak256(abi.encode(REGISTRATION_WITHOUT_SLASHING_TYPEHASH, operator, avs, nonce, expiry)));
    }
}
