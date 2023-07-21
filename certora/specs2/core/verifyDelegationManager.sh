if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/DelegationManagerHarness.sol \
    lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol lib/openzeppelin-contracts/contracts/mocks/ERC1271WalletMock.sol \
    certora/munged/pods/EigenPodManager.sol certora/munged/pods/EigenPod.sol certora/munged/strategies/StrategyBase.sol certora/munged/core/StrategyManager.sol \
    certora/munged/core/Slasher.sol certora/munged/permissions/PauserRegistry.sol \
    --verify DelegationManagerHarness:certora/specs2/core/DelegationManager.spec \
    --optimistic_loop \
    --send_only \
    --prover_args '-optimisticFallback true' \
    $RULE \
    --loop_iter 3 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "DelegationManager $1 $2" \
