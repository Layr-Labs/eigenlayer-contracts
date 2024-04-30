## PaymentCoordinator

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`PaymentCoordinator.sol`](../../src/contracts/core/PaymentCoordinator.sol) | Singleton | Transparent proxy |


The primary functions of the `PaymentCoordinator` contract are for (i) handling ERC20 payments from AVSs (Actively Validated Services) to their operators and delegated stakers for a given time range; (ii) enable the protocol to provide ERC20 tokens to all stakers over a specified time range; and (iii) allows stakers and operators to claim their accumulated earnings.

AVSs can submit payments to the `PaymentCoordinator` contract for a given time range with a specified ERC20 token and amount; this struct in particular is called a `RangePayment`. All AVS `RangePayments` are consolidated into a merkle tree and later distributed via posting of the merkle root by a permissioned entity of the protocol. These merkle trees are periodically calculated in our backend processing and all historical roots are stored in the contract. After a `DistributionRoot` is submitted and the activation delay has passed, stakers/operators(or their personally set claimers) can claim all their accumulated earnings by providing a merkle proof that gets verified against the latest `DistributionRoot`. This entire flow will repeat periodically as AVSs submit `RangePayment`s, `DistributionRoot`s are submitted, and stakers/operators claim their accumulated earnings. 

Note: stakers/operators aren't required to claim against every root unless they want immediate access to their funds. Since earnings are cumulative, they can be claimed all at once if proving against the latest root.

#### High-level Concepts

This document organizes methods according to the following themes (click each to be taken to the relevant section):
* [Submitting Payment Requests](#submitting-payment-requests)
* [Distributing and Claiming Payments](#distributing-and-claiming-payments)
* [System Configuration](#system-configuration)
* [Payments Merkle Tree Structure](#payments-merkle-tree-structure)


#### Important state variables

* `DistributionRoot[] public distributionRoots`:
    * The `DistributionRoot` array that stores historic payment merkle tree roots. The latest root can be retrieved by getting the last array index. The payment merkle tree stores for each earner, their cumulative earnings per ERC20 payment token. More details on the tree can be found in the Payment Merkle Tree Structure section below.
    * Claims made against the latest `DistributionRoot` will include all earnings up to that point.
    * `DistributionRoot` struct contains in addition to the payment merkle root, the `paymentCalculationEndTimestamp` and `activatedAt` timestamps in seconds. 
    * Only the `paymentUpdater` role can submit new `DistributionRoot`s via calling `submitRoot`.
    * related functions: `submitRoot`, `processClaim`
* `mapping(address => bool) public isPayAllForRangeSubmitter`:
    * Mapping for pay all for range submitters. Only a valid pay all for range submitter can call `payAllForRange` which distributes a RangePayment across all EigenLayer stakers.
    * related functions: `payAllForRange`, `setPayAllForRangeSubmitter`
* `mapping(address => address) public claimerFor`: earner => claimer
    * Mapping for earners(stakers/operators) to track their claimer address. The claimer is the address that can call `processClaim` on behalf of the earner. If the claimer is not set i.e `claimerFor[earner] == address(0)` , the earner themselves can call `processClaim` directly.
    * Note that the claimerFor isn't necessarily the address that the payment tokens are transferred to but rather have the authority to process the claims and direct the tokens to their recipient address. See the interface for `processClaim`.
    * related functions: `processClaim`, `setClaimerFor`
* `mapping(address => mapping(IERC20 => uint256)) public cumulativeClaimed`: earner => token => total amount claimed
    * Mapping for earners(stakers/operators) to track their total claimed earnings per ERC20 token. This mapping is used to calculate the difference between the cumulativeEarnings in the merkle tree and the previous total claimed amount. This difference is then transfered to the specified destination address.
    * related functions: `processClaim`

#### Helpful definitions

* `_checkClaim(PaymentMerkleClaim calldata claim, DistributionRoot memory root)`
    * Checks the merkle inclusion of a claim against a `DistributionRoot`
    * Reverts if any of the following are true:
        * mismatch input param lengths: tokenIndices, tokenTreeProofs, tokenLeaves
        * earner proof reverting from calling `_verifyEarnerClaimProof`
        * any of the token proofs reverting from calling `_verifyTokenClaimProof`
* Wherever AVS (Actively Validated Service) is mentioned, it refers to the contract entity that is submitting payments to the PaymentCoordinator. This is assumed to be a customized ServiceManager contract of some kind that is interfacing with the EigenLayer protocol. See base contract here [ServiceManager.sol](https://github.com/Layr-Labs/eigenlayer-middleware/blob/dev/src/ServiceManagerBase.sol)

---

### Submitting Payment Requests

Payment Requests are submitted through the following functions:

* [`PaymentCoordinator.payForRange`](#payforrange)
* [`PaymentCoordinator.payAllForRange`](#payallforrange)

#### `payForRange`

```solidity
function payForRange(RangePayment[] calldata rangePayments) external
```

Called by an AVS to submit a list of `RangePayment`s to the PaymentCoordinator to be distributed across all operators and its delegated stakers registered to the AVS. Each `RangePayment` struct contains the following fields:
* `StrategyAndMultiplier[] strategiesAndMultipliers`: an array of `StrategyAndMultiplier` structs that define the strategy and relative multiplier weighting. This linear combination of strategies of weights are the EigenLayer strategies that an AVS considers eligible for payment.
    * Each `StrategyAndMultiplier` struct contains the following fields:
        * `IStrategy strategy`: address of an EigenLayer strategy contract
        * `uint96 multiplier`: the relative weighting of the strategy in the linear combination. Recommended use here is to use 1e18 as the base multiplier and adjust the relative weightings accordingly
* `IERC20 token`: the address of ERC20 token being used for payment
* `uint256 amount`: amount of the ERC20 payment token to transfer to the PaymentCoordinator for distribution
* `uint32 startTimestamp`: the start of the payment time range for which the payment is being made for
* `uint32 duration`: the duration of the payment time range for which the payment is being made for

`payForRange` is called by AVSs to submit payments in the form of any ERC20 token over a specified time range. The payment distribution amongst the AVS's operators and delegated stakers is determined offchain by the strategies and multipliers provided in the `RangePayment` struct as well as the actual operator/staker shares for those defined strategies over that time range. Those operator/staker shares are read from the core contracts (see DelegationManager and StrategyManager).

*Effects*:
* For each `RangePayment` element
    * Transfers `amount` of `token` from the msg.sender (AVS) to the `PaymentCoordinator`
    * Hashes msg.sender(AVS), nonce, and `RangePayment` struct to create a unique payment hash and sets this mapping to true in `isRangePaymentHash`
    * Increments `paymentNonce[msg.sender]`
    * Emits a `RangePaymentCreated` event

*Requirements*:
* For each `RangePayment` element
    * `isRangePaymentHash[msg.sender][rangePaymentHash]` MUST be false, the hash keccak256(msg.sender[AVS], nonce, and `RangePayment`) should not have been submitted before
    * Requirements from calling internal function `_payForRange()`
        * `rangePayment.strategiesAndMultipliers.length > 0`
        * `rangePayment.amount > 0`
        * `rangePayment.duration <= MAX_PAYMENT_DURATION`
        * `rangePayment.duration % calculationIntervalSeconds == 0`
        * `rangePayment.startTimestamp % calculationIntervalSeconds == 0`
        * `block.timestamp - MAX_RETROACTIVE_LENGTH <= rangePayment.startTimestamp`
        * `GENESIS_PAYMENT_TIMESTAMP <= rangePayment.startTimestamp`
        * `rangePayment.startTimestamp <= block.timestamp + MAX_FUTURE_LENGTH`
        * Requirements for `rangePayment.strategiesAndMultipliers`
            * Each `strategy` is whitelisted for deposit in the StrategyManager or is the `beaconChainETHStrategy()`
            * `rangePayment.strategiesAndMultipliers` is sorted by ascending strategy address to prevent duplicate strategies
* Pause status MUST NOT be set: `PAUSED_PAY_FOR_RANGE`
* Function call is not reentered

The text diagram below better visualizes a valid start timestamp for a `RangePayment`
```
Sliding Window for valid RangePayment startTimestamp

Scenario A: GENESIS_PAYMENT_TIMESTAMP IS WITHIN RANGE
        <-----MAX_RETROACTIVE_LENGTH-----> t (block.timestamp) <---MAX_FUTURE_LENGTH--->
            <--------------------valid range for startTimestamp------------------------>
            ^
        GENESIS_PAYMENT_TIMESTAMP


Scenario B: GENESIS_PAYMENT_TIMESTAMP IS OUT OF RANGE
        <-----MAX_RETROACTIVE_LENGTH-----> t (block.timestamp) <---MAX_FUTURE_LENGTH--->
        <------------------------valid range for startTimestamp------------------------>
    ^
GENESIS_PAYMENT_TIMESTAMP
```

#### `payAllForRange`

```solidity
function payAllForRange(RangePayment[] calldata rangePayments) external onlyPayAllForRangeSubmitter
```

Called by an permissioned entity of protocol to submits a list of `RangePayment`s to the PaymentCoordinator to be distributed across all EigenLayer stakers. Just like `payForRange`, each `RangePayment` struct contains the following fields:
* `StrategyAndMultiplier[] strategiesAndMultipliers`: an array of `StrategyAndMultiplier` structs that define the strategy and relative multiplier weighting. This linear combination of strategies of weights are the EigenLayer strategies that an AVS considers eligible for payment.
    * Each `StrategyAndMultiplier` struct contains the following fields:
        * `IStrategy strategy`: address of an EigenLayer strategy contract
        * `uint96 multiplier`: the relative weighting of the strategy in the linear combination. Recommended use here is to use 1e18 as the base multiplier and adjust the relative weightings accordingly
* `IERC20 token`: the address of ERC20 token being used for payment
* `uint256 amount`: amount of the ERC20 payment token to transfer to the PaymentCoordinator for distribution
* `uint32 startTimestamp`: the start of the payment time range for which the payment is being made for
* `uint32 duration`: the duration of the payment time range for which the payment is being made for

`payAllForRange` is called by a valid submitter `isPayAllForRangeSubmitter`,  to submit payments in the form of any ERC20 token over a specified time range. The payment distribution amongst all EigenLayer stakers is determined offchain by the strategies and multipliers provided in the `RangePayment` struct as well as the actual staker shares for those defined strategies over that time range. Those staker shares are read from the core contracts (see DelegationManager and StrategyManager).

*Effects*:
* For each `RangePayment` element
    * Transfers `amount` of `token` from the msg.sender (AVS) to the `PaymentCoordinator`
    * Hashes msg.sender(payAllForRangeSubmitter), nonce, and `RangePayment` struct to create a unique payment hash and sets this mapping to true in `isRangePaymentForAllHash`
    * Increments `paymentNonce[msg.sender]`
    * Emits a `RangePaymentForAllCreated` event

*Requirements*:
* For each `RangePayment` element
    * `isRangePaymentForAllHash[msg.sender][rangePaymentHash]` MUST be false, the hash keccak256(msg.sender[payAllForRangeSubmitter], nonce, and `RangePayment`) should not have been submitted before
    * Requirements from calling internal function `_payForRange()`
        * `rangePayment.strategiesAndMultipliers.length > 0`
        * `rangePayment.amount > 0`
        * `rangePayment.duration <= MAX_PAYMENT_DURATION`
        * `rangePayment.duration % calculationIntervalSeconds == 0`
        * `rangePayment.startTimestamp % calculationIntervalSeconds == 0`
        * `block.timestamp - MAX_RETROACTIVE_LENGTH <= rangePayment.startTimestamp`
        * `GENESIS_PAYMENT_TIMESTAMP <= rangePayment.startTimestamp`
        * `rangePayment.startTimestamp <= block.timestamp + MAX_FUTURE_LENGTH`
        * Requirements for `rangePayment.strategiesAndMultipliers`
            * Each `strategy` is whitelisted for deposit in the StrategyManager or is the `beaconChainETHStrategy()`
            * `rangePayment.strategiesAndMultipliers` is sorted by ascending strategy address to prevent duplicate strategies
* Pause status MUST NOT be set: `PAUSED_PAY_ALL_FOR_RANGE`
* `isRangePaymentForAllHash[msg.sender] == true` checked in modifier onlyPayAllForRangeSubmitter
* Function call is not reentered

---

### Distributing and Claiming Payments

Distribution of Payments Payment Requests are submitted through the following functions:

* [`PaymentCoordinator.submitRoot`](#submitroot)
* [`PaymentCoordinator.setClaimerFor`](#setclaimerfor)
* [`PaymentCoordinator.processClaim`](#processclaim)

#### `submitRoot`

```solidity
function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) external
```

Called only by the `paymentUpdater` singleton address to submit a new `DistributionRoot` to the PaymentCoordinator. The `DistributionRoot` struct contains the following fields:
* `bytes32 root`: the merkle root of the payment merkle tree
* `uint32 paymentCalculationEndTimestamp`: the end of the payment time range for which the `DistributionRoot` is being submitted for
* `uint32 activatedAt`: the timestamp in seconds when the `DistributionRoot` is activated and can be claimed against

`submitRoot` pushes a new `DistributionRoot` to the `distributionRoots` array. This root is used to verify claims made by stakers/operators against the latest `DistributionRoot`. The `activatedAt` timestamp is set to `block.timestamp + activationDelay()` to allow for a delay before claims can be processed and for verification of the root submitted.

*Effects*:
* Pushes a new `DistributionRoot` to the `distributionRoots` array
* Emits a `DistributionRootSubmitted` event

*Requirements*:
* `paymentCalculationEndTimestamp > currPaymentCalculationEndTimestamp`
* `paymentCalculationEndTimestamp < block.timestamp`
* Pause status MUST NOT be set: `PAUSED_SUBMIT_ROOTS`
* msg.sender is `paymentUpdater`

#### `setClaimerFor`

```solidity
function setClaimerFor(address claimer) external
```

Called by an earner to set a claimer address that can call `processClaim` on their behalf. If the claimer is not set i.e `claimerFor[earner] == address(0)` , the earner themselves can call `processClaim` directly. Of course the claimer can be set to the earner themselves if they wish but it would functionally be the same as not setting a claimer and leaving this mapping set as `address(0)`.

*Effects*:
* Sets the `claimerFor[msg.sender]` to the input param `claimer`
* Emits a `ClaimerForSet` event

*Requirements*:
* Only the earner can set their claimer since the earner is read from `msg.sender`


#### `processClaim`

```solidity
function processClaim(PaymentMerkleClaim calldata claim, address recipient) external
```

Called by earners (stakers/operators) or their set claimers to claim their accumulated earnings. The `PaymentMerkleClaim` struct contains the following fields:
* `uint32 rootIndex`: the array index of the `DistributionRoot` to verify the claim against
* `uint32 earnerIndex`: the index of the earner's account root in the merkle tree
* `bytes earnerTreeProof`: the proof of the earner's EarnerTreeMerkleLeaf against the merkle root
* `EarnerTreeMerkleLeaf earnerLeaf`: The earner's EarnerTreeMerkleLeaf struct, defined by the following fields
    * `address earner`: the address of the earner
    * `bytes32 earnerTokenRoot`: the merkle root of the earner's token merkle tree
* `uint32[] tokenIndices`: the indices of the token leaves in the earner's subtree
* `bytes[] tokenTreeProofs`: the proofs of the token leaves against the earner's earnerTokenRoot
* `TokenTreeMerkleLeaf[] tokenLeaves`: the token leaves to be claimed, each `TokenTreeMerkleLeaf` defined by the following fields
    * `IERC20 token`: the ERC20 token to be claimed
    * `uint256 amount`: the amount of the ERC20 token to be claimed

`processClaim` will first call `_checkClaim` to verify the merkle proofs against the `DistributionRoot` at the specified `rootIndex`. This is done by first performing a merkle proof verification of the earner's `EarnerTreeMerkleLeaf` against the `DistributionRoot` and then for each tokenIndex, verifying each token leaf against the earner's `earnerTokenRoot`. 

The caller must be the set claimer address in the `claimerFor` mapping or the earner themselves if the claimer is not set. The claimer has the right to process the claim for a given earner.
After the claim is verified, for each token leaf, the difference between the cumulative earnings in the merkle tree and the previous total claimed amount last stored in the contract is calculated and transferred from the `PaymentCoordinator` contract to the address `recipient`.

*Effects*:
* For each `TokenTreeMerkleLeaf`, 
    * transfers the difference between the cumulative earnings in the merkle tree and the previous total claimed amount to the specified destination address
    * Updates the `cumulativeClaimed` mapping for the earner and token
    * Emits a `PaymentClaimed` event

*Requirements*:
* `_checkClaim(PaymentMerkleClaim calldata claim, DistributionRoot memory root)` does not revert:
    * Checks the merkle inclusion of a claim against a `DistributionRoot`
    * Reverts if any of the following are true:
        * mismatch input param lengths: tokenIndices, tokenTreeProofs, tokenLeaves
        * earner proof reverting from calling `_verifyEarnerClaimProof`
        * any of the token proofs reverting from calling `_verifyTokenClaimProof`
* msg.sender is NOT the `claimerFor[earner]` and msg.sender is NOT the earner themself
* For each `TokenTreeMerkleLeaf`, 
    * `tokenLeaf.cumulativeEarnings >= cumulativeClaimed[earner][token]`: cumulativeEarnings must be gte than cumulativeClaimed

---

### System Configuration

#### `setCalculationIntervalSeconds`

```solidity
function setCalculationIntervalSeconds(uint32 _calculationIntervalSeconds) external
```

Allows the Owner to set the calculation interval seconds. `RangePayment.duration` and `RangePayment.startTimestamp` must be in multiples of this interval. This requirement reduces offchain payment calculations and will be useful in the future with slashing.

*Effects*:
* Sets the global `calculationIntervalSeconds`
* Emits a `CalculationIntervalSecondsSet` event

*Requirements*:
* Caller MUST be the Owner

#### `setActivationDelay`

```solidity
function setActivationDelay(uint32 _activationDelay) external
```

Allows the owner to set the activation delay (seconds). The activation delay is the time in seconds after a `DistributionRoot` is submitted before it can be claimed against. This delay is to allow for the verification of the root submitted.

*Effects*:
* Sets the global `activationDelay`
* Emits a `ActivationDelaySet` event

*Requirements*:
* Caller MUST be the Owner

#### `setGlobalOperatorCommission`

```solidity
function setGlobalOperatorCommission(uint16 _globalCommissionBips) external
```

Allows the owner to set the global operator commission in basis points. This commission is applied to all operator earnings across all AVSs. The commission is calculated as a percentage of the operator earnings and is deducted from the operator's earnings before distribution to its delegated stakers. This is expected to be a flat 10% rate for the initial payments release.

*Effects*:
* Sets the global `globalOperatorCommissionBips`
* Emits a `GlobalCommissionBipsSet` event

*Requirements*:
* Caller MUST be the Owner

#### `setPaymentUpdater`

```solidity
function setPaymentUpdater(address _paymentUpdater) external
```

Allows the owner to set the `paymentUpdater` address. The `paymentUpdater` is the singleton address that can submit new `DistributionRoot`s to the PaymentCoordinator contract. This address is expected to be a separate entity that will periodically submit payment merkle roots to the contract.

*Effects*:
* Sets the global `paymentUpdater` address
* Emits a `PaymentUpdaterSet` event

*Requirements*:
* Caller MUST be the Owner

#### `setPayAllForRangeSubmitter`

```solidity
function setPayAllForRangeSubmitter(address _submitter, bool _newValue) external
```

Allows the owner to set the `isPayAllForRangeSubmitter` mapping. This mapping is used to determine if a given address is a valid submitter for `payAllForRange`. Only a valid submitter can call `payAllForRange` which submits a `RangePayment` that will be distributed across all EigenLayer stakers.

*Effects*:
* Sets the `isPayAllForRangeSubmitter` mapping for the address `_submitter` to the bool `_newValue`
* Emits a `PayAllForRangeSubmitterSet` event

*Requirements*:
* Caller MUST be the Owner

---

### Payments Merkle Tree Structure

This merkle tree is used to verify the claims made by earners' claimers against a `DistributionRoot`.
The `DistributionRoot` consolidates all `RangePayment`s submitted by AVSs since the previously submitted `DistributionRoot` into a merkle tree comprised of earners and their cumulative earnings for their respective payment tokens distributed.

When an earner or their designated claimer calls `processClaim`, they must provide a `PaymentMerkleClaim` struct that contains the necessary information to verify their claim against the latest `DistributionRoot`. The merkle proof verification is done in the internal `_checkClaim` helper function. This function verifies the merkle proof of the earner's `EarnerTreeMerkleLeaf` against the `DistributionRoot` and then for each tokenIndex, verifies each token leaf against the earner's `earnerTokenRoot`.

Claimers can selectively choose which token leaves to prove against and claim accumulated earnings. Each token payment claimed in a `processClaim` call will send tokens to the `recipient` address specified in the call. 

The payment merkle tree is structured in the diagram below:

![.](../images/PaymentCoordinator_Merkle_Tree.png)