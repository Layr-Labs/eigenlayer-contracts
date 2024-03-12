// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

contract M2_Deploy_Holesky_From_Scratch is ExistingDeploymentParser {
    function run() external {
        _parseInitialDeploymentParams("script/configs/holesky/M2_deploy_from_scratch.holesky.config.json");

        // // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        // vm.startBroadcast();
        // _deployFromScratch();

        // // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        // vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();

        logContractAddresses();
    }

    function _deployFromScratch() internal {}
}
