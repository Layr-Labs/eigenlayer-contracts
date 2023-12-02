## EigenLayer M2 Docs

**EigenLayer M2 is a testnet-only release** that extends the functionality of EigenLayer M1 (which is live on both Goerli and mainnet).

M1 enables very basic restaking: users that stake ETH natively or with a liquid staking token can opt-in to the M1 smart contracts, which currently support two basic operations: deposits and withdrawals. 

M2 adds several features, the most important of which is the basic support needed to create an AVS<!--(*link: ["what is an AVS?"](https://github.com/Layr-Labs/eigenlayer-contracts/blob/master/docs/AVS-Guide.md) TODO*)-->. The M2 release includes the first AVS, EigenDA <!--(*link: read more about EigenDA (TODO)*)-->. The other features of M2 support AVSs and pad out existing features of M1. A short list of new features includes:
* Anyone can register as an operator
* Operators can begin providing services to an AVS
* Stakers can delegate their stake to a single operator
* Native ETH restaking is now fully featured, using beacon chain state proofs to validate withdrawal credentials, validator balances, and validator exits
* Proofs are supported by beacon chain headers provided by an oracle (See [`EigenPodManager` docs](./core/EigenPodManager.md) for more info)
* TODO - multiquorums

### System Components

**EigenPodManager**:

| File | Type | Proxy | Goerli |
| -------- | -------- | -------- | -------- |
| [`EigenPodManager.sol`](../src/contracts/pods/EigenPodManager.sol) | Singleton | [`0x3093...C9a5`](https://goerli.etherscan.io/address/0x3093F3B560352F896F0e9567019902C9Aff8C9a5) | [`0x86bf...6CcA`](https://goerli.etherscan.io/address/0x86bf376E0C0c9c6D332E13422f35Aca75C106CcA) |
| [`EigenPod.sol`](../src/contracts/pods/EigenPod.sol) | Instanced, deployed per-user | [`0x3093...C9a5`](https://goerli.etherscan.io/address/0x3093F3B560352F896F0e9567019902C9Aff8C9a5) | [`0x86bf...6CcA`](https://goerli.etherscan.io/address/0x86bf376E0C0c9c6D332E13422f35Aca75C106CcA) |
| [`DelayedWithdrawalRouter.sol`](../src/contracts/pods/DelayedWithdrawalRouter.sol) | Singleton | [`0x895...388f`](https://goerli.etherscan.io/address/0x89581561f1F98584F88b0d57c2180fb89225388f) | [`0x607...7fe`](https://goerli.etherscan.io/address/0x60700ade3Bf48C437a5D01b6Da8d7483cffa27fe) |
| [`succinctlabs/EigenLayerBeaconOracle.sol`](https://github.com/succinctlabs/telepathy-contracts/blob/main/external/integrations/eigenlayer/EigenLayerBeaconOracle.sol) | Singleton | - | [`0x40B1...9f2c`](https://goerli.etherscan.io/address/0x40B10ddD29a2cfF33DBC420AE5bbDa0649049f2c) |

These contracts work together to enable native ETH restaking:
* Users deploy `EigenPods` via the `EigenPodManager`, which contain beacon chain state proof logic used to verify a validator's withdrawal credentials, balance, and exit. An `EigenPod's` main role is to serve as the withdrawal address for one or more of a user's validators.
* The `EigenPodManager` handles `EigenPod` creation and accounting+interactions between users with restaked native ETH and the `DelegationManager`.
* The `DelayedWithdrawalRouter` imposes a 7-day delay on completing certain withdrawals from an `EigenPod`. This is primarily to add a stopgap against a hack being able to instantly withdraw funds (note that most withdrawals are processed via the `DelegationManager`, not the `DelayedWithdrawalRouter`).
* The `EigenLayerBeaconOracle` provides beacon chain block roots for use in various proofs. The oracle is supplied by Succinct's Telepathy protocol ([docs link](https://docs.telepathy.xyz/)).

See full documentation in [`/core/EigenPodManager.md`](./core/EigenPodManager.md).

**StrategyManager**:

| File | Type | Proxy | Goerli |
| -------- | -------- | -------- | -------- |
| [`StrategyManager.sol`](../src/contracts/core/StrategyManager.sol) | Singleton | [`0x779...8E907`](https://goerli.etherscan.io/address/0x779d1b5315df083e3F9E94cB495983500bA8E907) | [`0x8676...0055`](https://goerli.etherscan.io/address/0x8676bb5f792ED407a237234Fe422aC6ed3540055) |
| [`StrategyBaseTVLLimits.sol`](../src/contracts/strategies/StrategyBaseTVLLimits.sol) | Instanced, stETH | [`0xB6...d14da`](https://goerli.etherscan.io/address/0xb613e78e2068d7489bb66419fb1cfa11275d14da) | [`0x81E9...F8ebA`](https://goerli.etherscan.io/address/0x81E94e16949AC397d508B5C2557a272faD2F8ebA) |
| [`StrategyBaseTVLLimits.sol`](../src/contracts/strategies/StrategyBaseTVLLimits.sol) | Instanced, rETH | [`0x8799...70b5`](https://goerli.etherscan.io/address/0x879944A8cB437a5f8061361f82A6d4EED59070b5) | [`0x81E9...F8ebA`](https://goerli.etherscan.io/address/0x81E94e16949AC397d508B5C2557a272faD2F8ebA) |

These contracts work together to enable restaking for LSTs:
* The `StrategyManager` acts as the entry and exit point for LSTs in EigenLayer. It handles deposits into each of the 3 LST-specific strategies, and manages accounting+interactions between users with restaked LSTs and the `DelegationManager`.
* `StrategyBaseTVLLimits` is deployed as three separate instances, one for each supported LST (cbETH, rETH, and stETH). When a user deposits into a strategy through the `StrategyManager`, this contract receives the tokens and awards the user with a proportional quantity of shares in the strategy. When a user withdraws, the strategy contract sends the LSTs back to the user.

See full documentation in [`/core/StrategyManager.md`](./core/StrategyManager.md).

**DelegationManager**:

| File | Type | Proxy | Goerli |
| -------- | -------- | -------- | -------- |
| [`DelegationManager.sol`](../src/contracts/core/DelegationManager.sol) | Singleton | [`0x1b7...Eb0a8`](https://goerli.etherscan.io/address/0x1b7b8F6b258f95Cf9596EabB9aa18B62940Eb0a8) | [`0x9b79...A99d`](https://goerli.etherscan.io/address/0x9b7980a32ceCe2Aa936DD2E43AF74af62581A99d) |

The `DelegationManager` sits between the `EigenPodManager` and `StrategyManager` to manage delegation and undelegation of Stakers to Operators. Its primary features are to allow Operators to register as Operators (`registerAsOperator`), to keep track of shares being delegated to Operators across different strategies, and to manage withdrawals on behalf of the `EigenPodManager` and `StrategyManager`.

See full documentation in [`/core/DelegationManager.md`](./core/DelegationManager.md).

**Slasher**:

| File | Type | Proxy | Goerli |
| -------- | -------- | -------- | -------- |
| [`Slasher.sol`](../src/contracts/core/Slasher.sol) | Singleton | [`0xD1...0C22`](https://goerli.etherscan.io/address/0xD11d60b669Ecf7bE10329726043B3ac07B380C22) | [`0x3865...8Be6`](https://goerli.etherscan.io/address/0x3865B5F5297f86c5295c7f818BAD1fA5286b8Be6) |

The `Slasher` is deployed, but will remain completely paused during M2. Its design is not finalized.