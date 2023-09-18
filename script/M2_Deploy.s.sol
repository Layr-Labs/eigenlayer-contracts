// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./utils/ExistingDeploymentParser.sol";

import "../src/contracts/pods/BeaconChainOracle.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/M2_Deploy.s.sol:Deployer_M2 --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract Deployer_M2 is ExistingDeploymentParser {
    // string public existingDeploymentInfoPath  = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));
    string public existingDeploymentInfoPath  = "script/output/M1_MOCK_deployment_data.json";
    string public deployConfigPath = "script/M2_deploy.config.json";
    string public outputPath = "output/M2_deploy.config.json";

    BeaconChainOracle public beaconChainOracle;
    address public oracleInitialOwner;
    uint256 public initialThreshold;
    address[] public initialOracleSigners;
    uint256 public currentChainId;


    function run() public {
        // get info on all the already-deployed contracts
        _parseDeployedContracts(existingDeploymentInfoPath);

        // read and log the chainID
        currentChainId= block.chainid;
        emit log_named_uint("You are deploying on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = stdJson.readUint(config_data, ".chainInfo.chainId");
        require(configChainId == currentChainId, "You are on the wrong chain for this config");

        oracleInitialOwner = executorMultisig;
        initialThreshold = stdJson.readUint(config_data, ".oracleInitialization.threshold");
        bytes memory oracleSignerListRaw = stdJson.parseRaw(config_data, ".oracleInitialization.signers");
        initialOracleSigners = abi.decode(oracleSignerListRaw, (address[]));

        require(initialThreshold <= initialOracleSigners.length, "invalid initialThreshold");
        require(initialThreshold >= 1, "invalid initialThreshold");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        beaconChainOracle = new BeaconChainOracle(oracleInitialOwner, initialThreshold, initialOracleSigners);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // additional check for correctness of deployment
        require(beaconChainOracle.owner() == executorMultisig, "beaconChainOracle owner not set correctly");


        _writeJson();
    }

    function _writeJson() internal {
        /// Parent json object to hold elements
        string memory parent_object = "parent object";
        /// JSON Keys
        string memory deployed_addresses = "addresses";
        string memory chain_info = "chainInfo";

         vm.serializeAddress(deployed_addresses, "beaconChainOracle", address(beaconChainOracle));
        string memory deployed_addresses_output = vm.serializeAddress(deployed_addresses, "initialOwner", oracleInitialOwner);

        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", currentChainId);

        // serialize all the data
        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        string memory finalJson = vm.serializeString(parent_object, chain_info, chain_info_output);

        vm.writeJson(finalJson, "script/output/M2_deployment_data.json");

        }
}



    

