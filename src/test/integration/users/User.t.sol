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
    function startValidators() public createSnapshot virtual returns (uint40[] memory, uint) {
        _logM("startValidators");

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
        uint totalBeaconBalance = address(this).balance - balanceWei;

        _log("- creating new validators", newValidators.length);
        _log("- depositing balance to beacon chain (wei)", totalBeaconBalance);

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

        // Advance forward one epoch and generate withdrawal and balance proofs for each validator
        beaconChain.advanceEpoch_NoRewards();

        return (newValidators, totalBeaconBalance);
    }

    function exitValidators(uint40[] memory _validators) public createSnapshot virtual returns (uint64 exitedBalanceGwei) {
        _logM("exitValidators");

        _log("- exiting num validators", _validators.length);

        for (uint i = 0; i < _validators.length; i++) {
            exitedBalanceGwei += beaconChain.exitValidator(_validators[i]);
        }

        _log("- exited balance to pod (gwei)", exitedBalanceGwei);

        return exitedBalanceGwei;
    }

    /*******************************************************************************
                                 EIGENPOD METHODS
    *******************************************************************************/

    function verifyWithdrawalCredentials(
        uint40[] memory _validators
    ) public createSnapshot virtual {
        _logM("verifyWithdrawalCredentials");

        CredentialProofs memory proofs = beaconChain.getCredentialProofs(_validators);

        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: _validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
    }

    function startCheckpoint() public createSnapshot virtual {
        _logM("startCheckpoint");

        pod.startCheckpoint(false);
    }

    function completeCheckpoint() public createSnapshot virtual {
        _logM("completeCheckpoint");

        _log("- active validator count", pod.activeValidatorCount());
        _log("- proofs remaining", pod.currentCheckpoint().proofsRemaining);

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators);
        _log("- submitting num checkpoint proofs", proofs.balanceProofs.length);

        pod.verifyCheckpointProofs({
            balanceContainerProof: proofs.balanceContainerProof,
            proofs: proofs.balanceProofs
        });
    }

    /*******************************************************************************
                              STRATEGYMANAGER METHODS
    *******************************************************************************/

    function depositLSTs(IStrategy[] memory strategies) public createSnapshot virtual {
        // _log("depositIntoStrategies");

        // for (uint i = 0; i < strategies.length; i++) {
        //     IStrategy strat = strategies[i];
        //     uint tokenBalance = tokenBalances[i];

        //     IERC20 underlyingToken = strat.underlyingToken();
        //     underlyingToken.approve(address(strategyManager), tokenBalance);
        //     strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
        // }
    }

    /// @dev For each strategy/token balance, call the relevant deposit method
    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot virtual {
        _logM("depositIntoEigenlayer");

        revert("unimplemented");
    }

    function updateBalances(IStrategy[] memory strategies, int[] memory tokenDeltas) public createSnapshot virtual {
        _logM("updateBalances");
        revert("fail - placeholder");

        // for (uint i = 0; i < strategies.length; i++) {
        //     IStrategy strat = strategies[i];
        //     int delta = tokenDeltas[i];

        //     if (strat == BEACONCHAIN_ETH_STRAT) {
        //         // TODO - right now, we just grab the first validator
        //         uint40 validator = getUpdatableValidator();
        //         BalanceUpdate memory update = beaconChain.updateBalance(validator, delta);

        //         int sharesBefore = eigenPodManager.podOwnerShares(address(this));

        //         pod.verifyBalanceUpdates({
        //             oracleTimestamp: update.oracleTimestamp,
        //             validatorIndices: update.validatorIndices,
        //             stateRootProof: update.stateRootProof,
        //             validatorFieldsProofs: update.validatorFieldsProofs,
        //             validatorFields: update.validatorFields
        //         });

        //         int sharesAfter = eigenPodManager.podOwnerShares(address(this));

        //         emit log_named_int("pod owner shares before: ", sharesBefore);
        //         emit log_named_int("pod owner shares after: ", sharesAfter);
        //     } else {
        //         uint tokens = uint(delta);
        //         IERC20 underlyingToken = strat.underlyingToken();
        //         underlyingToken.approve(address(strategyManager), tokens);
        //         strategyManager.depositIntoStrategy(strat, underlyingToken, tokens);
        //     }
        // }
    }

    function _completeQueuedWithdrawal(
        IDelegationManager.Withdrawal memory withdrawal, 
        bool receiveAsTokens
    ) internal virtual returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint i = 0; i < tokens.length; i++) {
            IStrategy strat = withdrawal.strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokens[i] = NATIVE_ETH;

                // If we're withdrawing as tokens, we need to process a withdrawal proof first
                if (receiveAsTokens) {
                    
                    emit log("exiting validators and processing withdrawals...");
                    revert("fail - placeholder");

                    // uint numValidators = validators.length;
                    // for (uint j = 0; j < numValidators; j++) {
                    //     emit log_named_uint("exiting validator ", j);

                    //     uint40 validatorIndex = validators[j];
                    //     BeaconWithdrawal memory proofs = beaconChain.exitValidator(validatorIndex);

                    //     uint64 withdrawableBefore = pod.withdrawableRestakedExecutionLayerGwei();

                    //     pod.verifyAndProcessWithdrawals({
                    //         oracleTimestamp: proofs.oracleTimestamp,
                    //         stateRootProof: proofs.stateRootProof,
                    //         withdrawalProofs: proofs.withdrawalProofs,
                    //         validatorFieldsProofs: proofs.validatorFieldsProofs,
                    //         validatorFields: proofs.validatorFields,
                    //         withdrawalFields: proofs.withdrawalFields
                    //     });

                    //     uint64 withdrawableAfter = pod.withdrawableRestakedExecutionLayerGwei();

                    //     emit log_named_uint("pod withdrawable before: ", withdrawableBefore);
                    //     emit log_named_uint("pod withdrawable after: ", withdrawableAfter);
                    // }
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

    function getUpdatableValidator() public view returns (uint40) {
        return validators[0];
    }
}

/// @notice A user contract that calls nonstandard methods (like xBySignature methods)
contract User_AltMethods is User {

    mapping(bytes32 => bool) public signedHashes;

    constructor(string memory name) User(name) {}

    function delegateTo(User operator) public createSnapshot override {
        _logM("delegateTo", operator.NAME());

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
        // emit log(_name(".depositIntoEigenlayer"));
        
        // uint256 expiry = type(uint256).max;
        // for (uint i = 0; i < strategies.length; i++) {
        //     IStrategy strat = strategies[i];
        //     uint tokenBalance = tokenBalances[i];

        //     if (strat == BEACONCHAIN_ETH_STRAT) {
        //         // We're depositing via `eigenPodManager.stake`, which only accepts
        //         // deposits of exactly 32 ether.
        //         require(tokenBalance % 32 ether == 0, "User.depositIntoEigenlayer: balance must be multiple of 32 eth");
                
        //         // For each multiple of 32 ether, deploy a new validator to the same pod
        //         uint numValidators = tokenBalance / 32 ether;
        //         for (uint j = 0; j < numValidators; j++) {
        //             eigenPodManager.stake{ value: 32 ether }("", "", bytes32(0));

        //             (uint40 newValidatorIndex, CredentialsProofs memory proofs) = 
        //                 beaconChain.newValidator({
        //                     balanceWei: 32 ether,
        //                     withdrawalCreds: _podWithdrawalCredentials()
        //                 });
                    
        //             validators.push(newValidatorIndex);

        //             pod.verifyWithdrawalCredentials({
        //                 beaconTimestamp: proofs.oracleTimestamp,
        //                 stateRootProof: proofs.stateRootProof,
        //                 validatorIndices: proofs.validatorIndices,
        //                 validatorFieldsProofs: proofs.validatorFieldsProofs,
        //                 validatorFields: proofs.validatorFields
        //             });
        //         }
        //     } else {
        //         // Approve token
        //         IERC20 underlyingToken = strat.underlyingToken();
        //         underlyingToken.approve(address(strategyManager), tokenBalance);

        //         // Get signature
        //         uint256 nonceBefore = strategyManager.nonces(address(this));
        //         bytes32 structHash = keccak256(
        //             abi.encode(strategyManager.DEPOSIT_TYPEHASH(), address(this), strat, underlyingToken, tokenBalance, nonceBefore, expiry)
        //         );
        //         bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));
        //         bytes memory signature = bytes(abi.encodePacked(digestHash)); // dummy sig data

        //         // Mark hash as signed
        //         signedHashes[digestHash] = true;

        //         // Deposit
        //         strategyManager.depositIntoStrategyWithSignature(
        //             strat,
        //             underlyingToken,
        //             tokenBalance,
        //             address(this),
        //             expiry,
        //             signature
        //         );

        //         // Mark hash as used
        //         signedHashes[digestHash] = false;
        //     }
        // }
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
