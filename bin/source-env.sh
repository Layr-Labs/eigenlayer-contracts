#!/bin/bash

# Check for arguments
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 [local]"
    return 1
fi

case $1 in
    local)
        CHAIN_ID=31337
        FOUNDRY_FUZZ_RUNS=256
        ;;
    *)
        echo "Invalid argument. Usage: $0 [local]"
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