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
       require(indexRegistry.totalOperators() == 1, "IndexRegistry.registerOperator: operator not registered correctly");
       require(indexRegistry.totalOperatorsForQuorum(1) == 1, "IndexRegistry.registerOperator: operator not registered correctly");
       (uint32 index, uint32 toBlockNumber) = indexRegistry.operatorIdToIndexHistory(operatorId, 1, 0);
       require(index == 0, "IndexRegistry.registerOperator: operator not registered correctly");
       require(toBlockNumber == 0, "block number should not be set");
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
        uint32[] memory quorumToOperatorListIndexes = new uint32[](1);

        cheats.startPrank(nonRegistryCoordinator);
        cheats.expectRevert(bytes("IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        indexRegistry.deregisterOperator(bytes32(0), quorumNumbers, quorumToOperatorListIndexes, 0);
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

        uint32[] memory quorumToOperatorListIndexes = new uint32[](2);
        quorumToOperatorListIndexes[0] = 0;
        quorumToOperatorListIndexes[1] = 0;

        cheats.roll(block.number + 1);

        // deregister the operatorId1, removing it from both quorum 1 and 2.
        cheats.startPrank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, quorumToOperatorListIndexes, 0);
        cheats.stopPrank();

        require(indexRegistry.totalOperators() == 1, "IndexRegistry.registerOperator: operator not registered correctly");
        require(indexRegistry.quorumToOperatorList(1, 0) == operatorId2, "IndexRegistry.registerOperator: operator not deregistered and swapped correctly");
        require(indexRegistry.quorumToOperatorList(2, 0) == operatorId2, "IndexRegistry.registerOperator: operator not deregistered and swapped correctly");
        require(indexRegistry.globalOperatorList(0) == operatorId2, "IndexRegistry.registerOperator: operator not deregistered and swapped correctly");

        (uint32 toBlockNumber1, uint32 index1) = indexRegistry.operatorIdToIndexHistory(operatorId1, defaultQuorumNumber, 0);

        require(toBlockNumber1 == block.number, "toBlockNumber not set correctly");
        require(index1 == 0, "incorrect index");

        (uint32 toBlockNumber2, uint32 index2) = indexRegistry.operatorIdToIndexHistory(operatorId2, defaultQuorumNumber, 1);
        require(toBlockNumber2 == 0, "toBlockNumber not set correctly");
        require(index2 == 0, "incorrect index");

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

    function testGettingOperatorIndexForQuorumAtBlockNumber(uint32 numOperators, uint32 operatorToDeregister) external {
        cheats.assume(numOperators > 0 && numOperators < 256);
        cheats.assume(operatorToDeregister >= 0 && operatorToDeregister < numOperators);

        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        for (uint256 i = 0; i < numOperators; i++) {
            bytes32 operatorId = _getRandomId(i);
            _registerOperator(operatorId, quorumNumbers);
        }
        cheats.roll(block.number + 100);

        bytes32 operator = _getRandomId(operatorToDeregister);

        uint32[] memory quorumToOperatorListIndexes = new uint32[](1);
        quorumToOperatorListIndexes[0] = uint32(_generateRandomNumber(operatorToDeregister, numOperators));
        _deregisterOperator(operator, quorumNumbers, quorumToOperatorListIndexes, operatorToDeregister);
        require(operatorToDeregister == indexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex(operator, defaultQuorumNumber, uint32(_generateRandomNumber(1, 100)), 0), "wrong index returned");
    }

    function _registerOperator(bytes32 operatorId, bytes memory quorumNumbers) public {
        cheats.startPrank(address(registryCoordinatorMock));
       indexRegistry.registerOperator(operatorId, quorumNumbers);
        cheats.stopPrank();
    }
    
    function _deregisterOperator(bytes32 operatorId, bytes memory quorumNumbers, uint32[] memory quorumToOperatorListIndexes, uint32 index) public {
        cheats.startPrank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId, quorumNumbers, quorumToOperatorListIndexes, index);
        cheats.stopPrank();
    }

    function _getRandomId(uint256 seed) internal view returns (bytes32) {
        return keccak256(abi.encodePacked(block.timestamp, seed));
    }

    function _generateRandomNumber(uint256 seed, uint256 modulus) internal view returns (uint256) {
        uint256 randomNumber = uint256(keccak256(abi.encodePacked(block.timestamp, seed)));
        return (randomNumber % modulus); 
    }
}
