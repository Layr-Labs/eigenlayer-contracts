// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";

contract IntegrationMainnetFork_UpgradeSetup is IntegrationCheckUtils {

    // /// @notice Test upgrade setup is correct
    // function test_mainnetUpgradeSetup(uint24 _random) public {
    //     _configRand({
    //         _randomSeed: _random,
    //         _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
    //         _userTypes: DEFAULT | ALT_METHODS,
    //         _forkTypes: MAINNET
    //     });

    //     // 1. Upgrade mainnet contracts
    //     _upgradeEigenLayerContracts();

    //     // 2. Verify contract pointers and setup
    //     _verifyContractPointers();
    //     _verifyImplementationsContracts();
    //     _verifyInitialOwners();
    // }

}
