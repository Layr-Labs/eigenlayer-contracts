# PaymentManager
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/middleware/PaymentManager.sol)

**Inherits:**
Initializable, [IPaymentManager](/docs/docgen/src/src/contracts/interfaces/IPaymentManager.sol/interface.IPaymentManager.md), [Pausable](/docs/docgen/src/src/contracts/permissions/Pausable.sol/contract.Pausable.md)

**Author:**
Layr Labs, Inc.

This contract is used for doing interactive payment challenges.

The contract is marked as abstract since it does not implement the `respondToPaymentChallengeFinal`
function -- see DataLayerPaymentManager for an example


## State Variables
### PAUSED_NEW_PAYMENT_COMMIT

```solidity
uint8 internal constant PAUSED_NEW_PAYMENT_COMMIT = 0;
```


### PAUSED_REDEEM_PAYMENT

```solidity
uint8 internal constant PAUSED_REDEEM_PAYMENT = 1;
```


### paymentFraudproofInterval
Challenge window for submitting fraudproof in the case of an incorrect payment claim by a registered operator.


```solidity
uint256 public constant paymentFraudproofInterval = 7 days;
```


### MAX_BIPS
Constant used as a divisor in dealing with BIPS amounts


```solidity
uint256 internal constant MAX_BIPS = 10000;
```


### LOW_LEVEL_GAS_BUDGET
Gas budget provided in calls to DelegationTerms contracts


```solidity
uint256 internal constant LOW_LEVEL_GAS_BUDGET = 1e5;
```


### delegationManager
The global EigenLayer Delegation contract, which is primarily used by
stakers to delegate their stake to operators who serve as middleware nodes.

*For more details, see DelegationManager.sol.*


```solidity
IDelegationManager public immutable delegationManager;
```


### serviceManager
The ServiceManager contract for this middleware, where tasks are created / initiated.


```solidity
IServiceManager public immutable serviceManager;
```


### registry
The Registry contract for this middleware, where operators register and deregister.


```solidity
IQuorumRegistry public immutable registry;
```


### paymentToken
the ERC20 token that will be used by the disperser to pay the service fees to middleware nodes.


```solidity
IERC20 public immutable paymentToken;
```


### paymentChallengeToken
Token used for placing a guarantee on challenges & payment commits


```solidity
IERC20 public immutable paymentChallengeToken;
```


### paymentChallengeAmount
Specifies the payment that has to be made as a guarantee for fraudproof during payment challenges.


```solidity
uint256 public paymentChallengeAmount;
```


### operatorToPayment
mapping between the operator and its current committed payment or last redeemed payment


```solidity
mapping(address => Payment) public operatorToPayment;
```


### operatorToPaymentChallenge
mapping from operator => PaymentChallenge


```solidity
mapping(address => PaymentChallenge) public operatorToPaymentChallenge;
```


### depositsOf
Deposits of future fees to be drawn against when paying for service from the middleware


```solidity
mapping(address => uint256) public depositsOf;
```


### allowances
depositors => addresses approved to spend deposits => allowance


```solidity
mapping(address => mapping(address => uint256)) public allowances;
```


## Functions
### onlyServiceManager

when applied to a function, ensures that the function is only callable by the `serviceManager`


```solidity
modifier onlyServiceManager();
```

### onlyRegistry

when applied to a function, ensures that the function is only callable by the `registry`


```solidity
modifier onlyRegistry();
```

### onlyServiceManagerOwner

when applied to a function, ensures that the function is only callable by the owner of the `serviceManager`


```solidity
modifier onlyServiceManagerOwner();
```

### constructor


```solidity
constructor(
    IDelegationManager _delegationManager,
    IServiceManager _serviceManager,
    IQuorumRegistry _registry,
    IERC20 _paymentToken,
    IERC20 _paymentChallengeToken
);
```

### initialize


```solidity
function initialize(IPauserRegistry _pauserReg, uint256 _paymentChallengeAmount) public initializer;
```

### depositFutureFees

deposit one-time fees by the `msg.sender` with this contract to pay for future tasks of this middleware


```solidity
function depositFutureFees(address depositFor, uint256 amount) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositFor`|`address`|could be the `msg.sender` themselves, or a different address for whom `msg.sender` is depositing these future fees|
|`amount`|`uint256`|is amount of futures fees being deposited|


### setAllowance

Allows the `allowed` address to spend up to `amount` of the `msg.sender`'s funds that have been deposited in this contract


```solidity
function setAllowance(address allowed, uint256 amount) external;
```

### setPaymentChallengeAmount

Modifies the `paymentChallengeAmount` amount.


```solidity
function setPaymentChallengeAmount(uint256 _paymentChallengeAmount) external virtual onlyServiceManagerOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_paymentChallengeAmount`|`uint256`|The new value for `paymentChallengeAmount` to take.|


### takeFee

Used for deducting the fees from the payer to the middleware


```solidity
function takeFee(address initiator, address payer, uint256 feeAmount) external virtual onlyServiceManager;
```

### commitPayment

This is used by an operator to make a claim on the amount that they deserve for their service from their last payment until `toTaskNumber`

*Once this payment is recorded, a fraud proof period commences during which a challenger can dispute the proposed payment.*


```solidity
function commitPayment(uint32 toTaskNumber, uint96 amount) external onlyWhenNotPaused(PAUSED_NEW_PAYMENT_COMMIT);
```

### redeemPayment

For the special case of this being the first payment that is being claimed by the operator,
the operator must be claiming payment starting from when they registered.

Called by an operator to redeem a payment that they previously 'committed' to by calling `commitPayment`.

*This function can only be called after the challenge window for the payment claim has completed.*


```solidity
function redeemPayment() external onlyWhenNotPaused(PAUSED_REDEEM_PAYMENT);
```

### _payForServiceHook


```solidity
function _payForServiceHook(IDelegationTerms dt, uint256 amount) internal;
```

### initPaymentChallenge

We use low-level call functionality here to ensure that an operator cannot maliciously make this function fail in order to prevent undelegation.
In particular, in-line assembly is also used to prevent the copying of uncapped return data which is also a potential DoS vector.

This function is called by a fraud prover to challenge a payment, initiating an interactive-type fraudproof.


```solidity
function initPaymentChallenge(address operator, uint96 amount1, uint96 amount2) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator against whose payment claim the fraudproof is being made|
|`amount1`|`uint96`|is the reward amount the challenger in that round claims is for the first half of tasks|
|`amount2`|`uint96`|is the reward amount the challenger in that round claims is for the second half of tasks|


### performChallengeBisectionStep

Perform a single bisection step in an existing interactive payment challenge.


```solidity
function performChallengeBisectionStep(address operator, bool secondHalf, uint96 amount1, uint96 amount2) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|The middleware operator who was challenged (used to look up challenge details)|
|`secondHalf`|`bool`|If true, then the caller wishes to challenge the amount claimed as payment in the *second half* of the previous bisection step. If false then the *first half* is indicated instead.|
|`amount1`|`uint96`|The amount that the caller asserts the operator is entitled to, for the first half *of the challenged half* of the previous bisection.|
|`amount2`|`uint96`|The amount that the caller asserts the operator is entitled to, for the second half *of the challenged half* of the previous bisection.|


### _updateStatus

Change the challenged interval to the one the challenger cares about.
If the difference between the current start and end is even, then the new interval has an endpoint halfway in-between
If the difference is odd = 2n + 1, the new interval has a "from" endpoint at (start + n = end - (n + 1)) if the second half is challenged,
or a "to" endpoint at (end - (2n + 2)/2 = end - (n + 1) = start + n) if the first half is challenged
In other words, it's simple when the difference is even, and when the difference is odd, we just always make the first half the smaller one.

This function is used for updating the status of the challenge in terms of who has to respon
to the interactive challenge mechanism next -  is it going to be challenger or the operator.

*If the challenge is over only one task, then the challenge is marked specially as a one step challenge –
the smallest unit over which a challenge can be proposed – and 'true' is returned.
Otherwise status is updated normally and 'false' is returned.*


```solidity
function _updateStatus(address operator, uint32 diff) internal returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator whose payment claim is being challenged|
|`diff`|`uint32`|is the number of tasks across which payment is being challenged in this iteration|


### _updateChallengeAmounts

Used to update challenge amounts when the operator (or challenger) breaks down the challenged amount (single bisection step)


```solidity
function _updateChallengeAmounts(address operator, DissectionType dissectionType, uint96 amount1, uint96 amount2)
    internal;
```

### resolveChallenge

resolve an existing PaymentChallenge for an operator


```solidity
function resolveChallenge(address operator) external;
```

### _resolve

Resolves a single payment challenge, paying the winner.

*If challenger is proven correct, then they are refunded their own challengeAmount plus the challengeAmount put up by the operator.
If operator is proven correct, then the challenger's challengeAmount is transferred to them, since the operator still hasn't been
proven right, and thus their challengeAmount is still required in case they are challenged again.*


```solidity
function _resolve(PaymentChallenge memory challenge, address winner) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`challenge`|`PaymentChallenge`|The challenge that is being resolved.|
|`winner`|`address`|Address of the winner of the challenge.|


### getChallengeStatus

Returns the ChallengeStatus for the `operator`'s payment claim.


```solidity
function getChallengeStatus(address operator) external view returns (ChallengeStatus);
```

### getAmount1

Returns the 'amount1' for the `operator`'s payment claim.


```solidity
function getAmount1(address operator) external view returns (uint96);
```

### getAmount2

Returns the 'amount2' for the `operator`'s payment claim.


```solidity
function getAmount2(address operator) external view returns (uint96);
```

### getToTaskNumber

Returns the 'toTaskNumber' for the `operator`'s payment claim.


```solidity
function getToTaskNumber(address operator) external view returns (uint48);
```

### getFromTaskNumber

Returns the 'fromTaskNumber' for the `operator`'s payment claim.


```solidity
function getFromTaskNumber(address operator) external view returns (uint48);
```

### getDiff

Returns the task number difference for the `operator`'s payment claim.


```solidity
function getDiff(address operator) external view returns (uint48);
```

### getPaymentChallengeAmount

Returns the active challengeAmount of the `operator` placed on their payment claim.


```solidity
function getPaymentChallengeAmount(address operator) external view returns (uint256);
```

### _taskNumber

Convenience function for fetching the current taskNumber from the `serviceManager`


```solidity
function _taskNumber() internal view returns (uint32);
```

### _setPaymentChallengeAmount

Modifies the `paymentChallengeAmount` amount.


```solidity
function _setPaymentChallengeAmount(uint256 _paymentChallengeAmount) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_paymentChallengeAmount`|`uint256`|The new value for `paymentChallengeAmount` to take.|


## Events
### PaymentChallengeAmountSet
Emitted when the `paymentChallengeAmount` variable is modified


```solidity
event PaymentChallengeAmountSet(uint256 previousValue, uint256 newValue);
```

### PaymentCommit
Emitted when an operator commits to a payment by calling the `commitPayment` function


```solidity
event PaymentCommit(address operator, uint32 fromTaskNumber, uint32 toTaskNumber, uint256 fee);
```

### PaymentChallengeInit
Emitted when a new challenge is created through a call to the `initPaymentChallenge` function


```solidity
event PaymentChallengeInit(address indexed operator, address challenger);
```

### PaymentRedemption
Emitted when an operator redeems a payment by calling the `redeemPayment` function


```solidity
event PaymentRedemption(address indexed operator, uint256 fee);
```

### PaymentBreakdown
Emitted when a bisection step is performed in a challenge, through a call to the `performChallengeBisectionStep` function


```solidity
event PaymentBreakdown(
    address indexed operator, uint32 fromTaskNumber, uint32 toTaskNumber, uint96 amount1, uint96 amount2
);
```

### PaymentChallengeResolution
Emitted upon successful resolution of a payment challenge, within a call to `resolveChallenge`


```solidity
event PaymentChallengeResolution(address indexed operator, bool operatorWon);
```

### OnPayForServiceCallFailure
*Emitted when a low-level call to `delegationTerms.payForService` fails, returning `returnData`*


```solidity
event OnPayForServiceCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);
```

