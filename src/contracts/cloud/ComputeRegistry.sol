// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SemVerMixin.sol";
import "./ComputeRegistryStorage.sol";

/**
 * @title ComputeRegistry
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This contract handles permissionless (de)registration of AVS operator sets to the EigenCompute Operator.
 * It enables AVSs to easily access managed operator infrastructure as part of EigenCloud for quick bootstrapping.
 */
contract ComputeRegistry is Initializable, ComputeRegistryStorage, PermissionControllerMixin, SemVerMixin {
    using OperatorSetLib for OperatorSet;
    using ECDSA for bytes32;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the contract with immutable values
     * @param _releaseManager The ReleaseManager contract address
     * @param _permissionController The PermissionController contract address
     * @param _version The semantic version of the contract
     */
    constructor(
        IReleaseManager _releaseManager,
        IPermissionController _permissionController,
        string memory _version
    ) ComputeRegistryStorage(_releaseManager) PermissionControllerMixin(_permissionController) SemVerMixin(_version) {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract
     * @param _tos The Terms of Service string that AVS operators must sign
     */
    function initialize(
        string memory _tos
    ) external initializer {
        TOS = _tos;
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /**
     * @inheritdoc IComputeRegistry
     */
    function registerForCompute(
        OperatorSet calldata operatorSet,
        bytes calldata tosSignature
    ) external checkCanCall(operatorSet.avs) {
        // Check if there is at least one release for the operator set
        // The ReleaseManager will revert with `NoReleases()`if there are no releases for the operator set
        releaseManager.getLatestRelease(operatorSet);

        // Verify the TOS signature
        bytes32 tosHash = keccak256(bytes(TOS));
        address signer = tosHash.toEthSignedMessageHash().recover(tosSignature);

        require(signer == msg.sender, InvalidTOSSignature());

        // Check if already registered
        bytes32 operatorSetKey = operatorSet.key();
        require(!isOperatorSetRegistered[operatorSetKey], OperatorSetAlreadyRegistered());

        // Register the operator set
        isOperatorSetRegistered[operatorSetKey] = true;
        operatorSetTosSignature[operatorSetKey] = tosSignature;

        emit OperatorSetRegistered(operatorSet, tosSignature);
    }

    /**
     * @inheritdoc IComputeRegistry
     */
    function deregisterFromCompute(
        OperatorSet calldata operatorSet
    ) external checkCanCall(operatorSet.avs) {
        bytes32 operatorSetKey = operatorSet.key();
        require(isOperatorSetRegistered[operatorSetKey], OperatorSetNotRegistered());

        // Deregister the operator set
        isOperatorSetRegistered[operatorSetKey] = false;

        emit OperatorSetDeregistered(operatorSet);
    }
}
