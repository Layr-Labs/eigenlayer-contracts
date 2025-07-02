// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/SlashingLib.sol";

contract SlashingLibHarness {
    DepositScalingFactor dsf;

    function unsafeUpdate(uint scalingFactor) public {
        dsf._scalingFactor = scalingFactor;
    }

    function update(uint prevDepositShares, uint addedShares, uint slashingFactor) public {
        dsf.update(prevDepositShares, addedShares, slashingFactor);
    }
}

contract SlashingLibUnitTests is Test {
    /// @dev We use a harness so that `vm.expectRevert()` can be used.
    SlashingLibHarness harness;

    function setUp() public {
        harness = new SlashingLibHarness();
    }

    function test_Revert_InvalidDepositScalingFactor() public {
        harness.unsafeUpdate(1);

        vm.expectRevert(SlashingLib.InvalidDepositScalingFactor.selector);
        harness.update({prevDepositShares: 1e18, addedShares: 1, slashingFactor: WAD - 1});
    }
}
