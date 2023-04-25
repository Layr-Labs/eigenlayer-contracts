# Solidity API

## DelegationManagerStorage

This storage contract is separate from the logic to simplify the upgrade process.

### LOW_LEVEL_GAS_BUDGET

```solidity
uint256 LOW_LEVEL_GAS_BUDGET
```

Gas budget provided in calls to DelegationTerms contracts

### DOMAIN_TYPEHASH

```solidity
bytes32 DOMAIN_TYPEHASH
```

The EIP-712 typehash for the contract's domain

### DELEGATION_TYPEHASH

```solidity
bytes32 DELEGATION_TYPEHASH
```

The EIP-712 typehash for the delegation struct used by the contract

### DOMAIN_SEPARATOR

```solidity
bytes32 DOMAIN_SEPARATOR
```

EIP-712 Domain separator

### strategyManager

```solidity
contract IStrategyManager strategyManager
```

The StrategyManager contract for EigenLayer

### slasher

```solidity
contract ISlasher slasher
```

The Slasher contract for EigenLayer

### operatorShares

```solidity
mapping(address => mapping(contract IStrategy => uint256)) operatorShares
```

Mapping: operator => strategy => total number of shares in the strategy delegated to the operator

### delegationTerms

```solidity
mapping(address => contract IDelegationTerms) delegationTerms
```

Mapping: operator => delegation terms contract

### delegatedTo

```solidity
mapping(address => address) delegatedTo
```

Mapping: staker => operator whom the staker has delegated to

### nonces

```solidity
mapping(address => uint256) nonces
```

Mapping: delegator => number of signed delegation nonce (used in delegateToBySignature)

### constructor

```solidity
constructor(contract IStrategyManager _strategyManager, contract ISlasher _slasher) internal
```

### __gap

```solidity
uint256[46] __gap
```

_This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps_

