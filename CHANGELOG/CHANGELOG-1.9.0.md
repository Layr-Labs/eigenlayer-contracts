# v1.9.0 Slashing UX Improvements

## Release Manager

@ypatil12 @eigenmikem @0xclandestine

# Overview

The Slashing UX improvement release is a tech debt-focused release that improves key parts of the Eigenlayer Core UX. This release will upgrade every core contract. 

The below release notes cover Core Contracts.

## Highlights

🚀 New Features

- The `AllocationManager` has been split into two contracts to address size limitations of the contract. The main contract handles state-mutating operations, while `AllocationManagerView` handles all read-only view functions. **This is not a breaking change for introspection as previous introspection calls fallback to `delegateCall` into the `AllocationManagerView` contract.**. For more information, see the [contract architecture](../docs/core/AllocationManager.md#contract-architecture). 
- The `ProtocolRegistry` is a new contract that stores all proxy contract addresses, global semver, and has the ability to pause the entire protocol. This contract will be deployed on all chains.
- Two new `createOperatorSets` functions (for redistributing and non redistributing operatorSets) have been added that take in a slasher address. This address is the *only* address that can slash an operatorSet. Changing the address is behind a `ALLOCATION_CONFIGURATION_DELAY` (17.5 days on mainnet). 

⛔ Breaking Changes

- The slasher permissions are set and stored in the `AllocationManager` instead of the `PermissionController`. Only one address can slash an operatorSet; the address is initially set upon creation of the operatorSet. OperatorSets created prior to this release will have their slasher migrated based on the following rules:
    - If there is no slasher set or the slasher in the `PermissionController` is the 0 address, the AVS address will be set as the slasher
    - If there are multiple slashers set in the `PermissionController`, the first address will be set as the slasher
- Semver (`SemverMixin.sol`) has been removed from all contracts, except from those that inherit the `SignatureUtilsMixin`. The version of the core protocol can be introspected via the `ProtocolRegistry`. 

📌 Deprecations

The old `createOperatorSets` functions in the leftmost column will be deprecated in Q2 2026 in favor of the newly specified functions. The old functions do not pass in a slasher. The new functions do pass in a slasher. If the old function is used, the slasher of the operatorSet is set to the avs address. 

| Function | MigrateTo | Notes |
| -------- | -------- | -------- |
| `createOperatorSets(avs, CreateSetParams[])` | `createOperatorSets(address avs, CreateSetParamsV2[])` | New function takes in a slasher address |
| `createRedistributingOperatorSets(avs, CreateSetParams[], redistributionRecipients[])` | `createRedistributingOperatorSets(avs, CreateSetParamsV2[], redistributionRecipients[])` | New function takes in a slasher address |

🔧 Improvements

- Added a non-revert `_canCall` in the `PermissionControllerMixin` for space savings. This function is used in the `AllocationManager` and `DelegationManager`. 
- The allocation delay for a newly created operator is active immediately. This allows operators to make allocations instantly after registering in the core. 
- The internal `SlashingLib.scaleForBurning` function has been deprecated in favor of `SlashingLib.calcSlashedAmount`, standardizing the calculation of slashed shares across the withdrawal queue and storage. See [PR #1502](https://github.com/Layr-Labs/eigenlayer-contracts/pull/1502) for more information. 


# Changelog

- feat: re-enable forge fmt + foundry v1.5.0 [PR #1669](https://github.com/layr-labs/eigenlayer-contracts/pull/1669)
- feat: substitute calcSlashedAmount for scaleForBurning [PR #1502](https://github.com/layr-labs/eigenlayer-contracts/pull/1502)
- fix: `v1.9.0` upgrade script [PR #1666](https://github.com/layr-labs/eigenlayer-contracts/pull/1666)
- feat: `v1.9.0` upgrade scripts + reusable upgrade helpers [PR #1665](https://github.com/layr-labs/eigenlayer-contracts/pull/1665)
- chore: update interface natspec for DM [PR #1664](https://github.com/layr-labs/eigenlayer-contracts/pull/1664)
- feat: slashing commitments [PR #1645](https://github.com/layr-labs/eigenlayer-contracts/pull/1645)
- feat: remove semver + minor optimizations [PR #1654](https://github.com/layr-labs/eigenlayer-contracts/pull/1654)
- feat: split `AllocationManager` [PR #1643](https://github.com/layr-labs/eigenlayer-contracts/pull/1643)
- feat: add protocol registry [PR #1655](https://github.com/layr-labs/eigenlayer-contracts/pull/1655)
- feat: instant alloc delay from dm [PR #1646](https://github.com/layr-labs/eigenlayer-contracts/pull/1646)
- chore: remove holesky [PR #1662](https://github.com/layr-labs/eigenlayer-contracts/pull/1662)
- chore: hardcode foundry ci to v1.3.5 [PR #1658](https://github.com/layr-labs/eigenlayer-contracts/pull/1658)
- chore: hardcode foundry to v1.3.5 in ci [PR #1657](https://github.com/layr-labs/eigenlayer-contracts/pull/1657)
- feat(audit): publish Hourglass + Multichain + RMS audit reports [PR #1644](https://github.com/layr-labs/eigenlayer-contracts/pull/1644)
- docs: add transport frequency for multichain [PR #1642](https://github.com/layr-labs/eigenlayer-contracts/pull/1642)
- chore: update readMe for multichain/hourglass [PR #1637](https://github.com/layr-labs/eigenlayer-contracts/pull/1637)