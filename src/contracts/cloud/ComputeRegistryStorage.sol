// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IComputeRegistry.sol";
import "../interfaces/IReleaseManager.sol";
import "../interfaces/IAllocationManager.sol";
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

    // Storage

    /// @notice The hash of the Terms of Service that AVS operators must sign
    bytes32 public tosHash;

    /// @notice Mapping to track if an operator set is registered for compute
    /// @dev operatorSetKey => isRegistered
    mapping(bytes32 => bool) public isOperatorSetRegistered;

    /// @notice Mapping to store the Terms of Service signature for each registered operator set
    /// @dev operatorSetKey => tosSignature
    mapping(bytes32 => bytes) public operatorSetTosSignature;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[47] private __gap;

    constructor(IReleaseManager _releaseManager, IAllocationManager _allocationManager) {
        RELEASE_MANAGER = _releaseManager;
        ALLOCATION_MANAGER = _allocationManager;
    }
}
