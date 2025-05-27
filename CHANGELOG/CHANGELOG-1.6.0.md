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

## Changelog

* feat: MOOCOW and ELIP5 ([#1375](https://github.com/Layr-Labs/eigenlayer-contracts/pull/1375))
