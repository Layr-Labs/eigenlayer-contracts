// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import {IReleaseManager} from "./IReleaseManager.sol";
import {IAllocationManager} from "./IAllocationManager.sol";

interface IComputeRegistryErrors {
    /// @dev Thrown when the provided signature does not match the expected Terms of Service signature
    error InvalidTOSSignature();

    /// @dev Thrown when an operator set is already registered for compute
    error OperatorSetAlreadyRegistered();

    /// @dev Thrown when an operator set is not registered but expected to be
    error OperatorSetNotRegistered();

    /// @dev Thrown when an invalid operator set is provided
    error InvalidOperatorSet();
}

interface IComputeRegistryEvents {
    /// @notice Emitted when an operator set is registered for compute
    /// @param operatorSet The operator set that was registered
    /// @param tosSignature The signature of the Terms of Service
    event OperatorSetRegistered(OperatorSet indexed operatorSet, bytes tosSignature);

    /// @notice Emitted when an operator set is deregistered from compute
    /// @param operatorSet The operator set that was deregistered
    event OperatorSetDeregistered(OperatorSet indexed operatorSet);
}

interface IComputeRegistry is IComputeRegistryErrors, IComputeRegistryEvents {
    /**
     *
     *                         WRITE FUNCTIONS
     *
     */

    /**
     * @notice Registers an operator set for compute services
     * @param operatorSet The operator set to register
     * @param tosSignature The EIP-712 signature of the Terms of Service
     * @dev Requires the caller to have permission to call on behalf of the operatorSet.avs
     * @dev The operator set must have at least one release available in the ReleaseManager
     * @dev The signature must be a valid EIP-712 signature of the Terms of Service with expiry set to MAX_EXPIRY
     */
    function registerForCompute(OperatorSet calldata operatorSet, bytes memory tosSignature) external;

    /**
     * @notice Deregisters an operator set from compute services
     * @param operatorSet The operator set to deregister
     * @dev Requires the caller to have permission to call on behalf of the operatorSet.avs
     * @dev The operator set must be registered
     */
    function deregisterFromCompute(
        OperatorSet calldata operatorSet
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Returns the EIP-712 type hash used for TOS agreements
     * @return The TOS_AGREEMENT_TYPEHASH constant
     */
    function TOS_AGREEMENT_TYPEHASH() external view returns (bytes32);

    /**
     * @notice Returns the maximum expiry value used for signatures
     * @return The MAX_EXPIRY constant (type(uint256).max)
     */
    function MAX_EXPIRY() external view returns (uint256);

    /**
     * @notice Returns the ReleaseManager contract
     * @return The ReleaseManager contract
     */
    function RELEASE_MANAGER() external view returns (IReleaseManager);

    /**
     * @notice Returns the AllocationManager contract
     * @return The AllocationManager contract
     */
    function ALLOCATION_MANAGER() external view returns (IAllocationManager);

    /**
     * @notice Returns the Terms of Service string
     * @return The Terms of Service that must be signed
     */
    function tos() external view returns (string memory);

    /**
     * @notice Checks if an operator set is registered for compute
     * @param operatorSetKey The key of the operator set to check
     * @return True if the operator set is registered, false otherwise
     */
    function isOperatorSetRegistered(
        bytes32 operatorSetKey
    ) external view returns (bool);

    /**
     * @notice Returns the Terms of Service signature for a registered operator set
     * @param operatorSetKey The key of the operator set to query
     * @return The Terms of Service signature
     */
    function operatorSetTosSignature(
        bytes32 operatorSetKey
    ) external view returns (bytes memory);

    /**
     * @notice Calculates the EIP-712 digest hash that should be signed
     * @param operatorSet The operator set that is agreeing to the TOS
     * @param signer The address that is signing the agreement
     * @return The EIP-712 digest hash ready for signing
     */
    function calculateTOSAgreementDigest(
        OperatorSet memory operatorSet,
        address signer
    ) external view returns (bytes32);
}
