// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "forge-std/console2.sol";

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/AllocationManager.sol";
import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/RewardsCoordinator.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/strategies/StrategyFactory.sol";
import "src/contracts/interfaces/IDurationVaultStrategy.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

import {MockAVSRegistrar} from "./MockAVSRegistrar.sol";

/// @notice End-to-end smoke test for DurationVault on Hoodi (or any live env).
///
/// What this script does on-chain (broadcast):
/// - Creates an AVS (your EOA address) in AllocationManager with a minimal registrar
/// - Creates a redistributing operator set (insurance recipient = you by default)
/// - Deploys an ERC20 test asset + DurationVaultStrategy via StrategyFactory
/// - Adds the vault strategy to the operator set
/// - Delegates + deposits into the vault
/// - Locks the vault, then slashes it and clears redistributable shares to insurance recipient
/// - Creates an AVS rewards submission (sidecar should pick this up and later submit a root)
///
/// This is designed to validate:
/// - vault deploy works against live core protocol
/// - delegation requirement works
/// - lock/allocations/slashability works
/// - slashing routes value to the operator set redistribution recipient
/// - rewards submission plumbing works end-to-end with the sidecar (root + claim is follow-up)
///
/// Address wiring:
/// - You can either pass the 5 core contract addresses as args (recommended), or set env vars:
///   - ALLOCATION_MANAGER
///   - DELEGATION_MANAGER
///   - STRATEGY_MANAGER
///   - STRATEGY_FACTORY
///   - REWARDS_COORDINATOR
contract DurationVaultHoodiE2E is Script {
    using OperatorSetLib for OperatorSet;

    struct PersistedState {
        address avs;
        uint32 operatorSetId;
        address insuranceRecipient;
        address registrar;
        address asset;
        address vault;
        uint32 allocationEffectBlock;
    }

    struct E2EContext {
        // core
        AllocationManager allocationManager;
        DelegationManager delegationManager;
        StrategyManager strategyManager;
        StrategyFactory strategyFactory;
        RewardsCoordinator rewardsCoordinator;
        // identities
        address eoa;
        address avs;
        address insuranceRecipient;
        uint32 operatorSetId;
        OperatorSet opSet;
        // deployed
        MockAVSRegistrar registrar;
        ERC20PresetFixedSupply asset;
        IDurationVaultStrategy vault;
        // config
        uint256 maxPerDeposit;
        uint256 stakeCap;
        uint32 durationSeconds;
        uint256 depositAmount;
        uint256 slashWad;
        // ephemeral
        IStrategy[] strategies;
    }

    /// Optional environment overrides:
    /// - E2E_OPERATOR_SET_ID (uint)
    /// - E2E_VAULT_DURATION_SECONDS (uint)
    /// - E2E_DEPOSIT_AMOUNT (uint)
    /// - E2E_STAKE_CAP (uint)
    /// - E2E_MAX_PER_DEPOSIT (uint)
    /// - E2E_SLASH_WAD (uint)            // 1e18 = 100%
    /// - E2E_INSURANCE_RECIPIENT (address)
    /// - E2E_REWARD_AMOUNT (uint)
    /// - E2E_REWARDS_START_TIMESTAMP (uint)   // defaults to previous CALCULATION_INTERVAL_SECONDS boundary
    /// - E2E_REWARDS_DURATION_SECONDS (uint)  // defaults to CALCULATION_INTERVAL_SECONDS
    /// - E2E_REWARDS_MODE (string): "avs" (default) | "operatorDirectedOperatorSet"
    /// - E2E_REQUIRE_NONZERO_REDISTRIBUTION (bool) // default false; if true, require slashing redistributes > 0
    /// - E2E_REQUIRE_NOW_IN_REWARDS_WINDOW (bool) // default true; if true, require now in [start, start+duration)
    /// - E2E_PHASE (string): "all" (default) | "setup" | "slash" | "rewards"
    /// - E2E_STATE_FILE (string): file to write/read persisted state (default "script/operations/e2e/e2e-state.json")

    /// @notice Run using addresses from environment variables.
    /// @dev Use `--sig "run()"` (default) and set ALLOCATION_MANAGER/DELEGATION_MANAGER/...
    function run() public {
        (
            AllocationManager allocationManager,
            DelegationManager delegationManager,
            StrategyManager strategyManager,
            StrategyFactory strategyFactory,
            RewardsCoordinator rewardsCoordinator
        ) = _loadCoreFromEnv();
        _runWithCore(allocationManager, delegationManager, strategyManager, strategyFactory, rewardsCoordinator);
    }

    /// @notice Run using addresses passed as arguments.
    /// @dev Use:
    /// forge script ... --sig "runWithCore(address,address,address,address,address)" -- <allocationManager> <delegationManager> <strategyManager> <strategyFactory> <rewardsCoordinator>
    function runWithCore(
        address allocationManager,
        address delegationManager,
        address strategyManager,
        address strategyFactory,
        address rewardsCoordinator
    ) public {
        _runWithCore(
            AllocationManager(allocationManager),
            DelegationManager(delegationManager),
            StrategyManager(strategyManager),
            StrategyFactory(strategyFactory),
            RewardsCoordinator(rewardsCoordinator)
        );
    }

    function _runWithCore(
        AllocationManager allocationManager,
        DelegationManager delegationManager,
        StrategyManager strategyManager,
        StrategyFactory strategyFactory,
        RewardsCoordinator rewardsCoordinator
    ) internal {
        // If you want fork-only simulation without passing a private key, you can set E2E_SENDER=<address>.
        address sender = vm.envOr("E2E_SENDER", address(0));
        if (sender != address(0)) vm.startBroadcast(sender);
        else vm.startBroadcast();

        string memory phase = vm.envOr("E2E_PHASE", string("all"));
        bytes32 phaseHash = keccak256(bytes(phase));

        if (phaseHash == keccak256("setup")) {
            E2EContext memory ctx = _initContext(
                allocationManager, delegationManager, strategyManager, strategyFactory, rewardsCoordinator
            );
            _setupAVS(ctx);
            _createOperatorSet(ctx);
            _deployVault(ctx);
            _depositAndLock(ctx);
            _persistState(ctx);
        } else if (phaseHash == keccak256("slash")) {
            E2EContext memory ctx = _initContext(
                allocationManager, delegationManager, strategyManager, strategyFactory, rewardsCoordinator
            );
            PersistedState memory st = _loadState();
            require(st.avs == ctx.avs, "state avs != broadcaster");
            require(st.operatorSetId == ctx.operatorSetId, "state operatorSetId mismatch");

            ctx.insuranceRecipient = st.insuranceRecipient;
            ctx.opSet = OperatorSet({avs: ctx.avs, id: st.operatorSetId});
            ctx.asset = ERC20PresetFixedSupply(st.asset);
            ctx.vault = IDurationVaultStrategy(st.vault);
            ctx.strategies = new IStrategy[](1);
            ctx.strategies[0] = IStrategy(st.vault);

            // Ensure the allocation has become effective, otherwise slashing can legitimately produce 0.
            IAllocationManagerTypes.Allocation memory alloc =
                ctx.allocationManager.getAllocation(address(ctx.vault), ctx.opSet, IStrategy(address(ctx.vault)));
            if (alloc.currentMagnitude == 0 && block.number < alloc.effectBlock) {
                console2.log("Allocation not effective yet. Wait until block >=", uint256(alloc.effectBlock));
                console2.log("Current block:", block.number);
                revert("allocation not effective yet");
            }

            _slashAndRedistribute(ctx);
            _createRewardsSubmission(ctx);
        } else if (phaseHash == keccak256("rewards")) {
            // Rewards-only: reuse the previously deployed vault and just create a new rewards submission in a window
            // that includes "now" (the script enforces this).
            E2EContext memory ctx = _initContext(
                allocationManager, delegationManager, strategyManager, strategyFactory, rewardsCoordinator
            );
            PersistedState memory st = _loadState();
            require(st.avs == ctx.avs, "state avs != broadcaster");
            require(st.operatorSetId == ctx.operatorSetId, "state operatorSetId mismatch");

            ctx.insuranceRecipient = st.insuranceRecipient;
            ctx.opSet = OperatorSet({avs: ctx.avs, id: st.operatorSetId});
            ctx.asset = ERC20PresetFixedSupply(st.asset);
            ctx.vault = IDurationVaultStrategy(st.vault);
            ctx.strategies = new IStrategy[](1);
            ctx.strategies[0] = IStrategy(st.vault);

            _createRewardsSubmission(ctx);
        } else {
            // "all" (default): run everything in one shot; slashing may produce 0 if allocation isn't effective yet.
            E2EContext memory ctx = _initContext(
                allocationManager, delegationManager, strategyManager, strategyFactory, rewardsCoordinator
            );
            _setupAVS(ctx);
            _createOperatorSet(ctx);
            _deployVault(ctx);
            _depositAndLock(ctx);
            _slashAndRedistribute(ctx);
            _createRewardsSubmission(ctx);
        }
        vm.stopBroadcast();
    }

    function _loadCoreFromEnv()
        internal
        view
        returns (AllocationManager, DelegationManager, StrategyManager, StrategyFactory, RewardsCoordinator)
    {
        address allocationManager = vm.envAddress("ALLOCATION_MANAGER");
        address delegationManager = vm.envAddress("DELEGATION_MANAGER");
        address strategyManager = vm.envAddress("STRATEGY_MANAGER");
        address strategyFactory = vm.envAddress("STRATEGY_FACTORY");
        address rewardsCoordinator = vm.envAddress("REWARDS_COORDINATOR");
        return (
            AllocationManager(allocationManager),
            DelegationManager(delegationManager),
            StrategyManager(strategyManager),
            StrategyFactory(strategyFactory),
            RewardsCoordinator(rewardsCoordinator)
        );
    }

    function _initContext(
        AllocationManager allocationManager,
        DelegationManager delegationManager,
        StrategyManager strategyManager,
        StrategyFactory strategyFactory,
        RewardsCoordinator rewardsCoordinator
    ) internal view returns (E2EContext memory ctx) {
        ctx.allocationManager = allocationManager;
        ctx.delegationManager = delegationManager;
        ctx.strategyManager = strategyManager;
        ctx.strategyFactory = strategyFactory;
        ctx.rewardsCoordinator = rewardsCoordinator;

        // In forge scripts, `msg.sender` is the script contract; use tx.origin as the broadcaster address.
        ctx.eoa = tx.origin;
        ctx.avs = ctx.eoa;

        ctx.operatorSetId = uint32(vm.envOr("E2E_OPERATOR_SET_ID", uint256(1)));
        ctx.insuranceRecipient = vm.envOr("E2E_INSURANCE_RECIPIENT", address(ctx.eoa));

        ctx.maxPerDeposit = vm.envOr("E2E_MAX_PER_DEPOSIT", uint256(200 ether));
        ctx.stakeCap = vm.envOr("E2E_STAKE_CAP", uint256(1000 ether));
        ctx.durationSeconds = uint32(vm.envOr("E2E_VAULT_DURATION_SECONDS", uint256(120))); // 2 minutes default
        ctx.depositAmount = vm.envOr("E2E_DEPOSIT_AMOUNT", uint256(100 ether));
        ctx.slashWad = vm.envOr("E2E_SLASH_WAD", uint256(0.25e18));

        ctx.opSet = OperatorSet({avs: ctx.avs, id: ctx.operatorSetId});
    }

    function _setupAVS(
        E2EContext memory ctx
    ) internal returns (E2EContext memory) {
        ctx.allocationManager.updateAVSMetadataURI(ctx.avs, "ipfs://e2e-duration-vault/avs-metadata");
        ctx.registrar = new MockAVSRegistrar(ctx.avs);
        ctx.allocationManager.setAVSRegistrar(ctx.avs, ctx.registrar);
        require(address(ctx.allocationManager.getAVSRegistrar(ctx.avs)) == address(ctx.registrar), "registrar not set");
        return ctx;
    }

    function _createOperatorSet(
        E2EContext memory ctx
    ) internal returns (E2EContext memory) {
        IAllocationManagerTypes.CreateSetParamsV2[] memory sets = new IAllocationManagerTypes.CreateSetParamsV2[](1);
        sets[0] = IAllocationManagerTypes.CreateSetParamsV2({
            operatorSetId: ctx.operatorSetId,
            strategies: new IStrategy[](0),
            slasher: ctx.eoa
        });
        address[] memory recipients = new address[](1);
        recipients[0] = ctx.insuranceRecipient;
        ctx.allocationManager.createRedistributingOperatorSets(ctx.avs, sets, recipients);
        require(ctx.allocationManager.isOperatorSet(ctx.opSet), "operator set not created");
        require(ctx.allocationManager.isRedistributingOperatorSet(ctx.opSet), "operator set not redistributing");
        require(
            ctx.allocationManager.getRedistributionRecipient(ctx.opSet) == ctx.insuranceRecipient,
            "bad redistribution recipient"
        );
        require(ctx.allocationManager.getSlasher(ctx.opSet) == ctx.eoa, "bad slasher");
        return ctx;
    }

    function _deployVault(
        E2EContext memory ctx
    ) internal returns (E2EContext memory) {
        ctx.asset = new ERC20PresetFixedSupply("Hoodi Duration Vault Asset", "HDVA", 1_000_000 ether, ctx.eoa);

        IDurationVaultStrategyTypes.VaultConfig memory cfg;
        cfg.underlyingToken = ctx.asset;
        cfg.vaultAdmin = ctx.eoa;
        cfg.arbitrator = ctx.eoa;
        cfg.duration = ctx.durationSeconds;
        cfg.maxPerDeposit = ctx.maxPerDeposit;
        cfg.stakeCap = ctx.stakeCap;
        cfg.metadataURI = "ipfs://e2e-duration-vault/metadata";
        cfg.operatorSet = ctx.opSet;
        cfg.operatorSetRegistrationData = "";
        cfg.delegationApprover = address(0);
        cfg.operatorMetadataURI = "ipfs://e2e-duration-vault/operator-metadata";

        ctx.vault = ctx.strategyFactory.deployDurationVaultStrategy(cfg);

        ctx.strategies = new IStrategy[](1);
        ctx.strategies[0] = IStrategy(address(ctx.vault));
        ctx.allocationManager.addStrategiesToOperatorSet(ctx.avs, ctx.operatorSetId, ctx.strategies);
        require(
            _operatorSetContainsStrategy(ctx.allocationManager, ctx.opSet, IStrategy(address(ctx.vault))),
            "vault not in operator set"
        );
        return ctx;
    }

    function _depositAndLock(
        E2EContext memory ctx
    ) internal returns (E2EContext memory) {
        require(ctx.depositAmount <= ctx.maxPerDeposit, "deposit > maxPerDeposit");

        IDelegationManager.SignatureWithExpiry memory emptySig;
        ctx.delegationManager.delegateTo(address(ctx.vault), emptySig, bytes32(0));

        ctx.asset.approve(address(ctx.strategyManager), ctx.depositAmount);
        uint256 depositShares =
            ctx.strategyManager.depositIntoStrategy(IStrategy(address(ctx.vault)), ctx.asset, ctx.depositAmount);
        require(depositShares != 0, "deposit shares = 0");
        require(
            _getDepositedShares(ctx.strategyManager, ctx.eoa, IStrategy(address(ctx.vault))) == depositShares,
            "unexpected shares"
        );

        ctx.vault.lock();
        require(ctx.vault.allocationsActive(), "allocations not active after lock");
        require(ctx.allocationManager.isOperatorSlashable(address(ctx.vault), ctx.opSet), "vault not slashable");
        return ctx;
    }

    function _slashAndRedistribute(
        E2EContext memory ctx
    ) internal returns (E2EContext memory) {
        require(ctx.slashWad > 0 && ctx.slashWad <= 1e18, "invalid slash wad");

        uint256 slashCountBefore = ctx.allocationManager.getSlashCount(ctx.opSet);

        IAllocationManagerTypes.SlashingParams memory slashParams;
        slashParams.operator = address(ctx.vault);
        slashParams.operatorSetId = ctx.operatorSetId;
        slashParams.strategies = ctx.strategies;
        slashParams.wadsToSlash = new uint256[](1);
        slashParams.wadsToSlash[0] = ctx.slashWad;
        slashParams.description = "e2e-duration-vault-slash";

        (uint256 slashId,) = ctx.allocationManager.slashOperator(ctx.avs, slashParams);
        require(ctx.allocationManager.getSlashCount(ctx.opSet) == slashCountBefore + 1, "slash count not incremented");
        require(slashId == slashCountBefore + 1, "unexpected slashId");

        uint256 sharesToRedistribute =
            ctx.strategyManager.getBurnOrRedistributableShares(ctx.opSet, slashId, IStrategy(address(ctx.vault)));

        uint256 insuranceBefore = ctx.asset.balanceOf(ctx.insuranceRecipient);
        uint256 redistributed = ctx.strategyManager
            .clearBurnOrRedistributableSharesByStrategy(ctx.opSet, slashId, IStrategy(address(ctx.vault)));
        uint256 insuranceAfter = ctx.asset.balanceOf(ctx.insuranceRecipient);

        require(insuranceAfter >= insuranceBefore, "insurance balance decreased");
        require(insuranceAfter - insuranceBefore == redistributed, "redistributed != insurance delta");

        // In some live-network conditions (notably allocation delay / effectBlock timing), the slash can legitimately
        // result in zero burn/redistributable shares. That still validates that the slashing plumbing works.
        //
        // For fork confidence, you can require a nonzero redistribution:
        //   export E2E_REQUIRE_NONZERO_REDISTRIBUTION=true
        bool requireNonzero = vm.envOr("E2E_REQUIRE_NONZERO_REDISTRIBUTION", false);
        if (requireNonzero) {
            require(sharesToRedistribute != 0, "no shares to redistribute");
            require(redistributed != 0, "no tokens redistributed");
        }
        return ctx;
    }

    function _createRewardsSubmission(
        E2EContext memory ctx
    ) internal {
        string memory mode = vm.envOr("E2E_REWARDS_MODE", string("avs"));
        bytes32 modeHash = keccak256(bytes(mode));

        ERC20PresetFixedSupply rewardToken =
            new ERC20PresetFixedSupply("Hoodi Rewards Token", "HDRW", 1_000_000 ether, ctx.eoa);
        uint256 rewardAmount = vm.envOr("E2E_REWARD_AMOUNT", uint256(10 ether));

        IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory sams =
            new IRewardsCoordinatorTypes.StrategyAndMultiplier[](1);
        sams[0] =
            IRewardsCoordinatorTypes.StrategyAndMultiplier({strategy: IStrategy(address(ctx.vault)), multiplier: 1e18});

        uint32 interval = ctx.rewardsCoordinator.CALCULATION_INTERVAL_SECONDS();

        // RewardsCoordinator constraints (enforced again on-chain):
        // - startTimestamp % CALCULATION_INTERVAL_SECONDS == 0
        // - duration % CALCULATION_INTERVAL_SECONDS == 0
        uint32 defaultDuration = interval;
        uint32 rewardsDuration = uint32(vm.envOr("E2E_REWARDS_DURATION_SECONDS", uint256(defaultDuration)));
        require(rewardsDuration % interval == 0, "rewardsDuration must be multiple of CALCULATION_INTERVAL_SECONDS");

        // Mode A: Standard AVS rewards submission (works for protocol; may not be indexed by Sidecar if it only
        // understands legacy AVSDirectory registrations).
        if (modeHash == keccak256("avs")) {
            // Defaults are aligned for live networks like preprod-hoodi (interval = 1 day).
            uint32 defaultStartTs = _defaultAlignedRewardsStartTimestamp(interval);
            uint32 startTs = uint32(vm.envOr("E2E_REWARDS_START_TIMESTAMP", uint256(defaultStartTs)));

            require(startTs % interval == 0, "startTs must be multiple of CALCULATION_INTERVAL_SECONDS");

            // Ensure the submission window overlaps the time we actually have stake (this script deposits immediately
            // before creating the rewards submission in the "all" / "slash" flows). If you set a window entirely in the
            // past, Sidecar will correctly omit this earner and you'll see "earner index not found" when generating a
            // claim.
            //
            // For live-network runs, it can be convenient to schedule a future-aligned submission window (e.g. next UTC
            // day) without waiting for the window to start. Set E2E_REQUIRE_NOW_IN_REWARDS_WINDOW=false to skip this.
            bool requireNowInWindow = vm.envOr("E2E_REQUIRE_NOW_IN_REWARDS_WINDOW", true);
            if (requireNowInWindow) {
                require(block.timestamp >= startTs, "now before rewards start");
                require(block.timestamp < uint256(startTs) + uint256(rewardsDuration), "now after rewards end");
            }

            IRewardsCoordinatorTypes.RewardsSubmission memory sub = IRewardsCoordinatorTypes.RewardsSubmission({
                strategiesAndMultipliers: sams,
                token: rewardToken,
                amount: rewardAmount,
                startTimestamp: startTs,
                duration: rewardsDuration
            });
            _submitAndVerifyRewards(ctx, rewardToken, rewardAmount, sub);
            return;
        }

        // Mode B: Operator-directed rewards submission for the operator set (strictly retroactive).
        // Sidecar can choose to index this path differently. This will revert unless the window is strictly in the past.
        if (modeHash == keccak256("operatorDirectedOperatorSet")) {
            uint256 nowTs = block.timestamp;
            uint256 endTs = (interval == 0) ? nowTs : (nowTs / interval) * interval;
            // Must be strictly retroactive: end < now.
            if (endTs >= nowTs) endTs = endTs - interval;

            uint32 defaultStartTs = uint32(endTs - uint256(rewardsDuration));
            uint32 startTs = uint32(vm.envOr("E2E_REWARDS_START_TIMESTAMP", uint256(defaultStartTs)));
            require(startTs % interval == 0, "startTs must be multiple of CALCULATION_INTERVAL_SECONDS");
            require(uint256(startTs) + uint256(rewardsDuration) < block.timestamp, "operator-directed must be retro");

            _submitAndVerifyOperatorDirectedOperatorSetRewards(
                ctx, rewardToken, rewardAmount, sams, startTs, rewardsDuration
            );
            return;
        }

        revert("unknown E2E_REWARDS_MODE");
    }

    function _submitAndVerifyOperatorDirectedOperatorSetRewards(
        E2EContext memory ctx,
        ERC20PresetFixedSupply rewardToken,
        uint256 rewardAmount,
        IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory sams,
        uint32 startTs,
        uint32 rewardsDuration
    ) internal {
        IRewardsCoordinatorTypes.OperatorReward[] memory ors = new IRewardsCoordinatorTypes.OperatorReward[](1);
        ors[0] = IRewardsCoordinatorTypes.OperatorReward({operator: address(ctx.vault), amount: rewardAmount});

        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission memory od =
            IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
                strategiesAndMultipliers: sams,
                token: rewardToken,
                operatorRewards: ors,
                startTimestamp: startTs,
                duration: rewardsDuration,
                description: "e2e-duration-vault-operator-directed"
            });

        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory ods =
            new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
        ods[0] = od;

        uint256 nonceBefore = ctx.rewardsCoordinator.submissionNonce(ctx.avs);
        uint256 rcBalBefore = rewardToken.balanceOf(address(ctx.rewardsCoordinator));
        bytes32 expectedHash = keccak256(abi.encode(ctx.avs, nonceBefore, od));

        rewardToken.approve(address(ctx.rewardsCoordinator), rewardAmount);
        ctx.rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(ctx.opSet, ods);

        require(ctx.rewardsCoordinator.submissionNonce(ctx.avs) == nonceBefore + 1, "submission nonce not incremented");
        require(
            ctx.rewardsCoordinator.isOperatorDirectedOperatorSetRewardsSubmissionHash(ctx.avs, expectedHash),
            "operator-directed hash not recorded"
        );
        require(
            rewardToken.balanceOf(address(ctx.rewardsCoordinator)) == rcBalBefore + rewardAmount,
            "reward token not transferred"
        );
    }

    function _submitAndVerifyRewards(
        E2EContext memory ctx,
        ERC20PresetFixedSupply rewardToken,
        uint256 rewardAmount,
        IRewardsCoordinatorTypes.RewardsSubmission memory sub
    ) internal {
        IRewardsCoordinatorTypes.RewardsSubmission[] memory subs = new IRewardsCoordinatorTypes.RewardsSubmission[](1);
        subs[0] = sub;

        uint256 nonceBefore = ctx.rewardsCoordinator.submissionNonce(ctx.avs);
        uint256 rcBalBefore = rewardToken.balanceOf(address(ctx.rewardsCoordinator));
        bytes32 expectedHash = keccak256(abi.encode(ctx.avs, nonceBefore, sub));

        rewardToken.approve(address(ctx.rewardsCoordinator), rewardAmount);
        ctx.rewardsCoordinator.createAVSRewardsSubmission(subs);

        require(ctx.rewardsCoordinator.submissionNonce(ctx.avs) == nonceBefore + 1, "submission nonce not incremented");
        require(
            ctx.rewardsCoordinator.isAVSRewardsSubmissionHash(ctx.avs, expectedHash), "submission hash not recorded"
        );
        require(
            rewardToken.balanceOf(address(ctx.rewardsCoordinator)) == rcBalBefore + rewardAmount,
            "reward token not transferred"
        );
    }

    function _defaultAlignedRewardsStartTimestamp(
        uint32 interval
    ) internal view returns (uint32) {
        // Use the current interval boundary so the window includes "now".
        // Example: interval=86400, now=10:15 UTC => start at today 00:00 UTC.
        uint256 nowTs = block.timestamp;
        if (interval == 0) return uint32(nowTs);
        uint256 floored = (nowTs / interval) * interval;
        return uint32(floored);
    }

    function _operatorSetContainsStrategy(
        AllocationManager allocationManager,
        OperatorSet memory operatorSet,
        IStrategy strategy
    ) internal view returns (bool) {
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        for (uint256 i = 0; i < strategies.length; i++) {
            if (address(strategies[i]) == address(strategy)) return true;
        }
        return false;
    }

    function _getDepositedShares(
        StrategyManager strategyManager,
        address staker,
        IStrategy strategy
    ) internal view returns (uint256) {
        (IStrategy[] memory strategies, uint256[] memory shares) = strategyManager.getDeposits(staker);
        for (uint256 i = 0; i < strategies.length; i++) {
            if (address(strategies[i]) == address(strategy)) return shares[i];
        }
        return 0;
    }

    function _stateFile() internal view returns (string memory) {
        return vm.envOr("E2E_STATE_FILE", string("script/operations/e2e/e2e-state.json"));
    }

    function _persistState(
        E2EContext memory ctx
    ) internal {
        // Capture the allocation effect block so the user knows how long to wait before running the slash phase.
        IAllocationManagerTypes.Allocation memory alloc =
            ctx.allocationManager.getAllocation(address(ctx.vault), ctx.opSet, IStrategy(address(ctx.vault)));

        string memory json = string.concat(
            "{",
            "\"avs\":\"",
            vm.toString(ctx.avs),
            "\",",
            "\"operatorSetId\":",
            vm.toString(uint256(ctx.operatorSetId)),
            ",",
            "\"insuranceRecipient\":\"",
            vm.toString(ctx.insuranceRecipient),
            "\",",
            "\"registrar\":\"",
            vm.toString(address(ctx.registrar)),
            "\",",
            "\"asset\":\"",
            vm.toString(address(ctx.asset)),
            "\",",
            "\"vault\":\"",
            vm.toString(address(ctx.vault)),
            "\",",
            "\"allocationEffectBlock\":",
            vm.toString(uint256(alloc.effectBlock)),
            "}"
        );

        string memory fullPath = string.concat(vm.projectRoot(), "/", _stateFile());
        vm.writeFile(fullPath, json);
        console2.log("Wrote E2E state to", fullPath);
        console2.log("Allocation effectBlock:", uint256(alloc.effectBlock));
        console2.log("Current block:", block.number);
    }

    function _loadState() internal view returns (PersistedState memory st) {
        string memory fullPath = string.concat(vm.projectRoot(), "/", _stateFile());
        string memory json = vm.readFile(fullPath);

        st.avs = abi.decode(vm.parseJson(json, ".avs"), (address));
        st.operatorSetId = uint32(abi.decode(vm.parseJson(json, ".operatorSetId"), (uint256)));
        st.insuranceRecipient = abi.decode(vm.parseJson(json, ".insuranceRecipient"), (address));
        st.registrar = abi.decode(vm.parseJson(json, ".registrar"), (address));
        st.asset = abi.decode(vm.parseJson(json, ".asset"), (address));
        st.vault = abi.decode(vm.parseJson(json, ".vault"), (address));
        st.allocationEffectBlock = uint32(abi.decode(vm.parseJson(json, ".allocationEffectBlock"), (uint256)));
    }
}

