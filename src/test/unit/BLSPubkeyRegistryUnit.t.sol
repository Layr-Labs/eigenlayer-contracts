// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/middleware/VoteWeigherBase.sol";
import "../mocks/ERC20Mock.sol";
import "../mocks/ServiceManagerMock.sol";
import "../mocks/StrategyManagerMock.sol";
import "../mocks/SlasherMock.sol";
import "../../contracts/strategies/StrategyWrapper.sol";
import "forge-std/Test.sol";


contract BLSPubkeyRegistrationUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    BLSPubkeyRegistry blsPubReg;
    ServiceManagerMock public serviceManagerMock;
    StrategyManagerMock public strategyManagerMock;
    StrategyWrapper public dummyStrat;
    SlasherMock public slasherMock;
    ERC20Mock public dummyToken;


    function setUp() external {
        dummyToken = new ERC20Mock();
        dummyStrat = new StrategyWrapper(strategyManagerMock, dummyToken);
        slasherMock = new SlasherMock();
        strategyManagerMock = new StrategyManagerMock();
        serviceManagerMock = new ServiceManagerMock(slasherMock);

        blsPubReg = new BLSPubkeyRegistry(strategyManagerMock, serviceManagerMock);

        uint96[] memory minimumStakeForQuorum = new uint96[](2);
        minimumStakeForQuorum[0] = 100;
        minimumStakeForQuorum[1] = 100;

        VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[][] memory quorumStrategiesConsideredAndMultipliers =
                new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[][](2);
        quorumStrategiesConsideredAndMultipliers[0][0].strategy = dummyStrat;
        quorumStrategiesConsideredAndMultipliers[0][0].multiplier = 1e18;
        quorumStrategiesConsideredAndMultipliers[1][0].strategy = dummyStrat;
        quorumStrategiesConsideredAndMultipliers[1][0].multiplier = 1e18;


        blsPubReg.initialize(minimumStakeForQuorum, quorumStrategiesConsideredAndMultipliers);
    }

    function testRegisterOperator() public {

    }
}