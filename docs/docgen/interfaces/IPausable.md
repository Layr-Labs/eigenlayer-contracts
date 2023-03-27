# Solidity API

## IPausable

### pauserRegistry

```solidity
function pauserRegistry() external view returns (contract IPauserRegistry)
```

Address of the `PauserRegistry` contract that this contract defers to for determining access control (for pausing).

### pause

```solidity
function pause(uint256 newPausedStatus) external
```

This function is used to pause an EigenLayer/DataLayer contract's functionality.
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

This function is used to unpause an EigenLayer/DataLayercontract's functionality.
It is permissioned to the `unpauser` address, which is expected to be a high threshold multisig or goverance contract.

_This function can only unpause functionality, and thus cannot 'flip' any bit in `_paused` from 0 to 1._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| newPausedStatus | uint256 | represents the new value for `_paused` to take, which means it may flip several bits at once. |

### paused

```solidity
function paused() external view returns (uint256)
```

Returns the current paused status as a uint256.

### paused

```solidity
function paused(uint8 index) external view returns (bool)
```

Returns 'true' if the `indexed`th bit of `_paused` is 1, and 'false' otherwise

