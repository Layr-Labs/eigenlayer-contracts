# BeaconChainOracleUnitTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/unit/BeaconChainOracleUnit.t.sol)

**Inherits:**
Test


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### beaconChainOracle

```solidity
BeaconChainOracle public beaconChainOracle;
```


### initialBeaconChainOwner

```solidity
address public initialBeaconChainOwner = address(this);
```


### initialBeaconChainOracleThreshold

```solidity
uint256 public initialBeaconChainOracleThreshold = 2;
```


### minThreshold

```solidity
uint256 public minThreshold;
```


### addressIsExcludedFromFuzzedInputs

```solidity
mapping(address => bool) public addressIsExcludedFromFuzzedInputs;
```


### numberPotentialOracleSigners

```solidity
uint256 numberPotentialOracleSigners = 16;
```


### potentialOracleSigners

```solidity
address[] public potentialOracleSigners;
```


### blockNumberToVoteFor

```solidity
uint64 public blockNumberToVoteFor = 5151;
```


### stateRootToVoteFor

```solidity
bytes32 public stateRootToVoteFor = bytes32(uint256(987));
```


## Functions
### filterFuzzedAddressInputs


```solidity
modifier filterFuzzedAddressInputs(address fuzzedAddress);
```

### setUp


```solidity
function setUp() external;
```

### testConstructor_RevertsOnThresholdTooLow


```solidity
function testConstructor_RevertsOnThresholdTooLow() external;
```

### testSetThreshold


```solidity
function testSetThreshold(uint256 newThreshold) public;
```

### testSetThreshold_RevertsOnThresholdTooLow


```solidity
function testSetThreshold_RevertsOnThresholdTooLow() external;
```

### testSetThreshold_RevertsOnCallingFromNotOwner


```solidity
function testSetThreshold_RevertsOnCallingFromNotOwner(address notOwner) external;
```

### testAddOracleSigner


```solidity
function testAddOracleSigner(address signerToAdd) public;
```

### testAddOracleSigners


```solidity
function testAddOracleSigners(uint8 amountSignersToAdd) public;
```

### testAddOracleSigners_SignerAlreadyInSet


```solidity
function testAddOracleSigners_SignerAlreadyInSet() external;
```

### testAddOracleSigners_RevertsOnCallingFromNotOwner


```solidity
function testAddOracleSigners_RevertsOnCallingFromNotOwner(address notOwner) external;
```

### testRemoveOracleSigner


```solidity
function testRemoveOracleSigner(address signerToRemove) public;
```

### testRemoveOracleSigners


```solidity
function testRemoveOracleSigners(uint8 amountSignersToAdd, uint8 amountSignersToRemove) external;
```

### testRemoveOracleSigners_SignerAlreadyNotInSet


```solidity
function testRemoveOracleSigners_SignerAlreadyNotInSet() external;
```

### testRemoveOracleSigners_RevertsOnCallingFromNotOwner


```solidity
function testRemoveOracleSigners_RevertsOnCallingFromNotOwner(address notOwner) external;
```

### testVoteForBeaconChainStateRoot


```solidity
function testVoteForBeaconChainStateRoot(address oracleSigner, uint64 _blockNumber, bytes32 _stateRoot) public;
```

### testVoteForBeaconChainStateRoot_VoteDoesNotCauseConfirmation


```solidity
function testVoteForBeaconChainStateRoot_VoteDoesNotCauseConfirmation() public;
```

### testVoteForBeaconChainStateRoot_VoteCausesConfirmation


```solidity
function testVoteForBeaconChainStateRoot_VoteCausesConfirmation(uint64 _blockNumber, bytes32 _stateRoot) public;
```

### testVoteForBeaconChainStateRoot_VoteCausesConfirmation_latestOracleBlockNumberDoesNotIncrease


```solidity
function testVoteForBeaconChainStateRoot_VoteCausesConfirmation_latestOracleBlockNumberDoesNotIncrease() external;
```

### testVoteForBeaconChainStateRoot_RevertsWhenCallerHasVoted


```solidity
function testVoteForBeaconChainStateRoot_RevertsWhenCallerHasVoted() external;
```

### testVoteForBeaconChainStateRoot_RevertsWhenStateRootAlreadyConfirmed


```solidity
function testVoteForBeaconChainStateRoot_RevertsWhenStateRootAlreadyConfirmed() external;
```

### testVoteForBeaconChainStateRoot_RevertsWhenCallingFromNotOracleSigner


```solidity
function testVoteForBeaconChainStateRoot_RevertsWhenCallingFromNotOracleSigner(address notOracleSigner) external;
```

