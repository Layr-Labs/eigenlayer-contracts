// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/fork/goerli/IntegrationGoerliFork.t.sol";
import "src/test/integration/User.t.sol";

contract IntegrationGoerliFork_UpgradeSetup is IntegrationGoerliFork {

    /// @notice Test upgrade setup is correct
    function test_goerliUpgradeSetup(uint24 _random) public {
        _verifyContractPointers();
        _verifyImplementationsContracts();
        _verifyInitialOwners();
    }
}
