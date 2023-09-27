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
import "../../contracts/middleware/BLSOperatorStateRetriever.sol";
import "../../contracts/middleware/BLSRegistryCoordinatorWithIndices.sol";
import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/middleware/IndexRegistry.sol";

import "../../contracts/libraries/BitmapUtils.sol";
import "../../contracts/libraries/MiddlewareUtils.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/EigenPodManagerMock.sol";
import "../mocks/ServiceManagerMock.sol";
import "../mocks/OwnableMock.sol";
import "../mocks/DelegationMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/BLSPublicKeyCompendiumMock.sol";
import "../mocks/EmptyContract.sol";

import "../harnesses/StakeRegistryHarness.sol";
import "../harnesses/BLSRegistryCoordinatorWithIndicesHarness.sol";

import "forge-std/Test.sol";

contract MockAVSDeployer is Test {
    using BN254 for BN254.G1Point;

    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    ISlasher public slasher = ISlasher(address(uint160(uint256(keccak256("slasher")))));
    Slasher public slasherImplementation;

    EmptyContract public emptyContract;
    BLSPublicKeyCompendiumMock public pubkeyCompendium;

    BLSRegistryCoordinatorWithIndicesHarness public registryCoordinatorImplementation;
    StakeRegistryHarness public stakeRegistryImplementation;
    IBLSPubkeyRegistry public blsPubkeyRegistryImplementation;
    IIndexRegistry public indexRegistryImplementation;

    BLSOperatorStateRetriever public operatorStateRetriever;
    BLSRegistryCoordinatorWithIndicesHarness public registryCoordinator;
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

    uint256 churnApproverPrivateKey = uint256(keccak256("churnApproverPrivateKey"));
    address churnApprover = cheats.addr(churnApproverPrivateKey);
    bytes32 defaultSalt = bytes32(uint256(keccak256("defaultSalt")));

    address ejector = address(uint160(uint256(keccak256("ejector"))));

    address defaultOperator = address(uint160(uint256(keccak256("defaultOperator"))));
    bytes32 defaultOperatorId;
    BN254.G1Point internal defaultPubKey =  BN254.G1Point(18260007818883133054078754218619977578772505796600400998181738095793040006897,3432351341799135763167709827653955074218841517684851694584291831827675065899);
    string defaultSocket = "69.69.69.69:420";
    uint96 defaultStake = 1 ether;
    uint8 defaultQuorumNumber = 0;

    uint32 defaultMaxOperatorCount = 10;
    uint16 defaultKickBIPsOfOperatorStake = 15000;
    uint16 defaultKickBIPsOfTotalStake = 150;
    uint8 numQuorums = 192;

    IBLSRegistryCoordinatorWithIndices.OperatorSetParam[] operatorSetParams;

    uint8 maxQuorumsToRegisterFor = 4;
    uint256 maxOperatorsToRegister = 4;
    uint32 registrationBlockNumber = 100;
    uint32 blocksBetweenRegistrations = 10;

    struct OperatorMetadata {
        uint256 quorumBitmap;
        address operator;
        bytes32 operatorId;
        BN254.G1Point pubkey;
        uint96[] stakes; // in every quorum for simplicity
    }

    function _deployMockEigenLayerAndAVS() internal {
        _deployMockEigenLayerAndAVS(numQuorums);
    }

    function _deployMockEigenLayerAndAVS(uint8 numQuorumsToAdd) internal {
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
        registryCoordinator = BLSRegistryCoordinatorWithIndicesHarness(address(
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
        uint96[] memory minimumStakeForQuorum = new uint96[](numQuorumsToAdd);
        for (uint256 i = 0; i < minimumStakeForQuorum.length; i++) {
            minimumStakeForQuorum[i] = uint96(i+1);
        }

        // setup the dummy quorum strategies
        IVoteWeigher.StrategyAndWeightingMultiplier[][] memory quorumStrategiesConsideredAndMultipliers =
            new IVoteWeigher.StrategyAndWeightingMultiplier[][](numQuorumsToAdd);
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

        registryCoordinatorImplementation = new BLSRegistryCoordinatorWithIndicesHarness(
            slasher,
            serviceManagerMock,
            stakeRegistry,
            blsPubkeyRegistry,
            indexRegistry
        );
        {
            delete operatorSetParams;
            for (uint i = 0; i < numQuorumsToAdd; i++) {
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
                    churnApprover,
                    ejector,
                    operatorSetParams,
                    pauserRegistry,
                    0/*initialPausedStatus*/
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

        operatorStateRetriever = new BLSOperatorStateRetriever();

        cheats.stopPrank();
    }

    /**
     * @notice registers operator with coordinator 
     */
    function _registerOperatorWithCoordinator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubKey) internal {
        _registerOperatorWithCoordinator(operator, quorumBitmap, pubKey, defaultStake);
    }

    /**
     * @notice registers operator with coordinator 
     */
    function _registerOperatorWithCoordinator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubKey, uint96 stake) internal {
        // quorumBitmap can only have 192 least significant bits
        quorumBitmap &= MiddlewareUtils.MAX_QUORUM_BITMAP;

        pubkeyCompendium.setBLSPublicKey(operator, pubKey);

        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            stakeRegistry.setOperatorWeight(uint8(quorumNumbers[i]), operator, stake);
        }

        cheats.prank(operator);
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, pubKey, defaultSocket);
    }

    /**
     * @notice registers operator with coordinator 
     */
    function _registerOperatorWithCoordinator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubKey, uint96[] memory stakes) internal {
        // quorumBitmap can only have 192 least significant bits
        quorumBitmap &= MiddlewareUtils.MAX_QUORUM_BITMAP;

        pubkeyCompendium.setBLSPublicKey(operator, pubKey);

        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            stakeRegistry.setOperatorWeight(uint8(quorumNumbers[i]), operator, stakes[uint8(quorumNumbers[i])]);
        }

        cheats.prank(operator);
        registryCoordinator.registerOperatorWithCoordinator(quorumNumbers, pubKey, defaultSocket);
    }

    function _registerRandomOperators(uint256 pseudoRandomNumber) internal returns(OperatorMetadata[] memory, uint256[][] memory) {
        OperatorMetadata[] memory operatorMetadatas = new OperatorMetadata[](maxOperatorsToRegister);
        for (uint i = 0; i < operatorMetadatas.length; i++) {
            // limit to 16 quorums so we don't run out of gas, make them all register for quorum 0 as well
            operatorMetadatas[i].quorumBitmap = uint256(keccak256(abi.encodePacked("quorumBitmap", pseudoRandomNumber, i))) & (1 << maxQuorumsToRegisterFor - 1) | 1;
            operatorMetadatas[i].operator = _incrementAddress(defaultOperator, i);
            operatorMetadatas[i].pubkey = BN254.hashToG1(keccak256(abi.encodePacked("pubkey", pseudoRandomNumber, i)));
            operatorMetadatas[i].operatorId = operatorMetadatas[i].pubkey.hashG1Point();
            operatorMetadatas[i].stakes = new uint96[](maxQuorumsToRegisterFor);
            for (uint j = 0; j < maxQuorumsToRegisterFor; j++) {
                operatorMetadatas[i].stakes[j] = uint96(uint64(uint256(keccak256(abi.encodePacked("stakes", pseudoRandomNumber, i, j)))));
            }
        }

        // get the index in quorumBitmaps of each operator in each quorum in the order they will register
        uint256[][] memory expectedOperatorOverallIndices = new uint256[][](numQuorums);
        for (uint i = 0; i < numQuorums; i++) {
            uint32 numOperatorsInQuorum;
            // for each quorumBitmap, check if the i'th bit is set
            for (uint j = 0; j < operatorMetadatas.length; j++) {
                if (operatorMetadatas[j].quorumBitmap >> i & 1 == 1) {
                    numOperatorsInQuorum++;
                }
            }
            expectedOperatorOverallIndices[i] = new uint256[](numOperatorsInQuorum);
            uint256 numOperatorCounter;
            for (uint j = 0; j < operatorMetadatas.length; j++) {
                if (operatorMetadatas[j].quorumBitmap >> i & 1 == 1) {
                    expectedOperatorOverallIndices[i][numOperatorCounter] = j;
                    numOperatorCounter++;
                }
            }
        }

        // register operators
        for (uint i = 0; i < operatorMetadatas.length; i++) {
            cheats.roll(registrationBlockNumber + blocksBetweenRegistrations * i);
            
            _registerOperatorWithCoordinator(operatorMetadatas[i].operator, operatorMetadatas[i].quorumBitmap, operatorMetadatas[i].pubkey, operatorMetadatas[i].stakes);
        }

        return (operatorMetadatas, expectedOperatorOverallIndices);
    }

    function _incrementAddress(address start, uint256 inc) internal pure returns(address) {
        return address(uint160(uint256(uint160(start) + inc)));
    }

    function _incrementBytes32(bytes32 start, uint256 inc) internal pure returns(bytes32) {
        return bytes32(uint256(start) + inc);
    }

    function _signOperatorChurnApproval(bytes32 registeringOperatorId, IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] memory operatorKickParams, bytes32 salt,  uint256 expiry) internal view returns(ISignatureUtils.SignatureWithSaltAndExpiry memory) {
        bytes32 digestHash = registryCoordinator.calculateOperatorChurnApprovalDigestHash(
            registeringOperatorId,
            operatorKickParams,
            salt,
            expiry
        );
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(churnApproverPrivateKey, digestHash);
        return ISignatureUtils.SignatureWithSaltAndExpiry({
            signature: abi.encodePacked(r, s, v),
            expiry: expiry,
            salt: salt
        });
    }
}