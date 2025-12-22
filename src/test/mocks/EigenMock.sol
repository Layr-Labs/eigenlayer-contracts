// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./BackingEigenMock.sol";

contract EigenMock is Test, ERC20 {
    BackingEigenMock backingEigen;

    constructor(BackingEigenMock _backingEigen) ERC20("EIGEN", "EIGEN") {
        backingEigen = _backingEigen;
    }

    function wrap(uint amount) external {
        backingEigen.transferFrom(msg.sender, address(this), amount);
        _mint(msg.sender, amount);
    }
}
