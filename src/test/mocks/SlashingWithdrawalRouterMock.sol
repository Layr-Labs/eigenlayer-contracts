// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";

contract SlashingWithdrawalRouterMock is Test {
    receive() external payable {}
    fallback() external payable {}
}
