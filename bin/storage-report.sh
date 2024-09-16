#!/bin/sh
mkdir docs/storage-report
for file in $(find src/contracts -name "*.sol" ! -path "src/contracts/interfaces/*" ! -path "src/contracts/libraries/*"); do
    contract_name=$(basename "$file" .sol)
    forge inspect "$contract_name" storage --pretty > docs/storage-report/"$contract_name".md
done