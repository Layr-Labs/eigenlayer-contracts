// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/User.t.sol";

contract IntegrationMainnetFork is IntegrationCheckUtils {
    // Select mainnet fork id
    function setUp() public override {
        // create mainnet fork that can be used later
        mainnetForkId = cheats.createSelectFork(cheats.rpcUrl("mainnet"));

        string memory deploymentInfoPath = "script/configs/mainnet/Mainnet_current_deploy.config.json";
        _parseDeployedContracts(deploymentInfoPath);
        _upgradeMainnetContracts();
    }
}
