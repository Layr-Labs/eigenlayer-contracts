#!/bin/bash

# Check if version argument is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <new_version>"
    exit 1
fi

NEW_VERSION="$1"
CONTRACTS_DIR="src/contracts"

# Find and update the version in matching contracts
find "$CONTRACTS_DIR" -type f -name "*.sol" | while read -r file; do
    if grep -q 'function version() public pure override' "$file"; then
        sed -i '' "s/return \".*\";/return \"$NEW_VERSION\";/" "$file"
        echo "Updated version in: $file"
    fi
done

echo "Done!"
