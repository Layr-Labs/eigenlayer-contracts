if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12 

certoraRun certora/harnesses/PausableHarness.sol \
    certora/munged/permissions/PauserRegistry.sol \
    --verify PausableHarness:certora/specs2/permissions/Pausable.spec \
    --optimistic_loop \
    --send_only \
    --prover_args '-optimisticFallback true -recursionErrorAsAssert false -recursionEntryLimit 3' \
    --loop_iter 3 \
    --link PausableHarness:pauserRegistry=PauserRegistry \
    $RULE \
    --msg "Pausable $1 $2" \