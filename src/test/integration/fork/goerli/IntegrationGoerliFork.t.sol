// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/User.t.sol";

contract IntegrationGoerliFork is IntegrationCheckUtils {
    // Select mainnet fork id
    function setUp() public override {
        // create mainnet fork that can be used later
        mainnetForkId = cheats.createSelectFork(cheats.rpcUrl("goerli"));

        string memory deploymentInfoPath = "script/configs/goerli/Goerli_current_deploy.config.json";
        _parseDeployedContracts(deploymentInfoPath);
        _upgradeGoerliContracts();
    }
}
