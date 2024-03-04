// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IEigenPod.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IEigenPodManager.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IStrategyManager.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IDelayedWithdrawalRouter.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";

contract IntegrationMainnetFork is IntegrationCheckUtils {
    using Strings for *;

    // M1 contracts with deprecated interfaces
    IStrategyManager_DeprecatedM1 strategyManager_M1;
    IEigenPodManager_DeprecatedM1 eigenPodManager_M1;
    IEigenPod_DeprecatedM1 eigenPod_M1;
    IDelayedWithdrawalRouter_DeprecatedM1 delayedWithdrawalRouter_M1;

    // Select mainnet fork id
    function setUp() public virtual override {
        // create mainnet fork that can be used later
        mainnetForkId = cheats.createSelectFork(cheats.rpcUrl("mainnet"));

        string memory deploymentInfoPath = "script/configs/mainnet/Mainnet_current_deploy.config.json";
        _parseDeployedContracts(deploymentInfoPath);
        _parseDeployedContractsAsM1();
    }

    /**
     * @dev Create a new user according to configured random variants.
     * This user is ready to deposit into some strategies and has some underlying token balances
     */
    function _newRandomStaker_M1() internal returns (User_M1, IStrategy[] memory, uint[] memory) {
        string memory stakerName = string.concat("- Staker", numStakers.toString());
        numStakers++;

        (User_M1 staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _randUser_M1(stakerName);
        
        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "_newRandomStaker: failed to award token balances");

        return (staker, strategies, tokenBalances);
    }

    function _newRandomOperator_M1() internal returns (User_M1, IStrategy[] memory, uint[] memory) {
        string memory operatorName = string.concat("- Operator", numOperators.toString());
        numOperators++;

        (User_M1 operator, IStrategy[] memory strategies, uint[] memory tokenBalances) = _randUser_M1(operatorName);
        
        operator.registerAsOperator();
        operator.depositIntoEigenlayer(strategies, tokenBalances);

        assert_Snap_Added_StakerShares(operator, strategies, tokenBalances, "_newRandomOperator: failed to add delegatable shares");
        assert_Snap_Added_OperatorShares(operator, strategies, tokenBalances, "_newRandomOperator: failed to award shares to operator");
        assertTrue(delegationManager.isOperator(address(operator)), "_newRandomOperator: operator should be registered");

        return (operator, strategies, tokenBalances);
    }

    function _randUser_M1(string memory name) internal returns (User_M1, IStrategy[] memory, uint[] memory) {
        // For the new user, select what type of assets they'll have and whether
        // they'll use `xWithSignature` methods.
        //
        // The values selected here are in the ranges configured via `_configRand`
        uint assetType = _randAssetType();
        uint userType = _randUserType();
        
        // Create User contract based on deposit type:
        User_M1 user = new User_M1(name);
        // if (userType == DEFAULT) {
        //     user = new User(name);
        // } else if (userType == ALT_METHODS) {
        //     // User will use nonstandard methods like:
        //     // `delegateToBySignature` and `depositIntoStrategyWithSignature`
        //     user = User(new User_AltMethods(name));
        // } else {
        //     revert("_randUser: unimplemented userType");
        // }

        // For the specific asset selection we made, get a random assortment of
        // strategies and deal the user some corresponding underlying token balances
        (IStrategy[] memory strategies, uint[] memory tokenBalances) = _dealRandAssets(user, assetType);

        _printUserInfo(assetType, userType, strategies, tokenBalances);

        return (user, strategies, tokenBalances);
    }


    /**
     * @notice Create contract instances as M1 contracts, assumes 
     * ExistingDeploymentParse._parseDeployedContracts has already parsed existing contracts on mainnet
     */
    function _parseDeployedContractsAsM1() internal {
        strategyManager_M1 = IStrategyManager_DeprecatedM1(address(strategyManager));
        eigenPodManager_M1 = IEigenPodManager_DeprecatedM1(address(eigenPodManager));
        eigenPod_M1 = IEigenPod_DeprecatedM1(address(eigenPodImplementation));
        delayedWithdrawalRouter_M1 = IDelayedWithdrawalRouter_DeprecatedM1(address(delayedWithdrawalRouter));
    }

    // function _randUser
}
