// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./StrategyBaseUnit.t.sol";
import "../../contracts/strategies/DurationVaultStrategy.sol";
import "../../contracts/interfaces/IDurationVaultStrategy.sol";
import "../../contracts/interfaces/IDelegationManager.sol";
import "../../contracts/interfaces/IAllocationManager.sol";
import "../../contracts/interfaces/IRewardsCoordinator.sol";
import "../../contracts/interfaces/ISignatureUtilsMixin.sol";
import "../../contracts/libraries/OperatorSetLib.sol";
import "../mocks/DelegationManagerMock.sol";
import "../mocks/AllocationManagerMock.sol";
import "../mocks/RewardsCoordinatorMock.sol";
import "../mocks/StrategyFactoryMock.sol";

contract DurationVaultStrategyUnitTests is StrategyBaseUnitTests {
    DurationVaultStrategy public durationVaultImplementation;
    DurationVaultStrategy public durationVault;
    DelegationManagerMock internal delegationManagerMock;
    AllocationManagerMock internal allocationManagerMock;
    RewardsCoordinatorMock internal rewardsCoordinatorMock;
    StrategyFactoryMock internal strategyFactoryMock;

    // TVL limits for tests
    uint internal maxTotalDeposits = 3200e18;
    uint internal maxPerDeposit = 32e18;

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
        rewardsCoordinatorMock = new RewardsCoordinatorMock();
        strategyFactoryMock = new StrategyFactoryMock();

        durationVaultImplementation = new DurationVaultStrategy(
            strategyManager,
            pauserRegistry,
            IDelegationManager(address(delegationManagerMock)),
            IAllocationManager(address(allocationManagerMock)),
            IRewardsCoordinator(address(rewardsCoordinatorMock)),
            IStrategyFactory(address(strategyFactoryMock))
        );

        IDurationVaultStrategyTypes.VaultConfig memory config = IDurationVaultStrategyTypes.VaultConfig({
            underlyingToken: underlyingToken,
            vaultAdmin: address(this),
            duration: defaultDuration,
            maxPerDeposit: maxPerDeposit,
            stakeCap: maxTotalDeposits,
            metadataURI: "ipfs://duration-vault",
            operatorSet: OperatorSet({avs: OPERATOR_SET_AVS, id: OPERATOR_SET_ID}),
            operatorSetRegistrationData: REGISTRATION_DATA,
            delegationApprover: DELEGATION_APPROVER,
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

        // Set the strategy for inherited tests
        strategy = StrategyBase(address(durationVault));
    }

    // ===================== OPERATOR INTEGRATION TESTS =====================

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
        assertEq(address(durationVault.rewardsCoordinator()), address(rewardsCoordinatorMock), "rewards coordinator mismatch");
        assertTrue(durationVault.operatorSetRegistered(), "operator set should be registered");

        // Verify rewards split is set to 0 (100% to stakers).
        RewardsCoordinatorMock.SetOperatorSetSplitCall memory splitCall = rewardsCoordinatorMock.lastSetOperatorSetSplitCall();
        assertEq(splitCall.operator, address(durationVault), "split operator mismatch");
        assertEq(splitCall.operatorSet.avs, OPERATOR_SET_AVS, "split AVS mismatch");
        assertEq(splitCall.operatorSet.id, OPERATOR_SET_ID, "split operatorSetId mismatch");
        assertEq(splitCall.split, 0, "operator split should be 0 for 100% to stakers");
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

    // ===================== VAULT STATE TESTS =====================

    function testDepositsBlockedAfterLock() public {
        durationVault.lock();

        uint depositAmount = 1e18;

        underlyingToken.transfer(address(durationVault), depositAmount);
        // The deposit() call succeeds (only validates token), but beforeAddShares() reverts
        cheats.startPrank(address(strategyManager));
        uint shares = durationVault.deposit(underlyingToken, depositAmount);
        cheats.expectRevert(IDurationVaultStrategyErrors.DepositsLocked.selector);
        durationVault.beforeAddShares(address(this), shares);
        cheats.stopPrank();
    }

    function testWithdrawalQueueingBlockedDuringAllocations() public {
        // prepare deposit
        uint depositAmount = 10 ether;
        underlyingToken.transfer(address(durationVault), depositAmount);
        cheats.prank(address(strategyManager));
        durationVault.deposit(underlyingToken, depositAmount);

        durationVault.lock();

        assertTrue(durationVault.isLocked(), "vault should be locked");
        assertFalse(durationVault.withdrawalsOpen(), "withdrawals should be closed before maturity");

        uint shares = durationVault.totalShares();

        // Attempt to queue withdrawal (beforeRemoveShares) during ALLOCATIONS - should revert
        cheats.startPrank(address(strategyManager));
        cheats.expectRevert(IDurationVaultStrategyErrors.WithdrawalsLockedDuringAllocations.selector);
        durationVault.beforeRemoveShares(address(this), shares);
        cheats.stopPrank();

        // After maturity, queuing should be allowed
        cheats.warp(block.timestamp + defaultDuration + 1);
        durationVault.markMatured();
        assertTrue(durationVault.withdrawalsOpen(), "withdrawals should open after maturity");

        // Queuing now works
        cheats.prank(address(strategyManager));
        durationVault.beforeRemoveShares(address(this), shares);

        // And completion works
        cheats.prank(address(strategyManager));
        durationVault.withdraw(address(this), underlyingToken, shares);
    }

    // ===================== TVL LIMITS TESTS =====================

    function testSetTVLLimits(uint newMaxPerDeposit, uint newMaxTotalDeposits) public {
        cheats.assume(newMaxPerDeposit <= newMaxTotalDeposits);
        durationVault.updateTVLLimits(newMaxPerDeposit, newMaxTotalDeposits);
        (uint _maxPerDeposit, uint _maxDeposits) = durationVault.getTVLLimits();

        assertEq(_maxPerDeposit, newMaxPerDeposit);
        assertEq(_maxDeposits, newMaxTotalDeposits);
    }

    function testSetInvalidMaxPerDepositAndMaxDeposits(uint newMaxPerDeposit, uint newMaxTotalDeposits) public {
        cheats.assume(newMaxTotalDeposits < newMaxPerDeposit);
        cheats.expectRevert(IStrategyErrors.MaxPerDepositExceedsMax.selector);
        durationVault.updateTVLLimits(newMaxPerDeposit, newMaxTotalDeposits);
    }

    function testDepositMoreThanMaxPerDeposit() public {
        uint newMaxPerDeposit = 1e18;
        uint newMaxTotalDeposits = 10e18;
        durationVault.updateTVLLimits(newMaxPerDeposit, newMaxTotalDeposits);

        // Set up delegation to vault
        address staker = address(0xBEEF);
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory emptySig;
        cheats.prank(staker);
        delegationManagerMock.delegateTo(address(durationVault), emptySig, bytes32(0));

        uint depositAmount = newMaxPerDeposit + 1;
        underlyingToken.transfer(address(durationVault), depositAmount);

        cheats.startPrank(address(strategyManager));
        uint shares = durationVault.deposit(underlyingToken, depositAmount);
        cheats.expectRevert(IStrategyErrors.MaxPerDepositExceedsMax.selector);
        durationVault.beforeAddShares(staker, shares);
        cheats.stopPrank();
    }

    function testDepositMoreThanMaxTotalDeposits() public {
        uint newMaxPerDeposit = 3e11;
        uint newMaxTotalDeposits = 1e12;
        uint numDeposits = newMaxTotalDeposits / newMaxPerDeposit;

        durationVault.updateTVLLimits(newMaxPerDeposit, newMaxTotalDeposits);

        // Set up delegation to vault
        address staker = address(0xBEEF);
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory emptySig;
        cheats.prank(staker);
        delegationManagerMock.delegateTo(address(durationVault), emptySig, bytes32(0));

        uint cumulativeShares = 0;
        for (uint i = 0; i < numDeposits; i++) {
            underlyingToken.transfer(address(durationVault), newMaxPerDeposit);
            cheats.startPrank(address(strategyManager));
            uint shares = durationVault.deposit(underlyingToken, newMaxPerDeposit);
            // beforeAddShares checks operatorShares + new shares, so we update after each
            durationVault.beforeAddShares(staker, shares);
            cheats.stopPrank();
            cumulativeShares += shares;
            delegationManagerMock.setOperatorShares(address(durationVault), IStrategy(address(durationVault)), cumulativeShares);
        }

        // Next deposit exceeds the cap
        underlyingToken.transfer(address(durationVault), newMaxPerDeposit);
        cheats.startPrank(address(strategyManager));
        uint shares = durationVault.deposit(underlyingToken, newMaxPerDeposit);
        cheats.expectRevert(IStrategyErrors.BalanceExceedsMaxTotalDeposits.selector);
        durationVault.beforeAddShares(staker, shares);
        cheats.stopPrank();
    }

    function testTVLLimitsCannotBeChangedAfterLock() public {
        durationVault.lock();
        cheats.expectRevert(IDurationVaultStrategyErrors.DepositsLocked.selector);
        durationVault.updateTVLLimits(1e18, 10e18);
    }
}
