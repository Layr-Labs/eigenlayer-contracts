// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../../contracts/libraries/Endian.sol";

contract TestOutput is Test {
    function test_BalanceRootOutput() public view {
        // Define balances in wei (1 ETH = 1e18 wei)
        uint64 val0Balance = 32e9;
        uint64 val1Balance = 48e9;
        uint64 val2Balance = 64e9;
        uint64 val3Balance = 96e9;

        // Calculate the balance root
        bytes32 balanceRoot = bytes32(
            uint256(toLittleEndianUint64(val0Balance)) |
            (uint256(toLittleEndianUint64(val1Balance)) >> 64) |
            (uint256(toLittleEndianUint64(val2Balance)) >> 128) |
            (uint256(toLittleEndianUint64(val3Balance)) >> 192)
        );

        // Print the output
        console.log("Balance Root (bytes32):");
        console.logBytes32(balanceRoot);
        
        // Also print individual balances for reference
        console.log("\nIndividual balances:");
        console.log("val0Balance: %s Gwei (%s wei)", val0Balance / 1e9, val0Balance);
        console.log("val1Balance: %s Gwei (%s wei)", val1Balance / 1e9, val1Balance);
        console.log("val2Balance: %s Gwei (%s wei)", val2Balance / 1e9, val2Balance);
        console.log("val3Balance: %s Gwei (%s wei)", val3Balance / 1e9, val3Balance);
        
        // Print the hex values of each balance for clarity
        console.log("\nBalances in hex:");
        console.log("val0Balance: 0x%x", val0Balance);
        console.log("val1Balance: 0x%x", val1Balance);
        console.log("val2Balance: 0x%x", val2Balance);
        console.log("val3Balance: 0x%x", val3Balance);
        
        // Show the shifted values
        console.log("\nShifted values:");
        console.log("val0Balance (no shift): 0x%x", uint256(val0Balance));
        console.log("val1Balance >> 64: 0x%x", uint256(val1Balance) >> 64);
        console.log("val2Balance >> 128: 0x%x", uint256(val2Balance) >> 128);
        console.log("val3Balance >> 192: 0x%x", uint256(val3Balance) >> 192);
    }

    /// @dev Opposite of Endian.fromLittleEndianUint64
    function toLittleEndianUint64(uint64 num) internal pure returns (bytes32) {
        uint lenum;

        // Rearrange the bytes from big-endian to little-endian format
        lenum |= uint((num & 0xFF) << 56);
        lenum |= uint((num & 0xFF00) << 40);
        lenum |= uint((num & 0xFF0000) << 24);
        lenum |= uint((num & 0xFF000000) << 8);
        lenum |= uint((num & 0xFF00000000) >> 8);
        lenum |= uint((num & 0xFF0000000000) >> 24);
        lenum |= uint((num & 0xFF000000000000) >> 40);
        lenum |= uint((num & 0xFF00000000000000) >> 56);

        // Shift the little-endian bytes to the end of the bytes32 value
        return bytes32(lenum << 192);
    }
}
