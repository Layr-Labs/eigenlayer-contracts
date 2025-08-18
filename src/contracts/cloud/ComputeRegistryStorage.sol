// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IComputeRegistry.sol";
import "../interfaces/IReleaseManager.sol";
import "../interfaces/IAllocationManager.sol";
import "../interfaces/IKeyRegistrar.sol";
import "../interfaces/ICrossChainRegistry.sol";
import "../libraries/OperatorSetLib.sol";

abstract contract ComputeRegistryStorage is IComputeRegistry {
    // Constants and Immutables

    /// @notice EIP-712 Type Hash for TOS Agreement
    bytes32 public constant TOS_AGREEMENT_TYPEHASH =
        keccak256("TOSAgreement(bytes32 tosHash,address avs,uint32 operatorSetId,address signer,uint256 expiry)");

    /// @notice Maximum expiry value for signatures (effectively never expires)
    uint256 public constant MAX_EXPIRY = type(uint256).max;

    /// @notice The ReleaseManager contract
    IReleaseManager public immutable RELEASE_MANAGER;

    /// @notice The AllocationManager contract
    IAllocationManager public immutable ALLOCATION_MANAGER;

    /// @notice The KeyRegistrar contract
    IKeyRegistrar public immutable KEY_REGISTRAR;

    /// @notice The CrossChainRegistry contract
    ICrossChainRegistry public immutable CROSS_CHAIN_REGISTRY;

    /// @notice The hash of the Terms of Service that AVS operators must sign
    bytes32 public immutable TOS_HASH;

    // Storage

    /// @notice Mapping to track if an operator set is registered for compute
    /// @dev operatorSetKey => isRegistered
    mapping(bytes32 operatorSetKey => bool isRegistered) public isOperatorSetRegistered;

    /// @notice Mapping to store the Terms of Service signature for each registered operator set
    /// @dev operatorSetKey => tosSignature
    mapping(bytes32 operatorSetKey => TOSSignature tosSignature) internal _operatorSetTosSignature;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;

    constructor(
        IReleaseManager _releaseManager,
        IAllocationManager _allocationManager,
        IKeyRegistrar _keyRegistrar,
        ICrossChainRegistry _crossChainRegistry,
        bytes32 _tosHash
    ) {
        RELEASE_MANAGER = _releaseManager;
        ALLOCATION_MANAGER = _allocationManager;
        KEY_REGISTRAR = _keyRegistrar;
        CROSS_CHAIN_REGISTRY = _crossChainRegistry;
        TOS_HASH = _tosHash;
    }
}
