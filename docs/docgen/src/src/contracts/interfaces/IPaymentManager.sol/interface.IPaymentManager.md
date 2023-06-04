# IPaymentManager
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IPaymentManager.sol)

**Author:**
Layr Labs, Inc.


## Functions
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

### takeFee

Used for deducting the fees from the payer to the middleware


```solidity
function takeFee(address initiator, address payer, uint256 feeAmount) external;
```

### setPaymentChallengeAmount

Modifies the `paymentChallengeAmount` amount.


```solidity
function setPaymentChallengeAmount(uint256 _paymentChallengeAmount) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_paymentChallengeAmount`|`uint256`|The new value for `paymentChallengeAmount` to take.|


### commitPayment

This is used by an operator to make a claim on the amount that they deserve for their service from their last payment until `toTaskNumber`

*Once this payment is recorded, a fraud proof period commences during which a challenger can dispute the proposed payment.*


```solidity
function commitPayment(uint32 toTaskNumber, uint96 amount) external;
```

### redeemPayment

Called by an operator to redeem a payment that they previously 'committed' to by calling `commitPayment`.

*This function can only be called after the challenge window for the payment claim has completed.*


```solidity
function redeemPayment() external;
```

### initPaymentChallenge

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


### resolveChallenge

resolve an existing PaymentChallenge for an operator


```solidity
function resolveChallenge(address operator) external;
```

### paymentFraudproofInterval

Challenge window for submitting fraudproof in the case of an incorrect payment claim by a registered operator.


```solidity
function paymentFraudproofInterval() external view returns (uint256);
```

### paymentChallengeAmount

Specifies the payment that has to be made as a guarantee for fraudproof during payment challenges.


```solidity
function paymentChallengeAmount() external view returns (uint256);
```

### paymentToken

the ERC20 token that will be used by the disperser to pay the service fees to middleware nodes.


```solidity
function paymentToken() external view returns (IERC20);
```

### paymentChallengeToken

Token used for placing a guarantee on challenges & payment commits


```solidity
function paymentChallengeToken() external view returns (IERC20);
```

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

Returns the active guarantee amount of the `operator` placed on their payment claim.


```solidity
function getPaymentChallengeAmount(address) external view returns (uint256);
```

## Structs
### Payment
used for storing information on the most recent payment made to the operator


```solidity
struct Payment {
    uint32 fromTaskNumber;
    uint32 toTaskNumber;
    uint32 confirmAt;
    uint96 amount;
    PaymentStatus status;
    uint256 challengeAmount;
}
```

### PaymentChallenge
used for storing information on the payment challenge as part of the interactive process


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
    ChallengeStatus status;
}
```

### TotalStakes

```solidity
struct TotalStakes {
    uint256 signedStakeFirstQuorum;
    uint256 signedStakeSecondQuorum;
}
```

## Enums
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

