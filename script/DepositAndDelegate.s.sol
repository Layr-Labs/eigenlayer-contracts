// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./EigenLayerParser.sol";

contract DepositAndDelegate is Script, DSTest, EigenLayerParser {

    //performs basic deployment before each test
    function run() external {
        parseEigenLayerParams();

        uint256 wethAmount = eigenTotalSupply / (numStaker + numDis + 50); // save 100 portions

        address dlnAddr;

        //get the corresponding dln
        //is there an easier way to do this?
        for (uint256 i = 0; i < numStaker; i++) {
            address stakerAddr =
                stdJson.readAddress(configJson, string.concat(".staker[", string.concat(vm.toString(i), "].address")));
            if (stakerAddr == msg.sender) {
                dlnAddr =
                    stdJson.readAddress(configJson, string.concat(".dln[", string.concat(vm.toString(i), "].address")));
            }
        }

        vm.startBroadcast(msg.sender);
        eigen.approve(address(strategyManager), wethAmount);
        strategyManager.depositIntoStrategy(eigenStrat, eigen, wethAmount);
        weth.approve(address(strategyManager), wethAmount);
        strategyManager.depositIntoStrategy(wethStrat, weth, wethAmount);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(dlnAddr, signatureWithExpiry);
        vm.stopBroadcast();
    }
}
