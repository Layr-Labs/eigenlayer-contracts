# Solidity API

## ECDSARegistry

This contract is used for
- registering new operators
- committing to and finalizing de-registration as an operator
- updating the stakes of the operator

### operatorWhitelister

```solidity
address operatorWhitelister
```

the address that can whitelist people

### operatorWhitelistEnabled

```solidity
bool operatorWhitelistEnabled
```

toggle of whether the operator whitelist is on or off

### whitelisted

```solidity
mapping(address => bool) whitelisted
```

operator => are they whitelisted (can they register with the middleware)

### Registration

```solidity
event Registration(address operator, string socket)
```

Emitted upon the registration of a new operator for the middleware

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | Address of the new operator |
| socket | string | The ip:port of the operator |

### OperatorWhitelisterTransferred

```solidity
event OperatorWhitelisterTransferred(address previousAddress, address newAddress)
```

Emitted when the `operatorWhitelister` role is transferred.

### onlyOperatorWhitelister

```solidity
modifier onlyOperatorWhitelister()
```

Modifier that restricts a function to only be callable by the `whitelister` role.

### constructor

```solidity
constructor(contract IStrategyManager _strategyManager, contract IServiceManager _serviceManager) public
```

### initialize

```solidity
function initialize(address _operatorWhitelister, bool _operatorWhitelistEnabled, uint256[] _quorumBips, struct VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] _quorumStrategiesConsideredAndMultipliers) public virtual
```

Initialize whitelister and the quorum strategies + multipliers.

### setOperatorWhitelister

```solidity
function setOperatorWhitelister(address _operatorWhitelister) external
```

Called by the service manager owner to transfer the whitelister role to another address

### setOperatorWhitelistStatus

```solidity
function setOperatorWhitelistStatus(bool _operatorWhitelistEnabled) external
```

Callable only by the service manager owner, this function toggles the whitelist on or off

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| _operatorWhitelistEnabled | bool | true if turning whitelist on, false otherwise |

### addToOperatorWhitelist

```solidity
function addToOperatorWhitelist(address[] operators) external
```

Called by the operatorWhitelister, adds a list of operators to the whitelist

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operators | address[] | the operators to add to the whitelist |

### removeFromWhitelist

```solidity
function removeFromWhitelist(address[] operators) external
```

Called by the operatorWhitelister, removes a list of operators to the whitelist

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operators | address[] | the operators to remove from the whitelist |

### registerOperator

```solidity
function registerOperator(string socket) external virtual
```

called for registering as an operator

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| socket | string | is the socket address of the operator |

### _registerOperator

```solidity
function _registerOperator(address operator, string socket) internal
```

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | is the node who is registering to be a operator |
| socket | string | is the socket address of the operator |

### deregisterOperator

```solidity
function deregisterOperator(uint32 index) external virtual returns (bool)
```

Used by an operator to de-register itself from providing service to the middleware.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| index | uint32 | is the sender's location in the dynamic array `operatorList` |

### _deregisterOperator

```solidity
function _deregisterOperator(address operator, uint32 index) internal
```

Used to process de-registering an operator from providing service to the middleware.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | The operator to be deregistered |
| index | uint32 | is the sender's location in the dynamic array `operatorList` |

### updateStakes

```solidity
function updateStakes(address[] operators, uint256[] prevElements) external
```

Used for updating information on deposits of nodes.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operators | address[] | are the nodes whose deposit information is getting updated |
| prevElements | uint256[] | are the elements before this middleware in the operator's linked list within the slasher |

### _setOperatorWhitelister

```solidity
function _setOperatorWhitelister(address _operatorWhitelister) internal
```

