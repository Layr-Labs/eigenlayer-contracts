if [[ "$2" ]]
then
    RULE="--rule $2"
fi

# solc-select use 0.8.27

# certoraRun certora/harnesses/EigenPodHarness.sol \
#     src/contracts/core/DelegationManager.sol src/contracts/pods/EigenPodManager.sol \
#     src/contracts/permissions/PauserRegistry.sol \
#     src/contracts/core/StrategyManager.sol \
#     src/contracts/strategies/StrategyBase.sol \
#     lib/openzeppelin-contracts-v4.9.0/contracts/token/ERC20/ERC20.sol \
#     lib/openzeppelin-contracts-v4.9.0/contracts/mocks/ERC1271WalletMock.sol \
#     --verify EigenPodHarness:certora/specs/pods/EigenPod.spec \
#     --optimistic_loop \
#     --prover_args '-recursionEntryLimit 3' \
#     --optimistic_hashing \
#     --parametric_contracts EigenPodHarness \
#     $RULE \
#     --loop_iter 1 \
#     --packages @openzeppelin=lib/openzeppelin-contracts-v4.9.0 @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable-v4.9.0 \
#     --msg "EigenPod $1 $2" \
