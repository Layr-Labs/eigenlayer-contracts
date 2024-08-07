// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../interfaces/IStrategy.sol";


interface IStakeRootCompendium {
    struct StrategyAndMultiplier {
        IStrategy strategy;
        uint96 multiplier;
    }
        
    /**
    * @notice called by an AVS to set their strategies and multipliers used to determine stakes for stake roots
    * @param operatorSetId the operator set to set the strategies and multipliers for
    * @dev msg.sender is used as the AVS in determining the operator set
    */
     function setStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata strategiesAndMultipliers
    ) external;

    event SnarkProofVerified(bytes journal, bytes seal);

    event VerifierChanged(address oldVerifier, address newVerifier);

    event ImageIdChanged(bytes32 oldImageId, bytes32 newImageId);

    function getOperatorSetRoot(address avs, uint32 operatorSetId, address[] calldata operators, IStrategy[] calldata strategies) external view returns (bytes32);

    function getStakeRoot(address[] calldata avss, bytes32[] calldata operatorSetRoots) external view returns (bytes32);

    function getMultipliers(uint32 operatorSetId, IStrategy[] calldata strategies) external view returns (uint96[] memory);

    function verifySnarkProof(
        bytes calldata _journal,
        bytes calldata _seal
    ) external;

    function setImageId(bytes32 _imageId) external;

    function setVerifier(address _verifier) external;
}