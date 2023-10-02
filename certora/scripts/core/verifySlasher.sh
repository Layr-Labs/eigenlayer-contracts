if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12   

certoraRun certora/harnesses/SlasherHarness.sol \
    lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol \
    certora/munged/pods/EigenPodManager.sol certora/munged/pods/EigenPod.sol certora/munged/strategies/StrategyBase.sol certora/munged/core/DelegationManager.sol \
    certora/munged/core/StrategyManager.sol certora/munged/permissions/PauserRegistry.sol \
    --verify SlasherHarness:certora/specs/core/Slasher.spec \
    --optimistic_loop \
    --prover_args '-optimisticFallback true -recursionErrorAsAssert false -recursionEntryLimit 2' \
    --loop_iter 2 \
    --link SlasherHarness:delegation=DelegationManager \
    $RULE \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "Slasher $1 $2" \
