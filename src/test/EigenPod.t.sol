// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../contracts/interfaces/IEigenPod.sol";
import "../contracts/pods/DelayedWithdrawalRouter.sol";
import "./utils/ProofParsing.sol";
import "./EigenLayerDeployer.t.sol";
import "../contracts/libraries/BeaconChainProofs.sol";
import "./mocks/BeaconChainOracleMock.sol";
import "./harnesses/EigenPodHarness.sol";

contract EigenPodTests is ProofParsing, EigenPodPausingConstants {
    using BytesLib for bytes;
    using BeaconChainProofs for *;


    uint256 internal constant GWEI_TO_WEI = 1e9;
    uint64 public constant DENEB_FORK_TIMESTAMP_GOERLI = 1705473120;


    bytes pubkey =
        hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";
    uint40 validatorIndex0 = 0;
    uint40 validatorIndex1 = 1;


    address podOwner = address(42000094993494);

    bool public IS_DENEB;

    Vm cheats = Vm(HEVM_ADDRESS);
    DelegationManager public delegation;
    IStrategyManager public strategyManager;
    Slasher public slasher;
    PauserRegistry public pauserReg;

    ProxyAdmin public eigenLayerProxyAdmin;
    IEigenPodManager public eigenPodManager;
    IEigenPod public podImplementation;
    IDelayedWithdrawalRouter public delayedWithdrawalRouter;
    IETHPOSDeposit public ethPOSDeposit;
    UpgradeableBeacon public eigenPodBeacon;
    EPInternalFunctions public podInternalFunctionTester;

    BeaconChainOracleMock public beaconChainOracle;
    address[] public slashingContracts;
    address pauser = address(69);
    address unpauser = address(489);
    address podManagerAddress = 0x212224D2F2d262cd093eE13240ca4873fcCBbA3C;
    address podAddress = address(123);
    uint256 stakeAmount = 32e18;
    mapping(address => bool) fuzzedAddressMapping;
    bytes signature;
    bytes32 depositDataRoot;

    bytes32[] withdrawalFields;
    bytes32[] validatorFields;

    uint32 WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
    IStrategy[] public initializeStrategiesToSetDelayBlocks;
    uint256[] public initializeWithdrawalDelayBlocks;
    uint64 MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32e9;
    uint64 RESTAKED_BALANCE_OFFSET_GWEI = 75e7;
    uint64 internal constant GOERLI_GENESIS_TIME = 1616508000;
    uint64 internal constant SECONDS_PER_SLOT = 12;

    // bytes validatorPubkey = hex"93a0dd04ccddf3f1b419fdebf99481a2182c17d67cf14d32d6e50fc4bf8effc8db4a04b7c2f3a5975c1b9b74e2841888";

    // EIGENPODMANAGER EVENTS
    /// @notice Emitted to notify the update of the beaconChainOracle address
    event BeaconOracleUpdated(address indexed newOracleAddress);

    /// @notice Emitted to notify the deployment of an EigenPod
    event PodDeployed(address indexed eigenPod, address indexed podOwner);

    /// @notice Emitted to notify a deposit of beacon chain ETH recorded in the strategy manager
    event BeaconChainETHDeposited(address indexed podOwner, uint256 amount);

    // EIGENPOD EVENTS
    /// @notice Emitted when an ETH validator stakes via this eigenPod
    event EigenPodStaked(bytes pubkey);

    /// @notice Emitted when an ETH validator's withdrawal credentials are successfully verified to be pointed to this eigenPod
    event ValidatorRestaked(uint40 validatorIndex);

    /// @notice Emitted when an ETH validator's balance is updated in EigenLayer
    event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newBalanceGwei);

    /// @notice Emitted when an ETH validator is prove to have withdrawn from the beacon chain
    event FullWithdrawalRedeemed(
        uint40 validatorIndex,
        uint64 withdrawalTimestamp,
        address indexed recipient,
        uint64 withdrawalAmountGwei
    );

    /// @notice Emitted when a partial withdrawal claim is successfully redeemed
    event PartialWithdrawalRedeemed(
        uint40 validatorIndex,
        uint64 withdrawalTimestamp,
        address indexed recipient,
        uint64 partialWithdrawalAmountGwei
    );

    /// @notice Emitted when restaked beacon chain ETH is withdrawn from the eigenPod.
    event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount);

    // DELAYED WITHDRAWAL ROUTER EVENTS
    /// @notice Emitted when the `withdrawalDelayBlocks` variable is modified from `previousValue` to `newValue`.
    event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue);

    /// @notice event for delayedWithdrawal creation
    event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index);

    /// @notice event for the claiming of delayedWithdrawals
    event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted);

    /// @notice Emitted when ETH that was previously received via the `receive` fallback is withdrawn
    event NonBeaconChainETHWithdrawn(address indexed recipient, uint256 amountWithdrawn);

    modifier fuzzedAddress(address addr) virtual {
        cheats.assume(fuzzedAddressMapping[addr] == false);
        _;
    }

    //performs basic deployment before each test
    function setUp() public {
        // deploy proxy admin for ability to upgrade proxy contracts
        eigenLayerProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserReg = new PauserRegistry(pausers, unpauser);

        /// weird workaround: check commit before this -- the call to `upgradeAndCall` the DelegationManager breaks without performing this step!
        /// seems to be foundry bug. the revert is ultimately for 'TransparentUpgradeableProxy: admin cannot fallback to proxy target', i.e.
        /// the simulated caller is somehow the ProxyAdmin itself.
        EmptyContract emptyContract = new EmptyContract();

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        emptyContract = new EmptyContract();
        delegation = DelegationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        strategyManager = StrategyManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        slasher = Slasher(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        delayedWithdrawalRouter = DelayedWithdrawalRouter(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        cheats.warp(GOERLI_GENESIS_TIME); // start at a sane timestamp
        ethPOSDeposit = new ETHPOSDepositMock();
        podImplementation = new EigenPod(
            ethPOSDeposit,
            delayedWithdrawalRouter,
            IEigenPodManager(podManagerAddress),
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );

        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));

        // this contract is deployed later to keep its address the same (for these tests)
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        DelegationManager delegationImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        StrategyManager strategyManagerImplementation = new StrategyManager(
            delegation,
            IEigenPodManager(podManagerAddress),
            slasher
        );
        Slasher slasherImplementation = new Slasher(strategyManager, delegation);
        EigenPodManager eigenPodManagerImplementation = new EigenPodManager(
            ethPOSDeposit,
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegation
        );

        //ensuring that the address of eigenpodmanager doesn't change
        bytes memory code = address(eigenPodManager).code;
        cheats.etch(podManagerAddress, code);
        eigenPodManager = IEigenPodManager(podManagerAddress);

        beaconChainOracle = new BeaconChainOracleMock();
        DelayedWithdrawalRouter delayedWithdrawalRouterImplementation = new DelayedWithdrawalRouter(
            IEigenPodManager(podManagerAddress)
        );

        address initialOwner = address(this);
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delegation))),
            address(delegationImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                initialOwner,
                pauserReg,
                0 /*initialPausedStatus*/,
                WITHDRAWAL_DELAY_BLOCKS,
                initializeStrategiesToSetDelayBlocks,
                initializeWithdrawalDelayBlocks
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation),
            abi.encodeWithSelector(
                StrategyManager.initialize.selector,
                initialOwner,
                initialOwner,
                pauserReg,
                0 /*initialPausedStatus*/
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation),
            abi.encodeWithSelector(Slasher.initialize.selector, initialOwner, pauserReg, 0 /*initialPausedStatus*/)
        );
        // TODO: add `cheats.expectEmit` calls for initialization events
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                beaconChainOracle,
                initialOwner,
                pauserReg,
                0 /*initialPausedStatus*/
            )
        );
        uint256 initPausedStatus = 0;
        uint256 withdrawalDelayBlocks = WITHDRAWAL_DELAY_BLOCKS;
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))),
            address(delayedWithdrawalRouterImplementation),
            abi.encodeWithSelector(
                DelayedWithdrawalRouter.initialize.selector,
                initialOwner,
                pauserReg,
                initPausedStatus,
                withdrawalDelayBlocks
            )
        );

        cheats.deal(address(podOwner), 5 * stakeAmount);

        fuzzedAddressMapping[address(0)] = true;
        fuzzedAddressMapping[address(eigenLayerProxyAdmin)] = true;
        fuzzedAddressMapping[address(strategyManager)] = true;
        fuzzedAddressMapping[address(eigenPodManager)] = true;
        fuzzedAddressMapping[address(delegation)] = true;
        fuzzedAddressMapping[address(slasher)] = true;
    }

    function testStaking() public {
        cheats.startPrank(podOwner);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
    }

    function testWithdrawBeforeRestaking() public {
        testStaking();
        IEigenPod pod = eigenPodManager.getPod(podOwner);

        //simulate that hasRestaked is set to false, so that we can test withdrawBeforeRestaking for pods deployed before M2 activation
        cheats.store(address(pod), bytes32(uint256(52)), bytes32(uint256(1)));
        require(pod.hasRestaked() == false, "Pod should not be restaked");

        // simulate a withdrawal
        cheats.deal(address(pod), stakeAmount);
        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(delayedWithdrawalRouter));
        emit DelayedWithdrawalCreated(
            podOwner,
            podOwner,
            stakeAmount,
            delayedWithdrawalRouter.userWithdrawalsLength(podOwner)
        );

        uint timestampBeforeTx = pod.mostRecentWithdrawalTimestamp();

        pod.withdrawBeforeRestaking();

        require(_getLatestDelayedWithdrawalAmount(podOwner) == stakeAmount, "Payment amount should be stake amount");
        require(
            pod.mostRecentWithdrawalTimestamp() == uint64(block.timestamp),
            "Most recent withdrawal block number not updated"
        );
        require(
            pod.mostRecentWithdrawalTimestamp() > timestampBeforeTx,
            "Most recent withdrawal block number not updated"
        );
    }

    function testDeployEigenPodWithoutActivateRestaking() public {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(getLatestBlockRoot());

        //this simulates that hasRestaking is set to false, as would be the case for deployed pods that have not yet restaked prior to M2
        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(uint256(0)));

        cheats.startPrank(podOwner);
        cheats.warp(GOERLI_GENESIS_TIME);
        cheats.expectRevert(bytes("EigenPod.hasEnabledRestaking: restaking is not enabled"));
        newPod.verifyWithdrawalCredentials(
            GOERLI_GENESIS_TIME,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        cheats.stopPrank();
    }

    // regression test for bug when activateRestaking -> verifyWC occur in the same epoch, in that order
    function testEigenPodWithdrawalCredentialTimingBug(uint16 epochNum) public {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(getLatestBlockRoot());

        //this simulates that hasRestaking is set to false, as would be the case for deployed pods that have not yet restaked prior to M2
        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(uint256(0)));
        assertTrue(newPod.hasRestaked() == false, "EigenPod should not be restaked");

        uint64 startTime = GOERLI_GENESIS_TIME + (BeaconChainProofs.SECONDS_PER_EPOCH * epochNum);
        uint64 startEpoch = _timestampToEpoch(startTime);
        // move to start time - this is the first slot in the epoch, and is where we will call activateRestaking
        cheats.warp(startTime);

        // activate restaking
        cheats.startPrank(podOwner);
        newPod.activateRestaking();

        // Ensure verifyWC fails for each slot remaining in the epoch
        for (uint i = 0; i < 32; i++) {
            // Move forward 0-31 slots
            uint64 slotTimestamp = uint64(startTime + (BeaconChainProofs.SECONDS_PER_SLOT * i));
            uint64 epoch = _timestampToEpoch(slotTimestamp);
            assertTrue(epoch == startEpoch, "calculated epoch should not change");

            cheats.warp(slotTimestamp);
            cheats.expectRevert(bytes("EigenPod.verifyWithdrawalCredentials: proof must be in the epoch after activation"));
            newPod.verifyWithdrawalCredentials(
                slotTimestamp,
                stateRootProofStruct,
                validatorIndices,
                proofsArray,
                validatorFieldsArray
            );
        }

        // finally, move to the next epoch
        cheats.warp(block.timestamp + BeaconChainProofs.SECONDS_PER_SLOT);
        uint64 endEpoch = _timestampToEpoch(uint64(block.timestamp));
        assertEq(startEpoch + 1, endEpoch, "should have advanced one epoch");

        // now verifyWC should succeed
        newPod.verifyWithdrawalCredentials(
            uint64(block.timestamp),
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        
        cheats.stopPrank();
    }

    function _timestampToEpoch(uint64 timestamp) internal view returns (uint64) {
        require(timestamp >= GOERLI_GENESIS_TIME, "Test._timestampToEpoch: timestamp is before genesis");
        return (timestamp - GOERLI_GENESIS_TIME) / BeaconChainProofs.SECONDS_PER_EPOCH;
    }

    // verifies that it is possible to subsequently prove WCs and withdraw funds when activateRestaking -> validator exit occur in the same epoch, in that order
    // this is similar to `testProveWithdrawalCredentialsAfterValidatorExit` but somewhat more specific / targeted
    function testEigenPodWithdrawalCredentialTimingSuccess() public {
        // upgrade EigenPods to an implementation that has the genesis time to **one slot earlier**
        // we do this so the validator exit at `GOERLI_GENESIS_TIME` is explicitly after genesis time, where we will active restaking
        uint64 modifiedGenesisTime = GOERLI_GENESIS_TIME - 12;
        podImplementation = new EigenPod(
            ethPOSDeposit,
            delayedWithdrawalRouter,
            IEigenPodManager(podManagerAddress),
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            modifiedGenesisTime
        );
        eigenPodBeacon.upgradeTo(address(podImplementation));

        cheats.startPrank(podOwner);
        IEigenPod newPod = IEigenPod(eigenPodManager.createPod());
        cheats.stopPrank();

        //this simulates that hasRestaking is set to false, as would be the case for deployed pods that have not yet restaked prior to M2
        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(uint256(0)));
        require(newPod.hasRestaked() == false, "Pod should not be restaked");

        // move to slot at modifiedGenesisTime. this is within epoch 0, like the withdrawal
        cheats.warp(modifiedGenesisTime);
        cheats.prank(podOwner);
        newPod.activateRestaking();

        // warp forward to start of next epoch (epoch 1), so we can do WC proof with zero effective balance
        cheats.warp(modifiedGenesisTime + BeaconChainProofs.SECONDS_PER_EPOCH);

        // get proof of validator 302913 having WCs pointed with 0 balance
        // ./solidityProofGen  -newBalance=0 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913_exited.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913_exited.json");

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        //set the oracle block root
        _setOracleBlockRoot();

        int256 beaconChainETHSharesBefore = eigenPodManager.podOwnerShares(podOwner);

        cheats.prank(podOwner);
        newPod.verifyWithdrawalCredentials(
            uint64(block.timestamp),
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );

        int256 beaconChainETHSharesAfter = eigenPodManager.podOwnerShares(podOwner);
        require(beaconChainETHSharesBefore == beaconChainETHSharesAfter,
            "effectiveBalance should be zero, no shares should be credited");

        //./solidityProofGen "WithdrawalFieldsProof" 302913 146 8092 true false "data/withdrawal_proof_goerli/goerli_block_header_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6397852.json" "data/withdrawal_proof_goerli/goerli_block_header_6397852.json" "data/withdrawal_proof_goerli/goerli_block_6397852.json" "fullWithdrawalProof_Latest.json" false
        // To get block header: curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v1/beacon/headers/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_header_6399000.json
        // To get block:  curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v2/beacon/blocks/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_6399000.json
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");

        _proveWithdrawalForPod(newPod);
    }

    function testWithdrawNonBeaconChainETHBalanceWei() public {
        IEigenPod pod = testDeployAndVerifyNewEigenPod();

        cheats.deal(address(podOwner), 10 ether);
        emit log_named_address("Pod:", address(pod));

        uint256 balanceBeforeDeposit = pod.nonBeaconChainETHBalanceWei();

        (bool sent, ) = payable(address(pod)).call{value: 1 ether}("");

        require(sent == true, "not sent");

        uint256 balanceAfterDeposit = pod.nonBeaconChainETHBalanceWei();

        require(
            balanceBeforeDeposit < balanceAfterDeposit 
            && (balanceAfterDeposit - balanceBeforeDeposit) == 1 ether, 
            "increment checks"
        );

        cheats.startPrank(podOwner, podOwner);
        cheats.expectEmit(true, true, true, true, address(pod));
        emit NonBeaconChainETHWithdrawn(podOwner, 1 ether);
        pod.withdrawNonBeaconChainETHBalanceWei(
            podOwner,
            1 ether
        );

        uint256 balanceAfterWithdrawal = pod.nonBeaconChainETHBalanceWei();

        require(
            balanceAfterWithdrawal < balanceAfterDeposit 
            && balanceAfterWithdrawal == balanceBeforeDeposit, 
            "decrement checks"
        );

        cheats.stopPrank();
    }

    function testWithdrawFromPod() public {
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        cheats.startPrank(podOwner);

        cheats.expectEmit(true, true, true, true, address(pod));
        emit EigenPodStaked(pubkey);

        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        cheats.deal(address(pod), stakeAmount);

        // this is testing if pods deployed before M2 that do not have hasRestaked initialized to true, will revert
        cheats.store(address(pod), bytes32(uint256(52)), bytes32(uint256(1)));

        cheats.startPrank(podOwner);
        uint256 userWithdrawalsLength = delayedWithdrawalRouter.userWithdrawalsLength(podOwner);
        // cheats.expectEmit(true, true, true, true, address(delayedWithdrawalRouter));
        //cheats.expectEmit(true, true, true, true);
        emit DelayedWithdrawalCreated(podOwner, podOwner, stakeAmount, userWithdrawalsLength);
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();
        require(address(pod).balance == 0, "Pod balance should be 0");
    }

    function testFullWithdrawalProof() public {
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        BeaconChainProofs.WithdrawalProof memory proofs = _getWithdrawalProof();
        bytes32 beaconStateRoot = getBeaconStateRoot();
        withdrawalFields = getWithdrawalFields();
        validatorFields = getValidatorFields();

        Relayer relay = new Relayer();

        relay.verifyWithdrawal(beaconStateRoot, withdrawalFields, proofs);
    }

    function testFullWithdrawalProofWithWrongIndices(
        uint64 wrongBlockRootIndex,
        uint64 wrongWithdrawalIndex,
        uint64 wrongHistoricalSummariesIndex
    ) public {
        uint256 BLOCK_ROOTS_TREE_HEIGHT = 13;
        uint256 WITHDRAWALS_TREE_HEIGHT = 4;
        uint256 HISTORICAL_SUMMARIES_TREE_HEIGHT = 24;
        cheats.assume(wrongBlockRootIndex > 2 ** BLOCK_ROOTS_TREE_HEIGHT);
        cheats.assume(wrongWithdrawalIndex > 2 ** WITHDRAWALS_TREE_HEIGHT);
        cheats.assume(wrongHistoricalSummariesIndex > 2 ** HISTORICAL_SUMMARIES_TREE_HEIGHT);

        Relayer relay = new Relayer();

        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        bytes32 beaconStateRoot = getBeaconStateRoot();
        validatorFields = getValidatorFields();
        withdrawalFields = getWithdrawalFields();

        {
            BeaconChainProofs.WithdrawalProof memory wrongProofs = _getWithdrawalProof();
            wrongProofs.blockRootIndex = wrongBlockRootIndex;
            cheats.expectRevert(bytes("BeaconChainProofs.verifyWithdrawal: blockRootIndex is too large"));
            relay.verifyWithdrawal(beaconStateRoot, withdrawalFields, wrongProofs);
        }

        {
            BeaconChainProofs.WithdrawalProof memory wrongProofs = _getWithdrawalProof();
            wrongProofs.withdrawalIndex = wrongWithdrawalIndex;
            cheats.expectRevert(bytes("BeaconChainProofs.verifyWithdrawal: withdrawalIndex is too large"));
            relay.verifyWithdrawal(beaconStateRoot, withdrawalFields, wrongProofs);
        }

        {
            BeaconChainProofs.WithdrawalProof memory wrongProofs = _getWithdrawalProof();
            wrongProofs.historicalSummaryIndex = wrongHistoricalSummariesIndex;
            cheats.expectRevert(bytes("BeaconChainProofs.verifyWithdrawal: historicalSummaryIndex is too large"));
            relay.verifyWithdrawal(beaconStateRoot, withdrawalFields, wrongProofs);
        }
    }

    /// @notice This test is to ensure the full withdrawal flow works
    function testFullWithdrawalFlowDeneb() public returns (IEigenPod) {
        eigenPodManager.setDenebForkTimestamp(DENEB_FORK_TIMESTAMP_GOERLI);
        IS_DENEB = true;
        //this call is to ensure that validator 302913 has proven their withdrawalcreds
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        //Deneb: ./solidityProofGen/solidityProofGen "WithdrawalFieldsProof" 302913 271 8191 true false "data/deneb_goerli_block_header_7431952.json" "data/deneb_goerli_slot_7431952.json" "data/deneb_goerli_slot_7421952.json" "data/deneb_goerli_block_header_7421951.json" "data/deneb_goerli_block_7421951.json" "fullWithdrawalProof_Latest.json" false false
        // To get block header: curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v1/beacon/headers/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_header_6399000.json
        // To get block:  curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v2/beacon/blocks/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_6399000.json
        setJSON("./src/test/test-data/fullWithdrawalDeneb.json");
        return _proveWithdrawalForPod(newPod);
    }

    function testFullWithdrawalFlowCapellaWithdrawalAgainstDenebRoot() public returns (IEigenPod) {
        IS_DENEB = false;
        //this call is to ensure that validator 302913 has proven their withdrawalcreds
        // ./solidityProofGen/solidityProofGen "WithdrawalFieldsProof" 302913 146 8092 true false "data/deneb_goerli_block_header_7431952.json" "data/deneb_goerli_slot_7431952.json" "data/goerli_slot_6397952.json" "data/goerli_block_header_6397852.json" "data/goerli_block_6397852.json" "fullWithdrawalProof_CapellaAgainstDeneb.json" false true
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        //Deneb: ./solidityProofGen/solidityProofGen "WithdrawalFieldsProof" 302913 271 8191 true false "data/deneb_goerli_block_header_7431952.json" "data/deneb_goerli_slot_7431952.json" "data/deneb_goerli_slot_7421952.json" "data/deneb_goerli_block_header_7421951.json" "data/deneb_goerli_block_7421951.json" "fullWithdrawalProof_Latest.json" false
        // To get block header: curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v1/beacon/headers/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_header_6399000.json
        // To get block:  curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v2/beacon/blocks/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_6399000.json
        setJSON("./src/test/test-data/fullWithdrawalCapellaAgainstDenebRoot.json");
        return _proveWithdrawalForPod(newPod);
    }

    function testFullWithdrawalFlow() public returns (IEigenPod) {
        //this call is to ensure that validator 302913 has proven their withdrawalcreds
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        //./solidityProofGen "WithdrawalFieldsProof" 302913 146 8092 true false "data/withdrawal_proof_goerli/goerli_block_header_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6397852.json" "data/withdrawal_proof_goerli/goerli_block_header_6397852.json" "data/withdrawal_proof_goerli/goerli_block_6397852.json" "fullWithdrawalProof_Latest.json" false
        // To get block header: curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v1/beacon/headers/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_header_6399000.json
        // To get block:  curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v2/beacon/blocks/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_6399000.json
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        return _proveWithdrawalForPod(newPod);
    }

    /**
     * @notice this test is to ensure that a full withdrawal can be made once a validator has processed their first full withrawal
     * This is specifically for the case where a validator has redeposited into their exited validator and needs to prove another withdrawal
     * to get their funds out
     */
    function testWithdrawAfterFullWithdrawal() external {
        _deployInternalFunctionTester();
        IEigenPod pod = testFullWithdrawalFlow();

        // ./solidityProofGen "WithdrawalFieldsProof" 302913 146 8092 true false "data/withdrawal_proof_goerli/goerli_block_header_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6397852.json" "data/withdrawal_proof_goerli/goerli_block_header_6397852.json" "data/withdrawal_proof_goerli/goerli_block_6397852.json" "fullWithdrawalProof_Latest_1SlotAdvanced.json" true
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest_1SlotAdvanced.json");
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(getLatestBlockRoot());

        withdrawalFields = getWithdrawalFields();
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(
            withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]
        );
        uint64 leftOverBalanceWEI = uint64(
            withdrawalAmountGwei - pod.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR()
        ) * uint64(GWEI_TO_WEI);
        cheats.deal(address(pod), leftOverBalanceWEI);
        {
            BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](
                1
            );
            withdrawalProofsArray[0] = _getWithdrawalProof();
            bytes[] memory validatorFieldsProofArray = new bytes[](1);
            validatorFieldsProofArray[0] = abi.encodePacked(getValidatorProof());
            bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
            validatorFieldsArray[0] = getValidatorFields();
            bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
            withdrawalFieldsArray[0] = withdrawalFields;

            BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

            pod.verifyAndProcessWithdrawals(
                0,
                stateRootProofStruct,
                withdrawalProofsArray,
                validatorFieldsProofArray,
                validatorFieldsArray,
                withdrawalFieldsArray
            );
        }
    }

    function testProvingFullWithdrawalForTheSameSlotFails() external {
        IEigenPod pod = testFullWithdrawalFlow();

        {
            BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](
                1
            );
            withdrawalProofsArray[0] = _getWithdrawalProof();
            bytes[] memory validatorFieldsProofArray = new bytes[](1);
            validatorFieldsProofArray[0] = abi.encodePacked(getValidatorProof());
            bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
            validatorFieldsArray[0] = getValidatorFields();
            bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
            withdrawalFieldsArray[0] = withdrawalFields;

            BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

            cheats.expectRevert(
                bytes("EigenPod._verifyAndProcessWithdrawal: withdrawal has already been proven for this timestamp")
            );
            pod.verifyAndProcessWithdrawals(
                0,
                stateRootProofStruct,
                withdrawalProofsArray,
                validatorFieldsProofArray,
                validatorFieldsArray,
                withdrawalFieldsArray
            );
        }
    }

    /// @notice This test is to ensure that the partial withdrawal flow works correctly
    function testPartialWithdrawalFlow() public returns (IEigenPod) {
        //this call is to ensure that validator 61068 has proven their withdrawalcreds
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        //generate partialWithdrawalProofs.json with:
        // ./solidityProofGen "WithdrawalFieldsProof" 302913 146 8092 true true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6397852.json" "data/withdrawal_proof_goerli/goerli_block_header_6397852.json" "data/withdrawal_proof_goerli/goerli_block_6397852.json" "partialWithdrawalProof_Latest.json" false
        setJSON("./src/test/test-data/partialWithdrawalProof_Latest.json");
        withdrawalFields = getWithdrawalFields();
        validatorFields = getValidatorFields();
        BeaconChainProofs.WithdrawalProof memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());

        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(getLatestBlockRoot());
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(
            withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]
        );
        uint40 validatorIndex = uint40(
            Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_INDEX_INDEX])
        );

        cheats.deal(address(newPod), stakeAmount);
        {
            BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](
                1
            );
            withdrawalProofsArray[0] = withdrawalProofs;
            bytes[] memory validatorFieldsProofArray = new bytes[](1);
            validatorFieldsProofArray[0] = validatorFieldsProof;
            bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
            validatorFieldsArray[0] = validatorFields;
            bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
            withdrawalFieldsArray[0] = withdrawalFields;

            uint256 delayedWithdrawalRouterContractBalanceBefore = address(delayedWithdrawalRouter).balance;

            BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

            //cheats.expectEmit(true, true, true, true, address(newPod));
            // cheats.expectEmit(validatorIndex, _computeTimestampAtSlot(Endian.fromLittleEndianUint64(withdrawalProofs.slotRoot)), podOwner, withdrawalAmountGwei, address(newPod));
            emit PartialWithdrawalRedeemed(
                validatorIndex,
                _computeTimestampAtSlot(Endian.fromLittleEndianUint64(withdrawalProofs.slotRoot)),
                podOwner,
                withdrawalAmountGwei
            );
            newPod.verifyAndProcessWithdrawals(
                0,
                stateRootProofStruct,
                withdrawalProofsArray,
                validatorFieldsProofArray,
                validatorFieldsArray,
                withdrawalFieldsArray
            );
            require(
                newPod.provenWithdrawal(
                    validatorFields[0],
                    _computeTimestampAtSlot(Endian.fromLittleEndianUint64(withdrawalProofs.slotRoot))
                ),
                "provenPartialWithdrawal should be true"
            );
            withdrawalAmountGwei = uint64(withdrawalAmountGwei * GWEI_TO_WEI);
            require(
                address(delayedWithdrawalRouter).balance - delayedWithdrawalRouterContractBalanceBefore ==
                    withdrawalAmountGwei,
                "pod delayed withdrawal balance hasn't been updated correctly"
            );
        }

        cheats.roll(block.number + WITHDRAWAL_DELAY_BLOCKS + 1);
        uint256 podOwnerBalanceBefore = address(podOwner).balance;
        delayedWithdrawalRouter.claimDelayedWithdrawals(podOwner, 1);
        require(
            address(podOwner).balance - podOwnerBalanceBefore == withdrawalAmountGwei,
            "Pod owner balance hasn't been updated correctly"
        );
        return newPod;
    }

    /// @notice verifies that multiple partial withdrawals can be made before a full withdrawal
    function testProvingMultiplePartialWithdrawalsForSameSlot() public /*uint256 numPartialWithdrawals*/ {
        IEigenPod newPod = testPartialWithdrawalFlow();

        BeaconChainProofs.WithdrawalProof memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());
        withdrawalFields = getWithdrawalFields();
        validatorFields = getValidatorFields();

        BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](1);
        withdrawalProofsArray[0] = withdrawalProofs;
        bytes[] memory validatorFieldsProofArray = new bytes[](1);
        validatorFieldsProofArray[0] = validatorFieldsProof;
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = validatorFields;
        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
        withdrawalFieldsArray[0] = withdrawalFields;

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        cheats.expectRevert(
            bytes("EigenPod._verifyAndProcessWithdrawal: withdrawal has already been proven for this timestamp")
        );
        newPod.verifyAndProcessWithdrawals(
            0,
            stateRootProofStruct,
            withdrawalProofsArray,
            validatorFieldsProofArray,
            validatorFieldsArray,
            withdrawalFieldsArray
        );
    }

    /// @notice verifies that multiple full withdrawals for a single validator fail
    function testDoubleFullWithdrawal() public returns (IEigenPod newPod) {
        newPod = testFullWithdrawalFlow();
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(
            withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]
        );
        uint64 leftOverBalanceWEI = uint64(withdrawalAmountGwei - newPod.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR()) *
            uint64(GWEI_TO_WEI);
        cheats.deal(address(newPod), leftOverBalanceWEI);

        BeaconChainProofs.WithdrawalProof memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());
        withdrawalFields = getWithdrawalFields();
        validatorFields = getValidatorFields();

        BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](1);
        withdrawalProofsArray[0] = withdrawalProofs;
        bytes[] memory validatorFieldsProofArray = new bytes[](1);
        validatorFieldsProofArray[0] = validatorFieldsProof;
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = validatorFields;
        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
        withdrawalFieldsArray[0] = withdrawalFields;

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        cheats.expectRevert(
            bytes("EigenPod._verifyAndProcessWithdrawal: withdrawal has already been proven for this timestamp")
        );
        newPod.verifyAndProcessWithdrawals(
            0,
            stateRootProofStruct,
            withdrawalProofsArray,
            validatorFieldsProofArray,
            validatorFieldsArray,
            withdrawalFieldsArray
        );

        return newPod;
    }

    function testDeployAndVerifyNewEigenPod() public returns (IEigenPod) {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        return _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
    }

    // test freezing operator after a beacon chain slashing event
    function testUpdateSlashedBeaconBalance() public {
        _deployInternalFunctionTester();
        //make initial deposit
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.warp(block.timestamp + 1);
        // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_updated_to_0ETH_302913.json");
        _proveOverCommittedStake(newPod);

        uint64 newValidatorBalance = _getValidatorUpdatedBalance(); 
        int256 beaconChainETHShares = eigenPodManager.podOwnerShares(podOwner);

        require(
            beaconChainETHShares == int256((newValidatorBalance) * GWEI_TO_WEI),
            "eigenPodManager shares not updated correctly"
        );
    }
    
    /// @notice Similar test done in EP unit test
    //test deploying an eigen pod with mismatched withdrawal credentials between the proof and the actual pod's address
    function testDeployNewEigenPodWithWrongWithdrawalCreds(address wrongWithdrawalAddress) public {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        cheats.startPrank(podOwner);
        IEigenPod newPod;
        newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();


        // make sure that wrongWithdrawalAddress is not set to actual pod address
        cheats.assume(wrongWithdrawalAddress != address(newPod));

        validatorFields = getValidatorFields();
        validatorFields[1] = abi.encodePacked(bytes1(uint8(1)), bytes11(0), wrongWithdrawalAddress).toBytes32(0);

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = validatorFields;
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(validatorIndex0);

        cheats.startPrank(podOwner);
        if (!newPod.hasRestaked()) {
            newPod.activateRestaking();
        }
        // set oracle block root
        _setOracleBlockRoot();

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        uint64 timestamp = _verifyWCStartTimestamp(newPod);
        cheats.warp(timestamp);

        cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: Proof is not for this EigenPod"));
        newPod.verifyWithdrawalCredentials(
            timestamp,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        cheats.stopPrank();
    }
    //ensures that a validator proving WC after they have exited the beacon chain is allowed to
    //prove their WC and process a withdrawal
    function testProveWithdrawalCredentialsAfterValidatorExit() public {
        // ./solidityProofGen  -newBalance=0 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913_exited.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913_exited.json");
               emit log("hello");

        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        //./solidityProofGen "WithdrawalFieldsProof" 302913 146 8092 true false "data/withdrawal_proof_goerli/goerli_block_header_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6397852.json" "data/withdrawal_proof_goerli/goerli_block_header_6397852.json" "data/withdrawal_proof_goerli/goerli_block_6397852.json" "fullWithdrawalProof_Latest.json" false
        // To get block header: curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v1/beacon/headers/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_header_6399000.json
        // To get block:  curl -H "Accept: application/json" 'https://eigenlayer.spiceai.io/goerli/beacon/eth/v2/beacon/blocks/6399000?api_key\="343035|f6ebfef661524745abb4f1fd908a76e8"' > block_6399000.json
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        _proveWithdrawalForPod(newPod);
    }

    function testVerifyWithdrawalCredsFromNonPodOwnerAddress(address nonPodOwnerAddress) public {
        // nonPodOwnerAddress must be different from podOwner
        cheats.assume(nonPodOwnerAddress != podOwner);
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        cheats.startPrank(podOwner);

        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);

        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(validatorIndex0);

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        uint64 timestamp = _verifyWCStartTimestamp(newPod);
        cheats.warp(timestamp);

        cheats.startPrank(nonPodOwnerAddress);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodOwner: not podOwner"));
        newPod.verifyWithdrawalCredentials(
            timestamp,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        cheats.stopPrank();
    }

    function testBalanceProofWithWrongTimestamp(uint64 timestamp) public {
        cheats.assume(timestamp > GOERLI_GENESIS_TIME);
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        IEigenPod newPod = testDeployAndVerifyNewEigenPod();

         // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_updated_to_0ETH_302913.json");
        // prove overcommitted balance
        cheats.warp(timestamp);
        _proveOverCommittedStake(newPod);


        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        bytes memory proof = abi.encodePacked(getBalanceUpdateProof());
        bytes[] memory proofs = new bytes[](1);
        proofs[0] = proof;

        bytes32 newLatestBlockRoot = getLatestBlockRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();      

        cheats.expectRevert(bytes("EigenPod.verifyBalanceUpdate: Validators balance has already been updated for this timestamp"));
        newPod.verifyBalanceUpdates(uint64(block.timestamp - 1), validatorIndices, stateRootProofStruct, proofs, validatorFieldsArray);
    }

    // // 3. Single withdrawal credential
    // // Test: Owner proves an withdrawal credential.
    // //                     validator status should be marked as ACTIVE

    function testProveSingleWithdrawalCredential() public {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();

        assertTrue(
            pod.validatorStatus(validatorPubkeyHash) == IEigenPod.VALIDATOR_STATUS.ACTIVE,
            "wrong validator status"
        );
    }

    // 5. Prove overcommitted balance
    // Setup: Run (3).
    // Test: Watcher proves an overcommitted balance for validator from (3).
    //                     validator status should be marked as OVERCOMMITTED

    function testProveOverCommittedBalance() public {
        _deployInternalFunctionTester();
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        // get beaconChainETH shares
        int256 beaconChainETHBefore = eigenPodManager.podOwnerShares(podOwner);

        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();
        uint256 validatorRestakedBalanceBefore = newPod
            .validatorPubkeyHashToInfo(validatorPubkeyHash)
            .restakedBalanceGwei;

        // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_updated_to_0ETH_302913.json");
        // prove overcommitted balance
        cheats.warp(block.timestamp + 1);
        _proveOverCommittedStake(newPod);

        uint256 validatorRestakedBalanceAfter = newPod
            .validatorPubkeyHashToInfo(validatorPubkeyHash)
            .restakedBalanceGwei;

        uint64 newValidatorBalance = _getValidatorUpdatedBalance();
        int256 shareDiff = beaconChainETHBefore - eigenPodManager.podOwnerShares(podOwner);
        assertTrue(
            eigenPodManager.podOwnerShares(podOwner) ==
                int256(newValidatorBalance * GWEI_TO_WEI),
            "hysterisis not working"
        );
        assertTrue(
            beaconChainETHBefore - eigenPodManager.podOwnerShares(podOwner) == shareDiff,
            "BeaconChainETHShares not updated"
        );
        assertTrue(
            int256(validatorRestakedBalanceBefore) - int256(validatorRestakedBalanceAfter) ==
                shareDiff / int256(GWEI_TO_WEI),
            "validator restaked balance not updated"
        );
    }

    function testVerifyUndercommittedBalance() public {
        _deployInternalFunctionTester();
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        // get beaconChainETH shares
        int256 beaconChainETHBefore = eigenPodManager.podOwnerShares(podOwner);
        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();
        uint256 validatorRestakedBalanceBefore = newPod
            .validatorPubkeyHashToInfo(validatorPubkeyHash)
            .restakedBalanceGwei;

        // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_updated_to_0ETH_302913.json");
        // prove overcommitted balance
        cheats.warp(block.timestamp + 1);
        _proveOverCommittedStake(newPod);

        cheats.warp(block.timestamp + 1);
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 100 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913_incrementedBlockBy100.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913_incrementedBlockBy100.json");
        _proveUnderCommittedStake(newPod);

        uint256 validatorRestakedBalanceAfter = newPod
            .validatorPubkeyHashToInfo(validatorPubkeyHash)
            .restakedBalanceGwei;

        int256 shareDiff = beaconChainETHBefore - eigenPodManager.podOwnerShares(podOwner);

        assertTrue(
            eigenPodManager.podOwnerShares(podOwner) ==
                int256((MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR) * GWEI_TO_WEI),
            "hysterisis not working"
        );
        assertTrue(
            beaconChainETHBefore - eigenPodManager.podOwnerShares(podOwner) == shareDiff,
            "BeaconChainETHShares not updated"
        );
        assertTrue(
            int256(uint256(validatorRestakedBalanceBefore)) - int256(uint256(validatorRestakedBalanceAfter)) ==
                shareDiff / int256(GWEI_TO_WEI),
            "validator restaked balance not updated"
        );
    }

    function testDeployingEigenPodRevertsWhenPaused() external {
        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_NEW_EIGENPODS);
        cheats.stopPrank();

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
    }

    function testCreatePodWhenPaused() external {
        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_NEW_EIGENPODS);
        cheats.stopPrank();

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        eigenPodManager.createPod();
        cheats.stopPrank();
    }

    function testCreatePodIfItReturnsPodAddress() external {
        cheats.startPrank(podOwner);
        address _podAddress = eigenPodManager.createPod();
        cheats.stopPrank();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(_podAddress == address(pod), "invalid pod address");
    }

    function testStakeOnEigenPodFromNonPodManagerAddress(address nonPodManager) external fuzzedAddress(nonPodManager) {
        cheats.assume(nonPodManager != address(eigenPodManager));

        cheats.startPrank(podOwner);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        cheats.deal(nonPodManager, stakeAmount);

        cheats.startPrank(nonPodManager);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodManager: not eigenPodManager"));
        newPod.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
    }

    function testCallWithdrawBeforeRestakingFromNonOwner(address nonPodOwner) external fuzzedAddress(nonPodOwner) {
        cheats.assume(nonPodOwner != podOwner);
        testStaking();
        IEigenPod pod = eigenPodManager.getPod(podOwner);

        // this is testing if pods deployed before M2 that do not have hasRestaked initialized to true, will revert
        cheats.store(address(pod), bytes32(uint256(52)), bytes32(0));
        require(pod.hasRestaked() == false, "Pod should not be restaked");

        //simulate a withdrawal
        cheats.startPrank(nonPodOwner);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodOwner: not podOwner"));
        pod.withdrawBeforeRestaking();
    }

    /* test deprecated since this is checked on the EigenPodManager level, rather than the EigenPod level
    TODO: @Sidu28 - check whether we have adequate coverage of the correct function
    function testWithdrawRestakedBeaconChainETHRevertsWhenPaused() external {
        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_WITHDRAW_RESTAKED_ETH);
        cheats.stopPrank();

        address recipient = address(this);
        uint256 amount = 1e18;
        IEigenPod eigenPod = eigenPodManager.getPod(podOwner);
        cheats.startPrank(address(eigenPodManager));
        cheats.expectRevert(bytes("Pausable: index is paused"));
        eigenPod.withdrawRestakedBeaconChainETH(recipient, amount);
        cheats.stopPrank();
    }
    */

    function testVerifyCorrectWithdrawalCredentialsRevertsWhenPaused() external {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newBeaconStateRoot);

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_CREDENTIALS);
        cheats.stopPrank();

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        uint64 timestamp = _verifyWCStartTimestamp(newPod);
        cheats.warp(timestamp);

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        newPod.verifyWithdrawalCredentials(
            timestamp,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        cheats.stopPrank();
    }

    function testVerifyOvercommittedStakeRevertsWhenPaused() external {
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_updated_to_0ETH_302913.json");
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        bytes[] memory proofs = new bytes[](1);
        proofs[0] = abi.encodePacked(getBalanceUpdateProof());

        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newBeaconStateRoot);
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE);
        cheats.stopPrank();

        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        newPod.verifyBalanceUpdates(0, validatorIndices, stateRootProofStruct, proofs, validatorFieldsArray);
    }

    function _proveOverCommittedStake(IEigenPod newPod) internal {
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        bytes[] memory proofs = new bytes[](1);
        proofs[0] = abi.encodePacked(getBalanceUpdateProof());

        bytes32 newLatestBlockRoot = getLatestBlockRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        newPod.verifyBalanceUpdates(
            uint64(block.timestamp),
            validatorIndices,
            stateRootProofStruct,
            proofs,
            validatorFieldsArray
        );
    }

    function _proveUnderCommittedStake(IEigenPod newPod) internal {
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        bytes[] memory proofs = new bytes[](1);
        proofs[0] = abi.encodePacked(getBalanceUpdateProof());

        bytes32 newLatestBlockRoot = getLatestBlockRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        newPod.verifyBalanceUpdates(
            uint64(block.timestamp),
            validatorIndices,
            stateRootProofStruct,
            proofs,
            validatorFieldsArray
        );
        require(newPod.validatorPubkeyHashToInfo(getValidatorPubkeyHash()).status == IEigenPod.VALIDATOR_STATUS.ACTIVE);
    }

    function _getValidatorUpdatedBalance() internal returns (uint64) {
        bytes32[] memory validatorFieldsToGet = getValidatorFields();
        return validatorFieldsToGet.getEffectiveBalanceGwei();
    }

    function testStake(bytes calldata _pubkey, bytes calldata _signature, bytes32 _depositDataRoot) public {
        // should fail if no/wrong value is provided
        cheats.startPrank(podOwner);
        cheats.expectRevert("EigenPod.stake: must initially stake for any validator with 32 ether");
        eigenPodManager.stake(_pubkey, _signature, _depositDataRoot);
        cheats.expectRevert("EigenPod.stake: must initially stake for any validator with 32 ether");
        eigenPodManager.stake{value: 12 ether}(_pubkey, _signature, _depositDataRoot);

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        // successful call
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(_pubkey);
        eigenPodManager.stake{value: 32 ether}(_pubkey, _signature, _depositDataRoot);
        cheats.stopPrank();
    }

    /// @notice Test that the Merkle proof verification fails when the proof length is 0
    function testVerifyInclusionSha256FailsForEmptyProof(bytes32 root, bytes32 leaf, uint256 index) public {
        bytes memory emptyProof = new bytes(0);
        cheats.expectRevert(
            bytes("Merkle.processInclusionProofSha256: proof length should be a non-zero multiple of 32")
        );
        Merkle.verifyInclusionSha256(emptyProof, root, leaf, index);
    }

    /// @notice Test that the Merkle proof verification fails when the proof length is not a multple of 32
    function testVerifyInclusionSha256FailsForNonMultipleOf32ProofLength(
        bytes32 root,
        bytes32 leaf,
        uint256 index,
        bytes memory proof
    ) public {
        cheats.assume(proof.length % 32 != 0);
        cheats.expectRevert(
            bytes("Merkle.processInclusionProofSha256: proof length should be a non-zero multiple of 32")
        );
        Merkle.verifyInclusionSha256(proof, root, leaf, index);
    }

    // verifies that the `numPod` variable increments correctly on a succesful call to the `EigenPod.stake` function
    function test_incrementNumPodsOnStake(
        bytes calldata _pubkey,
        bytes calldata _signature,
        bytes32 _depositDataRoot
    ) public {
        uint256 numPodsBefore = EigenPodManager(address(eigenPodManager)).numPods();
        testStake(_pubkey, _signature, _depositDataRoot);
        uint256 numPodsAfter = EigenPodManager(address(eigenPodManager)).numPods();
        require(numPodsAfter == numPodsBefore + 1, "numPods did not increment correctly");
    }

    // verifies that the `numPod` variable increments correctly on a succesful call to the `EigenPod.createPod` function
    function test_incrementNumPodsOnCreatePod() public {
        uint256 numPodsBefore = EigenPodManager(address(eigenPodManager)).numPods();
        eigenPodManager.createPod();
        uint256 numPodsAfter = EigenPodManager(address(eigenPodManager)).numPods();
        require(numPodsAfter == numPodsBefore + 1, "numPods did not increment correctly");
    }

    function test_createPodTwiceFails() public {
        eigenPodManager.createPod();
        cheats.expectRevert(bytes("EigenPodManager.createPod: Sender already has a pod"));
        eigenPodManager.createPod();
    }

    function test_validatorPubkeyToInfo() external {
        bytes memory _pubkey = hex"93a0dd04ccddf3f1b419fdebf99481a2182c17d67cf14d32d6e50fc4bf8effc8db4a04b7c2f3a5975c1b9b74e2841888";

        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod pod = eigenPodManager.getPod(podOwner);

        IEigenPod.ValidatorInfo memory info1 = pod.validatorPubkeyToInfo(_pubkey);
        IEigenPod.ValidatorInfo memory info2 = pod.validatorPubkeyHashToInfo(getValidatorPubkeyHash());

        require(info1.validatorIndex == info2.validatorIndex, "validatorIndex does not match");
        require(info1.restakedBalanceGwei > 0, "restakedBalanceGwei is 0");
        require(info1.restakedBalanceGwei == info2.restakedBalanceGwei, "restakedBalanceGwei does not match");
        require(info1.mostRecentBalanceUpdateTimestamp == info2.mostRecentBalanceUpdateTimestamp, "mostRecentBalanceUpdateTimestamp does not match");
        require(info1.status == info2.status, "status does not match");
    }

    function test_validatorStatus() external {
        bytes memory _pubkey = hex"93a0dd04ccddf3f1b419fdebf99481a2182c17d67cf14d32d6e50fc4bf8effc8db4a04b7c2f3a5975c1b9b74e2841888";

        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod pod = eigenPodManager.getPod(podOwner);

        IEigenPod.VALIDATOR_STATUS status1 = pod.validatorStatus(_pubkey);
        IEigenPod.VALIDATOR_STATUS status2 = pod.validatorStatus(getValidatorPubkeyHash());

        require(status1 == status2, "status does not match");
    }

    /* TODO: reimplement similar tests
    function testQueueBeaconChainETHWithdrawalWithoutProvingFullWithdrawal() external {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        uint256 shareAmount = 32e18;
        // expect revert from underflow
        cheats.expectRevert();
        _testQueueWithdrawal(podOwner, shareAmount);
    }

    function testQueueBeaconChainETHWithdrawal() external {
        IEigenPod pod = testFullWithdrawalFlow();

        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();

        uint256 withdrawableRestakedExecutionLayerGweiBefore = pod.withdrawableRestakedExecutionLayerGwei();
        
        uint256 shareAmount = (pod.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR()) * GWEI_TO_WEI;
        _verifyEigenPodBalanceSharesInvariant(podOwner, pod, validatorPubkeyHash);
        _testQueueWithdrawal(podOwner, shareAmount);
        _verifyEigenPodBalanceSharesInvariant(podOwner, pod, validatorPubkeyHash);

        require(withdrawableRestakedExecutionLayerGweiBefore - pod.withdrawableRestakedExecutionLayerGwei() == shareAmount/int256(GWEI_TO_WEI),
            "withdrawableRestakedExecutionLayerGwei not decremented correctly");
    }
*/
    function _verifyEigenPodBalanceSharesInvariant(
        address podowner,
        IEigenPod pod,
        bytes32 validatorPubkeyHash
    ) internal view {
        int256 shares = eigenPodManager.podOwnerShares(podowner);
        uint64 withdrawableRestakedExecutionLayerGwei = pod.withdrawableRestakedExecutionLayerGwei();

        EigenPod.ValidatorInfo memory info = pod.validatorPubkeyHashToInfo(validatorPubkeyHash);

        uint64 validatorBalanceGwei = info.restakedBalanceGwei;
        require(
            shares / int256(GWEI_TO_WEI) ==
                int256(uint256(validatorBalanceGwei)) + int256(uint256(withdrawableRestakedExecutionLayerGwei)),
            "EigenPod invariant violated: sharesInSM != withdrawableRestakedExecutionLayerGwei"
        );
    }

    function _proveWithdrawalForPod(IEigenPod newPod) internal returns (IEigenPod) {
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(getLatestBlockRoot());
        uint64 restakedExecutionLayerGweiBefore = newPod.withdrawableRestakedExecutionLayerGwei();

        withdrawalFields = getWithdrawalFields();
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(
            withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]
        );
        emit log_named_uint("withdrawalAmountGwei", withdrawalAmountGwei);
        uint64 leftOverBalanceWEI = uint64(withdrawalAmountGwei - newPod.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR()) *
            uint64(GWEI_TO_WEI);
        cheats.deal(address(newPod), leftOverBalanceWEI);
        emit log_named_uint("leftOverBalanceWEI", leftOverBalanceWEI);
        emit log_named_uint("address(newPod)", address(newPod).balance);
        emit log_named_uint("withdrawalAmountGwei", withdrawalAmountGwei);

        uint256 delayedWithdrawalRouterContractBalanceBefore = address(delayedWithdrawalRouter).balance;
        {
            BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](
                1
            );
            withdrawalProofsArray[0] = _getWithdrawalProof();
            bytes[] memory validatorFieldsProofArray = new bytes[](1);
            validatorFieldsProofArray[0] = abi.encodePacked(getValidatorProof());
            bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
            validatorFieldsArray[0] = getValidatorFields();
            bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
            withdrawalFieldsArray[0] = withdrawalFields;

            BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

            newPod.verifyAndProcessWithdrawals(
                0,
                stateRootProofStruct,
                withdrawalProofsArray,
                validatorFieldsProofArray,
                validatorFieldsArray,
                withdrawalFieldsArray
            );
        }
        require(
            newPod.withdrawableRestakedExecutionLayerGwei() - restakedExecutionLayerGweiBefore ==
                newPod.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR(),
            "restakedExecutionLayerGwei has not been incremented correctly"
        );
        require(
            address(delayedWithdrawalRouter).balance - delayedWithdrawalRouterContractBalanceBefore ==
                leftOverBalanceWEI,
            "pod delayed withdrawal balance hasn't been updated correctly"
        );
        require(
            newPod.validatorPubkeyHashToInfo(getValidatorPubkeyHash()).restakedBalanceGwei == 0,
            "balance not reset correctly"
        );

        cheats.roll(block.number + WITHDRAWAL_DELAY_BLOCKS + 1);
        uint256 podOwnerBalanceBefore = address(podOwner).balance;
        delayedWithdrawalRouter.claimDelayedWithdrawals(podOwner, 1);
        require(
            address(podOwner).balance - podOwnerBalanceBefore == leftOverBalanceWEI,
            "Pod owner balance hasn't been updated correctly"
        );
        return newPod;
    }

    // simply tries to register 'sender' as a delegate, setting their 'DelegationTerms' contract in DelegationManager to 'dt'
    // verifies that the storage of DelegationManager contract is updated appropriately
    function _testRegisterAsOperator(
        address sender,
        IDelegationManager.OperatorDetails memory operatorDetails
    ) internal {
        cheats.startPrank(sender);
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        assertTrue(delegation.isOperator(sender), "testRegisterAsOperator: sender is not a delegate");

        // TODO: FIX THIS
        // assertTrue(
        //     delegation.delegationTerms(sender) == dt, "_testRegisterAsOperator: delegationTerms not set appropriately"
        // );

        assertTrue(delegation.isDelegated(sender), "_testRegisterAsOperator: sender not marked as actively delegated");
        cheats.stopPrank();
    }

    function _testDelegateToOperator(address sender, address operator) internal {
        //delegator-specific information
        (IStrategy[] memory delegateStrategies, uint256[] memory delegateShares) = strategyManager.getDeposits(sender);

        uint256 numStrats = delegateShares.length;
        assertTrue(numStrats > 0, "_testDelegateToOperator: delegating from address with no deposits");
        uint256[] memory inititalSharesInStrats = new uint256[](numStrats);
        for (uint256 i = 0; i < numStrats; ++i) {
            inititalSharesInStrats[i] = delegation.operatorShares(operator, delegateStrategies[i]);
        }

        cheats.startPrank(sender);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(operator, signatureWithExpiry, bytes32(0));
        cheats.stopPrank();

        assertTrue(
            delegation.delegatedTo(sender) == operator,
            "_testDelegateToOperator: delegated address not set appropriately"
        );
        assertTrue(delegation.isDelegated(sender), "_testDelegateToOperator: delegated status not set appropriately");

        for (uint256 i = 0; i < numStrats; ++i) {
            uint256 operatorSharesBefore = inititalSharesInStrats[i];
            uint256 operatorSharesAfter = delegation.operatorShares(operator, delegateStrategies[i]);
            assertTrue(
                operatorSharesAfter == (operatorSharesBefore + delegateShares[i]),
                "_testDelegateToOperator: delegatedShares not increased correctly"
            );
        }
    }

    function _testDelegation(address operator, address staker) internal {
        if (!delegation.isOperator(operator)) {
            IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
                __deprecated_earningsReceiver: operator,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            });
            _testRegisterAsOperator(operator, operatorDetails);
        }

        //making additional deposits to the strategies
        assertTrue(!delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");
        _testDelegateToOperator(staker, operator);
        assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");

        IStrategy[] memory updatedStrategies;
        uint256[] memory updatedShares;
        (updatedStrategies, updatedShares) = strategyManager.getDeposits(staker);
    }

    function _testDeployAndVerifyNewEigenPod(
        address _podOwner,
        bytes memory _signature,
        bytes32 _depositDataRoot
    ) internal returns (IEigenPod) {
        IEigenPod newPod = eigenPodManager.getPod(_podOwner);

        cheats.startPrank(_podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, _signature, _depositDataRoot);
        cheats.stopPrank();

        return _verifyWithdrawalCredentials(newPod, _podOwner);
    }

    function _verifyWithdrawalCredentials(IEigenPod newPod, address _podOwner) internal returns (IEigenPod) {
        _deployInternalFunctionTester();
        uint64 timestamp = EigenPod(payable(address(newPod))).GENESIS_TIME();
        // cheats.expectEmit(true, true, true, true, address(newPod));
        // emit ValidatorRestaked(validatorIndex);

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        int256 beaconChainETHSharesBefore = eigenPodManager.podOwnerShares(_podOwner);

        cheats.startPrank(_podOwner);
        cheats.warp(timestamp);
        if (newPod.hasRestaked() == false) {
            newPod.activateRestaking();
        }
        //set the oracle block root
        _setOracleBlockRoot();

        emit log_named_bytes32(
            "restaking activated",
            BeaconChainOracleMock(address(beaconChainOracle)).mockBeaconChainStateRoot()
        );

        timestamp = _verifyWCStartTimestamp(newPod);
        cheats.warp(timestamp);
        newPod.verifyWithdrawalCredentials(
            timestamp,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        cheats.stopPrank();

        int256 beaconChainETHSharesAfter = eigenPodManager.podOwnerShares(_podOwner);
        uint256 effectiveBalance = uint256(Endian.fromLittleEndianUint64(validatorFieldsArray[0][2]) * GWEI_TO_WEI);

        emit log_named_uint("effectiveBalance", effectiveBalance);
        emit log_named_uint("MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR", MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR * GWEI_TO_WEI);
        emit log_named_uint("beaconChainETHSharesAfter", uint256(beaconChainETHSharesAfter));
        emit log_named_uint("beaconChainETHSharesBefore", uint256(beaconChainETHSharesBefore));

        if(effectiveBalance < MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR * GWEI_TO_WEI){
            require(
                (beaconChainETHSharesAfter - beaconChainETHSharesBefore) == int256(effectiveBalance),
                "eigenPodManager shares not updated correctly"
             );
        } else {
            emit log("here)");
            require(
                (beaconChainETHSharesAfter - beaconChainETHSharesBefore) == int256(uint256(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR * GWEI_TO_WEI)),
                "eigenPodManager shares not updated correctly"
            );
            emit log("here)");

        }
        return newPod;
    }

    function _verifyWCStartTimestamp(IEigenPod pod) internal  returns (uint64) {
        uint64 genesis = EigenPod(payable(address(pod))).GENESIS_TIME();
        uint64 activateRestakingTimestamp = pod.mostRecentWithdrawalTimestamp();

        // For pods deployed after M2, `mostRecentWithdrawalTimestamp` will always be 0
        // In order to give our `_nextEpochStartTimestamp` a sane calculation, we set it
        // to the genesis time
        if (activateRestakingTimestamp == 0) {
            activateRestakingTimestamp = uint64(block.timestamp);
        }

        emit log_named_uint("genesis", genesis);
        emit log_named_uint("activation", activateRestakingTimestamp);

        return _nextEpochStartTimestamp(genesis, activateRestakingTimestamp);
    }

    /// @dev Given a genesis time and a timestamp, converts the timestamp to an epoch
    /// then calculates the next epoch's start timestamp. Used for the verifyWC proof window.
    function _nextEpochStartTimestamp(uint64 genesisTime, uint64 timestamp) internal pure returns (uint64) {
        require(timestamp >= genesisTime, "_verifyWCStartTimestamp: timestamp is before genesis");
        uint64 epoch = (timestamp - genesisTime) / BeaconChainProofs.SECONDS_PER_EPOCH;

        return genesisTime + ((1 + epoch) * BeaconChainProofs.SECONDS_PER_EPOCH);
    }

    /* TODO: reimplement similar tests
    function _testQueueWithdrawal(
        address _podOwner,
        uint256 amountWei
    )
        internal
        returns (bytes32)
    {
        //make a call from _podOwner to queue the withdrawal
        cheats.startPrank(_podOwner);
        bytes32 withdrawalRoot = eigenPodManager.queueWithdrawal(
            amountWei,
            _podOwner
        );
        cheats.stopPrank();
        return withdrawalRoot;
    }
*/
    function _getLatestDelayedWithdrawalAmount(address recipient) internal view returns (uint256) {
        return
            delayedWithdrawalRouter
                .userDelayedWithdrawalByIndex(recipient, delayedWithdrawalRouter.userWithdrawalsLength(recipient) - 1)
                .amount;
    }

    function _getStateRootProof() internal returns (BeaconChainProofs.StateRootProof memory) {
        return BeaconChainProofs.StateRootProof(getBeaconStateRoot(), abi.encodePacked(getStateRootProof()));
    }

    /// @notice this function just generates a valid proof so that we can test other functionalities of the withdrawal flow
    function _getWithdrawalProof() internal returns (BeaconChainProofs.WithdrawalProof memory) {
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        //make initial deposit
        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        if(!IS_DENEB){
            emit log("NOT DENEB");
        }
        bytes memory withdrawalProof = IS_DENEB ? abi.encodePacked(getWithdrawalProofDeneb()) : abi.encodePacked(getWithdrawalProofCapella());
        bytes memory timestampProof = IS_DENEB ? abi.encodePacked(getTimestampProofDeneb()) : abi.encodePacked(getTimestampProofCapella());
        {
            bytes32 blockRoot = getBlockRoot();
            bytes32 slotRoot = getSlotRoot();
            bytes32 timestampRoot = getTimestampRoot();
            bytes32 executionPayloadRoot = getExecutionPayloadRoot();
            return
                BeaconChainProofs.WithdrawalProof(
                    abi.encodePacked(withdrawalProof),
                    abi.encodePacked(getSlotProof()),
                    abi.encodePacked(getExecutionPayloadProof()),
                    abi.encodePacked(timestampProof),
                    abi.encodePacked(getHistoricalSummaryProof()),
                    uint64(getBlockRootIndex()),
                    uint64(getHistoricalSummaryIndex()),
                    uint64(getWithdrawalIndex()),
                    blockRoot,
                    slotRoot,
                    timestampRoot,
                    executionPayloadRoot
                );
        }
    }
    

    function _setOracleBlockRoot() internal {
        bytes32 latestBlockRoot = getLatestBlockRoot();
        //set beaconStateRoot
        beaconChainOracle.setOracleBlockRootAtTimestamp(latestBlockRoot);
    }

    function _computeTimestampAtSlot(uint64 slot) internal pure returns (uint64) {
        return uint64(GOERLI_GENESIS_TIME + slot * SECONDS_PER_SLOT);
    }

    function _deployInternalFunctionTester() internal {
        podInternalFunctionTester = new EPInternalFunctions(
            ethPOSDeposit,
            delayedWithdrawalRouter,
            IEigenPodManager(podManagerAddress),
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );
    }
}

contract Relayer is Test {
    function verifyWithdrawal(
        bytes32 beaconStateRoot,
        bytes32[] calldata withdrawalFields,
        BeaconChainProofs.WithdrawalProof calldata proofs
    ) public view {
        BeaconChainProofs.verifyWithdrawal(beaconStateRoot, withdrawalFields, proofs, type(uint64).max);
    }
}


//TODO: Integration Tests from old EPM unit tests:
    // queues a withdrawal of "beacon chain ETH shares" from this address to itself
    // fuzzed input amountGwei is sized-down, since it must be in GWEI and gets sized-up to be WEI
// TODO: reimplement similar test
    // function testQueueWithdrawalBeaconChainETHToSelf(uint128 amountGwei)
    //     public returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory, bytes32 /*withdrawalRoot*/) 
    // {
    //     // scale fuzzed amount up to be a whole amount of GWEI
    //     uint256 amount = uint256(amountGwei) * 1e9;
    //     address staker = address(this);
    //     address withdrawer = staker;

    //     testRestakeBeaconChainETHSuccessfully(staker, amount);

    //     (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) =
    //         _createQueuedWithdrawal(staker, amount, withdrawer);

    //     return (queuedWithdrawal, withdrawalRoot);
    // }
// TODO: reimplement similar test
    // function testQueueWithdrawalBeaconChainETHToDifferentAddress(address withdrawer, uint128 amountGwei)
    //     public
    //     filterFuzzedAddressInputs(withdrawer)
    //     returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory, bytes32 /*withdrawalRoot*/) 
    // {
    //     // scale fuzzed amount up to be a whole amount of GWEI
    //     uint256 amount = uint256(amountGwei) * 1e9;
    //     address staker = address(this);

    //     testRestakeBeaconChainETHSuccessfully(staker, amount);

    //     (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) =
    //         _createQueuedWithdrawal(staker, amount, withdrawer);

    //     return (queuedWithdrawal, withdrawalRoot);
    // }
// TODO: reimplement similar test

    // function testQueueWithdrawalBeaconChainETHFailsNonWholeAmountGwei(uint256 nonWholeAmount) external {
    //     // this also filters out the zero case, which will revert separately
    //     cheats.assume(nonWholeAmount % GWEI_TO_WEI != 0);
    //     cheats.expectRevert(bytes("EigenPodManager._queueWithdrawal: cannot queue a withdrawal of Beacon Chain ETH for an non-whole amount of gwei"));
    //     eigenPodManager.queueWithdrawal(nonWholeAmount, address(this));
    // }

    // function testQueueWithdrawalBeaconChainETHFailsZeroAmount() external {
    //     cheats.expectRevert(bytes("EigenPodManager._queueWithdrawal: amount must be greater than zero"));
    //     eigenPodManager.queueWithdrawal(0, address(this));
    // }

// TODO: reimplement similar test
    // function testCompleteQueuedWithdrawal() external {
    //     address staker = address(this);
    //     uint256 withdrawalAmount = 1e18;

    //     // withdrawalAmount is converted to GWEI here
    //     (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) = 
    //         testQueueWithdrawalBeaconChainETHToSelf(uint128(withdrawalAmount / 1e9));

    //     IEigenPod eigenPod = eigenPodManager.getPod(staker);
    //     uint256 eigenPodBalanceBefore = address(eigenPod).balance;

    //     uint256 middlewareTimesIndex = 0;

    //     // actually complete the withdrawal
    //     cheats.startPrank(staker);
    //     cheats.expectEmit(true, true, true, true, address(eigenPodManager));
    //     emit BeaconChainETHWithdrawalCompleted(
    //         queuedWithdrawal.podOwner,
    //         queuedWithdrawal.shares,
    //         queuedWithdrawal.nonce,
    //         queuedWithdrawal.delegatedAddress,
    //         queuedWithdrawal.withdrawer,
    //         withdrawalRoot
    //     );
    //     eigenPodManager.completeQueuedWithdrawal(queuedWithdrawal, middlewareTimesIndex);
    //     cheats.stopPrank();

    //     // TODO: make EigenPodMock do something so we can verify that it gets called appropriately?
    //     uint256 eigenPodBalanceAfter = address(eigenPod).balance;

    //     // verify that the withdrawal root does bit exist after queuing
    //     require(!eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");
    // }

// TODO: reimplement similar test
    // // creates a queued withdrawal of "beacon chain ETH shares", from `staker`, of `amountWei`, "to" the `withdrawer`
    // function _createQueuedWithdrawal(address staker, uint256 amountWei, address withdrawer)
    //     internal
    //     returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot)
    // {
    //     // create the struct, for reference / to return
    //     queuedWithdrawal = IEigenPodManager.BeaconChainQueuedWithdrawal({
    //         shares: amountWei,
    //         podOwner: staker,
    //         nonce: eigenPodManager.cumulativeWithdrawalsQueued(staker),
    //         startBlock: uint32(block.number),
    //         delegatedTo: delegationManagerMock.delegatedTo(staker),
    //         withdrawer: withdrawer
    //     });

    //     // verify that the withdrawal root does not exist before queuing
    //     require(!eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");

    //     // get staker nonce and shares before queuing
    //     uint256 nonceBefore = eigenPodManager.cumulativeWithdrawalsQueued(staker);
    //     int256 sharesBefore = eigenPodManager.podOwnerShares(staker);

    //     // actually create the queued withdrawal, and check for event emission
    //     cheats.startPrank(staker);
    
    //     cheats.expectEmit(true, true, true, true, address(eigenPodManager));
    //     emit BeaconChainETHWithdrawalQueued(
    //         queuedWithdrawal.podOwner,
    //         queuedWithdrawal.shares,
    //         queuedWithdrawal.nonce,
    //         queuedWithdrawal.delegatedAddress,
    //         queuedWithdrawal.withdrawer,
    //         eigenPodManager.calculateWithdrawalRoot(queuedWithdrawal)
    //     );
    //     withdrawalRoot = eigenPodManager.queueWithdrawal(amountWei, withdrawer);
    //     cheats.stopPrank();

    //     // verify that the withdrawal root does exist after queuing
    //     require(eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is false!");

    //     // verify that staker nonce incremented correctly and shares decremented correctly
    //     uint256 nonceAfter = eigenPodManager.cumulativeWithdrawalsQueued(staker);
    //     int256 sharesAfter = eigenPodManager.podOwnerShares(staker);
    //     require(nonceAfter == nonceBefore + 1, "nonce did not increment correctly on queuing withdrawal");
    //     require(sharesAfter + amountWei == sharesBefore, "shares did not decrement correctly on queuing withdrawal");

    //     return (queuedWithdrawal, withdrawalRoot);
    // }

    //Integration Test 
    // function testFullWithdrawalProofWithWrongWithdrawalFields(bytes32[] memory wrongWithdrawalFields) public {
    //     Relayer relay = new Relayer();
    //     uint256  WITHDRAWAL_FIELD_TREE_HEIGHT = 2;

    //     setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
    //     BeaconChainProofs.WithdrawalProof memory proofs = _getWithdrawalProof();
    //     bytes32 beaconStateRoot = getBeaconStateRoot();
    //     cheats.assume(wrongWithdrawalFields.length !=  2 ** WITHDRAWAL_FIELD_TREE_HEIGHT);
    //     validatorFields = getValidatorFields();

    //     cheats.expectRevert(bytes("BeaconChainProofs.verifyWithdrawal: withdrawalFields has incorrect length"));
    //     relay.verifyWithdrawal(beaconStateRoot, wrongWithdrawalFields, proofs);
    // }

    // // Integration Test
    // function testMismatchedWithdrawalProofInputs(uint64 numValidators, uint64 numValidatorProofs) external {
    //     cheats.assume(numValidators < numValidatorProofs && numValidatorProofs < 5);
    //     setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
    //     bytes[] memory validatorFieldsProofArray = new bytes[](numValidatorProofs);
    //     for (uint256 index = 0; index < numValidators; index++) {
    //         validatorFieldsProofArray[index] = abi.encodePacked(getValidatorProof());
    //     }
    //     bytes32[][] memory validatorFieldsArray = new bytes32[][](numValidators);
    //     for (uint256 index = 0; index < validatorFieldsArray.length; index++) {
    //          validatorFieldsArray[index] = getValidatorFields();
    //     }
    //     BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
    //     BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](1);

    //     withdrawalProofsArray[0] = _getWithdrawalProof();

    //     bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
    //     withdrawalFieldsArray[0] = withdrawalFields;

    //     cheats.expectRevert(bytes("EigenPod.verifyAndProcessWithdrawals: inputs must be same length"));
    //     pod.verifyAndProcessWithdrawals(0, stateRootProofStruct, withdrawalProofsArray, validatorFieldsProofArray, validatorFieldsArray, withdrawalFieldsArray);
    // }
