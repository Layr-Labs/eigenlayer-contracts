# Use Cases

## AVS

1. As an AVS, I can register and deregister operators to operatorSets
2. As an AVS, I would like a unified payment interface for Eigenlayer, that minimizes the contract upgrades required for me to integrate with it
3. As an AVS, I can reward sets of stakers/operators that are

    1. doing a particular task
    2. staking a particular token

    These sets are known as operatorSets.
4. As an AVS, I can reward operators (and their stakers) within an operatorSet uniformly or proportional to their delegated stake
5. As an AVS, I can seamlessly migrate my operators to operator sets. 


### Operator

1. As an operator, I can register for a specific AVS's operatorSet by providing a signature
2. As an operator, I can unilaterally deregister from an operatorSet
3. As an operator, I would like to set and adjust my commission on a per reward type, per operatorSet basis for all the AVSs I am serving

### Staker

1. As a staker, I would like to view operator commission rates to inform my decision(at least partially) on delegation
2. As a staker, I would like to be protected from my delegated operator “baiting” me with a low initial commission that they then instantly ramp up
3. As a staker, I can expect to earn rewards from multiple AVSs if my stake is being used to secure different operatorSets from different AVSs
