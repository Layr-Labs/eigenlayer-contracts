// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/Slasher.sol";
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/core/RewardsCoordinator.sol";

import "../../src/contracts/strategies/StrategyFactory.sol";
import "../../src/contracts/strategies/StrategyBase.sol";
import "../../src/contracts/strategies/StrategyBaseTVLLimits.sol";
import "../../src/contracts/strategies/EigenStrategy.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";
import "../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../../src/contracts/permissions/PauserRegistry.sol";

import "../../src/test/mocks/EmptyContract.sol";

import "../../src/contracts/interfaces/IBackingEigen.sol";
import "../../src/contracts/interfaces/IEigen.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

struct StrategyUnderlyingTokenConfig {
    address tokenAddress;
    string tokenName;
    string tokenSymbol;
}

struct DeployedEigenPods {
    address[] multiValidatorPods;
    address[] singleValidatorPods;
    address[] inActivePods;
}

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
    RewardsCoordinator public rewardsCoordinator;
    RewardsCoordinator public rewardsCoordinatorImplementation;
    EigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    DelayedWithdrawalRouter public delayedWithdrawalRouter;
    DelayedWithdrawalRouter public delayedWithdrawalRouterImplementation;
    IBeaconChainOracle beaconOracle;
    UpgradeableBeacon public eigenPodBeacon;
    EigenPod public eigenPodImplementation;
    StrategyBase public baseStrategyImplementation;
    StrategyFactory public strategyFactory;
    StrategyFactory public strategyFactoryImplementation;
    UpgradeableBeacon public strategyBeacon;

    // Token
    ProxyAdmin public tokenProxyAdmin;
    IEigen public EIGEN;
    IEigen public EIGENImpl;
    IBackingEigen public bEIGEN;
    IBackingEigen public bEIGENImpl;
    EigenStrategy public eigenStrategy;
    EigenStrategy public eigenStrategyImpl;

    // EigenPods deployed
    address[] public multiValidatorPods;
    address[] public singleValidatorPods;
    address[] public inActivePods;
    // All eigenpods is just single array list of above eigenPods
    address[] public allEigenPods;

    EmptyContract public emptyContract;

    address executorMultisig;
    address operationsMultisig;
    address communityMultisig;
    address pauserMultisig;
    address timelock;

    // strategies deployed
    uint256 numStrategiesDeployed;
    StrategyBase[] public deployedStrategyArray;
    // Strategies to Deploy
    uint256 numStrategiesToDeploy;
    StrategyUnderlyingTokenConfig[] public strategiesToDeploy;

    /// @notice Initialization Params for first initial deployment scripts
    // StrategyManager
    uint256 STRATEGY_MANAGER_INIT_PAUSED_STATUS;
    address STRATEGY_MANAGER_WHITELISTER;
    // SLasher
    uint256 SLASHER_INIT_PAUSED_STATUS;
    // DelegationManager
    uint256 DELEGATION_MANAGER_INIT_PAUSED_STATUS;
    uint256 DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS;
    // AVSDirectory
    uint256 AVS_DIRECTORY_INIT_PAUSED_STATUS;
    // RewardsCoordinator
    uint256 REWARDS_COORDINATOR_INIT_PAUSED_STATUS;
    uint32 REWARDS_COORDINATOR_MAX_REWARDS_DURATION;
    uint32 REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH;
    uint32 REWARDS_COORDINATOR_MAX_FUTURE_LENGTH;
    uint32 REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP;
    address REWARDS_COORDINATOR_UPDATER;
    uint32 REWARDS_COORDINATOR_ACTIVATION_DELAY;
    uint32 REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS;
    uint32 REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS;
    // EigenPodManager
    uint256 EIGENPOD_MANAGER_INIT_PAUSED_STATUS;
    uint64 EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP;
    // EigenPod
    uint64 EIGENPOD_GENESIS_TIME;
    uint64 EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR;
    address ETHPOSDepositAddress;
    uint256 DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS;

    // one week in blocks -- 50400
    uint32 DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS;

    // Strategy Deployment
    uint256 STRATEGY_MAX_PER_DEPOSIT;
    uint256 STRATEGY_MAX_TOTAL_DEPOSITS;

    /// @notice use for parsing already deployed EigenLayer contracts
    function _parseDeployedContracts(string memory existingDeploymentInfoPath) internal virtual {
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
        timelock = stdJson.readAddress(existingDeploymentData, ".parameters.timelock");

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
        delegationManager = DelegationManager(
            stdJson.readAddress(existingDeploymentData, ".addresses.delegationManager")
        );
        delegationManagerImplementation = DelegationManager(
            stdJson.readAddress(existingDeploymentData, ".addresses.delegationManagerImplementation")
        );
        avsDirectory = AVSDirectory(stdJson.readAddress(existingDeploymentData, ".addresses.avsDirectory"));
        avsDirectoryImplementation = AVSDirectory(
            stdJson.readAddress(existingDeploymentData, ".addresses.avsDirectoryImplementation")
        );
        rewardsCoordinator = RewardsCoordinator(
            stdJson.readAddress(existingDeploymentData, ".addresses.rewardsCoordinator")
        );
        rewardsCoordinatorImplementation = RewardsCoordinator(
            stdJson.readAddress(existingDeploymentData, ".addresses.rewardsCoordinatorImplementation")
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
        beaconOracle = IBeaconChainOracle(stdJson.readAddress(existingDeploymentData, ".addresses.beaconOracle"));
        eigenPodBeacon = UpgradeableBeacon(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodBeacon"));
        eigenPodImplementation = EigenPod(
            payable(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodImplementation"))
        );
        baseStrategyImplementation = StrategyBase(
            stdJson.readAddress(existingDeploymentData, ".addresses.baseStrategyImplementation")
        );
        emptyContract = EmptyContract(stdJson.readAddress(existingDeploymentData, ".addresses.emptyContract"));

        // Strategies Deployed, load strategy list
        numStrategiesDeployed = stdJson.readUint(existingDeploymentData, ".addresses.numStrategiesDeployed");
        for (uint256 i = 0; i < numStrategiesDeployed; ++i) {
            // Form the key for the current element
            string memory key = string.concat(".addresses.strategyAddresses[", vm.toString(i), "]");

            // Use the key and parse the strategy address
            address strategyAddress = abi.decode(stdJson.parseRaw(existingDeploymentData, key), (address));
            deployedStrategyArray.push(StrategyBase(strategyAddress));
        }

        // token
        tokenProxyAdmin = ProxyAdmin(stdJson.readAddress(existingDeploymentData, ".addresses.token.tokenProxyAdmin"));
        EIGEN = IEigen(stdJson.readAddress(existingDeploymentData, ".addresses.token.EIGEN"));
        EIGENImpl = IEigen(stdJson.readAddress(existingDeploymentData, ".addresses.token.EIGENImpl"));
        bEIGEN = IBackingEigen(stdJson.readAddress(existingDeploymentData, ".addresses.token.bEIGEN"));
        bEIGENImpl = IBackingEigen(stdJson.readAddress(existingDeploymentData, ".addresses.token.bEIGENImpl"));
        eigenStrategy = EigenStrategy(stdJson.readAddress(existingDeploymentData, ".addresses.token.eigenStrategy"));
        eigenStrategyImpl = EigenStrategy(stdJson.readAddress(existingDeploymentData, ".addresses.token.eigenStrategyImpl"));
    }

    function _parseDeployedEigenPods(string memory existingDeploymentInfoPath) internal returns (DeployedEigenPods memory) {
        uint256 currentChainId = block.chainid;

        // READ JSON CONFIG DATA
        string memory existingDeploymentData = vm.readFile(existingDeploymentInfoPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = stdJson.readUint(existingDeploymentData, ".chainInfo.chainId");
        require(configChainId == currentChainId, "You are on the wrong chain for this config");

        multiValidatorPods = stdJson.readAddressArray(existingDeploymentData, ".eigenPods.multiValidatorPods");
        singleValidatorPods = stdJson.readAddressArray(existingDeploymentData, ".eigenPods.singleValidatorPods");
        inActivePods = stdJson.readAddressArray(existingDeploymentData, ".eigenPods.inActivePods");
        allEigenPods = stdJson.readAddressArray(existingDeploymentData, ".eigenPods.allEigenPods");
        return DeployedEigenPods({
            multiValidatorPods: multiValidatorPods,
            singleValidatorPods: singleValidatorPods,
            inActivePods: inActivePods
        });
    }

    /// @notice use for deploying a new set of EigenLayer contracts
    /// Note that this does require multisigs to already be deployed
    function _parseInitialDeploymentParams(string memory initialDeploymentParamsPath) internal virtual {
        // read and log the chainID
        uint256 currentChainId = block.chainid;
        emit log_named_uint("You are parsing on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory initialDeploymentData = vm.readFile(initialDeploymentParamsPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = stdJson.readUint(initialDeploymentData, ".chainInfo.chainId");
        require(configChainId == currentChainId, "You are on the wrong chain for this config");

        // read beacon oracle
        beaconOracle = IBeaconChainOracle(stdJson.readAddress(initialDeploymentData, ".beaconOracleAddress"));

        // read all of the deployed addresses
        executorMultisig = stdJson.readAddress(initialDeploymentData, ".multisig_addresses.executorMultisig");
        operationsMultisig = stdJson.readAddress(initialDeploymentData, ".multisig_addresses.operationsMultisig");
        communityMultisig = stdJson.readAddress(initialDeploymentData, ".multisig_addresses.communityMultisig");
        pauserMultisig = stdJson.readAddress(initialDeploymentData, ".multisig_addresses.pauserMultisig");

        // Strategies to Deploy, load strategy list
        numStrategiesToDeploy = stdJson.readUint(initialDeploymentData, ".strategies.numStrategies");
        STRATEGY_MAX_PER_DEPOSIT = stdJson.readUint(initialDeploymentData, ".strategies.MAX_PER_DEPOSIT");
        STRATEGY_MAX_TOTAL_DEPOSITS = stdJson.readUint(initialDeploymentData, ".strategies.MAX_TOTAL_DEPOSITS");
        for (uint256 i = 0; i < numStrategiesToDeploy; ++i) {
            // Form the key for the current element
            string memory key = string.concat(".strategies.strategiesToDeploy[", vm.toString(i), "]");

            // Use parseJson with the key to get the value for the current element
            bytes memory tokenInfoBytes = stdJson.parseRaw(initialDeploymentData, key);

            // Decode the token information into the Token struct
            StrategyUnderlyingTokenConfig memory tokenInfo = abi.decode(
                tokenInfoBytes,
                (StrategyUnderlyingTokenConfig)
            );

            strategiesToDeploy.push(tokenInfo);
        }

        // Read initialize params for upgradeable contracts
        STRATEGY_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(
            initialDeploymentData,
            ".strategyManager.init_paused_status"
        );
        STRATEGY_MANAGER_WHITELISTER = stdJson.readAddress(
            initialDeploymentData,
            ".strategyManager.init_strategy_whitelister"
        );
        // Slasher
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
        // RewardsCoordinator
        REWARDS_COORDINATOR_INIT_PAUSED_STATUS = stdJson.readUint(
            initialDeploymentData,
            ".rewardsCoordinator.init_paused_status"
        );
        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS = uint32(stdJson.readUint(initialDeploymentData, ".rewardsCoordinator.CALCULATION_INTERVAL_SECONDS"));
        REWARDS_COORDINATOR_MAX_REWARDS_DURATION = uint32(stdJson.readUint(initialDeploymentData, ".rewardsCoordinator.MAX_REWARDS_DURATION"));
        REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH = uint32(stdJson.readUint(initialDeploymentData, ".rewardsCoordinator.MAX_RETROACTIVE_LENGTH"));
        REWARDS_COORDINATOR_MAX_FUTURE_LENGTH = uint32(stdJson.readUint(initialDeploymentData, ".rewardsCoordinator.MAX_FUTURE_LENGTH"));
        REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP = uint32(stdJson.readUint(initialDeploymentData, ".rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP"));
        REWARDS_COORDINATOR_UPDATER = stdJson.readAddress(initialDeploymentData, ".rewardsCoordinator.rewards_updater_address");
        REWARDS_COORDINATOR_ACTIVATION_DELAY = uint32(stdJson.readUint(initialDeploymentData, ".rewardsCoordinator.activation_delay"));
        REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS = uint32(
            stdJson.readUint(initialDeploymentData, ".rewardsCoordinator.global_operator_commission_bips")
        );
        // AVSDirectory
        AVS_DIRECTORY_INIT_PAUSED_STATUS = stdJson.readUint(initialDeploymentData, ".avsDirectory.init_paused_status");
        // EigenPodManager
        EIGENPOD_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(
            initialDeploymentData,
            ".eigenPodManager.init_paused_status"
        );
        EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP = uint64(
            stdJson.readUint(initialDeploymentData, ".eigenPodManager.deneb_fork_timestamp")
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
    function _verifyContractPointers() internal view virtual {
        // AVSDirectory
        require(
            avsDirectory.delegation() == delegationManager,
            "avsDirectory: delegationManager address not set correctly"
        );
        // RewardsCoordinator
        require(
            rewardsCoordinator.delegationManager() == delegationManager,
            "rewardsCoordinator: delegationManager address not set correctly"
        );
        require(
            rewardsCoordinator.strategyManager() == strategyManager,
            "rewardsCoordinator: strategyManager address not set correctly"
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
    function _verifyImplementations() internal view virtual {
        require(
            eigenLayerProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(avsDirectory)))) ==
                address(avsDirectoryImplementation),
            "avsDirectory: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(rewardsCoordinator)))
            ) == address(rewardsCoordinatorImplementation),
            "rewardsCoordinator: implementation set incorrectly"
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
    function _verifyContractsInitialized(bool isInitialDeployment) internal virtual {
        // AVSDirectory
        vm.expectRevert(bytes("Initializable: contract is already initialized"));
        avsDirectory.initialize(address(0), eigenLayerPauserReg, AVS_DIRECTORY_INIT_PAUSED_STATUS);
        // RewardsCoordinator
        vm.expectRevert(bytes("Initializable: contract is already initialized"));
        rewardsCoordinator.initialize(
            address(0),
            eigenLayerPauserReg,
            0, // initialPausedStatus
            address(0), // rewardsUpdater
            0, // activationDelay
            0 // globalCommissionBips
        );
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
        eigenPodManager.initialize(beaconOracle, address(0), eigenLayerPauserReg, EIGENPOD_MANAGER_INIT_PAUSED_STATUS);
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
    function _verifyInitializationParams() internal view virtual {
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
        // RewardsCoordinator
        require(
            rewardsCoordinator.pauserRegistry() == eigenLayerPauserReg,
            "rewardsCoordinator: pauser registry not set correctly"
        );
        require(
            rewardsCoordinator.owner() == executorMultisig,
            "rewardsCoordinator: owner not set correctly"
        );
        require(
            rewardsCoordinator.paused() == REWARDS_COORDINATOR_INIT_PAUSED_STATUS,
            "rewardsCoordinator: init paused status set incorrectly"
        );
        require(
            rewardsCoordinator.MAX_REWARDS_DURATION() == REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            "rewardsCoordinator: maxRewardsDuration not set correctly"
        );
        require(
            rewardsCoordinator.MAX_RETROACTIVE_LENGTH() == REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            "rewardsCoordinator: maxRetroactiveLength not set correctly"
        );
        require(
            rewardsCoordinator.MAX_FUTURE_LENGTH() == REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            "rewardsCoordinator: maxFutureLength not set correctly"
        );
        require(
            rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP() == REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP,
            "rewardsCoordinator: genesisRewardsTimestamp not set correctly"
        );
        require(
            rewardsCoordinator.rewardsUpdater() == REWARDS_COORDINATOR_UPDATER,
            "rewardsCoordinator: rewardsUpdater not set correctly"
        );
        require(
            rewardsCoordinator.activationDelay() == REWARDS_COORDINATOR_ACTIVATION_DELAY,
            "rewardsCoordinator: activationDelay not set correctly"
        );
        require(
            rewardsCoordinator.CALCULATION_INTERVAL_SECONDS() == REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            "rewardsCoordinator: CALCULATION_INTERVAL_SECONDS not set correctly"
        );
        require(
            rewardsCoordinator.globalOperatorCommissionBips() == REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS,
            "rewardsCoordinator: globalCommissionBips not set correctly"
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
        if (block.chainid == 1) {
            require(
                strategyManager.strategyWhitelister() == operationsMultisig,
                "strategyManager: strategyWhitelister not set correctly"
            );
        } else if (block.chainid == 17000) {
            // On holesky, for ease of whitelisting we set to executorMultisig
            require(
                strategyManager.strategyWhitelister() == executorMultisig,
                "strategyManager: strategyWhitelister not set correctly"
            );    
        }
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
        require(
            eigenPodManager.denebForkTimestamp() == EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP,
            "eigenPodManager: denebForkTimestamp not set correctly"
        );
        require(
            eigenPodManager.beaconChainOracle() == beaconOracle,
            "eigenPodManager: beaconChainOracle not set correctly"
        );
        require(
            eigenPodManager.ethPOS() == IETHPOSDeposit(ETHPOSDepositAddress),
            "eigenPodManager: ethPOS not set correctly"
        );
        // EigenPodBeacon
        require(eigenPodBeacon.owner() == executorMultisig, "eigenPodBeacon: owner not set correctly");
        // EigenPodImplementation
        require(
            eigenPodImplementation.GENESIS_TIME() == EIGENPOD_GENESIS_TIME,
            "eigenPodImplementation: GENESIS TIME not set correctly"
        );
        require(
            eigenPodImplementation.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() ==
                EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR &&
                EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR % 1 gwei == 0,
            "eigenPodImplementation: MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR not set correctly"
        );
        require(
            eigenPodImplementation.ethPOS() == IETHPOSDeposit(ETHPOSDepositAddress),
            "eigenPodImplementation: ethPOS not set correctly"
        );
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
            require(
                strategyManager.strategyIsWhitelistedForDeposit(deployedStrategyArray[i]),
                "StrategyBaseTVLLimits: strategy should be whitelisted"
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
        emit log_named_address("STRATEGY_MANAGER_WHITELISTER", STRATEGY_MANAGER_WHITELISTER);
        emit log_named_uint("SLASHER_INIT_PAUSED_STATUS", SLASHER_INIT_PAUSED_STATUS);
        emit log_named_uint(
            "DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS",
            DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS
        );
        emit log_named_uint("DELEGATION_MANAGER_INIT_PAUSED_STATUS", DELEGATION_MANAGER_INIT_PAUSED_STATUS);
        emit log_named_uint("AVS_DIRECTORY_INIT_PAUSED_STATUS", AVS_DIRECTORY_INIT_PAUSED_STATUS);
        emit log_named_uint("REWARDS_COORDINATOR_INIT_PAUSED_STATUS", REWARDS_COORDINATOR_INIT_PAUSED_STATUS);
        // todo log all rewards coordinator params
        emit log_named_uint("EIGENPOD_MANAGER_INIT_PAUSED_STATUS", EIGENPOD_MANAGER_INIT_PAUSED_STATUS);
        emit log_named_uint("EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP", EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP);
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

        emit log_string("==== Strategies to Deploy ====");
        for (uint256 i = 0; i < numStrategiesToDeploy; ++i) {
            // Decode the token information into the Token struct
            StrategyUnderlyingTokenConfig memory tokenInfo = strategiesToDeploy[i];

            strategiesToDeploy.push(tokenInfo);
            emit log_named_address("TOKEN ADDRESS", tokenInfo.tokenAddress);
            emit log_named_string("TOKEN NAME", tokenInfo.tokenName);
            emit log_named_string("TOKEN SYMBOL", tokenInfo.tokenSymbol);
        }
    }

    /**
     * @notice Log contract addresses and write to output json file
     */
    function logAndOutputContractAddresses(string memory outputPath) public {
        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_strategies = "strategies";
        for (uint256 i = 0; i < numStrategiesToDeploy; ++i) {
            vm.serializeAddress(
                deployed_strategies,
                strategiesToDeploy[i].tokenSymbol,
                address(deployedStrategyArray[i])
            );
        }
        string memory deployed_strategies_output = numStrategiesToDeploy == 0
            ? ""
            : vm.serializeAddress(
                deployed_strategies,
                strategiesToDeploy[numStrategiesToDeploy - 1].tokenSymbol,
                address(deployedStrategyArray[numStrategiesToDeploy - 1])
            );

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "eigenLayerProxyAdmin", address(eigenLayerProxyAdmin));
        vm.serializeAddress(deployed_addresses, "eigenLayerPauserReg", address(eigenLayerPauserReg));
        vm.serializeAddress(deployed_addresses, "slasher", address(slasher));
        vm.serializeAddress(deployed_addresses, "slasherImplementation", address(slasherImplementation));
        vm.serializeAddress(deployed_addresses, "avsDirectory", address(avsDirectory));
        vm.serializeAddress(deployed_addresses, "avsDirectoryImplementation", address(avsDirectoryImplementation));
        vm.serializeAddress(deployed_addresses, "delegationManager", address(delegationManager));
        vm.serializeAddress(
            deployed_addresses,
            "delegationManagerImplementation",
            address(delegationManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "strategyManager", address(strategyManager));
        vm.serializeAddress(
            deployed_addresses,
            "strategyManagerImplementation",
            address(strategyManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "rewardsCoordinator", address(rewardsCoordinator));
        vm.serializeAddress(
            deployed_addresses,
            "rewardsCoordinatorImplementation",
            address(rewardsCoordinatorImplementation)
        );
        vm.serializeAddress(deployed_addresses, "eigenPodManager", address(eigenPodManager));
        vm.serializeAddress(
            deployed_addresses,
            "eigenPodManagerImplementation",
            address(eigenPodManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "delayedWithdrawalRouter", address(delayedWithdrawalRouter));
        vm.serializeAddress(
            deployed_addresses,
            "delayedWithdrawalRouterImplementation",
            address(delayedWithdrawalRouterImplementation)
        );
        vm.serializeAddress(deployed_addresses, "beaconOracle", address(beaconOracle));
        vm.serializeAddress(deployed_addresses, "eigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(deployed_addresses, "eigenPodImplementation", address(eigenPodImplementation));
        vm.serializeAddress(deployed_addresses, "baseStrategyImplementation", address(baseStrategyImplementation));
        vm.serializeAddress(deployed_addresses, "emptyContract", address(emptyContract));
        string memory deployed_addresses_output = vm.serializeString(
            deployed_addresses,
            "strategies",
            deployed_strategies_output
        );

        string memory parameters = "parameters";
        vm.serializeAddress(parameters, "executorMultisig", executorMultisig);
        vm.serializeAddress(parameters, "operationsMultisig", operationsMultisig);
        vm.serializeAddress(parameters, "communityMultisig", communityMultisig);
        vm.serializeAddress(parameters, "pauserMultisig", pauserMultisig);
        vm.serializeAddress(parameters, "timelock", timelock);
        string memory parameters_output = vm.serializeAddress(parameters, "operationsMultisig", operationsMultisig);

        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", block.chainid);

        // serialize all the data
        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        vm.serializeString(parent_object, chain_info, chain_info_output);
        string memory finalJson = vm.serializeString(parent_object, parameters, parameters_output);

        vm.writeJson(finalJson, outputPath);
    }
}
