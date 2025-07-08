// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IAllocationManager.sol";
import "../interfaces/IKeyRegistrar.sol";

abstract contract KeyRegistrarStorage is IKeyRegistrar {
    // Immutables

    /// @dev Reference to the AllocationManager contract
    IAllocationManager public immutable allocationManager;

    // Mutatables

    /// @dev Maps (operatorSetKey, operator) to their key info
    mapping(bytes32 operatorSetKey => mapping(address operator => KeyInfo keyInfo)) internal _operatorKeyInfo;

    /// @dev Maps operatorSetKey to the key type
    mapping(bytes32 operatorSetKey => CurveType curveType) internal _operatorSetCurveTypes;

    /// @dev Global mapping of key hash to registration status - enforces global uniqueness
    mapping(bytes32 keyHash => bool isRegistered) internal _globalKeyRegistry;

    /// @dev Mapping from (keyHash) to the operator
    mapping(bytes32 keyHash => address operator) internal _keyHashToOperator;

    // Construction

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[46] private __gap;
}
