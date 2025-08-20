// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract CrossChainRegistryMock {
    using OperatorSetLib for OperatorSet;

    mapping(bytes32 => bool) public generationReservations;

    function hasActiveGenerationReservation(OperatorSet memory operatorSet) external view returns (bool) {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        return generationReservations[key];
    }

    // Helper function for testing
    function setHasActiveGenerationReservation(OperatorSet memory operatorSet, bool hasReservation) external {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        generationReservations[key] = hasReservation;
    }
}
