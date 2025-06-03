# v1.6.0 Moocow and ELIP5

## Release Manager

@wadealexc @bowenli86

## Highlights

🚀 **New Features**
- New APIs supporting Pectra's validator consolidation and withdrawal features: `EigenPod.requestConsolidation(...)` and `EigenPod.requestWithdrawal(...)`
- New getters to support Pectra APIs: `EigenPod.getConsolidationRequestFee()` and `EigenPod.getWithdrawalRequestFee()`
- Added 4 new events to `EigenPod.sol` to track consolidation and withdrawal requests
- Added 2 new events to `Eigen.sol` to track token wraps/unwraps with `BackingEigen`

📌 **Deprecations**
- Removed `EigenPod.GENESIS_TIME()` getter. This method, though public, has been unused for over a year.

🔧 **Improvements**
- When finalizing an `EigenPod` checkpoint (`proofsRemaining == 0`), the contract will store the finalized checkpoint in storage. This can be queried via `EigenPod.currentCheckpoint()`. Starting a new checkpoint will overwrite this previously-finalized checkpoint.
- Added semver to `Eigen`
- Signatures of a few `EigenPod` events are changed to match the rest events and take validator pubkey hash instead of validator index, which standardized `EigenPod` events signature

🐛 Bug Fixes
- For Hoodi, updates fixes ethPOS deposit contract to point to `0x00000000219ab540356cBB839Cbe05303d7705Fa`

## Changelog

- feat: merge Moocow and ELIP5 into main [PR #1425](https://github.com/layr-labs/eigenlayer-contracts/pull/1425)
- docs: proper markdown [PR #1435](https://github.com/layr-labs/eigenlayer-contracts/pull/1435)
- docs: update readme
- chore: update testnet addresses for redistribution [PR #1428](https://github.com/layr-labs/eigenlayer-contracts/pull/1428)
- chore: remove User_M2.t.sol
- feat: update EIGEN binding
- chore: resolve conflicts in upgrade.json
- chore: update harness class formatting
- chore: complete v1.6.0 changelog
- chore: changelog and bindings
- test: add more script tests for Eigen and standardize semver
- feat: add semver to eigen [PR #1371](https://github.com/layr-labs/eigenlayer-contracts/pull/1371)
- feat: add TokenWrapped and TokenUnwrapped events in Eigen for observability [PR #1356](https://github.com/layr-labs/eigenlayer-contracts/pull/1356)
- feat: change eigenpod events to use pubkeyHash over index
- feat: release scripts for moocow and elip5
- feat: currentCheckpoint now returns finalized checkpoint
- feat: implement consolidation and withdrawal requests
