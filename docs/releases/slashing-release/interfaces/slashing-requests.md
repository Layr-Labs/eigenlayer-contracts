# Slashing Requests

TODO: EVENTS

```solidity
interface IAVSDirectory {
    /// EVENTS

    /// @notice emitted when slashOperator is called
    event OperatorSlashed(
        address operator,
        OperatorSet operatorSet,
        uint32 bipsToSlash,
        IStrategy[] strategies,
        uint64[] newTotalMagnitudes
    );

    /// EXTERNAL - STATE MODIFYING

    /**
     * @notice Called by an AVS to slash an operator for given operatorSetId, list of strategies, and bipsToSlash.
     * For each given (operator, operatorSetId, strategy) tuple, bipsToSlash
     * bips of the operatorSet's slashable stake allocation will be slashed
     *
     * @param operator the address to slash
     * @param operatorSetId the ID of the operatorSet the operator is being slashed on behalf of
     * @param strategies the set of strategies to slash
     * @param bipsToSlash the number of bips to slash, this will be proportional to the
     * operator's slashable stake allocation for the operatorSet
     *
     * @dev emits OperatorSlashed
     */
    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint32 bipsToSlash
    ) external;

    /// VIEW

    /**
     * @notice fetches the minimum slashable shares for a certain operator and operatorSet for a list of strategies
     * from the current timestamp until the given timestamp
     *
     * @param operator the operator to get the minimum slashable shares for
     * @param operatorSet the operatorSet to get the minimum slashable shares for
     * @param strategies the strategies to get the minimum slashable shares for
     * @param timestamp the timestamp to the minimum slashable shares before
     *
     * @dev used to get the slashable stakes of operators to weigh over a given slashability window
     *
     * @return the list of share amounts for each strategy
     */
    function getMinimumSlashableSharesBefore(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint256[] calldata);
}
```

### slashOperator

Called by an AVS to slash an operator for a given array of strategies, the corresponding operatorSet to slash from, and the `bipsToSlash`. The bips are with respect to the slashable stake allocation that has been set for the operatorSet, operator, and strategy e.g. `bipsToIncrease = 5000`  leads to half of the slashable stake that has been allocated the operatorSet on behalf of the operator being slashed.

For accounting reasons, this function also reduces the magnitude of the slashing operatorSet in the future and pending deallocations from the operatorSet in order maintain pending nominal stake guarantees in future forecasts. Overall, the total magnitude for the (operator, strategy) before the request is greater than the total magnitude after. 

Emits a `OperatorSlashed` event.

Reverts if:

1. `bipsToSlash == 0 || bipsToSlash > 10000`
2. `operator` is not registered or within 17.5 days of deregistering from an operatorSet

### getMinimumSlashableSharesBefore

Fetches the minimum slashable shares for a certain operator and operatorSet for a list of strategies from the current timestamp until the given timestamp. This called by AVSs to forecast their slashable stake over a slashability window.

Reverts if:

1. `timestamp` is more than 17.5 days in the future. Since the withdrawal period is 17.5 days, longer intervals cannot be forecasted.