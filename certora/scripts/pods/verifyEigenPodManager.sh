if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/EigenPodManagerHarness.sol \
    certora/munged/core/DelegationManager.sol certora/munged/pods/EigenPod.sol certora/munged/strategies/StrategyBase.sol certora/munged/core/StrategyManager.sol \
    certora/munged/core/Slasher.sol certora/munged/permissions/PauserRegistry.sol \
    --verify EigenPodManagerHarness:certora/specs/pods/EigenPodManager.spec \
    --optimistic_loop \
    --prover_args '-optimisticFallback true' \
    --optimistic_hashing \
    $RULE \
    --loop_iter 3 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "EigenPodManager $1 $2" \
