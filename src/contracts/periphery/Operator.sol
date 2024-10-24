// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.0;

import "./RolesAuthority.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";


contract Operator is RolesAuthority {
    constructor(address _owner) RolesAuthority(_owner) {}

    function multicall(address[] memory targets, bytes[] memory data) external returns (bytes[] memory results) {
        require(targets.length == data.length);

        results = new bytes[](targets.length);

        for (uint256 i = 0; i < targets.length; i++) {
            require(canCall(msg.sender, targets[i], bytes4(data[i])));
            results[i] = Address.functionCall(targets[i], data[i]);
        }
    }

    // TODO: implement 1271 support
    /// @dev One idea here is to update the 712 typehash to include the function signature
    /// of the function that is validating the signature. This way, users with a given role
    /// can sign for the given (target, selector).
    function isValidSignature(bytes32 hash, bytes memory signature) external view returns (bytes4 magicValue) {
        return IERC1271.isValidSignature.selector;
    }
}