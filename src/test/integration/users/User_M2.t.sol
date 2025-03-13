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
import "src/test/utils/ArrayLib.sol";

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
 * @dev UserM2 interacts with pre-pectra beacon chain
 */
contract User_M2 is User {
    using ArrayLib for *;
    using print for *;

    IDelegationManager_DeprecatedM2 delegationManager_M2;
    IStrategyManager_DeprecatedM2 strategyManager_M2;
    IEigenPodManager_DeprecatedM2 eigenPodManager_M2;

    constructor(string memory name) User(name) {
        IUserM2MainnetForkDeployer deployer = IUserM2MainnetForkDeployer(msg.sender);

        delegationManager_M2 = IDelegationManager_DeprecatedM2(address(deployer.delegationManager()));
        strategyManager_M2 = IStrategyManager_DeprecatedM2(address(deployer.strategyManager()));
        eigenPodManager_M2 = IEigenPodManager_DeprecatedM2(address(deployer.eigenPodManager()));
        cheats.label(address(this), NAME_COLORED());
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

    /// @dev Queues a single withdrawal for every share and strategy pair
    /// @dev Returns the withdrawal struct of the new slashing interface
    function queueWithdrawals(IStrategy[] memory strategies, uint[] memory shares)
        public
        virtual
        override
        createSnapshot
        returns (Withdrawal[] memory)
    {
        print.method("queueWithdrawals_M2");

        address operator = delegationManager_M2.delegatedTo(address(this));
        address withdrawer = address(this);
        uint nonce = delegationManager_M2.cumulativeWithdrawalsQueued(address(this));

        // Create queueWithdrawals params
        IDelegationManager_DeprecatedM2.QueuedWithdrawalParams[] memory params =
            new IDelegationManager_DeprecatedM2.QueuedWithdrawalParams[](1);
        params[0] = IDelegationManager_DeprecatedM2.QueuedWithdrawalParams({strategies: strategies, shares: shares, withdrawer: withdrawer});

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

        Withdrawal[] memory withdrawalsToReturn = new Withdrawal[](1);
        withdrawalsToReturn[0] = Withdrawal({
            staker: address(this),
            delegatedTo: operator,
            withdrawer: withdrawer,
            nonce: nonce,
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: shares
        });

        return (withdrawalsToReturn);
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

    function updateBalances(IStrategy[] memory strategies, int[] memory tokenDeltas) public virtual override createSnapshot {
        print.method("updateBalances_M2");

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            int delta = tokenDeltas[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // If any balance update has occurred, a checkpoint will pick it up
                _startCheckpoint();
                if (pod.activeValidatorCount() != 0) _completeCheckpoint_M2();
            } else {
                uint tokens = uint(delta);
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokens);
                strategyManager_M2.depositIntoStrategy(strat, underlyingToken, tokens);
            }
        }
    }

    /// -----------------------------------------------------------------------
    /// Internal Methods
    /// -----------------------------------------------------------------------

    /// @dev Uses any ETH held by the User to start validators on the beacon chain
    /// @return A list of created validator indices
    /// @return The amount of wei sent to the beacon chain
    /// @return The number of validators that have the MaxEB
    /// Note: If the user does not have enough ETH to start a validator, this method reverts
    /// Note: This method also advances one epoch forward on the beacon chain, so that
    /// withdrawal credential proofs are generated for each validator.
    function _startValidators() internal virtual override returns (uint40[] memory, uint64, uint) {
        uint balanceWei = address(this).balance;

        // Number of full validators: balance / 32 ETH
        uint numValidators = balanceWei / 32 ether;
        balanceWei -= (numValidators * 32 ether);

        // If we still have at least 1 ETH left over, we can create another (non-full) validator
        // Note that in the mock beacon chain this validator will generate rewards like any other.
        // The main point is to ensure pods are able to handle validators that have less than 32 ETH
        uint lastValidatorBalance;
        uint totalValidators = numValidators;
        if (balanceWei >= 1 ether) {
            lastValidatorBalance = balanceWei - (balanceWei % 1 gwei);
            balanceWei -= lastValidatorBalance;
            totalValidators++;
        }

        require(totalValidators != 0, "startValidators: not enough ETH to start a validator");
        uint40[] memory newValidators = new uint40[](totalValidators);
        uint64 totalBeaconBalanceGwei = uint64((address(this).balance - balanceWei) / GWEI_TO_WEI);

        console.log("- creating new validators", newValidators.length);
        console.log("- depositing balance to beacon chain (gwei)", totalBeaconBalanceGwei);

        // Create each of the full validators
        for (uint i = 0; i < numValidators; i++) {
            uint40 validatorIndex = beaconChain.newValidator{value: 32 ether}(_podWithdrawalCredentials());

            newValidators[i] = validatorIndex;
            validators.push(validatorIndex);
        }

        // If we had a remainder, create the final, non-full validator
        if (totalValidators == numValidators + 1) {
            uint40 validatorIndex = beaconChain.newValidator{value: lastValidatorBalance}(_podWithdrawalCredentials());

            newValidators[newValidators.length - 1] = validatorIndex;
            validators.push(validatorIndex);
        }

        return (newValidators, totalBeaconBalanceGwei, numValidators);
    }

    function _completeCheckpoint_M2() internal {
        cheats.pauseTracing();
        IEigenPod_DeprecatedM2 pod = IEigenPod_DeprecatedM2(address(pod));
        console.log("- active validator count", pod.activeValidatorCount());
        console.log("- proofs remaining", pod.currentCheckpoint().proofsRemaining);

        uint64 checkpointTimestamp = pod.currentCheckpointTimestamp();
        if (checkpointTimestamp == 0) revert("User._completeCheckpoint: no existing checkpoint");

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, checkpointTimestamp);
        console.log("- submitting num checkpoint proofs", proofs.balanceProofs.length);

        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        cheats.resumeTracing();
    }

    function _completeQueuedWithdrawal_M2(IDelegationManager_DeprecatedM2.Withdrawal memory withdrawal, bool receiveAsTokens)
        internal
        virtual
        returns (IERC20[] memory)
    {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint i = 0; i < tokens.length; i++) {
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
                    if (pod.activeValidatorCount() != 0) _completeCheckpoint();
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
    function _getExpectedM2WithdrawalStructsForStaker(address staker)
        internal
        view
        returns (IDelegationManager_DeprecatedM2.Withdrawal[] memory)
    {
        (IStrategy[] memory strategies, uint[] memory shares) = delegationManager_M2.getDelegatableShares(staker);

        IDelegationManager_DeprecatedM2.Withdrawal[] memory expectedWithdrawals =
            new IDelegationManager_DeprecatedM2.Withdrawal[](strategies.length);
        address delegatedTo = delegationManager_M2.delegatedTo(staker);
        uint nonce = delegationManager_M2.cumulativeWithdrawalsQueued(staker);

        for (uint i = 0; i < strategies.length; ++i) {
            IStrategy[] memory singleStrategy = new IStrategy[](1);
            uint[] memory singleShares = new uint[](1);
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

    constructor(string memory name) User_M2(name) {}

    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public override createSnapshot {
        print.method("depositIntoEigenlayer_ALT");

        uint expiry = type(uint).max;
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                (uint40[] memory newValidators,,) = _startValidators();
                // Advance forward one epoch and generate credential and balance proofs for each validator
                beaconChain.advanceEpoch_NoRewards();
                _verifyWithdrawalCredentials(newValidators);
            } else {
                // Approve token
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokenBalance);

                // Get signature
                uint nonceBefore = strategyManager_M2.nonces(address(this));
                bytes32 structHash = keccak256(
                    abi.encode(
                        strategyManager_M2.DEPOSIT_TYPEHASH(), address(this), strat, underlyingToken, tokenBalance, nonceBefore, expiry
                    )
                );
                bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager_M2.domainSeparator(), structHash));
                bytes memory signature = bytes(abi.encodePacked(digestHash)); // dummy sig data

                // Mark hash as signed
                signedHashes[digestHash] = true;

                // Deposit
                strategyManager_M2.depositIntoStrategyWithSignature(strat, underlyingToken, tokenBalance, address(this), expiry, signature);

                // Mark hash as used
                signedHashes[digestHash] = false;
            }
        }
    }

    bytes4 internal constant MAGIC_VALUE = 0x1626ba7e;

    function isValidSignature(bytes32 hash, bytes memory) external view returns (bytes4) {
        if (signedHashes[hash]) return MAGIC_VALUE;
        else return 0xffffffff;
    }
}
