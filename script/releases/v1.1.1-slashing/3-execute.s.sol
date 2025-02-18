// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {Queue} from "./2-multisig.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {MockAVSRegistrarAlt} from "src/test/mocks/MockAVSRegistrar.sol";

contract Execute is Queue {
    using Env for *;

    function _runAsMultisig() internal override(Queue) prank(Env.protocolCouncilMultisig()) {
        bytes memory calldata_to_executor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
    }

    function testScript() public virtual override(Queue) {
        // 0. Deploy Impls
        runAsEOA();

        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        // 1. Queue Upgrade
        Queue._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        // 2. Warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // 3- execute
        execute();

        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");

        // 4. Validate
        _validateNewImplAddresses(true);
        _validateProxyConstructors();
        _validateProxyInitialized();

        // 5. Validate that the AllocationManager has been correctly upgraded with the following changes:
        // - setRegistrar (should revert if the registrar does not support the AVS)
        // - registerOperator/deregisterOperator (should correctly call with updated interfaces)
        // - createOperatorSets (should revert if the AVS has no registered metadata)
        _validateUpgrade111();
    }

    function _validateProxyConstructors() internal view {
        AllocationManager allocationManager = Env.proxy.allocationManager();
        assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
        assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "alm.pR invalid");
        assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "alm.pc invalid");
        assertTrue(
            allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "alm.deallocationDelay invalid"
        );
        assertTrue(
            allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
            "alm.allocationConfigurationDelay invalid"
        );
    }

    /// @dev Call initialize on all deployed proxies to ensure initializers are disabled
    function _validateProxyInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        AllocationManager allocationManager = Env.proxy.allocationManager();
        vm.expectRevert(errInit);
        allocationManager.initialize(address(0), 0);
    }

    /**
     * @dev Using MockAVSRegistrarAlt which doesn't have fallback so invalid callbacks will revert
     * Tests the following:
     * - setAVSRegistrar (should revert if the registrar does not support the AVS)
     * - createOperatorSets (should revert if the AVS has no registered metadata)
     * - registerOperator/deregisterOperator (should correctly call with updated interfaces)
     */
    function _validateUpgrade111() internal {
        AllocationManager allocationManager = Env.proxy.allocationManager();
        // Deploy mockAVSRegistrar
        MockAVSRegistrarAlt mockAVSRegistrar = new MockAVSRegistrarAlt();
        address avs = address(0x09);
        // 1. setRegistrar (should revert if the registrar does not support the AVS)
        vm.startPrank(avs);
        vm.expectRevert(IAllocationManagerErrors.InvalidAVSRegistrar.selector);
        allocationManager.setAVSRegistrar(avs, IAVSRegistrar(address(mockAVSRegistrar)));

        // set all AVSs to be supported on the mockAVSRegistrar to enable the following call
        mockAVSRegistrar.setSupportsAVS(true);
        allocationManager.setAVSRegistrar(avs, IAVSRegistrar(address(mockAVSRegistrar)));

        // 2. createOperatorSets (should revert if the AVS has no registered metadata)
        // create input data for createOperatorSets
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = Env.instance.strategyBaseTVLLimits(0);
        IAllocationManagerTypes.CreateSetParams[] memory createSetParams =
            new IAllocationManagerTypes.CreateSetParams[](1);
        createSetParams[0] = IAllocationManagerTypes.CreateSetParams({operatorSetId: 0, strategies: strategies});

        vm.expectRevert(IAllocationManagerErrors.InvalidAVSWithNoMetadataRegistered.selector);
        allocationManager.createOperatorSets(avs, createSetParams);

        allocationManager.updateAVSMetadataURI(avs, "test");
        allocationManager.createOperatorSets(avs, createSetParams);

        vm.stopPrank();

        // 3. registerOperator/deregisterOperator (should correctly call with updated interfaces)
        address operator = address(0x0A);
        vm.startPrank(operator);

        // register operator on DM
        DelegationManager delegationManager = Env.proxy.delegationManager();
        delegationManager.registerAsOperator(address(0), 0, "");

        // register operator to OpSet
        uint32[] memory operatorSetIds = new uint32[](1);
        operatorSetIds[0] = 0;
        bytes memory emptyData;
        IAllocationManagerTypes.RegisterParams memory registerParams =
            IAllocationManagerTypes.RegisterParams({avs: avs, operatorSetIds: operatorSetIds, data: emptyData});
        allocationManager.registerForOperatorSets(operator, registerParams);
        // check that the operator is registered to the OpSet
        OperatorSet memory operatorSet = OperatorSet(avs, 0);
        assertTrue(allocationManager.isMemberOfOperatorSet(operator, operatorSet), "operator not registered to OpSet");

        // deregister operator from OpSet
        IAllocationManagerTypes.DeregisterParams memory deregisterParams =
            IAllocationManagerTypes.DeregisterParams({operator: operator, avs: avs, operatorSetIds: operatorSetIds});
        allocationManager.deregisterFromOperatorSets(deregisterParams);
        // check that the operator is not registered to the OpSet
        assertFalse(allocationManager.isMemberOfOperatorSet(operator, operatorSet), "operator still registered to OpSet");
        
        vm.stopPrank();
    }
}
