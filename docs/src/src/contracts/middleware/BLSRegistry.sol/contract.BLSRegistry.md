# BLSRegistry
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/middleware/BLSRegistry.sol)

**Inherits:**
[RegistryBase](/src/contracts/middleware/RegistryBase.sol/abstract.RegistryBase.md), [IBLSRegistry](/src/contracts/interfaces/IBLSRegistry.sol/interface.IBLSRegistry.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

This contract is used for
- registering new operators
- committing to and finalizing de-registration as an operator
- updating the stakes of the operator


## State Variables
### ZERO_PK_HASH

```solidity
bytes32 internal constant ZERO_PK_HASH = hex"012893657d8eb2efad4de0a91bcd0e39ad9837745dec3ea923737ea803fc8e3d";
```


### pubkeyCompendium
contract used for looking up operators' BLS public keys


```solidity
IBLSPublicKeyCompendium public immutable pubkeyCompendium;
```


### _apkUpdates
list of keccak256(apk_x, apk_y) of operators, and the block numbers at which the aggregate
pubkeys were updated. This occurs whenever a new operator registers or deregisters.


```solidity
ApkUpdate[] internal _apkUpdates;
```


### apk
used for storing current aggregate public key

*Initialized value of APK is the point at infinity: (0, 0)*


```solidity
BN254.G1Point public apk;
```


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
constructor(
    IStrategyManager _strategyManager,
    IServiceManager _serviceManager,
    uint8 _NUMBER_OF_QUORUMS,
    IBLSPublicKeyCompendium _pubkeyCompendium
) RegistryBase(_strategyManager, _serviceManager, _NUMBER_OF_QUORUMS);
```

### initialize

Initialize the APK, the payment split between quorums, and the quorum strategies + multipliers.


```solidity
function initialize(
    address _operatorWhitelister,
    bool _operatorWhitelistEnabled,
    uint256[] memory _quorumBips,
    StrategyAndWeightingMultiplier[] memory _firstQuorumStrategiesConsideredAndMultipliers,
    StrategyAndWeightingMultiplier[] memory _secondQuorumStrategiesConsideredAndMultipliers
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

Called by the whitelister, adds a list of operators to the whitelist


```solidity
function addToOperatorWhitelist(address[] calldata operators) external onlyOperatorWhitelister;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operators`|`address[]`|the operators to add to the whitelist|


### removeFromWhitelist

Called by the whitelister, removes a list of operators to the whitelist


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
function registerOperator(uint8 operatorType, BN254.G1Point memory pk, string calldata socket) external virtual;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operatorType`|`uint8`|specifies whether the operator want to register as staker for one or both quorums|
|`pk`|`BN254.G1Point`|is the operator's G1 public key|
|`socket`|`string`|is the socket address of the operator|


### _registerOperator


```solidity
function _registerOperator(address operator, uint8 operatorType, BN254.G1Point memory pk, string calldata socket)
    internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the node who is registering to be a operator|
|`operatorType`|`uint8`|specifies whether the operator want to register as staker for one or both quorums|
|`pk`|`BN254.G1Point`|is the operator's G1 public key|
|`socket`|`string`|is the socket address of the operator|


### deregisterOperator

Used by an operator to de-register itself from providing service to the middleware.


```solidity
function deregisterOperator(BN254.G1Point memory pkToRemove, uint32 index) external virtual returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`pkToRemove`|`BN254.G1Point`|is the sender's pubkey in affine coordinates|
|`index`|`uint32`|is the sender's location in the dynamic array `operatorList`|


### _deregisterOperator

Used to process de-registering an operator from providing service to the middleware.


```solidity
function _deregisterOperator(address operator, BN254.G1Point memory pkToRemove, uint32 index) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|The operator to be deregistered|
|`pkToRemove`|`BN254.G1Point`|is the sender's pubkey|
|`index`|`uint32`|is the sender's location in the dynamic array `operatorList`|


### updateStakes

Used for updating information on deposits of nodes.

*Fetch operator's stored pubkeyHash*

*Verify that the stored pubkeyHash matches the 'pubkeyToRemoveAff' input*


```solidity
function updateStakes(address[] calldata operators, uint256[] calldata prevElements) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operators`|`address[]`|are the nodes whose deposit information is getting updated|
|`prevElements`|`uint256[]`|are the elements before this middleware in the operator's linked list within the slasher|


### _processApkUpdate

Updates the stored APK to `newApk`, calculates its hash, and pushes new entries to the `_apkUpdates` array


```solidity
function _processApkUpdate(BN254.G1Point memory newApk) internal returns (bytes32);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newApk`|`BN254.G1Point`|The updated APK. This will be the `apk` after this function runs!|


### _setOperatorWhitelister


```solidity
function _setOperatorWhitelister(address _operatorWhitelister) internal;
```

### getCorrectApkHash

get hash of a historical aggregated public key corresponding to a given index;
called by checkSignatures in BLSSignatureChecker.sol.


```solidity
function getCorrectApkHash(uint256 index, uint32 blockNumber) external view returns (bytes32);
```

### getApkUpdatesLength

returns the total number of APK updates that have ever occurred (including one for initializing the pubkey as the generator)


```solidity
function getApkUpdatesLength() external view returns (uint256);
```

### apkUpdates

returns the `ApkUpdate` struct at `index` in the list of APK updates


```solidity
function apkUpdates(uint256 index) external view returns (ApkUpdate memory);
```

### apkHashes

returns the APK hash that resulted from the `index`th APK update


```solidity
function apkHashes(uint256 index) external view returns (bytes32);
```

### apkUpdateBlockNumbers

returns the block number at which the `index`th APK update occurred


```solidity
function apkUpdateBlockNumbers(uint256 index) external view returns (uint32);
```

## Events
### Registration
Emitted upon the registration of a new operator for the middleware


```solidity
event Registration(
    address indexed operator, bytes32 pkHash, BN254.G1Point pk, uint32 apkHashIndex, bytes32 apkHash, string socket
);
```

### OperatorWhitelisterTransferred
Emitted when the `operatorWhitelister` role is transferred.


```solidity
event OperatorWhitelisterTransferred(address previousAddress, address newAddress);
```

