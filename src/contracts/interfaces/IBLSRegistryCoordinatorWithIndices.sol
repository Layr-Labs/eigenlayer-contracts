// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IRegistryCoordinator.sol";
import "./IStakeRegistry.sol";
import "./IBLSPubkeyRegistry.sol";
import "./IIndexRegistry.sol";

/**
 * @title Minimal interface for the `IBLSStakeRegistryCoordinator` contract.
 * @author Layr Labs, Inc.
 */
interface IBLSRegistryCoordinatorWithIndices is IRegistryCoordinator {
    // STRUCT

    struct OperatorKickParam {
        address operator;
        BN254.G1Point pubkey; 
        bytes32[] operatorIdsToSwap; // should be a single length array when kicking
        uint32 globalOperatorListIndex;
    }

    /// @notice the stake registry for this corrdinator is the contract itself
    function stakeRegistry() external view returns (IStakeRegistry);
    /// @notice the BLS Pubkey Registry contract that will keep track of operators' BLS public keys
    function blsPubkeyRegistry() external view returns (IBLSPubkeyRegistry);
    /// @notice the Index Registry contract that will keep track of operators' indexes
    function indexRegistry() external view returns (IIndexRegistry);
}