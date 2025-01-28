#!/bin/bash

# Default values
RPC_URL="https://eth.llamarpc.com"
ETHERSCAN_API_KEY=""
INPUT_FILE="solari.json"
QUIET=false

# Help message
usage() {
    cat << EOF
Usage: bash $0 --rpc-url=$RPC_URL --etherscan-key=$ETHERSCAN_API_KEY [--input=$INPUT_FILE] [--quiet] [--help]

Detects storage layout incompatibilities that could cause issues during upgrades.

Required:
    -r, --rpc-url <url>         RPC endpoint URL for the target network (e.g. https://eth.llamarpc.com)
    -e, --etherscan-key <key>   API key for Etherscan to fetch contract data

Options:
    -i, --input <file>          JSON file containing contract details (see format below)
                                If not provided, reads from stdin
    -q, --quiet                 Suppress informational output
    -h, --help                  Show this help message


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
# This allows for both short (-r) and long (--rpc-url) argument formats.
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

# Function to process a single contract
process_contract() {
    local contract_name=$1
    local contract_address=$2
    
    # Temporary files
    local local_file=$(mktemp)
    local onchain_file=$(mktemp)

    # Generate storage layouts
    if ! forge inspect "$contract_name" storage --json > "$local_file" 2>/dev/null; then
        echo "Error: forge inspect failed for contract: $contract_name"
        rm "$local_file" "$onchain_file"
        return 1
    fi

    if ! cast storage "$contract_address" --rpc-url "$RPC_URL" --etherscan-api-key "$ETHERSCAN_API_KEY" --json > "$onchain_file" 2>/dev/null; then
        echo "Error: cast storage failed for address: $contract_address"
        rm "$local_file" "$onchain_file"
        return 1
    fi

    # Delete the first line of $onchain_file
    sed '1d' "$onchain_file" > "$onchain_file.tmp" && mv "$onchain_file.tmp" "$onchain_file"

    # Filter out astId and contract fields from local and onchain files
    jq 'del(.types, .values, .storage[].astId, .storage[].contract, .storage[].type)' "$local_file" > "${local_file}.tmp" && mv "${local_file}.tmp" "$local_file"
    jq 'del(.types, .values, .storage[].astId, .storage[].contract, .storage[].type)' "$onchain_file" > "${onchain_file}.tmp" && mv "${onchain_file}.tmp" "$onchain_file"

    # Compare storage layouts and provide feedback
    if ! diff "$onchain_file" "$local_file" > "${contract_name}.diff"; then
        echo "Storage Layout Differences Detected for $contract_name!"
        if [ "$QUIET" = false ]; then
            echo "----------------------------------------"
            echo "Local contract:  $contract_name"
            echo "Chain address:   $contract_address" 
            echo "Network RPC:     $RPC_URL"
            echo "----------------------------------------"
            # Format diff output with colors
            while IFS= read -r line; do
                if [[ $line == \<* ]]; then
                    echo -e "\033[31m$line\033[0m"  # Red for removed lines
                elif [[ $line == \>* ]]; then
                    echo -e "\033[32m$line\033[0m"  # Green for added lines
                elif [[ $line == ---* ]] || [[ $line == [0-9]* ]]; then
                    echo -e "\033[36m$line\033[0m"  # Cyan for metadata lines
                else
                    echo "   $line"
                fi
            done < "${contract_name}.diff"
        fi
    else
        echo "Storage layouts match for $contract_name!"
        rm -f "${contract_name}.diff"
    fi

    # Cleanup
    rm -f "$local_file" "$onchain_file"
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
done <<< "$CONTRACTS"
