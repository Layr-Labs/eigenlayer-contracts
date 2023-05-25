// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IRegistry.sol";
import "../libraries/BN254.sol";

/**
 * @title Minimal interface for a registry that keeps track of aggregate operator public keys for among many quorums.
 * @author Layr Labs, Inc.
 */
interface IBLSPubkeyRegistry is IRegistry {
    /// @notice Data structure used to track the history of the Aggregate Public Key of all operators
    struct ApkUpdate {
        // keccak256(apk_x0, apk_x1, apk_y0, apk_y1)
        bytes32 apkHash;
        // block number at which the update occurred
        uint32 updateBlockNumber;
        // block number at which the next update occurred
        uint32 nextUpdateBlockNumber;
    }
    
    /**
     * @notice registers `operator` with the given `pubkey` for the specified `quorumNumbers`
     * @dev access restricted to the RegistryCoordinator
     */
    function registerOperator(address operator, uint8[] memory quorumNumbers, BN254.G1Point memory pubkey) external returns(bytes32);

    /**
     * @notice deregisters `operator` with the given `pubkey` for the specified `quorumNumbers`
     * @dev access restricted to the RegistryCoordinator
     */    
    function deregisterOperator(address operator, uint8[] memory quorumNumbers, BN254.G1Point memory pubkey) external returns(bytes32);

    /// @notice returns the current stored APK for the `quorumNumber`
    function quorumApk(uint8 quorumNumber) external view returns (BN254.G1Point memory);

    /// @notice returns the current stored APK among all quorums
    function globalApk() external view returns (BN254.G1Point memory);

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates for the `quorumNumber`
    function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) external view returns (ApkUpdate memory);

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates that keep track of the sum of the APK amonng all quorums
    function globalApkUpdates(uint256 index) external view returns (ApkUpdate memory);

    /**
     * @notice get hash of the apk of `quorumNumber` at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (bytes32);

	/**
     * @notice get hash of the apk among all quourms at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getGlobalApkHashAtBlockNumberFromIndex(uint32 blockNumber, uint256 index) external view returns (bytes32);
    
    ///@notice returns the length of the APK history for `quorumNumber`
    function getQuorumApkHistoryLength(uint8 quorumNumber) external view returns(uint32);

    ///@notice returns the length of the global APK history
    function getGlobalApkHistoryLength() external view returns(uint32);
}