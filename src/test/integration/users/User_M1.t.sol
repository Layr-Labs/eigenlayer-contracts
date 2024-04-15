// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/deprecatedInterfaces/mainnet/IEigenPod.sol";
import "src/test/integration/deprecatedInterfaces/mainnet/IEigenPodManager.sol";
import "src/test/integration/deprecatedInterfaces/mainnet/IStrategyManager.sol";
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

    constructor(string memory name) User(name) {
        IUserMainnetForkDeployer deployer = IUserMainnetForkDeployer(msg.sender);

        strategyManager_M1 = IStrategyManager_DeprecatedM1(address(deployer.strategyManager()));
        eigenPodManager_M1 = IEigenPodManager_DeprecatedM1(address(deployer.eigenPodManager()));
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
        _logM("depositIntoEigenlayer_M1");

        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint256 tokenBalance = tokenBalances[i];

            // Skip BeaconChainStrat, since BeaconChainStrat doesn't exist on M1 mainnet
            if (strat == BEACONCHAIN_ETH_STRAT) {
                continue;
            }

            IERC20 underlyingToken = strat.underlyingToken();
            underlyingToken.approve(address(strategyManager), tokenBalance);
            strategyManager.depositIntoStrategy(strat, underlyingToken, tokenBalance);
        }
    }

    function _createPod() internal virtual override {
        IEigenPodManager_DeprecatedM1(address(eigenPodManager)).createPod();
        // get EigenPod address
        pod = EigenPod(
            payable(address(IEigenPodManager_DeprecatedM1(address(eigenPodManager)).ownerToPod(address(this))))
        );
    }
}

contract User_M1_AltMethods is User_M1 {
    mapping(bytes32 => bool) public signedHashes;

    constructor(string memory name) User_M1(name) {}

    function depositIntoEigenlayer_M1(
        IStrategy[] memory strategies,
        uint256[] memory tokenBalances
    ) public override createSnapshot {
        _logM(".depositIntoEigenlayer_M1_ALT");

        uint256 expiry = type(uint256).max;
        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint256 tokenBalance = tokenBalances[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                revert("Should not be depositing with BEACONCHAIN_ETH_STRAT for M1-mainnet User");
            }
            
            // Approve token
            IERC20 underlyingToken = strat.underlyingToken();
            underlyingToken.approve(address(strategyManager), tokenBalance);

            // Get signature
            uint256 nonceBefore = strategyManager.nonces(address(this));
            bytes32 structHash = keccak256(
                abi.encode(
                    strategyManager.DEPOSIT_TYPEHASH(),
                    strat,
                    underlyingToken,
                    tokenBalance,
                    nonceBefore,
                    expiry
                )
            );
            bytes32 domain_separator = keccak256(abi.encode(strategyManager.DOMAIN_TYPEHASH(), keccak256(bytes("EigenLayer")), block.chainid, address(strategyManager)));
            bytes32 digestHash =
                keccak256(abi.encodePacked("\x19\x01", domain_separator, structHash));
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

    bytes4 internal constant MAGIC_VALUE = 0x1626ba7e;

    function isValidSignature(bytes32 hash, bytes memory) external view returns (bytes4) {
        if (signedHashes[hash]) {
            return MAGIC_VALUE;
        } else {
            return 0xffffffff;
        }
    }
}
