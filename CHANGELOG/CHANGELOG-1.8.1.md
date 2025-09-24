# v1.8.1 MultiChain Hourglass Combined

## Release Manager

@ypatil12 @eigenmikem @0xrajath

# Multichain 

The multichain release enables AVSs to launch their services and make verified Operator outputs available on any EVM chain, meeting their customers where they are. AVSs can specify custom operator weights to be transported to any destination chain. The release has 3 components:

1. Core Contracts
2. AVS Contracts
3. Offchain Infrastructure

The below release notes cover Core Contracts. For more information on the end to end protocol, see our [docs](../docs/multichain/README.md) and [ELIP-008](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-008.md).

## Highlights

This multichain release only introduces new standards and contracts. As a result, there are no breaking changes or deprecations. 

🚀 New Features – Highlight major new functionality

Source-Chain Contracts
- `KeyRegistrar`: Manages cryptographic keys for operators across different operator sets. It supports both ECDSA and BN254 key types and ensures global uniqueness of keys across all operator sets
- `CrossChainRegistry`: Enables AVSs to register to have their operator stakes transported to supported destination chains
- `ReleaseManager`: Provides a standardized way for AVSs to publish software artifacts (binaries, docker images, etc.) that operators in their operator sets should upgrade to by specified deadlines

Destination Chain Contracts
- `CertificateVerifier`: Proves the offchain execution of a task, via a Certificate, by the operators of an operatorSet. Two types of key material are supported: ECDSA and BN254
- `OperatorTableUpdater`: Updates operator tables in the `CertificateVerifier` to have tasks validated against up-to-date operator stake weights 

🔧 Improvements – Enhancements to existing features.

- The multichain protocol has protocol-ized several AVS-deployed contracts, enabling an simpler AVS developer experience. These include:
    - `KeyRegistrar`: Manages BLS and ECDSA signing keys. AVSs no longer have to deploy a `BLSAPKRegistry`
    - `CertificateVerifier`: Handles signature verification for BLS and ECDSA keys. AVSs no longer have to deploy a `BLSSignatureChecker`
    - Offchain Multichain Transport: AVSs no longer have to maintain [avs-sync](https://github.com/Layr-Labs/avs-sync) to keep operator stakes fresh

# Hourglass

The Hourglass release consists of a framework that supports the creation of task-based AVSs. The task-based AVSs are enabled through a `TaskMailbox` core contract deployed to all chains that support a `CertificateVerifier`. Additionally AVSs deploy their `TaskAVSRegistrar`. The release has 3 components:

1. Core Contracts
2. AVS Contracts
3. Offchain Infrastructure

The below release notes cover Core Contracts. For more information on the end to end protocol, see our [hourglass docs](https://github.com/Layr-Labs/hourglass-monorepo/blob/master/README.md),  [core contract docs](../docs/avs/task/TaskMailbox.md), and [ELIP-010](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-010.md).

## Highlights

This hourglass release only introduces new contracts. As a result, there are no breaking changes or deprecations.

🚀 New Features

Destination Chain Contracts
- `TaskMailbox`: A core infrastructure contract that enables task-based AVS execution models. It provides a standardized way for AVSs to create tasks, have operators execute them, and submit verified results on-chain. The contract acts as a mailbox system where task creators post tasks with fees, and operators compete to execute and submit results with proper consensus verification.

# Changelog

## What's Changed
* feat: multichain by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1494
* chore: add final moocow audit by @wadealexc in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1505
* fix: strategy manager gap by @0xClandestine in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1508
* test(redistribution-changes): passing by @0xClandestine in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1511
* docs(redistribution-changes): cleanup by @0xClandestine in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1513
* docs: update CHANGELOG by @0xClandestine in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1514
* release(redistribution): post audit changes by @0xClandestine in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1461
* refactor: remove unnecessary signature validation and change param name by @eigenmikem in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1509
* fix: remove unused constants and add gap by @8sunyuan in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1519
* feat: key data reverse lookup  by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1520
* docs: add final audit reports by @antojoseph in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1523
* docs: clear up certificate verification  by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1515
* refactor: cleaner reverts for `ECDSACertificateVerifier` by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1521
* fix: internal audit fixes by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1524
* docs: update README for v1.5.0 & v1.6.0 by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1526
* refactor: table calc interface by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1525
* fix: multichain clarity updates by @nadir-akhtar in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1527
* test: multichain integration tests by @eigenmikem in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1528
* fix: multichain deploy scripts by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1510
* chore: update `v1.7.0` changelog by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1529
* fix: typo  by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1530
* feat: invalid staleness period prevention by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1536
* feat: cleaner generator updates by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1537
* chore: symmetric `BN254` and `ECDSA` checks by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1540
* feat: add calculateCertificateDigestBytes to ECDSA cert verifier (#1532) by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1542
* refactor: remove transport interface by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1512
* chore: clean up natspec round 1 by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1541
* fix: storage gap in `CrossChainRegistry` by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1544
* test: more integration test scenarios and checks file by @eigenmikem in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1548
* fix: prevent `globalTableRoot` certificate replay upon newly instantiated CertificateVerifiers/OperatorTables by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1547
* fix: Multichain (Pt 2) fixes by @nadir-akhtar in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1549
* chore: update `foundry.toml` to include `forge lint` config by @0xClandestine in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1539
* chore: `KeyRegistrar` clarifications by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1551
* chore: `CertificateVerifier` clarifications by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1552
* chore: `CrossChainRegistry` clarifications by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1553
* docs: update hoodi strat addresses by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1556
* fix: hardening workflows test by @anupsv in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1554
* docs: update for `v1.6.0` and `v1.5.0` by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1558
* chore: update bindings by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1565
* docs: multichain by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1567
* fix: Release Manager internal review fixes by @nadir-akhtar in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1571
* feat: hourglass by @0xrajath in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1576
* fix: multichain pt1 audit fixes by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1572
* feat: update generator script by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1581
* fix: remove unused imports by @0xClandestine in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1585
* fix: sp multichain audit fixes by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1598
* chore: clearer error message by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1602
* fix: multichain pt2 audit fixes by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1592
* fix(audit): merkle library audit fixes by @nadir-akhtar in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1606
* fix: hourglass part 1 and 2 audit fixes by @0xrajath in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1609
* fix: include timestamp with BN254CertificateVerifier certificate generation by @nadir-akhtar in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1610
* fix: enforce ordering of nonsigners in bn254CV by @eigenmikem in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1615
* feat: final testnet upgrade scripts by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1594
* feat: electra timing fix by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1620
* docs: update readme by @eigenmikem in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1623
* fix: task replay directed at same operator set by @0xrajath in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1629
* feat: combine v1.7.0 and v1.8.0 upgrade scripts by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1628
* chore: zeus script for task replay fix by @0xrajath in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1632
* chore: update bindings by @0xrajath in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1634
* fix: create generator script; initialization ordering by @ypatil12 in https://github.com/Layr-Labs/eigenlayer-contracts/pull/1630