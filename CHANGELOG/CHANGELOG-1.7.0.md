# v1.7.0 Multi Chain

The multichain release enables AVSs to launch their services and make verified Operator outputs available on any EVM chain, meeting their customers where they are. AVSs can specify custom operator weights to be transported to any destination chain. The release has 3 components:

1. Core Contracts
2. AVS Contracts
3. Offchain Infrastructure

The below release notes cover Core Contracts. For more information on the end to end protocol, see our [docs](../docs/multichain/README.md) and [ELIP-008](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-008.md).

## Release Manager

@ypatil12 @eigenmikem

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

## Changelog

- feat: multichain deploy scripts [PR #1487](https://github.com/layr-labs/eigenlayer-contracts/pull/1487)
- feat: operator table updater pauser [PR #1501](https://github.com/layr-labs/eigenlayer-contracts/pull/1501)
- feat: add `publishMetadataURI` [PR #1492](https://github.com/layr-labs/eigenlayer-contracts/pull/1492)
- refactor: `globalRootConfirmerSet` -> `generator` [PR #1500](https://github.com/layr-labs/eigenlayer-contracts/pull/1500)
- fix: circular dependency for initial global root update [PR #1499](https://github.com/layr-labs/eigenlayer-contracts/pull/1499)
- chore: remove stale bindings [PR #1498](https://github.com/layr-labs/eigenlayer-contracts/pull/1498)
- fix: zero length [PR #1490](https://github.com/layr-labs/eigenlayer-contracts/pull/1490)
- docs: multichain docs [PR #1488](https://github.com/layr-labs/eigenlayer-contracts/pull/1488)
- refactor: remove table calculators [PR #1493](https://github.com/layr-labs/eigenlayer-contracts/pull/1493)
- refactor: `KeyRegistry` unit testing [PR #1482](https://github.com/layr-labs/eigenlayer-contracts/pull/1482)
- feat: disable root [PR #1481](https://github.com/layr-labs/eigenlayer-contracts/pull/1481)
- refactor: `ECDSATableCalculator` testing [PR #1479](https://github.com/layr-labs/eigenlayer-contracts/pull/1479)
- refactor: operators can deregister keys if not slashable [PR #1480](https://github.com/layr-labs/eigenlayer-contracts/pull/1480)
- refactor: `ECDSACertificateVerifier` testing [PR #1478](https://github.com/layr-labs/eigenlayer-contracts/pull/1478)
- refactor: `Bn254CertificateVerifierUnitTests` [PR #1476](https://github.com/layr-labs/eigenlayer-contracts/pull/1476)
- feat: ecdsa table calculator [PR #1473](https://github.com/layr-labs/eigenlayer-contracts/pull/1473)
- feat: ecdsacv views [PR #1475](https://github.com/layr-labs/eigenlayer-contracts/pull/1475)
- refactor: `BN254OperatorTableCalculator` [PR #1463](https://github.com/layr-labs/eigenlayer-contracts/pull/1463)
- feat: add `referenceBlockNumber` [PR #1472](https://github.com/layr-labs/eigenlayer-contracts/pull/1472)
- feat: release manager [PR #1469](https://github.com/layr-labs/eigenlayer-contracts/pull/1469)
- feat: ecdsa cert verifier [PR #1470](https://github.com/layr-labs/eigenlayer-contracts/pull/1470)
- feat: add view function for `getGlobalConfirmerSetReferenceTimestamp` [PR #1471](https://github.com/layr-labs/eigenlayer-contracts/pull/1471)
- chore: add helper view function [PR #1465](https://github.com/layr-labs/eigenlayer-contracts/pull/1465)
- chore: add sig digest functions to interface [PR #1464](https://github.com/layr-labs/eigenlayer-contracts/pull/1464)
- refactor: `CrossChainRegistry` [PR #1457](https://github.com/layr-labs/eigenlayer-contracts/pull/1457)
- fix: global table update message hash [PR #1460](https://github.com/layr-labs/eigenlayer-contracts/pull/1460)
- refactor: sig verification into library [PR #1455](https://github.com/layr-labs/eigenlayer-contracts/pull/1455)
- fix: `OperatorTableUpdater` encoding [PR #1456](https://github.com/layr-labs/eigenlayer-contracts/pull/1456)
- chore: add latest `referenceTimestamp` to OTC interface [PR #1454](https://github.com/layr-labs/eigenlayer-contracts/pull/1454)
- chore: bindings [PR #1452](https://github.com/layr-labs/eigenlayer-contracts/pull/1452)
- feat: add operator table updater to CCR [PR #1451](https://github.com/layr-labs/eigenlayer-contracts/pull/1451)
- chore: multichain deploy scripts [PR #1449](https://github.com/layr-labs/eigenlayer-contracts/pull/1449)
- feat: cross chain registry [PR #1439](https://github.com/layr-labs/eigenlayer-contracts/pull/1439)
- chore: update BN254CertificateVerifier [PR #1447](https://github.com/layr-labs/eigenlayer-contracts/pull/1447)
- feat: bn254 operator table contracts [PR #1429](https://github.com/layr-labs/eigenlayer-contracts/pull/1429)
- feat: KeyRegistrar [PR #1421](https://github.com/layr-labs/eigenlayer-contracts/pull/1421)
- feat: operator table updater [PR #1436](https://github.com/layr-labs/eigenlayer-contracts/pull/1436)
- feat: bn254 certificate verifier [PR #1431](https://github.com/layr-labs/eigenlayer-contracts/pull/1431)
- chore: bindings + interface update [PR #1438](https://github.com/layr-labs/eigenlayer-contracts/pull/1438)
- chore: update multichain interfaces [PR #1433](https://github.com/layr-labs/eigenlayer-contracts/pull/1433)
- feat: multi chain interfaces [PR #1423](https://github.com/layr-labs/eigenlayer-contracts/pull/1423)
- docs: update version matrix [PR #1491](https://github.com/layr-labs/eigenlayer-contracts/pull/1491)
- chore: add multisend parser to scripts directory [PR #1486](https://github.com/layr-labs/eigenlayer-contracts/pull/1486)
- chore: update eigenpod and eigen impls addresses in holesky and hoodi [PR #1448](https://github.com/layr-labs/eigenlayer-contracts/pull/1448)