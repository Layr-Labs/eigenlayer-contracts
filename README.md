<a name="introduction"/></a>

# EigenLayer

EigenLayer is a set of smart contracts deployed on Ethereum that enable restaking of assets to secure new services. This repo contains the EigenLayer core contracts, whose currently-supported assets include beacon chain ETH and several liquid staking tokens (LSTs). Users use these contracts to deposit and withdraw these assets, as well as delegate them to operators providing services to AVSs.

## Getting Started

* [Branching](#branching)
* [Documentation](#documentation)
* [Building and Running Tests](#building-and-running-tests)
* [Deployments](#deployments)

## Branching

The main branches we use are:
* [`dev (default)`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/dev): The most up-to-date branch, containing the work-in-progress code for upcoming releases
* [`testnet-holesky`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/testnet-holesky): Our current testnet deployment
* [`mainnet`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/mainnet): Our current mainnet deloyment

## Documentation

### Basics

To get a basic understanding of EigenLayer, check out [You Could've Invented EigenLayer](https://www.blog.eigenlayer.xyz/ycie/). Note that some of the document's content describes features that do not exist yet (like the Slasher). To understand more about how restakers and operators interact with EigenLayer, check out these guides:
* [Restaking User Guide](https://docs.eigenlayer.xyz/restaking-guides/restaking-user-guide)
* [Operator Guide](https://docs.eigenlayer.xyz/operator-guides/operator-introduction)

### Deep Dive

The most up-to-date and technical documentation can be found in [/docs](/docs). If you're a shadowy super coder, this is a great place to get an overview of the contracts before diving into the code.

To get an idea of how users interact with these contracts, check out our integration tests: [/src/test/integration](./src/test/integration/).

## Building and Running Tests

This repository uses Foundry. See the [Foundry docs](https://book.getfoundry.sh/) for more info on installation and usage. If you already have foundry, you can build this project and run tests with these commands:

```
foundryup

forge build
forge test
```

### Running Fork Tests

We have a few fork tests against ETH mainnet. Passing these requires the environment variable `RPC_MAINNET` to be set. See `.env.example` for an example. Once you've set up your environment, `forge test` should show these fork tests passing.

Additionally, to run all tests in a forked environment, [install yq](https://mikefarah.gitbook.io/yq/v/v3.x/). Then, set up your environment using this script to read from `config.yml`:

`source source-env.sh [goerli|local]`

Then run the tests:

`forge test --fork-url [RPC_URL]`

### Running Static Analysis

1. Install [solhint](https://github.com/protofire/solhint), then run:

`solhint 'src/contracts/**/*.sol'`

2. Install [slither](https://github.com/crytic/slither), then run:

`slither .`

### Generate Inheritance and Control-Flow Graphs

1. Install [surya](https://github.com/ConsenSys/surya/) and graphviz:

```
npm i -g surya

apt install graphviz
```

2. Then, run:

```
surya inheritance ./src/contracts/**/*.sol | dot -Tpng > InheritanceGraph.png

surya mdreport surya_report.md ./src/contracts/**/*.sol
```

## Deployments

### Current Mainnet Deployment

The current mainnet deployment is our M1 release, and is from a much older version of this repo. You can view the deployed contract addresses below, or check out the code itself on the [`mainnet`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/mainnet) branch.

| Name | Solidity | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | -------- | 
| StrategyManager | [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/core/StrategyManager.sol) | [`0x8586...075A`](https://etherscan.io/address/0x858646372CC42E1A627fcE94aa7A7033e7CF075A) | [`0x70f4...619b`](https://etherscan.io/address/0x70f44c13944d49a236e3cd7a94f48f5dab6c619b) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: cbETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x5494...56bc`](https://etherscan.io/address/0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: stETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x93c4...564D`](https://etherscan.io/address/0x93c4b944D05dfe6df7645A86cd2206016c51564D) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: rETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x1BeE...dCD2`](https://etherscan.io/address/0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: ETHx | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9d7e...011d`](https://etherscan.io/address/0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: ankrETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x1376...58ff`](https://etherscan.io/address/0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: OETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xa4C6...d059`](https://etherscan.io/address/0xa4C637e0F704745D182e4D38cAb7E7485321d059) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: osETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x57ba...4c02`](https://etherscan.io/address/0x57ba429517c3473B6d34CA9aCd56c0e735b94c02) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: swETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x0Fe4...96d6`](https://etherscan.io/address/0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: wBETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7CA9...2184`](https://etherscan.io/address/0x7CA911E83dabf90C90dD3De5411a10F1A6112184) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: sfrxETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x8CA7...28b6`](https://etherscan.io/address/0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: lsETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xAe60...4473`](https://etherscan.io/address/0xAe60d8180437b5C34bB956822ac2710972584473) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: mETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x298a...6dd2`](https://etherscan.io/address/0x298aFB19A105D59E74658C4C334Ff360BadE6dd2) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| EigenPodManager | [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/pods/EigenPodManager.sol) | [`0x91E6...A338`](https://etherscan.io/address/0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338) | [`0xe429...5762`](https://etherscan.io/address/0xe4297e3dadbc7d99e26a2954820f514cb50c5762) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| EigenLayerBeaconOracle | [`succinctlabs/EigenLayerBeaconOracle.sol`](https://github.com/succinctlabs/eigenlayer-beacon-oracle/blob/main/contracts/src/EigenLayerBeaconOracle.sol) | - | [`0x3439...5442`](https://etherscan.io/address/0x343907185b71adf0eba9567538314396aa985442) | |
| EigenPod (beacon) | [`EigenPod`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/pods/EigenPod.sol) | [`0x5a2a...9073`](https://etherscan.io/address/0x5a2a4F2F3C18f09179B6703e63D9eDD165909073) | [`0x8ba4...a255`](https://etherscan.io/address/0x8ba40da60f0827d027f029acee62609f0527a255) | - Beacon: [OpenZeppelin BeaconProxy@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Deployed pods use [UpgradableBeacon@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| DelayedWithdrawalRouter | [`DelayedWithdrawalRouter`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/pods/DelayedWithdrawalRouter.sol) | [`0x7Fe7...23D8`](https://etherscan.io/address/0x7Fe7E9CC0F274d2435AD5d56D5fa73E47F6A23D8) | [`0x4bb6...4226`](https://etherscan.io/address/0x4bb6731b02314d40abbffbc4540f508874014226) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| DelegationManager | [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/core/DelegationManager.sol) | [`0x3905...f37A`](https://etherscan.io/address/0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A) | [`0x1784...9dda`](https://etherscan.io/address/0x1784be6401339fc0fedf7e9379409f5c1bfe9dda) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Slasher | [`Slasher`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/core/Slasher.sol) | [`0xD921...c3Cd`](https://etherscan.io/address/0xD92145c07f8Ed1D392c1B88017934E301CC1c3Cd) | [`0xf323...6614`](https://etherscan.io/address/0xf3234220163a757edf1e11a8a085638d9b236614) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| PauserRegistry | [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/permissions/PauserRegistry.sol) | - | [`0x0c43...7060`](https://etherscan.io/address/0x0c431C66F4dE941d089625E5B423D00707977060) | |
| Pauser Multisig | [`GnosisSafe@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x5050...2390`](https://etherscan.io/address/0x5050389572f2d220ad927CcbeA0D406831012390) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`GnosisSafeProxy@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| Community Multisig | [`GnosisSafe@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xFEA4...c598`](https://etherscan.io/address/0xFEA47018D632A77bA579846c840d5706705Dc598) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`GnosisSafeProxy@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| Executor Multisig | [`GnosisSafe@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x369e...9111`](https://etherscan.io/address/0x369e6F597e22EaB55fFb173C6d9cD234BD699111) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`GnosisSafeProxy@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| Operations Multisig | [`GnosisSafe@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xBE16...3e90`](https://etherscan.io/address/0xBE1685C81aA44FF9FB319dD389addd9374383e90) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`GnosisSafeProxy@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| Timelock | [Compound: `Timelock.sol`](https://github.com/compound-finance/compound-protocol/blob/a3214f67b73310d547e00fc578e8355911c9d376/contracts/Timelock.sol) | - | [`0xA6Db...0EAF`](https://etherscan.io/address/0xA6Db1A8C5a981d1536266D2a393c5F8dDb210EAF) | |
| Proxy Admin | [OpenZeppelin ProxyAdmin@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x8b95...2444`](https://etherscan.io/address/0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444) | |

### Current Testnet Deployment

The current testnet deployment is on holesky, and is from our M2 beta release. You can view the deployed contract addresses below, or check out the code itself on the [`testnet-holesky`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/testnet-holesky) branch.

| Name | Solidity | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | -------- | 
| StrategyManager | [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/core/StrategyManager.sol) | [`0xdfB5...D5b6`](https://holesky.etherscan.io/address/0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6) | [`0x59f7...3a18`](https://holesky.etherscan.io/address/0x59f766A603C53f3AC8Be43bBe158c1519b193a18) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: stETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7D70...63d3`](https://holesky.etherscan.io/address/0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: rETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x3A8f...c4E0`](https://holesky.etherscan.io/address/0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: WETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x8052...FaF9`](https://holesky.etherscan.io/address/0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: lsETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x5037...D943`](https://holesky.etherscan.io/address/0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: sfrxETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9281...631c`](https://holesky.etherscan.io/address/0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: ETHx | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x31B6...6af7`](https://holesky.etherscan.io/address/0x31B6F59e1627cEfC9fA174aD03859fC337666af7) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: osETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x4628...C6Ef`](https://holesky.etherscan.io/address/0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: cbETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x70EB...bAB6`](https://holesky.etherscan.io/address/0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: mETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xaccc...F8e5`](https://holesky.etherscan.io/address/0xaccc5A86732BE85b5012e8614AF237801636F8e5) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: ankrETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7673...d0ac`](https://holesky.etherscan.io/address/0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac) | [`0xFb83...3305`](https://holesky.etherscan.io/address/0xFb83e1D133D0157775eC4F19Ff81478Df1103305) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| EigenPodManager | [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/pods/EigenPodManager.sol) | [`0x3077...e315`](https://holesky.etherscan.io/address/0x30770d7E3e71112d7A6b7259542D1f680a70e315) | [`0x5265...4a7B`](https://holesky.etherscan.io/address/0x5265C162f7d5F3fE3175a78828ab16bf5E324a7B) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| EigenLayerBeaconOracle | [`succinctlabs/EigenLayerBeaconOracle.sol`](https://github.com/succinctlabs/eigenlayer-beacon-oracle/blob/main/contracts/src/EigenLayerBeaconOracle.sol) | - | [`0x4C11...8f25`](https://holesky.etherscan.io/address/0x4C116BB629bff7A8373c2378bBd919f8349B8f25) | |
| EigenPod (beacon) | [`EigenPod`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/pods/EigenPod.sol) | [`0x7261...832D`](https://holesky.etherscan.io/address/0x7261C2bd75a7ACE1762f6d7FAe8F63215581832D) | [`0xe98f...641c`](https://holesky.etherscan.io/address/0xe98f9298344527608A1BCC23907B8145F9Cb641c) | - Beacon: [OpenZeppelin BeaconProxy@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Deployed pods use [UpgradableBeacon@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| DelayedWithdrawalRouter | [`DelayedWithdrawalRouter`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/pods/DelayedWithdrawalRouter.sol) | [`0x642c...FA32`](https://holesky.etherscan.io/address/0x642c646053eaf2254f088e9019ACD73d9AE0FA32) | [`0xcE8b...3407`](https://holesky.etherscan.io/address/0xcE8b8D99773a718423F8040a6e52c06a4ce63407) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| DelegationManager | [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/core/DelegationManager.sol) | [`0xA441...48e7`](https://holesky.etherscan.io/address/0xA44151489861Fe9e3055d95adC98FbD462B948e7) | [`0x83f8...0D76`](https://holesky.etherscan.io/address/0x83f8F8f0BB125F7870F6bfCf76853f874C330D76) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| AVSDirectory | [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/core/AVSDirectory.sol) | [`0x0557...FeBf`](https://holesky.etherscan.io/address/0x055733000064333CaDDbC92763c58BF0192fFeBf) | [`0xEF5B...3e3a`](https://holesky.etherscan.io/address/0xEF5BA995Bc7722fd1e163edF8Dc09375de3d3e3a) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Slasher | [`Slasher`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/core/Slasher.sol) | [`0xcAe7...6e7C`](https://holesky.etherscan.io/address/0xcAe751b75833ef09627549868A04E32679386e7C) | [`0x9971...345A`](https://holesky.etherscan.io/address/0x99715D255E34a39bE9943b82F281CA734bcF345A) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| PauserRegistry | [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/permissions/PauserRegistry.sol) | - | [`0x85Ef...2F06`](https://holesky.etherscan.io/address/0x85Ef7299F8311B25642679edBF02B62FA2212F06) | |
| Timelock | [Compound: `Timelock.sol`](https://github.com/compound-finance/compound-protocol/blob/a3214f67b73310d547e00fc578e8355911c9d376/contracts/Timelock.sol) | - | [`0xcF19...0A7D`](https://holesky.etherscan.io/address/0xcF19CE0561052a7A7Ff21156730285997B350A7D) | |
| Proxy Admin | [OpenZeppelin ProxyAdmin@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xDB02...A6cf`](https://holesky.etherscan.io/address/0xDB023566064246399b4AE851197a97729C93A6cf) | |

