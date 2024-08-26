// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol";

library SafeERC20Harness is SafeERC20  {
    function callOptionalReturn(IERC20 token, bytes memory data) public {
        SafeERC20._callOptionalReturn(token, data);
    }
}
