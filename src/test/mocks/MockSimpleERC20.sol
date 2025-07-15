// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title MockSimpleERC20
 * @notice Mock ERC20 token for testing
 * @dev Simple ERC20 implementation with mint function for testing
 */
contract MockSimpleERC20 is ERC20 {
    constructor() ERC20("Mock Token", "MOCK") {}

    function mint(address to, uint amount) public {
        _mint(to, amount);
    }
}
