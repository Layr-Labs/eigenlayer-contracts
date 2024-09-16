#!/bin/sh

# Default output directory
OUTPUT_DIR=${1:-docs/storage-report}

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Loop through Solidity files and generate storage report
# NOTE: Ignores `src/contracts/interfaces` & `src/contracts/libraries` since they "should" not contain storage logic.
for file in $(find src/contracts -name "*.sol" ! -path "src/contracts/interfaces/*" ! -path "src/contracts/libraries/*"); do
    contract_name=$(basename "$file" .sol)
    forge inspect "$contract_name" storage --pretty > "$OUTPUT_DIR/$contract_name.md"
done
