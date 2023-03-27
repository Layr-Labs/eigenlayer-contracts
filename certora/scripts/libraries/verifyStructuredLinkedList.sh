if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/StructuredLinkedListHarness.sol \
    --verify StructuredLinkedListHarness:certora/specs/libraries/StructuredLinkedList.spec \
    --optimistic_loop \
    --send_only \
    --settings -optimisticFallback=true \
    $RULE \
    --rule_sanity \
    # --staging master \
    --loop_iter 3 \
    --msg "StructuredLinkedList $1 $2" \