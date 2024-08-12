// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IStrategy.sol";

interface IStakeRootCompendium {

    struct StrategyAndMultiplier {
        IStrategy strategy;
        uint96 multiplier;
    }

    struct StakeRootLeaf {
        IAVSDirectory.OperatorSet operatorSet;
        bytes32 operatorSetRoot;
    }

    event SnarkProofVerified(bytes journal, bytes seal);
    event VerifierChanged(address oldVerifier, address newVerifier);
    event ImageIdChanged(bytes32 oldImageId, bytes32 newImageId);

    function MAX_OPERATOR_SET_SIZE() external view returns (uint32);
    function MAX_NUM_OPERATOR_SETS() external view returns (uint32);
    function MAX_NUM_STRATEGIES() external view returns (uint32);

    function delegation() external view returns (IDelegationManager);
    function avsDirectory() external view returns (IAVSDirectory);
    function verifier() external view returns (address);
    function imageId() external view returns (bytes32);
    function numConfiguredOperatorSets() external view returns (uint256);

    function getStakeRoot(StakeRootLeaf[] calldata stakeRootLeaves) external view returns (bytes32);

    function getOperatorSetRoot(
        IAVSDirectory.OperatorSet calldata operatorSet, 
        address[] calldata operators
    ) external view returns (bytes32);

    function addStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata strategiesAndMultipliers
    ) external;

    function removeStrategiesAndMultipliers(
        uint32 operatorSetId,
        IStrategy[] calldata strategies
    ) external;

    function setVerifier(address _verifier) external;
    function setImageId(bytes32 _imageId) external;
}
