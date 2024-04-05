// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/interfaces/IAVSDirectory.sol";

interface IAVSDirectoryEvents {
    /**
     * @notice Emitted when @param avs indicates that they are updating their MetadataURI string
     * @dev Note that these strings are *never stored in storage* and are instead purely emitted in events for off-chain indexing
     */
    event AVSMetadataURIUpdated(address indexed avs, string metadataURI);

    /// @notice Emitted when an operator's registration status for an AVS is updated
    event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, IAVSDirectory.OperatorAVSRegistrationStatus status);
}
