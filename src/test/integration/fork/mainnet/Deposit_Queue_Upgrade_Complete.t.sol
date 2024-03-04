// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/fork/mainnet/IntegrationMainnetFork.t.sol";
import "src/test/integration/fork/mainnet/deprecatedInterfaces/IEigenPod.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";


contract IntegrationMainnetFork_Deposit_Queue_Upgrade_Complete is IntegrationMainnetFork {

    // Select mainnet fork id
    function setUp() public override {
        // create mainnet fork that can be used later
        mainnetForkId = cheats.createSelectFork(cheats.rpcUrl("mainnet"));

        string memory deploymentInfoPath = "script/configs/mainnet/Mainnet_current_deploy.config.json";
        _parseDeployedContracts(deploymentInfoPath);
        _upgradeMainnetContracts();
    }

    /// @notice Test migrated withdrawals
    function test_deposit_queue_upgrade_complete(uint24 _random) public {

        // // Create a staker with a nonzero balance and corresponding strategies
        // (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        // // 1. deposit strategy
        // staker.depositIntoEigenlayer(strategies, tokenBalances);

        // // 2. queue withdrawal through StrategyManager
        // (bytes32 withdrawalRoot, IStrategyManager_DeprecatedM1.QueuedWithdrawal memory withdrawal) = staker.queueWithdrawals_M1(strategies, tokenBalances);

        // // 3. upgrade mainnet contracts
        // _upgradeMainnetContracts();

        // // 4. complete withdrawal through strategyManager calling migrateQueuedWithdrawals
        // staker.completeWithdrawals_M1(withdrawal, true);
    }

}
