<a name="introduction"/></a>

# EigenLayer 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/Layr-Labs/eigenlayer-contracts/workflows/Tests/badge.svg)](https://github.com/Layr-Labs/eigenlayer-contracts/actions)

> EigenLayer - A restaking protocol for Ethereum that enables staked assets to secure additional services.

EigenLayer is a set of smart contracts deployed on Ethereum that enable restaking of assets to secure new services. This repo contains the EigenLayer core contracts, whose currently-supported assets include beacon chain ETH and several liquid staking tokens (LSTs). Users use these contracts to deposit and withdraw these assets, as well as delegate them to operators providing services to AVSs.

## 📑 Contents

* [Getting Started](#getting-started)
* [Branching](#branching) 
* [Documentation](#documentation)
* [Building and Testing](#building-and-running-tests)
* [Deployments](#deployments)
* [Contributing](#contributing)

## 🚀 Getting Started

To get started with EigenLayer:

1. Clone the repository
2. Install dependencies: `make deps`
3. Build the project: `forge build`
4. Run tests: `forge test`

## 🌿 Branching

The main branches we use are:
* [`dev (default)`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/dev): The most up-to-date branch, containing the work-in-progress code for upcoming releases
* [`testnet-holesky`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/testnet-holesky): Our current testnet deployment
* [`mainnet`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/mainnet): Our current mainnet deployment

## 📚 Documentation

### Basics

To get a basic understanding of EigenLayer, check out [You Could've Invented EigenLayer](https://www.blog.eigenlayer.xyz/ycie/). Note that some of the document's content describes features that do not exist yet (like the Slasher). To understand more about how restakers and operators interact with EigenLayer, check out these guides:

* [Restaking User Guide](https://docs.eigenlayer.xyz/restaking-guides/restaking-user-guide)
* [Operator Guide](https://docs.eigenlayer.xyz/operator-guides/operator-introduction)

### Deep Dive

The most up-to-date and technical documentation can be found in [/docs](/docs). If you're a seasoned developer, this is a great place to get an overview of the contracts before diving into the code.

To get an idea of how users interact with these contracts, check out our integration tests: [/src/test/integration](./src/test/integration/).

## 👥 Contributing

We welcome contributions to the project! Here are several ways you can help:

1. Create issues for bugs you find
2. Submit improvements via pull requests
3. Improve documentation
4. Share feedback about protocol usage

### Developer Environment Setup

```bash
# Install dependencies
make deps

# This will install:
# - pre-commit hook
# - foundry and tools
# - abigen
```

## 📋 Deployments

### Current Mainnet Deployment

Our current mainnet deployment is our M2 version. You can view the deployed contract addresses below or check the code on the [`mainnet`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/mainnet) branch.

###### Core Contracts

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/core/DelegationManager.sol) | [`0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A`](https://etherscan.io/address/0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A) | [`0x1784...9dda`](https://etherscan.io/address/0x1784be6401339fc0fedf7e9379409f5c1bfe9dda) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/core/StrategyManager.sol) | [`0x858646372CC42E1A627fcE94aa7A7033e7CF075A`](https://etherscan.io/address/0x858646372CC42E1A627fcE94aa7A7033e7CF075A) | [`0x70f4...619b`](https://etherscan.io/address/0x70f44c13944d49a236e3cd7a94f48f5dab6c619b) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/bda003385c5fec59e35196dc14d01f17d1eb7001/src/contracts/pods/EigenPodManager.sol) | [`0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338`](https://etherscan.io/address/0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338) | [`0x731A...3FEa`](https://etherscan.io/address/0x731A0aD160e407393Ff662231Add6Dd145AD3FEa) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/core/AVSDirectory.sol) | [`0x135dda560e946695d6f155dacafc6f1f25c1f5af`](https://etherscan.io/address/0x135dda560e946695d6f155dacafc6f1f25c1f5af) | [`0xdabd...a5b7`](https://etherscan.io/address/0xdabdb3cd346b7d5f5779b0b614ede1cc9dcba5b7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Slasher`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/core/Slasher.sol) | [`0xD92145c07f8Ed1D392c1B88017934E301CC1c3Cd`](https://etherscan.io/address/0xD92145c07f8Ed1D392c1B88017934E301CC1c3Cd) | [`0xf323...6614`](https://etherscan.io/address/0xf3234220163a757edf1e11a8a085638d9b236614) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/core/RewardsCoordinator.sol) | [`0x7750d328b314EfFa365A0402CcfD489B80B0adda`](https://etherscan.io/address/0x7750d328b314EfFa365A0402CcfD489B80B0adda) | [`0x5bf7...8785`](https://etherscan.io/address/0x5bf7c13D5FAdba224ECB3D5C0a67A231D1628785) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s using `StrategyFactory`, deployed at the address below (see [documentation](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from `StrategyFactory` are deployed using the transparent proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyFactory.sol) | [`0x5e4c39ad7a3e881585e383db9827eb4811f6f647`](https://etherscan.io/address/0x5e4c39ad7a3e881585e383db9827eb4811f6f647) | [`0x3e07...5c74`](https://etherscan.io/address/0x3e07cc2d34c8e0965f5ba45ac1e960e535155c74) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/testnet-holesky/src/contracts/strategies/StrategyBase.sol) | [`0x0ed6703C298d28aE0878d1b28e88cA87F9662fE9`](https://etherscan.io/address/0x0ed6703c298d28ae0878d1b28e88ca87f9662fe9) | [`0xe9fa...7827`](https://etherscan.io/address/0xe9fa8f904d97854c7389b68923262adcc6c27827#code) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

The following strategies were initially deployed and whitelisted outside of `StrategyFactory`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyBase (cbETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc`](https://etherscan.io/address/0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (stETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x93c4b944D05dfe6df7645A86cd2206016c51564D`](https://etherscan.io/address/0x93c4b944D05dfe6df7645A86cd2206016c51564D) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (rETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2`](https://etherscan.io/address/0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ETHx)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d`](https://etherscan.io/address/0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ankrETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff`](https://etherscan.io/address/0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (OETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xa4C637e0F704745D182e4D38cAb7E7485321d059`](https://etherscan.io/address/0xa4C637e0F704745D182e4D38cAb7E7485321d059) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (osETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x57ba429517c3473B6d34CA9aCd56c0e735b94c02`](https://etherscan.io/address/0x57ba429517c3473B6d34CA9aCd56c0e735b94c02) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (swETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6`](https://etherscan.io/address/0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (wBETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7CA911E83dabf90C90dD3De5411a10F1A6112184`](https://etherscan.io/address/0x7CA911E83dabf90C90dD3De5411a10F1A6112184) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (sfrxETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6`](https://etherscan.io/address/0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (lsETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xAe60d8180437b5C34bB956822ac2710972584473`](https://etherscan.io/address/0xAe60d8180437b5C34bB956822ac2710972584473) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (mETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x298aFB19A105D59E74658C4C334Ff360BadE6dd2`](https://etherscan.io/address/0x298aFB19A105D59E74658C4C334Ff360BadE6dd2) | [`0xdfdA...46d3`](https://etherscan.io/address/0xdfdA04f980bE6A64E3607c95Ca26012Ab9aA46d3) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies - Special

The following strategies significantly differ from the other strategies used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/strategies/EigenStrategy.sol) | [`0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7`](https://etherscan.io/address/0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7) | [`0x27e7...0428`](https://etherscan.io/address/0x27e7a3a81741b9fcc5ad7edcbf9f8a72a5c00428) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/bda003385c5fec59e35196dc14d01f17d1eb7001/src/contracts/pods/EigenPod.sol) | [`0x5a2a4F2F3C18f09179B6703e63D9eDD165909073`](https://etherscan.io/address/0x5a2a4F2F3C18f09179B6703e63D9eDD165909073) | [`0x6D22...6430`](https://etherscan.io/address/0x6D225e974Fa404D25Ffb84eD6E242Ffa18eF6430) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| [`DelayedWithdrawalRouter`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/4b15d68b7e16b5965bad398496bfce57f5a47e1b/src/contracts/pods/DelayedWithdrawalRouter.sol) | [`0x7Fe7E9CC0F274d2435AD5d56D5fa73E47F6A23D8`](https://etherscan.io/address/0x7Fe7E9CC0F274d2435AD5d56D5fa73E47F6A23D8) | [`0x4bb6...4226`](https://etherscan.io/address/0x4bb6731b02314d40abbffbc4540f508874014226) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/token/Eigen.sol) | [`0xec53bf9167f50cdeb3ae105f56099aaab9061f83`](https://etherscan.io/address/0xec53bf9167f50cdeb3ae105f56099aaab9061f83) | [`0x17f5...26A0`](https://etherscan.io/address/0x17f56E911C279bad67eDC08acbC9cf3DC4eF26A0) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/token/BackingEigen.sol) | [`0x83E9115d334D248Ce39a6f36144aEaB5b3456e75`](https://etherscan.io/address/0x83E9115d334D248Ce39a6f36144aEaB5b3456e75) | [`0xF2b2...9b17`](https://etherscan.io/address/0xF2b225815F70c9b327DC9db758A36c92A4279b17) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`SignedDistributor`](https://etherscan.io/address/0x035bdAeaB85E47710C27EdA7FD754bA80aD4ad02#code) | - | [`0x035b...ad02`](https://etherscan.io/address/0x035bdAeaB85E47710C27EdA7FD754bA80aD4ad02) | - |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/0139d6213927c0a7812578899ddd3dda58051928/src/contracts/permissions/PauserRegistry.sol) | - | [`0x0c43...7060`](https://etherscan.io/address/0x0c431C66F4dE941d089625E5B423D00707977060) | |
| [`Pauser Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x5050389572f2d220ad927CcbeA0D406831012390`](https://etherscan.io/address/0x5050389572f2d220ad927CcbeA0D406831012390) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Community Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xFEA47018D632A77bA579846c840d5706705Dc598`](https://etherscan.io/address/0xFEA47018D632A77bA579846c840d5706705Dc598) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Executor Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x369e6F597e22EaB55fFb173C6d9cD234BD699111`](https://etherscan.io/address/0x369e6F597e22EaB55fFb173C6d9cD234BD699111) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Operations Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xBE1685C81aA44FF9FB319dD389addd9374383e90`](https://etherscan.io/address/0xBE1685C81aA44FF9FB319dD389addd9374383e90) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Compound: Timelock`](https://github.com/compound-finance/compound-protocol/blob/a3214f67b73310d547e00fc578e8355911c9d376/contracts/Timelock.sol) | - | [`0xA6Db...0EAF`](https://etherscan.io/address/0xA6Db1A8C5a981d1536266D2a393c5F8dDb210EAF) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x8b95...2444`](https://etherscan.io/address/0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444) | |
