# WETH
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/LiquidStakingToken.sol)

**Inherits:**
ERC20

**Author:**
Anderson Singh.

An implementation of Wrapped Ether.


## Functions
### constructor


```solidity
constructor() ERC20("Wrapped Ether", "WETH");
```

### deposit

*mint tokens for sender based on amount of ether sent.*


```solidity
function deposit() public payable;
```

### withdraw

*withdraw ether based on requested amount and user balance.*


```solidity
function withdraw(uint256 _amount) external;
```

### fallback


```solidity
fallback() external payable;
```

### receive


```solidity
receive() external payable;
```

