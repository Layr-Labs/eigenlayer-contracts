// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/StrategyManager.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use cast:
//
// cast send <TOKEN_ADDRESS> "approve(address,uint256)" \
// <STRATEGY_MANAGER_ADDRESS> \
// <AMOUNT> \
// --private-key <YOUR_PRIVATE_KEY>
//
// cast send <STRATEGY_MANAGER_ADDRESS> "depositIntoStrategy(address,address,uint256)" \
// <STRATEGY_ADDRESS> \
// <TOKEN_ADDRESS> \
// <AMOUNT> \
// --private-key <YOUR_PRIVATE_KEY>

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/deposit_into_strategy.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address strategy,address token,uint256 amount)" -- <DEPLOYMENT_OUTPUT_JSON> <STRATEGY_ADDRESS> <TOKEN_ADDRESS> <AMOUNT>
// RUST_LOG=forge,foundry=trace forge script script/tasks/deposit_into_strategy.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,address strategy,address token,uint256 amount)" -- local/slashing_output.json 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0x67d269191c92Caf3cD7723F116c85e6E9bf55933 $DEPOSIT_SHARES
contract DepositIntoStrategy is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    function run(string memory configFile, address strategy, address token, uint256 amount) public {
        // Load config
        string memory deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        string memory config_data = vm.readFile(deployConfigPath);

        // Pull strategy manager address
        address strategyManager = stdJson.readAddress(config_data, ".addresses.strategyManager");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        IERC20 tkn = IERC20(token);
        StrategyManager sm = StrategyManager(strategyManager);

        // approve spend
        tkn.approve(strategyManager, amount);

        // do deposit
        sm.depositIntoStrategy(IStrategy(strategy), IERC20(token), amount);

        vm.stopBroadcast();
    }
}
