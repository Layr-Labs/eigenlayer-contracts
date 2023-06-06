// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/core/Slasher.sol";
import "../../contracts/permissions/PauserRegistry.sol";
import "../../contracts/interfaces/IStrategyManager.sol";
import "../../contracts/interfaces/IServiceManager.sol";
import "../../contracts/interfaces/IVoteWeigher.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/EigenPodManagerMock.sol";
import "../mocks/ServiceManagerMock.sol";
import "../mocks/OwnableMock.sol";
import "../mocks/DelegationMock.sol";
import "../mocks/SlasherMock.sol";

import "../harnesses/StakeRegistryHarness.sol";

import "forge-std/Test.sol";

contract StakeRegistryUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    ISlasher public slasher = ISlasher(address(uint160(uint256(keccak256("slasher")))));

    Slasher public slasherImplementation;
    StakeRegistryHarness public stakeRegistryImplementation;
    StakeRegistryHarness public stakeRegistry;

    ServiceManagerMock public serviceManagerMock;
    StrategyManagerMock public strategyManagerMock;
    DelegationMock public delegationMock;
    EigenPodManagerMock public eigenPodManagerMock;

    address public serviceManagerOwner = address(uint160(uint256(keccak256("serviceManagerOwner"))));
    address public registryCoordinator = address(uint160(uint256(keccak256("registryCoordinator"))));
    address public pauser = address(uint160(uint256(keccak256("pauser"))));
    address public unpauser = address(uint160(uint256(keccak256("unpauser"))));

    address defaultOperator = address(uint160(uint256(keccak256("defaultOperator"))));
    bytes32 defaultOperatorId = keccak256("defaultOperatorId");
    uint8 defaultQuorumNumber = 0;
    uint8 numQuorums = 128;

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();
        pauserRegistry = new PauserRegistry(pauser, unpauser);

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

        cheats.startPrank(serviceManagerOwner);
        // make the serviceManagerOwner the owner of the serviceManager contract
        serviceManagerMock = new ServiceManagerMock(slasher);
        stakeRegistryImplementation = new StakeRegistryHarness(
            IRegistryCoordinator(registryCoordinator),
            strategyManagerMock,
            IServiceManager(address(serviceManagerMock))
        );

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

        stakeRegistry = StakeRegistryHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(stakeRegistryImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        StakeRegistry.initialize.selector,
                        minimumStakeForQuorum,
                        quorumStrategiesConsideredAndMultipliers // initialize with 0ed out 128 quorums
                    )
                )
            )
        );

        cheats.stopPrank();
    }

    function testCorrectConstruction() public {
        // make sure the contract intializers are disabled
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        stakeRegistryImplementation.initialize(new uint96[](0), new IVoteWeigher.StrategyAndWeightingMultiplier[][](0));
    }

    function testSetMinimumStakeForQuorum_NotFromServiceManager_Reverts() public {
        cheats.expectRevert("VoteWeigherBase.onlyServiceManagerOwner: caller is not the owner of the serviceManager");
        stakeRegistry.setMinimumStakeForQuorum(defaultQuorumNumber, 0);
    }

    function testSetMinimumStakeForQuorum_Valid(uint8 quorumNumber, uint96 minimumStakeForQuorum) public {
        // set the minimum stake for quorum
        cheats.prank(serviceManagerOwner);
        stakeRegistry.setMinimumStakeForQuorum(quorumNumber, minimumStakeForQuorum);

        // make sure the minimum stake for quorum is as expected
        assertEq(stakeRegistry.minimumStakeForQuorum(quorumNumber), minimumStakeForQuorum);
    }

    function testRegisterOperator_NotFromRegistryCoordinator_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        cheats.expectRevert("StakeRegistry.onlyRegistryCoordinator: caller is not the RegistryCoordinator");
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function testRegisterOperator_NotOptedIntoSlashing_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        cheats.expectRevert("StakeRegistry._registerOperator: operator must be opted into slashing by the serviceManager");
        cheats.prank(registryCoordinator);
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function testRegisterOperator_MoreQuorumsThanQuorumCount_Reverts() public {
        // opt into slashing
        cheats.startPrank(defaultOperator);
        delegationMock.setIsOperator(defaultOperator, true);
        slasher.optIntoSlashing(address(serviceManagerMock));
        cheats.stopPrank();

        bytes memory quorumNumbers = new bytes(numQuorums+1);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            quorumNumbers[i] = bytes1(uint8(i));
        }

        // expect that it reverts when you register
        cheats.expectRevert("StakeRegistry._registerStake: greatest quorumNumber must be less than quorumCount");
        cheats.prank(registryCoordinator);
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function testRegisterOperator_LessThanMinimumStakeForQuorum_Reverts(
        uint96[] memory stakesForQuorum
    ) public {
        cheats.assume(stakesForQuorum.length > 0);
        // opt into slashing
        cheats.startPrank(defaultOperator);
        delegationMock.setIsOperator(defaultOperator, true);
        slasher.optIntoSlashing(address(serviceManagerMock));
        cheats.stopPrank();

        // set the weights of the operator
        stakeRegistry.setOperatorWeight()

        bytes memory quorumNumbers = new bytes(stakesForQuorum.length > 128 ? 128 : stakesForQuorum.length);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            quorumNumbers[i] = bytes1(uint8(i));
        }

        stakesForQuorum[stakesForQuorum.length - 1] = stakeRegistry.minimumStakeForQuorum(uint8(quorumNumbers.length - 1)) - 1;

        // expect that it reverts when you register
        cheats.expectRevert("StakeRegistry._registerStake: Operator does not meet minimum stake requirement for quorum");
        cheats.prank(registryCoordinator);
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function testOperatorStakeUpdate_Valid(
        uint24[] memory blocksPassed,
        uint96[] memory stakes
    ) public {
        cheats.assume(blocksPassed.length > 0);
        cheats.assume(blocksPassed.length <= stakes.length);
        // initialize at a non-zero block number
        uint32 intialBlockNumber = 100;
        cheats.roll(intialBlockNumber);
        uint32 cumulativeBlockNumber = intialBlockNumber;
        // loop through each one of the blocks passed, roll that many blocks, set the weight in the given quorum to the stake, and trigger a stake update
        for (uint256 i = 0; i < blocksPassed.length; i++) {
            stakeRegistry.setOperatorWeight(defaultQuorumNumber, defaultOperator, stakes[i]);
            stakeRegistry.updateOperatorStake(defaultOperator, defaultOperatorId, defaultQuorumNumber);
            cumulativeBlockNumber += blocksPassed[i];
            cheats.roll(cumulativeBlockNumber);
        }

        // reset for checking indices
        cumulativeBlockNumber = intialBlockNumber;
        // make sure that the stake updates are as expected
        for (uint256 i = 0; i < blocksPassed.length - 1; i++) {
            IStakeRegistry.OperatorStakeUpdate memory operatorStakeUpdate = stakeRegistry.getStakeUpdateForQuorumFromOperatorIdAndIndex(defaultQuorumNumber, defaultOperatorId, i);
            
            uint96 expectedStake = stakes[i];
            if (expectedStake < stakeRegistry.minimumStakeForQuorum(defaultQuorumNumber)) {
                expectedStake = 0;
            }

            assertEq(operatorStakeUpdate.stake, stakes[i]);
            assertEq(operatorStakeUpdate.updateBlockNumber, cumulativeBlockNumber);
            cumulativeBlockNumber += blocksPassed[i];
            assertEq(operatorStakeUpdate.nextUpdateBlockNumber, cumulativeBlockNumber);
        }

        // make sure that the last stake update is as expected
        IStakeRegistry.OperatorStakeUpdate memory lastOperatorStakeUpdate = stakeRegistry.getStakeUpdateForQuorumFromOperatorIdAndIndex(defaultQuorumNumber, defaultOperatorId, blocksPassed.length - 1);
        assertEq(lastOperatorStakeUpdate.stake, stakes[blocksPassed.length - 1]);
        assertEq(lastOperatorStakeUpdate.updateBlockNumber, cumulativeBlockNumber);
        assertEq(lastOperatorStakeUpdate.nextUpdateBlockNumber, uint32(0));
    }
}