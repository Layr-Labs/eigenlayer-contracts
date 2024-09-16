| Name                       | Type                                      | Slot | Offset | Bytes | Contract                                     |
|----------------------------|-------------------------------------------|------|--------|-------|----------------------------------------------|
| IS_TEST                    | bool                                      | 0    | 0      | 1     | src/test/utils/ProofParsing.sol:ProofParsing |
| _failed                    | bool                                      | 0    | 1      | 1     | src/test/utils/ProofParsing.sol:ProofParsing |
| stdChainsInitialized       | bool                                      | 0    | 2      | 1     | src/test/utils/ProofParsing.sol:ProofParsing |
| chains                     | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| defaultRpcUrls             | mapping(string => string)                 | 2    | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| idToAlias                  | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| fallbackToDefaultRpcUrls   | bool                                      | 4    | 0      | 1     | src/test/utils/ProofParsing.sol:ProofParsing |
| gasMeteringOff             | bool                                      | 4    | 1      | 1     | src/test/utils/ProofParsing.sol:ProofParsing |
| stdstore                   | struct StdStorage                         | 5    | 0      | 224   | src/test/utils/ProofParsing.sol:ProofParsing |
| _excludedContracts         | address[]                                 | 12   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| _excludedSenders           | address[]                                 | 13   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| _targetedContracts         | address[]                                 | 14   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| _targetedSenders           | address[]                                 | 15   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| _excludedArtifacts         | string[]                                  | 16   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| _targetedArtifacts         | string[]                                  | 17   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| _targetedArtifactSelectors | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| _targetedSelectors         | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| stdstore                   | struct StdStorage                         | 20   | 0      | 224   | src/test/utils/ProofParsing.sol:ProofParsing |
| proofConfigJson            | string                                    | 27   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| prefix                     | string                                    | 28   | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| blockHeaderProof           | bytes32[18]                               | 29   | 0      | 576   | src/test/utils/ProofParsing.sol:ProofParsing |
| slotProof                  | bytes32[3]                                | 47   | 0      | 96    | src/test/utils/ProofParsing.sol:ProofParsing |
| withdrawalProofDeneb       | bytes32[10]                               | 50   | 0      | 320   | src/test/utils/ProofParsing.sol:ProofParsing |
| withdrawalProofCapella     | bytes32[9]                                | 60   | 0      | 288   | src/test/utils/ProofParsing.sol:ProofParsing |
| validatorProof             | bytes32[46]                               | 69   | 0      | 1472  | src/test/utils/ProofParsing.sol:ProofParsing |
| historicalSummaryProof     | bytes32[44]                               | 115  | 0      | 1408  | src/test/utils/ProofParsing.sol:ProofParsing |
| executionPayloadProof      | bytes32[7]                                | 159  | 0      | 224   | src/test/utils/ProofParsing.sol:ProofParsing |
| timestampProofsCapella     | bytes32[5]                                | 166  | 0      | 160   | src/test/utils/ProofParsing.sol:ProofParsing |
| timestampProofsDeneb       | bytes32[4]                                | 171  | 0      | 128   | src/test/utils/ProofParsing.sol:ProofParsing |
| slotRoot                   | bytes32                                   | 175  | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
| executionPayloadRoot       | bytes32                                   | 176  | 0      | 32    | src/test/utils/ProofParsing.sol:ProofParsing |
