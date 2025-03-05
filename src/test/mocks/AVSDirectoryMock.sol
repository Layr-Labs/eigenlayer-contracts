// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IAVSDirectory.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract AVSDirectoryMock is Test {
    receive() external payable {}
    fallback() external payable {}
}
