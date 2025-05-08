<a name="introduction"/></a>

# EigenLayer

EigenLayer is a set of smart contracts deployed on Ethereum that enable restaking of assets to secure new services (AVSs).


## Get Started

See [CONTRIBUTING](CONTRIBUTING.md). 

Contributions that do not follow our fork base PR practices will be either rejected or immediately deleted based on your role, preventing branch pollution, keeping our repository clean, make it more readable and searchable.


## Branching

Branches we use:
* `main`: The canonical, most up-to-date branch, containing the work-in-progress code for upcoming releases
* `Vx.y.z`: Release branch with version `x.y.z` that matches a release of EigenLayer, release branch is always cut from `main` via cherry-picking
* `release-dev/xxx`: A development branch for a large feature to be released, the branch should eventually be deleted after merge to `main`


## Documentation

### Basics

To get a basic understanding of EigenLayer, check out [You Could've Invented EigenLayer](https://www.blog.eigenlayer.xyz/ycie/). Note that some of the document's content describes features that do not exist yet (like the Slasher). To understand more about how restakers and operators interact with EigenLayer, check out these guides:
* [Restaking User Guide](https://docs.eigenlayer.xyz/eigenlayer/restaking-guides/overview)
* [Operator Guide](https://docs.eigenlayer.xyz/operator-guides/operator-introduction)

### Deep Dive

The most up-to-date and technical documentation can be found in [/docs](/docs). If you're a shadowy super coder, this is a great place to get an overview of the contracts before diving into the code.

To get an idea of how users interact with these contracts, check out our integration tests: [/src/test/integration](./src/test/integration/).



## Deployments

### Deployment Version By Environments

The deployments on `mainnet`, `holesky`, and `sepolia` are on the below versions:

| Environment | Version |
| -------- | -------- |
| Mainnet Ethereum | [`v1.4.1`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.4.1) |
| Testnet Holesky | [`v1.4.2`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.4.2) |
| Testnet Sepolia | [`v1.3.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.3.0) |


### Current Deployment Contracts

<details>
    <summary>Mainnet Ethereum</summary>


###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/DelegationManager.sol) | [`0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A`](https://etherscan.io/address/0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A) | [`0xA751...7E73`](https://etherscan.io/address/0xA75112d1df37FA53a431525CD47A7d7faCEA7E73) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/StrategyManager.sol) | [`0x858646372CC42E1A627fcE94aa7A7033e7CF075A`](https://etherscan.io/address/0x858646372CC42E1A627fcE94aa7A7033e7CF075A) | [`0xba4b...b925`](https://etherscan.io/address/0xba4b2b8A076851A3044882493C2e36503d50b925) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/pods/EigenPodManager.sol) | [`0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338`](https://etherscan.io/address/0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338) | [`0x8dB4...B7E9`](https://etherscan.io/address/0x8dB49233e3b7691D68745A31e4A0Cd9Cf924B7E9) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/AVSDirectory.sol) | [`0x135dda560e946695d6f155dacafc6f1f25c1f5af`](https://etherscan.io/address/0x135dda560e946695d6f155dacafc6f1f25c1f5af) | [`0xA396...B6A2`](https://etherscan.io/address/0xA396D855D70e1A1ec1A0199ADB9845096683B6A2) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/RewardsCoordinator.sol) | [`0x7750d328b314EfFa365A0402CcfD489B80B0adda`](https://etherscan.io/address/0x7750d328b314EfFa365A0402CcfD489B80B0adda) | [`0xa505...0aB`](https://etherscan.io/address/0xa505c0116aD65071F0130061F94745b7853220aB) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/permissions/PermissionController.sol) | [`0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5`](https://etherscan.io/address/0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5) | [`0xe7f3...C6B1`](https://etherscan.io/address/0xe7f3705c9Addf2DE14e03C345fA982CAb2c1C6B1) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/AllocationManager.sol) | [`0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39`](https://etherscan.io/address/0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39) | [`0x7400...8a0b`](https://etherscan.io/address/0x740058839A1668Af5700e5d7B062007275e77D25) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyFactory.sol) | [`0x5e4C39Ad7A3E881585e383dB9827EB4811f6F647`](https://etherscan.io/address/0x5e4C39Ad7A3E881585e383dB9827EB4811f6F647) | [`0x1b97...c66`](https://etherscan.io/address/0x1b97d8F963179C0e17E5F3d85cdfd9a31A49bc66) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBase.sol) | [`0x0ed6703C298d28aE0878d1b28e88cA87F9662fE9`](https://etherscan.io/address/0x0ed6703c298d28ae0878d1b28e88ca87f9662fe9) | [`0x0EC1...F456`](https://etherscan.io/address/0x0EC17ef9c00F360DB28CA8008684a4796b11E456) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

The following strategies were originally deployed and whitelisted outside of the `StrategyFactory`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyBase (cbETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc`](https://etherscan.io/address/0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (stETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x93c4b944D05dfe6df7645A86cd2206016c51564D`](https://etherscan.io/address/0x93c4b944D05dfe6df7645A86cd2206016c51564D) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (rETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2`](https://etherscan.io/address/0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ETHx)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d`](https://etherscan.io/address/0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ankrETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff`](https://etherscan.io/address/0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (OETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xa4C637e0F704745D182e4D38cAb7E7485321d059`](https://etherscan.io/address/0xa4C637e0F704745D182e4D38cAb7E7485321d059) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (osETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x57ba429517c3473B6d34CA9aCd56c0e735b94c02`](https://etherscan.io/address/0x57ba429517c3473B6d34CA9aCd56c0e735b94c02) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (swETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6`](https://etherscan.io/address/0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (wBETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7CA911E83dabf90C90dD3De5411a10F1A6112184`](https://etherscan.io/address/0x7CA911E83dabf90C90dD3De5411a10F1A6112184) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (sfrxETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6`](https://etherscan.io/address/0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (lsETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xAe60d8180437b5C34bB956822ac2710972584473`](https://etherscan.io/address/0xAe60d8180437b5C34bB956822ac2710972584473) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (mETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x298aFB19A105D59E74658C4C334Ff360BadE6dd2`](https://etherscan.io/address/0x298aFB19A105D59E74658C4C334Ff360BadE6dd2) | [`0xaFDa...F178`](https://etherscan.io/address/0xaFDa870d4A94B9444F9F22A0e61806178b6Bf178) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/EigenStrategy.sol) | [`0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7`](https://etherscan.io/address/0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7) | [`0x90B0...8729`](https://etherscan.io/address/0x90B074DDD680bD06C72e28b09231A0F848205729) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/pods/EigenPod.sol) | [`0x5a2a4F2F3C18f09179B6703e63D9eDD165909073`](https://etherscan.io/address/0x5a2a4F2F3C18f09179B6703e63D9eDD165909073) | [`0xe2E2...46c3`](https://etherscan.io/address/0xe2E2dB234b0FFB9AFe41e52dB7d3c2B8585646c3) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/token/Eigen.sol) | [`0xec53bf9167f50cdeb3ae105f56099aaab9061f83`](https://etherscan.io/address/0xec53bf9167f50cdeb3ae105f56099aaab9061f83) | [`0x17f5...26A0`](https://etherscan.io/address/0x17f56E911C279bad67eDC08acbC9cf3DC4eF26A0) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/mainnet/src/contracts/token/BackingEigen.sol) | [`0x83E9115d334D248Ce39a6f36144aEaB5b3456e75`](https://etherscan.io/address/0x83E9115d334D248Ce39a6f36144aEaB5b3456e75) | [`0xF2b2...9b17`](https://etherscan.io/address/0xF2b225815F70c9b327DC9db758A36c92A4279b17) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`SignedDistributor`](https://etherscan.io/address/0x035bdAeaB85E47710C27EdA7FD754bA80aD4ad02#code) | - | [`0x035b...ad02`](https://etherscan.io/address/0x035bdAeaB85E47710C27EdA7FD754bA80aD4ad02) | - |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/permissions/PauserRegistry.sol) | - | [`0xB876...2806`](https://etherscan.io/address/0xB8765ed72235d279c3Fb53936E4606db0Ef12806) | |
| [`Pauser Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x5050389572f2d220ad927CcbeA0D406831012390`](https://etherscan.io/address/0x5050389572f2d220ad927CcbeA0D406831012390) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Community Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xFEA47018D632A77bA579846c840d5706705Dc598`](https://etherscan.io/address/0xFEA47018D632A77bA579846c840d5706705Dc598) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Executor Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x369e6F597e22EaB55fFb173C6d9cD234BD699111`](https://etherscan.io/address/0x369e6F597e22EaB55fFb173C6d9cD234BD699111) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Operations Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xBE1685C81aA44FF9FB319dD389addd9374383e90`](https://etherscan.io/address/0xBE1685C81aA44FF9FB319dD389addd9374383e90) | [`0xd9db...9552`](https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Compound: Timelock`](https://github.com/compound-finance/compound-protocol/blob/a3214f67b73310d547e00fc578e8355911c9d376/contracts/Timelock.sol) | - | [`0xA6Db...0EAF`](https://etherscan.io/address/0xA6Db1A8C5a981d1536266D2a393c5F8dDb210EAF) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x8b95...2444`](https://etherscan.io/address/0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444) | |


</details>



<details>
    <summary>Testnet Holesky</summary>


You can view the deployed contract addresses below, or check out the code itself on the [`testnet-holesky`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/testnet-holesky) branch.

###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/DelegationManager.sol) | [`0xA44151489861Fe9e3055d95adC98FbD462B948e7`](https://holesky.etherscan.io/address/0xA44151489861Fe9e3055d95adC98FbD462B948e7) | [`0x5c79...4106`](https://holesky.etherscan.io/address/0x5c798965208a6FddEc2eA8505c544065645c4106) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/StrategyManager.sol) | [`0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6`](https://holesky.etherscan.io/address/0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6) | [`0x7798...7679`](https://holesky.etherscan.io/address/0x7798625888ECf3EB2c3c74Dc2746e09d72747679) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/pods/EigenPodManager.sol) | [`0x30770d7E3e71112d7A6b7259542D1f680a70e315`](https://holesky.etherscan.io/address/0x30770d7E3e71112d7A6b7259542D1f680a70e315) | [`0x35b7...f304`](https://holesky.etherscan.io/address/0x35b7743633AcdaEB18a4894469fcdBF23E13f304) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/AVSDirectory.sol) | [`0x055733000064333CaDDbC92763c58BF0192fFeBf`](https://holesky.etherscan.io/address/0x055733000064333CaDDbC92763c58BF0192fFeBf) | [`0x6966...8a0b`](https://holesky.etherscan.io/address/0x69660e721e4013dd8FEf83dCE731E915d74b8a0b) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/RewardsCoordinator.sol) | [`0xAcc1fb458a1317E886dB376Fc8141540537E68fE`](https://holesky.etherscan.io/address/0xAcc1fb458a1317E886dB376Fc8141540537E68fE) | [`0xA3c3...0d69`](https://holesky.etherscan.io/address/0xA3c31d2FBAD3d924baA64f8789E03E9FA7d70d69) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/AllocationManager.sol) | [`0x78469728304326CBc65f8f95FA756B0B73164462`](https://holesky.etherscan.io/address/0x78469728304326CBc65f8f95FA756B0B73164462) | [`0xA248...968A`](https://holesky.etherscan.io/address/0xA24898174701fF888B127e65c827E441F4Bc968A) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/permissions/PermissionController.sol) | [`0x598cb226B591155F767dA17AfE7A2241a68C5C10`](https://holesky.etherscan.io/address/0x598cb226B591155F767dA17AfE7A2241a68C5C10) | [`0x7ab0...a2b9`](https://holesky.etherscan.io/address/0x7ab0ebd25d5ffe7527600ca5b2858c1a3faba2b9#code) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyFactory.sol) | [`0x9c01252B580efD11a05C00Aa42Dd3ac1Ec52DF6d`](https://holesky.etherscan.io/address/0x9c01252B580efD11a05C00Aa42Dd3ac1Ec52DF6d) | [`0x62e3...7aDE`](https://holesky.etherscan.io/address/0x62e328C554AD0F8eD4735C215Ff43d8b8a407aDE) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBase.sol) | [`0xd3c6C6BA4E40dB9288c6a2077e5635344F8aFA4F`](https://holesky.etherscan.io/address/0xd3c6C6BA4E40dB9288c6a2077e5635344F8aFA4F) | [`0xcaF7...C7E2`](https://holesky.etherscan.io/address/0xcaF7217Ca38F8262573E0cd3Ed660a78Ea19C7E2) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

The following strategies were originally deployed and whitelisted outside of the `StrategyFactory`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyBase (stETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3`](https://holesky.etherscan.io/address/0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (rETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0`](https://holesky.etherscan.io/address/0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (WETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9`](https://holesky.etherscan.io/address/0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (lsETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943`](https://holesky.etherscan.io/address/0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (sfrxETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c`](https://holesky.etherscan.io/address/0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ETHx)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x31B6F59e1627cEfC9fA174aD03859fC337666af7`](https://holesky.etherscan.io/address/0x31B6F59e1627cEfC9fA174aD03859fC337666af7) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (osETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef`](https://holesky.etherscan.io/address/0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (cbETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6`](https://holesky.etherscan.io/address/0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (mETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xaccc5A86732BE85b5012e8614AF237801636F8e5`](https://holesky.etherscan.io/address/0xaccc5A86732BE85b5012e8614AF237801636F8e5) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ankrETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac`](https://holesky.etherscan.io/address/0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (reALT)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xAD76D205564f955A9c18103C4422D1Cd94016899`](https://holesky.etherscan.io/address/0xAD76D205564f955A9c18103C4422D1Cd94016899) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (EO)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x78dBcbEF8fF94eC7F631c23d38d197744a323868`](https://holesky.etherscan.io/address/0x78dBcbEF8fF94eC7F631c23d38d197744a323868) | [`0x5FdD...3C1e`](https://holesky.etherscan.io/address/0x5FdD6a71a3C88111474C812Ca6d60942d7923C1e) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/EigenStrategy.sol) | [`0x43252609bff8a13dFe5e057097f2f45A24387a84`](https://holesky.etherscan.io/address/0x43252609bff8a13dFe5e057097f2f45A24387a84) | [`0x917F...8Fd0`](https://holesky.etherscan.io/address/0x917F70Dd0C97332024A556c6EFeD6B9a8be98Fd0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/pods/EigenPod.sol) | [`0x7261C2bd75a7ACE1762f6d7FAe8F63215581832D`](https://holesky.etherscan.io/address/0x7261C2bd75a7ACE1762f6d7FAe8F63215581832D) | [`0x68bd...6D37`](https://holesky.etherscan.io/address/0x68bd1e75E9863C9066B46B8a44E953F918466D37) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/token/Eigen.sol) | [`0x3B78576F7D6837500bA3De27A60c7f594934027E`](https://holesky.etherscan.io/address/0x3B78576F7D6837500bA3De27A60c7f594934027E) | [`0x01cb...3050`](https://holesky.etherscan.io/address/0x01cbB2ae8eFE46EEdB9f7575D91cA1EB38823050) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/token/BackingEigen.sol) | [`0x275cCf9Be51f4a6C94aBa6114cdf2a4c45B9cb27`](https://holesky.etherscan.io/address/0x275cCf9Be51f4a6C94aBa6114cdf2a4c45B9cb27) | [`0x05ad...E05c`](https://holesky.etherscan.io/address/0x05adA1C66DdDD7c36705bC23a4d50dBa72E4E05c) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/permissions/PauserRegistry.sol) | - | [`0x41Db...ec1D`](https://holesky.etherscan.io/address/0x41Dbe7BbacA97D986FCF6f5203b98Ec02412ec1D) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0x5e83...F6fD`](https://holesky.etherscan.io/address/0x5e83c7d195318A5acf46B29E5810DdC323b2F6fD) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xDB02...A6cf`](https://holesky.etherscan.io/address/0xDB023566064246399b4AE851197a97729C93A6cf) | |

</details>



<details>
    <summary>Testnet Sepolia</summary>

###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/DelegationManager.sol) | [`0xD4A7E1Bd8015057293f0D0A557088c286942e84b`](https://sepolia.etherscan.io/address/0xD4A7E1Bd8015057293f0D0A557088c286942e84b) | [`0xa722...67f3`](https://sepolia.etherscan.io/address/0xa7227485e6C693AC4566fe168C5E3647c5c267f3) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/StrategyManager.sol) | [`0x2E3D6c0744b10eb0A4e6F679F71554a39Ec47a5D`](https://sepolia.etherscan.io/address/0x2E3D6c0744b10eb0A4e6F679F71554a39Ec47a5D) | [`0x8845...EB1b`](https://sepolia.etherscan.io/address/0x88457741E1bDE012a36Fc154B9004384C1eAEB1b) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/pods/EigenPodManager.sol) | [`0x56BfEb94879F4543E756d26103976c567256034a`](https://sepolia.etherscan.io/address/0x56BfEb94879F4543E756d26103976c567256034a) | [`0x8b1D...11a0`](https://sepolia.etherscan.io/address/0x8b1DBbAa79507CD6b6e1d9FBe90E3FB18EFf11a0) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) | All EigenPod functionality is paused on Holesky | 
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/AVSDirectory.sol) | [`0xa789c91ECDdae96865913130B786140Ee17aF545`](https://sepolia.etherscan.io/address/0xa789c91ECDdae96865913130B786140Ee17aF545) | [`0xD88b...C188`](https://sepolia.etherscan.io/address/0xD88b96998325c3e74A74a0B0938BBFeA1395C188) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/RewardsCoordinator.sol) | [`0x5ae8152fb88c26ff9ca5C014c94fca3c68029349`](https://sepolia.etherscan.io/address/0x5ae8152fb88c26ff9ca5C014c94fca3c68029349) | [`0xcC30...7940`](https://sepolia.etherscan.io/address/0xcC305562B01bec562D13A40ef8781e313AFE7940) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/core/AllocationManager.sol) | [`0x42583067658071247ec8CE0A516A58f682002d07`](https://sepolia.etherscan.io/address/0x42583067658071247ec8CE0A516A58f682002d07) | [`0x742A...927b`](https://sepolia.etherscan.io/address/0x742A228482701d693061BfE9C5B3Eb3959Ea927b) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/permissions/PermissionController.sol) | [`0x44632dfBdCb6D3E21EF613B0ca8A6A0c618F5a37`](https://sepolia.etherscan.io/address/0x44632dfBdCb6D3E21EF613B0ca8A6A0c618F5a37) | [`0x59B1...f525`](https://sepolia.etherscan.io/address/0x59B11b191B572888703E150E45F5015e0fFcf525) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyFactory.sol) | [`0x066cF95c1bf0927124DFB8B02B401bc23A79730D`](https://sepolia.etherscan.io/address/0x066cF95c1bf0927124DFB8B02B401bc23A79730D) | [`0xEE41...ca1A`](https://sepolia.etherscan.io/address/0xEE41826B7D5B89e7F5eED6a831b4eFD69FC9ca1A) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBase.sol) | [`0x427e627Bc7E83cac0f84337d3Ad94230C32697D3`](https://sepolia.etherscan.io/address/0x427e627Bc7E83cac0f84337d3Ad94230C32697D3) | [`0x9E05...7C5c`](https://sepolia.etherscan.io/address/0x9E0540212b45FE44459cDAD25CD9077acFB77C5c) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| [`StETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBase.sol) | [`0x8b29d91e67b013e855EaFe0ad704aC4Ab086a574`](https://sepolia.etherscan.io/address/0x8b29d91e67b013e855EaFe0ad704aC4Ab086a574) | [`0x9E05...7C5c`](https://sepolia.etherscan.io/address/0x9E0540212b45FE44459cDAD25CD9077acFB77C5c) | Strategy Factory deployed |
| [`WETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/StrategyBase.sol) | [`0x424246eF71b01ee33aA33aC590fd9a0855F5eFbc`](https://sepolia.etherscan.io/address/0x424246eF71b01ee33aA33aC590fd9a0855F5eFbc) | [`0x9E05...7C5c`](https://sepolia.etherscan.io/address/0x9E0540212b45FE44459cDAD25CD9077acFB77C5c) | Strategy Factory deployed |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/strategies/EigenStrategy.sol) | [`0x8E93249a6C37a32024756aaBd813E6139b17D1d5`](https://sepolia.etherscan.io/address/0x8E93249a6C37a32024756aaBd813E6139b17D1d5) | [`0x46CF...16db`](https://sepolia.etherscan.io/address/0x46CFA3C2eaDe97D53739120b87A63F739B9616db) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

**NOTE: Due to the permissioned validator set on Sepolia, all EigenPod functionality is *PAUSED*.**

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/pods/EigenPod.sol) | [`0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4`](https://sepolia.etherscan.io/address/0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4) | [`0xd85d...6bCf`](https://sepolia.etherscan.io/address/0xd85d0D9e24dC9af8a517034Caab2db68aD936bCf) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/token/Eigen.sol) | [`0x0011FA2c512063C495f77296Af8d195F33A8Dd38`](https://sepolia.etherscan.io/address/0x0011FA2c512063C495f77296Af8d195F33A8Dd38) | [`0xF83a...8725`](https://sepolia.etherscan.io/address/0xF83a81117AE073B13ce70f37302392BA90F28725) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/token/BackingEigen.sol) | [`0xc5B857A92245f64e9D90cCc5b096Db82eB77eB5c`](https://sepolia.etherscan.io/address/0xc5B857A92245f64e9D90cCc5b096Db82eB77eB5c) | [`0x1298...3173`](https://sepolia.etherscan.io/address/0x12988B679AA497C30A8D1850eCC4Dc7700383173) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/src/contracts/permissions/PauserRegistry.sol) | - | [`0x63AA...20f3`](https://sepolia.etherscan.io/address/0x63AAe451780090f50Ad323aAEF155F63a29D20f3) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0x1BEF...1b5B`](https://sepolia.etherscan.io/address/0x1BEF05C7303d44e0E2FCD2A19d993eDEd4c51b5B) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x56E8...6Fa1`](https://sepolia.etherscan.io/address/0x56E88cb4f0136fC27D95499dE4BE2acf47946Fa1) | |

</details>
