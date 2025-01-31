if [[ "$2" ]]
then
    RULE="--rule $2"
fi

# solc-select use 0.8.27

# certoraRun certora/harnesses/EigenPodManagerHarness.sol \
#     src/contracts/core/DelegationManager.sol src/contracts/pods/EigenPod.sol src/contracts/strategies/StrategyBase.sol src/contracts/core/StrategyManager.sol \
#     src/contracts/permissions/PauserRegistry.sol \
#     --verify EigenPodManagerHarness:certora/specs/pods/EigenPodManager.spec \
#     --optimistic_loop \
#     --optimistic_fallback \
#     --optimistic_hashing \
#     --parametric_contracts EigenPodManagerHarness \
#     $RULE \
#     --loop_iter 3 \
#     --packages @openzeppelin=lib/openzeppelin-contracts-v4.9.0 @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable-v4.9.0 \
#     --msg "EigenPodManager $1 $2" \
