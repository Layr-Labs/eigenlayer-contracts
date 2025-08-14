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
     *                         MODIFIERS
     *
     */

    /**
     * @dev Validates that the operator set exists in the AllocationManager
     * @param operatorSet The operator set to validate
     */
    modifier isValidOperatorSet(
        OperatorSet memory operatorSet
    ) {
        require(ALLOCATION_MANAGER.isOperatorSet(operatorSet), InvalidOperatorSet());
        _;
    }

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the contract with immutable values
     * @param _releaseManager The ReleaseManager contract address
     * @param _allocationManager The AllocationManager contract address
     * @param _permissionController The PermissionController contract address
     * @param _tosHash The immutable hash of the Terms of Service
     * @param _version The semantic version of the contract
     */
    constructor(
        IReleaseManager _releaseManager,
        IAllocationManager _allocationManager,
        IPermissionController _permissionController,
        bytes32 _tosHash,
        string memory _version
    )
        ComputeRegistryStorage(_releaseManager, _allocationManager, _tosHash)
        PermissionControllerMixin(_permissionController)
        SignatureUtilsMixin(_version)
    {
        _disableInitializers();
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
        bytes memory signature
    ) external checkCanCall(operatorSet.avs) isValidOperatorSet(operatorSet) {
        // Check if there is at least one release for the operator set
        // The ReleaseManager will revert with `NoReleases()` if there are no releases for the operator set
        RELEASE_MANAGER.getLatestRelease(operatorSet);

        // Verify the signature
        _checkIsValidSignatureNow({
            signer: msg.sender,
            signableDigest: calculateTOSAgreementDigest(operatorSet, msg.sender),
            signature: signature,
            expiry: MAX_EXPIRY
        });

        // Check if already registered
        bytes32 operatorSetKey = operatorSet.key();
        require(!isOperatorSetRegistered[operatorSetKey], OperatorSetAlreadyRegistered());

        // Register the operator set
        isOperatorSetRegistered[operatorSetKey] = true;
        _operatorSetTosSignature[operatorSetKey] =
            TOSSignature({signer: msg.sender, tosHash: TOS_HASH, signature: signature});

        emit OperatorSetRegistered(operatorSet, msg.sender, TOS_HASH, signature);
    }

    /**
     * @inheritdoc IComputeRegistry
     */
    function deregisterFromCompute(
        OperatorSet memory operatorSet
    ) external checkCanCall(operatorSet.avs) isValidOperatorSet(operatorSet) {
        bytes32 operatorSetKey = operatorSet.key();
        require(isOperatorSetRegistered[operatorSetKey], OperatorSetNotRegistered());

        // Deregister the operator set
        isOperatorSetRegistered[operatorSetKey] = false;

        emit OperatorSetDeregistered(operatorSet);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     */

    /**
     * @inheritdoc IComputeRegistry
     */
    function getOperatorSetTosSignature(
        OperatorSet memory operatorSet
    ) external view returns (TOSSignature memory) {
        return _operatorSetTosSignature[operatorSet.key()];
    }

    /**
     * @notice Calculates the EIP-712 digest hash that should be signed
     * @param operatorSet The operator set that is agreeing to the TOS
     * @param signer The address that is signing the agreement
     * @return The EIP-712 digest hash ready for signing
     */
    function calculateTOSAgreementDigest(
        OperatorSet memory operatorSet,
        address signer
    ) public view returns (bytes32) {
        /// forgefmt: disable-next-item
        return _calculateSignableDigest(
            keccak256(
                abi.encode(
                    TOS_AGREEMENT_TYPEHASH,
                    TOS_HASH,
                    operatorSet.avs,
                    operatorSet.id,
                    signer,
                    MAX_EXPIRY
                )
            )
        );
    }
}
