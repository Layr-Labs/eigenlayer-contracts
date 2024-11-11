// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IPermissionController.sol";

contract PermissionControllerMock is Test {
    receive() external payable {}
    fallback() external payable {}

    bool canCallResult = true;

    // Return true for now
    function canCall(address account, address caller, address target, bytes4 selector) external view returns (bool) {
        return canCallResult;
    }

    function setCanCallResult(bool result) public {
        canCallResult = result;
    }
}