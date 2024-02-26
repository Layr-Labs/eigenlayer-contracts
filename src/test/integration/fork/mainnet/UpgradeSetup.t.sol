// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/fork/mainnet/IntegrationMainnetFork.t.sol";
import "src/test/integration/User.t.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IEigenPod.sol";

contract IntegrationMainnetFork_UpgradeSetup is IntegrationMainnetFork {

    /// @notice Test upgrade setup is correct
    function test_mainnetUpgradeSetup(uint24 _random) public {
        _verifyContractPointers();
        _verifyImplementationsContracts();
        _verifyInitialOwners();
    }

}
