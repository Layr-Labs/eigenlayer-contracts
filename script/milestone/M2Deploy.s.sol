// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/interfaces/IETHPOSDeposit.sol";
import "../../src/contracts/interfaces/IBeaconChainOracle.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/Slasher.sol";
import "../../src/contracts/core/DelegationManager.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";
import "../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../../src/contracts/permissions/PauserRegistry.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/milestone/M2Deploy.s.sol:M2Deploy --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract M2Deploy is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public m1DeploymentOutputPath;
    string public m2DeploymentOutputPath;

    // EigenLayer core contracts
    ISlasher public slasher;
    IDelegationManager public delegation;
    DelegationManager public delegationImplementation;
    IStrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;
    IEigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    IDelayedWithdrawalRouter public delayedWithdrawalRouter;
    IBeacon public eigenPodBeacon;
    EigenPod public eigenPodImplementation;

    // Eigenlayer Proxy Admin
    ProxyAdmin public eigenLayerProxyAdmin;

    // BeaconChain deposit contract & beacon chain oracle
    IETHPOSDeposit public ethPOS;
    address public beaconChainOracle;

    // RPC url to fork from for pre-upgrade state change tests
    string public rpcUrl;

    // Pre-upgrade values to check post-upgrade
    address public strategyWhitelister;
    bytes32 public withdrawalDelayBlocksStorageSlot = bytes32(uint256(204)); // 0xcc == 204
    uint256 public withdrawalsQueuedStorageSlot = 208; //0xd0 = 208
    uint256 public withdrawalDelayBlocks;
    bytes32 public delegationManagerDomainSeparator;
    uint256 public numPods;
    uint256 public maxPods;

    // Pointers to pre-upgrade values for lstDepositor
    address public lstDepositor;
    uint256 public stakerStrategyListLength;
    uint256[] public stakerStrategyShares; // Array of shares in each strategy
    IStrategy[] public stakerStrategyList; // Array of strategies staker has deposited into
    IStrategyManager.DeprecatedStruct_QueuedWithdrawal public queuedWithdrawalLst; // queuedWithdrawal for
    uint256 public m1PostWithdrawTokensReceived; // Number of tokens received after completing withdrawal on M1 contracts
    bytes32 public withdrawalRootBeforeUpgrade;
    IERC20[] public tokensToWithdraw;
    uint256 public lstDepositorNonceBefore;
    uint256 public lstDepositorNumWithdrawalsQueued;
    uint256 public lstDepositorBalancePreUpgrade; // balance after withdrawal on m1 contracts
    uint256 public lstDepositorSharesPreUpgrade; // shares after withdrawal on m1 contracts

    // Pointers to pre-upgrade values for eigenPodDepositor
    address public eigenPodDepositor;
    IEigenPod public eigenPod;
    address public eigenPodOwner;
    bool public hasPod;
    uint64 public mostRecentWithdrawalBlock;

    function run() external {
        // Read and log the chain ID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        // Update deployment path addresses if on mainnet
        if (chainId == 1) {
            m1DeploymentOutputPath = string(bytes("script/output/M1_deployment_mainnet_2023_6_9.json"));
            m2DeploymentOutputPath = "script/output/M2_deployment_data_mainnet.json";
            rpcUrl = "RPC_MAINNET";
        } else if (chainId == 5) {
            m1DeploymentOutputPath = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));
            m2DeploymentOutputPath = "script/output/M2_deployment_data_goerli.json";
            rpcUrl = "RPC_GOERLI";
        } else {
            revert("Chain not supported");
        }

        // Set beacon chain oracle, currently 0 address
        beaconChainOracle = 0x0000000000000000000000000000000000000000;

        // Read json data
        string memory deployment_data = vm.readFile(m1DeploymentOutputPath);
        slasher = Slasher(stdJson.readAddress(deployment_data, ".addresses.slasher"));
        delegation = slasher.delegation();
        strategyManager = slasher.strategyManager();
        eigenPodManager = strategyManager.eigenPodManager();
        eigenPodBeacon = eigenPodManager.eigenPodBeacon();
        ethPOS = eigenPodManager.ethPOS();

        eigenLayerProxyAdmin = ProxyAdmin(stdJson.readAddress(deployment_data, ".addresses.eigenLayerProxyAdmin"));

        // Store pre-upgrade values to check against later
        strategyWhitelister = strategyManager.strategyWhitelister();
        withdrawalDelayBlocks = m1StrategyManager(address(strategyManager)).withdrawalDelayBlocks();
        delegationManagerDomainSeparator = IDelegationManagerV0(address(delegation)).DOMAIN_SEPARATOR();
        numPods = eigenPodManager.numPods();
        maxPods = eigenPodManager.maxPods();
        delayedWithdrawalRouter = EigenPod(payable(eigenPodBeacon.implementation())).delayedWithdrawalRouter();

        // Set chain-specific values
        IStrategy[] memory strategyArray = new IStrategy[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        if (chainId == 1) {
            // no-op for now
        } else if (chainId == 5) {
            // Set LST Depositor values
            lstDepositor = 0x01e453D2465cEC1BD2ac9aed06115Fbf28482b33;
            strategyArray[0] = IStrategy(0x879944A8cB437a5f8061361f82A6d4EED59070b5);
            shareAmounts[0] = 188647761812080108;
            IStrategyManager.DeprecatedStruct_WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager
                .DeprecatedStruct_WithdrawerAndNonce({withdrawer: lstDepositor, nonce: uint96(0)});
            queuedWithdrawalLst = IStrategyManager.DeprecatedStruct_QueuedWithdrawal({
                strategies: strategyArray,
                shares: shareAmounts,
                staker: lstDepositor,
                withdrawerAndNonce: withdrawerAndNonce,
                withdrawalStartBlock: uint32(9083727),
                delegatedAddress: delegation.delegatedTo(lstDepositor)
            });
            tokensToWithdraw.push(IERC20(0x178E141a0E3b34152f73Ff610437A7bf9B83267A));

            // Set eigenPod owner values
            eigenPodDepositor = 0xE9D04433bac1bd584B0493cbaBa170CCCBDA8F00;
        } else {
            revert("chain ID not supported");
        }

        // Store LST depositor pre-upgrade values
        stakerStrategyListLength = strategyManager.stakerStrategyListLength(lstDepositor);
        (stakerStrategyList, stakerStrategyShares) = strategyManager.getDeposits(lstDepositor);
        withdrawalRootBeforeUpgrade = strategyManager.calculateWithdrawalRoot(queuedWithdrawalLst);
        lstDepositorNonceBefore = StrategyManagerStorage(address(strategyManager)).nonces(lstDepositor);
        lstDepositorNumWithdrawalsQueued = m1StrategyManager(address(strategyManager)).numWithdrawalsQueued(
            lstDepositor
        );

        // Store eigenPod owner pre-ugprade values
        eigenPod = eigenPodManager.ownerToPod(eigenPodDepositor);
        require(address(eigenPod).balance > 0, "eigenPod to test has balance of 0");
        hasPod = eigenPodManager.hasPod(eigenPodDepositor);
        eigenPodOwner = eigenPod.podOwner();
        mostRecentWithdrawalBlock = m1EigenPod(address(eigenPod)).mostRecentWithdrawalBlockNumber();

        // Complete queued withdrawals before upgrade to sanity check post upgrade
        _completeWithdrawalsPreUpgrade();

        // Begin deployment
        vm.startBroadcast();

        // Deploy new implmementation contracts
        delegationImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        strategyManagerImplementation = new StrategyManager(delegation, eigenPodManager, slasher);
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOS,
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegation
        );
        eigenPodImplementation = new EigenPod({
            _ethPOS: ethPOS,
            _delayedWithdrawalRouter: delayedWithdrawalRouter,
            _eigenPodManager: eigenPodManager,
            _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR: 32 gwei,
            _GENESIS_TIME: 1616508000
        });

        vm.stopBroadcast();

        // Write json data out
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";

        // Serialize proxy and non-deployed addresses
        vm.serializeAddress(deployed_addresses, "slasher", address(slasher));
        vm.serializeAddress(deployed_addresses, "delegation", address(delegation));
        vm.serializeAddress(deployed_addresses, "strategyManager", address(strategyManager));
        vm.serializeAddress(deployed_addresses, "delayedWithdrawalRouter", address(delayedWithdrawalRouter));
        vm.serializeAddress(deployed_addresses, "eigenPodManager", address(eigenPodManager));
        vm.serializeAddress(deployed_addresses, "eigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(deployed_addresses, "ethPOS", address(ethPOS));

        // Serialize new implementation addresses
        vm.serializeAddress(deployed_addresses, "delegationImplementation", address(delegationImplementation));
        vm.serializeAddress(
            deployed_addresses,
            "strategyManagerImplementation",
            address(strategyManagerImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "eigenPodManagerImplementation",
            address(eigenPodManagerImplementation)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "eigenPodImplementation",
            address(eigenPodImplementation)
        );

        // Add chain info
        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", chainId);

        // Save addresses
        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        string memory finalJson = vm.serializeString(parent_object, chain_info, chain_info_output);

        // Write output to file
        vm.writeJson(finalJson, m2DeploymentOutputPath);

        // Perform post-upgrade tests
        simulatePerformingUpgrade();
        checkUpgradeCorrectness();
    }

    function simulatePerformingUpgrade() public {
        cheats.startPrank(eigenLayerProxyAdmin.owner());
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(delegation))),
            address(delegationImplementation)
        );
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation)
        );
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation)
        );
        cheats.stopPrank();

        // Upgrade beacon
        cheats.prank(UpgradeableBeacon(address(eigenPodBeacon)).owner());
        UpgradeableBeacon(address(eigenPodBeacon)).upgradeTo(address(eigenPodImplementation));
    }

    function checkUpgradeCorrectness() public {
        _verifyStorageSlots();

        _verifyContractsInitialized();

        _verifyLSTDepositorCorrectness();

        _verifyEigenPodCorrectness();
    }

    // Call contracts to ensure that all simple view functions return the same values (e.g. the return value of `StrategyManager.delegation()` hasn’t changed)
    // StrategyManager: delegation, eigenPodManager, slasher, strategyWhitelister, withdrawalDelayBlocks all unchanged
    // DelegationManager: DOMAIN_SEPARATOR, strategyManager, slasher, eigenPodManager  all unchanged
    // EigenPodManager: ethPOS, eigenPodBeacon, strategyManager, slasher, beaconChainOracle, numPods, maxPods  all unchanged
    // delegationManager is now correct (added immutable)
    // Call contracts to make sure they are still “initialized” (ensure that trying to call initializer reverts)
    function _verifyStorageSlots() internal view {
        // StrategyManager: Check view functions return pre-upgraded values
        require(strategyManager.delegation() == delegation, "strategyManager.delegation incorrect");
        require(strategyManager.eigenPodManager() == eigenPodManager, "strategyManager.eigenPodManager incorrect");
        require(strategyManager.slasher() == slasher, "strategyManager.slasher incorrect");
        require(
            strategyManager.strategyWhitelister() == strategyWhitelister,
            "strategyManager.strategyWhitelister incorrect"
        );
        require(
            cheats.load(address(strategyManager), withdrawalDelayBlocksStorageSlot) == bytes32(withdrawalDelayBlocks),
            "strategyManager.withdrawalDelayBlocks incorrect"
        );
        // DelegationManager: Check view functions return pre-upgraded values
        require(DelegationManagerStorage(address(delegation)).strategyManager() == strategyManager, "delegation.strategyManager incorrect");
        require(
            delegation.domainSeparator() == delegationManagerDomainSeparator,
            "delegation.domainSeparator incorrect"
        );
        require(DelegationManagerStorage(address(delegation)).slasher() == slasher, "delegation.slasher incorrect");
        require(DelegationManagerStorage(address(delegation)).eigenPodManager() == eigenPodManager, "delegation.eigenPodManager incorrect");
        // EigenPodManager: check view functions return pre-upgraded values
        require(eigenPodManager.ethPOS() == ethPOS, "eigenPodManager.ethPOS incorrect");
        require(eigenPodManager.eigenPodBeacon() == eigenPodBeacon, "eigenPodManager.eigenPodBeacon incorrect");
        require(eigenPodManager.strategyManager() == strategyManager, "eigenPodManager.strategyManager incorrect");
        require(eigenPodManager.slasher() == slasher, "eigenPodManager.slasher incorrect");
        require(
            address(eigenPodManager.beaconChainOracle()) == beaconChainOracle,
            "eigenPodManager.beaconChainOracle incorrect"
        );
        require(eigenPodManager.numPods() == numPods, "eigenPodManager.numPods incorrect");
        require(eigenPodManager.maxPods() == maxPods, "eigenPodManager.maxPods incorrect");
        require(EigenPodManagerStorage(address(eigenPodManager)).delegationManager() == delegation, "eigenPodManager.delegationManager incorrect");
    }

    function _verifyContractsInitialized() internal {
        // Check that contracts are unable to be re-initialized
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        StrategyManager(address(strategyManager)).initialize(
            address(this),
            address(this),
            PauserRegistry(address(this)),
            0
        );

        IStrategy[] memory strategyArray = new IStrategy[](0);
        uint256[] memory withdrawalDelayBlocksArray = new uint256[](0);
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        DelegationManager(address(delegation)).initialize(
            address(this),
            PauserRegistry(address(this)),
            0, // initialPausedStatus
            0, // minWithdrawalDelayBLocks
            strategyArray,
            withdrawalDelayBlocksArray
        );

        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        EigenPodManager(address(eigenPodManager)).initialize(
            0,
            IBeaconChainOracle(address(this)),
            address(this),
            PauserRegistry(address(this)),
            0
        );
    }

    function _verifyLSTDepositorCorrectness() internal {
        // Check that LST depositor has the same shares in the same strategies
        require(
            strategyManager.stakerStrategyListLength(lstDepositor) == stakerStrategyListLength,
            "strategyManager.stakerStrategyListLength incorrect"
        );
        (IStrategy[] memory stakerStrategyListAfter, uint256[] memory stakerStrategySharesAfter) = strategyManager
            .getDeposits(lstDepositor);
        for (uint256 i = 0; i < stakerStrategyListAfter.length; i++) {
            require(
                stakerStrategyListAfter[i] == stakerStrategyList[i],
                "strategyManager.stakerStrategyList incorrect"
            );
            require(
                stakerStrategySharesAfter[i] == stakerStrategyShares[i],
                "strategyManager.stakerStrategyShares incorrect"
            );
        }

        // Check that withdrawal root resolves to prev root
        require(
            withdrawalRootBeforeUpgrade == strategyManager.calculateWithdrawalRoot(queuedWithdrawalLst),
            "strategyManager.calculateWithdrawalRoot does not resolve to previous root"
        );
        require(
            StrategyManagerStorage(address(strategyManager)).withdrawalRootPending(withdrawalRootBeforeUpgrade),
            "strategyManager.withdrawalRootPending incorrect"
        );

        // Check that nonce and numWithdrawalsQueued remains the same
        require(
            lstDepositorNonceBefore == StrategyManagerStorage(address(strategyManager)).nonces(lstDepositor),
            "strategyManager.nonces mismatch"
        );
        bytes32 withdrawalsQueuedSlot = keccak256(abi.encode(lstDepositor, withdrawalsQueuedStorageSlot));
        require(
            lstDepositorNumWithdrawalsQueued == uint256(cheats.load(address(strategyManager), withdrawalsQueuedSlot)),
            "strategyManager.numWithdrawalsQueued mismatch"
        );

        // Unpause delegationManager withdrawals
        uint256 paused = IPausable(address(delegation)).paused();
        cheats.prank(IPauserRegistry(IPausable(address(delegation)).pauserRegistry()).unpauser());
        IPausable(address(delegation)).unpause(paused ^ (1 << 2)); // Withdrawal queue on 2nd bit

        // Migrate queued withdrawal to delegationManger
        // Migrating the withdrawal root also verifies that the root has not been erroneously set to false
        IStrategyManager.DeprecatedStruct_QueuedWithdrawal[]
            memory queuedWithdrawals = new IStrategyManager.DeprecatedStruct_QueuedWithdrawal[](1);
        queuedWithdrawals[0] = queuedWithdrawalLst;
        delegation.migrateQueuedWithdrawals(queuedWithdrawals);

        // If successful, confirms that queuedWithdrawal root has not been corrupted between upgrades
        // Queue withdrawal on delegationManager
        IDelegationManager.Withdrawal memory delegationManagerWithdrawal = IDelegationManager.Withdrawal({
            staker: queuedWithdrawalLst.staker,
            delegatedTo: queuedWithdrawalLst.delegatedAddress,
            withdrawer: queuedWithdrawalLst.withdrawerAndNonce.withdrawer,
            nonce: 0, // first withdrawal, so 0 nonce
            startBlock: queuedWithdrawalLst.withdrawalStartBlock,
            strategies: queuedWithdrawalLst.strategies,
            shares: queuedWithdrawalLst.shares
        });
        cheats.prank(lstDepositor);
        delegation.completeQueuedWithdrawal(delegationManagerWithdrawal, tokensToWithdraw, 0, true);

        // Check balances and shares are the same as a withdrawal done pre-upgrade
        uint256 lstDepositorBalancePostUpgrade = tokensToWithdraw[0].balanceOf(lstDepositor);
        uint256 lstDepositorSharesPostUpgrade = strategyManager.stakerStrategyShares(
            lstDepositor,
            delegationManagerWithdrawal.strategies[0]
        );
        require(
            lstDepositorBalancePostUpgrade == lstDepositorBalancePreUpgrade,
            "delegationManager.completeQueuedWithdrawal incorrect post balance"
        );
        require(
            lstDepositorSharesPostUpgrade == lstDepositorSharesPreUpgrade,
            "delegationManager.completeQueuedWithdrawal incorrect post shares"
        );
    }

    function _verifyEigenPodCorrectness() public {
        // Check that state is correct
        require(
            address(eigenPodManager.ownerToPod(eigenPodDepositor)) == address(eigenPod),
            "eigenPodManager.ownerToPod incorrect"
        );
        require(eigenPodManager.hasPod(eigenPodDepositor) == hasPod, "eigenPodManager.hasPod incorrect");
        require(eigenPod.podOwner() == eigenPodOwner, "eigenPod.podOwner incorrect");
        require(
            eigenPod.mostRecentWithdrawalTimestamp() == mostRecentWithdrawalBlock,
            "eigenPod.mostRecentWithdrawalTimestamp incorrect"
        ); // Timestmap replace by block number in storage
        require(!eigenPod.hasRestaked(), "eigenPod.hasRestaked incorrect");

        // Unpause eigenpods verify credentials
        uint256 paused = IPausable(address(eigenPodManager)).paused();
        cheats.prank(IPauserRegistry(IPausable(address(eigenPodManager)).pauserRegistry()).unpauser());
        IPausable(address(eigenPodManager)).unpause(paused ^ (1 << 2)); // eigenpods verify credentials on 2nd bit

        // Get values to check post activating restaking
        uint256 podBalanceBefore = address(eigenPod).balance;
        uint256 userWithdrawalsLength = delayedWithdrawalRouter.userWithdrawalsLength(eigenPodDepositor);

        // Activate restaking and expect emit
        cheats.prank(eigenPodOwner);
        cheats.expectEmit(true, true, true, true);
        emit RestakingActivated(eigenPodOwner);
        eigenPod.activateRestaking();

        // Check updated storage values
        require(eigenPod.hasRestaked(), "eigenPod.hasRestaked not set to true");
        require(address(eigenPod).balance == 0, "eigenPod balance not 0 after activating restaking");
        require(eigenPod.nonBeaconChainETHBalanceWei() == 0, "non beacon chain eth balance not 0");
        require(
            eigenPod.mostRecentWithdrawalTimestamp() == block.timestamp,
            "eigenPod.mostRecentWithdrawalTimestamp not updated"
        );
        require(
            eigenPod.mostRecentWithdrawalTimestamp() > mostRecentWithdrawalBlock,
            "eigenPod.mostRecentWithdrawalTimestamp not updated"
        );

        // Check that delayed withdrawal has been created
        require(
            delayedWithdrawalRouter.userWithdrawalsLength(eigenPodDepositor) == userWithdrawalsLength + 1,
            "delayedWithdrawalRouter.userWithdrawalsLength not incremented"
        );
        IDelayedWithdrawalRouter.DelayedWithdrawal memory delayedWithdrawal = delayedWithdrawalRouter
            .userDelayedWithdrawalByIndex(eigenPodDepositor, userWithdrawalsLength);
        require(delayedWithdrawal.amount == podBalanceBefore, "delayedWithdrawal.amount incorrect");
        require(delayedWithdrawal.blockCreated == block.number, "delayedWithdrawal.blockCreated incorrect");
    }

    function _completeWithdrawalsPreUpgrade() public {
        // Save fork and create new fork
        uint256 forkId = cheats.activeFork();
        cheats.createSelectFork(cheats.envString(rpcUrl), block.number);

        // Complete lstDepositor withdrawal
        cheats.prank(lstDepositor);
        m1StrategyManager(address(strategyManager)).completeQueuedWithdrawal(
            queuedWithdrawalLst,
            tokensToWithdraw,
            0,
            true
        );
        lstDepositorBalancePreUpgrade = tokensToWithdraw[0].balanceOf(lstDepositor);
        lstDepositorSharesPreUpgrade = strategyManager.stakerStrategyShares(
            lstDepositor,
            queuedWithdrawalLst.strategies[0]
        );

        // Reload previous fork
        cheats.selectFork(forkId);
    }

    // Existing LST depositor – ensure that strategy length and shares are all identical
    // Existing LST depositor – ensure that an existing queued withdrawal remains queued
    // Check from stored root, and recalculate root and make sure it matches
    // Check that completing the withdrawal results in the same behavior (same transfer of ERC20 tokens)
    // Check that staker nonce & numWithdrawalsQueued remains the same as before the upgrade
    // Existing LST depositor – queuing a withdrawal before/after the upgrade has the same effects (same decrease in shares, resultant withdrawal root)
    // Existing EigenPod owner – EigenPodManager.ownerToPod remains the same
    // Existing EigenPod owner –  EigenPodManager.hasPod remains the same
    // Existing EigenPod owner –  EigenPod.podOwner remains the same
    // Existing EigenPod owner –  EigenPod.mostRecentWithdrawalTimestamp (after upgrade) == EigenPod.mostRecentWithdrawalBlock (before upgrade)
    // Existing EigenPod owner – EigenPod.hasRestaked remains false
    // Can call EigenPod.activateRestaking and it correctly:
    // Sends all funds in EigenPod (need to make sure it has nonzero balance beforehand)
    // Sets `hasRestaked` to ‘true’
    // Emits a ‘RestakingActivated’ event
    // EigenPod.mostRecentWithdrawalTimestamp updates correctly
    // EigenPod: ethPOS, delayedWithdrawalRouter, eigenPodManager
    event RestakingActivated(address indexed podOwner);
}

interface IDelegationManagerV0 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}

interface m1StrategyManager {
    function withdrawalDelayBlocks() external view returns (uint256);

    function numWithdrawalsQueued(address staker) external view returns (uint256);

    function completeQueuedWithdrawal(
        IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory queuedWithdrawal,
        IERC20[] memory tokensToWithdraw,
        uint256 middlewareTimesIndex,
        bool receiveAsTokens
    ) external;
}

interface m1EigenPod {
    function mostRecentWithdrawalBlockNumber() external view returns (uint64);
}
