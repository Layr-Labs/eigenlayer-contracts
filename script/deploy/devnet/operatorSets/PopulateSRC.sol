// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../../../src/contracts/core/StakeRootCompendium.sol";
import "../../../utils/ExistingDeploymentParser.sol";
import "forge-std/Test.sol";
import "forge-std/Script.sol";


contract PopulateSRC is Script, Test, ExistingDeploymentParser {
    string internal constant TEST_MNEMONIC = "hundred february vast fluid produce radar notice ridge armed glare panther balance";

    uint32 constant NUM_OPSETS = 1;
    uint32 constant NUM_OPERATORS_PER_OPSET = 2048;
    uint32 constant NUM_STRATS_PER_OPSET = 20;
    uint256 constant TOKEN_AMOUNT_PER_OPERATOR = 1 ether;

    
    function run() public {
        _parseDeployedContracts("script/output/devnet/M2_from_scratch_deployment_data.json");

        vm.startBroadcast();
        IStakeRootCompendium stakeRootCompendiumImplementation =  new StakeRootCompendium({
            _delegationManager: delegationManager,
            _avsDirectory: avsDirectory,
            _proofInterval: 1 hours,
            _blacklistWindow: 12 seconds
        });
        StakeRootCompendium stakeRootCompendium = StakeRootCompendium(payable(new TransparentUpgradeableProxy(
            address(stakeRootCompendiumImplementation),
            address(msg.sender),
            ""
        )));
        vm.stopBroadcast();

        emit log_named_address("stakeRootCompendium", address(stakeRootCompendium));

        address[] memory allStrategies = _parseDeployedStrategies("script/output/devnet/deployed_strategies.json");

        // list of strategies for each operatorSet
        IStrategy[][] memory strategies = new IStrategy[][](NUM_OPSETS);
        for (uint256 i = 0; i < strategies.length; ++i) {
            strategies[i] = new IStrategy[](NUM_STRATS_PER_OPSET);
            for (uint256 j = 0; j < strategies[i].length; ++j) {
                // fill in the strategies array
                strategies[i][j] = IStrategy(allStrategies[i * strategies[i].length + j]);
            }
        }

        vm.broadcast();
        OperatorFactory operatorFactory = new OperatorFactory(delegationManager, strategyManager, avsDirectory);
        address[][] memory operators = new address[][](strategies.length);
        for (uint i = 0; i < operators.length; i++) {
            // todo: send operators[i].length*1 ether of strategy token to operatorfactory
            // //transfer enough tokens for every operator in the operator set for all the strategies in the operator set
           
            for (uint j = 0; j < strategies[i].length; j++) {
                IERC20 token = strategies[i][j].underlyingToken();
                vm.startBroadcast();
                token.approve(msg.sender, type(uint256).max);
                token.transfer(address(operatorFactory), NUM_OPERATORS_PER_OPSET*TOKEN_AMOUNT_PER_OPERATOR);
                vm.stopBroadcast();
            }
            
            vm.startBroadcast();
            operators[i] = operatorFactory.createManyOperators(strategies[i], NUM_OPERATORS_PER_OPSET);
            for (uint j = 0; j < strategies[i].length; j++) {
                operatorFactory.depositForOperators(strategies[i][j], operators[i], TOKEN_AMOUNT_PER_OPERATOR);
            }
            vm.stopBroadcast();
        }

        uint64 magnitudeForOperators = 0.1 ether;
        vm.startBroadcast();
        AVS avs = new AVS(avsDirectory, stakeRootCompendium);
        for (uint i = 0; i < strategies.length; i++) {
            avs.createOperatorSetAndRegisterOperators(uint32(i), strategies[i], operators[i]);
            IAVSDirectory.OperatorSet memory operatorSet = IAVSDirectory.OperatorSet({
                avs: address(avs),
                operatorSetId: uint32(i)
            });
            for (uint j = 0; j < strategies[i].length; j++) {
                operatorFactory.allocateForOperators(strategies[i][j], operatorSet, operators[i], magnitudeForOperators);
            }
        }
        vm.stopBroadcast();

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
    IAVSDirectory avsDirectory;
    constructor(IDelegationManager _delegationManager, IStrategyManager _strategyManager, IAVSDirectory _avsDirectory) {
        delegationManager = _delegationManager;
        strategyManager = _strategyManager;
        avsDirectory = _avsDirectory;
    }

    uint256 constant AMOUNT_TOKEN = 1000;

    function createManyOperators(IStrategy[] memory strategies, uint256 numOperatorsPerOpset) public returns(address[] memory) {
        IERC20[] memory tokens = new IERC20[](strategies.length);
        // approve all transfers
        for (uint256 i = 0; i < strategies.length; ++i) {
            tokens[i] = strategies[i].underlyingToken();
            tokens[i].approve(address(strategyManager), type(uint256).max);
        }

        address[] memory operators = new address[](numOperatorsPerOpset);
        for (uint256 i = 0; i < operators.length; ++i) {
            operators[i] = address(new Operator(delegationManager));
        }
        return operators;
    }

    function depositForOperators(IStrategy strategy, address[] memory operators, uint256 tokenAmountPerOperator) public {
        IERC20 token = strategy.underlyingToken();
        token.approve(address(strategyManager), type(uint256).max);
        
        for (uint256 i = 0; i < operators.length; ++i) {
            strategyManager.depositIntoStrategyWithSignature(strategy, token, tokenAmountPerOperator, operators[i], type(uint256).max, hex"");
        }
    }

    function allocateForOperators(IStrategy strategy, IAVSDirectory.OperatorSet calldata operatorSet, address[] memory operators, uint64 magnitudeForOperators) public {
        uint64 expectedTotalMagnitude = avsDirectory.INITIAL_TOTAL_MAGNITUDE();

        IAVSDirectory.OperatorSet[] memory operatorSets = new IAVSDirectory.OperatorSet[](1);
        operatorSets[0] = operatorSet;

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = magnitudeForOperators;

        IAVSDirectory.MagnitudeAllocation[] memory allocations = new IAVSDirectory.MagnitudeAllocation[](1);
        allocations[0] = IAVSDirectory.MagnitudeAllocation({
            strategy: strategy,
            expectedTotalMagnitude: expectedTotalMagnitude,
            operatorSets: operatorSets,
            magnitudes: magnitudes
        });

        ISignatureUtils.SignatureWithSaltAndExpiry memory signature = ISignatureUtils.SignatureWithSaltAndExpiry({
            signature: hex"",
            salt: bytes32(0),
            expiry: type(uint256).max
        });

        for (uint256 i = 0; i < operators.length; ++i) {            
            avsDirectory.modifyAllocations(
                operators[i],
                allocations,
                signature
            );
        }
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