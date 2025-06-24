// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// OpenZeppelin Contracts
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

// Core Contracts
import "src/contracts/core/AllocationManager.sol";
import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/permissions/PauserRegistry.sol";
import "src/contracts/permissions/PermissionController.sol";

// Multichain Contracts
import "src/contracts/multichain/CrossChainRegistry.sol";
import "src/contracts/multichain/ECDSATableCalculator.sol";
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/test/mocks/EmptyContract.sol";
import "src/test/mocks/MockAVSRegistrar.sol";

// Test Utils
import "src/test/utils/OperatorWalletLib.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/devnet/mutlichain/deploy_operatorSet_ecdsa.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast --sig "run()" --verify $ETHERSCAN_API_KEY
contract DeployOperatorSetECDSA is Script, Test {
    using OperatorWalletLib for *;
    using Strings for uint256;

    Vm cheats = Vm(VM_ADDRESS);

    // Admin that can perform actions on behalf of the operatorSet
    address superAdmin = 0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293;
    uint32 operatorSetId = 1;  // Changed from 0 to 1 for ECDSA operator set

    // Contracts
    AllocationManager public allocationManager = AllocationManager(0xFdD5749e11977D60850E06bF5B13221Ad95eb6B4);
    DelegationManager public delegationManager = DelegationManager(0x75dfE5B44C2E530568001400D3f704bC8AE350CC);
    StrategyManager public strategyManager = StrategyManager(0xF9fbF2e35D8803273E214c99BF15174139f4E67a);
    PermissionController public permissionController = PermissionController(0xa2348c77802238Db39f0CefAa500B62D3FDD682b);
    IStrategy public strategy = IStrategy(0xD523267698C81a372191136e477fdebFa33D9FB4); // WETH strategy
    IERC20 public weth = IERC20(0x94373a4919B3240D86eA41593D5eBa789FEF3848);
    CrossChainRegistry public crossChainRegistry = CrossChainRegistry(0x0022d2014901F2AFBF5610dDFcd26afe2a65Ca6F);
    KeyRegistrar public keyRegistrar = KeyRegistrar(0x1C84Bb62fE7791e173014A879C706445fa893BbE);
    ECDSATableCalculator public tableCalculator = ECDSATableCalculator(0xDF4C89289A1B9De74550cE4dcFb867Ee1D2D3bF2); 

    // Storage for created operators (only using vmWallet)
    Wallet[] public operators;

    function run() public {
        // Deploy the AVS
        _deployAVS();

        // Create operators
        _createOperators();

        // Deposit for operators
        _depositOperators();

        // Register Keys and Create Reservations
        _configureKeysAndCCR();

        // Write deployment data to JSON
        _writeDeploymentOutput();
    }

    function _deployAVS() internal returns (address) {
        vm.startBroadcast();

        // Create AVS
        allocationManager.updateAVSMetadataURI(superAdmin, "test-ecdsa");

        // Create operatorSet
        IAllocationManagerTypes.CreateSetParams[] memory params = new IAllocationManagerTypes.CreateSetParams[](1);
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategy;
        params[0] = IAllocationManagerTypes.CreateSetParams({operatorSetId: operatorSetId, strategies: strategies});

        allocationManager.createOperatorSets(superAdmin, params);

        // Set AVS Registrar
        // Deploy AVS Registrar
        MockAVSRegistrar avsRegistrar = new MockAVSRegistrar();
        allocationManager.setAVSRegistrar(superAdmin, IAVSRegistrar(address(avsRegistrar)));

        vm.stopBroadcast();

        return address(avsRegistrar);
    }

    function _createOperators() internal {
        vm.startBroadcast();

        // Create 2 operators with ECDSA wallets only
        for (uint256 i = 0; i < 2; i++) {
            // Create wallet using unique name
            Wallet memory wallet =
                OperatorWalletLib.createWallet(uint256(keccak256(abi.encodePacked("ecdsa-operator", i.toString()))));
            operators.push(wallet);

            // Label the operator address
            vm.label(wallet.addr, string(abi.encodePacked("ECDSA Operator", i.toString())));

            // Send 1 ETH to operator
            payable(wallet.addr).transfer(1 ether);
        }

        vm.stopBroadcast();

        // Register each operator
        for (uint256 i = 0; i < operators.length; i++) {
            vm.startBroadcast(operators[i].privateKey);

            // Add superAdmin as pending admin for this operator
            permissionController.addPendingAdmin(operators[i].addr, superAdmin);

            // Register as operator with delegationManager
            delegationManager.registerAsOperator(
                operators[i].addr,
                0, // earningsReceiver set to operator address
                string(abi.encodePacked("ECDSA Operator", i.toString()))
            );

            vm.stopBroadcast();
        }

        // Accept admin role for each operator
        vm.startBroadcast();
        for (uint256 i = 0; i < operators.length; i++) {
            // Accept the admin role
            permissionController.acceptAdmin(operators[i].addr);
        }
        vm.stopBroadcast();
    }

    function _depositOperators() internal {
        for (uint256 i = 0; i < operators.length; i++) {
            vm.startBroadcast(operators[i].privateKey);

            // Convert 0.5 ETH to WETH
            (bool success,) = address(weth).call{value: 0.5 ether}(abi.encodeWithSignature("deposit()"));
            require(success, "WETH deposit failed");

            // Approve strategyManager to spend WETH
            weth.approve(address(strategyManager), 0.5 ether);

            // Deposit WETH into strategy
            strategyManager.depositIntoStrategy(strategy, weth, 0.5 ether);

            vm.stopBroadcast();
        }
    }

    function _configureKeysAndCCR() internal {
        // Create operator set struct
        OperatorSet memory operatorSet = OperatorSet({avs: superAdmin, id: operatorSetId});

        // Step 1: Configure key material - superAdmin configures operator set to use ECDSA
        vm.startBroadcast();
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        vm.stopBroadcast();

        // Step 2: Register operator keys
        vm.startBroadcast();
        for (uint256 i = 0; i < operators.length; i++) {
            // For ECDSA, the pubkey is just the operator's address packed into bytes
            bytes memory pubkey = abi.encodePacked(operators[i].addr);

            // Generate ECDSA signature
            bytes memory signature =
                _generateECDSASignature(operators[i].addr, operatorSet, operators[i].addr, operators[i].privateKey);

            // Register the key
            keyRegistrar.registerKey(operators[i].addr, operatorSet, pubkey, signature);
        }
        vm.stopBroadcast();

        // Step 3: Create generation reservation for cross chain registry
        vm.startBroadcast();

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory config =
            ICrossChainRegistryTypes.OperatorSetConfig({owner: superAdmin, maxStalenessPeriod: 1 days});

        // Create chain IDs array with chainID 17000
        uint256[] memory chainIDs = new uint256[](1);
        chainIDs[0] = 17_000;

        // Create generation reservation
        crossChainRegistry.createGenerationReservation(operatorSet, tableCalculator, config, chainIDs);

        vm.stopBroadcast();
    }

    function _generateECDSASignature(
        address operator,
        OperatorSet memory operatorSet,
        address keyAddress,
        uint256 privKey
    ) internal view returns (bytes memory) {
        // Get the typehash from KeyRegistrar
        bytes32 ECDSA_KEY_REGISTRATION_TYPEHASH = keyRegistrar.ECDSA_KEY_REGISTRATION_TYPEHASH();

        // Create the struct hash
        bytes32 structHash = keccak256(
            abi.encode(ECDSA_KEY_REGISTRATION_TYPEHASH, operator, operatorSet.avs, operatorSet.id, keyAddress)
        );

        // Create the domain separator message
        bytes32 domainSeparator = keyRegistrar.domainSeparator();
        bytes32 messageHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));

        // Sign the message
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, messageHash);
        return abi.encodePacked(r, s, v);
    }

    function _writeDeploymentOutput() internal {
        string memory parent_object = "parent object";

        // Serialize superAdmin
        vm.serializeAddress(parent_object, "superAdmin", superAdmin);

        // Serialize operatorSetId
        vm.serializeUint(parent_object, "operatorSetId", uint256(operatorSetId));

        // Serialize operators array
        string memory operators_array = "operators";
        for (uint256 i = 0; i < operators.length; i++) {
            string memory operator_object = string(abi.encodePacked("operator", i.toString()));
            vm.serializeAddress(operator_object, "address", operators[i].addr);
            string memory operatorOutput = vm.serializeUint(operator_object, "privateKey", operators[i].privateKey);
            vm.serializeString(operators_array, string(abi.encodePacked("[", i.toString(), "]")), operatorOutput);
        }

        // Serialize contract addresses
        string memory contracts_object = "contracts";
        vm.serializeAddress(contracts_object, "allocationManager", address(allocationManager));
        vm.serializeAddress(contracts_object, "delegationManager", address(delegationManager));
        vm.serializeAddress(contracts_object, "strategyManager", address(strategyManager));
        vm.serializeAddress(contracts_object, "strategy", address(strategy));
        vm.serializeAddress(contracts_object, "keyRegistrar", address(keyRegistrar));
        string memory contractsOutput = vm.serializeAddress(contracts_object, "ecdsaTableCalculator", address(tableCalculator));

        // Combine all outputs
        vm.serializeString(
            parent_object, "operators", vm.serializeString(operators_array, "length", operators.length.toString())
        );
        vm.serializeString(parent_object, "contracts", contractsOutput);
        string memory finalJson = vm.serializeAddress(parent_object, "avs", superAdmin);

        // Write to file
        vm.writeJson(finalJson, "script/output/devnet/multichain/avs_deployment_ecdsa.json");
    }
} 