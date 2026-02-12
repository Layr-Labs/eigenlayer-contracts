<a name="introduction"/></a>

# EigenLayer

**EigenLayer** is a protocol built on Ethereum that introduces Restaking, a primitive for app and service builders to make verifiable commitments to their users.

EigenLayer brings together Restakers, Operators, and Autonomous Verifiable Services (AVSs) to extend Ethereum's cryptoeconomic security with penalty and reward commitments (like slashing) on staked assets acting as security. The protocol supports permissionless security; EIGEN, Native ETH, LSTs, and ERC-20s. 

## Deployments

The deployments on `mainnet`, `base`, `sepolia`, `hoodi`, and `base sepolia` are on the below versions:

| Environment | Version | Core Protocol Deployed | Supports Native Restaking | Supports Multichain |
| -------- | -------- | -------- | -------- | -------- | 
| Mainnet Ethereum | [`v1.9.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.9.0) | Yes | Yes | Yes (source & destination) |
| Base | [`v1.9.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.9.0) | No | No | Yes (destination) |
| Testnet Sepolia | [`v1.9.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.9.0) | Yes | No | Yes (source & destination) |
| Testnet Hoodi | [`v1.9.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.9.0) | Yes | Yes | No |
| Testnet Base Sepolia | [`v1.9.0`](https://github.com/Layr-Labs/eigenlayer-contracts/releases/tag/v1.9.0) | No | No | Yes (destination) |

### Current Deployment Contracts

<details>
    <summary>Mainnet Ethereum</summary>


###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/DelegationManager.sol) | [`0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A`](https://etherscan.io/address/0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A) | [`0xE702...cE75`](https://etherscan.io/address/0xE7022a128Acd4C6cad7aFf6FA874D61f984BcE75) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/StrategyManager.sol) | [`0x858646372CC42E1A627fcE94aa7A7033e7CF075A`](https://etherscan.io/address/0x858646372CC42E1A627fcE94aa7A7033e7CF075A) | [`0xE09d...CC3c`](https://etherscan.io/address/0xE09d4a1717C936ef021e14E72328128268B0CC3c) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/pods/EigenPodManager.sol) | [`0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338`](https://etherscan.io/address/0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338) | [`0xd22d...9ADB`](https://etherscan.io/address/0xd22dd829779ADBf3869fb224F703452f7F95E9dB) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/AVSDirectory.sol) | [`0x135dda560e946695d6f155dacafc6f1f25c1f5af`](https://etherscan.io/address/0x135dda560e946695d6f155dacafc6f1f25c1f5af) | [`0xcD35...862b`](https://etherscan.io/address/0xcD35Cef328b496fA9d70a8d7C34EF3434614862b) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/RewardsCoordinator.sol) | [`0x7750d328b314EfFa365A0402CcfD489B80B0adda`](https://etherscan.io/address/0x7750d328b314EfFa365A0402CcfD489B80B0adda) | [`0x788E...b0cF`](https://etherscan.io/address/0x788E38bCe16Cd96E5588559703469efBA3Afb0cF) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/PermissionController.sol) | [`0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5`](https://etherscan.io/address/0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5) | [`0x36dd...EEd2`](https://etherscan.io/address/0x36dd260AbF606172875E6B5B7A96B435DC74EEd2) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/AllocationManager.sol) | [`0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39`](https://etherscan.io/address/0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39) | [`0xdA2A...97C4`](https://etherscan.io/address/0xdA2A68D318A571dD550F2EcbCb09bf50497e97C4) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ProtocolRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/ProtocolRegistry.sol) | [`0x27a84740FdDed5B7D66d9bb6E5d1DEA6eb0C0129`](https://etherscan.io/address/0x27a84740FdDed5B7D66d9bb6E5d1DEA6eb0C0129) | [`0x611C...9E29`](https://etherscan.io/address/0x611C5b2Fe677a3C09e4D58a1fE40FD6547d5e29E) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyFactory.sol) | [`0x5e4C39Ad7A3E881585e383dB9827EB4811f6F647`](https://etherscan.io/address/0x5e4C39Ad7A3E881585e383dB9827EB4811f6F647) | [`0x5194...c8AF`](https://etherscan.io/address/0x5194D2a6A0900796903503926E9CF775b926c8AF) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBase.sol) | [`0x0ed6703C298d28aE0878d1b28e88cA87F9662fE9`](https://etherscan.io/address/0x0ed6703c298d28ae0878d1b28e88ca87f9662fe9) | [`0xd33A...2780`](https://etherscan.io/address/0xd33AAccc7E1a29Bc8E09Af55F8fa6fF3301e2780) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

The following strategies were originally deployed and whitelisted outside of the `StrategyFactory`:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyBase (cbETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc`](https://etherscan.io/address/0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (stETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x93c4b944D05dfe6df7645A86cd2206016c51564D`](https://etherscan.io/address/0x93c4b944D05dfe6df7645A86cd2206016c51564D) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (rETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2`](https://etherscan.io/address/0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ETHx)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d`](https://etherscan.io/address/0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (ankrETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff`](https://etherscan.io/address/0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (OETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xa4C637e0F704745D182e4D38cAb7E7485321d059`](https://etherscan.io/address/0xa4C637e0F704745D182e4D38cAb7E7485321d059) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (osETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x57ba429517c3473B6d34CA9aCd56c0e735b94c02`](https://etherscan.io/address/0x57ba429517c3473B6d34CA9aCd56c0e735b94c02) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (swETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6`](https://etherscan.io/address/0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (wBETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x7CA911E83dabf90C90dD3De5411a10F1A6112184`](https://etherscan.io/address/0x7CA911E83dabf90C90dD3De5411a10F1A6112184) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (sfrxETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6`](https://etherscan.io/address/0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (lsETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0xAe60d8180437b5C34bB956822ac2710972584473`](https://etherscan.io/address/0xAe60d8180437b5C34bB956822ac2710972584473) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase (mETH)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBaseTVLLimits.sol) | [`0x298aFB19A105D59E74658C4C334Ff360BadE6dd2`](https://etherscan.io/address/0x298aFB19A105D59E74658C4C334Ff360BadE6dd2) | [`0xb8d2...DBa7`](https://etherscan.io/address/0xb8d2cc94A9d2a8Fd7fF499fbE64B0B209212DBa7) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/EigenStrategy.sol) | [`0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7`](https://etherscan.io/address/0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7) | [`0x4a0a...3532`](https://etherscan.io/address/0x4a0aee93BE6C87B227cA0B450E15245631233532) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/pods/EigenPod.sol) | [`0x5a2a4F2F3C18f09179B6703e63D9eDD165909073`](https://etherscan.io/address/0x5a2a4F2F3C18f09179B6703e63D9eDD165909073) | [`0x53cC...e868`](https://etherscan.io/address/0x53cC2D82E08370Fe1e44a96f69CEc7d5b54ae868) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/token/Eigen.sol) | [`0xec53bf9167f50cdeb3ae105f56099aaab9061f83`](https://etherscan.io/address/0xec53bf9167f50cdeb3ae105f56099aaab9061f83) | [`0x2C4A...2E50`](https://etherscan.io/address/0x2C4A81e257381F87F5A5C4bd525116466D972E50) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/token/BackingEigen.sol) | [`0x83E9115d334D248Ce39a6f36144aEaB5b3456e75`](https://etherscan.io/address/0x83E9115d334D248Ce39a6f36144aEaB5b3456e75) | [`0xF2b2...9b17`](https://etherscan.io/address/0xF2b225815F70c9b327DC9db758A36c92A4279b17) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`SignedDistributor`](https://etherscan.io/address/0x035bdAeaB85E47710C27EdA7FD754bA80aD4ad02#code) | - | [`0x035b...ad02`](https://etherscan.io/address/0x035bdAeaB85E47710C27EdA7FD754bA80aD4ad02) | - |

###### AVS

The following contracts are used by AVSs.

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`ReleaseManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/ReleaseManager.sol) | [`0xeDA3CAd031c0cf367cF3f517Ee0DC98F9bA80C8F`](https://etherscan.io/address/0xeDA3CAd031c0cf367cF3f517Ee0DC98F9bA80C8F) | [`0xD0cb...3bd`](https://etherscan.io/address/0xD0cb07Df397b122bB7ebaA453356F21f8Ff813bd) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`TaskMailbox`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/avs/task/TaskMailbox.sol) | [`0x132b466d9d5723531F68797519DfED701aC2C749`](https://etherscan.io/address/0x132b466d9d5723531F68797519DfED701aC2C749) | [`0xA2A9...c55e`](https://etherscan.io/address/0xA2A9D0F957D81A0f8134F68803240f1CAD81c55e) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multichain - Source

The multichain protocol expects AVSs to register on the source chain. AVS's stakes are then transported to supported destination chains. For the mainnet Ethereum network, the destination chains include Base.

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`CrossChainRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/CrossChainRegistry.sol) | [`0x9376A5863F2193cdE13e1aB7c678F22554E2Ea2b`](https://etherscan.io/address/0x9376A5863F2193cdE13e1aB7c678F22554E2Ea2b) | [`0xdc23...642C`](https://etherscan.io/address/0xdc2354FaDd4bf0f9857038381917D0089880642C) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`KeyRegistrar`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/KeyRegistrar.sol) | [`0x54f4bC6bDEbe479173a2bbDc31dD7178408A57A4`](https://etherscan.io/address/0x54f4bC6bDEbe479173a2bbDc31dD7178408A57A4) | [`0x0f93...C219`](https://etherscan.io/address/0x0f939726Ab8514c13546804311149a8CC244C219) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multichain - Destination

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`OperatorTableUpdater`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/OperatorTableUpdater.sol) | [`0x5557E1fE3068A1e823cE5Dcd052c6C352E2617B5`](https://etherscan.io/address/0x5557E1fE3068A1e823cE5Dcd052c6C352E2617B5) | [`0x8c4F...adF7`](https://etherscan.io/address/0x8c4F429e6d884899ebf4602bd2691920B056adF7) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ECDSACertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/ECDSACertificateVerifier.sol) | [`0xd0930ee96D07de4F9d493c259232222e46B6EC25`](https://etherscan.io/address/0xd0930ee96D07de4F9d493c259232222e46B6EC25) | [`0x2d52...0dC`](https://etherscan.io/address/0x2d52c39461795835b1F8F6cD242f4955AaC040dC) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`BN254CertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/BN254CertificateVerifier.sol) | [`0x3F55654b2b2b86bB11bE2f72657f9C33bf88120A`](https://etherscan.io/address/0x3F55654b2b2b86bB11bE2f72657f9C33bf88120A) | [`0x9907...eaD2`](https://etherscan.io/address/0x9907690007b7ECFE1E03F96beca1a957faE3eaD2) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/PauserRegistry.sol) | - | [`0xB876...2806`](https://etherscan.io/address/0xB8765ed72235d279c3Fb53936E4606db0Ef12806) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0xC06F...Aa2d`](https://etherscan.io/address/0xC06Fd4F821eaC1fF1ae8067b36342899b57BAa2d) | |
| [`OZ: TimelockController (BEIGEN)`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0x7381...bc53`](https://etherscan.io/address/0x738130BC8eADe1Bc65A9c056DEa636835896bc53) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x8b95...2444`](https://etherscan.io/address/0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444) | |
| [`Pauser Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x5050389572f2d220ad927CcbeA0D406831012390`](https://etherscan.io/address/0x5050389572f2d220ad927CcbeA0D406831012390) | - | |
| [`Community Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xFEA47018D632A77bA579846c840d5706705Dc598`](https://etherscan.io/address/0xFEA47018D632A77bA579846c840d5706705Dc598) | - | |
| [`Executor Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x369e6F597e22EaB55fFb173C6d9cD234BD699111`](https://etherscan.io/address/0x369e6F597e22EaB55fFb173C6d9cD234BD699111) | - | |
| [`Operations Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xBE1685C81aA44FF9FB319dD389addd9374383e90`](https://etherscan.io/address/0xBE1685C81aA44FF9FB319dD389addd9374383e90) | - | |
| [`Protocol Council Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x461854d84ee845f905e0ecf6c288ddeeb4a9533f`](https://etherscan.io/address/0x461854d84ee845f905e0ecf6c288ddeeb4a9533f) | - | |
| [`Foundation Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xbb00DDa2832850a43840A3A86515E3Fe226865F2`](https://etherscan.io/address/0xbb00DDa2832850a43840A3A86515E3Fe226865F2) | - | |
| [`BEIGEN Executor Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x942eaF324971440384e4cA0ffA39fC3bb369D67d`](https://etherscan.io/address/0x942eaF324971440384e4cA0ffA39fC3bb369D67d) | - | |
| [`Multichain Deployer Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xa3053EF25F1F7d9D55a7655372B8a31D0f40eCA9`](https://etherscan.io/address/0xa3053EF25F1F7d9D55a7655372B8a31D0f40eCA9) | - | |


</details>



<details>
    <summary>Base</summary>


**Note: Base only supports verification of tasks from stake that lives on Mainnet Ethereum. Standard core protocol functionality (restaking, slashing) does not exist on Base.**

###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`ProtocolRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/ProtocolRegistry.sol) | [`0x27a84740FdDed5B7D66d9bb6E5d1DEA6eb0C0129`](https://basescan.org/address/0x27a84740FdDed5B7D66d9bb6E5d1DEA6eb0C0129) | [`0x069f...c8AF`](https://basescan.org/address/0x069f49Be653522fd3F806920f3C21b15F351f9bC) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multichain - Destination

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`OperatorTableUpdater`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/OperatorTableUpdater.sol) | [`0x5557E1fE3068A1e823cE5Dcd052c6C352E2617B5`](https://basescan.org/address/0x5557E1fE3068A1e823cE5Dcd052c6C352E2617B5) | [`0x792F...419D`](https://basescan.org/address/0x792FfeA0D8734695670697f8b9f03DE3F666419D) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ECDSACertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/ECDSACertificateVerifier.sol) | [`0xd0930ee96D07de4F9d493c259232222e46B6EC25`](https://basescan.org/address/0xd0930ee96D07de4F9d493c259232222e46B6EC25) | [`0xf2cC...025F`](https://basescan.org/address/0xf2cC35A48561c2344577489a0873339F5B7E025F) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`BN254CertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/BN254CertificateVerifier.sol) | [`0x3F55654b2b2b86bB11bE2f72657f9C33bf88120A`](https://basescan.org/address/0x3F55654b2b2b86bB11bE2f72657f9C33bf88120A) | [`0xc6c7...5D7d`](https://basescan.org/address/0xc6c7a63518d0150D57C6b00CDC7F13eB4b275D7d) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### AVS

The following contracts are used by AVSs.

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`TaskMailbox`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/avs/task/TaskMailbox.sol) | [`0x132b466d9d5723531F68797519DfED701aC2C749`](https://basescan.org/address/0x132b466d9d5723531F68797519DfED701aC2C749) | [`0x6eD2...c1bd`](https://basescan.org/address/0x6eD28c7dfEa8F70a7edec02A1E295DA6307Fc1bd) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/PauserRegistry.sol) | - | [`0xb175...6ea`](https://basescan.org/address/0xb1754226917e866c1701f1d9F9E135D88f2e86ea) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0xE48D...ABd`](https://basescan.org/address/0xE48D7CaeC1790b293667e4bB2dE1E00536F2bABd) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xDb00...9CE`](https://basescan.org/address/0xDb0052E272D1D126617c36A9D6d65C801d1549CE) | |
| [`Pauser Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x1a051eF1524cbaEa57Ca04319ef93fE78903D5E6`](https://basescan.org/address/0x1a051eF1524cbaEa57Ca04319ef93fE78903D5E6) | - | |
| [`Operations Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x8eD55c7640497Db15aC32c698c1a06E2E604d865`](https://basescan.org/address/0x8eD55c7640497Db15aC32c698c1a06E2E604d865) | - | |
| [`Protocol Council Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x841B988aaEafce13b6456ff34015FBc42Aedb7e6`](https://basescan.org/address/0x841B988aaEafce13b6456ff34015FBc42Aedb7e6) | - | |
| [`Community Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xC107547924C7D1d3E2d10eA8DF534BBfC5F373e6`](https://basescan.org/address/0xC107547924C7D1d3E2d10eA8DF534BBfC5F373e6) | - | |
| [`Executor Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x5d808608Ad70873ca4dE50F83416264dc15264C2`](https://basescan.org/address/0x5d808608Ad70873ca4dE50F83416264dc15264C2) | - | |
| [`Multichain Deployer Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xa3053EF25F1F7d9D55a7655372B8a31D0f40eCA9`](https://basescan.org/address/0xa3053EF25F1F7d9D55a7655372B8a31D0f40eCA9) | - | |

</details>



<details>
    <summary>Testnet Sepolia</summary>


You can view the deployed contract addresses below, or check out the code itself on the [`testnet-sepolia`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/testnet-sepolia) branch.

###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/DelegationManager.sol) | [`0xD4A7E1Bd8015057293f0D0A557088c286942e84b`](https://sepolia.etherscan.io/address/0xD4A7E1Bd8015057293f0D0A557088c286942e84b) | [`0xa6Fe...C0f5`](https://sepolia.etherscan.io/address/0xa6FeC50d15FCaE6d3Ac8755e350C747F7733C0f5) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/StrategyManager.sol) | [`0x2E3D6c0744b10eb0A4e6F679F71554a39Ec47a5D`](https://sepolia.etherscan.io/address/0x2E3D6c0744b10eb0A4e6F679F71554a39Ec47a5D) | [`0x7D59...E6F1`](https://sepolia.etherscan.io/address/0x7D59F252Bd32733f8850c50bF6BB2e46BF37E6F1) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/pods/EigenPodManager.sol) | [`0x56BfEb94879F4543E756d26103976c567256034a`](https://sepolia.etherscan.io/address/0x56BfEb94879F4543E756d26103976c567256034a) | [`0x151e...75D4`](https://sepolia.etherscan.io/address/0x151eCe1662E530f4889F016a63FC58F4b72175D4) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) | All EigenPod functionality is paused on Sepolia |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/AVSDirectory.sol) | [`0xa789c91ECDdae96865913130B786140Ee17aF545`](https://sepolia.etherscan.io/address/0xa789c91ECDdae96865913130B786140Ee17aF545) | [`0x3216...6D69`](https://sepolia.etherscan.io/address/0x321604fDA757e8728d7b338C284613E2A0136D69) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/RewardsCoordinator.sol) | [`0x5ae8152fb88c26ff9ca5C014c94fca3c68029349`](https://sepolia.etherscan.io/address/0x5ae8152fb88c26ff9ca5C014c94fca3c68029349) | [`0x5Ac1...FF61`](https://sepolia.etherscan.io/address/0x5Ac1705931E999aa267167DaE57b41fA690bFF61) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/AllocationManager.sol) | [`0x42583067658071247ec8CE0A516A58f682002d07`](https://sepolia.etherscan.io/address/0x42583067658071247ec8CE0A516A58f682002d07) | [`0x87cf...4f31`](https://sepolia.etherscan.io/address/0x87cffb49a29d4dC1C0842c70a75eDBa9B62e4f31) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/PermissionController.sol) | [`0x44632dfBdCb6D3E21EF613B0ca8A6A0c618F5a37`](https://sepolia.etherscan.io/address/0x44632dfBdCb6D3E21EF613B0ca8A6A0c618F5a37) | [`0xEBDa...B720`](https://sepolia.etherscan.io/address/0xEBDa1FCe8527C0c9ac94a60D7ECA37640415B720) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`KeyRegistrar`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/KeyRegistrar.sol) | [`0xA4dB30D08d8bbcA00D40600bee9F029984dB162a`](https://sepolia.etherscan.io/address/0xA4dB30D08d8bbcA00D40600bee9F029984dB162a) | [`0x1AFc...68e6`](https://sepolia.etherscan.io/address/0x1AFc1E6c6bA44B21Ce621B5B41fAa085682c68e6) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ProtocolRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/ProtocolRegistry.sol) | [`0xfF2dF19481cC4d7fB30698e78c3aeEc3b1265e1e`](https://sepolia.etherscan.io/address/0xfF2dF19481cC4d7fB30698e78c3aeEc3b1265e1e) | [`0x832F...E5e9`](https://sepolia.etherscan.io/address/0x832F4761ab4584FB967c08921A9C6a0a2D09d4e9) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyFactory.sol) | [`0x066cF95c1bf0927124DFB8B02B401bc23A79730D`](https://sepolia.etherscan.io/address/0x066cF95c1bf0927124DFB8B02B401bc23A79730D) | [`0x3f98...b5Ac`](https://sepolia.etherscan.io/address/0x3f98d8B9CFa2102aD340C19648E8BB3C06fbc5Ac) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBase.sol) | [`0x427e627Bc7E83cac0f84337d3Ad94230C32697D3`](https://sepolia.etherscan.io/address/0x427e627Bc7E83cac0f84337d3Ad94230C32697D3) | [`0x8f31...5cEf`](https://sepolia.etherscan.io/address/0x8f31bFC631B51A39f027A7C9750F7B5cCe9E5cEf) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| [`StETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBase.sol) | [`0x8b29d91e67b013e855EaFe0ad704aC4Ab086a574`](https://sepolia.etherscan.io/address/0x8b29d91e67b013e855EaFe0ad704aC4Ab086a574) | [`0x8f31...5cEf`](https://sepolia.etherscan.io/address/0x8f31bFC631B51A39f027A7C9750F7B5cCe9E5cEf) | Strategy Factory deployed |
| [`WETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBase.sol) | [`0x424246eF71b01ee33aA33aC590fd9a0855F5eFbc`](https://sepolia.etherscan.io/address/0x424246eF71b01ee33aA33aC590fd9a0855F5eFbc) | [`0x8f31...5cEf`](https://sepolia.etherscan.io/address/0x8f31bFC631B51A39f027A7C9750F7B5cCe9E5cEf) | Strategy Factory deployed |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/EigenStrategy.sol) | [`0x8E93249a6C37a32024756aaBd813E6139b17D1d5`](https://sepolia.etherscan.io/address/0x8E93249a6C37a32024756aaBd813E6139b17D1d5) | [`0xE8E4...7afC`](https://sepolia.etherscan.io/address/0xE8E469151e8d561Be94a4838B582B7da25487afC) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

**NOTE: Due to the permissioned validator set on Sepolia, all EigenPod functionality is *PAUSED*.**

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/pods/EigenPod.sol) | [`0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4`](https://sepolia.etherscan.io/address/0x0e19E56E41D42137d00dD4f51EC2F613E50cAcf4) | [`0xb1B6...B533`](https://sepolia.etherscan.io/address/0xb1B6dAD80AD719BC5A700f8F3C9Da46224D3B533) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/token/Eigen.sol) | [`0x0011FA2c512063C495f77296Af8d195F33A8Dd38`](https://sepolia.etherscan.io/address/0x0011FA2c512063C495f77296Af8d195F33A8Dd38) | [`0x7ec6...BD6F`](https://sepolia.etherscan.io/address/0x7ec6a02235Bf8d8a1fDB894AD2E1573192bfBD6F) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/token/BackingEigen.sol) | [`0xc5B857A92245f64e9D90cCc5b096Db82eB77eB5c`](https://sepolia.etherscan.io/address/0xc5B857A92245f64e9D90cCc5b096Db82eB77eB5c) | [`0x1298...3173`](https://sepolia.etherscan.io/address/0x12988B679AA497C30A8D1850eCC4Dc7700383173) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### AVS

The following contracts are used by AVSs.

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`ReleaseManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/ReleaseManager.sol) | [`0x59c8D715DCa616e032B744a753C017c9f3E16bf4`](https://sepolia.etherscan.io/address/0x59c8D715DCa616e032B744a753C017c9f3E16bf4) | [`0x67cE...2ac6`](https://sepolia.etherscan.io/address/0x67cEc1bE4De9D4a96Bd6dB28F9ceD6A1bD562ac6) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`TaskMailbox`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/avs/task/TaskMailbox.sol) | [`0xB99CC53e8db7018f557606C2a5B066527bF96b26`](https://sepolia.etherscan.io/address/0xB99CC53e8db7018f557606C2a5B066527bF96b26) | [`0x49A8...5AdB`](https://sepolia.etherscan.io/address/0x49A800bfFa5B561daD94Adf4Ea9bfe35bBEB5AdB) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multichain - Source

The multichain protocol expects AVSs to register on the source chain. AVS's stakes are then transported to supported destination chains. For the sepolia network, the destination chains are sepolia and base-sepolia. 

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`CrossChainRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/CrossChainRegistry.sol) | [`0x287381B1570d9048c4B4C7EC94d21dDb8Aa1352a`](https://sepolia.etherscan.io/address/0x287381B1570d9048c4B4C7EC94d21dDb8Aa1352a) | [`0x4a93...3084`](https://sepolia.etherscan.io/address/0x4a93d276ea78B48be9D8BB864BACDD5D5A713084) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multichain - Destination

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`OperatorTableUpdater`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/OperatorTableUpdater.sol) | [`0xB02A15c6Bd0882b35e9936A9579f35FB26E11476`](https://sepolia.etherscan.io/address/0xB02A15c6Bd0882b35e9936A9579f35FB26E11476) | [`0x0F26...0F5C`](https://sepolia.etherscan.io/address/0x0F264E714a3c03309F4041Db26229EF4E9B00F5C) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ECDSACertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/ECDSACertificateVerifier.sol) | [`0xb3Cd1A457dEa9A9A6F6406c6419B1c326670A96F`](https://sepolia.etherscan.io/address/0xb3Cd1A457dEa9A9A6F6406c6419B1c326670A96F) | [`0xFc86...7B26`](https://sepolia.etherscan.io/address/0xFc868562d93c4DC192419e970C220279cdAb7B26) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`BN254CertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/BN254CertificateVerifier.sol) | [`0xff58A373c18268F483C1F5cA03Cf885c0C43373a`](https://sepolia.etherscan.io/address/0xff58A373c18268F483C1F5cA03Cf885c0C43373a) | [`0x7a2b...8f5d`](https://sepolia.etherscan.io/address/0x7a2b8c559a8C8c71A9d364Ad250Fce5A24b18f5d) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/PauserRegistry.sol) | - | [`0x63AA...20f3`](https://sepolia.etherscan.io/address/0x63AAe451780090f50Ad323aAEF155F63a29D20f3) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0x1BEF...1b5B`](https://sepolia.etherscan.io/address/0x1BEF05C7303d44e0E2FCD2A19d993eDEd4c51b5B) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0x56E8...6Fa1`](https://sepolia.etherscan.io/address/0x56E88cb4f0136fC27D95499dE4BE2acf47946Fa1) | |
| [`Pauser Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x0B415f75980D863872C3eb8caa76E6eC8Bc81536`](https://sepolia.etherscan.io/address/0x0B415f75980D863872C3eb8caa76E6eC8Bc81536) | [`0x4167...461a`](https://sepolia.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Community Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x6f8459810197cc9fE123BBeB918451757a4fBAc6`](https://sepolia.etherscan.io/address/0x6f8459810197cc9fE123BBeB918451757a4fBAc6) | [`0x4167...461a`](https://sepolia.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Executor Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0x4FDA8998EC3b7d4b4A612d45FeB8fB36734470f2`](https://sepolia.etherscan.io/address/0x4FDA8998EC3b7d4b4A612d45FeB8fB36734470f2) | [`0x4167...461a`](https://sepolia.etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |
| [`Operations Multisig`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/GnosisSafe.sol) | [`0xb094Ba769b4976Dc37fC689A76675f31bc4923b0`](https://sepolia.etherscan.io/address/0xb094Ba769b4976Dc37fC689A76675f31bc4923b0) | [`0x4167...461a`](https://sepolia.etherscan.io/address/0x41675C099F32341bf84BFc5382aF534df5C7461a) | Proxy: [`Gnosis@1.3.0`](https://github.com/safe-global/safe-contracts/blob/v1.3.0/contracts/proxies/GnosisSafeProxy.sol) |

</details>



<details>
    <summary>Testnet Hoodi</summary>


You can view the deployed contract addresses below, or check out the code itself on the [`testnet-hoodi`](https://github.com/Layr-Labs/eigenlayer-contracts/tree/testnet-hoodi) branch.

###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`DelegationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/DelegationManager.sol) | [`0x867837a9722C512e0862d8c2E15b8bE220E8b87d`](https://hoodi.etherscan.io/address/0x867837a9722C512e0862d8c2E15b8bE220E8b87d) | [`0x6FE5...1407`](https://hoodi.etherscan.io/address/0x6FE5eBE3ED0F4b2419263868efb5CC640e541407) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/StrategyManager.sol) | [`0xeE45e76ddbEDdA2918b8C7E3035cd37Eab3b5D41`](https://hoodi.etherscan.io/address/0xeE45e76ddbEDdA2918b8C7E3035cd37Eab3b5D41) | [`0x1B76...f62`](https://hoodi.etherscan.io/address/0x1B76a41b1128a5Cd92dAE7C4c0DEC0bfe12B0f62) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`EigenPodManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/pods/EigenPodManager.sol) | [`0xcd1442415Fc5C29Aa848A49d2e232720BE07976c`](https://hoodi.etherscan.io/address/0xcd1442415Fc5C29Aa848A49d2e232720BE07976c) | [`0x0F26...0F5C`](https://hoodi.etherscan.io/address/0x0F264E714a3c03309F4041Db26229EF4E9B00F5C) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AVSDirectory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/AVSDirectory.sol) | [`0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926`](https://hoodi.etherscan.io/address/0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926) | [`0xE368...13Eb`](https://hoodi.etherscan.io/address/0xE368218aDDc4d2dDC5b0AFd70f5878DD8c1713Eb) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`RewardsCoordinator`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/RewardsCoordinator.sol) | [`0x29e8572678e0c272350aa0b4B8f304E47EBcd5e7`](https://hoodi.etherscan.io/address/0x29e8572678e0c272350aa0b4B8f304E47EBcd5e7) | [`0x35a8...0525`](https://hoodi.etherscan.io/address/0x35a8198E98cB76f1DD81D977fBd2a8F2a62C0525) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`AllocationManager`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/AllocationManager.sol) | [`0x95a7431400F362F3647a69535C5666cA0133CAA0`](https://hoodi.etherscan.io/address/0x95a7431400F362F3647a69535C5666cA0133CAA0) | [`0xe1b7...BfB9`](https://hoodi.etherscan.io/address/0xe1b7A1249C71b538CC183B0080FfC3eFD02BFfB9) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`PermissionController`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/PermissionController.sol) | [`0xdcCF401fD121d8C542E96BC1d0078884422aFAD2`](https://hoodi.etherscan.io/address/0xdcCF401fD121d8C542E96BC1d0078884422aFAD2) | [`0x151e...75D4`](https://hoodi.etherscan.io/address/0x151eCe1662E530f4889F016a63FC58F4b72175D4) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ProtocolRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/ProtocolRegistry.sol) | [`0xff2df19481cc4d7fb30698e78c3aeec3b1265e1e`](https://hoodi.etherscan.io/address/0xff2df19481cc4d7fb30698e78c3aeec3b1265e1e) | [`0xEb3C...7841`](https://hoodi.etherscan.io/address/0xEb3CDe433CEd38fEe69073658061A1D6A7DE7841) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |


###### Strategies

Anyone can deploy and whitelist strategies for standard ERC20s by using the `StrategyFactory` deployed to the address below (see [docs](./docs/core/StrategyManager.md#strategyfactorydeploynewstrategy)). Strategies deployed from the `StrategyFactory` are deployed using the beacon proxy pattern:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`StrategyFactory`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyFactory.sol) | [`0xfB7d94501E4d4ACC264833Ef4ede70a11517422B`](https://hoodi.etherscan.io/address/0xfB7d94501E4d4ACC264833Ef4ede70a11517422B) | [`0x0d82...A3f3`](https://hoodi.etherscan.io/address/0x0d824D128eb5bBF5a0dEDD7D58332a00de57A3f3) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`StrategyBase`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBase.sol) | [`0x6d28cEC1659BC3a9BC814c3EFc1412878B406579`](https://hoodi.etherscan.io/address/0x6d28cEC1659BC3a9BC814c3EFc1412878B406579) | [`0xa621...063A`](https://hoodi.etherscan.io/address/0xa621Ebb0Df923e87e9C527C9fC55b9C6F6E0063A) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.0/contracts/proxy/beacon/BeaconProxy.sol) <br />- Strategies: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |
| [`StETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBase.sol) | [`0xf8a1a66130d614c7360e868576d5e59203475fe0`](https://hoodi.etherscan.io/address/0xf8a1a66130d614c7360e868576d5e59203475fe0) | [`0xa621...063A`](https://hoodi.etherscan.io/address/0xa621Ebb0Df923e87e9C527C9fC55b9C6F6E0063A) | Strategy Factory deployed |
| [`WETH Strategy`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/StrategyBase.sol) | [`0x24579aD4fe83aC53546E5c2D3dF5F85D6383420d`](https://hoodi.etherscan.io/address/0x24579aD4fe83aC53546E5c2D3dF5F85D6383420d) | [`0xa621...063A`](https://hoodi.etherscan.io/address/0xa621Ebb0Df923e87e9C527C9fC55b9C6F6E0063A) | Strategy Factory deployed |

###### Strategies - Special

The following strategies differ significantly from the other strategies deployed/used above:

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`EigenStrategy (EIGEN)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/strategies/EigenStrategy.sol) | [`0xB27b10291DBFE6576d17afF3e251c954Ae14f1D3`](https://hoodi.etherscan.io/address/0xB27b10291DBFE6576d17afF3e251c954Ae14f1D3) | [`0xd318...caE8`](https://hoodi.etherscan.io/address/0xd318ECAcaf8998FB4508FA07159B758253FbcaE8) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `Beacon Chain ETH` | `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0` | - | - Used for Beacon Chain ETH shares <br />- Not a real contract! |

###### EigenPods

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`EigenPod (beacon)`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/pods/EigenPod.sol) | [`0x5e1577f8efB21b229cD5Eb4C5Aa3d6C4b228f650`](https://hoodi.etherscan.io/address/0x5e1577f8efB21b229cD5Eb4C5Aa3d6C4b228f650) | [`0x49A8...5AdB`](https://hoodi.etherscan.io/address/0x49A800bfFa5B561daD94Adf4Ea9bfe35bBEB5AdB) | - Beacon: [`BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/BeaconProxy.sol) <br />- Pods: [`UpgradeableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.4.1/contracts/proxy/beacon/UpgradeableBeacon.sol) |

###### EIGEN/bEIGEN

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/token/Eigen.sol) | [`0x8ae2520954db7D80D66835cB71E692835bbA45bf`](https://hoodi.etherscan.io/address/0x8ae2520954db7D80D66835cB71E692835bbA45bf) | [`0x68F3...2FE9`](https://hoodi.etherscan.io/address/0x68F3E0B02b15F96B50aC73d7680F020B646D2FE9) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`Backing Eigen`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/token/BackingEigen.sol) | [`0x6e60888132Cc7e637488379B4B40c42b3751f63a`](https://hoodi.etherscan.io/address/0x6e60888132Cc7e637488379B4B40c42b3751f63a) | [`0x43e4...1C4`](https://hoodi.etherscan.io/address/0x43e4940aCeb1C1F5a57e307EEB212007F0f6a1C4) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/PauserRegistry.sol) | - | [`0x64D7...c13D`](https://hoodi.etherscan.io/address/0x64D78399B0fa32EA72959f33edCF313159F3c13D) | |
| [`OZ: TimelockController`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.7/contracts/governance/TimelockController.sol) | - | [`0xE332...2d15`](https://hoodi.etherscan.io/address/0xE3328cb5068924119d6170496c4AB2dD12c12d15) | |
| [`OZ: Proxy Admin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/ProxyAdmin.sol) | - | [`0xE7f4...2257`](https://hoodi.etherscan.io/address/0xE7f4E30D2619273468afe9EC0Acf805E55532257) | |

</details>



<details>
    <summary>Testnet Base Sepolia</summary>


**Note: Testnet Base Sepolia only supports verification of tasks from stake that lives on Sepolia. Standard core protocol functionality (restaking, slashing) does not exist on Base Sepolia.**

###### Core

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`ProtocolRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/core/ProtocolRegistry.sol) | [`0xfF2dF19481cC4d7fB30698e78c3aeEc3b1265e1e`](https://sepolia.basescan.org/address/0xfF2dF19481cC4d7fB30698e78c3aeEc3b1265e1e) | [`0xcC30...7940`](https://sepolia.basescan.org/address/0xcC305562B01bec562D13A40ef8781e313AFE7940) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multichain - Destination

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`OperatorTableUpdater`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/OperatorTableUpdater.sol) | [`0xB02A15c6Bd0882b35e9936A9579f35FB26E11476`](https://sepolia.basescan.org/address/0xB02A15c6Bd0882b35e9936A9579f35FB26E11476) | [`0x6514...2D6F`](https://sepolia.basescan.org/address/0x65147e9916152C1EbdBaD8A6f3e145b4BDeE2D6F) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`ECDSACertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/ECDSACertificateVerifier.sol) | [`0xb3Cd1A457dEa9A9A6F6406c6419B1c326670A96F`](https://sepolia.basescan.org/address/0xb3Cd1A457dEa9A9A6F6406c6419B1c326670A96F) | [`0xEE41...ca1A`](https://sepolia.basescan.org/address/0xEE41826B7D5B89e7F5eED6a831b4eFD69FC9ca1A) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| [`BN254CertificateVerifier`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/multichain/BN254CertificateVerifier.sol) | [`0xff58A373c18268F483C1F5cA03Cf885c0C43373a`](https://sepolia.basescan.org/address/0xff58A373c18268F483C1F5cA03Cf885c0C43373a) | [`0x59B1...f525`](https://sepolia.basescan.org/address/0x59B11b191B572888703E150E45F5015e0fFcf525) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### AVS

The following contracts are used by AVSs.

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`TaskMailbox`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/avs/task/TaskMailbox.sol) | [`0xB99CC53e8db7018f557606C2a5B066527bF96b26`](https://sepolia.basescan.org/address/0xB99CC53e8db7018f557606C2a5B066527bF96b26) | [`0x46CF...16db`](https://sepolia.basescan.org/address/0x46CFA3C2eaDe97D53739120b87A63F739B9616db) | Proxy: [`TUP@4.9.0`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

###### Multisigs

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- | 
| [`PauserRegistry`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v1.9.0/src/contracts/permissions/PauserRegistry.sol) | - | [`0x6ffE...672B`](https://sepolia.basescan.org/address/0x6ffE77C321a773e2A27B0B0a31C5e1bBda83672B) | |
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