// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Allocator {

    function allocate(IERC20 token, address[] memory recipients, uint256 amount) public {
        token.transferFrom(msg.sender, address(this), recipients.length * amount);
        for (uint i = 0; i < recipients.length; i++) {
            token.transfer(recipients[i], amount);
        }
    }
}
