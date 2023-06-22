// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./BLSRegistryCoordinatorWithIndices.sol";

import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";


contract BLSOperatorStateRetriever {
    struct Operator {
        bytes32 operatorId;
        uint96 stake;
    }

    struct OperatorState {
        Operator[][] operators;
        uint256 callingOperatorIndex;
    }

    IBLSRegistryCoordinatorWithIndices public registryCoordinator;
    IStakeRegistry public stakeRegistry;
    IBLSPubkeyRegistry public blsPubkeyRegistry;
    IIndexRegistry public indexRegistry;

    constructor(IBLSRegistryCoordinatorWithIndices _registryCoordinator) {
        registryCoordinator = _registryCoordinator;

        stakeRegistry = _registryCoordinator.stakeRegistry();
        blsPubkeyRegistry = _registryCoordinator.blsPubkeyRegistry();
        indexRegistry = _registryCoordinator.indexRegistry();
    }

    /**
     * @notice returns the ordered list of operators (id and stake) for each quorum
     * @param quorumNumbers the quorum numbers to get the operator state for
     */
    function getOperatorState(bytes32 callingOperatorId, bytes calldata quorumNumbers) external view returns (OperatorState memory) {
        Operator[][] memory operators = new Operator[][](quorumNumbers.length);
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = uint8(quorumNumbers[i]);
            bytes32[] memory operatorIds = indexRegistry.getOperatorListForQuorum(quorumNumber);
            operators[i] = new Operator[](operatorIds.length);
            for (uint256 j = 0; j < operatorIds.length; j++) {
                bytes32 operatorId = bytes32(operatorIds[j]);
                operators[i][j] = Operator({
                    operatorId: operatorId,
                    stake: stakeRegistry.getMostRecentStakeUpdateByOperatorId(operatorId, quorumNumber).stake
                });
            }
        }

        OperatorState memory operatorState = OperatorState({
            operators: operators,
            callingOperatorIndex: indexRegistry.getIndexOfOperatorIdInGlobalOperatorList(callingOperatorId)
        });

        return operatorState;
    }

    
}