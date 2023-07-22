# DelegationManagerStorage
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/core/DelegationManagerStorage.sol)

**Inherits:**
[IDelegationManager](/src/contracts/interfaces/IDelegationManager.sol/interface.IDelegationManager.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

This storage contract is separate from the logic to simplify the upgrade process.


## State Variables
### LOW_LEVEL_GAS_BUDGET
Gas budget provided in calls to DelegationTerms contracts


```solidity
uint256 internal constant LOW_LEVEL_GAS_BUDGET = 1e5;
```


### DOMAIN_TYPEHASH
The EIP-712 typehash for the contract's domain


```solidity
bytes32 public constant DOMAIN_TYPEHASH =
    keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
```


### DELEGATION_TYPEHASH
The EIP-712 typehash for the delegation struct used by the contract


```solidity
bytes32 public constant DELEGATION_TYPEHASH =
    keccak256("Delegation(address delegator,address operator,uint256 nonce,uint256 expiry)");
```


### DOMAIN_SEPARATOR
EIP-712 Domain separator


```solidity
bytes32 public DOMAIN_SEPARATOR;
```


### strategyManager
The StrategyManager contract for EigenLayer


```solidity
IStrategyManager public immutable strategyManager;
```


### slasher
The Slasher contract for EigenLayer


```solidity
ISlasher public immutable slasher;
```


### operatorShares
Mapping: operator => strategy => total number of shares in the strategy delegated to the operator


```solidity
mapping(address => mapping(IStrategy => uint256)) public operatorShares;
```


### delegationTerms
Mapping: operator => delegation terms contract


```solidity
mapping(address => IDelegationTerms) public delegationTerms;
```


### delegatedTo
Mapping: staker => operator whom the staker has delegated to


```solidity
mapping(address => address) public delegatedTo;
```


### nonces
Mapping: delegator => number of signed delegation nonce (used in delegateToBySignature)


```solidity
mapping(address => uint256) public nonces;
```


### __gap
*This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps*


```solidity
uint256[46] private __gap;
```


## Functions
### constructor


```solidity
constructor(IStrategyManager _strategyManager, ISlasher _slasher);
```

