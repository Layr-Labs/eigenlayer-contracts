// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/pods/EigenPodManager.sol";

import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IStrategy.sol";

import "src/test/integration/Global.t.sol";

contract User is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    DelegationManager delegationManager;
    StrategyManager strategyManager;
    EigenPodManager eigenPodManager;

    IStrategy[] depositedStrategies;

    Global global;

    IStrategy constant BEACONCHAIN_ETH_STRAT = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    constructor(
        DelegationManager _delegationManager,
        StrategyManager _strategyManager,
        EigenPodManager _eigenPodManager,
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
                // TODO handle this flow - need to deposit into EPM + prove credentials
                revert("depositIntoEigenlayer: unimplemented");
            } else {
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokenBalance);
                strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
            }
            depositedStrategies.push(strat);
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
                tokens[i] = IERC20(address(0));
            } else {
                tokens[i] = strat.underlyingToken();
            }
        }

        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0, receiveAsTokens);

        return tokens;
    }
}

contract User_SignedMethods is User {

    uint256 immutable privateKey;
    address immutable userAddress;

    constructor(
        DelegationManager _delegationManager,
        StrategyManager _strategyManager,
        EigenPodManager _eigenPodManager,
        Global _global
    ) User(_delegationManager, _strategyManager, _eigenPodManager, _global) {
        // Generate private key from contract address
        privateKey = uint256(keccak256(abi.encodePacked(block.timestamp, block.difficulty, address(this))));
        userAddress = cheats.addr(privateKey);
    }

    function delegateTo(User operator) public createSnapshot override {
        // Create signatures
        ISignatureUtils.SignatureWithExpiry memory emptySig;
        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getDelegateToSignature(address(operator), expiry);

        // Delegate
        delegationManager.delegateToBySignature(userAddress, address(operator), stakerSignatureAndExpiry, emptySig, bytes32(0));
    }

    function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot override {
        uint256 expiry = type(uint256).max;
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // TODO handle this flow - need to deposit into EPM + prove credentials
                revert("depositIntoEigenlayer: unimplemented");
            } else {
                IERC20 underlyingToken = strat.underlyingToken();
                underlyingToken.approve(address(strategyManager), tokenBalance);
                bytes memory signature = _getStrategyDepositSignature(strat, underlyingToken, tokenBalance, expiry);
                strategyManager.depositIntoStrategyWithSignature(
                    strat,
                    underlyingToken,
                    tokenBalance,
                    userAddress,
                    expiry,
                    signature
                );
            }
            depositedStrategies.push(strat);
        }
    }

    function _getStrategyDepositSignature(IStrategy strategy, IERC20 token, uint256 amount, uint256 expiry) internal returns (bytes) {
        uint256 nonceBefore = strategyManager.nonces(userAddress);
        bytes memory signature;

        bytes32 structHash = keccak256(
            abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry)
        );
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

        signature = abi.encodePacked(r, s, v);

        return signature;
    }
 

    function _getDelegateToSignature(address operator, uint256 expiry)
        internal view returns (ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry)
    {
        stakerSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateCurrentStakerDelegationDigestHash(userAddress, operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);
            stakerSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }
        return stakerSignatureAndExpiry;
    }
}

// contract User_MixedAssets is User {
//     /// Since this is the base setup, we don't need anything else
// }

// contract User_LSTOnly is User {

//     /// @dev For each strategy/token balance, call the relevant deposit method
//     function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot override {

//         for (uint i = 0; i < strategies.length; i++) {
//             IStrategy strat = strategies[i];
//             uint tokenBalance = tokenBalances[i];
//             IERC20 underlyingToken = strat.underlyingToken();

//             strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
//         }
//     }
// }

// contract User_NativeOnly is User {

//     /// @dev For each strategy/token balance, call the relevant deposit method
//     function depositIntoEigenlayer(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot override {

//         for (uint i = 0; i < strategies.length; i++) {
//             IStrategy strat = strategies[i];
//             uint tokenBalance = tokenBalances[i];
            
//             // TODO handle this flow - need to deposit into EPM + prove credentials
//             revert("TODO: unimplemented");
//         }
//     }
// }