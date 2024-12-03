// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./TransactionSubmitter.sol";
import "script/utils/ExistingDeploymentParser.sol";

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "zeus-templates/utils/ZeusScript.sol";

contract TestRewardsV2 is ZeusScript {
    uint256 maxOperatorIndex = 80;
    string MNEMONIC;
    TransactionSubmitter transactionSubmitter;
    string statePath = "script/utils/rewards_testing/preprod_state.json";


    function tx_prep() public parseState {
        // Deploy token 
        (address[] memory operatorsRegisteredToAVS, uint256[] memory privateKeys) = _getAllOperatorsRegisteredToAVS(address(_rewardingServiceManager()));

        // for the first 3/5 of the operators, set the AVS split to random values
        for (uint256 i = 0; i < operatorsRegisteredToAVS.length*3/5; i++) {
            vm.startBroadcast(privateKeys[i]);
            _rewardsCoordinator().setOperatorAVSSplit(operatorsRegisteredToAVS[i], address(_rewardingServiceManager()), uint16(vm.randomUint(1, 10000)));
            vm.stopBroadcast();
        }

        // for 1/5 through 4/5 of the operators, set the PI split to random values
        for (uint256 i = operatorsRegisteredToAVS.length*1/5; i < operatorsRegisteredToAVS.length*4/5; i++) {
            vm.startBroadcast(privateKeys[i]);
            _rewardsCoordinator().setOperatorPISplit(operatorsRegisteredToAVS[i], uint16(vm.randomUint(1, 10000)));
            vm.stopBroadcast();
        }

        // this leaves 1/5 of the operators with no splits set
    }

    function tx_1() public parseState {
        // Deploy token 
        string memory name = "RewardsV2_OperatorDirectedRewards_Test_1";
        string memory symbol = "R2_ODR_1";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.OperatorDirectedRewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.OperatorDirectedRewardsSubmission[](1);

        // Format strategies and multipliers
        IStrategy[] memory strategies = _getAVStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: 1e18
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % _rewardsCoordinator().CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - _rewardsCoordinator().MAX_REWARDS_DURATION();
        uint32 duration = _rewardsCoordinator().MAX_REWARDS_DURATION();

        (address[] memory operatorsRegisteredToAVS,) = _getAllOperatorsRegisteredToAVS(address(_rewardingServiceManager()));

        rewardsSubmission[0] = IRewardsCoordinator.OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            operatorRewards: _getRandomOperatorRewards(operatorsRegisteredToAVS, 1e18),
            startTimestamp: startTimestamp,
            duration: duration,
            description: "Test 1"
        });

        _submitOperatorDirectedRewards(rewardsSubmission);
    }

    // Send a regular rewards submission
    function tx_2() public parseState {
        // Deploy token 
        string memory name = "RewardsV2_AVSRewardsSubmission_Test_2";
        string memory symbol = "R2_AVSR_2";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // TODO: replace
        IStrategy[] memory strategies = _getAVStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: 1e18
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % _rewardsCoordinator().CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - 1 weeks;
        uint32 duration = 1 weeks;

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            amount: 1e36,
            startTimestamp: startTimestamp,
            duration: duration
        });

        _submitAVSRewardsSubmission(rewardsSubmission);
    }

    // Send a PI submission
    function tx_3() public parseState {
        // Deploy token 
        string memory name = "RewardsV2_RFA_Test_3";
        string memory symbol = "R2_RFA_3";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // Format strategies and multipliers
        // This is to all the other old stakers on preprod, so use the old strategies
        IStrategy[] memory strategies = _getAVStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: 1e18
            });
        }

        // Format Range
        uint32 calculationIntervalSeconds = _rewardsCoordinator().CALCULATION_INTERVAL_SECONDS();
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % calculationIntervalSeconds);
        uint32 startTimestamp = moddedCurrTimestamp;
        uint32 duration = calculationIntervalSeconds;

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            amount: 1e36,
            startTimestamp: startTimestamp,
            duration: duration
        });

        _submitRewardsForAllEarners(rewardsSubmission);
    }

    function _deployToken(string memory name, string memory symbol) public returns (address) {
        uint256 tokenInitialSupply = 1e36;
        vm.startBroadcast();
        ERC20PresetFixedSupply rewardToken = new ERC20PresetFixedSupply(
            name,
            symbol,
            tokenInitialSupply,
            msg.sender
        );
        rewardToken.approve(address(_rewardsCoordinator()), tokenInitialSupply);
        vm.stopBroadcast();
        return address(rewardToken);
    }

    function _submitOperatorDirectedRewards(
        IRewardsCoordinator.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions
    ) internal {
        ServiceManagerMock serviceManager = _rewardingServiceManager();
        vm.startBroadcast();
        for (uint256 i = 0; i < rewardsSubmissions.length; i++) {
            rewardsSubmissions[i].token.approve(address(serviceManager), type(uint256).max);
        }
        serviceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
        vm.stopBroadcast();
    }

    function _submitAVSRewardsSubmission(
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions
    ) internal {
        ServiceManagerMock serviceManager = _rewardingServiceManager();
        vm.startBroadcast();
        for (uint256 i = 0; i < rewardsSubmissions.length; i++) {
            rewardsSubmissions[i].token.approve(address(serviceManager), type(uint256).max);
        }
        serviceManager.createAVSRewardsSubmission(rewardsSubmissions);
        vm.stopBroadcast();
    }

    function _submitRewardsForAllEarners(
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions
    ) internal {
        vm.startBroadcast();
        for (uint256 i = 0; i < rewardsSubmissions.length; i++) {
            rewardsSubmissions[i].token.approve(address(_rewardsCoordinator()), type(uint256).max);
        }
        _rewardsCoordinator().createRewardsForAllEarners(rewardsSubmissions);
        vm.stopBroadcast();
    }

    function _getAllOperatorsRegisteredToAVS(address avs) internal returns (address[] memory, uint256[] memory) {
        AVSDirectory avsDirectory = AVSDirectory(zDeployedProxy(type(AVSDirectory).name));

        // Get all operators
        (address[] memory allOperators, uint256[] memory privateKeys) = _getAllRegisteredOperators();

        // Get which operators are registered to the AVS
        bool[] memory isRegisteredToAVS = new bool[](allOperators.length);
        uint256 numOperatorsRegisteredToAVS = 0;
        for (uint256 i = 0; i < allOperators.length; i++) {
            isRegisteredToAVS[i] = avsDirectory.avsOperatorStatus(avs, allOperators[i]) == IAVSDirectory.OperatorAVSRegistrationStatus.REGISTERED;
            if (isRegisteredToAVS[i]) {
                numOperatorsRegisteredToAVS++;
            }
        }

        // Filter the operators that are registered to the AVS
        address[] memory operatorsRegisteredToAVS = new address[](numOperatorsRegisteredToAVS);
        uint256 numRegistered = 0;
        for (uint256 i = 0; i < allOperators.length; i++) {
            if (isRegisteredToAVS[i]) {
                operatorsRegisteredToAVS[numRegistered] = allOperators[i];
                numRegistered++;
            }
            if (numRegistered == numOperatorsRegisteredToAVS) {
                break;
            }
        }

        // sort the operators by address
        for (uint256 i = 0; i < operatorsRegisteredToAVS.length; i++) {
            for (uint256 j = i + 1; j < operatorsRegisteredToAVS.length; j++) {
                if (operatorsRegisteredToAVS[i] > operatorsRegisteredToAVS[j]) {
                    address temp = operatorsRegisteredToAVS[i];
                    operatorsRegisteredToAVS[i] = operatorsRegisteredToAVS[j];
                    operatorsRegisteredToAVS[j] = temp;
                }
            }
        }
        return (operatorsRegisteredToAVS, privateKeys);
    }

    modifier parseState() {
        _parseState(statePath);
        _;
    }

    function _parseState(string memory statePathToParse) internal {
        // READ JSON CONFIG DATA
        string memory stateData = vm.readFile(statePathToParse);
        emit log_named_string("Using state file", statePathToParse);

        transactionSubmitter = TransactionSubmitter(payable(stdJson.readAddress(stateData, ".submitterProxy")));
        MNEMONIC = vm.envString("MNEMONIC");
    }

    function _rewardingServiceManager() internal view returns (ServiceManagerMock) {
        address[] memory avss = transactionSubmitter.getAllAVSs();

        return ServiceManagerMock(avss[0]);
    }

    function _rewardsCoordinator() internal view returns (IRewardsCoordinator) {
        return IRewardsCoordinator(zDeployedProxy(type(RewardsCoordinator).name));
    }

    /// @notice Strategies that are previously used in the system
    function _getOldStrategies() internal pure returns (IStrategy[] memory) {
        IStrategy[] memory strategies = new IStrategy[](6);
        strategies[0] = IStrategy(0x5C8b55722f421556a2AAfb7A3EA63d4c3e514312);
        strategies[1] = IStrategy(0x6E5D5060B33ca2090A78E9cb74Fe207453b30E49);
        strategies[2] = IStrategy(0x7fA77c321bf66e42eaBC9b10129304F7f90c5585);
        strategies[3] = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
        strategies[4] = IStrategy(0xD523267698C81a372191136e477fdebFa33D9FB4);
        strategies[5] = IStrategy(0xdcCF401fD121d8C542E96BC1d0078884422aFAD2);

        return strategies;
    }

    function _getAVStrategies() internal pure returns (IStrategy[] memory) {
        IStrategy[] memory strategies = new IStrategy[](15);

        strategies[0] = IStrategy(0x08f8544E61Ebfa22e7c6ef9af9eFd428091b27AF);
        strategies[1] = IStrategy(0x31741340ab31e90f0624d2C843B323a97755f43a);
        strategies[2] = IStrategy(0x845aDA440F3BA5652e6f7c6Bbcd8fCAd068Be9c9);
        strategies[3] = IStrategy(0xcD6EDf68a5Fc79BBc14A0eCa1C8bbDe495fB1a5f);
        strategies[4] = IStrategy(0x77335a08a877cd874165bDC766feeD951a0d84c8);
        strategies[5] = IStrategy(0x53C916338FcA4f5541bc2eA5E5f7F0Ca6f0dAf6a);
        strategies[6] = IStrategy(0x77335a08a877cd874165bDC766feeD951a0d84c8);
        strategies[7] = IStrategy(0x51366AEf35475e82B864f2E7867B293c3BC61D86);
        strategies[8] = IStrategy(0x5803f7D0D273aa0B9774C0aDAB98A23ec348Ea77);
        strategies[9] = IStrategy(0xe5e8BcBd1DDD07460fA67f617a2D018ce0f5b7bf);
        strategies[10] = IStrategy(0x419B6Ba569169826d9A45565E651F03B8DdfA332);
        strategies[11] = IStrategy(0x9b5F92ed40e1436Ac8E31fd7f7bc6083edBd5E71);
        strategies[12] = IStrategy(0xE34b1Da874E6B9f091540c850095b99B4D3B475D);
        strategies[13] = IStrategy(0x8EF42562E9F1b81010a54615673739dd91E10016);
        strategies[14] = IStrategy(0xdE2788cb747b51a1747Bd59342b6214652Ddeb10);

        // sort strategies by address
        for (uint256 i = 0; i < strategies.length; i++) {
            for (uint256 j = i + 1; j < strategies.length; j++) {
                if (address(strategies[i]) > address(strategies[j])) {
                    IStrategy temp = strategies[i];
                    strategies[i] = strategies[j];
                    strategies[j] = temp;
                }
            }
        }

        return strategies;
    }

    function _getAllRegisteredOperators() internal returns(address[] memory, uint256[] memory){
        address[] memory operators = new address[](maxOperatorIndex);
        uint256[] memory privateKeys = new uint256[](maxOperatorIndex);
        for(uint256 i = 0; i < operators.length; i++){
            (address operator, uint256 privateKey) = deriveRememberKey(MNEMONIC, uint32(i));
            operators[i] = operator;
            privateKeys[i] = privateKey;
        }
        return (operators, privateKeys);
    }

    function _getRandomOperatorRewards(address[] memory operators, uint256 max) internal returns (IRewardsCoordinator.OperatorReward[] memory) {
        IRewardsCoordinator.OperatorReward[] memory operatorRewards = new IRewardsCoordinator.OperatorReward[](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            operatorRewards[i] = IRewardsCoordinator.OperatorReward({
                operator: operators[i],
                amount: vm.randomUint(1, max)
            });
        }
        return operatorRewards;
    }

}

contract TestStablecoin is ERC20PresetFixedSupply {
    constructor(
        string memory name,
        string memory symbol,
        uint256 initialSupply,
        address owner
    ) ERC20PresetFixedSupply(name, symbol, initialSupply, owner) {}

    function decimals() public pure override returns (uint8) {
        return 6;
    }
}