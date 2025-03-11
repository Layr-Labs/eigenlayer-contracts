# StrategyBase Accounting

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
