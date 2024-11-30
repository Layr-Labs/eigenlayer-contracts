// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/test/integration/deprecatedInterfaces/mainnet/IEigenPod.sol";
import "src/test/integration/deprecatedInterfaces/mainnet/IEigenPodManager.sol";
import "src/test/integration/deprecatedInterfaces/mainnet/IStrategyManager.sol";
import "src/test/integration/deprecatedInterfaces/mainnet/IDelegationManager.sol";

import "src/test/integration/users/User.t.sol";
import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/mocks/BeaconChainMock.t.sol";
import "src/test/utils/Logger.t.sol";
import "src/test/utils/SingleItemArrayLib.sol";

interface IUserM2MainnetForkDeployer {
    function delegationManager() external view returns (DelegationManager);
    function strategyManager() external view returns (StrategyManager);
    function eigenPodManager() external view returns (EigenPodManager);
    function delegationManager_M2() external view returns (IDelegationManager_DeprecatedM2);
    function strategyManager_M2() external view returns (IStrategyManager_DeprecatedM2);
    function eigenPodManager_M2() external view returns (IEigenPodManager_DeprecatedM2);
    function timeMachine() external view returns (TimeMachine);
    function beaconChain() external view returns (BeaconChainMock);
}

/**
 * @dev User_M2 used for performing mainnet M2 contract methods but also inherits User
 * to perform current local contract methods after a upgrade of core contracts
 */
contract User_M2 is User {
    using SingleItemArrayLib for *;
    using print for *;

    IDelegationManager_DeprecatedM2 delegationManager_M2;
    IStrategyManager_DeprecatedM2 strategyManager_M2;
    IEigenPodManager_DeprecatedM2 eigenPodManager_M2;

    constructor(
        string memory name
    ) User(name) {
        IUserM2MainnetForkDeployer deployer = IUserM2MainnetForkDeployer(msg.sender);

        delegationManager_M2 = IDelegationManager_DeprecatedM2(address(deployer.delegationManager()));
        strategyManager_M2 = IStrategyManager_DeprecatedM2(address(deployer.strategyManager()));
        eigenPodManager_M2 = IEigenPodManager_DeprecatedM2(address(deployer.eigenPodManager()));

    }

    /// -----------------------------------------------------------------------
    /// Delegation Manager Methods
    /// -----------------------------------------------------------------------

    function registerAsOperator_M2() public virtual createSnapshot {
        print.method("registerAsOperator_M2");

        IDelegationManager_DeprecatedM2.OperatorDetails memory details = IDelegationManager_DeprecatedM2.OperatorDetails({
            __deprecated_earningsReceiver: address(this),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });

        delegationManager_M2.registerAsOperator(details, "metadata");
    }

    /// @dev Delegate to the operator without a signature
    function delegateTo_M2(
        User operator
    ) public virtual createSnapshot {
        print.method("delegateTo_M2", operator.NAME_COLORED());

        ISignatureUtils.SignatureWithExpiry memory emptySig;
        delegationManager_M2.delegateTo(address(operator), emptySig, bytes32(0));
    }

    /// @dev Undelegate from operator
    function undelegate_M2() public virtual createSnapshot returns (IDelegationManager_DeprecatedM2.Withdrawal[] memory) {
        print.method("undelegate_M2");

        IDelegationManager_DeprecatedM2.Withdrawal[] memory expectedWithdrawals = _getExpectedM2WithdrawalStructsForStaker(address(this));
        delegationManager_M2.undelegate(address(this));

        for (uint256 i = 0; i < expectedWithdrawals.length; i++) {
            emit log("expecting withdrawal:");
            emit log_named_uint("nonce: ", expectedWithdrawals[i].nonce);
            emit log_named_address("strat: ", address(expectedWithdrawals[i].strategies[0]));
            emit log_named_uint("shares: ", expectedWithdrawals[i].shares[0]);
        }

        return expectedWithdrawals;
    }

    /// @dev Force undelegate staker
    function forceUndelegate_M2(
        User staker
    ) public virtual createSnapshot returns (IDelegationManager_DeprecatedM2.Withdrawal[] memory) {
        print.method("forceUndelegate_M2", staker.NAME());

        IDelegationManager_DeprecatedM2.Withdrawal[] memory expectedWithdrawals = _getExpectedM2WithdrawalStructsForStaker(address(this));
        delegationManager_M2.undelegate(address(staker));
        return expectedWithdrawals;
    }

    /// @dev Queues a single withdrawal for every share and strategy pair
    function queueWithdrawals_M2(
        IStrategy[] memory strategies,
        uint256[] memory shares
    ) public virtual createSnapshot returns (IDelegationManager_DeprecatedM2.Withdrawal[] memory) {
        print.method("queueWithdrawals_M2");

        address operator = delegationManager_M2.delegatedTo(address(this));
        address withdrawer = address(this);
        uint256 nonce = delegationManager_M2.cumulativeWithdrawalsQueued(address(this));

        // Create queueWithdrawals params
        IDelegationManager_DeprecatedM2.QueuedWithdrawalParams[] memory params = new IDelegationManager_DeprecatedM2.QueuedWithdrawalParams[](1);
        params[0] =
            IDelegationManager_DeprecatedM2.QueuedWithdrawalParams({strategies: strategies, shares: shares, withdrawer: withdrawer});

        // Create Withdrawal struct using same info
        IDelegationManager_DeprecatedM2.Withdrawal[] memory withdrawals = new IDelegationManager_DeprecatedM2.Withdrawal[](1);
        withdrawals[0] = IDelegationManager_DeprecatedM2.Withdrawal({
            staker: address(this),
            delegatedTo: operator,
            withdrawer: withdrawer,
            nonce: nonce,
            startBlock: uint32(block.number),
            strategies: strategies,
            shares: shares
        });

        bytes32[] memory withdrawalRoots = delegationManager_M2.queueWithdrawals(params);

        // Basic sanity check - we do all other checks outside this file
        assertEq(withdrawals.length, withdrawalRoots.length, "User.queueWithdrawals: length mismatch");

        return (withdrawals);
    }

    function completeWithdrawalsAsTokens_M2(
        IDelegationManager_DeprecatedM2.Withdrawal[] memory withdrawals
    ) public virtual createSnapshot returns (IERC20[][] memory) {
        print.method("completeWithdrawalsAsTokens_M2");

        IERC20[][] memory tokens = new IERC20[][](withdrawals.length);

        for (uint256 i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal_M2(withdrawals[i], true);
        }

        return tokens;
    }

    function completeWithdrawalAsTokens_M2(
        IDelegationManager_DeprecatedM2.Withdrawal memory withdrawal
    ) public virtual createSnapshot returns (IERC20[] memory) {
        print.method("completeWithdrawalsAsTokens_M2");

        return _completeQueuedWithdrawal_M2(withdrawal, true);
    }

    function completeWithdrawalsAsShares_M2(
        IDelegationManager_DeprecatedM2.Withdrawal[] memory withdrawals
    ) public virtual createSnapshot returns (IERC20[][] memory) {
        print.method("completeWithdrawalAsShares_M2");

        IERC20[][] memory tokens = new IERC20[][](withdrawals.length);

        for (uint256 i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal_M2(withdrawals[i], false);
        }

        return tokens;
    }

    function completeWithdrawalAsShares_M2(
        IDelegationManager_DeprecatedM2.Withdrawal memory withdrawal
    ) public virtual createSnapshot returns (IERC20[] memory) {
        print.method("completeWithdrawalAsShares_M2");

        return _completeQueuedWithdrawal_M2(withdrawal, false);
    }

    /// -----------------------------------------------------------------------
    /// Eigenpod Methods
    /// -----------------------------------------------------------------------

    function completeCheckpoint() public virtual override createSnapshot {
        print.method("completeCheckpoint_M2");

        _completeCheckpoint_M2();
    }

    /// -----------------------------------------------------------------------
    /// Strategy Methods
    /// -----------------------------------------------------------------------

    /// @dev For each strategy/token balance, call the relevant deposit method
    function depositIntoEigenlayer_M2(
        IStrategy[] memory strategies,
        uint256[] memory tokenBalances
    ) public virtual createSnapshot {
        print.method("depositIntoEigenlayer_M2");

        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint256 tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                (uint40[] memory newValidators,) = _startValidators();
                // Advance forward one epoch and generate credential and balance proofs for each validator
                beaconChain.advanceEpoch_NoRewards();
                _verifyWithdrawalCredentials(newValidators);
            } else {
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokenBalance);
                strategyManager_M2.depositIntoStrategy(strat, underlyingToken, tokenBalance);
            }
        }
    }

    function updateBalances_M2(IStrategy[] memory strategies, int256[] memory tokenDeltas) public virtual createSnapshot {
        print.method("updateBalances_M2");

        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            int256 delta = tokenDeltas[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // If any balance update has occured, a checkpoint will pick it up
                _startCheckpoint();
                if (pod.activeValidatorCount() != 0) {
                    _completeCheckpoint();
                }
            } else {
                uint256 tokens = uint256(delta);
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokens);
                strategyManager_M2.depositIntoStrategy(strat, underlyingToken, tokens);
            }
        }
    }

    /// -----------------------------------------------------------------------
    /// Internal Methods
    /// -----------------------------------------------------------------------

    function _completeCheckpoint_M2() internal {
        cheats.pauseTracing();
        console.log("- active validator count", pod.activeValidatorCount());
        // Comment out this method due to interface incompatability
        // console.log("- proofs remaining", pod.currentCheckpoint().proofsRemaining);

        uint64 checkpointTimestamp = pod.currentCheckpointTimestamp();
        if (checkpointTimestamp == 0) {
            revert("User._completeCheckpoint: no existing checkpoint");
        }

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, checkpointTimestamp);
        console.log("- submitting num checkpoint proofs", proofs.balanceProofs.length);

        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        cheats.resumeTracing();
    }


    function _completeQueuedWithdrawal_M2(
        IDelegationManager_DeprecatedM2.Withdrawal memory withdrawal,
        bool receiveAsTokens
    ) internal virtual returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint256 i = 0; i < tokens.length; i++) {
            IStrategy strat = withdrawal.strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokens[i] = NATIVE_ETH;

                // If we're withdrawing native ETH as tokens, stop ALL validators
                // and complete a checkpoint
                if (receiveAsTokens) {
                    console.log("- exiting all validators and completing checkpoint");
                    _exitValidators(getActiveValidators());

                    beaconChain.advanceEpoch_NoRewards();

                    _startCheckpoint();
                    if (pod.activeValidatorCount() != 0) {
                        _completeCheckpoint();
                    }
                }
            } else {
                tokens[i] = strat.underlyingToken();
            }
        }

        delegationManager_M2.completeQueuedWithdrawal(withdrawal, tokens, 0, receiveAsTokens);

        return tokens;
    }

    /// @notice Gets the expected withdrawals to be created when the staker is undelegated via a call to `delegationManager_M2.undelegate()`
    /// @notice Assumes staker and withdrawer are the same and that all strategies and shares are withdrawn
    function _getExpectedM2WithdrawalStructsForStaker(
        address staker
    ) internal view returns (IDelegationManager_DeprecatedM2.Withdrawal[] memory) {
        (IStrategy[] memory strategies, uint256[] memory shares)
            = delegationManager_M2.getDelegatableShares(staker);

        IDelegationManager_DeprecatedM2.Withdrawal[] memory expectedWithdrawals = new IDelegationManager_DeprecatedM2.Withdrawal[](strategies.length);
        address delegatedTo = delegationManager_M2.delegatedTo(staker);
        uint256 nonce = delegationManager_M2.cumulativeWithdrawalsQueued(staker);

        for (uint256 i = 0; i < strategies.length; ++i) {
            IStrategy[] memory singleStrategy = new IStrategy[](1);
            uint256[] memory singleShares = new uint256[](1);
            singleStrategy[0] = strategies[i];
            singleShares[0] = shares[i];
            expectedWithdrawals[i] = IDelegationManager_DeprecatedM2.Withdrawal({
                staker: staker,
                delegatedTo: delegatedTo,
                withdrawer: staker,
                nonce: (nonce + i),
                startBlock: uint32(block.number),
                strategies: singleStrategy,
                shares: singleShares
            });
        }

        return expectedWithdrawals;
    }
}

/// @notice A user contract that calls nonstandard methods (like xBySignature methods)
contract User_M2_AltMethods is User_M2 {
    mapping(bytes32 => bool) public signedHashes;

    constructor(
        string memory name
    ) User_M2(name) {}

    function depositIntoEigenlayer(
        IStrategy[] memory strategies,
        uint256[] memory tokenBalances
    ) public override createSnapshot {
        print.method("depositIntoEigenlayer_ALT");

        uint256 expiry = type(uint256).max;
        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint256 tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                (uint40[] memory newValidators,) = _startValidators();
                // Advance forward one epoch and generate credential and balance proofs for each validator
                beaconChain.advanceEpoch_NoRewards();
                _verifyWithdrawalCredentials(newValidators);
            } else {
                // Approve token
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokenBalance);

                // Get signature
                uint256 nonceBefore = strategyManager_M2.nonces(address(this));
                bytes32 structHash = keccak256(
                    abi.encode(
                        strategyManager_M2.DEPOSIT_TYPEHASH(),
                        address(this),
                        strat,
                        underlyingToken,
                        tokenBalance,
                        nonceBefore,
                        expiry
                    )
                );
                bytes32 digestHash =
                    keccak256(abi.encodePacked("\x19\x01", strategyManager_M2.domainSeparator(), structHash));
                bytes memory signature = bytes(abi.encodePacked(digestHash)); // dummy sig data

                // Mark hash as signed
                signedHashes[digestHash] = true;

                // Deposit
                strategyManager_M2.depositIntoStrategyWithSignature(
                    strat, underlyingToken, tokenBalance, address(this), expiry, signature
                );

                // Mark hash as used
                signedHashes[digestHash] = false;
            }
        }
    }

    bytes4 internal constant MAGIC_VALUE = 0x1626ba7e;

    function isValidSignature(bytes32 hash, bytes memory) external view returns (bytes4) {
        if (signedHashes[hash]) {
            return MAGIC_VALUE;
        } else {
            return 0xffffffff;
        }
    }
}
