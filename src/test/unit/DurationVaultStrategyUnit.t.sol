// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./StrategyBaseTVLLimitsUnit.sol";
import "../../contracts/strategies/DurationVaultStrategy.sol";
import "../../contracts/interfaces/IDurationVaultStrategy.sol";
import "../../contracts/interfaces/IDelegationManager.sol";
import "../../contracts/interfaces/IAllocationManager.sol";
import "../../contracts/libraries/OperatorSetLib.sol";
import "../mocks/DelegationManagerMock.sol";
import "../mocks/AllocationManagerMock.sol";

contract DurationVaultStrategyUnitTests is StrategyBaseTVLLimitsUnitTests {
    DurationVaultStrategy public durationVaultImplementation;
    DurationVaultStrategy public durationVault;
    DelegationManagerMock internal delegationManagerMock;
    AllocationManagerMock internal allocationManagerMock;

    uint32 internal defaultDuration = uint32(30 days);
    address internal constant OPERATOR_SET_AVS = address(0xA11CE);
    uint32 internal constant OPERATOR_SET_ID = 42;
    address internal constant DELEGATION_APPROVER = address(0xB0B);
    uint32 internal constant OPERATOR_ALLOCATION_DELAY = 5;
    string internal constant OPERATOR_METADATA_URI = "ipfs://operator-metadata";
    bytes internal constant REGISTRATION_DATA = hex"1234";
    uint64 internal constant FULL_ALLOCATION = 1e18;

    function setUp() public virtual override {
        StrategyBaseUnitTests.setUp();

        delegationManagerMock = new DelegationManagerMock();
        // Configure min withdrawal delay so allocationDelayBlocks == OPERATOR_ALLOCATION_DELAY.
        delegationManagerMock.setMinWithdrawalDelayBlocks(OPERATOR_ALLOCATION_DELAY - 1);
        allocationManagerMock = new AllocationManagerMock();

        durationVaultImplementation = new DurationVaultStrategy(
            strategyManager,
            pauserRegistry,
            IDelegationManager(address(delegationManagerMock)),
            IAllocationManager(address(allocationManagerMock))
        );

        IDurationVaultStrategy.VaultConfig memory config = IDurationVaultStrategy.VaultConfig({
            underlyingToken: underlyingToken,
            vaultAdmin: address(this),
            duration: defaultDuration,
            maxPerDeposit: maxPerDeposit,
            stakeCap: maxTotalDeposits,
            metadataURI: "ipfs://duration-vault",
            operatorSet: OperatorSet({avs: OPERATOR_SET_AVS, id: OPERATOR_SET_ID}),
            operatorSetRegistrationData: REGISTRATION_DATA,
            delegationApprover: DELEGATION_APPROVER,
            operatorAllocationDelay: OPERATOR_ALLOCATION_DELAY,
            operatorMetadataURI: OPERATOR_METADATA_URI
        });

        durationVault = DurationVaultStrategy(
            address(
                new TransparentUpgradeableProxy(
                    address(durationVaultImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(DurationVaultStrategy.initialize.selector, config)
                )
            )
        );

        strategy = StrategyBase(address(durationVault));
        strategyWithTVLLimits = StrategyBaseTVLLimits(address(durationVault));
    }

    function testInitializeConfiguresOperatorIntegration() public {
        DelegationManagerMock.RegisterAsOperatorCall memory delegationCall = delegationManagerMock.lastRegisterAsOperatorCall();
        assertEq(delegationCall.operator, address(durationVault), "delegation operator mismatch");
        assertEq(delegationCall.delegationApprover, DELEGATION_APPROVER, "delegation approver mismatch");
        assertEq(delegationCall.allocationDelay, OPERATOR_ALLOCATION_DELAY, "allocation delay mismatch");
        assertEq(delegationCall.metadataURI, OPERATOR_METADATA_URI, "metadata mismatch");

        AllocationManagerMock.RegisterCall memory registerCall = allocationManagerMock.lastRegisterForOperatorSetsCall();
        assertEq(registerCall.operator, address(durationVault), "register operator mismatch");
        assertEq(registerCall.avs, OPERATOR_SET_AVS, "register AVS mismatch");
        assertEq(registerCall.operatorSetIds.length, 1, "unexpected operatorSetIds length");
        assertEq(registerCall.operatorSetIds[0], OPERATOR_SET_ID, "operatorSetId mismatch");
        assertEq(registerCall.data, REGISTRATION_DATA, "registration data mismatch");

        (address avs, uint32 operatorSetId) = durationVault.operatorSetInfo();
        assertEq(avs, OPERATOR_SET_AVS, "stored AVS mismatch");
        assertEq(operatorSetId, OPERATOR_SET_ID, "stored operatorSetId mismatch");
        assertEq(address(durationVault.delegationManager()), address(delegationManagerMock), "delegation manager mismatch");
        assertEq(address(durationVault.allocationManager()), address(allocationManagerMock), "allocation manager mismatch");
        assertTrue(durationVault.operatorSetRegistered(), "operator set should be registered");
    }

    function testLockAllocatesFullMagnitude() public {
        assertEq(allocationManagerMock.modifyAllocationsCallCount(), 0, "precondition failed");

        durationVault.lock();

        assertTrue(durationVault.allocationsActive(), "allocations should be active after lock");
        assertEq(allocationManagerMock.modifyAllocationsCallCount(), 1, "modifyAllocations not called");

        AllocationManagerMock.AllocateCall memory allocateCall = allocationManagerMock.lastModifyAllocationsCall();
        assertEq(allocateCall.operator, address(durationVault), "allocate operator mismatch");
        assertEq(allocateCall.avs, OPERATOR_SET_AVS, "allocate AVS mismatch");
        assertEq(allocateCall.operatorSetId, OPERATOR_SET_ID, "allocate operatorSetId mismatch");
        assertEq(address(allocateCall.strategy), address(durationVault), "allocate strategy mismatch");
        assertEq(allocateCall.magnitude, FULL_ALLOCATION, "allocate magnitude mismatch");
    }

    function testMarkMaturedDeallocatesAndDeregisters() public {
        durationVault.lock();
        cheats.warp(block.timestamp + defaultDuration + 1);

        durationVault.markMatured();

        assertEq(allocationManagerMock.modifyAllocationsCallCount(), 2, "deallocation not invoked");
        AllocationManagerMock.AllocateCall memory allocateCall = allocationManagerMock.lastModifyAllocationsCall();
        assertEq(allocateCall.magnitude, 0, "expected zero magnitude");

        AllocationManagerMock.DeregisterCall memory deregisterCall = allocationManagerMock.lastDeregisterFromOperatorSetsCall();
        assertEq(deregisterCall.operator, address(durationVault), "deregister operator mismatch");
        assertEq(deregisterCall.avs, OPERATOR_SET_AVS, "deregister AVS mismatch");
        assertEq(deregisterCall.operatorSetIds.length, 1, "unexpected deregister length");
        assertEq(deregisterCall.operatorSetIds[0], OPERATOR_SET_ID, "deregister operatorSetId mismatch");

        assertFalse(durationVault.allocationsActive(), "allocations should be inactive");
        assertFalse(durationVault.operatorSetRegistered(), "operator set should be deregistered");
    }

    function testDepositsBlockedAfterLock() public {
        durationVault.lock();

        uint depositAmount = 1e18;

        underlyingToken.transfer(address(durationVault), depositAmount);
        cheats.prank(address(strategyManager));
        cheats.expectRevert(IDurationVaultStrategy.DepositsLocked.selector);
        durationVault.deposit(underlyingToken, depositAmount);
    }

    function testWithdrawalsBlockedUntilMaturity() public {
        // prepare deposit
        uint depositAmount = 10 ether;
        underlyingToken.transfer(address(durationVault), depositAmount);
        cheats.prank(address(strategyManager));
        durationVault.deposit(underlyingToken, depositAmount);

        durationVault.lock();

        assertTrue(durationVault.isLocked(), "vault should be locked");
        assertFalse(durationVault.withdrawalsOpen(), "withdrawals should be closed before maturity");

        uint shares = durationVault.totalShares();

        // Attempt withdrawal before maturity
        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(IDurationVaultStrategy.WithdrawalsLocked.selector);
        durationVault.withdraw(address(this), underlyingToken, shares);
        cheats.stopPrank();

        cheats.warp(block.timestamp + defaultDuration + 1);
        durationVault.markMatured();
        assertTrue(durationVault.withdrawalsOpen(), "withdrawals should open after maturity");

        cheats.startPrank(address(strategyManager));
        durationVault.withdraw(address(this), underlyingToken, durationVault.totalShares());
        cheats.stopPrank();
    }
}
