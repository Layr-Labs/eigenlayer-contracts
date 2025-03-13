// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "../../contracts/token/Eigen.sol";

contract EigenHarness is Eigen {
    constructor(
        IERC20 _bEIGEN
    ) Eigen(_bEIGEN) {}

    /// expose internal mint function
    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }
}