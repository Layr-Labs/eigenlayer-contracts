// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";

interface IComputeRegistryErrors {
    /// @dev Thrown when the provided signature does not match the expected Terms of Service signature
    error InvalidTOSSignature();

    /// @dev Thrown when an operator set is already registered for compute
    error OperatorSetAlreadyRegistered();

    /// @dev Thrown when an operator set is not registered but expected to be
    error OperatorSetNotRegistered();

    /// @dev Thrown when an operator set has no releases available
    error NoReleasesForOperatorSet();
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
     * @param tosSignature The signature of the Terms of Service
     * @dev Requires the caller to have permission to call on behalf of the operatorSet.avs
     * @dev The operator set must have at least one release available
     * @dev The signature must be a valid ECDSA signature of the Terms of Service
     */
    function registerForCompute(OperatorSet calldata operatorSet, bytes calldata tosSignature) external;

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
     * @notice Returns the Terms of Service string
     * @return The Terms of Service that must be signed
     */
    function TOS() external view returns (string memory);

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
}
