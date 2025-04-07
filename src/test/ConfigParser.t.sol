// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "script/releases/Env.sol";
import "src/test/mocks/EmptyContract.sol";

// WARNING: We expect the layout to follow the order of the toml config file. Please see:
// - script/configs/mainnet/mainnet-addresses.toml
// - https://book.getfoundry.sh/cheatcodes/parse-toml?highlight=toml#decoding-toml-tables-into-solidity-structs

struct Tokens {
    IERC20 bEIGEN;
    IERC20 EIGEN;
}

struct Core {
    AllocationManager allocationManager;
    AVSDirectory avsDirectory;
    DelegationManager delegationManager;
    PermissionController permissionController;
    RewardsCoordinator rewardsCoordinator;
    StrategyManager strategyManager;
}

struct Governance {
    address communityMultisig;
    address executorMultisig;
    address operationsMultisig;
    address pauserMultisig;
    PauserRegistry pauserRegistry;
    ProxyAdmin proxyAdmin;
    TimelockController timelock;
}

struct Pods {
    UpgradeableBeacon eigenPodBeacon;
    EigenPodManager eigenPodManager;
    EigenStrategy eigenStrategy;
}

struct Strategies {
    IStrategy[] strategyAddresses;
    StrategyFactory strategyFactory;
    UpgradeableBeacon strategyFactoryBeacon;
}

struct Config {
    Core core;
    Governance governance;
    Pods pods;
    Strategies strategies;
    Tokens tokens;
}

library ConfigParser {
    using ConfigParser for *;
    using stdToml for string;

    Vm constant vm = Vm(address(uint160(uint(keccak256("hevm cheat code")))));

    /// -----------------------------------------------------------------------
    /// Config Parsing
    /// -----------------------------------------------------------------------

    /// @dev Returns the config from a given TOML file.
    function parse(string memory path) internal returns (Config memory) {
        return abi.decode(vm.readFile(path).parseRaw("."), (Config));
    }

    /// @dev Returns the config from a given Zeus enviorment.
    function parseZeus() internal returns (Config memory c) {
        // Governance addresses
        c.governance.communityMultisig = Env.protocolCouncilMultisig();
        c.governance.executorMultisig = Env.executorMultisig();
        c.governance.operationsMultisig = Env.opsMultisig();
        c.governance.pauserMultisig = Env.pauserMultisig();
        c.governance.pauserRegistry = Env.pauserRegistry(Env.impl);
        c.governance.proxyAdmin = ProxyAdmin(Env.proxyAdmin());
        c.governance.timelock = Env.timelockController();
        // Token addresses
        c.tokens.bEIGEN = Env.beigen(Env.proxy);
        c.tokens.EIGEN = Env.eigen(Env.proxy);
        // Core addresses
        c.core.allocationManager = Env.allocationManager(Env.proxy);
        c.core.avsDirectory = Env.avsDirectory(Env.proxy);
        c.core.delegationManager = Env.delegationManager(Env.proxy);
        c.core.permissionController = Env.permissionController(Env.proxy);
        c.core.rewardsCoordinator = Env.rewardsCoordinator(Env.proxy);
        c.core.strategyManager = Env.strategyManager(Env.proxy);
        // Pod addresses
        c.pods.eigenPodBeacon = Env.eigenPod(Env.beacon);
        c.pods.eigenPodManager = Env.eigenPodManager(Env.proxy);
        c.pods.eigenStrategy = Env.eigenStrategy(Env.proxy);
        // Strategy addresses
        c.strategies.strategyFactory = Env.strategyFactory(Env.proxy);
        c.strategies.strategyFactoryBeacon = Env.strategyBase(Env.beacon);
        // Get all strategy instances
        c.strategies.strategyAddresses = new IStrategy[](Env.strategyBaseTVLLimits_Count(Env.instance));
        for (uint i; i < c.strategies.strategyAddresses.length; ++i) {
            c.strategies.strategyAddresses[i] = Env.strategyBaseTVLLimits(Env.instance, i);
        }
    }

    /// -----------------------------------------------------------------------
    /// Proxy Storage
    /// -----------------------------------------------------------------------

    bytes32 constant ERC1967_IMPL_SLOT = bytes32(uint(keccak256("eip1967.proxy.implementation")) - 1);
    bytes32 constant ERC1967_BEACON_SLOT = bytes32(uint(keccak256("eip1967.proxy.beacon")) - 1);
    bytes32 constant ERC1967_ADMIN_SLOT = bytes32(uint(keccak256("eip1967.proxy.admin")) - 1);

    /// @dev Returns the implementation address of a ERC1967 proxy.
    function impl(address proxy) internal view returns (address) {
        return address(uint160(uint(vm.load(proxy, ERC1967_IMPL_SLOT))));
    }

    /// @dev Returns the beacon address of a ERC1967 proxy.
    function beacon(address proxy) internal view returns (address) {
        return address(uint160(uint(vm.load(proxy, ERC1967_BEACON_SLOT))));
    }

    /// @dev Returns the admin address of a ERC1967 proxy.
    function admin(address proxy) internal view returns (address) {
        return address(uint160(uint(vm.load(proxy, ERC1967_ADMIN_SLOT))));
    }

    function setImpl(address proxy, address impl) internal returns (address) {
        vm.store(proxy, ERC1967_IMPL_SLOT, bytes32(uint(uint160(impl))));
        return proxy;
    }

    function setBeacon(address proxy, address beacon) internal returns (address) {
        vm.store(proxy, ERC1967_BEACON_SLOT, bytes32(uint(uint160(beacon))));
        return proxy;
    }

    function setAdmin(address proxy, address admin) internal returns (address) {
        vm.store(proxy, ERC1967_ADMIN_SLOT, bytes32(uint(uint160(admin))));
        return proxy;
    }

    function emptyContract() external returns (address deployed) {
        assembly {
            // Deploy a contract with a 1 opcode (STOP).
            deployed := create(0x00, 0x00, 1)
        }
        require(deployed != address(0), "Deployment failed");
    }
}

contract ConfigParserTest is Test {
    using ConfigParser for *;

    function test_ParseTOML() public {
        vm.createSelectFork(vm.rpcUrl("mainnet"), 22_181_590);
        Config memory c = ConfigParser.parse("./script/configs/mainnet/mainnet-addresses.toml");

        assertEq(c.governance.communityMultisig, 0xFEA47018D632A77bA579846c840d5706705Dc598);
        assertEq(c.governance.executorMultisig, 0x369e6F597e22EaB55fFb173C6d9cD234BD699111);
        assertEq(c.governance.operationsMultisig, 0xBE1685C81aA44FF9FB319dD389addd9374383e90);
        assertEq(c.governance.pauserMultisig, 0x5050389572f2d220ad927CcbeA0D406831012390);
        assertEq(address(c.governance.pauserRegistry), 0x0c431C66F4dE941d089625E5B423D00707977060);
        assertEq(address(c.governance.proxyAdmin), 0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444);
        assertEq(address(c.governance.timelock), 0xA6Db1A8C5a981d1536266D2a393c5F8dDb210EAF);

        assertEq(address(c.tokens.bEIGEN), 0x83E9115d334D248Ce39a6f36144aEaB5b3456e75);
        assertEq(address(c.tokens.EIGEN), 0xec53bF9167f50cDEB3Ae105f56099aaaB9061F83);

        assertEq(address(c.core.allocationManager), 0x0000000000000000000000000000000000000000);
        assertEq(address(c.core.avsDirectory), 0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF);
        assertEq(address(c.core.delegationManager), 0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A);
        assertEq(address(c.core.permissionController), 0x0000000000000000000000000000000000000000);
        assertEq(address(c.core.rewardsCoordinator), 0x7750d328b314EfFa365A0402CcfD489B80B0adda);
        assertEq(address(c.core.strategyManager), 0x858646372CC42E1A627fcE94aa7A7033e7CF075A);

        assertEq(address(c.pods.eigenPodBeacon), 0x5a2a4F2F3C18f09179B6703e63D9eDD165909073);
        assertEq(address(c.pods.eigenPodManager), 0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338);
        assertEq(address(c.pods.eigenStrategy), 0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7);

        assertEq(address(c.strategies.strategyFactory), 0x5e4C39Ad7A3E881585e383dB9827EB4811f6F647);
        assertEq(address(c.strategies.strategyFactoryBeacon), 0x0ed6703C298d28aE0878d1b28e88cA87F9662fE9);
        assertEq(c.strategies.strategyAddresses.length, 12);
        assertEq(address(c.strategies.strategyAddresses[0]), 0x93c4b944D05dfe6df7645A86cd2206016c51564D); // stETH
        assertEq(address(c.strategies.strategyAddresses[1]), 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2); // rETH
        assertEq(address(c.strategies.strategyAddresses[2]), 0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc); // cbETH
        assertEq(address(c.strategies.strategyAddresses[3]), 0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d); // ETHx
        assertEq(address(c.strategies.strategyAddresses[4]), 0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff); // ankrETH
        assertEq(address(c.strategies.strategyAddresses[5]), 0xa4C637e0F704745D182e4D38cAb7E7485321d059); // oETH
        assertEq(address(c.strategies.strategyAddresses[6]), 0x57ba429517c3473B6d34CA9aCd56c0e735b94c02); // osETH
        assertEq(address(c.strategies.strategyAddresses[7]), 0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6); // swETH
        assertEq(address(c.strategies.strategyAddresses[8]), 0x7CA911E83dabf90C90dD3De5411a10F1A6112184); // wBETH
        assertEq(address(c.strategies.strategyAddresses[9]), 0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6); // sfrxETH
        assertEq(address(c.strategies.strategyAddresses[10]), 0xAe60d8180437b5C34bB956822ac2710972584473); // lsETH
        assertEq(address(c.strategies.strategyAddresses[11]), 0x298aFB19A105D59E74658C4C334Ff360BadE6dd2); // mETH
    }
}
