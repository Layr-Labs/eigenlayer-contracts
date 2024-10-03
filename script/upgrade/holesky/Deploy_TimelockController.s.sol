// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * forge script script/upgrade/holesky/Deploy_TimelockController.s.sol:Deploy_TimelockController -vvv
 * 
 */
contract Deploy_TimelockController is ExistingDeploymentParser {
    bytes32 public constant TIMELOCK_ADMIN_ROLE = keccak256("TIMELOCK_ADMIN_ROLE");
    bytes32 public constant PROPOSER_ROLE = keccak256("PROPOSER_ROLE");
    bytes32 public constant EXECUTOR_ROLE = keccak256("EXECUTOR_ROLE");
    bytes32 public constant CANCELLER_ROLE = keccak256("CANCELLER_ROLE");

    string deployedContractsConfig = "script/configs/holesky/eigenlayer_addresses_testnet.config.json";
    string intialDeploymentParams = "script/configs/holesky/eigenlayer_testnet.config.json";

    function run() external virtual {
        string memory forkUrl = vm.envString("RPC_HOLESKY");
        uint256 forkId = vm.createFork(forkUrl);
        vm.selectFork(forkId);

        // read and log the chainID
        uint256 currentChainId = block.chainid;
        emit log_named_uint("You are parsing on ChainID", currentChainId);
        require(currentChainId == 17000, "script is only for mainnet");

        _parseDeployedContracts(deployedContractsConfig);
        _parseInitialDeploymentParams(intialDeploymentParams);

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized(false);
        _verifyInitializationParams();

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        uint256 minDelay = 1;
        address[] memory proposers;
        address[] memory executors;
        TimelockController newTimelockController = new TimelockController({
            minDelay: minDelay,
            proposers: proposers,
            executors: executors
        });

        newTimelockController.grantRole(TIMELOCK_ADMIN_ROLE, executorMultisig);
        newTimelockController.grantRole(CANCELLER_ROLE, operationsMultisig);
        newTimelockController.grantRole(PROPOSER_ROLE, foundationMultisig);
        newTimelockController.grantRole(EXECUTOR_ROLE, foundationMultisig);
        newTimelockController.renounceRole(TIMELOCK_ADMIN_ROLE, msg.sender);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        require(newTimelockController.hasRole(TIMELOCK_ADMIN_ROLE, executorMultisig),
            "executorMultisig does not have TIMELOCK_ADMIN_ROLE");
        require(newTimelockController.hasRole(CANCELLER_ROLE, operationsMultisig),
            "operationsMultisig does not have CANCELLER_ROLE");
        require(newTimelockController.hasRole(PROPOSER_ROLE, foundationMultisig),
            "foundationMultisig does not have PROPOSER_ROLE");
        require(newTimelockController.hasRole(EXECUTOR_ROLE, foundationMultisig),
            "foundationMultisig does not have EXECUTOR_ROLE");
        require(!newTimelockController.hasRole(TIMELOCK_ADMIN_ROLE, msg.sender),
            "TIMELOCK_ADMIN_ROLE should have been renounced");

        emit log_named_address("address(newTimelockController)", address(newTimelockController));
    }

    // check governance configuration
    function checkDesiredGovernanceConfiguration() public {
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
        require(eigenTokenTimelockController != beigenTokenTimelockController,
            "tokens must have different timelock controllers");

        assertEq(Ownable(address(EIGEN)).owner(), address(eigenTokenTimelockController),
            "EIGEN.owner() != eigenTokenTimelockController");
        assertEq(Ownable(address(bEIGEN)).owner(), address(beigenTokenTimelockController),
            "bEIGEN.owner() != beigenTokenTimelockController");
        assertEq(eigenTokenProxyAdmin.owner(), foundationMultisig,
            "eigenTokenProxyAdmin.owner() != foundationMultisig");
        assertEq(beigenTokenProxyAdmin.owner(), foundationMultisig,
            "beigenTokenProxyAdmin.owner() != foundationMultisig");
    }
}
