if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/StrategyManagerHarness.sol \
    lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol lib/openzeppelin-contracts/contracts/mocks/ERC1271WalletMock.sol \
    src/contracts/pods/EigenPodManager.sol src/contracts/pods/EigenPod.sol src/contracts/pods/DelayedWithdrawalRouter.sol \
    src/contracts/strategies/StrategyBase.sol src/contracts/core/DelegationManager.sol \
    src/contracts/core/Slasher.sol src/contracts/permissions/PauserRegistry.sol \
    --verify StrategyManagerHarness:certora/specs/core/StrategyManager.spec \
    --solc_via_ir \
    --solc_optimize 1 \
    --optimistic_loop \
    --optimistic_fallback \
    --optimistic_hashing \
    --parametric_contracts StrategyManagerHarness \
    $RULE \
    --loop_iter 2 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "StrategyManager $1 $2" \
