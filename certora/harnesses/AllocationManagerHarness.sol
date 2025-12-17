// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AllocationManager.sol";
contract AllocationManagerHarness is AllocationManager {
    constructor(
        IDelegationManager _delegation,
        IStrategy _eigenStrategy,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY,
        string memory _version
    ) AllocationManager(
            _delegation,
            _eigenStrategy,
            _pauserRegistry,
            _permissionController,
            _DEALLOCATION_DELAY,
            _ALLOCATION_CONFIGURATION_DELAY,
            _version
        ) {}

    function getOperatorKey(address avs, uint32 operatorSetId) external view returns (bytes32) {
        return OperatorSetLib.key(OperatorSet(avs, operatorSetId));
    }

    function getOperatorSetFromKey(bytes32 key) external view returns (OperatorSet memory) {
        return OperatorSetLib.decode(key);
    }

    function getOperatorKeyFromSet(OperatorSet calldata os) external view returns (bytes32) {
        return OperatorSetLib.key(os);
    }
}