# Rewards

Notably we don't include a foundation usecase due to changes in the foundation rewards flow.

These functions are additional to the existing functions in the Rewards protocol.

```solidity
interface IRewardsCoordinator {
    
    /// EVENTS

    event OperatorCommissionUpdated(
        address indexed operator,
        IAVSDirectory.OperatorSet indexed operatorSet
        uint16 indexed newCommissionBips,
        uint32 effectTimestamp,
    );

    /// EXTERNAL - STATE MODIFYING

    /**
     * @notice Sets the commission an operator takes in bips for a given operatorSet
     *
     * @param operatorSet The operatorSet to update commission for
     * @param commissionBips The commision, in bips
     *
     * @dev The commission can range from 1 to 10000
     * @dev The commission update takes effect after 7 days
     * @dev Default commission for all operators is 1000 bips (10%)
     */
    function setOperatorCommission(
        IAVSDirectory.OperatorSet calldata operatorSet,
        uint16 commissionBips
    ) external returns (uint32 effectTimestamp);
}
```

### setOperatorCommission

Sets the commission an operator takes in bips for a given operatorSet. Commission is active after 7 days. Notably, this gives stakers the guarantee that the commission updates they see when delegating will remain for 7 days, but does not give them the ability to withdraw before new updates occur.

Reverts if:

1. The commission bips are not between \[1,10000]
   1. The minimum of 1 is used to interpret 0 as the default 1000 bip commission
2. The OperatorSet does not exist in the AVSDirectory
