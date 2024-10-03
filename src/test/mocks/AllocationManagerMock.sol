// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IAllocationManager.sol";

contract AllocationManagerMock is IAllocationManager, Test {
    
    function setAllocationDelay(address operator, uint32 delay) external { }

    function setAllocationDelay(
        uint32 delay
    ) external { }

    function modifyAllocations(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external { }

    function completePendingDeallocations(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToComplete
    ) external { }

    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint16 bipsToSlash
    ) external { }

    function cancelSalt(
        bytes32 salt
    ) external { }

    function allocationDelay(
        address operator
    ) external view returns (bool isSet, uint32 delay) { }

    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64) { }

    function getPendingAllocations(
        address operator,
        IStrategy strategy,
        OperatorSet[] calldata operatorSets
    ) external view returns (uint64[] memory, uint32[] memory) { }

    function getPendingDeallocations(
        address operator,
        IStrategy strategy,
        OperatorSet[] calldata operatorSets
    ) external view returns (uint64[] memory, uint32[] memory) { }

    function getSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (OperatorSet[] memory, uint64[][] memory) { }

    function getTotalMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory) { }

    function getTotalMagnitudesAtTimestamp(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint64[] memory) { }

    function getTotalMagnitude(address operator, IStrategy strategy) external view returns (uint64) { }

    function calculateMagnitudeAllocationDigestHash(
        address operator,
        MagnitudeAllocation[] calldata allocations,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32) { }

    function domainSeparator() external view returns (bytes32) { }
}