# StrategyManagerMock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/StrategyManagerMock.sol)

**Inherits:**
Initializable, [IStrategyManager](/docs/docgen/src/src/contracts/interfaces/IStrategyManager.sol/interface.IStrategyManager.md), OwnableUpgradeable, ReentrancyGuardUpgradeable, [Pausable](/docs/docgen/src/src/contracts/permissions/Pausable.sol/contract.Pausable.md)


## State Variables
### delegation

```solidity
IDelegationManager public delegation;
```


### eigenPodManager

```solidity
IEigenPodManager public eigenPodManager;
```


### slasher

```solidity
ISlasher public slasher;
```


## Functions
### setAddresses


```solidity
function setAddresses(IDelegationManager _delegation, IEigenPodManager _eigenPodManager, ISlasher _slasher) external;
```

### depositIntoStrategy


```solidity
function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount) external returns (uint256);
```

### depositBeaconChainETH


```solidity
function depositBeaconChainETH(address staker, uint256 amount) external;
```

### recordOvercommittedBeaconChainETH


```solidity
function recordOvercommittedBeaconChainETH(
    address overcommittedPodOwner,
    uint256 beaconChainETHStrategyIndex,
    uint256 amount
) external;
```

### depositIntoStrategyWithSignature


```solidity
function depositIntoStrategyWithSignature(
    IStrategy strategy,
    IERC20 token,
    uint256 amount,
    address staker,
    uint256 expiry,
    bytes memory signature
) external returns (uint256 shares);
```

### stakerStrategyShares

Returns the current shares of `user` in `strategy`


```solidity
function stakerStrategyShares(address user, IStrategy strategy) external view returns (uint256 shares);
```

### getDeposits

Get all details on the depositor's deposits and corresponding shares


```solidity
function getDeposits(address depositor) external view returns (IStrategy[] memory, uint256[] memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`IStrategy[]`|(depositor's strategies, shares in these strategies)|
|`<none>`|`uint256[]`||


### stakerStrats

Returns the array of strategies in which `staker` has nonzero shares


```solidity
function stakerStrats(address staker) external view returns (IStrategy[] memory);
```

### stakerStrategyListLength

Simple getter function that returns `stakerStrategyList[staker].length`.


```solidity
function stakerStrategyListLength(address staker) external view returns (uint256);
```

### queueWithdrawal


```solidity
function queueWithdrawal(
    uint256[] calldata strategyIndexes,
    IStrategy[] calldata strategies,
    uint256[] calldata shares,
    address withdrawer,
    bool undelegateIfPossible
) external returns (bytes32);
```

### completeQueuedWithdrawal


```solidity
function completeQueuedWithdrawal(
    QueuedWithdrawal calldata queuedWithdrawal,
    IERC20[] calldata tokens,
    uint256 middlewareTimesIndex,
    bool receiveAsTokens
) external;
```

### completeQueuedWithdrawals


```solidity
function completeQueuedWithdrawals(
    QueuedWithdrawal[] calldata queuedWithdrawals,
    IERC20[][] calldata tokens,
    uint256[] calldata middlewareTimesIndexes,
    bool[] calldata receiveAsTokens
) external;
```

### slashShares


```solidity
function slashShares(
    address slashedAddress,
    address recipient,
    IStrategy[] calldata strategies,
    IERC20[] calldata tokens,
    uint256[] calldata strategyIndexes,
    uint256[] calldata shareAmounts
) external;
```

### slashQueuedWithdrawal

Slashes an existing queued withdrawal that was created by a 'frozen' operator (or a staker delegated to one)


```solidity
function slashQueuedWithdrawal(
    address recipient,
    QueuedWithdrawal calldata queuedWithdrawal,
    IERC20[] calldata tokens,
    uint256[] calldata indicesToSkip
) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`recipient`|`address`|The funds in the slashed withdrawal are withdrawn as tokens to this address.|
|`queuedWithdrawal`|`QueuedWithdrawal`||
|`tokens`|`IERC20[]`||
|`indicesToSkip`|`uint256[]`||


### calculateWithdrawalRoot

Returns the keccak256 hash of `queuedWithdrawal`.


```solidity
function calculateWithdrawalRoot(QueuedWithdrawal memory queuedWithdrawal) external pure returns (bytes32);
```

### beaconChainETHStrategy

returns the enshrined beaconChainETH Strategy


```solidity
function beaconChainETHStrategy() external view returns (IStrategy);
```

### withdrawalDelayBlocks


```solidity
function withdrawalDelayBlocks() external view returns (uint256);
```

### addStrategiesToDepositWhitelist


```solidity
function addStrategiesToDepositWhitelist(IStrategy[] calldata) external pure;
```

### removeStrategiesFromDepositWhitelist


```solidity
function removeStrategiesFromDepositWhitelist(IStrategy[] calldata) external pure;
```

