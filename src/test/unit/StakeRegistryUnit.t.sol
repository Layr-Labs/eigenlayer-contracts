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

    StakeRegistryHarness public stakeRegistry;

    ServiceManagerMock public serviceManagerMock;

    address public serviceManagerOwner = address(uint160(uint256(keccak256("serviceManagerOwner"))));
    address public pauser = address(uint160(uint256(keccak256("pauser"))));
    address public unpauser = address(uint160(uint256(keccak256("unpauser"))));

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();
        pauserRegistry = new PauserRegistry(pauser, unpauser);

        strategyManager = new StrategyManagerMock();

        // make the serviceManagerOwner the owner of the serviceManager contract
        cheats.startPrank(serviceManagerOwner);
        slasher = new SlasherMock();
        serviceManagerMock = new ServiceManagerMock(slasher);

        stakeRegistry = new StakeRegistryHarness(
            strategyManager,
            IServiceManager(address(serviceManagerMock))
        );
    }

    function testCorrectConstruction() public {
        // make sure the contract intializers are disabled
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        stakeRegistry.initialize(new uint96[](0), new IVoteWeigher.StrategyAndWeightingMultiplier[][](0));
    }

    function testOperatorStakeUpdate_Valid(
        uint256 quorumNumber,
        address operator,
        bytes32 operatorId,
        uint24[] memory blocksPassed,
        uint96[] memory stakes
    ) public {
        // loop through each one of the blocks passed, roll that many blocks, set the weight in the given quorum to the stake, and trigger a stake update
        for (uint256 i = 0; i < blocksPassed.length; i++) {
            cheats.roll(blocksPassed[i]);
            stakeRegistry.setOperatorWeight(operator, stakes[i]);
            stakeRegistry.updateOperatorStake(operator, operatorId, quorumNumber);
        }
    }
}