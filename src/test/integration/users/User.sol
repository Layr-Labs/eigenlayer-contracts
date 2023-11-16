// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

abstract contract User {

    IDelegationManager delegationManager;
    IStrategyManager strategyManager;
    IEigenPodManager eigenPodManager;

    Global global;

    WithdrawalType withdrawalType;

    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
        IEigenPodManager _eigenPodManager,
        Global _global
    ) {
        delegationManager = _delegationManager;
        strategyManager = _strategyManager;
        eigenPodManager = _eigenPodManager;
        global = _global;
    }

    modifier createSnapshot() virtual {
        global.createSnapshot();
        _;
    }

    /// @dev For each strategy/token balance, call the relevant deposit method
    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot virtual {

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // TODO handle this flow - need to deposit into EPM + prove credentials
                revert("TODO: unimplemented");
            } else {
                IERC20 underlyingToken = strat.underlyingToken();
                strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
            }
        }
    }

    /// @dev Delegate to the operator without a signature
    function delegateTo(User operator) public createSnapshot virtual {
        ISignatureUtils.SignatureWithExpiry memory emptySig;
        delegationManager.delegateTo(operator, emptySig, bytes32(0));
    }

    /// @dev Queue withdrawals in the DelegationManager
    function queueWithdrawals(IDelegationManager.QueuedWithdrawalParams memory params) public createSnapshot virtual returns (bytes32[] memory) {
        return delegationManager.queueWithdrawals(params);
    }

    function completeQueuedWithdrawal(
        IDelegationManager.Withdrawal memory withdrawal, 
        bool receiveAsTokens
    ) public createSnapshot virtual returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint i = 0; i < tokens.length; i++) {
            IStrategy strat = withdrawal.strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokens[i] = IERC20(address(0));
            } else {
                tokens[i] = strat.underlyingToken();
            }
        }

        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0, receiveAsTokens);

        return tokens;
    }

    function _completeQueuedWithdrawal(
        IDelegationManager.Withdrawal memory withdrawal,
        bool receiveAsTokens
    ) public virtual {
        delegationManager.completeQueuedWithdrawal(withdrawal, )
    }
}

contract User_MixedAssets is User {
    /// Since this is the base setup, we don't need anything else
}

contract User_LSTOnly is User {

    /// @dev For each strategy/token balance, call the relevant deposit method
    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot override {

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];
            IERC20 underlyingToken = strat.underlyingToken();

            strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
        }
    }
}

contract User_NativeOnly is User {

    /// @dev For each strategy/token balance, call the relevant deposit method
    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot override {

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];
            
            // TODO handle this flow - need to deposit into EPM + prove credentials
            revert("TODO: unimplemented");
        }
    }
}