//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {IIndexRegistry} from "src/contracts/interfaces/IIndexRegistry.sol";
import {IndexRegistry} from "src/contracts/middleware/IndexRegistry.sol";
import {RegistryCoordinatorMock} from "src/test/mocks/RegistryCoordinatorMock.sol";

import {Test, console2} from "forge-std/Test.sol";

contract IndexRegistryUnitTests is Test {
    IndexRegistry internal indexRegistry;
    RegistryCoordinatorMock internal registryCoordinatorMock;
    bytes internal NOT_REGISTRY_COORDINATOR_ERROR =
        bytes("IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator");
    bytes internal OPERATOR_NOT_LAST_OPERATOR_IN_QUORUM_ERROR =
        bytes("IndexRegistry._processOperatorRemoval: operatorIdToSwap is not the last operator in the quorum");
    bytes internal QUORUM_NUMBERS_AND_SWAP_IDS_LENGTH_MISMATCH_ERROR =
        bytes("IndexRegistry.deregisterOperator: quorumNumbers and operatorIdsToSwap must be the same length");

    uint8 internal defaultQuorumNumber = 1;
    bytes32 internal defaultOperator = bytes32(uint256(34));
    bytes32 defaultOperatorId = keccak256("defaultOperatorId");

    function setUp() public {
        registryCoordinatorMock = new RegistryCoordinatorMock();
        indexRegistry = new IndexRegistry(registryCoordinatorMock);
    }

    /// CONSTRUCTOR
    function test_Constructor() public {
        // check that the registry coordinator is set correctly
        assertEq(
            address(indexRegistry.registryCoordinator()),
            address(registryCoordinatorMock),
            "IndexRegistry.constructor: registry coordinator not set correctly"
        );
    }

    function test_RegisterOperator() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        vm.prank(address(registryCoordinatorMock));
        indexRegistry.registerOperator(defaultOperatorId, quorumNumbers);

        assertTrue(indexRegistry.globalOperatorList(0) == defaultOperatorId, "operator not registered correctly");
        assertTrue(indexRegistry.totalOperatorsForQuorum(1) == 1, "total operators for quorum not updated correctly");
        IIndexRegistry.OperatorIndexUpdate memory indexUpdate =
            indexRegistry.getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(defaultOperatorId, 1, 0);
        assertTrue(indexUpdate.index == 0, "index not 0");
        assertTrue(indexUpdate.fromBlockNumber == block.number, "block number should not be set");
    }

    function test_DeregisterOperator() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(defaultOperatorId, quorumNumbers);

        assertTrue(indexRegistry.totalOperatorsForQuorum(1) == 1, "operator not registered correctly");

        vm.roll(block.number + 1);
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        vm.prank(address(registryCoordinatorMock));

        /// TODO: Should an operator that was mistakenly registered be able to deregister
        vm.expectRevert();
        indexRegistry.deregisterOperator(defaultOperatorId, quorumNumbers, operatorIdsToSwap);

        assertTrue(indexRegistry.totalOperatorsForQuorum(1) == 1, "operator not deregistered correctly");
    }

    /// Test view functions
    function test_totalOperatorsForQuorum() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        // Register 3 operators
        _registerOperator(bytes32(uint256(1)), quorumNumbers);
        _registerOperator(bytes32(uint256(2)), quorumNumbers);
        _registerOperator(bytes32(uint256(3)), quorumNumbers);

        // Check if the total number of operators for the quorum is 3
        assertTrue(
            indexRegistry.totalOperatorsForQuorum(defaultQuorumNumber) == 3,
            "total operators for quorum not updated correctly"
        );
    }

    function test_GetTotalOperatorsUpdateForQuorumAtBlockNumberByIndex() public {
        uint256 blockNumber = 10;
        vm.roll(blockNumber);
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        // Register 3 operators
        _registerOperator(bytes32(uint256(1)), quorumNumbers);
        _registerOperator(bytes32(uint256(2)), quorumNumbers);
        _registerOperator(bytes32(uint256(3)), quorumNumbers);

        vm.roll(block.number + 1);
        // Get the total operators update for the quorum at index 0
        IIndexRegistry.OperatorIndexUpdate memory update =
            indexRegistry.getTotalOperatorsUpdateForQuorumAtIndex(defaultQuorumNumber, 0);

        // Check if the total number of operators in the update is 3
        assertTrue(update.index == 1, "index for quorum not updated correctly at index 0");

        // Check if the block number in the update is the current block number
        assertTrue(update.fromBlockNumber == blockNumber, "block number in update not set correctly");
    }

    function test_RevertsWhen_IndexTooFarInPast_GetTotalOperatorsUpdateForQuorumAtBlockNumberByIndex() public {}

    function test_RevertsWhen_IndexTooFarInFuture_GetTotalOperatorsForQuorumAtBlockNumberByIndex() public {}

    function test_getTotalOperatorsForQuorumAtBlockNumber() public {}

    function test_When_BeforeQuorumExisted_getTotalOperatorsForQuorumAtBlockNumber() public {}

    function test_When_NoOperatorsInQuorum_getTotalOperatorsForQuorumAtBlockNumber() public {}

    function test_When_OnlyOneOperatorInQuorum_getTotalOperatorsForQuorumAtBlockNumber() public {}

    function test_getIndexOfOperatorForQuorumAtBlockNumber() public {}

    function test_When_OperatorNotRegisteredForQuorumAtBlock_getIndexOfOperatorForQuorumAtBlockNumber() public {}

    /// REGISTER OPERATOR
    function testFuzz_When_OperatorInIndexRegistry_RegisterOperator(bytes32 operatorId) public {
        // register an operator
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        vm.startPrank(address(registryCoordinatorMock));
        indexRegistry.registerOperator(operatorId, quorumNumbers);
        vm.stopPrank();

        assertTrue(indexRegistry.globalOperatorList(0) == operatorId, "operator not registered correctly");
        assertTrue(indexRegistry.totalOperatorsForQuorum(1) == 1, "total operators for quorum not updated correctly");
        IIndexRegistry.OperatorIndexUpdate memory indexUpdate =
            indexRegistry.getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId, 1, 0);
        assertTrue(indexUpdate.index == 0, "index not 0");
        assertTrue(indexUpdate.fromBlockNumber == block.number, "block number should not be set");
    }

    function testFuzz_RevertsWhen_NotRegistryCoordinator_RegisterOperator(address nonRegistryCoordinator) public {
        vm.assume(address(registryCoordinatorMock) != nonRegistryCoordinator);
        // register an operator
        bytes memory quorumNumbers = new bytes(defaultQuorumNumber);

        vm.startPrank(nonRegistryCoordinator);
        vm.expectRevert(NOT_REGISTRY_COORDINATOR_ERROR);
        indexRegistry.registerOperator(bytes32(0), quorumNumbers);
        vm.stopPrank();
    }

    function testFuzz_RevertsWhen_NotRegistryCoordinator_DeregistorOperator(address nonRegistryCoordinator) public {
        vm.assume(address(registryCoordinatorMock) != nonRegistryCoordinator);
        // register an operator
        bytes memory quorumNumbers = new bytes(defaultQuorumNumber);
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);

        vm.startPrank(nonRegistryCoordinator);
        vm.expectRevert(NOT_REGISTRY_COORDINATOR_ERROR);
        indexRegistry.deregisterOperator(bytes32(0), quorumNumbers, operatorIdsToSwap);
        vm.stopPrank();
    }

    /// DEREGISTER OPERATOR
    function testFuzz_When_OperatorInIndexRegistry_DeregisterOperator(bytes32 operatorId1, bytes32 operatorId2)
        public
    {
        vm.assume(operatorId1 != operatorId2);
        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        quorumNumbers[1] = bytes1(defaultQuorumNumber + 1);

        _registerOperator(operatorId1, quorumNumbers);
        _registerOperator(operatorId2, quorumNumbers);

        assertTrue(indexRegistry.totalOperatorsForQuorum(1) == 2, "operator not registered correctly");
        assertTrue(indexRegistry.totalOperatorsForQuorum(2) == 2, "operator not registered correctly");

        bytes32[] memory operatorIdsToSwap = new bytes32[](2);
        operatorIdsToSwap[0] = operatorId2;
        operatorIdsToSwap[1] = operatorId2;

        vm.roll(block.number + 1);

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        vm.startPrank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, operatorIdsToSwap);
        vm.stopPrank();

        assertTrue(indexRegistry.totalOperatorsForQuorum(1) == 1, "operator not deregistered correctly");
        assertTrue(indexRegistry.totalOperatorsForQuorum(2) == 1, "operator not deregistered correctly");

        IIndexRegistry.OperatorIndexUpdate memory indexUpdate =
            indexRegistry.getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId2, defaultQuorumNumber, 1);
        assertTrue(indexUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
        assertTrue(indexUpdate.index == 0, "incorrect index");
    }

    function testFuzz_RevertsWhen_OperatorToSwapIncorrect_DeregisterOperator(
        bytes32 operatorId1,
        bytes32 operatorId2,
        bytes32 operatorId3
    ) public {
        vm.assume(operatorId1 != operatorId2 && operatorId3 != operatorId2 && operatorId3 != operatorId1);

        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        _registerOperator(operatorId1, quorumNumbers);
        _registerOperator(operatorId2, quorumNumbers);
        _registerOperator(operatorId3, quorumNumbers);

        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = operatorId2;

        vm.roll(block.number + 1);

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        vm.startPrank(address(registryCoordinatorMock));
        vm.expectRevert(OPERATOR_NOT_LAST_OPERATOR_IN_QUORUM_ERROR);
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, operatorIdsToSwap);
        vm.stopPrank();
    }

    function test_RevertsWhen_LenghtMismatchForOperatorIdsAndQuorums_DeregisterOperator() public {
        bytes memory quorumNumbers = new bytes(1);
        bytes32[] memory operatorIdsToSwap = new bytes32[](2);

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        vm.startPrank(address(registryCoordinatorMock));
        vm.expectRevert(QUORUM_NUMBERS_AND_SWAP_IDS_LENGTH_MISMATCH_ERROR);
        indexRegistry.deregisterOperator(defaultOperator, quorumNumbers, operatorIdsToSwap);
        vm.stopPrank();
    }

    function testFuzz_WhenUpdatesForOneQuorum_TotalOperatorsForQuorum(uint8 numOperators) public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        uint256 lengthBefore = 0;
        for (uint256 i = 0; i < numOperators; i++) {
            _registerOperator(bytes32(i), quorumNumbers);
            assertTrue(indexRegistry.totalOperatorsForQuorum(1) - lengthBefore == 1, "incorrect update");
            lengthBefore++;
        }
    }

    function _registerOperator(bytes32 operatorId, bytes memory quorumNumbers) public {
        vm.startPrank(address(registryCoordinatorMock));
        indexRegistry.registerOperator(operatorId, quorumNumbers);
        vm.stopPrank();
    }
}
