# Use Cases

## Slashing

### Useful Terminology

**OperatorSet** \
OperatorSets provide a way for AVSs to uniquely identify subsets of their total set of registered operators. They are more strictly defined as the following struct:
```solidity
struct OperatorSet {
   address avs;
   uint32 operatorSetId;
}
```

**Proportion/Magnitude/Allocation** \
This is referring to the proportional values indicating how much slashable stake, for a given Strategy, an operator has allocated to a specific operatorSet. We call this magnitude/allocation/proportion interchangeably.


### AVSs

1. As an AVS, I can register and deregister operators to operatorSets
2. As an AVS, I can reward subsets of my AVS's stakers/operators. There sets are known
as operatorSets.
3. As an AVS, I can seamlessly migrate my existing operators to operator sets without the need for their signatures.
4. As an AVS, I can reward my operatorSets in order to incentivize them to stake slashable funds.
5. As an AVS, I can view the proportion of stake that an operator has allocated to one-or-many of my operatorSets for slashing.
6. As an AVS, I can slash a proportion of a operator's slashable funds in a list of certain strategies.
7. As an AVS, I have time to update to new changes made by operators to the slashable stake they have allocated to my AVS
8. As an AVS, I can reward the earners in my operatorSets proportional to their slashable stake

### Operators

1. As an operator, I can register for a specific AVS's operatorSet by providing a signature
2. As an operator, I can unilaterally deregister from an operatorSet
3. As an operator, I would like to set and adjust my commission on a per operatorSet basis for all the AVSs I am serving
4. As an operator, I can set a proportion of my funds that are available for slashing to a specific operatorSet. The rest is non-slashable.
5. As an operator, I would like to receive rewards from providing slashable stake allocations to an operatorSet.
6. As an operator, after registering with an AVS I can set my slashable proportions for each strategy that I am willing to provide as slashable to a specific AVS.
   1. This proportion is used as the “slashable stake weight” and is used in reward calculations for slashable stake.
   2. This proportion is the maximum proportion that can be slashed by the AVS
   3. This proportion is only uniquely slashable by this AVS and is unaffected from other AVSs slashing the same operator.
7. As an operator, I can unilaterally “rage quit” an AVS.
   1. During a protocol defined period, my stake remains slashable by the AVS
   2. After this period ends, my stake is no longer slashable by the AVS

### Stakers

1. As a staker, I would like to view operator commission rates to inform my decision(at least partially) on delegation
2. As a staker, I would like to be protected from my delegated operator “baiting” me with a low initial commission that they then instantly ramp up
3. As a staker, I can expect to earn rewards from multiple AVSs if my stake is being used to secure different operatorSets from different AVSs
4. As a staker, when withdrawing funds for my shares I must wait for the withdrawal/unstake period, where I am still subject to slashing.
5. As a staker, I want an enforced window in which to undelegate from my Operator, before they allocate slashable stake to a new operatorSet. This should guarantee that if I act reasonably quickly to undelegate, my funds cannot be slashed by the added operatorSet.
6. As a staker, this enforced window is one-time configured and fixed by operators and can help inform my decision on delgation.
7. As a staker, I can claim rewards that I earn for the slashable stake I am providing to operatorSets.