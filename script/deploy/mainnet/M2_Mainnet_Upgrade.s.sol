// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";
import "../../utils/TimelockEncoding.sol";
import "../../utils/Multisend.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * anvil --fork-url $RPC_MAINNET
 * forge script script/deploy/mainnet/M2_Mainnet_Upgrade.s.sol:M2_Mainnet_Upgrade --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * 
 * forge script script/deploy/mainnet/M2_Mainnet_Upgrade.s.sol:M2_Mainnet_Upgrade --rpc-url $RPC_MAINNET --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 */
contract M2_Mainnet_Upgrade is ExistingDeploymentParser {
    function run() external virtual {
        _parseDeployedContracts("script/output/mainnet/M1_deployment_mainnet_2023_6_9.json");
        _parseInitialDeploymentParams("script/configs/mainnet/M2_mainnet_upgrade.config.json");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _deployImplementationContracts();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Simulate upgrade of contracts to new implementations
        _simulateUpgrade();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/mainnet/M2_mainnet_upgrade.output.json");
    }

    /**
     * @notice Deploy EigenLayer contracts from scratch for Holesky
     */
    function _deployImplementationContracts() internal {
        // 1. Deploy New TUPS
        avsDirectoryImplementation = new AVSDirectory(delegationManager);
        avsDirectory = AVSDirectory(
            address(
                new TransparentUpgradeableProxy(
                    address(avsDirectoryImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        AVSDirectory.initialize.selector,
                        executorMultisig, // initialOwner
                        eigenLayerPauserReg,
                        AVS_DIRECTORY_INIT_PAUSED_STATUS
                    )
                )
            )
        );

        // 2. Deploy Implementations
        eigenPodImplementation = new EigenPod(
            IETHPOSDeposit(ETHPOSDepositAddress),
            delayedWithdrawalRouter,
            eigenPodManager,
            EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            EIGENPOD_GENESIS_TIME
        );
        delegationManagerImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenPodManager, slasher);
        slasherImplementation = new Slasher(strategyManager, delegationManager);
        eigenPodManagerImplementation = new EigenPodManager(
            IETHPOSDeposit(ETHPOSDepositAddress),
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegationManager
        );
        delayedWithdrawalRouterImplementation = new DelayedWithdrawalRouter(eigenPodManager);
    }

    function _simulateUpgrade() internal {

        vm.startPrank(executorMultisig);

        // First, upgrade the proxy contracts to point to the implementations
        // AVSDirectory
        // eigenLayerProxyAdmin.upgrade(
        //     TransparentUpgradeableProxy(payable(address(avsDirectory))),
        //     address(avsDirectoryImplementation)
        // );
        // DelegationManager
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(delegationManager))),
            address(delegationManagerImplementation)
        );
        // StrategyManager
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation)
        );
        // Slasher
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation)
        );
        // EigenPodManager
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation)
        );
        // Delayed Withdrawal Router
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))),
            address(delayedWithdrawalRouterImplementation)
        );

        // Second, configure additional settings and paused statuses
        delegationManager.setMinWithdrawalDelayBlocks(DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS);
        delegationManager.unpause(0);
        eigenPodManager.unpause(0);

        eigenPodManager.setDenebForkTimestamp(EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP);
        eigenPodManager.updateBeaconChainOracle(beaconOracle);
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));

        vm.stopPrank();
    }
}

// forge t --mt test_queueUpgrade --fork-url $RPC_MAINNET -vvvv
contract Queue_M2_Upgrade is M2_Mainnet_Upgrade, TimelockEncoding {
    Vm cheats = Vm(HEVM_ADDRESS);

    // Thurs Apr 08 2024 12:00:00 GMT-0700 (Pacific Daylight Time)
    uint256 timelockEta = 1712559600;

    function test_M2_Script_queueUpgrade() external {
        _parseDeployedContracts("script/output/mainnet/M2_mainnet_upgrade.output.json");
        _parseInitialDeploymentParams("script/configs/mainnet/M2_mainnet_upgrade.config.json");

        Tx[] memory txs = new Tx[](11);
        // upgrade the DelegationManager, Slasher, StrategyManager, DelayedWithdrawalRouter, EigenPodManager, & EigenPod contracts
        txs[0] = Tx(
            address(eigenLayerProxyAdmin),
            0,
            abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector, 
                TransparentUpgradeableProxy(payable(address(delegationManager))), 
                delegationManagerImplementation
            )
        );

        txs[1] = Tx(
            address(eigenLayerProxyAdmin),
            0,
            abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector, 
                TransparentUpgradeableProxy(payable(address(slasher))), 
                slasherImplementation
            )
        );

        txs[2] = Tx(
            address(eigenLayerProxyAdmin),
            0,
            abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector, 
                TransparentUpgradeableProxy(payable(address(strategyManager))), 
                strategyManagerImplementation
            )
        );

        txs[3] = Tx(
            address(eigenLayerProxyAdmin),
            0,
            abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector, 
                TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))), 
                delayedWithdrawalRouterImplementation
            )
        );

        txs[4] = Tx(
            address(eigenLayerProxyAdmin),
            0,
            abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector, 
                TransparentUpgradeableProxy(payable(address(eigenPodManager))), 
                eigenPodManagerImplementation
            )
        );

        txs[5] = Tx(
            address(eigenPodBeacon),
            0,
            abi.encodeWithSelector(
                UpgradeableBeacon.upgradeTo.selector, 
                eigenPodImplementation
            )
        );

        // set the min withdrawal delay blocks on the DelegationManager
        txs[6] = Tx(
            address(delegationManager), 
            0, // value
            abi.encodeWithSelector(DelegationManager.setMinWithdrawalDelayBlocks.selector, DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS)
        );

        // set beacon chain oracle on EigenPodManager
        txs[7] = Tx(
            address(eigenPodManager), 
            0, // value
            abi.encodeWithSelector(EigenPodManager.updateBeaconChainOracle.selector, beaconOracle)
        );

        // set Deneb fork timestamp on EigenPodManager
        txs[8] = Tx(
            address(eigenPodManager), 
            0, // value
            abi.encodeWithSelector(EigenPodManager.setDenebForkTimestamp.selector, EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP)
        );

        // unpause everything on DelegationManager
        txs[9] = Tx(
            address(delegationManager), 
            0, // value
            abi.encodeWithSelector(Pausable.unpause.selector, 0)
        );

        // unpause everything on EigenPodManager
        txs[10] = Tx(
            address(eigenPodManager), 
            0, // value
            abi.encodeWithSelector(Pausable.unpause.selector, 0)
        );

        bytes memory calldata_to_multisend_contract = abi.encodeWithSelector(MultiSendCallOnly.multiSend.selector, encodeMultisendTxs(txs));
        emit log_named_bytes("calldata_to_multisend_contract", calldata_to_multisend_contract);

        bytes memory final_calldata_to_executor_multisig = encodeForExecutor({
            // call to executor will be from the timelock
            from: timelock,
            // performing many operations at the same time
            to: multiSendCallOnly,
            // value to send in tx
            value: 0,
            // calldata for the operation
            data: calldata_to_multisend_contract,
            // operation type (for performing many operations at the same time)
            operation: ISafe.Operation.DelegateCall
        });

        (bytes memory calldata_to_timelock_queuing_action, bytes memory calldata_to_timelock_executing_action) = encodeForTimelock({
            // address to be called from the timelock
            to: executorMultisig,
            // value to send in tx
            value: 0,
            // calldata for the operation
            data: final_calldata_to_executor_multisig,
            // time at which the tx will become executable
            timelockEta: timelockEta
        });

        bytes32 expectedTxHash = getTxHash({
            target: executorMultisig,
            _value: 0,
            _data: final_calldata_to_executor_multisig,
            eta: timelockEta
        });
        emit log_named_bytes32("expectedTxHash", expectedTxHash);

        cheats.prank(operationsMultisig);
        (bool success, ) = timelock.call(calldata_to_timelock_queuing_action);
        require(success, "call to timelock queuing action failed");

        require(ITimelock(timelock).queuedTransactions(expectedTxHash), "expectedTxHash not queued");

        // test performing the upgrade
        cheats.warp(timelockEta);
        cheats.prank(operationsMultisig);
        (success, ) = timelock.call(calldata_to_timelock_executing_action);
        require(success, "call to timelock executing action failed");

        // Check correctness after upgrade
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();
        _postUpgradeChecks();
    }

    function test_M2_Script_executeUpgrade() external {
        _parseDeployedContracts("script/output/mainnet/M2_mainnet_upgrade.output.json");
        _parseInitialDeploymentParams("script/configs/mainnet/M2_mainnet_upgrade.config.json");

        bytes memory calldata_to_timelock_executing_action = hex"0825f38f000000000000000000000000369e6f597e22eab55ffb173c6d9cd234bd699111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000661395f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008246a76120200000000000000000000000040a2accbd92bca938b02010e17a5b8929b49130d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007a000000000000000000000000000000000000000000000000000000000000006248d80ff0a000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000005d3008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000039053d51b77dc0d36036fc1fcc8cb819df8ef37a0000000000000000000000001784be6401339fc0fedf7e9379409f5c1bfe9dda008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec4000000000000000000000000d92145c07f8ed1d392c1b88017934e301cc1c3cd000000000000000000000000f3234220163a757edf1e11a8a085638d9b236614008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec4000000000000000000000000858646372cc42e1a627fce94aa7a7033e7cf075a00000000000000000000000070f44c13944d49a236e3cd7a94f48f5dab6c619b008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec40000000000000000000000007fe7e9cc0f274d2435ad5d56d5fa73e47f6a23d80000000000000000000000004bb6731b02314d40abbffbc4540f508874014226008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000091e677b07f7af907ec9a428aafa9fc14a0d3a338000000000000000000000000e4297e3dadbc7d99e26a2954820f514cb50c5762005a2a4f2f3c18f09179b6703e63d9edd165909073000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000243659cfe60000000000000000000000008ba40da60f0827d027f029acee62609f0527a2550039053d51b77dc0d36036fc1fcc8cb819df8ef37a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024635bbd10000000000000000000000000000000000000000000000000000000000000c4e00091e677b07f7af907ec9a428aafa9fc14a0d3a33800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024c1de3aef000000000000000000000000343907185b71adf0eba9567538314396aa9854420091e677b07f7af907ec9a428aafa9fc14a0d3a33800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024463db0380000000000000000000000000000000000000000000000000000000065f1b0570039053d51b77dc0d36036fc1fcc8cb819df8ef37a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024fabc1cbc00000000000000000000000000000000000000000000000000000000000000000091e677b07f7af907ec9a428aafa9fc14a0d3a33800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024fabc1cbc000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041000000000000000000000000a6db1a8c5a981d1536266d2a393c5f8ddb210eaf0000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

        // test performing the upgrade
        cheats.warp(timelockEta);
        cheats.prank(operationsMultisig);
        (bool success, ) = timelock.call(calldata_to_timelock_executing_action);
        require(success, "call to timelock executing action failed");

        emit log("we made it here!");

        // Check correctness after upgrade
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();
        _postUpgradeChecks();
    }


    function _postUpgradeChecks() internal {
        // check that LST deposits are paused
        address rETH = 0xae78736Cd615f374D3085123A210448E74Fc6393;
        address rETH_Strategy = 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2;
        uint256 amount = 1e18;
        cheats.prank(rETH);
        // this works because rETH has more than 1 ETH of its own token at its address :)
        IERC20(rETH).transfer(address(this), amount);
        IERC20(rETH).approve(address(strategyManager), amount);
        cheats.expectRevert("Pausable: index is paused");
        strategyManager.depositIntoStrategy({
            strategy: IStrategy(rETH_Strategy),
            token: IERC20(rETH),
            amount: amount
        });

        // unpause LST deposits and check that a deposit works
        cheats.prank(executorMultisig);
        strategyManager.unpause(0);
        strategyManager.depositIntoStrategy({
            strategy: IStrategy(rETH_Strategy),
            token: IERC20(rETH),
            amount: amount
        });

        // check that EigenPod proofs are live (although this still reverts later in the call)
        EigenPod existingEigenPod = EigenPod(payable(0x0b347D5E38296277E829CE1D8C6b82e4c63C2Df3));
        BeaconChainProofs.StateRootProof memory stateRootProof;
        uint40[] memory validatorIndices;
        bytes[] memory validatorFieldsProofs;
        bytes32[][] memory validatorFields;
        cheats.startPrank(existingEigenPod.podOwner());
        existingEigenPod.activateRestaking();
        cheats.expectRevert("EigenPodManager.getBlockRootAtTimestamp: state root at timestamp not yet finalized");
        existingEigenPod.verifyWithdrawalCredentials(
            uint64(block.timestamp),
            stateRootProof,
            validatorIndices,
            validatorFieldsProofs,
            validatorFields
        );
    }

    function getTxHash(address target, uint256 _value, bytes memory _data, uint256 eta) public pure returns (bytes32) {
        // empty bytes
        bytes memory signature;
        bytes32 txHash = keccak256(abi.encode(target, _value, signature, _data, eta));
        return txHash;        
    }
}
