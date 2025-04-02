// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";

import "script/releases/Env.sol";

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
    address pauserRegistry;
    address proxyAdmin;
    address timelock;
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
    using StdStyle for string;

    Vm constant vm = Vm(address(uint160(uint(keccak256("hevm cheat code")))));
    bytes32 constant ERC1967_IMPLEMENTATION_SLOT = bytes32(uint(keccak256("eip1967.proxy.implementation")) - 1);
    bytes32 constant ERC1967_BEACON_SLOT = bytes32(uint(keccak256("eip1967.proxy.beacon")) - 1);
    bytes32 constant ERC1967_ADMIN_SLOT = bytes32(uint(keccak256("eip1967.proxy.admin")) - 1);

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
        c.governance.pauserRegistry = address(Env.pauserRegistry(Env.impl));
        c.governance.proxyAdmin = Env.proxyAdmin();
        c.governance.timelock = address(Env.timelockController());
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

    /// @dev Returns the implementation address of a proxy.
    function impl(address proxy) internal view returns (address) {
        return vm.load(proxy, ERC1967_IMPLEMENTATION_SLOT).toAddress();
    }

    /// @dev Returns the beacon address of a proxy.
    function beacon(address proxy) internal view returns (address) {
        return vm.load(proxy, ERC1967_BEACON_SLOT).toAddress();
    }

    /// @dev Returns the admin address of a proxy.
    function admin(address proxy) internal view returns (address) {
        return vm.load(proxy, ERC1967_ADMIN_SLOT).toAddress();
    }

    /// -----------------------------------------------------------------------
    /// Logging
    /// -----------------------------------------------------------------------

    function log(Config memory c) internal {
        console2.log("Block:", block.number);
        console2.log("Timestamp:", block.timestamp);
        // Log governance addresses
        console2.log("\nGovernance Addresses");
        c.governance.communityMultisig.log("communityMultisig");
        c.governance.executorMultisig.log("executorMultisig");
        c.governance.operationsMultisig.log("operationsMultisig");
        c.governance.pauserMultisig.log("pauserMultisig");
        c.governance.pauserRegistry.log("pauserRegistry");
        c.governance.proxyAdmin.log("proxyAdmin");
        c.governance.timelock.log("timelock");
        // Log token addresses
        console2.log("\nToken Addresses:");
        address(c.tokens.EIGEN).logProxy("EIGEN");
        address(c.tokens.bEIGEN).logProxy("bEIGEN");
        // Log core protocol addresses
        console2.log("\nCore Addresses:");
        address(c.core.allocationManager).logProxy("allocationManager");
        address(c.core.avsDirectory).logProxy("avsDirectory");
        address(c.core.delegationManager).logProxy("delegationManager");
        address(c.core.permissionController).logProxy("permissionController");
        address(c.core.rewardsCoordinator).logProxy("rewardsCoordinator");
        address(c.core.strategyManager).logProxy("strategyManager");
        // Log pod system addresses
        console2.log("\nPod Addresses:");
        address(c.pods.eigenPodBeacon).logBeaconImpl("eigenPodBeacon");
        address(c.pods.eigenPodManager).logProxy("eigenPodManager");
        address(c.pods.eigenStrategy).logProxy("eigenStrategy");
        // Log strategy system addresses
        console2.log("\nStrategy Addresses:");
        address(c.strategies.strategyFactory).logProxy("strategyFactory");
        address(c.strategies.strategyFactoryBeacon).logBeaconImpl("strategyFactoryBeacon");
        for (uint i; i < c.strategies.strategyAddresses.length; ++i) {
            address(c.strategies.strategyAddresses[i]).logProxy(string.concat("strategy", vm.toString(i)));
        }
    }

    function log(address addr, string memory name) internal {
        console2.log(name.yellow(), addr);
    }

    /// @dev Logs the implementation address of a beacon with a name.
    function logBeaconImpl(address beacon, string memory name) internal {
        console2.log(name.yellow(), beacon, vm.toString(IBeacon(beacon).implementation()).dim());
    }

    /// @dev Logs the implementation address of a proxy with a name.
    function logProxy(address proxy, string memory name) internal {
        console2.log(name.yellow(), proxy, vm.toString(proxy.impl()).dim());
    }

    function toAddress(bytes32 x) internal pure returns (address r) {
        assembly {
            r := x
        }
    }
}

contract ParseMainnetConfigTest is Test {
    using ConfigParser for *;

    function test() public {
        vm.createSelectFork(vm.rpcUrl("mainnet"), 22_181_590);
        ConfigParser.parse("./script/configs/mainnet/mainnet-addresses.toml").log();
    }
}
