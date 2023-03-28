if [[ "$2" ]]
then
    RULE="--rule $2"
fi


certoraRun certora/harnesses/PausableHarness.sol \
    certora/munged/permissions/PauserRegistry.sol \
    --verify PausableHarness:certora/specs/permissions/Pausable.spec \
    --optimistic_loop \
    --send_only \
    --settings -optimisticFallback=true,-recursionErrorAsAssert=false,-recursionEntryLimit=3 \
    --loop_iter 3 \
    --link PausableHarness:pauserRegistry=PauserRegistry \
    $RULE \
    --msg "Pausable $1 $2" \