if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.27

certoraRun src/contracts/strategies/StrategyBase.sol \
    lib/openzeppelin-contracts-v4.9.0/contracts/token/ERC20/ERC20.sol \
    src/contracts/core/StrategyManager.sol \
    src/contracts/permissions/PauserRegistry.sol \
    --verify StrategyBase:certora/specs/strategies/StrategyBase.spec \
    --solc_via_ir \
    --solc_optimize 1 \
    --optimistic_loop \
    --optimistic_fallback \
    --prover_args '-recursionErrorAsAssert false -recursionEntryLimit 3' \
    --loop_iter 3 \
    --packages @openzeppelin=lib/openzeppelin-contracts-v4.9.0 @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable-v4.9.0 \
    --link StrategyBase:strategyManager=StrategyManager \
    --parametric_contracts StrategyBase \
    $RULE \
    --msg "StrategyBase $1 $2" \