// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./ExistingDeploymentParser.sol";
import "./TimelockEncoding.sol";
import "../NoDelayTimelock.sol";

/**
 * forge script script/utils/CurrentConfigCheck.s.sol:CurrentConfigCheck -vvv --sig "run(string)" $NETWORK_NAME
 * NETWORK_NAME options are currently preprod-holesky, testnet-holesky, mainnet, local
 */
contract CurrentConfigCheck is ExistingDeploymentParser, TimelockEncoding {
    string deployedContractsConfig;
    string intialDeploymentParams;
    string forkUrl;  
    string emptyString;

    address public protocolCouncilMultisig;
    TimelockController public protocolTimelockController;
    TimelockController public protocolTimelockController_BEIGEN;

    address public beigenExecutorMultisig;

    function run(string memory networkName) public virtual {
        startChainFork(networkName);

        _parseDeployedContracts(deployedContractsConfig);
        _parseInitialDeploymentParams(intialDeploymentParams);

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized(false);
        _verifyInitializationParams();

        // bytes memory data = abi.encodeWithSelector(ProxyAdmin.changeProxyAdmin.selector, address(bEIGEN), address(beigenTokenProxyAdmin));
        // bytes memory data = abi.encodeWithSignature("swapOwner(address,address,address)",
        //     0xCb8d2f9e55Bc7B1FA9d089f9aC80C583D2BDD5F7,
        //     0xcF19CE0561052a7A7Ff21156730285997B350A7D,
        //     0xFddd03C169E3FD9Ea4a9548dDC4BedC6502FE239
        // );
        // bytes memory callToExecutor = encodeForExecutor({
        //     from: communityMultisig,
        //     to: address(executorMultisig),
        //     value: 0,
        //     data: data,
        //     operation: ISafe.Operation.Call 
        // });

        // vm.prank(communityMultisig);
        // (bool success, /*bytes memory returndata*/) = executorMultisig.call(callToExecutor);
        // require(success, "call to executorMultisig failed");

        // vm.startPrank(foundationMultisig);
        // Ownable(address(bEIGEN)).transferOwnership(address(eigenTokenTimelockController));
        // Ownable(address(EIGEN)).transferOwnership(address(eigenTokenTimelockController));
        // vm.stopPrank();

        checkGovernanceConfiguration_Current();

        // simulateProtocolCouncilUpgrade(networkName);
    }

    function startChainFork(string memory networkName) public virtual {
        if (keccak256(abi.encodePacked(networkName)) == keccak256(abi.encodePacked("preprod-holesky"))) {
            deployedContractsConfig = "script/configs/holesky/eigenlayer_addresses_preprod.config.json";
            intialDeploymentParams = "script/configs/holesky/eigenlayer_preprod.config.json";
            forkUrl = vm.envString("RPC_HOLESKY");
            uint256 forkId = vm.createFork(forkUrl);
            vm.selectFork(forkId);
        } else if (keccak256(abi.encodePacked(networkName)) == keccak256(abi.encodePacked("testnet-holesky"))) {
            deployedContractsConfig = "script/configs/holesky/eigenlayer_addresses_testnet.config.json";
            intialDeploymentParams = "script/configs/holesky/eigenlayer_testnet.config.json";
            forkUrl = vm.envString("RPC_HOLESKY");
            uint256 forkId = vm.createFork(forkUrl);
            vm.selectFork(forkId);
        } else if (keccak256(abi.encodePacked(networkName)) == keccak256(abi.encodePacked("mainnet"))) {
            deployedContractsConfig = "script/configs/mainnet/mainnet-addresses.config.json";
            intialDeploymentParams = "script/configs/mainnet/mainnet-config.config.json"; 
            forkUrl = vm.envString("RPC_MAINNET");
            uint256 forkId = vm.createFork(forkUrl);
            vm.selectFork(forkId);
        } else if (keccak256(abi.encodePacked(networkName)) == keccak256(abi.encodePacked("local"))) {
            deployedContractsConfig = "script/output/devnet/local_from_scratch_deployment_data.json";
            intialDeploymentParams = "script/configs/local/deploy_from_scratch.anvil.config.json";             
        }

        require(keccak256(abi.encodePacked(deployedContractsConfig)) != keccak256(abi.encodePacked(emptyString)),
            "deployedContractsConfig cannot be unset");
        require(keccak256(abi.encodePacked(intialDeploymentParams)) != keccak256(abi.encodePacked(emptyString)),
            "intialDeploymentParams cannot be unset");

        // read and log the chainID
        uint256 currentChainId = block.chainid;
        emit log_named_uint("You are parsing on ChainID", currentChainId);
        require(currentChainId == 1 || currentChainId == 17000 || currentChainId == 31337,
            "script is only for mainnet or holesky or local environment");
    }

    // check governance configuration
    function checkGovernanceConfiguration_Current() public {
        assertEq(eigenLayerProxyAdmin.owner(), executorMultisig,
            "eigenLayerProxyAdmin.owner() != executorMultisig");
        assertEq(delegationManager.owner(), executorMultisig,
            "delegationManager.owner() != executorMultisig");
        assertEq(strategyManager.owner(), executorMultisig,
            "strategyManager.owner() != executorMultisig");
        assertEq(strategyManager.strategyWhitelister(), address(strategyFactory),
            "strategyManager.strategyWhitelister() != address(strategyFactory)");
        assertEq(strategyFactory.owner(), operationsMultisig,
            "strategyFactory.owner() != operationsMultisig");
        assertEq(avsDirectory.owner(), executorMultisig,
            "avsDirectory.owner() != executorMultisig");
        assertEq(rewardsCoordinator.owner(), operationsMultisig,
            "rewardsCoordinator.owner() != operationsMultisig");
        assertEq(eigenLayerPauserReg.unpauser(), executorMultisig,
            "eigenLayerPauserReg.unpauser() != operationsMultisig");
        require(eigenLayerPauserReg.isPauser(operationsMultisig),
            "operationsMultisig does not have pausing permissions");
        require(eigenLayerPauserReg.isPauser(executorMultisig),
            "executorMultisig does not have pausing permissions");
        require(eigenLayerPauserReg.isPauser(pauserMultisig),
            "pauserMultisig does not have pausing permissions");

        address timelockAdmin = NoDelayTimelock(payable(timelock)).admin();
        assertEq(timelockAdmin, operationsMultisig,
            "timelockAdmin != operationsMultisig");

        (bool success, bytes memory returndata) = executorMultisig.staticcall(abi.encodeWithSignature("getOwners()"));
        require(success, "call to executorMultisig.getOwners() failed");
        address[] memory executorMultisigOwners = abi.decode(returndata, (address[]));
        require(executorMultisigOwners.length == 2,
            "executorMultisig owners wrong length");
        bool timelockInOwners;
        bool communityMultisigInOwners;
        for (uint256 i = 0; i < 2; ++i) {
            if (executorMultisigOwners[i] == timelock) {
                timelockInOwners = true;
            }
            if (executorMultisigOwners[i] == communityMultisig) {
                communityMultisigInOwners = true;
            }
        }
        require(timelockInOwners, "timelock not in executorMultisig owners");
        require(communityMultisigInOwners, "communityMultisig not in executorMultisig owners");

        require(eigenTokenProxyAdmin != beigenTokenProxyAdmin,
            "tokens must have different proxy admins to allow different timelock controllers");
        require(eigenTokenTimelockController != beigenTokenTimelockController,
            "tokens must have different timelock controllers");

        // note that proxy admin owners are different but _token_ owners are the same
        assertEq(Ownable(address(EIGEN)).owner(), address(eigenTokenTimelockController),
            "EIGEN.owner() != eigenTokenTimelockController");
        assertEq(Ownable(address(bEIGEN)).owner(), address(eigenTokenTimelockController),
            "bEIGEN.owner() != eigenTokenTimelockController");
        assertEq(eigenTokenProxyAdmin.owner(), address(eigenTokenTimelockController),
            "eigenTokenProxyAdmin.owner() != eigenTokenTimelockController");
        assertEq(beigenTokenProxyAdmin.owner(), address(beigenTokenTimelockController),
            "beigenTokenProxyAdmin.owner() != beigenTokenTimelockController");

        assertEq(eigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(EIGEN)))),
            address(eigenTokenProxyAdmin),
            "eigenTokenProxyAdmin is not actually the admin of the EIGEN token");
        assertEq(beigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(bEIGEN)))),
            address(beigenTokenProxyAdmin),
            "beigenTokenProxyAdmin is not actually the admin of the bEIGEN token");

        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.PROPOSER_ROLE(), foundationMultisig),
            "foundationMultisig does not have PROPOSER_ROLE on eigenTokenTimelockController");
        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.EXECUTOR_ROLE(), foundationMultisig),
            "foundationMultisig does not have EXECUTOR_ROLE on eigenTokenTimelockController");
        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on eigenTokenTimelockController");
        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.TIMELOCK_ADMIN_ROLE(), executorMultisig),
            "executorMultisig does not have TIMELOCK_ADMIN_ROLE on eigenTokenTimelockController");
        require(eigenTokenTimelockController.hasRole(eigenTokenTimelockController.TIMELOCK_ADMIN_ROLE(), address(eigenTokenTimelockController)),
            "eigenTokenTimelockController does not have TIMELOCK_ADMIN_ROLE on itself");

        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.PROPOSER_ROLE(), foundationMultisig),
            "foundationMultisig does not have PROPOSER_ROLE on beigenTokenTimelockController");
        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.EXECUTOR_ROLE(), foundationMultisig),
            "foundationMultisig does not have EXECUTOR_ROLE on beigenTokenTimelockController");
        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on beigenTokenTimelockController");
        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.TIMELOCK_ADMIN_ROLE(), executorMultisig),
            "executorMultisig does not have TIMELOCK_ADMIN_ROLE on beigenTokenTimelockController");
        require(beigenTokenTimelockController.hasRole(beigenTokenTimelockController.TIMELOCK_ADMIN_ROLE(), address(beigenTokenTimelockController)),
            "beigenTokenTimelockController does not have TIMELOCK_ADMIN_ROLE on itself");
    }

    function checkGovernanceConfiguration_WithProtocolCouncil() public {
        assertEq(eigenLayerProxyAdmin.owner(), executorMultisig,
            "eigenLayerProxyAdmin.owner() != executorMultisig");
        assertEq(delegationManager.owner(), executorMultisig,
            "delegationManager.owner() != executorMultisig");
        assertEq(strategyManager.owner(), executorMultisig,
            "strategyManager.owner() != executorMultisig");
        assertEq(strategyManager.strategyWhitelister(), address(strategyFactory),
            "strategyManager.strategyWhitelister() != address(strategyFactory)");
        assertEq(strategyFactory.owner(), operationsMultisig,
            "strategyFactory.owner() != operationsMultisig");
        assertEq(avsDirectory.owner(), executorMultisig,
            "avsDirectory.owner() != executorMultisig");
        assertEq(rewardsCoordinator.owner(), operationsMultisig,
            "rewardsCoordinator.owner() != operationsMultisig");
        assertEq(eigenLayerPauserReg.unpauser(), executorMultisig,
            "eigenLayerPauserReg.unpauser() != operationsMultisig");
        require(eigenLayerPauserReg.isPauser(operationsMultisig),
            "operationsMultisig does not have pausing permissions");
        require(eigenLayerPauserReg.isPauser(executorMultisig),
            "executorMultisig does not have pausing permissions");
        require(eigenLayerPauserReg.isPauser(pauserMultisig),
            "pauserMultisig does not have pausing permissions");

        // TODO: delete this? it should no longer matter going forwards
        address timelockAdmin = NoDelayTimelock(payable(timelock)).admin();
        assertEq(timelockAdmin, operationsMultisig,
            "timelockAdmin != operationsMultisig");

        require(eigenTokenProxyAdmin != beigenTokenProxyAdmin,
            "tokens must have different proxy admins to allow different timelock controllers");
        require(protocolTimelockController != protocolTimelockController_BEIGEN,
            "tokens must have different timelock controllers");

        // note that proxy admin owners are different but _token_ owners per se are the same
        assertEq(Ownable(address(EIGEN)).owner(), address(executorMultisig),
            "EIGEN.owner() != executorMultisig");
        assertEq(Ownable(address(bEIGEN)).owner(), address(executorMultisig),
            "bEIGEN.owner() != executorMultisig");
        assertEq(eigenLayerProxyAdmin.owner(), address(executorMultisig),
            "eigenLayerProxyAdmin.owner() != executorMultisig");
        assertEq(beigenTokenProxyAdmin.owner(), address(protocolTimelockController_BEIGEN),
            "beigenTokenProxyAdmin.owner() != protocolTimelockController_BEIGEN");

        assertEq(eigenLayerProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(EIGEN)))),
            address(eigenLayerProxyAdmin),
            "eigenLayerProxyAdmin is not actually the admin of the EIGEN token");
        assertEq(beigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(bEIGEN)))),
            address(beigenTokenProxyAdmin),
            "beigenTokenProxyAdmin is not actually the admin of the bEIGEN token");

        // check that community multisig and protocol timelock are the owners of the executorMultisig
        checkExecutorMultisigOwnership(executorMultisig, address(protocolTimelockController));
        // check that community multisig and bEIGEN protocol timelock are the owners of the beigenExecutorMultisig
        checkExecutorMultisigOwnership(beigenExecutorMultisig, address(protocolTimelockController_BEIGEN));

        checkTimelockControllerConfig(protocolTimelockController);
        checkTimelockControllerConfig(protocolTimelockController_BEIGEN);
    }

    function checkExecutorMultisigOwnership(address _executorMultisig, address timelockControllerAddress) public view {
        (bool success, bytes memory returndata) = _executorMultisig.staticcall(abi.encodeWithSignature("getOwners()"));
        require(success, "call to _executorMultisig.getOwners() failed");
        address[] memory _executorMultisigOwners = abi.decode(returndata, (address[]));
        require(_executorMultisigOwners.length == 2,
            "executorMultisig owners wrong length");
        bool timelockControllerInOwners;
        bool communityMultisigInOwners;
        for (uint256 i = 0; i < 2; ++i) {
            if (_executorMultisigOwners[i] == address(timelockControllerAddress)) {
                timelockControllerInOwners = true;
            }
            if (_executorMultisigOwners[i] == communityMultisig) {
                communityMultisigInOwners = true;
            }
        }
        require(timelockControllerInOwners, "timelockControllerAddress not in _executorMultisig owners");
        require(communityMultisigInOwners, "communityMultisig not in _executorMultisig owners");

    }

    function checkTimelockControllerConfig(TimelockController timelockController) public view {
        // check for proposer + executor rights on Protocol Council multisig
        require(timelockController.hasRole(timelockController.PROPOSER_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have PROPOSER_ROLE on timelockController");
        require(timelockController.hasRole(timelockController.EXECUTOR_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have EXECUTOR_ROLE on timelockController");

        // check for proposer + canceller rights on ops multisig
        require(timelockController.hasRole(timelockController.PROPOSER_ROLE(), operationsMultisig),
            "operationsMultisig does not have PROPOSER_ROLE on timelockController");
        require(timelockController.hasRole(timelockController.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on timelockController");

        // check that community multisig has admin rights
        require(timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), communityMultisig),
            "communityMultisig does not have TIMELOCK_ADMIN_ROLE on timelockController");

        // check for self-administration
        require(timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), address(timelockController)),
            "timelockController does not have TIMELOCK_ADMIN_ROLE on itself");

        // check that deployer has no rights
        require(!timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), msg.sender),
            "deployer erroenously retains TIMELOCK_ADMIN_ROLE on timelockController");
        require(!timelockController.hasRole(timelockController.PROPOSER_ROLE(), msg.sender),
            "deployer erroenously retains PROPOSER_ROLE on timelockController");
        require(!timelockController.hasRole(timelockController.EXECUTOR_ROLE(), msg.sender),
            "deployer erroenously retains EXECUTOR_ROLE on timelockController");
        require(!timelockController.hasRole(timelockController.CANCELLER_ROLE(), msg.sender),
            "deployer erroenously retains CANCELLER_ROLE on timelockController");
    }

    // forge script script/utils/CurrentConfigCheck.s.sol:CurrentConfigCheck -vvv --sig "simulateProtocolCouncilUpgrade(string)" $NETWORK_NAME
    function simulateProtocolCouncilUpgrade(string memory networkName) public virtual {
        run(networkName);

        // deployment steps
        deployTimelockControllers();
        deployBeigenExecutorMultisig();
        configureTimelockController(protocolTimelockController);
        configureTimelockController(protocolTimelockController_BEIGEN);

        // simulate transferring powers
        simulateLegacyTimelockActions();
        simulateEigenTokenTimelockControllerActions();
        simulateBEIGENTokenTimelockControllerActions();

        // check correctness after deployment + simulation
        checkGovernanceConfiguration_WithProtocolCouncil();
    }

    function deployTimelockControllers() public {
        // set up initially with sender as a proposer & executor, to be renounced prior to finalizing deployment
        address[] memory proposers = new address[](1);
        proposers[0] = msg.sender;

        address[] memory executors = new address[](1);
        executors[0] = msg.sender;

        vm.startBroadcast();
        protocolTimelockController = 
            new TimelockController(
                0, // no delay for setup
                proposers, 
                executors
            );
        protocolTimelockController_BEIGEN = 
            new TimelockController(
                0, // no delay for setup
                proposers, 
                executors
            );
        vm.stopBroadcast();
    }

    function deployBeigenExecutorMultisig() public {
        require(address(beigenTokenTimelockController) != address(0),
            "must deploy beigenTokenTimelockController before beigenExecutorMultisig");
        // TODO: solution for local network?
        // TODO: verify this also works for Holesky
        address safeFactory = 0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2;
        address safeSingleton = 0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552;
        address safeFallbackHandler = 0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99;
        address[] memory owners = new address[](2);
        bytes memory emptyData;
        owners[0] = address(protocolTimelockController_BEIGEN);
        owners[1] = communityMultisig;
        bytes memory initializerData = abi.encodeWithSignature(
            "setup(address[],uint256,address,bytes,address,address,uint256,address)",
            owners,
            1, /* threshold */
            address(0), /* to (used in setupModules) */
            emptyData, /* data (used in setupModules) */
            safeFallbackHandler,
            address(0), /* paymentToken */
            0, /* payment */
            payable(address(0)) /* paymentReceiver */
        );
        uint256 salt = 0;

        bytes memory calldataToFactory = abi.encodeWithSignature(
            "createProxyWithNonce(address,bytes,uint256)",
            safeSingleton,
            initializerData,
            salt
        );

        (bool success, bytes memory returndata) = safeFactory.call(calldataToFactory);
        require(success, "multisig deployment failed");
        beigenExecutorMultisig = abi.decode(returndata, (address));
        require(beigenExecutorMultisig != address(0), "something wrong in multisig deployment, zero address returned");
    }

    function configureTimelockController(TimelockController timelockController) public {
        vm.startBroadcast();
        uint256 tx_array_length = 10;
        address[] memory targets = new address[](tx_array_length);
        for (uint256 i = 0; i < targets.length; ++i) {
            targets[i] = address(timelockController);
        }

        uint256[] memory values = new uint256[](tx_array_length);

        bytes[] memory payloads = new bytes[](tx_array_length);
        // 1. remove sender as canceller
        payloads[0] = abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.CANCELLER_ROLE(), msg.sender);
        // 2. remove sender as executor
        payloads[1] = abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.EXECUTOR_ROLE(), msg.sender);
        // 3. remove sender as proposer
        payloads[2] = abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.PROPOSER_ROLE(), msg.sender);
        // 4. remove sender as admin
        payloads[3] = abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.TIMELOCK_ADMIN_ROLE(), msg.sender);

        // 5. add operationsMultisig as canceller
        payloads[4] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.CANCELLER_ROLE(), operationsMultisig);
        // 6. add operationsMultisig as proposer
        payloads[5] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.PROPOSER_ROLE(), operationsMultisig);

        // 7. add protocolCouncilMultisig as proposer
        payloads[6] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.PROPOSER_ROLE(), protocolCouncilMultisig);
        // 8. add protocolCouncilMultisig as executor
        payloads[7] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.EXECUTOR_ROLE(), protocolCouncilMultisig);

        // 9. add communityMultisig as admin
        payloads[8] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.TIMELOCK_ADMIN_ROLE(), communityMultisig);

        // TODO: get appropriate value for chain instead of hardcoding
        uint256 delayToSet;
        if (timelockController == protocolTimelockController) {
            delayToSet = 10 days;
        } else if (timelockController == protocolTimelockController_BEIGEN) {
            delayToSet = 14 days;
        }
        require(delayToSet != 0, "delay not calculated");
        // 10. set min delay to appropriate length
        payloads[9] = abi.encodeWithSelector(timelockController.updateDelay.selector, delayToSet);

        // schedule the batch
        timelockController.scheduleBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0), // no salt 
            0 // 0 enforced delay
        );

        // execute the batch
        timelockController.executeBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0) // no salt
        );
        vm.stopBroadcast();
    }

    function simulateLegacyTimelockActions() public {
        // TODO
        // give proxy admin ownership to timelock controller
        // eigenTokenProxyAdmin.transferOwnership(address(timelockController));

        // swapOwner(address previousOwner, address oldOwner, address newOwner)
        // TODO: figure out if this is the correct input for all chains or the communityMultisig is correct on some
        address previousOwner = address(1);
        // address previousOwner = communityMultisig;
        bytes memory data_swapTimelockToProtocolTimelock =
            abi.encodeWithSignature("swapOwner(address,address,address)", previousOwner, address(timelock), address(protocolTimelockController));
        bytes memory callToExecutor = encodeForExecutor({
            from: timelock,
            to: address(executorMultisig),
            value: 0,
            data: data_swapTimelockToProtocolTimelock,
            operation: ISafe.Operation.Call 
        });
        // TODO: get appropriate value for chain instead of hardcoding
        // uint256 timelockEta = block.timestamp + timelock.delay();
        uint256 timelockEta = block.timestamp + 10 days;
        (bytes memory calldata_to_timelock_queuing_action, bytes memory calldata_to_timelock_executing_action) =
            encodeForTimelock({
                to: address(executorMultisig),
                value: 0,
                data: callToExecutor,
                timelockEta: timelockEta
            });
        vm.startPrank(operationsMultisig);

        (bool success, /*bytes memory returndata*/) = timelock.call(calldata_to_timelock_queuing_action);
        require(success, "call to timelock queuing action 1 failed");

        vm.warp(timelockEta);
        (success, /*bytes memory returndata*/) = timelock.call(calldata_to_timelock_executing_action);
        require(success, "call to timelock executing action 1 failed");

        vm.stopPrank();
    }

    function simulateEigenTokenTimelockControllerActions() public {
        vm.startPrank(foundationMultisig);

        uint256 tx_array_length = 3;
        address[] memory targets = new address[](tx_array_length);
        uint256[] memory values = new uint256[](tx_array_length);
        bytes[] memory payloads = new bytes[](tx_array_length);
        // 1. transfer upgrade rights over EIGEN token from eigenTokenProxyAdmin to eigenLayerProxyAdmin
        targets[0] = address(eigenTokenProxyAdmin);
        payloads[0] = abi.encodeWithSelector(ProxyAdmin.changeProxyAdmin.selector, address(EIGEN), address(eigenLayerProxyAdmin));
        // 2. transfer ownership of EIGEN token to executorMultisig
        targets[1] = address(EIGEN);
        payloads[1] = abi.encodeWithSelector(Ownable.transferOwnership.selector, executorMultisig);
        // 3. transfer ownership of bEIGEN token to executorMultisig
        targets[2] = address(bEIGEN);
        payloads[2] = abi.encodeWithSelector(Ownable.transferOwnership.selector, executorMultisig);

        // schedule the batch
        uint256 minDelay = eigenTokenTimelockController.getMinDelay();
        eigenTokenTimelockController.scheduleBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0), // no salt 
            minDelay
        );

        vm.warp(block.timestamp + minDelay);

        // execute the batch
        eigenTokenTimelockController.executeBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0) // no salt
        );

        vm.stopPrank();
    }

    function simulateBEIGENTokenTimelockControllerActions() public {
        vm.startPrank(foundationMultisig);

        uint256 tx_array_length = 1;
        address[] memory targets = new address[](tx_array_length);
        uint256[] memory values = new uint256[](tx_array_length);
        bytes[] memory payloads = new bytes[](tx_array_length);
        // 1. transfer ownership rights over beigenTokenProxyAdmin to protocolTimelockController_BEIGEN
        targets[0] = address(beigenTokenProxyAdmin);
        payloads[0] = abi.encodeWithSelector(Ownable.transferOwnership.selector, address(protocolTimelockController_BEIGEN));

        // schedule the batch
        uint256 minDelay = beigenTokenTimelockController.getMinDelay();
        beigenTokenTimelockController.scheduleBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0), // no salt 
            minDelay
        );

        vm.warp(block.timestamp + minDelay);

        // execute the batch
        beigenTokenTimelockController.executeBatch(
            targets, 
            values, 
            payloads, 
            bytes32(0), // no predecessor needed
            bytes32(0) // no salt
        );

        vm.stopPrank();
    }

    function checkGovernanceConfiguration_WithProtocolCouncilMultisig() public {
        assertEq(eigenLayerProxyAdmin.owner(), protocolCouncilMultisig,
            "eigenLayerProxyAdmin.owner() != protocolCouncilMultisig");
        assertEq(delegationManager.owner(), executorMultisig,
            "delegationManager.owner() != executorMultisig");
        assertEq(strategyManager.owner(), executorMultisig,
            "strategyManager.owner() != executorMultisig");
        assertEq(strategyManager.strategyWhitelister(), address(strategyFactory),
            "strategyManager.strategyWhitelister() != address(strategyFactory)");
        assertEq(strategyFactory.owner(), operationsMultisig,
            "strategyFactory.owner() != operationsMultisig");
        assertEq(avsDirectory.owner(), executorMultisig,
            "avsDirectory.owner() != executorMultisig");
        assertEq(rewardsCoordinator.owner(), operationsMultisig,
            "rewardsCoordinator.owner() != operationsMultisig");
        assertEq(eigenLayerPauserReg.unpauser(), executorMultisig,
            "eigenLayerPauserReg.unpauser() != operationsMultisig");
        require(eigenLayerPauserReg.isPauser(operationsMultisig),
            "operationsMultisig does not have pausing permissions");
        require(eigenLayerPauserReg.isPauser(executorMultisig),
            "executorMultisig does not have pausing permissions");
        require(eigenLayerPauserReg.isPauser(pauserMultisig),
            "pauserMultisig does not have pausing permissions");

        (bool success, bytes memory returndata) = timelock.staticcall(abi.encodeWithSignature("admin()"));
        require(success, "call to timelock.admin() failed");
        address timelockAdmin = abi.decode(returndata, (address));
        assertEq(timelockAdmin, operationsMultisig,
            "timelockAdmin != operationsMultisig");

        (success, returndata) = executorMultisig.staticcall(abi.encodeWithSignature("getOwners()"));
        require(success, "call to executorMultisig.getOwners() failed");
        address[] memory executorMultisigOwners = abi.decode(returndata, (address[]));
        require(executorMultisigOwners.length == 2,
            "executorMultisig owners wrong length");
        bool timelockInOwners;
        bool communityMultisigInOwners;
        for (uint256 i = 0; i < 2; ++i) {
            if (executorMultisigOwners[i] == timelock) {
                timelockInOwners = true;
            }
            if (executorMultisigOwners[i] == communityMultisig) {
                communityMultisigInOwners = true;
            }
        }
        require(timelockInOwners, "timelock not in executorMultisig owners");
        require(communityMultisigInOwners, "communityMultisig not in executorMultisig owners");

        require(eigenTokenProxyAdmin != beigenTokenProxyAdmin,
            "tokens must have different proxy admins to allow different timelock controllers");
        require(protocolTimelockController != protocolTimelockController_BEIGEN,
            "tokens must have different timelock controllers");

        // note that proxy admin owners are different but _token_ owners are the same
        assertEq(Ownable(address(EIGEN)).owner(), address(protocolTimelockController),
            "EIGEN.owner() != protocolTimelockController");
        assertEq(Ownable(address(bEIGEN)).owner(), address(protocolTimelockController),
            "bEIGEN.owner() != protocolTimelockController");
        assertEq(eigenTokenProxyAdmin.owner(), address(protocolTimelockController),
            "eigenTokenProxyAdmin.owner() != protocolTimelockController");
        assertEq(beigenTokenProxyAdmin.owner(), address(protocolTimelockController_BEIGEN),
            "beigenTokenProxyAdmin.owner() != protocolTimelockController_BEIGEN");

        assertEq(eigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(EIGEN)))),
            address(eigenTokenProxyAdmin),
            "eigenTokenProxyAdmin is not actually the admin of the EIGEN token");
        assertEq(beigenTokenProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(bEIGEN)))),
            address(beigenTokenProxyAdmin),
            "beigenTokenProxyAdmin is not actually the admin of the bEIGEN token");

        require(protocolTimelockController.hasRole(protocolTimelockController.PROPOSER_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have PROPOSER_ROLE on protocolTimelockController");
        require(protocolTimelockController.hasRole(protocolTimelockController.EXECUTOR_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have EXECUTOR_ROLE on protocolTimelockController");
        require(protocolTimelockController.hasRole(protocolTimelockController.PROPOSER_ROLE(), operationsMultisig),
            "operationsMultisig does not have PROPOSER_ROLE on protocolTimelockController");
        require(protocolTimelockController.hasRole(protocolTimelockController.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on protocolTimelockController");

        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.PROPOSER_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have PROPOSER_ROLE on protocolTimelockController_BEIGEN");
        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.EXECUTOR_ROLE(), protocolCouncilMultisig),
            "protocolCouncilMultisig does not have EXECUTOR_ROLE on protocolTimelockController_BEIGEN");
        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.PROPOSER_ROLE(), operationsMultisig),
            "operationsMultisig does not have PROPOSER_ROLE on protocolTimelockController_BEIGEN");
        require(protocolTimelockController_BEIGEN.hasRole(protocolTimelockController_BEIGEN.CANCELLER_ROLE(), operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE on protocolTimelockController_BEIGEN");
    }
}
