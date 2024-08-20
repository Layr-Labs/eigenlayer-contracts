if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/PausableHarness.sol \
    src/contracts/permissions/PauserRegistry.sol \
    --verify PausableHarness:certora/specs/permissions/Pausable.spec \
    --solc_via_ir \
    --solc_optimize 1 \
    --optimistic_loop \
    --optimistic_fallback \
    --prover_args '-recursionErrorAsAssert false -recursionEntryLimit 3' \
    --loop_iter 3 \
    --link PausableHarness:pauserRegistry=PauserRegistry \
    $RULE \
    --msg "Pausable $1 $2" \