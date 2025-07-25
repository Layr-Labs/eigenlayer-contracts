<a name="introduction"/></a>

# EigenLayer

**EigenLayer** is a protocol built on Ethereum that introduces Restaking, a primitive for app and service builders to make verifiable commitments to their users.

EigenLayer brings together Restakers, Operators, and Autonomous Verifiable Services (AVSs) to extend Ethereum's cryptoeconomic security with penalty and reward commitments (like slashing) on staked assets acting as security. The protocol supports permissionless security; EIGEN, Native ETH, LSTs, and ERC-20s. 

## Deployments

The deployments on `mainnet`, `holesky`, `sepolia`, `hoodi`, and `base sepolia` are on the below versions:

| Environment | Version | Core Protocol Deployed | Supports Native Restaking | Supports Multichain |
| -------- | -------- | -------- | -------- | -------- | 
| Mainnet Ethereum | [`v1.6.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.6.0) | Yes | Yes | No (pending) |
| Testnet Holesky | [`v1.6.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.6.0) | Yes | Yes | No |
| Testnet Sepolia | [`v1.7.0-rc.4`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.7.0-rc.4) | Yes | No | Yes (source & destination) |
| Testnet Hoodi | [`v1.6.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.6.0) | Yes | Yes | No |
| Testnet Base Sepolia | [`v1.7.0-rc.4`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.7.0-rc.4) | No | No | Yes (destination) |

### Current Deployment Contracts

<details>
    <summary>Mainnet Ethereum</summary>


###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/core/DelegationManager.sol) | [`0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A`](https://etherscan.io/address/0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A) | [`0x6EEd...F89`](https://etherscan.io/address/0x6EEd6c2802dF347e05884857CdDB2D3E96D12F89) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/core/StrategyManager.sol) | [`0x858646372CC42E1A627fcE94aa7A7033e7CF075A`](https://etherscan.io/address/0x858646372CC42E1A627fcE94aa7A7033e7CF075A) | [`0x46ae...d123`](https://etherscan.io/address/0x46aefd30415be99e20169eE7046F65784B46d123) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/pods/EigenPodManager.sol) | [`0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338`](https://etherscan.io/address/0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338) | [`0xE48D...ABd`](https://etherscan.io/address/0xE48D7CaeC1790b293667e4bB2dE1E00536F2bABd) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/AVSDirectory.sol) | [`0x135dda560e946695d6f155dacafc6f1f25c1f5af`](https://etherscan.io/address/0x135dda560e946695d6f155dacafc6f1f25c1f5af) | [`0xA396...B6A2`](https://etherscan.io/address/0xA396D855D70e1A1ec1A0199ADB9845096683B6A2) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/RewardsCoordinator.sol) | [`0x7750d328b314EfFa365A0402CcfD489B80B0adda`](https://etherscan.io/address/0x7750d328b314EfFa365A0402CcfD489B80B0adda) | [`0xa505...0aB`](https://etherscan.io/address/0xa505c0116aD65071F0130061F94745b7853220aB) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/permissions/PermissionController.sol) | [`0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5`](https://etherscan.io/address/0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5) | [`0xe7f3...C6B1`](https://etherscan.io/address/0xe7f3705c9Addf2DE14e03C345fA982CAb2c1C6B1) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/core/AllocationManager.sol) | [`0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39`](https://etherscan.io/address/0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39) | [`0xC976...345`](https://etherscan.io/address/0xC97602648fA52F92B4ee2b0e5a54Bd15b6cB0345) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyFactory.sol) | [`0x5e4C39Ad7A3E881585e383dB9827EB4811f6F647`](https://etherscan.io/address/0x5e4C39Ad7A3E881585e383dB9827EB4811f6F647) | [`0x1b97...c66`](https://etherscan.io/address/0x1b97d8F963179C0e17E5F3d85cdfd9a31A49bc66) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBase.sol) | [`0x0ed6703C298d28aE0878d1b28e88cA87F9662fE9`](https://etherscan.io/address/0x0ed6703c298d28ae0878d1b28e88ca87f9662fe9) | [`0xD4d1...E5B`](https://etherscan.io/address/0xD4d1746142642Db4c1ab17b03B9c58baac913E5B) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

The following strategies were originally deployed and whitelisted outside of the `StrategyFactory`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyBase (cbETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc`](https://etherscan.io/address/0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (stETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x93c4b944D05dfe6df7645A86cd2206016c51564D`](https://etherscan.io/address/0x93c4b944D05dfe6df7645A86cd2206016c51564D) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (rETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2`](https://etherscan.io/address/0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ETHx)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d`](https://etherscan.io/address/0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ankrETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff`](https://etherscan.io/address/0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (OETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xa4C637e0F704745D182e4D38cAb7E7485321d059`](https://etherscan.io/address/0xa4C637e0F704745D182e4D38cAb7E7485321d059) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (osETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x57ba429517c3473B6d34CA9aCd56c0e735b94c02`](https://etherscan.io/address/0x57ba429517c3473B6d34CA9aCd56c0e735b94c02) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (swETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6`](https://etherscan.io/address/0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (wBETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7CA911E83dabf90C90dD3De5411a10F1A6112184`](https://etherscan.io/address/0x7CA911E83dabf90C90dD3De5411a10F1A6112184) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (sfrxETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6`](https://etherscan.io/address/0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (lsETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xAe60d8180437b5C34bB956822ac2710972584473`](https://etherscan.io/address/0xAe60d8180437b5C34bB956822ac2710972584473) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (mETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x298aFB19A105D59E74658C4C334Ff360BadE6dd2`](https://etherscan.io/address/0x298aFB19A105D59E74658C4C334Ff360BadE6dd2) | [`0x62F7...f4E0`](https://etherscan.io/address/0x62F7226Fb9d615590EadB539713b250fB2fdf4E0) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/strategies/EigenStrategy.sol) | [`0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7`](https://etherscan.io/address/0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7) | [`0x530f...6be`](https://etherscan.io/address/0x530fDB7AdF7d489DF49c27e3d3512c0dD64886be) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/pods/EigenPod.sol) | [`0x5a2a4F2F3C18f09179B6703e63D9eDD165909073`](https://etherscan.io/address/0x5a2a4F2F3C18f09179B6703e63D9eDD165909073) | [`0xcb27...3745`](https://etherscan.io/address/0xcb27A4819A64FBA93ABD4D480e4466aEc0503745) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/token/Eigen.sol) | [`0xec53bf9167f50cdeb3ae105f56099aaab9061f83`](https://etherscan.io/address/0xec53bf9167f50cdeb3ae105f56099aaab9061f83) | [`0x2C4A...2E50`](https://etherscan.io/address/0x2C4A81e257381F87F5A5C4bd525116466D972E50) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
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
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/DelegationManager.sol) | [`0xA44151489861Fe9e3055d95adC98FbD462B948e7`](https://holesky.etherscan.io/address/0xA44151489861Fe9e3055d95adC98FbD462B948e7) | [`0x39F2...b0CF4`](https://holesky.etherscan.io/address/0x39F23a044Ea7DB60b5392b964F2f9261B93b0CF4) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/StrategyManager.sol) | [`0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6`](https://holesky.etherscan.io/address/0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6) | [`0x4870...21a04`](https://holesky.etherscan.io/address/0x4870B0BbDD8360E49A15987B16590Ef917c21a04) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/pods/EigenPodManager.sol) | [`0x30770d7E3e71112d7A6b7259542D1f680a70e315`](https://holesky.etherscan.io/address/0x30770d7E3e71112d7A6b7259542D1f680a70e315) | [`0xf889...23AA`](https://holesky.etherscan.io/address/0xf889B045cdE04c9CB22FF0Ead6C66a72f3cb23AA) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/core/AVSDirectory.sol) | [`0x055733000064333CaDDbC92763c58BF0192fFeBf`](https://holesky.etherscan.io/address/0x055733000064333CaDDbC92763c58BF0192fFeBf) | [`0x331e...2506`](https://holesky.etherscan.io/address/0x331e3Bd9cf69562f6F9ade72BD3D9f271f1B2506) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/core/RewardsCoordinator.sol) | [`0xAcc1fb458a1317E886dB376Fc8141540537E68fE`](https://holesky.etherscan.io/address/0xAcc1fb458a1317E886dB376Fc8141540537E68fE) | [`0x4087...d41`](https://holesky.etherscan.io/address/0x40873dddA165258E1c5aE487f4842dac3946Ad41) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/AllocationManager.sol) | [`0x78469728304326CBc65f8f95FA756B0B73164462`](https://holesky.etherscan.io/address/0x78469728304326CBc65f8f95FA756B0B73164462) | [`0x45A1...A591`](https://holesky.etherscan.io/address/0x45A1bb44e2d8d2475185d0bcE480ce2e55c4A591) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/permissions/PermissionController.sol) | [`0x598cb226B591155F767dA17AfE7A2241a68C5C10`](https://holesky.etherscan.io/address/0x598cb226B591155F767dA17AfE7A2241a68C5C10) | [`0x7ab0...a2b9`](https://holesky.etherscan.io/address/0x7ab0ebd25d5ffe7527600ca5b2858c1a3faba2b9#code) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.2/src/contracts/strategies/StrategyFactory.sol) | [`0x9c01252B580efD11a05C00Aa42Dd3ac1Ec52DF6d`](https://holesky.etherscan.io/address/0x9c01252B580efD11a05C00Aa42Dd3ac1Ec52DF6d) | [`0x84aa...a7d`](https://holesky.etherscan.io/address/0x84aaD0F753b84Cd68F36Ff207DDfA0f2865b1a7d) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBase.sol) | [`0xd3c6C6BA4E40dB9288c6a2077e5635344F8aFA4F`](https://holesky.etherscan.io/address/0xd3c6C6BA4E40dB9288c6a2077e5635344F8aFA4F) | [`0x62cd...B86d`](https://holesky.etherscan.io/address/0x62cda01F3FbaB32FEe9E08094D9766D77566B86d) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

The following strategies were originally deployed and whitelisted outside of the `StrategyFactory`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyBase (stETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3`](https://holesky.etherscan.io/address/0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (rETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0`](https://holesky.etherscan.io/address/0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (WETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9`](https://holesky.etherscan.io/address/0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (lsETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943`](https://holesky.etherscan.io/address/0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (sfrxETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c`](https://holesky.etherscan.io/address/0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ETHx)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x31B6F59e1627cEfC9fA174aD03859fC337666af7`](https://holesky.etherscan.io/address/0x31B6F59e1627cEfC9fA174aD03859fC337666af7) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (osETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef`](https://holesky.etherscan.io/address/0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (cbETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6`](https://holesky.etherscan.io/address/0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (mETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xaccc5A86732BE85b5012e8614AF237801636F8e5`](https://holesky.etherscan.io/address/0xaccc5A86732BE85b5012e8614AF237801636F8e5) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ankrETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac`](https://holesky.etherscan.io/address/0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (reALT)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xAD76D205564f955A9c18103C4422D1Cd94016899`](https://holesky.etherscan.io/address/0xAD76D205564f955A9c18103C4422D1Cd94016899) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (EO)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x78dBcbEF8fF94eC7F631c23d38d197744a323868`](https://holesky.etherscan.io/address/0x78dBcbEF8fF94eC7F631c23d38d197744a323868) | [`0x332b...B518`](https://holesky.etherscan.io/address/0x332b384D7100bc8A6D1d1B99fa5D60834ac4B518) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/EigenStrategy.sol) | [`0x43252609bff8a13dFe5e057097f2f45A24387a84`](https://holesky.etherscan.io/address/0x43252609bff8a13dFe5e057097f2f45A24387a84) | [`0xd673...2b99`](https://holesky.etherscan.io/address/0xd673a9Ee9DD278aED0705161b9D264e3F2296b99) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/pods/EigenPod.sol) | [`0x7261C2bd75a7ACE1762f6d7FAe8F63215581832D`](https://holesky.etherscan.io/address/0x7261C2bd75a7ACE1762f6d7FAe8F63215581832D) | [`0x2742...0aD9`](https://holesky.etherscan.io/address/0x2742bd43b254f0199Bb5020CB477Ff9Ad3A80aD9) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/token/Eigen.sol) | [`0x3B78576F7D7230049bE2c915629b31122C3FbF88`](https://holesky.etherscan.io/address/0x3B78576F7D7230049bE2c915629b31122C3FbF88) | [`0xfF1e...68F0`](https://holesky.etherscan.io/address/0xfF1e23C37EC543684cf866785c0626c2Ac7468F0) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
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
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/DelegationManager.sol) | [`0xD4A7E1Bd8015057293f0D0A557088c286942e84b`](https://sepolia.etherscan.io/address/0xD4A7E1Bd8015057293f0D0A557088c286942e84b) | [`0xa982...1b823`](https://sepolia.etherscan.io/address/0xa9821F0620D347648f375c597761C7FD16C1b823) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/StrategyManager.sol) | [`0x2E3D6c0744b10eb0A4e6F679F71554a39Ec47a5D`](https://sepolia.etherscan.io/address/0x2E3D6c0744b10eb0A4e6F679F71554a39Ec47a5D) | [`0xf0A8...FCB3`](https://sepolia.etherscan.io/address/0xf0A8735c26121e6C488ebac65c8fa3fe37cBFCB3) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/pods/EigenPodManager.sol) | [`0x56BfEb94879F4543E756d26103976c567256034a`](https://sepolia.etherscan.io/address/0x56BfEb94879F4543E756d26103976c567256034a) | [`0x1084...88Ae`](https://sepolia.etherscan.io/address/0x10848d61cE044bcAEE301a28A4C30Ad0a40e88Ae) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) | All EigenPod functionality is paused on Holesky | 
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/AVSDirectory.sol) | [`0xa789c91ECDdae96865913130B786140Ee17aF545`](https://sepolia.etherscan.io/address/0xa789c91ECDdae96865913130B786140Ee17aF545) | [`0xD88b...C188`](https://sepolia.etherscan.io/address/0xD88b96998325c3e74A74a0B0938BBFeA1395C188) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/RewardsCoordinator.sol) | [`0x5ae8152fb88c26ff9ca5C014c94fca3c68029349`](https://sepolia.etherscan.io/address/0x5ae8152fb88c26ff9ca5C014c94fca3c68029349) | [`0xcC30...7940`](https://sepolia.etherscan.io/address/0xcC305562B01bec562D13A40ef8781e313AFE7940) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/AllocationManager.sol) | [`0x42583067658071247ec8CE0A516A58f682002d07`](https://sepolia.etherscan.io/address/0x42583067658071247ec8CE0A516A58f682002d07) | [`0xB87A...7Db7`](https://sepolia.etherscan.io/address/0xB87AeeA46BB3a3050D277272E33b011922537Db7) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/permissions/PermissionController.sol) | [`0x44632dfBdCb6D3E21EF613B0ca8A6A0c618F5a37`](https://sepolia.etherscan.io/address/0x44632dfBdCb6D3E21EF613B0ca8A6A0c618F5a37) | [`0x59B1...f525`](https://sepolia.etherscan.io/address/0x59B11b191B572888703E150E45F5015e0fFcf525) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`KeyRegistrar`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/permissions/KeyRegistrar.sol) | [`0xA4dB30D08d8bbcA00D40600bee9F029984dB162a`](https://sepolia.etherscan.io/address/0xA4dB30D08d8bbcA00D40600bee9F029984dB162a) | [`0xe61E...87CD`](https://sepolia.etherscan.io/address/0xe61E638650DC91Fe0f59bBBf5315517337d687CD) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/strategies/StrategyFactory.sol) | [`0x066cF95c1bf0927124DFB8B02B401bc23A79730D`](https://sepolia.etherscan.io/address/0x066cF95c1bf0927124DFB8B02B401bc23A79730D) | [`0xEE41...ca1A`](https://sepolia.etherscan.io/address/0xEE41826B7D5B89e7F5eED6a831b4eFD69FC9ca1A) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBase.sol) | [`0x427e627Bc7E83cac0f84337d3Ad94230C32697D3`](https://sepolia.etherscan.io/address/0x427e627Bc7E83cac0f84337d3Ad94230C32697D3) | [`0x3BA4...a43b`](https://sepolia.etherscan.io/address/0x3BA4F6973D9ed2f0617C4b15d8CE06155fd6a43b) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| [`StETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBase.sol) | [`0x8b29d91e67b013e855EaFe0ad704aC4Ab086a574`](https://sepolia.etherscan.io/address/0x8b29d91e67b013e855EaFe0ad704aC4Ab086a574) | [`0x3BA4...a43b`](https://sepolia.etherscan.io/address/0x3BA4F6973D9ed2f0617C4b15d8CE06155fd6a43b) | Strategy Factory deployed |
| [`WETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBase.sol) | [`0x424246eF71b01ee33aA33aC590fd9a0855F5eFbc`](https://sepolia.etherscan.io/address/0x424246eF71b01ee33aA33aC590fd9a0855F5eFbc) | [`0x3BA4...a43b`](https://sepolia.etherscan.io/address/0x3BA4F6973D9ed2f0617C4b15d8CE06155fd6a43b) | Strategy Factory deployed |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/EigenStrategy.sol) | [`0x8E93249a6C37a32024756aaBd813E6139b17D1d5`](https://sepolia.etherscan.io/address/0x8E93249a6C37a32024756aaBd813E6139b17D1d5) | [`0x5351...52B7`](https://sepolia.etherscan.io/address/0x535185d24cD1B2e4F7afDe4ee39CF717c00B52B7) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

**NOTE: Due to the permissioned validator set on Sepolia, all EigenPod functionality is *PAUSED*.**

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/pods/EigenPod.sol) | [`0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4`](https://sepolia.etherscan.io/address/0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4) | [`0x1073...b40e`](https://sepolia.etherscan.io/address/0x107339987a46F77598cB563bf63eb8a1C9f5b40e) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/token/Eigen.sol) | [`0x0011FA2c512063C495f77296Af8d195F33A8Dd38`](https://sepolia.etherscan.io/address/0x0011FA2c512063C495f77296Af8d195F33A8Dd38) | [`0x7ec6...BD6F`](https://sepolia.etherscan.io/address/0x7ec6a02235Bf8d8a1fDB894AD2E1573192bfBD6F) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/token/BackingEigen.sol) | [`0xc5B857A92245f64e9D90cCc5b096Db82eB77eB5c`](https://sepolia.etherscan.io/address/0xc5B857A92245f64e9D90cCc5b096Db82eB77eB5c) | [`0x1298...3173`](https://sepolia.etherscan.io/address/0x12988B679AA497C30A8D1850eCC4Dc7700383173) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### AVS

The following contracts are be used by AVSs to manage and declare releases. 

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`ReleaseManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/core/ReleaseManager.sol) | [`0x59c8D715DCa616e032B744a753C017c9f3E16bf4`](https://sepolia.etherscan.io/address/0x59c8D715DCa616e032B744a753C017c9f3E16bf4) | [`0xAe14...3A15`](https://sepolia.etherscan.io/address/0xAe145081D4fd04Af9e0Ca38b35DC0113c3064A15) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multichain - Source

The multichain protocol expects AVSs to register on the source chain. AVS's stakes are then transported to supported destination chains. For the sepolia network, the destination chains are sepolia and base-sepolia. 

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`CrossChainRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/multichain/source/CrossChainRegistry.sol) | [`0x287381B1570d9048c4B4C7EC94d21dDb8Aa1352a`](https://sepolia.etherscan.io/address/0x287381B1570d9048c4B4C7EC94d21dDb8Aa1352a) | [`0x2071...8f48b`](https://sepolia.etherscan.io/address/0x20715838d4Cf054577cd2b89648D9cFC8994f48b) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multichain - Destination

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`OperatorTableUpdater`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/multichain/destination/OperatorTableUpdater.sol) | [`0xB02A15c6Bd0882b35e9936A9579f35FB26E11476`](https://sepolia.etherscan.io/address/0xB02A15c6Bd0882b35e9936A9579f35FB26E11476) | [`0xABbc...9473`](https://sepolia.etherscan.io/address/0xABbc13111F059D64A5BB9056399018c11D5C9473) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ECDSACertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/multichain/destination/ECDSACertificateVerifier.sol) | [`0xb3Cd1A457dEa9A9A6F6406c6419B1c326670A96F`](https://sepolia.etherscan.io/address/0xb3Cd1A457dEa9A9A6F6406c6419B1c326670A96F) | [`0x184D...3Eb4f`](https://sepolia.etherscan.io/address/0x184Dcf2c01e237644363925FBB0C0Dce88a3Eb4f) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`BN254CertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/multichain/destination/BN254CertificateVerifier.sol) | [`0xff58A373c18268F483C1F5cA03Cf885c0C43373a`](https://sepolia.etherscan.io/address/0xff58A373c18268F483C1F5cA03Cf885c0C43373a) | [`0x3B18...B41fC`](https://sepolia.etherscan.io/address/0x3B1842EB1e785b975316A4161044A76CaE6B41fC) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

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


You can view the deployed contract addresses below, or check out the code itself on the [`testnet-hoodi`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/testnet-hoodi) branch.

###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/DelegationManager.sol) | [`0x867837a9722C512e0862d8c2E15b8bE220E8b87d`](https://hoodi.etherscan.io/address/0x867837a9722C512e0862d8c2E15b8bE220E8b87d) | [`0xcC30...7940`](https://hoodi.etherscan.io/address/0xcC305562B01bec562D13A40ef8781e313AFE7940) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/StrategyManager.sol) | [`0xeE45e76ddbEDdA2918b8C7E3035cd37Eab3b5D41`](https://hoodi.etherscan.io/address/0xeE45e76ddbEDdA2918b8C7E3035cd37Eab3b5D41) | [`0x742A...27b`](https://hoodi.etherscan.io/address/0x742A228482701d693061BfE9C5B3Eb3959Ea927b) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/pods/EigenPodManager.sol) | [`0xcd1442415Fc5C29Aa848A49d2e232720BE07976c`](https://hoodi.etherscan.io/address/0xcd1442415Fc5C29Aa848A49d2e232720BE07976c) | [`0x59B1...f525`](https://hoodi.etherscan.io/address/0x59B11b191B572888703E150E45F5015e0fFcf525) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/AVSDirectory.sol) | [`0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926`](https://hoodi.etherscan.io/address/0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926) | [`0xd290...Ed5c`](https://hoodi.etherscan.io/address/0xd2905B858cA5Ded115B61dd9E98F7dcF9aEE2d5c) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/core/RewardsCoordinator.sol) | [`0x29e8572678e0c272350aa0b4B8f304E47EBcd5e7`](https://hoodi.etherscan.io/address/0x29e8572678e0c272350aa0b4B8f304E47EBcd5e7) | [`0xe786...832`](https://hoodi.etherscan.io/address/0xe786FD0dE8a6001772386700318187Dc438a2832) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/core/AllocationManager.sol) | [`0x95a7431400F362F3647a69535C5666cA0133CAA0`](https://hoodi.etherscan.io/address/0x95a7431400F362F3647a69535C5666cA0133CAA0) | [`0x8b1D...11a0`](https://hoodi.etherscan.io/address/0x8b1DBbAa79507CD6b6e1d9FBe90E3FB18EFf11a0) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/permissions/PermissionController.sol) | [`0xdcCF401fD121d8C542E96BC1d0078884422aFAD2`](https://hoodi.etherscan.io/address/0xdcCF401fD121d8C542E96BC1d0078884422aFAD2) | [`0x2D73...A27`](https://hoodi.etherscan.io/address/0x2D731E7993a100afd19454B98eEEC7b90366eA27) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |


###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/strategies/StrategyFactory.sol) | [`0xfB7d94501E4d4ACC264833Ef4ede70a11517422B`](https://hoodi.etherscan.io/address/0xfB7d94501E4d4ACC264833Ef4ede70a11517422B) | [`0x798E...7216`](https://hoodi.etherscan.io/address/0x798EB817B7C109c6780264D5161183809C817216) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBase.sol) | [`0x6d28cEC1659BC3a9BC814c3EFc1412878B406579`](https://hoodi.etherscan.io/address/0x6d28cEC1659BC3a9BC814c3EFc1412878B406579) | [`0x6514...2D6F`](https://hoodi.etherscan.io/address/0x65147e9916152C1EbdBaD8A6f3e145b4BDeE2D6F) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| [`StETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBase.sol) | [`0xf8a1a66130d614c7360e868576d5e59203475fe0`](https://hoodi.etherscan.io/address/0xf8a1a66130d614c7360e868576d5e59203475fe0) | [`0x6514...2D6F`](https://hoodi.etherscan.io/address/0x65147e9916152C1EbdBaD8A6f3e145b4BDeE2D6F) | Strategy Factory deployed |
| [`WETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/StrategyBase.sol) | [`0x24579aD4fe83aC53546E5c2D3dF5F85D6383420d`](https://hoodi.etherscan.io/address/0x24579aD4fe83aC53546E5c2D3dF5F85D6383420d) | [`0x6514...2D6F`](https://hoodi.etherscan.io/address/0x65147e9916152C1EbdBaD8A6f3e145b4BDeE2D6F) | Strategy Factory deployed |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.5.0/src/contracts/strategies/EigenStrategy.sol) | [`0xB27b10291DBFE6576d17afF3e251c954Ae14f1D3`](https://hoodi.etherscan.io/address/0xB27b10291DBFE6576d17afF3e251c954Ae14f1D3) | [`0xEE41...ca1A`](https://hoodi.etherscan.io/address/0xEE41826B7D5B89e7F5eED6a831b4eFD69FC9ca1A) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/pods/EigenPod.sol) | [`0x5e1577f8efB21b229cD5Eb4C5Aa3d6C4b228f650`](https://hoodi.etherscan.io/address/0x5e1577f8efB21b229cD5Eb4C5Aa3d6C4b228f650) | [`0x2D6c...2869`](https://hoodi.etherscan.io/address/0x2D6c7f9862BD80Cf0d9d93FC6b513D69E7Db7869) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.6.0/src/contracts/token/Eigen.sol) | [`0x8ae2520954db7D80D66835cB71E692835bbA45bf`](https://hoodi.etherscan.io/address/0x8ae2520954db7D80D66835cB71E692835bbA45bf) | [`0x68F3...2FE9`](https://hoodi.etherscan.io/address/0x68F3E0B02b15F96B50aC73d7680F020B646D2FE9) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.3.0/src/contracts/token/BackingEigen.sol) | [`0x6e60888132Cc7e637488379B4B40c42b3751f63a`](https://hoodi.etherscan.io/address/0x6e60888132Cc7e637488379B4B40c42b3751f63a) | [`0x43e4...1C4`](https://hoodi.etherscan.io/address/0x43e4940aCeb1C1F5a57e307EEB212007F0f6a1C4) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.4.1/src/contracts/permissions/PauserRegistry.sol) | - | [`0x64D7...c13D`](https://hoodi.etherscan.io/address/0x64D78399B0fa32EA72959f33edCF313159F3c13D) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0xE332...2d15`](https://hoodi.etherscan.io/address/0xE3328cb5068924119d6170496c4AB2dD12c12d15) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xE7f4...2257`](https://hoodi.etherscan.io/address/0xE7f4E30D2619273468afe9EC0Acf805E55532257) | |

</details>



<details>
    <summary>Testnet Base Sepolia</summary>


**Note: Testnet Base Sepolia only supports verification of tasks from stake that lives on Sepolia. Standard core protocol functionality (restaking, slashing) does not exist on Base Sepolia.**

###### Multichain - Destination

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`OperatorTableUpdater`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/multichain/destination/OperatorTableUpdater.sol) | [`0xB02A15c6Bd0882b35e9936A9579f35FB26E11476`](https://sepolia.basescan.org/address/0xB02A15c6Bd0882b35e9936A9579f35FB26E11476) | [`0x1D4d...3e17C`](https://sepolia.basescan.org/address/0x1D4d6054BD11A5711ad7c5d3E376C987a603e17C) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ECDSACertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/multichain/destination/ECDSACertificateVerifier.sol) | [`0xb3Cd1A457dEa9A9A6F6406c6419B1c326670A96F`](https://sepolia.basescan.org/address/0xb3Cd1A457dEa9A9A6F6406c6419B1c326670A96F) | [`0x3D8c...2A477`](https://sepolia.basescan.org/address/0x3D8c4Ac89040aa25168F32407274F239Ab52A477) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`BN254CertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/multichain/destination/BN254CertificateVerifier.sol) | [`0xff58A373c18268F483C1F5cA03Cf885c0C43373a`](https://sepolia.basescan.org/address/0xff58A373c18268F483C1F5cA03Cf885c0C43373a) | [`0x854d...f00b5`](https://sepolia.basescan.org/address/0x854dc9e5d011B060bf77B1a492302C349f2f00b5) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.7.0-rc.4/src/contracts/permissions/PauserRegistry.sol) | - | [`0x6ffE...672B`](https://sepolia.basescan.org/address/0x6ffE77C321a773e2A27B0B0a31C5e1bBda83672B) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xbe2F...0b5`](https://sepolia.basescan.org/address/0xbe2F96Efff467c6773Dc91eA62Ab34C73195a0b5) | |
| [`Pauser Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x2DD9BDDf299C846f61C0105905f9a60ee99242d6`](https://sepolia.basescan.org/address/0x2DD9BDDf299C846f61C0105905f9a60ee99242d6) | - | |
| [`Operations Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x4ACA7e09eF90612941A9185A6690Dbe9e39aae8f`](https://sepolia.basescan.org/address/0x4ACA7e09eF90612941A9185A6690Dbe9e39aae8f) | - | |
| [`Multichain Deployer Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xA591635DE4C254BD3fa9C9Db9000eA6488344C28`](https://sepolia.basescan.org/address/0xA591635DE4C254BD3fa9C9Db9000eA6488344C28) | - | |

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
