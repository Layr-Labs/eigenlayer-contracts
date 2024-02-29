// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/fork/mainnet/deprecatedInterfaces/IEigenPod.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IEigenPodManager.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IStrategyManager.sol";
import "src/test/integration/User.t.sol";

interface IUserMainnetForkDeployer {
    function delegationManager() external view returns (DelegationManager);
    function strategyManager() external view returns (StrategyManager);
    function eigenPodManager() external view returns (EigenPodManager);
    function timeMachine() external view returns (TimeMachine);
    function beaconChain() external view returns (BeaconChainMock);
    function strategyManager_depM1() external view returns (IStrategyManager_DeprecatedM1);
    function eigenPodManager_depM1() external view returns (IEigenPodManager_DeprecatedM1);
}

contract User_M1 is User {
    IStrategyManager_DeprecatedM1 strategyManager_depM1;
    IEigenPodManager_DeprecatedM1 eigenPodManager_depM1;
    IEigenPod_DeprecatedM1 pod_depM1;

    constructor(string memory name) User(name) {
        IUserMainnetForkDeployer deployer = IUserMainnetForkDeployer(msg.sender);

        strategyManager_depM1 = deployer.strategyManager_depM1();
        eigenPodManager_depM1 = deployer.eigenPodManager_depM1();
        // TODO: Use existing pod from mainnet??
        // pod_depM1 = EigenPod(payable(address(eigenPodManager_depM1.createPod())));
    }


    /**
     * StrategyManager M1 mainnet methods:
     */

    /// @notice Deposit into EigenLayer with M1 mainnet methods
    function depositIntoEigenlayer_M1(
        IStrategy[] memory strategies,
        uint256[] memory tokenBalances
    ) public virtual createSnapshot {
        emit log(_name(".depositIntoEigenlayer"));

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            IERC20 underlyingToken = strat.underlyingToken();
            underlyingToken.approve(address(strategyManager), tokenBalance);
            strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
        }
    }

    function updateBalances_M1(
        IStrategy[] memory strategies,
        int256[] memory tokenDeltas
    ) public virtual createSnapshot {
        emit log(_name(".updateBalances"));

        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            int256 delta = tokenDeltas[i];

            uint256 tokens = uint256(delta);
            IERC20 underlyingToken = strat.underlyingToken();
            underlyingToken.approve(address(strategyManager), tokens);
            strategyManager.depositIntoStrategy(strat, underlyingToken, tokens);
        }
    }

    /// @dev Queues a single withdrawal for every share and strategy pair
    function queueWithdrawals_M1(
        IStrategy[] calldata strategies,
        uint256[] calldata shares
    ) public virtual createSnapshot returns (bytes32, IStrategyManager_DeprecatedM1.QueuedWithdrawal memory) {
        emit log(_name(".queueWithdrawals_M1"));

        uint256[] memory strategyIndexes = new uint256[](shares.length);
        address withdrawer = address(this);
        // delegation not enabled on mainnet
        bool undelegateIfPossible = false;

        IStrategyManager_DeprecatedM1.QueuedWithdrawal memory withdrawal = IStrategyManager_DeprecatedM1.QueuedWithdrawal({
            strategies: strategies,
            shares: shares,
            depositor: address(this),
            withdrawerAndNonce: IStrategyManager_DeprecatedM1.WithdrawerAndNonce({
                withdrawer: withdrawer,
                nonce: uint96(strategyManager_depM1.numWithdrawalsQueued(address(this)))
            }),
            withdrawalStartBlock: uint32(block.number),
            delegatedAddress: address(0)
        });

        bytes32 withdrawalRoot = IStrategyManager_DeprecatedM1(address(strategyManager)).queueWithdrawal(
            strategyIndexes, strategies, shares, withdrawer, undelegateIfPossible
        );

        // Basic sanity check - withdrawalRootPending is set to true
        assertTrue(
            strategyManager_depM1.withdrawalRootPending(withdrawalRoot),
            "StrategyManager withdrawalRootPending should be set to true"
        );

        assertTrue(
            strategyManager_depM1.withdrawalRootPending(
                strategyManager_depM1.calculateWithdrawalRoot(withdrawal)
            ),
            "StrategyManager withdrawal not pending"
        );

        return (withdrawalRoot, withdrawal);
    }

    function completeQueuedWithdrawal_M1(
        IStrategyManager_DeprecatedM1.QueuedWithdrawal memory withdrawal,
        bool receiveAsTokens
    ) public virtual createSnapshot returns (IERC20[] memory) {
        emit log(_name(".completeQueuedWithdrawal_M1"));

        return _completeQueuedWithdrawal(withdrawal, receiveAsTokens);
    }

    function completeQueuedWithdrawals_M1(
        IStrategyManager_DeprecatedM1.QueuedWithdrawal[] memory withdrawals,
        bool receiveAsTokens
    ) public virtual createSnapshot returns (IERC20[][] memory) {
        emit log(_name(".completeQueuedWithdrawals_M1"));

        IERC20[][] memory tokens = new IERC20[][](withdrawals.length);

        for (uint256 i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal(withdrawals[i], receiveAsTokens);
        }

        return tokens;
    }

    function _completeQueuedWithdrawal(
        IStrategyManager_DeprecatedM1.QueuedWithdrawal memory withdrawal,
        bool receiveAsTokens
    ) internal virtual returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint256 i = 0; i < tokens.length; i++) {
            IStrategy strat = withdrawal.strategies[i];
            tokens[i] = strat.underlyingToken();
        }

        strategyManager_depM1.completeQueuedWithdrawal(withdrawal, tokens, 0, receiveAsTokens);

        return tokens;
    }
}
