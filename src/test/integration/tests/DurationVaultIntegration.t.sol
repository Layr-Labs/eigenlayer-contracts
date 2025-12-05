// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/interfaces/IDurationVaultStrategy.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IRewardsCoordinator.sol";
import "src/contracts/strategies/StrategyBaseTVLLimits.sol";

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/AVS.t.sol";

contract Integration_DurationVault is IntegrationCheckUtils {
    struct DurationVaultContext {
        IDurationVaultStrategy vault;
        ERC20PresetFixedSupply asset;
        AVS avs;
        OperatorSet operatorSet;
    }

    uint32 internal constant DEFAULT_DURATION = 10 days;
    uint internal constant VAULT_MAX_PER_DEPOSIT = 200 ether;
    uint internal constant VAULT_STAKE_CAP = 1000 ether;

    bool internal durationVaultSupported;

    modifier vaultSupported() {
        if (!durationVaultSupported) {
            console.log("DurationVaultIntegration: duration vault beacon not configured, skipping");
            return;
        }
        _;
    }

    function setUp() public virtual override {
        super.setUp();
        durationVaultSupported = address(strategyFactory.durationVaultBeacon()) != address(0);
    }

    function test_durationVaultLifecycle_flow_deposit_lock_mature() public vaultSupported {
        DurationVaultContext memory ctx = _deployDurationVault(_randomInsuranceRecipient());
        User staker = new User("duration-staker");

        uint depositAmount = 120 ether;
        ctx.asset.transfer(address(staker), depositAmount);

        IStrategy[] memory strategies = _durationStrategyArray(ctx.vault);
        uint[] memory tokenBalances = _singleAmountArray(depositAmount);
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);

        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, depositShares);

        _delegateToVault(staker, ctx.vault);
        ctx.vault.lock();
        assertTrue(ctx.vault.allocationsActive(), "allocations should be active after lock");
        assertTrue(allocationManager.isOperatorSlashable(address(ctx.vault), ctx.operatorSet), "should be slashable");

        // Cannot deposit once locked.
        uint extraDeposit = 10 ether;
        User lateStaker = new User("duration-late-staker");
        ctx.asset.transfer(address(lateStaker), extraDeposit);
        uint[] memory lateTokenBalances = _singleAmountArray(extraDeposit);
        cheats.expectRevert(IDurationVaultStrategy.DepositsLocked.selector);
        lateStaker.depositIntoEigenlayer(strategies, lateTokenBalances);

        // Queue withdrawal prior to maturity.
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, withdrawableShares);

        // Mature the vault and allow withdrawals.
        cheats.warp(block.timestamp + ctx.vault.duration() + 1);
        ctx.vault.markMatured();
        assertTrue(ctx.vault.withdrawalsOpen(), "withdrawals must open after maturity");
        assertFalse(ctx.vault.allocationsActive(), "allocations disabled after maturity");

        // Advance past deallocation delay to ensure slashability cleared.
        cheats.roll(block.number + allocationManager.DEALLOCATION_DELAY() + 1);
        assertFalse(allocationManager.isOperatorSlashable(address(ctx.vault), ctx.operatorSet), "should not be slashable");

        _rollBlocksForCompleteWithdrawals(withdrawals);
        IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[0]);
        assertEq(address(tokens[0]), address(ctx.asset), "unexpected token returned");
        assertEq(ctx.asset.balanceOf(address(staker)), depositAmount, "staker should recover deposit");
    }

    function test_durationVault_operatorIntegrationAndMetadataUpdate() public vaultSupported {
        DurationVaultContext memory ctx = _deployDurationVault(_randomInsuranceRecipient());

        assertTrue(delegationManager.isOperator(address(ctx.vault)), "vault must self-register as operator");
        assertEq(delegationManager.delegatedTo(address(ctx.vault)), address(ctx.vault), "vault should self delegate");
        assertTrue(ctx.vault.operatorSetRegistered(), "operator set should be marked registered");

        (address avsAddress, uint32 operatorSetId) = ctx.vault.operatorSetInfo();
        assertEq(avsAddress, ctx.operatorSet.avs, "avs mismatch");
        assertEq(operatorSetId, ctx.operatorSet.id, "operator set id mismatch");
        assertTrue(allocationManager.isOperatorSlashable(address(ctx.vault), ctx.operatorSet), "vault must be slashable");

        string memory newURI = "ipfs://duration-vault/metadata";
        cheats.expectEmit(false, false, false, true, address(ctx.vault));
        emit IDurationVaultStrategy.MetadataURIUpdated(newURI);
        ctx.vault.updateMetadataURI(newURI);
        assertEq(ctx.vault.metadataURI(), newURI, "metadata not updated");
    }

    function test_durationVault_TVLLimits_enforced_and_frozen_after_lock() public vaultSupported {
        DurationVaultContext memory ctx = _deployDurationVault(_randomInsuranceRecipient());
        User staker = new User("duration-tvl-staker");

        (uint maxPerDepositBefore, uint maxStakeBefore) = StrategyBaseTVLLimits(address(ctx.vault)).getTVLLimits();
        assertEq(maxPerDepositBefore, VAULT_MAX_PER_DEPOSIT, "initial max per deposit mismatch");
        assertEq(maxStakeBefore, VAULT_STAKE_CAP, "initial stake cap mismatch");

        uint newMaxPerDeposit = 50 ether;
        uint newStakeCap = 100 ether;
        ctx.vault.updateTVLLimits(newMaxPerDeposit, newStakeCap);

        IStrategy[] memory strategies = _durationStrategyArray(ctx.vault);
        uint[] memory amounts = _singleAmountArray(60 ether);
        ctx.asset.transfer(address(staker), amounts[0]);
        cheats.expectRevert(IStrategyErrors.MaxPerDepositExceedsMax.selector);
        staker.depositIntoEigenlayer(strategies, amounts);

        // Deposit within limits.
        uint firstDeposit = 50 ether;
        ctx.asset.transfer(address(staker), firstDeposit);
        amounts[0] = firstDeposit;
        staker.depositIntoEigenlayer(strategies, amounts);

        uint secondDeposit = 40 ether;
        ctx.asset.transfer(address(staker), secondDeposit);
        amounts[0] = secondDeposit;
        staker.depositIntoEigenlayer(strategies, amounts);

        // Hitting total cap reverts.
        uint thirdDeposit = 20 ether;
        ctx.asset.transfer(address(staker), thirdDeposit);
        amounts[0] = thirdDeposit;
        cheats.expectRevert(IStrategyErrors.BalanceExceedsMaxTotalDeposits.selector);
        staker.depositIntoEigenlayer(strategies, amounts);

        ctx.vault.lock();
        cheats.expectRevert(IDurationVaultStrategy.DepositsLocked.selector);
        ctx.vault.updateTVLLimits(10 ether, 20 ether);
    }

    function test_durationVault_rewards_claim_while_locked() public vaultSupported {
        DurationVaultContext memory ctx = _deployDurationVault(_randomInsuranceRecipient());
        User staker = new User("duration-reward-staker");
        uint depositAmount = 80 ether;
        ctx.asset.transfer(address(staker), depositAmount);

        IStrategy[] memory strategies = _durationStrategyArray(ctx.vault);
        uint[] memory tokenBalances = _singleAmountArray(depositAmount);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        _delegateToVault(staker, ctx.vault);
        ctx.vault.lock();

        // Prepare reward token and fund RewardsCoordinator.
        ERC20PresetFixedSupply rewardToken = new ERC20PresetFixedSupply("RewardToken", "RWRD", 1e24, address(this));
        uint rewardAmount = 25 ether;
        rewardToken.transfer(address(rewardsCoordinator), rewardAmount);

        // Allow this test to submit roots.
        cheats.startPrank(executorMultisig);
        rewardsCoordinator.setRewardsUpdater(address(this));
        cheats.stopPrank();

        // Build single-earner, single-token merkle claim.
        (IRewardsCoordinatorTypes.RewardsMerkleClaim memory claim, bytes32 rootHash) =
            _buildSingleLeafClaim(address(staker), rewardToken, rewardAmount);

        rewardsCoordinator.submitRoot(rootHash, uint32(block.timestamp - 1));
        claim.rootIndex = uint32(rewardsCoordinator.getDistributionRootsLength() - 1);

        cheats.prank(address(staker));
        rewardsCoordinator.processClaim(claim, address(staker));
        assertEq(rewardToken.balanceOf(address(staker)), rewardAmount, "staker failed to claim rewards");
    }

    function test_durationVault_slashing_routes_to_insurance_and_blocks_after_maturity() public vaultSupported {
        address insuranceRecipient = _randomInsuranceRecipient();
        DurationVaultContext memory ctx = _deployDurationVault(insuranceRecipient);
        User staker = new User("duration-slash-staker");
        uint depositAmount = 200 ether;
        ctx.asset.transfer(address(staker), depositAmount);

        IStrategy[] memory strategies = _durationStrategyArray(ctx.vault);
        uint[] memory tokenBalances = _singleAmountArray(depositAmount);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        _delegateToVault(staker, ctx.vault);
        ctx.vault.lock();

        uint slashWad = 0.25e18;
        IAllocationManager.SlashingParams memory slashParams;
        slashParams.operator = address(ctx.vault);
        slashParams.operatorSetId = ctx.operatorSet.id;
        slashParams.strategies = strategies;
        slashParams.wadsToSlash = _singleAmountArray(slashWad);
        slashParams.description = "insurance event";

        (uint slashId,) = ctx.avs.slashOperator(slashParams);
        uint redistributed =
            strategyManager.clearBurnOrRedistributableSharesByStrategy(ctx.operatorSet, slashId, IStrategy(address(ctx.vault)));
        uint expectedRedistribution = (depositAmount * slashWad) / 1e18;
        assertEq(redistributed, expectedRedistribution, "unexpected redistribution amount");
        assertEq(ctx.asset.balanceOf(insuranceRecipient), expectedRedistribution, "insurance recipient not paid");

        // Mature the vault and advance beyond slashable window.
        cheats.warp(block.timestamp + ctx.vault.duration() + 1);
        ctx.vault.markMatured();
        cheats.roll(block.number + allocationManager.DEALLOCATION_DELAY() + 2);

        cheats.expectRevert(IAllocationManagerErrors.OperatorNotSlashable.selector);
        ctx.avs.slashOperator(slashParams);
        assertEq(ctx.asset.balanceOf(insuranceRecipient), expectedRedistribution, "post-maturity slash should not pay");
    }

    function test_durationVault_slashing_affectsQueuedWithdrawalsAndPaysInsurance() public vaultSupported {
        address insuranceRecipient = _randomInsuranceRecipient();
        DurationVaultContext memory ctx = _deployDurationVault(insuranceRecipient);
        User staker = new User("duration-slash-queued");
        uint depositAmount = 180 ether;
        ctx.asset.transfer(address(staker), depositAmount);

        IStrategy[] memory strategies = _durationStrategyArray(ctx.vault);
        uint[] memory tokenBalances = _singleAmountArray(depositAmount);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        _delegateToVault(staker, ctx.vault);
        ctx.vault.lock();

        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, withdrawableShares);

        uint slashWad = 0.3e18;
        IAllocationManager.SlashingParams memory slashParams;
        slashParams.operator = address(ctx.vault);
        slashParams.operatorSetId = ctx.operatorSet.id;
        slashParams.strategies = strategies;
        slashParams.wadsToSlash = _singleAmountArray(slashWad);
        slashParams.description = "queued withdrawal slash";

        (uint slashId,) = ctx.avs.slashOperator(slashParams);
        uint redistributed =
            strategyManager.clearBurnOrRedistributableSharesByStrategy(ctx.operatorSet, slashId, IStrategy(address(ctx.vault)));
        uint expectedRedistribution = (depositAmount * slashWad) / 1e18;
        assertEq(redistributed, expectedRedistribution, "queued slash redistribution mismatch");
        assertEq(ctx.asset.balanceOf(insuranceRecipient), expectedRedistribution, "insurance recipient incorrect");

        _rollBlocksForCompleteWithdrawals(withdrawals);
        cheats.expectRevert(IDurationVaultStrategy.WithdrawalsLocked.selector);
        staker.completeWithdrawalAsTokens(withdrawals[0]);

        cheats.warp(block.timestamp + ctx.vault.duration() + 1);
        ctx.vault.markMatured();
        cheats.roll(block.number + allocationManager.DEALLOCATION_DELAY() + 2);
        _rollBlocksForCompleteWithdrawals(withdrawals);

        IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[0]);
        assertEq(address(tokens[0]), address(ctx.asset), "unexpected withdrawal token");
        assertEq(
            ctx.asset.balanceOf(address(staker)), depositAmount - expectedRedistribution, "staker balance after queued slash incorrect"
        );
    }

    function _deployDurationVault(address insuranceRecipient) internal returns (DurationVaultContext memory ctx) {
        ERC20PresetFixedSupply asset = new ERC20PresetFixedSupply("Duration Asset", "DURA", 1e24, address(this));
        AVS avsInstance = new AVS("duration-avs");
        avsInstance.updateAVSMetadataURI("https://avs-metadata.local");

        avsInstance.createOperatorSet(new IStrategy[](0));
        OperatorSet memory opSet = avsInstance.createRedistributingOperatorSet(new IStrategy[](0), insuranceRecipient);

        IDurationVaultStrategy.VaultConfig memory config;
        config.underlyingToken = asset;
        config.vaultAdmin = address(this);
        config.duration = DEFAULT_DURATION;
        config.maxPerDeposit = VAULT_MAX_PER_DEPOSIT;
        config.stakeCap = VAULT_STAKE_CAP;
        config.metadataURI = "ipfs://duration-vault";
        config.operatorSet = opSet;
        config.operatorSetRegistrationData = "";
        config.delegationApprover = address(0);
        config.operatorAllocationDelay = 0;
        config.operatorMetadataURI = "ipfs://duration-vault/operator";

        IDurationVaultStrategy vault = IDurationVaultStrategy(address(strategyFactory.deployDurationVaultStrategy(config)));

        IStrategy[] memory strategies = _durationStrategyArray(vault);
        avsInstance.addStrategiesToOperatorSet(opSet.id, strategies);

        ctx = DurationVaultContext({vault: vault, asset: asset, avs: avsInstance, operatorSet: opSet});
    }

    function _durationStrategyArray(IDurationVaultStrategy vault) internal pure returns (IStrategy[] memory arr) {
        arr = new IStrategy[](1);
        arr[0] = IStrategy(address(vault));
    }

    function _singleAmountArray(uint amount) internal pure returns (uint[] memory arr) {
        arr = new uint[](1);
        arr[0] = amount;
    }

    function _delegateToVault(User staker, IDurationVaultStrategy vault) internal {
        IDelegationManager.SignatureWithExpiry memory emptySig;
        cheats.startPrank(address(staker));
        delegationManager.delegateTo(address(vault), emptySig, bytes32(0));
        cheats.stopPrank();
        assertEq(delegationManager.delegatedTo(address(staker)), address(vault), "delegation failed");
    }

    function _randomInsuranceRecipient() internal view returns (address) {
        return address(uint160(uint(keccak256(abi.encodePacked(block.timestamp, address(this))))));
    }

    function _buildSingleLeafClaim(address earner, IERC20 token, uint amount)
        internal
        pure
        returns (IRewardsCoordinatorTypes.RewardsMerkleClaim memory claim, bytes32 rootHash)
    {
        IRewardsCoordinatorTypes.TokenTreeMerkleLeaf[] memory tokenLeaves = new IRewardsCoordinatorTypes.TokenTreeMerkleLeaf[](1);
        tokenLeaves[0] = IRewardsCoordinatorTypes.TokenTreeMerkleLeaf({token: token, cumulativeEarnings: amount});
        bytes32 tokenLeafHash = keccak256(abi.encodePacked(uint8(1), address(token), amount));
        IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf memory earnerLeaf =
            IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf({earner: earner, earnerTokenRoot: tokenLeafHash});
        rootHash = keccak256(abi.encodePacked(uint8(0), earner, tokenLeafHash));

        uint32[] memory tokenIndices = new uint32[](1);
        tokenIndices[0] = 0;
        bytes[] memory tokenProofs = new bytes[](1);
        tokenProofs[0] = bytes("");

        claim = IRewardsCoordinatorTypes.RewardsMerkleClaim({
            rootIndex: 0,
            earnerIndex: 0,
            earnerTreeProof: bytes(""),
            earnerLeaf: earnerLeaf,
            tokenIndices: tokenIndices,
            tokenTreeProofs: tokenProofs,
            tokenLeaves: tokenLeaves
        });
    }
}

