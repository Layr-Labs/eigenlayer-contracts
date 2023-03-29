if [[ "$2" ]]
then
    RULE="--rule $2"
fi


certoraRun certora/munged/strategies/StrategyBase.sol \
    lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol \
    certora/munged/core/StrategyManager.sol \
    certora/munged/permissions/PauserRegistry.sol \
    certora/munged/core/Slasher.sol \
    --verify StrategyBase:certora/specs/strategies/StrategyBase.spec \
    --optimistic_loop \
    --send_only \
    --settings -optimisticFallback=true,-recursionErrorAsAssert=false,-recursionEntryLimit=3 \
    --loop_iter 3 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --link StrategyBase:strategyManager=StrategyManager \
    $RULE \
    --msg "Pausable $1 $2" \