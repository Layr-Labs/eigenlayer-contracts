//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../contracts/interfaces/IIndexRegistry.sol";
import "../../contracts/middleware/IndexRegistry.sol";
import "../mocks/RegistryCoordinatorMock.sol";

import "forge-std/Test.sol";


contract IndexRegistryUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    IndexRegistry indexRegistry;
    RegistryCoordinatorMock registryCoordinatorMock;

    uint8 defaultQuorumNumber = 1;
    bytes32 defaultOperator = bytes32(uint256(34));


    function setUp() public {
        // deploy the contract
        registryCoordinatorMock = new RegistryCoordinatorMock();
        indexRegistry = new IndexRegistry(registryCoordinatorMock);
    }

    function testConstructor() public {
        // check that the registry coordinator is set correctly
        assertEq(address(indexRegistry.registryCoordinator()), address(registryCoordinatorMock), "IndexRegistry.constructor: registry coordinator not set correctly");
    }

    function testRegisterOperatorInIndexRegistry(bytes32 operatorId) public {
        // register an operator
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        cheats.startPrank(address(registryCoordinatorMock));
       indexRegistry.registerOperator(operatorId, quorumNumbers);
        cheats.stopPrank();

       require(indexRegistry.globalOperatorList(0) == operatorId, "IndexRegistry.registerOperator: operator not registered correctly");
       require(indexRegistry.totalOperators() == 1, "IndexRegistry.registerOperator: total operators not updated correctly");
       require(indexRegistry.totalOperatorsForQuorum(1) == 1, "IndexRegistry.registerOperator: total operators for quorum not updated correctly");
       IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry.getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId, 1, 0);
       require(indexUpdate.index == 0, "IndexRegistry.registerOperator: index not 0");
       require(indexUpdate.toBlockNumber == 0, "block number should not be set");
    }

    function testRegisterOperatorFromNonRegisterCoordinator(address nonRegistryCoordinator) public {
        cheats.assume(address(registryCoordinatorMock) != nonRegistryCoordinator);
        // register an operator
        bytes memory quorumNumbers = new bytes(defaultQuorumNumber);

        cheats.startPrank(nonRegistryCoordinator);
        cheats.expectRevert(bytes("IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        indexRegistry.registerOperator(bytes32(0), quorumNumbers);
        cheats.stopPrank();
    }

    function testDeregisterOperatorFromNonRegisterCoordinator(address nonRegistryCoordinator) public {
        cheats.assume(address(registryCoordinatorMock) != nonRegistryCoordinator);
        // register an operator
        bytes memory quorumNumbers = new bytes(defaultQuorumNumber);
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);

        cheats.startPrank(nonRegistryCoordinator);
        cheats.expectRevert(bytes("IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        indexRegistry.deregisterOperator(bytes32(0), true, quorumNumbers, operatorIdsToSwap, 0);
        cheats.stopPrank();
    }

    function testDeregisterOperatorInIndexRegistry(bytes32 operatorId1, bytes32 operatorId2) public {
        cheats.assume(operatorId1 != operatorId2);
        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        quorumNumbers[1] = bytes1(defaultQuorumNumber + 1);

        _registerOperator(operatorId1, quorumNumbers);
        _registerOperator(operatorId2, quorumNumbers);

        require(indexRegistry.totalOperators() == 2, "IndexRegistry.registerOperator: operator not registered correctly");
        require(indexRegistry.totalOperatorsForQuorum(1) == 2, "IndexRegistry.registerOperator: operator not registered correctly");
        require(indexRegistry.totalOperatorsForQuorum(2) == 2, "IndexRegistry.registerOperator: operator not registered correctly");

        bytes32[] memory operatorIdsToSwap = new bytes32[](2);
        operatorIdsToSwap[0] = operatorId2;
        operatorIdsToSwap[1] = operatorId2;

        cheats.roll(block.number + 1);

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        cheats.startPrank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, true, quorumNumbers, operatorIdsToSwap, 0);
        cheats.stopPrank();

        require(indexRegistry.totalOperators() == 1, "IndexRegistry.registerOperator: operator not registered correctly");
        require(indexRegistry.globalOperatorList(0) == operatorId2, "IndexRegistry.registerOperator: operator not deregistered and swapped correctly");

        IIndexRegistry.OperatorIndexUpdate memory indexUpdate1 = indexRegistry.getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId1, defaultQuorumNumber, 0);
        require(indexUpdate1.toBlockNumber == block.number, "toBlockNumber not set correctly");
        require(indexUpdate1.index == 0, "incorrect index");

        IIndexRegistry.OperatorIndexUpdate memory indexUpdate2 = indexRegistry.getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId2, defaultQuorumNumber, 1);
        require(indexUpdate2.toBlockNumber == 0, "toBlockNumber not set correctly");
        require(indexUpdate2.index == 0, "incorrect index");

    }

    function testDeregisterOperatorWithIncorrectOperatorToSwap(bytes32 operatorId1, bytes32 operatorId2, bytes32 operatorId3) public {
        cheats.assume(operatorId1 != operatorId2 && operatorId3 != operatorId2 && operatorId3 != operatorId1);

        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        _registerOperator(operatorId1, quorumNumbers);
        _registerOperator(operatorId2, quorumNumbers);
        _registerOperator(operatorId3, quorumNumbers);

        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = operatorId2;

        cheats.roll(block.number + 1);

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        cheats.startPrank(address(registryCoordinatorMock));
        cheats.expectRevert(bytes("IndexRegistry._processOperatorRemoval: operatorIdToSwap is not the last operator in the quorum"));
        indexRegistry.deregisterOperator(operatorId1, true, quorumNumbers, operatorIdsToSwap, 0);
        cheats.stopPrank();
    }

    function testDeregisterOperatorWithMismatchInputLengths() public {
        bytes memory quorumNumbers = new bytes(1);
        bytes32[] memory operatorIdsToSwap = new bytes32[](2);

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        cheats.startPrank(address(registryCoordinatorMock));
        cheats.expectRevert(bytes("IndexRegistry.deregisterOperator: quorumNumbers and operatorIdsToSwap must be the same length"));
        indexRegistry.deregisterOperator(defaultOperator, true, quorumNumbers, operatorIdsToSwap, 0);
        cheats.stopPrank();
    }

    function testTotalOperatorUpdatesForOneQuorum(uint8 numOperators) public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        
        uint256 lengthBefore = 0;
        for (uint256 i = 0; i < numOperators; i++) {
            _registerOperator(bytes32(i), quorumNumbers);
            require(indexRegistry.totalOperatorsForQuorum(1) - lengthBefore == 1, "incorrect update");
            require(indexRegistry.totalOperators() - lengthBefore == 1, "incorrect update");
            lengthBefore++;
        }
    }



    function _registerOperator(bytes32 operatorId, bytes memory quorumNumbers) public {
        cheats.startPrank(address(registryCoordinatorMock));
       indexRegistry.registerOperator(operatorId, quorumNumbers);
        cheats.stopPrank();
    }
    
    // function _deregisterOperator(bytes32 operatorId, bytes memory quorumNumbers, uint32[] memory quorumToOperatorListIndexes, uint32 index) public {
    //     cheats.startPrank(address(registryCoordinatorMock));
    //     indexRegistry.deregisterOperator(operatorId, quorumNumbers, quorumToOperatorListIndexes, index);
    //     cheats.stopPrank();
    // }

    // function _getRandomId(uint256 seed) internal view returns (bytes32) {
    //     return keccak256(abi.encodePacked(block.timestamp, seed));
    // }

    // function _generateRandomNumber(uint256 seed, uint256 modulus) internal view returns (uint256) {
    //     uint256 randomNumber = uint256(keccak256(abi.encodePacked(block.timestamp, seed)));
    //     return (randomNumber % modulus); 
    // }
}
