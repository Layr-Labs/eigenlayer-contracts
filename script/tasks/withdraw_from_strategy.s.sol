// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AllocationManager.sol";
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/libraries/SlashingLib.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/withdraw_from_strategy.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address strategy,address token,uint256 amount)" -- <DEPLOYMENT_OUTPUT_JSON> <STRATEGY_ADDRESS> <TOKEN_ADDRESS> <AMOUNT>
// RUST_LOG=forge,foundry=trace forge script script/tasks/withdraw_from_strategy.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address strategy,address token,uint256 amount)" -- local/slashing_output.json 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0x67d269191c92Caf3cD7723F116c85e6E9bf55933 750
contract WithdrawFromStrategy is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(string memory configFile, address strategy, address token, uint256 amount) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);

        // Pull addresses from config
        // address allocationManager = stdJson.readAddress(config_data, ".addresses.allocationManager");
        address delegationManager = stdJson.readAddress(config_data, ".addresses.delegationManager");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        // Attach to DelegationManager
        // AllocationManager am = AllocationManager(allocationManager);
        DelegationManager dm = DelegationManager(delegationManager);

        // Add strategy to array
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(strategy);
        // Add shares to array
        uint256[] memory shares = new uint256[](1);
        shares[0] = amount;
        // Add token to array
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(token);

        // Get the current withdrawal nonce for sender
        uint256 nonce = dm.cumulativeWithdrawalsQueued(msg.sender);

        // Define QueuedWithdrawalParams struct instance
        IDelegationManagerTypes.QueuedWithdrawalParams[] memory queueWithdrawals =
            new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        queueWithdrawals[0] = IDelegationManagerTypes.QueuedWithdrawalParams({
            strategies: strategies,
            depositShares: shares,
            __deprecated_withdrawer: address(0)
        });

        // Withdrawal roots will be returned when we queue
        bytes32[] memory withdrawalRoots;

        // Log the details we need to reproduce the WithdrawalRoot
        emit log_named_uint("nonce", nonce);
        emit log_named_uint("startBlock", block.number + 1);

        // Queue withdrawal
        withdrawalRoots = dm.queueWithdrawals(queueWithdrawals);

        // Log the withdrawalRoot
        emit log_named_bytes32("withdrawalRoot", withdrawalRoots[0]);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }
}
