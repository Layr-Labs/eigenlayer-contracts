// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../src/contracts/interfaces/IDelegationManager.sol";
import "../src/contracts/core/DelegationManager.sol";

import "../src/contracts/core/StrategyManager.sol";
import "../src/contracts/strategies/StrategyBase.sol";
import "../src/contracts/core/Slasher.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

contract EigenLayerParser is Script, DSTest {
    Vm cheats = Vm(HEVM_ADDRESS);

    uint256 numDis;
    uint256 numDln;
    uint256 numStaker;
    uint256 numCha;

    uint256 public constant eigenTotalSupply = 1000e18;
    DelegationManager public delegation;
    StrategyManager public strategyManager;
    IERC20 public weth;
    StrategyBase public wethStrat;
    IERC20 public eigen;
    StrategyBase public eigenStrat;

    string internal configJson;
    string internal addressJson;

    function parseEigenLayerParams() internal {
        configJson = vm.readFile("data/participants.json");
        numDis = stdJson.readUint(configJson, ".numDis");
        numDln = stdJson.readUint(configJson, ".numDln");
        numStaker = stdJson.readUint(configJson, ".numStaker");

        addressJson = vm.readFile("data/addresses.json");
        delegation = DelegationManager(stdJson.readAddress(addressJson, ".delegation"));
        strategyManager = StrategyManager(stdJson.readAddress(addressJson, ".strategyManager"));
        weth = IERC20(stdJson.readAddress(addressJson, ".weth"));
        wethStrat = StrategyBase(stdJson.readAddress(addressJson, ".wethStrat"));
        eigen = IERC20(stdJson.readAddress(addressJson, ".eigen"));
        eigenStrat = StrategyBase(stdJson.readAddress(addressJson, ".eigenStrat"));
    }
}
