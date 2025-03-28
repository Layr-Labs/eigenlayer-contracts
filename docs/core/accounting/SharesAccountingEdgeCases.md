[elip-002]: https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md

# Shares Accounting Edge Cases

This document is meant to explore and analyze the different mathematical operations we are performing in the slashing release. Primarily we want to ensure safety on rounding and overflow situations. Prior reading of the [Shares Accounting](./SharesAccounting.md) is required to make sense of this document.

## Prior Reading

* [ELIP-002: Slashing via Unique Stake and Operator Sets](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md)
* [Shares Accounting](./SharesAccounting.md)


## Fully Slashed for a Strategy

Within the context of a single Strategy, recall that updates to the deposit scaling factor $k_n$ are defined as the following:

$$
k_{n+1} = \frac{s_n k_n m_n + d_n}{s_{n+1} l_{n+1} m_{n+1}}=\frac{s_n k_n l_n m_n + d_n}{(s_n+d_n)l_nm_n}
$$

We can see here that calculating $k_{n+1}$ can give us a divide by 0 error if any of $(s_n + d_n)$, $l_n$, or $m_n$ are equal to 0. The $(s_n + d_n) = 0$ case should not arise because the `EigenPodManager` and `StrategyManager` will not report share increases in this case. However, the other two terms may reach 0:
* When an operator is 100% slashed for a given strategy and their max magnitude $m_n = 0$
* When a staker's `EigenPod` native ETH balance is 0 _and_ their validators have all been slashed such that $l_n = 0$

In these cases, updates to a staker's deposit scaling factor will encounter a division by 0 error. In either case, we know that since either the operator was fully slashed or the staker was fully slashed for the `beaconChainETHStrategy` then their withdrawable shares $a_n = 0$.

In practice, if $m_n = 0$ for a given operator, then:
1. Any staker who is already delegated to this operator _will be unable to deposit additional assets into the corresponding strategy_ 
2. Any staker that currently holds deposit shares in this strategy and is NOT delegated to the operator _will be unable to delegate to the operator_

Note that in the first case, it _is_ possible for the staker to undelegate, queue, and complete withdrawals - though as $a_n = 0$, they will not receive any withdrawable shares as a result.

Additionally, if $l_n = 0$ for a given staker in the beacon chain ETH strategy, then **any further deposits of ETH or restaking of validators will not yield shares in EigenLayer.** This should only occur in extraordinary circumstances, as a beacon chain slashing factor of 0 means that a staker both has ~0 assets in their `EigenPod`, and ALL of their validators have been ~100% slashed on the beacon chain - something that happens only when coordinated groups of validators are slashed. If this case occurs, an `EigenPod` is essentially bricked - the pod owner should NOT send ETH to the pod, and should NOT point additional validators at the pod.

These are all expected edge cases and their occurrences and side effects are within acceptable tolerances.

## Upper Bound on Deposit Scaling Factor $k_n$

Let's examine potential overflow situations with respect to calculating a staker's withdrawable shares.
Below is the function in `SlashingLib.sol` which calculates $a_n = s_nk_nl_nm_n$. \
Note: `slashingFactor` = $l_nm_n$

```solidity
function calcWithdrawable(
    DepositScalingFactor memory dsf,
    uint256 depositShares,
    uint256 slashingFactor
) internal pure returns (uint256) {
    /// forgefmt: disable-next-item
    return depositShares
        .mulWad(dsf.scalingFactor())
        .mulWad(slashingFactor);
}
```

`depositShares` are the staker’s shares $s_n$ in storage. We know this can at max be 1e38 - 1 as this is the max total shares we allow in a strategy. $l_n ≤ 1e18$ and $m_n ≤ 1e18$ as they are montonically decreasing values. So a `mulWad` of the `slashingFactor` operation should never result in a overflow, it will always result in a smaller or equal number.

The question now comes to `depositShares.mulWad(dsf.scalingFactor())` and whether this term will overflow a `uint256`. Let's examine the math behind this. The function `SlashingLib.update` performs the following calculation:

$$
k_{n+1} =\frac{s_n k_n l_n m_n + d_n}{(s_n+d_n)l_nm_n}
$$

Assuming:
- $k_0 = 1$
- 0 < $l_0$ ≤ 1 and is monotonically decreasing but doesn’t reach 0
- 0 < $m_0$ ≤ 1 and is monotonically decreasing but doesn’t reach 0
- 0 ≤ $s_n, {s_{n+1}}$ ≤ 1e38 - 1 (`MAX_TOTAL_SHARES = 1e38 - 1` in StrategyBase.sol)
- 0 < $d_n$ ≤ 1e38 - 1
- ${s_{n+1}}={s_n} + {d_n}$

Rewriting above we can get the following by factoring out the k and cancelling out some terms.

$$
k_{n+1} = k_n\frac{s_n}{s_n + d_n} + \frac{d_n}{(s_n+d_n)l_nm_n}
$$

The first term  $\frac{s_n}{{{s_n} + {d_n}}}$ < 1 so when multiplied with $k_n$ will not contribute to the growth of ${k_{n+1}}$ if only considering this term. 

The second term $\frac{d_n}{({{s_n} + {d_n}}){l_n}{m_n}}$ however can make $k_n$ grow over time depending on how small ${l_n}{m_n}$ becomes and also how large $d_n$ is proportionally compared to $s_n$. We only care about the worst case scenario here so let’s assume the upper bound on the existing shares and new deposit amount by rounding the value up to 1.

Now in practice, the smallest values ${l_n}$ and ${m_n}$ could equal to is 1/1e18. Substituting this in the above second term gives the following:

$$
\frac{d_n}{(s_n+d_n)l_nm_n} = \frac{d_n}{s_n+d_n}*1e18^2
$$

So lets round up the first term $\frac{s_n}{{{s_n} + {d_n}}}$ to 1 and also $\frac{d_n}{{{s_n} + {d_n}}}$ in the second term to 1. We can simplify the recursive definition of k in this worst case scenario as the following.

$$
k_{n+1} = k_n\frac{s_n}{s_n + d_n} + \frac{d_n}{(s_n+d_n)l_nm_n}
$$

$$
=> k_{n+1} = k_n+ \frac{d_n}{(s_n+d_n)l_nm_n}
$$

$$
=> k_{n+1} = k_n + 1e36
$$

Because of the max shares in storage for a strategy is 1e38 - 1 and deposits must be non-zero we can actually come up with an upper bound on ${k_n}$ by having 1e38-1 deposits of amount 1, updating ${k_n}$ each time.

$$
k_{1e38-1} \approx (1e38-1)\cdot 1e36 < 1e74
$$

After 1e38-1 iterations/deposits, the upper bound on k we calculate is 1e74 in the _worst_ case scenario. This is technically possible if as a staker, you are delegated to an operator for the beaconChainStrategy where your operator has been slashed 99.9999999…% for native ETH but also as a staker you have had proportional EigenPod balance decreases up to 99.9999999…..%.

The max shares of 1e38-1 also accommodates the entire supply of ETH as well (only needs 27 bits). For normal StrategyManager strategies,  ${l_n} = 1$ and ${k_n}$ would not grow nearly to the same extent.

Clearly this value of 1e74 for ${k_n}$ fits within a uint256 storage slot.

Bringing this all back to the `calcWithdrawable` method used to calculate your actual withdrawable shares for a staker as well as the actual next ${k_{n+1}}$ value. We can see here that the shares is not expected to overflow given the constraints on all our variables and the use of the depositScalingFactor is safe.


The staker depositScalingFactor is unbounded on how it can increase over time but because of the lower bounds we have  ${l_n}$ and  ${m_n}$ as well as the upper bound on number of shares a strategy has (or amount of ETH in existence w.r.t beaconChainStrategy) we can see that it is infeasble for the deposit scaling factor $k_n$ to overflow in our contracts.  



## Rounding Behavior Considerations

The `SlashingLib.sol` introduces some small rounding precision errors due to the usage of `mulWad`/`divWad` operations in the contracts where we are doing a `x * y / denominator` operation. In Solidity, we round down to the nearest integer introducing an absolute error of up to 1 wei. Taking this into consideration, in certain portions of code, we will explicitly use either take the floor or ceiling value of `x * y / denominator`.

This has implications on several parts of the system. For example, completing a withdrawal as shares and having your updated withdrawable shares being less than what it was originally due to rounding. For stakers having a non-WAD beacon chain slashing factor(BCSF) this is essentially self induced from being penalized/slashed on the BC. For operator's have non-WAD maxMagnitudes for specific strategies, it is also a result of them being slashed by the OperatorSet(s) they are allocated to. Stakers should be wary of delegating to operators of low maxMagnitude for the strategies they they have deposits in. The impact of rounding error can result in a larger discrepancy between what they _should_ have withdrawable vs what they actually can withdraw.

### Rounding up on Slashing

When an operator is slashed by an operatorSet in the `AllocationManager`, we actually want to round up on slashing. Rather than calculating `floor(x * y / denominator)` from mulDiv, we want `ceiling(x * y / denominator)`. This is because we don’t want any kind of DOS scenario where an operatorSet attempting to slash an operator is rounded to 0; potentially possible if an operator registered for their own fake AVS and slashed themselves repeatedly to bring their maxMagnitude to a small enough value. This will ensure an operator is always slashed for some amount from their maxMagnitude which eventually, if they are slashed enough, can reach 0.

`AllocationManager.slashOperator`
```solidity
// 3. Calculate the amount of magnitude being slashed, and subtract from
// the operator's currently-allocated magnitude, as well as the strategy's
// max and encumbered magnitudes
uint64 slashedMagnitude = uint64(uint256(allocation.currentMagnitude).mulWadRoundUp(params.wadsToSlash[i]));
```

### Deposits actually _reducing_ withdrawableShares

There are some very particular edge cases where, due to rounding error, deposits can actually decrease withdrawble shares for a staker which is conceptually wrong.
The unit test `DelegationUnit.t.sol:test_increaseDelegatedShares_depositRepeatedly` exemplifies this where there is an increasing difference over the course of multiple deposits between a staker's withdrawable shares and the staker's delegated operator shares.
Essentially, what’s happening in this test case is that after the very first deposit of a large amount of shares, subsequent deposits of amount 1000 are causing the getWithdrawable shares to actually decrease for the staker.

Since the operatorShares are simply incrementing by the exact depositShares, the operatorShares mapping is increasing as expected. This ends up creating a very big discrepancy/drift between the two values after performing 1000 deposits. The difference between the operatorShares and the staker’s withdrawableShares ends up being `4.418e13`.

Granted the initial deposit amount was `4.418e28` which is magnitudes larger than the discrepancy here but this its important to note the side effects of the redesigned accounting model.
Instead of purely incremented/decremented amounts, we have introduced magnitudes and scaling factor variables which now result in small amounts of rounding error from division in several places. We deem this rounding behavior to be tolerable given the costs associated for the number of transactions to emulate this and the proportional error is very small.

### Slashing and Rounding Up Operator Shares and Rounding down on Staker Withdrawable Shares

As can be observed in the `SlashingLib.sol` library, we round up on the operatorShares when slashing and round down on the staker's withdrawableShares. If we look at a core invariant of the shares accounting model, we ideally want to preserve the following:

$$
op_n = \sum_{i=1}^{k} a_{n,i}
$$

where $op_n$ is the operatorShares at time $n$ and $a_{n,i}$ is the staker's withdrawableShares at time $n$ for the $i^{th}$ staker.

However due to rounding limitations, there will be some error introduced in calculating the amount of operator shares to slash above and also in calculating the staker's withdrawableShares. To prevent a situation where all stakers were to attempt to withdraw and the operatorShares underflows, we round up on the operatorShares when slashing and round down on the staker's withdrawableShares.

So in practice, the above invariant becomes.

$$
op_n \geq \sum_{i=1}^{k} a_{n,i}
$$

Upwards rounding on calculating the amount of operatorShares to give to an operator after slashing is intentionally performed in `SlashingLib.calcSlashedAmount`.
For calculating a staker's withdrawableShares, there are many different factors to consider such as calculating their depositScalingFactor, their slashingFactor, and calculating the amount of withdrawable shares altogether with their depositShares. These variables are all by default rounded down in calculation and is expected behavior for stakers.


## Upper bound on Residual Operator Shares

Related to the above rounding error on deposits, we want to calculate what is the worst case rounding error for a staker depositing shares into EigenLayer.
That is, what is the largest difference between the depositShares deposited and the resulting withdrawableShares? For a staker who initially deposits without getting slashed, these two values should conceptually be equal. Let's examine below.

Below is a code snippet of `SlashingLib.sol`
```solidity
function update(
    DepositScalingFactor storage dsf,
    uint256 prevDepositShares,
    uint256 addedShares,
    uint256 slashingFactor
) internal {
    // If this is the staker's first deposit, set the scaling factor to
    // the inverse of slashingFactor
    if (prevDepositShares == 0) {
        dsf._scalingFactor = uint256(WAD).divWad(slashingFactor);
        return;
    }

...

function calcWithdrawable(
    DepositScalingFactor memory dsf,
    uint256 depositShares,
    uint256 slashingFactor
) internal pure returns (uint256) {
    /// forgefmt: disable-next-item
    return depositShares
        .mulWad(dsf.scalingFactor())
        .mulWad(slashingFactor);
}
```

Mathematically, withdrawable shares can be represented as below

$$
withdrawableShares = d\space\cdot\space \frac{k}{WAD} \space\cdot\space \frac{slashingFactor}{WAD}
$$

Substituting $k$ with `WAD.divWad(slashingFactor)` (see update function above) if the staker only has done one single deposit of amount $d$. Also expanding out slashingFactor which is `maxMagnitude.mulWad(beaconChainScalingFactor)`

$$
= d\space\cdot\space \frac{\frac{WAD\space\cdot \space WAD}{m_{deposit} \cdot l_{deposit}}}{WAD} \space\cdot\space \frac{\frac{m \space\cdot\space l}{WAD}}{WAD}
$$

Above is the real true value of the amount of withdrawable shares a staker has but in practice, there are rounding implications at each division operation. It becomes the following

$$
withdrawableShares (rounded) =
\lfloor
\lfloor 
d \space\cdot\space 
\frac{\lfloor\frac{WAD\space\cdot \space WAD
}{m_{deposit} \space\cdot\space l_{deposit}}
\rfloor }{WAD}
\rfloor 
\space\cdot\space \frac{\lfloor \frac{m \space\cdot\space l}{WAD}\rfloor}{WAD}
\rfloor
$$

Each floor operation can introduce a rounding error of at most 1 wei. Because there are nested divisions however, this error can result in a total error thats larger than just off by 1 wei.
We can rewrite parts of above with epsilon $e$ which is in the range of [0,1].

1. First inner rounded term

$$
\frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit}} = \lfloor \frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit}} \rfloor + \epsilon_1
$$

$$
\frac{\lfloor \frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit}} \rfloor}{WAD} = \frac{\frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit}} - \epsilon_1}{WAD}
$$

2. Second rounded term

$$
\lfloor d \cdot \frac{\lfloor \frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit}} \rfloor}{WAD} \rfloor
$$

$$
= \lfloor d \cdot \frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit} \cdot WAD} - d \cdot \frac{\epsilon_1}{WAD} \rfloor
$$

$$
= d \cdot \frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit} \cdot WAD} - d \cdot \frac{\epsilon_1}{WAD} - \epsilon_2
$$

3. Third rounded term

$$
\lfloor \frac{m \cdot l}{WAD} \rfloor = \frac{m \cdot l}{WAD}  - \epsilon_3
$$

$$
=>
\frac{\lfloor \frac{m \cdot l}{WAD} \rfloor}{WAD} = \frac{\frac{m \cdot l}{WAD} - \epsilon_3}{WAD}
$$

$$
=>
\frac{\lfloor \frac{m \cdot l}{WAD} \rfloor}{WAD} = \frac{m \cdot l}{WAD^2} - \frac{\epsilon_3}{WAD}
$$

4. Now bringing it all back to the original equation

$$
withdrawableShares (rounded) =
\lfloor
\lfloor 
d \space\cdot\space 
\frac{\lfloor\frac{WAD\space\cdot \space WAD
}{m_{deposit} \space\cdot\space l_{deposit}}
\rfloor }{WAD}
\rfloor 
\space\cdot\space \frac{\lfloor \frac{m \space\cdot\space l}{WAD}\rfloor}{WAD}
\rfloor
$$

$$
= \lfloor\left(d \cdot \frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit} \cdot WAD} - d \cdot \frac{\epsilon_1}{WAD} - \epsilon_2\right)\cdot\left(\frac{m \cdot l}{WAD^2} - \frac{\epsilon_3}{WAD}\right)\rfloor
$$

$$
= \left(
d \cdot \frac{WAD \cdot WAD}{m_{deposit} \cdot l_{deposit} \cdot WAD} - d \cdot \frac{\epsilon_1}{WAD} - \epsilon_2
\right)
\cdot
\left(
\frac{m \cdot l}{WAD^2} - \frac{\epsilon_3}{WAD}
\right) - \epsilon_4
$$

After expansion and some simplification

$$
withdrawableShares (rounded) =
d \cdot \frac{m\cdot l}{m_{deposit} \cdot l_{deposit}\cdot WAD} - d \cdot \frac{\epsilon_1 \cdot m \cdot l}{WAD^3} - \frac{\epsilon_2 \cdot m \cdot l}{WAD^2} - d \cdot \frac{\epsilon_3}{m_{deposit} \cdot l_{deposit} } + \text{(higher-order terms)}
$$

Note that (higher-order terms) are the terms with multiple epsilon terms where the amounts become negligible, because each term $e$ is < 1.

The true value term is the following:

$$
withdrawableShares = d\space\cdot\space \frac{\frac{WAD \space\cdot\space WAD}{m_{deposit} \cdot l_{deposit}}}{WAD} \space\cdot\space \frac{\frac{m \space\cdot\space l}{WAD}}{WAD}
$$

$$
= d\space\cdot\space \frac{WAD }{m_{deposit} \cdot l_{deposit}}\space\cdot\space \frac{m \space\cdot\space l}{WAD^2}
$$

$$
d \cdot \frac{m\cdot l}{m_{deposit } \cdot l_{deposit}\cdot WAD}
$$

But we can see this term show in the withdrawableShares(rounded) above in the first term! Then we can see that we can represent the equations as the following. 

$$
withdrawableShares (rounded) =
withdrawableShares - d \cdot \frac{\epsilon_1 \cdot m \cdot l}{WAD^3} - \frac{\epsilon_2 \cdot m \cdot l}{WAD^2} - d \cdot \frac{\epsilon_3 }{m_{deposit} \cdot l_{deposit} } + \text{(higher-order terms)}
$$

This intuitively makes sense as all the rounding error comes from the epsilon terms and how they propagate out from being nested. Therefore the introduced error from rounding are all the rounding terms added up ignoring the higher-order terms.

$$
roundedError =d \cdot \frac{\epsilon_1 \cdot m \cdot l}{WAD^3} + \frac{\epsilon_2 \cdot m \cdot l}{WAD^2} + d \cdot \frac{\epsilon_3 }{m_{\text{deposit}} \cdot l_{deposit} }
$$

Now lets assume the worst case scenario of maximizing this sum above, if each epsilon $e$ is replaced with the value of 1 due to a full wei being rounded off we can get the following.

$$
d \cdot \frac{m \cdot l}{WAD^3} + \frac{ m \cdot l}{WAD^2} + \frac{ d}{m_{\text{deposit}} \cdot l_{deposit}}
$$

Assuming close to max values that results in rounding behaviour, we can maximize this total sum by having $d = 1e38$ ,  $m, m_{deposit}, l, l_{deposit}$ equal to WAD(1e18) then we get the following:

$$
\frac{1e38\cdot WAD^2}{WAD^3} + \frac{ WAD^2}{WAD^2} + \frac{1e38}{1e36}
$$

$$
=> \frac{1e38}{1e18} + 1 + 100
$$

$$
\approx 1e20
$$

Framed in another way, the amount of loss a staker can have is $\frac{1}{1e18}$ th of the deposit amount. This makes sense as a result of having nested flooring operations that are then multiplied against outer terms.
Over time, as stakers deposit and withdraw, they may not receive as many shares as their “real” withdrawable amount as this is rounded down and there could be residual/dust shares amount in the delegated operatorShares mapping AND in the original Strategy contract.
This is known and we specifically round down to avoid underflow of operatorShares if all their delegated stakers were to withdraw.