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
    uint96 defaultStake = 1 ether;
    uint8 defaultQuorumNumber = 0;
    uint8 numQuorums = 192;

    IBLSRegistryCoordinatorWithIndices.OperatorSetParam[] operatorSetParams;


    /// @notice emitted whenever the stake of `operator` is updated
    event StakeUpdate(
        bytes32 indexed operatorId,
        uint8 quorumNumber,
        uint96 stake
    );

    // Emitted when a new operator pubkey is registered
    event PubkeyAdded(
        address operator,
        BN254.G1Point pubkey,
        bytes quorumNumbers
    );

    // Emitted when an operator pubkey is deregistered
    event PubkeyRemoved(
        address operator,
        BN254.G1Point pubkey,
        bytes quorumNumbers
    );

    // emitted when an operator's index in the orderd operator list for the quorum with number `quorumNumber` is updated
    event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newIndex);
    // emitted when an operator's index in the global operator list is updated
    event GlobalIndexUpdate(bytes32 indexed operatorId, uint32 newIndex);

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
                    maxOperatorCount: 10000,
                    kickBIPsOfOperatorStake: 15000,
                    kickBIPsOfAverageStake: 5000,
                    kickBIPsOfTotalStake: 100
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
        registryCoordinator.registerOperatorWithCoordinator(emptyQuorumNumbers, defaultPubKey);
    }

    function testRegisterOperatorWithCoordinator_QuorumNumbersTooLarge_Reverts() public {
        bytes memory quorumNumbersTooLarge = new bytes(1);
        quorumNumbersTooLarge[0] = 0xC0;
        cheats.expectRevert("BLSIndexRegistryCoordinator._registerOperatorWithCoordinator: quorumBitmap cant have more than 192 set bits");
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbersTooLarge, defaultPubKey);
    }

    function testRegisterOperatorWithCoordinatorForSingleQuorum_Valid() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        stakeRegistry.setOperatorWeight(uint8(quorumNumbers[0]), defaultOperator, defaultStake);

        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit PubkeyAdded(defaultOperator, defaultPubKey, quorumNumbers);
        cheats.expectEmit(true, true, true, true, address(stakeRegistry));
        emit StakeUpdate(defaultOperatorId, defaultQuorumNumber, defaultStake);
        cheats.expectEmit(true, true, true, true, address(indexRegistry));
        emit QuorumIndexUpdate(defaultOperatorId, defaultQuorumNumber, 0);
        uint256 gasBefore = gasleft();
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, defaultPubKey);
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
        emit PubkeyAdded(defaultOperator, defaultPubKey, quorumNumbers);

        for (uint i = 0; i < quorumNumbers.length; i++) {
            cheats.expectEmit(true, true, true, true, address(stakeRegistry));
            emit StakeUpdate(defaultOperatorId, uint8(quorumNumbers[i]), defaultStake);
        }    

        for (uint i = 0; i < quorumNumbers.length; i++) {
            cheats.expectEmit(true, true, true, true, address(indexRegistry));
            emit QuorumIndexUpdate(defaultOperatorId, uint8(quorumNumbers[i]), 0);
        }    
        uint256 gasBefore = gasleft();
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, defaultPubKey);
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

    function testRegisterOperatorWithCoordinatorForSingleQuorum_Valid() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        stakeRegistry.setOperatorWeight(uint8(quorumNumbers[0]), defaultOperator, defaultStake);

        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(blsPubkeyRegistry));
        emit PubkeyAdded(defaultOperator, defaultPubKey, quorumNumbers);
        cheats.expectEmit(true, true, true, true, address(stakeRegistry));
        emit StakeUpdate(defaultOperatorId, defaultQuorumNumber, defaultStake);
        cheats.expectEmit(true, true, true, true, address(indexRegistry));
        emit QuorumIndexUpdate(defaultOperatorId, defaultQuorumNumber, 0);
        uint256 gasBefore = gasleft();
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, defaultPubKey);
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

    function _incrementAddress(address start, uint256 inc) internal pure returns(address) {
        return address(uint160(uint256(uint160(start) + inc)));
    }

    function _incrementBytes32(bytes32 start, uint256 inc) internal pure returns(bytes32) {
        return bytes32(uint256(start) + inc);
    }
}