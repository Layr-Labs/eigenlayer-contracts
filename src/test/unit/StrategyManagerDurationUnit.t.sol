// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/StrategyManager.sol";
import "src/contracts/strategies/DurationVaultStrategy.sol";
import "src/contracts/interfaces/IDurationVaultStrategy.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/ISignatureUtilsMixin.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IRewardsCoordinator.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/mocks/RewardsCoordinatorMock.sol";
import "src/test/mocks/StrategyFactoryMock.sol";

contract StrategyManagerDurationUnitTests is EigenLayerUnitTestSetup, IStrategyManagerEvents {
    StrategyManager public strategyManagerImplementation;
    StrategyManager public strategyManager;

    DurationVaultStrategy public durationVaultImplementation;
    IDurationVaultStrategy public durationVault;

    StrategyFactoryMock public strategyFactoryMock;
    ERC20PresetFixedSupply public underlyingToken;

    address internal constant STAKER = address(0xBEEF);
    uint internal constant INITIAL_SUPPLY = 1e36;
    address internal constant OPERATOR_SET_AVS = address(0xF00D);
    uint32 internal constant OPERATOR_SET_ID = 7;
    address internal constant DELEGATION_APPROVER = address(0xCAFE);
    uint32 internal constant OPERATOR_ALLOCATION_DELAY = 3;
    string internal constant OPERATOR_METADATA_URI = "ipfs://strategy-manager-vault";
    bytes internal constant REGISTRATION_DATA = hex"DEADBEEF";

    function setUp() public override {
        EigenLayerUnitTestSetup.setUp();
        // Align allocation delay with expected OPERATOR_ALLOCATION_DELAY for duration vault registration.
        delegationManagerMock.setMinWithdrawalDelayBlocks(OPERATOR_ALLOCATION_DELAY - 1);

        strategyManagerImplementation = new StrategyManager(
            IAllocationManager(address(allocationManagerMock)), IDelegationManager(address(delegationManagerMock)), pauserRegistry, "9.9.9"
        );

        strategyManager = StrategyManager(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyManagerImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyManager.initialize.selector, address(this), address(this), 0)
                )
            )
        );

        underlyingToken = new ERC20PresetFixedSupply("Mock Token", "MOCK", INITIAL_SUPPLY, address(this));
        rewardsCoordinatorMock = new RewardsCoordinatorMock();
        strategyFactoryMock = new StrategyFactoryMock();

        durationVaultImplementation = new DurationVaultStrategy(
            IStrategyManager(address(strategyManager)),
            pauserRegistry,
            IDelegationManager(address(delegationManagerMock)),
            IAllocationManager(address(allocationManagerMock)),
            IRewardsCoordinator(address(rewardsCoordinatorMock)),
            IStrategyFactory(address(strategyFactoryMock))
        );

        IDurationVaultStrategyTypes.VaultConfig memory cfg = IDurationVaultStrategyTypes.VaultConfig({
            underlyingToken: IERC20(address(underlyingToken)),
            vaultAdmin: address(this),
            arbitrator: address(this),
            duration: uint32(30 days),
            maxPerDeposit: 1_000_000 ether,
            stakeCap: 10_000_000 ether,
            metadataURI: "ipfs://duration-vault-test",
            operatorSet: OperatorSet({avs: OPERATOR_SET_AVS, id: OPERATOR_SET_ID}),
            operatorSetRegistrationData: REGISTRATION_DATA,
            delegationApprover: DELEGATION_APPROVER,
            operatorMetadataURI: OPERATOR_METADATA_URI
        });

        durationVault = IDurationVaultStrategy(
            address(
                new TransparentUpgradeableProxy(
                    address(durationVaultImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(DurationVaultStrategy.initialize.selector, cfg)
                )
            )
        );

        // Configure the mock to return the vault as a supported strategy in the operator set
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(address(durationVault));
        allocationManagerMock.setStrategiesInOperatorSet(OperatorSet({avs: OPERATOR_SET_AVS, id: OPERATOR_SET_ID}), strategies);

        IStrategy[] memory whitelist = new IStrategy[](1);
        whitelist[0] = IStrategy(address(durationVault));

        cheats.prank(strategyManager.owner());
        strategyManager.addStrategiesToDepositWhitelist(whitelist);
    }

    function testDepositIntoDurationVaultViaStrategyManager() public {
        uint amount = 10 ether;
        underlyingToken.transfer(STAKER, amount);

        ISignatureUtilsMixinTypes.SignatureWithExpiry memory emptySig;
        cheats.prank(STAKER);
        delegationManagerMock.delegateTo(address(durationVault), emptySig, bytes32(0));

        cheats.startPrank(STAKER);
        underlyingToken.approve(address(strategyManager), amount);
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit Deposit(STAKER, IStrategy(address(durationVault)), amount);
        strategyManager.depositIntoStrategy(IStrategy(address(durationVault)), IERC20(address(underlyingToken)), amount);
        cheats.stopPrank();

        uint shares = strategyManager.stakerDepositShares(STAKER, IStrategy(address(durationVault)));
        assertEq(shares, amount, "staker shares mismatch");
    }

    function testDepositRevertsAfterVaultLock() public {
        durationVault.lock();

        uint amount = 5 ether;
        underlyingToken.transfer(STAKER, amount);

        cheats.startPrank(STAKER);
        underlyingToken.approve(address(strategyManager), amount);
        cheats.expectRevert(IDurationVaultStrategyErrors.DepositsLocked.selector);
        strategyManager.depositIntoStrategy(IStrategy(address(durationVault)), IERC20(address(underlyingToken)), amount);
        cheats.stopPrank();
    }

    function _depositFor(address staker, uint amount) internal {
        underlyingToken.transfer(staker, amount);
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory emptySig;
        cheats.prank(staker);
        delegationManagerMock.delegateTo(address(durationVault), emptySig, bytes32(0));
        cheats.startPrank(staker);
        underlyingToken.approve(address(strategyManager), amount);
        strategyManager.depositIntoStrategy(IStrategy(address(durationVault)), IERC20(address(underlyingToken)), amount);
        cheats.stopPrank();
    }

    function testWithdrawalsBlockedViaStrategyManagerBeforeMaturity() public {
        uint amount = 8 ether;
        _depositFor(STAKER, amount);
        durationVault.lock();

        uint shares = strategyManager.stakerDepositShares(STAKER, IStrategy(address(durationVault)));

        // Withdrawal blocking now happens at the queueing stage (removeDepositShares -> beforeRemoveShares),
        // not at the withdrawal execution stage (withdrawSharesAsTokens).
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert(IDurationVaultStrategyErrors.WithdrawalsLockedDuringAllocations.selector);
        strategyManager.removeDepositShares(STAKER, IStrategy(address(durationVault)), shares);
    }

    function testWithdrawalsAllowedAfterMaturity() public {
        uint amount = 6 ether;
        _depositFor(STAKER, amount);
        durationVault.lock();

        cheats.warp(block.timestamp + durationVault.duration() + 1);
        durationVault.markMatured();

        uint shares = strategyManager.stakerDepositShares(STAKER, IStrategy(address(durationVault)));

        cheats.prank(address(delegationManagerMock));
        strategyManager.removeDepositShares(STAKER, IStrategy(address(durationVault)), shares);

        uint balanceBefore = underlyingToken.balanceOf(STAKER);

        cheats.prank(address(delegationManagerMock));
        strategyManager.withdrawSharesAsTokens(STAKER, IStrategy(address(durationVault)), IERC20(address(underlyingToken)), shares);

        assertEq(strategyManager.stakerDepositShares(STAKER, IStrategy(address(durationVault))), 0, "shares should be zero after removal");
        assertEq(underlyingToken.balanceOf(STAKER), balanceBefore + amount, "staker did not receive withdrawn tokens");
    }
}
