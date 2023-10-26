// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../contracts/interfaces/IEigenPod.sol";
import "../contracts/pods/DelayedWithdrawalRouter.sol";
import "./utils/ProofParsing.sol";
import "./EigenLayerDeployer.t.sol";
import "./mocks/MiddlewareRegistryMock.sol";
import "../contracts/libraries/BeaconChainProofs.sol";
import "./mocks/BeaconChainOracleMock.sol";
import "./harnesses/EigenPodHarness.sol";

contract EigenPodTests is ProofParsing, EigenPodPausingConstants {
    using BytesLib for bytes;

    uint256 internal constant GWEI_TO_WEI = 1e9;

    bytes pubkey =
        hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";
    uint40 validatorIndex0 = 0;
    uint40 validatorIndex1 = 1;
    //hash tree root of list of validators
    bytes32 validatorTreeRoot;

    //hash tree root of individual validator container
    bytes32 validatorRoot;

    address podOwner = address(42000094993494);

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
    IBeacon public eigenPodBeacon;
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

    /// @notice Emitted when `maxPods` value is updated from `previousValue` to `newValue`
    event MaxPodsUpdated(uint256 previousValue, uint256 newValue);

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

        /// weird workaround: check commit before this
        /// either related to foundry bug or emptty contract below this one
        EmptyContract emptyContract2 = new EmptyContract();

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        EmptyContract emptyContract = new EmptyContract();
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
                0 /*initialPausedStatus*/
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
                type(uint256).max, // maxPods
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
        pod.withdrawBeforeRestaking();
        require(_getLatestDelayedWithdrawalAmount(podOwner) == stakeAmount, "Payment amount should be stake amount");
        require(
            pod.mostRecentWithdrawalTimestamp() == uint64(block.timestamp),
            "Most recent withdrawal block number not updated"
        );
    }

    function testDeployEigenPodWithoutActivateRestaking() public {
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
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

    function testDeployEigenPodTooSoon() public {
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        uint64 timestamp = 0;
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        cheats.startPrank(podOwner);
        cheats.expectRevert(
            bytes(
                "EigenPod.proofIsForValidTimestamp: beacon chain proof must be for timestamp after mostRecentWithdrawalTimestamp"
            )
        );
        newPod.verifyWithdrawalCredentials(
            timestamp,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        cheats.stopPrank();
    }

    function testWithdrawFromPod() public {
        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod pod = eigenPodManager.getPod(podOwner);
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
    function testFullWithdrawalFlow() public returns (IEigenPod) {
        //this call is to ensure that validator 302913 has proven their withdrawalcreds
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
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
        IEigenPod pod = testFullWithdrawalFlow();

        // ./solidityProofGen "WithdrawalFieldsProof" 302913 146 8092 true false "data/withdrawal_proof_goerli/goerli_block_header_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "data/withdrawal_proof_goerli/goerli_slot_6397852.json" "data/withdrawal_proof_goerli/goerli_block_header_6397852.json" "data/withdrawal_proof_goerli/goerli_block_6397852.json" "fullWithdrawalProof_Latest_1SlotAdvanced.json" true
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest_1SlotAdvanced.json");
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(getLatestBlockRoot());

        withdrawalFields = getWithdrawalFields();
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(
            withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]
        );
        uint64 leftOverBalanceWEI = uint64(
            withdrawalAmountGwei - (pod.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR())
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
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
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
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        return _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
    }

    // //test freezing operator after a beacon chain slashing event
    function testUpdateSlashedBeaconBalance() public {
        //make initial deposit
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.warp(GOERLI_GENESIS_TIME);
        // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
        _proveOverCommittedStake(newPod);

        uint64 newValidatorBalance = BeaconChainProofs.getBalanceAtIndex(getBalanceRoot(), uint40(getValidatorIndex()));
        int256 beaconChainETHShares = eigenPodManager.podOwnerShares(podOwner);

        require(
            beaconChainETHShares == int256((MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR) * GWEI_TO_WEI),
            "eigenPodManager shares not updated correctly"
        );
    }

    //test deploying an eigen pod with mismatched withdrawal credentials between the proof and the actual pod's address
    function testDeployNewEigenPodWithWrongWithdrawalCreds(address wrongWithdrawalAddress) public {
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod newPod;
        newPod = eigenPodManager.getPod(podOwner);
        // make sure that wrongWithdrawalAddress is not set to actual pod address
        cheats.assume(wrongWithdrawalAddress != address(newPod));

        validatorFields = getValidatorFields();
        validatorFields[1] = abi.encodePacked(bytes1(uint8(1)), bytes11(0), wrongWithdrawalAddress).toBytes32(0);
        uint64 timestamp = 0;

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = validatorFields;
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(validatorIndex0);

        cheats.startPrank(podOwner);
        cheats.warp(timestamp);
        if (!newPod.hasRestaked()) {
            newPod.activateRestaking();
        }
        // set oracle block root
        _setOracleBlockRoot();

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        cheats.warp(timestamp += 1);
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
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913_exited.json");
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
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        uint64 timestamp = 1;

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(validatorIndex0);

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

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

    //test that when withdrawal credentials are verified more than once, it reverts
    function testDeployNewEigenPodWithActiveValidator() public {
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        uint64 timestamp = 1;

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        bytes[] memory proofsArray = new bytes[](1);
        proofsArray[0] = abi.encodePacked(getWithdrawalCredentialProof());

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        cheats.startPrank(podOwner);
        cheats.expectRevert(
            bytes(
                "EigenPod.verifyCorrectWithdrawalCredentials: Validator must be inactive to prove withdrawal credentials"
            )
        );
        pod.verifyWithdrawalCredentials(
            timestamp,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        cheats.stopPrank();
    }

    // // 3. Single withdrawal credential
    // // Test: Owner proves an withdrawal credential.
    // //                     validator status should be marked as ACTIVE

    function testProveSingleWithdrawalCredential() public {
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();

        assertTrue(
            pod.validatorStatus(validatorPubkeyHash) == IEigenPod.VALIDATOR_STATUS.ACTIVE,
            "wrong validator status"
        );
    }

    // // 5. Prove overcommitted balance
    // // Setup: Run (3).
    // // Test: Watcher proves an overcommitted balance for validator from (3).
    // //                     validator status should be marked as OVERCOMMITTED

    function testProveOverCommittedBalance() public {
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
        setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
        // prove overcommitted balance
        cheats.warp(GOERLI_GENESIS_TIME);
        _proveOverCommittedStake(newPod);

        uint256 validatorRestakedBalanceAfter = newPod
            .validatorPubkeyHashToInfo(validatorPubkeyHash)
            .restakedBalanceGwei;

        uint64 newValidatorBalance = BeaconChainProofs.getBalanceAtIndex(getBalanceRoot(), uint40(getValidatorIndex()));
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
            int256(validatorRestakedBalanceBefore) - int256(validatorRestakedBalanceAfter) ==
                shareDiff / int256(GWEI_TO_WEI),
            "validator restaked balance not updated"
        );
    }

    function testVerifyUndercommittedBalance() public {
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
        setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
        // prove overcommitted balance
        cheats.warp(GOERLI_GENESIS_TIME);
        _proveOverCommittedStake(newPod);

        cheats.warp(block.timestamp + 1);
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 100 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913_incrementedBlockBy100.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913_incrementedBlockBy100.json");
        _proveUnderCommittedStake(newPod);

        uint256 validatorRestakedBalanceAfter = newPod
            .validatorPubkeyHashToInfo(validatorPubkeyHash)
            .restakedBalanceGwei;

        uint64 newValidatorBalance = BeaconChainProofs.getBalanceAtIndex(getBalanceRoot(), uint40(getValidatorIndex()));
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

    function testTooSoonBalanceUpdate(uint64 oracleTimestamp, uint64 mostRecentBalanceUpdateTimestamp) external {
        cheats.assume(oracleTimestamp < mostRecentBalanceUpdateTimestamp);
        _deployInternalFunctionTester();

        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        validatorFields = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());
        BeaconChainProofs.BalanceUpdateProof memory proof = _getBalanceUpdateProof();

        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        emit log_named_bytes32("newBeaconStateRoot", newBeaconStateRoot);
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newBeaconStateRoot);

        cheats.expectRevert(
            bytes("EigenPod.verifyBalanceUpdate: Validators balance has already been updated for this timestamp")
        );
        podInternalFunctionTester.verifyBalanceUpdate(
            oracleTimestamp,
            0,
            bytes32(0),
            proof,
            validatorFields,
            mostRecentBalanceUpdateTimestamp
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

    function testStakeOnEigenPodFromNonPodManagerAddress(address nonPodManager) external fuzzedAddress(nonPodManager) {
        cheats.assume(nonPodManager != address(eigenPodManager));

        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

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
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newBeaconStateRoot);

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        uint64 timestamp = 1;

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
        setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        BeaconChainProofs.BalanceUpdateProof[] memory proofs = new BeaconChainProofs.BalanceUpdateProof[](1);
        proofs[0] = _getBalanceUpdateProof();

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

    function testFullWithdrawalProofWithWrongWithdrawalFields(bytes32[] memory wrongWithdrawalFields) public {
        Relayer relay = new Relayer();
        uint256 WITHDRAWAL_FIELD_TREE_HEIGHT = 2;

        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        BeaconChainProofs.WithdrawalProof memory proofs = _getWithdrawalProof();
        bytes32 beaconStateRoot = getBeaconStateRoot();
        cheats.assume(wrongWithdrawalFields.length != 2 ** WITHDRAWAL_FIELD_TREE_HEIGHT);
        validatorFields = getValidatorFields();

        cheats.expectRevert(bytes("BeaconChainProofs.verifyWithdrawal: withdrawalFields has incorrect length"));
        relay.verifyWithdrawal(beaconStateRoot, wrongWithdrawalFields, proofs);
    }

    function testCheckThatHasRestakedIsSetToTrue() public returns (IEigenPod) {
        testStaking();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should not be restaked");
        return pod;
    }

    function testActivateRestakingWithM2Pods() external {
        IEigenPod pod = testCheckThatHasRestakedIsSetToTrue();
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.activateRestaking();
        cheats.stopPrank();
    }

    function testWithdrawBeforeRestakingWithM2Pods() external {
        IEigenPod pod = testCheckThatHasRestakedIsSetToTrue();
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    function testAttemptedWithdrawalAfterVerifyingWithdrawalCredentials() public {
        testDeployAndVerifyNewEigenPod();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        IEigenPod(pod).withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    function testBalanceProofWithWrongTimestamp(uint64 timestamp) public {
        cheats.assume(timestamp > GOERLI_GENESIS_TIME);
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
        // prove overcommitted balance
        cheats.warp(timestamp);
        _proveOverCommittedStake(newPod);

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        BeaconChainProofs.BalanceUpdateProof[] memory proofs = new BeaconChainProofs.BalanceUpdateProof[](1);
        proofs[0] = _getBalanceUpdateProof();

        bytes32 newLatestBlockRoot = getLatestBlockRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();

        cheats.expectRevert(
            bytes("EigenPod.verifyBalanceUpdate: Validators balance has already been updated for this timestamp")
        );
        newPod.verifyBalanceUpdates(
            uint64(block.timestamp - 1),
            validatorIndices,
            stateRootProofStruct,
            proofs,
            validatorFieldsArray
        );
    }

    function testProcessFullWithdrawalForLessThanMaxRestakedBalance(uint64 withdrawalAmount) public {
        _deployInternalFunctionTester();
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount < MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR);
        IEigenPod.ValidatorInfo memory validatorInfo = IEigenPod.ValidatorInfo({
            validatorIndex: 0,
            restakedBalanceGwei: 0,
            mostRecentBalanceUpdateTimestamp: 0,
            status: IEigenPod.VALIDATOR_STATUS.ACTIVE
        });
        uint64 balanceBefore = podInternalFunctionTester.withdrawableRestakedExecutionLayerGwei();
        podInternalFunctionTester.processFullWithdrawal(0, bytes32(0), 0, podOwner, withdrawalAmount, validatorInfo);
        require(
            podInternalFunctionTester.withdrawableRestakedExecutionLayerGwei() - balanceBefore == withdrawalAmount,
            "withdrawableRestakedExecutionLayerGwei hasn't been updated correctly"
        );
    }

    function testWithdrawBeforeRestakingAfterRestaking() public {
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_510257.json"
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");

        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        cheats.startPrank(podOwner);
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    //post M2, all new pods deployed will have "hasRestaked = true".  THis tests that
    function testDeployedPodIsRestaked() public fuzzedAddress(podOwner) {
        cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();

        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should be restaked");
    }

    function testTryToActivateRestakingAfterHasRestakedIsSet() public {
        cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();

        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should be restaked");

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.activateRestaking();
    }

    function testTryToWithdrawBeforeRestakingAfterHasRestakedIsSet() public {
        cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();

        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should be restaked");

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.withdrawBeforeRestaking();
    }

    function testMismatchedWithdrawalProofInputs(uint64 numValidators, uint64 numValidatorProofs) external {
        cheats.assume(numValidators < numValidatorProofs && numValidatorProofs < 5);

        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        bytes[] memory validatorFieldsProofArray = new bytes[](numValidatorProofs);
        for (uint256 index = 0; index < numValidators; index++) {
            validatorFieldsProofArray[index] = abi.encodePacked(getValidatorProof());
        }
        bytes32[][] memory validatorFieldsArray = new bytes32[][](numValidators);
        for (uint256 index = 0; index < validatorFieldsArray.length; index++) {
            validatorFieldsArray[index] = getValidatorFields();
        }

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](1);
        withdrawalProofsArray[0] = _getWithdrawalProof();
        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
        withdrawalFieldsArray[0] = withdrawalFields;

        cheats.expectRevert(bytes("EigenPod.verifyAndProcessWithdrawals: inputs must be same length"));
        newPod.verifyAndProcessWithdrawals(
            0,
            stateRootProofStruct,
            withdrawalProofsArray,
            validatorFieldsProofArray,
            validatorFieldsArray,
            withdrawalFieldsArray
        );
    }

    function testProveWithdrawalFromBeforeLastWithdrawBeforeRestaking() external {
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod pod = eigenPodManager.getPod(podOwner);

        cheats.store(address(pod), bytes32(uint256(52)), bytes32(uint256(1)));
        require(pod.hasRestaked() != true, "Pod should not be restaked");

        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(getLatestBlockRoot());

        BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](1);
        withdrawalProofsArray[0] = _getWithdrawalProof();
        uint64 timestampOfWithdrawal = Endian.fromLittleEndianUint64(withdrawalProofsArray[0].timestampRoot);
        uint256 newTimestamp = timestampOfWithdrawal + 2500;
        cheats.warp(newTimestamp);
        cheats.startPrank(podOwner);
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();

        bytes[] memory validatorFieldsProofArray = new bytes[](1);
        validatorFieldsProofArray[0] = abi.encodePacked(getValidatorProof());
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
        withdrawalFieldsArray[0] = withdrawalFields;
        cheats.warp(timestampOfWithdrawal);

        cheats.expectRevert(
            bytes(
                "EigenPod.proofIsForValidTimestamp: beacon chain proof must be for timestamp after mostRecentWithdrawalTimestamp"
            )
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

    function testPodReceiveFallBack(uint256 amountETH) external {
        cheats.assume(amountETH > 0);
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        cheats.deal(address(this), amountETH);

        Address.sendValue(payable(address(pod)), amountETH);
        require(address(pod).balance == amountETH, "Pod should have received ETH");
    }

    /**
     * This is a regression test for a bug (EIG-14) found by Hexens.  Lets say podOwner sends 32 ETH to the EigenPod,
     * the nonBeaconChainETHBalanceWei increases by 32 ETH. podOwner calls withdrawBeforeRestaking, which
     * will simply send the entire ETH balance (32 ETH) to the owner. The owner activates restaking,
     * creates a validator and verifies the withdrawal credentials, receiving 32 ETH in shares.
     * They can exit the validator, the pod gets the 32ETH and they can call withdrawNonBeaconChainETHBalanceWei
     * And simply withdraw the 32ETH because nonBeaconChainETHBalanceWei is 32ETH.  This was an issue because
     * nonBeaconChainETHBalanceWei was never zeroed out in _processWithdrawalBeforeRestaking
     */
    function testValidatorBalanceCannotBeRemovedFromPodViaNonBeaconChainETHBalanceWei() external {
        cheats.startPrank(podOwner);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        uint256 amount = 32 ether;

        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(0));
        cheats.deal(address(this), amount);
        // simulate a withdrawal processed on the beacon chain, pod balance goes to 32 ETH
        Address.sendValue(payable(address(newPod)), amount);
        require(newPod.nonBeaconChainETHBalanceWei() == amount, "nonBeaconChainETHBalanceWei should be 32 ETH");
        //simulate that hasRestaked is set to false, so that we can test withdrawBeforeRestaking for pods deployed before M2 activation
        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(uint256(1)));
        //this is an M1 pod so hasRestaked should be false
        require(newPod.hasRestaked() == false, "Pod should be restaked");
        cheats.startPrank(podOwner);
        newPod.activateRestaking();
        cheats.stopPrank();
        require(newPod.nonBeaconChainETHBalanceWei() == 0, "nonBeaconChainETHBalanceWei should be 32 ETH");
    }

    /**
     * Regression test for a bug that allowed balance updates to be made for withdrawn validators.  Thus
     * the validator's balance could be maliciously proven to be 0 before the validator themselves are
     * able to prove their withdrawal.
     */
    function testBalanceUpdateMadeAfterWithdrawableEpochFails() external {
        //make initial deposit
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.roll(block.number + 1);
        // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        BeaconChainProofs.BalanceUpdateProof[] memory proofs = new BeaconChainProofs.BalanceUpdateProof[](1);
        proofs[0] = _getBalanceUpdateProof();
        bytes32 newLatestBlockRoot = getLatestBlockRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        proofs[0].balanceRoot = bytes32(uint256(0));

        validatorFieldsArray[0][7] = bytes32(uint256(0));
        cheats.warp(GOERLI_GENESIS_TIME + 1 days);
        uint64 oracleTimestamp = uint64(block.timestamp);
        cheats.expectRevert(bytes("EigenPod.verifyBalanceUpdate: validator is withdrawable but has not withdrawn"));
        newPod.verifyBalanceUpdates(
            oracleTimestamp,
            validatorIndices,
            stateRootProofStruct,
            proofs,
            validatorFieldsArray
        );
    }

    function testWithdrawlBeforeRestakingFromNonPodOwnerAddress(address nonPodOwner) external {
        cheats.assume(nonPodOwner != podOwner);
        cheats.startPrank(podOwner);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        cheats.startPrank(nonPodOwner);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodOwner: not podOwner"));
        newPod.withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    function testDelayedWithdrawalIsCreatedByWithdrawBeforeRestaking() external {
        cheats.startPrank(podOwner);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        uint256 amount = 32 ether;

        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(0));
        cheats.deal(address(this), amount);
        // simulate a withdrawal processed on the beacon chain, pod balance goes to 32 ETH
        Address.sendValue(payable(address(newPod)), amount);
        require(newPod.nonBeaconChainETHBalanceWei() == amount, "nonBeaconChainETHBalanceWei should be 32 ETH");

        cheats.startPrank(podOwner);
        newPod.withdrawBeforeRestaking();
        cheats.stopPrank();

        require(_getLatestDelayedWithdrawalAmount(podOwner) == amount, "Payment amount should be stake amount");
        require(newPod.nonBeaconChainETHBalanceWei() == 0, "nonBeaconChainETHBalanceWei should be 32 ETH");
    }

    function testFullWithdrawalAmounts(bytes32 pubkeyHash, uint64 withdrawalAmount) external {
        _deployInternalFunctionTester();
        IEigenPod.ValidatorInfo memory validatorInfo = IEigenPod.ValidatorInfo({
            validatorIndex: 0,
            restakedBalanceGwei: 0,
            mostRecentBalanceUpdateTimestamp: 0,
            status: IEigenPod.VALIDATOR_STATUS.ACTIVE
        });
        IEigenPod.VerifiedWithdrawal memory vw = podInternalFunctionTester.processFullWithdrawal(
            0,
            pubkeyHash,
            0,
            podOwner,
            withdrawalAmount,
            validatorInfo
        );

        if (withdrawalAmount > podInternalFunctionTester.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR()) {
            require(
                vw.amountToSendGwei ==
                    withdrawalAmount - podInternalFunctionTester.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR(),
                "newAmount should be MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR"
            );
        } else {
            require(vw.amountToSendGwei == 0, "newAmount should be withdrawalAmount");
        }
    }

    function testProcessPartialWithdrawal(
        uint40 validatorIndex,
        uint64 withdrawalTimestamp,
        address recipient,
        uint64 partialWithdrawalAmountGwei
    ) external {
        _deployInternalFunctionTester();
        cheats.expectEmit(true, true, true, true, address(podInternalFunctionTester));
        emit PartialWithdrawalRedeemed(validatorIndex, withdrawalTimestamp, recipient, partialWithdrawalAmountGwei);
        IEigenPod.VerifiedWithdrawal memory vw = podInternalFunctionTester.processPartialWithdrawal(
            validatorIndex,
            withdrawalTimestamp,
            recipient,
            partialWithdrawalAmountGwei
        );

        require(vw.amountToSendGwei == partialWithdrawalAmountGwei, "newAmount should be partialWithdrawalAmountGwei");
    }

    function testRecoverTokens(uint256 amount, address recipient) external {
        cheats.assume(amount > 0 && amount < 1e30);
        IEigenPod pod = testDeployAndVerifyNewEigenPod();
        IERC20 randomToken = new ERC20PresetFixedSupply("rand", "RAND", 1e30, address(this));

        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = randomToken;
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = amount;

        randomToken.transfer(address(pod), amount);
        require(randomToken.balanceOf(address(pod)) == amount, "randomToken balance should be amount");

        uint256 recipientBalanceBefore = randomToken.balanceOf(recipient);

        cheats.startPrank(podOwner);
        pod.recoverTokens(tokens, amounts, recipient);
        cheats.stopPrank();
        require(
            randomToken.balanceOf(address(recipient)) - recipientBalanceBefore == amount,
            "recipient should have received amount"
        );
    }

    function testRecoverTokensMismatchedInputs() external {
        uint256 tokenListLen = 5;
        uint256 amountsToWithdrawLen = 2;

        IEigenPod pod = testDeployAndVerifyNewEigenPod();

        IERC20[] memory tokens = new IERC20[](tokenListLen);
        uint256[] memory amounts = new uint256[](amountsToWithdrawLen);

        cheats.expectRevert(bytes("EigenPod.recoverTokens: tokenList and amountsToWithdraw must be same length"));
        cheats.startPrank(podOwner);
        pod.recoverTokens(tokens, amounts, address(this));
        cheats.stopPrank();
    }

    function _proveOverCommittedStake(IEigenPod newPod) internal {
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        BeaconChainProofs.BalanceUpdateProof[] memory proofs = new BeaconChainProofs.BalanceUpdateProof[](1);
        proofs[0] = _getBalanceUpdateProof();

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

        BeaconChainProofs.BalanceUpdateProof[] memory proofs = new BeaconChainProofs.BalanceUpdateProof[](1);
        proofs[0] = _getBalanceUpdateProof();

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

    /// @notice Test that the Merkle proof verification fails when the proof length is empty
    function testVerifyInclusionKeccakFailsForEmptyProof(bytes32 root, bytes32 leaf, uint256 index) public {
        bytes memory emptyProof = new bytes(0);
        cheats.expectRevert(
            bytes("Merkle.processInclusionProofKeccak: proof length should be a non-zero multiple of 32")
        );
        Merkle.verifyInclusionKeccak(emptyProof, root, leaf, index);
    }

    /// @notice Test that the Merkle proof verification fails when the proof length is not a multiple of 32
    function testVerifyInclusionKeccakFailsForNonMultipleOf32ProofLength(
        bytes32 root,
        bytes32 leaf,
        uint256 index,
        bytes memory proof
    ) public {
        cheats.assume(proof.length % 32 != 0);
        cheats.expectRevert(
            bytes("Merkle.processInclusionProofKeccak: proof length should be a non-zero multiple of 32")
        );
        Merkle.verifyInclusionKeccak(proof, root, leaf, index);
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

    // verifies that the `maxPods` variable is enforced on the `EigenPod.stake` function
    function test_maxPodsEnforcementOnStake(
        bytes calldata _pubkey,
        bytes calldata _signature,
        bytes32 _depositDataRoot
    ) public {
        // set pod limit to current number of pods
        cheats.startPrank(unpauser);
        EigenPodManager(address(eigenPodManager)).setMaxPods(EigenPodManager(address(eigenPodManager)).numPods());
        cheats.stopPrank();

        cheats.startPrank(podOwner);
        cheats.expectRevert("EigenPodManager._deployPod: pod limit reached");
        eigenPodManager.stake{value: 32 ether}(_pubkey, _signature, _depositDataRoot);
        cheats.stopPrank();

        // set pod limit to *one more than* current number of pods
        cheats.startPrank(unpauser);
        EigenPodManager(address(eigenPodManager)).setMaxPods(EigenPodManager(address(eigenPodManager)).numPods() + 1);
        cheats.stopPrank();

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        // successful call
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(_pubkey);
        eigenPodManager.stake{value: 32 ether}(_pubkey, _signature, _depositDataRoot);
        cheats.stopPrank();
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

    // verifies that the `maxPods` variable is enforced on the `EigenPod.createPod` function
    function test_maxPodsEnforcementOnCreatePod() public {
        // set pod limit to current number of pods
        cheats.startPrank(unpauser);
        uint256 previousValue = EigenPodManager(address(eigenPodManager)).maxPods();
        uint256 newValue = EigenPodManager(address(eigenPodManager)).numPods();
        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit MaxPodsUpdated(previousValue, newValue);
        EigenPodManager(address(eigenPodManager)).setMaxPods(newValue);
        cheats.stopPrank();

        cheats.expectRevert("EigenPodManager._deployPod: pod limit reached");
        eigenPodManager.createPod();

        // set pod limit to *one more than* current number of pods
        cheats.startPrank(unpauser);
        previousValue = EigenPodManager(address(eigenPodManager)).maxPods();
        newValue = EigenPodManager(address(eigenPodManager)).numPods() + 1;
        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit MaxPodsUpdated(previousValue, newValue);
        EigenPodManager(address(eigenPodManager)).setMaxPods(newValue);
        cheats.stopPrank();

        // successful call
        eigenPodManager.createPod();
    }

    function test_setMaxPods(uint256 newValue) public {
        cheats.startPrank(unpauser);
        uint256 previousValue = EigenPodManager(address(eigenPodManager)).maxPods();
        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit MaxPodsUpdated(previousValue, newValue);
        EigenPodManager(address(eigenPodManager)).setMaxPods(newValue);
        cheats.stopPrank();

        require(EigenPodManager(address(eigenPodManager)).maxPods() == newValue, "maxPods value not set correctly");
    }

    function test_setMaxPods_RevertsWhenNotCalledByUnpauser(address notUnpauser) public fuzzedAddress(notUnpauser) {
        cheats.assume(notUnpauser != unpauser);
        uint256 newValue = 0;
        cheats.startPrank(notUnpauser);
        cheats.expectRevert("msg.sender is not permissioned as unpauser");
        EigenPodManager(address(eigenPodManager)).setMaxPods(newValue);
        cheats.stopPrank();
    }

    /* TODO: reimplement similar tests
    function testQueueBeaconChainETHWithdrawalWithoutProvingFullWithdrawal() external {
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_block_header_6399998.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_302913.json"         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
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
                earningsReceiver: operator,
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
        uint64 timestamp = 0;
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

        cheats.warp(timestamp += 1);
        newPod.verifyWithdrawalCredentials(
            timestamp,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
        cheats.stopPrank();

        int256 beaconChainETHSharesAfter = eigenPodManager.podOwnerShares(_podOwner);
        uint256 valBalance = Endian.fromLittleEndianUint64(getValidatorFields()[2]);
        uint256 effectiveBalance = valBalance *
            GWEI_TO_WEI;
        require(
            (beaconChainETHSharesAfter - beaconChainETHSharesBefore) == int256(effectiveBalance),
            "eigenPodManager shares not updated correctly"
        );
        return newPod;
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

    function _getBalanceUpdateProof() internal returns (BeaconChainProofs.BalanceUpdateProof memory) {
        bytes32 balanceRoot = getBalanceRoot();
        BeaconChainProofs.BalanceUpdateProof memory proofs = BeaconChainProofs.BalanceUpdateProof(
            abi.encodePacked(getValidatorBalanceProof()),
            abi.encodePacked(getWithdrawalCredentialProof()), //technically this is to verify validator pubkey in the validator fields, but the WC proof is effectively the same so we use it here again.
            balanceRoot
        );

        return proofs;
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

        {
            bytes32 blockRoot = getBlockRoot();
            bytes32 slotRoot = getSlotRoot();
            bytes32 timestampRoot = getTimestampRoot();
            bytes32 executionPayloadRoot = getExecutionPayloadRoot();

            return
                BeaconChainProofs.WithdrawalProof(
                    abi.encodePacked(getWithdrawalProof()),
                    abi.encodePacked(getSlotProof()),
                    abi.encodePacked(getExecutionPayloadProof()),
                    abi.encodePacked(getTimestampProof()),
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

    function testEffectiveRestakedBalance() public {
        uint64 amountGwei = 29134000000;
        uint64 effectiveBalance = (amountGwei);
        emit log_named_uint("effectiveBalance", effectiveBalance);
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
        BeaconChainProofs.verifyWithdrawal(beaconStateRoot, withdrawalFields, proofs);
    }
}
