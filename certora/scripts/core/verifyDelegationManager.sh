if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/harnesses/DelegationManagerHarness.sol \
    lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol lib/openzeppelin-contracts/contracts/mocks/ERC1271WalletMock.sol \
    src/contracts/pods/EigenPodManager.sol src/contracts/pods/EigenPod.sol src/contracts/strategies/StrategyBase.sol src/contracts/core/StrategyManager.sol \
    src/contracts/core/Slasher.sol src/contracts/permissions/PauserRegistry.sol \
    --verify DelegationManagerHarness:certora/specs/core/DelegationManager.spec \
    --solc_via_ir \
    --solc_optimize 1 \
    --optimistic_loop \
    --optimistic_fallback \
    --optimistic_hashing \
    --parametric_contracts DelegationManagerHarness \
    $RULE \
    --loop_iter 2 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "DelegationManager $1 $2" \
