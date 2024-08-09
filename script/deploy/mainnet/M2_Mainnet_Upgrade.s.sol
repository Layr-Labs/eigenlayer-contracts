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
        _verifyContractsInitialized();
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
            eigenPodManager,
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

        // Second, configure additional settings and paused statuses
        delegationManager.setMinWithdrawalDelayBlocks(DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS);
        delegationManager.unpause(0);
        eigenPodManager.unpause(0);

        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));

        vm.stopPrank();
    }
}

// forge t --mt test_queueUpgrade --fork-url $RPC_MAINNET -vvvv
contract Queue_M2_Upgrade is M2_Mainnet_Upgrade, TimelockEncoding {
    Vm cheats = Vm(HEVM_ADDRESS);

    // Thurs Apr 08 2024 12:00:00 GMT-0700 (Pacific Daylight Time)
    uint256 timelockEta = 1712559600;

    function test_queueUpgrade() external {
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

        // txs[3] = Tx(
        //     address(eigenLayerProxyAdmin),
        //     0,
        //     abi.encodeWithSelector(
        //         ProxyAdmin.upgrade.selector, 
        //         TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))), 
        //         delayedWithdrawalRouterImplementation
        //     )
        // );

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
        // txs[7] = Tx(
        //     address(eigenPodManager), 
        //     0, // value
        //     abi.encodeWithSelector(EigenPodManager.updateBeaconChainOracle.selector, beaconOracle)
        // );

        // set Deneb fork timestamp on EigenPodManager
        // txs[8] = Tx(
        //     address(eigenPodManager), 
        //     0, // value
        //     abi.encodeWithSelector(EigenPodManager.setDenebForkTimestamp.selector, EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP)
        // );

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
        _verifyContractsInitialized();
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
        existingEigenPod.startCheckpoint(false);
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
