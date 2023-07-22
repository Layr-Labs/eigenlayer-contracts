# IEigenPodManager
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IEigenPodManager.sol)

**Inherits:**
[IPausable](/src/contracts/interfaces/IPausable.sol/interface.IPausable.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service


## Functions
### createPod

Creates an EigenPod for the sender.

*Function will revert if the `msg.sender` already has an EigenPod.*


```solidity
function createPod() external;
```

### stake

Stakes for a new beacon chain validator on the sender's EigenPod.
Also creates an EigenPod for the sender if they don't have one already.


```solidity
function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`pubkey`|`bytes`|The 48 bytes public key of the beacon chain validator.|
|`signature`|`bytes`|The validator's signature of the deposit data.|
|`depositDataRoot`|`bytes32`|The root/hash of the deposit data for the validator's deposit.|


### restakeBeaconChainETH

Deposits/Restakes beacon chain ETH in EigenLayer on behalf of the owner of an EigenPod.

*Callable only by the podOwner's EigenPod contract.*


```solidity
function restakeBeaconChainETH(address podOwner, uint256 amount) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`podOwner`|`address`|The owner of the pod whose balance must be deposited.|
|`amount`|`uint256`|The amount of ETH to 'deposit' (i.e. be credited to the podOwner).|


### recordOvercommittedBeaconChainETH

Removes beacon chain ETH from EigenLayer on behalf of the owner of an EigenPod, when the
balance of a validator is lower than how much stake they have committed to EigenLayer

*Callable only by the podOwner's EigenPod contract.*


```solidity
function recordOvercommittedBeaconChainETH(address podOwner, uint256 beaconChainETHStrategyIndex, uint256 amount)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`podOwner`|`address`|The owner of the pod whose balance must be removed.|
|`beaconChainETHStrategyIndex`|`uint256`|is the index of the beaconChainETHStrategy for the pod owner for the callback to the StrategyManager in case it must be removed from the list of the podOwner's strategies|
|`amount`|`uint256`|The amount of ETH to remove.|


### withdrawRestakedBeaconChainETH

Withdraws ETH from an EigenPod. The ETH must have first been withdrawn from the beacon chain.

*Callable only by the StrategyManager contract.*


```solidity
function withdrawRestakedBeaconChainETH(address podOwner, address recipient, uint256 amount) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`podOwner`|`address`|The owner of the pod whose balance must be withdrawn.|
|`recipient`|`address`|The recipient of the withdrawn ETH.|
|`amount`|`uint256`|The amount of ETH to withdraw.|


### updateBeaconChainOracle

Updates the oracle contract that provides the beacon chain state root

*Callable only by the owner of this contract (i.e. governance)*


```solidity
function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newBeaconChainOracle`|`IBeaconChainOracle`|is the new oracle contract being pointed to|


### ownerToPod

Returns the address of the `podOwner`'s EigenPod if it has been deployed.


```solidity
function ownerToPod(address podOwner) external view returns (IEigenPod);
```

### getPod

Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).


```solidity
function getPod(address podOwner) external view returns (IEigenPod);
```

### beaconChainOracle

Oracle contract that provides updates to the beacon chain's state


```solidity
function beaconChainOracle() external view returns (IBeaconChainOracle);
```

### getBeaconChainStateRoot

Returns the Beacon Chain state root at `blockNumber`. Reverts if the Beacon Chain state root at `blockNumber` has not yet been finalized.


```solidity
function getBeaconChainStateRoot(uint64 blockNumber) external view returns (bytes32);
```

### strategyManager

EigenLayer's StrategyManager contract


```solidity
function strategyManager() external view returns (IStrategyManager);
```

### slasher

EigenLayer's Slasher contract


```solidity
function slasher() external view returns (ISlasher);
```

### hasPod


```solidity
function hasPod(address podOwner) external view returns (bool);
```

