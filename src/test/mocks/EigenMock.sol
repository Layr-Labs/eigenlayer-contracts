// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract EigenMock is Test, ERC20 {
    constructor() ERC20("EIGEN", "EIGEN") {}

    function mint(address to, uint amount) external {
        _mint(to, amount);
    }
}

