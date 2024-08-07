// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../interfaces/IStrategy.sol";


interface IStakeRootCompendium {

    event SnarkProofVerified(bytes journal, bytes seal);

    event VerifierChanged(address oldVerifier, address newVerifier);

    event ImageIdChanged(bytes32 oldImageId, bytes32 newImageId);

    function getOperatorSetRoot(address avs, uint32 operatorSetId, address[] calldata operators, IStrategy[] calldata strategies) external view returns (bytes32);

    function getStakeRoot(address[] calldata avss, bytes32[] calldata operatorSetRoots) external view returns (bytes32);

    function verifySnarkProof(
        bytes calldata _journal,
        bytes calldata _seal
    ) external;

    function setImageId(bytes32 _imageId) external;

    function setVerifier(address _verifier) external;
}