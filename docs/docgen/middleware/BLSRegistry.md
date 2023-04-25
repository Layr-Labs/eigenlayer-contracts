# Solidity API

## BLSRegistry

This contract is used for
- registering new operators
- committing to and finalizing de-registration as an operator
- updating the stakes of the operator

### ZERO_PK_HASH

```solidity
bytes32 ZERO_PK_HASH
```

### pubkeyCompendium

```solidity
contract IBLSPublicKeyCompendium pubkeyCompendium
```

contract used for looking up operators' BLS public keys

### _apkUpdates

```solidity
struct IBLSRegistry.ApkUpdate[] _apkUpdates
```

list of keccak256(apk_x, apk_y) of operators, and the block numbers at which the aggregate
pubkeys were updated. This occurs whenever a new operator registers or deregisters.

### apk

```solidity
struct BN254.G1Point apk
```

used for storing current aggregate public key

_Initialized value of APK is the point at infinity: (0, 0)_

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
event Registration(address operator, bytes32 pkHash, struct BN254.G1Point pk, uint32 apkHashIndex, bytes32 apkHash, string socket)
```

Emitted upon the registration of a new operator for the middleware

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | Address of the new operator |
| pkHash | bytes32 | The keccak256 hash of the operator's public key |
| pk | struct BN254.G1Point | The operator's public key itself |
| apkHashIndex | uint32 | The index of the latest (i.e. the new) APK update |
| apkHash | bytes32 | The keccak256 hash of the new Aggregate Public Key |
| socket | string |  |

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
constructor(contract IStrategyManager _strategyManager, contract IServiceManager _serviceManager, uint8 _NUMBER_OF_QUORUMS, contract IBLSPublicKeyCompendium _pubkeyCompendium) public
```

### initialize

```solidity
function initialize(address _operatorWhitelister, bool _operatorWhitelistEnabled, uint256[] _quorumBips, struct VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] _firstQuorumStrategiesConsideredAndMultipliers, struct VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] _secondQuorumStrategiesConsideredAndMultipliers) public virtual
```

Initialize the APK, the payment split between quorums, and the quorum strategies + multipliers.

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

Called by the whitelister, adds a list of operators to the whitelist

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operators | address[] | the operators to add to the whitelist |

### removeFromWhitelist

```solidity
function removeFromWhitelist(address[] operators) external
```

Called by the whitelister, removes a list of operators to the whitelist

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operators | address[] | the operators to remove from the whitelist |

### registerOperator

```solidity
function registerOperator(uint8 operatorType, struct BN254.G1Point pk, string socket) external virtual
```

called for registering as an operator

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operatorType | uint8 | specifies whether the operator want to register as staker for one or both quorums |
| pk | struct BN254.G1Point | is the operator's G1 public key |
| socket | string | is the socket address of the operator |

### _registerOperator

```solidity
function _registerOperator(address operator, uint8 operatorType, struct BN254.G1Point pk, string socket) internal
```

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | is the node who is registering to be a operator |
| operatorType | uint8 | specifies whether the operator want to register as staker for one or both quorums |
| pk | struct BN254.G1Point | is the operator's G1 public key |
| socket | string | is the socket address of the operator |

### deregisterOperator

```solidity
function deregisterOperator(struct BN254.G1Point pkToRemove, uint32 index) external virtual returns (bool)
```

Used by an operator to de-register itself from providing service to the middleware.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pkToRemove | struct BN254.G1Point | is the sender's pubkey in affine coordinates |
| index | uint32 | is the sender's location in the dynamic array `operatorList` |

### _deregisterOperator

```solidity
function _deregisterOperator(address operator, struct BN254.G1Point pkToRemove, uint32 index) internal
```

Used to process de-registering an operator from providing service to the middleware.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | The operator to be deregistered |
| pkToRemove | struct BN254.G1Point | is the sender's pubkey |
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

### _processApkUpdate

```solidity
function _processApkUpdate(struct BN254.G1Point newApk) internal returns (bytes32)
```

Updates the stored APK to `newApk`, calculates its hash, and pushes new entries to the `_apkUpdates` array

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| newApk | struct BN254.G1Point | The updated APK. This will be the `apk` after this function runs! |

### _setOperatorWhitelister

```solidity
function _setOperatorWhitelister(address _operatorWhitelister) internal
```

### getCorrectApkHash

```solidity
function getCorrectApkHash(uint256 index, uint32 blockNumber) external view returns (bytes32)
```

get hash of a historical aggregated public key corresponding to a given index;
called by checkSignatures in BLSSignatureChecker.sol.

### getApkUpdatesLength

```solidity
function getApkUpdatesLength() external view returns (uint256)
```

returns the total number of APK updates that have ever occurred (including one for initializing the pubkey as the generator)

### apkUpdates

```solidity
function apkUpdates(uint256 index) external view returns (struct IBLSRegistry.ApkUpdate)
```

returns the `ApkUpdate` struct at `index` in the list of APK updates

### apkHashes

```solidity
function apkHashes(uint256 index) external view returns (bytes32)
```

returns the APK hash that resulted from the `index`th APK update

### apkUpdateBlockNumbers

```solidity
function apkUpdateBlockNumbers(uint256 index) external view returns (uint32)
```

returns the block number at which the `index`th APK update occurred

