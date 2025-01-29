#!/bin/bash

# Default values
RPC_URL="https://eth.llamarpc.com"
ETHERSCAN_API_KEY="1234567890123456789012345678901234567890"
INPUT_FILE="contracts.json"
QUIET=false
TOTAL_ISSUES=0

# Help message
usage() {
    cat << EOF
Usage: bash $0 --rpc-url $RPC_URL --etherscan-key $ETHERSCAN_API_KEY --input $INPUT_FILE [--quiet] [--help]

Detects storage layout incompatibilities that could cause issues during upgrades.

Required:
    -r, --rpc-url <url>         RPC endpoint URL for the target network (default: https://eth.llamarpc.com).
    -e, --etherscan-key <key>   API key for Etherscan to fetch contract data.

Options:
    -i, --input <file>          JSON file containing contract details, see format below (default: contracts.json).
                                If not provided, reads from stdin.
    -q, --quiet                 Suppress informational output.
    -h, --help                  Show this help message.


Input JSON format:
{
  "contracts": [
    {
      "name": "AVSDirectory",
      "address": "0x135dda560e946695d6f155dacafc6f1f25c1f5af"
    },
    {
      "name": "DelegationManager",
      "address": "0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A"
    }
  ]
}
EOF
    exit 1
}

# Process command line arguments using a while loop and case statement.
while [[ $# -gt 0 ]]; do
    case $1 in
        -r|--rpc-url)
            RPC_URL="$2"
            shift 2
            ;;
        -e|--etherscan-key)
            ETHERSCAN_API_KEY="$2"
            shift 2
            ;;
        -i|--input)
            INPUT_FILE="$2"
            shift 2
            ;;
        -q|--quiet)
            QUIET=true
            shift
            ;;
        -h|--help)
            usage
            ;;
        *)
            echo "Unknown option: $1"
            usage
            ;;
    esac
done

# Validate required arguments
if [ -z "$RPC_URL" ] || [ -z "$ETHERSCAN_API_KEY" ]; then
    echo "Error: RPC URL and Etherscan API key are required"
    usage
fi

# Read JSON input
if [ -n "$INPUT_FILE" ]; then
    if [ ! -f "$INPUT_FILE" ]; then
        echo "Error: Input file not found: $INPUT_FILE"
        exit 1
    fi
    json_input=$(cat "$INPUT_FILE")
else
    json_input=$(cat)
fi

# Parse JSON values using jq
CONTRACTS=$(echo "$json_input" | jq -c '.contracts[]')

# Verify contracts are specified
if [ -z "$CONTRACTS" ]; then
    echo "Error: No contracts specified in JSON input"
    exit 1
fi

# Function to calculate number of slots a variable type occupies
calculate_slots() {
    local var_type=$1

    # Handle basic types
    case $var_type in
        *"uint256"*|*"int256"*|*"bytes32"*|*"address"*)
            echo 1
            ;;
        *"mapping"*)
            echo 1  # Mappings use 1 slot for the starting position
            ;;
        *"bytes"*|*"string"*)
            echo 1  # Dynamic types use 1 slot for length/pointer
            ;;
        *"array"*)
            echo 1  # Arrays use 1 slot for length/pointer, need to figure out how to parse slots consumed.
            ;;
        *)
            # Default to 1 slot if unknown
            echo 1
            ;;
    esac
}

# Function to analyze storage changes
analyze_storage_changes() {
    local onchain_file=$1
    local local_file=$2
    local contract_name=$3
    local issues_found=0

    # Get the storage layouts as arrays
    local onchain_slots=$(jq -r '.storage[] | "\(.slot)|\(.label)|\(.offset)|\(.type)"' "$onchain_file")
    local local_slots=$(jq -r '.storage[] | "\(.slot)|\(.label)|\(.offset)|\(.type)"' "$local_file")

    echo "Storage Layout Analysis for $contract_name:"
    echo "----------------------------------------"

    # Create temporary files for our data structures
    local onchain_map_file=$(mktemp)
    local local_map_file=$(mktemp)
    local processed_slots_file=$(mktemp)
    local renamed_vars_file=$(mktemp)

    # Parse onchain slots
    echo "$onchain_slots" | while IFS='|' read -r slot label offset type; do
        if [[ -n "$slot" ]]; then
            echo "${slot}|${label}|${offset}|${type}" >> "$onchain_map_file"
        fi
    done

    # Parse local slots
    echo "$local_slots" | while IFS='|' read -r slot label offset type; do
        if [[ -n "$slot" ]]; then
            echo "${slot}|${label}|${offset}|${type}" >> "$local_map_file"
        fi
    done

    # First pass: Check for renames (same slot, same type, different name)
    while IFS='|' read -r slot local_label local_offset local_type; do
        if [[ -z "$slot" ]]; then continue; fi

        # Look for matching slot in onchain
        onchain_line=$(grep "^${slot}|" "$onchain_map_file")
        if [[ -n "$onchain_line" ]]; then
            IFS='|' read -r _ onchain_label onchain_offset onchain_type <<< "$onchain_line"

            if [[ "$local_label" != "$onchain_label" && "$local_type" == "$onchain_type" && "$local_offset" == "$onchain_offset" ]]; then
                echo "${slot}|${onchain_label}|${local_label}|${local_type}" >> "$renamed_vars_file"
                echo "$slot" >> "$processed_slots_file"
                issues_found=$((issues_found + 1))
            fi
        fi
    done < "$local_map_file"

    # Print renames first
    while IFS='|' read -r slot old_name new_name type; do
        if [[ -n "$slot" ]]; then
            echo -e "\033[36m📝 Variable renamed at slot $slot:\033[0m"
            echo -e "\033[36m   $old_name -> $new_name ($type)\033[0m"
        fi
    done < "$renamed_vars_file"

    # Analyze other differences
    while IFS='|' read -r slot local_label local_offset local_type; do
        if [[ -z "$slot" ]]; then continue; fi

        # Skip if this slot was processed as a rename
        if grep -q "^${slot}$" "$processed_slots_file"; then
            continue
        fi

        # Look for matching slot in onchain
        onchain_line=$(grep "^${slot}|" "$onchain_map_file")
        if [[ -z "$onchain_line" ]]; then
            # New variable added
            slots_needed=$(calculate_slots "$local_type")
            echo -e "\033[32m✨ New variable added: $local_label ($local_type) at slot $slot\033[0m"
            issues_found=$((issues_found + 1))
            if [ "$slots_needed" -gt 1 ]; then
                echo -e "\033[33m   📦 This variable occupies $slots_needed slots\033[0m"
            fi
        else
            IFS='|' read -r _ onchain_label onchain_offset onchain_type <<< "$onchain_line"

            if [[ "$local_label" != "$onchain_label" ]]; then
                echo -e "\033[31m🚨 Storage slot override detected at slot $slot:\033[0m"
                echo -e "\033[31m   Previous: $onchain_label ($onchain_type)\033[0m"
                echo -e "\033[32m   New: $local_label ($local_type)\033[0m"
                issues_found=$((issues_found + 1))

                # Calculate potential impact
                old_slots=$(calculate_slots "$onchain_type")
                new_slots=$(calculate_slots "$local_type")
                slot_diff=$((new_slots - old_slots))

                if [ "$slot_diff" -gt 0 ]; then
                    echo -e "\033[33m   ⚠️ This change will shift subsequent storage slots by +$slot_diff positions\033[0m"
                elif [ "$slot_diff" -lt 0 ]; then
                    echo -e "\033[33m   💡 This change will reduce storage usage by $((slot_diff * -1)) slots\033[0m"
                fi
            elif [[ "$local_type" != "$onchain_type" ]]; then
                echo -e "\033[33m🔄 Type change detected for $local_label at slot $slot:\033[0m"
                echo -e "\033[31m   Previous: $onchain_type\033[0m"
                echo -e "\033[32m   New: $local_type\033[0m"
                issues_found=$((issues_found + 1))
            fi
        fi
        echo "$slot" >> "$processed_slots_file"
    done < "$local_map_file"

    # Check for removed variables
    while IFS='|' read -r slot onchain_label onchain_offset onchain_type; do
        if [[ -z "$slot" ]]; then continue; fi

        # Skip if this slot was processed as a rename or already handled
        if grep -q "^${slot}$" "$processed_slots_file"; then
            continue
        fi

        # Look for matching slot in local
        if ! grep -q "^${slot}|" "$local_map_file"; then
            echo -e "\033[31m➖ Variable removed: $onchain_label ($onchain_type) from slot $slot\033[0m"
            issues_found=$((issues_found + 1))
        fi
    done < "$onchain_map_file"

    # Cleanup temporary files
    rm -f "$onchain_map_file" "$local_map_file" "$processed_slots_file" "$renamed_vars_file"

    echo "Issues found in $contract_name: $issues_found"
    return $issues_found
}

# Function to process a single contract
process_contract() {
    local contract_name=$1
    local contract_address=$2
    local issues_found=0

    # Create directories for storing layouts and diffs
    mkdir -p "storage-report/layouts" "storage-report/diffs"
    local local_file="storage-report/layouts/${contract_name}_local.json"
    local onchain_file="storage-report/layouts/${contract_name}_onchain.json"
    local diff_file="storage-report/diffs/${contract_name}.diff"

    # Generate storage layouts
    if ! forge inspect "$contract_name" storage --json > "$local_file" 2>/dev/null; then
        echo "Error: forge inspect failed for contract: $contract_name"
        return 1
    fi

    if ! cast storage "$contract_address" --rpc-url "$RPC_URL" --etherscan-api-key "$ETHERSCAN_API_KEY" --json > "$onchain_file" 2>/dev/null; then
        echo "Error: cast storage failed for address: $contract_address"
        return 1
    fi

    # Delete the first line of $onchain_file
    sed '1d' "$onchain_file" > "$onchain_file.tmp" && mv "$onchain_file.tmp" "$onchain_file"

    # Filter out astId and contract fields from local and onchain files, and normalize type identifiers
    jq 'del(.types, .values, .storage[].astId, .storage[].contract) | .storage[].type |= gsub("\\([^)]+\\)[0-9]+"; "")' "$local_file" > "${local_file}.tmp" && mv "${local_file}.tmp" "$local_file"
    jq 'del(.types, .values, .storage[].astId, .storage[].contract) | .storage[].type |= gsub("\\([^)]+\\)[0-9]+"; "")' "$onchain_file" > "${onchain_file}.tmp" && mv "${onchain_file}.tmp" "$onchain_file"

    if [ "$QUIET" = false ]; then
        echo "----------------------------------------"
        echo "Local contract:  $contract_name"
        echo "Chain address:   $contract_address"
        echo "Network RPC:     $RPC_URL"
        echo "JSON files stored in: storage-report/layouts/"
        echo "Diffs stored in: storage-report/diffs/"
        echo "----------------------------------------"
    fi

    # Generate and store diff
    diff -u "$onchain_file" "$local_file" > "$diff_file" 2>/dev/null || true

    # Analyze storage changes
    analyze_storage_changes "$onchain_file" "$local_file" "$contract_name"
    issues_found=$?

    return $issues_found
}

# Process each contract from the JSON input
while IFS= read -r contract; do
    contract_name=$(echo "$contract" | jq -r '.name')
    contract_address=$(echo "$contract" | jq -r '.address')

    if [ -z "$contract_name" ] || [ -z "$contract_address" ]; then
        echo "Error: Each contract must specify both name and address"
        continue
    fi

    if [ "$QUIET" = false ]; then
        echo "Processing contract: $contract_name at $contract_address"
    fi
    process_contract "$contract_name" "$contract_address"
    TOTAL_ISSUES=$((TOTAL_ISSUES + $?))
done <<< "$CONTRACTS"

if [ "$TOTAL_ISSUES" -gt 0 ]; then
    echo -e "\n\033[31m🚨 Total storage layout issues found: $TOTAL_ISSUES\033[0m"
    exit 1
else
    echo -e "\n\033[32m✅ No storage layout issues found\033[0m"
    exit 0
fi
