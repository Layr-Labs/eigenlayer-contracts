// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IStrategyManager.sol";
import "./IDelegationManager.sol";
import "../interfaces/ISignatureUtils.sol";

/**
 * @title Interface for the primary 'slashing' contract for EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice See the `Slasher` contract itself for implementation details.
 */
interface ISlasher {

    enum RegistrationStatus {
        DEREGISTERED,
        REGISTERED
    }

    // EVENTS

    /// @notice Emitted when an operator's registration status is updated.
    event OperatorRegistrationStatusUpdate(address operator, address avs, RegistrationStatus status);

    // FUNCTIONS

    /// @notice The StrategyManager contract of EigenLayer
    function strategyManager() external view returns (IStrategyManager);

    /// @notice The DelegationManager contract of EigenLayer
    function delegation() external view returns (IDelegationManager);

    /// @notice avs => operator => bool whether the operator is registered without slashing for the avs
    function registeredWithoutSlashing(address avs, address operator) external view returns (bool);

    /**
     * @notice External function called by AVSs to register an operator without the ability to slash it
     * @param operator The address of the registering operator
     * @param signatureWithSaltAndExpiry The signature, salt, and expiry of the operator's signature
     */
    function registerOperatorNoSlashingWithAVS(address operator, ISignatureUtils.SignatureWithSaltAndExpiry memory signatureWithSaltAndExpiry) external;

    /**
     * @notice Public function called by AVSs to deregister an operator once the operator is already registered without slashing
     * @param operator The address of the registering operator
     */
    function deregisterOperatorNoSlashingFromAVS(address operator) external;

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
    ) external view returns (bytes32);
}
