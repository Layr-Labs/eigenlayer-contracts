# EigenPodManagerNEW
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/SigP/EigenPodManagerNEW.sol)

**Inherits:**
Initializable, OwnableUpgradeable, [IEigenPodManager](/docs/docgen/src/src/contracts/interfaces/IEigenPodManager.sol/interface.IEigenPodManager.md)

**Author:**
Layr Labs, Inc.

The main functionalities are:
- creating EigenPods
- staking for new validators on EigenPods
- keeping track of the balances of all validators of EigenPods, and their stake in EigenLayer
- withdrawing eth when withdrawals are initiated


## State Variables
### ethPOS

```solidity
IETHPOSDeposit internal immutable ethPOS;
```


### eigenPodBeacon
Beacon proxy to which the EigenPods point


```solidity
IBeacon public immutable eigenPodBeacon;
```


### strategyManager
EigenLayer's StrategyManager contract


```solidity
IStrategyManager public immutable strategyManager;
```


### slasher
EigenLayer's Slasher contract


```solidity
ISlasher public immutable slasher;
```


### beaconChainOracle
Oracle contract that provides updates to the beacon chain's state


```solidity
IBeaconChainOracle public beaconChainOracle;
```


### podOwnerToUnwithdrawnPaidPenalties
Pod owner to the amount of penalties they have paid that are still in this contract


```solidity
mapping(address => uint256) public podOwnerToUnwithdrawnPaidPenalties;
```


## Functions
### getBeaconChainStateRoot


```solidity
function getBeaconChainStateRoot(uint64 slot) external view returns (bytes32);
```

### pause


```solidity
function pause(uint256 newPausedStatus) external;
```

### pauseAll


```solidity
function pauseAll() external;
```

### paused


```solidity
function paused() external view returns (uint256);
```

### paused


```solidity
function paused(uint8 index) external view returns (bool);
```

### pauserRegistry


```solidity
function pauserRegistry() external view returns (IPauserRegistry);
```

### unpause


```solidity
function unpause(uint256 newPausedStatus) external;
```

### ownerToPod


```solidity
function ownerToPod(address podOwner) external view returns (IEigenPod);
```

### onlyEigenPod


```solidity
modifier onlyEigenPod(address podOwner);
```

### onlyStrategyManager


```solidity
modifier onlyStrategyManager();
```

### constructor


```solidity
constructor(IETHPOSDeposit _ethPOS, IBeacon _eigenPodBeacon, IStrategyManager _strategyManager, ISlasher _slasher);
```

### initialize


```solidity
function initialize(IBeaconChainOracle _beaconChainOracle, address initialOwner) public initializer;
```

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
function restakeBeaconChainETH(address podOwner, uint256 amount) external onlyEigenPod(podOwner);
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
    external
    onlyEigenPod(podOwner);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`podOwner`|`address`|The owner of the pod whose balance must be removed.|
|`beaconChainETHStrategyIndex`|`uint256`||
|`amount`|`uint256`|The amount of beacon chain ETH to decrement from the podOwner's shares in the strategyManager.|


### withdrawRestakedBeaconChainETH

Withdraws ETH from an EigenPod. The ETH must have first been withdrawn from the beacon chain.

*Callable only by the StrategyManager contract.*


```solidity
function withdrawRestakedBeaconChainETH(address podOwner, address recipient, uint256 amount)
    external
    onlyStrategyManager;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`podOwner`|`address`|The owner of the pod whose balance must be withdrawn.|
|`recipient`|`address`|The recipient of the withdrawn ETH.|
|`amount`|`uint256`|The amount of ETH to withdraw.|


### payPenalties

Records receiving ETH from the `PodOwner`'s EigenPod, paid in order to fullfill the EigenPod's penalties to EigenLayer

*Callable only by the podOwner's EigenPod contract.*


```solidity
function payPenalties(address podOwner) external payable onlyEigenPod(podOwner);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`podOwner`|`address`|The owner of the pod whose balance is being sent.|


### withdrawPenalties

Withdraws paid penalties of the `podOwner`'s EigenPod, to the `recipient` address

*Callable only by the strategyManager.owner().*


```solidity
function withdrawPenalties(address podOwner, address recipient, uint256 amount) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`podOwner`|`address`||
|`recipient`|`address`|The recipient of withdrawn ETH.|
|`amount`|`uint256`|The amount of ETH to withdraw.|


### updateBeaconChainOracle

Updates the oracle contract that provides the beacon chain state root

*Callable only by the owner of this contract (i.e. governance)*


```solidity
function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newBeaconChainOracle`|`IBeaconChainOracle`|is the new oracle contract being pointed to|


### _deployPod


```solidity
function _deployPod() internal returns (IEigenPod);
```

### _updateBeaconChainOracle


```solidity
function _updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) internal;
```

### getPod

Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).


```solidity
function getPod(address podOwner) public view returns (IEigenPod);
```

### hasPod

Returns 'true' if the `podOwner` has created an EigenPod, and 'false' otherwise.


```solidity
function hasPod(address podOwner) public view returns (bool);
```

### getBeaconChainStateRoot


```solidity
function getBeaconChainStateRoot() external view returns (bytes32);
```

## Events
### BeaconOracleUpdated
Emitted to notify the update of the beaconChainOracle address


```solidity
event BeaconOracleUpdated(address indexed newOracleAddress);
```

### PodDeployed
Emitted to notify the deployment of an EigenPod


```solidity
event PodDeployed(address indexed eigenPod, address indexed podOwner);
```

### BeaconChainETHDeposited
Emitted to notify a deposit of beacon chain ETH recorded in the  manager


```solidity
event BeaconChainETHDeposited(address indexed podOwner, uint256 amount);
```

### PenaltiesPaid
Emitted when an EigenPod pays penalties, on behalf of its owner


```solidity
event PenaltiesPaid(address indexed podOwner, uint256 amountPaid);
```

