// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../test/EigenLayerTestHelper.t.sol";
import "src/test/utils/Utils.sol";
import "./mocks/ERC20Mock.sol";

///@notice This set of tests shadow forks the contracts from M1, queues withdrawals, and tests the migration functionality
contract WithdrawalMigrationTests is EigenLayerTestHelper, Utils {
    // Pointers to M1 contracts
    address private constant _M1_EIGENLAYER_PROXY_ADMIN = 0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444;
    address private constant _M1_STRATEGY_MANAGER = 0x858646372CC42E1A627fcE94aa7A7033e7CF075A;
    address private constant _M1_DELEGATION_MANAGER = 0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A;
    address private constant _M1_EIGENPOD_MANAGER = 0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338;
    address private constant _M1_SLASHER = 0xD92145c07f8Ed1D392c1B88017934E301CC1c3Cd;
    address private constant _M1_UNPAUSER = 0x369e6F597e22EaB55fFb173C6d9cD234BD699111;
    address private constant _M1_PAUSER_REGISTRY = 0x0c431C66F4dE941d089625E5B423D00707977060;
    address private constant _M1_CBETH_STRATEGY = 0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc;
    address private constant _CBETH_ADDRESS = 0xBe9895146f7AF43049ca1c1AE358B0541Ea49704;

    // Block to fork from, must be before M1 upgrade
    uint256 private constant _M1_BLOCK_FORK = 18322601;

    IM1StrategyManager m1StrategyManager = IM1StrategyManager(_M1_STRATEGY_MANAGER);
    ProxyAdmin m1EigenLayerProxyAdmin = ProxyAdmin(_M1_EIGENLAYER_PROXY_ADMIN);
    DelegationManager m1DelegationManager = DelegationManager(_M1_DELEGATION_MANAGER);
    IStrategy beaconChainETHStrategy;
    IStrategy cbETHStrategy = IStrategy(_M1_CBETH_STRATEGY);
    IERC20 cbETH = IERC20(_CBETH_ADDRESS);

    function setUp() public override {
        vm.createSelectFork("https://eth.llamarpc.com", _M1_BLOCK_FORK);
        beaconChainETHStrategy = m1StrategyManager.beaconChainETHStrategy();
        // Unpause strategyManager
        cheats.prank(_M1_UNPAUSER);
        m1StrategyManager.unpause(0);
    }

    function testMigrateWithdrawalBeaconChainETHStrategy(uint128 amountGwei) public {
        // Queue a withdrawal
        (
            bytes32 withdrawalRootSM,
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal
        ) = _queueWithdrawalBeaconChainETH(amountGwei);

        // Assert that withdrawal is queued
        assertTrue(m1StrategyManager.withdrawalRootPending(withdrawalRootSM));

        // Upgrade, migrate, and check
        _upgradeAndMigrate(withdrawalRootSM, queuedWithdrawal);
    }

    function testMigrateQueuedWithdrawalStrategy(
        uint256 withdrawalAmount,
        uint256 depositAmount,
        bool undelegateIfPossible
    ) public {
        depositAmount = bound(
            depositAmount,
            1e3,
            IMaxDepositAmount(address(cbETHStrategy)).maxTotalDeposits() - cbETH.balanceOf(address(cbETHStrategy))
        );
        withdrawalAmount = bound(withdrawalAmount, 1, cbETHStrategy.underlyingToShares(depositAmount));

        // Deal tokens
        deal(address(cbETH), address(this), depositAmount);
        cbETH.approve(address(m1StrategyManager), depositAmount);

        // Queue withdrawal
        (
            bytes32 withdrawalRootSM,
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal
        ) = _queueWithdrawalTokenStrategy(address(this), depositAmount, withdrawalAmount, undelegateIfPossible);

        // Assert that withdrawal is queued
        assertTrue(m1StrategyManager.withdrawalRootPending(withdrawalRootSM));

        // Upgrade, migrate, and check
        _upgradeAndMigrate(withdrawalRootSM, queuedWithdrawal);
    }

    function _upgradeAndMigrate(
        bytes32 withdrawalRootSM,
        IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal
    ) internal {
        // Upgrade M1 Contracts
        _upgradeM1Contracts();

        // Format withdrawal struct as represented in delegationManager
        IDelegationManager.Withdrawal memory migratedWithdrawal = IDelegationManager.Withdrawal({
            staker: queuedWithdrawal.staker,
            delegatedTo: queuedWithdrawal.delegatedAddress,
            withdrawer: queuedWithdrawal.withdrawerAndNonce.withdrawer,
            nonce: queuedWithdrawal.withdrawerAndNonce.nonce,
            startBlock: queuedWithdrawal.withdrawalStartBlock,
            strategies: queuedWithdrawal.strategies,
            shares: queuedWithdrawal.shares
        });
        bytes32 withdrawalRootDM = m1DelegationManager.calculateWithdrawalRoot(migratedWithdrawal);
        uint256 nonceBefore = m1DelegationManager.cumulativeWithdrawalsQueued(queuedWithdrawal.staker);

        // Migrate withdrawal
        cheats.expectEmit(true, true, true, true);
        emit WithdrawalQueued(withdrawalRootDM, migratedWithdrawal);
        cheats.expectEmit(true, true, true, true);
        emit WithdrawalMigrated(withdrawalRootSM, withdrawalRootDM);
        IStrategyManager.DeprecatedStruct_QueuedWithdrawal[]
            memory queuedWithdrawals = new IStrategyManager.DeprecatedStruct_QueuedWithdrawal[](1);
        queuedWithdrawals[0] = queuedWithdrawal;
        m1DelegationManager.migrateQueuedWithdrawals(queuedWithdrawals);

        // Assertions
        assertFalse(m1StrategyManager.withdrawalRootPending(withdrawalRootSM));
        assertTrue(m1DelegationManager.pendingWithdrawals(withdrawalRootDM));
        uint256 nonceAfter = m1DelegationManager.cumulativeWithdrawalsQueued(queuedWithdrawal.staker);
        assertTrue(nonceAfter == nonceBefore + 1);
    }

    function testMigrateMultipleQueuedWithdrawals() public {
        // Get a beaconChain queued withdrawal
        (
            bytes32 withdrawalRootSM1,
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal1
        ) = _queueWithdrawalBeaconChainETH(1e9);

        // Get a tokenStrategy queued withdrawal
        deal(address(cbETH), address(this), 1e19);
        cbETH.approve(address(m1StrategyManager), 1e19);
        (
            bytes32 withdrawalRootSM2,
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal2
        ) = _queueWithdrawalTokenStrategy(address(this), 1e19, 5e18, false);

        // Assert Pending Withdrawals
        assertTrue(m1StrategyManager.withdrawalRootPending(withdrawalRootSM1));
        assertTrue(m1StrategyManager.withdrawalRootPending(withdrawalRootSM2));

        // Upgrade M1 Contracts
        _upgradeM1Contracts();

        // Format withdrawal structs as represented in delegationManager
        IDelegationManager.Withdrawal memory migratedWithdrawal1 = IDelegationManager.Withdrawal({
            staker: queuedWithdrawal1.staker,
            delegatedTo: queuedWithdrawal1.delegatedAddress,
            withdrawer: queuedWithdrawal1.withdrawerAndNonce.withdrawer,
            nonce: queuedWithdrawal1.withdrawerAndNonce.nonce,
            startBlock: queuedWithdrawal1.withdrawalStartBlock,
            strategies: queuedWithdrawal1.strategies,
            shares: queuedWithdrawal1.shares
        });

        IDelegationManager.Withdrawal memory migratedWithdrawal2 = IDelegationManager.Withdrawal({
            staker: queuedWithdrawal2.staker,
            delegatedTo: queuedWithdrawal2.delegatedAddress,
            withdrawer: queuedWithdrawal2.withdrawerAndNonce.withdrawer,
            nonce: queuedWithdrawal2.withdrawerAndNonce.nonce,
            startBlock: queuedWithdrawal2.withdrawalStartBlock,
            strategies: queuedWithdrawal2.strategies,
            shares: queuedWithdrawal2.shares
        });

        bytes32 withdrawalRootDM1 = m1DelegationManager.calculateWithdrawalRoot(migratedWithdrawal1);
        bytes32 withdrawalRootDM2 = m1DelegationManager.calculateWithdrawalRoot(migratedWithdrawal2);

        uint256 nonceBefore = m1DelegationManager.cumulativeWithdrawalsQueued(queuedWithdrawal1.staker);

        // Migrate withdrawals
        IStrategyManager.DeprecatedStruct_QueuedWithdrawal[]
            memory queuedWithdrawals = new IStrategyManager.DeprecatedStruct_QueuedWithdrawal[](2);
        queuedWithdrawals[0] = queuedWithdrawal1;
        queuedWithdrawals[1] = queuedWithdrawal2;
        m1DelegationManager.migrateQueuedWithdrawals(queuedWithdrawals);

        // Assertions
        assertFalse(m1StrategyManager.withdrawalRootPending(withdrawalRootSM1));
        assertFalse(m1StrategyManager.withdrawalRootPending(withdrawalRootSM2));
        assertTrue(m1DelegationManager.pendingWithdrawals(withdrawalRootDM1));
        assertTrue(m1DelegationManager.pendingWithdrawals(withdrawalRootDM2));
        uint256 nonceAfter = m1DelegationManager.cumulativeWithdrawalsQueued(queuedWithdrawal1.staker);
        assertTrue(nonceAfter == nonceBefore + 2);
    }

    function _queueWithdrawalTokenStrategy(
        address staker,
        uint256 depositAmount,
        uint256 withdrawalAmount,
        bool undelegateIfPossible
    ) internal returns (bytes32, IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory) {
        // Deposit
        _depositIntoTokenStrategy(staker, depositAmount);

        // Setup payload
        (
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal,
            /*IERC20[] memory tokensArray*/,
            bytes32 withdrawalRootSM
        ) = _setUpQueuedWithdrawalStructSingleStrat(
                /*staker*/ address(this),
                /*withdrawer*/ address(this),
                cbETH,
                cbETHStrategy,
                withdrawalAmount
            );

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;
        m1StrategyManager.queueWithdrawal(
            strategyIndexes,
            queuedWithdrawal.strategies,
            queuedWithdrawal.shares,
            /*withdrawer*/ address(this),
            undelegateIfPossible
        );

        return (withdrawalRootSM, queuedWithdrawal);
    }

    function _depositIntoTokenStrategy(address staker, uint256 amount) internal {
        // filter out zero case since it will revert with "m1StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // sanity check / filter
        cheats.assume(amount <= cbETH.balanceOf(address(this)));

        uint256 sharesBefore = m1StrategyManager.stakerStrategyShares(staker, cbETHStrategy);
        uint256 stakerStrategyListLengthBefore = m1StrategyManager.stakerStrategyListLength(staker);

        cheats.startPrank(staker);
        uint256 shares = m1StrategyManager.depositIntoStrategy(cbETHStrategy, cbETH, amount);
        cheats.stopPrank();

        uint256 sharesAfter = m1StrategyManager.stakerStrategyShares(staker, cbETHStrategy);
        uint256 stakerStrategyListLengthAfter = m1StrategyManager.stakerStrategyListLength(staker);

        require(sharesAfter == sharesBefore + shares, "sharesAfter != sharesBefore + shares");
        if (sharesBefore == 0) {
            require(
                stakerStrategyListLengthAfter == stakerStrategyListLengthBefore + 1,
                "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
            );
            require(
                m1StrategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) == cbETHStrategy,
                "m1StrategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) != strategy"
            );
        }
    }

    function _queueWithdrawalBeaconChainETH(
        uint128 amountGwei
    ) internal returns (bytes32, IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory) {
        // scale fuzzed amount up to be a whole amount of GWEI
        uint256 amount = uint256(amountGwei) * 1e9;
        address staker = address(this);
        address withdrawer = staker;
        IStrategy strategy = beaconChainETHStrategy;
        IERC20 dummyToken;

        _depositBeaconChainETH(staker, amount);

        bool undelegateIfPossible = false;
        (
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal /*IERC20[] memory tokensArray*/,
            ,
            bytes32 withdrawalRootSM
        ) = _setUpQueuedWithdrawalStructSingleStrat(staker, withdrawer, dummyToken, strategy, amount);

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;
        m1StrategyManager.queueWithdrawal(
            strategyIndexes,
            queuedWithdrawal.strategies,
            queuedWithdrawal.shares,
            withdrawer,
            undelegateIfPossible
        );

        return (withdrawalRootSM, queuedWithdrawal);
    }

    function _depositBeaconChainETH(address staker, uint256 amount) internal {
        // filter out zero case since it will revert with "m1StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        uint256 sharesBefore = m1StrategyManager.stakerStrategyShares(staker, beaconChainETHStrategy);

        cheats.startPrank(address(m1StrategyManager.eigenPodManager()));
        m1StrategyManager.depositBeaconChainETH(staker, amount);
        cheats.stopPrank();

        uint256 sharesAfter = m1StrategyManager.stakerStrategyShares(staker, beaconChainETHStrategy);
        require(sharesAfter == sharesBefore + amount, "sharesAfter != sharesBefore + amount");
    }

    function _setUpQueuedWithdrawalStructSingleStrat(
        address staker,
        address withdrawer,
        IERC20 token,
        IStrategy strategy,
        uint256 shareAmount
    )
        internal
        view
        returns (
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal,
            IERC20[] memory tokensArray,
            bytes32 withdrawalRoot
        )
    {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = shareAmount;
        IStrategyManager.DeprecatedStruct_WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager
            .DeprecatedStruct_WithdrawerAndNonce({
                withdrawer: withdrawer,
                nonce: uint96(m1StrategyManager.numWithdrawalsQueued(staker))
            });
        queuedWithdrawal = IStrategyManager.DeprecatedStruct_QueuedWithdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            staker: staker,
            withdrawerAndNonce: withdrawerAndNonce,
            withdrawalStartBlock: uint32(block.number),
            delegatedAddress: m1StrategyManager.delegation().delegatedTo(staker)
        });
        // calculate the withdrawal root
        withdrawalRoot = m1StrategyManager.calculateWithdrawalRoot(queuedWithdrawal);
        return (queuedWithdrawal, tokensArray, withdrawalRoot);
    }

    function _upgradeM1Contracts() internal {
        // Deploy Implementation Contracts
        StrategyManager strategyManagerImplementation = new StrategyManager(
            IDelegationManager(_M1_DELEGATION_MANAGER),
            IEigenPodManager(_M1_EIGENPOD_MANAGER),
            ISlasher(_M1_SLASHER)
        );
        DelegationManager delegationManagerImplementation = new DelegationManager(
            IStrategyManager(_M1_STRATEGY_MANAGER),
            ISlasher(_M1_SLASHER),
            IEigenPodManager(_M1_EIGENPOD_MANAGER)
        );
        // Upgrade
        cheats.startPrank(m1EigenLayerProxyAdmin.owner());
        m1EigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(m1StrategyManager))),
            address(strategyManagerImplementation)
        );
        m1EigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(m1DelegationManager))),
            address(delegationManagerImplementation)
        );
        cheats.stopPrank();
    }

    // Events
    event WithdrawalMigrated(bytes32 oldWithdrawalRoot, bytes32 newWithdrawalRoot);
    event WithdrawalQueued(bytes32 withdrawalRoot, IDelegationManager.Withdrawal withdrawal);
}

interface IM1StrategyManager is IStrategyManager, IPausable {
    function depositBeaconChainETH(address staker, uint256 amount) external;

    function withdrawalRootPending(bytes32 withdrawalRoot) external view returns (bool);

    function numWithdrawalsQueued(address staker) external view returns (uint256);

    function beaconChainETHStrategy() external view returns (IStrategy);

    function queueWithdrawal(
        uint256[] calldata strategyIndexes,
        IStrategy[] calldata strategies,
        uint256[] calldata shares,
        address withdrawer,
        bool undelegateIfPossible
    ) external returns (bytes32);

    function strategyWhitelister() external view returns (address);

    function stakerStrategyList(address staker, uint256 index) external view returns (IStrategy);
}

interface IMaxDepositAmount {
    function maxPerDeposit() external view returns (uint256);

    function maxTotalDeposits() external view returns (uint256);
}
