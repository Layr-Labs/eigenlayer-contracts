// SPDX-License-Identifier: agpl-3.0
pragma solidity >=0.8.0;

/**
 * Dummy Weth token.
 */
contract DummyWeth {
    uint256 t;
    
    mapping(address => uint256) b;
    mapping(address => mapping(address => uint256)) a;

    string public name;
    string public symbol;
    uint public decimals;

    function myAddress() external view returns (address) {
        return address(this);
    }

    function totalSupply() external view returns (uint256) {
        return t;
    }

    function balanceOf(address account) external view returns (uint256) {
        return b[account];
    }

    function transfer(address recipient, uint256 amount) external returns (bool) {
        b[msg.sender] -= amount;
        b[recipient] += amount;
        return true;
    }

    function allowance(address owner, address spender) external view returns (uint256) {
        return a[owner][spender];
    }

    function approve(address spender, uint256 amount) external returns (bool) {
        a[msg.sender][spender] = amount;
        return true;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool) {
        b[sender] -= amount;
        b[recipient] +=  amount;
        a[sender][msg.sender] -= amount;
        return true;
    }
    
    // WETH
    function deposit() external payable {
        b[msg.sender] += msg.value;
    }

    function withdraw(uint256 amt) external {
        b[msg.sender] -= amt;
        payable(msg.sender).transfer(amt); // use optimistic_fallback here
    }
}
