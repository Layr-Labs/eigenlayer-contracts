// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../../../src/contracts/core/StakeRootCompendium.sol";
import "../../../utils/ExistingDeploymentParser.sol";
import "forge-std/Test.sol";
import "forge-std/Script.sol";


contract PopulateSRC is Script, Test, ExistingDeploymentParser {
    string internal constant TEST_MNEMONIC = "hundred february vast fluid produce radar notice ridge armed glare panther balance";

    uint32 constant NUM_OPSETS = 10;
    uint32 constant NUM_OPERATORS_PER_OPSET = 2048;
    uint256 constant AMOUNT_TOKEN = 1000;
    
    function run() public {
        _parseDeployedContracts("script/output/devnet/M2_from_scratch_deployment_data.json");

        vm.broadcast();
        IStakeRootCompendium stakeRootCompendium =  new StakeRootCompendium({
            _delegationManager: delegationManager,
            _avsDirectory: avsDirectory,
            _challengePeriod: 12 seconds,
            _minTimeSinceLastClaim: 1 days
        });

        emit log_named_address("stakeRootCompendium", address(stakeRootCompendium));

        address[] memory allStrategies = _parseDeployedStrategies("script/output/devnet/deployed_strategies.json");

        // list of strategies for each operatorSet
        IStrategy[][] memory strategies = new IStrategy[][](NUM_OPSETS);
        for (uint256 i = 0; i < strategies.length; ++i) {
            strategies[i] = new IStrategy[](stakeRootCompendium.MAX_NUM_STRATEGIES());
            for (uint256 j = 0; j < strategies[i].length; ++j) {
                // fill in the strategies array
                strategies[i][j] = IStrategy(allStrategies[i * strategies[i].length + j]);
            }
        }

        vm.broadcast();
        OperatorFactory operatorFactory = new OperatorFactory(delegationManager, strategyManager);
        address[][] memory operators = new address[][](strategies.length);
        for (uint i = 0; i < operators.length; i++) {
            // todo: send operators[i].length*1 ether of strategy token to operatorfactory
            // //transfer enough tokens for every operator in the operator set for all the strategies in the operator set
           
            for (uint j = 0; j < strategies[i].length; j++) {
                 vm.startBroadcast();
                IERC20 token = strategies[i][j].underlyingToken();
                token.approve(msg.sender, type(uint256).max);
                token.transferFrom(msg.sender, address(operatorFactory), NUM_OPERATORS_PER_OPSET*AMOUNT_TOKEN);
                vm.stopBroadcast();
            }


            operators[i] = operatorFactory.createManyOperators(strategies[i], NUM_OPERATORS_PER_OPSET);
        }

        vm.broadcast();
        AVS avs = new AVS(avsDirectory, stakeRootCompendium);
        for (uint i = 0; i < strategies.length; i++) {
            vm.startBroadcast();
            avs.createOperatorSetAndRegisterOperators(uint32(i), strategies[i], operators[i]);
            vm.stopBroadcast();
        }

        /// WRITE TO JSON

        // string memory output_path = "script/output/devnet/populate_src/";
        // FsMetadata[] memory metadata = vm.fsMetadata(output_path);
        // for (uint256 i = 0; i < entries.length; ++i) {
        //     vm.removeFile(entries[i].path);
        // }
        for (uint256 i = 0; i < operators.length; ++i) {
            address[] memory strategyAddresses = new address[](strategies[i].length);
            for (uint256 j = 0; j < strategies[i].length; ++j) {
                strategyAddresses[j] = address(strategies[i][j]);
            }

            string memory parent_object = "success";
            vm.serializeAddress(parent_object, "stakeRootCompendium", address(stakeRootCompendium));
            vm.serializeAddress(parent_object, "avs", address(avs));
            vm.serializeAddress(parent_object, "operators", operators[i]);
            vm.serializeAddress(parent_object, "strategies", strategyAddresses);
            string memory finalJson = vm.serializeString("success", "success", parent_object);
            vm.writeJson(finalJson, string.concat("script/output/devnet/populate_src/opset_", string.concat(vm.toString(i), ".json")));
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

    function createOperatorSetAndRegisterOperators(uint32 operatorSetId, IStrategy[] memory strategies, address[] memory operators) external {
        // create operator sets and become an AVS
        uint32[] memory operatorSetIdsToCreate = new uint32[](1);
        operatorSetIdsToCreate[0] = operatorSetId;
        avsDirectory.createOperatorSets(operatorSetIdsToCreate);

        // register operators to operator sets
        for (uint256 i = 0; i < operators.length; ++i) {
            avsDirectory.registerOperatorToOperatorSets(
                operators[i], 
                operatorSetIdsToCreate, 
                ISignatureUtils.SignatureWithSaltAndExpiry({
                    signature: hex"",
                    salt: bytes32(0),
                    expiry: type(uint256).max
                })
            );
        }

        // loop through strategies in batches of NUM_STRATS_PER_OPERATOR for each operator set
        IStakeRootCompendium.StrategyAndMultiplier[] memory strategiesAndMultipliers = new IStakeRootCompendium.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategiesAndMultipliers.length; ++i) {
            strategiesAndMultipliers[i] = IStakeRootCompendium.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: 1 ether
            });
        }
        stakeRootCompendium.addStrategiesAndMultipliers(operatorSetId, strategiesAndMultipliers); 
    }
}


contract OperatorFactory is Test {
    IDelegationManager delegationManager;
    IStrategyManager strategyManager;
    constructor(IDelegationManager _delegationManager, IStrategyManager _strategyManager) {
        delegationManager = _delegationManager;
        strategyManager = _strategyManager;
    }

    uint256 constant AMOUNT_TOKEN = 1000;

    function createManyOperators(IStrategy[] memory strategies, uint256 numOperatorsPerOpset) public returns(address[] memory) {
        address[] memory operators = new address[](numOperatorsPerOpset);
        for (uint256 i = 0; i < operators.length; ++i) {
            operators[i] = address(new Operator(delegationManager));
            for (uint256 j = 0; j < strategies.length; ++j) {
                vm.startBroadcast();
                // todo: deposit on behalf of operator for each strategy
                IERC20 token = strategies[j].underlyingToken();
                token.approve(address(this), type(uint256).max);
                token.approve(address(strategyManager), type(uint256).max);
                vm.stopBroadcast();

                vm.startBroadcast();
                strategyManager.depositIntoStrategyWithSignature(strategies[j], token, AMOUNT_TOKEN , operators[i], type(uint256).max, hex"");
                vm.stopBroadcast();
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