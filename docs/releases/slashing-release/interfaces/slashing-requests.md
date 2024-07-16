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
        uint64[] slashingRates // the parts per hundred million that have been slashed for the operator since the beginning of the protocol. used to denormalize shares during deposit/withdrawal.
    );

    event MagnitudeDecremented(
        address operator,
        OperatorSet operatorSet,
        IStrategy strategy,
        uint64 updatedSlashableMagnitude,
        uint32 effectTimestamp
    );

    event NonslashableMagnitudeDecremented(
        address operator,
        IStrategy strategy,
        uint64 updatedNonslashableMagnitude,
        uint32 effectTimestamp
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
     * @dev emits MagnitudeDecremented and/or NonslashableMagnitudeDecremented to convey which
     * magnitude has been decremented.
     */
    function slashOperator(
        address operator,
        bytes4 operatorSetId,
        IStrategy[] memory strategies,
        uint32 bipsToSlash
    ) external;

    /// VIEW

    /**
     * @notice fetches the requested parts per hundred million that have been slashed for the operator since the slashing protocol started
     *
     * @param operator the operator to get the requested slashing rate for
     * @param strategy the strategy to get the requested slashing rate for
     * @param operatorSet the operatorSet to get the requested requested slashing rate for
     * @param timestamp the timestamp to get the slashing rate for
     *
     * @dev used during deposits and withdrawals
     *
     * @return slashingRate parts per hundred million that have been slashed for the given
     * operator, strategy, until the given timestamp
     */
    function getSlashedRate(
        address operator,
        IStrategy strategy,
        uint32 timestamp
    ) external view returns (uint64);

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

For accounting reasons, this function also reduces the magnitude of the slashing operatorSet in the future in order maintain pending nominal stake guarantees in future forecasts.

Emits a `OperatorSlashed` event.
Either emits a `MagnitudeDecremented` or/and `NonslashableMagnitudeDecremented` to convey which magnitude have been decremented.

Reverts if:

1. `bipsToSlash == 0 || bipsToSlash > 10000`
2. `operator` is not registered or within 17.5 days of deregistering from an operatorSet

### getSlashedRate

Fetches the parts per hundred million slashed for the given operator and strategy from the protocol started until a given timestamp. This takes into account all operatorSet slashing events for the operator and strategy pairing. This is used when calculating how funds should allowed to be withdrawn during completion, deposits, and in several view functions due to slashing accounting.

The protocol does things in terms of slashing rates (parts per hundred million, 1e8) for these functions for accuracy.

### getMinimumSlashableSharesBefore

Fetches the minimum slashable shares for a certain operator and operatorSet for a list of strategies from the current timestamp until the given timestamp. This called by AVSs to forecast their slashable stake over a slashability window.

Reverts if:

1. `timestamp` is more than 17.5 days in the future. Since the withdrawal period is 17.5 days, longer intervals cannot be forecasted.