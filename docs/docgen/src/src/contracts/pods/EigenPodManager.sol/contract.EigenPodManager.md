# EigenPodManager
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/pods/EigenPodManager.sol)

**Inherits:**
Initializable, OwnableUpgradeable, [Pausable](/docs/docgen/src/src/contracts/permissions/Pausable.sol/contract.Pausable.md), [IEigenPodManager](/docs/docgen/src/src/contracts/interfaces/IEigenPodManager.sol/interface.IEigenPodManager.md), [EigenPodPausingConstants](/docs/docgen/src/src/contracts/pods/EigenPodPausingConstants.sol/abstract.EigenPodPausingConstants.md)

**Author:**
Layr Labs, Inc.

The main functionalities are:
- creating EigenPods
- staking for new validators on EigenPods
- keeping track of the balances of all validators of EigenPods, and their stake in EigenLayer
- withdrawing eth when withdrawals are initiated


## State Variables
### beaconProxyBytecode
Stored code of type(BeaconProxy).creationCode

*Maintained as a constant to solve an edge case - changes to OpenZeppelin's BeaconProxy code should not cause
addresses of EigenPods that are pre-computed with Create2 to change, even upon upgrading this contract, changing compiler version, etc.*


```solidity
bytes internal constant beaconProxyBytecode =
    hex"608060405260405161090e38038061090e83398101604081905261002291610460565b61002e82826000610035565b505061058a565b61003e83610100565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100fb576100f9836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e99190610520565b836102a360201b6100291760201c565b505b505050565b610113816102cf60201b6100551760201c565b6101725760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101e6816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d79190610520565b6102cf60201b6100551760201c565b61024b5760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610169565b806102827fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6102de60201b6100641760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606102c883836040518060600160405280602781526020016108e7602791396102e1565b9392505050565b6001600160a01b03163b151590565b90565b6060600080856001600160a01b0316856040516102fe919061053b565b600060405180830381855af49150503d8060008114610339576040519150601f19603f3d011682016040523d82523d6000602084013e61033e565b606091505b5090925090506103508683838761035a565b9695505050505050565b606083156103c65782516103bf576001600160a01b0385163b6103bf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610169565b50816103d0565b6103d083836103d8565b949350505050565b8151156103e85781518083602001fd5b8060405162461bcd60e51b81526004016101699190610557565b80516001600160a01b038116811461041957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561044f578181015183820152602001610437565b838111156100f95750506000910152565b6000806040838503121561047357600080fd5b61047c83610402565b60208401519092506001600160401b038082111561049957600080fd5b818501915085601f8301126104ad57600080fd5b8151818111156104bf576104bf61041e565b604051601f8201601f19908116603f011681019083821181831017156104e7576104e761041e565b8160405282815288602084870101111561050057600080fd5b610511836020830160208801610434565b80955050505050509250929050565b60006020828403121561053257600080fd5b6102c882610402565b6000825161054d818460208701610434565b9190910192915050565b6020815260008251806020840152610576816040850160208701610434565b601f01601f19169190910160400192915050565b61034e806105996000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610067565b610100565b565b606061004e83836040518060600160405280602781526020016102f260279139610124565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fb9190610249565b905090565b3660008037600080366000845af43d6000803e80801561011f573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161014191906102a2565b600060405180830381855af49150503d806000811461017c576040519150601f19603f3d011682016040523d82523d6000602084013e610181565b606091505b50915091506101928683838761019c565b9695505050505050565b6060831561020d578251610206576001600160a01b0385163b6102065760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610217565b610217838361021f565b949350505050565b81511561022f5781518083602001fd5b8060405162461bcd60e51b81526004016101fd91906102be565b60006020828403121561025b57600080fd5b81516001600160a01b038116811461004e57600080fd5b60005b8381101561028d578181015183820152602001610275565b8381111561029c576000848401525b50505050565b600082516102b4818460208701610272565b9190910192915050565b60208152600082518060208401526102dd816040850160208701610272565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d51e81d3bc5ed20a26aeb05dce7e825c503b2061aa78628027300c8d65b9d89a64736f6c634300080c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564";
```


### ethPOS
The ETH2 Deposit Contract


```solidity
IETHPOSDeposit public immutable ethPOS;
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


### ownerToPod
Pod owner to deployed EigenPod address


```solidity
mapping(address => IEigenPod) public ownerToPod;
```


### numPods
The number of EigenPods that have been deployed


```solidity
uint256 public numPods;
```


### maxPods
The maximum number of EigenPods that can be deployed


```solidity
uint256 public maxPods;
```


### __gap
*This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps*


```solidity
uint256[46] private __gap;
```


## Functions
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
function initialize(
    uint256 _maxPods,
    IBeaconChainOracle _beaconChainOracle,
    address initialOwner,
    IPauserRegistry _pauserRegistry,
    uint256 _initPausedStatus
) external initializer;
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
|`beaconChainETHStrategyIndex`|`uint256`|is the index of the beaconChainETHStrategy for the pod owner for the callback to the StrategyManager in case it must be removed from the list of the podOwner's strategies|
|`amount`|`uint256`|The amount of beacon chain ETH to decrement from the podOwner's shares in the strategyManager.|


### withdrawRestakedBeaconChainETH

Withdraws ETH from an EigenPod. The ETH must have first been withdrawn from the beacon chain.

*Callable only by the StrategyManager contract.*


```solidity
function withdrawRestakedBeaconChainETH(address podOwner, address recipient, uint256 amount)
    external
    onlyStrategyManager
    onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`podOwner`|`address`|The owner of the pod whose balance must be withdrawn.|
|`recipient`|`address`|The recipient of the withdrawn ETH.|
|`amount`|`uint256`|The amount of ETH to withdraw.|


### setMaxPods

Sets the maximum number of pods that can be deployed

*Callable by the pauser of this contract*


```solidity
function setMaxPods(uint256 newMaxPods) external onlyPauser;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newMaxPods`|`uint256`|The new maximum number of pods that can be deployed|


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
function _deployPod() internal onlyWhenNotPaused(PAUSED_NEW_EIGENPODS) returns (IEigenPod);
```

### _updateBeaconChainOracle

Internal setter for `beaconChainOracle` that also emits an event


```solidity
function _updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) internal;
```

### _setMaxPods

Internal setter for `maxPods` that also emits an event


```solidity
function _setMaxPods(uint256 _maxPods) internal;
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

Returns the Beacon Chain state root at `blockNumber`. Reverts if the Beacon Chain state root at `blockNumber` has not yet been finalized.


```solidity
function getBeaconChainStateRoot(uint64 blockNumber) external view returns (bytes32);
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
Emitted to notify a deposit of beacon chain ETH recorded in the strategy manager


```solidity
event BeaconChainETHDeposited(address indexed podOwner, uint256 amount);
```

### MaxPodsUpdated
Emitted when `maxPods` value is updated from `previousValue` to `newValue`


```solidity
event MaxPodsUpdated(uint256 previousValue, uint256 newValue);
```

