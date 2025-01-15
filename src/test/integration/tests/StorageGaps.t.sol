// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/users/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_StorageGaps is IntegrationCheckUtils {
    function assertSlotEq(address target, uint256 slot, uint256 value) internal {
        assertEq(uint256(vm.load(target, bytes32(slot))), value, "Slot value mismatch");
    }

    function assertSlotEq(address target, uint256 slot, address addr) internal {
        assertEq(address(uint160(uint256(vm.load(target, bytes32(slot))))), addr, "Slot value mismatch");
    }

    function testFuzz_upgrade_ProveOldSlotsPersist() public {
        vm.skip(block.chainid != 1);

        assertSlotEq(address(strategyFactory), 101, strategyFactory.owner());

        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(strategyFactory))),
            address(new StrategyFactory(strategyManager, eigenLayerPauserReg)),
            abi.encodeWithSelector(
                StrategyFactory.initialize.selector,
                executorMultisig,
                0, // initial paused status
                IBeacon(strategyBeacon)
            )
        );

        // If these slots match, we know old storage still persists...
        assertSlotEq(address(strategyFactory), 101, strategyFactory.owner());
        assertSlotEq(address(strategyFactory), 102, strategyFactory.owner());
    }
}
