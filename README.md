<a name="introduction"/></a>

# EigenLayer

EigenLayer is a set of smart contracts deployed on Ethereum that enable restaking of assets to secure new services. This repo contains the EigenLayer core contracts, whose currently-supported assets include beacon chain ETH and several liquid staking tokens (LSTs). Users use these contracts to deposit and withdraw these assets, as well as delegate them to operators providing services to AVSs.

## Getting Started

* [Documentation](#documentation)
* [Building and Running Tests](#building-and-running-tests)
* [Deployments](#deployments)

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

The current mainnet deployment is our M1 release, and is from a much older version of this repo. You can view the deployed contract addresses in [Current Mainnet Deployment](#current-mainnet-deployment) below, or check out the [`init-mainnet-deployment`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/mainnet-deployment) branch in "Releases".

| Name | Solidity | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | -------- | 
| StrategyManager | [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/core/StrategyManager.sol) | [`0x8586...075A`](https://etherscan.io/address/0x858646372CC42E1A627fcE94aa7A7033e7CF075A) | [`0x5d25...42Fb`](https://etherscan.io/address/0x5d25EEf8CfEdaA47d31fE2346726dE1c21e342Fb) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: cbETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x5494...56bc`](https://etherscan.io/address/0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: stETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x93c4...564D`](https://etherscan.io/address/0x93c4b944D05dfe6df7645A86cd2206016c51564D) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: rETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x1BeE...dCD2`](https://etherscan.io/address/0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| EigenPodManager | [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/pods/EigenPodManager.sol) | [`0x91E6...A338`](https://etherscan.io/address/0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338) | [`0xEB86...e111`](https://etherscan.io/address/0xEB86a5c40FdE917E6feC440aBbCDc80E3862e111) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| EigenPod (beacon) | [`EigenPod`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/pods/EigenPod.sol) | [`0x5a2a...9073`](https://etherscan.io/address/0x5a2a4F2F3C18f09179B6703e63D9eDD165909073) | [`0x5c86...9dA7`](https://etherscan.io/address/0x5c86e9609fbBc1B754D0FD5a4963Fdf0F5b99dA7) | - Beacon: [OpenZeppelin BeaconProxy@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Deployed pods use [UpgradableBeacon@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| DelayedWithdrawalRouter | [`DelayedWithdrawalRouter`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/pods/DelayedWithdrawalRouter.sol) | [`0x7Fe7...23D8`](https://etherscan.io/address/0x7Fe7E9CC0F274d2435AD5d56D5fa73E47F6A23D8) | [`0x44Bc...E2AF`](https://etherscan.io/address/0x44Bcb0E01CD0C5060D4Bb1A07b42580EF983E2AF) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| DelegationManager | [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/core/DelegationManager.sol) | [`0x3905...f37A`](https://etherscan.io/address/0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A) | [`0xf97E...75e4`](https://etherscan.io/address/0xf97E97649Da958d290e84E6D571c32F4b7F475e4) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Slasher | [`Slasher`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/core/Slasher.sol) | [`0xD921...c3Cd`](https://etherscan.io/address/0xD92145c07f8Ed1D392c1B88017934E301CC1c3Cd) | [`0xef31...d6d8`](https://etherscan.io/address/0xef31c292801f24f16479DD83197F1E6AeBb8d6d8) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| PauserRegistry | [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/permissions/PauserRegistry.sol) | - | [`0x0c43...7060`](https://etherscan.io/address/0x0c431C66F4dE941d089625E5B423D00707977060) | |
| Pauser Multisig | [`GnosisSafe@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x5050…2390`](https://etherscan.io/address/0x5050389572f2d220ad927CcbeA0D406831012390) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`GnosisSafeProxy@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| Community Multisig | [`GnosisSafe@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xFEA4...c598`](https://etherscan.io/address/0xFEA47018D632A77bA579846c840d5706705Dc598) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`GnosisSafeProxy@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| Executor Multisig | [`GnosisSafe@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x369e...9111`](https://etherscan.io/address/0x369e6F597e22EaB55fFb173C6d9cD234BD699111) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`GnosisSafeProxy@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| Operations Multisig | [`GnosisSafe@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xBE16...3e90`](https://etherscan.io/address/0xBE1685C81aA44FF9FB319dD389addd9374383e90) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`GnosisSafeProxy@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| Timelock | [Compound: `Timelock.sol`](https://github.com/compound-finance/compound-protocol/blob/a3214f67b73310d547e00fc578e8355911c9d376/contracts/Timelock.sol) | - | [`0xA6Db...0EAF`](https://etherscan.io/address/0xA6Db1A8C5a981d1536266D2a393c5F8dDb210EAF) | |
| Proxy Admin | [OpenZeppelin ProxyAdmin@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x8b95...2444`](https://etherscan.io/address/0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444) | |

### Current Testnet Deployment

The current testnet deployment is from our M2 beta release, which is a slightly older version of this repo. You can view the deployed contract addresses in [Current Testnet Deployment](#current-testnet-deployment), or check out the [`goerli-m2-deployment`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/goerli-m2-deployment) branch in "Releases".

| Name | Solidity | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | -------- | 
| StrategyManager | [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/core/StrategyManager.sol) | [`0x779d...E907`](https://goerli.etherscan.io/address/0x779d1b5315df083e3F9E94cB495983500bA8E907) | [`0x8676...0055`](https://goerli.etherscan.io/address/0x8676bb5f792ED407a237234Fe422aC6ed3540055) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: stETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xB613...14da`](https://goerli.etherscan.io/address/0xb613e78e2068d7489bb66419fb1cfa11275d14da) | [`0x81E9...8ebA`](https://goerli.etherscan.io/address/0x81E94e16949AC397d508B5C2557a272faD2F8ebA) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Strategy: rETH | [`StrategyBaseTVLLimits`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x8799...70b5`](https://goerli.etherscan.io/address/0x879944A8cB437a5f8061361f82A6d4EED59070b5) | [`0x81E9...8ebA`](https://goerli.etherscan.io/address/0x81E94e16949AC397d508B5C2557a272faD2F8ebA) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| EigenPodManager | [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/pods/EigenPodManager.sol) | [`0xa286...df41`](https://goerli.etherscan.io/address/0xa286b84C96aF280a49Fe1F40B9627C2A2827df41) | [`0xdD09...901b`](https://goerli.etherscan.io/address/0xdD09d95bD25299EDBF4f33d76F84dBc77b0B901b) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| EigenLayerBeaconOracle | [`succinctlabs/EigenLayerBeaconOracle.sol`](https://github.com/succinctlabs/telepathy-contracts/blob/main/external/integrations/eigenlayer/EigenLayerBeaconOracle.sol) | [`0x40B1...9f2c`](https://goerli.etherscan.io/address/0x40B10ddD29a2cfF33DBC420AE5bbDa0649049f2c) | [`0x0a6e...db01`](https://goerli.etherscan.io/address/0x0a6e235c30658dbdb53147fbb199878a4e34db01) | OpenZeppelin UUPS |
| EigenPod (beacon) | [`EigenPod`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/pods/EigenPod.sol) | [`0x3093...C9a5`](https://goerli.etherscan.io/address/0x3093F3B560352F896F0e9567019902C9Aff8C9a5) | [`0x86bf...6CcA`](https://goerli.etherscan.io/address/0x86bf376E0C0c9c6D332E13422f35Aca75C106CcA) | - Beacon: [OpenZeppelin BeaconProxy@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Deployed pods use [UpgradableBeacon@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| DelayedWithdrawalRouter | [`DelayedWithdrawalRouter`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/pods/DelayedWithdrawalRouter.sol) | [`0x8958...388f`](https://goerli.etherscan.io/address/0x89581561f1F98584F88b0d57c2180fb89225388f) | [`0x6070...27fe`](https://goerli.etherscan.io/address/0x60700ade3Bf48C437a5D01b6Da8d7483cffa27fe) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| DelegationManager | [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/core/DelegationManager.sol) | [`0x1b7b...b0a8`](https://goerli.etherscan.io/address/0x1b7b8F6b258f95Cf9596EabB9aa18B62940Eb0a8) | [`0x9b79...A99d`](https://goerli.etherscan.io/address/0x9b7980a32ceCe2Aa936DD2E43AF74af62581A99d) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| Slasher | [`Slasher`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/core/Slasher.sol) | [`0xD11d...0C22`](https://goerli.etherscan.io/address/0xD11d60b669Ecf7bE10329726043B3ac07B380C22) | [`0x3865...8Be6`](https://goerli.etherscan.io/address/0x3865B5F5297f86c5295c7f818BAD1fA5286b8Be6) | Proxy: [OpenZeppelin TUP@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| PauserRegistry | [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/goerli-m2-deployment/src/contracts/permissions/PauserRegistry.sol) | - | [`0x2588...0010`](https://goerli.etherscan.io/address/0x2588f9299871a519883D92dcd5092B4A0Cf70010) | |
| Timelock | [Compound: `Timelock.sol`](https://github.com/compound-finance/compound-protocol/blob/a3214f67b73310d547e00fc578e8355911c9d376/contracts/Timelock.sol) | - | [`0xa7e7...796e`](https://goerli.etherscan.io/address/0xa7e72a0564ebf25fa082fc27020225edeaf1796e) | |
| Proxy Admin | [OpenZeppelin ProxyAdmin@4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x28ce...02e2`](https://goerli.etherscan.io/address/0x28ceac2ff82B2E00166e46636e2A4818C29902e2) | |

