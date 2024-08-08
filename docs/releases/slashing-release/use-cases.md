# Use Cases

## Slashing

### AVSs

1. As an AVS, I can reward sets of stakers/operators that are on my AVS in order to incentivize them to stake slashable funds. These sets are known as operatorSets.
2. As an AVS, I can view the proportion of stake that an operator has allocated for slashing by my AVS
3. As an AVS, I can slash a proportion of an operator's slashable funds in a list of certain strategies.
4. As an AVS, I have time to update to new changes made by operators to the slashable stake they have allocated to my AVS
5. As an AVS, I can reward the earners in my operatorSets proportional to their slashable stake

### Operators

1. As an operator, I can set proportion of my funds that are available for slashing. The rest is non-slashable.
2. As an operator, I would like to receive rewards that are proportional to the slashable stake I am allocating slashing by an operatorSet
3. As an operator, after registering with an AVS I can set my slashable proportions for each strategy that I am willing to provide as slashable to a specific AVS.
   1. This proportion is used as the “slashable stake weight” and is used in reward calculations for slashable stake.
   2. This proportion is the maximum proportion that can be slashed by the AVS
4. As an operator, I can unilaterally “rage quit” an AVS.
   1. During a protocol defined period, my stake remains slashable by the AVS
   2. After this period ends, my stake is no longer slashable by the AVS

### Stakers

1. As a staker, when withdrawing funds for my shares I must wait for the withdrawal/unstake period, where I am still subject to slashing.
2. As a staker, I want an enforced window in which to undelegate from my Operator, before they allocate slashable stake to a new operatorSet. This should guarantee that if I act reasonably quickly to undelegate, my funds cannot be slashed by the added operatorSet.
3. As a staker, I can claim rewards that I earn for the slashable stake I am providing to operatorSets.
4. As a staker, I expect to earn more if my stake is being used to secure additional operatorSets
