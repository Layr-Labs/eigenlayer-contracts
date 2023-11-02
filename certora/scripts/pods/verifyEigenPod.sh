if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/EigenPodHarness.sol \
    certora/munged/core/DelegationManager.sol certora/munged/pods/EigenPodManager.sol certora/munged/strategies/StrategyBase.sol certora/munged/core/StrategyManager.sol \
    certora/munged/core/Slasher.sol certora/munged/permissions/PauserRegistry.sol \
    --verify EigenPodHarness:certora/specs/pods/EigenPod.spec \
    --optimistic_loop \
    --prover_args '-optimisticFallback true' \
    $RULE \
    --loop_iter 2 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "EigenPod $1 $2" \
