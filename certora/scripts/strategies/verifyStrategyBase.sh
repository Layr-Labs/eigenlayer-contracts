if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun src/contracts/strategies/StrategyBase.sol \
    lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol \
    src/contracts/core/StrategyManager.sol \
    src/contracts/permissions/PauserRegistry.sol \
    src/contracts/core/Slasher.sol \
    --verify StrategyBase:certora/specs/strategies/StrategyBase.spec \
    --solc_via_ir \
    --solc_optimize 1 \
    --optimistic_loop \
    --optimistic_fallback \
    --prover_args '-recursionErrorAsAssert false -recursionEntryLimit 3' \
    --loop_iter 3 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --link StrategyBase:strategyManager=StrategyManager \
    --parametric_contracts StrategyBase \
    $RULE \
    --msg "StrategyBase $1 $2" \