// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPod.sol";

import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IStrategy.sol";

import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/mocks/BeaconChainMock.t.sol";

interface IUserDeployer {
    function delegationManager() external view returns (DelegationManager);
    function strategyManager() external view returns (StrategyManager);
    function eigenPodManager() external view returns (EigenPodManager);
    function timeMachine() external view returns (TimeMachine);
    function beaconChain() external view returns (BeaconChainMock);
}

contract User is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    DelegationManager delegationManager;
    StrategyManager strategyManager;
    EigenPodManager eigenPodManager;

    TimeMachine timeMachine;

    /// @dev Native restaker state vars

    BeaconChainMock beaconChain;
    // User's EigenPod and each of their validator indices within that pod
    EigenPod pod;
    uint40[] validators;

    IStrategy constant BEACONCHAIN_ETH_STRAT = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
    IERC20 constant NATIVE_ETH = IERC20(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
    uint constant GWEI_TO_WEI = 1e9;

    constructor() {
        IUserDeployer deployer = IUserDeployer(msg.sender);

        delegationManager = deployer.delegationManager();
        strategyManager = deployer.strategyManager();
        eigenPodManager = deployer.eigenPodManager();
        timeMachine = deployer.timeMachine();
                
        beaconChain = deployer.beaconChain();
        pod = EigenPod(payable(eigenPodManager.createPod()));
    }

    modifier createSnapshot() virtual {
        timeMachine.createSnapshot();
        _;
    }

    receive() external payable {}

    /**
     * DelegationManager methods:
     */

    function registerAsOperator() public createSnapshot virtual {
        IDelegationManager.OperatorDetails memory details = IDelegationManager.OperatorDetails({
            earningsReceiver: address(this),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });

        delegationManager.registerAsOperator(details, "metadata");
    }

    /// @dev For each strategy/token balance, call the relevant deposit method
    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot virtual {

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // We're depositing via `eigenPodManager.stake`, which only accepts
                // deposits of exactly 32 ether.
                require(tokenBalance % 32 ether == 0, "User.depositIntoEigenlayer: balance must be multiple of 32 eth");
                
                // For each multiple of 32 ether, deploy a new validator to the same pod
                uint numValidators = tokenBalance / 32 ether;
                for (uint j = 0; j < numValidators; j++) {
                    eigenPodManager.stake{ value: 32 ether }("", "", bytes32(0));

                    (uint40 newValidatorIndex, CredentialsProofs memory proofs) = 
                        beaconChain.newValidator({
                            balanceWei: 32 ether,
                            withdrawalCreds: _podWithdrawalCredentials()
                        });
                    
                    validators.push(newValidatorIndex);

                    pod.verifyWithdrawalCredentials({
                        oracleTimestamp: proofs.oracleTimestamp,
                        stateRootProof: proofs.stateRootProof,
                        validatorIndices: proofs.validatorIndices,
                        validatorFieldsProofs: proofs.validatorFieldsProofs,
                        validatorFields: proofs.validatorFields
                    });
                }
            } else {
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokenBalance);
                strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
            }
        }
    }

    /// @dev Delegate to the operator without a signature
    function delegateTo(User operator) public createSnapshot virtual {
        ISignatureUtils.SignatureWithExpiry memory emptySig;
        delegationManager.delegateTo(address(operator), emptySig, bytes32(0));
    }

    /// @dev Queues a single withdrawal for every share and strategy pair
    function queueWithdrawals(
        IStrategy[] memory strategies, 
        uint[] memory shares
    ) public createSnapshot virtual returns (IDelegationManager.Withdrawal[] memory, bytes32[] memory) {

        address operator = delegationManager.delegatedTo(address(this));
        address withdrawer = address(this);
        uint nonce = delegationManager.cumulativeWithdrawalsQueued(address(this));
        
        bytes32[] memory withdrawalRoots;

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

        withdrawalRoots = delegationManager.queueWithdrawals(params);

        // Basic sanity check - we do all other checks outside this file
        assertEq(withdrawals.length, withdrawalRoots.length, "User.queueWithdrawals: length mismatch");

        return (withdrawals, withdrawalRoots);
    }

    function completeQueuedWithdrawal(
        IDelegationManager.Withdrawal memory withdrawal, 
        bool receiveAsTokens
    ) public createSnapshot virtual returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint i = 0; i < tokens.length; i++) {
            IStrategy strat = withdrawal.strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokens[i] = NATIVE_ETH;

                // If we're withdrawing as tokens, we need to process a withdrawal proof first
                if (receiveAsTokens) {
                    
                    emit log("exiting validators and processing withdrawals...");
                    
                    uint numValidators = validators.length;
                    for (uint j = 0; j < numValidators; j++) {
                        emit log_named_uint("exiting validator ", j);

                        uint40 validatorIndex = validators[j];
                        BeaconWithdrawal memory proofs = beaconChain.exitValidator(validatorIndex);

                        uint64 withdrawableBefore = pod.withdrawableRestakedExecutionLayerGwei();

                        pod.verifyAndProcessWithdrawals({
                            oracleTimestamp: proofs.oracleTimestamp,
                            stateRootProof: proofs.stateRootProof,
                            withdrawalProofs: proofs.withdrawalProofs,
                            validatorFieldsProofs: proofs.validatorFieldsProofs,
                            validatorFields: proofs.validatorFields,
                            withdrawalFields: proofs.withdrawalFields
                        });

                        uint64 withdrawableAfter = pod.withdrawableRestakedExecutionLayerGwei();

                        emit log_named_uint("pod withdrawable before: ", withdrawableBefore);
                        emit log_named_uint("pod withdrawable after: ", withdrawableAfter);
                    }
                }
            } else {
                tokens[i] = strat.underlyingToken();
            }
        }

        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0, receiveAsTokens);

        return tokens;
    }

    function _podWithdrawalCredentials() internal view returns (bytes memory) {
        return abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(pod));
    }
}

/// @notice A user contract that calls nonstandard methods (like xBySignature methods)
contract User_AltMethods is User {

    mapping(bytes32 => bool) public signedHashes;

    constructor() User() {}

    function delegateTo(User operator) public createSnapshot override {
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
        uint256 expiry = type(uint256).max;
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // We're depositing via `eigenPodManager.stake`, which only accepts
                // deposits of exactly 32 ether.
                require(tokenBalance % 32 ether == 0, "User.depositIntoEigenlayer: balance must be multiple of 32 eth");
                
                // For each multiple of 32 ether, deploy a new validator to the same pod
                uint numValidators = tokenBalance / 32 ether;
                for (uint j = 0; j < numValidators; j++) {
                    eigenPodManager.stake{ value: 32 ether }("", "", bytes32(0));

                    (uint40 newValidatorIndex, CredentialsProofs memory proofs) = 
                        beaconChain.newValidator({
                            balanceWei: 32 ether,
                            withdrawalCreds: _podWithdrawalCredentials()
                        });
                    
                    validators.push(newValidatorIndex);

                    pod.verifyWithdrawalCredentials({
                        oracleTimestamp: proofs.oracleTimestamp,
                        stateRootProof: proofs.stateRootProof,
                        validatorIndices: proofs.validatorIndices,
                        validatorFieldsProofs: proofs.validatorFieldsProofs,
                        validatorFields: proofs.validatorFields
                    });
                }
            } else {
                // Approve token
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokenBalance);

                // Get signature
                uint256 nonceBefore = strategyManager.nonces(address(this));
                bytes32 structHash = keccak256(
                    abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strat, underlyingToken, tokenBalance, nonceBefore, expiry)
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