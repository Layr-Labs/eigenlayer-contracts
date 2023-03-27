# Solidity API

## IPaymentManager

### DissectionType

```solidity
enum DissectionType {
  INVALID,
  FIRST_HALF,
  SECOND_HALF
}
```

### PaymentStatus

```solidity
enum PaymentStatus {
  REDEEMED,
  COMMITTED,
  CHALLENGED
}
```

### ChallengeStatus

```solidity
enum ChallengeStatus {
  RESOLVED,
  OPERATOR_TURN,
  CHALLENGER_TURN,
  OPERATOR_TURN_ONE_STEP,
  CHALLENGER_TURN_ONE_STEP
}
```

### Payment

```solidity
struct Payment {
  uint32 fromTaskNumber;
  uint32 toTaskNumber;
  uint32 confirmAt;
  uint96 amount;
  enum IPaymentManager.PaymentStatus status;
  uint256 challengeAmount;
}
```

### PaymentChallenge

```solidity
struct PaymentChallenge {
  address operator;
  address challenger;
  address serviceManager;
  uint32 fromTaskNumber;
  uint32 toTaskNumber;
  uint96 amount1;
  uint96 amount2;
  uint32 settleAt;
  enum IPaymentManager.ChallengeStatus status;
}
```

### TotalStakes

```solidity
struct TotalStakes {
  uint256 signedStakeFirstQuorum;
  uint256 signedStakeSecondQuorum;
}
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

### takeFee

```solidity
function takeFee(address initiator, address payer, uint256 feeAmount) external
```

Used for deducting the fees from the payer to the middleware

### setPaymentChallengeAmount

```solidity
function setPaymentChallengeAmount(uint256 _paymentChallengeAmount) external
```

Modifies the `paymentChallengeAmount` amount.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| _paymentChallengeAmount | uint256 | The new value for `paymentChallengeAmount` to take. |

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

### resolveChallenge

```solidity
function resolveChallenge(address operator) external
```

resolve an existing PaymentChallenge for an operator

### paymentFraudproofInterval

```solidity
function paymentFraudproofInterval() external view returns (uint256)
```

Challenge window for submitting fraudproof in the case of an incorrect payment claim by a registered operator.

### paymentChallengeAmount

```solidity
function paymentChallengeAmount() external view returns (uint256)
```

Specifies the payment that has to be made as a guarantee for fraudproof during payment challenges.

### paymentToken

```solidity
function paymentToken() external view returns (contract IERC20)
```

the ERC20 token that will be used by the disperser to pay the service fees to middleware nodes.

### paymentChallengeToken

```solidity
function paymentChallengeToken() external view returns (contract IERC20)
```

Token used for placing a guarantee on challenges & payment commits

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
function getPaymentChallengeAmount(address) external view returns (uint256)
```

Returns the active guarantee amount of the `operator` placed on their payment claim.

