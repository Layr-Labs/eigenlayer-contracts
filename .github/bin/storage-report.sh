#!/bin/sh

# Default output directory
OUTPUT_DIR=${1:-docs/storage-report}

# Function to print messages
log() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') [INFO] $1"
}

# Function to print error messages
error() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') [ERROR] $1" >&2
}

log "Starting the storage report generation."

# Create the output directory if it doesn't exist
if ! mkdir -p "$OUTPUT_DIR"; then
    error "Failed to create output directory: $OUTPUT_DIR"
    exit 1
fi

log "Output directory is set to: $OUTPUT_DIR"

# Loop through Solidity files and generate storage report
# NOTE: Ignores `src/contracts/interfaces` & `src/contracts/libraries` since they "should" not contain storage logic.
for file in $(find src/contracts -name "*.sol" ! -path "src/contracts/interfaces/*" ! -path "src/contracts/libraries/*"); do
    contract_name=$(basename "$file" .sol)
    
    # Check if the file exists and is readable
    if [ ! -r "$file" ]; then
        error "Cannot read file: $file"
        continue
    fi
    
    log "Processing contract: $contract_name"

    # Run forge inspect and capture errors
    if ! forge inspect "$contract_name" storage --pretty > "$OUTPUT_DIR/$contract_name.md"; then
        error "Failed to generate storage report for contract: $contract_name"
    else
        log "Storage report generated for contract: $contract_name"
    fi
done