if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/StructuredLinkedListHarness.sol \
    --verify StructuredLinkedListHarness:certora/specs/libraries/StructuredLinkedList.spec \
    --solc_via_ir \
    --solc_optimize 1 \
    --optimistic_loop \
    --optimistic_fallback \
    --parametric_contracts StructuredLinkedListHarness \
    $RULE \
    --rule_sanity \
    --loop_iter 3 \
    --msg "StructuredLinkedList $1 $2" \