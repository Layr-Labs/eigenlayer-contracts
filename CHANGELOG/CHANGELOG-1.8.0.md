# v1.8.0 Hourglass

The Hourglass release consists of a framework that supports the creation of task-based AVSs. The task-based AVSs are enabled through a `TaskMailbox` core contract deployed to all chains that support a `CertificateVerifier`. Additionally AVSs deploy their `TaskAVSRegistrar`. The release has 3 components:

1. Core Contracts
2. AVS Contracts
3. Offchain Infrastructure

The below release notes cover Core Contracts. For more information on the end to end protocol, see our [docs](https://github.com/Layr-Labs/hourglass-monorepo/blob/master/README.md) and [core contract docs](../docs/avs/task/TaskMailbox.md).

## Release Manager

@0xrajath

## Highlights

This hourglass release only introduces new contracts. As a result, there are no breaking changes or deprecations.

🚀 New Features

Destination Chain Contracts
- `TaskMailbox`: A core infrastructure contract that enables task-based AVS execution models. It provides a standardized way for AVSs to create tasks, have operators execute them, and submit verified results on-chain. The contract acts as a mailbox system where task creators post tasks with fees, and operators compete to execute and submit results with proper consensus verification.

## Changelog

- fix: correct ecdsa message hash check [PR #1563](https://github.com/layr-labs/eigenlayer-contracts/pull/1563)
- fix: missing assume in fuzz test
- fix: `submitResult` certificate checks [PR #1557](https://github.com/layr-labs/eigenlayer-contracts/pull/1557)
- chore: forge fmt
- fix: certificate verifier interface changes
- feat: hourglass zeus script [PR #1546](https://github.com/layr-labs/eigenlayer-contracts/pull/1546)
- fix: mock certificate verifiers [PR #1545](https://github.com/layr-labs/eigenlayer-contracts/pull/1545)
- feat: hourglass (task-based AVS framework) [PR #1534](https://github.com/layr-labs/eigenlayer-contracts/pull/1534)
