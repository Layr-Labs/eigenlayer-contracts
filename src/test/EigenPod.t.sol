// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../contracts/interfaces/IEigenPod.sol";
import "../contracts/pods/DelayedWithdrawalRouter.sol";
import "./utils/ProofParsing.sol";
import "./EigenLayerDeployer.t.sol";
import "../contracts/libraries/BeaconChainProofs.sol";
import "./integration/mocks/EIP_4788_Oracle_Mock.t.sol";
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
    EigenPodHarness public podInternalFunctionTester;

    // BeaconChainOracleMock public beaconChainOracle;

    // EIP-4788 oracle address
    EIP_4788_Oracle_Mock constant EIP_4788_ORACLE = EIP_4788_Oracle_Mock(0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02);
    uint256 constant SINGLE_VALIDATOR_STAKE = 32e18;
    // Generated proofs use this EigenPodManager address as the address for Create2 pod deployments
    address constant EPM_PROOFGEN_ADDRESS = 0x212224D2F2d262cd093eE13240ca4873fcCBbA3C;

    address pauser = address(69);
    address unpauser = address(489);
    
    address podAddress = address(123);
    mapping(address => bool) fuzzedAddressMapping;
    bytes signature;
    bytes32 depositDataRoot;

    uint32 WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
    IStrategy[] public initializeStrategiesToSetDelayBlocks;
    uint256[] public initializeWithdrawalDelayBlocks;
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
            IEigenPodManager(EPM_PROOFGEN_ADDRESS),
            GOERLI_GENESIS_TIME
        );

        // Create mock 4788 oracle
        cheats.etch(address(EIP_4788_ORACLE), type(EIP_4788_Oracle_Mock).runtimeCode);

        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));

        // this contract is deployed later to keep its address the same (for these tests)
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Ensure the EPM's address is consistent so created pods work with proofgen
        cheats.etch(EPM_PROOFGEN_ADDRESS, address(eigenPodManager).code);
        eigenPodManager = IEigenPodManager(EPM_PROOFGEN_ADDRESS);

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        DelegationManager delegationImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        StrategyManager strategyManagerImplementation = new StrategyManager(
            delegation,
            eigenPodManager,
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

        podInternalFunctionTester = new EigenPodHarness(
            ethPOSDeposit,
            eigenPodManager,
            GOERLI_GENESIS_TIME
        );

        DelayedWithdrawalRouter delayedWithdrawalRouterImplementation = new DelayedWithdrawalRouter(
            eigenPodManager
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

        cheats.deal(address(podOwner), 5 * SINGLE_VALIDATOR_STAKE);

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
        eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
    }

    function testDeployEigenPodWithoutActivateRestaking() public {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = getWithdrawalCredentialProof();
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());
        EIP_4788_ORACLE.setBlockRoot(GOERLI_GENESIS_TIME, getLatestBlockRoot());

        //this simulates that hasRestaking is set to false, as would be the case for deployed pods that have not yet restaked prior to M2
        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(uint256(0)));

        cheats.startPrank(podOwner);
        cheats.warp(GOERLI_GENESIS_TIME);
        cheats.expectRevert(bytes("EigenPod.verifyWithdrawalCredentials: restaking not active"));
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
        revert("TODO");
        // // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        // setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");

        // IEigenPod newPod = eigenPodManager.getPod(podOwner);

        // cheats.startPrank(podOwner);
        // cheats.expectEmit(true, true, true, true, address(newPod));
        // emit EigenPodStaked(pubkey);
        // eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
        // cheats.stopPrank();

        // bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        // validatorFieldsArray[0] = getValidatorFields();
        // bytes[] memory proofsArray = new bytes[](1);
        // proofsArray[0] = getWithdrawalCredentialProof();
        // BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        // uint40[] memory validatorIndices = new uint40[](1);
        // validatorIndices[0] = uint40(getValidatorIndex());
        // bytes32 beaconBlockRoot = getLatestBlockRoot();

        // //this simulates that hasRestaking is set to false, as would be the case for deployed pods that have not yet restaked prior to M2
        // cheats.store(address(newPod), bytes32(uint256(52)), bytes32(uint256(0)));
        // assertTrue(newPod.hasRestaked() == false, "EigenPod should not be restaked");

        // uint64 startTime = GOERLI_GENESIS_TIME + (BeaconChainProofs.SECONDS_PER_EPOCH * epochNum);
        // uint64 startEpoch = _timestampToEpoch(startTime);
        // // move to start time - this is the first slot in the epoch, and is where we will call activateRestaking
        // cheats.warp(startTime);

        // // activate restaking
        // EIP_4788_ORACLE.setBlockRoot(uint64(block.timestamp), beaconBlockRoot);
        // cheats.startPrank(podOwner);
        // newPod.startCheckpoint(false);

        // // Ensure verifyWC fails for each slot remaining in the epoch
        // for (uint i = 0; i < 32; i++) {
        //     // Move forward 0-31 slots
        //     uint64 slotTimestamp = uint64(startTime + (BeaconChainProofs.SECONDS_PER_SLOT * i));
        //     uint64 epoch = _timestampToEpoch(slotTimestamp);
        //     assertTrue(epoch == startEpoch, "calculated epoch should not change");
            
        //     EIP_4788_ORACLE.setBlockRoot(slotTimestamp, beaconBlockRoot);
        //     cheats.warp(slotTimestamp);
        //     cheats.expectRevert(bytes("EigenPod.verifyWithdrawalCredentials: proof must be in the epoch after activation"));
        //     newPod.verifyWithdrawalCredentials(
        //         slotTimestamp,
        //         stateRootProofStruct,
        //         validatorIndices,
        //         proofsArray,
        //         validatorFieldsArray
        //     );
        // }

        // // finally, move to the next epoch
        // cheats.warp(block.timestamp + BeaconChainProofs.SECONDS_PER_SLOT);
        // uint64 endEpoch = _timestampToEpoch(uint64(block.timestamp));
        // assertEq(startEpoch + 1, endEpoch, "should have advanced one epoch");

        // // now verifyWC should succeed
        // EIP_4788_ORACLE.setBlockRoot(uint64(block.timestamp), beaconBlockRoot);
        // newPod.verifyWithdrawalCredentials(
        //     uint64(block.timestamp),
        //     stateRootProofStruct,
        //     validatorIndices,
        //     proofsArray,
        //     validatorFieldsArray
        // );
        
        // cheats.stopPrank();
    }

    function _timestampToEpoch(uint64 timestamp) internal view returns (uint64) {
        require(timestamp >= GOERLI_GENESIS_TIME, "Test._timestampToEpoch: timestamp is before genesis");
        return (timestamp - GOERLI_GENESIS_TIME) / BeaconChainProofs.SECONDS_PER_EPOCH;
    }

    function testWithdrawNonBeaconChainETHBalanceWei() public {
        revert("TODO");
        // IEigenPod pod = testDeployAndVerifyNewEigenPod();

        // cheats.deal(address(podOwner), 10 ether);
        // emit log_named_address("Pod:", address(pod));

        // uint256 balanceBeforeDeposit = pod.nonBeaconChainETHBalanceWei();

        // (bool sent, ) = payable(address(pod)).call{value: 1 ether}("");

        // require(sent == true, "not sent");

        // uint256 balanceAfterDeposit = pod.nonBeaconChainETHBalanceWei();

        // require(
        //     balanceBeforeDeposit < balanceAfterDeposit 
        //     && (balanceAfterDeposit - balanceBeforeDeposit) == 1 ether, 
        //     "increment checks"
        // );

        // cheats.startPrank(podOwner, podOwner);
        // cheats.expectEmit(true, true, true, true, address(pod));
        // emit NonBeaconChainETHWithdrawn(podOwner, 1 ether);
        // pod.withdrawNonBeaconChainETHBalanceWei(
        //     podOwner,
        //     1 ether
        // );

        // uint256 balanceAfterWithdrawal = pod.nonBeaconChainETHBalanceWei();

        // require(
        //     balanceAfterWithdrawal < balanceAfterDeposit 
        //     && balanceAfterWithdrawal == balanceBeforeDeposit, 
        //     "decrement checks"
        // );

        // cheats.stopPrank();
    }

    function testDeployAndVerifyNewEigenPod() public returns (IEigenPod) {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        return _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
    }
    
    // /// @notice Similar test done in EP unit test
    // //test deploying an eigen pod with mismatched withdrawal credentials between the proof and the actual pod's address
    // function testDeployNewEigenPodWithWrongWithdrawalCreds(address wrongWithdrawalAddress) public {
    //     // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
    //     setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
    //     cheats.startPrank(podOwner);
    //     IEigenPod newPod;
    //     newPod = eigenPodManager.getPod(podOwner);
    //     cheats.expectEmit(true, true, true, true, address(newPod));
    //     emit EigenPodStaked(pubkey);
    //     eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
    //     cheats.stopPrank();


    //     // make sure that wrongWithdrawalAddress is not set to actual pod address
    //     cheats.assume(wrongWithdrawalAddress != address(newPod));

    //     bytes32[] memory validatorFields = getValidatorFields();
    //     validatorFields[1] = abi.encodePacked(bytes1(uint8(1)), bytes11(0), wrongWithdrawalAddress).toBytes32(0);

    //     bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
    //     validatorFieldsArray[0] = validatorFields;
    //     bytes[] memory proofsArray = new bytes[](1);
    //     proofsArray[0] = getWithdrawalCredentialProof();
    //     uint40[] memory validatorIndices = new uint40[](1);
    //     validatorIndices[0] = uint40(validatorIndex0);

    //     cheats.startPrank(podOwner);
    //     if (!newPod.hasRestaked()) {
    //         newPod.startCheckpoint(false);
    //     }
    //     // set oracle block root
    //     _setOracleBlockRoot();

    //     BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

    //     uint64 timestamp = _verifyWCStartTimestamp(newPod);
    //     cheats.warp(timestamp);

    //     cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: Proof is not for this EigenPod"));
    //     newPod.verifyWithdrawalCredentials(
    //         timestamp,
    //         stateRootProofStruct,
    //         validatorIndices,
    //         proofsArray,
    //         validatorFieldsArray
    //     );
    //     cheats.stopPrank();
    // }

    function testVerifyWithdrawalCredsFromNonPodOwnerAddress(address nonPodOwnerAddress) public {
        // nonPodOwnerAddress must be different from podOwner
        cheats.assume(nonPodOwnerAddress != podOwner);
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        cheats.startPrank(podOwner);

        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);

        eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = getWithdrawalCredentialProof();
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
        revert("TODO");
        // cheats.assume(timestamp > GOERLI_GENESIS_TIME);
        // // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        // setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        // IEigenPod newPod = testDeployAndVerifyNewEigenPod();

        //  // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        // setJSON("./src/test/test-data/balanceUpdateProof_updated_to_0ETH_302913.json");
        // // prove overcommitted balance
        // cheats.warp(timestamp);
        // _proveOverCommittedStake(newPod);


        // bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        // validatorFieldsArray[0] = getValidatorFields();

        // uint40[] memory validatorIndices = new uint40[](1);
        // validatorIndices[0] = uint40(getValidatorIndex());

        // bytes memory proof = abi.encodePacked(getBalanceUpdateProof());
        // bytes[] memory proofs = new bytes[](1);
        // proofs[0] = proof;

        // bytes32 newLatestBlockRoot = getLatestBlockRoot();
        // BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);
        // BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();      

        // cheats.expectRevert(bytes("EigenPod.verifyBalanceUpdate: Validators balance has already been updated for this timestamp"));
        // newPod.verifyBalanceUpdates(uint64(block.timestamp - 1), validatorIndices, stateRootProofStruct, proofs, validatorFieldsArray);
    }

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

    function testDeployingEigenPodRevertsWhenPaused() external {
        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_NEW_EIGENPODS);
        cheats.stopPrank();

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
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
        eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        cheats.deal(nonPodManager, SINGLE_VALIDATOR_STAKE);

        cheats.startPrank(nonPodManager);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodManager: not eigenPodManager"));
        newPod.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
    }

    function testVerifyCorrectWithdrawalCredentialsRevertsWhenPaused() external {
        // ./solidityProofGen  -newBalance=32000115173 "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_CREDENTIALS);
        cheats.stopPrank();

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = getWithdrawalCredentialProof();

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        uint64 timestamp = _verifyWCStartTimestamp(newPod);
        cheats.warp(timestamp);
        EIP_4788_ORACLE.setBlockRoot(timestamp, newBeaconStateRoot);

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
        require(info1.lastCheckpointedAt == info2.lastCheckpointedAt, "lastCheckpointedAt does not match");
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

    function _newPod_VerifiedWC(
        address _podOwner,
        bytes memory _signature,
        bytes32 _depositDataRoot
    ) internal returns (IEigenPod) {
        IEigenPod pod = eigenPodManager.getPod(_podOwner);

        cheats.prank(_podOwner);
        cheats.expectEmit(true, true, true, true, address(pod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, _signature, _depositDataRoot);

        _verifyWithdrawalCredentials(pod, _podOwner);

        return pod;
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
        eigenPodManager.stake{value: SINGLE_VALIDATOR_STAKE}(pubkey, _signature, _depositDataRoot);
        cheats.stopPrank();

        _verifyWithdrawalCredentials(newPod, _podOwner);

        return newPod;
    }

    function _verifyWithdrawalCredentials(IEigenPod pod, address _podOwner) internal {
        cheats.startPrank(podOwner);

        (
            bytes32 beaconBlockRoot,
            BeaconChainProofs.StateRootProof memory stateRootProof,
            uint40[] memory validatorIndices,
            bytes[] memory validatorFieldsProofs,
            bytes32[][] memory validatorFields
        ) = _getWithdrawalCredentialProof();

        // Set the 4788 oracle block root
        EIP_4788_ORACLE.setBlockRoot(uint64(block.timestamp), beaconBlockRoot);

        int beaconChainETHSharesBefore = eigenPodManager.podOwnerShares(_podOwner);

        emit log_named_address("pod", address(pod));
        emit log_named_bytes32("wc", BeaconChainProofs.getWithdrawalCredentials(validatorFields[0]));

        pod.verifyWithdrawalCredentials(
            uint64(block.timestamp),
            stateRootProof,
            validatorIndices,
            validatorFieldsProofs,
            validatorFields
        );

        int beaconChainETHSharesAfter = eigenPodManager.podOwnerShares(_podOwner);
        uint effectiveBalanceWei = uint(BeaconChainProofs.getEffectiveBalanceGwei(validatorFields[0])) * GWEI_TO_WEI;

        emit log_named_uint("effectiveBalance (wei)", effectiveBalanceWei);
        emit log_named_uint("beaconChainETHSharesAfter", uint(beaconChainETHSharesAfter));
        emit log_named_uint("beaconChainETHSharesBefore", uint(beaconChainETHSharesBefore));

        assertEq(beaconChainETHSharesBefore + int(effectiveBalanceWei), beaconChainETHSharesAfter, 
            "_verifyWithdrawalCredentials: shares not updated correctly");

        cheats.stopPrank();
    }

    function _verifyWCStartTimestamp(IEigenPod pod) internal  returns (uint64) {
        uint64 genesis = EigenPod(payable(address(pod))).GENESIS_TIME();
        uint64 activateRestakingTimestamp;

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

    function _getLatestDelayedWithdrawalAmount(address recipient) internal view returns (uint256) {
        return
            delayedWithdrawalRouter
                .userDelayedWithdrawalByIndex(recipient, delayedWithdrawalRouter.userWithdrawalsLength(recipient) - 1)
                .amount;
    }

    function _getStateRootProof() internal returns (BeaconChainProofs.StateRootProof memory) {
        return BeaconChainProofs.StateRootProof(getBeaconStateRoot(), getStateRootProof());
    }

    /// @dev Read various proving parameters from JSON
    function _getWithdrawalCredentialProof() 
        internal 
        returns (
            bytes32 beaconBlockRoot,
            BeaconChainProofs.StateRootProof memory stateRootProof,
            uint40[] memory validatorIndices,
            bytes[] memory validatorFieldsProofs,
            bytes32[][] memory validatorFields
        ) 
    {
        beaconBlockRoot = getLatestBlockRoot();

        stateRootProof = BeaconChainProofs.StateRootProof({
            beaconStateRoot: getBeaconStateRoot(),
            proof: getStateRootProof()
        });

        validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        validatorFieldsProofs = new bytes[](1);
        validatorFieldsProofs[0] = getWithdrawalCredentialProof();

        validatorFields = new bytes32[][](1);
        validatorFields[0] = getValidatorFields();
    }

    function _computeTimestampAtSlot(uint64 slot) internal pure returns (uint64) {
        return uint64(GOERLI_GENESIS_TIME + slot * SECONDS_PER_SLOT);
    }
}