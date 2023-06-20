// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./BLSRegistryCoordinatorWithIndices.sol";

import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";


contract BLSOperatorStateRetriever {
    IBLSRegistryCoordinatorWithIndices public registryCoordinator;
    IStakeRegistry public stakeRegistry;
    IBLSPubkeyRegistry public blsPubkeyRegistry;
    IIndexRegistry public indexRegistry;

    struct OperatorState {
        bytes32[] operatorIds;
        uint32 blockNumber;
        uint96[] stakePerQuorum;
        bytes32 apkHash;
        uint32 numOperators;
    }

    constructor(IBLSRegistryCoordinatorWithIndices _registryCoordinator) {
        registryCoordinator = _registryCoordinator;

        stakeRegistry = _registryCoordinator.stakeRegistry();
        blsPubkeyRegistry = _registryCoordinator.blsPubkeyRegistry();
        indexRegistry = _registryCoordinator.indexRegistry();
    }

    function getOperatorState(uint32 blockNumber, uint256 stakeHistoryIndex, uint256 globalApkHashIndex) external view returns (OperatorState memory) {
        OperatorState memory state;

        state.operatorIds = indexRegistry.getOperatorIds();
        state.blockNumber = blockNumber;

        uint96[] memory stakePerQuorum = new uint96[](256);

    
        for (uint quorumNumber = 0; quorumNumber < 256; quorumNumber++) {
            IStakeRegistry.OperatorStakeUpdate[] memory totalHistoryForQuorum = stakeRegistry.getTotalStakeHistoryForQuorum(quorumNumber);
            for (uint256 i = totalHistoryForQuorum.length - 1; i >= 0; i--) {
                IStakeRegistry.OperatorStakeUpdate memory update = totalHistoryForQuorum[i];

                if (blockNumber < update.blockNumber){
                    continue;
                }
                if (update.blockNumber == 0 || blockNumber < update.nextUpdateBlockNumber ) {
                    stakePerQuorum[quorumNumber] = update.stake;
                    break;
                }
            }
        }
        state.stakePerQuorum = stakePerQuorum;
        state.apkHash = blsPubkeyRegistry.getGlobalApkHashAtBlockNumberFromIndex(blockNumber, globalApkHashIndex);
        state.numOperators = indexRegistry.totalOperators();

        return state;
    }

    
}