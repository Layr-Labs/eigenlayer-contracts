# AllocationManager

| File | Notes |
| -------- | -------- |
| [`AllocationManager.sol`](../../src/contracts/core/AllocationManager.sol) |  |
| [`AllocationManagerStorage.sol`](../../src/contracts/core/AllocationManagerStorage.sol) | state variables |
| [`IAllocationManager.sol`](../../src/contracts/interfaces/IAllocationManager.sol) | interface |

Libraries and Mixins:

| File | Notes |
| -------- | -------- |
| [`PermissionControllerMixin.sol`](../../src/contracts/mixins/PermissionControllerMixin.sol) | account delegation |
| [`Pausable.sol`](../../src/contracts/permissions/Pausable.sol) | |
| [`SlashingLib.sol`](../../src/contracts/libraries/SlashingLib.sol) | slashing math |
| [`OperatorSetLib.sol`](../../src/contracts/libraries/OperatorSetLib.sol) | encode/decode operator sets |
| [`Snapshots.sol`](../../src/contracts/libraries/Snapshots.sol) | historical state |

## Prior Reading

* [ELIP-002: Slashing via Unique Stake and Operator Sets](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md)

## Overview

The `AllocationManager` manages AVS metadata registration, registration and deregistration of operators to operator sets, handles allocation and slashing of operators' slashable stake, and is the entry point an AVS uses to slash an operator. The `AllocationManager's` responsibilities are broken down into the following concepts:

* [AVS Metadata](#avs-metadata)
* [Operator Sets](#operator-sets)
* [Allocations and Slashing](#allocations-and-slashing)
* [Config](#config)

## Parameterization

* `ALLOCATION_CONFIGURATION_DELAY`: The delay in blocks before allocations take effect.
    * Mainnet: `126000 blocks` (17.5 days).
    * Testnet: `75 blocks` (15 minutes).
* `DEALLOCATION_DELAY`: The delay in blocks before deallocations take effect.
    * Mainnet: `100800 blocks` (14 days).
    * Testnet: `50 blocks` (10 minutes).

---

## AVS Metadata

AVSs must register their metadata to declare themselves who they are as the first step, before they can create operator sets or register operators into operator sets. `AllocationManager` keeps track of AVSs that have registered metadata.

**Methods:**
* [`updateAVSMetadataURI`](#updateavsmetadatauri)


#### `updateAVSMetadataURI`

```solidity
/**
 *  @notice Called by an AVS to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
 *
 *  @param metadataURI The URI for metadata associated with an AVS.
 *
 *  @dev Note that the `metadataURI` is *never stored* and is only emitted in the `AVSMetadataURIUpdated` event.
 */
function updateAVSMetadataURI(
    address avs, 
    string calldata metadataURI
) 
    external
    checkCanCall(avs)
```

_Note: this method can be called directly by an AVS, or by a caller authorized by the AVS. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

Below is the format AVSs should use when updating their metadata URI initially. This is not validated onchain.

```json
{
    "name": "AVS",
    "website": "https.avs.xyz/",
    "description": "Some description about",
    "logo": "http://github.com/logo.png",
    "twitter": "https://twitter.com/avs",
}
```


Later on, once AVSs have created operator sets, content in their metadata URI can be updated subsequently.

```json
{
    "name": "AVS",
    "website": "https.avs.xyz/",
    "description": "Some description about",
    "logo": "http://github.com/logo.png",
    "twitter": "https://twitter.com/avs",
    "operatorSets": [
        {
            "name": "ETH Set",
            "id": "1", // Note: we use this param to match the opSet id in the Allocation Manager
            "description": "The ETH operatorSet for AVS",
            "software": [
                {
                    "name": "NetworkMonitor",
                    "description": "",
                    "url": "https://link-to-binary-or-github.com"
                },
                {
                    "name": "ValidatorClient",
                    "description": "",
                    "url": "https://link-to-binary-or-github.com"
                }
            ],
            "slashingConditions": ["Condition A", "Condition B"]
        },
        {
            "name": "EIGEN Set",
            "id": "2", // Note: we use this param to match the opSet id in the Allocation Manager
            "description": "The EIGEN operatorSet for AVS",
            "software": [
                {
                    "name": "NetworkMonitor",
                    "description": "",
                    "url": "https://link-to-binary-or-github.com"
                },
                {
                    "name": "ValidatorClient",
                    "description": "",
                    "url": "https://link-to-binary-or-github.com"
                }
            ],
            "slashingConditions": ["Condition A", "Condition B"]
        }
    ]
}
```

*Effects*:
* Emits an `AVSMetadataURIUpdated` event for use in offchain services

*Requirements*:
* Caller MUST be authorized, either as the AVS itself or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))


## Operator Sets

Operator sets, as described in [ELIP-002](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md#operator-sets), are useful for AVSs to configure operator groupings which can be assigned different tasks, rewarded based on their strategy allocations, and slashed according to different rules. Operator sets are defined in [`libraries/OperatorSetLib.sol`](../../src/contracts/libraries/OperatorSetLib.sol):

```solidity
/**
 * @notice An operator set identified by the AVS address and an identifier
 * @param avs The address of the AVS this operator set belongs to
 * @param id The unique identifier for the operator set
 */
struct OperatorSet {
    address avs;
    uint32 id;
}
```

The `AllocationManager` tracks operator sets and members of operator sets in the following mappings:

```solidity
/// @dev Lists the operator set ids an AVS has created
mapping(address avs => EnumerableSet.UintSet) internal _operatorSets;

/// @dev Lists the members of an AVS's operator set
mapping(bytes32 operatorSetKey => EnumerableSet.AddressSet) internal _operatorSetMembers;
```

Every `OperatorSet` corresponds to a single AVS, as indicated by the `avs` parameter. On creation, the AVS provides an `id` (unique to that AVS), as well as a list of `strategies` the `OperatorSet` includes. Together, the `avs` and `id` form the `key` that uniquely identifies a given `OperatorSet`. Operators can register to and deregister from operator sets. In combination with allocating slashable magnitude, operator set registration forms the basis of operator slashability (discussed further in [Allocations and Slashing](#allocations-and-slashing)).

**Concepts:**
* [Registration Status](#registration-status)

**Methods:**
* [`createOperatorSets`](#createoperatorsets)
* [`addStrategiesToOperatorSet`](#addstrategiestooperatorset)
* [`removeStrategiesFromOperatorSet`](#removestrategiesfromoperatorset)
* [`registerForOperatorSets`](#registerforoperatorsets)
* [`deregisterFromOperatorSets`](#deregisterfromoperatorsets)

#### Registration Status

Operator registration and deregistration is tracked in the following state variables:

```solidity
/// @dev Lists the operator sets the operator is registered for. Note that an operator
/// can be registered without allocated stake. Likewise, an operator can allocate
/// without being registered.
mapping(address operator => EnumerableSet.Bytes32Set) internal registeredSets;

/**
 * @notice Contains registration details for an operator pertaining to an operator set
 * @param registered Whether the operator is currently registered for the operator set
 * @param slashableUntil If the operator is not registered, they are still slashable until
 * this block is reached.
 */
struct RegistrationStatus {
    bool registered;
    uint32 slashableUntil;
}

/// @dev Contains the operator's registration status for an operator set.
mapping(address operator => mapping(bytes32 operatorSetKey => RegistrationStatus)) internal registrationStatus;
```

For each operator, `registeredSets` keeps a list of `OperatorSet` `keys` for which the operator is currently registered. Each operator registration and deregistration respectively adds and removes the relevant `key` for a given operator. An additional factor in registration is the operator's `RegistrationStatus`.

The `RegistrationStatus.slashableUntil` value is used to ensure an operator remains slashable for a period of time after they initiate deregistration. This is to prevent an operator from committing a slashable offence and immediately deregistering to avoid penalty. This means that when an operator deregisters from an operator set, their `RegistrationStatus.slashableUntil` value is set to `block.number + DEALLOCATION_DELAY`.

#### `createOperatorSets`

```solidity
/**
 * @notice Parameters used by an AVS to create new operator sets
 * @param operatorSetId the id of the operator set to create
 * @param strategies the strategies to add as slashable to the operator set
 */
struct CreateSetParams {
    uint32 operatorSetId;
    IStrategy[] strategies;
}

/**
 * @notice Allows an AVS to create new operator sets, defining strategies that the operator set uses
 */
function createOperatorSets(
    address avs,
    CreateSetParams[] calldata params
)
    external
    checkCanCall(avs)
```

_Note: this method can be called directly by an AVS, or by a caller authorized by the AVS. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

AVSs use this method to create new operator sets. An AVS can create as many operator sets as they desire, depending on their needs. Once created, operators can [allocate slashable stake to](#modifyallocations) and [register for](#registerforoperatorsets) these operator sets.

On creation, the `avs` specifies an `operatorSetId` unique to the AVS. Together, the `avs` address and `operatorSetId` create a `key` that uniquely identifies this operator set throughout the `AllocationManager`.

Optionally, the `avs` can provide a list of `strategies`, specifying which strategies will be slashable for the new operator set. AVSs may create operator sets with various strategies based on their needs - and strategies may be added to more than one operator set.

*Effects*:
* For each `CreateSetParams` element:
    * For each `params.strategies` element:
        * Add `strategy` to `_operatorSetStrategies[operatorSetKey]`
        * Emits `StrategyAddedToOperatorSet` event

*Requirements*:
* Caller MUST be authorized, either as the AVS itself or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* AVS MUST have registered metadata via calling `updateAVSMetadataURI`
* For each `CreateSetParams` element:
    * Each `params.operatorSetId` MUST NOT already exist in `_operatorSets[avs]`
    
#### `addStrategiesToOperatorSet`

```solidity
/**
 * @notice Allows an AVS to add strategies to an operator set
 * @dev Strategies MUST NOT already exist in the operator set
 * @param avs the avs to set strategies for
 * @param operatorSetId the operator set to add strategies to
 * @param strategies the strategies to add
 */
function addStrategiesToOperatorSet(
    address avs,
    uint32 operatorSetId,
    IStrategy[] calldata strategies
)
    external
    checkCanCall(avs)
```

_Note: this method can be called directly by an AVS, or by a caller authorized by the AVS. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

This function allows an AVS to add slashable strategies to a given operator set. If any strategy is already registered for the given operator set, the entire call will fail.

*Effects*:
* For each `strategies` element:
    * Adds the strategy to `_operatorSetStrategies[operatorSetKey]`
    * Emits a `StrategyAddedToOperatorSet` event

*Requirements*:
* Caller MUST be authorized, either as the AVS itself or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* The operator set MUST be registered for the AVS
* Each proposed strategy MUST NOT be registered for the operator set

#### `removeStrategiesFromOperatorSet`

```solidity
/**
 * @notice Allows an AVS to remove strategies from an operator set
 * @dev Strategies MUST already exist in the operator set
 * @param avs the avs to remove strategies for
 * @param operatorSetId the operator set to remove strategies from
 * @param strategies the strategies to remove
 */
function removeStrategiesFromOperatorSet(
    address avs,
    uint32 operatorSetId,
    IStrategy[] calldata strategies
)
    external
    checkCanCall(avs)
```

_Note: this method can be called directly by an AVS, or by a caller authorized by the AVS. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

This function allows an AVS to remove slashable strategies from a given operator set. If any strategy is not registered for the given operator set, the entire call will fail.

*Effects*:
* For each `strategies` element:
    * Removes the strategy from `_operatorSetStrategies[operatorSetKey]`
    * Emits a `StrategyRemovedFromOperatorSet` event

*Requirements*:
* Caller MUST be authorized, either as the AVS itself or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* The operator set MUST be registered for the AVS
* Each proposed strategy MUST be registered for the operator set

#### `registerForOperatorSets`

```solidity
/**
 * @notice Parameters used to register for an AVS's operator sets
 * @param avs the AVS being registered for
 * @param operatorSetIds the operator sets within the AVS to register for
 * @param data extra data to be passed to the AVS to complete registration
 */
struct RegisterParams {
    address avs;
    uint32[] operatorSetIds;
    bytes data;
}

/**
 * @notice Allows an operator to register for one or more operator sets for an AVS. If the operator
 * has any stake allocated to these operator sets, it immediately becomes slashable.
 * @dev After registering within the ALM, this method calls the AVS Registrar's `IAVSRegistrar.
 * registerOperator` method to complete registration. This call MUST succeed in order for 
 * registration to be successful.
 */
function registerForOperatorSets(
    address operator,
    RegisterParams calldata params
)
    external
    onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION)
    checkCanCall(operator)
```

_Note: this method can be called directly by an operator, or by a caller authorized by the operator. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

An operator may call this function to register for any number of operator sets of a given AVS at once. There are two very important details to know about this method:
1. As part of registration, each operator set is added to the operator's `registeredSets`. Note that for each newly-registered set, **any stake allocations to the operator set become immediately slashable**.
2. Once all sets have been added, the AVS's configured `IAVSRegistrar` is called to confirm and complete registration. _This call MUST NOT revert,_ as **AVSs are expected to use this call to reject ineligible operators** (according to their own custom logic). Note that if the AVS has not configured a registrar, the `avs` itself is called.

This method makes an external call to the `IAVSRegistrar.registerOperator` method, passing in the registering `operator`, the `operatorSetIds` being registered for, and the input `params.data` provided during registration. From [`IAVSRegistrar.sol`](../../src/contracts/interfaces/IAVSRegistrar.sol):

```solidity
/**
 * @notice Called by the AllocationManager when an operator wants to register
 * for one or more operator sets. This method should revert if registration
 * is unsuccessful.
 * @param operator the registering operator
 * @param operatorSetIds the list of operator set ids being registered for
 * @param data arbitrary data the operator can provide as part of registration
 */
function registerOperator(address operator, uint32[] calldata operatorSetIds, bytes calldata data) external;
```

*Effects*:
* Adds the proposed operator sets to the operator's list of registered sets (`registeredSets`)
* Adds the operator to `_operatorSetMembers` for each operator set
* Marks the operator as registered for the given operator sets (in `registrationStatus`)
* Passes the `params` for registration to the AVS's `AVSRegistrar`, which can arbitrarily handle the registration request
* Emits an `OperatorAddedToOperatorSet` event for each operator

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION`
* `operator` MUST be registered as an operator in the `DelegationManager`
* Caller MUST be authorized, either the operator themselves, or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* Each `operatorSetId` MUST exist for the given AVS
* Operator MUST NOT already be registered for any proposed operator sets
* If operator has deregistered, operator MUST NOT be slashable anymore (i.e. the `DEALLOCATION_DELAY` must have passed)
* The call to the AVS's configured `IAVSRegistrar` MUST NOT revert

#### `deregisterFromOperatorSets`

```solidity
/**
 * @notice Parameters used to deregister from an AVS's operator sets
 * @param operator the operator being deregistered
 * @param avs the avs being deregistered from
 * @param operatorSetIds the operator sets within the AVS being deregistered from
 */
struct DeregisterParams {
    address operator;
    address avs;
    uint32[] operatorSetIds;
}

/**
 * @notice Allows an operator or AVS to deregister the operator from one or more of the AVS's operator sets.
 * If the operator has any slashable stake allocated to the AVS, it remains slashable until the
 * DEALLOCATION_DELAY has passed.
 * @dev After deregistering within the ALM, this method calls the AVS Registrar's `IAVSRegistrar.
 * deregisterOperator` method to complete deregistration. This call MUST succeed in order for
 * deregistration to be successful.
 */
function deregisterFromOperatorSets(
    DeregisterParams calldata params
)
    external
    onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION)
```

_Note: this method can be called directly by an operator/AVS, or by a caller authorized by the operator/AVS. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

This method may be called by EITHER an operator OR an AVS to which an operator is registered; it is intended to allow deregistration to be triggered by EITHER party. This method generally inverts the effects of `registerForOperatorSets`, with two specific exceptions:
1. As part of deregistration, each operator set is removed from the operator's `registeredSets`. HOWEVER, **any stake allocations to that operator set will remain slashable for `DEALLOCATION_DELAY` blocks.** The operator will not be allowed to register for the operator set again until this slashable window has passed.
2. Once all sets have been removed, the AVS's configured `IAVSRegistrar` is called to complete deregistration on the AVS side.

This method makes an external call to the `IAVSRegistrar.deregisterOperator` method, passing in the deregistering `operator` and the `operatorSetIds` being deregistered from. From [`IAVSRegistrar.sol`](../../src/contracts/interfaces/IAVSRegistrar.sol):

```solidity
/**
 * @notice Called by the AllocationManager when an operator is deregistered from
 * one or more operator sets. If this method reverts, it is ignored.
 * @param operator the deregistering operator
 * @param operatorSetIds the list of operator set ids being deregistered from
 */
function deregisterOperator(address operator, uint32[] calldata operatorSetIds) external;
```

*Effects*:
* Removes the proposed operator sets from the operator's list of registered sets (`registeredSets`)
* Removes the operator from `_operatorSetMembers` for each operator set
* Updates the operator's `registrationStatus` with:
    * `registered: false`
    * `slashableUntil: block.number + DEALLOCATION_DELAY`
        * As mentioned above, this allows for AVSs to slash deregistered operators until `block.number == slashableUntil`
* Emits an `OperatorRemovedFromOperatorSet` event for each operator
* Passes the `operator` and `operatorSetIds` to the AVS's `AVSRegistrar`, which can arbitrarily handle the deregistration request

*Requirements*:
<!-- * Address MUST be registered as an operator -->
* Pause status MUST NOT be set: `PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION`
* Caller MUST be authorized, either the operator/AVS themselves, or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* Each operator set ID MUST exist for the given AVS
* Operator MUST be registered for the given operator sets
* Note that, unlike `registerForOperatorSets`, the AVS's `AVSRegistrar` MAY revert and the deregistration will still succeed

---

## Allocations and Slashing

[Operator set registration](#operator-sets) is one step of preparing to participate in an AVS. When an operator successfully registers for an operator set, it is because the AVS in question is ready to assign them tasks. However, it follows that _before assigning tasks_ to an operator, an AVS will expect operators to allocate slashable stake to the operator set such that the AVS has some economic security.

For this reason, it is expected that many AVSs will require operators to **allocate slashable stake BEFORE registering for an operator set**. This is due to [`registerForOperatorSets`](#registerforoperatorsets) serving in part as an AVS's "consent mechanism," as calling `IAVSRegsitrar.registerOperator` allows the AVS to query the amount of slashable stake the operator can provide when assigned tasks.

It is only once an operator is both _registered for an operator set_ and _has an active allocation to that operator set_ that the associated AVS can slash actual stake from an operator.

See [ELIP-002#Unique Stake Allocation & Deallocation](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md#unique-stake-allocation--deallocation) for additional context.

**Concepts:**
* [Max vs Encumbered Magnitude](#max-vs-encumbered-magnitude)
* [Evaluating the "Current" Allocation](#evaluating-the-current-allocation)
* [Evaluating Whether an Allocation is "Slashable"](#evaluating-whether-an-allocation-is-slashable)

**Methods:**
* [`modifyAllocations`](#modifyallocations)
* [`clearDeallocationQueue`](#cleardeallocationqueue)
* [`slashOperator`](#slashoperator)

#### Max vs Encumbered Magnitude

Operators allocate _magnitude_, which represents a proportion of their total stake. For a given strategy, the `AllocationManager` tracks two quantities, _max magnitude_ and _encumbered magnitude_:

```solidity
/**
 * @notice Contains allocation info for a specific strategy
 * @param maxMagnitude the maximum magnitude that can be allocated between all operator sets
 * @param encumberedMagnitude the currently-allocated magnitude for the strategy
 */
struct StrategyInfo {
    uint64 maxMagnitude;
    uint64 encumberedMagnitude;
}

/// @dev Contains a history of the operator's maximum magnitude for a given strategy
mapping(address operator => mapping(IStrategy strategy => Snapshots.DefaultWadHistory)) internal _maxMagnitudeHistory;

/// @dev For a strategy, contains the amount of magnitude an operator has allocated to operator sets
/// @dev This value should be read with caution, as deallocations that are completable but not
///      popped off the queue are still included in the encumbered magnitude
mapping(address operator => mapping(IStrategy strategy => uint64)) public encumberedMagnitude;
```

An operator's max magnitude starts at `1 WAD` (`1e18`), and is decreased when they are slashed. Max magnitude represents "100%" of allocatable magnitude. When an operator allocates magnitude from a strategy to an operator set, their encumbered magnitude for that strategy increases. An operator cannot allocate > 100%; therefore, a strategy's encumbered magnitude can never exceed that strategy's max magnitude.

#### Evaluating the "Current" Allocation

As mentioned in the previous section, allocations and deallocations take place on a delay, and as such the `Allocation` struct has both a `currentMagnitude`, and `pendingDiff` / `effectBlock` fields:

```solidity
/**
 * @notice Defines allocation information from a strategy to an operator set, for an operator
 * @param currentMagnitude the current magnitude allocated from the strategy to the operator set
 * @param pendingDiff a pending change in magnitude, if it exists (0 otherwise)
 * @param effectBlock the block at which the pending magnitude diff will take effect
 */
struct Allocation {
    uint64 currentMagnitude;
    int128 pendingDiff;
    uint32 effectBlock;
}
```

Although the `allocations` mapping can be used to fetch an `Allocation` directly, you'll notice a convention in the `AllocationManager` of using the `_getUpdatedAllocation` helper, instead. This helper reads an existing `Allocation`, then evaluates `block.number` against `Allocation.effectBlock` to determine whether or not to apply the `pendingDiff`. 
* If the diff can be applied, the helper returns an `Allocation` with an updated `currentMagnitude` and zeroed out `pendingDiff` and `effectBlock` fields -- as if the modification has already been completed.
* Otherwise, the `Allocation` is returned from storage unmodified.

Generally, when an `Allocation` is mentioned in this doc (or used within the `AllocationManager.sol` contract), we are referring to the "Current" `Allocation` as defined above.

#### Evaluating Whether an Allocation is "Slashable"

Given an `operator` and an `Allocation` from a `strategy` to an AVS's `OperatorSet`, the `AllocationManager` uses the following criteria to determine whether the operator's allocation is slashable:
1. The `operator` must be registered for the operator set, or if they are deregistered, they must still be slashable (See [Registration Status](#registration-status)).
2. The AVS must have added the `strategy` to the operator set (See [`addStrategiesToOperatorSet`](#addstrategiestooperatorset) and [`removeStrategiesFromOperatorSet`](#removestrategiesfromoperatorset))
3. The existing `Allocation` must have a nonzero `Allocation.currentMagnitude`

If ALL of these are true, the `AllocationManager` will allow the AVS to slash the `operator's` `Allocation`.

#### `modifyAllocations`

```solidity
/**
 * @notice struct used to modify the allocation of slashable magnitude to an operator set
 * @param operatorSet the operator set to modify the allocation for
 * @param strategies the strategies to modify allocations for
 * @param newMagnitudes the new magnitude to allocate for each strategy to this operator set
 */
struct AllocateParams {
    OperatorSet operatorSet;
    IStrategy[] strategies;
    uint64[] newMagnitudes;
}

/**
 * @notice Modifies the proportions of slashable stake allocated to an operator set from a list of strategies
 * Note that deallocations remain slashable for DEALLOCATION_DELAY blocks therefore when they are cleared they may
 * free up less allocatable magnitude than initially deallocated.
 * @param operator the operator to modify allocations for
 * @param params array of magnitude adjustments for one or more operator sets
 * @dev Updates encumberedMagnitude for the updated strategies
 */
function modifyAllocations(
    address operator, 
    AllocateParams[] calldata params
) 
    external
    onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS)
```

_Note: this method can be called directly by an operator, or by a caller authorized by the operator. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

This function is called by an operator to EITHER increase OR decrease the slashable magnitude allocated from a strategy to an operator set. As input, the operator provides an operator set as the target, and a list of strategies and corresponding `newMagnitudes` to allocate. The `newMagnitude` value is compared against the operator's current `Allocation` for that operator set/strategy:
* If `newMagnitude` is _greater than_ `Allocation.currentMagnitude`, this is an allocation
* If `newMagnitude` is _less than_ `Allocation.currentMagnitude`, this is a deallocation
* If `newMagnitude` is _equal to_ `Allocation.currentMagnitude`, this is invalid (revert)

Allocation modifications play by different rules depending on a few factors. Recall that at all times, the `encumberedMagnitude` for a strategy may not exceed that strategy's `maxMagnitude`. Additionally, note that _before processing a modification for a strategy,_ the `deallocationQueue` for that strategy is first cleared. This ensures any completable deallocations are processed first, freeing up magnitude for allocation. This process is further explained in [`clearDeallocationQueue`](#cleardeallocationqueue). 

Finally, `modifyAllocations` does NOT require an allocation to consider whether its corresponding strategy is relevant to the operator set in question. This is primarily to cut down on complexity. Because [`removeStrategiesFromOperatorSet`](#removestrategiesfromoperatorset) always allows an AVS to _remove_ strategies from consideration, we always need to be sure an operator can initiate a _deallocation_ for such strategies. Although there's not a clear usecase for _allocating_ when a strategy is not included in an operator set, we elected not to check for this. It's possible some AVSs may announce a strategy is being added ahead of time specifically to encourage allocations in advance. **It is expected behavior** that an AVS adding a strategy to an operator set makes any existing allocations to that strategy instantly slashable.

**If we are handling an _increase in magnitude_ (allocation):**

* The increase in magnitude is immediately added to the strategy's `encumberedMagnitude`. This ensures that subsequent _allocations to other operator sets from the same strategy_ will not go above the strategy's `maxMagnitude`.
* The `allocation.pendingDiff` is set, with an `allocation.effectBlock` equal to the current block plus the operator's configured allocation delay.

**If we are handling a _decrease in magnitude_ (deallocation):**

First, evaluate whether the operator's _existing allocation is currently slashable_ by the AVS. This is important because the AVS might be using the existing allocation to secure a task given to this operator. See [Evaluating Whether an Allocation is "Slashable"](#evaluating-whether-an-allocation-is-slashable) for details.

Next, _if the existing allocation IS slashable_:

* The `allocation.pendingDiff` is set, with an `allocation.effectBlock` equal to the current block plus `DEALLOCATION_DELAY + 1`. This means the existing allocation _remains slashable_ for `DEALLOCATION_DELAY` blocks.
* The _operator set_ is pushed to the operator's `deallocationQueue` for that strategy, denoting that there is a pending deallocation for this `(operatorSet, strategy)`. This is an ordered queue that enforces deallocations are processed sequentially and is used both in this method and in [`clearDeallocationQueue`](#cleardeallocationqueue).

Alternatively, _if the existing allocation IS NOT slashable_, the deallocated amount is immediately **freed**. It is subtracted from the strategy's encumbered magnitude and can be used for subsequent allocations. This is the only type of update that does not result in a "pending modification." The rationale here is that if the existing allocation is not slashable, the AVS does not need it to secure tasks, and therefore does not need to enforce a deallocation delay.

Another point of consideration are race conditions involving a slashing event and a deallocation occurring for an operator. Consider the following scenario with an operator having an allocation of 500 magnitude and trying to deallocate setting it to 250. However in the same block _right_ before calling `modifyAllocations` the operator is slashed 100% by the OperatorSet, setting the current magnitude to 0. Now the operator's deallocation is considered an allocation and ends up allocating 250 magnitude when they were trying to _deallocate_. This is a potential griefing vector by malicious AVSs and a known shortcoming. In such scenarios, the operator should simply deallocate all their allocations to 0 so that they don't accidentally allocate more slashable stake. In general for non malicious AVSs, slashing is deemed to be a very occasional occurrence and this race condition to not be impacting to operators.

*Effects*:
* For each `AllocateParams` element:
    * Complete any existing deallocations (See [`clearDeallocationQueue`](#cleardeallocationqueue))
    * Update the operator's `encumberedMagnitude`, `allocations`, and `deallocationQueue` according to the rules described above. Additionally:
        * If `encumberedMagnitude` is updated, emits `EncumberedMagnitudeUpdated`
        * If a pending modification is created:
            * Adds the `strategy` to `allocatedStrategies[operator][operatorSetKey]` (if not present)
            * Adds the `operatorSetKey` to `allocatedSets[operator]` (if not present)
        * If the allocation now has a `currentMagnitude` of 0:
            * Removes `strategy` from the `allocatedStrategies[operator][operatorSetKey]` list
            * If this list now has a length of 0, remove `operatorSetKey` from `allocatedSets[operator]`
    * Emits an `AllocationUpdated` event

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_MODIFY_ALLOCATIONS`
* Caller MUST be authorized, either as the operator themselves or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* Operator MUST have already set an allocation delay (See [`setAllocationDelay`](#setallocationdelay))
* For each `AllocationParams` element:
    * Provided strategies MUST be of equal length to provided magnitudes for a given `AllocateParams` object
    * Operator set MUST exist for each specified AVS
    * Operator MUST NOT have pending, non-completable modifications for any given strategy
    * New magnitudes MUST NOT match existing ones
    * New encumbered magnitude MUST NOT exceed the operator's max magnitude for the given strategy

#### `clearDeallocationQueue`

```solidity
/**
 * @notice This function takes a list of strategies and for each strategy, removes from the deallocationQueue
 * all clearable deallocations up to max `numToClear` number of deallocations, updating the encumberedMagnitude
 * of the operator as needed.
 *
 * @param operator address to clear deallocations for
 * @param strategies a list of strategies to clear deallocations for
 * @param numToClear a list of number of pending deallocations to clear for each strategy
 *
 * @dev can be called permissionlessly by anyone
 */
function clearDeallocationQueue(
    address operator,
    IStrategy[] calldata strategies,
    uint16[] calldata numToClear
)
    external
    onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS)
```

This function is used to complete any eligible pending deallocations for an operator. The function takes an operator, a list of strategies, and a corresponding number of pending deallocations to complete. 

Clearing pending deallocations plays an important role in [`modifyAllocations`](#modifyallocations), as completable deallocations represent magnitude that can be freed for re-allocation to a different operator set. This method exists as a convenience for operators that want to complete pending deallocations as a standalone operation. However, `modifyAllocations` will _automatically_ clear any eligible deallocations when processing an allocation modification for a given strategy.

For each strategy, the method iterates over `deallocationQueue[operator][strategy]`:

```solidity
/// @dev For a strategy, keeps an ordered queue of operator sets that have pending deallocations
/// These must be completed in order to free up magnitude for future allocation
mapping(address operator => mapping(IStrategy strategy => DoubleEndedQueue.Bytes32Deque)) internal deallocationQueue;
```

This queue contains a per-strategy ordered list of operator sets that, due to prior calls by the `operator` to `modifyAllocations`, have a pending decrease in slashable magnitude. For each operator set in the queue, the corresponding allocation for that operator set is evaluated. If its `effectBlock` has been reached, the deallocation is completed, freeing up the deallocated magnitude by subtracting it from `encumberedMagnitude[operator][strategy]`. The corresponding entry is then popped from the front of the queue.

This method stops iterating when: the queue is empty, a deallocation is reached that cannot be completed yet, or when it has cleared `numToClear` entries from the queue.

*Effects*:
* For each `strategy` and _completable_ deallocation in `deallocationQueue[operator][strategy]`:
    * Pops the corresponding operator set from the `deallocationQueue`
    * Reduces `allocation.currentMagnitude` by the deallocated amount
    * Sets `allocation.pendingDiff` and `allocation.effectBlock` to 0
    * Adds the deallocated amount to the strategy's `encumberedMagnitude`
    * Emits `EncumberedMagnitudeUpdated`
    * Additionally, if the deallocation leaves `allocation.currentMagnitude` equal to zero:
        * Removes `strategy` from the `allocatedStrategies[operator][operatorSetKey]` list
        * If this list now has a length of 0, remove `operatorSetKey` from `allocatedSets[operator]`

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_MODIFY_ALLOCATIONS`
* Strategy list MUST be equal length of `numToClear` list

#### `slashOperator`

```solidity
/**
 * @notice Struct containing parameters to slashing
 * @param operator the address to slash
 * @param operatorSetId the ID of the operatorSet the operator is being slashed on behalf of
 * @param strategies the set of strategies to slash
 * @param wadsToSlash the parts in 1e18 to slash, this will be proportional to the operator's
 * slashable stake allocation for the operatorSet
 * @param description the description of the slashing provided by the AVS for legibility
 */
struct SlashingParams {
    address operator;
    uint32 operatorSetId;
    IStrategy[] strategies;
    uint256[] wadsToSlash;
    string description;
}

/**
 * @notice Called by an AVS to slash an operator in a given operator set. The operator must be registered
 * and have slashable stake allocated to the operator set.
 *
 * @param avs The AVS address initiating the slash.
 * @param params The slashing parameters, containing:
 *  - operator: The operator to slash.
 *  - operatorSetId: The ID of the operator set the operator is being slashed from.
 *  - strategies: Array of strategies to slash allocations from (must be in ascending order).
 *  - wadsToSlash: Array of proportions to slash from each strategy (must be between 0 and 1e18).
 *  - description: Description of why the operator was slashed.
 *
 * @dev For each strategy:
 *      1. Reduces the operator's current allocation magnitude by wadToSlash proportion.
 *      2. Reduces the strategy's max and encumbered magnitudes proportionally.
 *      3. If there is a pending deallocation, reduces it proportionally.
 *      4. Updates the operator's shares in the DelegationManager.
 *
 * @dev Small slashing amounts may not result in actual token burns due to
 *      rounding, which will result in small amounts of tokens locked in the contract
 *      rather than fully burning through the burn mechanism.
 */
function slashOperator(
    address avs,
    SlashingParams calldata params
)
    external
    onlyWhenNotPaused(PAUSED_OPERATOR_SLASHING)
    checkCanCall(avs)
```

_Note: this method can be called directly by an AVS, or by a caller authorized by the AVS. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

AVSs use slashing as a punitive disincentive for misbehavior. For details and examples of how slashing works, see [ELIP-002#Slashing of Unique Stake](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md#slashing-of-unique-stake). Note that whatever slashing criteria an AVS decides on, the only criteria enforced by the `AllocationManager` are those detailed above (see [Evaluating Whether an Allocation is "Slashable"](#evaluating-whether-an-allocation-is-slashable)).

In order to slash an eligible operator, the AVS specifies which operator set the operator belongs to, the `strategies` the operator should be slashed for, and for each strategy, the _proportion of the operator's allocated magnitude_ that should be slashed (given by `wadsToSlash`). An optional `description` string allows the AVS to add context to the slash.

Once triggered in the `AllocationManager`, slashing is instant and irreversible. For each slashed strategy, the operator's `maxMagnitude` and `encumberedMagnitude` are decreased, and the allocation made to the given operator set has its `currentMagnitude` reduced. See [TODO - Accounting Doc]() for details on how slashed amounts are calculated.

There are two edge cases to note for this method:
1. In the process of slashing an `operator` for a given `strategy`, if the `Allocation` being slashed has a `currentMagnitude` of 0, the call will NOT revert. Instead, the `strategy` is skipped and slashing continues with the next `strategy` listed. This is to prevent an edge case where slashing occurs on or around a deallocation's `effectBlock` -- if the call reverted, the entire slash would fail. Skipping allows any valid slashes to be processed without requiring resubmission.
2. If the `operator` has a pending, non-completable deallocation, the deallocation's `pendingDiff` is reduced proportional to the slash. This ensures that when the deallocation is completed, less `encumberedMagnitude` is freed.

Once slashing is processed for a strategy, [slashed stake is burned via the `DelegationManager`](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md#burning-of-slashed-funds).

*Effects*:
* Given an `operator` and `operatorSet`, then for each `params.strategies` element and its corresponding `allocation`:
    * Calculates magnitude to slash by multiplying current magnitude by the provided `wadsToSlash`
    * Reduce `allocation.currentMagnitude` by the slashed magnitude
        * Emit an `AllocationUpdated` event
    * Reduce the operator's `encumberedMagnitude` for this strategy by the slashed magnitude
        * Emit an `EncumberedMagnitudeUpdated` event
    * Push an entry to the operator's `maxMagnitudeHistory`, reducing their `maxMagnitude` by the slashed magnitude
        * Emit a `MaxMagnitudeUpdated` event
    * If the `allocation` has a pending, non-completable deallocation, additionally reduce `allocation.pendingDiff` by the same proportion and emit an `AllocationUpdated` event
    * If the `allocation` now has a `currentMagnitude` of 0:
        * Removes `strategy` from the `allocatedStrategies[operator][operatorSetKey]` list
        * If this list now has a length of 0, remove `operatorSetKey` from `allocatedSets[operator]`
    * Calls [`DelegationManager.slashOperatorShares`](./DelegationManager.md#slashoperatorshares)
* Emit an `OperatorSlashed` event

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_OPERATOR_SLASHING`
* Caller MUST be authorized, either as the AVS itself or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* Operator set MUST be registered for the AVS
* Operator MUST BE slashable, i.e.:
    * Operator is registered for the operator set, *OR*
    * The operator's `DEALLOCATION_DELAY` has not yet completed
* `params.strategies` MUST be in ascending order (to ensure no duplicates)
* `params.strategies.length` MUST be equal to `params.wadsToSlash.length`
* For each `params.strategies` element:
    * `wadsToSlash` MUST be within the bounds `(0, 1e18]`
    * Operator set MUST contain the strategy

---

## Config

**Methods:**
* [`setAllocationDelay`](#setallocationdelay)
* [`setAVSRegistrar`](#setavsregistrar)

#### `setAllocationDelay`

```solidity
/**
 * @notice Called by the delegation manager OR an operator to set an operator's allocation delay.
 * This is set when the operator first registers, and is the number of blocks between an operator
 * allocating magnitude to an operator set, and the magnitude becoming slashable.
 * @param operator The operator to set the delay on behalf of.
 * @param delay the allocation delay in blocks
 */
function setAllocationDelay(
    address operator,
    uint32 delay
)
    external
```

_Note: IF NOT CALLED BY THE `DelegationManager`, this method can be called directly by an operator, or by a caller authorized by the operator. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

This function sets an operator's allocation delay, in blocks. This delay can be updated by the operator once set. Both the initial setting of this value and any further updates _take `ALLOCATION_CONFIGURATION_DELAY` blocks_ to take effect. Because having a delay is a requirement to allocating slashable stake, this effectively means that once the slashing release goes live, no one will be able to allocate slashable stake for at least `ALLOCATION_CONFIGURATION_DELAY` blocks.

The `DelegationManager` calls this upon operator registration for all new operators created after the slashing release. For operators that existed in the `DelegationManager` _prior_ to the slashing release, **they will need to call this method to configure an allocation delay prior to allocating slashable stake to any AVS**.

The allocation delay's primary purpose is to give stakers delegated to an operator the chance to withdraw their stake before the operator can change the risk profile to something they're not comfortable with. However, operators can choose to configure this delay however they want - including setting it to 0.

*Effects*:
* Sets the operator's `pendingDelay` to the proposed `delay`, and save the `effectBlock` at which the `pendingDelay` can be activated
    * `effectBlock = uint32(block.number) + ALLOCATION_CONFIGURATION_DELAY + 1`
* If the operator has a `pendingDelay`, and if the `effectBlock` has passed, sets the operator's `delay` to the `pendingDelay` value
    * This also sets the `isSet` boolean to `true` to indicate that the operator's `delay`, even if 0, was set intentionally
* Emits an `AllocationDelaySet` event

*Requirements*:
* Caller MUST BE either the DelegationManager, or a registered operator
    * An admin and/or appointee for the operator can also call this function (see [`PermissionController.md`](../permissions/PermissionController.md))

#### `setAVSRegistrar`

```solidity
/**
 * @notice Called by an AVS to configure the address that is called when an operator registers
 * or is deregistered from the AVS's operator sets. If not set (or set to 0), defaults
 * to the AVS's address.
 * @param registrar the new registrar address
 */
function setAVSRegistrar(
    address avs,
    IAVSRegistrar registrar
)
    external
    checkCanCall(avs)
```

_Note: this method can be called directly by an AVS, or by a caller authorized by the AVS. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

Sets the `registrar` for a given `avs`. Note that if the registrar is set to 0, `getAVSRegistrar` will return the AVS's address.

The avs registrar is called when operators register to or deregister from an operator set. From [`IAVSRegistrar.sol`](../../src/contracts/interfaces/IAVSRegistrar.sol), the avs registrar should use the following interface:

```solidity
interface IAVSRegistrar {
    /**
     * @notice Called by the AllocationManager when an operator wants to register
     * for one or more operator sets. This method should revert if registration
     * is unsuccessful.
     * @param operator the registering operator
     * @param operatorSetIds the list of operator set ids being registered for
     * @param data arbitrary data the operator can provide as part of registration
     */
    function registerOperator(address operator, uint32[] calldata operatorSetIds, bytes calldata data) external;

    /**
     * @notice Called by the AllocationManager when an operator is deregistered from
     * one or more operator sets. If this method reverts, it is ignored.
     * @param operator the deregistering operator
     * @param operatorSetIds the list of operator set ids being deregistered from
     */
    function deregisterOperator(address operator, uint32[] calldata operatorSetIds) external;
}
```

Note that when an operator registers, registration will FAIL if the call to `IAVSRegistrar` reverts. However, when an operator deregisters, a revert in `deregisterOperator` is ignored.

*Effects*:
* Sets `_avsRegistrar[avs]` to `registrar`
* Emits an `AVSRegistrarSet` event

*Requirements*:
* Caller MUST be authorized, either as the AVS itself or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
