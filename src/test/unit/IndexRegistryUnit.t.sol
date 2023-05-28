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
        uint8[] memory quorumNumbers = new uint8[](1);
        quorumNumbers[0] = 1;

        cheats.startPrank(address(registryCoordinatorMock));
       indexRegistry.registerOperator(operatorId, quorumNumbers);
        cheats.stopPrank();

       require(indexRegistry.globalOperatorList(0) == operatorId, "IndexRegistry.registerOperator: operator not registered correctly");
       require(indexRegistry.totalOperators() == 1, "IndexRegistry.registerOperator: operator not registered correctly");
       require(indexRegistry.totalOperatorsForQuorum(1) == 1, "IndexRegistry.registerOperator: operator not registered correctly");
       (uint32 index, uint32 toBlockNumber) = indexRegistry.operatorIdToIndexHistory(operatorId, 1, 0);
       require(index == 0, "IndexRegistry.registerOperator: operator not registered correctly");
    }

    function testRegisterOperatorFromNonRegisterCoordinator(address nonRegistryCoordinator) public {
        cheats.assume(address(registryCoordinatorMock) != nonRegistryCoordinator);
        // register an operator
        uint8[] memory quorumNumbers = new uint8[](1);

        cheats.startPrank(nonRegistryCoordinator);
        cheats.expectRevert(bytes("IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        indexRegistry.registerOperator(bytes32(0), quorumNumbers);
        cheats.stopPrank();
    }

    function testDeregisterOperatorInIndexRegistry(bytes32 operatorId1, bytes32 operatorId2) public {
        uint8[] memory quorumNumbers = new uint8[](2);
        quorumNumbers[0] = 1;
        quorumNumbers[1] = 2;
        _registerOperator(operatorId1, quorumNumbers);
        _registerOperator(operatorId2, quorumNumbers);

        require(indexRegistry.totalOperators() == 2, "IndexRegistry.registerOperator: operator not registered correctly");
        require(indexRegistry.totalOperatorsForQuorum(1) == 2, "IndexRegistry.registerOperator: operator not registered correctly");
        require(indexRegistry.totalOperatorsForQuorum(2) == 2, "IndexRegistry.registerOperator: operator not registered correctly");

        uint32[] memory quorumToOperatorListIndexes = new uint32[](2);
        quorumToOperatorListIndexes[0] = 0;
        quorumToOperatorListIndexes[1] = 0;

        // deregister the operatorId1, removing it from both quorum 1 and 2.
        cheats.startPrank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, quorumToOperatorListIndexes, 0);
        cheats.stopPrank();

        require(indexRegistry.totalOperators() == 1, "IndexRegistry.registerOperator: operator not registered correctly");
        require(indexRegistry.quorumToOperatorList(1, 0) == operatorId2, "IndexRegistry.registerOperator: operator not deregistered and swapped correctly");
        require(indexRegistry.quorumToOperatorList(2, 0) == operatorId2, "IndexRegistry.registerOperator: operator not deregistered and swapped correctly");
        require(indexRegistry.globalOperatorList(0) == operatorId2, "IndexRegistry.registerOperator: operator not deregistered and swapped correctly");
    }

    function _registerOperator(bytes32 operatorId, uint8[] memory quorumNumbers) public {
        cheats.startPrank(address(registryCoordinatorMock));
       indexRegistry.registerOperator(operatorId, quorumNumbers);
        cheats.stopPrank();
    }
}
