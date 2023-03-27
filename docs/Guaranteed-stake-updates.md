# Design: Withdrawals From EigenLayer -- Guaranteed Stake Updates on Withdrawal
Withdrawals are one of the critical flows in the EigenLayer system.  Guaranteed stake updates ensure that all middlewares that an operator has opted-into (i.e. allowed to slash them) are notified at the appropriate time regarding any withdrawals initiated by an operator.  To put it simply, an operator can "queue" a withdrawal at any point in time.  In order to complete the withdrawal, the operator must first serve all existing obligations related to keeping their stake slashable.  The contract `Slasher.sol` keeps track of a historic record of each operator's  `latestServeUntil` time at various blocks, which is the timestamp after which their stake will have served its obligations which were created at or before the block in question. To complete a withdrawal, an operator (or a staker delegated to them) can point to a relevant point in the record which proves that the funds they are withdrawing are no longer "at stake" on any middleware tasks.
EigenLayer uses a 'push' model for it's own core contracts -- when a staker queues a withdrawal from EigenLayer (or deposits new funds into EigenLayer), their withdrawn shares are immediately decremented, both in the StrategyManager itself and in the DelegationManager contract. Middlewares, however, must 'pull' this data. Their worldview is stale until a call is made that triggers a 'stake update', updating the middleware's view on how much the operator has staked. The middleware then informs EigenLayer (either immediately or eventually) that the stake update has occurred.

## Storage Model

Below, a whitelisted contract refers to a contract that is a part of a middleware that is allowed to freeze the opted-in operators.

For each operator, the Slasher contract stores:

1. A `mapping(address => mapping(address => MiddlewareDetails))`, from operator address to contract whitelisted by the operator to slash them, to [details](https://github.com/Layr-Labs/eignlayr-contracts/blob/master/src/contracts/interfaces/ISlasher.sol) about that contract formatted as
```solidity
    struct MiddlewareDetails {
        // the UTC timestamp before which the contract is allowed to slash the user
        uint32 contractCanSlashOperatorUntil;
        // the block at which the middleware's view of the operator's stake was most recently updated
        uint32 latestUpdateBlock;
    }
```
2. A `mapping(address => LinkedList<address>) operatorToWhitelistedContractsByUpdate`, from operator address to a [linked list](../src/contracts/libraries/StructuredLinkedList.sol) of addresses of all whitelisted contracts, ordered by when their stakes were last updated by each middleware, from 'stalest'/earliest (at the 'HEAD' of the list) to most recent (at the 'TAIL' of the list)
3. A `mapping(address => MiddlewareTimes[]) middlewareTimes` from operators to a historic list of
```solidity
    struct MiddlewareTimes {
        // The update block for the middleware whose most recent update was earliest, i.e. the 'stalest' update out of all middlewares the operator is serving
        uint32 stalestUpdateBlock;
        // The latest 'serve until' time from all of the middleware that the operator is serving
        uint32 latestServeUntil;
    }
```

The reason we store an array of updates as opposed to one `MiddlewareTimes` struct with the most up-to-date values is that this would require pushing updates carefully to not disrupt existing withdrawals. This way, operators can select the `MiddlewareTimes` entry that is appropriate for their withdrawal.  Thus, the operator provides an entry from their `operatorMiddlewareTimes` based on which a withdrawal can be completed. The withdrawability is checked by `slasher.canWithdraw()`, which checks that for the queued withdrawal in question, `withdrawalStartBlock` is less than the provided `operatorMiddlewareTimes` entry's 'stalestUpdateBlock', i.e. the specified stake update occurred *strictly after* the withdrawal was queued.  It also checks that the current block.timestamp is greater than the `operatorMiddlewareTimes` entry's 'latestServeUntil' time, i.e. that the current time is *strictly after* the end of all obligations that the operator had, at the time of the specified stake update.  If these criteria are both met, then the withdrawal can be completed.

Note:
`remove`, `nodeExists`,`getHead`, `getNextNode`, and `pushBack` are all constant time operations on linked lists. This is gained at the sacrifice of getting any elements by specifying their *indices* in the list. Searching the linked list for the correct entry is linear-time with respect to the length of the list; this should only ever happen in a "worst-case" scenario, after the provided index input is determined to be incorrect.

## An Instructive Example

Let us say an operator has opted into serving a middleware, `Middleware A`. As a result of the operator's actions, `Middleware A` calls `recordFirstStakeUpdate`, adding  `Middleware A` to their linked list of middlewares, recording the `block.number` as the `updateBlock` and the middleware's specified `serveUntil` time as the `latestServeUntil` in a `MiddlewareTimes` struct that gets pushed to `operatorMiddlewareTimes`.  At later times, the operator registers with a second and third middleware, `Middleware B` and `Middleware C`, respectively.  At this point, the current state is as follows:

![Three Middlewares Timeline](images/three_middlewares.png?raw=true "Three Middlewares Timeline")

Based on this, the *current* latest serveUntil time is `serveUntil_B`, and the 'stalest' stake update from a middleware occurred at `updateBlock_A`.  So the most recent entry in the `operatorMiddlewareTimes` array for the operator will have `serveUntil = serveUntil_B` and `stalestUpdateBlock = updateBlock_A`.

In the meantime, let us say that the operator had also queued a withdrawal between opting-in to serve `Middleware A` and opting-in to serve `Middleware B`:

![Three Middlewares Timeline With Queued Withdrawal](images/three_middlewares_withdrawal_queued.png?raw=true "Three Middlewares Timeline With Queued Withdrawal")

Now that a withdrawal has been queued, the operator must wait till their obligations have been met before they can withdraw their stake.  At this point, in our example, the `operatorMiddlewareTimes` array looks like this:

```solidity
{
    {
        stalestUpdateBlock: updateBlock_A
        latestServeUntil: serveUntil_A
    },
    {
        stalestUpdateBlock: updateBlock_A
        latestServeUntil: serveUntil_B
    },
    {
        stalestUpdateBlock: updateBlock_A
        latestServeUntil: serveUntil_B
    }
}
```
  In order to complete a withdrawal in this example, the operator would have to record a stake update in `Middleware A`, signalling readiness for withdrawal.  Assuming this update was performed at roughly the time that the operator signed up to serve `Middleware B`, the state would now look like this:

![Updated Three Middlewares Timeline With Queued Withdrawal](images/withdrawal.png?raw=true "Updated Three Middlewares Timeline With Queued Withdrawal")

By recording a stake update in `Middleware A`, a new entry would be pushed to the operator's `operatorMiddlewareTimes` array, with `serveUntil = serveUntil_A` and `stalestUpdateBlock = updateBlock_B`. The queued withdrawal will then become completable after the current value of `serveUntil_A`, by referencing this entry in the array.

## Deep Dive, aka "The Function-by-Function Explanation"

### Internal Functions

#### `_recordUpdateAndAddToMiddlewareTimes`
```solidity
    function _recordUpdateAndAddToMiddlewareTimes(address operator, uint32 updateBlock, uint32 serveUntil) internal {
```

This function is called each time a middleware posts a stake update, through a call to `recordFirstStakeUpdate`, `recordStakeUpdate`, or `recordLastStakeUpdateAndRevokeSlashingAbility`. It records that the middleware has had a stake update and pushes a new entry to the operator's list of 'MiddlewareTimes', i.e. `operatorToMiddlewareTimes[operator]`, if *either* the `operator`'s stalestUpdateBlock' has decreased, *or* their latestServeUntil' has increased. An entry is also pushed in the special case of this being the first update of the first middleware that the operator has opted-in to serving.

### External Functions

#### `recordFirstStakeUpdate`
```solidity
    function recordFirstStakeUpdate(address operator, uint32 serveUntil) external onlyCanSlash(operator) {

```

This function is called by a whitelisted slashing contract during registration of a new operator. The middleware posts an initial update, passing in the time until which the `operator`'s stake is bonded -- `serveUntil`. The middleware is pushed to the end ('TAIL') of the linked list since in `operatorToWhitelistedContractsByUpdate[operator]`, since the new middleware must have been updated the most recently, i.e. at the present moment.


#### `recordStakeUpdate`
```solidity
recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntil, uint256 insertAfter) 
```

This function is called by a whitelisted slashing contract, passing in the time until which the operator's stake is bonded -- `serveUntil`, the block for which the stake update to the middleware is being recorded (which may be the current block or a past block) -- `updateBlock`, and an index specifying the element of the `operator`'s linked list that the currently updating middleware should be inserted after -- `insertAfter`. It makes a call to the internal function `_updateMiddlewareList` to actually update the linked list.

#### `recordLastStakeUpdateAndRevokeSlashingAbility`
```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntil) external onlyCanSlash(operator) {
```

This function is called by a whitelisted slashing contract on deregistration of an operator, passing in the time until which the operator's stake is bonded -- `serveUntil`. It assumes that the update is posted for the *current* block, rather than a past block, in contrast to `recordStakeUpdate`.


### View / "Helper" Functions

#### `canWithdraw`
```solidity
canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) external returns (bool) {
```

The biggest thing guaranteed stake updates do is to make sure that withdrawals only happen once the stake being withdrawn is no longer slashable in a non-optimistic way. This is done by calling the `canWithdraw` function on the Slasher contract, which returns 'true' if the `operator` can currently complete a withdrawal started at the `withdrawalStartBlock`, with `middlewareTimesIndex` used to specify the index of a `MiddlewareTimes` struct in the operator's list (i.e. an index in `operatorToMiddlewareTimes[operator]`). The specified struct is consulted as proof of the `operator`'s ability (or lack thereof) to complete the withdrawal.




















