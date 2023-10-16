//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../contracts/interfaces/IIndexRegistry.sol";
import "../../contracts/middleware/IndexRegistry.sol";
import "../mocks/RegistryCoordinatorMock.sol";
import "../harnesses/BitmapUtilsWrapper.sol";

import "forge-std/Test.sol";

contract IndexRegistryUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    IndexRegistry indexRegistry;
    RegistryCoordinatorMock registryCoordinatorMock;
    BitmapUtilsWrapper bitmapUtilsWrapper;

    uint8 defaultQuorumNumber = 1;
    bytes32 operatorId1 = bytes32(uint256(34));
    bytes32 operatorId2 = bytes32(uint256(35));
    bytes32 operatorId3 = bytes32(uint256(36));

    // Test 0 length operators in operators to remove
    function setUp() public {
        // deploy the contract
        registryCoordinatorMock = new RegistryCoordinatorMock();
        indexRegistry = new IndexRegistry(registryCoordinatorMock);
        bitmapUtilsWrapper = new BitmapUtilsWrapper();
    }

    function testConstructor() public {
        // check that the registry coordinator is set correctly
        assertEq(address(indexRegistry.registryCoordinator()), address(registryCoordinatorMock));
    }

    /*******************************************************************************
                            UNIT TESTS - REGISTRATION
    *******************************************************************************/

    /**
     * Preconditions for registration -> checks in BLSRegistryCoordinator
     * 1. quorumNumbers has no duplicates
     * 2. quorumNumbers ordered in ascending order
     * 3. quorumBitmap is <= uint192.max
     * 4. quorumNumbers.length != 0
     * 5. operator is not already registerd for any quorums being registered for
     */
    function testRegisterOperator() public {
        // register an operator
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        cheats.prank(address(registryCoordinatorMock));
        uint32[] memory numOperatorsPerQuorum = indexRegistry.registerOperator(operatorId1, quorumNumbers);

        // Check return value
        require(
            numOperatorsPerQuorum.length == 1,
            "IndexRegistry.registerOperator: numOperatorsPerQuorum length not 1"
        );
        require(numOperatorsPerQuorum[0] == 1, "IndexRegistry.registerOperator: numOperatorsPerQuorum[0] not 1");

        // Check globalOperatorList updates
        require(
            indexRegistry.globalOperatorList(0) == operatorId1,
            "IndexRegistry.registerOperator: operator not appened to globalOperatorList"
        );
        require(
            indexRegistry.getGlobalOperatorListLength() == 1,
            "IndexRegistry.registerOperator: globalOperatorList length incorrect"
        );

        // Check _operatorIdToIndexHistory updates
        IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
            .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId1, 1, 0);
        require(indexUpdate.index == 0, "IndexRegistry.registerOperator: index not 0");
        require(
            indexUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: fromBlockNumber not correct"
        );

        // Check _totalOperatorsHistory updates
        IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate = indexRegistry
            .getTotalOperatorsUpdateForQuorumAtIndex(1, 0);
        require(
            totalOperatorUpdate.index == 1,
            "IndexRegistry.registerOperator: totalOperatorsHistory index (num operators) not 1"
        );
        require(
            totalOperatorUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: totalOperatorsHistory fromBlockNumber not correct"
        );
        require(
            indexRegistry.totalOperatorsForQuorum(1) == 1,
            "IndexRegistry.registerOperator: total operators for quorum not updated correctly"
        );
    }

    function testRegisterOperatorMultipleQuorums() public {
        // Register operator for 1st quorum
        testRegisterOperator();

        // Register operator for 2nd quorum
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber + 1);

        cheats.prank(address(registryCoordinatorMock));
        uint32[] memory numOperatorsPerQuorum = indexRegistry.registerOperator(operatorId1, quorumNumbers);

        ///@notice The only value that should be different from before are what quorum we index into and the globalOperatorList
        // Check return value
        require(
            numOperatorsPerQuorum.length == 1,
            "IndexRegistry.registerOperator: numOperatorsPerQuorum length not 2"
        );
        require(numOperatorsPerQuorum[0] == 1, "IndexRegistry.registerOperator: numOperatorsPerQuorum[1] not 1");

        // Check globalOperatorList updates
        require(
            indexRegistry.globalOperatorList(1) == operatorId1,
            "IndexRegistry.registerOperator: operator not appened to globalOperatorList"
        );
        require(
            indexRegistry.getGlobalOperatorListLength() == 2,
            "IndexRegistry.registerOperator: globalOperatorList length incorrect"
        );

        // Check _operatorIdToIndexHistory updates
        IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
            .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId1, 2, 0);
        require(indexUpdate.index == 0, "IndexRegistry.registerOperator: index not 0");
        require(
            indexUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: fromBlockNumber not correct"
        );

        // Check _totalOperatorsHistory updates
        IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate = indexRegistry
            .getTotalOperatorsUpdateForQuorumAtIndex(2, 0);
        require(
            totalOperatorUpdate.index == 1,
            "IndexRegistry.registerOperator: totalOperatorsHistory index (num operators) not 1"
        );
        require(
            totalOperatorUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: totalOperatorsHistory fromBlockNumber not correct"
        );
        require(
            indexRegistry.totalOperatorsForQuorum(2) == 1,
            "IndexRegistry.registerOperator: total operators for quorum not updated correctly"
        );
    }

    function testRegisterOperatorMultipleQuorumsSingleCall() public {
        // Register operator for 1st and 2nd quorum
        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        quorumNumbers[1] = bytes1(defaultQuorumNumber + 1);

        cheats.prank(address(registryCoordinatorMock));
        uint32[] memory numOperatorsPerQuorum = indexRegistry.registerOperator(operatorId1, quorumNumbers);

        // Check return value
        require(
            numOperatorsPerQuorum.length == 2,
            "IndexRegistry.registerOperator: numOperatorsPerQuorum length not 2"
        );
        require(numOperatorsPerQuorum[0] == 1, "IndexRegistry.registerOperator: numOperatorsPerQuorum[0] not 1");
        require(numOperatorsPerQuorum[1] == 1, "IndexRegistry.registerOperator: numOperatorsPerQuorum[1] not 1");

        // Check globalOperatorList updates
        require(
            indexRegistry.globalOperatorList(0) == operatorId1,
            "IndexRegistry.registerOperator: operator not appened to globalOperatorList"
        );
        require(
            indexRegistry.getGlobalOperatorListLength() == 1,
            "IndexRegistry.registerOperator: globalOperatorList length incorrect"
        );

        // Check _operatorIdToIndexHistory updates for quorum 1
        IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
            .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId1, 1, 0);
        require(indexUpdate.index == 0, "IndexRegistry.registerOperator: index not 0");
        require(
            indexUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: fromBlockNumber not correct"
        );

        // Check _totalOperatorsHistory updates for quorum 1
        IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate = indexRegistry
            .getTotalOperatorsUpdateForQuorumAtIndex(1, 0);
        require(
            totalOperatorUpdate.index == 1,
            "IndexRegistry.registerOperator: totalOperatorsHistory index (num operators) not 1"
        );
        require(
            totalOperatorUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: totalOperatorsHistory fromBlockNumber not correct"
        );
        require(
            indexRegistry.totalOperatorsForQuorum(1) == 1,
            "IndexRegistry.registerOperator: total operators for quorum not updated correctly"
        );

        // Check _operatorIdToIndexHistory updates for quorum 2
        indexUpdate = indexRegistry.getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId1, 2, 0);
        require(indexUpdate.index == 0, "IndexRegistry.registerOperator: index not 0");
        require(
            indexUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: fromBlockNumber not correct"
        );

        // Check _totalOperatorsHistory updates for quorum 2
        totalOperatorUpdate = indexRegistry.getTotalOperatorsUpdateForQuorumAtIndex(2, 0);
        require(
            totalOperatorUpdate.index == 1,
            "IndexRegistry.registerOperator: totalOperatorsHistory index (num operators) not 1"
        );
        require(
            totalOperatorUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: totalOperatorsHistory fromBlockNumber not correct"
        );
        require(
            indexRegistry.totalOperatorsForQuorum(2) == 1,
            "IndexRegistry.registerOperator: total operators for quorum not updated correctly"
        );
    }

    function testRegisterMultipleOperatorsSingleQuorum() public {
        // Register operator for first quorum
        testRegisterOperator();

        // Register another operator
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        cheats.prank(address(registryCoordinatorMock));
        uint32[] memory numOperatorsPerQuorum = indexRegistry.registerOperator(operatorId2, quorumNumbers);

        // Check return value
        require(
            numOperatorsPerQuorum.length == 1,
            "IndexRegistry.registerOperator: numOperatorsPerQuorum length not 1"
        );
        require(numOperatorsPerQuorum[0] == 2, "IndexRegistry.registerOperator: numOperatorsPerQuorum[0] not 2");

        // Check globalOperatorList updates
        require(
            indexRegistry.globalOperatorList(1) == operatorId2,
            "IndexRegistry.registerOperator: operator not appened to globalOperatorList"
        );
        require(
            indexRegistry.getGlobalOperatorListLength() == 2,
            "IndexRegistry.registerOperator: globalOperatorList length incorrect"
        );

        // Check _operatorIdToIndexHistory updates
        IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
            .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId2, 1, 0);
        require(indexUpdate.index == 1, "IndexRegistry.registerOperator: index not 1");
        require(
            indexUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: fromBlockNumber not correct"
        );

        // Check _totalOperatorsHistory updates
        IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate = indexRegistry
            .getTotalOperatorsUpdateForQuorumAtIndex(1, 1);
        require(
            totalOperatorUpdate.index == 2,
            "IndexRegistry.registerOperator: totalOperatorsHistory index (num operators) not 2"
        );
        require(
            totalOperatorUpdate.fromBlockNumber == block.number,
            "IndexRegistry.registerOperator: totalOperatorsHistory fromBlockNumber not correct"
        );
        require(
            indexRegistry.totalOperatorsForQuorum(1) == 2,
            "IndexRegistry.registerOperator: total operators for quorum not updated correctly"
        );
    }

    /*******************************************************************************
                            UNIT TESTS - DEREGISTRATION
    *******************************************************************************/

    function testDeregisterOperatorRevertMismatchInputLengths() public {
        bytes memory quorumNumbers = new bytes(1);
        bytes32[] memory operatorIdsToSwap = new bytes32[](2);

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        cheats.prank(address(registryCoordinatorMock));
        cheats.expectRevert(
            bytes("IndexRegistry.deregisterOperator: quorumNumbers and operatorIdsToSwap must be the same length")
        );
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, operatorIdsToSwap);
    }

    function testDeregisterOperatorSingleOperator() public {
        // Register operator
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);

        // Deregister operator
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = operatorId1;
        cheats.prank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, operatorIdsToSwap);

        // Check operator's index
        IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
            .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId1, defaultQuorumNumber, 1);
        require(indexUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
        require(indexUpdate.index == type(uint32).max, "incorrect index");

        // Check total operators
        IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate = indexRegistry
            .getTotalOperatorsUpdateForQuorumAtIndex(defaultQuorumNumber, 1);
        require(totalOperatorUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
        require(totalOperatorUpdate.index == 0, "incorrect total number of operators");
        require(indexRegistry.totalOperatorsForQuorum(1) == 0, "operator not deregistered correctly");
    }

    function testDeregisterOperatorRevertIncorrectOperatorToSwap() public {
        // Register 3 operators
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);
        _registerOperator(operatorId2, quorumNumbers);
        _registerOperator(operatorId3, quorumNumbers);

        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = operatorId2;

        //deregister the operatorId1, removing it from both quorum 1 and 2.
        cheats.prank(address(registryCoordinatorMock));
        cheats.expectRevert(
            bytes("IndexRegistry._processOperatorRemoval: operatorIdToSwap is not the last operator in the quorum")
        );
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, operatorIdsToSwap);
    }

    function testDeregisterOperatorMultipleQuorums() public {
        // Register 3 operators to two quorums
        bytes memory quorumNumbers = new bytes(3);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        quorumNumbers[1] = bytes1(defaultQuorumNumber + 1);
        quorumNumbers[2] = bytes1(defaultQuorumNumber + 2);
        _registerOperator(operatorId1, quorumNumbers);
        _registerOperator(operatorId2, quorumNumbers);
        _registerOperator(operatorId3, quorumNumbers);

        // Deregister operator from quorums 1 and 2
        bytes memory quorumsToRemove = new bytes(2);
        quorumsToRemove[0] = bytes1(defaultQuorumNumber);
        quorumsToRemove[1] = bytes1(defaultQuorumNumber + 1);
        bytes32[] memory operatorIdsToSwap = new bytes32[](2);
        operatorIdsToSwap[0] = operatorId3;
        operatorIdsToSwap[1] = operatorId3;

        cheats.prank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumsToRemove, operatorIdsToSwap);

        // Check operator's index for removed quorums
        for (uint256 i = 0; i < quorumsToRemove.length; i++) {
            IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
                .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId1, uint8(quorumsToRemove[i]), 1); // 2 indexes -> 1 update and 1 remove
            require(indexUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
            require(indexUpdate.index == type(uint32).max, "incorrect index");
        }

        // Check total operators for removed quorums
        for (uint256 i = 0; i < quorumsToRemove.length; i++) {
            IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate = indexRegistry
                .getTotalOperatorsUpdateForQuorumAtIndex(uint8(quorumsToRemove[i]), 3); // 4 updates total
            require(totalOperatorUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
            require(totalOperatorUpdate.index == 2, "incorrect total number of operators");
            require(
                indexRegistry.totalOperatorsForQuorum(uint8(quorumsToRemove[i])) == 2,
                "operator not deregistered correctly"
            );
        }

        // Check swapped operator's index for removed quorums
        for (uint256 i = 0; i < quorumsToRemove.length; i++) {
            IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
                .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId3, uint8(quorumsToRemove[i]), 1); // 2 indexes -> 1 update and 1 swap
            require(indexUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
            require(indexUpdate.index == 0, "incorrect index");
        }
    }

    /*******************************************************************************
                                UNIT TESTS - GETTERS
    *******************************************************************************/

    function testGetOperatorIndexForQuorumAtBlockNumberByIndex_revert_earlyBlockNumber() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);
        cheats.expectRevert(
            bytes(
                "IndexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex: provided index is too far in the past for provided block number"
            )
        );
        indexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex(
            operatorId1,
            defaultQuorumNumber,
            uint32(block.number - 1),
            0
        );
    }

    function testGetOperatorIndexForQuorumAtBlockNumberByIndex_revert_indexBlockMismatch() public {
        // Register operator for first quorum
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);

        // Deregister operator from first quorum
        vm.roll(block.number + 10);
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = operatorId1;
        cheats.prank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, operatorIdsToSwap);

        // Fails due to block number corresponding to a later index (from the deregistration)
        cheats.expectRevert(
            bytes(
                "IndexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex: provided index is too far in the future for provided block number"
            )
        );
        indexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex(
            operatorId1,
            defaultQuorumNumber,
            uint32(block.number),
            0
        );
    }

    function testGetOperatorIndexForQuorumAtBlockNumberByIndex() public {
        // Register operator for first quorum
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);

        // Deregister operator from first quorum
        vm.roll(block.number + 10);
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = operatorId1;
        cheats.prank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, operatorIdsToSwap);

        // Check that the first index is correct
        uint32 firstIndex = indexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex(
            operatorId1,
            defaultQuorumNumber,
            uint32(block.number - 10),
            0
        );
        require(firstIndex == 0, "IndexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex: first index not 0");

        // Check that the index is correct
        uint32 index = indexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex(
            operatorId1,
            defaultQuorumNumber,
            uint32(block.number),
            1
        );
        require(
            index == type(uint32).max,
            "IndexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex: second index not deregistered"
        );
    }

    function testGetTotalOperatorsForQuorumAtBlockNumberByIndex_revert_indexTooEarly() public {
        // Add operator
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);

        cheats.expectRevert(
            "IndexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex: provided index is too far in the past for provided block number"
        );
        indexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex(defaultQuorumNumber, uint32(block.number - 1), 0);
    }

    function testGetTotalOperatorsForQuorumAtBlockNumberByIndex_revert_indexBlockMismatch() public {
        // Add two operators
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);
        vm.roll(block.number + 10);
        _registerOperator(operatorId2, quorumNumbers);

        cheats.expectRevert(
            "IndexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex: provided index is too far in the future for provided block number"
        );
        indexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex(defaultQuorumNumber, uint32(block.number), 0);
    }

    function testGetTotalOperatorsForQuorumAtBlockNumberByIndex() public {
        // Add two operators
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);
        vm.roll(block.number + 10);
        _registerOperator(operatorId2, quorumNumbers);

        // Check that the first total is correct
        uint32 prevTotal = indexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex(
            defaultQuorumNumber,
            uint32(block.number - 10),
            0
        );
        require(prevTotal == 1, "IndexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex: prev total not 1");

        // Check that the total is correct
        uint32 currentTotal = indexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex(
            defaultQuorumNumber,
            uint32(block.number),
            1
        );
        require(currentTotal == 2, "IndexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex: current total not 2");
    }

    function testGetOperatorListForQuorumAtBlockNumber() public {
        // Register two operators
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        _registerOperator(operatorId1, quorumNumbers);
        vm.roll(block.number + 10);
        _registerOperator(operatorId2, quorumNumbers);

        // Deregister first operator
        vm.roll(block.number + 10);
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = operatorId2;
        cheats.prank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumNumbers, operatorIdsToSwap);

        // Check the operator list after first registration
        bytes32[] memory operatorList = indexRegistry.getOperatorListForQuorumAtBlockNumber(
            defaultQuorumNumber,
            uint32(block.number - 20)
        );
        require(
            operatorList.length == 1,
            "IndexRegistry.getOperatorListForQuorumAtBlockNumber: operator list length not 1"
        );
        require(
            operatorList[0] == operatorId1,
            "IndexRegistry.getOperatorListForQuorumAtBlockNumber: operator list incorrect"
        );

        // Check the operator list after second registration
        operatorList = indexRegistry.getOperatorListForQuorumAtBlockNumber(
            defaultQuorumNumber,
            uint32(block.number - 10)
        );
        require(
            operatorList.length == 2,
            "IndexRegistry.getOperatorListForQuorumAtBlockNumber: operator list length not 2"
        );
        require(
            operatorList[0] == operatorId1,
            "IndexRegistry.getOperatorListForQuorumAtBlockNumber: operator list incorrect"
        );
        require(
            operatorList[1] == operatorId2,
            "IndexRegistry.getOperatorListForQuorumAtBlockNumber: operator list incorrect"
        );

        // Check the operator list after deregistration
        operatorList = indexRegistry.getOperatorListForQuorumAtBlockNumber(defaultQuorumNumber, uint32(block.number));
        require(
            operatorList.length == 1,
            "IndexRegistry.getOperatorListForQuorumAtBlockNumber: operator list length not 1"
        );
        require(
            operatorList[0] == operatorId2,
            "IndexRegistry.getOperatorListForQuorumAtBlockNumber: operator list incorrect"
        );
    }

    /*******************************************************************************
                                    FUZZ TESTS
    *******************************************************************************/

    function testFuzzRegisterOperatorRevertFromNonRegisterCoordinator(address nonRegistryCoordinator) public {
        cheats.assume(address(registryCoordinatorMock) != nonRegistryCoordinator);
        bytes memory quorumNumbers = new bytes(defaultQuorumNumber);

        cheats.prank(nonRegistryCoordinator);
        cheats.expectRevert(bytes("IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        indexRegistry.registerOperator(bytes32(0), quorumNumbers);
    }

    function testFuzzTotalOperatorUpdatesForOneQuorum(uint8 numOperators) public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        uint256 lengthBefore = 0;
        for (uint256 i = 0; i < numOperators; i++) {
            _registerOperator(bytes32(i), quorumNumbers);
            require(indexRegistry.totalOperatorsForQuorum(1) - lengthBefore == 1, "incorrect update");
            lengthBefore++;
        }
    }

    /**
     * Preconditions for registration -> checks in BLSRegistryCoordinator
     * 1. quorumNumbers has no duplicates
     * 2. quorumNumbers ordered in ascending order
     * 3. quorumBitmap is <= uint192.max
     * 4. quorumNumbers.length != 0
     * 5. operator is not already registerd for any quorums being registered for
     */
    function testFuzzRegisterOperatorMultipleQuorums(bytes memory quorumNumbers) public {
        // Validate quorumNumbers
        cheats.assume(quorumNumbers.length > 0);
        cheats.assume(bitmapUtilsWrapper.isArrayStrictlyAscendingOrdered(quorumNumbers));
        uint256 bitmap = bitmapUtilsWrapper.orderedBytesArrayToBitmap(quorumNumbers);
        cheats.assume(bitmap <= type(uint192).max);

        // Register operator
        cheats.prank(address(registryCoordinatorMock));
        uint32[] memory numOperatorsPerQuorum = indexRegistry.registerOperator(operatorId1, quorumNumbers);

        // Check return value
        require(
            numOperatorsPerQuorum.length == quorumNumbers.length,
            "IndexRegistry.registerOperator: numOperatorsPerQuorum length not correct"
        );
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            require(numOperatorsPerQuorum[i] == 1, "IndexRegistry.registerOperator: numOperatorsPerQuorum not 1");
        }

        // Check globalOperatorList updates
        require(
            indexRegistry.globalOperatorList(0) == operatorId1,
            "IndexRegistry.registerOperator: operator not appened to globalOperatorList"
        );
        require(
            indexRegistry.getGlobalOperatorListLength() == 1,
            "IndexRegistry.registerOperator: globalOperatorList length incorrect"
        );

        // Check _operatorIdToIndexHistory updates
        IIndexRegistry.OperatorIndexUpdate memory indexUpdate;
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            indexUpdate = indexRegistry.getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(
                operatorId1,
                uint8(quorumNumbers[i]),
                0
            );
            require(indexUpdate.index == 0, "IndexRegistry.registerOperator: index not 0");
            require(
                indexUpdate.fromBlockNumber == block.number,
                "IndexRegistry.registerOperator: fromBlockNumber not correct"
            );
        }

        // Check _totalOperatorsHistory updates
        IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate;
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            totalOperatorUpdate = indexRegistry.getTotalOperatorsUpdateForQuorumAtIndex(uint8(quorumNumbers[i]), 0);
            require(
                totalOperatorUpdate.index == 1,
                "IndexRegistry.registerOperator: totalOperatorsHistory index (num operators) not 1"
            );
            require(
                totalOperatorUpdate.fromBlockNumber == block.number,
                "IndexRegistry.registerOperator: totalOperatorsHistory fromBlockNumber not correct"
            );
            require(
                indexRegistry.totalOperatorsForQuorum(uint8(quorumNumbers[i])) == 1,
                "IndexRegistry.registerOperator: total operators for quorum not updated correctly"
            );
        }
    }

    function testFuzzRegsiterMultipleOperatorsMultipleQuorums(bytes memory quorumNumbers) public {
        // Validate quorumNumbers
        cheats.assume(quorumNumbers.length > 0);
        cheats.assume(bitmapUtilsWrapper.isArrayStrictlyAscendingOrdered(quorumNumbers));
        uint256 bitmap = bitmapUtilsWrapper.orderedBytesArrayToBitmap(quorumNumbers);
        cheats.assume(bitmap <= type(uint192).max);

        // Register operators 1,2,3
        _registerOperator(operatorId1, quorumNumbers);
        vm.roll(block.number + 10);
        _registerOperator(operatorId2, quorumNumbers);
        vm.roll(block.number + 10);
        _registerOperator(operatorId3, quorumNumbers);

        // Check globalOperatorList updates
        require(
            indexRegistry.getGlobalOperatorListLength() == 3,
            "IndexRegistry.registerOperator: globalOperatorList length incorrect"
        );

        // Check history of _totalOperatorsHistory updates at each blockNumber
        IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate;
        uint256 numOperators = 1;
        for (uint256 blockNumber = block.number - 20; blockNumber <= block.number; blockNumber += 10) {
            for (uint256 i = 0; i < quorumNumbers.length; i++) {
                totalOperatorUpdate = indexRegistry.getTotalOperatorsUpdateForQuorumAtIndex(
                    uint8(quorumNumbers[i]),
                    uint32(numOperators - 1)
                );
                require(
                    totalOperatorUpdate.index == numOperators,
                    "IndexRegistry.registerOperator: totalOperatorsHistory index (num operators) not correct"
                );
                require(
                    totalOperatorUpdate.fromBlockNumber == blockNumber,
                    "IndexRegistry.registerOperator: totalOperatorsHistory fromBlockNumber not correct"
                );
            }
            numOperators++;
        }
    }

    function testFuzzDeregisterOperatorRevertFromNonRegisterCoordinator(address nonRegistryCoordinator) public {
        cheats.assume(address(registryCoordinatorMock) != nonRegistryCoordinator);
        // de-register an operator
        bytes memory quorumNumbers = new bytes(defaultQuorumNumber);
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);

        cheats.prank(nonRegistryCoordinator);
        cheats.expectRevert(bytes("IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        indexRegistry.deregisterOperator(bytes32(0), quorumNumbers, operatorIdsToSwap);
    }

    function testFuzzDeregisterOperator(bytes memory quorumsToAdd, uint256 bitsToFlip) public {
        // Validate quorumsToAdd
        cheats.assume(quorumsToAdd.length > 0);
        cheats.assume(bitmapUtilsWrapper.isArrayStrictlyAscendingOrdered(quorumsToAdd));
        uint256 bitmap = bitmapUtilsWrapper.orderedBytesArrayToBitmap(quorumsToAdd);
        cheats.assume(bitmap <= type(uint192).max);

        // Format quorumsToRemove
        bitsToFlip = bound(bitsToFlip, 1, quorumsToAdd.length);
        uint256 bitsFlipped = 0;
        uint256 bitPosition = 0;
        uint256 bitmapQuorumsToRemove = bitmap;
        while (bitsFlipped < bitsToFlip && bitPosition < 192) {
            uint256 bitMask = 1 << bitPosition;
            if (bitmapQuorumsToRemove & bitMask != 0) {
                bitmapQuorumsToRemove ^= bitMask;
                bitsFlipped++;
            }
            bitPosition++;
        }
        bytes memory quorumsToRemove = bitmapUtilsWrapper.bitmapToBytesArray(bitmapQuorumsToRemove);
        // Sanity check quorumsToRemove
        cheats.assume(bitmapUtilsWrapper.isArrayStrictlyAscendingOrdered(quorumsToRemove));

        // Register operators
        _registerOperator(operatorId1, quorumsToAdd);
        _registerOperator(operatorId2, quorumsToAdd);

        // Deregister operator
        bytes32[] memory operatorIdsToSwap = new bytes32[](quorumsToRemove.length);
        for (uint256 i = 0; i < quorumsToRemove.length; i++) {
            operatorIdsToSwap[i] = operatorId2;
        }
        cheats.prank(address(registryCoordinatorMock));
        indexRegistry.deregisterOperator(operatorId1, quorumsToRemove, operatorIdsToSwap);

        // Check operator's index for removed quorums
        for (uint256 i = 0; i < quorumsToRemove.length; i++) {
            IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
                .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId1, uint8(quorumsToRemove[i]), 1); // 2 indexes -> 1 update and 1 remove
            require(indexUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
            require(indexUpdate.index == type(uint32).max, "incorrect index");
        }

        // Check total operators for removed quorums
        for (uint256 i = 0; i < quorumsToRemove.length; i++) {
            IIndexRegistry.OperatorIndexUpdate memory totalOperatorUpdate = indexRegistry
                .getTotalOperatorsUpdateForQuorumAtIndex(uint8(quorumsToRemove[i]), 2); // 3 updates total
            require(totalOperatorUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
            require(totalOperatorUpdate.index == 1, "incorrect total number of operators");
            require(
                indexRegistry.totalOperatorsForQuorum(uint8(quorumsToRemove[i])) == 1,
                "operator not deregistered correctly"
            );
        }

        // Check swapped operator's index for removed quorums
        for (uint256 i = 0; i < quorumsToRemove.length; i++) {
            IIndexRegistry.OperatorIndexUpdate memory indexUpdate = indexRegistry
                .getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId2, uint8(quorumsToRemove[i]), 1); // 2 indexes -> 1 update and 1 swap
            require(indexUpdate.fromBlockNumber == block.number, "fromBlockNumber not set correctly");
            require(indexUpdate.index == 0, "incorrect index");
        }
    }

    function _registerOperator(bytes32 operatorId, bytes memory quorumNumbers) internal {
        cheats.prank(address(registryCoordinatorMock));
        indexRegistry.registerOperator(operatorId, quorumNumbers);
    }
}
