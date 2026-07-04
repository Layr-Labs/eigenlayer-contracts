// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import {IReleaseManager} from "./IReleaseManager.sol";
import {IAllocationManager} from "./IAllocationManager.sol";
import {IKeyRegistrar} from "./IKeyRegistrar.sol";
import {ICrossChainRegistry} from "./ICrossChainRegistry.sol";

interface IComputeRegistryTypes {
    /// @notice The Terms of Service signature
    struct TOSSignature {
        address signer;
        bytes32 tosHash;
        bytes signature;
    }
}

interface IComputeRegistryErrors {
    /// @notice Error thrown when the provided signature does not match the expected Terms of Service signature
    /// @dev Error code: 0x04bf729c
    error InvalidTOSSignature();

    /// @notice Error thrown when an operator set is already registered for compute
    /// @dev Error code: 0x1503562a
    error OperatorSetAlreadyRegistered();

    /// @notice Error thrown when an operator set is not registered but expected to be
    /// @dev Error code: 0x3a2e3ac6
    error OperatorSetNotRegistered();

    /// @notice Error thrown when an invalid operator set is provided
    /// @dev Error code: 0x7ec5c154
    error InvalidOperatorSet();

    /// @notice Error thrown when the curve type for an operator set has not been set
    /// @dev Error code: 0x3104b8e7
    error CurveTypeNotSet();

    /// @notice Error thrown when the operator set does not have an active generation reservation
    /// @dev Error code: 0xd0147d2d
    error NoActiveGenerationReservation();
}

interface IComputeRegistryEvents {
    /// @notice Emitted when an operator set is registered for compute
    /// @param operatorSet The operator set that was registered
    /// @param signer The address that signed the Terms of Service
    /// @param tosHash The hash of the Terms of Service
    /// @param signature The signature of the Terms of Service
    event OperatorSetRegistered(
        OperatorSet indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature
    );

    /// @notice Emitted when an operator set is deregistered from compute
    /// @param operatorSet The operator set that was deregistered
    event OperatorSetDeregistered(OperatorSet indexed operatorSet);
}

interface IComputeRegistry is IComputeRegistryErrors, IComputeRegistryEvents, IComputeRegistryTypes {
    /**
     *
     *                         WRITE FUNCTIONS
     *
     */

    /**
     * @notice Registers an operator set for compute services
     * @param operatorSet The operator set to register
     * @param signature The EIP-712 signature of the Terms of Service agreement
     * @dev The caller must have permission to call on behalf of the operatorSet.avs through the PermissionController
     * @dev The signature must be a valid EIP-712 signature of the Terms of Service with expiry set to MAX_EXPIRY
     * @dev Reverts for:
     *      - InvalidPermissions: Caller does not have permission to call on behalf of operatorSet.avs
     *      - InvalidOperatorSet: The operator set does not exist in the AllocationManager
     *      - OperatorSetAlreadyRegistered: The operator set is already registered for compute
     *      - CurveTypeNotSet: The operator set has not configured a curve type in the KeyRegistrar
     *      - NoActiveGenerationReservation: The operator set does not have an active generation reservation in the CrossChainRegistry
     *      - NoReleases: The operator set does not have any releases in the ReleaseManager
     *      - InvalidSignature: The provided signature is invalid or does not match the expected signer
     * @dev Emits the following events:
     *      - OperatorSetRegistered: When the operator set is successfully registered with the TOS signature
     */
    function registerForCompute(OperatorSet memory operatorSet, bytes memory signature) external;

    /**
     * @notice Deregisters an operator set from compute services
     * @param operatorSet The operator set to deregister
     * @dev The caller must have permission to call on behalf of the operatorSet.avs through the PermissionController
     * @dev Reverts for:
     *      - InvalidPermissions: Caller does not have permission to call on behalf of operatorSet.avs
     *      - InvalidOperatorSet: The operator set does not exist in the AllocationManager
     *      - OperatorSetNotRegistered: The operator set is not registered for compute
     * @dev Emits the following events:
     *      - OperatorSetDeregistered: When the operator set is successfully deregistered
     */
    function deregisterFromCompute(
        OperatorSet memory operatorSet
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Returns the EIP-712 type hash used for TOS agreements
     * @return The TOS_AGREEMENT_TYPEHASH constant used in EIP-712 signature verification
     * @dev This type hash is: keccak256("TOSAgreement(bytes32 tosHash,address avs,uint32 operatorSetId,address signer,uint256 expiry)")
     */
    function TOS_AGREEMENT_TYPEHASH() external view returns (bytes32);

    /**
     * @notice Returns the maximum expiry value used for signatures
     * @return The MAX_EXPIRY constant (type(uint256).max)
     * @dev Signatures use MAX_EXPIRY to indicate they never expire
     */
    function MAX_EXPIRY() external view returns (uint256);

    /**
     * @notice Returns the ReleaseManager contract
     * @return The ReleaseManager contract address
     * @dev Used to verify operator sets have at least one release before registration
     */
    function RELEASE_MANAGER() external view returns (IReleaseManager);

    /**
     * @notice Returns the AllocationManager contract
     * @return The AllocationManager contract address
     * @dev Used to verify operator sets exist and are valid
     */
    function ALLOCATION_MANAGER() external view returns (IAllocationManager);

    /**
     * @notice Returns the KeyRegistrar contract
     * @return The KeyRegistrar contract address
     * @dev Used to verify operator sets have configured curve types
     */
    function KEY_REGISTRAR() external view returns (IKeyRegistrar);

    /**
     * @notice Returns the CrossChainRegistry contract
     * @return The CrossChainRegistry contract address
     * @dev Used to verify operator sets have active generation reservations
     */
    function CROSS_CHAIN_REGISTRY() external view returns (ICrossChainRegistry);

    /**
     * @notice Returns the hash of the Terms of Service
     * @return The immutable hash of the Terms of Service that must be signed
     * @dev This hash is set at contract deployment and cannot be changed
     */
    function TOS_HASH() external view returns (bytes32);

    /**
     * @notice Checks if an operator set is registered for compute
     * @param operatorSetKey The key of the operator set to check
     * @return True if the operator set is registered, false otherwise
     */
    function isOperatorSetRegistered(
        bytes32 operatorSetKey
    ) external view returns (bool);

    /**
     * @notice Returns the Terms of Service signature for an operator set
     * @param operatorSet The operator set to query
     * @return The Terms of Service signature including signer, TOS hash, and signature bytes
     * @dev Returns an empty struct if the operator set is not registered
     */
    function getOperatorSetTosSignature(
        OperatorSet memory operatorSet
    ) external view returns (TOSSignature memory);

    /**
     * @notice Calculates the EIP-712 digest hash that should be signed
     * @param operatorSet The operator set that is agreeing to the TOS
     * @param signer The address that is signing the agreement
     * @return The EIP-712 digest hash ready for signing
     * @dev The digest includes the TOS hash, operator set details, signer, and MAX_EXPIRY
     * @dev This digest should be signed according to EIP-712 standards
     */
    function calculateTOSAgreementDigest(
        OperatorSet memory operatorSet,
        address signer
    ) external view returns (bytes32);
}
