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

contract MarketplaceStateInitializer_Simple is Script {
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
    string statePathToParse = "script/utils/seedMarketplacePreprod/state.json";

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

    /// @notice Deploys upgradeable AVSs with the `ServiceManagerMock` contract
    /// @notice All ALM functions can be called by the superAdmin
    /// @notice The resulting avs address will be in `state.json`
    /// @notice forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "deployAVS" --verify --etherscan-api-key $ETHERSCAN_API_KEY --broadcast
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

    /// @notice Create operatorSets wth strategies
    /// @notice forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "createOperatorSetsAvs1" --broadcast
    /// @notice Update the strategy and avs address for creating more operatorSets
    function createOperatorSetsAvs1() external parseState {
        // Create operatorSets
        IAllocationManager.CreateSetParams[] memory params = new IAllocationManager.CreateSetParams[](3);
        // First opset
        IStrategy[] memory strategies = new IStrategy[](2);
        strategies[0] = IStrategy(strategyAddresses[0]);
        strategies[1] = IStrategy(strategyAddresses[1]);
        params[0].operatorSetId = uint32(0);
        params[0].strategies = strategies;

        // Second opset
        strategies = new IStrategy[](2);
        strategies[0] = IStrategy(strategyAddresses[1]);
        strategies[1] = IStrategy(strategyAddresses[2]);
        params[1].operatorSetId = uint32(1);
        params[1].strategies = strategies;

        // Third opset
        strategies = new IStrategy[](2);
        strategies[0] = IStrategy(strategyAddresses[2]);
        strategies[1] = IStrategy(strategyAddresses[3]);
        params[2].operatorSetId = uint32(2);
        params[2].strategies = strategies;

        _broadcastSuperAdmin();
        allocationManager.createOperatorSets(avsAddresses[0], params);
        vm.stopBroadcast();
    }

    /// @notice Initialize an M2 AVS with restakeable strategeis & operatorRestakedStrategies
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "initializeM2Avs(uint256)" x --broadcast
    function initializeM2AVS(
        uint256 avsIndex
    ) external parseState {
        ServiceManagerMock avs = ServiceManagerMock(avsAddresses[avsIndex]);

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
    uint256 operatorStartIndex = 10;
    uint256 numOperatorsDeployed;

    /// @notice Deploys an operator and sets its allocation delay to be instant
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "deployOperators(uint256)" 5 --broadcast
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
    /// @notice We are going to deploy 3 new strategies
    /// @notice We'll also use BC and stETH strategies, giving a total of 5 strategies

    uint256 immutable MAX_TOKEN_SUPPLY = 1e37;

    /// @notice Deploys custom strategies for custom tokens. SuperAdmin has initial supply
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "deployStrategies(uint256)" 3 --verify --etherscan-api-key $ETHERSCAN_API_KEY --broadcast
    function deployStrategies(
        uint256 numToDeploy
    ) external parseState {
        for (uint256 i = 0; i < numToDeploy; i++) {
            string memory name = string.concat("SlashingStrategyToken_", uint256(i).toString());
            string memory symbol = string.concat("SST_", uint256(i).toString());

            _broadcastSuperAdmin();
            StrategyToken token = new StrategyToken(name, symbol, MAX_TOKEN_SUPPLY, superAdmin);
            IStrategy strategy = strategyFactory.deployNewStrategy(IERC20(token));
            console.log("Token: ", address(token));
            console.log("Strategy: ", address(strategy));

            vm.stopBroadcast();
        }
    }

    ///
    ///                         Operator Registration
    ///

    /// @notice Registers operators to given operatorSets
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "registerOperator1"  -vvv --broadcast
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

    /// @notice Registers all operators to M2 AVSs
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "registerOperatorsToM2Avs(uint256)" x  -vvv --broadcast
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

    /// @notice Registers all operators to the given avsIndex and opSetId
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "registerAllOperatorsToOpset(uint256,uint32)" x x  -vvv --broadcast
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

    /// @notice Allocates operator to given operatorSet for strategies and wads
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "allocateOperator3"  -vvv --broadcast
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

    /// @notice Slashes operator for misbehavior
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "slashOperator3"  -vvv --broadcast
    function slashOperator3() external parseState {
        IStrategy[] memory strategies = new IStrategy[](5);
        strategies[0] = IStrategy(strategyAddresses[3]);
        strategies[1] = IStrategy(strategyAddresses[1]);
        strategies[2] = IStrategy(strategyAddresses[0]);
        strategies[3] = IStrategy(strategyAddresses[2]);
        strategies[4] = IStrategy(strategyAddresses[4]);

        uint256[] memory wadsToSlash = new uint256[](5);
        for (uint256 i = 0; i < wadsToSlash.length; i++) {
            wadsToSlash[i] = 10e16;
        }

        IAllocationManagerTypes.SlashingParams memory params = IAllocationManagerTypes.SlashingParams({
            operator: operatorAddresses[2],
            operatorSetId: 5,
            strategies: strategies,
            wadsToSlash: wadsToSlash,
            description: "Misbehavior"
        });

        _broadcastSuperAdmin();
        allocationManager.slashOperator(avsAddresses[2], params);
        vm.stopBroadcast();
    }

    ///
    ///                        Stakers
    ///

    uint256 stakerStartIndex = 100;

    /// @notice Initializes stakers with random amounts in strategies. Also delegates staker to given operator
    /// forge script script/utils/seedMarketplacePreprod/txSubmitter.s.sol --rpc-url $RPC_HOLESKY --sig "initializeStakers(uint256, uint256)" x x  -vvv --broadcast
    function initializeStakers(
        uint256 numToInitialize,
        uint256 operatorIndex
    ) external parseState {
        // Get strategies
        IStrategy[] memory strategies = new IStrategy[](4);
        strategies[0] = IStrategy(strategyAddresses[0]);
        strategies[1] = IStrategy(strategyAddresses[1]);
        strategies[2] = IStrategy(strategyAddresses[3]); // Skip native ETH, use pods for that
        strategies[3] = IStrategy(strategyAddresses[4]);

        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;

        for (uint256 i = 0; i < numToInitialize; i++) {
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
                uint256 amount = vm.randomUint(1e18, 4e18);
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
            console.log("operator delegated to: ", operatorAddresses[operatorIndex]);

            vm.startBroadcast(staker);
            delegationManager.delegateTo(operatorAddresses[operatorIndex], approverSignatureAndExpiry, "");
            vm.stopBroadcast();
        }
    }

    /// @notice Deposits staker into strategies
    function depositStakerByIndex(
        uint256 stakerIndex
    ) external parseState {
        // Get strategies
        IStrategy[] memory strategies = new IStrategy[](4);
        strategies[0] = IStrategy(strategyAddresses[0]);
        strategies[1] = IStrategy(strategyAddresses[1]);
        strategies[2] = IStrategy(strategyAddresses[3]); // Skip native ETH, use pods for that
        strategies[3] = IStrategy(strategyAddresses[4]);

        // Get staker
        (address staker, uint256 stakerPrivateKey) = deriveRememberKey(MNEMONIC, uint32(stakerStartIndex + stakerIndex));
        console.log("staker: ", staker);

        // Deposit to strategies
        for (uint256 j = 0; j < strategies.length; j++) {
            uint256 amount = vm.randomUint(1e18, 4e18);
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
}

interface IStrategyManagerNonces {
    function nonces(
        address staker
    ) external view returns (uint256);
}
