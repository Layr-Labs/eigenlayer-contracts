# v1.5.0 Redistribution

## Release Manager

@0xClandestine @ypatil12

## Highlights

🚀 New features

- Redistribution is a feature that gives Service Builders a means to not just burn, but repurpose slashed funds.
- We introduce a new operatorSet creation mechanism: [`AllocationManager.createRedistributingOperatorSets`](../docs/core/AllocationManager.md#createredistributingoperatorsets), which allows slashed funds to be redistributed to a `RedistributionRecipient`. *Note: The redistribution recipient can be set only once and is immutable*. 
- *All slashed funds will now be routed to individual `SlashEscrow` contracts.* The release of funds from escrow is gated by the `SlashEscrowFactory`. The `SlashEscrowFactory` deploys individual `SlashEscrow` contracts per slash, enforces a global delay for all escrowed funds, and handles pausing/unpausing of escrowed funds. 
- The original `createOperatorSets` function still exists. This function creates operatorSets whose slashed funds will eventually be burned. There is no mechanism to convert an operatorSet to be redistributing. 
- See [ELIP-006](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-006.md) for a full description. 

⛔ Breaking changes
- Funds marked for burning now go through a 4-day escrow period via `SlashEscrow` contracts. These funds are burned by calling [`SlashEscrowFactory.releaseSlashEscrow`](../docs/core/SlashEscrowFactory.md#releaseslashescrow).

📌 Future Deprecations 
- The pre-redistribution burn pathway [`StrategyManager.decreaseBurnableShares`](../docs/core/StrategyManager.md#burnshares) will be deprecated in an upgrade *after* the redistribution release. This function can still be used to burn shares that have been slashed at any point prior to the redistribution upgrade. 

🛠️ Security Updates
- The slashing of burned funds is no longer instantaneous. All slashed funds (burned or redistributed) now go through a 4-day escrow delay. The eventual burning or redistribution of slashed funds can be paused by the `PauserMultisig`. 
- The upgradability of the `SlashEscrowFactory` is controlled by the `CommunityMultisig`. The contract will have a separate `ProxyAdmin` from the rest of the EigenLayer core protocol. Each individual `SlashEscrow` contract is an immutable clone. 

🔧 Improvements
- The [`AllocationManager.slashOperator`](../docs/core/AllocationManager.md#slashoperator) function now returns a `slashId` and array of `shares` to be burned/redistributed. **The function selector remains the same.**
- OperatorSets now have a `slashCount` field, which returns the number of slashes completed by the operatorSet. This value only reflects the number of slashes after the redistribution upgrade.
- `StrategyBase` returns an `amountOut` upon withdrawal to comply with standard ERC-4626 vaults. 
- The `AllocationManager` and `DelegationManager` no longer use ownable. Thus, they now inherit the `Deprecated_OwnableUpgradeable` mixin in its place to reduce codesize. 

🐛 Bug Fixes
- `SemVerMixin` is updated to only return the first character of `majorVersion`. We currently return `1.` and will return `1` after this upgrade. 


## Changelog

- feat(draft): `AllocationManager` redistribution support [PR #1346](https://github.com/layr-labs/eigenlayer-contracts/pull/1346)
- feat: redistribution upgrade script [PR #1396](https://github.com/layr-labs/eigenlayer-contracts/pull/1396)
- chore: bindings [PR #1422](https://github.com/layr-labs/eigenlayer-contracts/pull/1422)
- test: redistribution upgrade [PR #1410](https://github.com/layr-labs/eigenlayer-contracts/pull/1410)
- test: redistribution integration [PR #1415](https://github.com/layr-labs/eigenlayer-contracts/pull/1415)
- docs: redistribution [PR #1409](https://github.com/layr-labs/eigenlayer-contracts/pull/1409)
- chore: address redistribution nits [PR #1420](https://github.com/layr-labs/eigenlayer-contracts/pull/1420)
- chore: style updates [PR #1416](https://github.com/layr-labs/eigenlayer-contracts/pull/1416)
- perf: avoid binary search [PR #1417](https://github.com/layr-labs/eigenlayer-contracts/pull/1417)
- feat: release escrow by strategy [PR #1412](https://github.com/layr-labs/eigenlayer-contracts/pull/1412)
- refactor: review changes [PR #1411](https://github.com/layr-labs/eigenlayer-contracts/pull/1411)
- refactor: `decreaseBurnOrRedistributableShares` [PR #1414](https://github.com/layr-labs/eigenlayer-contracts/pull/1414)
- feat: deploy escrow in `initiateSlashEscrow` [PR #1413](https://github.com/layr-labs/eigenlayer-contracts/pull/1413)
- chore: update naming [PR #1408](https://github.com/layr-labs/eigenlayer-contracts/pull/1408)
- feat: simplify escrow delay; add convenience functions [PR #1406](https://github.com/layr-labs/eigenlayer-contracts/pull/1406)
- fix: enumerable map overwrite [PR #1399](https://github.com/layr-labs/eigenlayer-contracts/pull/1399)
- chore: decrease dm diff further  [PR #1404](https://github.com/layr-labs/eigenlayer-contracts/pull/1404)
- test: full coverage `SlashEscrowFactory` + `SlashEscrow` [PR #1403](https://github.com/layr-labs/eigenlayer-contracts/pull/1403)
- chore: remove dm/alm code size optimizations [PR #1398](https://github.com/layr-labs/eigenlayer-contracts/pull/1398)
- chore: rename burnable -> burnOrRedistributable; fix storage gap; remove poc code [PR #1397](https://github.com/layr-labs/eigenlayer-contracts/pull/1397)
- chore: use internal getters; update `isOperatorRedistributable` [PR #1401](https://github.com/layr-labs/eigenlayer-contracts/pull/1401)
- fix: storage checker [PR #1394](https://github.com/layr-labs/eigenlayer-contracts/pull/1394)
- fix: review issues [PR #1391](https://github.com/layr-labs/eigenlayer-contracts/pull/1391)
- feat: escrow funds in unique clone contracts [PR #1387](https://github.com/layr-labs/eigenlayer-contracts/pull/1387)
- refactor: remove `v` prefix from `SemVerMixin` [PR #1385](https://github.com/layr-labs/eigenlayer-contracts/pull/1385)
- test(redistribution): add unit tests  [PR #1383](https://github.com/layr-labs/eigenlayer-contracts/pull/1383)
- feat: add `SlashingWithdrawalRouter` [PR #1358](https://github.com/layr-labs/eigenlayer-contracts/pull/1358)
- feat: simplify removeDepositShares in StrategyManager [PR #1373](https://github.com/layr-labs/eigenlayer-contracts/pull/1373)
- feat(draft): `AllocationManager` redistribution support [PR #1346](https://github.com/layr-labs/eigenlayer-contracts/pull/1346)
- ci: add explicit permissions to workflows to  mitigate security concerns [PR #1392](https://github.com/layr-labs/eigenlayer-contracts/pull/1392)
- ci: remove branch constraint for foundry coverage job
- docs: add release managers to changelogs
- docs: add templates for changelog and release notes [PR #1382](https://github.com/layr-labs/eigenlayer-contracts/pull/1382)
- docs: add doc for steps to write deploy scripts [PR #1380](https://github.com/layr-labs/eigenlayer-contracts/pull/1380)
- ci: add testnet envs sepolia and hoodi to validate-deployment-scripts [PR #1378](https://github.com/layr-labs/eigenlayer-contracts/pull/1378)
- docs: update MAINTENANCE to include practices of merging multiple release-dev branches
- docs: updating readme for dead links, readability, new language, and more [PR #1377](https://github.com/layr-labs/eigenlayer-contracts/pull/1377)
- docs: bump deployment matrix to top of README [PR #1376](https://github.com/layr-labs/eigenlayer-contracts/pull/1376)
- ci: add CI to auto validate deployment scripts [PR #1360](https://github.com/layr-labs/eigenlayer-contracts/pull/1360)
- chore: update readme for v1.4.1 [PR #1361](https://github.com/layr-labs/eigenlayer-contracts/pull/1361)
- ci: add cron to auto remove stale branches [PR #1348](https://github.com/layr-labs/eigenlayer-contracts/pull/1348)
- chore: Update README for Holesky v1.4.2 release [PR #1351](https://github.com/layr-labs/eigenlayer-contracts/pull/1351)
- docs: remove fork-pr instructions from CONTRIBUTING.md and MAINTENANCE.md
- ci: disable delete unauthorized branches
- docs: update addresses for mainnet [PR #1341](https://github.com/layr-labs/eigenlayer-contracts/pull/1341)
- docs: enrich MAINTENANCE.md re: release branches [PR #1340](https://github.com/layr-labs/eigenlayer-contracts/pull/1340)
- ci: enable auto delete branch upon eigengit launch [PR #1339](https://github.com/layr-labs/eigenlayer-contracts/pull/1339)