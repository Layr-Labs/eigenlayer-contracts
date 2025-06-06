<a name="introduction"/></a>

# EigenLayer

**EigenLayer** is a protocol built on Ethereum that introduces Restaking, a primitive for app and service builders to make verifiable commitments to their users.

EigenLayer brings together Restakers, Operators, and Autonomous Verifiable Services (AVSs) to extend Ethereum's cryptoeconomic security with penalty and reward commitments (like slashing) on staked assets acting as security. The protocol supports permissionless security; EIGEN, Native ETH, LSTs, and ERC-20s. 

## Deployments

The deployments on `mainnet`, `holesky`, `sepolia`, and `hoodi` are on the below versions:

| Environment | Version |
| -------- | -------- |
| Mainnet Ethereum | [`v1.4.1`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.4.1) |
| Testnet Holesky | [`v1.6.0-rc.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.6.0-rc.0) |
| Testnet Sepolia | [`v1.5.0-rc.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.5.0-rc.0) |
| Testnet Hoodi | [`v1.6.0-rc.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.6.0-rc.0) |

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
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob//v1.5.0-rc.0/src/contracts/core/DelegationManager.sol) | [`0xA44151489861Fe9e3055d95adC98FbD462B948e7`](https://holesky.etherscan.io/address/0xA44151489861Fe9e3055d95adC98FbD462B948e7) | [`0xd563...67EC`](https://holesky.etherscan.io/address/0xd5639169277e342582Ccd3A67C9fcAFFEF3867EC) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob//v1.5.0-rc.0/src/contracts/core/StrategyManager.sol) | [`0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6`](https://holesky.etherscan.io/address/0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6) | [`0x84aa...b1a7d`](https://holesky.etherscan.io/address/0x84aaD0F753b84Cd68F36Ff207DDfA0f2865b1a7d) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob//v1.5.0-rc.0/src/contracts/pods/EigenPodManager.sol) | [`0x30770d7E3e71112d7A6b7259542D1f680a70e315`](https://holesky.etherscan.io/address/0x30770d7E3e71112d7A6b7259542D1f680a70e315) | [`0xCc81...72c1`](https://holesky.etherscan.io/address/0xCc817FfA0C2AdED97DdEDF309b52Aea252f772c1) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/core/AVSDirectory.sol) | [`0x055733000064333CaDDbC92763c58BF0192fFeBf`](https://holesky.etherscan.io/address/0x055733000064333CaDDbC92763c58BF0192fFeBf) | [`0x331e...2506`](https://holesky.etherscan.io/address/0x331e3Bd9cf69562f6F9ade72BD3D9f271f1B2506) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/core/RewardsCoordinator.sol) | [`0xAcc1fb458a1317E886dB376Fc8141540537E68fE`](https://holesky.etherscan.io/address/0xAcc1fb458a1317E886dB376Fc8141540537E68fE) | [`0x4087...d41`](https://holesky.etherscan.io/address/0x40873dddA165258E1c5aE487f4842dac3946Ad41) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob//v1.5.0-rc.0/src/contracts/core/AllocationManager.sol) | [`0x78469728304326CBc65f8f95FA756B0B73164462`](https://holesky.etherscan.io/address/0x78469728304326CBc65f8f95FA756B0B73164462) | [`0x5912...A04Ea`](https://holesky.etherscan.io/address/0x5912005643201649542F5AE6CCE96C12b4DA04Ea) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/permissions/PermissionController.sol) | [`0x598cb226B591155F767dA17AfE7A2241a68C5C10`](https://holesky.etherscan.io/address/0x598cb226B591155F767dA17AfE7A2241a68C5C10) | [`0x7ab0...a2b9`](https://holesky.etherscan.io/address/0x7ab0ebd25d5ffe7527600ca5b2858c1a3faba2b9#code) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Slashing

These contracts handle the burning/redistribution of slashed funds. The `SlashEscrowFactory` is upgradeable by the `SlashEscrowProxyAdmin`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`OZ Proxy Admin (SlashEscrowProxyAdmin)`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x18dc...966b`](https://holesky.etherscan.io/address/0x0AA4F4791872211374d5912B67F5673E757CE430) | |
| [`SlashEscrowFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/SlashEscrowFactory.sol) | [`0xcc444eccD13E29033A46D3cbd4d30a2f70c10cbe`](https://holesky..etherscan.io/address/0xA5022befe84Ad0f5aAdc12e9c59230bc076083A5) | [`0xB643...348B`](https://holesky.etherscan.io/address/0xB64333C42F3c187744ad9F5d317C243A7788348B) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`SlashEscrow (Clone Implementation)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/SlashEscrow.sol) | - | [`0xa84b...ab2d`](https://holesky.etherscan.io/address/0x9c4cAc1e205cB33B4596E4f612eFdFDAe278A9CC) | [`EIP-1167 Clone`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/ClonesUpgradeable.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/strategies/StrategyFactory.sol) | [`0x9c01252B580efD11a05C00Aa42Dd3ac1Ec52DF6d`](https://holesky.etherscan.io/address/0x9c01252B580efD11a05C00Aa42Dd3ac1Ec52DF6d) | [`0x84aa...a7d`](https://holesky.etherscan.io/address/0x84aaD0F753b84Cd68F36Ff207DDfA0f2865b1a7d) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/strategies/StrategyBase.sol) | [`0xd3c6C6BA4E40dB9288c6a2077e5635344F8aFA4F`](https://holesky.etherscan.io/address/0xd3c6C6BA4E40dB9288c6a2077e5635344F8aFA4F) | [`0x0xbD16...3427`](https://holesky.etherscan.io/address/0xbD161189bCdb7d948761D74A84E34AB193B83427) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

The following strategies were originally deployed and whitelisted outside of the `StrategyFactory`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyBase (stETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3`](https://holesky.etherscan.io/address/0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (rETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0`](https://holesky.etherscan.io/address/0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (WETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9`](https://holesky.etherscan.io/address/0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (lsETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943`](https://holesky.etherscan.io/address/0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (sfrxETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c`](https://holesky.etherscan.io/address/0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ETHx)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x31B6F59e1627cEfC9fA174aD03859fC337666af7`](https://holesky.etherscan.io/address/0x31B6F59e1627cEfC9fA174aD03859fC337666af7) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (osETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef`](https://holesky.etherscan.io/address/0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (cbETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6`](https://holesky.etherscan.io/address/0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (mETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xaccc5A86732BE85b5012e8614AF237801636F8e5`](https://holesky.etherscan.io/address/0xaccc5A86732BE85b5012e8614AF237801636F8e5) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ankrETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac`](https://holesky.etherscan.io/address/0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (reALT)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xAD76D205564f955A9c18103C4422D1Cd94016899`](https://holesky.etherscan.io/address/0xAD76D205564f955A9c18103C4422D1Cd94016899) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (EO)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x78dBcbEF8fF94eC7F631c23d38d197744a323868`](https://holesky.etherscan.io/address/0x78dBcbEF8fF94eC7F631c23d38d197744a323868) | [`0x8037...3bc01`](https://holesky.etherscan.io/address/0x8037874C44Be3fe27Fa45c415163D38Ceab3bc01) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/EigenStrategy.sol) | [`0x43252609bff8a13dFe5e057097f2f45A24387a84`](https://holesky.etherscan.io/address/0x43252609bff8a13dFe5e057097f2f45A24387a84) | [`0x0x6377...8459`](https://holesky.etherscan.io/address/0x63772B1e4965a40dD658Ad824F841378CF518459) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0-rc.0/src/contracts/pods/EigenPod.sol) | [`0x7261C2bd75a7ACE1762f6d7FAe8F63215581832D`](https://holesky.etherscan.io/address/0x7672F1f72d30bBf5A6781aC183EDCCA5e3003AAD) | [`0x2742...0aD9`](https://holesky.etherscan.io/address/0x2742bd43b254f0199Bb5020CB477Ff9Ad3A80aD9) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0-rc.0/src/contracts/token/Eigen.sol) | [`0x3B78576F7D4230A49bE2c915629b31122C3FbF88`](https://holesky.etherscan.io/address/0x3B78576F7D4230A49bE2c915629b31122C3FbF88) | [`0xfF1e...468F0`](https://holesky.etherscan.io/address/0xfF1e23C37EC543684cf866785c0626c2Ac7468F0) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/token/BackingEigen.sol) | [`0x275cCf9Be51f4a6C94aBa6114cdf2a4c45B9cb27`](https://holesky.etherscan.io/address/0x275cCf9Be51f4a6C94aBa6114cdf2a4c45B9cb27) | [`0x05ad...E05c`](https://holesky.etherscan.io/address/0x05adA1C66DdDD7c36705bC23a4d50dBa72E4E05c) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/permissions/PauserRegistry.sol) | - | [`0x41Db...ec1D`](https://holesky.etherscan.io/address/0x41Dbe7BbacA97D986FCF6f5203b98Ec02412ec1D) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0x5e83c7d195318A5acf46B29E5810DdC323b2F6fD`](https://holesky.etherscan.io/address/0x5e83c7d195318A5acf46B29E5810DdC323b2F6fD) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xDB023566064246399b4AE851197a97729C93A6cf`](https://holesky.etherscan.io/address/0xDB023566064246399b4AE851197a97729C93A6cf) | |
| [`Pauser Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x53410249ec7d3a3F9F1ba3912D50D6A3Df6d10A7`](https://holesky.etherscan.io/address/0x53410249ec7d3a3F9F1ba3912D50D6A3Df6d10A7) | [`0xd9db...9552`](https://holesky.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Community Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xCb8d2f9e55Bc7B1FA9d089f9aC80C583D2BDD5F7`](https://holesky.etherscan.io/address/0xCb8d2f9e55Bc7B1FA9d089f9aC80C583D2BDD5F7) | [`0xd9db...9552`](https://holesky.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Executor Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x28Ade60640fdBDb2609D8d8734D1b5cBeFc0C348`](https://holesky.etherscan.io/address/0x28Ade60640fdBDb2609D8d8734D1b5cBeFc0C348) | [`0xd9db...9552`](https://holesky.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Operations Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xfaEF7338b7490b9E272d80A1a39f4657cAf2b97d`](https://holesky.etherscan.io/address/0xfaEF7338b7490b9E272d80A1a39f4657cAf2b97d) | [`0xd9db...9552`](https://holesky.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |

</details>



<details>
    <summary>Testnet Sepolia</summary>


You can view the deployed contract addresses below, or check out the code itself on the [`testnet-sepolia`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/testnet-sepolia) branch.

###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/DelegationManager.sol) | [`0xD4A7E1Bd8015057293f0D0A557088c286942e84b`](https://sepolia.etherscan.io/address/0xD4A7E1Bd8015057293f0D0A557088c286942e84b) | [`0x8a62...fdDb`](https://sepolia.etherscan.io/address/0x8a62419F507CF5FAFaCAfad1e0c20c49a0BbfdDb) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/StrategyManager.sol) | [`0x2E3D6c0744b10eb0A4e6F679F71554a39Ec47a5D`](https://sepolia.etherscan.io/address/0x2E3D6c0744b10eb0A4e6F679F71554a39Ec47a5D) | [`0x9942...6e42`](https://sepolia.etherscan.io/address/0x994213Ce37DB1d6AC856aD014d4392F60EB16e42) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/pods/EigenPodManager.sol) | [`0x56BfEb94879F4543E756d26103976c567256034a`](https://sepolia.etherscan.io/address/0x56BfEb94879F4543E756d26103976c567256034a) | [`0xB417...B499`](https://sepolia.etherscan.io/address/0xB417F83047DB5197D52AF36302a23b4878F2B499) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) | All EigenPod functionality is paused on Holesky | 
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/AVSDirectory.sol) | [`0xa789c91ECDdae96865913130B786140Ee17aF545`](https://sepolia.etherscan.io/address/0xa789c91ECDdae96865913130B786140Ee17aF545) | [`0xD88b...C188`](https://sepolia.etherscan.io/address/0xD88b96998325c3e74A74a0B0938BBFeA1395C188) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/RewardsCoordinator.sol) | [`0x5ae8152fb88c26ff9ca5C014c94fca3c68029349`](https://sepolia.etherscan.io/address/0x5ae8152fb88c26ff9ca5C014c94fca3c68029349) | [`0xcC30...7940`](https://sepolia.etherscan.io/address/0xcC305562B01bec562D13A40ef8781e313AFE7940) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/AllocationManager.sol) | [`0x42583067658071247ec8CE0A516A58f682002d07`](https://sepolia.etherscan.io/address/0x42583067658071247ec8CE0A516A58f682002d07) | [`0xb368...DAd6`](https://sepolia.etherscan.io/address/0xb36883818b5a4D25C409A81946DE9067cdC8DAd6) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/permissions/PermissionController.sol) | [`0x44632dfBdCb6D3E21EF613B0ca8A6A0c618F5a37`](https://sepolia.etherscan.io/address/0x44632dfBdCb6D3E21EF613B0ca8A6A0c618F5a37) | [`0x59B1...f525`](https://sepolia.etherscan.io/address/0x59B11b191B572888703E150E45F5015e0fFcf525) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Slashing

These contracts handle the burning/redistribution of slashed funds. The `SlashEscrowFactory` is upgradeable by the `SlashEscrowProxyAdmin`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`OZ Proxy Admin (SlashEscrowProxyAdmin)`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x18dc...966b`](https://sepolia.etherscan.io/address/0x18dc7D96d26b4F43ac464349D5D4af0310Ca966b) | `SlashEscrowFactory` proxy admin | 
| [`SlashEscrowFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/SlashEscrowFactory.sol) | [`0xA5022befe84Ad0f5aAdc12e9c59230bc076083A5`](https://sepolia.etherscan.io/address/0xA5022befe84Ad0f5aAdc12e9c59230bc076083A5) | [`0x7A0D...883E`](https://sepolia.etherscan.io/address/0x7A0D6553941BFc3864E5EEdEa7B2d9EA6Eb5883E) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) | - |
| [`SlashEscrow`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/SlashEscrow.sol) | - | [`0xa84b596F9456f473AD3241431fde8C135a63ab2d`](https://sepolia.etherscan.io/address/0xa84b596F9456f473AD3241431fde8C135a63ab2d) | [`EIP-1167 Clone`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/ClonesUpgradeable.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyFactory.sol) | [`0x066cF95c1bf0927124DFB8B02B401bc23A79730D`](https://sepolia.etherscan.io/address/0x066cF95c1bf0927124DFB8B02B401bc23A79730D) | [`0xEE41...ca1A`](https://sepolia.etherscan.io/address/0xEE41826B7D5B89e7F5eED6a831b4eFD69FC9ca1A) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBase.sol) | [`0x427e627Bc7E83cac0f84337d3Ad94230C32697D3`](https://sepolia.etherscan.io/address/0x427e627Bc7E83cac0f84337d3Ad94230C32697D3) | [`0x427e...97D3`](https://sepolia.etherscan.io/address/0x427e627Bc7E83cac0f84337d3Ad94230C32697D3) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| [`StETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBase.sol) | [`0x8b29d91e67b013e855EaFe0ad704aC4Ab086a574`](https://sepolia.etherscan.io/address/0x8b29d91e67b013e855EaFe0ad704aC4Ab086a574) | [`0x427e...97D3`](https://sepolia.etherscan.io/address/0x427e627Bc7E83cac0f84337d3Ad94230C32697D3) | Strategy Factory deployed |
| [`WETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBase.sol) | [`0x424246eF71b01ee33aA33aC590fd9a0855F5eFbc`](https://sepolia.etherscan.io/address/0x424246eF71b01ee33aA33aC590fd9a0855F5eFbc) | [`0x427e...97D3`](https://sepolia.etherscan.io/address/0x427e627Bc7E83cac0f84337d3Ad94230C32697D3) | Strategy Factory deployed |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/EigenStrategy.sol) | [`0x8E93249a6C37a32024756aaBd813E6139b17D1d5`](https://sepolia.etherscan.io/address/0x8E93249a6C37a32024756aaBd813E6139b17D1d5) | [`0xD2DC...6A68`](https://sepolia.etherscan.io/address/0xD2DCAf17966337fDb2AE586A3068C4Df2F516A68) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

**NOTE: Due to the permissioned validator set on Sepolia, all EigenPod functionality is *PAUSED*.**

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/pods/EigenPod.sol) | [`0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4`](https://sepolia.etherscan.io/address/0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4) | [`0xd85d...6bCf`](https://sepolia.etherscan.io/address/0xd85d0D9e24dC9af8a517034Caab2db68aD936bCf) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/token/Eigen.sol) | [`0x0011FA2c512063C495f77296Af8d195F33A8Dd38`](https://sepolia.etherscan.io/address/0x0011FA2c512063C495f77296Af8d195F33A8Dd38) | [`0xF83a...8725`](https://sepolia.etherscan.io/address/0xF83a81117AE073B13ce70f37302392BA90F28725) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/token/BackingEigen.sol) | [`0xc5B857A92245f64e9D90cCc5b096Db82eB77eB5c`](https://sepolia.etherscan.io/address/0xc5B857A92245f64e9D90cCc5b096Db82eB77eB5c) | [`0x1298...3173`](https://sepolia.etherscan.io/address/0x12988B679AA497C30A8D1850eCC4Dc7700383173) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/permissions/PauserRegistry.sol) | - | [`0x63AA...20f3`](https://sepolia.etherscan.io/address/0x63AAe451780090f50Ad323aAEF155F63a29D20f3) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0x1BEF...1b5B`](https://sepolia.etherscan.io/address/0x1BEF05C7303d44e0E2FCD2A19d993eDEd4c51b5B) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x56E8...6Fa1`](https://sepolia.etherscan.io/address/0x56E88cb4f0136fC27D95499dE4BE2acf47946Fa1) | |
| [`Pauser Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x0B415f75980D863872C3eb8caa76E6eC8Bc81536`](https://sepolia.etherscan.io/address/0x0B415f75980D863872C3eb8caa76E6eC8Bc81536) | [`0x4167...461a`](https://holesky.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Community Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x6f8459810197cc9fE123BBeB918451757a4fBAc6`](https://sepolia.etherscan.io/address/0x6f8459810197cc9fE123BBeB918451757a4fBAc6) | [`0x4167...461a`](https://holesky.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Executor Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x4FDA8998EC3b7d4b4A612d45FeB8fB36734470f2`](https://sepolia.etherscan.io/address/0x4FDA8998EC3b7d4b4A612d45FeB8fB36734470f2) | [`0x4167...461a`](https://holesky.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Operations Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xb094Ba769b4976Dc37fC689A76675f31bc4923b0`](https://sepolia.etherscan.io/address/0xb094Ba769b4976Dc37fC689A76675f31bc4923b0) | [`0x4167...461a`](https://holesky.etherscan.io/address/0x41675C099F32341bf84BFc5382aF534df5C7461a) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |

</details>



<details>
    <summary>Testnet Hoodi</summary>


###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/DelegationManager.sol) | [`0x867837a9722C512e0862d8c2E15b8bE220E8b87d`](https://hoodi.etherscan.io/address/0x867837a9722C512e0862d8c2E15b8bE220E8b87d) | [`0xC908...bfa0`](https://hoodi.etherscan.io/address/0xC908fAFAE29B5C9F0b5E0Da1d3025b8d6D42bfa0) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/StrategyManager.sol) | [`0xeE45e76ddbEDdA2918b8C7E3035cd37Eab3b5D41`](https://hoodi.etherscan.io/address/0xeE45e76ddbEDdA2918b8C7E3035cd37Eab3b5D41) | [`0x56Bf...034a`](https://hoodi.etherscan.io/address/0x56BfEb94879F4543E756d26103976c567256034a) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/pods/EigenPodManager.sol) | [`0xcd1442415Fc5C29Aa848A49d2e232720BE07976c`](https://hoodi.etherscan.io/address/0xcd1442415Fc5C29Aa848A49d2e232720BE07976c) | [`0x0A98...c49d`](https://hoodi.etherscan.io/address/0x0A987C508b0f56154CA534b7Fa5b84863cbcc49d) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/core/AVSDirectory.sol) | [`0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926`](https://hoodi.etherscan.io/address/0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926) | [`0xd290...5d5c`](https://hoodi.etherscan.io/address/0xd2905B858cA5Ded115B61dd9E98F7dcF9aEE2d5c) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/core/RewardsCoordinator.sol) | [`0x29e8572678e0c272350aa0b4B8f304E47EBcd5e7`](https://hoodi.etherscan.io/address/0x29e8572678e0c272350aa0b4B8f304E47EBcd5e7) | [`0xe786...2832`](https://hoodi.etherscan.io/address/0xe786FD0dE8a6001772386700318187Dc438a2832) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/AllocationManager.sol) | [`0x95a7431400F362F3647a69535C5666cA0133CAA0`](https://hoodi.etherscan.io/address/0x95a7431400F362F3647a69535C5666cA0133CAA0) | [`0x5ae8...9349`](https://hoodi.etherscan.io/address/0x5ae8152fb88c26ff9ca5C014c94fca3c68029349) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/permissions/PermissionController.sol) | [`0xdcCF401fD121d8C542E96BC1d0078884422aFAD2`](https://hoodi.etherscan.io/address/0xdcCF401fD121d8C542E96BC1d0078884422aFAD2) | [`0x2D73...eA27`](https://hoodi.etherscan.io/address/0x2D731E7993a100afd19454B98eEEC7b90366eA27) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Slashing

These contracts handle the burning/redistribution of slashed funds. The `SlashEscrowFactory` is upgradeable by the `SlashEscrowProxyAdmin`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`OZ Proxy Admin (SlashEscrowProxyAdmin)`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xa789...F545`](https://hoodi.etherscan.io/address/0xa789c91ECDdae96865913130B786140Ee17aF545) | |
| [`SlashEscrowFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/SlashEscrowFactory.sol) | [`0x885C0CC8118E428a2C04de58A93eB15Ed4F0e064`](https://hoodi..etherscan.io/address/0x885C0CC8118E428a2C04de58A93eB15Ed4F0e064) | [`0x4258...2d07`](https://hoodi.etherscan.io/address/0x42583067658071247ec8CE0A516A58f682002d07) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`SlashEscrow (Clone Implementation)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/core/SlashEscrow.sol) | - | [`0x889B...420d`](https://hoodi.etherscan.io/address/0x889B040116f453D89e9d6d692Ad70Edd7357420d) | [`EIP-1167 Clone`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/ClonesUpgradeable.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/strategies/StrategyFactory.sol) | [`0xfB7d94501E4d4ACC264833Ef4ede70a11517422B`](https://hoodi.etherscan.io/address/0xfB7d94501E4d4ACC264833Ef4ede70a11517422B) | [`0x798E...7216`](https://hoodi.etherscan.io/address/0x798EB817B7C109c6780264D5161183809C817216) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBase.sol) | [`0x6d28cEC1659BC3a9BC814c3EFc1412878B406579`](https://hoodi.etherscan.io/address/0x6d28cEC1659BC3a9BC814c3EFc1412878B406579) | [`0xB4ba...6C67`](https://hoodi.etherscan.io/address/0xB4baAfee917fb4449f5ec64804217bccE9f46C67) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| [`StETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBase.sol) | [`0x41525d38f59895d606e8c05c229864f2db6b64fd`](https://hoodi.etherscan.io/address/0x41525d38f59895d606e8c05c229864f2db6b64fd) | [`0xB4ba...6C67`](https://hoodi.etherscan.io/address/0xB4baAfee917fb4449f5ec64804217bccE9f46C67) | Strategy Factory deployed |
| [`WETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0-rc.0/src/contracts/strategies/StrategyBase.sol) | [`0x24579ad4fe83ac53546e5c2d3df5f85d6383420d`](https://hoodi.etherscan.io/address/0x24579ad4fe83ac53546e5c2d3df5f85d6383420d) | [`0xB4ba...6C67`](https://hoodi.etherscan.io/address/0xB4baAfee917fb4449f5ec64804217bccE9f46C67) | Strategy Factory deployed |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/strategies/EigenStrategy.sol) | [`0xB27b10291DBFE6576d17afF3e251c954Ae14f1D3`](https://hoodi.etherscan.io/address/0xB27b10291DBFE6576d17afF3e251c954Ae14f1D3) | [`0x09A4...f929`](https://hoodi.etherscan.io/address/0x09A4c2a9C3Ac0AF6E89BD768220493eBF2F4f929) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0-rc.0/src/contracts/pods/EigenPod.sol) | [`0x5e1577f8efB21b29cD5Eb4C5Aa3d6C4b228f650`](https://hoodi.etherscan.io/address/0x5e1577f8efB21b29cD5Eb4C5Aa3d6C4b228f650) | [`0x0e19...cAcf4`](https://hoodi.etherscan.io/address/0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0-rc.0/src/contracts/token/Eigen.sol) | [`0x8ae2520954db7D80D66835cB71E692835bbA45bf`](https://hoodi.etherscan.io/address/0x8ae2520954db7D80D66835cB71E692835bbA45bf) | [`0x63AA...20f3`](https://hoodi.etherscan.io/address/0x63AAe451780090f50Ad323aAEF155F63a29D20f3) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/token/BackingEigen.sol) | [`0x6e60888132Cc7e637488379B4B40c42b3751f63a`](https://hoodi.etherscan.io/address/0x6e60888132Cc7e637488379B4B40c42b3751f63a) | [`0x43e4...a1C4`](https://hoodi.etherscan.io/address/0x43e4940aCeb1C1F5a57e307EEB212007F0f6a1C4) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/permissions/PauserRegistry.sol) | - | [`0x64D7...c13D`](https://hoodi.etherscan.io/address/0x64D78399B0fa32EA72959f33edCF313159F3c13D) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0xE332...2d15`](https://hoodi.etherscan.io/address/0xE3328cb5068924119d6170496c4AB2dD12c12d15) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xE7f4...2257`](https://hoodi.etherscan.io/address/0xE7f4E30D2619273468afe9EC0Acf805E55532257) | |

</details>

## Branching

Branches we use:
* `main`: The canonical, most up-to-date branch, containing the work-in-progress code for upcoming releases
* `Vx.y.z`: Release branch with version `x.y.z` that matches a release of EigenLayer, release branch is always cut from `main` via cherry-picking
* `release-dev/xxx`: A development branch for a large feature to be released, the branch should eventually be deleted after merge to `main`

## Building Apps & Services on EigenLayer

### Developer & User Basics

[Our documentation](https://docs.eigenlayer.xyz/) is a great place to start to understand EigenLayer and the Eigen ecosystem. If you are an audio/visual learner, check out the "[You Could've Invented EigenLayer](https://www.blog.eigenlayer.xyz/ycie/)" video. 

To understand more about how our ecosystem and its participants work, checkout the guides below. These are split out by who you are: AVS developers, restakers and Operators. All have different roles and interactioons with EigenLayer:

* [AVS Developer Guide](https://docs.eigenlayer.xyz/developers/Concepts/avs-developer-guide)
* [Operator Guide](https://docs.eigenlayer.xyz/operators/concepts/operator-introduction)
* [Restaker Guide](https://docs.eigenlayer.xyz/restakers/concepts/overview)

### Contract Docs & Deep Dive

The most up-to-date and technical documentation on our core contracts can be found in [/docs](/docs). If you're a shadowy super coder, this is a great place to get an overview of the contracts before diving into the code.

To learn more about interfacing with the EigenLayer core contracts onchain, see our [middleware repo.](https://github.com/Layr-Labs/eigenlayer-middleware)

To get an idea of how users interact with these contracts, check out our integration tests: [/src/test/integration](./src/test/integration/).

## Contribute to the Core

See [CONTRIBUTING](CONTRIBUTING.md). 

Contributions that do not follow our fork base PR practices will be either rejected or immediately deleted based on your role, preventing branch pollution, keeping our repository clean, make it more readable and searchable.
