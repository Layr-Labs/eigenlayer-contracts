mkdir storage-report
for file in $(find src/contracts -name "*.sol"); do
    contract_name=$(basename "$file" .sol)
    forge inspect "$contract_name" storage --pretty > storage-report/"$contract_name".md
done