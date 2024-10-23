// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "./IOperatorController.sol";

/// @notice A roles based access control contract that supports up to 256 roles.
/// @author Modified from Solmate (https://github.com/transmissions11/solmate/blob/main/src/auth/authorities/RolesAuthority.sol)

contract OperatorController is IOperatorController, Ownable {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

    /// @notice Mapping from a delegate to the list of encoded target & selectors
    mapping(address delegate => EnumerableSet.Bytes32Set) internal _delegatePermissions;

    /// @notice Mapping from encoded target & selector to the list of delegates
    mapping(address target => mapping(bytes4 selector => EnumerableSet.AddressSet)) internal _permissionDelegates;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */
    constructor(
        address _owner
    ) {
        _transferOwnership(_owner);
    }

    /**
     *
     *                         Delegation Configuration
     *
     */

    /// @inheritdoc IOperatorController
    function setDelegatedRole(address delegate, address target, bytes4 selector) external override onlyOwner {
        _delegatePermissions[delegate].add(_encodeTargetSelector(target, selector));
        _permissionDelegates[target][selector].add(delegate);

        emit DelegateSet(delegate, target, selector);
    }

    /// @inheritdoc IOperatorController
    function revokeDelegatedRole(address delegate, address target, bytes4 selector) external override onlyOwner {
        _delegatePermissions[delegate].remove(_encodeTargetSelector(target, selector));
        _permissionDelegates[target][selector].remove(delegate);

        emit DelegateRevoked(delegate, target, selector);
    }

    /**
     *
     *                         View Functions
     *
     */

    /// @inheritdoc IOperatorController
    function hasPermission(address delegate, address target, bytes4 selector) public view override returns (bool) {
        return _permissionDelegates[target][selector].contains(delegate);
    }

    /// @inheritdoc IOperatorController
    function getDelegatePermissions(
        address delegate
    ) external view override returns (address[] memory, bytes4[] memory) {
        uint256 length = _delegatePermissions[delegate].length();

        address[] memory targets = new address[](length);
        bytes4[] memory selectors = new bytes4[](length);

        for (uint256 i = 0; i < length; i++) {
            (address target, bytes4 selector) = _decodeTargetSelector(_delegatePermissions[delegate].at(i));
            targets[i] = target;
            selectors[i] = selector;
        }

        return (targets, selectors);
    }

    /// @inheritdoc IOperatorController
    function getDelegates(address target, bytes4 selector) external view override returns (address[] memory) {
        return _permissionDelegates[target][selector].values();
    }

    /**
     *
     *                         Helper Functions
     *
     */
    function _encodeTargetSelector(address target, bytes4 selector) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(target, uint96(bytes12((selector)))));
    }

    function _decodeTargetSelector(
        bytes32 key
    ) internal pure returns (address, bytes4) {
        address target = address(uint160(uint256(key) >> 96));
        bytes4 selector = bytes4(uint32((uint256(key) & type(uint96).max)));

        return (target, selector);
    }
}
