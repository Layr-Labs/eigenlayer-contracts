// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/Slasher.sol";
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/core/AVSDirectory.sol";

import "../../src/contracts/strategies/StrategyBase.sol";
import "../../src/contracts/strategies/StrategyBaseTVLLimits.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";
import "../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../../src/contracts/permissions/PauserRegistry.sol";

import "../../src/test/mocks/EmptyContract.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

contract ExistingDeploymentParser is Script, Test {
    // EigenLayer Contracts
    ProxyAdmin public eigenLayerProxyAdmin;
    PauserRegistry public eigenLayerPauserReg;
    Slasher public slasher;
    Slasher public slasherImplementation;
    AVSDirectory public avsDirectory;
    AVSDirectory public avsDirectoryImplementation;
    DelegationManager public delegationManager;
    DelegationManager public delegationManagerImplementation;
    StrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;
    EigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    DelayedWithdrawalRouter public delayedWithdrawalRouter;
    DelayedWithdrawalRouter public delayedWithdrawalRouterImplementation;
    IBeaconChainOracle beaconOracle;
    UpgradeableBeacon public eigenPodBeacon;
    EigenPod public eigenPodImplementation;
    StrategyBase public baseStrategyImplementation;

    EmptyContract public emptyContract;

    address executorMultisig;
    address operationsMultisig;
    address communityMultisig;
    address pauserMultisig;

    // strategies deployed
    StrategyBase[] public deployedStrategyArray;

    // the ETH2 deposit contract -- if not on mainnet, we deploy a mock as stand-in
    IETHPOSDeposit public ethPOSDeposit;

    // // IMMUTABLES TO SET
    // uint64 MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR;
    // uint64 GOERLI_GENESIS_TIME = 1616508000;

    /// @notice Initialization Params for first initial deployment scripts
    // StrategyManager
    uint256 STRATEGY_MANAGER_INIT_PAUSED_STATUS;
    // SLasher
    uint256 SLASHER_INIT_PAUSED_STATUS;
    // DelegationManager
    uint256 DELEGATION_MANAGER_INIT_PAUSED_STATUS;
    uint256 DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS;
    // AVSDirectory
    uint256 AVS_DIRECTORY_INIT_PAUSED_STATUS;
    // EigenPodManager
    uint256 EIGENPOD_MANAGER_INIT_PAUSED_STATUS;
    uint256 EIGENPOD_MANAGER_MAX_PODS;
    // EigenPod
    uint64 EIGENPOD_GENESIS_TIME;
    uint64 EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR;
    address ETHPOSDepositAddress;
    uint256 DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS;

    // one week in blocks -- 50400
    uint32 DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS;

    /// @notice use for parsing already deployed EigenLayer contracts
    function _parseDeployedContracts(string memory existingDeploymentInfoPath) internal {
        // read and log the chainID
        uint256 currentChainId = block.chainid;
        emit log_named_uint("You are parsing on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory existingDeploymentData = vm.readFile(existingDeploymentInfoPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = stdJson.readUint(existingDeploymentData, ".chainInfo.chainId");
        require(configChainId == currentChainId, "You are on the wrong chain for this config");

        // read all of the deployed addresses
        executorMultisig = stdJson.readAddress(existingDeploymentData, ".parameters.executorMultisig");
        operationsMultisig = stdJson.readAddress(existingDeploymentData, ".parameters.operationsMultisig");
        communityMultisig = stdJson.readAddress(existingDeploymentData, ".parameters.communityMultisig");
        pauserMultisig = stdJson.readAddress(existingDeploymentData, ".parameters.pauserMultisig");

        eigenLayerProxyAdmin = ProxyAdmin(
            stdJson.readAddress(existingDeploymentData, ".addresses.eigenLayerProxyAdmin")
        );
        eigenLayerPauserReg = PauserRegistry(
            stdJson.readAddress(existingDeploymentData, ".addresses.eigenLayerPauserReg")
        );
        slasher = Slasher(stdJson.readAddress(existingDeploymentData, ".addresses.slasher"));
        slasherImplementation = Slasher(
            stdJson.readAddress(existingDeploymentData, ".addresses.slasherImplementation")
        );
        delegationManager = DelegationManager(stdJson.readAddress(existingDeploymentData, ".addresses.delegation"));
        delegationManagerImplementation = DelegationManager(
            stdJson.readAddress(existingDeploymentData, ".addresses.delegationImplementation")
        );
        avsDirectory = AVSDirectory(stdJson.readAddress(existingDeploymentData, ".addresses.avsDirectory"));
        avsDirectoryImplementation = AVSDirectory(
            stdJson.readAddress(existingDeploymentData, ".addresses.avsDirectoryImplementation")
        );
        strategyManager = StrategyManager(stdJson.readAddress(existingDeploymentData, ".addresses.strategyManager"));
        strategyManagerImplementation = StrategyManager(
            stdJson.readAddress(existingDeploymentData, ".addresses.strategyManagerImplementation")
        );
        eigenPodManager = EigenPodManager(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodManager"));
        eigenPodManagerImplementation = EigenPodManager(
            stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodManagerImplementation")
        );
        delayedWithdrawalRouter = DelayedWithdrawalRouter(
            stdJson.readAddress(existingDeploymentData, ".addresses.delayedWithdrawalRouter")
        );
        delayedWithdrawalRouterImplementation = DelayedWithdrawalRouter(
            stdJson.readAddress(existingDeploymentData, ".addresses.delayedWithdrawalRouterImplementation")
        );
        beaconOracle = IBeaconChainOracle(
            stdJson.readAddress(existingDeploymentData, ".addresses.beaconOracleAddress")
        );
        eigenPodBeacon = UpgradeableBeacon(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodBeacon"));
        eigenPodImplementation = EigenPod(
            payable(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodImplementation"))
        );
        baseStrategyImplementation = StrategyBase(
            stdJson.readAddress(existingDeploymentData, ".addresses.baseStrategyImplementation")
        );
        emptyContract = EmptyContract(stdJson.readAddress(existingDeploymentData, ".addresses.emptyContract"));

        /*
        commented out -- needs JSON formatting of the form:
        strategies": [
      {"WETH": "0x7CA911E83dabf90C90dD3De5411a10F1A6112184"},
      {"rETH": "0x879944A8cB437a5f8061361f82A6d4EED59070b5"},
      {"tsETH": "0xcFA9da720682bC4BCb55116675f16F503093ba13"},
      {"wstETH": "0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff"}]
        // load strategy list
        bytes memory strategyListRaw = stdJson.parseRaw(existingDeploymentData, ".addresses.strategies");
        address[] memory strategyList = abi.decode(strategyListRaw, (address[]));
        for (uint256 i = 0; i < strategyList.length; ++i) {
            deployedStrategyArray.push(StrategyBase(strategyList[i]));
        }
        */
    }

    /// @notice use for deploying a new set of EigenLayer contracts
    /// Note that this does require multisigs to already be deployed
    function _parseInitialDeploymentParams(string memory initialDeploymentParamsPath) internal {
        // read and log the chainID
        uint256 currentChainId = block.chainid;
        emit log_named_uint("You are parsing on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory initialDeploymentData = vm.readFile(initialDeploymentParamsPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = stdJson.readUint(initialDeploymentData, ".chainInfo.chainId");
        require(configChainId == currentChainId, "You are on the wrong chain for this config");

        // read beacon oracle
        beaconOracle = IBeaconChainOracle(stdJson.readAddress(initialDeploymentData, ".addresses.beaconOracleAddress"));

        // read all of the deployed addresses
        executorMultisig = stdJson.readAddress(initialDeploymentData, ".multisig_addresses.executorMultisig");
        operationsMultisig = stdJson.readAddress(initialDeploymentData, ".multisig_addresses.operationsMultisig");
        communityMultisig = stdJson.readAddress(initialDeploymentData, ".multisig_addresses.communityMultisig");
        pauserMultisig = stdJson.readAddress(initialDeploymentData, ".multisig_addresses.pauserMultisig");

        // Read initialize params for upgradeable contracts
        STRATEGY_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(
            initialDeploymentData,
            ".strategyManager.init_paused_status"
        );
        SLASHER_INIT_PAUSED_STATUS = stdJson.readUint(initialDeploymentData, ".slasher.init_paused_status");
        // DelegationManager
        DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS = stdJson.readUint(
            initialDeploymentData,
            ".delegationManager.init_minWithdrawalDelayBlocks"
        );
        DELEGATION_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(
            initialDeploymentData,
            ".delegationManager.init_paused_status"
        );
        // AVSDirectory
        AVS_DIRECTORY_INIT_PAUSED_STATUS = stdJson.readUint(initialDeploymentData, ".avsDirectory.init_paused_status");
        // EigenPodManager
        EIGENPOD_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(
            initialDeploymentData,
            ".eigenPodManager.init_paused_status"
        );

        // EigenPod
        EIGENPOD_GENESIS_TIME = uint64(stdJson.readUint(initialDeploymentData, ".eigenPod.GENESIS_TIME"));
        EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = uint64(
            stdJson.readUint(initialDeploymentData, ".eigenPod.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR")
        );
        ETHPOSDepositAddress = stdJson.readAddress(initialDeploymentData, ".ethPOSDepositAddress");
        // DelayedWithdrawalRouter
        DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS = stdJson.readUint(
            initialDeploymentData,
            ".delayedWithdrawalRouter.init_paused_status"
        );

        // both set to one week in blocks 50400
        DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS = uint32(
            stdJson.readUint(initialDeploymentData, ".delayedWithdrawalRouter.init_withdrawalDelayBlocks")
        );

        logInitialDeploymentParams();
    }

    /// @notice Ensure contracts point at each other correctly via constructors
    function _verifyContractPointers() internal view {
        // AVSDirectory
        require(
            avsDirectory.delegation() == delegationManager,
            "avsDirectory: delegationManager address not set correctly"
        );
        // DelegationManager
        require(delegationManager.slasher() == slasher, "delegationManager: slasher address not set correctly");
        require(
            delegationManager.strategyManager() == strategyManager,
            "delegationManager: strategyManager address not set correctly"
        );
        require(
            delegationManager.eigenPodManager() == eigenPodManager,
            "delegationManager: eigenPodManager address not set correctly"
        );
        // StrategyManager
        require(strategyManager.slasher() == slasher, "strategyManager: slasher address not set correctly");
        require(
            strategyManager.delegation() == delegationManager,
            "strategyManager: delegationManager address not set correctly"
        );
        require(
            strategyManager.eigenPodManager() == eigenPodManager,
            "strategyManager: eigenPodManager address not set correctly"
        );
        // EPM
        require(
            address(eigenPodManager.ethPOS()) == ETHPOSDepositAddress,
            "eigenPodManager: ethPOSDeposit contract address not set correctly"
        );
        require(
            eigenPodManager.eigenPodBeacon() == eigenPodBeacon,
            "eigenPodManager: eigenPodBeacon contract address not set correctly"
        );
        require(
            eigenPodManager.strategyManager() == strategyManager,
            "eigenPodManager: strategyManager contract address not set correctly"
        );
        require(eigenPodManager.slasher() == slasher, "eigenPodManager: slasher contract address not set correctly");
        require(
            eigenPodManager.delegationManager() == delegationManager,
            "eigenPodManager: delegationManager contract address not set correctly"
        );
        // DelayedWithdrawalRouter
        require(
            delayedWithdrawalRouter.eigenPodManager() == eigenPodManager,
            "delayedWithdrawalRouterContract: eigenPodManager address not set correctly"
        );
    }

    /// @notice verify implementations for Transparent Upgradeable Proxies
    function _verifyImplementations() internal view {
        require(
            eigenLayerProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(avsDirectory)))) ==
                address(avsDirectoryImplementation),
            "avsDirectory: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(delegationManager)))
            ) == address(delegationManagerImplementation),
            "delegationManager: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(strategyManager)))
            ) == address(strategyManagerImplementation),
            "strategyManager: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(slasher)))) ==
                address(slasherImplementation),
            "slasher: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(eigenPodManager)))
            ) == address(eigenPodManagerImplementation),
            "eigenPodManager: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter)))
            ) == address(delayedWithdrawalRouterImplementation),
            "delayedWithdrawalRouter: implementation set incorrectly"
        );

        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            require(
                eigenLayerProxyAdmin.getProxyImplementation(
                    TransparentUpgradeableProxy(payable(address(deployedStrategyArray[i])))
                ) == address(baseStrategyImplementation),
                "strategy: implementation set incorrectly"
            );
        }

        require(
            eigenPodBeacon.implementation() == address(eigenPodImplementation),
            "eigenPodBeacon: implementation set incorrectly"
        );
    }

    /**
     * @notice Verify initialization of Transparent Upgradeable Proxies. Also check
     * initialization params if this is the first deployment.
     * @param isInitialDeployment True if this is the first deployment of contracts from scratch
     */
    function _verifyContractsInitialized(bool isInitialDeployment) internal {
        // AVSDirectory
        vm.expectRevert(bytes("Initializable: contract is already initialized"));
        avsDirectory.initialize(address(0), eigenLayerPauserReg, AVS_DIRECTORY_INIT_PAUSED_STATUS);
        // DelegationManager
        vm.expectRevert(bytes("Initializable: contract is already initialized"));
        IStrategy[] memory initializeStrategiesToSetDelayBlocks = new IStrategy[](0);
        uint256[] memory initializeWithdrawalDelayBlocks = new uint256[](0);
        delegationManager.initialize(
            address(0),
            eigenLayerPauserReg,
            0,
            0, // minWithdrawalDelayBLocks
            initializeStrategiesToSetDelayBlocks,
            initializeWithdrawalDelayBlocks
        );
        // StrategyManager
        vm.expectRevert(bytes("Initializable: contract is already initialized"));
        strategyManager.initialize(address(0), address(0), eigenLayerPauserReg, STRATEGY_MANAGER_INIT_PAUSED_STATUS);
        // EigenPodManager
        vm.expectRevert(bytes("Initializable: contract is already initialized"));
        eigenPodManager.initialize(
            EIGENPOD_MANAGER_MAX_PODS,
            beaconOracle,
            address(0),
            eigenLayerPauserReg,
            EIGENPOD_MANAGER_INIT_PAUSED_STATUS
        );
        // DelayedWithdrawalRouter
        vm.expectRevert(bytes("Initializable: contract is already initialized"));
        delayedWithdrawalRouter.initialize(
            address(0),
            eigenLayerPauserReg,
            DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS,
            DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS
        );
        // Strategies
        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            vm.expectRevert(bytes("Initializable: contract is already initialized"));
            StrategyBaseTVLLimits(address(deployedStrategyArray[i])).initialize(
                0,
                0,
                IERC20(address(0)),
                eigenLayerPauserReg
            );
        }
    }

    /// @notice Verify params based on config constants that are updated from calling `_parseInitialDeploymentParams`
    function _verifyInitializationParams() internal {
        // AVSDirectory
        require(
            avsDirectory.pauserRegistry() == eigenLayerPauserReg,
            "avsdirectory: pauser registry not set correctly"
        );
        require(avsDirectory.owner() == executorMultisig, "avsdirectory: owner not set correctly");
        require(
            avsDirectory.paused() == AVS_DIRECTORY_INIT_PAUSED_STATUS,
            "avsdirectory: init paused status set incorrectly"
        );
        // DelegationManager
        require(
            delegationManager.pauserRegistry() == eigenLayerPauserReg,
            "delegationManager: pauser registry not set correctly"
        );
        require(delegationManager.owner() == executorMultisig, "delegationManager: owner not set correctly");
        require(
            delegationManager.paused() == DELEGATION_MANAGER_INIT_PAUSED_STATUS,
            "delegationManager: init paused status set incorrectly"
        );
        require(
            delegationManager.minWithdrawalDelayBlocks() == DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS,
            "delegationManager: minWithdrawalDelayBlocks not set correctly"
        );
        // StrategyManager
        require(
            strategyManager.pauserRegistry() == eigenLayerPauserReg,
            "strategyManager: pauser registry not set correctly"
        );
        require(strategyManager.owner() == executorMultisig, "strategyManager: owner not set correctly");
        require(
            strategyManager.paused() == STRATEGY_MANAGER_INIT_PAUSED_STATUS,
            "strategyManager: init paused status set incorrectly"
        );
        // EigenPodManager
        require(
            eigenPodManager.pauserRegistry() == eigenLayerPauserReg,
            "eigenPodManager: pauser registry not set correctly"
        );
        require(eigenPodManager.owner() == executorMultisig, "eigenPodManager: owner not set correctly");
        require(
            eigenPodManager.paused() == EIGENPOD_MANAGER_INIT_PAUSED_STATUS,
            "eigenPodManager: init paused status set incorrectly"
        );
        // EigenPodBeacon
        require(eigenPodBeacon.owner() == executorMultisig, "eigenPodBeacon: owner not set correctly");
        // DelayedWithdrawalRouter
        require(
            delayedWithdrawalRouter.pauserRegistry() == eigenLayerPauserReg,
            "delayedWithdrawalRouter: pauser registry not set correctly"
        );
        require(
            delayedWithdrawalRouter.owner() == executorMultisig,
            "delayedWithdrawalRouter: owner not set correctly"
        );
        require(
            delayedWithdrawalRouter.paused() == DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS,
            "delayedWithdrawalRouter: init paused status set incorrectly"
        );
        require(
            delayedWithdrawalRouter.withdrawalDelayBlocks() == DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS,
            "delayedWithdrawalRouter: withdrawalDelayBlocks not set correctly"
        );
        // Strategies
        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            require(
                deployedStrategyArray[i].pauserRegistry() == eigenLayerPauserReg,
                "StrategyBaseTVLLimits: pauser registry not set correctly"
            );
            require(
                deployedStrategyArray[i].paused() == 0,
                "StrategyBaseTVLLimits: init paused status set incorrectly"
            );
        }

        // Pausing Permissions
        require(eigenLayerPauserReg.isPauser(operationsMultisig), "pauserRegistry: operationsMultisig is not pauser");
        require(eigenLayerPauserReg.isPauser(executorMultisig), "pauserRegistry: executorMultisig is not pauser");
        require(eigenLayerPauserReg.isPauser(pauserMultisig), "pauserRegistry: pauserMultisig is not pauser");
        require(eigenLayerPauserReg.unpauser() == executorMultisig, "pauserRegistry: unpauser not set correctly");
    }

    function logInitialDeploymentParams() public {
        emit log_string("==== Parsed Initilize Params for Initial Deployment ====");

        emit log_named_address("executorMultisig", executorMultisig);
        emit log_named_address("operationsMultisig", operationsMultisig);
        emit log_named_address("communityMultisig", communityMultisig);
        emit log_named_address("pauserMultisig", pauserMultisig);

        emit log_named_uint("STRATEGY_MANAGER_INIT_PAUSED_STATUS", STRATEGY_MANAGER_INIT_PAUSED_STATUS);
        emit log_named_uint("SLASHER_INIT_PAUSED_STATUS", SLASHER_INIT_PAUSED_STATUS);
        emit log_named_uint(
            "DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS",
            DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS
        );
        emit log_named_uint("DELEGATION_MANAGER_INIT_PAUSED_STATUS", DELEGATION_MANAGER_INIT_PAUSED_STATUS);
        emit log_named_uint("AVS_DIRECTORY_INIT_PAUSED_STATUS", AVS_DIRECTORY_INIT_PAUSED_STATUS);
        emit log_named_uint("EIGENPOD_MANAGER_INIT_PAUSED_STATUS", EIGENPOD_MANAGER_INIT_PAUSED_STATUS);
        emit log_named_uint("EIGENPOD_MANAGER_MAX_PODS", EIGENPOD_MANAGER_MAX_PODS);
        emit log_named_uint("EIGENPOD_GENESIS_TIME", EIGENPOD_GENESIS_TIME);
        emit log_named_uint(
            "EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR",
            EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR
        );
        emit log_named_address("ETHPOSDepositAddress", ETHPOSDepositAddress);
        emit log_named_uint(
            "DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS",
            DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS
        );
        emit log_named_uint(
            "DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS",
            DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS
        );
    }

    function logContractAddresses() public {
        emit log_string("==== Contract Addresses from Deployment/Upgrade ====");
    }
}
