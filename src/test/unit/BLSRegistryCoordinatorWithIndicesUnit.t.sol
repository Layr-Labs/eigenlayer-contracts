// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../utils/MockAVSDeployer.sol";

contract BLSRegistryCoordinatorWithIndicesUnit is MockAVSDeployer {
    using BN254 for BN254.G1Point;

    event OperatorSocketUpdate(bytes32 indexed operatorId, string socket);

    /// @notice emitted whenever the stake of `operator` is updated
    event StakeUpdate(
        bytes32 indexed operatorId,
        uint8 quorumNumber,
        uint96 stake
    );

    // Emitted when a new operator pubkey is registered for a set of quorums
    event OperatorAddedToQuorums(
        address operator,
        bytes quorumNumbers
    );

    // Emitted when an operator pubkey is removed from a set of quorums
    event OperatorRemovedFromQuorums(
        address operator, 
        bytes quorumNumbers
    );

    // emitted when an operator's index in the orderd operator list for the quorum with number `quorumNumber` is updated
    event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newIndex);

    event OperatorSetParamsUpdated(uint8 indexed quorumNumber, IBLSRegistryCoordinatorWithIndices.OperatorSetParam operatorSetParams);

    event ChurnApproverUpdated(address churnApprover);

    function setUp() virtual public {
        _deployMockEigenLayerAndAVS();
    }

    function testCorrectConstruction() public {
        assertEq(address(registryCoordinator.stakeRegistry()), address(stakeRegistry));
        assertEq(address(registryCoordinator.blsPubkeyRegistry()), address(blsPubkeyRegistry));
        assertEq(address(registryCoordinator.indexRegistry()), address(indexRegistry));
        assertEq(address(registryCoordinator.slasher()), address(slasher));

        for (uint i = 0; i < numQuorums; i++) {
            assertEq(
                keccak256(abi.encode(registryCoordinator.getOperatorSetParams(uint8(i)))), 
                keccak256(abi.encode(operatorSetParams[i]))
            );
        }

        // make sure the contract intializers are disabled
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        registryCoordinator.initialize(churnApprover, operatorSetParams);
    }

    function testSetOperatorSetParams_NotServiceManagerOwner_Reverts() public {
        cheats.expectRevert("BLSRegistryCoordinatorWithIndices.onlyServiceManagerOwner: caller is not the service manager owner");
        cheats.prank(defaultOperator);
        registryCoordinator.setOperatorSetParams(0, operatorSetParams[0]);
    }

    function testSetOperatorSetParams_Valid() public {
        cheats.prank(serviceManagerOwner);
        cheats.expectEmit(true, true, true, true, address(registryCoordinator));
        emit OperatorSetParamsUpdated(0, operatorSetParams[1]);
        registryCoordinator.setOperatorSetParams(0, operatorSetParams[1]);
    }

    function testSetChurnApprover_NotServiceManagerOwner_Reverts() public {
        address newChurnApprover = address(uint160(uint256(keccak256("newChurnApprover"))));
        cheats.expectRevert("BLSRegistryCoordinatorWithIndices.onlyServiceManagerOwner: caller is not the service manager owner");
        cheats.prank(defaultOperator);
        registryCoordinator.setChurnApprover(newChurnApprover);
    }

    function testSetChurnApprover_Valid() public {
        address newChurnApprover = address(uint160(uint256(keccak256("newChurnApprover"))));
        cheats.prank(serviceManagerOwner);
        cheats.expectEmit(true, true, true, true, address(registryCoordinator));
        emit ChurnApproverUpdated(newChurnApprover);
        registryCoordinator.setChurnApprover(newChurnApprover);
    }

    function testRegisterOperatorWithCoordinator_EmptyQuorumNumbers_Reverts() public {
        bytes memory emptyQuorumNumbers = new bytes(0);
        cheats.expectRevert("BLSRegistryCoordinatorWithIndices._registerOperatorWithCoordinator: quorumBitmap cannot be 0");
        cheats.prank(defaultOperator);
        registryCoordinator.registerOperatorWithCoordinator(emptyQuorumNumbers, defaultPubKey, defaultSocket);
    }

    function testRegisterOperatorWithCoordinator_QuorumNumbersTooLarge_Reverts() public {
        bytes memory quorumNumbersTooLarge = new bytes(1);
        quorumNumbersTooLarge[0] = 0xC0;
        cheats.expectRevert("BLSRegistryCoordinatorWithIndices._registerOperatorWithCoordinator: quorumBitmap exceeds of max bitmap size");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbersTooLarge, defaultPubKey, defaultSocket);
    }

    function testRegisterOperatorWithCoordinator_QuorumNotCreated_Reverts() public {
        _deployMockEigenLayerAndAVS(10);
        bytes memory quorumNumbersNotCreated = new bytes(1);
        quorumNumbersNotCreated[0] = 0x0B;
        cheats.prank(defaultOperator);
        cheats.expectRevert("StakeRegistry._registerOperator: greatest quorumNumber must be less than quorumCount");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbersNotCreated, defaultPubKey, defaultSocket);
    }

    function testRegisterOperatorWithCoordinatorForSingleQuorum_Valid() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        stakeRegistry.setOperatorWeight(uint8(quorumNumbers[0]), defaultOperator, defaultStake);

        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit OperatorAddedToQuorums(defaultOperator, quorumNumbers);
        cheats.expectEmit(true, true, true, true, address(stakeRegistry));
        emit StakeUpdate(defaultOperatorId, defaultQuorumNumber, defaultStake);
        cheats.expectEmit(true, true, true, true, address(indexRegistry));
        emit QuorumIndexUpdate(defaultOperatorId, defaultQuorumNumber, 0);
        cheats.expectEmit(true, true, true, true, address(registryCoordinator));
        emit OperatorSocketUpdate(defaultOperatorId, defaultSocket);

        uint256 gasBefore = gasleft();
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, defaultPubKey, defaultSocket);
        uint256 gasAfter = gasleft();
        emit log_named_uint("gasUsed", gasBefore - gasAfter);

        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        assertEq(registryCoordinator.getOperatorId(defaultOperator), defaultOperatorId);
        assertEq(
            keccak256(abi.encode(registryCoordinator.getOperator(defaultOperator))), 
            keccak256(abi.encode(IRegistryCoordinator.Operator({
                operatorId: defaultOperatorId,
                status: IRegistryCoordinator.OperatorStatus.REGISTERED
            })))
        );
        assertEq(registryCoordinator.getCurrentQuorumBitmapByOperatorId(defaultOperatorId), quorumBitmap);
        assertEq(
            keccak256(abi.encode(registryCoordinator.getQuorumBitmapUpdateByOperatorIdByIndex(defaultOperatorId, 0))), 
            keccak256(abi.encode(IRegistryCoordinator.QuorumBitmapUpdate({
                quorumBitmap: uint192(quorumBitmap),
                updateBlockNumber: uint32(block.number),
                nextUpdateBlockNumber: 0
            })))
        );
    }

    function testRegisterOperatorWithCoordinatorForFuzzedQuorums_Valid(uint256 quorumBitmap) public {
        quorumBitmap = quorumBitmap & MiddlewareUtils.MAX_QUORUM_BITMAP;
        cheats.assume(quorumBitmap != 0);
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        for (uint i = 0; i < quorumNumbers.length; i++) {
            stakeRegistry.setOperatorWeight(uint8(quorumNumbers[i]), defaultOperator, defaultStake);
        }

        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit OperatorAddedToQuorums(defaultOperator, quorumNumbers);

        for (uint i = 0; i < quorumNumbers.length; i++) {
            cheats.expectEmit(true, true, true, true, address(stakeRegistry));
            emit StakeUpdate(defaultOperatorId, uint8(quorumNumbers[i]), defaultStake);
        }    

        for (uint i = 0; i < quorumNumbers.length; i++) {
            cheats.expectEmit(true, true, true, true, address(indexRegistry));
            emit QuorumIndexUpdate(defaultOperatorId, uint8(quorumNumbers[i]), 0);
        }    
        cheats.expectEmit(true, true, true, true, address(registryCoordinator));
        emit OperatorSocketUpdate(defaultOperatorId, defaultSocket);

        uint256 gasBefore = gasleft();
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, defaultPubKey, defaultSocket);
        uint256 gasAfter = gasleft();
        emit log_named_uint("gasUsed", gasBefore - gasAfter);
        emit log_named_uint("numQuorums", quorumNumbers.length);

        assertEq(registryCoordinator.getOperatorId(defaultOperator), defaultOperatorId);
        assertEq(
            keccak256(abi.encode(registryCoordinator.getOperator(defaultOperator))), 
            keccak256(abi.encode(IRegistryCoordinator.Operator({
                operatorId: defaultOperatorId,
                status: IRegistryCoordinator.OperatorStatus.REGISTERED
            })))
        );
        assertEq(registryCoordinator.getCurrentQuorumBitmapByOperatorId(defaultOperatorId), quorumBitmap);
        assertEq(
            keccak256(abi.encode(registryCoordinator.getQuorumBitmapUpdateByOperatorIdByIndex(defaultOperatorId, 0))), 
            keccak256(abi.encode(IRegistryCoordinator.QuorumBitmapUpdate({
                quorumBitmap: uint192(quorumBitmap),
                updateBlockNumber: uint32(block.number),
                nextUpdateBlockNumber: 0
            })))
        );
    }

    function testRegisterOperatorWithCoordinator_OverFilledQuorum_Reverts(uint256 pseudoRandomNumber) public {
        uint32 numOperators = defaultMaxOperatorCount;
        uint32 registrationBlockNumber = 200;

        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        cheats.roll(registrationBlockNumber);

        for (uint i = 0; i < numOperators; i++) {
            BN254.G1Point memory pubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, i)));
            address operator = _incrementAddress(defaultOperator, i);
            
            _registerOperatorWithCoordinator(operator, quorumBitmap, pubKey);
        }

        address operatorToRegister = _incrementAddress(defaultOperator, numOperators);
        BN254.G1Point memory operatorToRegisterPubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, numOperators)));
    
        pubkeyCompendium.setBLSPublicKey(operatorToRegister, operatorToRegisterPubKey);

        stakeRegistry.setOperatorWeight(defaultQuorumNumber, operatorToRegister, defaultStake);

        cheats.prank(operatorToRegister);
        cheats.expectRevert("BLSRegistryCoordinatorWithIndices._registerOperatorWithCoordinatorAndNoOverfilledQuorums: quorum is overfilled");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, operatorToRegisterPubKey, defaultSocket);
    }

    function testDeregisterOperatorWithCoordinator_NotRegistered_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        cheats.expectRevert("BLSRegistryCoordinatorWithIndices._deregisterOperatorWithCoordinator: operator is not registered");
        cheats.prank(defaultOperator);
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbers, defaultPubKey, new bytes32[](0));
    }

    function testDeregisterOperatorWithCoordinator_IncorrectPubkey_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        _registerOperatorWithCoordinator(defaultOperator, quorumBitmap, defaultPubKey);

        BN254.G1Point memory incorrectPubKey = BN254.hashToG1(bytes32(uint256(123)));

        cheats.expectRevert("BLSRegistryCoordinatorWithIndices._deregisterOperatorWithCoordinator: operatorId does not match pubkey hash");
        cheats.prank(defaultOperator);
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbers, incorrectPubKey, new bytes32[](0));
    }

    function testDeregisterOperatorWithCoordinator_InvalidQuorums_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        _registerOperatorWithCoordinator(defaultOperator, quorumBitmap, defaultPubKey);

        quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        quorumNumbers[1] = bytes1(uint8(192));

        cheats.expectRevert("BLSRegistryCoordinatorWithIndices._deregisterOperatorWithCoordinator: quorumsToRemoveBitmap exceeds of max bitmap size");
        cheats.prank(defaultOperator);
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbers, defaultPubKey, new bytes32[](0));
    }

    function testDeregisterOperatorWithCoordinator_IncorrectQuorums_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        _registerOperatorWithCoordinator(defaultOperator, quorumBitmap, defaultPubKey);

        quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        quorumNumbers[1] = bytes1(defaultQuorumNumber + 1);

        cheats.expectRevert("BLSRegistryCoordinatorWithIndices._deregisterOperatorWithCoordinator: cannot deregister operator for quorums that it is not a part of");
        cheats.prank(defaultOperator);
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbers, defaultPubKey, new bytes32[](0));
    }

    function testDeregisterOperatorWithCoordinatorForSingleQuorumAndSingleOperator_Valid() public {
        uint32 registrationBlockNumber = 100;
        uint32 deregistrationBlockNumber = 200;

        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        stakeRegistry.setOperatorWeight(uint8(quorumNumbers[0]), defaultOperator, defaultStake);

        cheats.startPrank(defaultOperator);
        
        cheats.roll(registrationBlockNumber);
        
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, defaultPubKey, defaultSocket);

        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        operatorIdsToSwap[0] = defaultOperatorId;

        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit OperatorAddedToQuorums(defaultOperator, quorumNumbers);
        cheats.expectEmit(true, true, true, true, address(stakeRegistry));
        emit StakeUpdate(defaultOperatorId, defaultQuorumNumber, 0);

        cheats.roll(deregistrationBlockNumber);

        uint256 gasBefore = gasleft();
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbers, defaultPubKey, operatorIdsToSwap);
        uint256 gasAfter = gasleft();
        emit log_named_uint("gasUsed", gasBefore - gasAfter);

        assertEq(
            keccak256(abi.encode(registryCoordinator.getOperator(defaultOperator))), 
            keccak256(abi.encode(IRegistryCoordinator.Operator({
                operatorId: defaultOperatorId,
                status: IRegistryCoordinator.OperatorStatus.DEREGISTERED
            })))
        );
        assertEq(registryCoordinator.getCurrentQuorumBitmapByOperatorId(defaultOperatorId), 0);
        assertEq(
            keccak256(abi.encode(registryCoordinator.getQuorumBitmapUpdateByOperatorIdByIndex(defaultOperatorId, 0))), 
            keccak256(abi.encode(IRegistryCoordinator.QuorumBitmapUpdate({
                quorumBitmap: uint192(quorumBitmap),
                updateBlockNumber: registrationBlockNumber,
                nextUpdateBlockNumber: deregistrationBlockNumber
            })))
        );
    }

    function testDeregisterOperatorWithCoordinatorForFuzzedQuorumAndSingleOperator_Valid(uint256 quorumBitmap) public {
        uint32 registrationBlockNumber = 100;
        uint32 deregistrationBlockNumber = 200;

        quorumBitmap = quorumBitmap & MiddlewareUtils.MAX_QUORUM_BITMAP;
        cheats.assume(quorumBitmap != 0);
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        for (uint i = 0; i < quorumNumbers.length; i++) {
            stakeRegistry.setOperatorWeight(uint8(quorumNumbers[i]), defaultOperator, defaultStake);
        }

        cheats.startPrank(defaultOperator);
        
        cheats.roll(registrationBlockNumber);
        
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, defaultPubKey, defaultSocket);

        bytes32[] memory operatorIdsToSwap = new bytes32[](quorumNumbers.length);
        for (uint i = 0; i < operatorIdsToSwap.length; i++) {
            operatorIdsToSwap[i] = defaultOperatorId;
        }

        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit OperatorAddedToQuorums(defaultOperator, quorumNumbers);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            cheats.expectEmit(true, true, true, true, address(stakeRegistry));
            emit StakeUpdate(defaultOperatorId, uint8(quorumNumbers[i]), 0);
        }

        cheats.roll(deregistrationBlockNumber);

        uint256 gasBefore = gasleft();
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbers, defaultPubKey, operatorIdsToSwap);
        uint256 gasAfter = gasleft();
        emit log_named_uint("gasUsed", gasBefore - gasAfter);
        emit log_named_uint("numQuorums", quorumNumbers.length);

        assertEq(
            keccak256(abi.encode(registryCoordinator.getOperator(defaultOperator))), 
            keccak256(abi.encode(IRegistryCoordinator.Operator({
                operatorId: defaultOperatorId,
                status: IRegistryCoordinator.OperatorStatus.DEREGISTERED
            })))
        );
        assertEq(registryCoordinator.getCurrentQuorumBitmapByOperatorId(defaultOperatorId), 0);
        assertEq(
            keccak256(abi.encode(registryCoordinator.getQuorumBitmapUpdateByOperatorIdByIndex(defaultOperatorId, 0))), 
            keccak256(abi.encode(IRegistryCoordinator.QuorumBitmapUpdate({
                quorumBitmap: uint192(quorumBitmap),
                updateBlockNumber: registrationBlockNumber,
                nextUpdateBlockNumber: deregistrationBlockNumber
            })))
        );
    }

    function testDeregisterOperatorWithCoordinatorForFuzzedQuorumAndManyOperators_Valid(uint256 pseudoRandomNumber) public {
        uint32 numOperators = defaultMaxOperatorCount;
        
        uint32 registrationBlockNumber = 100;
        uint32 deregistrationBlockNumber = 200;

        // pad quorumBitmap with 1 until it has numOperators elements
        uint256[] memory quorumBitmaps = new uint256[](numOperators);
        for (uint i = 0; i < numOperators; i++) {
            // limit to maxQuorumsToRegisterFor quorums via mask so we don't run out of gas, make them all register for quorum 0 as well
            quorumBitmaps[i] = uint256(keccak256(abi.encodePacked("quorumBitmap", pseudoRandomNumber, i))) & (1 << maxQuorumsToRegisterFor - 1) | 1;
        }

        cheats.roll(registrationBlockNumber);
        
        bytes32[] memory lastOperatorInQuorum = new bytes32[](numQuorums);
        for (uint i = 0; i < numOperators; i++) {
            emit log_named_uint("i", i);
            BN254.G1Point memory pubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, i)));
            bytes32 operatorId = pubKey.hashG1Point();
            address operator = _incrementAddress(defaultOperator, i);
            
            _registerOperatorWithCoordinator(operator, quorumBitmaps[i], pubKey);

            // for each quorum the operator is in, save the operatorId
            bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmaps[i]);
            for (uint j = 0; j < quorumNumbers.length; j++) {
                lastOperatorInQuorum[uint8(quorumNumbers[j])] = operatorId;
            }
        }

        uint256 indexOfOperatorToDerigister = pseudoRandomNumber % numOperators;
        address operatorToDerigister = _incrementAddress(defaultOperator, indexOfOperatorToDerigister);
        BN254.G1Point memory operatorToDeregisterPubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, indexOfOperatorToDerigister)));
        bytes32 operatorToDerigisterId = operatorToDeregisterPubKey.hashG1Point();
        uint256 operatorToDeregisterQuorumBitmap = quorumBitmaps[indexOfOperatorToDerigister];
        bytes memory operatorToDeregisterQuorumNumbers = BitmapUtils.bitmapToBytesArray(operatorToDeregisterQuorumBitmap);

        bytes32[] memory operatorIdsToSwap = new bytes32[](operatorToDeregisterQuorumNumbers.length);
        for (uint i = 0; i < operatorToDeregisterQuorumNumbers.length; i++) {
            operatorIdsToSwap[i] = lastOperatorInQuorum[uint8(operatorToDeregisterQuorumNumbers[i])];
        }

        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit OperatorAddedToQuorums(operatorToDerigister, operatorToDeregisterQuorumNumbers);
        
        for (uint i = 0; i < operatorToDeregisterQuorumNumbers.length; i++) {
            cheats.expectEmit(true, true, true, true, address(stakeRegistry));
            emit StakeUpdate(operatorToDerigisterId, uint8(operatorToDeregisterQuorumNumbers[i]), 0);
        }

        // expect events from the index registry
        for (uint i = 0; i < operatorToDeregisterQuorumNumbers.length; i++) {
            if(operatorIdsToSwap[i] != operatorToDerigisterId) {
                cheats.expectEmit(true, true, false, false, address(indexRegistry));
                emit QuorumIndexUpdate(operatorIdsToSwap[i], uint8(operatorToDeregisterQuorumNumbers[i]), 0);
            }
        }

        cheats.roll(deregistrationBlockNumber);

        cheats.prank(operatorToDerigister);
        registryCoordinator.deregisterOperatorWithCoordinator(operatorToDeregisterQuorumNumbers, operatorToDeregisterPubKey, operatorIdsToSwap);

        assertEq(
            keccak256(abi.encode(registryCoordinator.getOperator(operatorToDerigister))), 
            keccak256(abi.encode(IRegistryCoordinator.Operator({
                operatorId: operatorToDerigisterId,
                status: IRegistryCoordinator.OperatorStatus.DEREGISTERED
            })))
        );
        assertEq(registryCoordinator.getCurrentQuorumBitmapByOperatorId(defaultOperatorId), 0);
        assertEq(
            keccak256(abi.encode(registryCoordinator.getQuorumBitmapUpdateByOperatorIdByIndex(operatorToDerigisterId, 0))), 
            keccak256(abi.encode(IRegistryCoordinator.QuorumBitmapUpdate({
                quorumBitmap: uint192(operatorToDeregisterQuorumBitmap),
                updateBlockNumber: registrationBlockNumber,
                nextUpdateBlockNumber: deregistrationBlockNumber
            })))
        );
    }


    function testRegisterOperatorWithCoordinatorWithKicks_Valid(uint256 pseudoRandomNumber) public {
        uint32 numOperators = defaultMaxOperatorCount;
        uint32 kickRegistrationBlockNumber = 100;
        uint32 registrationBlockNumber = 200;

        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        cheats.roll(kickRegistrationBlockNumber);

        for (uint i = 0; i < numOperators - 1; i++) {
            BN254.G1Point memory pubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, i)));
            address operator = _incrementAddress(defaultOperator, i);
            
            _registerOperatorWithCoordinator(operator, quorumBitmap, pubKey);
        }

        address operatorToRegister = _incrementAddress(defaultOperator, numOperators);
        BN254.G1Point memory operatorToRegisterPubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, numOperators)));
        bytes32 operatorToRegisterId = operatorToRegisterPubKey.hashG1Point();
        bytes32 operatorToKickId;
        address operatorToKick;
        
        // register last operator before kick
        IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] memory operatorKickParams = new IBLSRegistryCoordinatorWithIndices.OperatorKickParam[](1);
        {
            BN254.G1Point memory pubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, numOperators - 1)));
            operatorToKickId = pubKey.hashG1Point();
            operatorToKick = _incrementAddress(defaultOperator, numOperators - 1);

            _registerOperatorWithCoordinator(operatorToKick, quorumBitmap, pubKey);

            bytes32[] memory operatorIdsToSwap = new bytes32[](1);
            // operatorIdsToSwap[0] = operatorToRegisterId
            operatorIdsToSwap[0] = operatorToRegisterId;

            operatorKickParams[0] = IBLSRegistryCoordinatorWithIndices.OperatorKickParam({
                operator: operatorToKick,
                pubkey: pubKey,
                operatorIdsToSwap: operatorIdsToSwap
            });
        }

        pubkeyCompendium.setBLSPublicKey(operatorToRegister, operatorToRegisterPubKey);

        uint96 registeringStake = defaultKickBIPsOfOperatorStake * defaultStake;
        stakeRegistry.setOperatorWeight(defaultQuorumNumber, operatorToRegister, registeringStake);

        cheats.roll(registrationBlockNumber);
        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit OperatorAddedToQuorums(operatorToRegister, quorumNumbers);
        cheats.expectEmit(true, true, true, true, address(stakeRegistry));
        emit StakeUpdate(operatorToRegisterId, defaultQuorumNumber, registeringStake);
        cheats.expectEmit(true, true, true, true, address(indexRegistry));
        emit QuorumIndexUpdate(operatorToRegisterId, defaultQuorumNumber, numOperators);

        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit OperatorAddedToQuorums(operatorKickParams[0].operator, quorumNumbers);
        cheats.expectEmit(true, true, true, true, address(stakeRegistry));
        emit StakeUpdate(operatorToKickId, defaultQuorumNumber, 0);
        cheats.expectEmit(true, true, true, true, address(indexRegistry));
        emit QuorumIndexUpdate(operatorToRegisterId, defaultQuorumNumber, numOperators - 1);

        {
            ISignatureUtils.SignatureWithSaltAndExpiry memory signatureWithExpiry = _signOperatorChurnApproval(operatorToRegisterId, operatorKickParams, defaultSalt, block.timestamp + 10);
            cheats.prank(operatorToRegister);
            uint256 gasBefore = gasleft();
            registryCoordinator.registerOperatorWithCoordinator(
                quorumNumbers, 
                operatorToRegisterPubKey, 
                defaultSocket, 
                operatorKickParams, 
                signatureWithExpiry
            );
            uint256 gasAfter = gasleft();
            emit log_named_uint("gasUsed", gasBefore - gasAfter);
        }

        assertEq(
            keccak256(abi.encode(registryCoordinator.getOperator(operatorToRegister))), 
            keccak256(abi.encode(IRegistryCoordinator.Operator({
                operatorId: operatorToRegisterId,
                status: IRegistryCoordinator.OperatorStatus.REGISTERED
            })))
        );
        assertEq(
            keccak256(abi.encode(registryCoordinator.getOperator(operatorToKick))), 
            keccak256(abi.encode(IRegistryCoordinator.Operator({
                operatorId: operatorToKickId,
                status: IRegistryCoordinator.OperatorStatus.DEREGISTERED
            })))
        );
        assertEq(
            keccak256(abi.encode(registryCoordinator.getQuorumBitmapUpdateByOperatorIdByIndex(operatorToKickId, 0))), 
            keccak256(abi.encode(IRegistryCoordinator.QuorumBitmapUpdate({
                quorumBitmap: uint192(quorumBitmap),
                updateBlockNumber: kickRegistrationBlockNumber,
                nextUpdateBlockNumber: registrationBlockNumber
            })))
        );
    }

    function testRegisterOperatorWithCoordinatorWithKicks_LessThanKickBIPsOfOperatorStake_Reverts(uint256 pseudoRandomNumber) public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        uint96 operatorToKickStake = defaultMaxOperatorCount * defaultStake;
        (   
            address operatorToRegister, 
            BN254.G1Point memory operatorToRegisterPubKey,
            IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] memory operatorKickParams
        ) = _testRegisterOperatorWithKicks_SetUp(pseudoRandomNumber, quorumNumbers, defaultStake);
        bytes32 operatorToRegisterId = operatorToRegisterPubKey.hashG1Point();

        stakeRegistry.setOperatorWeight(defaultQuorumNumber, operatorToRegister, defaultStake);

        cheats.roll(registrationBlockNumber);
        ISignatureUtils.SignatureWithSaltAndExpiry memory signatureWithExpiry = _signOperatorChurnApproval(operatorToRegisterId, operatorKickParams, defaultSalt, block.timestamp + 10);
        cheats.prank(operatorToRegister);
        cheats.expectRevert("BLSRegistryCoordinatorWithIndices.registerOperatorWithCoordinator: registering operator has less than kickBIPsOfOperatorStake");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, operatorToRegisterPubKey, defaultSocket, operatorKickParams, signatureWithExpiry);
    }

    function testRegisterOperatorWithCoordinatorWithKicks_LessThanKickBIPsOfTotalStake_Reverts(uint256 pseudoRandomNumber) public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        uint96 operatorToKickStake = defaultMaxOperatorCount * defaultStake;
        (   
            address operatorToRegister, 
            BN254.G1Point memory operatorToRegisterPubKey,
            IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] memory operatorKickParams
        ) = _testRegisterOperatorWithKicks_SetUp(pseudoRandomNumber, quorumNumbers, operatorToKickStake);
        bytes32 operatorToRegisterId = operatorToRegisterPubKey.hashG1Point();


        // set the stake of the operator to register to the defaultKickBIPsOfOperatorStake multiple of the operatorToKickStake
        stakeRegistry.setOperatorWeight(defaultQuorumNumber, operatorToRegister, operatorToKickStake * defaultKickBIPsOfOperatorStake / 10000 + 1);

        cheats.roll(registrationBlockNumber);
        ISignatureUtils.SignatureWithSaltAndExpiry memory signatureWithExpiry = _signOperatorChurnApproval(operatorToRegisterId, operatorKickParams, defaultSalt, block.timestamp + 10);
        cheats.prank(operatorToRegister);
        cheats.expectRevert("BLSRegistryCoordinatorWithIndices.registerOperatorWithCoordinator: operator to kick has more than kickBIPSOfTotalStake");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, operatorToRegisterPubKey, defaultSocket, operatorKickParams, signatureWithExpiry);
    }

    function testRegisterOperatorWithCoordinatorWithKicks_InvalidSignatures_Reverts(uint256 pseudoRandomNumber) public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        (   
            address operatorToRegister, 
            BN254.G1Point memory operatorToRegisterPubKey,
            IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] memory operatorKickParams
        ) = _testRegisterOperatorWithKicks_SetUp(pseudoRandomNumber, quorumNumbers, defaultStake);

        uint96 registeringStake = defaultKickBIPsOfOperatorStake * defaultStake;
        stakeRegistry.setOperatorWeight(defaultQuorumNumber, operatorToRegister, registeringStake);

        cheats.roll(registrationBlockNumber);
        ISignatureUtils.SignatureWithSaltAndExpiry memory signatureWithSaltAndExpiry;
        signatureWithSaltAndExpiry.expiry = block.timestamp + 10;
        signatureWithSaltAndExpiry.signature = hex"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001B";
        signatureWithSaltAndExpiry.salt = defaultSalt;
        cheats.prank(operatorToRegister);
        cheats.expectRevert("ECDSA: invalid signature");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, operatorToRegisterPubKey, defaultSocket, operatorKickParams, signatureWithSaltAndExpiry);
    }

    function testRegisterOperatorWithCoordinatorWithKicks_ExpiredSignatures_Reverts(uint256 pseudoRandomNumber) public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        (   
            address operatorToRegister, 
            BN254.G1Point memory operatorToRegisterPubKey,
            IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] memory operatorKickParams
        ) = _testRegisterOperatorWithKicks_SetUp(pseudoRandomNumber, quorumNumbers, defaultStake);
        bytes32 operatorToRegisterId = operatorToRegisterPubKey.hashG1Point();

        uint96 registeringStake = defaultKickBIPsOfOperatorStake * defaultStake;
        stakeRegistry.setOperatorWeight(defaultQuorumNumber, operatorToRegister, registeringStake);

        cheats.roll(registrationBlockNumber);
        ISignatureUtils.SignatureWithSaltAndExpiry memory signatureWithSaltAndExpiry = _signOperatorChurnApproval(operatorToRegisterId, operatorKickParams, defaultSalt, block.timestamp - 1);
        cheats.prank(operatorToRegister);
        cheats.expectRevert("BLSRegistryCoordinatorWithIndices._verifyChurnApproverSignatureOnOperatorChurnApproval: churnApprover signature expired");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, operatorToRegisterPubKey, defaultSocket, operatorKickParams, signatureWithSaltAndExpiry);
    }

    function testUpdateSocket() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);
        _registerOperatorWithCoordinator(defaultOperator, quorumBitmap, defaultPubKey);

        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(registryCoordinator));
        emit OperatorSocketUpdate(defaultOperatorId, "localhost:32004");
        registryCoordinator.updateSocket("localhost:32004");

    }

    function testUpdateSocket_NotRegistered_Reverts() public {
        cheats.prank(defaultOperator);
        cheats.expectRevert("BLSRegistryCoordinatorWithIndicies.updateSocket: operator is not registered");
        registryCoordinator.updateSocket("localhost:32004");
    }

    function _testRegisterOperatorWithKicks_SetUp(uint256 pseudoRandomNumber, bytes memory quorumNumbers, uint96 operatorToKickStake) internal returns(address operatorToRegister, BN254.G1Point memory operatorToRegisterPubKey, IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] memory operatorKickParams) {
        uint32 kickRegistrationBlockNumber = 100;
        uint32 registrationBlockNumber = 200;

        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        cheats.roll(kickRegistrationBlockNumber);

        for (uint i = 0; i < defaultMaxOperatorCount - 1; i++) {
            BN254.G1Point memory pubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, i)));
            address operator = _incrementAddress(defaultOperator, i);
            
            _registerOperatorWithCoordinator(operator, quorumBitmap, pubKey);
        }

        operatorToRegister = _incrementAddress(defaultOperator, defaultMaxOperatorCount);
        operatorToRegisterPubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, defaultMaxOperatorCount)));
        bytes32 operatorToRegisterId = operatorToRegisterPubKey.hashG1Point();
        bytes32 operatorToKickId;
        address operatorToKick;
        
        // register last operator before kick
        operatorKickParams = new IBLSRegistryCoordinatorWithIndices.OperatorKickParam[](1);
        {
            BN254.G1Point memory pubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, defaultMaxOperatorCount - 1)));
            operatorToKickId = pubKey.hashG1Point();
            operatorToKick = _incrementAddress(defaultOperator, defaultMaxOperatorCount - 1);

            // register last operator with much more than the kickBIPsOfTotalStake stake
            _registerOperatorWithCoordinator(operatorToKick, quorumBitmap, pubKey, operatorToKickStake);

            bytes32[] memory operatorIdsToSwap = new bytes32[](1);
            // operatorIdsToSwap[0] = operatorToRegisterId
            operatorIdsToSwap[0] = operatorToRegisterId;

            operatorKickParams[0] = IBLSRegistryCoordinatorWithIndices.OperatorKickParam({
                operator: operatorToKick,
                pubkey: pubKey,
                operatorIdsToSwap: operatorIdsToSwap
            });
        }

        pubkeyCompendium.setBLSPublicKey(operatorToRegister, operatorToRegisterPubKey);
    }
}