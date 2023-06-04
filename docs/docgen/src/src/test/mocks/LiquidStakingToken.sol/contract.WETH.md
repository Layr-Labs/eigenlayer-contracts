# WETH
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/LiquidStakingToken.sol)

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

