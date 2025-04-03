// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import "src/test/ConfigParser.t.sol";

abstract contract ConfigGetters {
    using ConfigParser for *;

    Config public config;

    /// -----------------------------------------------------------------------
    /// Governance
    /// -----------------------------------------------------------------------

    function communityMultisig() public view virtual returns (address) {
        return config.governance.communityMultisig;
    }

    function executorMultisig() public view virtual returns (address) {
        return config.governance.executorMultisig;
    }

    function operationsMultisig() public view virtual returns (address) {
        return config.governance.operationsMultisig;
    }

    function pauserMultisig() public view virtual returns (address) {
        return config.governance.pauserMultisig;
    }

    function pauserRegistry() public view virtual returns (PauserRegistry) {
        return config.governance.pauserRegistry;
    }

    function proxyAdmin() public view virtual returns (ProxyAdmin) {
        return config.governance.proxyAdmin;
    }

    function timelock() public view virtual returns (TimelockController) {
        return config.governance.timelock;
    }

    /// -----------------------------------------------------------------------
    /// Tokens
    /// -----------------------------------------------------------------------

    function bEIGEN() public view virtual returns (IERC20) {
        return config.tokens.bEIGEN;
    }

    function EIGEN() public view virtual returns (IERC20) {
        return config.tokens.EIGEN;
    }

    /// -----------------------------------------------------------------------
    /// Core
    /// -----------------------------------------------------------------------

    function allocationManager() public view virtual returns (AllocationManager) {
        return config.core.allocationManager;
    }

    function allocationManagerImpl() public view virtual returns (address) {
        return address(config.core.allocationManager).impl();
    }

    function avsDirectory() public view virtual returns (AVSDirectory) {
        return config.core.avsDirectory;
    }

    function avsDirectoryImpl() public view virtual returns (address) {
        return address(config.core.avsDirectory).impl();
    }

    function delegationManager() public view virtual returns (DelegationManager) {
        return config.core.delegationManager;
    }

    function delegationManagerImpl() public view virtual returns (address) {
        return address(config.core.delegationManager).impl();
    }

    function permissionController() public view virtual returns (PermissionController) {
        return config.core.permissionController;
    }

    function permissionControllerImpl() public view virtual returns (address) {
        return address(config.core.permissionController).impl();
    }

    function rewardsCoordinator() public view virtual returns (RewardsCoordinator) {
        return config.core.rewardsCoordinator;
    }

    function rewardsCoordinatorImpl() public view virtual returns (address) {
        return address(config.core.rewardsCoordinator).impl();
    }

    function strategyManager() public view virtual returns (StrategyManager) {
        return config.core.strategyManager;
    }

    function strategyManagerImpl() public view virtual returns (address) {
        return address(config.core.strategyManager).impl();
    }

    /// -----------------------------------------------------------------------
    /// Pods
    /// -----------------------------------------------------------------------

    function eigenPodBeacon() public view virtual returns (UpgradeableBeacon) {
        return config.pods.eigenPodBeacon;
    }

    function eigenPodManager() public view virtual returns (EigenPodManager) {
        return config.pods.eigenPodManager;
    }

    function eigenPodManagerImpl() public view virtual returns (address) {
        return address(config.pods.eigenPodManager).impl();
    }

    function eigenStrategy() public view virtual returns (EigenStrategy) {
        return config.pods.eigenStrategy;
    }

    function eigenStrategyImpl() public view virtual returns (address) {
        return address(config.pods.eigenStrategy).impl();
    }

    /// -----------------------------------------------------------------------
    /// Strategies
    /// -----------------------------------------------------------------------

    function strategyFactory() public view virtual returns (StrategyFactory) {
        return config.strategies.strategyFactory;
    }

    function strategyFactoryBeacon() public view virtual returns (UpgradeableBeacon) {
        return config.strategies.strategyFactoryBeacon;
    }

    function strategyAddresses() public view virtual returns (IStrategy[] memory) {
        return config.strategies.strategyAddresses;
    }

    function strategyAddresses(uint index) public view virtual returns (IStrategy) {
        return config.strategies.strategyAddresses[index];
    }

    function totalStrategies() public view virtual returns (uint) {
        return config.strategies.strategyAddresses.length;
    }
}

contract ConfigGettersTest is ConfigGetters, Test {
    function setUp() public {
        vm.createSelectFork(vm.rpcUrl("mainnet"), 22_181_590);
        config = ConfigParser.parse("./script/configs/mainnet/mainnet-addresses.toml");
    }

    function test_parseTOML() public {
        // Governance addresses
        assertEq(communityMultisig(), 0xFEA47018D632A77bA579846c840d5706705Dc598);
        assertEq(executorMultisig(), 0x369e6F597e22EaB55fFb173C6d9cD234BD699111);
        assertEq(operationsMultisig(), 0xBE1685C81aA44FF9FB319dD389addd9374383e90);
        assertEq(pauserMultisig(), 0x5050389572f2d220ad927CcbeA0D406831012390);
        assertEq(address(pauserRegistry()), 0x0c431C66F4dE941d089625E5B423D00707977060);
        assertEq(address(proxyAdmin()), 0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444);
        assertEq(address(timelock()), 0xA6Db1A8C5a981d1536266D2a393c5F8dDb210EAF);

        // Token addresses
        assertEq(address(bEIGEN()), 0x83E9115d334D248Ce39a6f36144aEaB5b3456e75);
        assertEq(address(EIGEN()), 0xec53bF9167f50cDEB3Ae105f56099aaaB9061F83);

        // Core addresses
        assertEq(address(allocationManager()), 0x0000000000000000000000000000000000000000);
        assertEq(address(avsDirectory()), 0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF);
        assertEq(address(delegationManager()), 0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A);
        assertEq(address(permissionController()), 0x0000000000000000000000000000000000000000);
        assertEq(address(rewardsCoordinator()), 0x7750d328b314EfFa365A0402CcfD489B80B0adda);
        assertEq(address(strategyManager()), 0x858646372CC42E1A627fcE94aa7A7033e7CF075A);

        // Pod addresses
        assertEq(address(eigenPodBeacon()), 0x5a2a4F2F3C18f09179B6703e63D9eDD165909073);
        assertEq(address(eigenPodManager()), 0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338);
        assertEq(address(eigenStrategy()), 0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7);

        // Strategy addresses
        assertEq(address(strategyFactory()), 0x5e4C39Ad7A3E881585e383dB9827EB4811f6F647);
        assertEq(address(strategyFactoryBeacon()), 0x0ed6703C298d28aE0878d1b28e88cA87F9662fE9);
        assertEq(totalStrategies(), 12);
        assertEq(address(strategyAddresses(0)), 0x93c4b944D05dfe6df7645A86cd2206016c51564D); // stETH
        assertEq(address(strategyAddresses(1)), 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2); // rETH
        assertEq(address(strategyAddresses(2)), 0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc); // cbETH
        assertEq(address(strategyAddresses(3)), 0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d); // ETHx
        assertEq(address(strategyAddresses(4)), 0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff); // ankrETH
        assertEq(address(strategyAddresses(5)), 0xa4C637e0F704745D182e4D38cAb7E7485321d059); // oETH
        assertEq(address(strategyAddresses(6)), 0x57ba429517c3473B6d34CA9aCd56c0e735b94c02); // osETH
        assertEq(address(strategyAddresses(7)), 0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6); // swETH
        assertEq(address(strategyAddresses(8)), 0x7CA911E83dabf90C90dD3De5411a10F1A6112184); // wBETH
        assertEq(address(strategyAddresses(9)), 0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6); // sfrxETH
        assertEq(address(strategyAddresses(10)), 0xAe60d8180437b5C34bB956822ac2710972584473); // lsETH
        assertEq(address(strategyAddresses(11)), 0x298aFB19A105D59E74658C4C334Ff360BadE6dd2); // mETH
    }
}
