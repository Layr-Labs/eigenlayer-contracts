// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.0;

import "./OperatorController.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";


contract Operator is OperatorController, IERC1271 {
    constructor(
        address _owner
    ) OperatorController(_owner) {}

    function multicall(address[] memory targets, bytes[] memory data) external returns (bytes[] memory results) {
        require(targets.length == data.length);

        results = new bytes[](targets.length);

        for (uint256 i = 0; i < targets.length; i++) {
            require(hasPermission(msg.sender, targets[i], bytes4(data[i])));
            results[i] = Address.functionCall(targets[i], data[i]);
        }
    }

    /// TODO: Any way to make this cleaner? We cannot extract function signature from the caller, so we are using isValidSignature's selector as the selector
    function isValidSignature(bytes32 hash, bytes memory signature) public view override returns (bytes4) {
        address signer = ECDSA.recover(hash, signature);
        require(hasPermission(signer, msg.sender, IERC1271.isValidSignature.selector));

        return IERC1271.isValidSignature.selector;
    }
}
