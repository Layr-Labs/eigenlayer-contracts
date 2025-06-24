// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/SlashingLib.sol";

contract SlashingLibUnitTests is Test {
    using SlashingLib for *;

    function test_mulWad() public {
        assertEq(SlashingLib.mulWad(1e18, 5e17), 5e17);
        assertEq(SlashingLib.mulWad(2e18, 3e17), 6e17);
    }

    function test_mulWadRoundUp() public {
        // Test rounding up behavior
        assertGt(SlashingLib.mulWadRoundUp(1, 1), 0);
    }

    function test_calcSlashedAmount() public {
        uint256 operatorShares = 1000e18;
        uint64 prevMaxMagnitude = 1e18;
        uint64 newMaxMagnitude = 5e17; // 50% slash
        
        uint256 slashed = SlashingLib.calcSlashedAmount(
            operatorShares, 
            prevMaxMagnitude, 
            newMaxMagnitude
        );
        
        assertEq(slashed, 500e18); // Should slash 50%
    }

    function testFuzz_calcSlashedAmount(
        uint256 operatorShares,
        uint64 prevMaxMagnitude,
        uint64 newMaxMagnitude
    ) public {
        vm.assume(operatorShares <= type(uint128).max);
        vm.assume(prevMaxMagnitude > 0);
        vm.assume(newMaxMagnitude <= prevMaxMagnitude);
        
        uint256 slashed = SlashingLib.calcSlashedAmount(
            operatorShares, 
            prevMaxMagnitude, 
            newMaxMagnitude
        );
        
        assertLe(slashed, operatorShares);
    }
}