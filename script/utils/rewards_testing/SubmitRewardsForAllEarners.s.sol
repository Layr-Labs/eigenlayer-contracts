// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/utils/ExistingDeploymentParser.sol";
import "./TransactionSubmitter.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";


contract SubmitRewardsForAllEarners is ExistingDeploymentParser {
    string contractsPath = "script/configs/holesky/eigenlayer_addresses_preprod.config.json";

    function tx_1() public parseState {
        // Deploy token 
        string memory name = "RewardForAllEarner_Test_1";
        string memory symbol = "RFA_1";
        IERC20 rewardToken = IERC20(_deployStablecoin(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // Format strategies and multipliers
        IStrategy[] memory strategies = _getOldStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: 1
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - rewardsCoordinator.MAX_REWARDS_DURATION();
        uint32 duration = rewardsCoordinator.MAX_REWARDS_DURATION();

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            amount: 10_000e6,
            startTimestamp: startTimestamp,
            duration: duration
        });

        _submitRewardsForAllEarners(rewardsSubmission);
    }

    function tx_2() public parseState {
        // Deploy token 
        string memory name = "RewardForAllEarner_Test_2";
        string memory symbol = "RFA_2";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // Format strategies and multipliers
        IStrategy[] memory strategies = _getOldStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: type(uint96).max
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - (rewardsCoordinator.MAX_REWARDS_DURATION()/2);
        uint32 duration = rewardsCoordinator.MAX_REWARDS_DURATION();

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            amount: 10_000e18,
            startTimestamp: startTimestamp,
            duration: duration
        });

        _submitRewardsForAllEarners(rewardsSubmission);
    }

    function tx_3() public parseState {
        // Deploy token 
        string memory name = "RewardForAllEarner_Test_3";
        string memory symbol = "RFA_3";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // Format strategies and multipliers
        IStrategy[] memory strategies = _getAllStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: uint96(vm.randomUint(1e18, 10e18))
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - 1 weeks;
        uint32 duration = 1 weeks;

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            amount: 1_000_000e18,
            startTimestamp: startTimestamp,
            duration: duration
        });

        _submitRewardsForAllEarners(rewardsSubmission);
    }

    function tx_4() public parseState {
        // Deploy token 
        string memory name = "RewardForAllEarner_Test_4";
        string memory symbol = "RFA_4";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // Format strategies and multipliers
        IStrategy[] memory strategies = _getOldStrategies();
        uint96[] memory multipliers = _getEigenDAMultipliers();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: multipliers[i]
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - 1 weeks;
        uint32 duration = 2 weeks;

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            amount: 100_000e18,
            startTimestamp: startTimestamp,
            duration: duration
        });

        _submitRewardsForAllEarners(rewardsSubmission);
    }

    function tx_8() public parseState {
        // Deploy token 
        string memory name = "RewardForAllEarner_Test_8";
        string memory symbol = "RFA_8";
        IERC20 rewardToken = IERC20(_deployToken(name, symbol));
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](1);

        // Format strategies and multipliers
        IStrategy[] memory strategies = _getTxSubmitterStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: uint96(vm.randomUint(1e18, 10e18))
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = 1724889600;
        uint32 duration = 6 weeks;

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: rewardToken,
            amount: 1_000_000e18,
            startTimestamp: startTimestamp,
            duration: duration
        });

        _submitRewardsForAllEarners(rewardsSubmission);
    }

    function tx_4_and_5() public parseState {
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmission = new IRewardsCoordinator.RewardsSubmission[](2);

        // Tx 4
        IERC20 transaction4RewardToken = IERC20(0x826FaE505F3652eD85A5Fa9cB9ab98AD1AE434de);
        IStrategy[] memory strategies = _getOldStrategies();
        uint96[] memory multipliers = _getEigenDAMultipliers();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliers[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: multipliers[i]
            });
        }

        // Format Range
        uint32 moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        uint32 startTimestamp = moddedCurrTimestamp - 1 weeks;
        uint32 duration = 2 weeks;

        rewardsSubmission[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: transaction4RewardToken,
            amount: 100_000e18,
            startTimestamp: startTimestamp,
            duration: duration
        });


        // Tx 5
        IERC20 transaction5RewardToken = IERC20(0x23bC9Db0e7976B68835C12338715C2DB0ebC9C1e);
        strategies = _getTxSubmitterStrategies();
        IRewardsCoordinator.StrategyAndMultiplier[] memory strategyAndMultipliersTx5 = new IRewardsCoordinator.StrategyAndMultiplier[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategyAndMultipliersTx5[i] = IRewardsCoordinator.StrategyAndMultiplier({
                strategy: strategies[i],
                multiplier: uint96(vm.randomUint(1e18, 10e18))
            });
        }

        // Format Range
        moddedCurrTimestamp = uint32(block.timestamp) - (uint32(block.timestamp) % rewardsCoordinator.CALCULATION_INTERVAL_SECONDS());
        startTimestamp = moddedCurrTimestamp - 1 weeks;
        duration = 2 weeks;

        rewardsSubmission[1] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliersTx5,
            token: transaction5RewardToken,
            amount: 1_000_000e18,
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

    function _deployStablecoin(string memory name, string memory symbol) public returns (address) {
        uint256 tokenInitialSupply = 1_000_000_000e6;
        vm.startBroadcast();
        TestStablecoin rewardToken = new TestStablecoin(
            name,
            symbol,
            tokenInitialSupply,
            msg.sender
        );
        rewardToken.approve(address(rewardsCoordinator), tokenInitialSupply);
        vm.stopBroadcast();
        return address(rewardToken);
    }

    function _submitRewardsForAllEarners(
        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions
    ) internal {
        vm.startBroadcast();
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);
        vm.stopBroadcast();
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