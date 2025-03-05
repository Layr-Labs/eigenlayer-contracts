// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodPausingConstants.sol";

import "src/test/mocks/ETHDepositMock.sol";
import "src/test/mocks/ERC20Mock.sol";
import "src/test/harnesses/EigenPodHarness.sol";
import "src/test/utils/ProofParsing.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

import "src/test/integration/mocks/BeaconChainMock.t.sol";
import "src/test/integration/mocks/EIP_4788_Oracle_Mock.t.sol";
import "src/test/utils/EigenPodUser.t.sol";

contract EigenPodUnitTests is EigenLayerUnitTestSetup, EigenPodPausingConstants, IEigenPodEvents {
    using BytesLib for bytes;
    using BeaconChainProofs for *;

    // Contract Under Test: EigenPod
    EigenPod public eigenPod;
    EigenPod public podImplementation;
    IBeacon public eigenPodBeacon;

    // BeaconChain Mock Setup
    TimeMachine public timeMachine;
    ETHPOSDepositMock ethPOSDepositMock;
    BeaconChainMock public beaconChain;
    EIP_4788_Oracle_Mock constant EIP_4788_ORACLE = EIP_4788_Oracle_Mock(0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02);

    uint public numStakers;

    address defaultProofSubmitter = cheats.addr(uint(0xDEADBEEF));

    // Beacon chain genesis time when running locally
    // Multiple of 12 for sanity's sake
    uint64 constant GENESIS_TIME_LOCAL = 1 hours * 12;
    uint constant TIME_TILL_STALE_BALANCE = 2 weeks;

    bytes internal constant beaconProxyBytecode =
        hex"608060405260405161090e38038061090e83398101604081905261002291610460565b61002e82826000610035565b505061058a565b61003e83610100565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100fb576100f9836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e99190610520565b836102a360201b6100291760201c565b505b505050565b610113816102cf60201b6100551760201c565b6101725760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101e6816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d79190610520565b6102cf60201b6100551760201c565b61024b5760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610169565b806102827fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6102de60201b6100641760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606102c883836040518060600160405280602781526020016108e7602791396102e1565b9392505050565b6001600160a01b03163b151590565b90565b6060600080856001600160a01b0316856040516102fe919061053b565b600060405180830381855af49150503d8060008114610339576040519150601f19603f3d011682016040523d82523d6000602084013e61033e565b606091505b5090925090506103508683838761035a565b9695505050505050565b606083156103c65782516103bf576001600160a01b0385163b6103bf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610169565b50816103d0565b6103d083836103d8565b949350505050565b8151156103e85781518083602001fd5b8060405162461bcd60e51b81526004016101699190610557565b80516001600160a01b038116811461041957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561044f578181015183820152602001610437565b838111156100f95750506000910152565b6000806040838503121561047357600080fd5b61047c83610402565b60208401519092506001600160401b038082111561049957600080fd5b818501915085601f8301126104ad57600080fd5b8151818111156104bf576104bf61041e565b604051601f8201601f19908116603f011681019083821181831017156104e7576104e761041e565b8160405282815288602084870101111561050057600080fd5b610511836020830160208801610434565b80955050505050509250929050565b60006020828403121561053257600080fd5b6102c882610402565b6000825161054d818460208701610434565b9190910192915050565b6020815260008251806020840152610576816040850160208701610434565b601f01601f19169190910160400192915050565b61034e806105996000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610067565b610100565b565b606061004e83836040518060600160405280602781526020016102f260279139610124565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fb9190610249565b905090565b3660008037600080366000845af43d6000803e80801561011f573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161014191906102a2565b600060405180830381855af49150503d806000811461017c576040519150601f19603f3d011682016040523d82523d6000602084013e610181565b606091505b50915091506101928683838761019c565b9695505050505050565b6060831561020d578251610206576001600160a01b0385163b6102065760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610217565b610217838361021f565b949350505050565b81511561022f5781518083602001fd5b8060405162461bcd60e51b81526004016101fd91906102be565b60006020828403121561025b57600080fd5b81516001600160a01b038116811461004e57600080fd5b60005b8381101561028d578181015183820152602001610275565b8381111561029c576000848401525b50505050565b600082516102b4818460208701610272565b9190910192915050565b60208152600082518060208401526102dd816040850160208701610272565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d51e81d3bc5ed20a26aeb05dce7e825c503b2061aa78628027300c8d65b9d89a64736f6c634300080c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564";

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Create time machine and beacon chain. Set block time to beacon chain genesis time
        // beaconChainMock will also etch 4788 precompile
        ethPOSDepositMock = new ETHPOSDepositMock();
        cheats.warp(GENESIS_TIME_LOCAL);
        timeMachine = new TimeMachine();
        beaconChain = new BeaconChainMock(EigenPodManager(address(eigenPodManagerMock)), GENESIS_TIME_LOCAL);

        // Deploy EigenPod
        podImplementation = new EigenPod(ethPOSDepositMock, IEigenPodManager(address(eigenPodManagerMock)), GENESIS_TIME_LOCAL, "v9.9.9");

        // Deploy Beacon
        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));

        // Deploy Proxy same way as EigenPodManager does
        eigenPod = EigenPod(
            payable(
                Create2.deploy(
                    0,
                    bytes32(uint(uint160(address(this)))),
                    // set the beacon address to the eigenPodBeacon
                    abi.encodePacked(beaconProxyBytecode, abi.encode(eigenPodBeacon, ""))
                )
            )
        );

        // Etch 4788 precompile
        cheats.etch(address(EIP_4788_ORACLE), type(EIP_4788_Oracle_Mock).runtimeCode);

        // Store the eigenPodBeacon address in the eigenPod beacon proxy
        bytes32 beaconSlot = 0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50;
        cheats.store(address(eigenPod), beaconSlot, bytes32(uint(uint160(address(eigenPodBeacon)))));

        // Initialize pod
        eigenPod.initialize(address(this));

        // Set a proof submitter
        eigenPod.setProofSubmitter(defaultProofSubmitter);
    }

    /**
     *
     *                     EIGENPOD Helpers
     *
     */
    modifier timewarp() {
        uint curState = timeMachine.travelToLast();
        _;
        timeMachine.travel(curState);
    }

    function _seedPodWithETH(uint ethAmount) internal {
        cheats.deal(address(this), ethAmount);
        bool result;
        bytes memory data;
        (result, data) = address(eigenPod).call{value: ethAmount}("");
    }

    function _newEigenPodStaker(uint rand) internal returns (EigenPodUser, uint) {
        string memory stakerName;

        EigenPodUser staker;

        stakerName = string.concat("Staker", cheats.toString(numStakers));
        staker = new EigenPodUser(stakerName);

        uint amount = bound(rand, 1 ether, 640 ether);
        cheats.deal(address(staker), amount);

        numStakers++;
        return (staker, amount);
    }

    /// @dev Opposite of Endian.fromLittleEndianUint64
    function _toLittleEndianUint64(uint64 num) internal pure returns (bytes32) {
        uint lenum;

        // Rearrange the bytes from big-endian to little-endian format
        lenum |= uint((num & 0xFF) << 56);
        lenum |= uint((num & 0xFF00) << 40);
        lenum |= uint((num & 0xFF0000) << 24);
        lenum |= uint((num & 0xFF000000) << 8);
        lenum |= uint((num & 0xFF00000000) >> 8);
        lenum |= uint((num & 0xFF0000000000) >> 24);
        lenum |= uint((num & 0xFF000000000000) >> 40);
        lenum |= uint((num & 0xFF00000000000000) >> 56);

        // Shift the little-endian bytes to the end of the bytes32 value
        return bytes32(lenum << 192);
    }

    /**
     *
     *                     verifyWithdrawalCredentials Assertions
     *
     */
    function assert_Snap_Added_ActiveValidatorCount(EigenPodUser staker, uint addedValidators, string memory err) internal {
        uint curActiveValidatorCount = _getActiveValidatorCount(staker);
        uint prevActiveValidatorCount = _getPrevActiveValidatorCount(staker);

        assertEq(prevActiveValidatorCount + addedValidators, curActiveValidatorCount, err);
    }

    function assert_Snap_Removed_ActiveValidatorCount(EigenPodUser staker, uint removedValidators, string memory err) internal {
        uint curActiveValidatorCount = _getActiveValidatorCount(staker);
        uint prevActiveValidatorCount = _getPrevActiveValidatorCount(staker);

        assertEq(curActiveValidatorCount + removedValidators, prevActiveValidatorCount, err);
    }

    function assert_Snap_Unchanged_ActiveValidatorCount(EigenPodUser staker, string memory err) internal {
        uint curActiveValidatorCount = _getActiveValidatorCount(staker);
        uint prevActiveValidatorCount = _getPrevActiveValidatorCount(staker);

        assertEq(curActiveValidatorCount, prevActiveValidatorCount, err);
    }

    function _getActiveValidatorCount(EigenPodUser staker) internal view returns (uint) {
        EigenPod pod = staker.pod();
        return pod.activeValidatorCount();
    }

    function _getPrevActiveValidatorCount(EigenPodUser staker) internal timewarp returns (uint) {
        return _getActiveValidatorCount(staker);
    }

    function assert_Snap_Added_ActiveValidators(EigenPodUser staker, uint40[] memory addedValidators, string memory err) internal {
        bytes32[] memory pubkeyHashes = beaconChain.getPubkeyHashes(addedValidators);

        IEigenPodTypes.VALIDATOR_STATUS[] memory curStatuses = _getValidatorStatuses(staker, pubkeyHashes);
        IEigenPodTypes.VALIDATOR_STATUS[] memory prevStatuses = _getPrevValidatorStatuses(staker, pubkeyHashes);

        for (uint i = 0; i < curStatuses.length; i++) {
            assertTrue(prevStatuses[i] == IEigenPodTypes.VALIDATOR_STATUS.INACTIVE, err);
            assertTrue(curStatuses[i] == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE, err);
        }
    }

    function assert_Snap_Removed_ActiveValidators(EigenPodUser staker, uint40[] memory removedValidators, string memory err) internal {
        bytes32[] memory pubkeyHashes = beaconChain.getPubkeyHashes(removedValidators);

        IEigenPodTypes.VALIDATOR_STATUS[] memory curStatuses = _getValidatorStatuses(staker, pubkeyHashes);
        IEigenPodTypes.VALIDATOR_STATUS[] memory prevStatuses = _getPrevValidatorStatuses(staker, pubkeyHashes);

        for (uint i = 0; i < curStatuses.length; i++) {
            assertTrue(prevStatuses[i] == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE, err);
            assertTrue(curStatuses[i] == IEigenPodTypes.VALIDATOR_STATUS.WITHDRAWN, err);
        }
    }

    function _getValidatorStatuses(EigenPodUser staker, bytes32[] memory pubkeyHashes)
        internal
        view
        returns (IEigenPodTypes.VALIDATOR_STATUS[] memory)
    {
        EigenPod pod = staker.pod();
        IEigenPodTypes.VALIDATOR_STATUS[] memory statuses = new IEigenPodTypes.VALIDATOR_STATUS[](pubkeyHashes.length);

        for (uint i = 0; i < statuses.length; i++) {
            statuses[i] = pod.validatorStatus(pubkeyHashes[i]);
        }

        return statuses;
    }

    function _getPrevValidatorStatuses(EigenPodUser staker, bytes32[] memory pubkeyHashes)
        internal
        timewarp
        returns (IEigenPodTypes.VALIDATOR_STATUS[] memory)
    {
        return _getValidatorStatuses(staker, pubkeyHashes);
    }

    /**
     *
     *                     startCheckpoint Assertions
     *
     */
    function check_StartCheckpoint_State(EigenPodUser staker) internal {
        assert_ProofsRemainingEqualsActive(staker, "checkpoint proofs remaining should equal active validator count");
        assert_Snap_Created_Checkpoint(staker, "staker should have created a new checkpoint");
    }

    function assert_ProofsRemainingEqualsActive(EigenPodUser staker, string memory err) internal view {
        EigenPod pod = staker.pod();
        assertEq(pod.currentCheckpoint().proofsRemaining, pod.activeValidatorCount(), err);
    }

    function assert_Snap_Created_Checkpoint(EigenPodUser staker, string memory err) internal {
        uint64 curCheckpointTimestamp = _getCheckpointTimestamp(staker);
        uint64 prevCheckpointTimestamp = _getPrevCheckpointTimestamp(staker);

        assertEq(prevCheckpointTimestamp, 0, err);
        assertTrue(curCheckpointTimestamp != 0, err);
    }

    function _getCheckpointTimestamp(EigenPodUser staker) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.currentCheckpointTimestamp();
    }

    function _getPrevCheckpointTimestamp(EigenPodUser staker) internal timewarp returns (uint64) {
        return _getCheckpointTimestamp(staker);
    }

    /**
     *
     *                     verifyCheckpointProofs
     *
     */

    /// @notice assumes positive rewards and that the checkpoint will be finalized
    function _expectEventsVerifyCheckpointProofs(
        EigenPodUser staker,
        uint40[] memory validators,
        BeaconChainProofs.BalanceProof[] memory proofs
    ) internal {
        EigenPod pod = staker.pod();
        int totalBalanceDeltaGWei = 0;
        uint64 checkpointTimestamp = pod.currentCheckpointTimestamp();
        for (uint i = 0; i < validators.length; i++) {
            IEigenPodTypes.ValidatorInfo memory info = pod.validatorPubkeyHashToInfo(proofs[i].pubkeyHash);
            uint64 prevBalanceGwei = info.restakedBalanceGwei;
            uint64 newBalanceGwei = BeaconChainProofs.getBalanceAtIndex(proofs[i].balanceRoot, validators[i]);
            int128 balanceDeltaGwei = _calcBalanceDelta({newAmountGwei: newBalanceGwei, previousAmountGwei: prevBalanceGwei});
            if (newBalanceGwei != prevBalanceGwei) {
                cheats.expectEmit(true, true, true, true, address(pod));
                emit ValidatorBalanceUpdated(validators[i], checkpointTimestamp, newBalanceGwei);
            }

            if (newBalanceGwei == 0) {
                cheats.expectEmit(true, true, true, true, address(pod));
                emit ValidatorWithdrawn(checkpointTimestamp, validators[i]);
            }

            cheats.expectEmit(true, true, true, true, address(pod));
            emit ValidatorCheckpointed(checkpointTimestamp, validators[i]);

            totalBalanceDeltaGWei += balanceDeltaGwei;
        }
        int totalShareDeltaWei = (int128(uint128(pod.currentCheckpoint().podBalanceGwei)) + totalBalanceDeltaGWei) * int(1 gwei);
        cheats.expectEmit(true, true, true, true, address(pod));
        emit CheckpointFinalized(checkpointTimestamp, totalShareDeltaWei);
    }

    /// @dev Calculates the delta between two Gwei amounts and returns as an int256
    function _calcBalanceDelta(uint64 newAmountGwei, uint64 previousAmountGwei) internal pure returns (int128) {
        return int128(uint128(newAmountGwei)) - int128(uint128(previousAmountGwei));
    }
}

contract EigenPodUnitTests_Initialization is EigenPodUnitTests {
    function test_constructor() public {
        EigenPod pod = new EigenPod(ethPOSDepositMock, IEigenPodManager(address(eigenPodManagerMock)), GENESIS_TIME_LOCAL, "v9.9.9");

        assertTrue(pod.ethPOS() == ethPOSDepositMock, "should have set ethPOS correctly");
        assertTrue(address(pod.eigenPodManager()) == address(eigenPodManagerMock), "should have set eigenpodmanager correctly");
        assertTrue(pod.GENESIS_TIME() == GENESIS_TIME_LOCAL, "should have set genesis time correctly");
    }

    function test_initialization() public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();

        // Check podOwner and restaked
        assertEq(pod.podOwner(), address(staker), "Pod owner incorrectly set");
        // Check immutable storage
        assertEq(address(pod.ethPOS()), address(ethPOSDepositMock), "EthPOS incorrectly set");
        assertEq(address(pod.eigenPodManager()), address(eigenPodManagerMock), "EigenPodManager incorrectly set");
        assertEq(pod.GENESIS_TIME(), GENESIS_TIME_LOCAL, "LOCAL genesis time incorrectly set");
    }

    function test_initialize_revert_alreadyInitialized() public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();

        cheats.expectRevert("Initializable: contract is already initialized");
        pod.initialize(address(this));
    }

    function test_initialize_revert_emptyPodOwner() public {
        EigenPod pod = new EigenPod(ethPOSDepositMock, IEigenPodManager(address(eigenPodManagerMock)), GENESIS_TIME_LOCAL, "v9.9.9");
        // un-initialize pod
        cheats.store(address(pod), 0, 0);

        cheats.expectRevert(IEigenPodErrors.InputAddressZero.selector);
        pod.initialize(address(0));
    }

    function test_setProofSubmitter_revert_notPodOwner(address invalidCaller) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();

        cheats.assume(invalidCaller != address(staker));
        cheats.prank(invalidCaller);
        cheats.expectRevert(IEigenPodErrors.OnlyEigenPodOwner.selector);
        pod.setProofSubmitter(invalidCaller);
    }

    function test_setProofSubmitter(address newProofSubmitter) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();
        address prevProofSubmitter = pod.proofSubmitter();

        cheats.prank(address(staker));
        cheats.expectEmit(true, true, true, true, address(pod));
        emit ProofSubmitterUpdated(prevProofSubmitter, newProofSubmitter);
        pod.setProofSubmitter(newProofSubmitter);

        assertEq(pod.proofSubmitter(), newProofSubmitter, "did not update proof submitter");
    }
}

contract EigenPodUnitTests_EPMFunctions is EigenPodUnitTests {
    /**
     *
     *                         stake() tests
     *
     */

    // Beacon chain staking constnats
    bytes public constant pubkey = hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";
    bytes public signature;
    bytes32 public depositDataRoot;

    function testFuzz_stake_revert_notEigenPodManager(address invalidCaller) public {
        cheats.assume(invalidCaller != address(eigenPodManagerMock));
        cheats.deal(invalidCaller, 32 ether);

        cheats.prank(invalidCaller);
        cheats.expectRevert(IEigenPodErrors.OnlyEigenPodManager.selector);
        eigenPod.stake{value: 32 ether}(pubkey, signature, depositDataRoot);
    }

    function testFuzz_stake_revert_invalidValue(uint value) public {
        cheats.assume(value != 32 ether);
        cheats.deal(address(eigenPodManagerMock), value);

        cheats.prank(address(eigenPodManagerMock));
        cheats.expectRevert(IEigenPodErrors.MsgValueNot32ETH.selector);
        eigenPod.stake{value: value}(pubkey, signature, depositDataRoot);
    }

    function test_stake() public {
        cheats.deal(address(eigenPodManagerMock), 32 ether);

        // Expect emit
        vm.expectEmit(true, true, true, true);
        emit EigenPodStaked(pubkey);

        // Stake
        cheats.prank(address(eigenPodManagerMock));
        eigenPod.stake{value: 32 ether}(pubkey, signature, depositDataRoot);

        // Check eth transferred
        assertEq(address(ethPOSDepositMock).balance, 32 ether, "Incorrect amount transferred");
    }

    /**
     *
     *                         withdrawRestakedBeaconChainETH() tests
     *
     */
    function testFuzz_withdrawRestakedBeaconChainETH_revert_notEigenPodManager(address invalidCaller, address recipient, uint randAmount)
        public
        filterFuzzedAddressInputs(invalidCaller)
    {
        // Setup EigenPod Staker
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();

        // ensure invalid caller causing revert
        cheats.assume(invalidCaller != address(eigenPodManagerMock));
        cheats.prank(invalidCaller);
        cheats.expectRevert(IEigenPodErrors.OnlyEigenPodManager.selector);
        pod.withdrawRestakedBeaconChainETH(recipient, randAmount);
    }

    function testFuzz_withdrawRestakedBeaconChainETH_revert_withdrawAmountTooLarge(uint rand, address recipient, uint randAmountWei)
        public
    {
        // Setup EigenPod Staker
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        beaconChain.advanceEpoch();
        staker.startCheckpoint();
        staker.completeCheckpoint();

        // ensure amount is too large
        uint64 withdrawableRestakedExecutionLayerGwei = pod.withdrawableRestakedExecutionLayerGwei();
        randAmountWei = randAmountWei - (randAmountWei % 1 gwei);
        cheats.assume((randAmountWei / 1 gwei) > withdrawableRestakedExecutionLayerGwei);
        cheats.expectRevert(IEigenPodErrors.InsufficientWithdrawableBalance.selector);
        cheats.prank(address(eigenPodManagerMock));
        pod.withdrawRestakedBeaconChainETH(recipient, randAmountWei);
    }

    function testFuzz_withdrawRestakedBeaconChainETH(uint rand, uint randAmountWei) public {
        // Setup EigenPod Staker
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        beaconChain.advanceEpoch();
        staker.startCheckpoint();
        staker.completeCheckpoint();

        // ensure valid fuzzed wei amounts
        uint64 withdrawableRestakedExecutionLayerGwei = pod.withdrawableRestakedExecutionLayerGwei();
        randAmountWei = randAmountWei - (randAmountWei % 1 gwei);
        cheats.assume((randAmountWei / 1 gwei) <= withdrawableRestakedExecutionLayerGwei);

        address recipient = cheats.addr(uint(123_456_789));

        cheats.prank(address(eigenPodManagerMock));
        cheats.expectEmit(true, true, true, true, address(pod));
        emit RestakedBeaconChainETHWithdrawn(recipient, randAmountWei);
        pod.withdrawRestakedBeaconChainETH(recipient, randAmountWei);

        assertEq(address(recipient).balance, randAmountWei, "recipient should have received withdrawn balance");
        assertEq(
            address(pod).balance,
            uint(withdrawableRestakedExecutionLayerGwei * 1 gwei) - randAmountWei,
            "pod balance should have decreased by withdrawn eth"
        );
        assertEq(
            pod.withdrawableRestakedExecutionLayerGwei(),
            withdrawableRestakedExecutionLayerGwei - uint64(randAmountWei / 1 gwei),
            "withdrawableRestakedExecutionLayerGwei should have decreased by amount withdrawn"
        );
    }

    function testFuzz_withdrawRestakedBeaconChainETH_AmountGweiNotDivisibleByGwei(uint rand, uint randAmountWei) public {
        // Setup EigenPod Staker
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        beaconChain.advanceEpoch();
        staker.startCheckpoint();
        staker.completeCheckpoint();

        // ensure valid fuzzed wei amounts
        uint64 withdrawableRestakedExecutionLayerGwei = pod.withdrawableRestakedExecutionLayerGwei();
        uint randAmountWeiAdjusted = randAmountWei - (randAmountWei % 1 gwei);
        cheats.assume((randAmountWei / 1 gwei) <= withdrawableRestakedExecutionLayerGwei);

        address recipient = cheats.addr(uint(123_456_789));

        cheats.prank(address(eigenPodManagerMock));
        cheats.expectEmit(true, true, true, true, address(pod));
        emit RestakedBeaconChainETHWithdrawn(recipient, randAmountWeiAdjusted);
        pod.withdrawRestakedBeaconChainETH(recipient, randAmountWei);

        assertEq(address(recipient).balance, randAmountWeiAdjusted, "recipient should have received withdrawn balance");
        assertEq(
            address(pod).balance,
            uint(withdrawableRestakedExecutionLayerGwei * 1 gwei) - randAmountWeiAdjusted,
            "pod balance should have decreased by withdrawn eth"
        );
        assertEq(
            pod.withdrawableRestakedExecutionLayerGwei(),
            withdrawableRestakedExecutionLayerGwei - uint64(randAmountWeiAdjusted / 1 gwei),
            "withdrawableRestakedExecutionLayerGwei should have decreased by amount withdrawn"
        );
    }
}

contract EigenPodUnitTests_recoverTokens is EigenPodUnitTests {
    /**
     *
     *                         recoverTokens() tests
     *
     */
    function testFuzz_recoverTokens_revert_notPodOwner(address invalidCaller) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();
        address podOwner = pod.podOwner();

        cheats.assume(invalidCaller != podOwner);
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(address(0x123));
        uint[] memory amounts = new uint[](1);
        amounts[0] = 1;

        cheats.prank(invalidCaller);
        cheats.expectRevert(IEigenPodErrors.OnlyEigenPodOwner.selector);
        pod.recoverTokens(tokens, amounts, podOwner);
    }

    function test_recoverTokens_revert_whenPaused() public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();
        address podOwner = pod.podOwner();

        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(address(0x123));
        uint[] memory amounts = new uint[](1);
        amounts[0] = 1;

        // pause recoverTokens
        cheats.prank(pauser);
        eigenPodManagerMock.pause(1 << PAUSED_NON_PROOF_WITHDRAWALS);

        cheats.prank(podOwner);
        cheats.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        pod.recoverTokens(tokens, amounts, podOwner);
    }

    function test_recoverTokens_revert_invalidLengths() public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();
        address podOwner = pod.podOwner();

        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(address(0x123));
        uint[] memory amounts = new uint[](2);
        amounts[0] = 1;
        amounts[1] = 1;

        cheats.startPrank(podOwner);
        cheats.expectRevert(IEigenPodErrors.InputArrayLengthMismatch.selector);
        pod.recoverTokens(tokens, amounts, podOwner);
        cheats.stopPrank();
    }

    function test_recoverTokens() public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        EigenPod pod = staker.pod();
        address podOwner = pod.podOwner();

        // Deploy dummy token
        IERC20 dummyToken = new ERC20Mock();
        dummyToken.transfer(address(pod), 1e18);

        // Recover tokens
        address recipient = address(0x123);
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = dummyToken;
        uint[] memory amounts = new uint[](1);
        amounts[0] = 1e18;

        cheats.prank(podOwner);
        pod.recoverTokens(tokens, amounts, recipient);

        // Checks
        assertEq(dummyToken.balanceOf(recipient), 1e18, "Incorrect amount recovered");
    }
}

contract EigenPodUnitTests_verifyWithdrawalCredentials is EigenPodUnitTests, ProofParsing {
    /**
     *
     *                         verifyWithdrawalCredentials() tests
     *
     */

    /// @notice revert when verify wc is not called by pod owner
    function testFuzz_revert_callerIsNotPodOwnerOrProofSubmitter(address invalidCaller) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});

        (uint40[] memory validators,) = staker.startValidators();
        EigenPod pod = staker.pod();
        address podOwner = pod.podOwner();
        address proofSubmitter = pod.proofSubmitter();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        cheats.assume(invalidCaller != podOwner && invalidCaller != proofSubmitter);
        cheats.prank(invalidCaller);
        cheats.expectRevert(IEigenPodErrors.OnlyEigenPodOwnerOrProofSubmitter.selector);

        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
    }

    /// @notice test verify wc reverts when paused
    function test_revert_verifyWithdrawalCredentialsPaused() public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: 0});
        (uint40[] memory validators,) = staker.startValidators();

        cheats.prank(pauser);
        eigenPodManagerMock.pause(2 ** PAUSED_EIGENPODS_VERIFY_CREDENTIALS);

        cheats.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        staker.verifyWithdrawalCredentials(validators);
    }

    /// @notice beaconTimestamp must be after the current checkpoint
    function testFuzz_revert_beaconTimestampInvalid(uint rand) public {
        cheats.warp(10 days);
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        // Ensure we have more than one validator (_newEigenPodStaker allocates a nonzero amt of eth)
        cheats.deal(address(staker), address(staker).balance + 32 ether);
        (uint40[] memory validators,) = staker.startValidators();

        uint40[] memory firstValidator = new uint40[](1);
        firstValidator[0] = validators[0];
        staker.verifyWithdrawalCredentials(firstValidator);

        // Start a checkpoint so `currentCheckpointTimestamp` is nonzero
        staker.startCheckpoint();

        // Try to verify withdrawal credentials at the current block
        cheats.expectRevert(IEigenPodErrors.BeaconTimestampTooFarInPast.selector);
        staker.verifyWithdrawalCredentials(validators);
    }

    /// @notice Check for revert on input array mismatch lengths
    function testFuzz_revert_inputArrayLengthsMismatch(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();
        EigenPod pod = staker.pod();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        uint40[] memory invalidValidatorIndices = new uint40[](validators.length + 1);
        bytes[] memory invalidValidatorFieldsProofs = new bytes[](proofs.validatorFieldsProofs.length + 1);
        bytes32[][] memory invalidValidatorFields = new bytes32[][](proofs.validatorFields.length + 1);

        cheats.startPrank(address(staker));
        cheats.expectRevert(IEigenPodErrors.InputArrayLengthMismatch.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: invalidValidatorIndices,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });

        cheats.expectRevert(IEigenPodErrors.InputArrayLengthMismatch.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: invalidValidatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });

        cheats.expectRevert(IEigenPodErrors.InputArrayLengthMismatch.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: invalidValidatorFields
        });

        cheats.stopPrank();
    }

    /// @notice Check beaconStateRootProof reverts on invalid length or invalid proof
    function testFuzz_revert_beaconStateRootProofInvalid(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();
        EigenPod pod = staker.pod();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        bytes memory proofWithInvalidLength = new bytes(proofs.stateRootProof.proof.length + 1);
        BeaconChainProofs.StateRootProof memory invalidStateRootProof =
            BeaconChainProofs.StateRootProof({beaconStateRoot: proofs.stateRootProof.beaconStateRoot, proof: proofWithInvalidLength});

        cheats.startPrank(address(staker));
        cheats.expectRevert(BeaconChainProofs.InvalidProofLength.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: invalidStateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });

        // Change the proof to have an invalid value
        bytes1 randValue = bytes1(keccak256(abi.encodePacked(proofs.stateRootProof.proof[0])));
        uint proofLength = proofs.stateRootProof.proof.length;
        uint randIndex = bound(rand, 0, proofLength - 1);
        proofs.stateRootProof.proof[randIndex] = randValue;
        cheats.expectRevert(BeaconChainProofs.InvalidProof.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        cheats.stopPrank();
    }

    /// @notice attempt to verify validator credentials in both ACTIVE and WITHDRAWN states
    /// check reverts
    function testFuzz_revert_validatorsWithdrawn(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);

        // now that validators are ACTIVE, ensure we can't verify them again
        cheats.expectRevert(IEigenPodErrors.CredentialsAlreadyVerified.selector);
        staker.verifyWithdrawalCredentials(validators);

        staker.exitValidators(validators);
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();
        staker.completeCheckpoint();
        beaconChain.advanceEpoch_NoRewards();

        // now that validators are WITHDRAWN, ensure we can't verify them again
        cheats.expectRevert(IEigenPodErrors.CredentialsAlreadyVerified.selector);
        staker.verifyWithdrawalCredentials(validators);
    }

    /// @notice attempt to verify validator credentials after they have exited
    /// check reverts
    function testFuzz_revert_validatorsExited(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();

        // Exit validators from beacon chain and withdraw to pod
        staker.exitValidators(validators);
        beaconChain.advanceEpoch();

        // now that validators are exited, ensure we can't verify them
        cheats.expectRevert(IEigenPodErrors.ValidatorIsExitingBeaconChain.selector);
        staker.verifyWithdrawalCredentials(validators);
    }

    /// @notice modify withdrawal credentials to cause a revert
    function testFuzz_revert_invalidWithdrawalAddress(uint rand, bytes32 invalidWithdrawalCredentials) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();
        EigenPod pod = staker.pod();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        // Set invalid withdrawal credentials in validatorFields
        uint VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX = 1;
        proofs.validatorFields[0][VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX] = invalidWithdrawalCredentials;

        cheats.startPrank(address(staker));
        cheats.expectRevert(IEigenPodErrors.WithdrawalCredentialsNotForEigenPod.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        cheats.stopPrank();
    }

    /// @notice modify validator field length to cause a revert
    function testFuzz_revert_invalidValidatorFields(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();
        EigenPod pod = staker.pod();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        // change validator field length to invalid value
        bytes32[] memory invalidValidatorFields = new bytes32[](BeaconChainProofs.VALIDATOR_FIELDS_LENGTH + 1);
        for (uint i = 0; i < BeaconChainProofs.VALIDATOR_FIELDS_LENGTH; i++) {
            invalidValidatorFields[i] = proofs.validatorFields[0][i];
        }
        proofs.validatorFields[0] = invalidValidatorFields;

        cheats.startPrank(address(staker));
        cheats.expectRevert(BeaconChainProofs.InvalidValidatorFieldsLength.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        cheats.stopPrank();
    }

    /// @notice modify validator activation epoch to cause a revert
    function testFuzz_revert_activationEpochNotSet(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();
        EigenPod pod = staker.pod();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        proofs.validatorFields[0][BeaconChainProofs.VALIDATOR_ACTIVATION_EPOCH_INDEX] =
            _toLittleEndianUint64(BeaconChainProofs.FAR_FUTURE_EPOCH);

        cheats.startPrank(address(staker));
        cheats.expectRevert(IEigenPodErrors.ValidatorInactiveOnBeaconChain.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        cheats.stopPrank();
    }

    /// @notice modify validator proof length to cause a revert
    function testFuzz_revert_invalidValidatorProofLength(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();
        EigenPod pod = staker.pod();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        // add an element to the proof
        proofs.validatorFieldsProofs[0] = new bytes(proofs.validatorFieldsProofs[0].length + 32);

        cheats.startPrank(address(staker));
        cheats.expectRevert(BeaconChainProofs.InvalidProofLength.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        cheats.stopPrank();
    }

    /// @notice modify validator pubkey to cause a revert
    function testFuzz_revert_invalidValidatorProof(uint rand, bytes32 randPubkey) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        (uint40[] memory validators,) = staker.startValidators();
        EigenPod pod = staker.pod();
        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        // change validator pubkey to an invalid value causing a revert
        uint VALIDATOR_PUBKEY_INDEX = 0;
        proofs.validatorFields[0][VALIDATOR_PUBKEY_INDEX] = randPubkey;

        cheats.startPrank(address(staker));
        cheats.expectRevert(BeaconChainProofs.InvalidProof.selector);
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        cheats.stopPrank();
    }

    /// @notice fuzz test a eigenPod with multiple validators. Using timemachine to assert values over time
    function testFuzz_verifyWithdrawalCredentials(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        // Complete a quick empty checkpoint so we have a nonzero value for `lastCheckpointedAt`
        staker.startCheckpoint();
        beaconChain.advanceEpoch_NoRewards();

        CredentialProofs memory proofs = beaconChain.getCredentialProofs(validators);

        for (uint i; i < validators.length; i++) {
            cheats.expectEmit(true, true, true, true, address(pod));
            emit ValidatorRestaked(validators[i]);
            cheats.expectEmit(true, true, true, true, address(pod));
            emit ValidatorBalanceUpdated(validators[i], pod.lastCheckpointTimestamp(), beaconChain.effectiveBalance(validators[i]));
        }
        // staker.verifyWithdrawalCredentials(validators);
        // Prank either the staker or proof submitter
        cheats.prank(rand % 2 == 1 ? address(staker) : pod.proofSubmitter());
        pod.verifyWithdrawalCredentials({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            validatorIndices: validators,
            validatorFieldsProofs: proofs.validatorFieldsProofs,
            validatorFields: proofs.validatorFields
        });
        assert_Snap_Added_ActiveValidatorCount(staker, validators.length, "staker should have increased active validator count");
        assert_Snap_Added_ActiveValidators(staker, validators, "validators should each be active");
        // Check ValidatorInfo values for each validator
        for (uint i = 0; i < validators.length; i++) {
            bytes32 pubkeyHash = beaconChain.pubkeyHash(validators[i]);
            bytes memory pubkey = beaconChain.pubkey(validators[i]);

            IEigenPodTypes.ValidatorInfo memory info = pod.validatorPubkeyHashToInfo(pubkeyHash);
            IEigenPodTypes.ValidatorInfo memory pkInfo = pod.validatorPubkeyToInfo(pubkey);
            assertTrue(pod.validatorStatus(pubkey) == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE, "validator status should be active");
            assertEq(keccak256(abi.encode(info)), keccak256(abi.encode(pkInfo)), "validator info should be identical");
            assertEq(info.validatorIndex, validators[i], "should have assigned correct validator index");
            assertEq(info.restakedBalanceGwei, beaconChain.effectiveBalance(validators[i]), "should have restaked full effective balance");
            assertEq(info.lastCheckpointedAt, pod.lastCheckpointTimestamp(), "should have recorded correct update time");
        }
    }
}

contract EigenPodUnitTests_startCheckpoint is EigenPodUnitTests {
    /**
     *
     *                         startCheckpoint() tests
     *
     */

    /// @notice revert when startCheckpoint is not called by pod owner
    function testFuzz_revert_callerIsNotPodOwnerOrProofSubmitter(uint rand, address invalidCaller) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});

        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        EigenPod pod = staker.pod();
        address podOwner = pod.podOwner();
        address proofSubmitter = pod.proofSubmitter();

        cheats.assume(invalidCaller != podOwner && invalidCaller != proofSubmitter);
        cheats.prank(invalidCaller);
        cheats.expectRevert(IEigenPodErrors.OnlyEigenPodOwnerOrProofSubmitter.selector);
        pod.startCheckpoint({revertIfNoBalance: false});
    }

    /// @notice test startCheckpoint reverts when paused
    function testFuzz_revert_startCheckpointPaused(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});

        staker.startValidators();

        cheats.prank(pauser);
        eigenPodManagerMock.pause(2 ** PAUSED_START_CHECKPOINT);

        cheats.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        staker.startCheckpoint();
    }

    /// @notice startCheckpoint should revert if another checkpoint already in progress
    function testFuzz_revert_checkpointAlreadyStarted(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});

        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        staker.startCheckpoint();
        cheats.expectRevert(IEigenPodErrors.CheckpointAlreadyActive.selector);
        staker.startCheckpoint();
    }

    /// @notice startCheckpoint should revert if a checkpoint has already been completed this block
    function testFuzz_revert_checkpointTwicePerBlock(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});

        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        staker.startCheckpoint();
        staker.completeCheckpoint();

        cheats.expectRevert(IEigenPodErrors.CannotCheckpointTwiceInSingleBlock.selector);
        staker.startCheckpoint();
    }

    /// @notice if no rewards and revertIfNoBalance is set, startCheckpoint should revert
    function testFuzz_revert_revertIfNoBalanceIsSet(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});

        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        EigenPod pod = staker.pod();
        beaconChain.advanceEpoch_NoRewards();

        cheats.prank(pod.podOwner());
        cheats.expectRevert(IEigenPodErrors.NoBalanceToCheckpoint.selector);
        pod.startCheckpoint({revertIfNoBalance: true});
    }

    /// @notice fuzz test an eigenpod with multiple validators and starting a checkpoint
    function testFuzz_startCheckpoint(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();

        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);

        cheats.expectEmit(true, true, true, true, address(staker.pod()));
        emit CheckpointCreated(uint64(block.timestamp), EIP_4788_ORACLE.timestampToBlockRoot(block.timestamp), validators.length);
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);
        assertEq(
            pod.currentCheckpoint().proofsRemaining, uint24(validators.length), "should have one proof remaining pre verified validator"
        );
    }

    /// @notice fuzz test an eigenpod with multiple validators and starting a checkpoint
    function testFuzz_startCheckpoint_AsProofSubmitter(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();

        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);

        cheats.expectEmit(true, true, true, true, address(staker.pod()));
        emit CheckpointCreated(uint64(block.timestamp), EIP_4788_ORACLE.timestampToBlockRoot(block.timestamp), validators.length);
        cheats.prank(pod.proofSubmitter());
        pod.startCheckpoint(false);
        check_StartCheckpoint_State(staker);
        assertEq(
            pod.currentCheckpoint().proofsRemaining, uint24(validators.length), "should have one proof remaining pre verified validator"
        );
    }
}

contract EigenPodUnitTests_verifyCheckpointProofs is EigenPodUnitTests {
    /**
     *
     *                         verifyCheckpointProofs() tests
     *
     */

    /// @notice test verifyCheckpointProofs reverts when paused
    function testFuzz_revert_verifyCheckpointProofsPaused(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());

        cheats.prank(pauser);
        eigenPodManagerMock.pause(2 ** PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS);
        cheats.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
    }

    /// @notice verifyCheckpointProofs should revert if checkpoint not in progress
    function testFuzz_revert_checkpointNotStarted(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());
        cheats.expectRevert(IEigenPodErrors.NoActiveCheckpoint.selector);
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
    }

    /// @notice invalid proof length should revert
    function testFuzz_revert_verifyBalanceContainerInvalidLengths(uint rand) public {
        // Setup verifyCheckpointProofs
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        beaconChain.advanceEpoch();
        staker.startCheckpoint();
        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());

        // change the length of balanceContainerProof to cause a revert
        proofs.balanceContainerProof.proof = new bytes(proofs.balanceContainerProof.proof.length + 1);

        cheats.expectRevert(BeaconChainProofs.InvalidProofLength.selector);
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
    }

    /// @notice change one of the bytes in the balanceContainer proof to cause a revert
    function testFuzz_revert_verifyBalanceContainerInvalidProof(uint rand) public {
        // Setup verifyCheckpointProofs
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        beaconChain.advanceEpoch();
        staker.startCheckpoint();

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());
        // randomly change one of the bytes in the proof to make the proof invalid
        bytes1 randValue = bytes1(keccak256(abi.encodePacked(proofs.balanceContainerProof.proof[0])));
        proofs.balanceContainerProof.proof[0] = randValue;

        cheats.expectRevert(BeaconChainProofs.InvalidProof.selector);
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
    }

    /// @notice invalid balance proof length should revert
    function testFuzz_revert_verifyValidatorBalanceInvalidLength(uint rand) public {
        // Setup verifyCheckpointProofs
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        beaconChain.advanceEpoch();
        staker.startCheckpoint();
        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());

        // change the length of balance proof to cause a revert
        proofs.balanceProofs[0].proof = new bytes(proofs.balanceProofs[0].proof.length + 1);

        cheats.expectRevert(BeaconChainProofs.InvalidProofLength.selector);
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
    }

    /// @notice change one of the bytes in one of the balance proofs to cause a revert
    function testFuzz_revert_verifyValidatorBalanceInvalidProof(uint rand) public {
        // Setup verifyCheckpointProofs
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        beaconChain.advanceEpoch();
        staker.startCheckpoint();

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());
        // randomly change one of the bytes in the first proof to make the proof invalid
        bytes1 randValue = bytes1(keccak256(abi.encodePacked(proofs.balanceProofs[0].proof[0])));
        proofs.balanceProofs[0].proof[0] = randValue;

        cheats.expectRevert(BeaconChainProofs.InvalidProof.selector);
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
    }

    /// @notice test that verifyCheckpointProofs skips proofs submitted for non-ACTIVE validators
    function testFuzz_verifyCheckpointProofs_skipIfNotActive(uint rand) public {
        // Setup verifyCheckpointProofs
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        // Ensure we have more than one validator (_newEigenPodStaker allocates a nonzero amt of eth)
        cheats.deal(address(staker), address(staker).balance + 32 ether);
        EigenPod pod = staker.pod();

        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);

        // Exit a validator and advance epoch so the exit is picked up next checkpoint
        uint40[] memory exitedValidator = new uint40[](1);
        exitedValidator[0] = validators[0];
        uint64 exitedBalanceGwei = staker.exitValidators(exitedValidator);
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(exitedValidator, pod.currentCheckpointTimestamp());

        // verify checkpoint proof for one exited validator
        // manually create a snapshot here for Snap checks
        timeMachine.createSnapshot();
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        assertEq(0, pod.withdrawableRestakedExecutionLayerGwei(), "should not have updated withdrawable balance");
        assertEq(pod.currentCheckpoint().proofsRemaining, validators.length - 1, "should have decreased proofs remaining by 1");
        assert_Snap_Removed_ActiveValidatorCount(staker, 1, "should have removed one validator from active set");
        assert_Snap_Removed_ActiveValidators(staker, exitedValidator, "should have set validator status to WITHDRAWN");

        // attempt to submit the same proof and ensure that checkpoint did not progress
        // the call should succeed, but nothing should happen
        // manually create a snapshot here for Snap checks
        timeMachine.createSnapshot();
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});

        assertEq(0, pod.withdrawableRestakedExecutionLayerGwei(), "should not have updated withdrawable balance");
        assertEq(pod.currentCheckpoint().proofsRemaining, validators.length - 1, "should not have decreased proofs remaining");
        assert_Snap_Unchanged_ActiveValidatorCount(staker, "should have the same active validator count");

        // finally, finish the checkpoint by submitting all proofs
        proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());
        // manually create a snapshot here for Snap checks
        timeMachine.createSnapshot();
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});

        assert_Snap_Unchanged_ActiveValidatorCount(staker, "should have the same active validator count after completing checkpoint");
        assertEq(exitedBalanceGwei, pod.withdrawableRestakedExecutionLayerGwei(), "exited balance should now be withdrawable");
        assertEq(pod.currentCheckpointTimestamp(), 0, "checkpoint should be complete");
    }

    /// @notice test that verifyCheckpointProofs skips duplicate checkpoint proofs
    function testFuzz_verifyCheckpointProofs_skipIfAlreadyProven(uint rand) public {
        // Setup verifyCheckpointProofs
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        // Ensure we have more than one validator (_newEigenPodStaker allocates a nonzero amt of eth)
        cheats.deal(address(staker), address(staker).balance + 32 ether);
        EigenPod pod = staker.pod();

        (uint40[] memory validators,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdraw(); // generate rewards on the beacon chain
        staker.verifyWithdrawalCredentials(validators);

        // select a single validator to submit multiple times
        uint40[] memory singleValidator = new uint40[](1);
        singleValidator[0] = validators[0];

        staker.startCheckpoint();

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(singleValidator, pod.currentCheckpointTimestamp());

        // verify checkpoint proof for one validator
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        assertEq(pod.currentCheckpoint().proofsRemaining, validators.length - 1, "should have decreased proofs remaining by 1");

        // attempt to submit the same proof and ensure that checkpoint did not progress
        // the call should succeed, but nothing should happen
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        assertEq(pod.currentCheckpoint().proofsRemaining, validators.length - 1, "should not have decreased proofs remaining");

        // finally, finish the checkpoint by submitting all proofs
        proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});

        assertEq(pod.currentCheckpointTimestamp(), 0, "checkpoint should be complete");
        // Check ValidatorInfo values for each validator
        for (uint i = 0; i < validators.length; i++) {
            bytes32 pubkeyHash = beaconChain.pubkeyHash(validators[i]);

            IEigenPodTypes.ValidatorInfo memory info = pod.validatorPubkeyHashToInfo(pubkeyHash);
            assertEq(info.restakedBalanceGwei, beaconChain.currentBalance(validators[i]), "should have restaked full current balance");
            assertEq(info.lastCheckpointedAt, pod.lastCheckpointTimestamp(), "should have recorded correct update time");
        }
    }

    /// @notice test that verifyCheckpointProofs sets validators to WITHDRAWN if they are exited
    function testFuzz_verifyCheckpointProofs_validatorExits(uint rand) public {
        // Setup verifyCheckpointProofs
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);

        // Exit validators and advance epoch so exits are picked up in next checkpoint
        uint64 exitedBalanceGwei = staker.exitValidators(validators);
        beaconChain.advanceEpoch_NoRewards();

        staker.startCheckpoint();

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());

        // Verify checkpoint proofs emit the expected values
        _expectEventsVerifyCheckpointProofs(staker, validators, proofs.balanceProofs);
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        assertEq(pod.currentCheckpointTimestamp(), 0, "checkpoint should be complete");

        assertEq(pod.withdrawableRestakedExecutionLayerGwei(), exitedBalanceGwei, "exited balance should be withdrawable");

        // Check ValidatorInfo values for each validator
        for (uint i = 0; i < validators.length; i++) {
            bytes32 pubkeyHash = beaconChain.pubkeyHash(validators[i]);

            IEigenPodTypes.ValidatorInfo memory info = pod.validatorPubkeyHashToInfo(pubkeyHash);
            assertEq(info.restakedBalanceGwei, 0, "should have 0 restaked balance");
            assertEq(info.lastCheckpointedAt, pod.lastCheckpointTimestamp(), "should have recorded correct update time");
            assertTrue(info.status == IEigenPodTypes.VALIDATOR_STATUS.WITHDRAWN, "should have recorded correct update time");
        }
    }

    /// @notice fuzz test an eigenPod with multiple validators and verifyCheckpointProofs
    function testFuzz_verifyCheckpointProofs(uint rand, bool epochRewards) public {
        // Setup verifyCheckpointProofs
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        staker.verifyWithdrawalCredentials(validators);
        if (epochRewards) beaconChain.advanceEpoch();
        else beaconChain.advanceEpoch_NoRewards();
        staker.startCheckpoint();

        CheckpointProofs memory proofs = beaconChain.getCheckpointProofs(validators, pod.currentCheckpointTimestamp());

        // Verify checkpoint proofs emit the expected values
        _expectEventsVerifyCheckpointProofs(staker, validators, proofs.balanceProofs);
        pod.verifyCheckpointProofs({balanceContainerProof: proofs.balanceContainerProof, proofs: proofs.balanceProofs});
        assertEq(pod.currentCheckpointTimestamp(), 0, "checkpoint should be complete");
    }
}

contract EigenPodUnitTests_verifyStaleBalance is EigenPodUnitTests {
    /// @notice test verifyStaleBalance reverts when paused
    function testFuzz_revert_verifyStaleBalancePaused(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validators[0]);

        cheats.prank(pauser);
        eigenPodManagerMock.pause(2 ** PAUSED_VERIFY_STALE_BALANCE);
        cheats.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
    }

    /// @notice test verifyStaleBalance reverts when paused via the PAUSED_START_CHECKPOINT flag
    function testFuzz_revert_verifyStaleBalancePausedViaStartCheckpoint(uint rand) public {
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validators[0]);

        cheats.prank(pauser);
        eigenPodManagerMock.pause(2 ** PAUSED_START_CHECKPOINT);
        cheats.expectRevert(IEigenPodErrors.CurrentlyPaused.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
    }

    /// @notice verifyStaleBalance should revert if validator balance too stale
    function testFuzz_revert_validatorBalanceNotStale(uint rand) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];
        beaconChain.advanceEpoch();
        staker.verifyWithdrawalCredentials(validators);

        staker.startCheckpoint();
        staker.completeCheckpoint();

        uint64 lastCheckpointTimestamp = pod.lastCheckpointTimestamp();
        // proof for given beaconTimestamp is not yet stale, this should revert
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);
        cheats.expectRevert(IEigenPodErrors.BeaconTimestampTooFarInPast.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: lastCheckpointTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
    }

    /// @notice checks staleness condition when a pod has never completed a checkpoint before
    /// The only value that will result in a revert here is `beaconTimestamp == 0`
    function testFuzz_revert_validatorBalanceNotStale_NeverCheckpointed(uint rand) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];
        beaconChain.advanceEpoch();
        staker.verifyWithdrawalCredentials(validators);

        // proof for given beaconTimestamp is not yet stale, this should revert
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);
        cheats.expectRevert(IEigenPodErrors.BeaconTimestampTooFarInPast.selector);
        pod.verifyStaleBalance({beaconTimestamp: 0, stateRootProof: proofs.stateRootProof, proof: proofs.validatorProof});
    }

    /// @notice verifyStaleBalance should revert if validator status is not ACTIVE
    function testFuzz_revert_validatorStatusNotActive(uint rand) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];

        // Advance epoch and use stale balance proof without verifyingWithdrawalCredentials
        // validator should be INACTIVE and cause a revert
        beaconChain.advanceEpoch();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);

        cheats.expectRevert(IEigenPodErrors.ValidatorNotActiveInPod.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
    }

    /// @notice verifyStaleBalance should revert if validator is not slashed
    function testFuzz_revert_validatorNotSlashed(uint rand) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];
        staker.verifyWithdrawalCredentials(validators);

        // Advance epoch and use stale balance proof where the validator has not been slashed
        // this should cause a revert
        beaconChain.advanceEpoch();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);

        cheats.expectRevert(IEigenPodErrors.ValidatorNotSlashedOnBeaconChain.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
    }

    /// @notice verifyStaleBalance should revert with invalid beaconStateRoot proof length
    function testFuzz_revert_beaconStateRootProofInvalidLength(uint rand) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];
        staker.verifyWithdrawalCredentials(validators);

        // Slash validators and advance epoch
        beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);

        // change the proof to have an invalid length
        bytes memory proofWithInvalidLength = new bytes(proofs.stateRootProof.proof.length + 1);
        BeaconChainProofs.StateRootProof memory invalidStateRootProof =
            BeaconChainProofs.StateRootProof({beaconStateRoot: proofs.stateRootProof.beaconStateRoot, proof: proofWithInvalidLength});

        cheats.expectRevert(BeaconChainProofs.InvalidProofLength.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: invalidStateRootProof,
            proof: proofs.validatorProof
        });
    }

    /// @notice verifyStaleBalance should revert with invalid beaconStateRoot proof
    function testFuzz_revert_beaconStateRootProofInvalid(uint rand) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];
        staker.verifyWithdrawalCredentials(validators);

        // Slash validators and advance epoch
        beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);

        // change the proof to have an invalid value
        bytes1 randValue = bytes1(keccak256(abi.encodePacked(proofs.stateRootProof.proof[0])));
        uint proofLength = proofs.stateRootProof.proof.length;
        uint randIndex = bound(rand, 0, proofLength - 1);
        proofs.stateRootProof.proof[randIndex] = randValue;

        cheats.expectRevert(BeaconChainProofs.InvalidProof.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
    }

    /// @notice verifyStaleBalance should revert with invalid validatorFields and validator proof length
    function testFuzz_revert_validatorContainerProofInvalidLength(uint rand) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];
        staker.verifyWithdrawalCredentials(validators);

        // Slash validators and advance epoch
        beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);

        // change the proof to have an invalid length
        uint proofLength = proofs.validatorProof.proof.length;
        bytes memory invalidProof = new bytes(proofLength + 1);
        BeaconChainProofs.ValidatorProof memory invalidValidatorProof =
            BeaconChainProofs.ValidatorProof({validatorFields: proofs.validatorProof.validatorFields, proof: invalidProof});
        cheats.expectRevert(BeaconChainProofs.InvalidProofLength.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: invalidValidatorProof
        });

        // Change the validator fields to have an invalid length
        bytes32[] memory validatorFieldsInvalidLength = new bytes32[](proofs.validatorProof.validatorFields.length + 1);
        for (uint i = 0; i < proofs.validatorProof.validatorFields.length; i++) {
            validatorFieldsInvalidLength[i] = proofs.validatorProof.validatorFields[i];
        }
        proofs.validatorProof.validatorFields = validatorFieldsInvalidLength;

        cheats.expectRevert(BeaconChainProofs.InvalidValidatorFieldsLength.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
    }

    /// @notice verifyStaleBalance should revert with invalid validatorContainer proof
    function testFuzz_revert_validatorContainerProofInvalid(uint rand, bytes32 randWithdrawalCredentials) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];
        staker.verifyWithdrawalCredentials(validators);

        // Slash validators and advance epoch
        beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);

        // change validator withdrawal creds to an invalid value causing a revert
        uint VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX = 1;
        proofs.validatorProof.validatorFields[VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX] = randWithdrawalCredentials;

        cheats.expectRevert(BeaconChainProofs.InvalidProof.selector);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
    }

    function testFuzz_verifyStaleBalance(uint rand) public {
        // setup eigenpod staker and validators
        (EigenPodUser staker,) = _newEigenPodStaker({rand: rand});
        EigenPod pod = staker.pod();
        (uint40[] memory validators,) = staker.startValidators();
        uint40 validator = validators[0];
        staker.verifyWithdrawalCredentials(validators);

        // Slash validators and advance epoch
        beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch();
        StaleBalanceProofs memory proofs = beaconChain.getStaleBalanceProofs(validator);

        cheats.expectEmit(true, true, true, true, address(staker.pod()));
        emit CheckpointCreated(uint64(block.timestamp), EIP_4788_ORACLE.timestampToBlockRoot(block.timestamp), validators.length);
        pod.verifyStaleBalance({
            beaconTimestamp: proofs.beaconTimestamp,
            stateRootProof: proofs.stateRootProof,
            proof: proofs.validatorProof
        });
        check_StartCheckpoint_State(staker);
    }
}

contract EigenPodHarnessSetup is EigenPodUnitTests {
    // Harness that exposes internal functions for test
    EigenPodHarness public eigenPodHarnessImplementation;
    EigenPodHarness public eigenPodHarness;

    function setUp() public virtual override {
        EigenPodUnitTests.setUp();

        // Deploy EP Harness
        eigenPodHarnessImplementation =
            new EigenPodHarness(ethPOSDepositMock, IEigenPodManager(address(eigenPodManagerMock)), GENESIS_TIME_LOCAL, "v9.9.9");

        // Upgrade eigenPod to harness
        UpgradeableBeacon(address(eigenPodBeacon)).upgradeTo(address(eigenPodHarnessImplementation));
        eigenPodHarness = EigenPodHarness(payable(eigenPod));
    }
}

/// @notice No unit tests as of now but would be good to add specific unit tests using proofs from our proofGen library
/// for a EigenPod on Holesky
contract EigenPodUnitTests_proofParsingTests is EigenPodHarnessSetup, ProofParsing {
    using BytesLib for bytes;
    using BeaconChainProofs for *;

    // Params to _verifyWithdrawalCredentials, can be set in test or helper function
    uint64 oracleTimestamp;
    bytes32 beaconStateRoot;
    uint40 validatorIndex;
    bytes validatorFieldsProof;
    bytes32[] validatorFields;

    function _assertWithdrawalCredentialsSet(uint restakedBalanceGwei) internal view {
        IEigenPodTypes.ValidatorInfo memory validatorInfo = eigenPodHarness.validatorPubkeyHashToInfo(validatorFields[0]);
        assertEq(uint8(validatorInfo.status), uint8(IEigenPodTypes.VALIDATOR_STATUS.ACTIVE), "Validator status should be active");
        assertEq(validatorInfo.validatorIndex, validatorIndex, "Validator index incorrectly set");
        assertEq(validatorInfo.lastCheckpointedAt, oracleTimestamp, "Last checkpointed at timestamp incorrectly set");
        assertEq(validatorInfo.restakedBalanceGwei, restakedBalanceGwei, "Restaked balance gwei not set correctly");
    }

    function _setWithdrawalCredentialParams() public {
        // Set beacon state root, validatorIndex
        beaconStateRoot = getBeaconStateRoot();
        validatorIndex = uint40(getValidatorIndex());
        validatorFieldsProof = getWithdrawalCredentialProof(); // Validator fields are proven here
        validatorFields = getValidatorFields();

        // Get an oracle timestamp
        cheats.warp(GENESIS_TIME_LOCAL + 1 days);
        oracleTimestamp = uint64(block.timestamp);
    }

    ///@notice Effective balance is > 32 ETH
    modifier setWithdrawalCredentialsExcess() {
        // Set JSON and params
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        // Set beacon state root, validatorIndex
        beaconStateRoot = getBeaconStateRoot();
        validatorIndex = uint40(getValidatorIndex());
        validatorFieldsProof = getWithdrawalCredentialProof(); // Validator fields are proven here
        validatorFields = getValidatorFields();

        // Get an oracle timestamp
        cheats.warp(GENESIS_TIME_LOCAL + 1 days);
        oracleTimestamp = uint64(block.timestamp);

        eigenPodHarness.verifyWithdrawalCredentials(beaconStateRoot, validatorIndex, validatorFieldsProof, validatorFields);
        _;
    }
}
