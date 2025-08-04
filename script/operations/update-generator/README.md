# Generator Update

The `generator`, signs off on roots. In order to generate this operatorSet, we use the following [script](../../deploy/multichain/deploy_globalRootConfirmerSet.s.sol) The script outputs two items:

1. `network.wallet.json`: Private keys and BLS sig info (should be kept secure)
2. `network.toml`: A toml file passed into initialization on `operatorTableUpdater`

See our [docs](../../../docs/multichain/) for more information. 