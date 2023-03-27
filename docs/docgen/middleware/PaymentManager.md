# Solidity API

## PaymentManager

This contract is used for doing interactive payment challenges.
The contract is marked as abstract since it does not implement the `respondToPaymentChallengeFinal`
function -- see DataLayerPaymentManager for an example

### PAUSED_NEW_PAYMENT_COMMIT

```solidity
uint8 PAUSED_NEW_PAYMENT_COMMIT
```

### PAUSED_REDEEM_PAYMENT

```solidity
uint8 PAUSED_REDEEM_PAYMENT
```

### paymentFraudproofInterval

```solidity
uint256 paymentFraudproofInterval
```

Challenge window for submitting fraudproof in the case of an incorrect payment claim by a registered operator.

### MAX_BIPS

```solidity
uint256 MAX_BIPS
```

Constant used as a divisor in dealing with BIPS amounts

### LOW_LEVEL_GAS_BUDGET

```solidity
uint256 LOW_LEVEL_GAS_BUDGET
```

Gas budget provided in calls to DelegationTerms contracts

### delegationManager

```solidity
contract IDelegationManager delegationManager
```

The global EigenLayer Delegation contract, which is primarily used by
stakers to delegate their stake to operators who serve as middleware nodes.

_For more details, see DelegationManager.sol._

### serviceManager

```solidity
contract IServiceManager serviceManager
```

The ServiceManager contract for this middleware, where tasks are created / initiated.

### registry

```solidity
contract IQuorumRegistry registry
```

The Registry contract for this middleware, where operators register and deregister.

### paymentToken

```solidity
contract IERC20 paymentToken
```

the ERC20 token that will be used by the disperser to pay the service fees to middleware nodes.

### paymentChallengeToken

```solidity
contract IERC20 paymentChallengeToken
```

Token used for placing a guarantee on challenges & payment commits

### paymentChallengeAmount

```solidity
uint256 paymentChallengeAmount
```

Specifies the payment that has to be made as a guarantee for fraudproof during payment challenges.

### operatorToPayment

```solidity
mapping(address => struct IPaymentManager.Payment) operatorToPayment
```

mapping between the operator and its current committed payment or last redeemed payment

### operatorToPaymentChallenge

```solidity
mapping(address => struct IPaymentManager.PaymentChallenge) operatorToPaymentChallenge
```

mapping from operator => PaymentChallenge

### depositsOf

```solidity
mapping(address => uint256) depositsOf
```

Deposits of future fees to be drawn against when paying for service from the middleware

### allowances

```solidity
mapping(address => mapping(address => uint256)) allowances
```

depositors => addresses approved to spend deposits => allowance

### PaymentChallengeAmountSet

```solidity
event PaymentChallengeAmountSet(uint256 previousValue, uint256 newValue)
```

Emitted when the `paymentChallengeAmount` variable is modified

### PaymentCommit

```solidity
event PaymentCommit(address operator, uint32 fromTaskNumber, uint32 toTaskNumber, uint256 fee)
```

Emitted when an operator commits to a payment by calling the `commitPayment` function

### PaymentChallengeInit

```solidity
event PaymentChallengeInit(address operator, address challenger)
```

Emitted when a new challenge is created through a call to the `initPaymentChallenge` function

### PaymentRedemption

```solidity
event PaymentRedemption(address operator, uint256 fee)
```

Emitted when an operator redeems a payment by calling the `redeemPayment` function

### PaymentBreakdown

```solidity
event PaymentBreakdown(address operator, uint32 fromTaskNumber, uint32 toTaskNumber, uint96 amount1, uint96 amount2)
```

Emitted when a bisection step is performed in a challenge, through a call to the `performChallengeBisectionStep` function

### PaymentChallengeResolution

```solidity
event PaymentChallengeResolution(address operator, bool operatorWon)
```

Emitted upon successful resolution of a payment challenge, within a call to `resolveChallenge`

### OnPayForServiceCallFailure

```solidity
event OnPayForServiceCallFailure(contract IDelegationTerms delegationTerms, bytes32 returnData)
```

_Emitted when a low-level call to `delegationTerms.payForService` fails, returning `returnData`_

### onlyServiceManager

```solidity
modifier onlyServiceManager()
```

when applied to a function, ensures that the function is only callable by the `serviceManager`

### onlyRegistry

```solidity
modifier onlyRegistry()
```

when applied to a function, ensures that the function is only callable by the `registry`

### onlyServiceManagerOwner

```solidity
modifier onlyServiceManagerOwner()
```

when applied to a function, ensures that the function is only callable by the owner of the `serviceManager`

### constructor

```solidity
constructor(contract IDelegationManager _delegationManager, contract IServiceManager _serviceManager, contract IQuorumRegistry _registry, contract IERC20 _paymentToken, contract IERC20 _paymentChallengeToken) internal
```

### initialize

```solidity
function initialize(contract IPauserRegistry _pauserReg, uint256 _paymentChallengeAmount) public
```

### depositFutureFees

```solidity
function depositFutureFees(address depositFor, uint256 amount) external
```

deposit one-time fees by the `msg.sender` with this contract to pay for future tasks of this middleware

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| depositFor | address | could be the `msg.sender` themselves, or a different address for whom `msg.sender` is depositing these future fees |
| amount | uint256 | is amount of futures fees being deposited |

### setAllowance

```solidity
function setAllowance(address allowed, uint256 amount) external
```

Allows the `allowed` address to spend up to `amount` of the `msg.sender`'s funds that have been deposited in this contract

### setPaymentChallengeAmount

```solidity
function setPaymentChallengeAmount(uint256 _paymentChallengeAmount) external virtual
```

Modifies the `paymentChallengeAmount` amount.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| _paymentChallengeAmount | uint256 | The new value for `paymentChallengeAmount` to take. |

### takeFee

```solidity
function takeFee(address initiator, address payer, uint256 feeAmount) external virtual
```

Used for deducting the fees from the payer to the middleware

### commitPayment

```solidity
function commitPayment(uint32 toTaskNumber, uint96 amount) external
```

This is used by an operator to make a claim on the amount that they deserve for their service from their last payment until `toTaskNumber`

_Once this payment is recorded, a fraud proof period commences during which a challenger can dispute the proposed payment._

### redeemPayment

```solidity
function redeemPayment() external
```

Called by an operator to redeem a payment that they previously 'committed' to by calling `commitPayment`.

_This function can only be called after the challenge window for the payment claim has completed._

### _payForServiceHook

```solidity
function _payForServiceHook(contract IDelegationTerms dt, uint256 amount) internal
```

### initPaymentChallenge

```solidity
function initPaymentChallenge(address operator, uint96 amount1, uint96 amount2) external
```

This function is called by a fraud prover to challenge a payment, initiating an interactive-type fraudproof.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | is the operator against whose payment claim the fraudproof is being made |
| amount1 | uint96 | is the reward amount the challenger in that round claims is for the first half of tasks |
| amount2 | uint96 | is the reward amount the challenger in that round claims is for the second half of tasks |

### performChallengeBisectionStep

```solidity
function performChallengeBisectionStep(address operator, bool secondHalf, uint96 amount1, uint96 amount2) external
```

Perform a single bisection step in an existing interactive payment challenge.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | The middleware operator who was challenged (used to look up challenge details) |
| secondHalf | bool | If true, then the caller wishes to challenge the amount claimed as payment in the *second half* of the previous bisection step. If false then the *first half* is indicated instead. |
| amount1 | uint96 | The amount that the caller asserts the operator is entitled to, for the first half *of the challenged half* of the previous bisection. |
| amount2 | uint96 | The amount that the caller asserts the operator is entitled to, for the second half *of the challenged half* of the previous bisection. |

### _updateStatus

```solidity
function _updateStatus(address operator, uint32 diff) internal returns (bool)
```

This function is used for updating the status of the challenge in terms of who has to respon
to the interactive challenge mechanism next -  is it going to be challenger or the operator.

_If the challenge is over only one task, then the challenge is marked specially as a one step challenge –
the smallest unit over which a challenge can be proposed – and 'true' is returned.
Otherwise status is updated normally and 'false' is returned._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | is the operator whose payment claim is being challenged |
| diff | uint32 | is the number of tasks across which payment is being challenged in this iteration |

### _updateChallengeAmounts

```solidity
function _updateChallengeAmounts(address operator, enum IPaymentManager.DissectionType dissectionType, uint96 amount1, uint96 amount2) internal
```

Used to update challenge amounts when the operator (or challenger) breaks down the challenged amount (single bisection step)

### resolveChallenge

```solidity
function resolveChallenge(address operator) external
```

resolve an existing PaymentChallenge for an operator

### _resolve

```solidity
function _resolve(struct IPaymentManager.PaymentChallenge challenge, address winner) internal
```

Resolves a single payment challenge, paying the winner.

_If challenger is proven correct, then they are refunded their own challengeAmount plus the challengeAmount put up by the operator.
If operator is proven correct, then the challenger's challengeAmount is transferred to them, since the operator still hasn't been
proven right, and thus their challengeAmount is still required in case they are challenged again._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| challenge | struct IPaymentManager.PaymentChallenge | The challenge that is being resolved. |
| winner | address | Address of the winner of the challenge. |

### getChallengeStatus

```solidity
function getChallengeStatus(address operator) external view returns (enum IPaymentManager.ChallengeStatus)
```

Returns the ChallengeStatus for the `operator`'s payment claim.

### getAmount1

```solidity
function getAmount1(address operator) external view returns (uint96)
```

Returns the 'amount1' for the `operator`'s payment claim.

### getAmount2

```solidity
function getAmount2(address operator) external view returns (uint96)
```

Returns the 'amount2' for the `operator`'s payment claim.

### getToTaskNumber

```solidity
function getToTaskNumber(address operator) external view returns (uint48)
```

Returns the 'toTaskNumber' for the `operator`'s payment claim.

### getFromTaskNumber

```solidity
function getFromTaskNumber(address operator) external view returns (uint48)
```

Returns the 'fromTaskNumber' for the `operator`'s payment claim.

### getDiff

```solidity
function getDiff(address operator) external view returns (uint48)
```

Returns the task number difference for the `operator`'s payment claim.

### getPaymentChallengeAmount

```solidity
function getPaymentChallengeAmount(address operator) external view returns (uint256)
```

Returns the active challengeAmount of the `operator` placed on their payment claim.

### _taskNumber

```solidity
function _taskNumber() internal view returns (uint32)
```

Convenience function for fetching the current taskNumber from the `serviceManager`

### _setPaymentChallengeAmount

```solidity
function _setPaymentChallengeAmount(uint256 _paymentChallengeAmount) internal
```

Modifies the `paymentChallengeAmount` amount.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| _paymentChallengeAmount | uint256 | The new value for `paymentChallengeAmount` to take. |

