if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/EigenPodHarness.sol \
    src/contracts/core/DelegationManager.sol src/contracts/pods/EigenPodManager.sol \
    src/contracts/core/Slasher.sol src/contracts/permissions/PauserRegistry.sol \
    src/contracts/core/StrategyManager.sol \
    src/contracts/strategies/StrategyBase.sol \
    lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol \
    lib/openzeppelin-contracts/contracts/mocks/ERC1271WalletMock.sol \
    --verify EigenPodHarness:certora/specs/pods/EigenPod.spec \
    --solc_via_ir \
    --solc_optimize 1 \
    --optimistic_loop \
    --prover_args '-recursionEntryLimit 3' \
    --optimistic_hashing \
    --parametric_contracts EigenPodHarness \
    $RULE \
    --loop_iter 1 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "EigenPod $1 $2" \
