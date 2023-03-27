# Solidity API

## IEigenPodManager

### createPod

```solidity
function createPod() external
```

Creates an EigenPod for the sender.

_Function will revert if the `msg.sender` already has an EigenPod._

### stake

```solidity
function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) external payable
```

Stakes for a new beacon chain validator on the sender's EigenPod. 
Also creates an EigenPod for the sender if they don't have one already.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pubkey | bytes | The 48 bytes public key of the beacon chain validator. |
| signature | bytes | The validator's signature of the deposit data. |
| depositDataRoot | bytes32 | The root/hash of the deposit data for the validator's deposit. |

### restakeBeaconChainETH

```solidity
function restakeBeaconChainETH(address podOwner, uint256 amount) external
```

Deposits/Restakes beacon chain ETH in EigenLayer on behalf of the owner of an EigenPod.

_Callable only by the podOwner's EigenPod contract._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| podOwner | address | The owner of the pod whose balance must be deposited. |
| amount | uint256 | The amount of ETH to 'deposit' (i.e. be credited to the podOwner). |

### recordOvercommittedBeaconChainETH

```solidity
function recordOvercommittedBeaconChainETH(address podOwner, uint256 beaconChainETHStrategyIndex, uint256 amount) external
```

Removes beacon chain ETH from EigenLayer on behalf of the owner of an EigenPod, when the
        balance of a validator is lower than how much stake they have committed to EigenLayer

_Callable only by the podOwner's EigenPod contract._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| podOwner | address | The owner of the pod whose balance must be removed. |
| beaconChainETHStrategyIndex | uint256 |  |
| amount | uint256 | The amount of ETH to remove. |

### withdrawRestakedBeaconChainETH

```solidity
function withdrawRestakedBeaconChainETH(address podOwner, address recipient, uint256 amount) external
```

Withdraws ETH from an EigenPod. The ETH must have first been withdrawn from the beacon chain.

_Callable only by the StrategyManager contract._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| podOwner | address | The owner of the pod whose balance must be withdrawn. |
| recipient | address | The recipient of the withdrawn ETH. |
| amount | uint256 | The amount of ETH to withdraw. |

### updateBeaconChainOracle

```solidity
function updateBeaconChainOracle(contract IBeaconChainOracle newBeaconChainOracle) external
```

Updates the oracle contract that provides the beacon chain state root

_Callable only by the owner of this contract (i.e. governance)_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| newBeaconChainOracle | contract IBeaconChainOracle | is the new oracle contract being pointed to |

### getPod

```solidity
function getPod(address podOwner) external view returns (contract IEigenPod)
```

Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).

### beaconChainOracle

```solidity
function beaconChainOracle() external view returns (contract IBeaconChainOracle)
```

Oracle contract that provides updates to the beacon chain's state

### getBeaconChainStateRoot

```solidity
function getBeaconChainStateRoot(uint64 blockNumber) external view returns (bytes32)
```

Returns the Beacon Chain state root at `blockNumber`. Reverts if the Beacon Chain state root at `blockNumber` has not yet been finalized.

### strategyManager

```solidity
function strategyManager() external view returns (contract IStrategyManager)
```

EigenLayer's StrategyManager contract

### slasher

```solidity
function slasher() external view returns (contract ISlasher)
```

EigenLayer's Slasher contract

### hasPod

```solidity
function hasPod(address podOwner) external view returns (bool)
```

