// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/BLSRegistryCoordinatorWithIndices.sol";

// wrapper around the BLSRegistryCoordinatorWithIndices contract that exposes the internal functions for unit testing.
contract BLSRegistryCoordinatorWithIndicesHarness is BLSRegistryCoordinatorWithIndices {
    constructor(
        ISlasher _slasher,
        IServiceManager _serviceManager,
        IStakeRegistry _stakeRegistry,
        IBLSPubkeyRegistry _blsPubkeyRegistry,
        IIndexRegistry _indexRegistry
    ) BLSRegistryCoordinatorWithIndices(_slasher, _serviceManager, _stakeRegistry, _blsPubkeyRegistry, _indexRegistry) {
    }

    function recordOperatorQuorumBitmapUpdate(bytes32 operatorId, uint192 quorumBitmap) external {
        uint256 operatorQuorumBitmapHistoryLength = _operatorIdToQuorumBitmapHistory[operatorId].length;
        if (operatorQuorumBitmapHistoryLength != 0) {
            _operatorIdToQuorumBitmapHistory[operatorId][operatorQuorumBitmapHistoryLength - 1].nextUpdateBlockNumber = uint32(block.number);
        }

        _operatorIdToQuorumBitmapHistory[operatorId].push(QuorumBitmapUpdate({
            updateBlockNumber: uint32(block.number),
            nextUpdateBlockNumber: 0,
            quorumBitmap: quorumBitmap
        }));
    }

    
}