# Solidity API

## Pausable

Contracts that inherit from this contract may define their own `pause` and `unpause` (and/or related) functions.
These functions should be permissioned as "onlyPauser" which defers to a `PauserRegistry` for determining access control.

_Pausability is implemented using a uint256, which allows up to 256 different single bit-flags; each bit can potentially pause different functionality.
Inspiration for this was taken from the NearBridge design here https://etherscan.io/address/0x3FEFc5A4B1c02f21cBc8D3613643ba0635b9a873#code.
For the `pause` and `unpause` functions we've implemented, if you pause, you can only flip (any number of) switches to on/1 (aka "paused"), and if you unpause,
you can only flip (any number of) switches to off/0 (aka "paused").
If you want a pauseXYZ function that just flips a single bit / "pausing flag", it will:
1) 'bit-wise and' (aka `&`) a flag with the current paused state (as a uint256)
2) update the paused state to this new value
We note as well that we have chosen to identify flags by their *bit index* as opposed to their numerical value, so, e.g. defining `DEPOSITS_PAUSED = 3`
indicates specifically that if the *third bit* of `_paused` is flipped -- i.e. it is a '1' -- then deposits should be paused_

### pauserRegistry

```solidity
contract IPauserRegistry pauserRegistry
```

Address of the `PauserRegistry` contract that this contract defers to for determining access control (for pausing).

### _paused

```solidity
uint256 _paused
```

_whether or not the contract is currently paused_

### UNPAUSE_ALL

```solidity
uint256 UNPAUSE_ALL
```

### PAUSE_ALL

```solidity
uint256 PAUSE_ALL
```

### Paused

```solidity
event Paused(address account, uint256 newPausedStatus)
```

Emitted when the pause is triggered by `account`, and changed to `newPausedStatus`.

### Unpaused

```solidity
event Unpaused(address account, uint256 newPausedStatus)
```

Emitted when the pause is lifted by `account`, and changed to `newPausedStatus`.

### onlyPauser

```solidity
modifier onlyPauser()
```

@notice

### onlyUnpauser

```solidity
modifier onlyUnpauser()
```

### whenNotPaused

```solidity
modifier whenNotPaused()
```

Throws if the contract is paused, i.e. if any of the bits in `_paused` is flipped to 1.

### onlyWhenNotPaused

```solidity
modifier onlyWhenNotPaused(uint8 index)
```

Throws if the `indexed`th bit of `_paused` is 1, i.e. if the `index`th pause switch is flipped.

### _initializePauser

```solidity
function _initializePauser(contract IPauserRegistry _pauserRegistry, uint256 initPausedStatus) internal
```

One-time function for setting the `pauserRegistry` and initializing the value of `_paused`.

### pause

```solidity
function pause(uint256 newPausedStatus) external
```

This function is used to pause an EigenLayer contract's functionality.
It is permissioned to the `pauser` address, which is expected to be a low threshold multisig.

_This function can only pause functionality, and thus cannot 'unflip' any bit in `_paused` from 1 to 0._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| newPausedStatus | uint256 | represents the new value for `_paused` to take, which means it may flip several bits at once. |

### pauseAll

```solidity
function pauseAll() external
```

Alias for `pause(type(uint256).max)`.

### unpause

```solidity
function unpause(uint256 newPausedStatus) external
```

This function is used to unpause an EigenLayer contract's functionality.
It is permissioned to the `unpauser` address, which is expected to be a high threshold multisig or governance contract.

_This function can only unpause functionality, and thus cannot 'flip' any bit in `_paused` from 0 to 1._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| newPausedStatus | uint256 | represents the new value for `_paused` to take, which means it may flip several bits at once. |

### paused

```solidity
function paused() public view virtual returns (uint256)
```

Returns the current paused status as a uint256.

### paused

```solidity
function paused(uint8 index) public view virtual returns (bool)
```

Returns 'true' if the `indexed`th bit of `_paused` is 1, and 'false' otherwise

### __gap

```solidity
uint256[48] __gap
```

_This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps_

