// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/core/AllocationManager.sol";
import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPod.sol";

import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/mocks/BeaconChainMock.t.sol";
import "src/test/utils/Logger.t.sol";
import "src/test/utils/SingleItemArrayLib.sol";

struct Validator {
    uint40 index;
}

interface IUserDeployer {
    function allocationManager() external view returns (AllocationManager);
    function delegationManager() external view returns (DelegationManager);
    function strategyManager() external view returns (StrategyManager);
    function eigenPodManager() external view returns (EigenPodManager);
    function timeMachine() external view returns (TimeMachine);
    function beaconChain() external view returns (BeaconChainMock);
}

contract User is Logger, IDelegationManagerTypes, IAllocationManagerTypes {
<<<<<<< HEAD
    using ArrayLib for *;
=======
    using SingleItemArrayLib for *;
>>>>>>> b0193bfe (feat: alm tests)
    using print for *;

    AllocationManager allocationManager;
    DelegationManager delegationManager;
    StrategyManager strategyManager;
    EigenPodManager eigenPodManager;
    TimeMachine timeMachine;
    BeaconChainMock beaconChain;

    string _NAME;

    // User's EigenPod and each of their validator indices within that pod
    EigenPod public pod;
    uint40[] validators;

    constructor(
        string memory name
    ) {
        IUserDeployer deployer = IUserDeployer(msg.sender);

        allocationManager = deployer.allocationManager();
        delegationManager = deployer.delegationManager();
        strategyManager = deployer.strategyManager();
        eigenPodManager = deployer.eigenPodManager();
        timeMachine = deployer.timeMachine();

        beaconChain = deployer.beaconChain();
        _createPod();

        _NAME = name;
    }

    modifier createSnapshot() virtual {
        timeMachine.createSnapshot();
        _;
    }

    receive() external payable {}

    function NAME() public view override returns (string memory) {
        return _NAME;
    }

    /// -----------------------------------------------------------------------
    /// Allocation Manager Methods
    /// -----------------------------------------------------------------------
    
    /// @dev Allocates randomly accross the operator set's strategies with a sum of `magnitudeSum`.
    /// NOTE: Calling more than once will lead to deallocations...
    function modifyAllocations(OperatorSet memory operatorSet, uint64[] memory magnitudes) public virtual createSnapshot {
        print.method(
            "modifyAllocations",
            string.concat(
                "{avs: ",
                Logger(operatorSet.avs).NAME_COLORED(),
                ", operatorSetId: ",
                cheats.toString(operatorSet.id),
                "}"
            )
        );

        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);

        require(strategies.length == magnitudes.length, "User.modifyAllocations: length mismatch");

        AllocateParams[] memory params = AllocateParams({
            operatorSet: operatorSet,
            strategies: strategies,
            newMagnitudes: magnitudes
        }).toArray();

        allocationManager.modifyAllocations(address(this), params);
        (, uint32 delay) = allocationManager.getAllocationDelay(address(this));
        rollForward({blocks: delay});
    }
    
    function deallocateAll(
        OperatorSet memory operatorSet
    ) public virtual {
        modifyAllocations(operatorSet, new uint64[](allocationManager.getStrategiesInOperatorSet(operatorSet).length));
    }

    function registerForOperatorSets(
        OperatorSet[] memory operatorSets
    ) public virtual createSnapshot {
        for (uint256 i; i < operatorSets.length; ++i) {
            registerForOperatorSet(operatorSets[i]);
        }
    }

    function registerForOperatorSet(
        OperatorSet memory operatorSet
    ) public virtual createSnapshot {
        print.method(
            "registerForOperatorSet",
            string.concat(
                "{avs: ",
                Logger(operatorSet.avs).NAME_COLORED(),
                ", operatorSetId: ",
                cheats.toString(operatorSet.id),
                "}"
            )
        );

        allocationManager.registerForOperatorSets(
            address(this),
            RegisterParams({avs: operatorSet.avs, operatorSetIds: operatorSet.id.toArrayU32(), data: ""})
        );
    }

    function deregisterFromOperatorSet(
        OperatorSet memory operatorSet
    ) public virtual createSnapshot {
        print.method(
            "deregisterFromOperatorSet",
            string.concat(
                "{avs: ",
                Logger(operatorSet.avs).NAME_COLORED(),
                ", operatorSetId: ",
                cheats.toString(operatorSet.id),
                "}"
            )
        );

        allocationManager.deregisterFromOperatorSets(
            DeregisterParams({
                operator: address(this),
                avs: operatorSet.avs,
                operatorSetIds: operatorSet.id.toArrayU32()
            })
        );
    }

    function setAllocationDelay(uint32 delay) public virtual createSnapshot {
        print.method("setAllocationDelay");
        allocationManager.setAllocationDelay(address(this), delay);
        rollForward({blocks: allocationManager.ALLOCATION_CONFIGURATION_DELAY()});
    }

    /// -----------------------------------------------------------------------
    /// Delegation Manager Methods
    /// -----------------------------------------------------------------------

    uint32 withdrawalDelay = 1;

    function registerAsOperator() public virtual createSnapshot {
        print.method("registerAsOperator");

        delegationManager.registerAsOperator(address(0), withdrawalDelay, "metadata");
    }

    /// @dev Delegate to the operator without a signature
    function delegateTo(
        User operator
    ) public virtual createSnapshot {
        print.method("delegateTo", operator.NAME_COLORED());

        ISignatureUtils.SignatureWithExpiry memory emptySig;
        delegationManager.delegateTo(address(operator), emptySig, bytes32(0));
    }

    /// @dev Undelegate from operator
    function undelegate() public virtual createSnapshot returns (Withdrawal[] memory) {
        print.method("undelegate");

        Withdrawal[] memory expectedWithdrawals = _getExpectedWithdrawalStructsForStaker(address(this));
        delegationManager.undelegate(address(this));

        for (uint256 i = 0; i < expectedWithdrawals.length; i++) {
            emit log("expecting withdrawal:");
            emit log_named_uint("nonce: ", expectedWithdrawals[i].nonce);
            emit log_named_address("strat: ", address(expectedWithdrawals[i].strategies[0]));
            emit log_named_uint("shares: ", expectedWithdrawals[i].scaledShares[0]);
        }

        return expectedWithdrawals;
    }

    /// @dev Force undelegate staker
    function forceUndelegate(
        User staker
    ) public virtual createSnapshot returns (Withdrawal[] memory) {
        print.method("forceUndelegate", staker.NAME());

        Withdrawal[] memory expectedWithdrawals = _getExpectedWithdrawalStructsForStaker(address(staker));
        delegationManager.undelegate(address(staker));
        return expectedWithdrawals;
    }

    /// @dev Queues a single withdrawal for every share and strategy pair
    function queueWithdrawals(
        IStrategy[] memory strategies,
        uint256[] memory depositShares
    ) public virtual createSnapshot returns (Withdrawal[] memory) {
        print.method("queueWithdrawals");

        address operator = delegationManager.delegatedTo(address(this));
        address withdrawer = address(this);
        uint256 nonce = delegationManager.cumulativeWithdrawalsQueued(address(this));

        // Create queueWithdrawals params
        QueuedWithdrawalParams[] memory params = new QueuedWithdrawalParams[](1);
        params[0] =
            QueuedWithdrawalParams({strategies: strategies, depositShares: depositShares, withdrawer: withdrawer});

        // Create Withdrawal struct using same info
        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = Withdrawal({
            staker: address(this),
            delegatedTo: operator,
            withdrawer: withdrawer,
            nonce: nonce,
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: depositShares // TODO: convert depositShares to shares and then scale in withdrawal
        });

        bytes32[] memory withdrawalRoots = delegationManager.queueWithdrawals(params);

        // Basic sanity check - we do all other checks outside this file
        assertEq(withdrawals.length, withdrawalRoots.length, "User.queueWithdrawals: length mismatch");

        return (withdrawals);
    }

    function completeWithdrawalsAsTokens(
        Withdrawal[] memory withdrawals
    ) public virtual createSnapshot returns (IERC20[][] memory) {
        print.method("completeWithdrawalsAsTokens");

        IERC20[][] memory tokens = new IERC20[][](withdrawals.length);

        for (uint256 i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal(withdrawals[i], true);
        }

        return tokens;
    }

    function completeWithdrawalAsTokens(
        Withdrawal memory withdrawal
    ) public virtual createSnapshot returns (IERC20[] memory) {
        print.method("completeWithdrawalsAsTokens");

        return _completeQueuedWithdrawal(withdrawal, true);
    }

    function completeWithdrawalsAsShares(
        Withdrawal[] memory withdrawals
    ) public virtual createSnapshot returns (IERC20[][] memory) {
        print.method("completeWithdrawalAsShares");

        IERC20[][] memory tokens = new IERC20[][](withdrawals.length);

        for (uint256 i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal(withdrawals[i], false);
        }

        return tokens;
    }

    function completeWithdrawalAsShares(
        Withdrawal memory withdrawal
    ) public virtual createSnapshot returns (IERC20[] memory) {
        print.method("completeWithdrawalAsShares");

        return _completeQueuedWithdrawal(withdrawal, false);
    }

    /// -----------------------------------------------------------------------
    /// Beacon Chain Methods
    /// -----------------------------------------------------------------------

    /// @dev Uses any ETH held by the User to start validators on the beacon chain
    /// @return A list of created validator indices
    /// @return The amount of wei sent to the beacon chain
    /// Note: If the user does not have enough ETH to start a validator, this method reverts
    /// Note: This method also advances one epoch forward on the beacon chain, so that
    /// withdrawal credential proofs are generated for each validator.
    function startValidators() public virtual createSnapshot returns (uint40[] memory, uint64) {
        print.method("startValidators");

        return _startValidators();
    }

    function exitValidators(
        uint40[] memory _validators
    ) public virtual createSnapshot returns (uint64 exitedBalanceGwei) {
        print.method("exitValidators");

        return _exitValidators(_validators);
    }

    /// -----------------------------------------------------------------------
    /// Eigenpod Methods
    /// -----------------------------------------------------------------------

    function verifyWithdrawalCredentials(
        uint40[] memory _validators
    ) public virtual createSnapshot {
        print.method("verifyWithdrawalCredentials");

        _verifyWithdrawalCredentials(_validators);
    }

    function startCheckpoint() public virtual createSnapshot {
        print.method("startCheckpoint");

        _startCheckpoint();
    }

    function completeCheckpoint() public virtual createSnapshot {
        print.method("completeCheckpoint");

        _completeCheckpoint();
    }

    function verifyStaleBalance(
        uint40 validatorIndex
    ) public virtual createSnapshot {
        print.method("verifyStaleBalance");

        StaleBalanceProofs memory proof = beaconChain.getStaleBalanceProofs(validatorIndex);

        try pod.verifyStaleBalance({
            beaconTimestamp: proof.beaconTimestamp,
            stateRootProof: proof.stateRootProof,
            proof: proof.validatorProof
        }) {} catch (bytes memory err) {
            _revert(err);
        }
    }

    /// -----------------------------------------------------------------------
    /// Strategy Methods
    /// -----------------------------------------------------------------------

    /// @dev For each strategy/token balance, call the relevant deposit method
    function depositIntoEigenlayer(
        IStrategy[] memory strategies,
        uint256[] memory tokenBalances
    ) public virtual createSnapshot {
        print.method("depositIntoEigenlayer");

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
                strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
            }
        }
    }

    function updateBalances(IStrategy[] memory strategies, int256[] memory tokenDeltas) public virtual createSnapshot {
        print.method("updateBalances");

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
                strategyManager.depositIntoStrategy(strat, underlyingToken, tokens);
            }
        }
    }

    /// -----------------------------------------------------------------------
    /// Internal Methods
    /// -----------------------------------------------------------------------

    function _completeQueuedWithdrawal(
        Withdrawal memory withdrawal,
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

        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);

        return tokens;
    }

    function _createPod() internal virtual {
        pod = EigenPod(payable(eigenPodManager.createPod()));
    }

    /// @dev Uses any ETH held by the User to start validators on the beacon chain
    /// @return A list of created validator indices
    /// @return The amount of wei sent to the beacon chain
    /// Note: If the user does not have enough ETH to start a validator, this method reverts
    /// Note: This method also advances one epoch forward on the beacon chain, so that
    /// withdrawal credential proofs are generated for each validator.
    function _startValidators() internal returns (uint40[] memory, uint64) {
        uint256 balanceWei = address(this).balance;

        // Number of full validators: balance / 32 ETH
        uint256 numValidators = balanceWei / 32 ether;
        balanceWei -= (numValidators * 32 ether);

        // If we still have at least 1 ETH left over, we can create another (non-full) validator
        // Note that in the mock beacon chain this validator will generate rewards like any other.
        // The main point is to ensure pods are able to handle validators that have less than 32 ETH
        uint256 lastValidatorBalance;
        uint256 totalValidators = numValidators;
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
        for (uint256 i = 0; i < numValidators; i++) {
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

        return (newValidators, totalBeaconBalanceGwei);
    }

    function _exitValidators(
        uint40[] memory _validators
    ) internal returns (uint64 exitedBalanceGwei) {
        console.log("- exiting num validators", _validators.length);

        for (uint256 i = 0; i < _validators.length; i++) {
            exitedBalanceGwei += beaconChain.exitValidator(_validators[i]);
        }

        console.log("- exited balance to pod (gwei)", exitedBalanceGwei);

        return exitedBalanceGwei;
    }

    function _startCheckpoint() internal {
        try pod.startCheckpoint(false) {}
        catch (bytes memory err) {
            _revert(err);
        }
    }

    function _completeCheckpoint() internal {
        cheats.pauseTracing();
        console.log("- active validator count", pod.activeValidatorCount());
        console.log("- proofs remaining", pod.currentCheckpoint().proofsRemaining);

        uint64 checkpointTimestamp = pod.currentCheckpointTimestamp();
        if (checkpointTimestamp == 0) {
            revert("User._completeCheckpoint: no existing checkpoint");
        }

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, checkpointTimestamp);
        console.log("- submitting num checkpoint proofs", proofs.balanceProofs.length);

        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        cheats.resumeTracing();
    }

    function _verifyWithdrawalCredentials(
        uint40[] memory _validators
    ) internal {
        cheats.pauseTracing();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(_validators);

        try pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: _validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        }) {} catch (bytes memory err) {
            _revert(err);
        }
        cheats.resumeTracing();
    }

    /// @dev Revert, passing through an error message
    function _revert(
        bytes memory err
    ) internal pure {
        if (err.length != 0) {
            assembly {
                revert(add(32, err), mload(err))
            }
        }
        revert("reverted with unknown error");
    }

    function _podWithdrawalCredentials() internal view returns (bytes memory) {
        return abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(pod));
    }

    /// @notice Gets the expected withdrawals to be created when the staker is undelegated via a call to `DelegationManager.undelegate()`
    /// @notice Assumes staker and withdrawer are the same and that all strategies and shares are withdrawn
    function _getExpectedWithdrawalStructsForStaker(
        address staker
    ) internal view returns (Withdrawal[] memory) {
        (IStrategy[] memory strategies, uint256[] memory shares) = delegationManager.getDepositedShares(staker);

        Withdrawal[] memory expectedWithdrawals = new Withdrawal[](strategies.length);
        address delegatedTo = delegationManager.delegatedTo(staker);
        uint256 nonce = delegationManager.cumulativeWithdrawalsQueued(staker);

        for (uint256 i = 0; i < strategies.length; ++i) {
            IStrategy[] memory singleStrategy = new IStrategy[](1);
            uint256[] memory singleShares = new uint256[](1);
            singleStrategy[0] = strategies[i];
            singleShares[0] = shares[i];
            expectedWithdrawals[i] = Withdrawal({
                staker: staker,
                delegatedTo: delegatedTo,
                withdrawer: staker,
                nonce: (nonce + i),
                startBlock: uint32(block.number),
                strategies: singleStrategy,
                scaledShares: singleShares
            });
        }

        return expectedWithdrawals;
    }

    function getActiveValidators() public view returns (uint40[] memory) {
        uint40[] memory activeValidators = new uint40[](validators.length);

        uint256 numActive;
        uint256 pos;
        for (uint256 i = 0; i < validators.length; i++) {
            if (beaconChain.isActive(validators[i])) {
                activeValidators[pos] = validators[i];
                numActive++;
                pos++;
            }
        }

        // Manually update length
        assembly {
            mstore(activeValidators, numActive)
        }

        return activeValidators;
    }
}

/// @notice A user contract that calls nonstandard methods (like xBySignature methods)
contract User_AltMethods is User {
    mapping(bytes32 => bool) public signedHashes;

    constructor(
        string memory name
    ) User(name) {}

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
                uint256 nonceBefore = strategyManager.nonces(address(this));
                bytes32 structHash = keccak256(
                    abi.encode(
                        strategyManager.DEPOSIT_TYPEHASH(),
                        address(this),
                        strat,
                        underlyingToken,
                        tokenBalance,
                        nonceBefore,
                        expiry
                    )
                );
                bytes32 digestHash =
                    keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));
                bytes memory signature = bytes(abi.encodePacked(digestHash)); // dummy sig data

                // Mark hash as signed
                signedHashes[digestHash] = true;

                // Deposit
                strategyManager.depositIntoStrategyWithSignature(
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
