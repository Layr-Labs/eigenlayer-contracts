// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";

contract OperatorTableUpdaterMock is Test {
    receive() external payable {}
    fallback() external payable {}
}
