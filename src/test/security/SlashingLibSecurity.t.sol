// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.27;

import {Test} from "forge-std/Test.sol";
import {SlashingLib} from "../../contracts/libraries/SlashingLib.sol";
import {DepositScalingFactor} from "../../contracts/libraries/SlashingLib.sol";

contract SlashingLibSecurityTest is Test {
    uint256 constant WAD = 1e18;

    function testFuzz_SlashingLib_ExtremeValues(uint256 depositShares, uint256 slashingFactor) public {
        // Bound inputs to realistic but extreme ranges
        depositShares = bound(depositShares, 1, type(uint128).max);
        slashingFactor = bound(slashingFactor, 1, WAD);
        
        // Test calcWithdrawable with extreme values
        uint256 withdrawableShares = SlashingLib.calcWithdrawable(
            DepositScalingFactor(WAD), 
            depositShares, 
            slashingFactor
        );
        
        // Verify result is within expected bounds
        assertLe(withdrawableShares, depositShares);
    }
}
