// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/fork/mainnet/deprecatedInterfaces/IEigenPod.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IEigenPodManager.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IStrategyManager.sol";
import "src/test/integration/users/User.t.sol";

interface IUserMainnetForkDeployer {
    function delegationManager() external view returns (DelegationManager);
    function strategyManager() external view returns (StrategyManager);
    function eigenPodManager() external view returns (EigenPodManager);
    function timeMachine() external view returns (TimeMachine);
    function beaconChain() external view returns (BeaconChainMock);
    function strategyManager_M1() external view returns (IStrategyManager_DeprecatedM1);
    function eigenPodManager_M1() external view returns (IEigenPodManager_DeprecatedM1);
}

/**
 * @dev User_M1 used for performing mainnet M1 contract methods but also inherits User
 * to perform current local contract methods after a upgrade of core contracts
 */
contract User_M1 is User {
    IStrategyManager_DeprecatedM1 strategyManager_M1;
    IEigenPodManager_DeprecatedM1 eigenPodManager_M1;
    IEigenPod_DeprecatedM1 pod_M1;

    constructor(string memory name) User(name) {
        IUserMainnetForkDeployer deployer = IUserMainnetForkDeployer(msg.sender);

        strategyManager_M1 = deployer.strategyManager_M1();
        eigenPodManager_M1 = deployer.eigenPodManager_M1();
        // TODO: Use existing pod from mainnet??
        // pod_M1 = EigenPod(payable(address(eigenPodManager_M1.createPod())));
    }


    /**
     * StrategyManager M1 mainnet methods:
     */

    /// @notice Deposit into EigenLayer with M1 mainnet methods, only concerns LSTs
    /// Note that this should not be called with BeaconChainStrat
    function depositIntoEigenlayer_M1(
        IStrategy[] memory strategies,
        uint256[] memory tokenBalances
    ) public virtual createSnapshot {
        emit log(_name(".depositIntoEigenlayer_M1"));

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint tokenBalance = tokenBalances[i];

            // Skip BeaconChainStrat, since BeaconChainStrat doesn't exist on M1 mainnet
            if (strat == BEACONCHAIN_ETH_STRAT) {
                continue;
            }

            IERC20 underlyingToken = strat.underlyingToken();
            underlyingToken.approve(address(strategyManager), tokenBalance);
            strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
        }
    }
}

contract User_M1_AltMethods is User_M1 {
    mapping(bytes32 => bool) public signedHashes;

    constructor(string memory name) User_M1("User_M1_AltMethods") {}

    function depositIntoEigenlayer_M1(IStrategy[] memory strategies, uint[] memory tokenBalances) public createSnapshot override {
        emit log(_name(".depositIntoEigenlayer"));
        
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
