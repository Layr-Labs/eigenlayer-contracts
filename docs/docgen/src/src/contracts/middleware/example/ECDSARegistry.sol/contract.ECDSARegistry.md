# ECDSARegistry
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/middleware/example/ECDSARegistry.sol)

**Inherits:**
[RegistryBase](/docs/docgen/src/src/contracts/middleware/RegistryBase.sol/abstract.RegistryBase.md)

**Author:**
Layr Labs, Inc.

This contract is used for
- registering new operators
- committing to and finalizing de-registration as an operator
- updating the stakes of the operator


## State Variables
### operatorWhitelister
the address that can whitelist people


```solidity
address public operatorWhitelister;
```


### operatorWhitelistEnabled
toggle of whether the operator whitelist is on or off


```solidity
bool public operatorWhitelistEnabled;
```


### whitelisted
operator => are they whitelisted (can they register with the middleware)


```solidity
mapping(address => bool) public whitelisted;
```


## Functions
### onlyOperatorWhitelister

Modifier that restricts a function to only be callable by the `whitelister` role.


```solidity
modifier onlyOperatorWhitelister();
```

### constructor


```solidity
constructor(IStrategyManager _strategyManager, IServiceManager _serviceManager)
    RegistryBase(_strategyManager, _serviceManager, 1);
```

### initialize

Initialize whitelister and the quorum strategies + multipliers.


```solidity
function initialize(
    address _operatorWhitelister,
    bool _operatorWhitelistEnabled,
    uint256[] memory _quorumBips,
    StrategyAndWeightingMultiplier[] memory _quorumStrategiesConsideredAndMultipliers
) public virtual initializer;
```

### setOperatorWhitelister

Called by the service manager owner to transfer the whitelister role to another address


```solidity
function setOperatorWhitelister(address _operatorWhitelister) external onlyServiceManagerOwner;
```

### setOperatorWhitelistStatus

Callable only by the service manager owner, this function toggles the whitelist on or off


```solidity
function setOperatorWhitelistStatus(bool _operatorWhitelistEnabled) external onlyServiceManagerOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_operatorWhitelistEnabled`|`bool`|true if turning whitelist on, false otherwise|


### addToOperatorWhitelist

Called by the operatorWhitelister, adds a list of operators to the whitelist


```solidity
function addToOperatorWhitelist(address[] calldata operators) external onlyOperatorWhitelister;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operators`|`address[]`|the operators to add to the whitelist|


### removeFromWhitelist

Called by the operatorWhitelister, removes a list of operators to the whitelist


```solidity
function removeFromWhitelist(address[] calldata operators) external onlyOperatorWhitelister;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operators`|`address[]`|the operators to remove from the whitelist|


### registerOperator

called for registering as an operator


```solidity
function registerOperator(string calldata socket) external virtual;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`socket`|`string`|is the socket address of the operator|


### _registerOperator


```solidity
function _registerOperator(address operator, string calldata socket) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the node who is registering to be a operator|
|`socket`|`string`|is the socket address of the operator|


### deregisterOperator

Used by an operator to de-register itself from providing service to the middleware.


```solidity
function deregisterOperator(uint32 index) external virtual returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`index`|`uint32`|is the sender's location in the dynamic array `operatorList`|


### _deregisterOperator

Used to process de-registering an operator from providing service to the middleware.


```solidity
function _deregisterOperator(address operator, uint32 index) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|The operator to be deregistered|
|`index`|`uint32`|is the sender's location in the dynamic array `operatorList`|


### updateStakes

Used for updating information on deposits of nodes.


```solidity
function updateStakes(address[] calldata operators, uint256[] calldata prevElements) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operators`|`address[]`|are the nodes whose deposit information is getting updated|
|`prevElements`|`uint256[]`|are the elements before this middleware in the operator's linked list within the slasher|


### _setOperatorWhitelister


```solidity
function _setOperatorWhitelister(address _operatorWhitelister) internal;
```

## Events
### Registration
Emitted upon the registration of a new operator for the middleware


```solidity
event Registration(address indexed operator, string socket);
```

### OperatorWhitelisterTransferred
Emitted when the `operatorWhitelister` role is transferred.


```solidity
event OperatorWhitelisterTransferred(address previousAddress, address newAddress);
```

