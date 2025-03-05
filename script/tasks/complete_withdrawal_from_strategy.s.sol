// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/AllocationManager.sol";
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/pods/EigenPodManager.sol";
import "../../src/contracts/libraries/SlashingLib.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// use forge:
// RUST_LOG=forge,foundry=trace forge script script/tasks/complete_withdrawal_from_strategy.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,,address strategy,address token,uint256 amount)" -- <DEPLOYMENT_OUTPUT_JSON> <STRATEGY_ADDRESS> <TOKEN_ADDRESS> <AMOUNT> <NONCE> <START_BLOCK_NUMBER>
// RUST_LOG=forge,foundry=trace forge script script/tasks/complete_withdrawal_from_strategy.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile,,address strategy,address token,uint256 amount,uint256 nonce,uint32 startBlock)" -- local/slashing_output.json 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0x67d269191c92Caf3cD7723F116c85e6E9bf55933 750 0 630
contract CompleteWithdrawFromStrategy is Script, Test {
    using SlashingLib for *;

    Vm cheats = Vm(VM_ADDRESS);

    string public deployConfigPath;
    string public config_data;

    function run(
        string memory configFile,
        address strategy,
        address token,
        uint256 amount,
        uint256 nonce,
        uint32 startBlock
    ) public {
        // Load config
        deployConfigPath = string(bytes(string.concat("script/output/", configFile)));
        config_data = vm.readFile(deployConfigPath);

        // Pull addresses from config
        AllocationManager am = AllocationManager(stdJson.readAddress(config_data, ".addresses.allocationManager"));
        DelegationManager dm = DelegationManager(stdJson.readAddress(config_data, ".addresses.delegationManager"));
        EigenPodManager em = EigenPodManager(stdJson.readAddress(config_data, ".addresses.eigenPodManager"));

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        // Add token to array
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(token);

        // Get the withdrawal struct
        IDelegationManagerTypes.Withdrawal memory withdrawal =
            getWithdrawalStruct(am, dm, em, strategy, amount, nonce, startBlock);

        // complete
        dm.completeQueuedWithdrawal(withdrawal, tokens, true);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }

    function getWithdrawalStruct(
        AllocationManager am,
        DelegationManager dm,
        EigenPodManager em,
        address strategy,
        uint256 amount,
        uint256 nonce,
        uint32 startBlock
    ) internal returns (IDelegationManagerTypes.Withdrawal memory) {
        // Add strategy to array
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(strategy);
        // Add shares to array
        uint256[] memory shares = new uint256[](1);
        shares[0] = amount;

        // Get DSF for Staker in strategy
        DepositScalingFactor memory dsf = DepositScalingFactor(dm.depositScalingFactor(msg.sender, strategies[0]));

        // Get TM for Operator in strategies
        uint64[] memory maxMagnitudes = am.getMaxMagnitudesAtBlock(msg.sender, strategies, startBlock);
        uint256 slashingFactor = _getSlashingFactor(em, msg.sender, strategies[0], maxMagnitudes[0]);
        uint256 sharesToWithdraw = dsf.calcWithdrawable(amount, slashingFactor);

        // Get scaled shares for the given amount
        uint256[] memory scaledShares = new uint256[](1);
        scaledShares[0] = dsf.scaleForQueueWithdrawal(sharesToWithdraw);

        // Log the current state before completing
        emit log_uint(dsf.scalingFactor());
        emit log_uint(maxMagnitudes[0]);
        emit log_uint(scaledShares[0]);

        // Create the withdrawal struct
        IDelegationManagerTypes.Withdrawal memory withdrawal = IDelegationManagerTypes.Withdrawal({
            staker: msg.sender,
            delegatedTo: msg.sender,
            withdrawer: msg.sender,
            nonce: nonce,
            startBlock: startBlock,
            strategies: strategies,
            scaledShares: scaledShares
        });

        // Return the withdrawal struct
        return withdrawal;
    }

    function _getSlashingFactor(
        EigenPodManager em,
        address staker,
        IStrategy strategy,
        uint64 operatorMaxMagnitude
    ) internal view returns (uint256) {
        if (strategy == em.beaconChainETHStrategy()) {
            uint64 beaconChainSlashingFactor = em.beaconChainSlashingFactor(staker);
            return operatorMaxMagnitude.mulWad(beaconChainSlashingFactor);
        }

        return operatorMaxMagnitude;
    }
}
