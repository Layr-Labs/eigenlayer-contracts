// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/interfaces/IAVSDirectory.sol";

interface IAVSDirectoryEvents {
    /// @notice Emitted when an operator set is created by an AVS.
    event OperatorSetCreated(IAVSDirectory.OperatorSet operatorSet);

    /**
     *  @notice Emitted when an operator's registration status with an AVS id udpated
     *  @notice Only used by legacy M2 AVSs that have not integrated with operatorSets.
     */
    event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, IAVSDirectory.OperatorAVSRegistrationStatus status);

    /// @notice Emitted when an operator is added to an operator set.
    event OperatorAddedToOperatorSet(address indexed operator, IAVSDirectory.OperatorSet operatorSet);

    /// @notice Emitted when an operator is removed from an operator set.
    event OperatorRemovedFromOperatorSet(address indexed operator, IAVSDirectory.OperatorSet operatorSet);

    /// @notice Emitted when an AVS updates their metadata URI (Uniform Resource Identifier).
    /// @dev The URI is never stored; it is simply emitted through an event for off-chain indexing.
    event AVSMetadataURIUpdated(address indexed avs, string metadataURI);

    /// @notice Emitted when an AVS migrates to using operator sets
    event AVSMigratedToOperatorSets(address indexed avs);

    /// @notice Emitted when an operator is migrated from M2 registration to operator sets.
    event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds);
}
