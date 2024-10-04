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

    // Holesky - testnet
    // string deployedContractsConfig = "script/configs/holesky/eigenlayer_addresses_testnet.config.json";
    // string intialDeploymentParams = "script/configs/holesky/eigenlayer_testnet.config.json";

    // Holesky - preprod
    string deployedContractsConfig = "script/configs/holesky/eigenlayer_addresses_preprod.config.json";
    string intialDeploymentParams = "script/configs/holesky/eigenlayer_preprod.config.json";

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
}
