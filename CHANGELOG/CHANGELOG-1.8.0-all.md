# v1.8.0-testnet-final

The below release notes cover the updated version release candidate for multichain and hourglass

# Release Manager

@ypatil12 @eigenmikem @rajath0x

# Multichain

## Highlights

⛔ Breaking Changes
- The leaves of merkle trees used by the `OperatorTableUpdater` and `BN254CertificateVerifier` are now salted. The [`LeafCalculatorMixin`]() is used by:
   - `OperatorTableUpdater`: To salt the `operatorTableLeaf` via `calculateOperatorTableLeaf`. This change is also reflected in the [offchain transporter](https://github.com/Layr-Labs/multichain-go/pull/19)
   - `BN254CertificateVerifier`: To salt the `operatorInfoLeaf` via `calculateOperatorInfoLeaf`. BN254 OperatorSets MUST update their table calculators to use the new `BN254TableCalculatorBase` in the [`middleware repo`](https://github.com/Layr-Labs/eigenlayer-middleware/blob/dev/src/middlewareV2/tableCalculator/BN254TableCalculatorBase.sol)
- Nonsigners in the `BN254CertificateVerifier` are now sorted by operator index. See [PR #1615](https://github.com/layr-labs/eigenlayer-contracts/pull/1615). All offchain aggregators MUST sort nonsigners by operator index
- The `BN254CertificateVerifier` now requires signatures over the `referenceTimestamp` via `calculateCertificateDigest`. See [PR #1610](https://github.com/layr-labs/eigenlayer-contracts/pull/1610)

🛠️ Security Fixes
- The merkle library has been updated to address minor audit issues. No breaking changes. See [PR #1606](https://github.com/layr-labs/eigenlayer-contracts/pull/1606)
- Audit fixes for the `ReleaseManager`. See [PR #1608](https://github.com/layr-labs/eigenlayer-contracts/pull/1608)

🔧 Improvements
- Introspection for operatorSets with active generation reservations: `hasActiveGenerationReservation`. See [PR #1589](https://github.com/layr-labs/eigenlayer-contracts/pull/1589)
- Clearer error messages/natspec:
   - [PR #1602](https://github.com/layr-labs/eigenlayer-contracts/pull/1602)
   - [PR #1587](https://github.com/layr-labs/eigenlayer-contracts/pull/1587)
   - [PR #1585](https://github.com/layr-labs/eigenlayer-contracts/pull/1585)
   - [PR #1562](https://github.com/layr-labs/eigenlayer-contracts/pull/1562)
   - [PR #1567](https://github.com/layr-labs/eigenlayer-contracts/pull/1567)
- Require `KeyType` to be set when creating a generation reservation. See [PR #1561](https://github.com/layr-labs/eigenlayer-contracts/pull/1561)

🐛 Bug Fixes
- Add pagination for querying active generation reservations. See [PR #1569](https://github.com/layr-labs/eigenlayer-contracts/pull/1569)
- Fix race conditions on offchain table updates. See [PR #1575](https://github.com/layr-labs/eigenlayer-contracts/pull/1575)
- Remove restrictive check on ECDSA certificates required to be confirmed against the latest `referenceTimestamp`. See [PR #1582](https://github.com/layr-labs/eigenlayer-contracts/pull/1582)

## What's Changed
- fix: enforce ordering of nonsigners in bn254CV [PR #1615](https://github.com/layr-labs/eigenlayer-contracts/pull/1615)
- fix: releaseManager audit fixes [PR #1608](https://github.com/layr-labs/eigenlayer-contracts/pull/1608)
- fix: include timestamp with BN254CertificateVerifier certificate generation [PR #1610](https://github.com/layr-labs/eigenlayer-contracts/pull/1610)
- fix(audit): merkle library audit fixes [PR #1606](https://github.com/layr-labs/eigenlayer-contracts/pull/1606)
- chore: clearer error message [PR #1602](https://github.com/layr-labs/eigenlayer-contracts/pull/1602)
- chore: fmt and bindings
- fix: elm-08
- fix: elm-11(2) middleware ref typo
- chore: bindings
- fix: add `hasActiveGenerationReservation` [PR #1589](https://github.com/layr-labs/eigenlayer-contracts/pull/1589)
- fix(h-03): add pagination [PR #1569](https://github.com/layr-labs/eigenlayer-contracts/pull/1569)
- fix(docs): correct hash value in natspec [PR #1587](https://github.com/layr-labs/eigenlayer-contracts/pull/1587)
- fix(h-04): race condition [PR #1575](https://github.com/layr-labs/eigenlayer-contracts/pull/1575)
- docs: multichain natspec [PR #1584](https://github.com/layr-labs/eigenlayer-contracts/pull/1584)
- fix(audit): add salt to Merkle leaf hashing [PR #1580](https://github.com/layr-labs/eigenlayer-contracts/pull/1580)
- fix: check for key upon gen reservation [PR #1561](https://github.com/layr-labs/eigenlayer-contracts/pull/1561)
- fix: remove unused imports [PR #1585](https://github.com/layr-labs/eigenlayer-contracts/pull/1585)
- feat: update generator script [PR #1581](https://github.com/layr-labs/eigenlayer-contracts/pull/1581)
- fix(m-01): remove restrictive check [PR #1582](https://github.com/layr-labs/eigenlayer-contracts/pull/1582)
- fix(I-01): check input lengths [PR #1564](https://github.com/layr-labs/eigenlayer-contracts/pull/1564)
- chore: clean up interfaces [PR #1562](https://github.com/layr-labs/eigenlayer-contracts/pull/1562)
- fix: addressing pr comments [PR #1568](https://github.com/layr-labs/eigenlayer-contracts/pull/1568)
- fix: correct ecdsa message hash check [PR #1563](https://github.com/layr-labs/eigenlayer-contracts/pull/1563)
- fix: Release Manager internal review fixes [PR #1571](https://github.com/layr-labs/eigenlayer-contracts/pull/1571)
- docs: multichain [PR #1567](https://github.com/layr-labs/eigenlayer-contracts/pull/1567)