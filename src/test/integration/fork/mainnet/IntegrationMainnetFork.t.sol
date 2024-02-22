// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/User.t.sol";

contract IntegrationMainnetFork is IntegrationCheckUtils {
    // Select mainnet fork id
    function setUp() public override {
        cheats.selectFork(mainnetForkId);
        
        string memory deploymentInfoPath = "script/config/mainnet/Mainnet_current_deploy.config.json";
        _parseDeployedContracts(deploymentInfoPath);
        _upgradeMainnetContracts();
    }
}