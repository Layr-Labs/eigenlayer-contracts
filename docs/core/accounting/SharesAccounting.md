
[magnitude-doc]: https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md
[elip-002]: https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md

# Shares Accounting

## Prerequisite Documents

## Terminology

The word "shares" in EigenLayer has historically referred to the amount of shares a Staker receives upon depositing assets either through the `StrategyManager` or `EigenPodManager`. Outside of some conversion ratios in the `StrategyManager` to account for rebasing tokens, shares roughly correspond 1:1 with deposit amounts (i.e. 1e18 shares in the `beaconChainETHStrategy` corresponds to 1 ETH of assets). When delegating to an operator or queueing a withdrawal, the `DelegationManager` reads deposit shares from the `StrategyManager` or `EigenPodManager` to determine how many shares to delegate (or undelegate).

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

We can write this a bit more formally with the following:

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
The remaining portions of this document will assume understanding of Allocations/Deallocations, and Max Magnitudes, and OperatorSets.
For more information on this, there is the [ELIP-002][elip-002] which has a high-level but detailed overview of the Slashing upgrade as well as the separate [Magnitude document][magnitude-doc] here.

We now introduce a few new types of Shares concepts:

1. **deposit shares**: \
    For a Staker, this is the amount of Strategy shares deposited
2. **withdrawable shares**: \
    For a Staker, this is the actual amount of Strategy shares they are eligible to withdraw. \
    This does not live in storage but is read through the view function `DelegationManager.getWithdrawableShares`. Note that this amount is <= deposit shares as the Staker may have had their shares slashed.
3. **operator/delegated shares**: \
    This still remains the same definition as before, the amount of delegated shares of an operator from all their delegated stakers.
    However, this is now equal to the summation of all their staker's withdrawable shares.

Notice that these definitions also apply to the shares model prior to the Slashing upgrade but with the caveat that for all Stakers, withdrawable shares equals the deposit shares. After the Slashing upgrade this is not neccesarily the case if a Staker's delegated Operator were slashed resulting in less withdrawable shares for the Staker. \
Now lets look at these updated definitions in detail and how the accounting math works with deposits, withdrawals, and slashing.

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

#### Operator Level

The operator magnitude doesn’t change.

$$
m_{n+1} = m_n
$$

For the operator,

$$
op_{n+1} = op_n+d_n
$$

#### Staker Level

From the conceptual level,

$$
a_{n+1} = a_n + d_n
$$

$$
s_{n+1} = s_n +d_n
$$

$$
l_{n+1} = l_n
$$

Expanding the $a_{n+1}$ calculation

$$
s_{n+1} k_{n+1} l_{n+1} m_{n+1} = s_n k_n l_n m_n + d_n
$$

Which yields

$$
k_{n+1} = \frac{s_n k_n m_n + d_n}{s_{n+1} l_{n+1} m_{n+1}}=\frac{s_n k_n l_n m_n + d_n}{(s_n+d_n)m_n}
$$


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

However, since we don't overwrite `operatorShares` directly in storage and perform increments/decrements we will calculate the amount of $sharesToDecrement$.

$$
    sharesToDecrement = op_n - op_{n+1}
$$

$$
    = op_n - op_n \frac {m_{n+1}} {m_n}
$$

which is exactly how we calculate sharesToDecrement in our library `SlashingLib.sol`

```solidity
function calcSlashedAmount(
    uint256 operatorShares,
    uint256 prevMaxMagnitude,
    uint256 newMaxMagnitude
) internal pure returns (uint256) {
    // round up mulDiv so we don't overslash
    return operatorShares - operatorShares.mulDiv(newMaxMagnitude, prevMaxMagnitude, Math.Rounding.Up);
}
```

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


### Queue Withdrawal

Withdrawals are queued by inputting a `depositShares` amount $x_n$ which corresponds to the amount stored in $s_n$.
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

The DelegationManager will tell the EigenPodManager/StrategyManager to decrement the depositShares the staker is withdrawing.
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