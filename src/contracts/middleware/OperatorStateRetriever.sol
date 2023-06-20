// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./BLSIndexRegistryCoordinator.sol";

import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IBLSStakeRegistryCoordinator.sol";


contract OperatorStateRetriever {
    IBLSStakeRegistryCoordinator public registryCoordinator;
    IStakeRegistry public stakeRegistry;
    IBLSPubkeyRegistry public blsPubkeyRegistry;
    IIndexRegistry public indexRegistry;

    struct OperatorView {
        IBLSStakeRegistryCoordinator.Operator operator;
        uint8[] quorumNumbers;
        uint256[] quorumStake;
    }

    struct OperatorState {
        OperatorView[] operatorView;
        uint32[] operatorIndices;
        uint32 blockNumber;
        uint256[] stakePerQuorum;
        bytes32 apkHash;
        BN254.G1Point apkG1;
        uint32 numOperators;
    }

    constructor(IBLSStakeRegistryCoordinator _registryCoordinator) {
        registryCoordinator = _registryCoordinator;

        stakeRegistry = _registryCoordinator.stakeRegistry();
        blsPubkeyRegistry = _registryCoordinator.blsPubkeyRegistry();
        indexRegistry = _registryCoordinator.indexRegistry();
    }

    function getOperatorState(address operator) external view returns (OperatorState memory) {

    }

    
}