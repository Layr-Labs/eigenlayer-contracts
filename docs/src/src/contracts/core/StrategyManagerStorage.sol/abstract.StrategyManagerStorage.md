# StrategyManagerStorage
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/core/StrategyManagerStorage.sol)

**Inherits:**
[IStrategyManager](/src/contracts/interfaces/IStrategyManager.sol/interface.IStrategyManager.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

This storage contract is separate from the logic to simplify the upgrade process.


## State Variables
### DOMAIN_TYPEHASH
The EIP-712 typehash for the contract's domain


```solidity
bytes32 public constant DOMAIN_TYPEHASH =
    keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
```


### DEPOSIT_TYPEHASH
The EIP-712 typehash for the deposit struct used by the contract


```solidity
bytes32 public constant DEPOSIT_TYPEHASH =
    keccak256("Deposit(address strategy,address token,uint256 amount,uint256 nonce,uint256 expiry)");
```


### DOMAIN_SEPARATOR
EIP-712 Domain separator


```solidity
bytes32 public DOMAIN_SEPARATOR;
```


### nonces

```solidity
mapping(address => uint256) public nonces;
```


### MAX_STAKER_STRATEGY_LIST_LENGTH

```solidity
uint8 internal constant MAX_STAKER_STRATEGY_LIST_LENGTH = 32;
```


### delegation

```solidity
IDelegationManager public immutable delegation;
```


### eigenPodManager

```solidity
IEigenPodManager public immutable eigenPodManager;
```


### slasher

```solidity
ISlasher public immutable slasher;
```


### strategyWhitelister
Permissioned role, which can be changed by the contract owner. Has the ability to edit the strategy whitelist


```solidity
address public strategyWhitelister;
```


### withdrawalDelayBlocks
Minimum delay enforced by this contract for completing queued withdrawals. Measured in blocks, and adjustable by this contract's owner,
up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).

*Note that the withdrawal delay is not enforced on withdrawals of 'beaconChainETH', as the EigenPods have their own separate delay mechanic
and we want to avoid stacking multiple enforced delays onto a single withdrawal.*


```solidity
uint256 public withdrawalDelayBlocks;
```


### MAX_WITHDRAWAL_DELAY_BLOCKS

```solidity
uint256 public constant MAX_WITHDRAWAL_DELAY_BLOCKS = 50400;
```


### stakerStrategyShares
Mapping: staker => Strategy => number of shares which they currently hold


```solidity
mapping(address => mapping(IStrategy => uint256)) public stakerStrategyShares;
```


### stakerStrategyList
Mapping: staker => array of strategies in which they have nonzero shares


```solidity
mapping(address => IStrategy[]) public stakerStrategyList;
```


### withdrawalRootPending
Mapping: hash of withdrawal inputs, aka 'withdrawalRoot' => whether the withdrawal is pending


```solidity
mapping(bytes32 => bool) public withdrawalRootPending;
```


### numWithdrawalsQueued
Mapping: staker => cumulative number of queued withdrawals they have ever initiated. only increments (doesn't decrement)


```solidity
mapping(address => uint256) public numWithdrawalsQueued;
```


### strategyIsWhitelistedForDeposit
Mapping: strategy => whether or not stakers are allowed to deposit into it


```solidity
mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit;
```


### beaconChainETHSharesToDecrementOnWithdrawal

```solidity
mapping(address => uint256) public beaconChainETHSharesToDecrementOnWithdrawal;
```


### beaconChainETHStrategy

```solidity
IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
```


### __gap
*This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps*


```solidity
uint256[40] private __gap;
```


## Functions
### constructor


```solidity
constructor(IDelegationManager _delegation, IEigenPodManager _eigenPodManager, ISlasher _slasher);
```

