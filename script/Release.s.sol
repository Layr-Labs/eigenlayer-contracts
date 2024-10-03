// // SPDX-License-Identifier: BUSL-1.1
// pragma solidity ^0.8.12;

// import "./utils/Releasoor.s.sol";
// import "./utils/Encoders.sol";

// contract Release is Releasoor {

//     using TxBuilder for *;
//     using AddressUtils for *;

//     struct ReleaseStep {
//         function () external step;
//     }

//     // Public, standard release:
//     // - EOA deploy
//     // - queue/execute via ops multisig/timelock
//     function manifest() public returns (ReleaseStep[] memory steps) {
//         return _newRelease(ReleaseType.OPS_TIMELOCK)
//             .append(this.deploy)
//             .append(this.queueUpgrade)
//             .wait(10 days)
//             .append(this.executeUpgrade);
//     }

//     // Private, emergency release:
//     // - EOA deploy
//     // - execute instantly via community msig
//     function manifest() public returns (ReleaseStep[] memory steps) {
//         return _newRelease(ReleaseType.COMMUNITY_MSIG)
//             .append(this.deploy)
//             .append(this.executeUpgrade);
//     }

//     struct Deployment {
//         string name;
//         address deployedTo;
//     }

//     function deploy(
//         Addresses memory addrs,
//         Environment memory env,
//         Params memory params
//     ) internal override returns (Deployment[] memory deployments) {


//         vm.startBroadcast();

//         deployments[0] = Deployment({
//             name: type(EigenPod).name,
//             deployedTo: address(new EigenPod(
//                 IETHPOSDeposit(params.ethPOS),
//                 IEigenPodManager(addrs.eigenPodManager.proxy),
//                 params.EIGENPOD_GENESIS_TIME
//             ))
//         });

//         EigenPodManager newEPMImpl = new EigenPodManager(
//             IETHPOSDeposit(params.ethPOS),
//             IUpgradeableBeacon(addrs.eigenPod.beacon),
//             IStrategyManager(addrs.strategyManager.proxy),
//             ISlasher(addrs.slasher.proxy),
//             IDelegationManager(addrs.delegationManager.proxy)
//         );

//         vm.stopBroadcast();

//         type(EigenPod).contractName;



//         _logDeploy(EIGENPOD, address(newEigenPodImpl));

//         // addrs.eigenPod.setPending(address(newEigenPodImpl));
//         // addrs.eigenPodManager.setPending(address(newEPMImpl));
//     }

//     function queueUpgrade(Addresses memory addrs) internal override {
//         Txs storage txns = _newTxs();
//         eta = env.isMainnet() ? 12351235 : 0;

//         txns.append({
//             to: addrs.eigenPod.beacon,
//             data: EncBeacon.upgradeTo(addrs.eigenPod.getPending())
//         });

//         txns.append({
//             to: addrs.proxyAdmin,
//             data: EncProxyAdmin.upgrade(addrs.eigenPodManager.proxy, addrs.eigenPodManager.getPending())
//         });

//         (
//             bytes memory calldata_to_timelock_queueing_action,
//             bytes memory calldata_to_timelock_executing_action,
//             bytes memory final_calldata_to_executor_multisig
//         ) = EncTimelock.queueTransaction({
//             timelock: addrs.timelock,
//             multisend: params.multiSendCallOnly,
//             executor: addrs.executorMultisig,
//             executorTxns: txns,
//             eta: eta
//         });

//         _log("calldata_to_timelock_queueing_action", calldata_to_timelock_queueing_action);
//         _log("calldata_to_timelock_executing_action", calldata_to_timelock_executing_action);
//         _log("final_calldata_to_executor_multisig", final_calldata_to_executor_multisig);

//         vm.startBroadcast(addrs.operationsMultisig);

//         (success, ) = addrs.timelock.call(calldata_to_timelock_queueing_action);
//         require(success, "queueing transaction in timelock failed");

//         vm.stopBroadcast();
//     }

//     function executeUpgrade(Addresses memory addrs) internal override {
//         Txs storage txs = _newTxs();

//         txs.append({
//             to: addrs.eigenPod.beacon,
//             data: EncUpgradeableBeacon.upgradeTo(addrs.eigenPod.getPending())
//         });

//         txs.append({
//             to: addrs.eigenPodManager.proxy,
//             data: EncProxyAdmin.upgrade(addrs.eigenPodManager.proxy, addrs.eigenPodManager.getPending())
//         });

//         bytes memory calldata_to_multisend_contract = EncMultiSendCallOnly.multiSend(txs);

//         bytes memory final_calldata_to_executor_multisig = EncGnosisSafe.execTransaction({
//             from: addrs.timelock,
//             to: params.multiSendCallOnly,
//             data: calldata_to_multisend_contract,
//             op: EncGnosisSafe.Operation.DelegateCall
//         });

//         bytes memory calldata_to_timelock_executing_action = EncTimelock.executeTransaction({
//             target: addrs.executorMultisig,
//             data: final_calldata_to_executor_multisig,
//             eta: 0
//         });

//         emit log_named_bytes("calldata_to_timelock_executing_action", calldata_to_timelock_executing_action);

//         Txs storage txns = _newTxs();

//         txns.append({
//             to: addrs.timelock,
//             data: calldata_to_timelock_executing_action
//         });

//         txns.append({
//             to: addrs.strategyFactory.proxy,
//             data: EncStrategyFactory.whitelistStrategies(strats)
//         });

//         txns.broadcastFrom(addrs.opsMultisig);

//         // Update config
//         addrs.eigenPod.updateFromPending();
//         addrs.eigenPodManager.updateFromPending();
//     }

//     function queueUpgrade(Addresses memory addrs) internal override returns (Txns storage) {
//         Txns storage txns = addrs.opsMultisig.queueTimelock({
//             timelock: addrs.timelock,
//             eta: 0,
//             _txns: addrs.executorMultisig
//                 .append({
//                     to: addrs.strategyManager.proxy,
//                     data: EncStrategyManager.setStrategyWhitelister(addrs.opsMultisig)
//                 })
//                 .append({
//                     to: addrs.eigenPod.beacon,
//                     data: EncBeacon.upgradeTo(addrs.eigenPod.getPending())
//                 })
//                 .append({
//                     to: addrs.proxyAdmin,
//                     data: EncProxyAdmin.upgrade(addrs.eigenPodManager.proxy, addrs.eigenPodManager.getPending())
//                 })
//                 .asSafeExecTxn()
//         });

//         txns.printSummary();

//         return txns;
//     }

//     function executeUpgrade(Addresses memory addrs) internal override returns (Txns storage) {
//         Txns storage txns = addrs.opsMultisig.executeTimelock({
//             timelock: addrs.timelock,
//             eta: 0,
//             to: addrs.executorMultisig,
//             data: addrs.executorMultisig
//                 .append({
//                     to: addrs.strategyManager.proxy,
//                     data: EncStrategyManager.setStrategyWhitelister(addrs.opsMultisig)
//                 })
//                 .append({
//                     to: addrs.eigenPod.beacon,
//                     data: EncBeacon.upgradeTo(addrs.eigenPod.pendingImpl)
//                 })
//                 .append({
//                     to: addrs.proxyAdmin,
//                     data: EncProxyAdmin.upgrade(addrs.eigenPodManager.proxy, addrs.eigenPodManager.pendingImpl)
//                 })
//         })
//         .append({
//             to: addrs.strategyManager.proxy,
//             data: EncStrategyManager.addStrategiesToDepositWhitelist(strats, bools)
//         });

//         txns.printSummary();

//         // Update config
//         addrs.eigenPod.updateFromPending();
//         addrs.eigenPodManager.updateFromPending();

//         return txns;
//     }

//     // function _deployPEPE(Addresses memory addrs) internal returns (EigenPod, EigenPodManager) {
//     //     // Deploy EigenPod
//     //     eigenPodImplementation = new EigenPod(
//     //         IETHPOSDeposit(cfg.ETHPOSDepositAddress),
//     //         IEigenPodManager(addrs.eigenPodManager.proxy),
//     //         cfg.EIGENPOD_GENESIS_TIME
//     //     );

//     //     // Deploy EigenPodManager
//     //     eigenPodManagerImplementation = new EigenPodManager(
//     //         IETHPOSDeposit(cfg.ETHPOSDepositAddress),
//     //         IBeacon(addrs.eigenPodBeacon.beacon),
//     //         IStrategyManager(addrs.strategyManager.proxy),
//     //         ISlasher(addrs.slasher.proxy),
//     //         IDelegationManager(addrs.delegationManager.proxy)
//     //     );

//     //     return (eigenPodImplementation, eigenPodManagerImplementation);
//     // }

//     // function test_Release() public {
//     //     _readEnvironment(ENV_MAINNET);

//     //     _printAddrs();
//     //     emit log("=====");

//     //     _readEnvironment(ENV_HOLESKY);

//     //     _printAddrs();
//     //     emit log("=====");

//     //     _readEnvironment(ENV_PREPROD);

//     //     _printAddrs();
//     //     emit log("=====");
//     // }
// }
