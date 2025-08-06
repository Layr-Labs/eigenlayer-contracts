// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

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
contract ComputeRegistry is
    Initializable,
    OwnableUpgradeable,
    ComputeRegistryStorage,
    PermissionControllerMixin,
    SignatureUtilsMixin
{
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
     * @param _version The semantic version of the contract
     */
    constructor(
        IReleaseManager _releaseManager,
        IAllocationManager _allocationManager,
        IPermissionController _permissionController,
        string memory _version
    )
        ComputeRegistryStorage(_releaseManager, _allocationManager)
        PermissionControllerMixin(_permissionController)
        SignatureUtilsMixin(_version)
    {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract
     * @param _owner The owner of the contract
     * @param _tosHash The hash of the Terms of Service that AVS operators must sign
     */
    function initialize(address _owner, bytes32 _tosHash) external initializer {
        __Ownable_init();
        transferOwnership(_owner);
        _setTosHash(_tosHash);
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
            TOSSignature({signer: msg.sender, tosHash: tosHash, signature: signature});

        emit OperatorSetRegistered(operatorSet, msg.sender, tosHash, signature);
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
     * @notice Updates the Terms of Service hash
     * @param _tosHash The new Terms of Service hash
     * @dev Only callable by the contract owner
     */
    function setTosHash(
        bytes32 _tosHash
    ) external onlyOwner {
        _setTosHash(_tosHash);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @dev Internal function to set the Terms of Service hash
     * @param _tosHash The new Terms of Service hash
     */
    function _setTosHash(
        bytes32 _tosHash
    ) internal {
        tosHash = _tosHash;
        emit TosHashSet(_tosHash);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
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
                    tosHash,
                    operatorSet.avs,
                    operatorSet.id,
                    signer,
                    MAX_EXPIRY
                )
            )
        );
    }
}
