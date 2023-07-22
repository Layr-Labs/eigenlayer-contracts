# EigenLayerTestHelper
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/EigenLayerTestHelper.t.sol)

**Inherits:**
[EigenLayerDeployer](/src/test/EigenLayerDeployer.t.sol/contract.EigenLayerDeployer.md)


## State Variables
### durationToInit

```solidity
uint8 durationToInit = 2;
```


### SECP256K1N_MODULUS

```solidity
uint256 public SECP256K1N_MODULUS = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141;
```


### SECP256K1N_MODULUS_HALF

```solidity
uint256 public SECP256K1N_MODULUS_HALF = 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0;
```


### sharesBefore

```solidity
uint256[] sharesBefore;
```


### balanceBefore

```solidity
uint256[] balanceBefore;
```


### priorTotalShares

```solidity
uint256[] priorTotalShares;
```


### strategyTokenBalance

```solidity
uint256[] strategyTokenBalance;
```


## Functions
### _testInitiateDelegation

Helper function to test `initiateDelegation` functionality.  Handles registering as an operator, depositing tokens
into both Weth and Eigen strategies, as well as delegating assets from "stakers" to the operator.


```solidity
function _testInitiateDelegation(uint8 operatorIndex, uint256 amountEigenToDeposit, uint256 amountEthToDeposit)
    public
    returns (uint256 amountEthStaked, uint256 amountEigenStaked);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operatorIndex`|`uint8`|is the index of the operator to use from the test-data/operators.json file|
|`amountEigenToDeposit`|`uint256`|amount of eigen token to deposit|
|`amountEthToDeposit`|`uint256`|amount of eth to deposit|


### _testRegisterAsOperator

Register 'sender' as an operator, setting their 'DelegationTerms' contract in DelegationManager to 'dt', verifies
that the storage of DelegationManager contract is updated appropriately


```solidity
function _testRegisterAsOperator(address sender, IDelegationTerms dt) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|is the address being registered as an operator|
|`dt`|`IDelegationTerms`|is the sender's DelegationTerms contract|


### _testDepositWeth

Deposits `amountToDeposit` of WETH from address `sender` into `wethStrat`.


```solidity
function _testDepositWeth(address sender, uint256 amountToDeposit) internal returns (uint256 amountDeposited);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address to spoof calls from using `cheats.startPrank(sender)`|
|`amountToDeposit`|`uint256`|Amount of WETH that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`|


### _testDepositEigen

Deposits `amountToDeposit` of EIGEN from address `sender` into `eigenStrat`.


```solidity
function _testDepositEigen(address sender, uint256 amountToDeposit) internal returns (uint256 amountDeposited);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address to spoof calls from using `cheats.startPrank(sender)`|
|`amountToDeposit`|`uint256`|Amount of EIGEN that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`|


### _testDepositToStrategy

Deposits `amountToDeposit` of `underlyingToken` from address `sender` into `stratToDepositTo`.
If*  `sender` has zero shares prior to deposit, *then* checks that `stratToDepositTo` is correctly added to their `stakerStrategyList` array.


```solidity
function _testDepositToStrategy(
    address sender,
    uint256 amountToDeposit,
    IERC20 underlyingToken,
    IStrategy stratToDepositTo
) internal returns (uint256 amountDeposited);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address to spoof calls from using `cheats.startPrank(sender)`|
|`amountToDeposit`|`uint256`|Amount of WETH that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`|
|`underlyingToken`|`IERC20`||
|`stratToDepositTo`|`IStrategy`||


### _testDelegateToOperator

tries to delegate from 'staker' to 'operator', verifies that staker has at least some shares
delegatedShares update correctly for 'operator' and delegated status is updated correctly for 'staker'


```solidity
function _testDelegateToOperator(address staker, address operator) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`staker`|`address`|the staker address to delegate from|
|`operator`|`address`|the operator address to delegate to|


### _testDepositStrategies

deploys 'numStratsToAdd' strategies contracts and initializes them to treat `underlyingToken` as their underlying token
and then deposits 'amountToDeposit' to each of them from 'sender'


```solidity
function _testDepositStrategies(address sender, uint256 amountToDeposit, uint8 numStratsToAdd) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|address that is depositing into the strategies|
|`amountToDeposit`|`uint256`|amount being deposited|
|`numStratsToAdd`|`uint8`|number of strategies that are being deployed and deposited into|


### _createQueuedWithdrawal

Creates a queued withdrawal from `staker`. Begins by registering the staker as a delegate (if specified), then deposits `amountToDeposit`
into the WETH strategy, and then queues a withdrawal using
`strategyManager.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, withdrawer)`

After initiating a queued withdrawal, this test checks that `strategyManager.canCompleteQueuedWithdrawal` immediately returns the correct
response depending on whether `staker` is delegated or not.


```solidity
function _createQueuedWithdrawal(
    address staker,
    bool registerAsOperator,
    uint256 amountToDeposit,
    IStrategy[] memory strategyArray,
    uint256[] memory shareAmounts,
    uint256[] memory strategyIndexes,
    address withdrawer
) internal returns (bytes32 withdrawalRoot, IStrategyManager.QueuedWithdrawal memory queuedWithdrawal);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`staker`|`address`|The address to initiate the queued withdrawal|
|`registerAsOperator`|`bool`|If true, `staker` will also register as a delegate in the course of this function|
|`amountToDeposit`|`uint256`|The amount of WETH to deposit|
|`strategyArray`|`IStrategy[]`||
|`shareAmounts`|`uint256[]`||
|`strategyIndexes`|`uint256[]`||
|`withdrawer`|`address`||


### getVSfromVandS

Helper for ECDSA signatures: combines V and S into VS - if S is greater than SECP256K1N_MODULUS_HALF, then we
get the modulus, so that the leading bit of s is always 0.  Then we set the leading
bit to be either 0 or 1 based on the value of v, which is either 27 or 28


```solidity
function getVSfromVandS(uint8 v, bytes32 s) internal view returns (bytes32);
```

### _testDelegation

registers a fixed address as an operator, delegates to it from a second address,
and checks that the operator's voteWeights increase properly


```solidity
function _testDelegation(
    address operator,
    address staker,
    uint256 ethAmount,
    uint256 eigenAmount,
    IVoteWeigher voteWeigher
) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator being delegated to.|
|`staker`|`address`|is the staker delegating stake to the operator.|
|`ethAmount`|`uint256`||
|`eigenAmount`|`uint256`||
|`voteWeigher`|`IVoteWeigher`|is the VoteWeigher-type contract to consult for stake weight changes|


### _testCompleteQueuedWithdrawalShares

Helper function to complete an existing queued withdrawal in shares


```solidity
function _testCompleteQueuedWithdrawalShares(
    address depositor,
    IStrategy[] memory strategyArray,
    IERC20[] memory tokensArray,
    uint256[] memory shareAmounts,
    address delegatedTo,
    IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce,
    uint32 withdrawalStartBlock,
    uint256 middlewareTimesIndex
) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|is the address of the staker who queued a withdrawal|
|`strategyArray`|`IStrategy[]`|is the array of strategies to withdraw from|
|`tokensArray`|`IERC20[]`|is the array of tokens to withdraw from said strategies|
|`shareAmounts`|`uint256[]`|is the array of shares to be withdrawn from said strategies|
|`delegatedTo`|`address`|is the address the staker has delegated their shares to|
|`withdrawerAndNonce`|`IStrategyManager.WithdrawerAndNonce`|is a struct containing the withdrawer and the nonce of the withdrawal|
|`withdrawalStartBlock`|`uint32`|the block number of the original queued withdrawal|
|`middlewareTimesIndex`|`uint256`|index in the middlewareTimes array used to queue this withdrawal|


### _testCompleteQueuedWithdrawalTokens

Helper function to complete an existing queued withdrawal in tokens


```solidity
function _testCompleteQueuedWithdrawalTokens(
    address depositor,
    IStrategy[] memory strategyArray,
    IERC20[] memory tokensArray,
    uint256[] memory shareAmounts,
    address delegatedTo,
    IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce,
    uint32 withdrawalStartBlock,
    uint256 middlewareTimesIndex
) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|is the address of the staker who queued a withdrawal|
|`strategyArray`|`IStrategy[]`|is the array of strategies to withdraw from|
|`tokensArray`|`IERC20[]`|is the array of tokens to withdraw from said strategies|
|`shareAmounts`|`uint256[]`|is the array of shares to be withdrawn from said strategies|
|`delegatedTo`|`address`|is the address the staker has delegated their shares to|
|`withdrawerAndNonce`|`IStrategyManager.WithdrawerAndNonce`|is a struct containing the withdrawer and the nonce of the withdrawal|
|`withdrawalStartBlock`|`uint32`|the block number of the original queued withdrawal|
|`middlewareTimesIndex`|`uint256`|index in the middlewareTimes array used to queue this withdrawal|


### _testQueueWithdrawal


```solidity
function _testQueueWithdrawal(
    address depositor,
    uint256[] memory strategyIndexes,
    IStrategy[] memory strategyArray,
    uint256[] memory shareAmounts,
    address withdrawer,
    bool undelegateIfPossible
) internal returns (bytes32);
```

