
[elip-002]: https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md

# Shares Accounting

This document outlines the changes to the Staker and Operator Shares accounting resulting from the Slashing Upgrade. There are several introduced variables such as the _deposit scaling factor_ ($k_n$), _max magnitude_ ($m_n$), and _beacon chain slashing factor_ ($l_n$). How these interact with the Operator and Staker events like deposits, slashing, withdrawals will all be described below.

## Prior Reading

* [ELIP-002: Slashing via Unique Stake and Operator Sets](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md)

## Terminology

The word "shares" in EigenLayer has historically referred to the amount of shares a Staker receives upon depositing assets through the `StrategyManager` or `EigenPodManager`. Outside of some conversion ratios in the `StrategyManager` to account for rebasing tokens, shares roughly correspond 1:1 with deposit amounts (i.e. 1e18 shares in the `beaconChainETHStrategy` corresponds to 1 ETH of assets). When delegating to an operator or queueing a withdrawal, the `DelegationManager` reads deposit shares from the `StrategyManager` or `EigenPodManager` to determine how many shares to delegate (or undelegate).

With the slashing release, there is a need to differentiate "classes" of shares.

**Deposit shares**: 

Formerly known as "shares," these are the same shares used before the slashing release. They continue to be managed by the `StrategyManager` and `EigenPodManager`, and roughly correspond 1:1 with deposited assets.

**Withdrawable shares**: 

When an operator is slashed, the slash is applied to their stakers _asynchronously_ (otherwise, slashing would require iterating over each of an operator's stakers; this is prohibitively expensive). 

The `DelegationManager` must find a common representation for the deposit shares of many stakers, each of which may have experienced different amounts of slashing depending on which operator they are delegated to, and when they delegated. This common representation is achieved in part through a value called the `depositScalingFactor`: a per-staker, per-strategy value that scales a staker's deposit shares as they deposit assets over time.

When a staker does just about anything (changing their delegated operator, queueing/completing a withdrawal, depositing new assets), the `DelegationManager` converts their _deposit shares_ to _withdrawable shares_ by applying the staker's `depositScalingFactor` and the current _slashing factor_ (a per-strategy scalar primarily derived from the amount of slashing an operator has received in the `AllocationManager`).

These _withdrawable shares_ are used to determine how many of a staker's deposit shares are actually able to be withdrawn from the protocol, as well as how many shares can be delegated to an operator. A staker's withdrawable shares are not reflected anywhere in storage; they are calculated on-demand.

**Operator shares**:

_Operator shares_ are derivative of _withdrawable shares_. When a staker delegates to an operator, they are delegating their _withdrawable shares_. Thus, an operator's _operator shares_ represent the sum of all of their stakers' _withdrawable shares_. Note that when a staker first delegates to an operator, this is a special case where _deposit shares_ == _withdrawable shares_. If the staker deposits additional assets later, this case will not hold if slashing was experienced in the interim.

## Pre-Slashing Upgrade

We'll look at the "shares" model as historically defined prior to the Slashing upgrade. We can write this a bit more formally:

### Staker Level

$s_n$ - The amount of shares in the storage of the `StrategyManager`/`EigenPodManager` at time n.

### Operator Level

$op_n$ - The operator shares in the storage of the `DelegationManager` at time n which can also be rewritten as \
$op_n = \sum_{i=1}^{k} s_{n,i}$ where the operator has $k$ number of stakers delegated to them.


### Staker Deposits 

Upon each Staker deposit of amount $d_n$ at time $n$, the Staker's shares and delegated Operator's shares are updated as follows:

$$
 s_{n+1} = s_{n} + d_{n}
$$

$$
 op_{n+1} = op_{n} + d_{n}
$$

### Staker Withdrawals 

Similarly for Staker withdrawals, given an amount $w_n$ to withdraw at time $n$, the Staker and Operator's shares are decremented at the point of the withdrawal being queued:

$$
 s_{n+1} = s_{n} - w_{n}
$$

$$  
 op_{n+1} = op_{n} - w_{n}
$$

Later after the withdrawal delay has passed, the Staker can complete their withdrawal to withdraw the full amount $w_n$ of shares. 


## Slashing Upgrade Changes

As of release `v1.0.0` and the introduction of Unique Stake and Operator Sets, programmatic slashing will be enabled in the core EigenLayer protocol where Staker deposits can be subject to slashing. 
The remaining portions of this document will assume understanding of Allocations/Deallocations, Max Magnitudes, and OperatorSets.
For more information on this, there is the [ELIP-002][elip-002] which has a high-level but detailed overview of the Slashing upgrade.

We now introduce a few new types of Shares concepts:

1. **deposit shares**: \
 For a Staker, this is the amount of Strategy shares deposited
2. **withdrawable shares**: \
 For a Staker, this is the actual amount of Strategy shares they are eligible to withdraw. \
 This does not live in storage but is read through the view function `DelegationManager.getWithdrawableShares`. Note that this amount is <= deposit shares as the Staker may have had their shares slashed.
3. **operator/delegated shares**: \
 This still remains the same definition as before, the amount of delegated shares of an operator from all their delegated stakers.
 However, this is now equal to the summation of all their staker's withdrawable shares.

Notice that these definitions also apply to the shares model prior to the Slashing upgrade but with the caveat that for all Stakers, withdrawable shares equal the deposit shares. After the Slashing upgrade this is not necessarily the case if a Staker's delegated Operator were slashed resulting in less withdrawable shares for the Staker. \
Now let's look at these updated definitions in detail and how the accounting math works with deposits, withdrawals, and slashing.

---

### Stored Variables

Note that these variables are all defined within the context of a single Strategy.
#### Staker Level

$s_n$ - The amount of deposit shares in the storage of the `StrategyManager`/`EigenPodManager` at time $n$. \
 Exists in storage: `StrategyManager.stakerDepositShares`, `EigenPodManager.stakerDepositShares` for beaconChainETHStrategy\
$k_n$ - The Staker “deposit scaling factor” at time $n$. This is initialized to 1. \
 Exists in storage: `DelegationManager.depositScalingFactor` \
$l_n$ - The Staker's "beacon chain slashing factor" at time $n$. This is initialized to 1 and for any non-native ETH Strategies always is always fixed to 1 rather than calculating withdrawable shares completely different depending on native versus non-native ETH.
 Exists in storage: `EigenPodManager.beaconChainSlashingFactor`

#### Operator Level

$m_n$ - The operator magnitude at time n. This is initialized to 1. \
$op_n$ - The operator shares in the storage of the `DelegationManager` at time n which can also be rewritten as $op_n = \sum_{i=1}^{k} a_{n,i}$ \
 Exists in storage: `DelegationManager.operatorShares`

### Conceptual Variables

$a_n = s_n k_n l_n m_n$ - The withdrawable shares that the staker owns at time $n$.
 Read from view function `DelegationManager.getWithdrawableShares`



---

### Deposits

For an amount of newly deposited shares $d_n$,

#### Staker Level

From the conceptual level, the staker's deposit shares and withdrawable shares should increase by the deposited amount $d_n$. Let's workout how this math impacts some of the deposit scaling factor $k_n$.

$$
a_{n+1} = a_n + d_n
$$

$$
s_{n+1} = s_n +d_n
$$

$$
l_{n+1} = l_n
$$

$$
m_{n+1} = m_n
$$

Expanding the $a_{n+1}$ calculation

$$
s_{n+1} k_{n+1} l_{n+1} m_{n+1} = s_n k_n l_n m_n + d_n
$$

Which yields

$$
k_{n+1} = \frac{s_n k_n m_n + d_n}{s_{n+1} l_{n+1} m_{n+1}}=\frac{s_n k_n l_n m_n + d_n}{(s_n+d_n)l_nm_n}
$$

#### Operator Level

For the operator (if the staker is delegated), the delegated operator shares should increase by the exact amount
the staker just deposited. Therefore $op_n$ is updated as follows:

$$
op_{n+1} = op_n+d_n
$$


See implementation in:
* [`StrategyManager.depositIntoStrategy`](../../../src/contracts/core/StrategyManager.sol)
* [`EigenPod`](../../../src/contracts/pods/EigenPod.sol)
* [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](../../../src/contracts/pods/EigenPodManager.sol)

<br> 

---

### Slashing

Given a proportion to slash $p_n = \frac {m_{n+1}}{m_n}$ ,

#### Operator Level

From a conceptual level, operator shares should be decreased by the proportion according to the following:

$$
 op_{n+1} = op_n p_n
$$

$$
 => op_{n+1} = op_n \frac {m_{n+1}} {m_n}
$$ 

However, since we don't overwrite `operatorShares` directly in storage and perform increment/decrement operations we will calculate the amount of $sharesToDecrement$.

$$
 sharesToDecrement = op_n - op_{n+1}
$$

$$
 = op_n - op_n \frac {m_{n+1}} {m_n}
$$

_sharesToDecrement_ is calculated to be decremented from $op_n$.

#### Staker Level

From the conceptual level, a Staker's withdrawable shares should also be proportionally slashed so the following must be true:

$$
a_{n+1} = a_n p_n
$$

We don't want to update storage at the Staker level during slashing as this would be computationally too expensive given an operator has a 1-many relationship with its delegated stakers.

Therefore we want to prove $a_{n+1} = a_n p_n$ since withdrawable shares are slashed by $p_n$ given the following:

$l_{n+1} = l_n$ \
$k_{n+1} = k_n$ \
$s_{n+1} = s_n$


because we don’t want to update EigenPodManager,StrategyManager storage.

Expanding the $a_{n+1}$ equation

$$
a_{n+1} = s_{n+1} k_{n+1} l_{n+1} m_{n+1}
$$

$$
=> s_{n} k_{n} l_{n} m_{n+1}
$$

We know that  $p_n = \frac {m_{n+1}}{m_n}$ =>  $m_{n+1} = m_n p_n$ 

$$
=>  s_n k_n l_n m_n p_n
$$

$$
=> a_n p_n
$$

Which is exactly as wanted so a Staker's withdrawable shares are immediately affected upon their operator's maxMagnitude being slashed(decreased).


See implementation in:
* [`AllocationManager.slashOperator`](../../../src/contracts/core/AllocationManager.sol)
* [`DelegationManager.burnOperatorShares`](../../../src/contracts/core/DelegationManager.sol)

<br>

---

### Queue Withdrawal

Withdrawals are queued by inputting a `depositShares` amount $x_n <= s_n$ which corresponds to the deposit shares amount stored in $s_n$.
The actual withdrawable amount $w_n$ corresponding to $x_n$ is given by the following:

$$
 w_n = x_n k_n l_n m_n
$$

This conceptually makes sense as the amount being withdrawn $w_n$ is some amount <= $a_n$ which is the total withdrawable shares amount for the Staker. 


#### Operator Level

The operator shares are reduced accordingly by the reduced delegated shares of the Staker.

$$
 op_{n+1} = op_n - w_n
$$


#### Staker Level

$$
 a_{n+1} = a_n - w_n
$$

$$
 s_{n+1} = s_n - x_n
$$

The DelegationManager will call the EigenPodManager/StrategyManager to decrement the depositShares the staker is withdrawing.
We want to show that the withdrawable shares for the Staker are decreased accordingly where $a_{n+1} = a_n - w_n$.

$$
 a_{n+1} = s_{n+1} k_{n+1} l_{n+1} m_{n+1}
$$


$$
 =>  (s_{n} - x_n) k_{n+1} l_{n+1} m_{n+1}
$$

$$
 =  (s_{n} - x_n) k_n l_n m_n
$$

$$
 = s_n k_n l_n m_n - x_n k_n l_n m_n
$$

$$
 = a_n - w_n
$$

Within the `Withdrawal` struct, scaled shares are defined as $q_n = x_n k_n$ where $n$ is the time of the queuing. The reason we define and store scaled shares like this will be clearer in the 'Complete Withdrawal' section below.

See implementation in:
* [`DelegationManager.queueWithdrawals`](../../../src/contracts/core/DelegationManager.sol)
* [`SlashingLib.scaleForQueueWithdrawal`](../../../src/contracts/libraries/SlashingLib.sol)
* [`SlashingLib.calcWithdrawable`](../../../src/contracts/libraries/SlashingLib.sol)

<br>

---

### Complete Withdrawal

Now the staker completes a withdrawal $(q_t, t)$ which was queued at time $t$.

#### Operator Level

The operator shares were already decremented at the time of withdrawal and remain unchanged.


#### Staker Level

There are no storage updates for the staker outside of needing to calculate the shares to send the staker.

The shares that were attempted to be withdrawn by the staker is equal is $w_t$

$$
w_t = q_t m_t l_t
$$

$$
= x_t k_t l_t m_t
$$

However, the staker's shares in their withdrawal may have been slashed both from in EigenLayer during the queued withdrawal period and from the BeaconChain (if the Strategy is native ETH). The amount of shares they actually receive is proportionally the following:

$$
    \frac{m_{t+delay} l_{now} }{m_t l_t}
$$

So the actual amount of shares withdrawn on completion is calculated to be:

$$
sharesWithdrawn = w_t (\frac{m_{t+delay} l_{now}}{m_t l_t} )
$$

$$ 
= x_t k_t l_t m_t (\frac{m_{t+delay} l_{now}}{m_t l_t} )
$$

$$ 
= x_t k_t m_{t+delay} l_{now}
$$

Now we know that $q_t = x_t k_t$ so we can substitute this value in here. 

$$ 
= q_t m_{t+delay} l_{now}
$$

From the above equations the known values we have during the time of queue withdrawal is $x_t k_t$ and we only know $m_{t+delay} l_{now}$ when the queued withdrawal is completable. This is why we store scaled shares as $q_t = x_t k_t$ and $m_{t+delay} l_{now}$ has to be later read during the completing transaction of the withdrawal.

Note: Reading $m_{t+delay}$ is performed by a historical Snapshot lookup of the max magnitude in the `AllocationManager` while $l_{now}$, the current beacon chain slashing factor, is done through the EigenPodManager(default to value of 1 if the Strategy is not native ETH).

The definition of scaled shares is used solely for handling withdrawals and accounting for slashing
that may have occurred(both on EigenLayer and on the BeaconChain) during the queue period.

See implementation in:
* [`DelegationManager.completeQueuedWithdrawal`](../../../src/contracts/core/DelegationManager.sol)
* [`SlashingLib.scaleForCompleteWithdrawal`](../../../src/contracts/libraries/SlashingLib.sol)

<br>

---

### EigenPod BeaconChain Slashing (Negative Shares decrements)

Accounting handling Beacon Chain slashing, aka negative share delta's for EigenPods, is handled differently after the Slashing upgrade with the introduction of $l_n$ the Beacon Chain slashing factor. Prior to the upgrade, any decreases in an EigenPod balance for a staker as a result of completing a checkpoint immediately decrements from the Staker's shares in the `EigenPodManager`. The introduction of $l_n$ the Beacon Chain slashing factor allows us to slash from both the Staker's shares in the `EigenPodManager` and also any queued withdrawals. 
The below diagram helps visualizing this.

![.](../../images/slashing-model.png)

Now lets consider when Beacon Chain slashing occurs and we have a total negative delta in a Staker's EigenPod.


#### Added Definitions

$welw$ is `withdrawableRestakedExecutionLayerGwei` \
Note: this is what is withdrawable in the EigenPod without considering any slashing that has occurred in EigenLayer. `DelegationManager.getWithdrawableShares` can be called  to account for EigenLayer and Beacon Chain slashing. \
$before\text{ }complete$ is time just before the completeCheckpoint transaction is included in the chain \
$before\text{ }start$ is time just before the the checkpoint is started (whether voluntarily or involuntarily) \
$after\text{ }complete$ is the time just after the checkpoint is completed \
For a completed checkpoint that results in a decrease from \
$g_n = welw_{before\text{ }complete}+\sum_i validator_i.balance_{before\text{ }start}$ \
$h_n = welw_{after\text{ }complete}+\sum_i validator_i.balance_{after\text{ }complete}$


#### Staker Level

From a conceptual level, the above logic specifies that we decrease the owned share by the proportion slashed

$$
a_{n+1} = \frac{h_n}{g_n}a_n
$$

We implement this by setting

$$
l_{n+1}=\frac{h_n}{g_n}l_n
$$

Clearly, $m_{n+1}=m_n$ (since we don’t want to affect the operator’s other stakers) and we keep $s_{n+1} = s_n$ **(no subtraction)** and $k_{n+1}=k_n$:

$$
a_{n+1} = s_{n+1}k_{n+1}l_{n+1}m_{n+1}
$$

$$
=s_nk_n\frac{h_n}{g_n}l_nm_n
$$

$$
= \frac{h_n}{g_n}a_n
$$

as wanted.

#### Operator Level

Now we want to update the operator's shares accordingly. At a conceptual level $op_{n+1}$ should be the following:

$$
 op_{n+1} = op_n - a_n + a_{n+1}
$$

We can simplify this further


$$
 =op_{n}-s_nk_nl_nm_n + s_nk_nl_{n+1}m_n
$$


$$
 = op_{n}+s_nk_nm_n(l_{n+1}-l_n)
$$

See implementation in:
* [`EigenPod`](../../../src/contracts/pods/EigenPod.sol)
* [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](../../../src/contracts/pods/EigenPodManager.sol)
* [`DelegationManager.decreaseDelegatedShares`](../../../src/contracts/core/DelegationManager.sol)

## Implementation Details

In practice, we can’t actually have floating values so we will substitute all $k_n, l_n, m_n$ terms with $m_n$/1e18  $\frac{k_n}{1e18},\frac{l_n}{1e18} ,\frac{m_n}{1e18}$ respectively where $k_n, l_n, m_n$ are the values in storage, all initialized to 1e18. This allows us to conceptually have values in the range $[0,1]$.

We make use of OpenZeppelin's Math library and `mulDiv` for calculating $floor(\frac{x \cdot y}{denominator})$ with full precision. Sometimes for specific rounding edge cases, $ceiling(\frac{x \cdot y}{denominator})$ is explicitly used.

#### Multiplication and Division Operations
For all the equations in the above document, we substitute any product operations of $k_n, l_n, m_n$ with the `mulWad` pure function.
```solidity
function mulWad(uint256 x, uint256 y) internal pure returns (uint256) {
    return x.mulDiv(y, WAD);
}
```

Conversely, for any divisions of $k_n, l_n, m_n$ we use the `divWad` pure function.

```solidity
function divWad(uint256 x, uint256 y) internal pure returns (uint256) {
    return x.mulDiv(WAD, y);
}
```
