// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";

contract OperatorTableUpdaterMock is Test {
    receive() external payable {}
    fallback() external payable {}

    mapping(uint32 referenceTimestamp => bool valid) internal _invalidRoots;

    function isRootValidByTimestamp(uint32 referenceTimestamp) external view returns (bool) {
        // If the root is invalid, return false
        return !_invalidRoots[referenceTimestamp];
    }

    function invalidateRoot(uint32 referenceTimestamp) external {
        _invalidRoots[referenceTimestamp] = true;
    }
}
