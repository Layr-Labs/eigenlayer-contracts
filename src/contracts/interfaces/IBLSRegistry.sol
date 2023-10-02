// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IQuorumRegistry.sol";
import "../libraries/BN254.sol";


/**
 * @title Minimal interface extension to `IQuorumRegistry`.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Adds BLS-specific functions to the base interface.
 */
interface IBLSRegistry is IQuorumRegistry {
    /// @notice Data structure used to track the history of the Aggregate Public Key of all operators
    struct ApkUpdate {
        // keccak256(apk_x0, apk_x1, apk_y0, apk_y1)
        bytes32 apkHash;
        // block number at which the update occurred
        uint32 blockNumber;
    }

    // EVENTS
    /**
     * @notice Emitted upon the registration of a new operator for the middleware
     * @param operator Address of the new operator
     * @param pkHash The keccak256 hash of the operator's public key
     * @param pk The operator's public key itself
     * @param apkHashIndex The index of the latest (i.e. the new) APK update
     * @param apkHash The keccak256 hash of the new Aggregate Public Key
     */
    event Registration(
        address indexed operator,
        bytes32 pkHash,
        BN254.G1Point pk,
        uint32 apkHashIndex,
        bytes32 apkHash,
        string socket
    );

    /// @notice Emitted when the `operatorWhitelister` role is transferred.
    event OperatorWhitelisterTransferred(address previousAddress, address newAddress);

    /**
     * @notice get hash of a historical aggregated public key corresponding to a given index;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getCorrectApkHash(uint256 index, uint32 blockNumber) external returns (bytes32);

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates
    function apkUpdates(uint256 index) external view returns (ApkUpdate memory);

    /// @notice returns the APK hash that resulted from the `index`th APK update
    function apkHashes(uint256 index) external view returns (bytes32);

    /// @notice returns the block number at which the `index`th APK update occurred
    function apkUpdateBlockNumbers(uint256 index) external view returns (uint32);

    function operatorWhitelister() external view returns (address);

    function operatorWhitelistEnabled() external view returns (bool);

    function whitelisted(address) external view returns (bool);

    function setOperatorWhitelistStatus(bool _operatorWhitelistEnabled) external;

    function addToOperatorWhitelist(address[] calldata) external;

    function removeFromWhitelist(address[] calldata operators) external;
}
