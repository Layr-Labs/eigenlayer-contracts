// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/middleware/VoteWeigherBase.sol";
import "../mocks/ERC20Mock.sol";
import "../../contracts/middleware/StrategyWrapper.sol";
import "forge-std/Test.sol";


contract BLSPubkeyRegistrationUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    BLSPubkeyRegistry blsPubReg;
    ServiceManagerMock public serviceManagerMock;
    StrategyManagerMock public strategyManagerMock;
    StrategyWrapper public dummyStrat;
    ERC20Mock public dummyToken;


    function setUp() external {
        dummyToken = new ERC20Mock();
        dummyStrat = new StrategyWrapper(strategyManagerMock, dummyToken);

        blsPubReg = new BLSPubkeyRegistry(serviceManagerMock, strategManagerMock);

        uint96[] memory minimumStakeForQuorum = new uint96[](2);
        minimumStakeForQuorum[0] = 100;
        minimumStakeForQuorum[1] = 100;

        VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory quorumStrategiesConsideredAndMultipliers =
                new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](2);
        ethStratsAndMultipliers[0].strategy = dummyStrat;
        ethStratsAndMultipliers[0].multiplier = 1e18;
        eigenStratsAndMultipliers[1].strategy = dummyStrat;
        eigenStratsAndMultipliers[1].multiplier = 1e18;


        BLSPubkeyRegistry.initialize(minimumStakeForQuorum, quorumStrategiesConsideredAndMultipliers);
    }

    function testRegisterOperator() public {
        
    }
}