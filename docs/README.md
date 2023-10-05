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

| File | Type | Proxy? | Goerli |
| -------- | -------- | -------- | -------- |
| [`EigenPodManager.sol`](../src/contracts/pods/EigenPodManager.sol) | Singleton | Transparent proxy | TODO |
| [`EigenPod.sol`](../src/contracts/pods/EigenPod.sol) | Instanced, deployed per-user | Beacon proxy | TODO |
| [`DelayedWithdrawalRouter.sol`](../src/contracts/pods/DelayedWithdrawalRouter.sol) | Singleton | Transparent proxy | TODO |
| [`succinctlabs/EigenLayerBeaconOracle.sol`](https://github.com/succinctlabs/telepathy-contracts/blob/main/external/integrations/eigenlayer/EigenLayerBeaconOracle.sol) | Singleton | UUPS proxy | [`0x40B1...9f2c`](https://goerli.etherscan.io/address/0x40B10ddD29a2cfF33DBC420AE5bbDa0649049f2c) |

These contracts work together to enable native ETH restaking:
* Users deploy `EigenPods` via the `EigenPodManager`, which contain beacon chain state proof logic used to verify a validator's withdrawal credentials, balance, and exit. An `EigenPod's` main role is to serve as the withdrawal address for one or more of a user's validators.
* The `EigenPodManager` handles `EigenPod` creation, validator withdrawal, and accounting+interactions between users with restaked native ETH and the `DelegationManager`.
* The `DelayedWithdrawalRouter` imposes a 7-day delay on completing withdrawals from an `EigenPod`. This is primarily to add a stopgap against a hack being able to instantly withdraw funds.
* The `EigenLayerBeaconOracle` provides beacon chain block roots for use in various proofs. The oracle is supplied by Succinct's Telepathy protocol ([docs link](https://docs.telepathy.xyz/)).

See full documentation in [`/core/EigenPodManager.md`](./core/EigenPodManager.md).

**StrategyManager**:

| File | Type | Proxy? | Goerli |
| -------- | -------- | -------- | -------- |
| [`StrategyManager.sol`](../src/contracts/core/StrategyManager.sol) | Singleton | Transparent proxy | TODO |
| [`StrategyBaseTVLLimits.sol`](../src/contracts/strategies/StrategyBaseTVLLimits.sol) | 3 instances (for cbETH, rETH, stETH) | Transparent proxy | TODO |

These contracts work together to enable restaking for LSTs:
* The `StrategyManager` acts as the entry and exit point for LSTs in EigenLayer. It handles deposits and withdrawals from each of the 3 LST-specific strategies, and manages interactions between users with restaked LSTs and the `DelegationManager`.
* `StrategyBaseTVLLimits` is deployed as three separate instances, one for each supported LST (cbETH, rETH, and stETH). When a user deposits into a strategy through the `StrategyManager`, this contract receives the tokens and awards the user with a proportional quantity of shares in the strategy. When a user withdraws, the strategy contract sends the LSTs back to the user.

See full documentation in [`/core/StrategyManager.md`](./core/StrategyManager.md).

**DelegationManager**:

| File | Type | Proxy? | Goerli |
| -------- | -------- | -------- | -------- |
| [`DelegationManager.sol`](../src/contracts/core/DelegationManager.sol) | Singleton | Transparent proxy | TODO |

The `DelegationManager` sits between the `EigenPodManager` and `StrategyManager` to manage delegation and undelegation of Stakers to Operators. Its primary features are to allow Operators to register as Operators (`registerAsOperator`), and to keep track of shares being delegated to Operators across different strategies.

See full documentation in [`/core/DelegationManager.md`](./core/DelegationManager.md).

**Slasher**:

| File | Type | Proxy? | Goerli |
| -------- | -------- | -------- | -------- |
| [`Slasher.sol`](../src/contracts/core/Slasher.sol) | Singleton | Transparent proxy | TODO |

The `Slasher` is deployed, but will remain completely paused during M2. Though some contracts make calls to the `Slasher`, they are all currently no-ops.

#### Supporting Components

**PauserRegistry and MSig**: TODO

**ExecutorMsig and Comm/Ops/Timelock Msigs**

**Proxies and ProxyAdmin**