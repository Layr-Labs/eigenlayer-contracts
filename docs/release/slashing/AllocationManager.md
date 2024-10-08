# AllocationManager

## Prerequisites

- [The Mechanics of Allocating and Slashing Unique Stake](https://forum.eigenlayer.xyz/t/the-mechanics-of-allocating-and-slashing-unique-stake/13870)

## Overview
The AllocationManager contract manages the allocation and reallocation of operators' slashable stake across various strategies and operator sets. It enforces allocation and deallocation delays and handles the slashing process initiated by AVSs.

## Parameterization

- `ALLOCATION_CONFIGURATION_DELAY`: The delay in seconds before allocations take effect.
    - Mainnet: `21 days`. Very TBD
    - Testnet: `1 hour`. Very TBD
- `DEALLOCATION_DELAY`: The delay in seconds before deallocations take effect.
    - Mainnet: `17.5 days`. Slightly TBD
    - Testnet: `3 days`. Very TBD

## `setAllocationDelay` 

```solidity
/**
 * @notice Called by the delagation manager to set delay when operators register.
 * @param operator The operator to set the delay on behalf of.
 * @param delay The allocation delay in seconds.
 * @dev msg.sender is assumed to be the delegation manager.
 */
function setAllocationDelay(address operator, uint32 delay) external;

/**
 * @notice Called by operators to set their allocation delay.
 * @param delay the allocation delay in seconds
 * @dev msg.sender is assumed to be the operator
 */
function setAllocationDelay(uint32 delay) external;
```

These functions allow operators to set their allocation delay. The first variant is called by the DelegationManager upon operator registration for all new operators created after the slashing release. The second variant is called by operators themselves to update their allocation delay or set it for the first time if they joined before the slashing release.

The allocation delay takes effect in `ALLOCATION_CONFIGURATION_DELAY` seconds.

The allocation delay can be any positive uint32.

The allocation delay's primary purpose is to give stakers delegated to an operator the chance to withdraw their stake before the operator can change the risk profile to something they're not comfortable with.

## `modifyAllocations`

```solidity
/**
 * @notice struct used to modify the allocation of slashable magnitude to list of operatorSets
 * @param strategy the strategy to allocate magnitude for
 * @param expectedTotalMagnitude the expected total magnitude of the operator used to combat against race conditions with slashing
 * @param operatorSets the operatorSets to allocate magnitude for
 * @param magnitudes the magnitudes to allocate for each operatorSet
 */
struct MagnitudeAllocation {
    IStrategy strategy;
    uint64 expectedTotalMagnitude;
    OperatorSet[] operatorSets;
    uint64[] magnitudes;
}

/**
 * @notice Modifies the propotions of slashable stake allocated to a list of operatorSets for a set of strategies
 * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
 * @dev Updates encumberedMagnitude for the updated strategies
 * @dev msg.sender is used as operator
 */
function modifyAllocations(MagnitudeAllocation[] calldata allocations) external
```

This function is called by operators to adjust the proportions of their slashable stake allocated to different operator sets for different strategies.

The operator provides their expected total magnitude for each strategy they're adjusting the allocation for. This is used to combat race conditions with slashings for the strategy, which may result in larger than expected slashable proportions allocated to operator sets.

Each `(operator, operatorSet, strategy)` tuple can have at most 1 pending modification at a time. The function will revert is there is a pending modification for any of the tuples in the input. 

The contract limits keeps track of the total magnitude in pending allocations, active allocations, and pending deallocations. This is called the **_encumbered magnitude_** for a strategy. The contract verifies that the allocations made in this call do not make the encumbered magnitude exceed the operator's total magnitude for the strategy. If the encumbered magnitude exceeds the total magnitude, the function reverts.

Any _allocations_ (i.e. increases in the proportion of slashable stake allocated to an AVS) take effect after the operator's allocation delay. The allocation delay must be set for the operator before they can call this function.

Any _deallocations_ (i.e. decreases in the proportion of slashable stake allocated to an AVS) take after `DEALLOCATION_DELAY` seconds. This enables AVSs enough time to update their view of stakes to the new proportions and have any tasks created against previous stakes to expire.

## `clearPendingModifications`

```solidity
/**
 * @notice This function takes a list of strategies and adds all completable modifications for each strategy, 
 * updating the encumberedMagnitude of the operator as needed.
 *
 * @param operator address to complete modifications for
 * @param strategies a list of strategies to complete modifications for
 * @param numToClear a list of number of pending modifications to complete for each strategy
 *
 * @dev can be called permissionlessly by anyone
 */
function clearModificationQueue(
    address operator,
    IStrategy[] calldata strategies,
    uint16[] calldata numToClear
) external;
```

This function is used to complete pending modifications for a list of strategies for an operator. The function takes a list of strategies and the number of pending modifications to complete for each strategy. For each strategy, the function completes a modification if its effect timestamp has passed. 

Completing an allocation doesn't have any material change to the protocol other than cleaning up some state and increasing code readability. 

Completing a deallocation decreases the encumbered magnitude for the strategy, allowing them to make allocations with that magnitude. Encumbered magnitude must be decreased only upon completion because pending deallocations can be slashed before they are completable.

## `slashOperator`

```solidity
/**
 * @notice Called by an AVS to slash an operator for given operatorSetId, list of strategies, and bipsToSlash.
 * For each given (operator, operatorSetId, strategy) tuple, bipsToSlash
 * bips of the operatorSet's slashable stake allocation will be slashed
 *
 * @param operator the address to slash
 * @param operatorSetId the ID of the operatorSet the operator is being slashed on behalf of
 * @param strategies the set of strategies to slash
 * @param wadToSlash the parts in 1e18 to slash, this will be proportional to the
 * @param description the description of the slashing provided by the AVS for legibility
 * operator's slashable stake allocation for the operatorSet
 */
function slashOperator(
    address operator,
    uint32 operatorSetId,
    IStrategy[] calldata strategies,
    uint256 wadToSlash
) external
```

This function is called by AVSs to slash an operator for a given operator set and list of strategies. The AVS provides the proportion of the operator's slashable stake allocation to slash for each strategy. The proportion is given in parts in 1e18 and is with respect to the operator's _current_ slashable stake allocation for the operator set (i.e. `wadsToSlash=5e17` means 50% of the operator's slashable stake allocation for the operator set will be slashed). The AVS also provides a description of the slashing for legibility by outside integrations.

Slashing is instant and irreversable. Slashed funds remain unrecoverable in the protocol but will be burned/redistributed in a future release. Slashing by one operatorSet does not effect the slashable stake allocation of other operatorSets for the same operator and strategy.

Slashing updates storage in a way that instantly updates all view functions to reflect the correct values.

## View Functions 

### `getMinDelegatedAndSlashableOperatorShares`

```solidity
/**
 * @notice returns the minimum operatorShares and the slashableOperatorShares for an operator, list of strategies, 
 * and an operatorSet before a given timestamp
 * @param operator the operator to get the shares for
 * @param operatorSet the operatorSet to get the shares for
 * @param strategies the strategies to get the shares for
 * @param beforeTimestamp the timestamp to get the shares at
 */
function getMinDelegatedAndSlashableOperatorShares(
    address operator,
    OperatorSet calldata operatorSet,
    IStrategy[] calldata strategies,
    uint32 beforeTimestamp
) external view returns (uint256[] memory, uint256[] memory);
```

This function returns the minimum operator shares and the slashable operator shares for an operator, list of strategies, and an operator set before a given timestamp. This is used by AVSs to pessimistically estimate the operator's slashable stake allocation for a given strategy and operator set within their slashability windows.

### Additional View Functions

See the [AllocationManager Interface](../../../src/contracts/interfaces/IAllocationManager.sol) for additional view functions.