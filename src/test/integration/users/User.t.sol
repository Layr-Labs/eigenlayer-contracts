// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/test/Config.t.sol";
import "src/test/utils/ArrayLib.sol";
import "src/test/utils/Constants.t.sol";
import "src/test/utils/Logger.t.sol";

struct Validator {
    uint40 index;
}

contract User is Logger, IDelegationManagerTypes, IAllocationManagerTypes {
    using StdStyle for *;
    using SlashingLib for *;
    using ArrayLib for *;
    using print for *;

    uint32 public allocationDelay = 1;

    string _NAME;

    // User's EigenPod and each of their validator indices within that pod
    EigenPod public pod;
    uint40[] validators;

    ConfigGetters public config;

    function eq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }

    constructor(string memory name) {
        config = ConfigGetters(address(msg.sender));

        string memory chain = cheats.envOr(string("FORK_CHAIN"), string("local"));

        if (eq(chain, "mainnet") || eq(chain, "local")) _createPod();

        _NAME = name;
        cheats.label(address(this), NAME_COLORED());
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

    function modifyAllocations(AllocateParams memory params) public virtual createSnapshot {
        print.method(
            "modifyAllocations",
            string.concat(
                "{avs: ", Logger(params.operatorSet.avs).NAME_COLORED(), ", operatorSetId: ", cheats.toString(params.operatorSet.id), "}"
            )
        );

        _tryPrankAppointee_AllocationManager(IAllocationManager.modifyAllocations.selector);
        config.allocationManager().modifyAllocations(address(this), params.toArray());
        print.gasUsed();
    }

    function deallocateAll(OperatorSet memory operatorSet) public virtual returns (AllocateParams memory) {
        AllocateParams memory params;
        params.operatorSet = operatorSet;
        params.strategies = config.allocationManager().getStrategiesInOperatorSet(operatorSet);
        params.newMagnitudes = new uint64[](params.strategies.length);

        modifyAllocations(params);

        return params;
    }

    function registerForOperatorSets(OperatorSet[] memory operatorSets) public virtual createSnapshot {
        for (uint i; i < operatorSets.length; ++i) {
            registerForOperatorSet(operatorSets[i]);
        }
    }

    function registerForOperatorSet(OperatorSet memory operatorSet) public virtual createSnapshot {
        print.method(
            "registerForOperatorSet",
            string.concat("{avs: ", Logger(operatorSet.avs).NAME_COLORED(), ", operatorSetId: ", cheats.toString(operatorSet.id), "}")
        );

        _tryPrankAppointee_AllocationManager(IAllocationManager.registerForOperatorSets.selector);
        config.allocationManager().registerForOperatorSets(
            address(this), RegisterParams({avs: operatorSet.avs, operatorSetIds: operatorSet.id.toArrayU32(), data: ""})
        );
        print.gasUsed();
    }

    function deregisterFromOperatorSet(OperatorSet memory operatorSet) public virtual createSnapshot {
        print.method(
            "deregisterFromOperatorSet",
            string.concat("{avs: ", Logger(operatorSet.avs).NAME_COLORED(), ", operatorSetId: ", cheats.toString(operatorSet.id), "}")
        );

        _tryPrankAppointee_AllocationManager(IAllocationManager.deregisterFromOperatorSets.selector);
        config.allocationManager().deregisterFromOperatorSets(
            DeregisterParams({operator: address(this), avs: operatorSet.avs, operatorSetIds: operatorSet.id.toArrayU32()})
        );
        print.gasUsed();
    }

    function setAllocationDelay(uint32 delay) public virtual createSnapshot {
        print.method("setAllocationDelay");
        _tryPrankAppointee_AllocationManager(IAllocationManager.setAllocationDelay.selector);
        config.allocationManager().setAllocationDelay(address(this), delay);
        print.gasUsed();

        allocationDelay = delay;
    }

    /// -----------------------------------------------------------------------
    /// Delegation Manager Methods
    /// -----------------------------------------------------------------------

    function registerAsOperator() public virtual createSnapshot {
        print.method("registerAsOperator");
        config.delegationManager().registerAsOperator(address(0), allocationDelay, "metadata");
        print.gasUsed();
    }

    /// @dev Delegate to the operator without a signature
    function delegateTo(User operator) public virtual createSnapshot {
        print.method("delegateTo", operator.NAME_COLORED());

        ISignatureUtilsMixinTypes.SignatureWithExpiry memory emptySig;
        config.delegationManager().delegateTo(address(operator), emptySig, bytes32(0));
        print.gasUsed();
    }

    /// @dev Undelegate from operator
    function undelegate() public virtual createSnapshot returns (Withdrawal[] memory) {
        print.method("undelegate");

        Withdrawal[] memory expectedWithdrawals = _getExpectedWithdrawalStructsForStaker(address(this));
        _tryPrankAppointee_DelegationManager(IDelegationManager.undelegate.selector);
        config.delegationManager().undelegate(address(this));
        print.gasUsed();

        for (uint i = 0; i < expectedWithdrawals.length; i++) {
            IStrategy strat = expectedWithdrawals[i].strategies[0];

            string memory name = strat == BEACONCHAIN_ETH_STRAT ? "Native ETH" : IERC20Metadata(address(strat.underlyingToken())).name();

            console.log(
                "   Expecting withdrawal with nonce %s of %s for %s scaled shares.",
                expectedWithdrawals[i].nonce,
                name,
                expectedWithdrawals[i].scaledShares[0]
            );
        }

        return expectedWithdrawals;
    }

    /// @dev Redelegate to a new operator
    function redelegate(User newOperator) public virtual createSnapshot returns (Withdrawal[] memory) {
        print.method("redelegate", newOperator.NAME_COLORED());
        Withdrawal[] memory expectedWithdrawals = _getExpectedWithdrawalStructsForStaker(address(this));
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory emptySig;
        _tryPrankAppointee_DelegationManager(IDelegationManager.redelegate.selector);
        config.delegationManager().redelegate(address(newOperator), emptySig, bytes32(0));
        print.gasUsed();

        for (uint i = 0; i < expectedWithdrawals.length; i++) {
            IStrategy strat = expectedWithdrawals[i].strategies[0];

            string memory name = strat == BEACONCHAIN_ETH_STRAT ? "Native ETH" : IERC20Metadata(address(strat.underlyingToken())).name();

            console.log(
                "   Expecting withdrawal with nonce %s of %s for %s scaled shares.",
                expectedWithdrawals[i].nonce,
                name,
                expectedWithdrawals[i].scaledShares[0]
            );
        }
        return expectedWithdrawals;
    }

    /// @dev Force undelegate staker
    function forceUndelegate(User staker) public virtual createSnapshot returns (Withdrawal[] memory) {
        print.method("forceUndelegate", staker.NAME());

        Withdrawal[] memory expectedWithdrawals = _getExpectedWithdrawalStructsForStaker(address(staker));
        config.delegationManager().undelegate(address(staker));
        print.gasUsed();

        return expectedWithdrawals;
    }

    /// @dev Queues a single withdrawal for every share and strategy pair
    function queueWithdrawals(IStrategy[] memory strategies, uint[] memory depositShares)
        public
        virtual
        createSnapshot
        returns (Withdrawal[] memory)
    {
        print.method("queueWithdrawals");

        address operator = config.delegationManager().delegatedTo(address(this));
        uint nonce = config.delegationManager().cumulativeWithdrawalsQueued(address(this));

        // Create queueWithdrawals params
        QueuedWithdrawalParams[] memory params = new QueuedWithdrawalParams[](1);
        params[0] = QueuedWithdrawalParams({strategies: strategies, depositShares: depositShares, __deprecated_withdrawer: address(0)});

        uint[] memory scaledSharesForWithdrawal = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; ++i) {
            DepositScalingFactor memory dsf =
                DepositScalingFactor(config.delegationManager().depositScalingFactor(address(this), strategies[i]));

            scaledSharesForWithdrawal[i] = dsf.scaleForQueueWithdrawal(depositShares[i]);
        }

        // Create Withdrawal struct using same info
        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = Withdrawal({
            staker: address(this),
            delegatedTo: operator,
            withdrawer: address(this),
            nonce: nonce,
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: scaledSharesForWithdrawal
        });

        bytes32[] memory withdrawalRoots = config.delegationManager().queueWithdrawals(params);
        print.gasUsed();

        // Basic sanity check - we do all other checks outside this file
        assertEq(withdrawals.length, withdrawalRoots.length, "User.queueWithdrawals: length mismatch");

        return (withdrawals);
    }

    function completeWithdrawalsAsTokens(Withdrawal[] memory withdrawals)
        public
        virtual
        createSnapshot
        returns (IERC20[][] memory tokens)
    {
        print.method("completeWithdrawalsAsTokens");
        tokens = new IERC20[][](withdrawals.length);
        for (uint i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal(withdrawals[i], true);
        }
    }

    function completeWithdrawalAsTokens(Withdrawal memory withdrawal) public virtual createSnapshot returns (IERC20[] memory) {
        print.method("completeWithdrawalsAsTokens");
        return _completeQueuedWithdrawal(withdrawal, true);
    }

    function completeWithdrawalsAsShares(Withdrawal[] memory withdrawals)
        public
        virtual
        createSnapshot
        returns (IERC20[][] memory tokens)
    {
        print.method("completeWithdrawalsAsShares");
        tokens = new IERC20[][](withdrawals.length);
        for (uint i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal(withdrawals[i], false);
        }
    }

    function completeWithdrawalAsShares(Withdrawal memory withdrawal) public virtual createSnapshot returns (IERC20[] memory) {
        print.method("completeWithdrawalAsShares");
        return _completeQueuedWithdrawal(withdrawal, false);
    }

    /// -----------------------------------------------------------------------
    /// Beacon Chain Methods
    /// -----------------------------------------------------------------------

    /// @dev Uses any ETH held by the User to start validators on the beacon chain
    /// @return A list of created validator indices
    /// @return The amount of wei sent to the beacon chain
    /// @return The number of validators that have the MaxEB
    /// Note: If the user does not have enough ETH to start a validator, this method reverts
    /// Note: This method also advances one epoch forward on the beacon chain, so that
    /// withdrawal credential proofs are generated for each validator.
    function startValidators() public virtual createSnapshot returns (uint40[] memory, uint64, uint) {
        print.method("startValidators");
        return _startValidators();
    }

    /// @dev Starts a specified number of validators for 32 ETH each on the beacon chain
    /// @param numValidators The number of validators to start
    /// @return A list of created validator indices
    /// @return The amount of wei sent to the beacon chain
    function startValidators(uint8 numValidators) public virtual createSnapshot returns (uint40[] memory, uint64, uint) {
        require(numValidators > 0 && numValidators <= 10, "startValidators: numValidators must be between 1 and 10");
        uint balanceWei = address(this).balance;

        // given a number of validators, the current balance, calculate the amount of ETH needed to start that many validators
        uint ethNeeded = numValidators * 32 ether - balanceWei;
        cheats.deal(address(this), ethNeeded);

        print.method("startValidators");
        return _startValidators();
    }

    function exitValidators(uint40[] memory _validators) public virtual createSnapshot returns (uint64 exitedBalanceGwei) {
        print.method("exitValidators");
        return _exitValidators(_validators);
    }

    /// -----------------------------------------------------------------------
    /// Eigenpod Methods
    /// -----------------------------------------------------------------------

    function verifyWithdrawalCredentials(uint40[] memory _validators) public virtual createSnapshot {
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

    function verifyStaleBalance(uint40 validatorIndex) public virtual createSnapshot {
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
    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public virtual createSnapshot {
        print.method("depositIntoEigenlayer");

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                (uint40[] memory newValidators,,) = _startValidators();
                // Advance forward one epoch and generate credential and balance proofs for each validator
                beaconChain.advanceEpoch_NoRewards();
                _verifyWithdrawalCredentials(newValidators);
            } else {
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(config.strategyManager()), tokenBalance);
                config.strategyManager().depositIntoStrategy(strat, underlyingToken, tokenBalance);
                print.gasUsed();
            }
        }
    }

    function updateBalances(IStrategy[] memory strategies, int[] memory tokenDeltas) public virtual createSnapshot {
        print.method("updateBalances");

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            int delta = tokenDeltas[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // If any balance update has occurred, a checkpoint will pick it up
                _startCheckpoint();
                if (pod.activeValidatorCount() != 0) _completeCheckpoint();
            } else {
                uint tokens = uint(delta);
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(config.strategyManager()), tokens);
                config.strategyManager().depositIntoStrategy(strat, underlyingToken, tokens);
                print.gasUsed();
            }
        }
    }

    /// -----------------------------------------------------------------------
    /// View Methods
    /// -----------------------------------------------------------------------

    function getSlashingFactor(IStrategy strategy) public view returns (uint) {
        return _getSlashingFactor(address(this), strategy);
    }

    /// -----------------------------------------------------------------------
    /// Internal Methods
    /// -----------------------------------------------------------------------

    function _completeQueuedWithdrawal(Withdrawal memory withdrawal, bool receiveAsTokens) internal virtual returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint i = 0; i < tokens.length; i++) {
            IStrategy strat = withdrawal.strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokens[i] = NATIVE_ETH;

                // If we're withdrawing native ETH as tokens && do not have negative shares
                // stop ALL validators and complete a checkpoint
                if (receiveAsTokens && config.eigenPodManager().podOwnerDepositShares(address(this)) >= 0) {
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

        config.delegationManager().completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);
        print.gasUsed();

        return tokens;
    }

    function _createPod() internal virtual {
        pod = EigenPod(payable(config.eigenPodManager().createPod()));
    }

    /// @dev Uses any ETH held by the User to start validators on the beacon chain
    /// @dev Creates validators between 32 and 2048 ETH
    /// @return A list of created validator indices
    /// @return The amount of wei sent to the beacon chain
    /// @return The number of validators that have the MaxEB
    /// Note: If the user does not have enough ETH to start a validator, this method reverts
    /// Note: This method also advances one epoch forward on the beacon chain, so that
    /// withdrawal credential proofs are generated for each validator.
    function _startValidators() internal virtual returns (uint40[] memory, uint64, uint) {
        uint originalBalance = address(this).balance;
        uint balanceWei = address(this).balance;
        uint numValidators = 0;
        uint maxEBValidators = 0;
        // Get maximum possible number of validators. Add 1 to account for a validator with < 32 ETH
        uint maxValidators = balanceWei / 32 ether;
        uint40[] memory newValidators = new uint40[](maxValidators + 1);

        // Create validators between 32 and 2048 ETH until we can't create more
        while (balanceWei >= 32 ether) {
            // Generate random validator balance between 32 and 2048 ETH
            uint validatorEth = uint(keccak256(abi.encodePacked(block.timestamp, balanceWei, numValidators))) % 64 + 1; // 1-64 multiplier
            validatorEth *= 32 ether; // Results in 32-2048 ETH

            // If we don't have enough ETH for the random amount, use remaining balance
            // as long as it's >= 32 ETH
            if (balanceWei < validatorEth) {
                if (balanceWei >= 32 ether) validatorEth = balanceWei - (balanceWei % 32 ether);
                else break;
            }

            // Track validators with maximum effective balance
            if (validatorEth == 2048 ether) maxEBValidators++;

            // Create the validator
            bytes memory withdrawalCredentials =
                validatorEth == 32 ether ? _podWithdrawalCredentials() : _podCompoundingWithdrawalCredentials();

            uint40 validatorIndex = beaconChain.newValidator{value: validatorEth}(withdrawalCredentials);

            newValidators[numValidators] = validatorIndex;
            validators.push(validatorIndex);

            balanceWei -= validatorEth;
            numValidators++;
        }

        // If we still have at least 1 ETH left over, we can create another (non-full) validator
        // Note that in the mock beacon chain this validator will generate rewards like any other.
        // The main point is to ensure pods are able to handle validators that have less than 32 ETH
        if (balanceWei >= 1 ether) {
            uint lastValidatorBalance = balanceWei - (balanceWei % 1 gwei);

            uint40 validatorIndex = beaconChain.newValidator{value: lastValidatorBalance}(_podWithdrawalCredentials());

            newValidators[numValidators] = validatorIndex;
            validators.push(validatorIndex);

            balanceWei -= lastValidatorBalance;
            numValidators++;
        }

        // Resize the array to actual number of validators created
        assembly {
            mstore(newValidators, numValidators)
        }

        require(numValidators != 0, "startValidators: not enough ETH to start a validator");

        uint64 totalBeaconBalanceGwei = uint64((originalBalance - balanceWei) / GWEI_TO_WEI);
        console.log("- created new validators", newValidators.length);
        console.log("- deposited balance to beacon chain (gwei)", totalBeaconBalanceGwei);

        // Advance forward one epoch and generate withdrawal and balance proofs for each validator
        beaconChain.advanceEpoch_NoRewards();

        return (newValidators, totalBeaconBalanceGwei, maxEBValidators);
    }

    function _exitValidators(uint40[] memory _validators) internal returns (uint64 exitedBalanceGwei) {
        console.log("- exiting num validators", _validators.length);

        for (uint i = 0; i < _validators.length; i++) {
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
        if (checkpointTimestamp == 0) revert("User._completeCheckpoint: no existing checkpoint");

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, checkpointTimestamp);
        console.log("- submitting num checkpoint proofs", proofs.balanceProofs.length);

        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        cheats.resumeTracing();
    }

    function _verifyWithdrawalCredentials(uint40[] memory _validators) internal {
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
    function _revert(bytes memory err) internal pure {
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

    function _podCompoundingWithdrawalCredentials() internal view returns (bytes memory) {
        return abi.encodePacked(bytes1(uint8(2)), bytes11(0), address(pod));
    }

    function _getSlashingFactor(address staker, IStrategy strategy) internal view returns (uint) {
        address operator = config.delegationManager().delegatedTo(staker);
        uint64 maxMagnitude = config.allocationManager().getMaxMagnitudes(operator, strategy.toArray())[0];
        if (strategy == BEACONCHAIN_ETH_STRAT) return maxMagnitude.mulWad(config.eigenPodManager().beaconChainSlashingFactor(staker));
        return maxMagnitude;
    }

    /// @notice Gets the expected withdrawals to be created when the staker is undelegated via a call to `DelegationManager.undelegate()`
    /// @notice Assumes staker and withdrawer are the same and that all strategies and shares are withdrawn
    function _getExpectedWithdrawalStructsForStaker(address staker) internal view returns (Withdrawal[] memory expectedWithdrawals) {
        (IStrategy[] memory strategies,) = config.delegationManager().getDepositedShares(staker);

        expectedWithdrawals = new Withdrawal[](strategies.length);

        (, uint[] memory depositShares) = config.delegationManager().getWithdrawableShares(staker, strategies);

        address delegatedTo = config.delegationManager().delegatedTo(staker);
        uint nonce = config.delegationManager().cumulativeWithdrawalsQueued(staker);

        for (uint i = 0; i < strategies.length; ++i) {
            DepositScalingFactor memory dsf = DepositScalingFactor(config.delegationManager().depositScalingFactor(staker, strategies[i]));

            uint scaledShares = dsf.scaleForQueueWithdrawal(depositShares[i]);

            expectedWithdrawals[i] = Withdrawal({
                staker: staker,
                delegatedTo: delegatedTo,
                withdrawer: staker,
                nonce: (nonce + i),
                startBlock: uint32(block.number),
                strategies: strategies[i].toArray(),
                scaledShares: scaledShares.toArrayU256()
            });
        }
    }

    function getActiveValidators() public view returns (uint40[] memory) {
        uint40[] memory activeValidators = new uint40[](validators.length);

        uint numActive;
        uint pos;
        for (uint i = 0; i < validators.length; i++) {
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

    function _tryPrankAppointee(address target, bytes4 selector) internal {
        address[] memory appointees = config.permissionController().getAppointees(address(this), target, selector);
        if (appointees.length != 0) cheats.prank(appointees[0]);
    }

    function _tryPrankAppointee_AllocationManager(bytes4 selector) internal {
        return _tryPrankAppointee(address(config.allocationManager()), selector);
    }

    function _tryPrankAppointee_DelegationManager(bytes4 selector) internal {
        return _tryPrankAppointee(address(config.delegationManager()), selector);
    }
}

/// @notice A user contract that calls nonstandard methods (like xBySignature methods)
contract User_AltMethods is User {
    mapping(bytes32 => bool) public signedHashes;

    constructor(string memory name) User(name) {}

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
                underlyingToken.approve(address(config.strategyManager()), tokenBalance);

                // Get signature
                uint nonceBefore = config.strategyManager().nonces(address(this));
                bytes32 structHash = keccak256(
                    abi.encode(
                        config.strategyManager().DEPOSIT_TYPEHASH(),
                        address(this),
                        strat,
                        underlyingToken,
                        tokenBalance,
                        nonceBefore,
                        expiry
                    )
                );
                bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", config.strategyManager().domainSeparator(), structHash));
                bytes memory signature = bytes(abi.encodePacked(digestHash)); // dummy sig data

                // Mark hash as signed
                signedHashes[digestHash] = true;

                // Deposit
                config.strategyManager().depositIntoStrategyWithSignature(
                    strat, underlyingToken, tokenBalance, address(this), expiry, signature
                );

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
