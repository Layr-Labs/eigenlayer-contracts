// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../contracts/interfaces/IEigenPod.sol";
import "../contracts/interfaces/IBLSPublicKeyCompendium.sol";
import "../contracts/middleware/BLSPublicKeyCompendium.sol";
import "../contracts/pods/DelayedWithdrawalRouter.sol";
import "./utils/ProofParsing.sol";
import "./EigenLayerDeployer.t.sol";
import "./mocks/MiddlewareRegistryMock.sol";
import "./mocks/ServiceManagerMock.sol";
import "../contracts/libraries/BeaconChainProofs.sol";
import "./mocks/BeaconChainOracleMock.sol";


contract EigenPodTests is ProofParsing, EigenPodPausingConstants {
    using BytesLib for bytes;

    uint256 internal constant GWEI_TO_WEI = 1e9;

    bytes pubkey = hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";
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
    IBLSPublicKeyCompendium public blsPkCompendium;
    IEigenPodManager public eigenPodManager;
    IEigenPod public podImplementation;
    IDelayedWithdrawalRouter public delayedWithdrawalRouter;
    IETHPOSDeposit public ethPOSDeposit;
    IBeacon public eigenPodBeacon;
    IBeaconChainOracleMock public beaconChainOracle;
    MiddlewareRegistryMock public generalReg1;
    ServiceManagerMock public generalServiceManager1;
    address[] public slashingContracts;
    address pauser = address(69);
    address unpauser = address(489);
    address podManagerAddress = 0x212224D2F2d262cd093eE13240ca4873fcCBbA3C;
    address podAddress = address(123);
    uint256 stakeAmount = 32e18;
    mapping (address => bool) fuzzedAddressMapping;
    bytes signature;
    bytes32 depositDataRoot;

    bytes32[] withdrawalFields;
    bytes32[] validatorFields;


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

    /// @notice Emitted when an ETH validator is proven to have a balance less than `REQUIRED_BALANCE_GWEI` in the beacon chain
    event ValidatorOvercommitted(uint40 validatorIndex);
    
    /// @notice Emitted when an ETH validator is prove to have withdrawn from the beacon chain
    event FullWithdrawalRedeemed(uint40 validatorIndex, address indexed recipient, uint64 withdrawalAmountGwei);

    /// @notice Emitted when a partial withdrawal claim is successfully redeemed
    event PartialWithdrawalRedeemed(uint40 validatorIndex, address indexed recipient, uint64 partialWithdrawalAmountGwei);

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


    uint32 WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
    uint256 REQUIRED_BALANCE_WEI = 31 ether;

    //performs basic deployment before each test
    function setUp() public {
        // deploy proxy admin for ability to upgrade proxy contracts
        eigenLayerProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserReg= new PauserRegistry(pausers, unpauser);

        blsPkCompendium = new BLSPublicKeyCompendium();

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
                REQUIRED_BALANCE_WEI
        );
        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));

        // this contract is deployed later to keep its address the same (for these tests)
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        DelegationManager delegationImplementation = new DelegationManager(strategyManager, slasher);
        StrategyManager strategyManagerImplementation = new StrategyManager(delegation, IEigenPodManager(podManagerAddress), slasher);
        Slasher slasherImplementation = new Slasher(strategyManager, delegation);
        EigenPodManager eigenPodManagerImplementation = new EigenPodManager(ethPOSDeposit, eigenPodBeacon, strategyManager, slasher);

        //ensuring that the address of eigenpodmanager doesn't change
        bytes memory code = address(eigenPodManager).code;
        cheats.etch(podManagerAddress, code);
        eigenPodManager = IEigenPodManager(podManagerAddress);

        beaconChainOracle = new BeaconChainOracleMock();
        DelayedWithdrawalRouter delayedWithdrawalRouterImplementation = new DelayedWithdrawalRouter(IEigenPodManager(podManagerAddress));

        address initialOwner = address(this);
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delegation))),
            address(delegationImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                initialOwner,
                pauserReg,
                0/*initialPausedStatus*/
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
                0/*initialPausedStatus*/,
                0/*withdrawalDelayBlocks*/
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation),
            abi.encodeWithSelector(
                Slasher.initialize.selector,
                initialOwner,
                pauserReg,
                0/*initialPausedStatus*/
            )
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
                0/*initialPausedStatus*/
            )
        );
        uint256 initPausedStatus = 0;
        uint256 withdrawalDelayBlocks = WITHDRAWAL_DELAY_BLOCKS;
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))),
            address(delayedWithdrawalRouterImplementation),
            abi.encodeWithSelector(DelayedWithdrawalRouter.initialize.selector, initialOwner, pauserReg, initPausedStatus, withdrawalDelayBlocks)
        );
        generalServiceManager1 = new ServiceManagerMock(slasher);

        generalReg1 = new MiddlewareRegistryMock(
             generalServiceManager1,
             strategyManager
        );

        cheats.deal(address(podOwner), 5*stakeAmount);     

        fuzzedAddressMapping[address(0)] = true;
        fuzzedAddressMapping[address(eigenLayerProxyAdmin)] = true;
        fuzzedAddressMapping[address(strategyManager)] = true;
        fuzzedAddressMapping[address(eigenPodManager)] = true;
        fuzzedAddressMapping[address(delegation)] = true;
        fuzzedAddressMapping[address(slasher)] = true;
        fuzzedAddressMapping[address(generalServiceManager1)] = true;
        fuzzedAddressMapping[address(generalReg1)] = true;
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
        require(pod.hasRestaked() == false, "Pod should not be restaked");

        // simulate a withdrawal
        cheats.deal(address(pod), stakeAmount);
        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(delayedWithdrawalRouter));
        emit DelayedWithdrawalCreated(podOwner, podOwner, stakeAmount, delayedWithdrawalRouter.userWithdrawalsLength(podOwner));
        pod.withdrawBeforeRestaking();
        require(_getLatestDelayedWithdrawalAmount(podOwner) == stakeAmount, "Payment amount should be stake amount");
        require(pod.mostRecentWithdrawalBlockNumber() == uint64(block.number), "Most recent withdrawal block number not updated");
    }

    function testWithdrawBeforeRestakingAfterRestaking() public {
        // ./solidityProofGen "ValidatorFieldsProof" 61336 true "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61336.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61336.json");
        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        cheats.startPrank(podOwner);
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    function testWithdrawFromPod() public {
        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod pod = eigenPodManager.getPod(podOwner);
        cheats.deal(address(pod), stakeAmount);

        cheats.startPrank(podOwner);
        uint256 userWithdrawalsLength = delayedWithdrawalRouter.userWithdrawalsLength(podOwner);
        // cheats.expectEmit(true, true, true, true, address(delayedWithdrawalRouter));
        cheats.expectEmit(true, true, true, true);
        emit DelayedWithdrawalCreated(podOwner, podOwner, stakeAmount, userWithdrawalsLength);
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();
        require(address(pod).balance == 0, "Pod balance should be 0");
    }

    function testAttemptedWithdrawalAfterVerifyingWithdrawalCredentials() public {
        testDeployAndVerifyNewEigenPod();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        IEigenPod(pod).withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    function testFullWithdrawalProof() public {
        setJSON("./src/test/test-data/fullWithdrawalProof.json");
        BeaconChainProofs.WithdrawalProofs memory proofs = _getWithdrawalProof();
        withdrawalFields = getWithdrawalFields();   
        validatorFields = getValidatorFields();

        Relayer relay = new Relayer();

        bytes32 beaconStateRoot = getBeaconStateRoot();
        relay.verifyWithdrawalProofs(beaconStateRoot, proofs, withdrawalFields);

    }

    /// @notice This test is to ensure the full withdrawal flow works
    function testFullWithdrawalFlow() public returns (IEigenPod) {
        //this call is to ensure that validator 61336 has proven their withdrawalcreds
        // ./solidityProofGen "ValidatorFieldsProof" 61336 true "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61336.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61336.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        // ./solidityProofGen "WithdrawalFieldsProof" 61336 2262 "data/slot_43222/oracle_capella_beacon_state_43300.ssz" "data/slot_43222/capella_block_header_43222.json" "data/slot_43222/capella_block_43222.json" fullWithdrawalProof.json
        setJSON("./src/test/test-data/fullWithdrawalProof.json");
        BeaconChainProofs.WithdrawalProofs memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());
        withdrawalFields = getWithdrawalFields();   
        validatorFields = getValidatorFields();
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);

        uint64 restakedExecutionLayerGweiBefore = newPod.restakedExecutionLayerGwei();
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]);
        uint64 leftOverBalanceWEI = uint64(withdrawalAmountGwei - newPod.REQUIRED_BALANCE_GWEI()) * uint64(GWEI_TO_WEI);
        uint40 validatorIndex = uint40(Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_INDEX_INDEX]));
        cheats.deal(address(newPod), leftOverBalanceWEI);
        
        uint256 delayedWithdrawalRouterContractBalanceBefore = address(delayedWithdrawalRouter).balance;
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit FullWithdrawalRedeemed(validatorIndex, podOwner, withdrawalAmountGwei);
        newPod.verifyAndProcessWithdrawal(withdrawalProofs, validatorFieldsProof, validatorFields, withdrawalFields, 0, 0);
        require(newPod.restakedExecutionLayerGwei() -  restakedExecutionLayerGweiBefore == newPod.REQUIRED_BALANCE_GWEI(),
            "restakedExecutionLayerGwei has not been incremented correctly");
        require(address(delayedWithdrawalRouter).balance - delayedWithdrawalRouterContractBalanceBefore == leftOverBalanceWEI,
            "pod delayed withdrawal balance hasn't been updated correctly");

        cheats.roll(block.number + WITHDRAWAL_DELAY_BLOCKS + 1);
        uint podOwnerBalanceBefore = address(podOwner).balance;
        delayedWithdrawalRouter.claimDelayedWithdrawals(podOwner, 1);
        require(address(podOwner).balance - podOwnerBalanceBefore == leftOverBalanceWEI, "Pod owner balance hasn't been updated correctly");
        return newPod;
    }

    /// @notice This test is to ensure that the partial withdrawal flow works correctly
    function testPartialWithdrawalFlow() public returns(IEigenPod) {
        //this call is to ensure that validator 61068 has proven their withdrawalcreds
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        //generate partialWithdrawalProofs.json with: 
        // ./solidityProofGen "WithdrawalFieldsProof" 61068 656 "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "data/slot_58000/capella_block_header_58000.json" "data/slot_58000/capella_block_58000.json" "partialWithdrawalProof.json"
        setJSON("./src/test/test-data/partialWithdrawalProof.json");
        BeaconChainProofs.WithdrawalProofs memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());

        withdrawalFields = getWithdrawalFields();   
        validatorFields = getValidatorFields();
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);

        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]);
        uint64 slot = Endian.fromLittleEndianUint64(withdrawalProofs.slotRoot);
        uint40 validatorIndex = uint40(Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_INDEX_INDEX]));

        cheats.deal(address(newPod), stakeAmount);    

        uint256 delayedWithdrawalRouterContractBalanceBefore = address(delayedWithdrawalRouter).balance;
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit PartialWithdrawalRedeemed(validatorIndex, podOwner, withdrawalAmountGwei);
        newPod.verifyAndProcessWithdrawal(withdrawalProofs, validatorFieldsProof, validatorFields, withdrawalFields, 0, 0);
        require(newPod.provenPartialWithdrawal(validatorIndex, slot), "provenPartialWithdrawal should be true");
        withdrawalAmountGwei = uint64(withdrawalAmountGwei*GWEI_TO_WEI);
        require(address(delayedWithdrawalRouter).balance - delayedWithdrawalRouterContractBalanceBefore == withdrawalAmountGwei,
            "pod delayed withdrawal balance hasn't been updated correctly");

        cheats.roll(block.number + WITHDRAWAL_DELAY_BLOCKS + 1);
        uint podOwnerBalanceBefore = address(podOwner).balance;
        delayedWithdrawalRouter.claimDelayedWithdrawals(podOwner, 1);
        require(address(podOwner).balance - podOwnerBalanceBefore == withdrawalAmountGwei, "Pod owner balance hasn't been updated correctly");
        return newPod;
    }

    /// @notice verifies that multiple partial withdrawals can be made before a full withdrawal
    function testProvingMultipleWithdrawalsForSameSlot(/*uint256 numPartialWithdrawals*/) public {
        IEigenPod newPod = testPartialWithdrawalFlow();

        BeaconChainProofs.WithdrawalProofs memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());
        withdrawalFields = getWithdrawalFields();   
        validatorFields = getValidatorFields();

        cheats.expectRevert(bytes("EigenPod._processPartialWithdrawal: partial withdrawal has already been proven for this slot"));
        newPod.verifyAndProcessWithdrawal(withdrawalProofs, validatorFieldsProof, validatorFields, withdrawalFields, 0, 0);
    }

    /// @notice verifies that multiple full withdrawals for a single validator fail
    function testDoubleFullWithdrawal() public {
        IEigenPod newPod = testFullWithdrawalFlow();
        BeaconChainProofs.WithdrawalProofs memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());
        withdrawalFields = getWithdrawalFields();   
        validatorFields = getValidatorFields();
        cheats.expectRevert(bytes("EigenPod.verifyBeaconChainFullWithdrawal: VALIDATOR_STATUS is WITHDRAWN or invalid VALIDATOR_STATUS"));
        newPod.verifyAndProcessWithdrawal(withdrawalProofs, validatorFieldsProof, validatorFields, withdrawalFields, 0, 0);
    }

    function testDeployAndVerifyNewEigenPod() public returns(IEigenPod) {
        // ./solidityProofGen "ValidatorFieldsProof" 61068 false "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61068.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        return _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
    }

    // //test freezing operator after a beacon chain slashing event
    function testUpdateSlashedBeaconBalance() public {
        //make initial deposit
        // ./solidityProofGen "ValidatorFieldsProof" 61511 true "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "withdrawalCredentialAndBalanceProof_61511.json"
        setJSON("./src/test/test-data/slashedProofs/notOvercommittedBalanceProof_61511.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        // ./solidityProofGen "ValidatorFieldsProof" 61511 false  "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "withdrawalCredentialAndBalanceProof_61511.json"
        setJSON("./src/test/test-data/slashedProofs/overcommittedBalanceProof_61511.json");
        _proveOverCommittedStake(newPod);
        
        uint256 beaconChainETHShares = strategyManager.stakerStrategyShares(podOwner, strategyManager.beaconChainETHStrategy());

        require(beaconChainETHShares == 0, "strategyManager shares not updated correctly");
    }

    //test deploying an eigen pod with mismatched withdrawal credentials between the proof and the actual pod's address
    function testDeployNewEigenPodWithWrongWithdrawalCreds(address wrongWithdrawalAddress) public {
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod newPod;
        newPod = eigenPodManager.getPod(podOwner);
        // make sure that wrongWithdrawalAddress is not set to actual pod address
        cheats.assume(wrongWithdrawalAddress != address(newPod));

        validatorFields = getValidatorFields();
        validatorFields[1] = abi.encodePacked(bytes1(uint8(1)), bytes11(0), wrongWithdrawalAddress).toBytes32(0);
        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = _getValidatorFieldsAndBalanceProof();
        uint64 blockNumber = 1;

        cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: Proof is not for this EigenPod"));
        newPod.verifyWithdrawalCredentialsAndBalance(blockNumber, validatorIndex0, proofs, validatorFields);
    }

    //test that when withdrawal credentials are verified more than once, it reverts
    function testDeployNewEigenPodWithActiveValidator() public {
        // ./solidityProofGen "ValidatorFieldsProof" 61068 false "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61068.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        uint64 blockNumber = 1;
        uint40 validatorIndex = uint40(getValidatorIndex());
        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = _getValidatorFieldsAndBalanceProof();
        validatorFields = getValidatorFields();
        cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: Validator must be inactive to prove withdrawal credentials"));
        pod.verifyWithdrawalCredentialsAndBalance(blockNumber, validatorIndex, proofs, validatorFields);
    }

    function testVerifyWithdrawalCredentialsWithInadequateBalance() public {
         // ./solidityProofGen "ValidatorFieldsProof" 61068 false "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61068.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = _getValidatorFieldsAndBalanceProof();
        validatorFields = getValidatorFields();
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        uint40 validatorIndex = uint40(getValidatorIndex());
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);


        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        uint64 blockNumber = 1;

        //set the validator balance to less than REQUIRED_BALANCE_WEI
        proofs.balanceRoot = bytes32(0);

        cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: ETH validator's balance must be greater than or equal to the restaked balance per validator"));
        newPod.verifyWithdrawalCredentialsAndBalance(blockNumber, validatorIndex, proofs, validatorFields);
    }

    function testProveOverComittedStakeOnWithdrawnValidator() public {
        // ./solidityProofGen "ValidatorFieldsProof" 61511 true "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "withdrawalCredentialAndBalanceProof_61511.json"
        setJSON("./src/test/test-data/slashedProofs/notOvercommittedBalanceProof_61511.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        // ./solidityProofGen "ValidatorFieldsProof" 61511 false  "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "withdrawalCredentialAndBalanceProof_61511.json"
        //setJSON("./src/test/test-data/slashedProofs/overcommittedBalanceProof_61511.json");
        emit log_named_address("podOwner", podOwner);
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);
        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = _getValidatorFieldsAndBalanceProof();
        //set slashed status to false, and balance to 0
        proofs.balanceRoot = bytes32(0);
        validatorFields[3] = bytes32(0);
        cheats.expectRevert(bytes("EigenPod.verifyOvercommittedStake: Validator must be slashed to be overcommitted"));
        newPod.verifyOvercommittedStake(validatorIndex, proofs, validatorFields, 0, uint64(block.number));

    }

    function getBeaconChainETHShares(address staker) internal view returns(uint256) {
        return strategyManager.stakerStrategyShares(staker, strategyManager.beaconChainETHStrategy());
    }

    // // 3. Single withdrawal credential
    // // Test: Owner proves an withdrawal credential.
    // // Expected Behaviour: beaconChainETH shares should increment by REQUIRED_BALANCE_WEI
    // //                     validator status should be marked as ACTIVE

    function testProveSingleWithdrawalCredential() public {
        // get beaconChainETH shares
        uint256 beaconChainETHBefore = getBeaconChainETHShares(podOwner);

        // ./solidityProofGen "ValidatorFieldsProof" 61068 false "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61068.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        uint40 validatorIndex = uint40(getValidatorIndex());

        uint256 beaconChainETHAfter = getBeaconChainETHShares(pod.podOwner());
        assertTrue(beaconChainETHAfter - beaconChainETHBefore == pod.REQUIRED_BALANCE_WEI());
        assertTrue(pod.validatorStatus(validatorIndex) == IEigenPod.VALIDATOR_STATUS.ACTIVE);
    }

    // // 5. Prove overcommitted balance
    // // Setup: Run (3). 
    // // Test: Watcher proves an overcommitted balance for validator from (3).
    // // Expected Behaviour: beaconChainETH shares should decrement by REQUIRED_BALANCE_WEI
    // //                     validator status should be marked as OVERCOMMITTED

    function testProveOverCommittedBalance() public {
        // ./solidityProofGen "ValidatorFieldsProof" 61511 true "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "withdrawalCredentialAndBalanceProof_61511.json"
        setJSON("./src/test/test-data/slashedProofs/notOvercommittedBalanceProof_61511.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        // get beaconChainETH shares
        uint256 beaconChainETHBefore = getBeaconChainETHShares(podOwner);

        // ./solidityProofGen "ValidatorFieldsProof" 61511 false  "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "withdrawalCredentialAndBalanceProof_61511.json"
        setJSON("./src/test/test-data/slashedProofs/overcommittedBalanceProof_61511.json");
        // prove overcommitted balance
        _proveOverCommittedStake(newPod);

        uint40 validatorIndex = uint40(getValidatorIndex());

        assertTrue(beaconChainETHBefore - getBeaconChainETHShares(podOwner) == newPod.REQUIRED_BALANCE_WEI(), "BeaconChainETHShares not updated");
        assertTrue(newPod.validatorStatus(validatorIndex) == IEigenPod.VALIDATOR_STATUS.OVERCOMMITTED, "validator status not set correctly");
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
        require(pod.hasRestaked() == false, "Pod should not be restaked");

        //simulate a withdrawal
        cheats.startPrank(nonPodOwner);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodOwner: not podOwner"));
        pod.withdrawBeforeRestaking();
    }
    

    function testWithdrawRestakedBeaconChainETHRevertsWhenPaused() external {
        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_WITHDRAW_RESTAKED_ETH);
        cheats.stopPrank();

        address recipient = address(this);
        uint256 amount = 1e18;
        cheats.startPrank(address(eigenPodManager.strategyManager()));
        cheats.expectRevert(bytes("Pausable: index is paused"));
        eigenPodManager.withdrawRestakedBeaconChainETH(podOwner, recipient, amount);
        cheats.stopPrank();
    }

    function testVerifyCorrectWithdrawalCredentialsRevertsWhenPaused() external {
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = _getValidatorFieldsAndBalanceProof();
        validatorFields = getValidatorFields();
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        uint40 validatorIndex = uint40(getValidatorIndex());
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        uint64 blockNumber = 1;

        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_CREDENTIALS);
        cheats.stopPrank();

        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        newPod.verifyWithdrawalCredentialsAndBalance(blockNumber, validatorIndex, proofs, validatorFields);
    }

    function testVerifyOvercommittedStakeRevertsWhenPaused() external {
        // ./solidityProofGen "ValidatorFieldsProof" 61511 true "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "withdrawalCredentialAndBalanceProof_61511.json"
         setJSON("./src/test/test-data/slashedProofs/notOvercommittedBalanceProof_61511.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        // ./solidityProofGen "ValidatorFieldsProof" 61511 false  "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "withdrawalCredentialAndBalanceProof_61511.json"
        setJSON("./src/test/test-data/slashedProofs/overcommittedBalanceProof_61511.json");
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);
        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = _getValidatorFieldsAndBalanceProof();
        

        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_OVERCOMMITTED);
        cheats.stopPrank();

        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        newPod.verifyOvercommittedStake(validatorIndex, proofs, validatorFields, 0, 0);    
    }


    function _proveOverCommittedStake(IEigenPod newPod) internal {
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);
        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = _getValidatorFieldsAndBalanceProof();
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit ValidatorOvercommitted(validatorIndex);
        newPod.verifyOvercommittedStake(validatorIndex, proofs, validatorFields, 0, uint64(block.number));
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
    function testVerifyInclusionSha256FailsForEmptyProof(
        bytes32 root,
        bytes32 leaf,
        uint256 index
    ) public {
        bytes memory emptyProof = new bytes(0);
        cheats.expectRevert(bytes("Merkle.processInclusionProofSha256: proof length should be a non-zero multiple of 32"));
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
        cheats.expectRevert(bytes("Merkle.processInclusionProofSha256: proof length should be a non-zero multiple of 32"));
        Merkle.verifyInclusionSha256(proof, root, leaf, index);
    }

    /// @notice Test that the Merkle proof verification fails when the proof length is empty
    function testVerifyInclusionKeccakFailsForEmptyProof(
        bytes32 root,
        bytes32 leaf,
        uint256 index
    ) public {
        bytes memory emptyProof = new bytes(0);
        cheats.expectRevert(bytes("Merkle.processInclusionProofKeccak: proof length should be a non-zero multiple of 32"));
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
        cheats.expectRevert(bytes("Merkle.processInclusionProofKeccak: proof length should be a non-zero multiple of 32"));
        Merkle.verifyInclusionKeccak(proof, root, leaf, index);
    }

    // verifies that the `numPod` variable increments correctly on a succesful call to the `EigenPod.stake` function
    function test_incrementNumPodsOnStake(bytes calldata _pubkey, bytes calldata _signature, bytes32 _depositDataRoot) public {
        uint256 numPodsBefore = EigenPodManager(address(eigenPodManager)).numPods();
        testStake(_pubkey, _signature, _depositDataRoot);
        uint256 numPodsAfter = EigenPodManager(address(eigenPodManager)).numPods();
        require(numPodsAfter == numPodsBefore + 1, "numPods did not increment correctly");
    }

    // verifies that the `maxPods` variable is enforced on the `EigenPod.stake` function
    function test_maxPodsEnforcementOnStake(bytes calldata _pubkey, bytes calldata _signature, bytes32 _depositDataRoot) public {
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

    // simply tries to register 'sender' as an operator, setting their 'OperatorDetails' in DelegationManager to 'operatorDetails'
    // verifies that the storage of DelegationManager contract is updated appropriately
    function _testRegisterAsOperator(address sender, IDelegationManager.OperatorDetails memory operatorDetails) internal {
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
        (IStrategy[] memory delegateStrategies, uint256[] memory delegateShares) =
            strategyManager.getDeposits(sender);

        uint256 numStrats = delegateShares.length;
        assertTrue(numStrats > 0, "_testDelegateToOperator: delegating from address with no deposits");
        uint256[] memory inititalSharesInStrats = new uint256[](numStrats);
        for (uint256 i = 0; i < numStrats; ++i) {
            inititalSharesInStrats[i] = delegation.operatorShares(operator, delegateStrategies[i]);
        }

        cheats.startPrank(sender);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(operator, signatureWithExpiry);
        cheats.stopPrank();

        assertTrue(
            delegation.delegatedTo(sender) == operator,
            "_testDelegateToOperator: delegated address not set appropriately"
        );
        assertTrue(
            delegation.isDelegated(sender),
            "_testDelegateToOperator: delegated status not set appropriately"
        );

        for (uint256 i = 0; i < numStrats; ++i) {
            uint256 operatorSharesBefore = inititalSharesInStrats[i];
            uint256 operatorSharesAfter = delegation.operatorShares(operator, delegateStrategies[i]);
            assertTrue(
                operatorSharesAfter == (operatorSharesBefore + delegateShares[i]),
                "_testDelegateToOperator: delegatedShares not increased correctly"
            );
        }
    }
    function _testDelegation(address operator, address staker)
        internal
    {   
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
        (updatedStrategies, updatedShares) =
            strategyManager.getDeposits(staker);
    }

    function _testDeployAndVerifyNewEigenPod(address _podOwner, bytes memory _signature, bytes32 _depositDataRoot)
        internal returns (IEigenPod)
    {
        // (beaconStateRoot, beaconStateMerkleProofForValidators, validatorContainerFields, validatorMerkleProof, validatorTreeRoot, validatorRoot) =
        //     getInitialDepositProof(validatorIndex);

        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = _getValidatorFieldsAndBalanceProof();
        validatorFields = getValidatorFields();
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        uint40 validatorIndex = uint40(getValidatorIndex());
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);

        IEigenPod newPod = eigenPodManager.getPod(_podOwner);

        cheats.startPrank(_podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, _signature, _depositDataRoot);
        cheats.stopPrank();

        uint64 blockNumber = 1;
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit ValidatorRestaked(validatorIndex);
        newPod.verifyWithdrawalCredentialsAndBalance(blockNumber, validatorIndex, proofs, validatorFields);

        IStrategy beaconChainETHStrategy = strategyManager.beaconChainETHStrategy();

        uint256 beaconChainETHShares = strategyManager.stakerStrategyShares(_podOwner, beaconChainETHStrategy);
        require(beaconChainETHShares == REQUIRED_BALANCE_WEI, "strategyManager shares not updated correctly");
        return newPod;
    }

    function _testQueueWithdrawal(
        address depositor,
        uint256[] memory strategyIndexes,
        IStrategy[] memory strategyArray,
        uint256[] memory shareAmounts,
        bool undelegateIfPossible
    )
        internal
        returns (bytes32)
    {
        cheats.startPrank(depositor);

        //make a call with depositor aka podOwner also as withdrawer.
        bytes32 withdrawalRoot = strategyManager.queueWithdrawal(
            strategyIndexes,
            strategyArray,
            shareAmounts,
            depositor,
            // TODO: make this an input
            undelegateIfPossible
        );

        cheats.stopPrank();
        return withdrawalRoot;
    }

    function _getLatestDelayedWithdrawalAmount(address recipient) internal view returns (uint256) {
        return delayedWithdrawalRouter.userDelayedWithdrawalByIndex(recipient, delayedWithdrawalRouter.userWithdrawalsLength(recipient) - 1).amount;
    }

    function _getValidatorFieldsAndBalanceProof() internal returns (BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory) {

        bytes32 balanceRoot = getBalanceRoot();
        BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs = BeaconChainProofs.ValidatorFieldsAndBalanceProofs(
            abi.encodePacked(getWithdrawalCredentialProof()),
            abi.encodePacked(getValidatorBalanceProof()),
            balanceRoot
        );

        return proofs;
    }

    /// @notice this function just generates a valid proof so that we can test other functionalities of the withdrawal flow
    function _getWithdrawalProof() internal returns(BeaconChainProofs.WithdrawalProofs memory) {
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        //make initial deposit
        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        
        {
            bytes32 beaconStateRoot = getBeaconStateRoot();
            //set beaconStateRoot
            beaconChainOracle.setBeaconChainStateRoot(beaconStateRoot);
            bytes32 blockHeaderRoot = getBlockHeaderRoot();
            bytes32 blockBodyRoot = getBlockBodyRoot();
            bytes32 slotRoot = getSlotRoot();
            bytes32 blockNumberRoot = getBlockNumberRoot();
            bytes32 executionPayloadRoot = getExecutionPayloadRoot();



            uint256 withdrawalIndex = getWithdrawalIndex();
            uint256 blockHeaderRootIndex = getBlockHeaderRootIndex();


            BeaconChainProofs.WithdrawalProofs memory proofs = BeaconChainProofs.WithdrawalProofs(
                abi.encodePacked(getBlockHeaderProof()),
                abi.encodePacked(getWithdrawalProof()),
                abi.encodePacked(getSlotProof()),
                abi.encodePacked(getExecutionPayloadProof()),
                abi.encodePacked(getBlockNumberProof()),
                uint64(blockHeaderRootIndex),
                uint64(withdrawalIndex),
                blockHeaderRoot,
                blockBodyRoot,
                slotRoot,
                blockNumberRoot,
                executionPayloadRoot
            );
            return proofs;
        }
    }

    function _getValidatorFieldsProof() internal returns(BeaconChainProofs.ValidatorFieldsProof memory) {
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        //make initial deposit
        cheats.startPrank(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        
        {
            bytes32 beaconStateRoot = getBeaconStateRoot();
            //set beaconStateRoot
            beaconChainOracle.setBeaconChainStateRoot(beaconStateRoot);
            uint256 validatorIndex = getValidatorIndex(); 
            BeaconChainProofs.ValidatorFieldsProof memory proofs = BeaconChainProofs.ValidatorFieldsProof(
                abi.encodePacked(getValidatorProof()),
                uint40(validatorIndex)
            );
            return proofs;
        }
    }

 }


 contract Relayer is Test {
    function verifyWithdrawalProofs(
        bytes32 beaconStateRoot,
        BeaconChainProofs.WithdrawalProofs calldata proofs,
        bytes32[] calldata withdrawalFields
    ) public view {
        BeaconChainProofs.verifyWithdrawalProofs(beaconStateRoot, proofs, withdrawalFields);
    }
 }