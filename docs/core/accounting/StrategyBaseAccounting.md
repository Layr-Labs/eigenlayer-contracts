# StrategyBase Accounting

- [StrategyBase Accounting](#strategybase-accounting)
  - [Overview](#overview)
  - [Why are shares needed?](#why-are-shares-needed)
  - [Can I mess up the accounting? Say, with a direct transfer?](#can-i-mess-up-the-accounting-say-with-a-direct-transfer)
    - [Arbitrary direct transfers](#arbitrary-direct-transfers)
      - [Simple Scenario](#simple-scenario)
      - [Scenario with More Depositors](#scenario-with-more-depositors)
    - [Inflation attacks](#inflation-attacks)
      - [Exploit Scenario](#exploit-scenario)
      - [Mitigation: Virtual Shares](#mitigation-virtual-shares)
      - [Attack Variation: Diluting the Virtual Depositor](#attack-variation-diluting-the-virtual-depositor)
      - [Attack Variation: Flash Loans](#attack-variation-flash-loans)
      - [Mitigation Side Effects](#mitigation-side-effects)
  - [Conclusion](#conclusion)

## Overview

The `StrategyBase` contract is used to manage the accounting of deposit shares for a specific token by collecting tokens and producing shares. Shares represent a proportional claim to the `StrategyBase`'s token balance, ensuring that users can withdraw their intended amount of tokens.

This document serves *specifically* to describe the accounting behavior for the `StrategyBase` contract. General documentation on the `StrategyBase` contract can be found [here](../StrategyManager.md#strategybase).

## Why are shares needed?

At first glance, one may wonder why we need shares to be minted and burned when a staker deposits and withdraws tokens. Why not just track the token balance directly?

The primary reason is **rebase tokens**. Rebase tokens are tokens whose supply can change. An example of a rebase token's behavior is as follows:

* A user holds 1000 tokens
* The token undergoes a 2x rebase (doubling what `balanceOf` returns)
* The token balance of the user is now 2000 tokens

If we were to track the token balance directly, then this 2x increase would not be reflected by the user's withdrawable amount.

Consider the following scenario, where a user deposits a rebase token into EigenLayer:

* A user deposits 1000 tokens, and this value is recorded
* The token undergoes a 2x rebase
* The user's original 1000 tokens are now worth 2000 tokens
* The user *can only withdraw 1000 tokens*, as the recorded deposit is 1000 tokens

**This is where shares come in.** Shares represent a proportional claim to the `StrategyBase`'s token balance, ensuring that users can withdraw their intended amount of tokens.

When a user deposits tokens, they receive shares proportional to the amount of tokens they deposited. Similarly, when a user withdraws tokens, they receive tokens proportional to the amount of shares they burned. Even though the underlying token balance may change, the number of shares a user holds will always represent their proportion of the underlying `StrategyBase` token balance.

With shares, the above scenario would play out as follows:

* A user deposits 1000 tokens, and they receive 1000 shares. Assume this is the first deposit for this strategy.
* The token undergoes a 2x rebase
* The user's original 1000 tokens are now worth 2000 tokens
* The user's 1000 shares are *also* now worth 2000 tokens
* The user can withdraw 2000 tokens, as expected

In short, shares allow for tracking a user's proportional claim to the `StrategyBase`'s token balance, ensuring that users can withdraw their intended amount of tokens even in the presence of rebase or other token behavior.

## Can I mess up the accounting? Say, with a direct transfer?

*TL;DR: No, you cannot arbitrarily mess up the accounting with a direct transfer -- at least, not without leveraging or destroying much capital to do so with minimal benefit.*

### Arbitrary direct transfers

As mentioned before, shares represent a proportional claim to the `StrategyBase`'s token balance. When a user deposits tokens, they receive shares proportional to the amount of tokens they deposited. Similarly, when a user withdraws tokens, they receive tokens proportional to the amount of shares they burned.

However, one can impact the accounting by sending tokens directly to the `StrategyBase` contract. For example:

* A user deposits 1000 tokens, receiving 1000 shares. Assume this is the first deposit for this strategy.
* The user sends 1000 tokens directly to the `StrategyBase` contract
* The user still has 1000 shares, but the `StrategyBase` now has 2000 tokens
* The user can now withdraw 2000 tokens, as their shares are 100% of the total number of shares, so they can withdraw 100% of the `StrategyBase`'s token balance

In this case, even though the user circumvented the expected flow, they still receive their intended amount of tokens.

This is due to the concept of an ***exchange rate***, which is the ratio of total shares to total tokens. This exchange rate by default is 1:1, but implicitly changes as the number of shares deviates from the number of tokens. In the above scenario, the exchange rate is now 2 tokens : 1 share.

However, can this behavior be used by an attacker to profit, or otherwise impact the accounting? Let's walk through a scenario.

#### Simple Scenario

Consider Alice and Bob, who are both aspirational stakers for EigenLayer. Alice has 1000 tokens, and Bob has 1000 tokens.

Alice deposits 500 tokens into the `StrategyBase` contract for this token, the first staker to do so. She receives 500 shares in return, a 1:1 ratio. Presume for demonstration purposes that this is not a rebase token, and that no other users are interacting with this `StrategyBase` contract.

The current state:
* `StrategyBase` total shares: **500**
* `StrategyBase` token balance: **500**
* Alice's shares: **500**
* Alice's "deserved" token balance: **500**
* Bob's shares: 0
* Bob's "deserved" token balance: 0

Bob is about to deposit 500 tokens into the `StrategyBase` contract, but Alice notices this. To increase her proportion of the shares, she sends 500 tokens *directly* to the `StrategyBase` contract. She does not go through the `deposit` function; she does not receive any additional shares for depositing these tokens into the contract.

The current state:
* `StrategyBase` total shares: 500
* `StrategyBase` token balance: **1000**
* Alice's shares: 500
* Alice's "deserved" token balance: **1000**
* Bob's shares: 0
* Bob's "deserved" token balance: 0

As we can see, Alice deposited a net total of 1000 tokens, and she is entitled to those 1000 tokens. Even though she only has 500 shares, she still has 100% of the total shares, and therefore is entitled to 100% of the `StrategyBase`'s token balance.

Bob continues with his day, depositing 500 tokens into the `StrategyBase` contract. However, he does not receive 500 shares, but instead receives 250 shares. This is fewer shares than he expected -- is this an issue?

The current state:
* `StrategyBase` total shares: **750**
* `StrategyBase` token balance: **1500**
* Alice's shares: 500
* Alice's "deserved" token balance: 1000
* Bob's shares: **250**
* Bob's "deserved" token balance: **500**

Let's investigate the accounting -- specifically, the ratio of shares to tokens.

As we can see, the `StrategyBase` contract for this token has 750 shares, and 1500 tokens. The ratio of tokens to shares is 2:1. Here you'll notice the concept of an ***exchange rate***, which is the ratio of total shares to total tokens. This exchange rate by default is 1:1, but implicitly changes as the number of shares deviates from the number of tokens.

Does this exchange rate correctly reflect what Alice and Bob are entitled to?

* Alice's 500 shares * 2 tokens / 1 share = 1000 tokens. **Correct!**
* Bob's 250 shares * 2 tokens / 1 share = 500 tokens. **Correct!**

Both Alice and Bob are still able to withdraw their "deserved" token amounts, as their shares are still correctly proportional in relation to each other. As such, despite the direct transfer of tokens to the `StrategyBase` contract, the accounting is still correct.

#### Scenario with More Depositors

Consider the same scenario, but with more depositors. Alice and Bob each deposit 500 tokens, receiving 500 shares each.

The current state:
* `StrategyBase` total shares: **1000**
* `StrategyBase` token balance: **1000**
* Alice's shares: **500**
* Alice's "deserved" token balance: **500**
* Bob's shares: **500**
* Bob's "deserved" token balance: **500**

A third entity, Charlie, prepares to deposit 500 tokens into the `StrategyBase` contract. However, Alice notices this. To increase her proportion of the shares, she sends 500 tokens *directly* to the `StrategyBase` contract. She does not go through the `deposit` function; she does not receive any additional shares for depositing these tokens into the contract.

The current state:
* `StrategyBase` total shares: **1500**
* `StrategyBase` token balance: **2000**
* Alice's shares: 500
* Alice's "deserved" token balance: **1000**
* Bob's shares: 500
* Bob's "deserved" token balance: 500
* Charlie's shares: 0
* Charlie's "deserved" token balance: 0

As we can see, the exchange rate is now 4 tokens : 3 shares.

As such:
* Alice's 500 shares * 4 tokens / 3 shares = 666 tokens. ***Lower* than the expected 1000 tokens.**
* Bob's 500 shares * 4 tokens / 3 shares = 666 tokens. ***Higher* than the expected 500 tokens. (Good for Bob!)**

From the onset, Bob benefits from Alice's direct deposit. He is still entitled to the same proportion of tokens, but the total tokens have increased. Alice, on the other hand, is worse off, as she is entitled to fewer tokens than she overall deposited or transferred.

As Charlie deposits 500 tokens, he receives 375 shares due to the exchange rate, as 500 tokens * 3 shares / 4 tokens = 375 shares.

The current state:
* `StrategyBase` total shares: **1875**
* `StrategyBase` token balance: **2500**
* Alice's shares: 500
* Alice's "deserved" token balance: 1000
* Bob's shares: 500
* Bob's "deserved" token balance: 500
* Charlie's shares: **375**
* Charlie's "deserved" token balance: **500**

As we can see, the exchange rate is now 4 tokens : 3 shares.

We can see that Charlie's 375 shares are correctly worth 500 tokens, as 375 shares * 4 tokens / 3 shares = 500 tokens. *Note that, since the exchange rate hasn't changed, we do not need to recalculate Alice and Bob's eligible tokens.*

Therefore, even though the exchange rate was not 1:1 prior to Charlie's deposit, the accounting is still correct, and he receives his intended amount of tokens. The only victim here is Alice, who effectively donated tokens to the existing depositors of the `StrategyBase` contract.

### Inflation attacks

An inflation attack is a more *specific* scenario that may impact depositors during the first deposits into a `StrategyBase` contract. This kind of attack is possible when the `StrategyBase` contract is first created, before or after the very first deposit. At this stage, the exchange rate of the `StrategyBase` is highly susceptible to manipulation, giving the first depositor the ability to steal funds from later depositors.

#### Exploit Scenario

Say Alice deposits 1 token into a `StrategyBase` contract, the first depositor to do so, and receives 1 share. This is an intentionally minimal amount so that she can perform an inflation attack.

The current state:
* `StrategyBase` total shares: **1**
* `StrategyBase` token balance: **1**
* Alice's shares: **1**
* Alice's "deserved" token balance: **1**

She notices that Bob is about to deposit 1000 tokens into the `StrategyBase` contract. She wants to manipulate the exchange rate to be so high that, due to rounding, Bob receives no shares for his deposit. In other words, she wants to set the exchange rate to some value where the number of tokens that Bob will be depositing is less than the number of tokens that would be required to receive 1 share. This would leave Bob with no shares, and no way to withdraw his tokens.

For example, say that Alice sends a *million* tokens to the `StrategyBase` contract. She does not go through the `deposit` function; she does not receive any additional shares for depositing these tokens into the contract. The exchange rate is now 1e6 + 1 tokens : 1 share.

The current state:
* `StrategyBase` total shares: 1
* `StrategyBase` token balance: **1e6 + 1**
* Alice's shares: **1**
* Alice's "deserved" token balance: **1e6 + 1**

Note the large difference between the `StrategyBase`'s token balance and the number of shares. As mentioned before, the exchange rate is now 1e6 + 1 tokens : 1 share.

When Bob deposits 1000 tokens, he is depositing less than 1e6 tokens, meaning that *he receives no shares for his deposit*. Calculating Bob's total shares:

* Bob's 1000 tokens * 1 share / 1e6 + 1 tokens = 1e-3 shares = 0 shares

Due to the large divisor, Bob's received shares are now 1e-3, a very small number, which is *rounded down to 0* due to EVM division. Thus, **Bob receives no shares for his token deposit**. This is a problem, as he is entitled to 1000 tokens, but has no shares for withdrawing it.

The current state:
* `StrategyBase` total shares: 1
* `StrategyBase` token balance: **1e6 + 1001**
* Alice's shares: 1
* Alice's "deserved" token balance: 1e6 + 1
* Bob's shares: 0
* Bob's "deserved" token balance: **1000**

As we can see, Bob's token balance increased, but his shares remained at 0, losing his deposited tokens to Alice. Alice has all the shares of the `StrategyBase` contract, and can withdraw all of the tokens, including Bob's 1000 tokens, even though she does not "deserve" those tokens as she is not the rightful owner.

#### Mitigation: Virtual Shares

To mitigate this, **we use a "virtual shares" mechanism**. Every created `StrategyBase` contract is initialized with a certain number of virtual shares (1e3) and virtual tokens (1e3), which simulate the "first deposit" into this `StrategyBase` contract. This prevents a first depositor from manipulating the exchange rate to their benefit, as they lose the advantages typically associated with the first depositor.

Consider Alice trying to perform the same attack as before, but with the virtual shares mechanism. The virtual shares are now 1e3, and the virtual tokens are 1e3. The total shares and token balance reflect this.

The current state:
* `StrategyBase` total shares: **1,000**
* `StrategyBase` token balance: **1,000**
* Alice's shares: 0
* Alice's "deserved" token balance: 0

Alice deposits her 1 token into the `StrategyBase` contract. She receives 1 share, as expected.

The current state:
* `StrategyBase` total shares: **1,001**
* `StrategyBase` token balance: **1,001**
* Alice's shares: **1**
* Alice's "deserved" token balance: **1**

Note immediately that Alice has 1 share, which is less than 0.1% of the total shares. She no longer has 100% of the shares due to this "virtual depositor".

Bob again intends to deposit 100 tokens into the `StrategyBase` contract. However, Alice beats him to it, depositing 1 million tokens into the `StrategyBase` contract. She does not receive any additional shares for depositing these tokens into the contract.

The current state:
* `StrategyBase` total shares: **1,001**
* `StrategyBase` token balance: **1,001,001**
* Alice's shares: **1**
* Alice's "deserved" token balance: **1,000,001**

Remember how Alice only has 1 share? Notice how she cannot withdraw the million tokens she deposited!

* Alice's 1 share * 1,000,001 tokens / 1,001 shares = 999 tokens. ***Lower* than the expected 1,000,001 tokens.**

Given the virtual depositor, Alice is "donating" her tokens to the `StrategyBase` contract, and is not able to withdraw the majority of the tokens she deposited.

Bob now deposits his 1000 tokens, at an exchange rate of ~1000 tokens : 1 share. Given that the exchange rate is now 1000 tokens : 1 share, Bob receives 1 share for his deposit, as he has deposited enough tokens to not have his shares rounded down to 0.

The current state:
* `StrategyBase` total shares: **1,002**
* `StrategyBase` token balance: **1,002,001**
* Alice's shares: 1
* Alice's "deserved" token balance: 1,000,001
* Bob's shares: **1**
* Bob's "deserved" token balance: **1,000**

Hilariously enough, not only is Bob able to withdraw his 1000 tokens, but Alice is also *only* able to withdraw 1000 tokens! Alice's attack is rendered unsuccessful.

#### Attack Variation: Diluting the Virtual Depositor

What if Alice attempts to dilute the initial 1e3 virtual shares and tokens? Clearly, since she didn't have enough shares with her minimal initial deposit, her attack's capital was dilulted by the "virtual depositor".

Imagine that Alice instead deposits 1 million tokens upfront. She receives 1 million shares, as expected.

The current state:
* `StrategyBase` total shares: **1,001,000**
* `StrategyBase` token balance: **1,001,000**
* Alice's shares: **1,000,000**
* Alice's "deserved" token balance: **1,000,000**

As you can see here, Alice has a vast majority of the shares. She tries again to manipulate the exchange rate before Bob deposits by depositing yet *another* million tokens. Remember that she does not receive any additional shares for depositing these tokens into the contract.

The current state:
* `StrategyBase` total shares: 1,001,000
* `StrategyBase` token balance: **2,001,000**
* Alice's shares: 1,000,000
* Alice's "deserved" token balance: **2,000,000**

The exchange rate is now ~2 tokens : 1 share. More accurately, it is 2,001,000 tokens : 1,001,000 shares, or ~1.999 tokens : 1 share.

So let's ask the question: how many tokens can Alice withdraw given her 1 million shares?

* Alice's 1 million shares * 2,001,000 tokens / 1,001,000 shares = 1,999,000 tokens. **Lower* than the expected 2,000,000 tokens.**

As we can see, Alice is not able to withdraw her intended amount of tokens, even though she has a vast majority of the shares. She actually *loses* 1000 tokens due to the virtual depositor.

Let's see what happens when Bob deposits his 1000 tokens. Given the exchange rate of ~2 tokens : 1 share, we expect him to receive 500 shares for his deposit. We calculate this as follows:

* Bob's 1000 tokens * 1,001,000 shares / 2,001,000 tokens = 500.25 shares = 500 shares (due to rounding)

The current state:
* `StrategyBase` total shares: **1,001,500**
* `StrategyBase` token balance: **2,002,000**
* Alice's shares: 1,000,000
* Alice's "deserved" token balance: 2,000,000
* Bob's shares: **500**
* Bob's "deserved" token balance: **1,000**

As we can see, Bob receives 500 shares for his deposit, as expected. If he attempts to withdraw:

* Bob's 500 shares * 2,002,000 tokens / 1,001,500 shares = 999.5 tokens = 999 tokens (due to rounding). ***Correct!**

Alice, the attempted attacker, is the one who loses in this scenario. Bob loses one token, but Alice loses 1000 tokens, and locks up significant capital for her troubles. An attacker **is not economically incentivized** to perform an inflation attack, as they will lose out in the end.

As such, the virtual depositor is a useful mechanism to protect against inflation attacks, even when the attacker has a vast majority of the shares.

#### Attack Variation: Flash Loans

You may be wondering, what if Alice performs a flash loan attack? These provide large amounts of capital on demand, and perhaps enough capital can make an inflation attack profitable.

First, similar to how Alice lost capital to the virtual depositor in the previous scenario, she will lose capital when performing the flash loan attack. This alone prevents her attack from being profitable, even in the best case scenario.

Say that Alice is, for lack of a better term, "insane" and chooses to disobey economic incentives. Note that typical flash loans only provide capital [within a given transaction](https://aave.com/docs/developers/flash-loans#:~:text=the%20borrowed%20amount%20(and%20a%20fee)%20is%20returned%20before%20the%20end%20of%20the%20transaction), and are not able to be borrowed over any larger unit of time. As such, flash loans are not a viable route for performing this attack in the first place.

#### Mitigation Side Effects

The virtual depositor has a few side effects that are important to note.

* **Rebase dilution:** In the event of a token rebase, user token balances will typically increase by the rebase factor. However, the virtual depositor's token balance will not increase by the same factor, as it is a fixed amount. This means that user gains will be mildly diluted over time.
  * However, as the virtual depositor only has 1e3 shares and tokens, this effect is negligible (estimated to be 1 part in 1e20).
* **Negative rebase:** In the event of a "negative rebase," where the token balance decreases, not all users may be able to withdraw. The `StrategyBase` contract will have more shares than assets due to this loss of principal. As a result, the last depositor(s) will not be able to withdraw. This is because the virtual depositor's shares and tokens are fixed, and are not subject to the loss of principal. Thus, the last withdrawal(s) will attempt to withdraw more tokens than the `StrategyBase` contract has.
  * However, this is expected to occur infrequently, if ever. For example, many rebasing tokens such as LSTs only undergo negative rebases in the event of a beacon chain slash, which is a rare event. Given this minimal impact, we do not consider this a significant issue.

## Conclusion

Shares are a useful mechanism to manage the accounting of a `StrategyBase` contract. They allow for tracking a user's proportional claim to the `StrategyBase`'s token balance, ensuring that users can withdraw their intended amount of tokens even in the presence of rebase or other token behavior.

Typically, this model is vulnerable to an "inflation attack," but the virtual depositor mitigation protects against this. It is a simple and effective mechanism to prevent a first depositor from manipulating the exchange rate to their benefit, as they lose the advantages typically associated with the first depositor.

Any attacker attempting to perform an inflation attack will lose out in the end. Even if they seek to grief other users, the amount of capital required to perform the attack in the first place is extremely high. Though there are small side effects to the virtual depositor, they are negligible and do not impact the core functionality of the `StrategyBase` contract.
