// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./ServiceManagerMock.sol";
import "./StrategyToken.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IStrategyFactory.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/interfaces/IPermissionController.sol";
import {OperatorSet} from "src/contracts/libraries/OperatorSetLib.sol";

contract MarketplaceStateInitializer is Script {
    using Strings for uint256;

    // Contracts
    IDelegationManager delegationManager;
    IStrategyFactory strategyFactory;
    IAllocationManager allocationManager;
    IStrategyManager strategyManager;
    IPermissionController permissionController;
    IAVSDirectory avsDirectory;
    IRewardsCoordinator rewardsCoordinator;

    // Local deploy helpers
    string MNEMONIC;
    address proxyAdmin;
    address superAdmin;
    address[] avsAddresses;
    address[] operatorAddresses;
    address[] strategyAddresses;
    string statePathToParse = "script/utils/seedPreprod/state.json";

    ///
    ///                        Helpers/Modifiers
    ///

    // Parse state for the transaction submitter script
    modifier parseState() {
        // Set contracts
        delegationManager = IDelegationManager(0xC9366ab4A299e0937EC15A6C256C4481C05A24fD);
        strategyFactory = IStrategyFactory(0x462ac07BE300865D774A67C35F94065e8b2b8bDB);
        allocationManager = IAllocationManager(0x720b23A1fd296ae030EF56B435f25b263E3c2632);
        strategyManager = IStrategyManager(0xb22Ef643e1E067c994019A4C19e403253C05c2B0);
        permissionController = IPermissionController(0xf95e6E3400475F6000Fbb0557B776D25067128a2);
        avsDirectory = IAVSDirectory(0xA5a53a6Df5A0807FBDFe70eAe1E7b0fB984F6710);
        rewardsCoordinator = IRewardsCoordinator(0xfc67A24cDb9EE82F6111082274c62D8b2Aae3abC);

        // Set mnemonic
        MNEMONIC = vm.envString("MNEMONIC");

        // Set avs addresses from state.json
        string memory stateData = vm.readFile(statePathToParse);
        proxyAdmin = stdJson.readAddress(stateData, ".avsProxyAdmin");
        superAdmin = stdJson.readAddress(stateData, ".superAdmin");

        // AVSs deployed
        uint256 numAVSsDeployed = stdJson.readUint(stateData, ".numAVSsDeployed");
        for (uint256 i = 0; i < numAVSsDeployed; ++i) {
            // Form the key for the current element
            string memory key = string.concat(".avss[", vm.toString(i), "]");
            // Use the key and parse the strategy address
            address avs = abi.decode(stdJson.parseRaw(stateData, key), (address));
            avsAddresses.push(avs);
        }

        // Operators deployed
        numOperatorsDeployed = stdJson.readUint(stateData, ".numOperatorsDeployed");
        for (uint256 i = 0; i < numOperatorsDeployed; ++i) {
            // Form the key for the current element
            string memory key = string.concat(".operators[", vm.toString(i), "]");
            // Use the key and parse the strategy address
            address operator = abi.decode(stdJson.parseRaw(stateData, key), (address));
            operatorAddresses.push(operator);
        }

        // Strategies deployed
        uint256 numStrategiesDeployed = stdJson.readUint(stateData, ".numStrategiesDeployed");
        for (uint256 j = 0; j < numStrategiesDeployed; ++j) {
            // Form the key for the current element
            string memory key = string.concat(".strategies[", vm.toString(j), "]");
            // Use the key and parse the strategy address
            address strategy = abi.decode(stdJson.parseRaw(stateData, key), (address));
            strategyAddresses.push(strategy);
        }
        _;
    }

    /// @notice The superadmin has admin rights over every user in the system:
    /// 1. Can act on behalf of operators
    /// 2. Can act on behalf of AVSs
    /// 3. Can upgrade AVS contracts
    function _broadcastSuperAdmin() internal {
        (address admin,) = deriveRememberKey(MNEMONIC, uint32(0));
        vm.startBroadcast(admin);
    }

    ///
    ///                         AVS Deployer
    ///

    /// @notice This script should only be called once. Don't call if proxyAdmin is not set
    /// forge script script/utils/seedPreprod/txSubmitter_full.s.sol --rpc-url $RPC_URL --sig "deployAVSProxyAdmin" --verify --etherscan-api-key $ETHERSCAN_API_KEY --broadcast
    function deployAVSProxyAdmin() external parseState {
        require(address(proxyAdmin) == address(0), "Proxy admin already set");
        _broadcastSuperAdmin();
        new ProxyAdmin();
        vm.stopBroadcast();
    }

    /// @notice Deploys upgradeable AVSs with the `ServiceManagerMock` contract
    /// @notice All ALM functions can be called by the superAdmin
    /// @notice The resulting avs address will be in `state.json`
    /// @notice forge script script/utils/seedPreprod/txSubmitter_full.s.sol --rpc-url $RPC_URL --sig "deployAVS" --verify --etherscan-api-key $ETHERSCAN_API_KEY --broadcast
    function deployAVS() external parseState {
        _broadcastSuperAdmin();
        // Deploy impl
        ServiceManagerMock serviceManager =
            new ServiceManagerMock(avsDirectory, rewardsCoordinator, permissionController);

        // Deploy proxy
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy({
            _logic: address(serviceManager),
            admin_: proxyAdmin,
            _data: abi.encodeWithSelector(ServiceManagerMock.initialize.selector, superAdmin)
        });

        // Accept adminhood
        permissionController.acceptAdmin(address(proxy));

        vm.stopBroadcast();
    }

    function createRewardsSubmission() external parseState {
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions =
            new IRewardsCoordinator.RewardsSubmission[](1);
        // Use sst0 as the reward token
        IERC20 rewardToken = IERC20(0xdE36107C2E514ff0f8a5CE92F4CEC6de65485860);
        rewardsSubmissions[0].token = rewardToken;
        rewardsSubmissions[0].amount = 100_000e18;
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers =
            new IRewardsCoordinator.StrategyAndMultiplier[](5);
        for (uint256 i = 0; i < 5; i++) {
            strategyAndMultipliers[i].multiplier = 1e18;
        }
        strategyAndMultipliers[0].strategy = IStrategy(strategyAddresses[3]);
        strategyAndMultipliers[1].strategy = IStrategy(strategyAddresses[1]);
        strategyAndMultipliers[2].strategy = IStrategy(strategyAddresses[0]);
        strategyAndMultipliers[3].strategy = IStrategy(strategyAddresses[2]);
        strategyAndMultipliers[4].strategy = IStrategy(strategyAddresses[4]);

        rewardsSubmissions[0].strategiesAndMultipliers = strategyAndMultipliers;
        rewardsSubmissions[0].duration = 1 weeks;
        rewardsSubmissions[0].startTimestamp = 1_734_048_000;

        _broadcastSuperAdmin();
        rewardToken.approve(address(avsAddresses[2]), rewardsSubmissions[0].amount);
        ServiceManagerMock(avsAddresses[2]).createAVSRewardsSubmission(rewardsSubmissions);
        vm.stopBroadcast();
    }

    /// @notice Initialize avs 0 opsets
    function createOperatorSetsAvs1() external parseState {
        // Create operatorSets
        IAllocationManager.CreateSetParams[] memory params = new IAllocationManager.CreateSetParams[](3);
        // First opset
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = IStrategy(strategyAddresses[0]);
        strategies[1] = IStrategy(strategyAddresses[1]);
        strategies[2] = IStrategy(strategyAddresses[2]);
        params[0].operatorSetId = uint32(0);
        params[0].strategies = strategies;

        // Second opset
        params[1].operatorSetId = uint32(1);
        params[1].strategies = strategies;

        // Third opset
        params[2].operatorSetId = uint32(2);
        params[2].strategies = strategies;

        _broadcastSuperAdmin();
        allocationManager.updateAVSMetadataURI(avsAddresses[0], "https://avs.eigen.xyz/avs1");
        allocationManager.setAVSRegistrar(avsAddresses[0], IAVSRegistrar(avsAddresses[0]));
        allocationManager.createOperatorSets(avsAddresses[0], params);
        vm.stopBroadcast();
    }

    function createOperatorSetsAvs2() external parseState {
        // Create operatorSets
        IAllocationManager.CreateSetParams[] memory params = new IAllocationManager.CreateSetParams[](2);
        // First opset
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = IStrategy(strategyAddresses[0]);
        strategies[1] = IStrategy(strategyAddresses[1]);
        strategies[2] = IStrategy(strategyAddresses[2]);
        params[0].operatorSetId = uint32(3);
        params[0].strategies = strategies;

        // Second opset
        params[1].operatorSetId = uint32(4);
        params[1].strategies = strategies;

        _broadcastSuperAdmin();
        allocationManager.updateAVSMetadataURI(avsAddresses[1], "https://avs.eigen.xyz/avs2");
        allocationManager.setAVSRegistrar(avsAddresses[1], IAVSRegistrar(avsAddresses[1]));
        allocationManager.createOperatorSets(avsAddresses[1], params);
        vm.stopBroadcast();
    }

    function createOperatorSetsAvs3() external parseState {
        // Create operatorSets
        IAllocationManager.CreateSetParams[] memory params = new IAllocationManager.CreateSetParams[](1);

        // First opset
        IStrategy[] memory strategies = new IStrategy[](3);
        for (uint256 i = 0; i < 3; i++) {
            strategies[i] = IStrategy(strategyAddresses[i]);
        }
        params[0].operatorSetId = uint32(5);
        params[0].strategies = strategies;

        _broadcastSuperAdmin();
        allocationManager.updateAVSMetadataURI(avsAddresses[2], "https://avs.eigen.xyz/avs3");
        allocationManager.setAVSRegistrar(avsAddresses[2], IAVSRegistrar(avsAddresses[2]));
        allocationManager.createOperatorSets(avsAddresses[2], params);
        vm.stopBroadcast();
    }

    function createOperatorSetsAvs4() external parseState {
        // Create operatorSets
        IAllocationManager.CreateSetParams[] memory params = new IAllocationManager.CreateSetParams[](1);

        // First opset
        IStrategy[] memory strategies = new IStrategy[](3);
        for (uint256 i = 0; i < 3; i++) {
            strategies[i] = IStrategy(strategyAddresses[i]);
        }
        params[0].operatorSetId = uint32(6);
        params[0].strategies = strategies;

        _broadcastSuperAdmin();
        allocationManager.updateAVSMetadataURI(avsAddresses[3], "https://avs.eigen.xyz/avs4");
        allocationManager.setAVSRegistrar(avsAddresses[3], IAVSRegistrar(avsAddresses[3]));
        allocationManager.createOperatorSets(avsAddresses[3], params);
        vm.stopBroadcast();
    }

    /// @notice Initialize as an M2 AVS
    function initializeAvs3() external parseState {
        ServiceManagerMock avs = ServiceManagerMock(avsAddresses[2]);

        _broadcastSuperAdmin();
        // Set restakeable strategies
        avs.setRestakeableStrategies(strategyAddresses);

        // Set restaked strategies
        for (uint256 i = 0; i < operatorAddresses.length; i++) {
            avs.setOperatorRestakedStrategies(operatorAddresses[i], strategyAddresses);
        }
        vm.stopBroadcast();
    }

    function initializeAvs4() external parseState {
        ServiceManagerMock avs = ServiceManagerMock(avsAddresses[3]);

        _broadcastSuperAdmin();
        // Set restakeable strategies
        avs.setRestakeableStrategies(strategyAddresses);

        // Set restaked strategies
        for (uint256 i = 0; i < operatorAddresses.length; i++) {
            avs.setOperatorRestakedStrategies(operatorAddresses[i], strategyAddresses);
        }
        vm.stopBroadcast();
    }

    ///
    ///                         Operators
    ///
    /// @notice Indices 10-100 are reserved for operators
    uint256 operatorStartIndex = 15;
    uint256 numOperatorsDeployed = 5;

    /// @notice Deploys an operator and sets its allocation delay to be instant
    /// forge script script/utils/seedPreprod/txSubmitter_full.s.sol --rpc-url $RPC_URL --sig "deployOperators(uint256)" 5 --broadcast
    function deployOperators(
        uint256 numToDeploy
    ) external parseState {
        uint256 startIndex = operatorStartIndex + numOperatorsDeployed;
        for (uint256 i = 0; i < numToDeploy; ++i) {
            (address operator,) = deriveRememberKey(MNEMONIC, uint32(startIndex + i));

            console.log("operator: ", operator);

            // Seed operator with 0.1 ETH
            _broadcastSuperAdmin();
            (bool success,) = operator.call{value: 0.1 ether}("");
            require(success, "ETH transfer failed");
            vm.stopBroadcast();

            // Get private key
            deriveRememberKey(MNEMONIC, uint32(startIndex + i));

            vm.startBroadcast(operator);

            // Register operator
            delegationManager.registerAsOperator(address(0), 0, "");

            // Add pending admin
            permissionController.addPendingAdmin(operator, superAdmin);

            vm.stopBroadcast();

            // Accept adminhood
            _broadcastSuperAdmin();
            permissionController.acceptAdmin(operator);
            vm.stopBroadcast();
        }
    }

    ///
    ///                         Strategies
    ///
    /// @notice We are going to use stETH, wETH, and BC strategies

    uint256 immutable MAX_TOKEN_SUPPLY = 1e37;

    ///
    ///                         Operator Registration
    ///

    function registerOperator1() external parseState {
        IAllocationManager.RegisterParams[] memory params = new IAllocationManager.RegisterParams[](3);

        // Opset 0/1 on AVS 1
        params[0].avs = avsAddresses[0];
        params[0].operatorSetIds = new uint32[](2);
        params[0].operatorSetIds[0] = uint32(0);
        params[0].operatorSetIds[1] = uint32(1);

        // Opset 3 on AVS 2
        params[1].avs = avsAddresses[1];
        params[1].operatorSetIds = new uint32[](1);
        params[1].operatorSetIds[0] = uint32(3);

        // Opset 5 on AVS 3
        params[2].avs = avsAddresses[2];
        params[2].operatorSetIds = new uint32[](1);
        params[2].operatorSetIds[0] = uint32(5);

        _broadcastSuperAdmin();
        allocationManager.registerForOperatorSets(operatorAddresses[0], params[0]);
        allocationManager.registerForOperatorSets(operatorAddresses[0], params[1]);
        allocationManager.registerForOperatorSets(operatorAddresses[0], params[2]);
        vm.stopBroadcast();
    }

    function registerOperator2() external parseState {
        IAllocationManager.RegisterParams[] memory params = new IAllocationManager.RegisterParams[](3);

        // Opset 0/1 on AVS 1
        params[0].avs = avsAddresses[0];
        params[0].operatorSetIds = new uint32[](2);
        params[0].operatorSetIds[0] = uint32(0);
        params[0].operatorSetIds[1] = uint32(1);

        // Opset 3/4 on AVS 2
        params[1].avs = avsAddresses[1];
        params[1].operatorSetIds = new uint32[](2);
        params[1].operatorSetIds[0] = uint32(3);
        params[1].operatorSetIds[1] = uint32(4);

        // Opset 5 on AVS 3
        params[2].avs = avsAddresses[2];
        params[2].operatorSetIds = new uint32[](1);
        params[2].operatorSetIds[0] = uint32(5);

        _broadcastSuperAdmin();
        allocationManager.registerForOperatorSets(operatorAddresses[1], params[0]);
        allocationManager.registerForOperatorSets(operatorAddresses[1], params[1]);
        allocationManager.registerForOperatorSets(operatorAddresses[1], params[2]);
        vm.stopBroadcast();
    }

    function registerOperator3() external parseState {
        IAllocationManager.RegisterParams[] memory params = new IAllocationManager.RegisterParams[](4);

        // Opset 1/2 on AVS 1
        params[0].avs = avsAddresses[0];
        params[0].operatorSetIds = new uint32[](2);
        params[0].operatorSetIds[0] = uint32(1);
        params[0].operatorSetIds[1] = uint32(2);

        // // Opset 3/4 on AVS 2
        params[1].avs = avsAddresses[1];
        params[1].operatorSetIds = new uint32[](2);
        params[1].operatorSetIds[0] = uint32(3);
        params[1].operatorSetIds[1] = uint32(4);

        // Opset 5 on AVS 3
        params[2].avs = avsAddresses[2];
        params[2].operatorSetIds = new uint32[](1);
        params[2].operatorSetIds[0] = uint32(5);

        // Opset 6 on AVS 4
        params[3].avs = avsAddresses[3];
        params[3].operatorSetIds = new uint32[](1);
        params[3].operatorSetIds[0] = uint32(6);

        _broadcastSuperAdmin();
        allocationManager.registerForOperatorSets(operatorAddresses[2], params[0]);
        allocationManager.registerForOperatorSets(operatorAddresses[2], params[1]);
        allocationManager.registerForOperatorSets(operatorAddresses[2], params[2]);
        vm.stopBroadcast();
    }

    function registerOperator4() external parseState {
        IAllocationManager.RegisterParams[] memory params = new IAllocationManager.RegisterParams[](4);

        // Opset 0/1/2 on AVS 1
        params[0].avs = avsAddresses[0];
        params[0].operatorSetIds = new uint32[](3);
        params[0].operatorSetIds[0] = uint32(0);
        params[0].operatorSetIds[1] = uint32(1);
        params[0].operatorSetIds[2] = uint32(2);

        // Opset 3/4 on AVS 2
        params[1].avs = avsAddresses[1];
        params[1].operatorSetIds = new uint32[](2);
        params[1].operatorSetIds[0] = uint32(3);
        params[1].operatorSetIds[1] = uint32(4);

        // Opset 5 on AVS 3
        params[2].avs = avsAddresses[2];
        params[2].operatorSetIds = new uint32[](1);
        params[2].operatorSetIds[0] = uint32(5);

        // Opset 6 on AVS 4
        params[3].avs = avsAddresses[3];
        params[3].operatorSetIds = new uint32[](1);
        params[3].operatorSetIds[0] = uint32(6);

        _broadcastSuperAdmin();
        allocationManager.registerForOperatorSets(operatorAddresses[3], params[0]);
        allocationManager.registerForOperatorSets(operatorAddresses[3], params[1]);
        allocationManager.registerForOperatorSets(operatorAddresses[3], params[2]);
        vm.stopBroadcast();
    }

    function registerOperator5() external parseState {
        IAllocationManager.RegisterParams[] memory params = new IAllocationManager.RegisterParams[](3);

        // Opset 0/2 on AVS 1
        params[0].avs = avsAddresses[0];
        params[0].operatorSetIds = new uint32[](2);
        params[0].operatorSetIds[0] = uint32(0);
        params[0].operatorSetIds[1] = uint32(2);

        // Opset 4 on AVS 2
        params[1].avs = avsAddresses[1];
        params[1].operatorSetIds = new uint32[](1);
        params[1].operatorSetIds[0] = uint32(4);

        // Opset 6 on AVS 4
        params[2].avs = avsAddresses[3];
        params[2].operatorSetIds = new uint32[](1);
        params[2].operatorSetIds[0] = uint32(6);

        _broadcastSuperAdmin();
        allocationManager.registerForOperatorSets(operatorAddresses[4], params[0]);
        allocationManager.registerForOperatorSets(operatorAddresses[4], params[1]);
        allocationManager.registerForOperatorSets(operatorAddresses[4], params[2]);
        vm.stopBroadcast();
    }

    // Register to M2 AVSs
    function registerOperatorsToM2Avs(
        uint256 avsIndex
    ) external parseState {
        ServiceManagerMock avs = ServiceManagerMock(avsAddresses[avsIndex]);
        // Register all operators to M2 version of Avs at avsIndex
        for (uint256 i = 0; i < operatorAddresses.length; i++) {
            (address operator, uint256 privateKey) = deriveRememberKey(MNEMONIC, uint32(operatorStartIndex + i));

            ISignatureUtilsMixinTypes.SignatureWithSaltAndExpiry memory operatorSignature;
            operatorSignature.expiry = type(uint32).max;
            operatorSignature.salt = bytes32(vm.randomUint(1, type(uint32).max));
            {
                bytes32 digestHash = avsDirectory.calculateOperatorAVSRegistrationDigestHash(
                    operator, address(avs), operatorSignature.salt, operatorSignature.expiry
                );
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, digestHash);
                operatorSignature.signature = abi.encodePacked(r, s, v);
            }

            vm.startBroadcast(operator);
            avs.registerOperatorToAVS(operatorAddresses[i], operatorSignature);
            vm.stopBroadcast();
        }
    }

    function registerAllOperatorsToOpset(
        uint256 avsIndex,
        uint32 opSetId
    ) external parseState {
        IAllocationManagerTypes.RegisterParams memory params = IAllocationManagerTypes.RegisterParams({
            avs: avsAddresses[avsIndex],
            operatorSetIds: new uint32[](1),
            data: ""
        });
        params.operatorSetIds[0] = opSetId;
        for (uint256 i = 0; i < operatorAddresses.length; i++) {
            _broadcastSuperAdmin();
            allocationManager.registerForOperatorSets(operatorAddresses[i], params);
            vm.stopBroadcast();
        }
    }

    // Note is
    function allocateOperator1() external parseState {
        OperatorSet memory opSet = OperatorSet({avs: avsAddresses[2], id: 5});

        IStrategy[] memory strategies = new IStrategy[](4);
        for (uint256 i = 0; i < 5; i++) {
            // Skip BC ETH
            if (strategyAddresses[i] == address(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0)) {
                break;
            }
            strategies[i] = IStrategy(strategyAddresses[i]);
        }
        strategies[3] = IStrategy(strategyAddresses[4]);

        uint64[] memory magnitudes = new uint64[](4);
        for (uint256 i = 0; i < magnitudes.length; i++) {
            // magnitudes[i] = 10e16;
            magnitudes[i] = 0;
        }

        // Operator 1
        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = opSet;
        params[0].strategies = strategies;
        params[0].newMagnitudes = magnitudes;

        _broadcastSuperAdmin();
        allocationManager.modifyAllocations(operatorAddresses[0], params);
        vm.stopBroadcast();
    }

    function allocateOperator2() external parseState {
        OperatorSet memory opSet = OperatorSet({avs: avsAddresses[2], id: 5});

        IStrategy[] memory strategies = new IStrategy[](5);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategies[i] = IStrategy(strategyAddresses[i]);
        }

        uint64[] memory magnitudes = new uint64[](5);
        for (uint256 i = 0; i < magnitudes.length; i++) {
            magnitudes[i] = 25e16;
        }

        // Operator 2
        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = opSet;
        params[0].strategies = strategies;
        params[0].newMagnitudes = magnitudes;

        _broadcastSuperAdmin();
        allocationManager.modifyAllocations(operatorAddresses[1], params);
        vm.stopBroadcast();
    }

    function slashOperator2() external parseState {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
        uint256[] memory wadsToSlash = new uint256[](1);
        wadsToSlash[0] = 5e17;
        IAllocationManagerTypes.SlashingParams memory params = IAllocationManagerTypes.SlashingParams({
            operator: operatorAddresses[1],
            operatorSetId: 5,
            strategies: strategies,
            wadsToSlash: wadsToSlash,
            description: "Equivocation"
        });

        _broadcastSuperAdmin();
        allocationManager.slashOperator(avsAddresses[2], params);
        vm.stopBroadcast();
    }

    function allocateOperator3() external parseState {
        OperatorSet memory opSet = OperatorSet({avs: avsAddresses[2], id: 5});

        IStrategy[] memory strategies = new IStrategy[](5);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategies[i] = IStrategy(strategyAddresses[i]);
        }

        uint64[] memory magnitudes = new uint64[](5);
        for (uint256 i = 0; i < magnitudes.length; i++) {
            magnitudes[i] = 40e16;
        }

        // Operator 3
        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = opSet;
        params[0].strategies = strategies;
        params[0].newMagnitudes = magnitudes;

        _broadcastSuperAdmin();
        allocationManager.modifyAllocations(operatorAddresses[2], params);
        vm.stopBroadcast();
    }

    function slashOperator3() external parseState {
        IStrategy[] memory strategies = new IStrategy[](5);
        strategies[0] = IStrategy(strategyAddresses[3]);
        strategies[1] = IStrategy(strategyAddresses[1]);
        strategies[2] = IStrategy(strategyAddresses[0]);
        strategies[3] = IStrategy(strategyAddresses[2]);
        strategies[4] = IStrategy(strategyAddresses[4]);

        uint256[] memory wadsToSlash = new uint256[](5);
        for (uint256 i = 0; i < wadsToSlash.length; i++) {
            wadsToSlash[i] = 20e16;
        }

        IAllocationManagerTypes.SlashingParams memory params = IAllocationManagerTypes.SlashingParams({
            operator: operatorAddresses[2],
            operatorSetId: 5,
            strategies: strategies,
            wadsToSlash: wadsToSlash,
            description: "Misbehavior Test 2"
        });

        _broadcastSuperAdmin();
        allocationManager.slashOperator(avsAddresses[2], params);
        vm.stopBroadcast();
    }

    function allocateOperator4() external parseState {
        OperatorSet memory opSet = OperatorSet({avs: avsAddresses[2], id: 5});

        IStrategy[] memory strategies = new IStrategy[](5);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategies[i] = IStrategy(strategyAddresses[i]);
        }

        uint64[] memory magnitudes = new uint64[](5);
        for (uint256 i = 0; i < magnitudes.length; i++) {
            magnitudes[i] = 75e15;
        }

        // Operator 4
        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = opSet;
        params[0].strategies = strategies;
        params[0].newMagnitudes = magnitudes;

        _broadcastSuperAdmin();
        allocationManager.modifyAllocations(operatorAddresses[3], params);
        vm.stopBroadcast();
    }

    function slashOperator4() external parseState {
        IStrategy[] memory strategies = new IStrategy[](5);
        strategies[0] = IStrategy(strategyAddresses[3]);
        strategies[1] = IStrategy(strategyAddresses[1]);
        strategies[2] = IStrategy(strategyAddresses[0]);
        strategies[3] = IStrategy(strategyAddresses[2]);
        strategies[4] = IStrategy(strategyAddresses[4]);

        uint256[] memory wadsToSlash = new uint256[](5);
        for (uint256 i = 0; i < wadsToSlash.length; i++) {
            wadsToSlash[i] = 5e17;
        }

        IAllocationManagerTypes.SlashingParams memory params = IAllocationManagerTypes.SlashingParams({
            operator: operatorAddresses[3],
            operatorSetId: 5,
            strategies: strategies,
            wadsToSlash: wadsToSlash,
            description: "Misbehavior"
        });

        _broadcastSuperAdmin();
        allocationManager.slashOperator(avsAddresses[2], params);
        vm.stopBroadcast();
    }

    function allocateOperator5() external parseState {
        OperatorSet memory opSet = OperatorSet({avs: avsAddresses[2], id: 5});

        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(0x5C8b55722f421556a2AAfb7A3EA63d4c3e514312); // only stETh

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = 1e17;

        // Operator 5
        IAllocationManager.AllocateParams[] memory params = new IAllocationManager.AllocateParams[](1);
        params[0].operatorSet = opSet;
        params[0].strategies = strategies;
        params[0].newMagnitudes = magnitudes;

        _broadcastSuperAdmin();
        allocationManager.modifyAllocations(operatorAddresses[4], params);
        vm.stopBroadcast();
    }

    ///
    ///                        Stakers
    ///

    uint256 stakerStartIndex = 100;

    function initializeStakers(
        uint256 numToInitialize
    ) external parseState {
        // Get strategies
        IStrategy[] memory strategies = new IStrategy[](2);
        strategies[0] = IStrategy(strategyAddresses[1]);
        strategies[1] = IStrategy(strategyAddresses[0]);

        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;
        uint256 startIndex = 14;

        for (uint256 i = startIndex; i < startIndex + numToInitialize; i++) {
            // Get staker
            (address staker, uint256 stakerPrivateKey) = deriveRememberKey(MNEMONIC, uint32(stakerStartIndex + i));
            console.log("staker: ", staker);

            // Seed staker with 0.1 ETH
            _broadcastSuperAdmin();
            (bool success,) = staker.call{value: 0.1 ether}("");
            require(success, "ETH transfer failed");
            vm.stopBroadcast();

            // Deposit to strategies
            for (uint256 j = 0; j < strategies.length; j++) {
                uint256 amount = vm.randomUint(1e16, 1e18);
                IERC20 token = strategies[j].underlyingToken();
                bytes memory signature =
                    _getStakerStrategyInfoManualNonce(stakerPrivateKey, staker, strategies[j], token, amount);

                _broadcastSuperAdmin();
                // Approve amount to SM
                token.approve(address(strategyManager), amount);

                // Deposit on behalf of staker
                strategyManager.depositIntoStrategyWithSignature(
                    strategies[j], token, amount, staker, type(uint32).max, signature
                );

                vm.stopBroadcast();
            }

            deriveRememberKey(MNEMONIC, uint32(stakerStartIndex + i));

            // Delegate to operator
            // uint256 operatorIndex = (i / 2) + 1;
            uint256 operatorIndex = 4;
            console.log("operator delegated to: ", operatorAddresses[operatorIndex]);

            vm.startBroadcast(staker);
            delegationManager.delegateTo(operatorAddresses[operatorIndex], approverSignatureAndExpiry, "");
            vm.stopBroadcast();
        }
    }

    function depositStakerByIndex(
        uint256 stakerIndex
    ) external parseState {
        // Get strategies
        IStrategy[] memory strategies = new IStrategy[](4);
        strategies[0] = IStrategy(strategyAddresses[0]);
        strategies[1] = IStrategy(strategyAddresses[1]);

        // Get staker
        (address staker, uint256 stakerPrivateKey) = deriveRememberKey(MNEMONIC, uint32(stakerStartIndex + stakerIndex));
        console.log("staker: ", staker);

        // Deposit to strategies
        for (uint256 j = 0; j < strategies.length; j++) {
            uint256 amount = vm.randomUint(1e16, 1e18);
            IERC20 token = strategies[j].underlyingToken();
            bytes memory signature =
                _getStakerStrategyInfoManualNonce(stakerPrivateKey, staker, strategies[j], token, amount);

            _broadcastSuperAdmin();
            // Approve amount to SM
            token.approve(address(strategyManager), amount);

            // Deposit on behalf of staker
            strategyManager.depositIntoStrategyWithSignature(
                strategies[j], token, amount, staker, type(uint32).max, signature
            );

            vm.stopBroadcast();
        }
    }

    function _getStakerStrategyInfoManualNonce(
        uint256 stakerPrivateKey,
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) internal view returns (bytes memory signature) {
        uint256 expiry = type(uint32).max;
        bytes32 digestHash = strategyManager.calculateStrategyDepositDigestHash(
            staker, strategy, token, amount, IStrategyManagerNonces(address(strategyManager)).nonces(staker), expiry
        );

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(stakerPrivateKey, digestHash);

        signature = abi.encodePacked(r, s, v);
    }

    /// TODO: need to call this
    function createRewardsV1Submission() external parseState {
        // Deploy token
        string memory name = "RewardsV1_Submission";
        string memory symbol = "RV1T";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission =
            new IRewardsCoordinator.RewardsSubmission[](1);
        // Format strategies and multipliers
        // This is to the slashing strategies
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers =
            new IRewardsCoordinator.StrategyAndMultiplier[](strategyAddresses.length);
        for (uint256 i = 0; i < strategyAddresses.length; i++) {
            strategyAndMultipliers[i].multiplier = 1e18;
        }
        strategyAndMultipliers[0].strategy = IStrategy(strategyAddresses[2]);
        strategyAndMultipliers[1].strategy = IStrategy(strategyAddresses[1]);
        strategyAndMultipliers[2].strategy = IStrategy(strategyAddresses[0]);

        // Format Range
        uint32 calculationIntervalSeconds = rewardsCoordinator.CALCULATION_INTERVAL_SECONDS();
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % calculationIntervalSeconds);
        uint32 startTimestamp = moddedCurrTimestamp - 4 days;
        uint32 duration = 28 days;

        rewardsSubmission[0].strategiesAndMultipliers = strategyAndMultipliers;
        rewardsSubmission[0].token = rewardToken;
        rewardsSubmission[0].amount = 100_000e18;
        rewardsSubmission[0].startTimestamp = startTimestamp;
        rewardsSubmission[0].duration = duration;

        _broadcastSuperAdmin();
        for (uint256 i = 0; i < rewardsSubmission.length; i++) {
            rewardsSubmission[i].token.approve(address(avsAddresses[2]), type(uint256).max);
        }
        ServiceManagerMock(avsAddresses[2]).createAVSRewardsSubmission(rewardsSubmission);
        vm.stopBroadcast();
    }

    /// TODO: need to call this eventually for retroactive rewards
    function createRewardsV2Submission() external parseState {
        // Deploy token
        string memory name = "RewardsV2_Submission";
        string memory symbol = "SRT";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.OperatorDirectedRewardsSubmission[] memory rewardsSubmission =
            new IRewardsCoordinator.OperatorDirectedRewardsSubmission[](1);
        // Format strategies and multipliers
        // This is to the slashing strategies
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers =
            new IRewardsCoordinator.StrategyAndMultiplier[](strategyAddresses.length);
        for (uint256 i = 0; i < strategyAddresses.length; i++) {
            strategyAndMultipliers[i].multiplier = 1e18;
        }
        strategyAndMultipliers[0].strategy = IStrategy(strategyAddresses[2]);
        strategyAndMultipliers[1].strategy = IStrategy(strategyAddresses[1]);
        strategyAndMultipliers[2].strategy = IStrategy(strategyAddresses[0]);

        // Format Range
        uint32 calculationIntervalSeconds = rewardsCoordinator.CALCULATION_INTERVAL_SECONDS();
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % calculationIntervalSeconds);
        uint32 startTimestamp = moddedCurrTimestamp - 4 days;
        uint32 duration = 28 days;

        IRewardsCoordinator.OperatorReward[] memory operatorRewards = new IRewardsCoordinator.OperatorReward[](5);
        operatorRewards[0].operator = operatorAddresses[1];
        operatorRewards[0].amount = 10e18;
        operatorRewards[1].operator = operatorAddresses[2];
        operatorRewards[1].amount = 20e18;
        operatorRewards[2].operator = operatorAddresses[3];
        operatorRewards[2].amount = 30e18;
        operatorRewards[3].operator = operatorAddresses[4];
        operatorRewards[3].amount = 40e18;
        operatorRewards[4].operator = operatorAddresses[0];
        operatorRewards[4].amount = 50e18;

        rewardsSubmission[0].strategiesAndMultipliers = strategyAndMultipliers;
        rewardsSubmission[0].token = rewardToken;
        rewardsSubmission[0].startTimestamp = startTimestamp;
        rewardsSubmission[0].duration = duration;
        rewardsSubmission[0].operatorRewards = operatorRewards;

        _broadcastSuperAdmin();
        for (uint256 i = 0; i < rewardsSubmission.length; i++) {
            rewardsSubmission[i].token.approve(address(rewardsCoordinator), type(uint256).max);
        }
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(avsAddresses[2], rewardsSubmission);
        vm.stopBroadcast();
    }

    function createPIRewardSubmission() external parseState {
        // Deploy token
        string memory name = "PI_Rewards_Test";
        string memory symbol = "PIRT";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission =
            new IRewardsCoordinator.RewardsSubmission[](1);
        // Format strategies and multipliers
        // This is to the slashing strategies
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers =
            new IRewardsCoordinator.StrategyAndMultiplier[](strategyAddresses.length);
        for (uint256 i = 0; i < strategyAddresses.length; i++) {
            strategyAndMultipliers[i].multiplier = 1e18;
        }
        strategyAndMultipliers[0].strategy = IStrategy(strategyAddresses[2]);
        strategyAndMultipliers[1].strategy = IStrategy(strategyAddresses[1]);
        strategyAndMultipliers[2].strategy = IStrategy(strategyAddresses[0]);

        // Format Range
        uint32 calculationIntervalSeconds = rewardsCoordinator.CALCULATION_INTERVAL_SECONDS();
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % calculationIntervalSeconds);
        uint32 startTimestamp = moddedCurrTimestamp - 3 days;
        uint32 duration = 84 days;

        rewardsSubmission[0].strategiesAndMultipliers = strategyAndMultipliers;
        rewardsSubmission[0].token = rewardToken;
        rewardsSubmission[0].amount = 1e36;
        rewardsSubmission[0].startTimestamp = startTimestamp;
        rewardsSubmission[0].duration = duration;

        vm.startBroadcast();
        for (uint256 i = 0; i < rewardsSubmission.length; i++) {
            rewardsSubmission[i].token.approve(address(rewardsCoordinator), type(uint256).max);
        }
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmission);
        vm.stopBroadcast();
    }

    function _deployToken(
        string memory name,
        string memory symbol
    ) public returns (address) {
        uint256 tokenInitialSupply = 1e36;
        _broadcastSuperAdmin();
        ERC20PresetFixedSupply rewardToken = new ERC20PresetFixedSupply(name, symbol, tokenInitialSupply, superAdmin);
        rewardToken.approve(address(rewardsCoordinator), tokenInitialSupply);
        vm.stopBroadcast();
        return address(rewardToken);
    }
}

interface IStrategyManagerNonces {
    function nonces(
        address staker
    ) external view returns (uint256);
}
