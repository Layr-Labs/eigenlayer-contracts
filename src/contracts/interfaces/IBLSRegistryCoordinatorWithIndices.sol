// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./ISignatureUtils.sol";
import "./IRegistryCoordinator.sol";
import "./IStakeRegistry.sol";
import "./IBLSPubkeyRegistry.sol";
import "./IIndexRegistry.sol";

/**
 * @title Minimal interface for the `IBLSStakeRegistryCoordinator` contract.
 * @author Layr Labs, Inc.
 */
interface IBLSRegistryCoordinatorWithIndices is ISignatureUtils, IRegistryCoordinator {
    // STRUCT

    /**
     * @notice Data structure for storing operator set params for a given quorum. Specifically the 
     * `maxOperatorCount` is the maximum number of operators that can be registered for the quorum,
     * `kickBIPsOfOperatorStake` is the basis points of a new operator needs to have of an operator they are trying to kick from the quorum,
     * and `kickBIPsOfTotalStake` is the basis points of the total stake of the quorum that an operator needs to be below to be kicked.
     */ 
    struct OperatorSetParam {
        uint32 maxOperatorCount;
        uint16 kickBIPsOfOperatorStake;
        uint16 kickBIPsOfTotalStake;
    }

    /**
     * @notice Data structure for the parameters needed to kick an operator from a quorum with number `quorumNumber`, used during registration churn.
     * Specifically the `operator` is the address of the operator to kick, `pubkey` is the BLS public key of the operator,
     * `operatorIdsToSwap` is the list of operatorIds to swap with the operator being kicked in the indexRegistry,
     * and `globalOperatorListIndex` is the index of the operator in the global operator list in the indexRegistry.
     */
    struct OperatorKickParam {
        uint8 quorumNumber;
        address operator;
        BN254.G1Point pubkey; 
        bytes32[] operatorIdsToSwap; // should be a single length array when kicking
    }

    // EVENTS

    event OperatorSetParamsUpdated(uint8 indexed quorumNumber, OperatorSetParam operatorSetParams);

    event ChurnApproverUpdated(address churnApprover);

    /// @notice Returns the operator set params for the given `quorumNumber`
    function getOperatorSetParams(uint8 quorumNumber) external view returns (OperatorSetParam memory);
    /// @notice the stake registry for this corrdinator is the contract itself
    function stakeRegistry() external view returns (IStakeRegistry);
    /// @notice the BLS Pubkey Registry contract that will keep track of operators' BLS public keys
    function blsPubkeyRegistry() external view returns (IBLSPubkeyRegistry);
    /// @notice the Index Registry contract that will keep track of operators' indexes
    function indexRegistry() external view returns (IIndexRegistry);
}