// SPDX-License-Identifier: BUSL-1.1
// modified version of https://github.com/itstargetconfirmed/wrapped-ether/blob/master/contracts/WETH.sol
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/// @notice An implementation of Wrapped Ether.
/// @author Anderson Singh.

contract WETH is ERC20 {
    constructor() ERC20("Wrapped Ether", "WETH") {}

    /// @dev mint tokens for sender based on amount of ether sent.
    function deposit() public payable {
        _mint(msg.sender, msg.value);
    }

    /// @dev withdraw ether based on requested amount and user balance.
    function withdraw(uint _amount) external {
        require(balanceOf(msg.sender) >= _amount, "insufficient balance.");
        _burn(msg.sender, _amount);
        payable(msg.sender).transfer(_amount);
    }

    fallback() external payable {
        deposit();
    }

    receive() external payable {
        deposit();
    }
}
