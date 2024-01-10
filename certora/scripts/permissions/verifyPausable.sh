if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/PausableHarness.sol \
    src/contracts/permissions/PauserRegistry.sol \
    --verify PausableHarness:certora/specs/permissions/Pausable.spec \
    --optimistic_loop \
    --prover_args '-optimisticFallback true -recursionErrorAsAssert false -recursionEntryLimit 3' \
    --loop_iter 3 \
    --link PausableHarness:pauserRegistry=PauserRegistry \
    $RULE \
    --msg "Pausable $1 $2" \