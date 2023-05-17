#!/bin/bash

# Check if yq is installed
if ! command -v yq &> /dev/null
then
    echo "yq is not installed. Please install it and try again."
    return 1
fi

# Check for arguments
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 [goerli|local]"
    return 1
fi

# Read the YAML file
CONFIG_FILE="config.yml"

case $1 in
    goerli)
        CHAIN_ID=$(yq e '.goerli.CHAIN_ID' $CONFIG_FILE)
        EXECUTOR_MULTISIG=$(yq e '.goerli.EXECUTOR_MULTISIG' $CONFIG_FILE)
        FOUNDRY_FUZZ_RUNS=$(yq e '.goerli.FOUNDRY_FUZZ_RUNS' $CONFIG_FILE)
        ;;
    local)
        CHAIN_ID=$(yq e '.local.CHAIN_ID' $CONFIG_FILE)
        FOUNDRY_FUZZ_RUNS=$(yq e '.local.FOUNDRY_FUZZ_RUNS' $CONFIG_FILE)
        ;;
    *)
        echo "Invalid argument. Usage: $0 [goerli|local]"
        return 1
        ;;
esac

# Export environment variables
export CHAIN_ID=$CHAIN_ID
export EXECUTOR_MULTISIG=$EXECUTOR_MULTISIG
export FOUNDRY_FUZZ_RUNS=$FOUNDRY_FUZZ_RUNS

# Print environment variables
echo "Environment variables set:"
echo "CHAIN_ID: $CHAIN_ID"
echo "EXECUTOR_MULTISIG: $EXECUTOR_MULTISIG"
echo "FOUNDRY_FUZZ_RUNS: $FOUNDRY_FUZZ_RUNS"