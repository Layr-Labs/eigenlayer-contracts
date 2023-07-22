# EigenLayerDeployer
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/EigenLayerDeployer.t.sol)

**Inherits:**
[Operators](/src/test/utils/Operators.sol/contract.Operators.md)


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### eigenLayerProxyAdmin

```solidity
ProxyAdmin public eigenLayerProxyAdmin;
```


### eigenLayerPauserReg

```solidity
PauserRegistry public eigenLayerPauserReg;
```


### slasher

```solidity
Slasher public slasher;
```


### delegation

```solidity
DelegationManager public delegation;
```


### strategyManager

```solidity
StrategyManager public strategyManager;
```


### eigenPodManager

```solidity
EigenPodManager public eigenPodManager;
```


### pod

```solidity
IEigenPod public pod;
```


### delayedWithdrawalRouter

```solidity
IDelayedWithdrawalRouter public delayedWithdrawalRouter;
```


### ethPOSDeposit

```solidity
IETHPOSDeposit public ethPOSDeposit;
```


### eigenPodBeacon

```solidity
IBeacon public eigenPodBeacon;
```


### beaconChainOracle

```solidity
IBeaconChainOracle public beaconChainOracle;
```


### eigenToken

```solidity
IERC20 public eigenToken;
```


### weth

```solidity
IERC20 public weth;
```


### wethStrat

```solidity
StrategyBase public wethStrat;
```


### eigenStrat

```solidity
StrategyBase public eigenStrat;
```


### baseStrategyImplementation

```solidity
StrategyBase public baseStrategyImplementation;
```


### emptyContract

```solidity
EmptyContract public emptyContract;
```


### strategies

```solidity
mapping(uint256 => IStrategy) public strategies;
```


### priv_key_0

```solidity
bytes32 priv_key_0 = 0x1234567812345678123456781234567812345678123456781234567812345678;
```


### priv_key_1

```solidity
bytes32 priv_key_1 = 0x1234567812345678123456781234567812345698123456781234567812348976;
```


### strategyIndexes

```solidity
uint256[] public strategyIndexes;
```


### stakers

```solidity
address[2] public stakers;
```


### sample_registrant

```solidity
address sample_registrant = cheats.addr(436364636);
```


### slashingContracts

```solidity
address[] public slashingContracts;
```


### wethInitialSupply

```solidity
uint256 wethInitialSupply = 10e50;
```


### eigenTotalSupply

```solidity
uint256 public constant eigenTotalSupply = 1000e18;
```


### nonce

```solidity
uint256 nonce = 69;
```


### gasLimit

```solidity
uint256 public gasLimit = 750000;
```


### PARTIAL_WITHDRAWAL_FRAUD_PROOF_PERIOD_BLOCKS

```solidity
uint32 PARTIAL_WITHDRAWAL_FRAUD_PROOF_PERIOD_BLOCKS = 7 days / 12 seconds;
```


### REQUIRED_BALANCE_WEI

```solidity
uint256 REQUIRED_BALANCE_WEI = 31 ether;
```


### MAX_PARTIAL_WTIHDRAWAL_AMOUNT_GWEI

```solidity
uint64 MAX_PARTIAL_WTIHDRAWAL_AMOUNT_GWEI = 1 ether / 1e9;
```


### pauser

```solidity
address pauser;
```


### unpauser

```solidity
address unpauser;
```


### theMultiSig

```solidity
address theMultiSig = address(420);
```


### operator

```solidity
address operator = address(0x4206904396bF2f8b173350ADdEc5007A52664293);
```


### acct_0

```solidity
address acct_0 = cheats.addr(uint256(priv_key_0));
```


### acct_1

```solidity
address acct_1 = cheats.addr(uint256(priv_key_1));
```


### _challenger

```solidity
address _challenger = address(0x6966904396bF2f8b173350bCcec5007A52669873);
```


### eigenLayerReputedMultisig

```solidity
address public eigenLayerReputedMultisig = address(this);
```


### eigenLayerProxyAdminAddress

```solidity
address eigenLayerProxyAdminAddress;
```


### eigenLayerPauserRegAddress

```solidity
address eigenLayerPauserRegAddress;
```


### slasherAddress

```solidity
address slasherAddress;
```


### delegationAddress

```solidity
address delegationAddress;
```


### strategyManagerAddress

```solidity
address strategyManagerAddress;
```


### eigenPodManagerAddress

```solidity
address eigenPodManagerAddress;
```


### podAddress

```solidity
address podAddress;
```


### delayedWithdrawalRouterAddress

```solidity
address delayedWithdrawalRouterAddress;
```


### eigenPodBeaconAddress

```solidity
address eigenPodBeaconAddress;
```


### beaconChainOracleAddress

```solidity
address beaconChainOracleAddress;
```


### emptyContractAddress

```solidity
address emptyContractAddress;
```


### operationsMultisig

```solidity
address operationsMultisig;
```


### executorMultisig

```solidity
address executorMultisig;
```


### initialBeaconChainOracleThreshold

```solidity
uint256 public initialBeaconChainOracleThreshold = 3;
```


### goerliDeploymentConfig

```solidity
string internal goerliDeploymentConfig = vm.readFile("script/output/M1_deployment_goerli_2023_3_23.json");
```


### fuzzedAddressMapping

```solidity
mapping(address => bool) fuzzedAddressMapping;
```


## Functions
### fuzzedAddress


```solidity
modifier fuzzedAddress(address addr) virtual;
```

### cannotReinit


```solidity
modifier cannotReinit();
```

### setUp


```solidity
function setUp() public virtual;
```

### _deployEigenLayerContractsGoerli


```solidity
function _deployEigenLayerContractsGoerli() internal;
```

### _deployEigenLayerContractsLocal


```solidity
function _deployEigenLayerContractsLocal() internal;
```

### _setAddresses

First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.


```solidity
function _setAddresses(string memory config) internal;
```

