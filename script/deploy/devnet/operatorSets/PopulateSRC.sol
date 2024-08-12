// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../../../src/contracts/core/StakeRootCompendium.sol";
import "../../../utils/ExistingDeploymentParser.sol";
import "forge-std/Test.sol";
import "forge-std/Script.sol";


contract PopulateSRC is Script, Test, ExistingDeploymentParser {
    string internal constant TEST_MNEMONIC = "hundred february vast fluid produce radar notice ridge armed glare panther balance";

    uint32 constant NUM_OPSETS = 50;
    uint32 constant NUM_OPERATORS_PER_OPSET = 2048;
    
    function run() public {
        _parseDeployedContracts("script/output/devnet/M2_from_scratch_deployment_data.json");

        vm.broadcast();
        IStakeRootCompendium stakeRootCompendium =  new StakeRootCompendium(delegationManager, avsDirectory);

        emit log_named_address("stakeRootCompendium", address(stakeRootCompendium));

        address[] memory allStrategies = _parseDeployedStrategies("script/output/devnet/deployed_strategies.json");

        uint32 opsetsPerChunk = 1;
        uint32 numChunks = NUM_OPSETS/opsetsPerChunk;
        IStrategy[][][] memory strategies = new IStrategy[][][](numChunks);
        for (uint256 i = 0; i < strategies.length; ++i) {
            strategies[i] = new IStrategy[][](opsetsPerChunk);
            for (uint256 j = 0; j < strategies[i].length; ++j) {
                // if we are at the end of the strategies array, break
                if (i * strategies[i].length + j == NUM_OPSETS) {
                    break;
                }

                // fill in the strategies array
                strategies[i][j] = new IStrategy[](stakeRootCompendium.MAX_NUM_STRATEGIES());
                for (uint256 k = 0; k < strategies[i][j].length; ++k) {
                    strategies[i][j][k] = IStrategy(allStrategies[i * strategies[i].length * strategies[i][j].length + j * strategies[i][j].length + k]);
                }
            }
        }

        vm.startBroadcast();

        OperatorFactory operatorFactory = new OperatorFactory(delegationManager, strategyManager);
        address[][][] memory operators = new address[][][](strategies.length);
        for (uint i = 0; i < operators.length; i++) {
            operators[i] = operatorFactory.createManyOperators(strategies[i], NUM_OPERATORS_PER_OPSET);
        }

        AVS avs = new AVS(avsDirectory, stakeRootCompendium);
        for (uint i = 0; i < strategies.length; i++) {
            avs.createOperatorSetsAndRegisterOperators(uint32(i * strategies[i].length), strategies[i], operators[i]);
        }

        vm.stopBroadcast();

        /// WRITE TO JSON

        // string memory output_path = "script/output/devnet/populate_src/";
        // FsMetadata[] memory metadata = vm.fsMetadata(output_path);
        // for (uint256 i = 0; i < entries.length; ++i) {
        //     vm.removeFile(entries[i].path);
        // }
        for (uint256 i = 0; i < operators.length; ++i) {
            for (uint256 j = 0; j < operators[i].length; ++j) {
                address[] memory strategyAddresses = new address[](strategies[i][j].length);
                for (uint256 k = 0; k < strategies[i][j].length; ++k) {
                    strategyAddresses[k] = address(strategies[i][j][k]);
                }

                string memory parent_object = "success";
                vm.serializeAddress(parent_object, "stakeRootCompendium", address(stakeRootCompendium));
                vm.serializeAddress(parent_object, "avs", address(avs));
                vm.serializeAddress(parent_object, "operators", operators[i][j]);
                vm.serializeAddress(parent_object, "strategies", strategyAddresses);
                string memory finalJson = vm.serializeString("success", "success", parent_object);
                uint32 opsetId = uint32(i * strategies[i].length + j);
                vm.writeJson(finalJson, string.concat("script/output/devnet/populate_src/opset_", string.concat(vm.toString(opsetId), ".json")));
            }
        }
    }
}

contract AVS {
    IAVSDirectory avsDirectory;
    IStakeRootCompendium stakeRootCompendium;

    // creates an operator set for each list of strategies
    constructor(IAVSDirectory _avsDirectory, IStakeRootCompendium _stakeRootCompendium) {
        avsDirectory = _avsDirectory;
        stakeRootCompendium = _stakeRootCompendium;
        avsDirectory.becomeOperatorSetAVS();
    }

    function createOperatorSetsAndRegisterOperators(uint32 startOperatorSetId, IStrategy[][] memory strategies, address[][] memory operators) external {
        uint256 numOperatorSets = strategies.length;
        // create operator sets and become an AVS
        uint32[] memory operatorSetIdsToCreate = new uint32[](numOperatorSets);
        for (uint256 i = 0; i < operatorSetIdsToCreate.length; ++i) {
            operatorSetIdsToCreate[i] = uint32(startOperatorSetId + i);
        }
        avsDirectory.createOperatorSets(operatorSetIdsToCreate);
        for (uint256 i = 0; i < operatorSetIdsToCreate.length; ++i) {
            uint32[] memory operatorSetIdsToRegisterWith = new uint32[](1);
            operatorSetIdsToRegisterWith[0] = operatorSetIdsToCreate[i];
            // register operators to operator sets
            for (uint256 j = 0; j < operators[i].length; ++j) {
                avsDirectory.registerOperatorToOperatorSets(
                    operators[i][j], 
                    operatorSetIdsToRegisterWith, 
                    ISignatureUtils.SignatureWithSaltAndExpiry({
                        signature: hex"",
                        salt: bytes32(0),
                        expiry: type(uint256).max
                    })
                );
            }

            // loop through strategies in batches of NUM_STRATS_PER_OPERATOR for each operator set
            IStakeRootCompendium.StrategyAndMultiplier[] memory strategiesAndMultipliers = new IStakeRootCompendium.StrategyAndMultiplier[](strategies[i].length);
            for (uint256 j = 0; j < strategiesAndMultipliers.length; ++j) {
                strategiesAndMultipliers[j] = IStakeRootCompendium.StrategyAndMultiplier({
                    strategy: strategies[i][j],
                    multiplier: 1 ether
                });
            }
            stakeRootCompendium.addStrategiesAndMultipliers(operatorSetIdsToCreate[i], strategiesAndMultipliers); 
        }
    }
}


contract OperatorFactory {
    IDelegationManager delegationManager;
    IStrategyManager strategyManager;
    constructor(IDelegationManager _delegationManager, IStrategyManager _strategyManager) {
        delegationManager = _delegationManager;
        strategyManager = _strategyManager;
    }

    function createManyOperators(IStrategy[][] memory strategies, uint256 numOperatorsPerOpset) public returns(address[][] memory) {
        address[][] memory operators = new address[][](strategies.length);
        for (uint256 i = 0; i < operators.length; ++i) {
            operators[i] = new address[](numOperatorsPerOpset);
            for (uint256 j = 0; j < operators[i].length; ++j) {
                operators[i][j] = address(new Operator(delegationManager));
            }
        }
        return operators;
    }
}

contract Operator is IERC1271 {
    constructor(IDelegationManager delegationManager) {
        // register as operator
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: address(this),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(
            operatorDetails,
            ""
        );
    }

    // sign everything
    function isValidSignature(bytes32, bytes memory) external pure returns (bytes4 magicValue) {
        return EIP1271SignatureUtils.EIP1271_MAGICVALUE;
    }
}