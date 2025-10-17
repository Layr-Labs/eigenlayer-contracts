// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

/**
 * @title AllocationManagerUpgrade
 * @notice Upgrade test for AllocationManager that verifies state consistency across upgrades
 */
contract Integration_Upgrade_AllocationManager is UpgradeTest {
    using ArrayLib for *;
    using StdStyle for *;

    function test_AllocationManagerUpgrade_StateConsistency() public {
        console.log("Testing AllocationManager upgrade...".green());

        // Test basic AllocationManager functionality before upgrade
        console.log("Pre-upgrade: Testing basic functionality...".green());

        // Just test that we can read basic AllocationManager state
        uint32 deallocationDelay = allocationManager.DEALLOCATION_DELAY();
        uint32 allocationConfigDelay = allocationManager.ALLOCATION_CONFIGURATION_DELAY();

        console.log("Pre-upgrade: DEALLOCATION_DELAY = %d", deallocationDelay);
        console.log("Pre-upgrade: ALLOCATION_CONFIGURATION_DELAY = %d", allocationConfigDelay);

        // Perform upgrade
        console.log("Performing upgrade...".green());
        _upgradeEigenLayerContracts();
        console.log("Upgrade completed".green());

        // Test basic functionality after upgrade
        console.log("Post-upgrade: Testing basic functionality...".green());

        // Verify the same values are still there
        uint32 deallocationDelayAfter = allocationManager.DEALLOCATION_DELAY();
        uint32 allocationConfigDelayAfter = allocationManager.ALLOCATION_CONFIGURATION_DELAY();

        console.log("Post-upgrade: DEALLOCATION_DELAY = %d", deallocationDelayAfter);
        console.log("Post-upgrade: ALLOCATION_CONFIGURATION_DELAY = %d", allocationConfigDelayAfter);

        // Verify they're the same
        assertEq(deallocationDelay, deallocationDelayAfter, "DEALLOCATION_DELAY should be the same after upgrade");
        assertEq(allocationConfigDelay, allocationConfigDelayAfter, "ALLOCATION_CONFIGURATION_DELAY should be the same after upgrade");

        console.log("AllocationManager upgrade test passed!".green().bold());
    }
}
