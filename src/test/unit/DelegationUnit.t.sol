// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/libraries/SlashingLib.sol";
import "src/test/utils/ArrayLib.sol";
import "src/test/harnesses/DelegationManagerHarness.sol";

/**
 * @notice Unit testing of the DelegationManager contract. Withdrawals are tightly coupled
 * with EigenPodManager and StrategyManager and are part of integration tests.
 * Contracts tested: DelegationManager
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract DelegationManagerUnitTests is EigenLayerUnitTestSetup, IDelegationManagerEvents, IDelegationManagerErrors {
    using SlashingLib for *;
    using ArrayLib for *;
    using Math for *;

    /// -----------------------------------------------------------------------
    /// Contracts and Mocks
    /// -----------------------------------------------------------------------

    DelegationManagerHarness delegationManager;
    DelegationManagerHarness delegationManagerImplementation;
    StrategyBase strategyImplementation;
    StrategyBase strategyMock;
    IERC20 tokenMock;
    uint tokenMockInitialSupply = 10e50;

    /// -----------------------------------------------------------------------
    /// Constants
    /// -----------------------------------------------------------------------

    uint32 constant MIN_WITHDRAWAL_DELAY_BLOCKS = 126_000; // 17.5 days in blocks
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;
    uint8 internal constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;
    uint8 internal constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;
    // Fuzz bound checks
    uint constant MIN_FUZZ_SHARES = 10_000;
    uint constant MIN_FUZZ_MAGNITUDE = 10_000;
    uint constant APPROX_REL_DIFF = 1e8; // 0.0.0000000100000000% relative difference for assertion checks. Needed due to rounding errors
    // Max shares in a strategy, see StrategyBase.sol
    uint constant MAX_STRATEGY_SHARES = 1e38 - 1;
    uint constant MAX_ETH_SUPPLY = 120_400_000 ether;

    /// -----------------------------------------------------------------------
    /// Defaults & Mappings for Stack too deep errors
    /// -----------------------------------------------------------------------

    // Delegation signer
    uint delegationSignerPrivateKey = uint(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
    address defaultApprover = cheats.addr(delegationSignerPrivateKey);
    uint stakerPrivateKey = uint(123_456_789);
    uint staker2PrivateKey = uint(234_567_891);
    address defaultStaker = cheats.addr(stakerPrivateKey);
    address defaultStaker2 = cheats.addr(staker2PrivateKey);
    address defaultOperator = address(this);
    address defaultOperator2 = address(0x123);
    address defaultAVS = address(this);
    string emptyStringForMetadataURI;
    ISignatureUtilsMixinTypes.SignatureWithExpiry emptyApproverSignatureAndExpiry;
    bytes32 emptySalt;
    // Helper to use in storage
    DepositScalingFactor dsf;
    uint stakerDSF;

    /// @notice mappings used to handle duplicate entries in fuzzed address array input
    mapping(address => uint) public totalSharesForStrategyInArray;
    mapping(IStrategy => uint) public totalSharesDecreasedForStrategy;
    mapping(IStrategy => uint) public delegatedSharesBefore;
    mapping(address => uint) public stakerDepositShares;
    // Keep track of queued withdrawals
    mapping(address => Withdrawal[]) public stakerQueuedWithdrawals;

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        delegationManager =
            DelegationManagerHarness(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));

        // Redeploy StrategyManagerMock with DM
        strategyManagerMock = StrategyManagerMock(payable(address(new StrategyManagerMock(delegationManager))));

        // Deploy DelegationManager implementation and upgrade proxy
        delegationManagerImplementation = new DelegationManagerHarness(
            IStrategyManager(address(strategyManagerMock)),
            IEigenPodManager(address(eigenPodManagerMock)),
            IAllocationManager(address(allocationManagerMock)),
            pauserRegistry,
            IPermissionController(address(permissionController)),
            MIN_WITHDRAWAL_DELAY_BLOCKS
        );

        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(delegationManager))),
            address(delegationManagerImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                address(this),
                0 // 0 is initial paused status
            )
        );

        // Deploy mock token and strategy
        tokenMock = new ERC20PresetFixedSupply("Mock Token", "MOCK", tokenMockInitialSupply, address(this));
        strategyImplementation = new StrategyBase(IStrategyManager(address(strategyManagerMock)), pauserRegistry, "v9.9.9");
        strategyMock = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, tokenMock)
                )
            )
        );

        // Roll blocks forward so that block.number - MIN_WITHDRAWAL_DELAY_BLOCKS doesn't revert
        // in _getSlashableSharesInQueue
        cheats.roll(MIN_WITHDRAWAL_DELAY_BLOCKS + 1);

        // Exclude delegation manager from fuzzed tests
        isExcludedFuzzAddress[address(delegationManager)] = true;
        isExcludedFuzzAddress[address(strategyManagerMock)] = true;
        isExcludedFuzzAddress[defaultApprover] = true;
    }

    /**
     * INTERNAL / HELPER FUNCTIONS
     */

    /**
     * @notice internal function to deploy mock tokens and strategies and have the staker deposit into them.
     * Since we are mocking the strategyManager we call strategyManagerMock.setDeposits so that when
     * DelegationManager calls getDeposits, we can have these share amounts returned.
     */
    function _deployAndDepositIntoStrategies(address staker, uint[] memory sharesAmounts, bool depositBeaconChainShares)
        internal
        returns (IStrategy[] memory)
    {
        uint numStrats = sharesAmounts.length;
        IStrategy[] memory strategies = new IStrategy[](numStrats);
        for (uint8 i = 0; i < numStrats; i++) {
            // If depositing beaconShares, then for last index of shareAmount, set shares into EPM instead
            if (depositBeaconChainShares && i == numStrats - 1) {
                strategies[i] = beaconChainETHStrategy;
                eigenPodManagerMock.setPodOwnerShares(staker, int(sharesAmounts[numStrats - 1]));
                break;
            }
            ERC20PresetFixedSupply token = new ERC20PresetFixedSupply(
                string(abi.encodePacked("Mock Token ", i)), string(abi.encodePacked("MOCK", i)), tokenMockInitialSupply, address(this)
            );
            strategies[i] = StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(strategyImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, token)
                    )
                )
            );
            strategyManagerMock.addDeposit(staker, strategies[i], sharesAmounts[i]);
        }
        return strategies;
    }

    /**
     * @notice internal function to deploy mock tokens and strategies and have the staker deposit into them.
     * Since we are mocking the strategyManager we call strategyManagerMock.setDeposits so that when
     * DelegationManager calls getDeposits, we can have these share amounts returned.
     */
    function _depositIntoStrategies(address staker, IStrategy[] memory strategies, uint[] memory sharesAmounts) internal {
        uint numStrats = strategies.length;
        require(numStrats == sharesAmounts.length, "DelegationManagerUnitTests: length mismatch");
        for (uint8 i = 0; i < numStrats; i++) {
            // If depositing beaconShares, then for last index of shareAmount, set shares into EPM instead
            if (strategies[i] == beaconChainETHStrategy) eigenPodManagerMock.setPodOwnerShares(staker, int(sharesAmounts[i]));
            else strategyManagerMock.addDeposit(staker, strategies[i], sharesAmounts[i]);
        }
    }

    /**
     * @notice internal function for calculating a signature from the delegationSigner corresponding to `_delegationSignerPrivateKey`, approving
     * the `staker` to delegate to `operator`, with the specified `salt`, and expiring at `expiry`.
     */
    function _getApproverSignature(uint _delegationSignerPrivateKey, address staker, address operator, bytes32 salt, uint expiry)
        internal
        view
        returns (ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry)
    {
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateDelegationApprovalDigestHash(
                staker, operator, delegationManager.delegationApprover(operator), salt, expiry
            );
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(_delegationSignerPrivateKey, digestHash);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }
        return approverSignatureAndExpiry;
    }

    // @notice Assumes operator does not have a delegation approver & staker != approver
    function _delegateToOperatorWhoAcceptsAllStakers(address staker, address operator) internal {
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;
        cheats.prank(staker);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
    }

    function _delegateToOperatorWhoRequiresSig(address staker, address operator, bytes32 salt) internal {
        uint expiry = type(uint).max;
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);
        cheats.prank(staker);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
    }

    function _delegateToOperatorWhoRequiresSig(address staker, address operator) internal {
        _delegateToOperatorWhoRequiresSig(staker, operator, emptySalt);
    }

    function _registerOperatorWithBaseDetails(address operator) internal {
        _registerOperator(operator, address(0), emptyStringForMetadataURI);
    }

    function _registerOperatorWithDelegationApprover(address operator) internal {
        _registerOperator(operator, defaultApprover, emptyStringForMetadataURI);
    }

    function _registerOperatorWith1271DelegationApprover(address operator) internal returns (ERC1271WalletMock) {
        address delegationSigner = defaultApprover;
        /**
         * deploy a ERC1271WalletMock contract with the `delegationSigner` address as the owner,
         * so that we can create valid signatures from the `delegationSigner` for the contract to check when called
         */
        ERC1271WalletMock wallet = new ERC1271WalletMock(delegationSigner);
        _registerOperator(operator, address(wallet), emptyStringForMetadataURI);

        return wallet;
    }

    function _registerOperator(address operator, address delegationApprover, string memory metadataURI)
        internal
        filterFuzzedAddressInputs(operator)
    {
        cheats.prank(operator);
        delegationManager.registerAsOperator(delegationApprover, 0, metadataURI);
    }

    /**
     * @notice Using this helper function to fuzz withdrawalAmounts since fuzzing two dynamic sized arrays of equal lengths
     * reject too many inputs.
     */
    function _fuzzDepositWithdrawalAmounts(Randomness r, uint32 numStrategies)
        internal
        returns (
            uint[] memory depositAmounts,
            uint[] memory withdrawalAmounts,
            uint64[] memory prevMagnitudes,
            uint64[] memory newMagnitudes
        )
    {
        withdrawalAmounts = new uint[](numStrategies);
        depositAmounts = new uint[](numStrategies);
        prevMagnitudes = new uint64[](numStrategies);
        newMagnitudes = new uint64[](numStrategies);
        for (uint i = 0; i < numStrategies; i++) {
            depositAmounts[i] = r.Uint256(1, MAX_STRATEGY_SHARES);
            // generate withdrawal amount within range s.t withdrawAmount <= depositAmount
            withdrawalAmounts[i] = r.Uint256(1, depositAmounts[i]);

            prevMagnitudes[i] = r.Uint64(2, WAD);
            newMagnitudes[i] = r.Uint64(1, prevMagnitudes[i]);
        }
        return (depositAmounts, withdrawalAmounts, prevMagnitudes, newMagnitudes);
    }

    function _setUpQueueWithdrawalsSingleStrat(address staker, IStrategy strategy, uint depositSharesToWithdraw)
        internal
        view
        returns (QueuedWithdrawalParams[] memory, Withdrawal memory, bytes32)
    {
        IStrategy[] memory strategyArray = strategy.toArray();
        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        {
            uint[] memory withdrawalAmounts = depositSharesToWithdraw.toArrayU256();
            queuedWithdrawalParams[0] =
                QueuedWithdrawalParams({strategies: strategyArray, depositShares: withdrawalAmounts, __deprecated_withdrawer: address(0)});
        }

        // Get scaled shares to withdraw
        uint[] memory scaledSharesArray = _getScaledShares(staker, strategy, depositSharesToWithdraw).toArrayU256();

        Withdrawal memory withdrawal = Withdrawal({
            staker: staker,
            delegatedTo: delegationManager.delegatedTo(staker),
            withdrawer: staker,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategyArray,
            scaledShares: scaledSharesArray
        });
        bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);

        return (queuedWithdrawalParams, withdrawal, withdrawalRoot);
    }

    function _setUpQueueWithdrawals(address staker, IStrategy[] memory strategies, uint[] memory depositWithdrawalAmounts)
        internal
        view
        returns (QueuedWithdrawalParams[] memory, Withdrawal memory, bytes32)
    {
        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        {
            queuedWithdrawalParams[0] = QueuedWithdrawalParams({
                strategies: strategies,
                depositShares: depositWithdrawalAmounts,
                __deprecated_withdrawer: address(0)
            });
        }

        // Get scaled shares to withdraw
        uint[] memory scaledSharesArray = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            scaledSharesArray[i] = _getScaledShares(staker, strategies[i], depositWithdrawalAmounts[i]);
        }

        Withdrawal memory withdrawal = Withdrawal({
            staker: staker,
            delegatedTo: delegationManager.delegatedTo(staker),
            withdrawer: staker,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: scaledSharesArray
        });
        bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);

        return (queuedWithdrawalParams, withdrawal, withdrawalRoot);
    }

    /// @notice Call queue withdrawals and push the Withdrawal to storage for testing purposes and
    /// later assertions
    function _queueWithdrawals(address staker, QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal)
        internal
    {
        stakerQueuedWithdrawals[staker].push(withdrawal);
        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function _getScaledShares(address staker, IStrategy strategy, uint depositSharesToWithdraw) internal view returns (uint) {
        DepositScalingFactor memory _dsf = DepositScalingFactor(delegationManager.depositScalingFactor(staker, strategy));

        return _dsf.scaleForQueueWithdrawal(depositSharesToWithdraw);
    }

    /// @notice get the shares expected to be withdrawn given the staker, strategy, maxMagnitude, and depositSharesToWithdraw
    function _getWithdrawableShares(
        address staker,
        IStrategy[] memory strategies,
        uint64[] memory maxMagnitudes,
        uint[] memory depositSharesToWithdraw
    ) internal view returns (uint[] memory) {
        require(strategies.length == depositSharesToWithdraw.length, "DelegationManagerUnitTests: length mismatch");
        uint[] memory withdrawnShares = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            uint slashingFactor = _getSlashingFactor(staker, strategies[i], maxMagnitudes[i]);
            withdrawnShares[i] = _calcWithdrawableShares(
                depositSharesToWithdraw[i], delegationManager.depositScalingFactor(staker, strategies[i]), slashingFactor
            );
        }
        return withdrawnShares;
    }

    function _getSlashingFactor(address staker, IStrategy strategy, uint64 operatorMaxMagnitude) internal view returns (uint) {
        if (strategy == beaconChainETHStrategy) {
            uint64 beaconChainSlashingFactor = eigenPodManagerMock.beaconChainSlashingFactor(staker);
            return operatorMaxMagnitude.mulWad(beaconChainSlashingFactor);
        }

        return operatorMaxMagnitude;
    }

    /**
     * Deploy and deposit staker into a single strategy, then set up a queued withdrawal for the staker
     * Assumptions:
     * - operator is already a registered operator.
     * - withdrawalAmount <= depositAmount
     */
    function _setUpCompleteQueuedWithdrawalSingleStrat(
        address staker,
        uint depositAmount,
        uint withdrawalAmount,
        bool isBeaconChainStrategy
    ) internal returns (Withdrawal memory, IERC20[] memory, bytes32) {
        uint[] memory depositAmounts = new uint[](1);
        depositAmounts[0] = depositAmount;
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts, isBeaconChainStrategy);
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: staker, strategy: strategies[0], depositSharesToWithdraw: withdrawalAmount});

        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
        // Set the current deposits to be the depositAmount - withdrawalAmount
        uint[] memory currentAmounts = uint(depositAmount - withdrawalAmount).toArrayU256();
        strategyManagerMock.setDeposits(staker, strategies, currentAmounts);

        IERC20[] memory tokens = new IERC20[](strategies.length);
        for (uint i = 0; i < tokens.length; i++) {
            tokens[i] = strategies[i].underlyingToken();
        }

        return (withdrawal, tokens, withdrawalRoot);
    }

    /**
     * Deploy and deposit staker into a single strategy, then set up multiple queued withdrawals for the staker
     * Assumptions:
     * - operator is already a registered operator.
     * - total deposit amount = depositAmount * numWithdrawals
     * - this will fully withdraw from the single strategy
     */
    function _setUpCompleteQueuedWithdrawalsSingleStrat(address staker, uint depositAmount, uint numWithdrawals)
        internal
        returns (Withdrawal[] memory withdrawals, IERC20[][] memory tokens, bytes32[] memory withdrawalRoots)
    {
        uint[] memory depositAmounts = new uint[](1);
        depositAmounts[0] = depositAmount * numWithdrawals;
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts, false);

        withdrawals = new Withdrawal[](numWithdrawals);
        tokens = new IERC20[][](numWithdrawals);
        withdrawalRoots = new bytes32[](numWithdrawals);

        for (uint i = 0; i < numWithdrawals; i++) {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
                _setUpQueueWithdrawalsSingleStrat({staker: staker, strategy: strategies[0], depositSharesToWithdraw: depositAmount});

            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);

            withdrawals[i] = withdrawal;
            tokens[i] = new IERC20[](1);
            tokens[i][0] = strategies[0].underlyingToken();
            withdrawalRoots[i] = withdrawalRoot;
        }

        {
            // Set the current deposits to be 0
            uint[] memory currentAmounts = new uint[](1);
            currentAmounts[0] = 0;
            strategyManagerMock.setDeposits(staker, strategies, currentAmounts);
        }

        return (withdrawals, tokens, withdrawalRoots);
    }

    /**
     * Deploy and deposit staker into strategies, then set up a queued withdrawal for the staker
     * Assumptions:
     * - operator is already a registered operator.
     * - for each i, withdrawalAmount[i] <= depositAmount[i] (see filterFuzzedDepositWithdrawInputs above)
     */
    function _setUpCompleteQueuedWithdrawal(
        address staker,
        uint[] memory depositAmounts,
        uint[] memory withdrawalAmounts,
        bool depositBeaconChainShares
    ) internal returns (Withdrawal memory, IERC20[] memory, bytes32) {
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts, depositBeaconChainShares);

        IERC20[] memory tokens = new IERC20[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            tokens[i] = strategies[i].underlyingToken();
        }

        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawals({staker: staker, strategies: strategies, depositWithdrawalAmounts: withdrawalAmounts});

        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        return (withdrawal, tokens, withdrawalRoot);
    }

    function _setOperatorMagnitude(address operator, IStrategy strategy, uint64 magnitude) internal {
        allocationManagerMock.setMaxMagnitude(operator, strategy, magnitude);
    }

    function _setNewBeaconChainSlashingFactor(address staker, int beaconShares, uint sharesDecrease)
        internal
        returns (uint64 prevBeaconSlashingFactor, uint64 newBeaconSlashingFactor)
    {
        uint newRestakedBalanceWei = uint(beaconShares) - sharesDecrease;
        prevBeaconSlashingFactor = eigenPodManagerMock.beaconChainSlashingFactor(staker);
        newBeaconSlashingFactor = uint64(prevBeaconSlashingFactor.mulDiv(newRestakedBalanceWei, uint(beaconShares)));
        eigenPodManagerMock.setBeaconChainSlashingFactor(staker, newBeaconSlashingFactor);
    }

    function _decreaseBeaconChainShares(address staker, int beaconShares, uint sharesDecrease)
        internal
        returns (uint64 prevBeaconSlashingFactor, uint64 newBeaconSlashingFactor)
    {
        (prevBeaconSlashingFactor, newBeaconSlashingFactor) = _setNewBeaconChainSlashingFactor(staker, beaconShares, sharesDecrease);

        cheats.prank(address(eigenPodManagerMock));
        delegationManager.decreaseDelegatedShares({
            staker: staker,
            curDepositShares: uint(beaconShares),
            beaconChainSlashingFactorDecrease: prevBeaconSlashingFactor - newBeaconSlashingFactor
        });
    }

    /// -----------------------------------------------------------------------
    /// Event helpers
    /// -----------------------------------------------------------------------

    struct RegisterAsOperatorEmitStruct {
        address operator;
        address delegationApprover;
        string metadataURI;
    }

    function _registerOperator_expectEmit(RegisterAsOperatorEmitStruct memory params) internal {
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DelegationApproverUpdated(params.operator, params.delegationApprover);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(params.operator, params.operator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorRegistered(params.operator, params.delegationApprover);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(params.operator, params.metadataURI);
    }

    struct DelegateToEmitStruct {
        address staker;
        address operator;
        IStrategy[] strategies;
        uint[] depositShares;
        uint[] depositScalingFactors;
    }

    function _delegateTo_expectEmit(DelegateToEmitStruct memory params) internal {
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(params.staker, params.operator);
        for (uint i = 0; i < params.strategies.length; i++) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(params.staker, params.strategies[i], params.depositScalingFactors[i]);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(params.operator, params.staker, params.strategies[i], params.depositShares[i]);
        }
    }

    struct DelegateToSingleStratEmitStruct {
        address staker;
        address operator;
        IStrategy strategy;
        uint depositShares;
        uint depositScalingFactor;
    }

    function _delegateTo_expectEmit_singleStrat(DelegateToSingleStratEmitStruct memory params) internal {
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(params.staker, params.operator);
        if (params.depositShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(params.staker, params.strategy, params.depositScalingFactor);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(params.operator, params.staker, params.strategy, params.depositShares);
        }
    }

    struct UndelegateEmitStruct {
        address staker;
        address operator;
        IStrategy strategy;
        uint depositSharesQueued;
        uint operatorSharesDecreased;
        Withdrawal withdrawal;
        bytes32 withdrawalRoot;
        uint depositScalingFactor;
        bool forceUndelegated;
    }

    /// @notice Assumes only single strategy for staker being withdrawn, only checks for single strategy if
    /// param.strategy address is not 0x0
    function _undelegate_expectEmit_singleStrat(UndelegateEmitStruct memory params) internal {
        if (params.forceUndelegated) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit StakerForceUndelegated(params.staker, params.operator);
        }
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(params.staker, params.operator);

        if (address(params.strategy) != address(0)) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesDecreased(params.operator, params.staker, params.strategy, params.operatorSharesDecreased);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit SlashingWithdrawalQueued(params.withdrawalRoot, params.withdrawal, params.operatorSharesDecreased.toArrayU256());
        }
    }

    struct IncreaseDelegatedSharesEmitStruct {
        address staker;
        address operator;
        IStrategy strategy;
        uint sharesToIncrease;
        uint depositScalingFactor;
    }

    function _increaseDelegatedShares_expectEmit(IncreaseDelegatedSharesEmitStruct memory params) internal {
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(params.staker, params.strategy, params.depositScalingFactor);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(params.operator, params.staker, params.strategy, params.sharesToIncrease);
    }

    struct DecreaseDelegatedSharesEmitStruct {
        address staker;
        address operator;
        uint sharesToDecrease;
    }

    function _decreaseDelegatedShares_expectEmit(DecreaseDelegatedSharesEmitStruct memory params) internal {
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(params.operator, params.staker, beaconChainETHStrategy, params.sharesToDecrease);
    }

    struct QueueWithdrawalsEmitStruct {
        address staker;
        address operator;
        QueuedWithdrawalParams[] queuedWithdrawalParams;
        Withdrawal withdrawal;
        bytes32 withdrawalRoot;
    }

    function _queueWithdrawals_expectEmit(QueueWithdrawalsEmitStruct memory params) internal {
        for (uint i = 0; i < params.queuedWithdrawalParams.length; i++) {
            uint[] memory sharesToWithdraw = new uint[](params.queuedWithdrawalParams[i].strategies.length);
            for (uint j = 0; j < params.queuedWithdrawalParams[i].strategies.length; j++) {
                uint depositScalingFactor =
                    delegationManager.depositScalingFactor(defaultStaker, params.queuedWithdrawalParams[i].strategies[j]);
                uint newMaxMagnitude =
                    allocationManagerMock.getMaxMagnitudes(params.operator, params.queuedWithdrawalParams[i].strategies)[j];
                sharesToWithdraw[j] =
                    _calcWithdrawableShares(params.queuedWithdrawalParams[i].depositShares[j], depositScalingFactor, newMaxMagnitude);

                cheats.expectEmit(true, true, true, true, address(delegationManager));
                emit OperatorSharesDecreased(
                    params.operator, params.staker, params.queuedWithdrawalParams[i].strategies[j], sharesToWithdraw[j]
                );
            }
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit SlashingWithdrawalQueued(params.withdrawalRoot, params.withdrawal, sharesToWithdraw);
        }
    }

    struct CompleteQueuedWithdrawalEmitStruct {
        Withdrawal withdrawal;
        IERC20[] tokens;
        bool receiveAsTokens;
    }

    function _completeQueuedWithdrawal_expectEmit(CompleteQueuedWithdrawalEmitStruct memory params) internal {
        if (!params.receiveAsTokens) {
            address operator = delegationManager.delegatedTo(params.withdrawal.staker);
            uint64[] memory slashingFactors = new uint64[](params.withdrawal.strategies.length);
            slashingFactors = allocationManagerMock.getMaxMagnitudes(operator, params.withdrawal.strategies);

            // receiving as shares so check for OperatorSharesIncrease and DepositScalingFactor updated
            for (uint i = 0; i < params.withdrawal.strategies.length; i++) {
                // Get updated deposit scaling factor
                uint curDepositShares;
                if (params.withdrawal.strategies[i] == beaconChainETHStrategy) {
                    curDepositShares = uint(eigenPodManagerMock.stakerDepositShares(params.withdrawal.staker, address(0)));
                    slashingFactors[i] =
                        uint64(slashingFactors[i].mulWad(eigenPodManagerMock.beaconChainSlashingFactor(params.withdrawal.staker)));
                } else {
                    curDepositShares = strategyManagerMock.stakerDepositShares(params.withdrawal.staker, params.withdrawal.strategies[i]);
                }

                uint sharesToWithdraw = _calcCompletedWithdrawnShares(params.withdrawal.scaledShares[i], slashingFactors[i]);

                uint expectedDepositScalingFactor = _calcDepositScalingFactor({
                    prevDsf: delegationManager.depositScalingFactor(params.withdrawal.staker, params.withdrawal.strategies[i]),
                    prevDepositShares: curDepositShares,
                    addedDepositShares: sharesToWithdraw,
                    slashingFactor: slashingFactors[i]
                });
                cheats.expectEmit(true, true, true, true, address(delegationManager));
                emit DepositScalingFactorUpdated(params.withdrawal.staker, params.withdrawal.strategies[i], expectedDepositScalingFactor);

                if (operator != address(0)) {
                    cheats.expectEmit(true, true, true, true, address(delegationManager));
                    emit OperatorSharesIncreased(operator, params.withdrawal.staker, params.withdrawal.strategies[i], sharesToWithdraw);
                }
            }
        }

        emit SlashingWithdrawalCompleted(delegationManager.calculateWithdrawalRoot(params.withdrawal));
    }

    struct CompleteQueuedWithdrawalsEmitStruct {
        Withdrawal[] withdrawals;
        IERC20[][] tokens;
        bool[] receiveAsTokens;
    }

    function _completeQueuedWithdrawals_expectEmit(CompleteQueuedWithdrawalsEmitStruct memory params) internal {
        for (uint i = 0; i < params.withdrawals.length; i++) {
            _completeQueuedWithdrawal_expectEmit(
                CompleteQueuedWithdrawalEmitStruct({
                    withdrawal: params.withdrawals[i],
                    tokens: params.tokens[i],
                    receiveAsTokens: params.receiveAsTokens[i]
                })
            );
        }
    }

    struct SlashOperatorSharesEmitStruct {
        address operator;
        IStrategy strategy;
        uint sharesToDecrease;
        uint sharesToBurn;
    }

    function _slashOperatorShares_expectEmit(SlashOperatorSharesEmitStruct memory params) internal {
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(params.operator, address(0), params.strategy, params.sharesToDecrease);
    }

    /// -----------------------------------------------------------------------
    /// Slashing Lib helpers
    /// Logic is essentially copied from SlashingLib to test the calculations
    /// and to avoid using the same library in the tests
    /// -----------------------------------------------------------------------

    /// @notice Calculates the exact withdrawable shares
    function _calcWithdrawableShares(uint depositShares, uint depositScalingFactor, uint slashingFactor) internal pure returns (uint) {
        return depositShares.mulWad(depositScalingFactor).mulWad(slashingFactor);
    }

    function _calcCompletedWithdrawnShares(uint scaledShares, uint slashingFactor) internal pure returns (uint) {
        return scaledShares.mulWad(slashingFactor);
    }

    /// @notice Calculates the new deposit scaling factor after a deposit
    function _calcDepositScalingFactor(uint prevDsf, uint prevDepositShares, uint addedDepositShares, uint slashingFactor)
        internal
        pure
        returns (uint)
    {
        if (prevDepositShares == 0) return uint(WAD).divWad(slashingFactor);

        uint currWithdrawableShares = _calcWithdrawableShares(prevDepositShares, prevDsf, slashingFactor);

        uint newWithdrawableShares = currWithdrawableShares + addedDepositShares;

        uint newDsf = newWithdrawableShares.divWad(prevDepositShares + addedDepositShares).divWad(slashingFactor);

        return newDsf;
    }

    function _calcSlashedAmount(uint operatorShares, uint64 prevMaxMagnitude, uint64 newMaxMagnitude)
        internal
        pure
        returns (uint slashedAmount, uint operatorSharesAfterSlash)
    {
        operatorSharesAfterSlash = operatorShares.mulDiv(newMaxMagnitude, prevMaxMagnitude, Math.Rounding.Up);
        slashedAmount = operatorShares - operatorSharesAfterSlash;
    }

    /// -----------------------------------------------------------------------
    /// Helper Assertions
    /// -----------------------------------------------------------------------

    /// @notice Asserts for depositShares, withdrawableShares, and depositScalingFactor after a deposit
    function _assertDeposit(
        address staker,
        address operator,
        IStrategy strategy,
        uint operatorSharesBefore,
        uint withdrawableSharesBefore,
        uint depositSharesBefore,
        uint prevDsf,
        uint depositAmount
    ) internal view {
        (uint[] memory withdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(staker, strategy.toArray());
        // Check deposit shares added correctly
        assertEq(
            depositShares[0], depositSharesBefore + depositAmount, "depositShares should be equal to depositSharesBefore + depositAmount"
        );
        // Check withdrawable shares are increased, with rounding error
        assertApproxEqRel(
            withdrawableShares[0],
            withdrawableSharesBefore + depositAmount,
            APPROX_REL_DIFF,
            "withdrawableShares should be equal to existingDepositShares - depositShares"
        );
        // Check the new dsf is accurate
        uint expectedWithdrawableShares;
        uint expectedDsf;
        {
            uint64 maxMagnitude = allocationManagerMock.getMaxMagnitude(operator, strategy);
            uint slashingFactor = _getSlashingFactor(staker, strategy, maxMagnitude);
            expectedDsf = _calcDepositScalingFactor(prevDsf, depositSharesBefore, depositAmount, slashingFactor);
            expectedWithdrawableShares = _calcWithdrawableShares(depositSharesBefore + depositAmount, expectedDsf, slashingFactor);
        }
        // Check the new dsf is accurate
        assertEq(
            expectedDsf, delegationManager.depositScalingFactor(staker, strategy), "depositScalingFactor should be equal to expectedDsf"
        );
        // Check new operatorShares increased correctly
        if (operator != address(0)) {
            assertEq(
                operatorSharesBefore + depositAmount,
                delegationManager.operatorShares(operator, strategy),
                "OperatorShares not increased correctly"
            );
        }
        // Check the newly calculated withdrawable shares are correct
        assertEq(withdrawableShares[0], expectedWithdrawableShares, "withdrawableShares should be equal to expectedWithdrawableShares");
    }

    /// @notice Asserts for depositShares, withdrawableShares, and depositScalingFactor after a delegation
    function _assertDelegation(
        address staker,
        address operator,
        IStrategy strategy,
        uint operatorSharesBefore,
        uint withdrawableSharesBefore,
        uint depositSharesBefore,
        uint prevDsf
    ) internal view {
        (uint[] memory withdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(staker, strategy.toArray());
        // Check deposit shares don't change
        assertEq(depositShares[0], depositSharesBefore, "depositShares should be equal to depositSharesBefore");
        // Check withdrawable shares don't change
        assertApproxEqRel(
            withdrawableShares[0],
            withdrawableSharesBefore,
            APPROX_REL_DIFF,
            "withdrawableShares should be equal to withdrawableSharesBefore"
        );
        // Check the new dsf is accurate
        uint expectedWithdrawableShares;
        uint expectedDsf;
        {
            uint64 maxMagnitude = allocationManagerMock.getMaxMagnitude(operator, strategy);
            expectedDsf = prevDsf.divWad(maxMagnitude);
            uint slashingFactor = _getSlashingFactor(staker, strategy, maxMagnitude);
            expectedWithdrawableShares = _calcWithdrawableShares(depositSharesBefore, expectedDsf, slashingFactor);
        }
        // Check the new dsf is accurate
        assertEq(
            expectedDsf, delegationManager.depositScalingFactor(staker, strategy), "depositScalingFactor should be equal to expectedDsf"
        );
        // Check new operatorShares increased correctly
        if (operator != address(0)) {
            assertEq(
                operatorSharesBefore + withdrawableSharesBefore,
                delegationManager.operatorShares(operator, strategy),
                "OperatorShares not increased correctly"
            );
        }
        // Check the newly calculated withdrawable shares are correct
        assertEq(withdrawableShares[0], expectedWithdrawableShares, "withdrawableShares should be equal to expectedWithdrawableShares");
    }

    /// @notice Asserts for depositShares, and operatorShares decremented properly after a withdrawal
    function _assertWithdrawal(
        address staker,
        address operator,
        IStrategy strategy,
        uint operatorSharesBefore,
        uint depositSharesBefore,
        uint depositSharesWithdrawn,
        uint depositScalingFactor,
        uint slashingFactor
    ) internal view {
        (uint[] memory withdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(staker, strategy.toArray());
        // Check deposit shares decreased correctly
        assertEq(
            depositShares[0],
            depositSharesBefore - depositSharesWithdrawn,
            "depositShares should be equal to depositSharesBefore - depositSharesWithdrawn"
        );
        // Check withdrawable shares are decreased, with rounding error
        uint expectedWithdrawableShares =
            _calcWithdrawableShares(depositSharesBefore - depositSharesWithdrawn, depositScalingFactor, slashingFactor);
        assertEq(withdrawableShares[0], expectedWithdrawableShares, "withdrawableShares should be equal to expectedWithdrawableShares");
        // Check operatorShares decreased properly
        uint expectedWithdrawnShares = _calcWithdrawableShares(depositSharesWithdrawn, depositScalingFactor, slashingFactor);
        assertEq(
            operatorSharesBefore - expectedWithdrawnShares,
            delegationManager.operatorShares(operator, strategy),
            "OperatorShares not decreased correctly"
        );
    }

    struct AssertCompletedWithdrawalStruct {
        address staker;
        address currentOperator;
        Withdrawal withdrawal;
        bool receiveAsTokens;
        uint[] operatorSharesBefore;
        uint[] withdrawableSharesBefore;
        uint[] depositSharesBefore;
        uint[] prevDepositScalingFactors;
        uint[] slashingFactors;
        uint64 beaconChainSlashingFactor;
    }

    /**
     * @notice Asserts for a queuedWithdrawal that its root is no longer pending and the withdrawal no longer exists
     * Also checks if the withdrawal is completed as shares that the current operator shares are increased appropriately
     * with the staker's depositScalingFactor updated.
     * NOTE: assumes no duplicate strategies in the withdrawal
     */
    function _assertCompletedWithdrawal(AssertCompletedWithdrawalStruct memory params) internal view {
        assertTrue(delegationManager.delegatedTo(params.staker) == params.currentOperator, "staker should be delegated to currentOperator");

        _assertWithdrawalRootsComplete(params.staker, params.withdrawal.toArray());
        // Check operator and staker shares if receiving as shares
        if (params.receiveAsTokens) {
            for (uint i = 0; i < params.withdrawal.strategies.length; i++) {
                {
                    // assert deposit and withdrawable shares unchanged
                    (uint[] memory withdrawableShares, uint[] memory depositShares) =
                        delegationManager.getWithdrawableShares(params.staker, params.withdrawal.strategies[i].toArray());
                    assertEq(
                        params.withdrawableSharesBefore[i],
                        withdrawableShares[0],
                        "withdrawableShares should be equal to withdrawableSharesBefore"
                    );
                    assertEq(params.depositSharesBefore[i], depositShares[0], "depositShares should be equal to depositSharesBefore");
                }
                // assert operatorShares unchanged
                assertEq(
                    params.operatorSharesBefore[i],
                    delegationManager.operatorShares(params.currentOperator, params.withdrawal.strategies[i]),
                    "OperatorShares should be equal to operatorSharesBefore"
                );
                // assert dsf is unchanged
                assertEq(
                    params.prevDepositScalingFactors[i],
                    delegationManager.depositScalingFactor(params.staker, params.withdrawal.strategies[i]),
                    "depositScalingFactor should be equal to prevDepositScalingFactors"
                );
            }
        } else {
            for (uint i = 0; i < params.withdrawal.strategies.length; i++) {
                // calculate shares to complete withdraw and add back as shares
                if (params.withdrawal.strategies[i] == beaconChainETHStrategy) {
                    params.slashingFactors[i] = uint64(params.slashingFactors[i].mulWad(params.beaconChainSlashingFactor));
                }
                uint sharesToAddBack = _calcCompletedWithdrawnShares(params.withdrawal.scaledShares[i], params.slashingFactors[i]);
                // assert deposit shares, withdrawable shares, and operator shares, and depositScalingFactor
                _assertDeposit({
                    staker: params.staker,
                    operator: params.currentOperator,
                    strategy: params.withdrawal.strategies[i],
                    operatorSharesBefore: params.operatorSharesBefore[i],
                    withdrawableSharesBefore: params.withdrawableSharesBefore[i],
                    depositSharesBefore: params.depositSharesBefore[i],
                    prevDsf: params.prevDepositScalingFactors[i],
                    depositAmount: sharesToAddBack
                });
            }
        }
    }

    /// @notice assert withdrawals completed are reflected as completed in storage for the withdrawal root and staker
    function _assertWithdrawalRootsComplete(address staker, Withdrawal[] memory withdrawals) internal view {
        for (uint i = 0; i < withdrawals.length; ++i) {
            // Check the withdrawal root is no longer pending
            // and also doesn't exist in storage for the staker
            bytes32 withdrawalRootToCheck = delegationManager.calculateWithdrawalRoot(withdrawals[i]);
            assertFalse(delegationManager.pendingWithdrawals(withdrawalRootToCheck), "withdrawalRoot not pending");
            (Withdrawal[] memory withdrawalsInStorage,) = delegationManager.getQueuedWithdrawals(staker);
            for (uint j = 0; j < withdrawalsInStorage.length; ++j) {
                assertTrue(
                    withdrawalRootToCheck != delegationManager.calculateWithdrawalRoot(withdrawalsInStorage[j]),
                    "withdrawal should not exist in storage"
                );
            }
        }
    }

    function _assertOperatorSharesAfterSlash(
        address operator,
        IStrategy strategy,
        uint operatorSharesBefore,
        uint64 prevMaxMagnitude,
        uint64 newMaxMagnitude
    ) internal view returns (uint sharesToDecrement, uint operatorSharesAfterSlash) {
        (sharesToDecrement, operatorSharesAfterSlash) =
            _calcSlashedAmount({operatorShares: operatorSharesBefore, prevMaxMagnitude: prevMaxMagnitude, newMaxMagnitude: newMaxMagnitude});

        assertEq(
            operatorSharesAfterSlash,
            delegationManager.operatorShares(operator, strategy),
            "OperatorShares should be equal to operatorSharesAfterSlash"
        );
        assertEq(
            delegationManager.operatorShares(operator, strategy) + sharesToDecrement,
            operatorSharesBefore,
            "OperatorShares + sharesToDecrement should be equal to operatorSharesBefore"
        );
    }

    function _assertSharesAfterSlash(
        address staker,
        IStrategy strategy,
        uint withdrawableSharesBefore,
        uint expectedWithdrawableShares,
        uint prevMaxMagnitude,
        uint currMaxMagnitude
    ) internal view {
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategy.toArray());

        assertApproxEqRel(
            uint(withdrawableSharesBefore).mulDiv(currMaxMagnitude, prevMaxMagnitude),
            withdrawableShares[0],
            APPROX_REL_DIFF,
            "withdrawableShares should be equal to withdrawableSharesBefore * currMaxMagnitude / prevMaxMagnitude"
        );

        assertEq(withdrawableShares[0], expectedWithdrawableShares, "withdrawableShares should be equal to expectedWithdrawableShares");
    }

    function _assertSharesAfterBeaconSlash(
        address staker,
        uint withdrawableSharesBefore,
        uint expectedWithdrawableShares,
        uint prevBeaconSlashingFactor
    ) internal view {
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, beaconChainETHStrategy.toArray());
        uint currBeaconSlashingFactor = eigenPodManagerMock.beaconChainSlashingFactor(defaultStaker);
        assertEq(withdrawableShares[0], expectedWithdrawableShares, "withdrawableShares should be equal to expectedWithdrawableShares");
        assertApproxEqRel(
            uint(withdrawableSharesBefore).mulDiv(currBeaconSlashingFactor, prevBeaconSlashingFactor),
            withdrawableShares[0],
            APPROX_REL_DIFF,
            "withdrawableShares should be equal to withdrawableSharesBefore * currBeaconSlashingFactor / prevBeaconChainSlashingFactor"
        );
    }

    /// @notice Due to rounding, withdrawable shares and operator shares may not align even if the operator
    /// only has the single staker with deposits.
    function _assertWithdrawableAndOperatorShares(uint withdrawableShares, uint operatorShares, string memory errorMessage) internal pure {
        if (withdrawableShares > 0) {
            assertApproxEqRel(withdrawableShares, operatorShares, APPROX_REL_DIFF, errorMessage);
        } else {}
        assertLe(withdrawableShares, operatorShares, "withdrawableShares should be less than or equal to operatorShares");
    }

    /**
     * @notice Assertion checks after queuing a withdrawal. Reads withdrawals set in storage in test
     * - Asserts exact match of Withdrawal struct exists in storage
     * - Asserts Withdrawal root is pending
     */
    function _assertQueuedWithdrawalExists(address staker) internal view {
        for (uint i = 0; i < stakerQueuedWithdrawals[staker].length; ++i) {
            Withdrawal memory withdrawal = stakerQueuedWithdrawals[staker][i];
            bytes32 withdrawalRootToCheck = delegationManager.calculateWithdrawalRoot(withdrawal);
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRootToCheck), "withdrawalRoot not pending");

            (Withdrawal[] memory withdrawals,) = delegationManager.getQueuedWithdrawals(staker);
            for (uint j = 0; j < withdrawals.length; ++j) {
                if (withdrawalRootToCheck == delegationManager.calculateWithdrawalRoot(withdrawals[j])) {
                    assertEq(withdrawals[j].staker, withdrawal.staker);
                    assertEq(withdrawals[j].withdrawer, withdrawal.withdrawer);
                    assertEq(withdrawals[j].delegatedTo, withdrawal.delegatedTo);
                    assertEq(withdrawals[j].nonce, withdrawal.nonce);
                    assertEq(withdrawals[j].startBlock, withdrawal.startBlock);
                    assertEq(withdrawals[j].scaledShares.length, withdrawal.scaledShares.length);
                    for (uint k = 0; k < withdrawal.scaledShares.length; ++k) {
                        assertEq(withdrawals[j].scaledShares[k], withdrawal.scaledShares[k]);
                        assertEq(address(withdrawals[j].strategies[k]), address(withdrawal.strategies[k]));
                    }
                }
            }
        }
    }

    /**
     * @notice Assertion checks after queuing a withdrawal. Reads withdrawals set in storage in test
     * - Asserts exact match of Withdrawal struct exists in storage
     * - Asserts Withdrawal root is pending
     */
    function _assertQueuedWithdrawalExists(address staker, Withdrawal memory withdrawal) internal view {
        bytes32 withdrawalRootToCheck = delegationManager.calculateWithdrawalRoot(withdrawal);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRootToCheck), "withdrawalRoot not pending");

        (Withdrawal[] memory withdrawals,) = delegationManager.getQueuedWithdrawals(staker);
        for (uint i = 0; i < withdrawals.length; ++i) {
            assertEq(withdrawals[i].staker, withdrawal.staker);
            assertEq(withdrawals[i].withdrawer, withdrawal.withdrawer);
            assertEq(withdrawals[i].delegatedTo, withdrawal.delegatedTo);
            assertEq(withdrawals[i].nonce, withdrawal.nonce);
            assertEq(withdrawals[i].startBlock, withdrawal.startBlock);
            assertEq(withdrawals[i].scaledShares.length, withdrawal.scaledShares.length);
            for (uint j = 0; j < withdrawal.scaledShares.length; ++j) {
                assertEq(withdrawals[i].scaledShares[j], withdrawal.scaledShares[j]);
                assertEq(address(withdrawals[i].strategies[j]), address(withdrawal.strategies[j]));
            }
        }
    }
}

contract DelegationManagerUnitTests_Initialization_Setters is DelegationManagerUnitTests {
    function test_initialization() public view {
        assertTrue(delegationManager.domainSeparator() != bytes32(0), "sanity check");
        assertEq(
            address(delegationManager.strategyManager()),
            address(strategyManagerMock),
            "constructor / initializer incorrect, strategyManager set wrong"
        );
        assertEq(
            address(delegationManager.pauserRegistry()),
            address(pauserRegistry),
            "constructor / initializer incorrect, pauserRegistry set wrong"
        );
        assertEq(
            address(delegationManager.eigenPodManager()),
            address(eigenPodManagerMock),
            "constructor / initializer incorrect, eigenPodManager set wrong"
        );
        assertEq(
            address(delegationManager.allocationManager()),
            address(allocationManagerMock),
            "constructor / initializer incorrect, allocationManager set wrong"
        );
        assertEq(
            delegationManager.minWithdrawalDelayBlocks(),
            MIN_WITHDRAWAL_DELAY_BLOCKS,
            "constructor / initializer incorrect, MIN_WITHDRAWAL_DELAY set wrong"
        );
        assertEq(delegationManager.owner(), address(this), "constructor / initializer incorrect, owner set wrong");
        assertEq(delegationManager.paused(), 0, "constructor / initializer incorrect, paused status set wrong");

        bytes memory v = bytes(delegationManager.version());
        bytes32 expectedDomainSeparator = keccak256(
            abi.encode(
                EIP712_DOMAIN_TYPEHASH,
                keccak256(bytes("EigenLayer")),
                keccak256(bytes.concat(v[0], v[1])),
                block.chainid,
                address(delegationManager)
            )
        );

        assertEq(delegationManager.domainSeparator(), expectedDomainSeparator, "sanity check");
    }

    /// @notice Verifies that the DelegationManager cannot be initialized multiple times
    function test_initialize_revert_reinitialization() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        delegationManager.initialize(address(this), 0);
    }
}

contract DelegationManagerUnitTests_RegisterModifyOperator is DelegationManagerUnitTests {
    using ArrayLib for *;

    function test_registerAsOperator_revert_paused() public {
        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.registerAsOperator(address(0), 0, emptyStringForMetadataURI);
    }

    // @notice Verifies that someone cannot successfully call `DelegationManager.registerAsOperator(delegationApprover)` again after registering for the first time
    function testFuzz_registerAsOperator_revert_cannotRegisterMultipleTimes(address delegationApprover) public {
        // Register once
        cheats.startPrank(defaultOperator);
        delegationManager.registerAsOperator(delegationApprover, 0, emptyStringForMetadataURI);

        // Expect revert when register again
        cheats.expectRevert(ActivelyDelegated.selector);
        delegationManager.registerAsOperator(delegationApprover, 0, emptyStringForMetadataURI);
        cheats.stopPrank();
    }

    /**
     * @notice `operator` registers via calling `DelegationManager.registerAsOperator(delegationApprover, metadataURI)`
     * Should be able to set any parameters, other than too high value for `stakerOptOutWindowBlocks`
     * The set parameters should match the desired parameters (correct storage update)
     * Operator becomes delegated to themselves
     * Properly emits events – especially the `OperatorRegistered` event, but also `StakerDelegated` & `DelegationApproverUpdated` events
     * Reverts appropriately if operator was already delegated to someone (including themselves, i.e. they were already an operator)
     * @param operator and @param delegationApprover are fuzzed inputs
     */
    function testFuzz_registerAsOperator(address operator, address delegationApprover, string memory metadataURI)
        public
        filterFuzzedAddressInputs(operator)
    {
        _registerOperator_expectEmit(
            RegisterAsOperatorEmitStruct({operator: operator, delegationApprover: delegationApprover, metadataURI: metadataURI})
        );
        cheats.prank(operator);
        delegationManager.registerAsOperator(delegationApprover, 0, metadataURI);

        // Storage checks
        assertEq(delegationApprover, delegationManager.delegationApprover(operator), "delegationApprover not set correctly");
        assertEq(delegationManager.delegatedTo(operator), operator, "operator not delegated to self");
    }

    /// @notice Register two separate operators shouldn't revert
    function testFuzz_registerAsOperator_TwoSeparateOperatorsRegister(
        address operator1,
        address delegationApprover1,
        address operator2,
        address delegationApprover2
    ) public {
        cheats.assume(operator1 != operator2);
        // register operator1 with expected emits
        _registerOperator_expectEmit(
            RegisterAsOperatorEmitStruct({
                operator: operator1,
                delegationApprover: delegationApprover1,
                metadataURI: emptyStringForMetadataURI
            })
        );
        _registerOperator(operator1, delegationApprover1, emptyStringForMetadataURI);
        // register operator2 with expected emits
        _registerOperator_expectEmit(
            RegisterAsOperatorEmitStruct({
                operator: operator2,
                delegationApprover: delegationApprover2,
                metadataURI: emptyStringForMetadataURI
            })
        );
        _registerOperator(operator2, delegationApprover2, emptyStringForMetadataURI);
        assertTrue(delegationManager.isOperator(operator1), "operator1 not registered");
        assertTrue(delegationManager.isOperator(operator2), "operator2 not registered");
    }

    // @notice Verifies that a staker who is actively delegated to an operator cannot register as an operator (without first undelegating, at least)
    function testFuzz_Revert_registerAsOperator_cannotRegisterWhileDelegated(address staker, address delegationApprover)
        public
        filterFuzzedAddressInputs(staker)
    {
        cheats.assume(staker != defaultOperator);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);

        // expect revert if attempt to register as operator
        cheats.expectRevert(ActivelyDelegated.selector);
        delegationManager.registerAsOperator(delegationApprover, 0, emptyStringForMetadataURI);

        cheats.stopPrank();
    }

    /// @notice Add test for registerAsOperator where the operator has existing deposits in strategies
    /// Assert:
    ///     depositShares == operatorShares == withdrawableShares
    ///     check operatorDetails hash encode matches the operatorDetails hash stored (call view function)
    function testFuzz_registerAsOperator_withDeposits(Randomness r) public rand(r) {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        // Set staker shares in StrategyManager
        IStrategy[] memory strategiesToReturn = strategyMock.toArray();
        uint[] memory sharesToReturn = new uint[](1);
        sharesToReturn[0] = shares;
        strategyManagerMock.setDeposits(defaultOperator, strategiesToReturn, sharesToReturn);

        // register operator, their own staker depositShares should increase their operatorShares
        _registerOperator_expectEmit(
            RegisterAsOperatorEmitStruct({operator: defaultOperator, delegationApprover: address(0), metadataURI: emptyStringForMetadataURI})
        );
        _registerOperatorWithBaseDetails(defaultOperator);
        uint operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);

        // check depositShares == operatorShares == withdrawableShares
        assertEq(operatorSharesAfter, shares, "operator shares not set correctly");
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultOperator, strategiesToReturn);
        assertEq(withdrawableShares[0], shares, "withdrawable shares not set correctly");
        assertEq(strategyManagerMock.stakerDepositShares(defaultOperator, strategyMock), shares, "staker deposit shares not set correctly");
    }

    /**
     * @notice Tests that an operator can modify their OperatorDetails by calling `DelegationManager.modifyOperatorDetails`
     * Should be able to set any parameters, other than setting their `earningsReceiver` to the zero address or too high value for `stakerOptOutWindowBlocks`
     * The set parameters should match the desired parameters (correct storage update)
     * Properly emits an `DelegationApproverUpdated` event
     * Reverts appropriately if the caller is not an operator
     * Reverts if operator tries to decrease their `stakerOptOutWindowBlocks` parameter
     * @param delegationApprover1 and @param delegationApprover2 are fuzzed inputs
     */
    function testFuzz_modifyOperatorParameters(address delegationApprover1, address delegationApprover2) public {
        _registerOperator_expectEmit(
            RegisterAsOperatorEmitStruct({
                operator: defaultOperator,
                delegationApprover: delegationApprover1,
                metadataURI: emptyStringForMetadataURI
            })
        );
        _registerOperator(defaultOperator, delegationApprover1, emptyStringForMetadataURI);

        cheats.startPrank(defaultOperator);

        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DelegationApproverUpdated(defaultOperator, delegationApprover2);
        delegationManager.modifyOperatorDetails(defaultOperator, delegationApprover2);

        assertEq(delegationApprover2, delegationManager.delegationApprover(defaultOperator), "delegationApprover not set correctly");
        assertEq(delegationManager.delegatedTo(defaultOperator), defaultOperator, "operator not delegated to self");
        // or else the transition is disallowed

        cheats.stopPrank();
    }

    // @notice Tests that an address which is not an operator cannot successfully call `updateOperatorMetadataURI`.
    function test_Revert_updateOperatorMetadataUri_notRegistered() public {
        assertFalse(delegationManager.isOperator(defaultOperator), "bad test setup");

        cheats.prank(defaultOperator);
        cheats.expectRevert(OperatorNotRegistered.selector);
        delegationManager.updateOperatorMetadataURI(defaultOperator, emptyStringForMetadataURI);
    }

    function test_Revert_updateOperatorMetadataUri_notOperator() public {
        cheats.expectRevert(OperatorNotRegistered.selector);
        delegationManager.modifyOperatorDetails(defaultOperator, defaultOperator);
    }

    /**
     * @notice Verifies that a staker cannot call cannot modify their `OperatorDetails` without first registering as an operator
     * @dev This is an important check to ensure that our definition of 'operator' remains consistent, in particular for preserving the
     * invariant that 'operators' are always delegated to themselves
     */
    function testFuzz_UpdateOperatorMetadataURI(string memory metadataURI) public {
        _registerOperatorWithBaseDetails(defaultOperator);

        // call `updateOperatorMetadataURI` and check for event
        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(defaultOperator, metadataURI);
        delegationManager.updateOperatorMetadataURI(defaultOperator, metadataURI);
    }

    function testFuzz_UAM_modifyOperatorDetails(address delegationApprover) public {
        // Set admin
        cheats.prank(defaultOperator);
        permissionController.setAppointee(
            defaultOperator, address(this), address(delegationManager), IDelegationManager.modifyOperatorDetails.selector
        );

        _registerOperatorWithBaseDetails(defaultOperator);

        // Modify operator details
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DelegationApproverUpdated(defaultOperator, delegationApprover);
        delegationManager.modifyOperatorDetails(defaultOperator, delegationApprover);

        // Check storage
        assertEq(delegationApprover, delegationManager.delegationApprover(defaultOperator), "delegationApprover not set correctly");
    }

    function testFuzz_UAM_updateOperatorMetadataURI(string memory metadataURI) public {
        // Set admin
        cheats.prank(defaultOperator);
        permissionController.setAppointee(
            defaultOperator, address(this), address(delegationManager), IDelegationManager.updateOperatorMetadataURI.selector
        );

        _registerOperatorWithBaseDetails(defaultOperator);

        // call `updateOperatorMetadataURI` and check for event
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(defaultOperator, metadataURI);
        delegationManager.updateOperatorMetadataURI(defaultOperator, metadataURI);
    }
}

contract DelegationManagerUnitTests_delegateTo is DelegationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;

    function test_Revert_WhenPaused() public {
        cheats.prank(defaultOperator);
        delegationManager.registerAsOperator(address(0), 0, emptyStringForMetadataURI);

        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;
        cheats.prank(defaultStaker);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
    }

    /**
     * @notice Delegates from `staker` to an operator, then verifies that the `staker` cannot delegate to another `operator` (at least without first undelegating)
     */
    function testFuzz_Revert_WhenDelegateWhileDelegated(
        Randomness r,
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry
    ) public rand(r) {
        address staker = r.Address();
        address operator = r.Address();
        bytes32 salt = r.Bytes32();

        // delegate from the staker to an operator
        _registerOperatorWithBaseDetails(operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // try to delegate again and check that the call reverts
        cheats.prank(staker);
        cheats.expectRevert(ActivelyDelegated.selector);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
    }

    /// @notice Verifies that `staker` cannot delegate to an unregistered `operator`
    function testFuzz_Revert_WhenDelegateToUnregisteredOperator(Randomness r) public rand(r) {
        address staker = r.Address();
        address operator = r.Address();
        assertFalse(delegationManager.isOperator(operator), "incorrect test input?");

        // try to delegate and check that the call reverts
        cheats.prank(staker);
        cheats.expectRevert(OperatorNotRegistered.selector);
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * Assertion checks
     * - Properly emitted events from `delegateTo`
     * - depositShares incremented for staker correctly
     * - withdrawableShares are correct
     * - depositScalingFactor is updated correctly
     * - operatorShares increase by depositShares amount
     * - defaultOperator is an operator, staker is delegated to defaultOperator, staker is not an operator
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_StrategyManagerShares(
        Randomness r,
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry
    ) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);

        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // Set staker shares in StrategyManager
        strategyManagerMock.addDeposit(staker, strategyMock, shares);
        // delegate from the `staker` to the operator
        cheats.prank(staker);
        _delegateTo_expectEmit_singleStrat(
            DelegateToSingleStratEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: strategyMock,
                depositShares: shares,
                depositScalingFactor: WAD
            })
        );
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`. `staker` holds beaconChainETHStrategy Shares
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * Assertion Checks
     * - Properly emitted events from `delegateTo`
     * - depositShares incremented for staker correctly
     * - withdrawableShares are correct
     * - depositScalingFactor is updated correctly
     * - operatorShares increase by depositShares amount
     * - defaultOperator is an operator, staker is delegated to defaultOperator, staker is not an operator
     * - That the staker withdrawableShares is <= operatorShares (less due to rounding from non-WAD maxMagnitude)
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_beaconChainStrategyShares(
        Randomness r,
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry
    ) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));

        _registerOperatorWithBaseDetails(defaultOperator);
        // Set the operators magnitude
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, WAD);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        _delegateTo_expectEmit_singleStrat(
            DelegateToSingleStratEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: beaconChainETHStrategy,
                depositShares: uint(beaconShares),
                depositScalingFactor: WAD
            })
        );
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: uint(beaconShares)
        });
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, beaconChainETHStrategy.toArray());
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "withdrawableShares not set correctly"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * but it should revert as the strategy has been fully slashed for the operator.
     * Assertion checks
     * - staker is not delegated to defaultOperator afterwards
     * - staker is not delegated
     * - staker is not registered as an operator
     * - salt is not spent
     */
    function testFuzz_Revert_OperatorWhoAcceptsAllStakers_AlreadySlashed100Percent_StrategyManagerShares(Randomness r) public rand(r) {
        address staker = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);

        _registerOperatorWithBaseDetails(defaultOperator);

        // Set staker shares in StrategyManager
        IStrategy[] memory strategiesToReturn = strategyMock.toArray();
        uint[] memory sharesToReturn = shares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);

        // Set the operators magnitude to be 0
        _setOperatorMagnitude(defaultOperator, strategyMock, 0);

        // delegate from the `staker` to the operator
        cheats.prank(staker);
        cheats.expectRevert(FullySlashed.selector);
        delegationManager.delegateTo(defaultOperator, emptyApproverSignatureAndExpiry, emptySalt);

        assertTrue(delegationManager.delegatedTo(staker) != defaultOperator, "staker should not be delegated to the operator");
        assertFalse(delegationManager.isDelegated(staker), "staker should not be delegated");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), emptySalt),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * but it should revert as the beaconChainStrategy has been fully slashed for the operator.
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * Assertion checks
     * - beaconChainETHStrategy shares are unchanged for the operator
     * - staker is not delegated to defaultOperator afterwards
     * - staker is not delegated
     * - staker is not registered as an operator
     * - salt is not spent
     */
    function testFuzz_Revert_OperatorWhoAcceptsAllStakers_AlreadySlashed100Percent_BeaconChainStrategyShares(
        Randomness r,
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry
    ) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));

        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        uint beaconSharesBefore = delegationManager.operatorShares(staker, beaconChainETHStrategy);

        // Set the operators magnitude for native restaking to be 0
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, 0);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(FullySlashed.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        uint beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);

        assertEq(beaconSharesBefore, beaconSharesAfter, "operator beaconchain shares should not have increased with negative shares");
        assertTrue(delegationManager.delegatedTo(staker) != defaultOperator, "staker should not be delegated to the operator");
        assertFalse(delegationManager.isDelegated(staker), "staker should not be delegated");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * and the strategy has already been slashed for the operator.
     * Assertion Checks
     * - Properly emitted events from `delegateTo`
     * - depositShares incremented for staker correctly
     * - withdrawableShares are correct
     * - depositScalingFactor is updated correctly
     * - operatorShares increase by depositShares amount
     * - defaultOperator is an operator, staker is delegated to defaultOperator, staker is not an operator
     * - That the staker withdrawableShares is <= operatorShares (less due to rounding from non-WAD maxMagnitude)
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_StrategyManagerShares(Randomness r) public rand(r) {
        address staker = r.Address();
        uint shares = r.Uint256(1 gwei, MAX_STRATEGY_SHARES);
        uint64 maxMagnitude = r.Uint64(1, WAD);

        _registerOperatorWithBaseDetails(defaultOperator);
        strategyManagerMock.addDeposit(staker, strategyMock, shares);
        _setOperatorMagnitude(defaultOperator, strategyMock, maxMagnitude);

        // Expected staker scaling factor
        uint stakerScalingFactor = uint(WAD).divWad(maxMagnitude);

        // delegate from the `staker` to the operator
        cheats.prank(staker);
        _delegateTo_expectEmit_singleStrat(
            DelegateToSingleStratEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: strategyMock,
                depositShares: shares,
                depositScalingFactor: stakerScalingFactor
            })
        );
        delegationManager.delegateTo(defaultOperator, emptyApproverSignatureAndExpiry, emptySalt);

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), emptySalt),
            "salt somehow spent too early?"
        );

        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategyMock.toArray());
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0], delegationManager.operatorShares(defaultOperator, strategyMock), "withdrawableShares not set correctly"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * and the strategy has already been slashed for the operator. `staker` holds beaconChainETHStrategy Shares
     * Assertion Checks
     * - Properly emitted events from `delegateTo`
     * - depositShares incremented for staker correctly
     * - withdrawableShares are correct
     * - depositScalingFactor is updated correctly
     * - operatorShares increase by depositShares amount
     * - defaultOperator is an operator, staker is delegated to defaultOperator, staker is not an operator
     * - That the staker withdrawableShares is <= operatorShares (less due to rounding from non-WAD maxMagnitude)
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_beaconChainStrategyShares(Randomness r) public rand(r) {
        address staker = r.Address();
        uint64 maxMagnitude = r.Uint64(1, WAD);
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));

        // Register and set operator's magnitude
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, maxMagnitude);

        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);

        // delegate from the `staker` to the operator, check for events emitted
        cheats.startPrank(staker);
        _delegateTo_expectEmit_singleStrat(
            DelegateToSingleStratEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: beaconChainETHStrategy,
                depositShares: beaconShares > 0 ? uint(beaconShares) : 0,
                depositScalingFactor: uint(WAD).divWad(maxMagnitude)
            })
        );
        delegationManager.delegateTo(defaultOperator, emptyApproverSignatureAndExpiry, emptySalt);

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: uint(beaconShares)
        });
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), emptySalt),
            "salt somehow spent too early?"
        );

        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, beaconChainETHStrategy.toArray());
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "withdrawableShares not set correctly"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * and the strategy has already been slashed for the operator. `staker` holds beaconChainETHStrategy Shares and has been
     * slashed on the beaconChain resulting in a non-WAD beaconChainSlashingFactor.
     * Assertion Checks
     * - Properly emitted events from `delegateTo`
     * - depositShares incremented for staker correctly
     * - withdrawableShares are correct
     * - depositScalingFactor is updated correctly
     * - operatorShares increase by withdrawableShares amount
     * - defaultOperator is an operator, staker is delegated to defaultOperator, staker is not an operator
     * - That the staker withdrawableShares is <= operatorShares (less due to rounding from non-WAD maxMagnitude)
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashedAVSAndBeaconChain_beaconChainStrategyShares(Randomness r) public rand(r) {
        address staker = r.Address();
        uint64 maxMagnitude = r.Uint64(1, WAD);
        uint64 beaconChainSlashingFactor = r.Uint64(1, WAD - 1);
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));

        // Register and set operator's magnitude
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, maxMagnitude);
        eigenPodManagerMock.setBeaconChainSlashingFactor(staker, beaconChainSlashingFactor);
        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, beaconChainETHStrategy.toArray());

        // delegate from the `staker` to the operator, check for events emitted
        cheats.startPrank(staker);
        _delegateTo_expectEmit_singleStrat(
            DelegateToSingleStratEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: beaconChainETHStrategy,
                depositShares: beaconShares > 0 ? withdrawableShares[0] : 0,
                depositScalingFactor: uint(WAD).divWad(maxMagnitude)
            })
        );
        delegationManager.delegateTo(defaultOperator, emptyApproverSignatureAndExpiry, emptySalt);

        _assertDelegation({
            staker: staker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: withdrawableShares[0],
            depositSharesBefore: uint(beaconShares),
            prevDsf: WAD
        });
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), emptySalt),
            "salt somehow spent too early?"
        );

        (uint[] memory withdrawableSharesAfter,) = delegationManager.getWithdrawableShares(staker, beaconChainETHStrategy.toArray());
        _assertWithdrawableAndOperatorShares(
            withdrawableSharesAfter[0],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "withdrawableShares not set correctly"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`
     * Similar to tests above but now with staker who has both EigenPod and StrategyManager shares.
     * Assertion Checks for strategyMock and beaconChainETHStrategy
     * - Properly emitted events from `delegateTo`
     * - depositShares incremented for staker correctly
     * - withdrawableShares are correct
     * - depositScalingFactor is updated correctly
     * - operatorShares increase by depositShares amount
     * - defaultOperator is an operator, staker is delegated to defaultOperator, staker is not an operator
     * - That the staker withdrawableShares is <= operatorShares (less due to rounding from non-WAD maxMagnitude)
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_BeaconChainAndStrategyManagerShares(
        Randomness r,
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry
    ) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);

        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // Set staker shares in BeaconChainStrategy and StrategyManager
        strategyManagerMock.addDeposit(staker, strategyMock, shares);
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        (IStrategy[] memory strategiesToReturn, uint[] memory sharesToReturn) = delegationManager.getDepositedShares(staker);
        uint[] memory depositScalingFactors = new uint[](2);
        depositScalingFactors[0] = uint(WAD);
        depositScalingFactors[1] = uint(WAD);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        _delegateTo_expectEmit(
            DelegateToEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategies: strategiesToReturn,
                depositShares: sharesToReturn,
                depositScalingFactors: depositScalingFactors
            })
        );
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: uint(beaconShares)
        });
        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategiesToReturn);
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0], delegationManager.operatorShares(defaultOperator, strategyMock), "withdrawableShares not set correctly"
        );
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[1],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "withdrawableShares not set correctly"
        );

        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice `defaultStaker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `defaultStaker` calling `DelegationManager.delegateTo`
     * Similar to tests above but now with staker who has both EigenPod and StrategyManager shares.
     * The operator has been slashed prior to deposit for both strategies.
     * Assertion Checks for strategyMock and beaconChainETHStrategy
     * - Properly emitted events from `delegateTo`
     * - depositShares incremented for staker correctly
     * - withdrawableShares are correct
     * - depositScalingFactor is updated correctly
     * - operatorShares increase by depositShares amount
     * - defaultOperator is an operator, defaultStaker is delegated to defaultOperator, defaultStaker is not an operator
     * - That the defaultStaker withdrawableShares is <= operatorShares (less due to rounding from non-WAD maxMagnitude)
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_BeaconChainAndStrategyManagerShares(Randomness r) public rand(r) {
        // 1. register operator and setup values, magnitudes
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));
        uint64 maxMagnitudeBeacon = r.Uint64(1, WAD);
        uint64 maxMagnitudeStrategy = r.Uint64(1, WAD);
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, maxMagnitudeBeacon);
        _setOperatorMagnitude(defaultOperator, strategyMock, maxMagnitudeStrategy);

        // 2. Set staker shares in BeaconChainStrategy and StrategyManager
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, shares);
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, beaconShares);
        (IStrategy[] memory strategiesToReturn, uint[] memory sharesToReturn) = delegationManager.getDepositedShares(defaultStaker);

        // 3. delegate from the `staker` to the operator with expected emitted events
        cheats.startPrank(defaultStaker);
        uint[] memory depositScalingFactors = new uint[](2);
        depositScalingFactors[0] = uint(WAD).divWad(maxMagnitudeStrategy);
        depositScalingFactors[1] = uint(WAD).divWad(maxMagnitudeBeacon);
        _delegateTo_expectEmit(
            DelegateToEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategies: strategiesToReturn,
                depositShares: sharesToReturn,
                depositScalingFactors: depositScalingFactors
            })
        );
        delegationManager.delegateTo(defaultOperator, emptyApproverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();

        // 4. Assert correct end state values
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: uint(beaconShares)
        });
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });
        assertTrue(delegationManager.isOperator(defaultOperator), "defaultStaker not registered as operator");
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "defaultStaker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(defaultStaker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), emptySalt),
            "salt somehow spent too early?"
        );
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategiesToReturn);
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0],
            delegationManager.operatorShares(defaultOperator, strategyMock),
            "withdrawable strategy shares not set correctly"
        );
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[1],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "withdrawable beacon shares not set correctly"
        );
    }

    /**
     * @notice `staker` delegates to a operator who does not require any signature verification similar to test above.
     * In this scenario, staker doesn't have any delegatable shares and operator shares should not increase. Staker
     * should still be correctly delegated to the operator after the call.
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_ZeroDelegatableShares(
        Randomness r,
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry
    ) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but using an invalid expiry on purpose and checking that reversion occurs
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresECDSASignature_ExpiredDelegationApproverSignature(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(0, block.timestamp - 1);
        // roll to a very late timestamp
        skip(type(uint).max / 2);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(ISignatureUtilsMixinErrors.SignatureExpired.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but undelegating after delegating and trying the same approveSignature
     * and checking that reversion occurs with the same salt
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresECDSASignature_PreviouslyUsedSalt(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(block.timestamp + 1, type(uint).max);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // delegate from the `staker` to the operator, undelegate, and then try to delegate again with same approversalt
        // to check that call reverts
        cheats.startPrank(staker);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        assertTrue(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent not spent?"
        );
        delegationManager.undelegate(staker);
        cheats.expectRevert(SaltSpent.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but using an incorrect signature on purpose and checking that reversion occurs
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresECDSASignature_WithBadSignature(Randomness random) public rand(random) {
        address staker = random.Address();
        uint expiry = random.Uint256(block.timestamp + 1, type(uint).max);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // calculate the signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateDelegationApprovalDigestHash(
                staker, defaultOperator, delegationManager.delegationApprover(defaultOperator), emptySalt, expiry
            );
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(delegationSignerPrivateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        cheats.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresECDSASignature(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(block.timestamp, type(uint).max);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        if (staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Operator shares should increase by the amount of shares delegated
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresECDSASignature_StrategyManagerShares(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(block.timestamp, type(uint).max);
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // Set staker shares in StrategyManager
        strategyManagerMock.addDeposit(staker, strategyMock, shares);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        _delegateTo_expectEmit_singleStrat(
            DelegateToSingleStratEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: strategyMock,
                depositShares: shares,
                depositScalingFactor: WAD
            })
        );
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategyMock.toArray());
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0], delegationManager.operatorShares(defaultOperator, strategyMock), "withdrawableShares not set correctly"
        );

        if (staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Operator beaconShares should increase by the amount of shares delegated if beaconShares > 0
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresECDSASignature_BeaconChainStrategyShares(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(block.timestamp, type(uint).max);
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        _delegateTo_expectEmit_singleStrat(
            DelegateToSingleStratEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: beaconChainETHStrategy,
                depositShares: uint(beaconShares),
                depositScalingFactor: WAD
            })
        );
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: uint(beaconShares)
        });
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, beaconChainETHStrategy.toArray());
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "withdrawableShares not set correctly"
        );
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        if (staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Operator beaconshares should increase by the amount of beaconShares delegated if beaconShares > 0
     * Operator strategy manager shares should icnrease by amount of shares
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresECDSASignature_BeaconChainAndStrategyManagerShares(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(block.timestamp, type(uint).max);
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);

        // filter inputs, since this will fail when the staker is already registered as an operator
        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // Set staker shares in BeaconChainStrategy and StrategyManager
        uint[] memory depositScalingFactors = new uint[](2);
        depositScalingFactors[0] = uint(WAD);
        depositScalingFactors[1] = uint(WAD);
        strategyManagerMock.addDeposit(staker, strategyMock, shares);
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        (IStrategy[] memory strategiesToReturn, uint[] memory sharesToReturn) = delegationManager.getDepositedShares(staker);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        _delegateTo_expectEmit(
            DelegateToEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategies: strategiesToReturn,
                depositShares: sharesToReturn,
                depositScalingFactors: depositScalingFactors
            })
        );
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: uint(beaconShares)
        });
        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategiesToReturn);
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0],
            delegationManager.operatorShares(defaultOperator, strategyMock),
            "withdrawableShares for strategy not set correctly"
        );
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[1],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "withdrawableShares for beacon strategy not set correctly"
        );
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        if (staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice delegateTo test with operator's delegationApprover address set to a contract address
     * and check that reversion occurs when the signature is expired
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresEIP1271Signature_ExpiredDelegationApproverSignature(Randomness r) public rand(r) {
        address staker = r.Address();
        uint expiry = r.Uint256(0, block.timestamp - 1);
        uint currTimestamp = r.Uint256(block.timestamp, type(uint).max);

        // roll to a late timestamp
        skip(currTimestamp);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // create the signature struct
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        cheats.expectRevert(ISignatureUtilsMixinErrors.SignatureExpired.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
    }

    /**
     * @notice delegateTo test with operator's delegationApprover address set to a contract address
     * and check that reversion occurs when the signature approverSalt is already used.
     * Performed by delegating to operator, undelegating, and trying to reuse the same signature
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresEIP1271Signature_PreviouslyUsedSalt(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(block.timestamp, type(uint).max);

        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        _registerOperatorWith1271DelegationApprover(defaultOperator);

        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        delegationManager.undelegate(staker);
        // Reusing same signature should revert with salt already being used
        cheats.expectRevert(SaltSpent.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    /**
     * @notice delegateTo test with operator's delegationApprover address set to a contract address that
     * is non compliant with EIP1271
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresEIP1271Signature_NonCompliantWallet(Randomness r) public rand(r) {
        address staker = r.Address();
        uint expiry = r.Uint256(block.timestamp, type(uint).max);

        // deploy a ERC1271MaliciousMock contract that will return an incorrect value when called
        ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();
        _registerOperator(defaultOperator, address(wallet), emptyStringForMetadataURI);

        // create the signature struct
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        // because the ERC1271MaliciousMock contract returns the wrong amount of data, we get a low-level "EvmError: Revert" message here rather than the error message bubbling up
        cheats.expectRevert();
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
    }

    /**
     * @notice delegateTo test with operator's delegationApprover address set to a contract address that
     * returns a value other than the EIP1271 "magic bytes" and checking that reversion occurs appropriately
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresEIP1271Signature_IsValidSignatureFails(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(block.timestamp, type(uint).max);

        // deploy a ERC1271WalletMock contract that will return an incorrect value when called
        // owner is the 0 address
        ERC1271WalletMock wallet = new ERC1271WalletMock(address(1));
        _registerOperator(defaultOperator, address(wallet), emptyStringForMetadataURI);

        // calculate the delegationSigner's but this is not the correct signature from the wallet contract
        // since the wallet owner is address(1)
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        // Signature should fail as the wallet will not return EIP1271_MAGICVALUE
        cheats.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EIP1271-compliant contract (i.e. the operator’s `delegationApprover` address is
     * set to a nonzero and code-containing address) via the `staker` calling `DelegationManager.delegateTo`
     * The function uses OZ's ERC1271WalletMock contract, and thus should pass *only when a valid ECDSA signature from the `owner` of the ERC1271WalletMock contract,
     * OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresEIP1271Signature(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        uint expiry = r.Uint256(block.timestamp, type(uint).max);

        _registerOperatorWith1271DelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        assertTrue(delegationManager.isDelegated(staker), "staker not delegated correctly");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the nonce incremented appropriately
        if (staker == defaultOperator || staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt),
                "salt somehow spent not spent?"
            );
        }
    }
}

contract DelegationManagerUnitTests_increaseDelegatedShares is DelegationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;
    using Math for *;

    /// @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts if not called by the StrategyManager nor EigenPodManager
    function testFuzz_Revert_increaseDelegatedShares_invalidCaller(Randomness r) public rand(r) {
        address invalidCaller = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        cheats.expectRevert(IDelegationManagerErrors.OnlyStrategyManagerOrEigenPodManager.selector);
        delegationManager.increaseDelegatedShares(invalidCaller, strategyMock, 0, shares);
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts when operator slashed 100% for a strategy
     * and the staker has deposits in that strategy
     */
    function testFuzz_Revert_increaseDelegatedShares_slashedOperator100Percent(Randomness r) public rand(r) {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        address staker = r.Address();

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        // Set operator magnitude
        _setOperatorMagnitude({operator: defaultOperator, strategy: strategyMock, magnitude: 0});
        // delegate from the `staker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        uint _delegatedSharesBefore = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategyMock);

        cheats.prank(address(strategyManagerMock));
        cheats.expectRevert(FullySlashed.selector);
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, shares);

        uint delegatedSharesAfter = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategyMock);

        assertEq(delegatedSharesAfter, _delegatedSharesBefore, "delegated shares incremented incorrectly");
        assertEq(_delegatedSharesBefore, 0, "nonzero shares delegated to zero address!");
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts when operator slashed 100% for a strategy
     * and the staker has deposits in that strategy. In this test case, the staker was initially deposited and delegated
     * to the operator before the operator was slashed 100%.
     * @dev Checks that withdrawable shares after 100% slashed is 0
     * @dev Checks that as a staker, redepositing after 100% slashed reverts
     */
    function testFuzz_Revert_increaseDelegatedShares_slashedOperator100PercentWithExistingStaker(Randomness r) public rand(r) {
        address staker = r.Address();
        uint64 initialMagnitude = r.Uint64(1, WAD);
        uint existingShares = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint shares = r.Uint256(existingShares, MAX_STRATEGY_SHARES);

        // 1. Register operator with initial operator magnitude and delegate staker to operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude({operator: defaultOperator, strategy: strategyMock, magnitude: initialMagnitude});
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        // 2. set staker initial shares and increase delegated shares
        IStrategy[] memory strategiesDeposited = strategyMock.toArray();
        uint[] memory sharesToReturn = existingShares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategiesDeposited, sharesToReturn);

        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, existingShares);

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: existingShares
        });
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategiesDeposited);
        _assertWithdrawableAndOperatorShares(
            withdrawableShares[0], delegationManager.operatorShares(defaultOperator, strategyMock), "Shares not increased correctly"
        );
        // 3. Now set operator magnitude to 0 (100% slashed)
        _setOperatorMagnitude({operator: defaultOperator, strategy: strategyMock, magnitude: 0});

        // 4. Try to "redeposit" and expect a revert since strategy is 100% slashed
        // staker's withdrawable shares should also be 0 now
        cheats.prank(address(strategyManagerMock));
        cheats.expectRevert(FullySlashed.selector);
        delegationManager.increaseDelegatedShares(staker, strategyMock, existingShares, shares);

        (withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategiesDeposited);
        assertEq(withdrawableShares[0], 0, "All existing shares should be slashed");
    }

    /// @notice Verifies that there is no change in operatorShares if the staker is not delegated
    function testFuzz_increaseDelegatedShares_noop(Randomness r) public rand(r) {
        address staker = r.Address();
        _registerOperatorWithBaseDetails(defaultOperator);
        assertFalse(delegationManager.isDelegated(staker), "bad test setup");

        cheats.prank(address(strategyManagerMock));
        vm.recordLogs();
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, 0);
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "shares should not have changed");

        Vm.Log[] memory entries = vm.getRecordedLogs();
        assertEq(0, entries.length, "should not have emitted any events");
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` properly increases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategy
     * Asserts:
     * - depositScalingFactor, depositShares, withdrawableShares, operatorShares after deposit
     * - correct operator shares after deposit
     *
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testFuzz_increaseDelegatedShares(Randomness r) public rand(r) {
        address staker = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        bool delegateFromStakerToOperator = r.Boolean();

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint delegatedSharesBefore = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategyMock);

        // deposit and increase delegated shares
        strategyManagerMock.addDeposit(staker, strategyMock, shares);
        if (delegationManager.isDelegated(staker)) {
            _increaseDelegatedShares_expectEmit(
                IncreaseDelegatedSharesEmitStruct({
                    staker: staker,
                    operator: delegationManager.delegatedTo(staker),
                    strategy: strategyMock,
                    sharesToIncrease: shares,
                    depositScalingFactor: WAD
                })
            );
        }
        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, shares);
        _assertDeposit({
            staker: staker,
            operator: delegationManager.delegatedTo(staker),
            strategy: strategyMock,
            operatorSharesBefore: delegatedSharesBefore,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });

        // Assert correct end state values
        uint delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategyMock.toArray());
        if (delegationManager.isDelegated(staker)) {
            _assertWithdrawableAndOperatorShares(withdrawableShares[0], delegatedSharesAfter, "Invalid withdrawable shares");
        } else {
            assertEq(delegatedSharesAfter, delegatedSharesBefore, "delegated shares incremented incorrectly");
            assertEq(delegatedSharesBefore, 0, "nonzero shares delegated to zero address!");
        }
    }

    function testFuzz_increaseDelegatedShares_beaconChainShares(Randomness r) public rand(r) {
        address staker = r.Address();
        uint shares = r.Uint256(1, MAX_ETH_SUPPLY);
        uint64 beaconChainSlashingFactor = r.Uint64(1, WAD);

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint delegatedSharesBefore = delegationManager.operatorShares(delegationManager.delegatedTo(staker), beaconChainETHStrategy);

        // deposit and increase delegated shares
        eigenPodManagerMock.setPodOwnerShares(staker, int(shares));
        eigenPodManagerMock.setBeaconChainSlashingFactor(staker, beaconChainSlashingFactor);
        _increaseDelegatedShares_expectEmit(
            IncreaseDelegatedSharesEmitStruct({
                staker: staker,
                operator: delegationManager.delegatedTo(staker),
                strategy: beaconChainETHStrategy,
                sharesToIncrease: shares,
                depositScalingFactor: uint(WAD).divWad(beaconChainSlashingFactor)
            })
        );
        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, beaconChainETHStrategy, 0, shares);
        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: delegatedSharesBefore,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });

        // Assert correct end state values
        uint delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, beaconChainETHStrategy.toArray());
        _assertWithdrawableAndOperatorShares(withdrawableShares[0], delegatedSharesAfter, "Invalid withdrawable shares");
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` properly increases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategy
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testFuzz_increaseDelegatedShares_slashedOperator(Randomness r) public rand(r) {
        address staker = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint64 magnitude = r.Uint64(1, WAD);
        bool delegateFromStakerToOperator = r.Boolean();

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Set operator magnitude
        _setOperatorMagnitude(defaultOperator, strategyMock, magnitude);

        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint delegatedSharesBefore = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategyMock);

        strategyManagerMock.addDeposit(staker, strategyMock, shares);
        if (delegationManager.isDelegated(staker)) {
            _increaseDelegatedShares_expectEmit(
                IncreaseDelegatedSharesEmitStruct({
                    staker: staker,
                    operator: defaultOperator,
                    strategy: strategyMock,
                    sharesToIncrease: shares,
                    depositScalingFactor: uint(WAD).divWad(magnitude)
                })
            );
        }
        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, shares);

        _assertDeposit({
            staker: staker,
            operator: delegationManager.delegatedTo(staker),
            strategy: strategyMock,
            operatorSharesBefore: delegatedSharesBefore,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });

        // Assert correct values
        uint delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategyMock.toArray());

        if (delegationManager.isDelegated(staker)) {
            _assertWithdrawableAndOperatorShares(withdrawableShares[0], delegatedSharesAfter, "Invalid withdrawable shares");
        } else {
            assertEq(delegatedSharesAfter, delegatedSharesBefore, "delegated shares incremented incorrectly");
            assertEq(delegatedSharesBefore, 0, "nonzero shares delegated to zero address!");
        }
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` properly increases the delegated `shares` for the
     * `defaultOperator` who the staker is delegated to. Asserts for proper events emitted and correct withdrawable shares,
     * despoitScalingFactor for the staker, and operator shares after deposit.
     */
    function testFuzz_increaseDelegatedShares_slashedOperatorAndBeaconChainShares(Randomness r) public rand(r) {
        address staker = r.Address();
        uint shares = r.Uint256(1, MAX_ETH_SUPPLY);
        uint64 maxMagnitude = r.Uint64(1, WAD);
        uint64 beaconChainSlashingFactor = r.Uint64(1, WAD);

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        // Set operator magnitude
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, maxMagnitude);
        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint delegatedSharesBefore = delegationManager.operatorShares(delegationManager.delegatedTo(staker), beaconChainETHStrategy);

        // deposit and increase delegated shares
        eigenPodManagerMock.setPodOwnerShares(staker, int(shares));
        eigenPodManagerMock.setBeaconChainSlashingFactor(staker, beaconChainSlashingFactor);
        _increaseDelegatedShares_expectEmit(
            IncreaseDelegatedSharesEmitStruct({
                staker: staker,
                operator: delegationManager.delegatedTo(staker),
                strategy: beaconChainETHStrategy,
                sharesToIncrease: shares,
                depositScalingFactor: uint(WAD).divWad(maxMagnitude.mulWad(beaconChainSlashingFactor))
            })
        );
        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, beaconChainETHStrategy, 0, shares);
        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: delegatedSharesBefore,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });

        // Assert correct end state values
        uint delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, beaconChainETHStrategy.toArray());
        _assertWithdrawableAndOperatorShares(withdrawableShares[0], delegatedSharesAfter, "Invalid withdrawable shares");
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` doesn't revert when operator slashed 100% for a strategy
     * and the staker has deposits in a separate strategy
     */
    function testFuzz_increaseDelegatedShares_slashedOperator100Percent(Randomness r) public rand(r) {
        address staker = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint64 magnitude = r.Uint64(1, WAD);
        IStrategy[] memory strategyArray = r.StrategyArray(1);
        IStrategy strategy = strategyArray[0];

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        // Set operator magnitude for 100% slashed strategy
        _setOperatorMagnitude({operator: defaultOperator, strategy: strategyMock, magnitude: 0});
        // Set operator magnitude for non-100% slashed strategy
        _setOperatorMagnitude({operator: defaultOperator, strategy: strategy, magnitude: magnitude});
        // delegate from the `staker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        uint delegatedSharesBefore = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategy);

        // deposit and increaseDelegatedShares
        strategyManagerMock.addDeposit(staker, strategy, shares);
        uint slashingFactor = _getSlashingFactor(staker, strategy, magnitude);
        dsf.update(0, shares, slashingFactor);
        _increaseDelegatedShares_expectEmit(
            IncreaseDelegatedSharesEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: strategy,
                sharesToIncrease: shares,
                depositScalingFactor: dsf.scalingFactor()
            })
        );
        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategy, 0, shares);

        _assertDeposit({
            staker: staker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: delegatedSharesBefore,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });

        // Assert correct end state values
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(staker, strategyArray);
        uint delegatedSharesAfter = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategy);
        _assertWithdrawableAndOperatorShares(withdrawableShares[0], delegatedSharesAfter, "Invalid withdrawable shares");
    }

    /**
     * @notice A unique test setup where impact of rounding can clearly be observed here.
     * After making the initial deposit of 44182209037560531097078597505 shares, and the operator's magnitude is set to 999999999999990009,
     * Each subsequent deposit amount of 1000 actually results in LESS withdrawable shares for the staker. There in an increasing drift
     * between the operator's shares and the staker's withdrawable shares.
     * The test below results in a drift difference of 4.418e13
     */
    function test_increaseDelegatedShares_depositRepeatedly() public {
        uint64 initialMagnitude = 999_999_999_999_990_009;
        uint shares = 44_182_209_037_560_531_097_078_597_505;

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initialMagnitude);

        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, shares);

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        {
            for (uint i = 0; i < 1000; ++i) {
                cheats.prank(address(strategyManagerMock));
                delegationManager.increaseDelegatedShares(defaultStaker, strategyMock, shares, 1000);
                shares += 1000;
                uint[] memory newDepositSharesArray = new uint[](1);
                newDepositSharesArray[0] = shares;

                strategyManagerMock.setDeposits(defaultStaker, strategies, newDepositSharesArray);
            }
        }

        (uint[] memory withdrawableShares, uint[] memory depositShares) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertEq(depositShares[0], shares, "staker deposit shares not reset correctly");
        assertEq(
            delegationManager.operatorShares(defaultOperator, strategyMock) - withdrawableShares[0],
            44_182_209_037_566,
            "drift should be 4.418e13 from previous tests"
        );
    }
}

contract DelegationManagerUnitTests_decreaseDelegatedShares is DelegationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;
    using Math for *;

    function testFuzz_Revert_decreaseDelegatedShares_invalidCaller(Randomness r) public rand(r) {
        address invalidCaller = r.Address();
        address staker = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint64 beaconChainSlashingFactorDecrease = uint64(r.Uint256(0, WAD));
        cheats.expectRevert(IDelegationManagerErrors.OnlyEigenPodManager.selector);
        cheats.prank(invalidCaller);
        delegationManager.decreaseDelegatedShares(staker, shares, beaconChainSlashingFactorDecrease);
    }

    /// @notice Verifies that there is no change in operatorShares if the staker is not delegated
    function testFuzz_decreaseDelegatedShares_noop(Randomness r) public rand(r) {
        address staker = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint64 beaconChainSlashingFactorDecrease = uint64(r.Uint256(0, WAD));

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        assertFalse(delegationManager.isDelegated(staker), "bad test setup");

        cheats.prank(address(eigenPodManagerMock));
        delegationManager.decreaseDelegatedShares(staker, shares, beaconChainSlashingFactorDecrease);
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "shares should not have changed");
    }

    /**
     * @notice Verifies that `decreaseDelegatedShares` properly updates the staker's withdrawable shares
     * and their delegated operator's shares are decreased by the correct amount.
     * Ensures that after the decrease, the staker's withdrawableShares <= operatorShares,
     * preventing any underflow for the operator's shares if they were all to be withdrawn.
     */
    function testFuzz_decreaseDelegatedShares_nonSlashedOperator(Randomness r) public rand(r) {
        int beaconShares = int(r.Uint256(1, MAX_ETH_SUPPLY));
        uint sharesDecrease = r.Uint256(0, uint(beaconShares) - 1);
        uint64 beaconChainSlashingFactor = r.Uint64(1, WAD);

        // 1. Setup staker and delegate to operator
        _registerOperatorWithBaseDetails(defaultOperator);
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, beaconShares);
        eigenPodManagerMock.setBeaconChainSlashingFactor(defaultStaker, beaconChainSlashingFactor);
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDelegation({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: withdrawableShares[0],
            depositSharesBefore: uint(beaconShares),
            prevDsf: WAD
        });

        // 2. Perform beaconChain slash + decreaseDelegatedShares()
        (uint64 prevBeaconSlashingFactor, uint64 newBeaconSlashingFactor) =
            _setNewBeaconChainSlashingFactor(defaultStaker, beaconShares, sharesDecrease);
        uint64 beaconChainSlashingFactorDecrease = prevBeaconSlashingFactor - newBeaconSlashingFactor;
        assertEq(beaconChainSlashingFactor, prevBeaconSlashingFactor, "Bad test setup");
        uint depositScalingFactor = uint(WAD);
        // expected operatorShares decreased for event
        uint operatorSharesToDecrease = _calcWithdrawableShares({
            depositShares: uint(beaconShares),
            depositScalingFactor: depositScalingFactor,
            slashingFactor: beaconChainSlashingFactorDecrease
        });
        // expected events
        _decreaseDelegatedShares_expectEmit(
            DecreaseDelegatedSharesEmitStruct({staker: defaultStaker, operator: defaultOperator, sharesToDecrease: operatorSharesToDecrease})
        );
        cheats.prank(address(eigenPodManagerMock));
        delegationManager.decreaseDelegatedShares(defaultStaker, uint(beaconShares), beaconChainSlashingFactorDecrease);

        // 3. Assert correct values
        uint expectedWithdrawableShares = _calcWithdrawableShares({
            depositShares: uint(beaconShares),
            depositScalingFactor: depositScalingFactor,
            slashingFactor: newBeaconSlashingFactor
        });
        _assertSharesAfterBeaconSlash({
            staker: defaultStaker,
            withdrawableSharesBefore: withdrawableShares[0],
            expectedWithdrawableShares: expectedWithdrawableShares,
            prevBeaconSlashingFactor: prevBeaconSlashingFactor
        });
        // Assert correct end state values
        (uint[] memory withdrawableSharesAfter,) = delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());

        assertEq(
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy) + operatorSharesToDecrease,
            withdrawableShares[0],
            "operator shares not decreased correctly"
        );

        _assertWithdrawableAndOperatorShares(
            withdrawableSharesAfter[0],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "Invalid withdrawable shares"
        );
    }

    /**
     * @notice Similar test to `testFuzz_decreaseDelegatedShares_nonSlashedOperator` but with
     * a pre-slashed operator (maxMagnitude < WAD).
     * Verifies that `decreaseDelegatedShares` properly updates the staker's withdrawable shares
     * and their delegated operator's shares are decreased by the correct amount.
     * Ensures that after the decrease, the staker's withdrawableShares <= operatorShares,
     * preventing any underflow for the operator's shares if they were all to be withdrawn.
     */
    function testFuzz_decreaseDelegatedShares_slashedOperator(Randomness r) public rand(r) {
        int beaconShares = int(r.Uint256(1, MAX_ETH_SUPPLY));
        uint sharesDecrease = r.Uint256(0, uint(beaconShares) - 1);
        uint64 maxMagnitude = r.Uint64(1, WAD - 1);
        uint64 beaconChainSlashingFactor = r.Uint64(1, WAD);

        // 1. Setup staker and delegate to operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, maxMagnitude);
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, beaconShares);
        eigenPodManagerMock.setBeaconChainSlashingFactor(defaultStaker, beaconChainSlashingFactor);
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDelegation({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: withdrawableShares[0],
            depositSharesBefore: uint(beaconShares),
            prevDsf: WAD
        });

        // 2. Perform beaconChain slash + decreaseDelegatedShares()
        (uint64 prevBeaconSlashingFactor, uint64 newBeaconSlashingFactor) =
            _setNewBeaconChainSlashingFactor(defaultStaker, beaconShares, sharesDecrease);
        uint64 beaconChainSlashingFactorDecrease = prevBeaconSlashingFactor - newBeaconSlashingFactor;
        assertEq(beaconChainSlashingFactor, prevBeaconSlashingFactor, "Bad test setup");
        uint depositScalingFactor = uint(WAD).divWad(maxMagnitude);
        // expected operatorShares decreased for event
        uint operatorSharesToDecrease = _calcWithdrawableShares({
            depositShares: uint(beaconShares),
            depositScalingFactor: depositScalingFactor,
            slashingFactor: maxMagnitude.mulWad(beaconChainSlashingFactorDecrease)
        });
        // expected events
        _decreaseDelegatedShares_expectEmit(
            DecreaseDelegatedSharesEmitStruct({staker: defaultStaker, operator: defaultOperator, sharesToDecrease: operatorSharesToDecrease})
        );
        cheats.prank(address(eigenPodManagerMock));
        delegationManager.decreaseDelegatedShares(defaultStaker, uint(beaconShares), beaconChainSlashingFactorDecrease);

        // 3. Assert correct values
        uint expectedWithdrawableShares = _calcWithdrawableShares({
            depositShares: uint(beaconShares),
            depositScalingFactor: depositScalingFactor,
            slashingFactor: maxMagnitude.mulWad(newBeaconSlashingFactor)
        });
        _assertSharesAfterBeaconSlash({
            staker: defaultStaker,
            withdrawableSharesBefore: withdrawableShares[0],
            expectedWithdrawableShares: expectedWithdrawableShares,
            prevBeaconSlashingFactor: prevBeaconSlashingFactor
        });
        // Assert correct end state values
        (uint[] memory withdrawableSharesAfter,) = delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());

        assertEq(
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy) + operatorSharesToDecrease,
            withdrawableShares[0],
            "operator shares not decreased correctly"
        );

        _assertWithdrawableAndOperatorShares(
            withdrawableSharesAfter[0],
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy),
            "Invalid withdrawable shares"
        );
    }

    /**
     * @notice Verifies that if a staker's beaconChainSlashingFactor is reduced to 0 if their entire balance
     * is slashed. Their withdrawable shares should be 0 afterwards and decreasing operatorShares should
     * not underflow and revert either.
     */
    function testFuzz_decreaseDelegatedShares_entireBalance(Randomness r) public rand(r) {
        int beaconShares = int(r.Uint256(1, MAX_ETH_SUPPLY));
        uint64 maxMagnitude = r.Uint64(1, WAD);
        uint64 beaconChainSlashingFactor = r.Uint64(1, WAD);

        // 1. Setup staker and delegate to operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, maxMagnitude);
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, beaconShares);
        eigenPodManagerMock.setBeaconChainSlashingFactor(defaultStaker, beaconChainSlashingFactor);
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDelegation({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: withdrawableShares[0],
            depositSharesBefore: uint(beaconShares),
            prevDsf: WAD
        });

        // 2. Perform beaconChain slash + decreaseDelegatedShares()
        (uint64 prevBeaconSlashingFactor, uint64 newBeaconSlashingFactor) =
            _setNewBeaconChainSlashingFactor(defaultStaker, beaconShares, uint(beaconShares));
        assertEq(beaconChainSlashingFactor, prevBeaconSlashingFactor, "Bad test setup");
        uint64 beaconChainSlashingFactorDecrease = prevBeaconSlashingFactor - newBeaconSlashingFactor;
        uint depositScalingFactor = uint(WAD).divWad(maxMagnitude);
        // expected operatorShares decreased for event
        uint operatorSharesToDecrease = _calcWithdrawableShares({
            depositShares: uint(beaconShares),
            depositScalingFactor: depositScalingFactor,
            slashingFactor: maxMagnitude.mulWad(beaconChainSlashingFactorDecrease)
        });
        // expected events
        _decreaseDelegatedShares_expectEmit(
            DecreaseDelegatedSharesEmitStruct({staker: defaultStaker, operator: defaultOperator, sharesToDecrease: operatorSharesToDecrease})
        );
        cheats.prank(address(eigenPodManagerMock));
        delegationManager.decreaseDelegatedShares(defaultStaker, uint(beaconShares), prevBeaconSlashingFactor);

        // 3. Assert correct values
        uint expectedWithdrawableShares = _calcWithdrawableShares({
            depositShares: uint(beaconShares),
            depositScalingFactor: depositScalingFactor,
            slashingFactor: maxMagnitude.mulWad(newBeaconSlashingFactor)
        });
        assertEq(expectedWithdrawableShares, 0, "All shares should be slashed");
        assertEq(eigenPodManagerMock.beaconChainSlashingFactor(defaultStaker), 0, "beaconChainSlashingFactor should be 0");
        _assertSharesAfterBeaconSlash({
            staker: defaultStaker,
            withdrawableSharesBefore: withdrawableShares[0],
            expectedWithdrawableShares: expectedWithdrawableShares,
            prevBeaconSlashingFactor: prevBeaconSlashingFactor
        });
        assertEq(
            delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy) + operatorSharesToDecrease,
            withdrawableShares[0],
            "operator shares not decreased correctly"
        );
    }
}

contract DelegationManagerUnitTests_undelegate is DelegationManagerUnitTests {
    using SlashingLib for uint;
    using ArrayLib for *;
    using Math for uint;

    // @notice Verifies that undelegating is not possible when the "undelegation paused" switch is flipped
    function testFuzz_Revert_undelegate_paused(Randomness r) public rand(r) {
        address staker = r.Address();
        address operator = r.Address();
        _registerOperatorWithBaseDetails(operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);
        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_ENTER_WITHDRAWAL_QUEUE);

        cheats.prank(staker);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.undelegate(staker);
    }

    function testFuzz_Revert_undelegate_notDelegated(Randomness r) public rand(r) {
        address undelegatedStaker = r.Address();
        assertFalse(delegationManager.isDelegated(undelegatedStaker), "bad test setup");

        cheats.prank(undelegatedStaker);
        cheats.expectRevert(NotActivelyDelegated.selector);
        delegationManager.undelegate(undelegatedStaker);
    }

    // @notice Verifies that an operator cannot undelegate from themself (this should always be forbidden)
    function testFuzz_Revert_undelegate_stakerIsOperator(Randomness r) public rand(r) {
        address operator = r.Address();
        _registerOperatorWithBaseDetails(operator);

        cheats.prank(operator);
        cheats.expectRevert(OperatorsCannotUndelegate.selector);
        delegationManager.undelegate(operator);
    }

    /**
     * @notice verifies that `DelegationManager.undelegate` reverts if trying to undelegate an operator from themselves
     */
    function testFuzz_Revert_undelegate_operatorCannotForceUndelegateThemself(Randomness r) public rand(r) {
        address delegationApprover = r.Address();
        bool callFromOperatorOrApprover = r.Boolean();

        // register *this contract* as an operator with the default `delegationApprover`
        _registerOperatorWithDelegationApprover(defaultOperator);

        address caller;
        if (callFromOperatorOrApprover) caller = delegationApprover;
        else caller = defaultOperator;

        // try to call the `undelegate` function and check for reversion
        cheats.prank(caller);
        cheats.expectRevert(OperatorsCannotUndelegate.selector);
        delegationManager.undelegate(defaultOperator);
    }

    /**
     * @notice Verifies that the `undelegate` function has proper access controls (can only be called by the operator who the `staker` has delegated
     * to or the operator's `delegationApprover`), or the staker themselves
     */
    function testFuzz_Revert_undelegate_invalidCaller(Randomness r) public rand(r) {
        address invalidCaller = r.Address();
        address staker = r.Address();

        _registerOperatorWithDelegationApprover(defaultOperator);
        _delegateToOperatorWhoRequiresSig(staker, defaultOperator);

        cheats.prank(invalidCaller);
        cheats.expectRevert(CallerCannotUndelegate.selector);
        delegationManager.undelegate(staker);
    }

    /**
     * Staker is undelegated from an operator, via a call to `undelegate`, properly originating from the staker's address.
     * Reverts if the staker is themselves an operator (i.e. they are delegated to themselves)
     * Does nothing if the staker is already undelegated
     * Properly undelegates the staker, i.e. the staker becomes “delegated to” the zero address, and `isDelegated(staker)` returns ‘false’
     * Emits a `StakerUndelegated` event
     */
    function testFuzz_undelegate_noDelegateableShares(Randomness r) public rand(r) {
        address staker = r.Address();

        // register *this contract* as an operator and delegate from the `staker` to them
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(staker, delegationManager.delegatedTo(staker));
        cheats.prank(staker);
        bytes32[] memory withdrawalRoots = delegationManager.undelegate(staker);

        assertEq(withdrawalRoots.length, 0, "withdrawalRoot should be an empty array");
        assertEq(delegationManager.delegatedTo(staker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(staker), "staker not undelegated");
    }

    /**
     * @notice Verifies that the `undelegate` function allows for a force undelegation
     */
    function testFuzz_undelegate_forceUndelegation_noDelegateableShares(Randomness r) public rand(r) {
        address staker = r.Address();
        bytes32 salt = r.Bytes32();
        bool callFromOperatorOrApprover = r.Boolean();

        _registerOperatorWithDelegationApprover(defaultOperator);
        _delegateToOperatorWhoRequiresSig(staker, defaultOperator, salt);

        address caller;
        if (callFromOperatorOrApprover) caller = defaultApprover;
        else caller = defaultOperator;

        Withdrawal memory withdrawal;
        _undelegate_expectEmit_singleStrat(
            UndelegateEmitStruct({
                staker: staker,
                operator: defaultOperator,
                strategy: IStrategy(address(0)),
                depositSharesQueued: 0,
                operatorSharesDecreased: 0,
                withdrawal: withdrawal,
                withdrawalRoot: bytes32(0),
                depositScalingFactor: WAD,
                forceUndelegated: true
            })
        );
        cheats.prank(caller);
        bytes32[] memory withdrawalRoots = delegationManager.undelegate(staker);

        assertEq(withdrawalRoots.length, 0, "withdrawalRoot should be an empty array");
        assertEq(delegationManager.delegatedTo(staker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(staker), "staker not undelegated");
    }

    /**
     * @notice Verifies that the `undelegate` function properly queues a withdrawal for all shares of the staker
     */
    function testFuzz_undelegate_nonSlashedOperator(Randomness r) public rand(r) {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        IStrategy[] memory strategyArray = r.StrategyArray(1);
        IStrategy strategy = strategyArray[0];

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategy, shares);

        // register *this contract* as an operator and delegate from the `staker` to them
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });

        // Format queued withdrawal
        (, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategy, depositSharesToWithdraw: shares});

        // Undelegate the staker
        _undelegate_expectEmit_singleStrat(
            UndelegateEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategy,
                depositSharesQueued: shares,
                operatorSharesDecreased: shares,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot,
                depositScalingFactor: WAD,
                forceUndelegated: false
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(delegationManager.delegatedTo(defaultStaker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");
        // Checks - operator & staker shares
        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: shares,
            depositSharesBefore: shares,
            depositSharesWithdrawn: shares,
            depositScalingFactor: uint(WAD),
            slashingFactor: uint(WAD)
        });
    }

    /**
     * @notice Verifies that the `undelegate` function properly queues a withdrawal for all shares of the staker
     * @notice The operator should have its shares slashed prior to the staker's deposit
     */
    function testFuzz_undelegate_preSlashedOperator(Randomness r) public rand(r) {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint64 operatorMagnitude = r.Uint64(1, WAD);
        IStrategy strategy = IStrategy(r.Address());

        // register *this contract* as an operator & set its slashed magnitude
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategy, operatorMagnitude);

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategy, shares);

        // delegate from the `staker` to them
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: uint(WAD).divWad(operatorMagnitude),
            depositAmount: shares
        });

        // Format queued withdrawal
        (, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategy, depositSharesToWithdraw: shares});

        // Calculate operatorShares decreased, may be off of shares due to rounding
        uint depositScalingFactor = delegationManager.depositScalingFactor(defaultStaker, strategy);
        assertTrue(depositScalingFactor > WAD, "bad test setup");
        uint operatorSharesDecreased = _calcWithdrawableShares(shares, depositScalingFactor, operatorMagnitude);
        assertLe(operatorSharesDecreased, shares, "operatorSharesDecreased should be <= shares");

        // Undelegate the staker
        _undelegate_expectEmit_singleStrat(
            UndelegateEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategy,
                depositSharesQueued: shares,
                operatorSharesDecreased: operatorSharesDecreased,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot,
                depositScalingFactor: WAD,
                forceUndelegated: false
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(delegationManager.delegatedTo(defaultStaker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: shares,
            depositSharesBefore: shares,
            depositSharesWithdrawn: shares,
            depositScalingFactor: uint(WAD).divWad(operatorMagnitude),
            slashingFactor: uint(operatorMagnitude)
        });
        (uint[] memory stakerWithdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategy.toArray());
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
    }

    /**
     * @notice Verifies that the `undelegate` function properly queues a withdrawal for all shares of the staker
     * @notice The operator should have its shares slashed prior to the staker's deposit
     */
    function testFuzz_undelegate_slashedWhileStaked(Randomness r) public rand(r) {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint64 prevMaxMagnitude = r.Uint64(2, WAD);
        uint64 newMaxMagnitude = r.Uint64(1, prevMaxMagnitude - 1);
        IStrategy strategy = IStrategy(r.Address());

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategy, prevMaxMagnitude);

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategy, shares);

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });

        assertEq(delegationManager.operatorShares(defaultOperator, strategy), shares, "operatorShares should increment correctly");

        // Set operator magnitude
        {
            (uint[] memory withdrawableSharesBefore,) = delegationManager.getWithdrawableShares(defaultStaker, strategy.toArray());
            uint delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategy);
            _setOperatorMagnitude(defaultOperator, strategy, newMaxMagnitude);
            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, strategy, prevMaxMagnitude, newMaxMagnitude);
            (, uint operatorSharesAfterSlash) = _assertOperatorSharesAfterSlash({
                operator: defaultOperator,
                strategy: strategy,
                operatorSharesBefore: delegatedSharesBefore,
                prevMaxMagnitude: prevMaxMagnitude,
                newMaxMagnitude: newMaxMagnitude
            });

            uint expectedWithdrawable = _calcWithdrawableShares(
                shares, uint(WAD).divWad(prevMaxMagnitude), _getSlashingFactor(defaultStaker, strategy, newMaxMagnitude)
            );
            _assertSharesAfterSlash({
                staker: defaultStaker,
                strategy: strategy,
                withdrawableSharesBefore: withdrawableSharesBefore[0],
                expectedWithdrawableShares: expectedWithdrawable,
                prevMaxMagnitude: prevMaxMagnitude,
                currMaxMagnitude: newMaxMagnitude
            });

            // Get withdrawable shares
            (uint[] memory withdrawableSharesAfter, uint[] memory depositSharesAfter) =
                delegationManager.getWithdrawableShares(defaultStaker, strategy.toArray());
            _assertWithdrawableAndOperatorShares(withdrawableSharesAfter[0], operatorSharesAfterSlash, "Invalid withdrawable shares");
            assertEq(depositSharesAfter[0], shares, "Invalid deposit shares");
            assertEq(delegationManager.depositScalingFactor(defaultStaker, strategy), uint(WAD).divWad(prevMaxMagnitude), "bad test setup");
        }

        // Format queued withdrawal
        (uint[] memory withdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(defaultStaker, strategy.toArray());
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategy);
        {
            (, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
                _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategy, depositSharesToWithdraw: shares});

            // Undelegate the staker
            _undelegate_expectEmit_singleStrat(
                UndelegateEmitStruct({
                    staker: defaultStaker,
                    operator: defaultOperator,
                    strategy: strategy,
                    depositSharesQueued: shares,
                    operatorSharesDecreased: withdrawableShares[0],
                    withdrawal: withdrawal,
                    withdrawalRoot: withdrawalRoot,
                    depositScalingFactor: WAD,
                    forceUndelegated: false
                })
            );
            cheats.prank(defaultStaker);
            delegationManager.undelegate(defaultStaker);
        }

        // Checks - delegation status
        assertEq(delegationManager.delegatedTo(defaultStaker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: operatorSharesBefore,
            depositSharesBefore: shares,
            depositSharesWithdrawn: shares,
            depositScalingFactor: uint(WAD).divWad(prevMaxMagnitude),
            slashingFactor: uint(newMaxMagnitude)
        });

        (withdrawableShares, depositShares) = delegationManager.getWithdrawableShares(defaultStaker, strategy.toArray());
        assertEq(withdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
        assertEq(depositShares[0], 0, "staker deposit shares not reset correctly");
    }

    /**
     * @notice Verifies that the `undelegate` function properly undelegates a staker even though their shares
     * were slashed entirely.
     */
    function testFuzz_undelegate_slashedOperator100PercentWhileStaked(Randomness r) public rand(r) {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        IStrategy[] memory strategyArray = r.StrategyArray(1);
        IStrategy strategy = strategyArray[0];

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategy, shares);

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });

        // Set operator magnitude
        uint64 operatorMagnitude = 0;
        uint operatorSharesAfterSlash;
        {
            _setOperatorMagnitude(defaultOperator, strategy, operatorMagnitude);
            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, strategy, WAD, 0);
            operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategy);
            assertEq(operatorSharesAfterSlash, 0, "operator shares not fully slashed");
        }

        (, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategy, depositSharesToWithdraw: shares});

        uint depositScalingFactor = delegationManager.depositScalingFactor(defaultStaker, strategy);
        assertEq(depositScalingFactor, WAD, "bad test setup");
        // Get withdrawable and deposit shares
        {
            (uint[] memory withdrawableSharesBefore, uint[] memory depositSharesBefore) =
                delegationManager.getWithdrawableShares(defaultStaker, strategyArray);
            assertEq(withdrawableSharesBefore[0], 0, "withdrawable shares should be 0 after being slashed fully");
            assertEq(depositSharesBefore[0], shares, "deposit shares should be unchanged after being slashed fully");
        }

        // Undelegate the staker
        _undelegate_expectEmit_singleStrat(
            UndelegateEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategy,
                depositSharesQueued: shares,
                operatorSharesDecreased: 0,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot,
                depositScalingFactor: WAD,
                forceUndelegated: false
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(delegationManager.delegatedTo(defaultStaker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategy,
            operatorSharesBefore: 0,
            depositSharesBefore: shares,
            depositSharesWithdrawn: shares,
            depositScalingFactor: uint(WAD),
            slashingFactor: 0
        });

        assertEq(delegationManager.operatorShares(defaultOperator, strategy), 0, "operator shares not decreased correctly");
        (uint[] memory stakerWithdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(defaultStaker, strategyArray);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
        assertEq(depositShares[0], 0, "staker deposit shares not reset correctly");
    }

    function testFuzz_undelegate_slashedOperatorCloseTo100(Randomness r) public rand(r) {
        address[] memory stakers = r.StakerArray(r.Uint32(1, 8));
        uint64 prevMaxMagnitude = r.Uint64(2, WAD);
        uint64 newMaxMagnitude = 1;

        // 1. register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, prevMaxMagnitude);

        // 2. Stakers deposits in the strategyMock
        {
            for (uint i = 0; i < stakers.length; ++i) {
                uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
                strategyManagerMock.addDeposit(stakers[i], strategyMock, shares);
                stakerDepositShares[stakers[i]] = shares;
            }
        }

        // 3. Delegate from the `stakers` to the operator
        {
            uint totalWithdrawable = 0;
            for (uint i = 0; i < stakers.length; ++i) {
                {
                    uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
                    _delegateToOperatorWhoAcceptsAllStakers(stakers[i], defaultOperator);
                    _assertDeposit({
                        staker: stakers[i],
                        operator: defaultOperator,
                        strategy: strategyMock,
                        operatorSharesBefore: operatorSharesBefore,
                        withdrawableSharesBefore: 0,
                        depositSharesBefore: 0,
                        prevDsf: WAD,
                        depositAmount: stakerDepositShares[stakers[i]]
                    });
                }

                (uint[] memory withdrawableSharesBefore,) = delegationManager.getWithdrawableShares(stakers[i], strategyMock.toArray());
                totalWithdrawable += withdrawableSharesBefore[0];
            }
            assertLe(
                totalWithdrawable, delegationManager.operatorShares(defaultOperator, strategyMock), "should be <= op shares due to rounding"
            );
        }

        // 4. Slash operator - Set operator magnitude and call burnOperatorShares
        {
            uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
            _setOperatorMagnitude(defaultOperator, strategyMock, newMaxMagnitude);

            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, strategyMock, prevMaxMagnitude, newMaxMagnitude);
            _assertOperatorSharesAfterSlash({
                operator: defaultOperator,
                strategy: strategyMock,
                operatorSharesBefore: operatorSharesBefore,
                prevMaxMagnitude: prevMaxMagnitude,
                newMaxMagnitude: newMaxMagnitude
            });
        }

        // 5. Undelegate the stakers with expected events
        uint totalOperatorSharesDecreased = 0;
        for (uint i = 0; i < stakers.length; ++i) {
            (, Withdrawal memory withdrawal, bytes32 withdrawalRoot) = _setUpQueueWithdrawalsSingleStrat({
                staker: stakers[i],
                strategy: strategyMock,
                depositSharesToWithdraw: stakerDepositShares[stakers[i]]
            });
            uint operatorSharesDecreased = _calcWithdrawableShares(
                stakerDepositShares[stakers[i]], delegationManager.depositScalingFactor(stakers[i], strategyMock), newMaxMagnitude
            );
            _undelegate_expectEmit_singleStrat(
                UndelegateEmitStruct({
                    staker: stakers[i],
                    operator: defaultOperator,
                    strategy: strategyMock,
                    depositSharesQueued: stakerDepositShares[stakers[i]],
                    operatorSharesDecreased: operatorSharesDecreased,
                    withdrawal: withdrawal,
                    withdrawalRoot: withdrawalRoot,
                    depositScalingFactor: WAD,
                    forceUndelegated: false
                })
            );

            cheats.prank(stakers[i]);
            delegationManager.undelegate(stakers[i]);

            totalOperatorSharesDecreased += operatorSharesDecreased;
        }

        // 6. Checks - delegation status and staker,operator shares
        assertEq(delegationManager.delegatedTo(defaultStaker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");
        for (uint i = 0; i < stakers.length; ++i) {
            (uint[] memory stakerWithdrawableShares, uint[] memory stakerDepositShares) =
                delegationManager.getWithdrawableShares(stakers[i], strategyMock.toArray());
            assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
            assertEq(stakerDepositShares[0], 0, "staker deposit shares not reset correctly");
        }
    }

    /**
     * @notice Given an operator with slashed magnitude, delegate, undelegate, and then delegate back to the same operator with
     * completing withdrawals as shares. This should result in the operatorShares after the second delegation being <= the shares from the first delegation.
     */
    function testFuzz_undelegate_delegateAgainWithRounding(Randomness r) public rand(r) {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        // set magnitude to 66% to ensure rounding when calculating `toShares`
        uint64 operatorMagnitude = 333_333_333_333_333_333;

        // register *this contract* as an operator & set its slashed magnitude
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, shares);

        // delegate from the `staker` to them
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: shares
        });
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);

        // Format queued withdrawal
        (, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: shares});

        uint slashingFactor = _getSlashingFactor(defaultStaker, strategyMock, operatorMagnitude);
        uint operatorSharesDecreased =
            _calcWithdrawableShares(shares, delegationManager.depositScalingFactor(defaultStaker, strategyMock), slashingFactor);

        // Undelegate the staker
        cheats.prank(defaultStaker);
        _undelegate_expectEmit_singleStrat(
            UndelegateEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategyMock,
                depositSharesQueued: shares,
                operatorSharesDecreased: operatorSharesDecreased,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot,
                depositScalingFactor: WAD,
                forceUndelegated: false
            })
        );
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(delegationManager.delegatedTo(defaultStaker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");
        // Checks - operator & staker shares
        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: operatorSharesBefore,
            depositSharesBefore: shares,
            depositSharesWithdrawn: shares,
            depositScalingFactor: uint(WAD).divWad(operatorMagnitude),
            slashingFactor: operatorMagnitude
        });
        (uint[] memory stakerWithdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategyMock.toArray());
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");

        // // Re-delegate the staker to the operator again. The shares should have increased but may be less than from before due to rounding
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        // complete withdrawal as shares, should add back delegated shares to operator due to delegating again
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(strategyMock.underlyingToken());
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, false);

        uint operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        assertLe(
            operatorSharesAfter, operatorSharesBefore, "operator shares should be less than or equal to before due to potential rounding"
        );
    }
}

contract DelegationManagerUnitTests_redelegate is DelegationManagerUnitTests {
    using ArrayLib for *;

    // @notice Verifies that redelegating is not possible when the "delegation paused" switch is flipped
    function testFuzz_Revert_redelegate_delegatePaused(Randomness r) public {
        address staker = r.Address();
        address newOperator = r.Address();

        // register *this contract* as an operator and delegate from the `staker` to them
        _registerOperatorWithBaseDetails(defaultOperator);
        _registerOperatorWithBaseDetails(newOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        cheats.prank(staker);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.redelegate(newOperator, emptyApproverSignatureAndExpiry, emptySalt);
    }

    // @notice Verifies that redelegating is not possible when the "undelegation paused" switch is flipped
    function testFuzz_Revert_redelegate_undelegatePaused(Randomness r) public {
        address staker = r.Address();
        address newOperator = r.Address();

        // register *this contract* as an operator and delegate from the `staker` to them
        _registerOperatorWithBaseDetails(defaultOperator);
        _registerOperatorWithBaseDetails(newOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_ENTER_WITHDRAWAL_QUEUE);

        cheats.prank(staker);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.redelegate(newOperator, emptyApproverSignatureAndExpiry, emptySalt);
    }

    function testFuzz_Revert_redelegate_notDelegated(Randomness r) public {
        address undelegatedStaker = r.Address();
        assertFalse(delegationManager.isDelegated(undelegatedStaker), "bad test setup");

        _registerOperatorWithBaseDetails(defaultOperator);

        cheats.prank(undelegatedStaker);
        cheats.expectRevert(NotActivelyDelegated.selector);
        delegationManager.redelegate(defaultOperator, emptyApproverSignatureAndExpiry, emptySalt);
    }

    // @notice Verifies that an operator cannot undelegate from themself (this should always be forbidden)
    function testFuzz_Revert_redelegate_stakerIsOperator(Randomness r) public {
        address operator = r.Address();
        _registerOperatorWithBaseDetails(operator);
        _registerOperatorWithBaseDetails(defaultOperator);

        cheats.prank(operator);
        cheats.expectRevert(OperatorsCannotUndelegate.selector);
        delegationManager.redelegate(defaultOperator, emptyApproverSignatureAndExpiry, emptySalt);
    }

    /// @notice Verifies that `staker` cannot redelegate to an unregistered `operator`
    function testFuzz_Revert_redelegateToUnregisteredOperator(Randomness r) public {
        address staker = r.Address();
        address operator = r.Address();
        assertFalse(delegationManager.isOperator(operator), "incorrect test input?");

        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        // try to delegate and check that the call reverts
        cheats.prank(staker);
        cheats.expectRevert(OperatorNotRegistered.selector);
        delegationManager.redelegate(operator, emptyApproverSignatureAndExpiry, emptySalt);
    }

    function testFuzz_Revert_redelegate_ExpiredSignature(Randomness r) public {
        // roll to a very late timestamp
        skip(type(uint).max / 2);

        address staker = r.Address();
        address newOperator = r.Address();
        uint expiry = r.Uint256(0, block.timestamp - 1);
        bytes32 salt = r.Bytes32();

        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        _registerOperatorWithDelegationApprover(newOperator);

        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, newOperator, salt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(ISignatureUtilsMixinErrors.SignatureExpired.selector);
        delegationManager.redelegate(newOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    function testFuzz_Revert_redelegate_SpentSalt(Randomness r) public {
        address staker = r.Address();
        address newOperator = r.Address();
        uint expiry = r.Uint256(block.timestamp, block.timestamp + 100);
        bytes32 salt = r.Bytes32();

        _registerOperatorWithBaseDetails(defaultOperator);
        _registerOperatorWithDelegationApprover(newOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(newOperator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtilsMixinTypes.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, newOperator, salt, expiry);

        // Spend salt by delegating normally first
        cheats.startPrank(staker);
        delegationManager.delegateTo(newOperator, approverSignatureAndExpiry, salt);
        assertTrue(
            delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(newOperator), salt),
            "salt somehow spent not spent?"
        );

        // redelegate to a different operator
        delegationManager.redelegate(defaultOperator, emptyApproverSignatureAndExpiry, emptySalt);

        // Now try to redelegate to the original operator using the invalid signature
        cheats.expectRevert(SaltSpent.selector);
        delegationManager.redelegate(newOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that the `redelegate` function properly queues a withdrawal for all shares of the staker
     * ... and delegates to a new operator
     */
    function testFuzz_redelegate_noSlashing(Randomness r) public {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        IStrategy[] memory strategyArray = r.StrategyArray(1);
        IStrategy strategy = strategyArray[0];

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategy, shares);

        // register *this contract* as an operator and delegate from the `staker` to them
        address newOperator = r.Address();
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _registerOperatorWithBaseDetails(newOperator);

        // Format queued withdrawal
        (, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategy, depositSharesToWithdraw: shares});

        // Redelegate the staker
        _undelegate_expectEmit_singleStrat(
            UndelegateEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategy,
                depositSharesQueued: shares,
                operatorSharesDecreased: shares,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot,
                depositScalingFactor: WAD,
                forceUndelegated: false
            })
        );
        _delegateTo_expectEmit_singleStrat(
            DelegateToSingleStratEmitStruct({
                staker: defaultStaker,
                operator: newOperator,
                strategy: strategyMock,
                depositShares: 0,
                depositScalingFactor: WAD
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.redelegate(newOperator, emptyApproverSignatureAndExpiry, emptySalt);

        // Checks - delegation status
        assertEq(delegationManager.delegatedTo(defaultStaker), newOperator, "undelegated staker should be delegated to new operator");
        assertTrue(delegationManager.isDelegated(defaultStaker), "staker should still be delegated");

        // Checks - operator & staker shares
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "operator shares not decreased correctly");
        assertEq(delegationManager.operatorShares(newOperator, strategyMock), 0, "operator shares should not have been added");
        (uint[] memory stakerWithdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategyArray);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
    }

    /**
     * @notice This function tests to ensure that a delegator can re-delegate to an operator after undelegating.
     * Asserts the shares after re-delegating are the same as originally. No slashing is done in this test.
     */
    function testFuzz_undelegate_redelegateWithSharesBack(Randomness r) public rand(r) {
        address staker = r.Address();
        address operator = r.Address();
        uint strategyShares = r.Uint256(1, MAX_STRATEGY_SHARES);
        int beaconShares = int(r.Uint256(1 gwei, MAX_ETH_SUPPLY));
        bool completeAsShares = r.Boolean();

        // 1. Set staker shares
        strategyManagerMock.addDeposit(staker, strategyMock, strategyShares);
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        (IStrategy[] memory strategiesToReturn,) = delegationManager.getDepositedShares(staker);
        // 2. register operator and delegate staker to operator
        _registerOperatorWithBaseDetails(operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);
        _assertDeposit({
            staker: staker,
            operator: operator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: uint(WAD),
            depositAmount: strategyShares
        });
        _assertDeposit({
            staker: staker,
            operator: operator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: uint(WAD),
            depositAmount: uint(beaconShares)
        });

        // 3. Setup queued withdrawals from `undelegate`
        // queued withdrawals done for single strat as this is how undelegate queue withdraws
        (, Withdrawal memory strategyWithdrawal,) =
            _setUpQueueWithdrawalsSingleStrat({staker: staker, strategy: strategyMock, depositSharesToWithdraw: strategyShares});
        (, Withdrawal memory beaconWithdrawal,) = _setUpQueueWithdrawalsSingleStrat({
            staker: staker,
            strategy: IStrategy(address(beaconChainETHStrategy)),
            depositSharesToWithdraw: uint(beaconShares)
        });
        beaconWithdrawal.nonce = 1; // Ensure nonce is greater for second withdrawal
        cheats.prank(staker);
        delegationManager.undelegate(staker);
        // 4. Delegate to operator again with shares added back
        {
            cheats.roll(block.number + delegationManager.minWithdrawalDelayBlocks() + 1);
            IERC20[] memory strategyTokens = new IERC20[](1);
            strategyTokens[0] = IERC20(strategyMock.underlyingToken());
            IERC20[] memory beaconTokens = new IERC20[](1);
            beaconTokens[0] = IERC20(address(beaconChainETHStrategy));
            if (completeAsShares) {
                // delegate first and complete withdrawal
                _delegateToOperatorWhoAcceptsAllStakers(staker, operator);
                cheats.startPrank(staker);
                delegationManager.completeQueuedWithdrawal(strategyWithdrawal, strategyTokens, false);
                delegationManager.completeQueuedWithdrawal(beaconWithdrawal, beaconTokens, false);
                cheats.stopPrank();
            } else {
                // complete withdrawal first and then delegate
                cheats.startPrank(staker);
                delegationManager.completeQueuedWithdrawal(strategyWithdrawal, strategyTokens, false);
                delegationManager.completeQueuedWithdrawal(beaconWithdrawal, beaconTokens, false);
                cheats.stopPrank();
                _delegateToOperatorWhoAcceptsAllStakers(staker, operator);
            }
        }

        // 5. assert correct shares and delegation state
        assertTrue(delegationManager.isDelegated(staker), "staker should be delegated");
        assertEq(delegationManager.delegatedTo(staker), operator, "staker should be delegated to operator");
        (uint[] memory stakerShares,) = delegationManager.getWithdrawableShares(staker, strategiesToReturn);
        assertEq(
            delegationManager.operatorShares(operator, strategyMock), stakerShares[0], "operator shares should be equal to strategyShares"
        );
        assertEq(uint(eigenPodManagerMock.podOwnerDepositShares(staker)), stakerShares[1], "beacon shares should be equal to beaconShares");
    }
}

contract DelegationManagerUnitTests_queueWithdrawals is DelegationManagerUnitTests {
    using SlashingLib for *;
    using ArrayLib for *;

    function test_Revert_WhenEnterQueueWithdrawalsPaused() public {
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_ENTER_WITHDRAWAL_QUEUE);
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams,,) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: 100});
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenQueueWithdrawalParamsLengthMismatch() public {
        IStrategy[] memory strategyArray = strategyMock.toArray();
        uint[] memory shareAmounts = new uint[](2);
        shareAmounts[0] = 100;
        shareAmounts[1] = 100;

        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] =
            QueuedWithdrawalParams({strategies: strategyArray, depositShares: shareAmounts, __deprecated_withdrawer: address(0)});

        cheats.expectRevert(InputArrayLengthMismatch.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function testFuzz_IgnoresWithdrawerField(address withdrawer) public {
        _depositIntoStrategies(defaultStaker, strategyMock.toArray(), uint(100).toArrayU256());
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams,,) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: 100});

        // set the ignored field to a different address. the dm should ignore this.
        queuedWithdrawalParams[0].__deprecated_withdrawer = withdrawer;

        cheats.prank(defaultStaker);
        bytes32 withdrawalRoot = delegationManager.queueWithdrawals(queuedWithdrawalParams)[0];

        (Withdrawal memory withdrawal,) = delegationManager.getQueuedWithdrawal(withdrawalRoot);
        assertEq(withdrawal.staker, defaultStaker, "staker should be msg.sender");
        assertEq(withdrawal.withdrawer, defaultStaker, "withdrawer should be msg.sender");
    }

    function test_Revert_WhenEmptyStrategiesArray() public {
        IStrategy[] memory strategyArray = new IStrategy[](0);
        uint[] memory shareAmounts = new uint[](0);

        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] =
            QueuedWithdrawalParams({strategies: strategyArray, depositShares: shareAmounts, __deprecated_withdrawer: address(0)});

        cheats.expectRevert(InputArrayLengthZero.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount`.
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_SingleStrat_nonSlashedOperator(Randomness r) public rand(r) {
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint withdrawalAmount = r.Uint256(1, depositAmount);
        bool depositBeaconChainShares = r.Boolean();
        // sharesAmounts is single element so returns single strategy
        IStrategy[] memory strategies =
            _deployAndDepositIntoStrategies(defaultStaker, depositAmount.toArrayU256(), depositBeaconChainShares);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategies[0],
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: uint(WAD),
            depositAmount: depositAmount
        });
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategies[0], depositSharesToWithdraw: withdrawalAmount});
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[0]);

        // queueWithdrawals
        _queueWithdrawals_expectEmit(
            QueueWithdrawalsEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                queuedWithdrawalParams: queuedWithdrawalParams,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategies[0],
            operatorSharesBefore: delegatedSharesBefore,
            depositSharesBefore: depositAmount,
            depositSharesWithdrawn: withdrawalAmount,
            depositScalingFactor: uint(WAD),
            slashingFactor: uint(WAD)
        });
        _assertQueuedWithdrawalExists(defaultStaker, withdrawal);
        uint nonceAfter = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount`. Operator is slashed prior to the staker's deposit
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_SingleStrat_preSlashedOperator(Randomness r) public rand(r) {
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint withdrawalAmount = r.Uint256(1, depositAmount);
        uint64 maxMagnitude = r.Uint64(1, WAD);

        // Slash the operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, maxMagnitude);

        // Deposit for staker & delegate
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: depositAmount
        });

        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: withdrawalAmount});

        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);

        // queueWithdrawals
        _queueWithdrawals_expectEmit(
            QueueWithdrawalsEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                queuedWithdrawalParams: queuedWithdrawalParams,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: delegatedSharesBefore,
            depositSharesBefore: depositAmount,
            depositSharesWithdrawn: withdrawalAmount,
            depositScalingFactor: delegationManager.depositScalingFactor(defaultStaker, strategyMock),
            slashingFactor: uint(maxMagnitude)
        });
        _assertQueuedWithdrawalExists(defaultStaker, withdrawal);

        uint nonceAfter = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount`. Operator is slashed while the staker is deposited
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_SingleStrat_slashedWhileStaked(Randomness r) public rand(r) {
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint withdrawalAmount = r.Uint256(1, depositAmount);
        uint64 prevMaxMagnitude = r.Uint64(2, WAD);
        uint64 newMaxMagnitude = r.Uint64(1, prevMaxMagnitude - 1);

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, prevMaxMagnitude);

        // Deposit for staker & delegate
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: depositAmount
        });

        // Slash the operator
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        _setOperatorMagnitude(defaultOperator, strategyMock, newMaxMagnitude);
        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, prevMaxMagnitude, newMaxMagnitude);
        // Assertions on amount burned
        (uint operatorSharesSlashed,) = _assertOperatorSharesAfterSlash({
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: operatorSharesBefore,
            prevMaxMagnitude: prevMaxMagnitude,
            newMaxMagnitude: newMaxMagnitude
        });
        uint nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);

        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: withdrawalAmount});

            // queueWithdrawals
            _queueWithdrawals_expectEmit(
                QueueWithdrawalsEmitStruct({
                    staker: defaultStaker,
                    operator: defaultOperator,
                    queuedWithdrawalParams: queuedWithdrawalParams,
                    withdrawal: withdrawal,
                    withdrawalRoot: withdrawalRoot
                })
            );
            cheats.prank(defaultStaker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            _assertQueuedWithdrawalExists(defaultStaker, withdrawal);
        }

        uint slashingFactor = _getSlashingFactor(defaultStaker, strategyMock, newMaxMagnitude);
        assertEq(nonceBefore + 1, delegationManager.cumulativeWithdrawalsQueued(defaultStaker), "staker nonce should have incremented");
        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: depositAmount - operatorSharesSlashed,
            depositSharesBefore: depositAmount,
            depositSharesWithdrawn: withdrawalAmount,
            depositScalingFactor: uint(WAD).divWad(prevMaxMagnitude),
            slashingFactor: slashingFactor
        });
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` queues an empty withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount` since the Operator is slashed 100% while the staker is deposited
     * - Asserts that queuing a withdrawal results in an empty withdrawal when the operator is slashed 100%
     * - Asserts that staker withdrawableShares after is 0
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_SingleStrat_slashed100PercentWhileStaked(Randomness r) public rand(r) {
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);

        // Register operator, deposit for staker & delegate
        _registerOperatorWithBaseDetails(defaultOperator);
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: depositAmount
        });

        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
        _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            strategy: strategyMock,
            depositSharesToWithdraw: 0 // expected 0 since slashed 100%
        });

        // Slash the operator
        uint64 operatorMagnitude = 0;
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, WAD, 0);
        _assertOperatorSharesAfterSlash({
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: depositAmount,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: operatorMagnitude
        });
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");

        // queueWithdrawals should result in an empty withdrawal
        _queueWithdrawals_expectEmit(
            QueueWithdrawalsEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                queuedWithdrawalParams: queuedWithdrawalParams,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategyMock.toArray());
        assertEq(withdrawableShares[0], 0, "withdrawable shares should be 0 after being slashed fully");
        _assertWithdrawal({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: strategyMock,
            operatorSharesBefore: 0,
            depositSharesBefore: depositAmount,
            depositSharesWithdrawn: 0,
            depositScalingFactor: uint(WAD),
            slashingFactor: 0
        });
        _assertQueuedWithdrawalExists(defaultStaker, withdrawal);
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * with multiple strategies and sharesAmounts. Operator has default WAD maxMagnitude for all strategies.
     * Depending on number of strategies randomized, deposits sharesAmounts into each strategy for the staker and delegates to operator.
     * For each strategy,
     * - Asserts that staker is delegated to the operator
     * - Asserts that the staker withdrawal is queued both with the root and actual Withdrawal struct in storage
     * - Asserts that the operator shares decrease by the expected withdrawn shares
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_MultipleStrats_nonSlashedOperator(Randomness r) public rand(r) {
        uint32 numStrategies = r.Uint32(1, 32);
        bool depositBeaconChainShares = r.Boolean();

        (uint[] memory depositAmounts, uint[] memory withdrawalAmounts,,) =
            _fuzzDepositWithdrawalAmounts({r: r, numStrategies: numStrategies});
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(defaultStaker, depositAmounts, depositBeaconChainShares);

        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        for (uint i = 0; i < strategies.length; ++i) {
            _assertDeposit({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: 0,
                withdrawableSharesBefore: 0,
                depositSharesBefore: 0,
                prevDsf: 0,
                depositAmount: depositAmounts[i]
            });
        }

        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawals({staker: defaultStaker, strategies: strategies, depositWithdrawalAmounts: withdrawalAmounts});
        // Before queueWithdrawal state values
        uint nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint[] memory delegatedSharesBefore = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            delegatedSharesBefore[i] = delegationManager.operatorShares(defaultOperator, strategies[i]);
        }

        // queueWithdrawals
        _queueWithdrawals_expectEmit(
            QueueWithdrawalsEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                queuedWithdrawalParams: queuedWithdrawalParams,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Post queueWithdrawal state values
        for (uint i = 0; i < strategies.length; i++) {
            _assertWithdrawal({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: delegatedSharesBefore[i],
                depositSharesBefore: depositAmounts[i],
                depositSharesWithdrawn: withdrawalAmounts[i],
                depositScalingFactor: uint(WAD),
                slashingFactor: uint(WAD)
            });
        }
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint nonceAfter = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
        _assertQueuedWithdrawalExists(defaultStaker, withdrawal);
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * with multiple strategies and sharesAmounts. Operator has random maxMagnitudes for each strategy.
     * Depending on number of strategies randomized, deposits sharesAmounts into each strategy for the staker and delegates to operator.
     * For each strategy,
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `depositAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_MultipleStrats_preSlashedOperator(Randomness r) public rand(r) {
        // 1. Setup
        // - fuzz numbers of strategies, deposit and withdraw amounts, and prev/new magnitudes for each strategy respectively
        // - deposit into strategies, delegate to operator
        bool depositBeaconChainShares = r.Boolean();
        IStrategy[] memory strategies = r.StrategyArray(r.Uint32(1, 32));
        if (depositBeaconChainShares) strategies[strategies.length - 1] = beaconChainETHStrategy;

        (uint[] memory depositAmounts, uint[] memory withdrawalAmounts, uint64[] memory prevMaxMagnitudes,) =
            _fuzzDepositWithdrawalAmounts({r: r, numStrategies: uint32(strategies.length)});
        _registerOperatorWithBaseDetails(defaultOperator);
        allocationManagerMock.setMaxMagnitudes(defaultOperator, strategies, prevMaxMagnitudes);
        _depositIntoStrategies(defaultStaker, strategies, depositAmounts);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        // Check deposit state for all strategies after delegating
        for (uint i = 0; i < strategies.length; ++i) {
            _assertDeposit({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: 0,
                withdrawableSharesBefore: 0,
                depositSharesBefore: 0,
                prevDsf: WAD,
                depositAmount: depositAmounts[i]
            });
        }
        uint nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);

        // 2. Setup and call queued withdrawals
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
                _setUpQueueWithdrawals({staker: defaultStaker, strategies: strategies, depositWithdrawalAmounts: withdrawalAmounts});
            // expected events emitted
            _queueWithdrawals_expectEmit(
                QueueWithdrawalsEmitStruct({
                    staker: defaultStaker,
                    operator: defaultOperator,
                    queuedWithdrawalParams: queuedWithdrawalParams,
                    withdrawal: withdrawal,
                    withdrawalRoot: withdrawalRoot
                })
            );
            // 3. call `DelegationManager.queueWithdrawals`
            _queueWithdrawals(defaultStaker, queuedWithdrawalParams, withdrawal);
        }

        // 4. Post queueWithdrawal state values
        // Post queueWithdrawal state values
        for (uint i = 0; i < strategies.length; i++) {
            _assertWithdrawal({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: depositAmounts[i],
                depositSharesBefore: depositAmounts[i],
                depositSharesWithdrawn: withdrawalAmounts[i],
                depositScalingFactor: uint(WAD).divWad(prevMaxMagnitudes[i]),
                slashingFactor: prevMaxMagnitudes[i]
            });
        }
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        assertEq(nonceBefore + 1, delegationManager.cumulativeWithdrawalsQueued(defaultStaker), "staker nonce should have incremented");
        _assertQueuedWithdrawalExists(defaultStaker);
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * with multiple strategies and sharesAmounts. Operator has random maxMagnitudes for each strategy.
     * Depending on number of strategies randomized, deposits sharesAmounts into each strategy for the staker and delegates to operator.
     * After depositing, the operator gets slashed for each of the strategies and has new maxMagnitudes set.
     * For each strategy,
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `depositAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_MultipleStrats_slashedWhileStaked(Randomness r) public rand(r) {
        // 1. Setup
        // - fuzz numbers of strategies, deposit and withdraw amounts, and prev/new magnitudes for each strategy respectively
        // - deposit into strategies, delegate to operator
        IStrategy[] memory strategies = r.StrategyArray(r.Uint32(1, 32));
        bool depositBeaconChainShares = r.Boolean();
        if (depositBeaconChainShares) strategies[strategies.length - 1] = beaconChainETHStrategy;
        (uint[] memory depositAmounts, uint[] memory withdrawalAmounts, uint64[] memory prevMaxMagnitudes, uint64[] memory newMaxMagnitudes)
        = _fuzzDepositWithdrawalAmounts({r: r, numStrategies: uint32(strategies.length)});
        _registerOperatorWithBaseDetails(defaultOperator);
        allocationManagerMock.setMaxMagnitudes(defaultOperator, strategies, prevMaxMagnitudes);
        _depositIntoStrategies(defaultStaker, strategies, depositAmounts);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        // Check deposit state for all strategies after delegating
        for (uint i = 0; i < strategies.length; ++i) {
            _assertDeposit({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: 0,
                withdrawableSharesBefore: 0,
                depositSharesBefore: 0,
                prevDsf: WAD,
                depositAmount: depositAmounts[i]
            });
        }

        // 2. Slash operator while staker is delegated and staked
        allocationManagerMock.setMaxMagnitudes(defaultOperator, strategies, newMaxMagnitudes);
        cheats.startPrank(address(allocationManagerMock));
        uint nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint[] memory slashedOperatorShares = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[i]);
            delegationManager.slashOperatorShares(defaultOperator, strategies[i], prevMaxMagnitudes[i], newMaxMagnitudes[i]);
            // Assert correct amount of shares slashed from operator
            (slashedOperatorShares[i],) = _assertOperatorSharesAfterSlash({
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: operatorSharesBefore,
                prevMaxMagnitude: prevMaxMagnitudes[i],
                newMaxMagnitude: newMaxMagnitudes[i]
            });
        }
        cheats.stopPrank();

        // 3. Setup and call queued withdrawals
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
                _setUpQueueWithdrawals({staker: defaultStaker, strategies: strategies, depositWithdrawalAmounts: withdrawalAmounts});
            // expected events emitted
            _queueWithdrawals_expectEmit(
                QueueWithdrawalsEmitStruct({
                    staker: defaultStaker,
                    operator: defaultOperator,
                    queuedWithdrawalParams: queuedWithdrawalParams,
                    withdrawal: withdrawal,
                    withdrawalRoot: withdrawalRoot
                })
            );
            // 4. call `DelegationManager.queueWithdrawals`
            _queueWithdrawals(defaultStaker, queuedWithdrawalParams, withdrawal);
            _assertQueuedWithdrawalExists(defaultStaker, withdrawal);
        }

        // 5. Post queueWithdrawal state values
        for (uint i = 0; i < strategies.length; i++) {
            _assertWithdrawal({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: depositAmounts[i] - slashedOperatorShares[i],
                depositSharesBefore: depositAmounts[i],
                depositSharesWithdrawn: withdrawalAmounts[i],
                depositScalingFactor: uint(WAD).divWad(prevMaxMagnitudes[i]),
                slashingFactor: newMaxMagnitudes[i]
            });
        }
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        assertEq(nonceBefore + 1, delegationManager.cumulativeWithdrawalsQueued(defaultStaker), "staker nonce should have incremented");
    }

    /**
     * @notice Same test as `testFuzz_queueWithdrawal_MultipleStrats_slashedWhileStaked` but with one strategy having 0 newMaxMagnitude
     * - Asserts that the strategy with 0 newMaxMagnitude has 0 delegated shares before and after withdrawal
     * - Asserts that the staker withdrawn shares for the strategy with 0 newMaxMagnitude is 0
     */
    function testFuzz_queueWithdrawal_MultipleStrats__slashed100PercentWhileStaked(Randomness r) public rand(r) {
        // 1. Setup
        // - fuzz numbers of strategies, deposit and withdraw amounts, and prev/new magnitudes for each strategy respectively
        // - deposit into strategies, delegate to operator
        uint32 numStrats = r.Uint32(1, 32);
        IStrategy[] memory strategies = r.StrategyArray(numStrats);
        bool depositBeaconChainShares = r.Boolean();
        if (depositBeaconChainShares) strategies[numStrats - 1] = beaconChainETHStrategy;
        (uint[] memory depositAmounts, uint[] memory withdrawalAmounts, uint64[] memory prevMaxMagnitudes, uint64[] memory newMaxMagnitudes)
        = _fuzzDepositWithdrawalAmounts({r: r, numStrategies: numStrats});
        // randomly choose strategy to have 0 newMaxMagnitude
        uint zeroMagnitudeIndex = r.Uint256(0, numStrats - 1);
        newMaxMagnitudes[zeroMagnitudeIndex] = 0;

        _registerOperatorWithBaseDetails(defaultOperator);
        allocationManagerMock.setMaxMagnitudes(defaultOperator, strategies, prevMaxMagnitudes);
        _depositIntoStrategies(defaultStaker, strategies, depositAmounts);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        // Check deposit state for all strategies after delegating
        for (uint i = 0; i < strategies.length; ++i) {
            _assertDeposit({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: 0,
                withdrawableSharesBefore: 0,
                depositSharesBefore: 0,
                prevDsf: WAD,
                depositAmount: depositAmounts[i]
            });
        }

        // 2. Slash operator while staker is delegated and staked
        uint nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint[] memory slashedOperatorShares = new uint[](strategies.length);
        allocationManagerMock.setMaxMagnitudes(defaultOperator, strategies, newMaxMagnitudes);
        cheats.startPrank(address(allocationManagerMock));
        for (uint i = 0; i < strategies.length; i++) {
            uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[i]);
            delegationManager.slashOperatorShares(defaultOperator, strategies[i], prevMaxMagnitudes[i], newMaxMagnitudes[i]);

            // Assertions on amount burned
            (slashedOperatorShares[i],) = _assertOperatorSharesAfterSlash({
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: operatorSharesBefore,
                prevMaxMagnitude: prevMaxMagnitudes[i],
                newMaxMagnitude: newMaxMagnitudes[i]
            });
            // additional assertion checks for strategy that was slashed 100%
            if (zeroMagnitudeIndex == i) {
                assertEq(slashedOperatorShares[i], operatorSharesBefore, "expected slashed operator shares to be full amount");
                assertEq(delegationManager.operatorShares(defaultOperator, strategies[i]), 0, "expected operator shares to be 0");
            }
        }
        cheats.stopPrank();

        // 3. Setup and call queued withdrawals
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
                _setUpQueueWithdrawals({staker: defaultStaker, strategies: strategies, depositWithdrawalAmounts: withdrawalAmounts});
            // expected events emitted
            _queueWithdrawals_expectEmit(
                QueueWithdrawalsEmitStruct({
                    staker: defaultStaker,
                    operator: defaultOperator,
                    queuedWithdrawalParams: queuedWithdrawalParams,
                    withdrawal: withdrawal,
                    withdrawalRoot: withdrawalRoot
                })
            );
            // 4. call `DelegationManager.queueWithdrawals`
            _queueWithdrawals(defaultStaker, queuedWithdrawalParams, withdrawal);
            _assertQueuedWithdrawalExists(defaultStaker, withdrawal);
        }

        // 5. Post queueWithdrawal state values
        for (uint i = 0; i < strategies.length; i++) {
            if (zeroMagnitudeIndex == i) assertEq(newMaxMagnitudes[i], 0, "expected new max magnitude to be 0");
            _assertWithdrawal({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategies[i],
                operatorSharesBefore: depositAmounts[i] - slashedOperatorShares[i],
                depositSharesBefore: depositAmounts[i],
                depositSharesWithdrawn: withdrawalAmounts[i],
                depositScalingFactor: uint(WAD).divWad(prevMaxMagnitudes[i]),
                slashingFactor: zeroMagnitudeIndex == i ? 0 : newMaxMagnitudes[i]
            });
        }
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        assertEq(nonceBefore + 1, delegationManager.cumulativeWithdrawalsQueued(defaultStaker), "staker nonce should have incremented");
    }
}

contract DelegationManagerUnitTests_completeQueuedWithdrawal is DelegationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;
    using Math for uint;

    function test_Revert_WhenExitWithdrawalQueuePaused() public {
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_EXIT_WITHDRAWAL_QUEUE);
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100,
            isBeaconChainStrategy: false
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // single withdrawal interface
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, false);

        IERC20[][] memory tokensArray = new IERC20[][](1);
        tokensArray[0] = tokens;

        bool[] memory receiveAsTokens = new bool[](1);
        receiveAsTokens[0] = false;

        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = withdrawal;

        // multiple Withdrawal interface
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.completeQueuedWithdrawals(withdrawals, tokensArray, receiveAsTokens);
    }

    function test_Revert_WhenInputArrayLengthMismatch() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100,
            isBeaconChainStrategy: false
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Roll to completion block
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);

        // resize tokens array
        IERC20[] memory newTokens = new IERC20[](0);

        cheats.prank(defaultStaker);
        cheats.expectRevert(InputArrayLengthMismatch.selector);
        delegationManager.completeQueuedWithdrawal(withdrawal, newTokens, false);

        // check that the withdrawal completes otherwise
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, true);
    }

    function test_Revert_WhenWithdrawerNotCaller(Randomness r) public rand(r) {
        address invalidCaller = r.Address();

        _registerOperatorWithBaseDetails(defaultOperator);
        (Withdrawal memory withdrawal, IERC20[] memory tokens,) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100,
            isBeaconChainStrategy: false
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        cheats.expectRevert(WithdrawerNotCaller.selector);
        cheats.prank(invalidCaller);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, false);
    }

    function test_Revert_WhenInvalidWithdrawalRoot() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (Withdrawal memory withdrawal, IERC20[] memory tokens, bytes32 withdrawalRoot) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100,
            isBeaconChainStrategy: false
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, true);
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");

        cheats.expectRevert(WithdrawalNotQueued.selector);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, false);
    }

    /**
     * @notice should revert if MIN_WITHDRAWAL_DELAY_BLOCKS has not passed, and if
     * delegationManager.getCompletableTimestamp returns a value greater than MIN_WITHDRAWAL_DELAY_BLOCKS
     * then it should revert if the validBlockNumber has not passed either.
     */
    function test_Revert_WhenWithdrawalDelayNotPassed(Randomness r) public rand(r) {
        uint32 numStrategies = r.Uint32(1, 32);
        bool receiveAsTokens = r.Boolean();
        (uint[] memory depositAmounts, uint[] memory withdrawalAmounts,,) = _fuzzDepositWithdrawalAmounts(r, numStrategies);

        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawal({
            staker: defaultStaker,
            depositAmounts: depositAmounts,
            withdrawalAmounts: withdrawalAmounts,
            depositBeaconChainShares: false
        });

        // prank as withdrawer address
        cheats.roll(withdrawal.startBlock + MIN_WITHDRAWAL_DELAY_BLOCKS);
        cheats.expectRevert(WithdrawalDelayNotElapsed.selector);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);
    }

    /// @notice Verifies that when we complete a withdrawal as shares after a full slash, we clear the withdrawal
    function test_clearWithdrawal_fullySlashed() public {
        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, WAD);

        // Set the staker deposits in the strategies
        uint depositAmount = 100e18;
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawal
        uint withdrawalAmount = depositAmount;
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal,) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: withdrawalAmount});
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Warp to just before the MIN_WITHDRAWAL_DELAY_BLOCKS
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks());

        // Slash all of operator's shares
        _setOperatorMagnitude(defaultOperator, strategyMock, 0);
        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, WAD, 0);

        // Complete withdrawal as shares and check that withdrawal was cleared
        cheats.roll(block.number + 1);
        IERC20[] memory tokens = strategyMock.underlyingToken().toArray();

        bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawal should be pending before completion");

        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, false);

        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawal should be cleared after completion");

        // Assert that no shares were added back
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "operator shares should remain 0");
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategyMock.toArray());
        assertEq(withdrawableShares[0], 0, "withdrawable shares should be 0");
    }

    /**
     * Test completing multiple queued withdrawals for a single strategy by passing in the withdrawals
     */
    function test_completeQueuedWithdrawals_MultipleWithdrawals(Randomness r) public rand(r) {
        address staker = r.Address();
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint numWithdrawals = r.Uint256(2, 20);
        bool receiveAsTokens = r.Boolean();

        (Withdrawal[] memory withdrawals, IERC20[][] memory tokens, bytes32[] memory withdrawalRoots) =
            _setUpCompleteQueuedWithdrawalsSingleStrat({staker: staker, depositAmount: depositAmount, numWithdrawals: numWithdrawals});

        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, withdrawals[0].strategies[0]);

        for (uint i = 0; i < withdrawalRoots.length; i++) {
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoots[i]), "withdrawalRoots should be pending");
        }
        bool[] memory receiveAsTokensArray = receiveAsTokens.toArray(numWithdrawals);

        // completeQueuedWithdrawal
        cheats.roll(withdrawals[0].startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        _completeQueuedWithdrawals_expectEmit(
            CompleteQueuedWithdrawalsEmitStruct({withdrawals: withdrawals, tokens: tokens, receiveAsTokens: receiveAsTokensArray})
        );
        cheats.prank(staker);
        delegationManager.completeQueuedWithdrawals(withdrawals, tokens, receiveAsTokensArray);

        // assertion checks
        (uint[] memory withdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(staker, withdrawals[0].strategies);
        uint operatorSharesAfter = delegationManager.operatorShares(defaultOperator, withdrawals[0].strategies[0]);
        if (receiveAsTokens) {
            assertEq(withdrawableShares[0], 0, "withdrawable shares should be 0 from withdrawing all");
            assertEq(depositShares[0], 0, "deposit shares should be 0 from withdrawing all");
            assertEq(operatorSharesAfter, operatorSharesBefore, "operator shares should be unchanged");
        } else {
            assertEq(withdrawableShares[0], depositAmount * numWithdrawals, "withdrawable shares should be added back as shares");
            assertEq(depositShares[0], depositAmount * numWithdrawals, "deposit shares should be added back as shares");
            assertEq(operatorSharesAfter, operatorSharesBefore + depositAmount * numWithdrawals, "operator shares should be unchanged");
        }
        _assertWithdrawalRootsComplete(staker, withdrawals);
        assertEq(
            delegationManager.depositScalingFactor(staker, withdrawals[0].strategies[0]), uint(WAD), "deposit scaling factor should be WAD"
        );
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for a single strategy.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts that event `WithdrawalCompleted` is emitted with withdrawalRoot
     * if receiveAsTokens is true
     * - Asserts operatorShares is unchanged after `completeQueuedWithdrawal`
     * - Asserts that the staker's withdrawable shares, deposit shares, and depositScalingFactors are unchanged
     * if receiveAsTokens is false
     * - Asserts operatorShares is increased correctly after `completeQueuedWithdrawal`
     * - Asserts that the staker's withdrawable shares, deposit shares, and depositScalingFactors are updated correctly
     */
    function test_completeQueuedWithdrawal_SingleStrat(Randomness r) public rand(r) {
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint withdrawalAmount = r.Uint256(1, depositAmount);
        bool receiveAsTokens = r.Boolean();

        _registerOperatorWithBaseDetails(defaultOperator);
        (Withdrawal memory withdrawal, IERC20[] memory tokens, bytes32 withdrawalRoot) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            depositAmount: depositAmount,
            withdrawalAmount: withdrawalAmount,
            isBeaconChainStrategy: false
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");

        // completeQueuedWithdrawal
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        _completeQueuedWithdrawal_expectEmit(
            CompleteQueuedWithdrawalEmitStruct({withdrawal: withdrawal, tokens: tokens, receiveAsTokens: receiveAsTokens})
        );
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);

        _assertCompletedWithdrawal(
            AssertCompletedWithdrawalStruct({
                staker: defaultStaker,
                currentOperator: defaultOperator,
                withdrawal: withdrawal,
                receiveAsTokens: receiveAsTokens,
                operatorSharesBefore: operatorSharesBefore.toArrayU256(),
                withdrawableSharesBefore: (depositAmount - withdrawalAmount).toArrayU256(),
                depositSharesBefore: (depositAmount - withdrawalAmount).toArrayU256(),
                prevDepositScalingFactors: uint(WAD).toArrayU256(),
                slashingFactors: uint(WAD).toArrayU256(),
                beaconChainSlashingFactor: WAD
            })
        );
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for a single strategy. Withdraws as tokens so there are no operator shares increase.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts operatorShares is decreased after the operator is slashed
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     * - Asserts that the shares the staker completed withdrawal for are less than what is expected since its operator is slashed
     */
    function test_completeQueuedWithdrawal_SingleStrat_slashOperatorDuringQueue(Randomness r) public rand(r) {
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint withdrawalAmount = r.Uint256(1, depositAmount);
        bool receiveAsTokens = r.Boolean();
        uint64 prevMaxMagnitude = r.Uint64(2, WAD);
        uint64 newMaxMagnitude = r.Uint64(1, prevMaxMagnitude - 1);

        // Deposit Staker
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, prevMaxMagnitude);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawal
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: withdrawalAmount});
        {
            uint operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, strategyMock);
            cheats.prank(defaultStaker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);

            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
            uint operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, strategyMock);
            uint sharesWithdrawn = _calcWithdrawableShares({
                depositShares: withdrawalAmount,
                depositScalingFactor: uint(WAD).divWad(prevMaxMagnitude),
                slashingFactor: uint(prevMaxMagnitude)
            });
            assertEq(
                operatorSharesAfterQueue, operatorSharesBeforeQueue - sharesWithdrawn, "operator shares should be decreased after queue"
            );
        }

        // Slash operator while staker has queued withdrawal
        {
            uint operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, strategyMock);
            (uint sharesToDecrement,) = _calcSlashedAmount({
                operatorShares: operatorSharesAfterQueue,
                prevMaxMagnitude: prevMaxMagnitude,
                newMaxMagnitude: newMaxMagnitude
            });
            _setOperatorMagnitude(defaultOperator, strategyMock, newMaxMagnitude);
            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, withdrawal.strategies[0], prevMaxMagnitude, newMaxMagnitude);
            uint operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
            assertEq(
                operatorSharesAfterSlash, operatorSharesAfterQueue - sharesToDecrement, "operator shares should be decreased after slash"
            );
        }

        // Complete queue withdrawal
        (uint[] memory withdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(defaultStaker, withdrawal.strategies);
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        {
            IERC20[] memory tokens = new IERC20[](1);
            tokens[0] = IERC20(strategyMock.underlyingToken());
            cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
            _completeQueuedWithdrawal_expectEmit(
                CompleteQueuedWithdrawalEmitStruct({withdrawal: withdrawal, tokens: tokens, receiveAsTokens: receiveAsTokens})
            );
            cheats.prank(defaultStaker);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);
        }

        _assertCompletedWithdrawal(
            AssertCompletedWithdrawalStruct({
                staker: defaultStaker,
                currentOperator: defaultOperator,
                withdrawal: withdrawal,
                receiveAsTokens: receiveAsTokens,
                operatorSharesBefore: operatorSharesBefore.toArrayU256(),
                withdrawableSharesBefore: withdrawableShares,
                depositSharesBefore: depositShares,
                prevDepositScalingFactors: uint(WAD).divWad(prevMaxMagnitude).toArrayU256(),
                slashingFactors: uint(newMaxMagnitude).toArrayU256(),
                beaconChainSlashingFactor: WAD
            })
        );
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for the BeaconChainStrategy. Withdraws as tokens so there are no operator shares increase.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts operatorShares is decreased after staker is slashed
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     * - Asserts that the shares the staker completed withdrawal for are less than what is expected since the staker is slashed during queue
     */
    function test_completeQueuedWithdrawal_BeaconStrat_slashStakerDuringQueue(Randomness r) public rand(r) {
        int depositAmount = int(r.Uint256(1, MAX_ETH_SUPPLY));
        uint withdrawalAmount = r.Uint256(1, uint(depositAmount));
        uint64 initialBCSF = r.Uint64(2, WAD);
        uint sharesDecrease = r.Uint256(1, uint(depositAmount) - withdrawalAmount);
        bool receiveAsTokens = r.Boolean();

        // Deposit Staker
        eigenPodManagerMock.setBeaconChainSlashingFactor(defaultStaker, initialBCSF);
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, depositAmount);

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawal
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
        _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            strategy: beaconChainETHStrategy,
            depositSharesToWithdraw: withdrawalAmount
        });
        uint64 prevBeaconSlashingFactor;
        uint64 newBeaconSlashingFactor;
        {
            uint sharesToDecrement = _calcWithdrawableShares({
                depositShares: withdrawalAmount,
                depositScalingFactor: uint(WAD),
                slashingFactor: uint(initialBCSF)
            });
            uint operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            cheats.prank(defaultStaker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
            uint operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            assertEq(
                operatorSharesAfterQueue, operatorSharesBeforeQueue - sharesToDecrement, "operator shares should be decreased after queue"
            );

            // Slash the staker for beacon chain shares while it has queued a withdrawal
            // simulate the operations done in EigenPodManager._reduceSlashingFactor
            (uint[] memory withdrawableSharesBefore,) =
                delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());

            uint currentPodShares = uint(depositAmount) - withdrawalAmount;
            (prevBeaconSlashingFactor, newBeaconSlashingFactor) =
                _decreaseBeaconChainShares({staker: defaultStaker, beaconShares: int(currentPodShares), sharesDecrease: sharesDecrease});

            uint expectedWithdrawableShares = _calcWithdrawableShares({
                depositShares: uint(currentPodShares),
                depositScalingFactor: uint(WAD),
                slashingFactor: uint(newBeaconSlashingFactor)
            });
            _assertSharesAfterBeaconSlash(defaultStaker, withdrawableSharesBefore[0], expectedWithdrawableShares, prevBeaconSlashingFactor);
        }

        // Complete queue withdrawal
        (uint[] memory withdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);

        {
            IERC20[] memory tokens = new IERC20[](1);
            cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
            _completeQueuedWithdrawal_expectEmit(
                CompleteQueuedWithdrawalEmitStruct({withdrawal: withdrawal, tokens: tokens, receiveAsTokens: receiveAsTokens})
            );
            cheats.prank(defaultStaker);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);
        }

        _assertCompletedWithdrawal(
            AssertCompletedWithdrawalStruct({
                staker: defaultStaker,
                currentOperator: defaultOperator,
                withdrawal: withdrawal,
                receiveAsTokens: receiveAsTokens,
                operatorSharesBefore: operatorSharesBefore.toArrayU256(),
                withdrawableSharesBefore: withdrawableShares,
                depositSharesBefore: depositShares,
                prevDepositScalingFactors: uint(WAD).toArrayU256(),
                slashingFactors: uint(WAD).toArrayU256(), // beaconChainSlashingFactor is separate from slashingFactors input
                beaconChainSlashingFactor: newBeaconSlashingFactor
            })
        );
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for the BeaconChainStrategy. Withdraws as tokens so there are no operator shares increase.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts operatorShares is decreased after staker is slashed and after the operator is slashed
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     * - Asserts that the shares the staker completed withdrawal for are less than what is expected since both the staker and its operator are slashed during queue
     */
    function test_completeQueuedWithdrawal_BeaconStratWithdrawAsTokens_slashStakerAndOperator(Randomness r) public rand(r) {
        int depositAmount = int(r.Uint256(1, MAX_ETH_SUPPLY));
        uint withdrawalAmount = r.Uint256(1, uint(depositAmount));
        bool receiveAsTokens = r.Boolean();

        // Deposit Staker
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, int(depositAmount));

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);

        // Queue withdrawal
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
        _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            strategy: beaconChainETHStrategy,
            depositSharesToWithdraw: withdrawalAmount
        });

        uint64 newBeaconSlashingFactor;
        {
            cheats.prank(defaultStaker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
            uint operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            assertEq(
                operatorSharesAfterQueue, operatorSharesBeforeQueue - withdrawalAmount, "operator shares should be decreased after queue"
            );

            // Slash the staker for beacon chain shares while it has queued a withdrawal
            // simulate the operations done in EigenPodManager._reduceSlashingFactor
            (, newBeaconSlashingFactor) = _decreaseBeaconChainShares({
                staker: defaultStaker,
                beaconShares: depositAmount - int(withdrawalAmount),
                sharesDecrease: (uint(depositAmount) - withdrawalAmount) / 2
            });
            uint operatorSharesAfterBeaconSlash = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            assertEq(
                operatorSharesAfterBeaconSlash,
                operatorSharesAfterQueue.ceilDiv(2),
                "operator shares should be decreased after beaconChain slash"
            );

            // Slash the operator for beacon chain shares
            uint64 operatorMagnitude = 5e17;
            _setOperatorMagnitude(defaultOperator, withdrawal.strategies[0], operatorMagnitude);
            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, withdrawal.strategies[0], WAD, operatorMagnitude);
            uint operatorSharesAfterAVSSlash = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            assertApproxEqAbs(
                operatorSharesAfterAVSSlash, operatorSharesAfterBeaconSlash / 2, 1, "operator shares should be decreased after AVS slash"
            );
        }

        // Complete queue withdrawal
        (uint[] memory withdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        IERC20[] memory tokens = new IERC20[](1);
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        _completeQueuedWithdrawal_expectEmit(
            CompleteQueuedWithdrawalEmitStruct({withdrawal: withdrawal, tokens: tokens, receiveAsTokens: receiveAsTokens})
        );
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);

        _assertCompletedWithdrawal(
            AssertCompletedWithdrawalStruct({
                staker: defaultStaker,
                currentOperator: defaultOperator,
                withdrawal: withdrawal,
                receiveAsTokens: receiveAsTokens,
                operatorSharesBefore: operatorSharesBefore.toArrayU256(),
                withdrawableSharesBefore: withdrawableShares,
                depositSharesBefore: depositShares,
                prevDepositScalingFactors: uint(WAD).toArrayU256(),
                slashingFactors: 5e17.toArrayU256(),
                beaconChainSlashingFactor: newBeaconSlashingFactor
            })
        );
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for a single strategy. Withdraws as shares so if the withdrawer is delegated, operator shares increase. In the test case, this only
     * happens if staker and withdrawer are fuzzed the same address (i.e. staker == withdrawer)
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts if staker == withdrawer, operatorShares increase, otherwise operatorShares are unchanged
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     */
    function testFuzz_completeQueuedWithdrawal_SingleStratWithdrawAsShares_nonSlashedOperator(Randomness r) public rand(r) {
        address staker = r.Address();
        uint128 depositAmount = r.Uint128();
        uint128 withdrawalAmount = r.Uint128(1, depositAmount);

        _registerOperatorWithBaseDetails(defaultOperator);

        (Withdrawal memory withdrawal, IERC20[] memory tokens, bytes32 withdrawalRoot) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: staker,
            depositAmount: depositAmount,
            withdrawalAmount: withdrawalAmount,
            isBeaconChainStrategy: false
        });
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint operatorSharesBefore = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");

        // Set delegationManager on strategyManagerMock so it can call back into delegationManager
        strategyManagerMock.setDelegationManager(delegationManager);

        // completeQueuedWithdrawal
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        _completeQueuedWithdrawal_expectEmit(
            CompleteQueuedWithdrawalEmitStruct({withdrawal: withdrawal, tokens: tokens, receiveAsTokens: false})
        );
        cheats.prank(staker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, false);

        uint operatorSharesAfter = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        // Since staker is delegated, operatorShares get incremented
        assertEq(operatorSharesAfter, operatorSharesBefore + withdrawalAmount, "operator shares not increased correctly");
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");
    }

    function testFuzz_completeQueuedWithdrawals_OutOfOrderBlocking(Randomness r) public {
        uint totalDepositShares = r.Uint256(4, 100 ether);
        uint depositSharesPerWithdrawal = totalDepositShares / 4;

        _registerOperatorWithBaseDetails(defaultOperator);
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, totalDepositShares);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        QueuedWithdrawalParams[] memory queuedParams = new QueuedWithdrawalParams[](4);
        Withdrawal[] memory withdrawals = new Withdrawal[](4);

        uint startBlock = block.number;

        uint nonce = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        for (uint i; i < 4; ++i) {
            cheats.roll(startBlock + i);
            (QueuedWithdrawalParams[] memory params, Withdrawal memory withdrawal,) =
                _setUpQueueWithdrawalsSingleStrat(defaultStaker, strategyMock, depositSharesPerWithdrawal);
            withdrawal.nonce = nonce;
            nonce += 1;

            (queuedParams[i], withdrawals[i]) = (params[0], withdrawal);
        }

        uint delay = delegationManager.minWithdrawalDelayBlocks();

        cheats.startPrank(defaultStaker);
        cheats.roll(startBlock);

        delegationManager.queueWithdrawals(queuedParams[0].toArray());
        cheats.roll(startBlock + 1);
        delegationManager.queueWithdrawals(queuedParams[1].toArray());

        (Withdrawal[] memory firstWithdrawals,) = delegationManager.getQueuedWithdrawals(defaultStaker);

        cheats.roll(startBlock + 2);
        delegationManager.queueWithdrawals(queuedParams[2].toArray());
        cheats.roll(startBlock + 3);
        delegationManager.queueWithdrawals(queuedParams[3].toArray());

        IERC20[][] memory tokens = new IERC20[][](2);
        for (uint i; i < 2; ++i) {
            tokens[i] = strategyMock.underlyingToken().toArray();
        }

        bytes32 root1 = delegationManager.calculateWithdrawalRoot(withdrawals[0]);
        bytes32 root2 = delegationManager.calculateWithdrawalRoot(withdrawals[1]);

        bytes32 root1_view = delegationManager.calculateWithdrawalRoot(firstWithdrawals[0]);
        bytes32 root2_view = delegationManager.calculateWithdrawalRoot(firstWithdrawals[1]);

        assertEq(root1, root1_view, "withdrawal root should be the same");

        assertEq(root2, root2_view, "withdrawal root should be the same");

        cheats.roll(startBlock + delay + 2);
        delegationManager.completeQueuedWithdrawals(firstWithdrawals, tokens, true.toArray(2));

        // Throws `WithdrawalNotQueued`.
        cheats.roll(startBlock + delay + 3);
        delegationManager.completeQueuedWithdrawals(withdrawals[2].toArray(), tokens, true.toArray());
        cheats.stopPrank();
    }
}

contract DelegationManagerUnitTests_slashingShares is DelegationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;
    using Math for *;

    /// @notice Verifies that `DelegationManager.slashOperatorShares` reverts if not called by the AllocationManager
    function testFuzz_Revert_slashOperatorShares_invalidCaller(Randomness r) public rand(r) {
        address invalidCaller = r.Address();

        cheats.startPrank(invalidCaller);
        cheats.expectRevert(IDelegationManagerErrors.OnlyAllocationManager.selector);
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, 0, 0);
    }

    /// @notice Verifies that there is no change in shares if the staker is not delegatedd
    function testFuzz_Revert_slashOperatorShares_noop() public {
        _registerOperatorWithBaseDetails(defaultOperator);

        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, WAD, WAD / 2);
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "shares should not have changed");
    }

    /// @notice Verifies that shares are burnable for a withdrawal slashed just before the MIN_WITHDRAWAL_DELAY_BLOCKS is hit
    function test_sharesBurnableAtMinDelayBlocks() public {
        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, WAD);

        // Set the staker deposits in the strategies
        uint depositAmount = 100e18;
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawal
        uint withdrawalAmount = depositAmount;
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal,) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: withdrawalAmount});
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Warp to just before the MIN_WITHDRAWAL_DELAY_BLOCKS
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks());

        uint slashableSharesInQueue = delegationManager.getSlashableSharesInQueue(defaultOperator, strategyMock);

        // Slash all of operator's shares
        _setOperatorMagnitude(defaultOperator, strategyMock, 0);
        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, WAD, 0);

        uint slashableSharesInQueueAfter = delegationManager.getSlashableSharesInQueue(defaultOperator, strategyMock);

        // Complete withdrawal as tokens and assert that nothing is returned and withdrawal is cleared
        cheats.roll(block.number + 1);
        IERC20[] memory tokens = strategyMock.underlyingToken().toArray();

        bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawal should be pending before completion");

        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, true);

        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawal should be cleared after completion");

        assertEq(
            slashableSharesInQueue,
            depositAmount,
            "the withdrawal in queue from block.number - minWithdrawalDelayBlocks should still be included"
        );
        assertEq(slashableSharesInQueueAfter, 0, "slashable shares in queue should be 0 after burning");
    }

    /// @notice Verifies that shares are NOT burnable for a withdrawal queued just before the MIN_WITHDRAWAL_DELAY_BLOCKS
    function test_sharesNotBurnableWhenWithdrawalCompletable() public {
        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, WAD);

        // Set the staker deposits in the strategies
        uint depositAmount = 100e18;
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawal
        uint withdrawalAmount = depositAmount;
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal,) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: withdrawalAmount});
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Warp to completion time
        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        uint slashableShares = delegationManager.getSlashableSharesInQueue(defaultOperator, strategyMock);
        assertEq(slashableShares, 0, "shares should not be slashable");

        // Slash all of operator's shares
        _setOperatorMagnitude(defaultOperator, strategyMock, 0);
        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, WAD, 0);

        // Complete withdrawal as tokens and assert that we call back into the SM with 100 tokens
        IERC20[] memory tokens = strategyMock.underlyingToken().toArray();
        cheats.expectCall(
            address(strategyManagerMock),
            abi.encodeWithSelector(
                IShareManager.withdrawSharesAsTokens.selector, defaultStaker, strategyMock, strategyMock.underlyingToken(), 100e18
            )
        );
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, true);
    }

    /**
     * @notice Queues 5 withdrawals at different blocks. Then, warps such that the first 2 are completable. Validates the slashable shares
     */
    function test_slashableSharesInQueue() public {
        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, WAD);

        // Set the staker deposits in the strategies
        uint depositAmount = 120e18;
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue 5 withdrawals
        uint startBlock = block.number;
        uint withdrawalAmount = depositAmount / 6;
        for (uint i = 0; i < 5; i++) {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams,,) = _setUpQueueWithdrawalsSingleStrat({
                staker: defaultStaker,
                strategy: strategyMock,
                depositSharesToWithdraw: withdrawalAmount
            });
            cheats.prank(defaultStaker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            cheats.roll(startBlock + i + 1);
        }

        // Warp to completion time for the first 2 withdrawals
        // First withdrawal queued at startBlock. Second queued at startBlock + 1
        cheats.roll(startBlock + 1 + delegationManager.minWithdrawalDelayBlocks() + 1);

        // Get slashable shares
        uint slashableSharesInQueue = delegationManager.getSlashableSharesInQueue(defaultOperator, strategyMock);
        assertEq(slashableSharesInQueue, depositAmount / 6 * 3, "slashable shares in queue should be 3/6 of the deposit amount");

        // Slash all of operator's shares
        _setOperatorMagnitude(defaultOperator, strategyMock, 0);
        cheats.prank(address(allocationManagerMock));
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(
            defaultOperator,
            address(0),
            strategyMock,
            depositAmount / 6 // 1 withdrawal not queued so decreased
        );
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, WAD, 0);

        // Assert slashable shares
        slashableSharesInQueue = delegationManager.getSlashableSharesInQueue(defaultOperator, strategyMock);
        assertEq(slashableSharesInQueue, 0);
    }

    /**
     * @notice Verifies that `DelegationManager.slashOperatorShares` properly decreases the delegated `shares` that the operator
     * who the `defaultStaker` is delegated to has in the strategies
     */
    function testFuzz_slashOperatorShares_slashedOperator(Randomness r) public rand(r) {
        // sanity-filtering on fuzzed input length & staker
        IStrategy[] memory strategies = r.StrategyArray(16);
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint64 prevMaxMagnitude = r.Uint64(2, WAD);
        uint64 newMaxMagnitude = r.Uint64(1, prevMaxMagnitude);
        bool hasBeaconChainStrategy = r.Boolean();
        if (hasBeaconChainStrategy) {
            // Set last strategy in array as  beacon chain strategy
            strategies[strategies.length - 1] = beaconChainETHStrategy;
        }

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Set the staker deposits in the strategies
        uint[] memory sharesToSet = new uint[](strategies.length);
        uint[] memory depositScalingFactors = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            strategies[i] = IStrategy(random().Address());
            sharesToSet[i] = shares;
            depositScalingFactors[i] = uint(WAD).divWad(uint(prevMaxMagnitude));
            _setOperatorMagnitude(defaultOperator, strategies[i], prevMaxMagnitude);
        }

        // Okay to set beacon chain shares in SM mock, wont' be called by DM
        strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        if (hasBeaconChainStrategy) eigenPodManagerMock.setPodOwnerShares(defaultStaker, int(uint(shares)));

        // events expected emitted for each strategy
        _delegateTo_expectEmit(
            DelegateToEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategies: strategies,
                depositShares: sharesToSet,
                depositScalingFactors: depositScalingFactors
            })
        );
        // delegate from the `staker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        address delegatedTo = delegationManager.delegatedTo(defaultStaker);

        // check shares before call to `slashOperatorShares`
        for (uint i = 0; i < strategies.length; ++i) {
            // store delegated shares in a mapping
            delegatedSharesBefore[strategies[i]] = delegationManager.operatorShares(delegatedTo, strategies[i]);
            // also construct an array which we'll use in another loop
            totalSharesForStrategyInArray[address(strategies[i])] += shares;
        }

        // for each strategy in `strategies`, decrease delegated shares by `shares`
        {
            cheats.startPrank(address(allocationManagerMock));
            for (uint i = 0; i < strategies.length; ++i) {
                uint currentShares = delegationManager.operatorShares(defaultOperator, strategies[i]);
                uint sharesToDecrease = SlashingLib.calcSlashedAmount({
                    operatorShares: currentShares,
                    prevMaxMagnitude: prevMaxMagnitude,
                    newMaxMagnitude: newMaxMagnitude
                });

                cheats.expectEmit(true, true, true, true, address(delegationManager));
                emit OperatorSharesDecreased(defaultOperator, address(0), strategies[i], sharesToDecrease);
                delegationManager.slashOperatorShares(defaultOperator, strategies[i], prevMaxMagnitude, newMaxMagnitude);

                // Also update maxMagnitude in ALM mock
                _setOperatorMagnitude(defaultOperator, strategies[i], newMaxMagnitude);

                totalSharesDecreasedForStrategy[strategies[i]] += sharesToDecrease;
            }
            cheats.stopPrank();
        }

        // check shares after call to `slashOperatorShares`
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        for (uint i = 0; i < strategies.length; ++i) {
            uint delegatedSharesAfter = delegationManager.operatorShares(delegatedTo, strategies[i]);
            assertEq(
                delegatedSharesAfter,
                delegatedSharesBefore[strategies[i]] - totalSharesDecreasedForStrategy[strategies[i]],
                "delegated shares did not decrement correctly"
            );

            _assertWithdrawableAndOperatorShares(
                withdrawableShares[i], delegatedSharesAfter, "withdrawable and operator shares not decremented correctly"
            );
        }
    }

    /**
     * @notice Test burning shares for an operator with no queued withdrawals
     * - Asserts slashable shares before and after in queue is 0
     * - Asserts operator shares are decreased by half
     */
    function testFuzz_slashOperatorShares_NoQueuedWithdrawals(Randomness r) public rand(r) {
        address operator = r.Address();
        address staker = r.Address();
        uint64 initMagnitude = WAD;
        uint64 newMagnitude = 5e17;
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(operator);
        _setOperatorMagnitude(operator, strategyMock, initMagnitude);
        // Set the staker deposits in the strategies
        IStrategy[] memory strategyArray = strategyMock.toArray();
        uint[] memory sharesArray = shares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategyArray, sharesArray);
        // delegate from the `staker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);
        uint operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
        uint queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

        // calculate burned shares, should be halved
        uint sharesToBurn = shares / 2;

        // Burn shares
        _slashOperatorShares_expectEmit(
            SlashOperatorSharesEmitStruct({
                operator: operator,
                strategy: strategyMock,
                sharesToDecrease: sharesToBurn,
                sharesToBurn: sharesToBurn
            })
        );

        // Assert OperatorSharesSlashed event was emitted with correct params
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesSlashed(operator, strategyMock, sharesToBurn);

        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: initMagnitude,
            newMaxMagnitude: newMagnitude
        });

        uint queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, strategyMock);
        uint operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(queuedSlashableSharesBefore, 0, "there should be no slashable shares in queue");
        assertEq(queuedSlashableSharesAfter, 0, "there should be no slashable shares in queue");
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToBurn, "operator shares should be decreased by sharesToBurn");
    }

    /**
     * @notice Test burning shares for an operator with no slashable queued withdrawals in past MIN_WITHDRAWAL_DELAY_BLOCKS window.
     * There does exist past queued withdrawals but nothing in the queue is slashable.
     * - Asserts slashable shares in queue right after queuing a withdrawal is the withdrawal amount
     * and then checks that after the withdrawal window the slashable shares is 0 again.
     * - Asserts operator shares are decreased by half after burning
     * - Asserts that the slashable shares in queue before/after burning are 0
     */
    function testFuzz_slashOperatorShares_NoQueuedWithdrawalsInWindow(Randomness r) public rand(r) {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 5e17;
        // First staker
        address staker1 = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        // Second Staker, will queue withdraw shares
        address staker2 = r.Address();
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint withdrawAmount = r.Uint256(1, depositAmount);

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        strategyManagerMock.addDeposit(staker1, strategyMock, shares);
        strategyManagerMock.addDeposit(staker2, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(staker1, operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker2, operator);

        // 3. Queue withdrawal for staker2 and roll blocks forward so that the withdrawal is not slashable
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal,) =
                _setUpQueueWithdrawalsSingleStrat({staker: staker2, strategy: strategyMock, depositSharesToWithdraw: withdrawAmount});
            cheats.prank(staker2);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                withdrawAmount,
                "there should be withdrawAmount slashable shares in queue"
            );
            cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        }

        uint operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
        uint queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

        // calculate burned shares, should be halved
        // staker2 queue withdraws shares and we roll blocks to after the withdrawal is no longer slashable.
        // Therefore amount of shares to burn should be what the staker still has remaining + staker1 shares and then
        // divided by 2 since the operator was slashed 50%
        uint sharesToBurn = (shares + depositAmount - withdrawAmount) / 2;

        // 4. Burn shares
        _setOperatorMagnitude(operator, strategyMock, newMagnitude);
        _slashOperatorShares_expectEmit(
            SlashOperatorSharesEmitStruct({
                operator: operator,
                strategy: strategyMock,
                sharesToDecrease: sharesToBurn,
                sharesToBurn: sharesToBurn
            })
        );

        // Assert OperatorSharesSlashed event was emitted with correct params
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesSlashed(operator, strategyMock, sharesToBurn);

        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, strategyMock);
        uint operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(queuedSlashableSharesBefore, 0, "there should be no slashable shares in queue");
        assertEq(queuedSlashableSharesAfter, 0, "there should be no slashable shares in queue");
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToBurn, "operator shares should be decreased by sharesToBurn");
    }

    /**
     * @notice Test burning shares for an operator with slashable queued withdrawals in past MIN_WITHDRAWAL_DELAY_BLOCKS window.
     * There exists a single withdrawal that is slashable.
     */
    function testFuzz_slashOperatorShares_SingleSlashableWithdrawal(Randomness r) public rand(r) {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 25e16;
        // First staker
        address staker1 = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        // Second Staker, will queue withdraw shares
        address staker2 = r.Address();
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint withdrawAmount = r.Uint256(1, depositAmount);

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        strategyManagerMock.addDeposit(staker1, strategyMock, shares);
        strategyManagerMock.addDeposit(staker2, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(staker1, operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker2, operator);

        // 3. Queue withdrawal for staker2 so that the withdrawal is slashable
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams,,) =
                _setUpQueueWithdrawalsSingleStrat({staker: staker2, strategy: strategyMock, depositSharesToWithdraw: withdrawAmount});
            cheats.prank(staker2);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                withdrawAmount,
                "there should be withdrawAmount slashable shares in queue"
            );
        }

        uint operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
        uint queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

        // calculate burned shares, should be 3/4 of the original shares
        // staker2 queue withdraws shares
        // Therefore amount of shares to burn should be what the staker still has remaining + staker1 shares and then
        // divided by 2 since the operator was slashed 50%
        uint sharesToDecrease = (shares + depositAmount - withdrawAmount) * 3 / 4;
        uint sharesToBurn = sharesToDecrease + withdrawAmount * 3 / 4;

        // 4. Burn shares
        _setOperatorMagnitude(operator, strategyMock, newMagnitude);
        _slashOperatorShares_expectEmit(
            SlashOperatorSharesEmitStruct({
                operator: operator,
                strategy: strategyMock,
                sharesToDecrease: sharesToDecrease,
                sharesToBurn: sharesToBurn
            })
        );

        // Assert OperatorSharesSlashed event was emitted with correct params
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesSlashed(operator, strategyMock, sharesToBurn);

        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, strategyMock);
        uint operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(queuedSlashableSharesBefore, withdrawAmount, "Slashable shares in queue should be full withdraw amount");
        assertEq(queuedSlashableSharesAfter, withdrawAmount / 4, "Slashable shares in queue should be 1/4 withdraw amount after slashing");
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToDecrease, "operator shares should be decreased by sharesToBurn");
    }

    /**
     * @notice Test burning shares for an operator with slashable queued withdrawals in past MIN_WITHDRAWAL_DELAY_BLOCKS window.
     * There exists multiple withdrawals that are slashable.
     */
    function testFuzz_slashOperatorShares_MultipleSlashableWithdrawals(Randomness r) public rand(r) {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 25e16;
        // Staker and withdrawing amounts
        address staker = r.Address();
        uint depositAmount = r.Uint256(3, MAX_STRATEGY_SHARES);
        uint withdrawAmount1 = r.Uint256(2, depositAmount);
        uint withdrawAmount2 = r.Uint256(1, depositAmount - withdrawAmount1);

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        strategyManagerMock.addDeposit(staker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // 3. Queue withdrawal for staker and roll blocks forward so that the withdrawal is not slashable
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal,) =
                _setUpQueueWithdrawalsSingleStrat({staker: staker, strategy: strategyMock, depositSharesToWithdraw: withdrawAmount1});
            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                withdrawAmount1,
                "there should be withdrawAmount slashable shares in queue"
            );

            (queuedWithdrawalParams, withdrawal,) =
                _setUpQueueWithdrawalsSingleStrat({staker: staker, strategy: strategyMock, depositSharesToWithdraw: withdrawAmount2});
            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                withdrawAmount2 + withdrawAmount1,
                "there should be withdrawAmount slashable shares in queue"
            );
        }

        uint operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
        uint queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

        // calculate burned shares, should be halved for both operatorShares and slashable shares in queue
        // staker queue withdraws shares twice and both withdrawals should be slashed 75%.
        uint sharesToDecrease = (depositAmount - withdrawAmount1 - withdrawAmount2) * 3 / 4;
        uint sharesToBurn = sharesToDecrease + (delegationManager.getSlashableSharesInQueue(operator, strategyMock) * 3 / 4);

        // 4. Burn shares
        _setOperatorMagnitude(operator, strategyMock, newMagnitude);
        _slashOperatorShares_expectEmit(
            SlashOperatorSharesEmitStruct({
                operator: operator,
                strategy: strategyMock,
                sharesToDecrease: sharesToDecrease,
                sharesToBurn: sharesToBurn
            })
        );

        // Assert OperatorSharesSlashed event was emitted with correct params
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesSlashed(operator, strategyMock, sharesToBurn);

        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, strategyMock);
        uint operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(
            queuedSlashableSharesBefore, (withdrawAmount1 + withdrawAmount2), "Slashable shares in queue should be full withdraw amount"
        );
        assertEq(
            queuedSlashableSharesAfter,
            (withdrawAmount1 + withdrawAmount2) / 4,
            "Slashable shares in queue should be 1/4 withdraw amount after slashing"
        );
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToDecrease, "operator shares should be decreased by sharesToBurn");
    }

    /**
     * @notice Test burning shares for an operator with slashable queued withdrawals in past MIN_WITHDRAWAL_DELAY_BLOCKS window.
     * There exists multiple withdrawals that are slashable but queued with different maxMagnitudes at
     * time of queuing.
     *
     * Test Setup:
     * - staker1 deposits, queues withdrawal for some amount,
     * - operator slashed 50%
     * - staker 2 deposits, queues withdrawal for some amount
     * - operator is then slashed another 50%
     * slashed amount for staker 1 should be 75% and staker 2 should be 50% where the total
     * slashed amount is the sum of both
     */
    function testFuzz_slashOperatorShares_MultipleWithdrawalsMultipleSlashings(Randomness r) public rand(r) {
        address operator = r.Address();
        address staker = r.Address();
        uint depositAmount = r.Uint256(3, MAX_STRATEGY_SHARES);
        uint depositSharesToWithdraw1 = r.Uint256(1, depositAmount);
        uint depositSharesToWithdraw2 = r.Uint256(1, depositAmount - depositSharesToWithdraw1);

        uint64 newMagnitude = 50e16;

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        strategyManagerMock.addDeposit(staker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // 3. Queue withdrawal for staker and slash operator for 50%
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams,,) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker,
                strategy: strategyMock,
                depositSharesToWithdraw: depositSharesToWithdraw1
            });

            // 3.1 queue a withdrawal for the staker
            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            uint operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
            uint queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

            uint sharesToDecrease = (depositAmount - depositSharesToWithdraw1) / 2;
            uint sharesToBurn = sharesToDecrease + depositSharesToWithdraw1 / 2;

            // 3.2 Burn shares
            _setOperatorMagnitude(operator, strategyMock, newMagnitude);
            _slashOperatorShares_expectEmit(
                SlashOperatorSharesEmitStruct({
                    operator: operator,
                    strategy: strategyMock,
                    sharesToDecrease: sharesToDecrease,
                    sharesToBurn: sharesToBurn
                })
            );

            // Assert OperatorSharesSlashed event was emitted with correct params
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesSlashed(operator, strategyMock, sharesToBurn);

            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares({
                operator: operator,
                strategy: strategyMock,
                prevMaxMagnitude: WAD,
                newMaxMagnitude: newMagnitude
            });

            // 3.3 Assert slashable shares and operator shares
            assertEq(queuedSlashableSharesBefore, depositSharesToWithdraw1, "Slashable shares in queue should be full withdraw1 amount");
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                depositSharesToWithdraw1 / 2,
                "Slashable shares in queue should be 1/2 withdraw1 amount after slashing"
            );
            assertEq(
                delegationManager.operatorShares(operator, strategyMock),
                operatorSharesBefore - sharesToDecrease,
                "operator shares should be decreased by sharesToBurn"
            );
        }

        // 4. Queue withdrawal for staker and slash operator for 60% again
        newMagnitude = 25e16;
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams,,) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker,
                strategy: strategyMock,
                depositSharesToWithdraw: depositSharesToWithdraw2
            });

            // 4.1 queue a withdrawal for the staker
            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            uint operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
            uint queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

            uint sharesToDecrease = operatorSharesBefore / 2;
            uint sharesToBurn = sharesToDecrease + (depositSharesToWithdraw1 + depositSharesToWithdraw2) / 4;

            // 4.2 Burn shares
            _setOperatorMagnitude(operator, strategyMock, newMagnitude);
            _slashOperatorShares_expectEmit(
                SlashOperatorSharesEmitStruct({
                    operator: operator,
                    strategy: strategyMock,
                    sharesToDecrease: sharesToDecrease,
                    sharesToBurn: sharesToBurn
                })
            );

            // Assert OperatorSharesSlashed event was emitted with correct params
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesSlashed(operator, strategyMock, sharesToBurn);

            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares({
                operator: operator,
                strategy: strategyMock,
                prevMaxMagnitude: newMagnitude * 2,
                newMaxMagnitude: newMagnitude
            });

            // 4.3 Assert slashable shares and operator shares
            assertEq(
                queuedSlashableSharesBefore,
                (depositSharesToWithdraw1 + depositSharesToWithdraw2) / 2,
                "Slashable shares in queue before should be both queued withdrawal amounts halved"
            );
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                queuedSlashableSharesBefore / 2,
                "Slashable shares in queue should be halved again after slashing"
            );
            assertEq(
                delegationManager.operatorShares(operator, strategyMock),
                operatorSharesBefore - sharesToDecrease,
                "operator shares should be decreased by sharesToBurn"
            );
        }
    }

    /**
     * @notice Ensure that when a withdrawal is completable then there are no slashable shares in the queue.
     * However if the withdrawal is not completable and the withdrawal delay hasn't elapsed, then the withdrawal
     * should be counted as slashable.
     */
    function testFuzz_slashOperatorShares_Timings(Randomness r) public rand(r) {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 25e16;
        // staker
        address staker = r.Address();
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);

        // 2. Register the operator, set the staker deposits, and delegate the staker to them
        _registerOperatorWithBaseDetails(operator);
        strategyManagerMock.addDeposit(staker, strategyMock, depositAmount);
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // 3. Queue withdrawal for staker and roll blocks forward so that the withdrawal is completable
        uint completableBlock;
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal,) =
                _setUpQueueWithdrawalsSingleStrat({staker: staker, strategy: strategyMock, depositSharesToWithdraw: depositAmount});
            cheats.startPrank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            // 3.1 after queuing the withdrawal, check that there are slashable shares in queue
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                depositAmount,
                "there should be depositAmount slashable shares in queue"
            );
            // Check slashable shares in queue before and when the withdrawal is completable
            completableBlock = withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1;
            IERC20[] memory tokenArray = strategyMock.underlyingToken().toArray();

            // 3.2 roll to right before withdrawal is completable, check that slashable shares are still there
            // attempting to complete a withdrawal should revert
            cheats.roll(completableBlock - 1);
            cheats.expectRevert(WithdrawalDelayNotElapsed.selector);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokenArray, true);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                depositAmount,
                "there should still be depositAmount slashable shares in queue"
            );

            // 3.3 roll to blocknumber that the withdrawal is completable, there should be no slashable shares in queue
            cheats.roll(completableBlock);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokenArray, true);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                0,
                "there should be no slashable shares in queue when the withdrawal is completable"
            );

            cheats.stopPrank();
        }

        uint operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);

        // 4. Burn 0 shares when new magnitude is set
        _setOperatorMagnitude(operator, strategyMock, newMagnitude);
        _slashOperatorShares_expectEmit(
            SlashOperatorSharesEmitStruct({operator: operator, strategy: strategyMock, sharesToDecrease: 0, sharesToBurn: 0})
        );

        // Assert OperatorSharesSlashed event was emitted with correct params
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesSlashed(operator, strategyMock, 0);

        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(
            delegationManager.getSlashableSharesInQueue(operator, strategyMock),
            0,
            "there should still be no slashable shares in queue after burning 0 shares"
        );
        assertEq(operatorSharesAfter, operatorSharesBefore, "operator shares should be unchanged and equal to 0");
        assertEq(operatorSharesBefore, 0, "operator shares should be unchanged and equal to 0");
    }

    /**
     * @notice Ensure that no burning takes place for the beaconChainETHStrategy when the operator is slashed
     * and there are no slashable shares in the queue. Note: this will be implemented in a future release with
     * consideration of the Pectra upgrade.
     */
    function testFuzz_slashOperatorShares_BeaconChainStrategy(Randomness r) public rand(r) {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 25e16;
        // First staker
        address staker1 = r.Address();
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        // Second Staker, will queue withdraw shares
        address staker2 = r.Address();
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint withdrawAmount = r.Uint256(1, depositAmount);

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        eigenPodManagerMock.setPodOwnerShares(staker1, int(shares));
        eigenPodManagerMock.setPodOwnerShares(staker2, int(depositAmount));
        _delegateToOperatorWhoAcceptsAllStakers(staker1, operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker2, operator);

        // 3. Queue withdrawal for staker2 so that the withdrawal is slashable
        {
            (QueuedWithdrawalParams[] memory queuedWithdrawalParams,,) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker2,
                strategy: beaconChainETHStrategy,
                depositSharesToWithdraw: withdrawAmount
            });
            cheats.prank(staker2);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, beaconChainETHStrategy),
                withdrawAmount,
                "there should be withdrawAmount slashable shares in queue"
            );
        }

        uint operatorSharesBefore = delegationManager.operatorShares(operator, beaconChainETHStrategy);
        uint queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, beaconChainETHStrategy);

        // calculate burned shares, should be 3/4 of the original shares
        // staker2 queue withdraws shares
        // Therefore amount of shares to burn should be what the staker still has remaining + staker1 shares and then
        // divided by 2 since the operator was slashed 50%
        uint sharesToDecrease = (shares + depositAmount - withdrawAmount) * 3 / 4;
        uint sharesToBurn = sharesToDecrease + (delegationManager.getSlashableSharesInQueue(operator, beaconChainETHStrategy) * 3 / 4);

        // 4. Burn shares
        _setOperatorMagnitude(operator, beaconChainETHStrategy, newMagnitude);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(operator, address(0), beaconChainETHStrategy, sharesToDecrease);

        // Assert OperatorSharesSlashed event was emitted with correct params
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesSlashed(operator, beaconChainETHStrategy, sharesToBurn);

        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares({
            operator: operator,
            strategy: beaconChainETHStrategy,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, beaconChainETHStrategy);
        uint operatorSharesAfter = delegationManager.operatorShares(operator, beaconChainETHStrategy);
        assertEq(queuedSlashableSharesBefore, withdrawAmount, "Slashable shares in queue should be full withdraw amount");
        assertEq(queuedSlashableSharesAfter, withdrawAmount / 4, "Slashable shares in queue should be 1/4 withdraw amount after slashing");
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToDecrease, "operator shares should be decreased by sharesToDecrease");
    }

    /**
     * @notice This test demonstrates that the rate that withdrawable shares decrease from slashing is at LEAST
     * greater than or equal to the rate that the operator shares decrease from slashing.
     * We want this property otherwise undelegating/queue withdrawing all shares as a staker could lead to a underflow revert.
     * Note: If the SlashingLib.calcSlashedAmount function were to round down (overslash) then this test would fail.
     */
    function test_slashOperatorShares_slashedRepeatedly() public {
        uint64 initialMagnitude = 90_009;
        uint shares = 40_000_000_004_182_209_037_560_531_097_078_597_505;

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initialMagnitude);

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, shares);

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Set operator magnitude
        uint64 newOperatorMagnitude = initialMagnitude;

        for (uint i = 0; i < 100; ++i) {
            uint64 slashMagnitude = 100;
            newOperatorMagnitude -= slashMagnitude;
            _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);

            // Assert OperatorSharesSlashed event was emitted with correct params
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesSlashed(defaultOperator, strategyMock, 44_440_000_449_046_438_731_194_137_360_795_695);

            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(
                defaultOperator, strategyMock, newOperatorMagnitude + slashMagnitude, newOperatorMagnitude
            );

            uint operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
            (uint[] memory withdrawableShares, uint[] memory depositShares) =
                delegationManager.getWithdrawableShares(defaultStaker, strategyMock.toArray());
            assertEq(depositShares[0], shares, "staker deposit shares not reset correctly");
            assertLe(withdrawableShares[0], operatorSharesAfterSlash, "withdrawable should always be <= operatorShares even after rounding");
        }
    }

    /**
     * @notice This unit test will slash a staker's beaconChainETH strategy shares from both on EigenLayer
     * and also on the beaconChain. This test ensures that the order of slashing does not matter and nets
     * the same withdrawableShares for the staker whether slashing occurred on the beaconChain, or on EigenLayer first.
     */
    function testFuzz_beaconSlashAndAVSSlash(Randomness r) public rand(r) {
        uint64 initialMagnitude = r.Uint64(2, WAD);
        uint64 newMaxMagnitude = r.Uint64(1, initialMagnitude);
        // note: beaconShares only goes negative when performing withdrawal -- and this will change post-migration
        // so it's ok to make this assumption of positive shares
        int beaconShares = int(r.Uint256(2, MAX_ETH_SUPPLY));
        uint sharesDecrease = r.Uint256(1, uint(beaconShares) - 1);

        ////////////////////////////
        // 0. setup operator and staker with Beacon Chain stake
        ////////////////////////////
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, initialMagnitude);
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, beaconShares);
        // delegate staker to operator with expected events emitted
        _delegateTo_expectEmit(
            DelegateToEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategies: beaconChainETHStrategy.toArray(),
                depositShares: uint(beaconShares).toArrayU256(),
                depositScalingFactors: uint(WAD).divWad(initialMagnitude).toArrayU256()
            })
        );
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _assertDeposit({
            staker: defaultStaker,
            operator: defaultOperator,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: uint(beaconShares)
        });

        uint[] memory withdrawableShares;
        uint64 newBeaconSlashingFactor;
        // withdrawable shares after both slashing, this will be checked with the other scenario when
        // slashing in reverse order
        uint sharesAfterAllSlashing;

        ////////////////////////////
        // 1. do beacon chain slash then AVS slash
        ////////////////////////////
        {
            // Slash beaconChain first
            {
                (withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());
                uint beaconSharesBeforeSlash = withdrawableShares[0];

                uint64 prevBeaconChainSlashingFactor;
                (prevBeaconChainSlashingFactor, newBeaconSlashingFactor) =
                    _decreaseBeaconChainShares(defaultStaker, beaconShares, sharesDecrease);

                uint expectedWithdrawableShares = _calcWithdrawableShares({
                    depositShares: uint(beaconShares),
                    depositScalingFactor: uint(WAD).divWad(initialMagnitude),
                    slashingFactor: initialMagnitude.mulWad(newBeaconSlashingFactor)
                });
                _assertSharesAfterBeaconSlash({
                    staker: defaultStaker,
                    withdrawableSharesBefore: beaconSharesBeforeSlash,
                    expectedWithdrawableShares: expectedWithdrawableShares,
                    prevBeaconSlashingFactor: prevBeaconChainSlashingFactor
                });
            }
            // Slash on EigenLayer second
            {
                (withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());
                uint beaconSharesBeforeSlash = withdrawableShares[0];

                // do a slash via an AVS
                _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, newMaxMagnitude);
                cheats.prank(address(allocationManagerMock));
                delegationManager.slashOperatorShares(defaultOperator, beaconChainETHStrategy, initialMagnitude, newMaxMagnitude);

                // save the outcome
                (withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, beaconChainETHStrategy.toArray());
                uint beaconSharesAfterSecondSlash = withdrawableShares[0];
                uint expectedWithdrawable = _calcWithdrawableShares(
                    uint(beaconShares),
                    uint(WAD).divWad(initialMagnitude),
                    _getSlashingFactor(defaultStaker, beaconChainETHStrategy, newMaxMagnitude)
                );

                _assertSharesAfterSlash({
                    staker: defaultStaker,
                    strategy: beaconChainETHStrategy,
                    withdrawableSharesBefore: beaconSharesBeforeSlash,
                    expectedWithdrawableShares: expectedWithdrawable,
                    prevMaxMagnitude: initialMagnitude,
                    currMaxMagnitude: newMaxMagnitude
                });

                sharesAfterAllSlashing = beaconSharesAfterSecondSlash;
            }
        }

        ////////////////////////////
        // 2. do AVS slash then beacon chain slash
        ////////////////////////////

        // initialize new staker and operator with same initial conditions
        delegationManager.undelegate(defaultStaker);
        _registerOperatorWithBaseDetails(defaultOperator2);
        _setOperatorMagnitude(defaultOperator2, beaconChainETHStrategy, initialMagnitude);
        eigenPodManagerMock.setPodOwnerShares(defaultStaker2, beaconShares);
        eigenPodManagerMock.setBeaconChainSlashingFactor(defaultStaker2, WAD);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker2, defaultOperator2);
        _assertDeposit({
            staker: defaultStaker2,
            operator: defaultOperator2,
            strategy: beaconChainETHStrategy,
            operatorSharesBefore: 0,
            withdrawableSharesBefore: 0,
            depositSharesBefore: 0,
            prevDsf: WAD,
            depositAmount: uint(beaconShares)
        });

        {
            // Slash on EigenLayer first
            {
                (withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker2, beaconChainETHStrategy.toArray());
                uint beaconSharesBeforeSlash = withdrawableShares[0];

                _setOperatorMagnitude(defaultOperator2, beaconChainETHStrategy, newMaxMagnitude);
                cheats.prank(address(allocationManagerMock));
                delegationManager.slashOperatorShares(defaultOperator2, beaconChainETHStrategy, initialMagnitude, newMaxMagnitude);

                uint expectedWithdrawable = _calcWithdrawableShares(
                    uint(beaconShares),
                    uint(WAD).divWad(initialMagnitude),
                    _getSlashingFactor(defaultStaker2, beaconChainETHStrategy, newMaxMagnitude)
                );

                _assertSharesAfterSlash({
                    staker: defaultStaker2,
                    strategy: beaconChainETHStrategy,
                    withdrawableSharesBefore: beaconSharesBeforeSlash,
                    expectedWithdrawableShares: expectedWithdrawable,
                    prevMaxMagnitude: initialMagnitude,
                    currMaxMagnitude: newMaxMagnitude
                });
            }

            // Slash beaconChain second
            {
                (withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker2, beaconChainETHStrategy.toArray());
                uint beaconSharesBeforeSlash = withdrawableShares[0];

                uint64 prevBeaconChainSlashingFactor;
                (prevBeaconChainSlashingFactor, newBeaconSlashingFactor) =
                    _decreaseBeaconChainShares(defaultStaker2, beaconShares, sharesDecrease);

                uint expectedWithdrawableShares = _calcWithdrawableShares({
                    depositShares: uint(beaconShares),
                    depositScalingFactor: uint(WAD).divWad(initialMagnitude),
                    slashingFactor: newMaxMagnitude.mulWad(newBeaconSlashingFactor)
                });
                _assertSharesAfterBeaconSlash({
                    staker: defaultStaker2,
                    withdrawableSharesBefore: beaconSharesBeforeSlash,
                    expectedWithdrawableShares: expectedWithdrawableShares,
                    prevBeaconSlashingFactor: prevBeaconChainSlashingFactor
                });
            }
        }

        ////////////////////////////
        // 3. Confirm withdrawable shares are the same regardless of order of operations in Test 1 or Test 2
        ////////////////////////////
        (withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker2, beaconChainETHStrategy.toArray());
        assertEq(withdrawableShares[0], sharesAfterAllSlashing, "shares after all slashing should be the same");
    }
}

/// @notice Fuzzed Unit tests to compare totalWitdrawable shares for an operator vs their actual operatorShares.
/// Requires the WRITE_CSV_TESTS env variable to be set to true to output to a test file
contract DelegationManagerUnitTests_SharesUnderflowChecks is DelegationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;

    /**
     * @notice Fuzzed tests
     * Single staker with fuzzed starting shares and magnitude.
     * Slash 100 magnitude and deposit 100 shares for 100 iterations.
     */
    /// forge-config: default.fuzz.runs = 50
    function testFuzz_slashDepositRepeatedly(Randomness r) public rand(r) {
        uint64 initMagnitude = r.Uint64(10_000, WAD);
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        cheats.assume(initMagnitude % 2 != 0);
        cheats.assume(shares % 2 != 0);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initMagnitude);

        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint[] memory sharesToSet = new uint[](1);
            sharesToSet[0] = shares;

            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        }

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Slash and deposit more for each iteration
        uint64 currMagnitude = initMagnitude;
        {
            uint newDepositShares = shares;
            for (uint i = 0; i < 100; ++i) {
                // 1. slash operator for 100 magnitude
                uint64 slashMagnitude = 100;
                currMagnitude -= slashMagnitude;
                _setOperatorMagnitude(defaultOperator, strategyMock, currMagnitude);
                cheats.prank(address(allocationManagerMock));
                delegationManager.slashOperatorShares({
                    operator: defaultOperator,
                    strategy: strategyMock,
                    prevMaxMagnitude: currMagnitude + slashMagnitude,
                    newMaxMagnitude: currMagnitude
                });

                // 2. deposit again
                uint sharesIncrease = 1000;
                cheats.prank(address(strategyManagerMock));
                delegationManager.increaseDelegatedShares(defaultStaker, strategyMock, newDepositShares, sharesIncrease);
                newDepositShares += sharesIncrease;

                uint[] memory newDepositSharesArray = new uint[](1);
                newDepositSharesArray[0] = newDepositShares;

                strategyManagerMock.setDeposits(defaultStaker, strategies, newDepositSharesArray);
            }
        }

        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertLe(
            withdrawableShares[0],
            delegationManager.operatorShares(defaultOperator, strategyMock),
            "withdrawableShares should be less than or equal to operatorShares"
        );

        if (cheats.envOr("WRITE_CSV_TESTS", false)) {
            cheats.writeLine(
                "./test.csv",
                string(
                    abi.encodePacked(
                        cheats.toString(initMagnitude),
                        ", ",
                        cheats.toString(shares),
                        ", ",
                        cheats.toString(delegationManager.operatorShares(defaultOperator, strategyMock)),
                        ", ",
                        cheats.toString(withdrawableShares[0]),
                        ", ",
                        cheats.toString(
                            stdMath.delta(delegationManager.operatorShares(defaultOperator, strategyMock), withdrawableShares[0])
                        )
                    )
                )
            );
        }
    }

    /**
     * @notice Fuzzed tests
     * Single staker with fuzzed starting shares and magnitude.
     * Slash 100 magnitude and fuzz deposit amount for 100 iterations.
     */
    /// forge-config: default.fuzz.runs = 50
    function testFuzz_slashDepositRepeatedly_randDeposits(Randomness r) public rand(r) {
        uint64 initMagnitude = r.Uint64(10_000, WAD);
        uint depositAmount = r.Uint256(1, 1e34);
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES / 1e4);
        cheats.assume(initMagnitude % 2 != 0);
        cheats.assume(shares % 2 != 0);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initMagnitude);

        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint[] memory sharesToSet = new uint[](1);
            sharesToSet[0] = shares;

            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        }

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Slash and deposit more for each iteration
        uint64 currMagnitude = initMagnitude;
        {
            uint newDepositShares = shares;
            for (uint i = 0; i < 100; ++i) {
                // 1. slash operator for 100 magnitude
                uint64 slashMagnitude = 100;
                currMagnitude -= slashMagnitude;
                _setOperatorMagnitude(defaultOperator, strategyMock, currMagnitude);
                cheats.prank(address(allocationManagerMock));
                delegationManager.slashOperatorShares({
                    operator: defaultOperator,
                    strategy: strategyMock,
                    prevMaxMagnitude: currMagnitude + slashMagnitude,
                    newMaxMagnitude: currMagnitude
                });

                // 2. deposit again
                cheats.prank(address(strategyManagerMock));
                delegationManager.increaseDelegatedShares(defaultStaker, strategyMock, newDepositShares, depositAmount);
                newDepositShares += depositAmount;

                uint[] memory newDepositSharesArray = new uint[](1);
                newDepositSharesArray[0] = newDepositShares;

                strategyManagerMock.setDeposits(defaultStaker, strategies, newDepositSharesArray);
            }
        }

        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertLe(
            withdrawableShares[0],
            delegationManager.operatorShares(defaultOperator, strategyMock),
            "withdrawableShares should be less than or equal to operatorShares"
        );

        if (cheats.envOr("WRITE_CSV_TESTS", false)) {
            cheats.writeLine(
                "./test2.csv",
                string(
                    abi.encodePacked(
                        cheats.toString(initMagnitude),
                        ", ",
                        cheats.toString(shares),
                        ", ",
                        cheats.toString(depositAmount),
                        ", ",
                        cheats.toString(delegationManager.operatorShares(defaultOperator, strategyMock)),
                        ", ",
                        cheats.toString(withdrawableShares[0]),
                        ", ",
                        cheats.toString(
                            stdMath.delta(delegationManager.operatorShares(defaultOperator, strategyMock), withdrawableShares[0])
                        )
                    )
                )
            );
        }
    }

    /**
     * @notice Fuzzed tests
     * For 500 stakers, deposit `shares` amount and delegate to the operator. After each staker delegates,
     * slash 100 magnitude.
     */
    /// forge-config: default.fuzz.runs = 50
    function testFuzz_depositMultipleStakers_slash_repeatedly(Randomness r) public rand(r) {
        uint64 initMagnitude = r.Uint64(50_000, WAD);
        uint shares = r.Uint256(MAX_STRATEGY_SHARES / 1e7, MAX_STRATEGY_SHARES / 1e4);
        cheats.assume(initMagnitude % 2 != 0);
        cheats.assume(shares % 2 != 0);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initMagnitude);

        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        uint[] memory sharesToSet = new uint[](1);
        sharesToSet[0] = shares;

        uint numStakers = 500;

        address[] memory stakers = new address[](numStakers);
        // Slash and deposit more for each iteration
        uint64 currMagnitude = initMagnitude;
        {
            for (uint i = 0; i < numStakers; ++i) {
                // 1. deposit and delegate new staker
                stakers[i] = random().Address();
                strategyManagerMock.setDeposits(stakers[i], strategies, sharesToSet);
                _delegateToOperatorWhoAcceptsAllStakers(stakers[i], defaultOperator);

                // 2. slash operator for 100 magnitude
                uint64 slashMagnitude = 100;
                currMagnitude -= slashMagnitude;
                _setOperatorMagnitude(defaultOperator, strategyMock, currMagnitude);
                cheats.prank(address(allocationManagerMock));
                delegationManager.slashOperatorShares({
                    operator: defaultOperator,
                    strategy: strategyMock,
                    prevMaxMagnitude: currMagnitude + slashMagnitude,
                    newMaxMagnitude: currMagnitude
                });
            }
        }

        uint operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint totalWithdrawableShares = 0;
        for (uint i = 0; i < numStakers; ++i) {
            (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(stakers[i], strategies);
            totalWithdrawableShares += withdrawableShares[0];
        }
        assertLe(totalWithdrawableShares, operatorSharesAfter, "withdrawableShares should be less than or equal to operatorShares");

        if (cheats.envOr("WRITE_CSV_TESTS", false)) {
            cheats.writeLine(
                "./test3.csv",
                string(
                    abi.encodePacked(
                        cheats.toString(initMagnitude),
                        ", ", // initial magnitude
                        cheats.toString(shares),
                        ", ", // amount each staker deposits
                        cheats.toString(operatorSharesAfter),
                        ", ", // operator shares after all slashing and deposits
                        cheats.toString(totalWithdrawableShares),
                        ", ", // total withdrawable shares from all stakers
                        cheats.toString(stdMath.delta(operatorSharesAfter, totalWithdrawableShares)) // delta difference between opShares and total withdrawable
                    )
                )
            );
        }
    }

    /**
     * @notice Fuzzed tests
     * For 500 stakers, deposit `shares` amount and delegate to the operator. After each staker delegates,
     * slash 1000 magnitude. Initial magnitude is very small so this will slash larger proportions.
     */
    /// forge-config: default.fuzz.runs = 50
    function testFuzz_depositMultipleStakers_slashLargeMagnitudes(Randomness r) public rand(r) {
        uint64 initMagnitude = r.Uint64(50_000, WAD);
        uint shares = r.Uint256(MAX_STRATEGY_SHARES / 1e7, MAX_STRATEGY_SHARES / 1e4);
        cheats.assume(initMagnitude % 2 != 0);
        cheats.assume(shares % 2 != 0);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initMagnitude);

        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        uint[] memory sharesToSet = new uint[](1);
        sharesToSet[0] = shares;

        uint numStakers = 500;

        address[] memory stakers = new address[](numStakers);
        // Slash and deposit more for each iteration
        uint64 currMagnitude = initMagnitude;
        {
            for (uint i = 0; i < numStakers; ++i) {
                // 1. deposit and delegate new staker
                stakers[i] = random().Address();
                strategyManagerMock.setDeposits(stakers[i], strategies, sharesToSet);
                _delegateToOperatorWhoAcceptsAllStakers(stakers[i], defaultOperator);

                // 2. slash operator for 100 magnitude
                uint64 slashMagnitude = 100;
                currMagnitude -= slashMagnitude;
                _setOperatorMagnitude(defaultOperator, strategyMock, currMagnitude);
                cheats.prank(address(allocationManagerMock));
                delegationManager.slashOperatorShares({
                    operator: defaultOperator,
                    strategy: strategyMock,
                    prevMaxMagnitude: currMagnitude + slashMagnitude,
                    newMaxMagnitude: currMagnitude
                });
            }
        }

        uint operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint totalWithdrawableShares = 0;
        for (uint i = 0; i < numStakers; ++i) {
            (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(stakers[i], strategies);
            totalWithdrawableShares += withdrawableShares[0];
        }
        assertLe(totalWithdrawableShares, operatorSharesAfter, "withdrawableShares should be less than or equal to operatorShares");

        if (cheats.envOr("WRITE_CSV_TESTS", false)) {
            cheats.writeLine(
                "./test4.csv",
                string(
                    abi.encodePacked(
                        cheats.toString(initMagnitude),
                        ", ", // initial magnitude
                        cheats.toString(shares),
                        ", ", // amount each staker deposits
                        cheats.toString(operatorSharesAfter),
                        ", ", // operator shares after all slashing and deposits
                        cheats.toString(totalWithdrawableShares),
                        ", ", // total withdrawable shares from all stakers
                        cheats.toString(stdMath.delta(operatorSharesAfter, totalWithdrawableShares)) // delta difference between opShares and total withdrawable
                    )
                )
            );
        }
    }

    /**
     * @notice Same as above `testFuzz_depositMultipleStakers_slashLargeMagnitudes` test but with slashing
     * 1 magnitude instead of 100.
     */
    /// forge-config: default.fuzz.runs = 50
    function testFuzz_depositMultipleStakers_slashSmallMagnitudes(Randomness r) public rand(r) {
        uint64 initMagnitude = r.Uint64(1000, WAD);
        uint shares = r.Uint256(MAX_STRATEGY_SHARES / 1e7, MAX_STRATEGY_SHARES / 1e4);
        cheats.assume(initMagnitude % 2 != 0);
        cheats.assume(shares % 2 != 0);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initMagnitude);

        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        uint[] memory sharesToSet = new uint[](1);
        sharesToSet[0] = shares;

        uint numStakers = 500;

        address[] memory stakers = new address[](numStakers);
        // Slash and deposit more for each iteration
        uint64 currMagnitude = initMagnitude;
        {
            for (uint i = 0; i < numStakers; ++i) {
                // 1. deposit and delegate new staker
                stakers[i] = random().Address();
                strategyManagerMock.setDeposits(stakers[i], strategies, sharesToSet);
                _delegateToOperatorWhoAcceptsAllStakers(stakers[i], defaultOperator);

                // 2. slash operator for 100 magnitude
                uint64 slashMagnitude = 1;
                currMagnitude -= slashMagnitude;
                _setOperatorMagnitude(defaultOperator, strategyMock, currMagnitude);
                cheats.prank(address(allocationManagerMock));
                delegationManager.slashOperatorShares({
                    operator: defaultOperator,
                    strategy: strategyMock,
                    prevMaxMagnitude: currMagnitude + slashMagnitude,
                    newMaxMagnitude: currMagnitude
                });
            }
        }

        uint operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint totalWithdrawableShares = 0;
        for (uint i = 0; i < numStakers; ++i) {
            (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(stakers[i], strategies);
            totalWithdrawableShares += withdrawableShares[0];
        }
        assertLe(totalWithdrawableShares, operatorSharesAfter, "withdrawableShares should be less than or equal to operatorShares");

        if (cheats.envOr("WRITE_CSV_TESTS", false)) {
            cheats.writeLine(
                "./test5.csv",
                string(
                    abi.encodePacked(
                        cheats.toString(initMagnitude),
                        ", ", // initial magnitude
                        cheats.toString(shares),
                        ", ", // amount each staker deposits
                        cheats.toString(operatorSharesAfter),
                        ", ", // operator shares after all slashing and deposits
                        cheats.toString(totalWithdrawableShares),
                        ", ", // total withdrawable shares from all stakers
                        cheats.toString(stdMath.delta(operatorSharesAfter, totalWithdrawableShares)) // delta difference between opShares and total withdrawable
                    )
                )
            );
        }
    }

    /**
     * @notice Setup 500 delegated stakers who each deposit `shares` amount.
     * Then slash 1 magnitude 500 times and then compare amount of shares that can be withdrawn vs operatorShares
     */
    /// forge-config: default.fuzz.runs = 50
    function testFuzz_depositMultipleStakersOnce_slashSmallMagnitudes(Randomness r) public rand(r) {
        uint64 initMagnitude = r.Uint64(1000, WAD);
        uint shares = r.Uint256(MAX_STRATEGY_SHARES / 1e7, MAX_STRATEGY_SHARES / 1e4);
        cheats.assume(initMagnitude % 2 != 0);
        cheats.assume(shares % 2 != 0);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initMagnitude);

        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        uint[] memory sharesToSet = new uint[](1);
        sharesToSet[0] = shares;

        uint numStakers = 500;

        address[] memory stakers = new address[](numStakers);
        // deposit all stakers one time
        for (uint i = 0; i < numStakers; ++i) {
            // 1. deposit and delegate new staker
            stakers[i] = random().Address();
            strategyManagerMock.setDeposits(stakers[i], strategies, sharesToSet);
            _delegateToOperatorWhoAcceptsAllStakers(stakers[i], defaultOperator);
        }

        // Slash and deposit more for each iteration
        uint64 currMagnitude = initMagnitude;
        {
            for (uint i = 0; i < numStakers; ++i) {
                // 2. slash operator for 100 magnitude
                uint64 slashMagnitude = 1;
                currMagnitude -= slashMagnitude;
                _setOperatorMagnitude(defaultOperator, strategyMock, currMagnitude);
                cheats.prank(address(allocationManagerMock));
                delegationManager.slashOperatorShares({
                    operator: defaultOperator,
                    strategy: strategyMock,
                    prevMaxMagnitude: currMagnitude + slashMagnitude,
                    newMaxMagnitude: currMagnitude
                });
            }
        }

        uint operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint totalWithdrawableShares = 0;
        for (uint i = 0; i < numStakers; ++i) {
            (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(stakers[i], strategies);
            totalWithdrawableShares += withdrawableShares[0];
        }
        assertLe(totalWithdrawableShares, operatorSharesAfter, "withdrawableShares should be less than or equal to operatorShares");

        if (cheats.envOr("WRITE_CSV_TESTS", false)) {
            cheats.writeLine(
                "./test6.csv",
                string(
                    abi.encodePacked(
                        cheats.toString(initMagnitude),
                        ", ", // initial magnitude
                        cheats.toString(shares),
                        ", ", // amount each staker deposits
                        cheats.toString(operatorSharesAfter),
                        ", ", // operator shares after all slashing and deposits
                        cheats.toString(totalWithdrawableShares),
                        ", ", // total withdrawable shares from all stakers
                        cheats.toString(stdMath.delta(operatorSharesAfter, totalWithdrawableShares)) // delta difference between opShares and total withdrawable
                    )
                )
            );
        }
    }
}

contract DelegationManagerUnitTests_Rounding is DelegationManagerUnitTests {}

/**
 * @notice TODO Lifecycle tests - These tests combine multiple functionalities of the DelegationManager
 *    1. Old SigP test - registerAsOperator, separate staker delegate to operator, as operator undelegate (reverts),
 *     checks that staker is still delegated and operator still registered, staker undelegates, checks staker not delegated and operator
 *     is still registered
 *    2. RegisterOperator, Deposit, Delegate, Queue, Complete
 *    3. RegisterOperator, Mock Slash(set maxMagnitudes), Deposit/Delegate, Queue, Complete
 *    4. RegisterOperator, Deposit/Delegate, Mock Slash(set maxMagnitudes), Queue, Complete
 *    5. RegisterOperator, Mock Slash(set maxMagnitudes), Deposit/Delegate, Queue, Mock Slash(set maxMagnitudes), Complete
 *    7. RegisterOperator, Deposit/Delegate, Mock Slash 100% (set maxMagnitudes), Undelegate, Complete non 100% slashed strategies
 *    8. RegisterOperator, Deposit/Delegate, Undelegate, Re delegate to another operator, Mock Slash 100% (set maxMagnitudes), Complete as shares
 *     (withdrawals should have been slashed even though delegated to a new operator)
 *    9. Invariant check getWithdrawableShares = sum(deposits), Multiple deposits with operator who has never been slashed
 *    10. Invariant check getWithdrawableShares = sum(deposits), Multiple deposits with operator who HAS been been slashed
 */
contract DelegationManagerUnitTests_Lifecycle is DelegationManagerUnitTests {
    using ArrayLib for *;

    // 2. RegisterOperator, Deposit, Delegate, Queue, Complete
    function test_register_operator_deposit_delegate_queue_complete(Randomness r) public rand(r) {
        address operator = r.Address();
        address staker = r.Address();
        IStrategy[] memory strategies = strategyMock.toArray();
        uint[] memory depositShares = uint(100 ether).toArrayU256();

        // 1) Register operator.
        _registerOperatorWithBaseDetails(operator);

        // 2) Mock deposit into SM.
        strategyManagerMock.setDeposits(staker, strategies, depositShares);

        // 3) Staker delegates to operator.
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // 3) Staker queues withdrawals.
        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] =
            QueuedWithdrawalParams({strategies: strategies, depositShares: depositShares, __deprecated_withdrawer: address(0)});
        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // 4) Complete queued withdrawals.
        Withdrawal memory withdrawal = Withdrawal({
            staker: staker,
            delegatedTo: operator,
            withdrawer: staker,
            nonce: 0,
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: depositShares
        });

        bytes32 withdrawalRoot = keccak256(abi.encode(withdrawal));
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
        cheats.roll(block.number + delegationManager.minWithdrawalDelayBlocks() + 1);

        cheats.prank(staker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokenMock.toArray(), false);

        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should not be pending");

        // Checks
        assertEq(delegationManager.cumulativeWithdrawalsQueued(staker), 1, "staker nonce should have incremented");
        assertEq(delegationManager.operatorShares(operator, strategies[0]), 100 ether, "operator shares should be 0 after withdrawal");
    }

    /**
     * @notice While delegated to an operator who becomes 100% slashed. When the staker undelegates and queues a withdrawal
     * for all their shares which are now 0, the withdrawal should be completed with 0 shares even if they delegate to a new operator
     * who has not been slashed.
     * Note: This specifically tests that the completeQueuedWithdrawal is looking up the correct maxMagnitude for the operator
     */
    function testFuzz_undelegate_slashOperator100Percent_delegate_complete(Randomness r) public rand(r) {
        uint shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        address newOperator = r.Address();
        IStrategy[] memory strategyArray = r.StrategyArray(1);
        IStrategy strategy = strategyArray[0];

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategy, shares);

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Set operator magnitude
        uint64 operatorMagnitude = 0;
        uint operatorSharesAfterSlash;
        {
            _setOperatorMagnitude(defaultOperator, strategy, operatorMagnitude);
            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, strategy, WAD, 0);
            operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategy);
            assertEq(operatorSharesAfterSlash, 0, "operator shares not fully slashed");
        }

        (, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategy, depositSharesToWithdraw: shares});

        uint depositScalingFactor = delegationManager.depositScalingFactor(defaultStaker, strategy);
        assertEq(depositScalingFactor, WAD, "bad test setup");
        // Get withdrawable and deposit shares
        {
            (uint[] memory withdrawableSharesBefore, uint[] memory depositSharesBefore) =
                delegationManager.getWithdrawableShares(defaultStaker, strategyArray);
            assertEq(withdrawableSharesBefore[0], 0, "withdrawable shares should be 0 after being slashed fully");
            assertEq(depositSharesBefore[0], shares, "deposit shares should be unchanged after being slashed fully");
        }

        // Undelegate the staker
        _undelegate_expectEmit_singleStrat(
            UndelegateEmitStruct({
                staker: defaultStaker,
                operator: defaultOperator,
                strategy: strategy,
                depositSharesQueued: shares,
                operatorSharesDecreased: 0,
                withdrawal: withdrawal,
                withdrawalRoot: withdrawalRoot,
                depositScalingFactor: WAD,
                forceUndelegated: false
            })
        );
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(delegationManager.delegatedTo(defaultStaker), address(0), "undelegated staker should be delegated to zero address");
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        assertEq(delegationManager.operatorShares(defaultOperator, strategy), 0, "operator shares not decreased correctly");
        (uint[] memory stakerWithdrawableShares, uint[] memory depositShares) =
            delegationManager.getWithdrawableShares(defaultStaker, strategyArray);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
        assertEq(depositShares[0], 0, "staker deposit shares not reset correctly");

        // delegate to the `newOperator` who has never been slashed
        // Ensure that completing withdrawal now still results in 0 shares
        _registerOperatorWithBaseDetails(newOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, newOperator);

        (stakerWithdrawableShares, depositShares) = delegationManager.getWithdrawableShares(defaultStaker, strategyArray);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
        assertEq(depositShares[0], 0, "staker deposit shares not reset correctly");

        cheats.roll(withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokenMock.toArray(), false);

        (stakerWithdrawableShares, depositShares) = delegationManager.getWithdrawableShares(defaultStaker, strategyArray);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
        assertEq(depositShares[0], 0, "staker deposit shares not reset correctly");
        assertEq(delegationManager.operatorShares(newOperator, strategy), 0, "new operator shares should be unchanged");
    }
}

contract DelegationManagerUnitTests_ConvertToDepositShares is DelegationManagerUnitTests {
    using ArrayLib for *;

    function test_convertToDepositShares_noSlashing() public {
        uint shares = 100 ether;

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, shares);
        _checkDepositSharesConvertCorrectly(strategyMock.toArray(), shares.toArrayU256());
    }

    function test_convertToDepositShares_withSlashing() public {
        IStrategy[] memory strategies = strategyMock.toArray();
        uint[] memory shares = uint(100 ether).toArrayU256();

        // Set the staker deposits in the strategies
        strategyManagerMock.addDeposit(defaultStaker, strategies[0], shares[0]);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, WAD / 3);

        _checkDepositSharesConvertCorrectly(strategies, shares);

        // queue and complete a withdrawal for half the deposit shares
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        _queueAndCompleteWithdrawalForSingleStrategy(strategies[0], shares[0] / 2);

        // queued a withdrawal for half the deposit shares, and added back as withdrawable shares
        shares[0] = shares[0] / 2 + withdrawableShares[0] / 2;
        _checkDepositSharesConvertCorrectly(strategies, shares);
    }

    function test_convertToDepositShares_beaconChainETH() public {
        IStrategy[] memory strategies = beaconChainETHStrategy.toArray();
        uint[] memory shares = uint(100 ether).toArrayU256();

        // Set the staker deposits in the strategies
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, int(shares[0]));

        uint[] memory depositShares = delegationManager.convertToDepositShares(defaultStaker, strategies, shares);
        assertEq(depositShares[0], shares[0], "deposit shares not converted correctly");

        // delegate to an operator and slash
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, WAD / 3);

        _checkDepositSharesConvertCorrectly(strategies, shares);

        // slash on beacon chain by 1/3
        _decreaseBeaconChainShares(defaultStaker, int(shares[0]), shares[0] / 3);

        _checkDepositSharesConvertCorrectly(strategies, shares);

        // queue and complete a withdrawal for half the deposit shares
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        _queueAndCompleteWithdrawalForSingleStrategy(strategies[0], shares[0] / 2);

        // queued a withdrawal for half the deposit shares, and added back as withdrawable shares
        shares[0] = shares[0] / 2 + withdrawableShares[0] / 2;
        _checkDepositSharesConvertCorrectly(strategies, shares);
    }

    function _checkDepositSharesConvertCorrectly(IStrategy[] memory strategies, uint[] memory expectedDepositShares) public view {
        (uint[] memory withdrawableShares,) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        // get the deposit shares
        uint[] memory depositShares = delegationManager.convertToDepositShares(defaultStaker, strategies, withdrawableShares);

        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqRel(expectedDepositShares[i], depositShares[i], APPROX_REL_DIFF, "deposit shares not converted correctly");

            // make sure that the deposit shares are less than or equal to the shares,
            // so this value is sane to input into `completeQueuedWithdrawals`
            assertLe(depositShares[i], expectedDepositShares[i], "deposit shares should be less than or equal to expected deposit shares");
        }

        // get the deposit shares
        uint[] memory oneThirdWithdrawableShares = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            oneThirdWithdrawableShares[i] = withdrawableShares[i] / 3;
        }
        uint[] memory oneThirdDepositShares =
            delegationManager.convertToDepositShares(defaultStaker, strategies, oneThirdWithdrawableShares);
        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqRel(
                expectedDepositShares[i] / 3, oneThirdDepositShares[i], APPROX_REL_DIFF, "deposit shares not converted correctly"
            );
        }
    }

    function _queueAndCompleteWithdrawalForSingleStrategy(IStrategy strategy, uint shares) public {
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal,) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategy, depositSharesToWithdraw: shares});

        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        cheats.roll(block.number + delegationManager.minWithdrawalDelayBlocks() + 1);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokenMock.toArray(), false);
    }
}

contract DelegationManagerUnitTests_getQueuedWithdrawals is DelegationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;

    function _withdrawalRoot(Withdrawal memory withdrawal) internal pure returns (bytes32) {
        return keccak256(abi.encode(withdrawal));
    }

    function test_getQueuedWithdrawals_Correctness(Randomness r) public rand(r) {
        uint numStrategies = r.Uint256(2, 8);
        uint[] memory depositShares = r.Uint256Array({len: numStrategies, min: 2, max: 100 ether});

        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(defaultStaker, depositShares, false);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        for (uint i; i < numStrategies; ++i) {
            uint newStakerShares = depositShares[i] / 2;
            _setOperatorMagnitude(defaultOperator, strategies[i], 0.5 ether);
            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, strategies[i], WAD, 0.5 ether);
            uint afterSlash = delegationManager.operatorShares(defaultOperator, strategies[i]);
            assertApproxEqAbs(afterSlash, newStakerShares, 1, "bad operator shares after slash");
        }

        // Queue withdrawals.
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawals({staker: defaultStaker, strategies: strategies, depositWithdrawalAmounts: depositShares});

        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Get queued withdrawals.
        (Withdrawal[] memory withdrawals, uint[][] memory shares) = delegationManager.getQueuedWithdrawals(defaultStaker);
        // Checks
        for (uint i; i < strategies.length; ++i) {
            uint newStakerShares = depositShares[i] / 2;
            assertApproxEqAbs(shares[0][i], newStakerShares, 1, "staker shares should be decreased by half +- 1");
        }

        assertEq(
            _withdrawalRoot(withdrawal), _withdrawalRoot(withdrawals[0]), "_withdrawalRoot(withdrawal) != _withdrawalRoot(withdrawals[0])"
        );
        assertEq(_withdrawalRoot(withdrawal), withdrawalRoot, "_withdrawalRoot(withdrawal) != withdrawalRoot");
    }

    function test_getQueuedWithdrawals_TotalQueuedGreaterThanTotalStrategies(Randomness r) public rand(r) {
        uint totalDepositShares = r.Uint256(2, 100 ether);

        _registerOperatorWithBaseDetails(defaultOperator);
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, totalDepositShares);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        uint newStakerShares = totalDepositShares / 2;
        _setOperatorMagnitude(defaultOperator, strategyMock, 0.5 ether);
        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, WAD, 0.5 ether);
        uint afterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
        assertApproxEqAbs(afterSlash, newStakerShares, 1, "bad operator shares after slash");

        // Queue withdrawals.
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams0, Withdrawal memory withdrawal0, bytes32 withdrawalRoot0) =
        _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: totalDepositShares / 2});

        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams0);

        (QueuedWithdrawalParams[] memory queuedWithdrawalParams1, Withdrawal memory withdrawal1, bytes32 withdrawalRoot1) =
        _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: totalDepositShares / 2});

        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams1);

        // Get queued withdrawals.
        (Withdrawal[] memory withdrawals, uint[][] memory shares) = delegationManager.getQueuedWithdrawals(defaultStaker);

        // Sanity
        assertEq(withdrawals.length, 2, "withdrawal.length != 2");
        assertEq(withdrawals[0].strategies.length, 1, "withdrawals[0].strategies.length != 1");
        assertEq(withdrawals[1].strategies.length, 1, "withdrawals[1].strategies.length != 1");

        // Checks
        assertApproxEqAbs(shares[0][0], newStakerShares / 2, 1, "shares[0][0] != newStakerShares");
        assertApproxEqAbs(shares[1][0], newStakerShares / 2, 1, "shares[1][0] != newStakerShares");
        assertEq(_withdrawalRoot(withdrawal0), _withdrawalRoot(withdrawals[0]), "withdrawal0 != withdrawals[0]");
        assertEq(_withdrawalRoot(withdrawal1), _withdrawalRoot(withdrawals[1]), "withdrawal1 != withdrawals[1]");
        assertEq(_withdrawalRoot(withdrawal0), withdrawalRoot0, "_withdrawalRoot(withdrawal0) != withdrawalRoot0");
        assertEq(_withdrawalRoot(withdrawal1), withdrawalRoot1, "_withdrawalRoot(withdrawal1) != withdrawalRoot1");
    }

    /**
     * @notice Assert that the shares returned in the view function `getQueuedWithdrawals` are unaffected from a
     * slash that occurs after the withdrawal is completed. Also assert that completing the withdrawal matches the
     * expected withdrawn shares from the view function.
     * Slashing on the completableBlock of the withdrawal should have no affect on the withdrawn shares.
     */
    function test_getQueuedWithdrawals_SlashAfterWithdrawalCompletion(Randomness r) public rand(r) {
        uint depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);

        // Deposit Staker
        strategyManagerMock.addDeposit(defaultStaker, strategyMock, depositAmount);

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawal
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, Withdrawal memory withdrawal, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: depositAmount});
        {
            uint operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, strategyMock);
            cheats.prank(defaultStaker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);

            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
            uint operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, strategyMock);
            uint sharesWithdrawn =
                _calcWithdrawableShares({depositShares: depositAmount, depositScalingFactor: uint(WAD), slashingFactor: uint(WAD)});
            assertEq(
                operatorSharesAfterQueue, operatorSharesBeforeQueue - sharesWithdrawn, "operator shares should be decreased after queue"
            );
        }

        // Slash operator 50% while staker has queued withdrawal
        {
            uint operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, strategyMock);
            (uint sharesToDecrement,) =
                _calcSlashedAmount({operatorShares: operatorSharesAfterQueue, prevMaxMagnitude: uint64(WAD), newMaxMagnitude: 50e16});
            _setOperatorMagnitude(defaultOperator, strategyMock, 50e16);
            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, withdrawal.strategies[0], uint64(WAD), 50e16);
            uint operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
            assertEq(
                operatorSharesAfterSlash, operatorSharesAfterQueue - sharesToDecrement, "operator shares should be decreased after slash"
            );
        }

        // Assert that the getQueuedWithdrawals returns shares that are halved as a result of being slashed 50%
        {
            (Withdrawal[] memory withdrawals, uint[][] memory shares) = delegationManager.getQueuedWithdrawals(defaultStaker);
            assertEq(withdrawals.length, 1, "withdrawals.length != 1");
            assertEq(withdrawals[0].strategies.length, 1, "withdrawals[0].strategies.length != 1");
            assertEq(shares[0][0], depositAmount / 2, "shares[0][0] != depositAmount / 2");
        }

        // Roll blocks to after withdrawal completion
        uint32 completableBlock = withdrawal.startBlock + delegationManager.minWithdrawalDelayBlocks() + 1;
        cheats.roll(completableBlock);

        // slash operator 50% again
        {
            uint operatorShares = delegationManager.operatorShares(defaultOperator, strategyMock);
            (uint sharesToDecrement,) =
                _calcSlashedAmount({operatorShares: operatorShares, prevMaxMagnitude: 50e16, newMaxMagnitude: 25e16});
            _setOperatorMagnitude(defaultOperator, strategyMock, 25e16);
            cheats.prank(address(allocationManagerMock));
            delegationManager.slashOperatorShares(defaultOperator, withdrawal.strategies[0], 50e16, 25e16);
            uint operatorSharesAfterSecondSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
            assertEq(operatorSharesAfterSecondSlash, operatorShares - sharesToDecrement, "operator shares should be decreased after slash");
        }

        // Assert that the getQueuedWithdrawals returns shares that are halved as a result of being slashed 50% and hasn't been
        // affected by the second slash
        uint expectedSharesIncrease = depositAmount / 2;
        uint queuedWithdrawableShares;
        {
            (Withdrawal[] memory withdrawals, uint[][] memory shares) = delegationManager.getQueuedWithdrawals(defaultStaker);
            queuedWithdrawableShares = shares[0][0];
            assertEq(withdrawals.length, 1, "withdrawals.length != 1");
            assertEq(withdrawals[0].strategies.length, 1, "withdrawals[0].strategies.length != 1");
            assertEq(queuedWithdrawableShares, depositAmount / 2, "queuedWithdrawableShares != withdrawalAmount / 2");
        }

        // Complete queued Withdrawal with shares added back. Since total deposit slashed by 50% and not 75%
        (uint[] memory withdrawableSharesBefore,) = delegationManager.getWithdrawableShares(defaultStaker, withdrawal.strategies);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokenMock.toArray(), false);
        (uint[] memory withdrawableSharesAfter,) = delegationManager.getWithdrawableShares(defaultStaker, withdrawal.strategies);

        // Added shares
        assertEq(
            withdrawableSharesAfter[0],
            withdrawableSharesBefore[0] + expectedSharesIncrease,
            "withdrawableShares should be increased by expectedSharesIncrease"
        );
        assertEq(expectedSharesIncrease, queuedWithdrawableShares, "expectedSharesIncrease should be equal to queuedWithdrawableShares");
        assertEq(block.number, completableBlock, "block.number should be the completableBlock");
    }

    function test_getQueuedWithdrawals_UsesCorrectOperatorMagnitude() public {
        // Alice deposits 100 shares into strategy
        uint depositAmount = 100e18;
        _depositIntoStrategies(defaultStaker, strategyMock.toArray(), depositAmount.toArrayU256());

        // Register operator with magnitude of 0.5 and delegate Alice to them
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, 0.5 ether);

        // Alice queues withdrawal of all 100 shares while operator magnitude is 0.5
        // This means she should get back 50 shares (100 * 0.5)
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams,, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: depositAmount});

        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Alice undelegates, which would normally update operator's magnitude to 1.0
        // This tests that the withdrawal still uses the original 0.5 magnitude from when it was queued
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Get shares from withdrawal - should return 50 shares (100 * 0.5) using original magnitude
        // rather than incorrectly returning 100 shares (100 * 1.0) using new magnitude
        (, uint[] memory shares) = delegationManager.getQueuedWithdrawal(withdrawalRoot);
        assertEq(shares[0], 50e18, "shares should be 50e18 (100e18 * 0.5) using original magnitude");
    }
}

contract DelegationManagerUnitTests_getQueuedWithdrawal is DelegationManagerUnitTests {
    using ArrayLib for *;
    using SlashingLib for *;

    function test_getQueuedWithdrawal_Correctness(Randomness r) public rand(r) {
        // Set up initial deposit
        uint depositAmount = r.Uint256(1 ether, 100 ether);
        _depositIntoStrategies(defaultStaker, strategyMock.toArray(), depositAmount.toArrayU256());

        // Register operator and delegate
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawal
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams,, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: depositAmount});

        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Get shares from queued withdrawal
        (, uint[] memory shares) = delegationManager.getQueuedWithdrawal(withdrawalRoot);

        // Verify withdrawal details match
        assertEq(shares.length, 1, "incorrect shares array length");
        assertEq(shares[0], depositAmount, "incorrect shares amount");
    }

    function test_getQueuedWithdrawal_AfterSlashing(Randomness r) public rand(r) {
        // Set up initial deposit
        uint depositAmount = r.Uint256(1 ether, 100 ether);
        _depositIntoStrategies(defaultStaker, strategyMock.toArray(), depositAmount.toArrayU256());

        // Register operator and delegate
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawal
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams,, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawalsSingleStrat({staker: defaultStaker, strategy: strategyMock, depositSharesToWithdraw: depositAmount});

        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Slash operator by 50%
        _setOperatorMagnitude(defaultOperator, strategyMock, 0.5 ether);
        cheats.prank(address(allocationManagerMock));
        delegationManager.slashOperatorShares(defaultOperator, strategyMock, WAD, 0.5 ether);

        // Get shares from queued withdrawal
        (, uint[] memory shares) = delegationManager.getQueuedWithdrawal(withdrawalRoot);

        // Verify withdrawal details match and shares are slashed
        assertEq(shares.length, 1, "incorrect shares array length");
        assertEq(shares[0], depositAmount / 2, "shares not properly slashed");
    }

    function test_getQueuedWithdrawal_NonexistentWithdrawal() public view {
        bytes32 nonexistentRoot = bytes32(uint(1));
        (, uint[] memory shares) = delegationManager.getQueuedWithdrawal(nonexistentRoot);
        assertEq(shares.length, 0, "shares array should be empty");
    }

    function test_getQueuedWithdrawal_MultipleStrategies(Randomness r) public rand(r) {
        // Set up multiple strategies with deposits
        uint numStrategies = r.Uint256(2, 5);
        uint[] memory depositShares = r.Uint256Array({len: numStrategies, min: 1 ether, max: 100 ether});

        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(defaultStaker, depositShares, false);

        // Register operator and delegate
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Queue withdrawals for multiple strategies
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams,, bytes32 withdrawalRoot) =
            _setUpQueueWithdrawals({staker: defaultStaker, strategies: strategies, depositWithdrawalAmounts: depositShares});

        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Get shares from queued withdrawal
        (, uint[] memory shares) = delegationManager.getQueuedWithdrawal(withdrawalRoot);

        // Verify withdrawal details and shares for each strategy
        assertEq(shares.length, numStrategies, "incorrect shares array length");
        for (uint i = 0; i < numStrategies; i++) {
            assertEq(shares[i], depositShares[i], "incorrect shares amount for strategy");
        }
    }

    function testFuzz_getQueuedWithdrawal_EmptyWithdrawal(bytes32 withdrawalRoot) public view {
        (, uint[] memory shares) = delegationManager.getQueuedWithdrawal(withdrawalRoot);
        assertEq(shares.length, 0, "sanity check");
    }
}
