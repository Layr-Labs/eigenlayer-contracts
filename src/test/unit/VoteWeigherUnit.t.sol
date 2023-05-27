// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/permissions/PauserRegistry.sol";
import "../../contracts/interfaces/IStrategyManager.sol";
import "../../contracts/interfaces/IServiceManager.sol";
import "../../contracts/interfaces/IVoteWeigher.sol";
import "../../contracts/middleware/VoteWeigherBase.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/OwnableMock.sol";

import "forge-std/Test.sol";

contract VoteWeigherUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;
    IStrategyManager public strategyManager;

    address serviceManagerOwner;
    IServiceManager public serviceManager;

    IVoteWeigher public voteWeigher;

    IVoteWeigher public voteWeigherImplementation;

    address public pauser = address(555);
    address public unpauser = address(999);

    uint256 initialSupply = 1e36;
    address initialOwner = address(this);

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();

        pauserRegistry = new PauserRegistry(pauser, unpauser);

        strategyManager = new StrategyManagerMock();

        // make the serviceManagerOwner the owner of the serviceManager contract
        cheats.prank(serviceManagerOwner);
        serviceManager = IServiceManager(address(new OwnableMock()));

        voteWeigherImplementation = new VoteWeigherBase(strategyManager, serviceManager);

        voteWeigher = IVoteWeigher(address(new TransparentUpgradeableProxy(address(voteWeigherImplementation), address(proxyAdmin), "")));
    }

    function testCorrectConstructionParameters() public {
        assertEq(address(voteWeigherImplementation.strategyManager()), address(strategyManager));
        assertEq(address(voteWeigherImplementation.slasher()), address(strategyManager.slasher()));
        assertEq(address(voteWeigherImplementation.delegation()), address(strategyManager.delegation()));
        assertEq(address(voteWeigherImplementation.serviceManager()), address(serviceManager));
    }

    function testValidCreateQuorum() public {
        // create a quorum from the serviceManagerOwner
        cheats.prank(serviceManagerOwner);
        voteWeigher.createQuorum()
    }

}