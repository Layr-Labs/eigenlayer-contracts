// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ExistingDeploymentParser.sol";
import "./TransactionSubmitter.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";


contract SubmitRewardsForAllEarners is ExistingDeploymentParser {
    string contractsPath = "script/configs/holesky/eigenlayer_addresses_preprod.config.json";

    uint256 maxOperatorIndex = 800;
    string MNEMONIC = vm.envString("MNEMONIC");

    // hardcode one service manager for now
    ServiceManagerMock rewardingServiceManager = ServiceManagerMock(payable(0xa5d5E9bcdDC1dACe96E5d8f7536A97900550BbB2));

    function tx_prep() public parseState {
        // Deploy token 
        (address[] memory operatorsRegisteredToAVS, uint256[] memory privateKeys) = _getAllOperatorsRegisteredToAVS(address(rewardingServiceManager));

        // for the first 3/5 of the operators, set the AVS split to random values
        for (uint256 i = 0; i < operatorsRegisteredToAVS.length*3/5; i++) {
            vm.broadcast(privateKeys[i]);
            rewardsCoordinator.setOperatorAVSSplit(operatorsRegisteredToAVS[i], address(rewardingServiceManager), uint16(vm.randomUint(1, 10000)));
        }

        // for 1/5 through 4/5 of the operators, set the PI split to random values
        for (uint256 i = operatorsRegisteredToAVS.length*1/5; i < operatorsRegisteredToAVS.length*4/5; i++) {
            vm.broadcast(privateKeys[i]);
            rewardsCoordinator.setOperatorPISplit(operatorsRegisteredToAVS[i], uint16(vm.randomUint(1, 10000)));
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
        IStrategy[] memory strategies = _getOldStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: 1e18
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - rewardsCoordinator.MAX_REWARDS_DURATION();
        uint32 duration = rewardsCoordinator.MAX_REWARDS_DURATION();

        (address[] memory operatorsRegisteredToAVS,) = _getAllOperatorsRegisteredToAVS(address(rewardingServiceManager));

        rewardsSubmission[0] = IRewardsCoordinator.OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            operatorRewards: _getRandomOperatorRewards(operatorsRegisteredToAVS, 1e18),
            startTimestamp: startTimestamp,
            duration: duration,
            description: "Test 1"
        });

        _submitOperatorDirectedRewards(rewardingServiceManager, rewardsSubmission);
    }

    // Send a regular rewards submission
    function tx_2() public parseState {
        // Deploy token 
        string memory name = "RewardsV2_AVSRewardsSubmission_Test_2";
        string memory symbol = "R2_AVSR_2";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // TODO: replace
        IStrategy[] memory strategies = _getOldStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: 1e18
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - 1 weeks;
        uint32 duration = 1 weeks;

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            amount: 1e36,
            startTimestamp: startTimestamp,
            duration: duration
        });

        _submitAVSRewardsSubmission(rewardingServiceManager, rewardsSubmission);
    }

    // Send a PI submission
    function tx_3() public parseState {
        // Deploy token 
        string memory name = "RewardsV2_RFA_Test_3";
        string memory symbol = "R2_RFA_3";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // TODO: replace
        // Format strategies and multipliers
        IStrategy[] memory strategies = _getOldStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: 1e18
            });
        }

        // Format Range
        uint32 calculationIntervalSeconds = rewardsCoordinator.CALCULATION_INTERVAL_SECONDS();
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
        rewardToken.approve(address(rewardsCoordinator), tokenInitialSupply);
        vm.stopBroadcast();
        return address(rewardToken);
    }

    function _submitOperatorDirectedRewards(
        ServiceManagerMock serviceManager,
        IRewardsCoordinator.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions
    ) internal {
        vm.startBroadcast();
        for (uint256 i = 0; i < rewardsSubmissions.length; i++) {
            rewardsSubmissions[i].token.approve(address(serviceManager), type(uint256).max);
        }
        serviceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
        vm.stopBroadcast();
    }

    function _submitAVSRewardsSubmission(
        ServiceManagerMock serviceManager,
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions
    ) internal {
        vm.startBroadcast();
        serviceManager.createAVSRewardsSubmission(rewardsSubmissions);
        vm.stopBroadcast();
    }

    function _submitRewardsForAllEarners(
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions
    ) internal {
        vm.startBroadcast();
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);
        vm.stopBroadcast();
    }

    function _getAllOperatorsRegisteredToAVS(address avs) internal returns (address[] memory, uint256[] memory) {
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

    function _getTxSubmitterStrategies() public view returns (IStrategy[] memory) {
        TransactionSubmitter txSubmitter = TransactionSubmitter(payable(0x5a80a187a159A5d8dfC5bD16236E5657e945ec3d));
        IStrategy[] memory strategies = txSubmitter.getAllStrategies();

        // Sort strategies by address
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

    function _getEigenDAMultipliers() internal pure returns (uint96[] memory) {
        uint96[] memory multipliers = new uint96[](6);

        multipliers[0] = 1152393415227598758;
        multipliers[1] = 1102456657360376283;
        multipliers[2] = 1028802524926876401;
        multipliers[3] = 1000000000000000000;
        multipliers[4] = 1012495275290785447;
        multipliers[5] = 1035345726488000000;

        return multipliers;
    }

    function _getAllStrategies() internal view returns (IStrategy[] memory) {
        IStrategy[] memory oldStrategies = _getOldStrategies();
        IStrategy[] memory newStrategies = _getTxSubmitterStrategies();
        IStrategy[] memory allStrategies = new IStrategy[](oldStrategies.length + newStrategies.length);

        // Create a sorted array, allStrategies, that returns the concatenation of old and new strateiges but in sorted order
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;
        while (i < oldStrategies.length && j < newStrategies.length) {
            if (address(oldStrategies[i]) < address(newStrategies[j])) {
                allStrategies[k] = oldStrategies[i];
                i++;
            } else {
                allStrategies[k] = newStrategies[j];
                j++;
            }
            k++;
        }

        while (i < oldStrategies.length) {
            allStrategies[k] = oldStrategies[i];
            i++;
            k++;
        }

        while (j < newStrategies.length) {
            allStrategies[k] = newStrategies[j];
            j++;
            k++;
        }

        return allStrategies;
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

    modifier parseState() {
        _parseDeployedContracts(contractsPath);
        _;
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