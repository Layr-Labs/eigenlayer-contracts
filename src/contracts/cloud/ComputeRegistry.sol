// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SignatureUtilsMixin.sol";
import "./ComputeRegistryStorage.sol";

/**
 * @title ComputeRegistry
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This contract handles permissionless (de)registration of AVS operator sets to the EigenCompute Operator.
 * It enables AVSs to easily access managed operator infrastructure as part of EigenCloud for quick bootstrapping.
 */
contract ComputeRegistry is Initializable, ComputeRegistryStorage, PermissionControllerMixin, SignatureUtilsMixin {
    using OperatorSetLib for OperatorSet;

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
    )
        ComputeRegistryStorage(_releaseManager)
        PermissionControllerMixin(_permissionController)
        SignatureUtilsMixin(_version)
    {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract
     * @param _tos The Terms of Service string that AVS operators must sign
     */
    function initialize(
        string memory _tos
    ) external initializer {
        tos = _tos;
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
        OperatorSet memory operatorSet,
        bytes memory tosSignature
    ) external checkCanCall(operatorSet.avs) {
        // Check if there is at least one release for the operator set
        // The ReleaseManager will revert with `NoReleases()` if there are no releases for the operator set
        RELEASE_MANAGER.getLatestRelease(operatorSet);

        // Verify the signature
        _checkIsValidSignatureNow({
            signer: msg.sender,
            signableDigest: calculateTOSAgreementDigest(operatorSet, msg.sender),
            signature: tosSignature,
            expiry: MAX_EXPIRY
        });

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
        OperatorSet memory operatorSet
    ) external checkCanCall(operatorSet.avs) {
        bytes32 operatorSetKey = operatorSet.key();
        require(isOperatorSetRegistered[operatorSetKey], OperatorSetNotRegistered());

        // Deregister the operator set
        isOperatorSetRegistered[operatorSetKey] = false;

        emit OperatorSetDeregistered(operatorSet);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Calculates the EIP-712 struct hash for a tos agreement
     * @param operatorSet The operator set that is agreeing to the tos
     * @param signer The address that is signing the agreement
     * @return The EIP-712 struct hash
     */
    function calculateTOSAgreementHash(OperatorSet memory operatorSet, address signer) public view returns (bytes32) {
        return keccak256(
            abi.encode(
                TOS_AGREEMENT_TYPEHASH, keccak256(bytes(tos)), operatorSet.avs, operatorSet.id, signer, MAX_EXPIRY
            )
        );
    }

    /**
     * @notice Calculates the EIP-712 digest hash that should be signed
     * @param operatorSet The operator set that is agreeing to the tos
     * @param signer The address that is signing the agreement
     * @return The EIP-712 digest hash ready for signing
     */
    function calculateTOSAgreementDigest(
        OperatorSet memory operatorSet,
        address signer
    ) public view returns (bytes32) {
        bytes32 structHash = calculateTOSAgreementHash(operatorSet, signer);
        return _calculateSignableDigest(structHash);
    }
}
