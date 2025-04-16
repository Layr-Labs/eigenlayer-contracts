// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract EigenPodTest is IntegrationChecks {
    function _init() internal virtual override {
        string memory chain = cheats.envString("FORK_CHAIN");
        if (eq(chain, "sepolia") || eq(chain, "hoodi")) cheats.skip(true);
    }
}
