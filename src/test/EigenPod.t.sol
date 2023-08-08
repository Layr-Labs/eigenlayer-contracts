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

    /// @notice Emitted when an ETH validator's balance is updated in EigenLayer
    event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 newBalanceGwei);

    
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
    uint256 REQUIRED_BALANCE_WEI = 32 ether;
    uint64  MAX_VALIDATOR_BALANCE_GWEI = 32e9;
    uint64  EFFECTIVE_RESTAKED_BALANCE_OFFSET = 75e7;

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
                MAX_VALIDATOR_BALANCE_GWEI,
                EFFECTIVE_RESTAKED_BALANCE_OFFSET
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
        require(pod.mostRecentWithdrawalTimestamp() == uint64(block.timestamp), "Most recent withdrawal block number not updated");
    }

    function testDeployEigenPodWithoutActivateRestaking() public {
        // ./solidityProofGen "ValidatorFieldsProof" 61336 true "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61336.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61336.json");

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        uint64 timestamp = 0;
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        BeaconChainProofs.WithdrawalCredentialProofs[] memory proofsArray = new BeaconChainProofs.WithdrawalCredentialProofs[](1);
        proofsArray[0] = _getWithdrawalCredentialProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());


        cheats.startPrank(podOwner);
        cheats.warp(timestamp += 1);
        cheats.expectRevert(bytes("EigenPod.hasEnabledRestaking: restaking is not enabled"));
        newPod.verifyWithdrawalCredentials(timestamp, validatorIndices, proofsArray, validatorFieldsArray);
        cheats.stopPrank();
    }

    function testDeployEigenPodTooSoon() public {
        // ./solidityProofGen "ValidatorFieldsProof" 61336 true "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61336.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61336.json");

        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        uint64 timestamp = 0;
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        BeaconChainProofs.WithdrawalCredentialProofs[] memory proofsArray = new BeaconChainProofs.WithdrawalCredentialProofs[](1);
        proofsArray[0] = _getWithdrawalCredentialProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());


        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.proofIsForValidTimestamp: beacon chain proof must be for timestamp after mostRecentWithdrawalTimestamp"));
        newPod.verifyWithdrawalCredentials(timestamp, validatorIndices, proofsArray, validatorFieldsArray);
        cheats.stopPrank();
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
        relay.verifyWithdrawalProofs(beaconStateRoot, withdrawalFields, proofs);
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
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(getLatestBlockHeaderRoot());
        uint64 restakedExecutionLayerGweiBefore = newPod.withdrawableRestakedExecutionLayerGwei();
       
        withdrawalFields = getWithdrawalFields();
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]);

        uint64 leftOverBalanceWEI = uint64(withdrawalAmountGwei - _getEffectiveRestakedBalanceGwei(newPod.MAX_VALIDATOR_BALANCE_GWEI())) * uint64(GWEI_TO_WEI);
        uint40 validatorIndex = uint40(Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_INDEX_INDEX]));
        cheats.deal(address(newPod), leftOverBalanceWEI);
        emit log_named_uint("leftOverBalanceWEI", leftOverBalanceWEI);
        emit log_named_uint("address(newPod)", address(newPod).balance);
        
        uint256 delayedWithdrawalRouterContractBalanceBefore = address(delayedWithdrawalRouter).balance;
        {
            BeaconChainProofs.WithdrawalProofs[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProofs[](1);
            withdrawalProofsArray[0] = _getWithdrawalProof();
            bytes[] memory validatorFieldsProofArray = new bytes[](1);
            validatorFieldsProofArray[0] = abi.encodePacked(getValidatorProof());
            bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
            validatorFieldsArray[0] = getValidatorFields();
            bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
            withdrawalFieldsArray[0] = withdrawalFields;


            //cheats.expectEmit(true, true, true, true, address(newPod));
            emit FullWithdrawalRedeemed(validatorIndex, podOwner, withdrawalAmountGwei);
            newPod.verifyAndProcessWithdrawals(withdrawalProofsArray, validatorFieldsProofArray, validatorFieldsArray, withdrawalFieldsArray, 0, 0);
        }
        require(newPod.withdrawableRestakedExecutionLayerGwei() -  restakedExecutionLayerGweiBefore == _getEffectiveRestakedBalanceGwei(newPod.MAX_VALIDATOR_BALANCE_GWEI()),
            "restakedExecutionLayerGwei has not been incremented correctly");
        require(address(delayedWithdrawalRouter).balance - delayedWithdrawalRouterContractBalanceBefore == leftOverBalanceWEI,
            "pod delayed withdrawal balance hasn't been updated correctly");
        require(newPod.validatorPubkeyHashToInfo(getValidatorPubkeyHash()).restakedBalanceGwei == 0, "balance not reset correctly");

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
        withdrawalFields = getWithdrawalFields();
        validatorFields = getValidatorFields();
        BeaconChainProofs.WithdrawalProofs memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());

        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(getLatestBlockHeaderRoot());
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]);
        uint40 validatorIndex = uint40(Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_INDEX_INDEX]));

        cheats.deal(address(newPod), stakeAmount);    
        {
            BeaconChainProofs.WithdrawalProofs[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProofs[](1);
            withdrawalProofsArray[0] = withdrawalProofs;
            bytes[] memory validatorFieldsProofArray = new bytes[](1);
            validatorFieldsProofArray[0] = validatorFieldsProof;
            bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
            validatorFieldsArray[0] = validatorFields;
            bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
            withdrawalFieldsArray[0] = withdrawalFields;

            uint256 delayedWithdrawalRouterContractBalanceBefore = address(delayedWithdrawalRouter).balance;
            cheats.expectEmit(true, true, true, true, address(newPod));
            emit PartialWithdrawalRedeemed(validatorIndex, podOwner, withdrawalAmountGwei);
            newPod.verifyAndProcessWithdrawals(withdrawalProofsArray, validatorFieldsProofArray, validatorFieldsArray, withdrawalFieldsArray, 0, 0);
            require(newPod.provenWithdrawal(validatorFields[0], Endian.fromLittleEndianUint64(withdrawalProofs.slotRoot)), "provenPartialWithdrawal should be true");
            withdrawalAmountGwei = uint64(withdrawalAmountGwei*GWEI_TO_WEI);
            require(address(delayedWithdrawalRouter).balance - delayedWithdrawalRouterContractBalanceBefore == withdrawalAmountGwei,
                "pod delayed withdrawal balance hasn't been updated correctly");
        }

        cheats.roll(block.number + WITHDRAWAL_DELAY_BLOCKS + 1);
        uint podOwnerBalanceBefore = address(podOwner).balance;
        delayedWithdrawalRouter.claimDelayedWithdrawals(podOwner, 1);
        require(address(podOwner).balance - podOwnerBalanceBefore == withdrawalAmountGwei, "Pod owner balance hasn't been updated correctly");
        return newPod;
    }

    /// @notice verifies that multiple partial withdrawals can be made before a full withdrawal
    function testProvingMultiplePartialWithdrawalsForSameSlot(/*uint256 numPartialWithdrawals*/) public {
        IEigenPod newPod = testPartialWithdrawalFlow();

        BeaconChainProofs.WithdrawalProofs memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());
        withdrawalFields = getWithdrawalFields();   
        validatorFields = getValidatorFields();

        BeaconChainProofs.WithdrawalProofs[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProofs[](1);
        withdrawalProofsArray[0] = withdrawalProofs;
        bytes[] memory validatorFieldsProofArray = new bytes[](1);
        validatorFieldsProofArray[0] = validatorFieldsProof;
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = validatorFields;
        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
        withdrawalFieldsArray[0] = withdrawalFields;

        cheats.expectRevert(bytes("EigenPod._verifyAndProcessWithdrawal: withdrawal has already been proven for this slot"));
        newPod.verifyAndProcessWithdrawals(withdrawalProofsArray, validatorFieldsProofArray, validatorFieldsArray, withdrawalFieldsArray, 0, 0);
    }

    /// @notice verifies that multiple full withdrawals for a single validator fail
    function testDoubleFullWithdrawal() public returns(IEigenPod newPod) {
        IEigenPod newPod = testFullWithdrawalFlow();
        BeaconChainProofs.WithdrawalProofs memory withdrawalProofs = _getWithdrawalProof();
        bytes memory validatorFieldsProof = abi.encodePacked(getValidatorProof());
        withdrawalFields = getWithdrawalFields();   
        validatorFields = getValidatorFields();
        uint256 beaconChainSharesBefore = getBeaconChainETHShares(podOwner);
        uint256 withdrawableRestakedGwei = newPod.withdrawableRestakedExecutionLayerGwei();
        
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(withdrawalFields[BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]);
        uint64 leftOverBalanceWEI = uint64(withdrawalAmountGwei - newPod.MAX_VALIDATOR_BALANCE_GWEI()) * uint64(GWEI_TO_WEI);
        cheats.deal(address(newPod), leftOverBalanceWEI);

        BeaconChainProofs.WithdrawalProofs[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProofs[](1);
        withdrawalProofsArray[0] = withdrawalProofs;
        bytes[] memory validatorFieldsProofArray = new bytes[](1);
        validatorFieldsProofArray[0] = validatorFieldsProof;
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = validatorFields;
        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
        withdrawalFieldsArray[0] = withdrawalFields;

        cheats.expectRevert(bytes("EigenPod._verifyAndProcessWithdrawal: withdrawal has already been proven for this slot"));
        newPod.verifyAndProcessWithdrawals(withdrawalProofsArray, validatorFieldsProofArray, validatorFieldsArray, withdrawalFieldsArray, 0, 0);
        
        return newPod;
    }

    function testDeployAndVerifyNewEigenPod() public returns(IEigenPod) {
        // ./solidityProofGen "ValidatorFieldsProof" 61068 false "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61068.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        return _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
    }

    // //test freezing operator after a beacon chain slashing event
    function testUpdateSlashedBeaconBalance() public {
        //make initial deposit
        // ./solidityProofGen "BalanceProof" 61511 true "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_notOvercommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_notOvercommitted_61511.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        // ./solidityProofGen "BalanceProof" 61511 false  "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_Overcommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_Overcommitted_61511.json");
        _proveOverCommittedStake(newPod);

        uint64 newValidatorBalance = BeaconChainProofs.getBalanceFromBalanceRoot(uint40(getValidatorIndex()), getBalanceRoot());        
        uint256 beaconChainETHShares = strategyManager.stakerStrategyShares(podOwner, strategyManager.beaconChainETHStrategy());

        require(beaconChainETHShares == _getEffectiveRestakedBalanceGwei(newValidatorBalance) * GWEI_TO_WEI, "strategyManager shares not updated correctly");
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
        uint64 timestamp = 0;

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = validatorFields;
        BeaconChainProofs.WithdrawalCredentialProofs[] memory proofsArray = new BeaconChainProofs.WithdrawalCredentialProofs[](1);
        proofsArray[0] = _getWithdrawalCredentialProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(validatorIndex0);


        cheats.startPrank(podOwner);
        cheats.warp(timestamp);
        newPod.activateRestaking();
        cheats.warp(timestamp += 1);
        cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: Proof is not for this EigenPod"));
        newPod.verifyWithdrawalCredentials(timestamp, validatorIndices, proofsArray, validatorFieldsArray);
        cheats.stopPrank();
    }

    function testVerifyWithdrawalCredsFromNonPodOwnerAddress(address nonPodOwnerAddress) public {
        require(nonPodOwnerAddress != podOwner, "nonPodOwnerAddress must be different from podOwner");
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        cheats.startPrank(podOwner);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();

        IEigenPod newPod = eigenPodManager.getPod(podOwner);
 
        uint64 timestamp = 1;

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        BeaconChainProofs.WithdrawalCredentialProofs[] memory proofsArray = new BeaconChainProofs.WithdrawalCredentialProofs[](1);
        proofsArray[0] = _getWithdrawalCredentialProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(validatorIndex0);

        cheats.startPrank(nonPodOwnerAddress);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodOwner: not podOwner"));
        newPod.verifyWithdrawalCredentials(timestamp, validatorIndices, proofsArray, validatorFieldsArray);
        cheats.stopPrank();
    }

    //test that when withdrawal credentials are verified more than once, it reverts
    function testDeployNewEigenPodWithActiveValidator() public {
        // ./solidityProofGen "ValidatorFieldsProof" 61068 false "data/slot_58000/oracle_capella_beacon_state_58100.ssz" "withdrawalCredentialAndBalanceProof_61068.json"
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61068.json");
        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        uint64 timestamp = 1;

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        BeaconChainProofs.WithdrawalCredentialProofs[] memory proofsArray = new BeaconChainProofs.WithdrawalCredentialProofs[](1);
        proofsArray[0] = _getWithdrawalCredentialProof();
        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.verifyCorrectWithdrawalCredentials: Validator must be inactive to prove withdrawal credentials"));
        pod.verifyWithdrawalCredentials(timestamp, validatorIndices, proofsArray, validatorFieldsArray);
        cheats.stopPrank();
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
        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();


        uint256 beaconChainETHAfter = getBeaconChainETHShares(pod.podOwner());
        emit log_named_uint("beaconChainETHBefore", beaconChainETHBefore);
        emit log_named_uint("beaconChainETHAfter", beaconChainETHAfter);
        assertTrue(beaconChainETHAfter - beaconChainETHBefore == _getEffectiveRestakedBalanceGwei(pod.MAX_VALIDATOR_BALANCE_GWEI())*GWEI_TO_WEI, "pod balance not updated correcty");
        assertTrue(pod.validatorStatus(validatorPubkeyHash) == IEigenPod.VALIDATOR_STATUS.ACTIVE, "wrong validator status");
    }

    // // 5. Prove overcommitted balance
    // // Setup: Run (3). 
    // // Test: Watcher proves an overcommitted balance for validator from (3).
    // // Expected Behaviour: beaconChainETH shares should decrement by REQUIRED_BALANCE_WEI
    // //                     validator status should be marked as OVERCOMMITTED

    function testProveOverCommittedBalance() public {
        // ./solidityProofGen "BalanceUpdateProof" 61511 true 0 "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_notOvercommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_notOvercommitted_61511.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        // get beaconChainETH shares
        uint256 beaconChainETHBefore = getBeaconChainETHShares(podOwner);

        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();
        uint256 validatorRestakedBalanceBefore = newPod.validatorPubkeyHashToInfo(validatorPubkeyHash).restakedBalanceGwei;
        

        // ./solidityProofGen "BalanceUpdateProof" 61511 false 0 "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_Overcommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_Overcommitted_61511.json");
        // prove overcommitted balance
        _proveOverCommittedStake(newPod);

        uint256 validatorRestakedBalanceAfter = newPod.validatorPubkeyHashToInfo(validatorPubkeyHash).restakedBalanceGwei;

        uint64 newValidatorBalance = BeaconChainProofs.getBalanceFromBalanceRoot(uint40(getValidatorIndex()), getBalanceRoot());        
        uint256 shareDiff = beaconChainETHBefore - getBeaconChainETHShares(podOwner);
 
        assertTrue(getBeaconChainETHShares(podOwner) == _getEffectiveRestakedBalanceGwei(newValidatorBalance) * GWEI_TO_WEI, "hysterisis not working");
        assertTrue(beaconChainETHBefore - getBeaconChainETHShares(podOwner) == shareDiff, "BeaconChainETHShares not updated");
        assertTrue(validatorRestakedBalanceBefore - validatorRestakedBalanceAfter  == shareDiff/GWEI_TO_WEI, "validator restaked balance not updated");
    }

    function testVerifyUndercommittedBalance() public {
        // ./solidityProofGen "BalanceUpdateProof" 61511 true "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_notOvercommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_notOvercommitted_61511.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        // get beaconChainETH shares
        uint256 beaconChainETHBefore = getBeaconChainETHShares(podOwner);
        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();
        uint256 validatorRestakedBalanceBefore = newPod.validatorPubkeyHashToInfo(validatorPubkeyHash).restakedBalanceGwei;

        // ./solidityProofGen "BalanceUpdateProof" 61511 false  "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_Overcommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_Overcommitted_61511.json");
        // prove overcommitted balance
        _proveOverCommittedStake(newPod);

        // ./solidityProofGen "BalanceUpdateProof" 61511 true 100  "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_notOvercommitted_61511_incrementedBlockBy100.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_notOvercommitted_61511_incrementedBlockBy100.json");
        _proveUnderCommittedStake(newPod);

        uint256 validatorRestakedBalanceAfter = newPod.validatorPubkeyHashToInfo(validatorPubkeyHash).restakedBalanceGwei;

        uint64 newValidatorBalance = BeaconChainProofs.getBalanceFromBalanceRoot(uint40(getValidatorIndex()), getBalanceRoot());        
        uint256 shareDiff = beaconChainETHBefore - getBeaconChainETHShares(podOwner);
        
        assertTrue(getBeaconChainETHShares(podOwner) == _getEffectiveRestakedBalanceGwei(newValidatorBalance) * GWEI_TO_WEI, "hysterisis not working");
        assertTrue(beaconChainETHBefore - getBeaconChainETHShares(podOwner) == shareDiff, "BeaconChainETHShares not updated");
        assertTrue(validatorRestakedBalanceBefore - validatorRestakedBalanceAfter  == shareDiff/GWEI_TO_WEI, "validator restaked balance not updated");    
    }

    function testTooSoonBalanceUpdates() public {
        // ./solidityProofGen "BalanceUpdateProof" 61511 true 0 "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_notOvercommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_notOvercommitted_61511.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        // get beaconChainETH shares
        uint256 beaconChainETHBefore = getBeaconChainETHShares(podOwner);
        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();
        uint256 validatorRestakedBalanceBefore = newPod.validatorPubkeyHashToInfo(validatorPubkeyHash).restakedBalanceGwei;

        // ./solidityProofGen "BalanceUpdateProof" 61511 false 0 "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_Overcommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_Overcommitted_61511.json");
        // prove overcommitted balance
        _proveOverCommittedStake(newPod);

        // ./solidityProofGen "BalanceUpdateProof" 61511 true 100  "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_notOvercommitted_61511_incrementedBlockBy100.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_notOvercommitted_61511.json");
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);
        BeaconChainProofs.BalanceUpdateProofs memory proofs = _getBalanceUpdateProofs();

        cheats.expectRevert(bytes("EigenPod.verifyBalanceUpdate: Validator's balance has already been updated for this slot"));
        newPod.verifyBalanceUpdate(validatorIndex, proofs, validatorFields, 0, uint64(block.number));
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
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);

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

        BeaconChainProofs.WithdrawalCredentialProofs[] memory proofsArray = new BeaconChainProofs.WithdrawalCredentialProofs[](1);
        proofsArray[0] = _getWithdrawalCredentialProof();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        newPod.verifyWithdrawalCredentials(timestamp, validatorIndices, proofsArray, validatorFieldsArray);
        cheats.stopPrank();
    }

    function testVerifyOvercommittedStakeRevertsWhenPaused() external {
        // ./solidityProofGen "BalanceUpdateProof" 61511 true "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_notOvercommitted_61511.json"
         setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_notOvercommitted_61511.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        // ./solidityProofGen "BalanceUpdateProof" 61511 false "data/slot_209635/oracle_capella_beacon_state_209635.ssz" "balanceUpdateProof_Overcommitted_61511.json"
        setJSON("./src/test/test-data/slashedProofs/balanceUpdateProof_Overcommitted_61511.json");
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newBeaconStateRoot = getBeaconStateRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);
        BeaconChainProofs.BalanceUpdateProofs memory proofs = _getBalanceUpdateProofs();        

        // pause the contract
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE);
        cheats.stopPrank();

        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        newPod.verifyBalanceUpdate(validatorIndex, proofs, validatorFields, 0, 0);    
    }


    function _proveOverCommittedStake(IEigenPod newPod) internal {
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newLatestBlockHeaderRoot = getLatestBlockHeaderRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newLatestBlockHeaderRoot);
        BeaconChainProofs.BalanceUpdateProofs memory proofs = _getBalanceUpdateProofs();
        //cheats.expectEmit(true, true, true, true, address(newPod));
        emit ValidatorBalanceUpdated(validatorIndex, _getEffectiveRestakedBalanceGwei(Endian.fromLittleEndianUint64(validatorFields[BeaconChainProofs.VALIDATOR_BALANCE_INDEX])));
        newPod.verifyBalanceUpdate(validatorIndex, proofs, validatorFields, 0, uint64(block.number));
    }

    function _proveUnderCommittedStake(IEigenPod newPod) internal {
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newLatestBlockHeaderRoot = getLatestBlockHeaderRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newLatestBlockHeaderRoot);
        BeaconChainProofs.BalanceUpdateProofs memory proofs = _getBalanceUpdateProofs();
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit ValidatorBalanceUpdated(validatorIndex, _getEffectiveRestakedBalanceGwei(Endian.fromLittleEndianUint64(validatorFields[BeaconChainProofs.VALIDATOR_BALANCE_INDEX])));
        newPod.verifyBalanceUpdate(validatorIndex, proofs, validatorFields, 0, uint64(block.number));
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

    function testQueueBeaconChainETHWithdrawalWithoutProvingFullWithdrawal(bytes memory signature, bytes32 depositDataRoot) external {
        setJSON("./src/test/test-data/withdrawalCredentialAndBalanceProof_61336.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategyManager.beaconChainETHStrategy();
        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = 31e18;
        bool undelegateIfPossible = false;
        cheats.expectRevert("EigenPod.decrementWithdrawableRestakedExecutionLayerGwei: amount to decrement is greater than current withdrawableRestakedRxecutionLayerGwei balance");
        _testQueueWithdrawal(podOwner, strategyIndexes, strategyArray, shareAmounts, undelegateIfPossible);
    }

    function testQueueBeaconChainETHWithdrawal123(bytes memory signature, bytes32 depositDataRoot) external {
        IEigenPod pod = testFullWithdrawalFlow();

        bytes32 validatorPubkeyHash = getValidatorPubkeyHash();

        uint256 withdrawableRestakedExecutionLayerGweiBefore = pod.withdrawableRestakedExecutionLayerGwei();
        
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategyManager.beaconChainETHStrategy();
        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = _getEffectiveRestakedBalanceGwei(pod.MAX_VALIDATOR_BALANCE_GWEI()) * GWEI_TO_WEI;
        bool undelegateIfPossible = false;
        _verifyEigenPodBalanceSharesInvariant(podOwner, pod, validatorPubkeyHash);
        _testQueueWithdrawal(podOwner, strategyIndexes, strategyArray, shareAmounts, undelegateIfPossible);
        _verifyEigenPodBalanceSharesInvariant(podOwner, pod, validatorPubkeyHash);

        require(withdrawableRestakedExecutionLayerGweiBefore - pod.withdrawableRestakedExecutionLayerGwei() == shareAmounts[0]/GWEI_TO_WEI, "withdrawableRestakedExecutionLayerGwei not decremented correctly");
    }

    function _verifyEigenPodBalanceSharesInvariant(address podOwner, IEigenPod pod, bytes32 validatorPubkeyHash) internal {
        uint256 sharesInSM = strategyManager.stakerStrategyShares(podOwner, strategyManager.beaconChainETHStrategy());
        uint64 withdrawableRestakedExecutionLayerGwei = pod.withdrawableRestakedExecutionLayerGwei();
        
        EigenPod.ValidatorInfo memory info = pod.validatorPubkeyHashToInfo(validatorPubkeyHash);

        uint64 validatorBalanceGwei = info.restakedBalanceGwei;
        emit log_named_uint("shares", sharesInSM/GWEI_TO_WEI);
        emit log_named_uint("validatorBalanceGwei", validatorBalanceGwei);
        emit log_named_uint("withdrawableRestakedExecutionLayerGwei", withdrawableRestakedExecutionLayerGwei);
        require(sharesInSM/GWEI_TO_WEI == validatorBalanceGwei + withdrawableRestakedExecutionLayerGwei, "EigenPod invariant violated: sharesInSM != withdrawableRestakedExecutionLayerGwei");
    }

    // simply tries to register 'sender' as a delegate, setting their 'DelegationTerms' contract in DelegationManager to 'dt'
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

        // bytes32 newBeaconStateRoot = getBeaconStateRoot();
        // BeaconChainOracleMock(address(beaconChainOracle)).setBeaconChainStateRoot(newBeaconStateRoot);

        IEigenPod newPod = eigenPodManager.getPod(_podOwner);

        cheats.startPrank(_podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, _signature, _depositDataRoot);
        cheats.stopPrank();

        uint64 timestamp = 0;
        // cheats.expectEmit(true, true, true, true, address(newPod));
        // emit ValidatorRestaked(validatorIndex);

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();

        BeaconChainProofs.WithdrawalCredentialProofs[] memory proofsArray = new BeaconChainProofs.WithdrawalCredentialProofs[](1);
        proofsArray[0] = _getWithdrawalCredentialProof();

        uint40[] memory validatorIndices = new uint40[](1);
        validatorIndices[0] = uint40(getValidatorIndex());


        cheats.startPrank(_podOwner);
        cheats.warp(timestamp);
        newPod.activateRestaking();
        emit log_named_bytes32("restaking activated", BeaconChainOracleMock(address(beaconChainOracle)).mockBeaconChainStateRoot());
        cheats.warp(timestamp += 1);
        newPod.verifyWithdrawalCredentials(timestamp, validatorIndices, proofsArray, validatorFieldsArray);
        IStrategy beaconChainETHStrategy = strategyManager.beaconChainETHStrategy();
        cheats.stopPrank();

        uint256 beaconChainETHShares = strategyManager.stakerStrategyShares(_podOwner, beaconChainETHStrategy);
        uint256 effectiveBalance = uint256(_getEffectiveRestakedBalanceGwei(uint64(REQUIRED_BALANCE_WEI/GWEI_TO_WEI))) * GWEI_TO_WEI;
        require(beaconChainETHShares == effectiveBalance, "strategyManager shares not updated correctly");
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

    function _getBalanceUpdateProofs() internal returns (BeaconChainProofs.BalanceUpdateProofs memory) {

        bytes32 beaconStateRoot = getBeaconStateRoot();
        bytes32 latestBlockHeaderRoot = getLatestBlockHeaderRoot();

        bytes32 balanceRoot = getBalanceRoot();
        bytes32 slotRoot = getSlotRoot();
        BeaconChainProofs.BalanceUpdateProofs memory proofs = BeaconChainProofs.BalanceUpdateProofs(
            beaconStateRoot,
            abi.encodePacked(getLatestBlockHeaderProof()),
            abi.encodePacked(getValidatorBalanceProof()),
            abi.encodePacked(getWithdrawalCredentialProof()),  //technically this is to verify validator pubkey in the validator fields, but the WC proof is effectively the same so we use it here again.
            abi.encodePacked(getBalanceUpdateSlotProof()),
            balanceRoot,
            slotRoot
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
            bytes32 latestBlockHeaderRoot = getLatestBlockHeaderRoot();
            //set beaconStateRoot
            beaconChainOracle.setBeaconChainStateRoot(latestBlockHeaderRoot);
            bytes32 blockHeaderRoot = getBlockHeaderRoot();
            bytes32 blockBodyRoot = getBlockBodyRoot();
            bytes32 slotRoot = getSlotRoot();
            bytes32 timestampRoot = getTimestampRoot();
            bytes32 executionPayloadRoot = getExecutionPayloadRoot();



            uint256 withdrawalIndex = getWithdrawalIndex();
            uint256 blockHeaderRootIndex = getBlockHeaderRootIndex();


            BeaconChainProofs.WithdrawalProofs memory proofs = BeaconChainProofs.WithdrawalProofs(
                beaconStateRoot,
                abi.encodePacked(getLatestBlockHeaderProof()),
                abi.encodePacked(getBlockHeaderProof()),
                abi.encodePacked(getWithdrawalProof()),
                abi.encodePacked(getSlotProof()),
                abi.encodePacked(getExecutionPayloadProof()),
                abi.encodePacked(getTimestampProof()),
                uint64(blockHeaderRootIndex),
                uint64(withdrawalIndex),
                blockHeaderRoot,
                blockBodyRoot,
                slotRoot,
                timestampRoot,
                executionPayloadRoot
            );
            return proofs;
        }
    }

    function _getWithdrawalCredentialProof() internal returns(BeaconChainProofs.WithdrawalCredentialProofs memory) {        
        {
            bytes32 latestBlockHeaderRoot = getLatestBlockHeaderRoot();
            //set beaconStateRoot
            beaconChainOracle.setBeaconChainStateRoot(latestBlockHeaderRoot);
            uint256 validatorIndex = getValidatorIndex(); 

            BeaconChainProofs.WithdrawalCredentialProofs memory proofs = BeaconChainProofs.WithdrawalCredentialProofs(
                getBeaconStateRoot(),
                abi.encodePacked(getLatestBlockHeaderProof()),
                abi.encodePacked(getWithdrawalCredentialProof())
            );
            return proofs;
        }
    }

    function testEffectiveRestakedBalance() public {
        uint64 amountGwei = 29134000000;
        uint64 offset =       750000000;
        uint64 effectiveBalance = _getEffectiveRestakedBalanceGwei(amountGwei);
        emit log_named_uint("effectiveBalance", effectiveBalance);
    }

    function _getEffectiveRestakedBalanceGwei(uint64 amountGwei) internal returns (uint64){
        if(amountGwei < 75e7) {
            return 0;
        }
        //calculates the "floor" of amountGwei - EFFECTIVE_RESTAKED_BALANCE_OFFSET
        uint64 effectiveBalance = uint64((amountGwei - 75e7) / GWEI_TO_WEI * GWEI_TO_WEI);
        return uint64(MathUpgradeable.min(32e9, effectiveBalance));
    }

 }


 contract Relayer is Test {
    function verifyWithdrawalProofs(
        bytes32 beaconStateRoot,
        bytes32[] calldata withdrawalFields,
        BeaconChainProofs.WithdrawalProofs calldata proofs
    ) public view {
        BeaconChainProofs.verifyWithdrawalProofs(beaconStateRoot, withdrawalFields, proofs);
    }
 }