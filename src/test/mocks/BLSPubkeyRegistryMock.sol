// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/libraries/BN254.sol";
import "../../contracts/interfaces/IBLSPubkeyRegistry.sol";
import "../../contracts/interfaces/IRegistryCoordinator.sol";

/**
 * @title Minimal interface for a registry that keeps track of aggregate operator public keys for among many quorums.
 * @author Layr Labs, Inc.
 */
contract BLSPubkeyRegistryMock is IBLSPubkeyRegistry {
    IRegistryCoordinator public registryCoordinator;

    function registerOperator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey) external returns(bytes32){}

    function deregisterOperator(address operator, bytes calldata quorumNumbers, BN254.G1Point memory pubkey) external returns(bytes32){}

    /// @notice Returns the current APK for the provided `quorumNumber `
    function getApkForQuorum(uint8 quorumNumber) external view returns (BN254.G1Point memory){}

    /// @notice Returns the `ApkUpdate` struct at `index` in the list of APK updates for the `quorumNumber`
    function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) external view returns (ApkUpdate memory){}

    function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (bytes32){
        return bytes32(0);
    }

	/**
     * @notice get hash of the apk among all quourums at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     * @param blockNumber is the number of the block for which the latest ApkHash muust be retrieved
     * @param index is the provided witness of the onchain index calculated offchain
     */
    function getGlobalApkHashAtBlockNumberFromIndex(uint32 blockNumber, uint256 index) external view returns (bytes32){}
    
    /// @notice Returns the length of ApkUpdates for the provided `quorumNumber`
    function getQuorumApkHistoryLength(uint8 quorumNumber) external view returns(uint32){}

    /// @notice Returns the length of ApkUpdates for the global APK
    function getGlobalApkHistoryLength() external view returns(uint32){}
}