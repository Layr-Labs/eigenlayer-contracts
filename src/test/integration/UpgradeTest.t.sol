// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationDeployer.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

abstract contract UpgradeTest is IntegrationCheckUtils {
    /// Only run upgrade tests on mainnet forks
    function setUp() public virtual override {
        if (!isForktest()) {
            cheats.skip(true);
        } else {
            isUpgraded = false;
            super.setUp();
        }
    }

    /// Deploy current implementation contracts and upgrade existing proxies
    function _upgradeEigenLayerContracts() public virtual {
        require(forkType == MAINNET, "_upgradeEigenLayerContracts: somehow running upgrade test locally");
        require(!isUpgraded, "_upgradeEigenLayerContracts: already performed upgrade");

        emit log("_upgradeEigenLayerContracts: upgrading mainnet to redistribution");

        _upgradeMainnetContracts();

        emit log("======");

        isUpgraded = true;
        emit log("_upgradeEigenLayerContracts: redistribution upgrade complete");
    }
}
