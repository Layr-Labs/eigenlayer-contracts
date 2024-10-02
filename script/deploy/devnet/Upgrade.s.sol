// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "../../utils/ExistingDeploymentParser.sol";

// # To load the variables in the .env file
// source .env

// Generic upgrade script, DOES NOT UPDATE IMPLEMENTATION IN OUTPUT FILE
// forge script script/deploy/devnet/Upgrade.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY
contract Upgrade is ExistingDeploymentParser {

    function run() external {
        // EDIT this for your script
        _parseDeployedContracts("script/output/holesky/pre_preprod_slashing.holesky.json");

        vm.startBroadcast();
        AVSDirectory newAVSDirectoryImplementation = new AVSDirectory(delegationManager, eigenLayerPauserReg);
        eigenLayerProxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(avsDirectory))), address(newAVSDirectoryImplementation));
        vm.stopBroadcast();

    }
}