// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/core/Slasher.sol";
import "../../contracts/permissions/PauserRegistry.sol";
import "../../contracts/interfaces/IStrategyManager.sol";
import "../../contracts/interfaces/IStakeRegistry.sol";
import "../../contracts/interfaces/IServiceManager.sol";
import "../../contracts/interfaces/IVoteWeigher.sol";

import "../../contracts/middleware/BLSPublicKeyCompendium.sol";
import "../../contracts/middleware/BLSRegistryCoordinatorWithIndices.sol";
import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/middleware/IndexRegistry.sol";

import "../../contracts/libraries/BitmapUtils.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/EigenPodManagerMock.sol";
import "../mocks/ServiceManagerMock.sol";
import "../mocks/OwnableMock.sol";
import "../mocks/DelegationMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/BLSPublicKeyCompendiumMock.sol";
import "../mocks/EmptyContract.sol";

import "../harnesses/StakeRegistryHarness.sol";

import "forge-std/Test.sol";

contract BLSRegistryCoordinatorWithIndicesUnit is Test {
    using BN254 for BN254.G1Point;

    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    ISlasher public slasher = ISlasher(address(uint160(uint256(keccak256("slasher")))));
    Slasher public slasherImplementation;

    EmptyContract public emptyContract;
    BLSPublicKeyCompendiumMock public pubkeyCompendium;

    IBLSRegistryCoordinatorWithIndices public registryCoordinatorImplementation;
    StakeRegistryHarness public stakeRegistryImplementation;
    IBLSPubkeyRegistry public blsPubkeyRegistryImplementation;
    IIndexRegistry public indexRegistryImplementation;

    BLSRegistryCoordinatorWithIndices public registryCoordinator;
    StakeRegistryHarness public stakeRegistry;
    IBLSPubkeyRegistry public blsPubkeyRegistry;
    IIndexRegistry public indexRegistry;

    ServiceManagerMock public serviceManagerMock;
    StrategyManagerMock public strategyManagerMock;
    DelegationMock public delegationMock;
    EigenPodManagerMock public eigenPodManagerMock;

    address public proxyAdminOwner = address(uint160(uint256(keccak256("proxyAdminOwner"))));
    address public serviceManagerOwner = address(uint160(uint256(keccak256("serviceManagerOwner"))));
    address public pauser = address(uint160(uint256(keccak256("pauser"))));
    address public unpauser = address(uint160(uint256(keccak256("unpauser"))));

    address defaultOperator = address(uint160(uint256(keccak256("defaultOperator"))));
    bytes32 defaultOperatorId;
    BN254.G1Point internal defaultPubKey =  BN254.G1Point(18260007818883133054078754218619977578772505796600400998181738095793040006897,3432351341799135763167709827653955074218841517684851694584291831827675065899);
    string defaultSocket = "69.69.69.69:420";
    uint96 defaultStake = 1 ether;
    uint8 defaultQuorumNumber = 0;

    uint32 defaultMaxOperatorCount = 100;
    uint16 defaultKickBIPsOfOperatorStake = 15000;
    uint16 defaultKickBIPsOfTotalStake = 150;
    uint8 numQuorums = 192;

    IBLSRegistryCoordinatorWithIndices.OperatorSetParam[] operatorSetParams;

    event OperatorSocketUpdate(bytes32 operatorId, string socket);

    /// @notice emitted whenever the stake of `operator` is updated
    event StakeUpdate(
        bytes32 indexed operatorId,
        uint8 quorumNumber,
        uint96 stake
    );

    // Emitted when a new operator pubkey is registered for a set of quorums
    event PubkeyAddedToQuorums(
        address operator,
        bytes quorumNumbers
    );

    // Emitted when an operator pubkey is removed from a set of quorums
    event PubkeyRemovedFromQuorums(
        address operator, 
        bytes quorumNumbers
    );

    // emitted when an operator's index in the orderd operator list for the quorum with number `quorumNumber` is updated
    event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newIndex);

    function setUp() virtual public {
        emptyContract = new EmptyContract();

        defaultOperatorId = defaultPubKey.hashG1Point();

        cheats.startPrank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        delegationMock = new DelegationMock();
        eigenPodManagerMock = new EigenPodManagerMock();
        strategyManagerMock = new StrategyManagerMock();
        slasherImplementation = new Slasher(strategyManagerMock, delegationMock);
        slasher = Slasher(
            address(
                new TransparentUpgradeableProxy(
                    address(slasherImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(Slasher.initialize.selector, msg.sender, pauserRegistry, 0/*initialPausedStatus*/)
                )
            )
        );

        strategyManagerMock.setAddresses(
            delegationMock,
            eigenPodManagerMock,
            slasher
        );

        pubkeyCompendium = new BLSPublicKeyCompendiumMock();
        pubkeyCompendium.setBLSPublicKey(defaultOperator, defaultPubKey);

        cheats.stopPrank();

        cheats.startPrank(serviceManagerOwner);
        // make the serviceManagerOwner the owner of the serviceManager contract
        serviceManagerMock = new ServiceManagerMock(slasher);
        registryCoordinator = BLSRegistryCoordinatorWithIndices(address(
            new TransparentUpgradeableProxy(
                address(emptyContract),
                address(proxyAdmin),
                ""
            )
        ));

        stakeRegistry = StakeRegistryHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(proxyAdmin),
                    ""
                )
            )
        );

        indexRegistry = IndexRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(proxyAdmin),
                    ""
                )
            )
        );

        blsPubkeyRegistry = BLSPubkeyRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(proxyAdmin),
                    ""
                )
            )
        );

        stakeRegistryImplementation = new StakeRegistryHarness(
            IRegistryCoordinator(registryCoordinator),
            strategyManagerMock,
            IServiceManager(address(serviceManagerMock))
        );

        cheats.stopPrank();
        cheats.startPrank(proxyAdminOwner);

        // setup the dummy minimum stake for quorum
        uint96[] memory minimumStakeForQuorum = new uint96[](numQuorums);
        for (uint256 i = 0; i < minimumStakeForQuorum.length; i++) {
            minimumStakeForQuorum[i] = uint96(i+1);
        }

        // setup the dummy quorum strategies
        IVoteWeigher.StrategyAndWeightingMultiplier[][] memory quorumStrategiesConsideredAndMultipliers =
            new IVoteWeigher.StrategyAndWeightingMultiplier[][](numQuorums);
        for (uint256 i = 0; i < quorumStrategiesConsideredAndMultipliers.length; i++) {
            quorumStrategiesConsideredAndMultipliers[i] = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
            quorumStrategiesConsideredAndMultipliers[i][0] = IVoteWeigher.StrategyAndWeightingMultiplier(
                IStrategy(address(uint160(i))),
                uint96(i+1)
            );
        }

        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(stakeRegistry))),
            address(stakeRegistryImplementation),
            abi.encodeWithSelector(
                StakeRegistry.initialize.selector,
                minimumStakeForQuorum,
                quorumStrategiesConsideredAndMultipliers
            )
        );

        registryCoordinatorImplementation = new BLSRegistryCoordinatorWithIndices(
            slasher,
            serviceManagerMock,
            stakeRegistry,
            blsPubkeyRegistry,
            indexRegistry
        );
        {
            for (uint i = 0; i < numQuorums; i++) {
                // hard code these for now
                operatorSetParams.push(IBLSRegistryCoordinatorWithIndices.OperatorSetParam({
                    maxOperatorCount: defaultMaxOperatorCount,
                    kickBIPsOfOperatorStake: defaultKickBIPsOfOperatorStake,
                    kickBIPsOfTotalStake: defaultKickBIPsOfTotalStake
                }));
            }
            proxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(registryCoordinator))),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    BLSRegistryCoordinatorWithIndices.initialize.selector,
                    operatorSetParams
                )
            );
        }

        blsPubkeyRegistryImplementation = new BLSPubkeyRegistry(
            registryCoordinator,
            BLSPublicKeyCompendium(address(pubkeyCompendium))
        );

        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(blsPubkeyRegistry))),
            address(blsPubkeyRegistryImplementation)
        );

        indexRegistryImplementation = new IndexRegistry(
            registryCoordinator
        );

        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(indexRegistry))),
            address(indexRegistryImplementation)
        );

        cheats.stopPrank();
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
        registryCoordinator.initialize(operatorSetParams);
    }

    function testRegisterOperatorWithCoordinator_EmptyQuorumNumbers_Reverts() public {
        bytes memory emptyQuorumNumbers = new bytes(0);
        cheats.expectRevert("BLSIndexRegistryCoordinator._registerOperatorWithCoordinator: quorumBitmap cannot be 0");
        cheats.prank(defaultOperator);
        registryCoordinator.registerOperatorWithCoordinator(emptyQuorumNumbers, defaultPubKey, defaultSocket);
    }

    function testRegisterOperatorWithCoordinator_QuorumNumbersTooLarge_Reverts() public {
        bytes memory quorumNumbersTooLarge = new bytes(1);
        quorumNumbersTooLarge[0] = 0xC0;
        cheats.expectRevert("BLSIndexRegistryCoordinator._registerOperatorWithCoordinator: quorumBitmap cant have more than 192 set bits");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbersTooLarge, defaultPubKey, defaultSocket);
    }

    function testRegisterOperatorWithCoordinatorForSingleQuorum_Valid() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        stakeRegistry.setOperatorWeight(uint8(quorumNumbers[0]), defaultOperator, defaultStake);

        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit PubkeyAddedToQuorums(defaultOperator, quorumNumbers);
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
        quorumBitmap = quorumBitmap & type(uint192).max;
        cheats.assume(quorumBitmap != 0);
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        for (uint i = 0; i < quorumNumbers.length; i++) {
            stakeRegistry.setOperatorWeight(uint8(quorumNumbers[i]), defaultOperator, defaultStake);
        }

        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit PubkeyAddedToQuorums(defaultOperator, quorumNumbers);

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

    function testDeregisterOperatorWithCoordinator_NotRegistered_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        cheats.expectRevert("BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: operator is not registered");
        cheats.prank(defaultOperator);
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbers, defaultPubKey, new bytes32[](0));
    }

    function testDeregisterOperatorWithCoordinator_IncorrectPubkey_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        uint256 quorumBitmap = BitmapUtils.orderedBytesArrayToBitmap(quorumNumbers);

        _registerOperatorWithCoordinator(defaultOperator, quorumBitmap, defaultPubKey);

        BN254.G1Point memory incorrectPubKey = BN254.hashToG1(bytes32(uint256(123)));

        cheats.expectRevert("BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: operatorId does not match pubkey hash");
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

        cheats.expectRevert("BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: quorumsToRemoveBitmap cant have more than 192 set bits");
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

        cheats.expectRevert("BLSIndexRegistryCoordinator._deregisterOperatorWithCoordinator: cannot deregister operator for quorums that it is not a part of");
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
        emit PubkeyRemovedFromQuorums(defaultOperator, quorumNumbers);
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
        cheats.expectRevert("BLSRegistryCoordinator.getCurrentQuorumBitmapByOperatorId: operator is not registered");
        registryCoordinator.getCurrentQuorumBitmapByOperatorId(defaultOperatorId);
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

        quorumBitmap = quorumBitmap & type(uint192).max;
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
        emit PubkeyRemovedFromQuorums(defaultOperator, quorumNumbers);
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
        cheats.expectRevert("BLSRegistryCoordinator.getCurrentQuorumBitmapByOperatorId: operator is not registered");
        registryCoordinator.getCurrentQuorumBitmapByOperatorId(defaultOperatorId);
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
        uint8 numOperators = 100;
        uint32 registrationBlockNumber = 100;
        uint32 deregistrationBlockNumber = 200;

        // pad quorumBitmap with 1 until it has numOperators elements
        uint256[] memory quorumBitmaps = new uint256[](numOperators);
        for (uint i = 0; i < numOperators; i++) {
            quorumBitmaps[i] = uint256(keccak256(abi.encodePacked("quorumBitmap", pseudoRandomNumber, i))) & type(uint192).max;
            if (quorumBitmaps[i] == 0) {
                quorumBitmaps[i] = 1;
            }
        }

        cheats.roll(registrationBlockNumber);
        
        bytes32[] memory lastOperatorInQuorum = new bytes32[](numQuorums);
        for (uint i = 0; i < numOperators; i++) {
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
        emit PubkeyRemovedFromQuorums(operatorToDerigister, operatorToDeregisterQuorumNumbers);
        
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
        cheats.expectRevert("BLSRegistryCoordinator.getCurrentQuorumBitmapByOperatorId: operator is not registered");
        registryCoordinator.getCurrentQuorumBitmapByOperatorId(operatorToDerigisterId);
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
        bytes32[] memory operatorIdsToSwap = new bytes32[](1);
        // operatorIdsToSwap[0] = operatorToRegisterId
        operatorIdsToSwap[0] = operatorToRegisterPubKey.hashG1Point();
        bytes32 operatorToKickId;
        address operatorToKick;
        
        // register last operator before kick
        IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] memory operatorKickParams = new IBLSRegistryCoordinatorWithIndices.OperatorKickParam[](1);
        {
            BN254.G1Point memory pubKey = BN254.hashToG1(keccak256(abi.encodePacked(pseudoRandomNumber, numOperators - 1)));
            operatorToKickId = pubKey.hashG1Point();
            operatorToKick = _incrementAddress(defaultOperator, numOperators - 1);

            _registerOperatorWithCoordinator(operatorToKick, quorumBitmap, pubKey);

            operatorKickParams[0] = IBLSRegistryCoordinatorWithIndices.OperatorKickParam({
                operator: operatorToKick,
                pubkey: pubKey,
                operatorIdsToSwap: operatorIdsToSwap
            });
        }

        pubkeyCompendium.setBLSPublicKey(operatorToRegister, operatorToRegisterPubKey);

        uint96 registeringStake = defaultKickBIPsOfOperatorStake * defaultStake;
        stakeRegistry.setOperatorWeight(defaultQuorumNumber, operatorToRegister, registeringStake);

        cheats.prank(operatorToRegister);
        cheats.roll(registrationBlockNumber);
        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit PubkeyAddedToQuorums(operatorToRegister, quorumNumbers);
        cheats.expectEmit(true, true, true, true, address(stakeRegistry));
        emit StakeUpdate(operatorIdsToSwap[0], defaultQuorumNumber, registeringStake);
        cheats.expectEmit(true, true, true, true, address(indexRegistry));
        emit QuorumIndexUpdate(operatorIdsToSwap[0], defaultQuorumNumber, numOperators);

        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit PubkeyRemovedFromQuorums(operatorKickParams[0].operator, quorumNumbers);
        cheats.expectEmit(true, true, true, true, address(stakeRegistry));
        emit StakeUpdate(operatorToKickId, defaultQuorumNumber, 0);
        cheats.expectEmit(true, true, true, true, address(indexRegistry));
        emit QuorumIndexUpdate(operatorIdsToSwap[0], defaultQuorumNumber, numOperators - 1);

        {
            uint256 gasBefore = gasleft();
            registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, operatorToRegisterPubKey, defaultSocket, operatorKickParams);
            uint256 gasAfter = gasleft();
            emit log_named_uint("gasUsed", gasBefore - gasAfter);
        }

        assertEq(
            keccak256(abi.encode(registryCoordinator.getOperator(operatorToRegister))), 
            keccak256(abi.encode(IRegistryCoordinator.Operator({
                operatorId: operatorIdsToSwap[0],
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

    /**
     * @notice registers operator with coordinator 
     */
    function _registerOperatorWithCoordinator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubKey) internal {
        // quorumBitmap can only have 192 least significant bits
        quorumBitmap &= type(uint192).max;

        pubkeyCompendium.setBLSPublicKey(operator, pubKey);

        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            stakeRegistry.setOperatorWeight(uint8(quorumNumbers[i]), operator, defaultStake);
        }

        cheats.prank(operator);
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, pubKey, defaultSocket);
    }

    function _incrementAddress(address start, uint256 inc) internal pure returns(address) {
        return address(uint160(uint256(uint160(start) + inc)));
    }

    function _incrementBytes32(bytes32 start, uint256 inc) internal pure returns(bytes32) {
        return bytes32(uint256(start) + inc);
    }
}