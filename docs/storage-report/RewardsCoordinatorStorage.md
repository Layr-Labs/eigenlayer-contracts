| Name                                 | Type                                                    | Slot | Offset | Bytes | Contract                                                                   |
|--------------------------------------|---------------------------------------------------------|------|--------|-------|----------------------------------------------------------------------------|
| __deprecated_DOMAIN_SEPARATOR        | bytes32                                                 | 0    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| _distributionRoots                   | struct IRewardsCoordinatorTypes.DistributionRoot[]      | 1    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| rewardsUpdater                       | address                                                 | 2    | 0      | 20    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| activationDelay                      | uint32                                                  | 2    | 20     | 4     | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| currRewardsCalculationEndTimestamp   | uint32                                                  | 2    | 24     | 4     | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| globalOperatorCommissionBips         | uint16                                                  | 2    | 28     | 2     | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| claimerFor                           | mapping(address => address)                             | 3    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| cumulativeClaimed                    | mapping(address => mapping(contract IERC20 => uint256)) | 4    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| submissionNonce                      | mapping(address => uint256)                             | 5    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| isAVSRewardsSubmissionHash           | mapping(address => mapping(bytes32 => bool))            | 6    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| isRewardsSubmissionForAllHash        | mapping(address => mapping(bytes32 => bool))            | 7    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| isRewardsForAllSubmitter             | mapping(address => bool)                                | 8    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| isRewardsSubmissionForAllEarnersHash | mapping(address => mapping(bytes32 => bool))            | 9    | 0      | 32    | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
| __gap                                | uint256[40]                                             | 10   | 0      | 1280  | src/contracts/core/RewardsCoordinatorStorage.sol:RewardsCoordinatorStorage |
