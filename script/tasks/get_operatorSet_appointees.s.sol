// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/interfaces/IPermissionController.sol";
import "../../src/contracts/interfaces/IAllocationManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

/// @title GetOperatorSetAppointees
/// @notice Queries PermissionController to find appointees for slashOperator across operator sets
/// @dev Usage:
///   1. Fetch operator sets from API and save to JSON:
///      curl -H "Accept: application/json" \
///        "https://sidecar-rpc.eigenlayer.xyz/mainnet/v1/operatorSets" \
///        > script/input/operatorSets.json
///
///   2. Run the script:
///      forge script script/tasks/get_operatorSet_appointees.s.sol \
///        --rpc-url $RPC_MAINNET \
///        --sig "run(string memory)" \
///        -- script/input/operatorSets.json

/// @notice As of 1/27 there are no appointees. You can sanity check that there is at least one for a different selector with the following command:
/// cast call 0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5 "getAppointees(address,address,bytes4)(address[])" 0xF07F83Ff977Dd004060F00ecefB80A9F92775098 0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39 0x36352057 --rpc-url $RPC_MAINNET
contract GetOperatorSetAppointees is Script, Test {
    // Mainnet addresses
    address constant PERMISSION_CONTROLLER = 0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5;
    address constant ALLOCATION_MANAGER = 0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39;

    IPermissionController permissionController = IPermissionController(PERMISSION_CONTROLLER);

    function run(
        string memory inputFile
    ) public view {
        // Read JSON file
        string memory json = vm.readFile(inputFile);

        // Parse operator sets from JSON
        // Expected format: { "operatorSets": [{ "id": 0, "avs": "0x..." }, ...] }
        bytes memory operatorSetsRaw = vm.parseJson(json, ".operatorSets");

        // Decode into array of OperatorSetData
        OperatorSetData[] memory operatorSets = abi.decode(operatorSetsRaw, (OperatorSetData[]));

        console.log("Total operator sets in input:", operatorSets.length);
        console.log("---");

        // Get unique AVS addresses
        address[] memory uniqueAvs = _getUniqueAvs(operatorSets);
        console.log("Unique AVS addresses:", uniqueAvs.length);
        console.log("---");

        // Get the slashOperator selector
        // slashOperator(address,SlashingParams) => 0x38c29160
        bytes4 slashOperatorSelector = IAllocationManagerActions.slashOperator.selector;

        uint256 totalAppointees = 0;
        uint256 avsWithAppointees = 0;

        // Query appointees for each AVS
        for (uint256 i = 0; i < uniqueAvs.length; i++) {
            address avs = uniqueAvs[i];

            address[] memory appointees =
                permissionController.getAppointees(avs, ALLOCATION_MANAGER, slashOperatorSelector);

            if (appointees.length > 0) {
                avsWithAppointees++;
                totalAppointees += appointees.length;

                console.log("AVS:", avs);
                console.log("  Appointees:", appointees.length);

                for (uint256 j = 0; j < appointees.length; j++) {
                    console.log("    -", appointees[j]);
                }
            }
        }

        console.log("---");
        console.log("Summary:");
        console.log("  Total operator sets:", operatorSets.length);
        console.log("  Unique AVS addresses:", uniqueAvs.length);
        console.log("  AVS with slasher appointees:", avsWithAppointees);
        console.log("  Total slasher appointees:", totalAppointees);
    }

    /// @notice Struct to decode operator set data from JSON
    struct OperatorSetData {
        uint256 id;
        address avs;
    }

    /// @notice Extract unique AVS addresses from operator sets
    function _getUniqueAvs(
        OperatorSetData[] memory operatorSets
    ) internal pure returns (address[] memory) {
        // First pass: count unique addresses
        address[] memory temp = new address[](operatorSets.length);
        uint256 uniqueCount = 0;

        for (uint256 i = 0; i < operatorSets.length; i++) {
            address avs = operatorSets[i].avs;
            bool found = false;

            for (uint256 j = 0; j < uniqueCount; j++) {
                if (temp[j] == avs) {
                    found = true;
                    break;
                }
            }

            if (!found) {
                temp[uniqueCount] = avs;
                uniqueCount++;
            }
        }

        // Second pass: copy to correctly sized array
        address[] memory result = new address[](uniqueCount);
        for (uint256 i = 0; i < uniqueCount; i++) {
            result[i] = temp[i];
        }

        return result;
    }
}
