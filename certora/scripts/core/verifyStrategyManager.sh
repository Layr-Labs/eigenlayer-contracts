if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/StrategyManagerHarness.sol \
    lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol lib/openzeppelin-contracts/contracts/mocks/ERC1271WalletMock.sol \
    certora/munged/pods/EigenPodManager.sol certora/munged/pods/EigenPod.sol certora/munged/pods/DelayedWithdrawalRouter.sol \
    certora/munged/strategies/StrategyBase.sol certora/munged/core/DelegationManager.sol \
    certora/munged/core/Slasher.sol certora/munged/permissions/PauserRegistry.sol \
    --verify StrategyManagerHarness:certora/specs/core/StrategyManager.spec \
    --optimistic_loop \
    --prover_args '-optimisticFallback true' \
    --optimistic_hashing \
    --parametric_contracts StrategyManagerHarness \
    $RULE \
    --loop_iter 2 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "StrategyManager $1 $2" \
