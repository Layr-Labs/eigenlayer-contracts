// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/permissions/PauserRegistry.sol";
import "../../contracts/interfaces/IStrategyManager.sol";
import "../../contracts/interfaces/IServiceManager.sol";
import "../../contracts/interfaces/IVoteWeigher.sol";

import "../mocks/StrategyManagerMock.sol";
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

    IStrategyManager public strategyManager;
    ISlasher public slasher = ISlasher(address(uint160(uint256(keccak256("slasher")))));

    StakeRegistryHarness public stakeRegistryImplementation;
    StakeRegistryHarness public stakeRegistry;

    ServiceManagerMock public serviceManagerMock;

    address public serviceManagerOwner = address(uint160(uint256(keccak256("serviceManagerOwner"))));
    address public registryCoordinator = address(uint160(uint256(keccak256("registryCoordinator"))));
    address public pauser = address(uint160(uint256(keccak256("pauser"))));
    address public unpauser = address(uint160(uint256(keccak256("unpauser"))));

    address defaultOperator = address(uint160(uint256(keccak256("defaultOperator"))));
    bytes32 defaultOperatorId = keccak256("defaultOperatorId");
    uint8 defaultQuorumNumber = 0;

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();
        pauserRegistry = new PauserRegistry(pauser, unpauser);

        strategyManager = new StrategyManagerMock();

        // make the serviceManagerOwner the owner of the serviceManager contract
        cheats.startPrank(serviceManagerOwner);
        slasher = new SlasherMock();
        serviceManagerMock = new ServiceManagerMock(slasher);

        stakeRegistryImplementation = new StakeRegistryHarness(
            IRegistryCoordinator(registryCoordinator),
            strategyManager,
            IServiceManager(address(serviceManagerMock))
        );

        // setup the dummy minimum stake for quorum
        uint96[] memory minimumStakeForQuorum = new uint96[](128);
        for (uint256 i = 0; i < minimumStakeForQuorum.length; i++) {
            minimumStakeForQuorum[i] = uint96(i+1);
        }

        // setup the dummy quorum strategies
        IVoteWeigher.StrategyAndWeightingMultiplier[][] memory quorumStrategiesConsideredAndMultipliers =
            new IVoteWeigher.StrategyAndWeightingMultiplier[][](128);
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