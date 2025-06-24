// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/SlashingLib.sol";

contract SlashingLibUnitTests is Test {
    using SlashingLib for *;
    
    string csvHeader = "scaled_shares,prev_max_mag,new_max_mag,slashed_amount";
    string outputFile;
    
    function setUp() public {
        outputFile = vm.envOr("CSV_OUTPUT_FILE", string("script/utils/queuedShares/slashing_test_output.csv"));
        vm.writeFile(outputFile, "");
        vm.writeLine(outputFile, csvHeader);
        
        console.log("Writing CSV to:", outputFile);
    }
    
    function testFuzz_scaleForBurning(
        uint256 scaledShares, 
        uint64 prevMaxMag, 
        uint64 newMaxMag
    ) public {
        vm.assume(scaledShares > 0 && scaledShares <= type(uint128).max);
        vm.assume(prevMaxMag <= WAD && prevMaxMag > 0);
        vm.assume(newMaxMag < prevMaxMag);
        
        uint256 slashedAmount = scaledShares.scaleForBurning(
            prevMaxMag, 
            newMaxMag
        );
        
        vm.writeLine(outputFile, string.concat(
            vm.toString(scaledShares), ",",
            vm.toString(prevMaxMag), ",",
            vm.toString(newMaxMag), ",",
            vm.toString(slashedAmount)
        ));
    }
}