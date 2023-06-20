// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./BLSIndexRegistryCoordinator.sol";

import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IBLSStakeRegistryCoordinator.sol";


contract BLSOperatorStateRetriever {
    IBLSStakeRegistryCoordinator public registryCoordinator;
    IStakeRegistry public stakeRegistry;
    IBLSPubkeyRegistry public blsPubkeyRegistry;
    IIndexRegistry public indexRegistry;

    struct OperatorState {
        uint32 blockNumber;
        uint96[] stakePerQuorum;
        bytes32 apkHash;
        uint32 numOperators;
    }

    constructor(IBLSStakeRegistryCoordinator _registryCoordinator) {
        registryCoordinator = _registryCoordinator;

        stakeRegistry = _registryCoordinator.stakeRegistry();
        blsPubkeyRegistry = _registryCoordinator.blsPubkeyRegistry();
        indexRegistry = _registryCoordinator.indexRegistry();
    }

    function getOperatorState(uint32 blockNumber, uint256 stakeHistoryIndex, uint256 globalApkHashIndex) external view returns (OperatorState memory) {
        OperatorState memory state;
        state.blockNumber = blockNumber;

        uint96[] memory stakePerQuorum = new uint96[](256);
        for (uint quorumNumber = 0; quorumNumber < 256; quorumNumber++) {
            stakePerQuorum[quorumNumber] = stakeRegistry.getTotalStakeAtBlockNumberFromIndex(quorumNumber, blockNumber, stakeHistoryIndex);
        }
        state.stakePerQuorum = stakePerQuorum;

        state.apkHash = blsPubkeyRegistry.getGlobalApkHashAtBlockNumberFromIndex(blockNumber, globalApkHashIndex);
        state.numOperators = indexRegistry.totalOperators();

        return state;
    }

    
}