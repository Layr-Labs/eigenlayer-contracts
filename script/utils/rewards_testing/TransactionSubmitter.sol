// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/interfaces/IRewardsCoordinator.sol";
import "src/contracts/interfaces/IAVSDirectory.sol";
import "src/contracts/interfaces/IStrategyFactory.sol";
import "./ServiceManagerMock.sol";
import {ISignatureUtils} from "src/contracts/interfaces/ISignatureUtils.sol";
import "./StrategyToken.sol";


/**
 * @title TransactionSubmitter
 * @notice A contract that batch submits transactions to seed EL core contracts
 * @notice We cannot seed 
 */
contract TransactionSubmitter {
    // Pointers to Core Contracts
    IDelegationManager public immutable delegation;
    IStrategyManager public immutable strategyManager;
    IAVSDirectory public immutable avsDirectory;
    IRewardsCoordinator public immutable rewardsCoordinator;
    IStrategyFactory public immutable strategyFactory;

    // Configs
    uint256 immutable MAX_TOKEN_SUPPLY = 1e37;

    // AVS Info
    mapping(uint256 => address) public avss;
    uint256 public numAVSs;

    // Strategy Info
    mapping(uint256 => IStrategy) public deprecated_customStrats;
    uint256 public deprecated_numCustomStrats;

    mapping(uint256 => IStrategy) public customStrats;
    uint256 public numCustomStrats;


    // Registration Utils 
    struct OperatorAVSRegistration {
        address operator;
        address[] avss;
        ISignatureUtils.SignatureWithSaltAndExpiry[] sigs;
    }

    struct StakerDelegation {
        address staker;
        address operator;
        ISignatureUtils.SignatureWithExpiry stakerSignatureAndExpiry;
        ISignatureUtils.SignatureWithExpiry approverSignatureAndExpiry;
        bytes32 approverSalt;
    }

    struct PermitInfo {
        address owner;
        address spender;
        uint256 value;
        uint256 deadline;
        uint8 v;
        bytes32 r;
        bytes32 s;
    }

    struct StrategyInfo {
        IStrategy strategy;
        IERC20 token;
        uint256 amount;
        address staker;
        uint256 expiry;
        bytes signature;
    }

    struct StakerDeposit {
        StrategyInfo[] strategyInfos;
    }

    constructor(IDelegationManager _delegation, IAVSDirectory _avsDirectory, IStrategyManager _strategyManager, IRewardsCoordinator _rewardsCoordinator, IStrategyFactory _strategyFactory) {
        delegation = _delegation;
        avsDirectory = _avsDirectory;
        strategyManager = _strategyManager;
        rewardsCoordinator = _rewardsCoordinator;
        strategyFactory = _strategyFactory;
    }

    /**
     * @notice Deploys `numAVSsToDeploy` AVSs by deploy a ServiceManagerMock contract
     * @param numAVSsToDeploy The number of AVSs to deploy
     */
    function deployAVSs(uint16 numAVSsToDeploy) external {
        for (uint16 i = 0; i < numAVSsToDeploy; i++) {
            address avs = address(new ServiceManagerMock(avsDirectory, rewardsCoordinator));
            avss[numAVSs] = avs;
            numAVSs++;
        }
    }

    /**
     * @notice Registers a list of operators to a list of AVSs
     * @param operatorAVSRegistrations A list of operators and their corresponding AVSs to register for
     */
    function registerOperatorsToAVSs(OperatorAVSRegistration[] memory operatorAVSRegistrations) external {
        for (uint256 i = 0; i < operatorAVSRegistrations.length; i++) {
            OperatorAVSRegistration memory operatorAVSRegistration = operatorAVSRegistrations[i];
            for (uint256 j = 0; j < operatorAVSRegistration.avss.length; j++) {
                ServiceManagerMock(operatorAVSRegistration.avss[j]).registerOperatorToAVS(operatorAVSRegistration.operator, operatorAVSRegistration.sigs[j]);
            }
        }
    }

    /**
     * @notice Deploys `numCustomStratsToDeploy` custom strategies by calling the `StrategyFactory`
     * @param names The names of tokens to deploy
     * @param symbols The symbols of tokens to deploy
     * @dev A new token is deployed for each custom strategy and the entire supply is transferred to the TransactionSubmitter
     */
    function deployCustomStrategies(string[] memory names, string[] memory symbols) external {
        for (uint16 i = 0; i < names.length; i++) {
            // Deploy Strategy Token 
            // Create a token name based on numCustomStrats
            StrategyToken token = new StrategyToken(
                names[i],
                symbols[i],
                MAX_TOKEN_SUPPLY,
                address(this)
            );

            // Deploy Custom Strategy
            IStrategy customStrat = strategyFactory.deployNewStrategy(IERC20(token));
            customStrats[numCustomStrats] = customStrat;
            numCustomStrats++;
        }
    }

    /**
     * @notice Delegates a list of stakers by signature
     * @param stakerDelegations A list of calldatas for staker to delegate
     */
    function delegateStakers(StakerDelegation[] memory stakerDelegations) external {
        for (uint256 i = 0; i < stakerDelegations.length; i++) {
            StakerDelegation memory stakerDelegation = stakerDelegations[i];
            delegation.delegateToBySignature(
                stakerDelegation.staker,
                stakerDelegation.operator,
                stakerDelegation.stakerSignatureAndExpiry,
                stakerDelegation.approverSignatureAndExpiry,
                stakerDelegation.approverSalt
            );
        }
    }

    function depositStakers(StakerDeposit[] memory stakerDeposits) external {
        for (uint256 i = 0; i < stakerDeposits.length; i++) {
            for(uint256 j = 0; j < stakerDeposits[i].strategyInfos.length; j++) {
                StrategyInfo memory strategyInfo = stakerDeposits[i].strategyInfos[j];

                // Approve
                strategyInfo.token.approve(address(strategyManager), strategyInfo.amount);

                // Deposit
                strategyManager.depositIntoStrategyWithSignature(
                    strategyInfo.strategy,
                    strategyInfo.token,
                    strategyInfo.amount,
                    strategyInfo.staker,
                    strategyInfo.expiry,
                    strategyInfo.signature
                );
            }
        }
    }

    /**
     * Sends `amount` of ETH to each address in `tos`
     * @param tos List of addresses to send ETH to
     * @param amount of ETH to send
     */
    function sendEth(address[] memory tos, uint256 amount) external {
        for (uint256 i = 0; i < tos.length; i++) {
            (bool success, ) = payable(tos[i]).call{value: amount}("");
            require(success, "TransactionSubmitter: sendEth failed");
        }
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Returns all AVSs that have been deployed by this contract
     */
    function getAllAVSs() external view returns (address[] memory) {
        address[] memory allAVSs = new address[](numAVSs);
        for (uint256 i = 0; i < numAVSs; i++) {
            allAVSs[i] = avss[i];
        }
        return allAVSs;
    }

    function getAllStrategies() external view returns(IStrategy[] memory) {
        IStrategy[] memory allStrategies = new IStrategy[](numCustomStrats);
        for (uint256 i = 0; i < numCustomStrats; i++) {
            allStrategies[i] = customStrats[i];
        }
        return allStrategies;
    }

    fallback() external {}

    receive() external payable {}
}