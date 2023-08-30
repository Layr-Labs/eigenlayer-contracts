if [[ "$2" ]]
then
    RULE="--rule $2"
fi

solc-select use 0.8.12

certoraRun certora/munged/middleware/BLSRegistryCoordinatorWithIndices.sol \
    lib/openzeppelin-contracts/contracts/utils/cryptography/draft-EIP712.sol lib/openzeppelin-contracts/contracts/mocks/ERC1271WalletMock.sol \
    certora/munged/middleware/StakeRegistry.sol certora/munged/middleware/BLSPubkeyRegistry.sol certora/munged/middleware/IndexRegistry.sol \
    certora/munged/core/Slasher.sol \
    --verify BLSRegistryCoordinatorWithIndices:certora/specs/middleware/BLSRegistryCoordinatorWithIndices.spec \
    --optimistic_loop \
    --prover_args '-optimisticFallback true' \
    $RULE \
    --loop_iter 3 \
    --packages @openzeppelin=lib/openzeppelin-contracts @openzeppelin-upgrades=lib/openzeppelin-contracts-upgradeable \
    --msg "BLSRegistryCoordinatorWithIndices $1 $2" \

# TODO: import a ServiceManager contract