// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPod.sol";

import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IStrategy.sol";

import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/mocks/BeaconChainMock.t.sol";
import "src/test/integration/utils/PrintUtils.t.sol";

struct Validator {
    uint40 index;
}

interface IUserDeployer {
    function delegationManager() external view returns (DelegationManager);
    function strategyManager() external view returns (StrategyManager);
    function eigenPodManager() external view returns (EigenPodManager);
    function timeMachine() external view returns (TimeMachine);
    function beaconChain() external view returns (BeaconChainMock);
}

contract User is PrintUtils {

    Vm cheats = Vm(HEVM_ADDRESS);

    DelegationManager delegationManager;
    StrategyManager strategyManager;
    EigenPodManager eigenPodManager;
    TimeMachine timeMachine;
    BeaconChainMock beaconChain;

    string _NAME;

    // User's EigenPod and each of their validator indices within that pod
    EigenPod public pod;
    uint40[] validators;

    IStrategy constant BEACONCHAIN_ETH_STRAT = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
    IERC20 constant NATIVE_ETH = IERC20(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
    uint constant GWEI_TO_WEI = 1e9;

    constructor(string memory name) {
        IUserDeployer deployer = IUserDeployer(msg.sender);

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

    /*******************************************************************************
                            DELEGATIONMANAGER METHODS
    *******************************************************************************/

    function registerAsOperator() public createSnapshot virtual {
        _logM("registerAsOperator");

        IDelegationManager.OperatorDetails memory details = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: address(this),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });

        delegationManager.registerAsOperator(details, "metadata");
    }

    /// @dev Delegate to the operator without a signature
    function delegateTo(User operator) public createSnapshot virtual {
        _logM("delegateTo", operator.NAME());

        ISignatureUtils.SignatureWithExpiry memory emptySig;
        delegationManager.delegateTo(address(operator), emptySig, bytes32(0));
    }

    /// @dev Undelegate from operator
    function undelegate() public createSnapshot virtual returns(IDelegationManager.Withdrawal[] memory){
        _logM("undelegate");

        IDelegationManager.Withdrawal[] memory expectedWithdrawals = _getExpectedWithdrawalStructsForStaker(address(this));
        delegationManager.undelegate(address(this));

        for (uint i = 0; i < expectedWithdrawals.length; i++) {
            emit log("expecting withdrawal:");
            emit log_named_uint("nonce: ", expectedWithdrawals[i].nonce);
            emit log_named_address("strat: ", address(expectedWithdrawals[i].strategies[0]));
            emit log_named_uint("shares: ", expectedWithdrawals[i].shares[0]);
        }
        
        return expectedWithdrawals;
    }

    /// @dev Force undelegate staker
    function forceUndelegate(User staker) public createSnapshot virtual returns(IDelegationManager.Withdrawal[] memory){
        _logM("forceUndelegate", staker.NAME());

        IDelegationManager.Withdrawal[] memory expectedWithdrawals = _getExpectedWithdrawalStructsForStaker(address(staker));
        delegationManager.undelegate(address(staker));
        return expectedWithdrawals;
    }

    /// @dev Queues a single withdrawal for every share and strategy pair
    function queueWithdrawals(
        IStrategy[] memory strategies, 
        uint[] memory shares
    ) public createSnapshot virtual returns (IDelegationManager.Withdrawal[] memory) {
        _logM("queueWithdrawals");

        address operator = delegationManager.delegatedTo(address(this));
        address withdrawer = address(this);
        uint nonce = delegationManager.cumulativeWithdrawalsQueued(address(this));

        // Create queueWithdrawals params
        IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        params[0] = IDelegationManager.QueuedWithdrawalParams({
            strategies: strategies,
            shares: shares,
            withdrawer: withdrawer
        });

        // Create Withdrawal struct using same info
        IDelegationManager.Withdrawal[] memory withdrawals = new IDelegationManager.Withdrawal[](1);
        withdrawals[0] = IDelegationManager.Withdrawal({
            staker: address(this),
            delegatedTo: operator,
            withdrawer: withdrawer,
            nonce: nonce,
            startBlock: uint32(block.number),
            strategies: strategies,
            shares: shares
        });

        bytes32[] memory withdrawalRoots = delegationManager.queueWithdrawals(params);

        // Basic sanity check - we do all other checks outside this file
        assertEq(withdrawals.length, withdrawalRoots.length, "User.queueWithdrawals: length mismatch");

        return (withdrawals);
    }

    function completeWithdrawalsAsTokens(IDelegationManager.Withdrawal[] memory withdrawals) public createSnapshot virtual returns (IERC20[][] memory) {
        _logM("completeWithdrawalsAsTokens");

        IERC20[][] memory tokens = new IERC20[][](withdrawals.length);

        for (uint i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal(withdrawals[i], true);
        }

        return tokens;
    }
    
    function completeWithdrawalAsTokens(IDelegationManager.Withdrawal memory withdrawal) public createSnapshot virtual returns (IERC20[] memory) {
        _logM("completeWithdrawalsAsTokens");

        return _completeQueuedWithdrawal(withdrawal, true);
    }

    function completeWithdrawalsAsShares(IDelegationManager.Withdrawal[] memory withdrawals) public createSnapshot virtual returns (IERC20[][] memory) {
        _logM("completeWithdrawalAsShares");
        
        IERC20[][] memory tokens = new IERC20[][](withdrawals.length);

        for (uint i = 0; i < withdrawals.length; i++) {
            tokens[i] = _completeQueuedWithdrawal(withdrawals[i], false);
        }

        return tokens;
    }

    function completeWithdrawalAsShares(IDelegationManager.Withdrawal memory withdrawal) public createSnapshot virtual returns (IERC20[] memory) {
        _logM("completeWithdrawalAsShares");

        return _completeQueuedWithdrawal(withdrawal, false);
    }

    /*******************************************************************************
                                BEACON CHAIN METHODS
    *******************************************************************************/

    /// @dev Uses any ETH held by the User to start validators on the beacon chain
    /// @return A list of created validator indices
    /// @return The amount of wei sent to the beacon chain
    /// Note: If the user does not have enough ETH to start a validator, this method reverts
    /// Note: This method also advances one epoch forward on the beacon chain, so that
    /// withdrawal credential proofs are generated for each validator.
    function startValidators() public createSnapshot virtual returns (uint40[] memory, uint64) {
        _logM("startValidators");

        return _startValidators();
    }

    function exitValidators(uint40[] memory _validators) public createSnapshot virtual returns (uint64 exitedBalanceGwei) {
        _logM("exitValidators");

        return _exitValidators(_validators);
    }

    /*******************************************************************************
                                 EIGENPOD METHODS
    *******************************************************************************/

    function verifyWithdrawalCredentials(
        uint40[] memory _validators
    ) public createSnapshot virtual {
        _logM("verifyWithdrawalCredentials");

        _verifyWithdrawalCredentials(_validators);
    }

    function startCheckpoint() public createSnapshot virtual {
        _logM("startCheckpoint");

        _startCheckpoint();
    }

    function completeCheckpoint() public createSnapshot virtual {
        _logM("completeCheckpoint");

        _completeCheckpoint();
    }

    function verifyStaleBalance(uint40 validatorIndex) public createSnapshot virtual {
        _logM("verifyStaleBalance");

        StaleBalanceProofs memory proof = beaconChain.getStaleBalanceProofs(validatorIndex);

        try pod.verifyStaleBalance({
            beaconTimestamp: proof.beaconTimestamp,
            stateRootProof: proof.stateRootProof,
            proof: proof.validatorProof
        }) { } catch (bytes memory err) {
            _revert(err);
        }
    }

    /*******************************************************************************
                                STRATEGY METHODS
    *******************************************************************************/

    /// @dev For each strategy/token balance, call the relevant deposit method
    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot virtual {
        _logM("depositIntoEigenlayer");

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                (uint40[] memory newValidators, ) = _startValidators();
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

    function updateBalances(IStrategy[] memory strategies, int[] memory tokenDeltas) public createSnapshot virtual {
        _logM("updateBalances");

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            int delta = tokenDeltas[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // If any balance update has occured, a checkpoint will pick it up
                _startCheckpoint();
                if (pod.activeValidatorCount() != 0) {
                    _completeCheckpoint();
                }
            } else {
                uint tokens = uint(delta);
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokens);
                strategyManager.depositIntoStrategy(strat, underlyingToken, tokens);
            }
        }
    }

    /*******************************************************************************
                                INTERNAL METHODS
    *******************************************************************************/

    function _completeQueuedWithdrawal(
        IDelegationManager.Withdrawal memory withdrawal, 
        bool receiveAsTokens
    ) internal virtual returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint i = 0; i < tokens.length; i++) {
            IStrategy strat = withdrawal.strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokens[i] = NATIVE_ETH;

                // If we're withdrawing native ETH as tokens, stop ALL validators
                // and complete a checkpoint
                if (receiveAsTokens) {
                    
                    _log("- exiting all validators and completing checkpoint");
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

        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0, receiveAsTokens);

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

        _log("- creating new validators", newValidators.length);
        _log("- depositing balance to beacon chain (gwei)", totalBeaconBalanceGwei);

        // Create each of the full validators
        for (uint i = 0; i < numValidators; i++) {
            uint40 validatorIndex = beaconChain.newValidator{ 
                value: 32 ether 
            }(_podWithdrawalCredentials());

            newValidators[i] = validatorIndex;
            validators.push(validatorIndex);
        }

        // If we had a remainder, create the final, non-full validator
        if (totalValidators == numValidators + 1) {
            uint40 validatorIndex = beaconChain.newValidator{ 
                value: lastValidatorBalance 
            }(_podWithdrawalCredentials());

            newValidators[newValidators.length - 1] = validatorIndex;
            validators.push(validatorIndex);
        }

        return (newValidators, totalBeaconBalanceGwei);
    }

    function _exitValidators(uint40[] memory _validators) internal returns (uint64 exitedBalanceGwei) {
        _log("- exiting num validators", _validators.length);

        for (uint i = 0; i < _validators.length; i++) {
            exitedBalanceGwei += beaconChain.exitValidator(_validators[i]);
        }

        _log("- exited balance to pod (gwei)", exitedBalanceGwei);

        return exitedBalanceGwei;
    }

    function _startCheckpoint() internal {
        try pod.startCheckpoint(false) {} catch (bytes memory err) {
            _revert(err);
        }
    }

    function _completeCheckpoint() internal {
        _log("- active validator count", pod.activeValidatorCount());
        _log("- proofs remaining", pod.currentCheckpoint().proofsRemaining);

        uint64 checkpointTimestamp = pod.currentCheckpointTimestamp();
        if (checkpointTimestamp == 0) {
            revert("User._completeCheckpoint: no existing checkpoint");
        }

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, checkpointTimestamp);
        _log("- submitting num checkpoint proofs", proofs.balanceProofs.length);

        pod.verifyCheckpointProofs({
            balanceContainerProof: proofs.balanceContainerProof,
            proofs: proofs.balanceProofs
        });
    }

    function _verifyWithdrawalCredentials(uint40[] memory _validators) internal {
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(_validators);

        try pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: _validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        }) { } catch (bytes memory err) {
            _revert(err);            
        }
    }

    /// @dev Revert, passing through an error message
    function _revert(bytes memory err) internal pure {
        if (err.length != 0) {
            assembly { revert(add(32, err), mload(err)) }
        }
        revert("reverted with unknown error");
    }

    function _podWithdrawalCredentials() internal view returns (bytes memory) {
        return abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(pod));
    }

    /// @notice Gets the expected withdrawals to be created when the staker is undelegated via a call to `DelegationManager.undelegate()`
    /// @notice Assumes staker and withdrawer are the same and that all strategies and shares are withdrawn
    function _getExpectedWithdrawalStructsForStaker(address staker) internal view returns (IDelegationManager.Withdrawal[] memory) {
        (IStrategy[] memory strategies, uint256[] memory shares)
            = delegationManager.getDelegatableShares(staker);

        IDelegationManager.Withdrawal[] memory expectedWithdrawals = new IDelegationManager.Withdrawal[](strategies.length);
        address delegatedTo = delegationManager.delegatedTo(staker);
        uint256 nonce = delegationManager.cumulativeWithdrawalsQueued(staker);
        
        for (uint256 i = 0; i < strategies.length; ++i) {
            IStrategy[] memory singleStrategy = new IStrategy[](1);
            uint256[] memory singleShares = new uint256[](1);
            singleStrategy[0] = strategies[i];
            singleShares[0] = shares[i];
            expectedWithdrawals[i] = IDelegationManager.Withdrawal({
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

    function getActiveValidators() public view returns (uint40[] memory) {
        uint40[] memory activeValidators = new uint40[](validators.length);

        uint numActive;
        uint pos;
        for(uint i = 0; i < validators.length; i++) {
            if (beaconChain.isActive(validators[i])) {
                activeValidators[pos] = validators[i];
                numActive++;
                pos++;
            }
        }

        // Manually update length
        assembly { mstore(activeValidators, numActive) }

        return activeValidators;
    }
}

/// @notice A user contract that calls nonstandard methods (like xBySignature methods)
contract User_AltMethods is User {

    mapping(bytes32 => bool) public signedHashes;

    constructor(string memory name) User(name) {}

    function delegateTo(User operator) public createSnapshot override {
        _logM("delegateTo_ALT", operator.NAME());

        // Create empty data
        ISignatureUtils.SignatureWithExpiry memory emptySig;
        uint256 expiry = type(uint256).max;

        // Get signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry;
        stakerSignatureAndExpiry.expiry = expiry;
        bytes32 digestHash = delegationManager.calculateCurrentStakerDelegationDigestHash(address(this), address(operator), expiry);
        stakerSignatureAndExpiry.signature = bytes(abi.encodePacked(digestHash)); // dummy sig data

        // Mark hash as signed
        signedHashes[digestHash] = true;

        // Delegate
        delegationManager.delegateToBySignature(address(this), address(operator), stakerSignatureAndExpiry, emptySig, bytes32(0));

        // Mark hash as used
        signedHashes[digestHash] = false;
    }

    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot override {
        _logM("depositIntoEigenlayer_ALT");
        
        uint256 expiry = type(uint256).max;
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                (uint40[] memory newValidators, ) = _startValidators();
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
                    abi.encode(strategyManager.DEPOSIT_TYPEHASH(), address(this), strat, underlyingToken, tokenBalance, nonceBefore, expiry)
                );
                bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));
                bytes memory signature = bytes(abi.encodePacked(digestHash)); // dummy sig data

                // Mark hash as signed
                signedHashes[digestHash] = true;

                // Deposit
                strategyManager.depositIntoStrategyWithSignature(
                    strat,
                    underlyingToken,
                    tokenBalance,
                    address(this),
                    expiry,
                    signature
                );

                // Mark hash as used
                signedHashes[digestHash] = false;
            }
        }
    }
 
    bytes4 internal constant MAGIC_VALUE = 0x1626ba7e;
    function isValidSignature(bytes32 hash, bytes memory) external view returns (bytes4) {
        if(signedHashes[hash]){
            return MAGIC_VALUE;
        } else {
            return 0xffffffff;
        }
    } 
}
