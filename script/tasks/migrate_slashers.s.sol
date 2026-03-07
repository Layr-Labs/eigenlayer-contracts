// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/interfaces/IAllocationManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

/// @title MigrateSlashers
/// @notice Migrates slashers from PermissionController to AllocationManager for operator sets created prior to v1.9.0
/// @dev Usage:
///   1. Fetch operator sets from API and save to JSON:
///      curl -H "Accept: application/json" \
///        "https://sidecar-rpc.eigenlayer.xyz/mainnet/v1/operatorSets" \
///        > script/input/operatorSets.json
///
///   2. Run the script (dry run):
///      forge script script/tasks/migrate_slashers.s.sol \
///        --rpc-url $RPC_MAINNET \
///        --sig "run(string memory)" \
///        -- script/input/operatorSets.json
///
///   3. Run the script (broadcast):
///      forge script script/tasks/migrate_slashers.s.sol \
///        --rpc-url $RPC_MAINNET \
///        --sig "run(string memory)" \
///        --broadcast \
///        --private-key $PRIVATE_KEY \
///        -- script/input/operatorSets.json
contract MigrateSlashers is Script, Test {
    // Mainnet addresses
    address constant ALLOCATION_MANAGER = 0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39;

    uint256 constant BATCH_SIZE = 20;

    IAllocationManager allocationManager = IAllocationManager(ALLOCATION_MANAGER);

    /// @notice Struct to decode operator set data from JSON
    struct OperatorSetData {
        uint256 id;
        address avs;
    }

    function run(
        string memory inputFile
    ) public {
        // Read JSON file
        string memory json = vm.readFile(inputFile);

        // Parse operator sets from JSON
        // Expected format: { "operatorSets": [{ "id": 0, "avs": "0x..." }, ...] }
        bytes memory operatorSetsRaw = vm.parseJson(json, ".operatorSets");

        // Decode into array of OperatorSetData
        OperatorSetData[] memory operatorSetData = abi.decode(operatorSetsRaw, (OperatorSetData[]));

        console.log("Total operator sets:", operatorSetData.length);

        // Convert to OperatorSet[] (casting id from uint256 to uint32)
        OperatorSet[] memory operatorSets = new OperatorSet[](operatorSetData.length);
        for (uint256 i = 0; i < operatorSetData.length; i++) {
            operatorSets[i] = OperatorSet({avs: operatorSetData[i].avs, id: uint32(operatorSetData[i].id)});
        }

        // Calculate number of batches
        uint256 numBatches = (operatorSets.length + BATCH_SIZE - 1) / BATCH_SIZE;
        console.log("Number of batches:", numBatches);
        console.log("---");

        vm.startBroadcast();

        // Process in batches
        for (uint256 i = 0; i < numBatches; i++) {
            uint256 start = i * BATCH_SIZE;
            uint256 end = _min(start + BATCH_SIZE, operatorSets.length);
            uint256 batchLength = end - start;

            // Create batch array
            OperatorSet[] memory batch = new OperatorSet[](batchLength);
            for (uint256 j = 0; j < batchLength; j++) {
                batch[j] = operatorSets[start + j];
            }

            console.log("Processing batch", i + 1, "of", numBatches);
            console.log("  Operator sets", start, "to", end - 1);

            // Call migrateSlashers for this batch
            allocationManager.migrateSlashers(batch);

            console.log("  Batch complete");
        }

        vm.stopBroadcast();

        console.log("---");
        console.log("Migration complete");
        console.log("  Total operator sets processed:", operatorSets.length);
        console.log("  Total batches:", numBatches);
    }

    function _min(
        uint256 a,
        uint256 b
    ) internal pure returns (uint256) {
        return a < b ? a : b;
    }
}
