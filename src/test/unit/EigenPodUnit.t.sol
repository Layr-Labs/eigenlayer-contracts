// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

import "src/contracts/pods/EigenPod.sol";

import "src/test/mocks/ETHDepositMock.sol";
import "src/test/mocks/DelayedWithdrawalRouterMock.sol";
import "src/test/mocks/ERC20Mock.sol";
import "src/test/harnesses/EigenPodHarness.sol";
import "src/test/utils/ProofParsing.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/events/IEigenPodEvents.sol";

contract EigenPodUnitTests is EigenLayerUnitTestSetup {
    // Contract Under Test: EigenPod
    EigenPod public eigenPod;
    EigenPod public podImplementation;
    IBeacon public eigenPodBeacon;
    
    // Mocks
    IETHPOSDeposit public ethPOSDepositMock;
    IDelayedWithdrawalRouter public delayedWithdrawalRouterMock;
    
    // Address of pod for which proofs were generated
    address podAddress = address(0x49c486E3f4303bc11C02F952Fe5b08D0AB22D443);


    bool IS_DENEB = false;
    
    // Constants
    // uint32 public constant WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
    uint64 public constant MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32e9;
    // uint64 public constant RESTAKED_BALANCE_OFFSET_GWEI = 75e7;
    uint64 public constant GOERLI_GENESIS_TIME = 1616508000;
    // uint64 public constant SECONDS_PER_SLOT = 12;
    
    bytes internal constant beaconProxyBytecode =
        hex"608060405260405161090e38038061090e83398101604081905261002291610460565b61002e82826000610035565b505061058a565b61003e83610100565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100fb576100f9836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e99190610520565b836102a360201b6100291760201c565b505b505050565b610113816102cf60201b6100551760201c565b6101725760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101e6816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d79190610520565b6102cf60201b6100551760201c565b61024b5760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610169565b806102827fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6102de60201b6100641760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606102c883836040518060600160405280602781526020016108e7602791396102e1565b9392505050565b6001600160a01b03163b151590565b90565b6060600080856001600160a01b0316856040516102fe919061053b565b600060405180830381855af49150503d8060008114610339576040519150601f19603f3d011682016040523d82523d6000602084013e61033e565b606091505b5090925090506103508683838761035a565b9695505050505050565b606083156103c65782516103bf576001600160a01b0385163b6103bf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610169565b50816103d0565b6103d083836103d8565b949350505050565b8151156103e85781518083602001fd5b8060405162461bcd60e51b81526004016101699190610557565b80516001600160a01b038116811461041957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561044f578181015183820152602001610437565b838111156100f95750506000910152565b6000806040838503121561047357600080fd5b61047c83610402565b60208401519092506001600160401b038082111561049957600080fd5b818501915085601f8301126104ad57600080fd5b8151818111156104bf576104bf61041e565b604051601f8201601f19908116603f011681019083821181831017156104e7576104e761041e565b8160405282815288602084870101111561050057600080fd5b610511836020830160208801610434565b80955050505050509250929050565b60006020828403121561053257600080fd5b6102c882610402565b6000825161054d818460208701610434565b9190910192915050565b6020815260008251806020840152610576816040850160208701610434565b601f01601f19169190910160400192915050565b61034e806105996000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610067565b610100565b565b606061004e83836040518060600160405280602781526020016102f260279139610124565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fb9190610249565b905090565b3660008037600080366000845af43d6000803e80801561011f573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161014191906102a2565b600060405180830381855af49150503d806000811461017c576040519150601f19603f3d011682016040523d82523d6000602084013e610181565b606091505b50915091506101928683838761019c565b9695505050505050565b6060831561020d578251610206576001600160a01b0385163b6102065760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610217565b610217838361021f565b949350505050565b81511561022f5781518083602001fd5b8060405162461bcd60e51b81526004016101fd91906102be565b60006020828403121561025b57600080fd5b81516001600160a01b038116811461004e57600080fd5b60005b8381101561028d578181015183820152602001610275565b8381111561029c576000848401525b50505050565b600082516102b4818460208701610272565b9190910192915050565b60208152600082518060208401526102dd816040850160208701610272565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d51e81d3bc5ed20a26aeb05dce7e825c503b2061aa78628027300c8d65b9d89a64736f6c634300080c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564";
    address public podOwner = address(this);

    function setUp() public override virtual {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy mocks
        ethPOSDepositMock = new ETHPOSDepositMock();
        delayedWithdrawalRouterMock = new DelayedWithdrawalRouterMock();

        // Deploy EigenPod
        podImplementation = new EigenPod(
            ethPOSDepositMock,
            delayedWithdrawalRouterMock,
            eigenPodManagerMock,
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );

        // Deploy Beacon
        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));

        // Deploy Proxy same way as EigenPodManager does
        eigenPod = EigenPod(payable(
            Create2.deploy(
                0,
                bytes32(uint256(uint160(address(this)))),
                // set the beacon address to the eigenPodBeacon
                abi.encodePacked(beaconProxyBytecode, abi.encode(eigenPodBeacon, ""))
        )));

        // Etch the eigenPod code to the address for which proofs are generated
        bytes memory code = address(eigenPod).code;
        cheats.etch(podAddress, code);
        eigenPod = EigenPod(payable(podAddress));

        // Store the eigenPodBeacon address in the eigenPod beacon proxy
        bytes32 beaconSlot = 0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50;   
        cheats.store(address(eigenPod), beaconSlot, bytes32(uint256(uint160(address(eigenPodBeacon)))));

        // Initialize pod
        eigenPod.initialize(address(this));
    }


    /// @notice Post-M2, all new deployed eigen pods will have restaked set to true
    modifier hasNotRestaked() {
        // Write hasRestaked as false. hasRestaked in slot 52
        bytes32 slot = bytes32(uint256(52)); 
        bytes32 value = bytes32(0); // 0 == false
        cheats.store(address(eigenPod), slot, value);
        _;
    }
}

contract EigenPodUnitTests_Initialization is EigenPodUnitTests, IEigenPodEvents {

    function test_initialization() public {
        // Check podOwner and restaked
        assertEq(eigenPod.podOwner(), podOwner, "Pod owner incorrectly set");
        assertTrue(eigenPod.hasRestaked(), "hasRestaked incorrectly set");
        // Check immutable storage
        assertEq(address(eigenPod.ethPOS()), address(ethPOSDepositMock), "EthPOS incorrectly set");
        assertEq(address(eigenPod.delayedWithdrawalRouter()), address(delayedWithdrawalRouterMock), "DelayedWithdrawalRouter incorrectly set");
        assertEq(address(eigenPod.eigenPodManager()), address(eigenPodManagerMock), "EigenPodManager incorrectly set");
        assertEq(eigenPod.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR(), MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Max restaked balance incorrectly set");
        assertEq(eigenPod.GENESIS_TIME(), GOERLI_GENESIS_TIME, "Goerli genesis time incorrectly set");
    }

    function test_initialize_revert_alreadyInitialized() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        eigenPod.initialize(podOwner);
    }

    function test_initialize_eventEmitted() public {
        address newPodOwner = address(0x123);

        // Deploy new pod
        EigenPod newEigenPod = EigenPod(payable(
            Create2.deploy(
                0,
                bytes32(uint256(uint160(newPodOwner))),
                // set the beacon address to the eigenPodBeacon
                abi.encodePacked(beaconProxyBytecode, abi.encode(eigenPodBeacon, ""))
        )));

        // Expect emit and Initialize new pod
        vm.expectEmit(true, true, true, true);
        emit RestakingActivated(newPodOwner);
        newEigenPod.initialize(newPodOwner);
    }
}

contract EigenPodUnitTests_Stake is EigenPodUnitTests, IEigenPodEvents {

    // Beacon chain staking constnats
    bytes public constant pubkey =
        hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";
    bytes public signature;
    bytes32 public depositDataRoot;

    function testFuzz_stake_revert_notEigenPodManager(address invalidCaller) public {
        cheats.assume(invalidCaller != address(eigenPodManagerMock));
        cheats.deal(invalidCaller, 32 ether);

        cheats.prank(invalidCaller);
        cheats.expectRevert("EigenPod.onlyEigenPodManager: not eigenPodManager");
        eigenPod.stake{value: 32 ether}(pubkey, signature, depositDataRoot);
    }


    function testFuzz_stake_revert_invalidValue(uint256 value) public {
        cheats.assume(value != 32 ether);
        cheats.deal(address(eigenPodManagerMock), value);

        cheats.prank(address(eigenPodManagerMock));
        cheats.expectRevert("EigenPod.stake: must initially stake for any validator with 32 ether");
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
}

contract EigenPodUnitTests_PodOwnerFunctions is EigenPodUnitTests, IEigenPodEvents {
    
    /*******************************************************************************
                            Withdraw Non Beacon Chain ETH Tests
    *******************************************************************************/

    function testFuzz_withdrawNonBeaconChainETH_revert_notPodOwner(address invalidCaller) public {
        cheats.assume(invalidCaller != podOwner);

        cheats.prank(invalidCaller);
        cheats.expectRevert("EigenPod.onlyEigenPodOwner: not podOwner");
        eigenPod.withdrawNonBeaconChainETHBalanceWei(invalidCaller, 1 ether);
    }

    function test_withdrawNonBeaconChainETH_revert_tooMuchWithdrawn() public {
        // Send EigenPod 0.9 ether
        _seedPodWithETH(0.9 ether);

        // Withdraw 1 ether
        cheats.expectRevert("EigenPod.withdrawnonBeaconChainETHBalanceWei: amountToWithdraw is greater than nonBeaconChainETHBalanceWei");
        eigenPod.withdrawNonBeaconChainETHBalanceWei(podOwner, 1 ether);
    }

    function testFuzz_withdrawNonBeaconChainETH(uint256 ethAmount) public {
        _seedPodWithETH(ethAmount);
        assertEq(eigenPod.nonBeaconChainETHBalanceWei(), ethAmount, "Incorrect amount incremented in receive function");

        cheats.expectEmit(true, true, true, true);
        emit NonBeaconChainETHWithdrawn(podOwner, ethAmount);
        eigenPod.withdrawNonBeaconChainETHBalanceWei(podOwner, ethAmount);

        // Checks
        assertEq(address(eigenPod).balance, 0, "Incorrect amount withdrawn from eigenPod");
        assertEq(address(delayedWithdrawalRouterMock).balance, ethAmount, "Incorrect amount set to delayed withdrawal router");
    }

    /*******************************************************************************
                                  Recover Tokens Tests
    *******************************************************************************/
    
    function testFuzz_recoverTokens_revert_notPodOwner(address invalidCaller) public {
        cheats.assume(invalidCaller != podOwner);

        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(address(0x123));
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = 1;
        
        cheats.prank(invalidCaller);
        cheats.expectRevert("EigenPod.onlyEigenPodOwner: not podOwner");
        eigenPod.recoverTokens(tokens, amounts, podOwner);
    }

    function test_recoverTokens_revert_invalidLengths() public {
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(address(0x123));
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 1;
        amounts[1] = 1;

        cheats.expectRevert("EigenPod.recoverTokens: tokenList and amountsToWithdraw must be same length");
        eigenPod.recoverTokens(tokens, amounts, podOwner);
    }

    function test_recoverTokens() public {
        // Deploy dummy token
        IERC20 dummyToken = new ERC20Mock();
        dummyToken.transfer(address(eigenPod), 1e18);

        // Recover tokens
        address recipient = address(0x123);
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = dummyToken;
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = 1e18;

        eigenPod.recoverTokens(tokens, amounts, recipient);

        // Checks
        assertEq(dummyToken.balanceOf(recipient), 1e18, "Incorrect amount recovered");
    }

    /*******************************************************************************
                             Activate Restaking Tests
    *******************************************************************************/

    function testFuzz_activateRestaking_revert_notPodOwner(address invalidCaller) public {
        cheats.assume(invalidCaller != podOwner);

        cheats.prank(invalidCaller);
        cheats.expectRevert("EigenPod.onlyEigenPodOwner: not podOwner");
        eigenPod.activateRestaking();
    }

    function test_activateRestaking_revert_alreadyRestaked() public {
        cheats.expectRevert("EigenPod.hasNeverRestaked: restaking is enabled");
        eigenPod.activateRestaking();
    }

    function testFuzz_activateRestaking(uint256 ethAmount) public hasNotRestaked {
        // Seed some ETH
        _seedPodWithETH(ethAmount);

        // Activate restaking
        vm.expectEmit(true, true, true, true);
        emit RestakingActivated(podOwner);
        eigenPod.activateRestaking();

        // Checks
        assertTrue(eigenPod.hasRestaked(), "hasRestaked incorrectly set");
        _assertWithdrawalProcessed(ethAmount);
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
    function test_regression_validatorBalance_cannotBeRemoved_viaNonBeaconChainETHBalanceWei() external hasNotRestaked {
        // Assert that the pod has not restaked
        assertFalse(eigenPod.hasRestaked(), "hasRestaked should be false");

        // Simulate podOwner sending 32 ETH to eigenPod
        _seedPodWithETH(32 ether);
        assertEq(eigenPod.nonBeaconChainETHBalanceWei(), 32 ether, "nonBeaconChainETHBalanceWei should be 32 ETH");

        // Pod owner calls withdrawBeforeRestaking, sending 32 ETH to owner
        eigenPod.withdrawBeforeRestaking();
        assertEq(address(eigenPod).balance, 0, "eigenPod balance should be 0");
        assertEq(address(delayedWithdrawalRouterMock).balance, 32 ether, "withdrawal router balance should be 32 ETH");

        // Upgrade from m1 to m2

        // Activate restaking on the pod
        eigenPod.activateRestaking();

        // Simulate a withdrawal by increasing eth balance without code execution
        cheats.deal(address(eigenPod), 32 ether);

        // Try calling withdrawNonBeaconChainETHBalanceWei, should fail since `nonBeaconChainETHBalanceWei` 
        // was set to 0 when calling `_processWithdrawalBeforeRestaking` from `activateRestaking` 
        cheats.expectRevert("EigenPod.withdrawnonBeaconChainETHBalanceWei: amountToWithdraw is greater than nonBeaconChainETHBalanceWei");
        eigenPod.withdrawNonBeaconChainETHBalanceWei(podOwner, 32 ether);
    }
    
    /*******************************************************************************
                        Withdraw Before Restaking Tests
    *******************************************************************************/

    function testFuzz_withdrawBeforeRestaking_revert_notPodOwner(address invalidCaller) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.assume(invalidCaller != podOwner);

        cheats.prank(invalidCaller);
        cheats.expectRevert("EigenPod.onlyEigenPodOwner: not podOwner");
        eigenPod.withdrawBeforeRestaking();
    }

    function test_withdrawBeforeRestaking_revert_alreadyRestaked() public {
        cheats.expectRevert("EigenPod.hasNeverRestaked: restaking is enabled");
        eigenPod.withdrawBeforeRestaking();
    }

    function testFuzz_withdrawBeforeRestaking(uint256 ethAmount) public hasNotRestaked {
        // Seed some ETH
        _seedPodWithETH(ethAmount);

        // Withdraw
        eigenPod.withdrawBeforeRestaking();

        // Checks
        _assertWithdrawalProcessed(ethAmount);
    }

    // Helpers
    function _assertWithdrawalProcessed(uint256 amount) internal {
        assertEq(eigenPod.mostRecentWithdrawalTimestamp(), uint32(block.timestamp), "Incorrect mostRecentWithdrawalTimestamp");
        assertEq(eigenPod.nonBeaconChainETHBalanceWei(), 0, "Incorrect nonBeaconChainETHBalanceWei");
        assertEq(address(delayedWithdrawalRouterMock).balance, amount, "Incorrect amount sent to delayed withdrawal router");
    }

    function _seedPodWithETH(uint256 ethAmount) internal {
        cheats.deal(address(this), ethAmount);
        bool result;
        bytes memory data;
        (result, data) = address(eigenPod).call{value: ethAmount}("");
    }
}

contract EigenPodHarnessSetup is EigenPodUnitTests {
    // Harness that exposes internal functions for test
    EPInternalFunctions public eigenPodHarnessImplementation;
    EPInternalFunctions public eigenPodHarness;

    function setUp() public virtual override {
        EigenPodUnitTests.setUp();

        // Deploy EP Harness
        eigenPodHarnessImplementation = new EPInternalFunctions(
            ethPOSDepositMock,
            delayedWithdrawalRouterMock,
            eigenPodManagerMock,
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );

        // Upgrade eigenPod to harness
        UpgradeableBeacon(address(eigenPodBeacon)).upgradeTo(address(eigenPodHarnessImplementation));
        eigenPodHarness = EPInternalFunctions(payable(eigenPod));
    }    
}

contract EigenPodUnitTests_VerifyWithdrawalCredentialsTests is EigenPodHarnessSetup, ProofParsing, IEigenPodEvents {
    using BytesLib for bytes;
    using BeaconChainProofs for *;

    // Params to _verifyWithdrawalCredentials, can be set in test or helper function
    uint64 oracleTimestamp;
    bytes32 beaconStateRoot;
    uint40 validatorIndex;
    bytes validatorFieldsProof;
    bytes32[] validatorFields;

    function test_revert_validatorActive() public {
        // Set JSON & params
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _setWithdrawalCredentialParams();

        // Set validator status to active
        eigenPodHarness.setValidatorStatus(validatorFields[0], IEigenPod.VALIDATOR_STATUS.ACTIVE);

        // Expect revert
        cheats.expectRevert(
            "EigenPod.verifyCorrectWithdrawalCredentials: Validator must be inactive to prove withdrawal credentials"
        );
        eigenPodHarness.verifyWithdrawalCredentials(
            oracleTimestamp,
            beaconStateRoot,
            validatorIndex,
            validatorFieldsProof,
            validatorFields
        );
    }

    function testFuzz_revert_invalidValidatorFields(address wrongWithdrawalAddress) public {
        cheats.assume(wrongWithdrawalAddress != address(eigenPodHarness));
        // Set JSON and params
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _setWithdrawalCredentialParams();

        // Change the withdrawal credentials in validatorFields, which is at index 1
        validatorFields[1] = abi.encodePacked(bytes1(uint8(1)), bytes11(0), wrongWithdrawalAddress).toBytes32(0);

        // Expect revert
        cheats.expectRevert(
            "EigenPod.verifyCorrectWithdrawalCredentials: Proof is not for this EigenPod"
        );
        eigenPodHarness.verifyWithdrawalCredentials(
            oracleTimestamp,
            beaconStateRoot,
            validatorIndex,
            validatorFieldsProof,
            validatorFields
        );
    }

    function test_effectiveBalanceGreaterThan32ETH() public {
        // Set JSON and params
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _setWithdrawalCredentialParams();
        
        // Check that restaked balance greater than 32 ETH
        uint64 effectiveBalanceGwei = validatorFields.getEffectiveBalanceGwei();
        assertGt(effectiveBalanceGwei, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Proof file has an effective balance less than 32 ETH");

        uint activeValidatorCountBefore = eigenPodHarness.getActiveValidatorCount();

        // Verify withdrawal credentials
        vm.expectEmit(true, true, true, true);
        emit ValidatorRestaked(validatorIndex);
        vm.expectEmit(true, true, true, true);
        emit ValidatorBalanceUpdated(validatorIndex, oracleTimestamp, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR);
        uint256 restakedBalanceWei = eigenPodHarness.verifyWithdrawalCredentials(
            oracleTimestamp,
            beaconStateRoot,
            validatorIndex,
            validatorFieldsProof,
            validatorFields
        );

        // Checks
        uint activeValidatorCountAfter = eigenPodHarness.getActiveValidatorCount();
        assertEq(activeValidatorCountAfter, activeValidatorCountBefore + 1, "active validator count should increase when proving withdrawal credentials");
        assertEq(restakedBalanceWei, uint256(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR) * uint256(1e9), "Returned restaked balance gwei should be max");
        _assertWithdrawalCredentialsSet(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR);
    }

    function test_effectiveBalanceLessThan32ETH() public {
        // Set JSON and params
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913_30ETHBalance.json");
        _setWithdrawalCredentialParams();
        
        // Check that restaked balance less than 32 ETH
        uint64 effectiveBalanceGwei = validatorFields.getEffectiveBalanceGwei();
        assertLt(effectiveBalanceGwei, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Proof file has an effective balance greater than 32 ETH");

        uint activeValidatorCountBefore = eigenPodHarness.getActiveValidatorCount();

        // Verify withdrawal credentials
        vm.expectEmit(true, true, true, true);
        emit ValidatorRestaked(validatorIndex);
        vm.expectEmit(true, true, true, true);
        emit ValidatorBalanceUpdated(validatorIndex, oracleTimestamp, effectiveBalanceGwei);
        uint256 restakedBalanceWei = eigenPodHarness.verifyWithdrawalCredentials(
            oracleTimestamp,
            beaconStateRoot,
            validatorIndex,
            validatorFieldsProof,
            validatorFields
        );

        // Checks
        uint activeValidatorCountAfter = eigenPodHarness.getActiveValidatorCount();
        assertEq(activeValidatorCountAfter, activeValidatorCountBefore + 1, "active validator count should increase when proving withdrawal credentials");
        assertEq(restakedBalanceWei, uint256(effectiveBalanceGwei) * uint256(1e9), "Returned restaked balance gwei incorrect");
        _assertWithdrawalCredentialsSet(effectiveBalanceGwei);
    }

    function _assertWithdrawalCredentialsSet(uint256 restakedBalanceGwei) internal {
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPodHarness.validatorPubkeyHashToInfo(validatorFields[0]);
        assertEq(uint8(validatorInfo.status), uint8(IEigenPod.VALIDATOR_STATUS.ACTIVE), "Validator status should be active");
        assertEq(validatorInfo.validatorIndex, validatorIndex, "Validator index incorrectly set");
        assertEq(validatorInfo.mostRecentBalanceUpdateTimestamp, oracleTimestamp, "Most recent balance update timestamp incorrectly set");
        assertEq(validatorInfo.restakedBalanceGwei, restakedBalanceGwei, "Restaked balance gwei not set correctly");
    }


    function _setWithdrawalCredentialParams() public {
        // Set beacon state root, validatorIndex
        beaconStateRoot = getBeaconStateRoot();
        validatorIndex = uint40(getValidatorIndex());
        validatorFieldsProof = abi.encodePacked(getWithdrawalCredentialProof()); // Validator fields are proven here
        validatorFields = getValidatorFields();

        // Get an oracle timestamp
        cheats.warp(GOERLI_GENESIS_TIME + 1 days);
        oracleTimestamp = uint64(block.timestamp);
    }
}

/// @notice In practice, this function should be called after a validator has verified their withdrawal credentials
contract EigenPodUnitTests_VerifyBalanceUpdateTests is EigenPodHarnessSetup, ProofParsing, IEigenPodEvents {
    using BeaconChainProofs for *;

    // Params to verifyBalanceUpdate, can be set in test or helper function
    uint64 oracleTimestamp;
    uint40 validatorIndex;
    bytes32 beaconStateRoot;
    bytes validatorFieldsProof;
    bytes32[] validatorFields;

    function testFuzz_revert_oracleTimestampStale(uint64 oracleFuzzTimestamp, uint64 mostRecentBalanceUpdateTimestamp) public {
        // Constain inputs and set proof file
        cheats.assume(oracleFuzzTimestamp < mostRecentBalanceUpdateTimestamp);
        setJSON("src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        
        // Get validator fields and balance update root
        validatorFields = getValidatorFields();
        validatorFieldsProof = abi.encodePacked(getBalanceUpdateProof());

        // Balance update reversion
        cheats.expectRevert(
            "EigenPod.verifyBalanceUpdate: Validators balance has already been updated for this timestamp"
        );
        eigenPodHarness.verifyBalanceUpdate(
            oracleFuzzTimestamp,
            0,
            bytes32(0),
            validatorFieldsProof,
            validatorFields,
            mostRecentBalanceUpdateTimestamp
        );
    }

    function test_revert_validatorInactive() public {
        // Set proof file
        setJSON("src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");

        // Set proof params
        _setBalanceUpdateParams();

        // Set validator status to inactive
        eigenPodHarness.setValidatorStatus(validatorFields[0], IEigenPod.VALIDATOR_STATUS.INACTIVE);

        // Balance update reversion
        cheats.expectRevert(
            "EigenPod.verifyBalanceUpdate: Validator not active"
        );
        eigenPodHarness.verifyBalanceUpdate(
            oracleTimestamp,
            validatorIndex,
            beaconStateRoot,
            validatorFieldsProof,
            validatorFields,
            0 // Most recent balance update timestamp set to 0
        );
    }

    /**
     * Regression test for a bug that allowed balance updates to be made for withdrawn validators. Thus
     * the validator's balance could be maliciously proven to be 0 before the validator themselves are
     * able to prove their withdrawal.
     */
    function test_revert_balanceUpdateAfterWithdrawableEpoch() external {
        // Set Json proof
        setJSON("src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        
        // Set proof params
        _setBalanceUpdateParams();
        
        // Set effective balance and  withdrawable epoch
        validatorFields[2] = bytes32(uint256(0)); // per consensus spec, slot 2 is effective balance
        validatorFields[7] = bytes32(uint256(0)); // per consensus spec, slot 7 is withdrawable epoch == 0
        
        console.log("withdrawable epoch: ", validatorFields.getWithdrawableEpoch());
        // Expect revert on balance update 
        cheats.expectRevert(bytes("EigenPod.verifyBalanceUpdate: validator is withdrawable but has not withdrawn"));
        eigenPodHarness.verifyBalanceUpdate(oracleTimestamp, validatorIndex, beaconStateRoot, validatorFieldsProof, validatorFields, 0);
    }

    /// @notice Rest of tests assume beacon chain proofs are correct; Now we update the validator's balance

    ///@notice Balance of validator is >= 32e9
    function test_positiveSharesDelta() public {
        // Set JSON
        setJSON("src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");

        // Set proof params
        _setBalanceUpdateParams();

        // Verify balance update
        vm.expectEmit(true, true, true, true);
        emit ValidatorBalanceUpdated(validatorIndex, oracleTimestamp, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR);
        int256 sharesDeltaGwei = eigenPodHarness.verifyBalanceUpdate(
            oracleTimestamp,
            validatorIndex,
            beaconStateRoot,
            validatorFieldsProof,
            validatorFields,
            0 // Most recent balance update timestamp set to 0
        );

        // Checks
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPodHarness.validatorPubkeyHashToInfo(validatorFields[0]);
        assertEq(validatorInfo.restakedBalanceGwei, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Restaked balance gwei should be max");
        assertGt(sharesDeltaGwei, 0, "Shares delta should be positive");
        assertEq(sharesDeltaGwei, int256(uint256(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR)), "Shares delta should be equal to restaked balance");
    }
    
    function test_negativeSharesDelta() public {
        // Set JSON
        setJSON("src/test/test-data/balanceUpdateProof_balance28ETH_302913.json");

        // Set proof params
        _setBalanceUpdateParams();
        uint64 newValidatorBalance = validatorFields.getEffectiveBalanceGwei();

        // Set balance of validator to max ETH
        eigenPodHarness.setValidatorRestakedBalance(validatorFields[0], MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR);

        // Verify balance update
        int256 sharesDeltaGwei = eigenPodHarness.verifyBalanceUpdate(
            oracleTimestamp,
            validatorIndex,
            beaconStateRoot,
            validatorFieldsProof,
            validatorFields,
            0 // Most recent balance update timestamp set to 0
        );

        // Checks
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPodHarness.validatorPubkeyHashToInfo(validatorFields[0]);
        assertEq(validatorInfo.restakedBalanceGwei, newValidatorBalance, "Restaked balance gwei should be max");
        assertLt(sharesDeltaGwei, 0, "Shares delta should be negative");
        int256 expectedSharesDiff = int256(uint256(newValidatorBalance)) - int256(uint256(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR));
        assertEq(sharesDeltaGwei, expectedSharesDiff, "Shares delta should be equal to restaked balance");
    }

    function test_zeroSharesDelta() public {
        // Set JSON
        setJSON("src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");

        // Set proof params
        _setBalanceUpdateParams();

        // Set previous restaked balance to max restaked balance
        eigenPodHarness.setValidatorRestakedBalance(validatorFields[0], MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR);

        // Verify balance update
        int256 sharesDeltaGwei = eigenPodHarness.verifyBalanceUpdate(
            oracleTimestamp,
            validatorIndex,
            beaconStateRoot,
            validatorFieldsProof,
            validatorFields,
            0 // Most recent balance update timestamp set to 0
        );

        // Checks
        assertEq(sharesDeltaGwei, 0, "Shares delta should be 0");
    }

    function _setBalanceUpdateParams() internal {
        // Set validator index, beacon state root, balance update proof, and validator fields
        validatorIndex = uint40(getValidatorIndex());
        beaconStateRoot = getBeaconStateRoot();
        validatorFieldsProof = abi.encodePacked(getBalanceUpdateProof());
        validatorFields = getValidatorFields();

        // Get an oracle timestamp
        cheats.warp(GOERLI_GENESIS_TIME + 1 days);
        oracleTimestamp = uint64(block.timestamp);

        // Set validator status to active
        eigenPodHarness.setValidatorStatus(validatorFields[0], IEigenPod.VALIDATOR_STATUS.ACTIVE);
    }
}

contract EigenPodUnitTests_WithdrawalTests is EigenPodHarnessSetup, ProofParsing, IEigenPodEvents {
    using BeaconChainProofs for *;

    // Params to process withdrawal 
    bytes32 beaconStateRoot;
    BeaconChainProofs.WithdrawalProof withdrawalToProve;
    bytes validatorFieldsProof;
    bytes32[] validatorFields;
    bytes32[] withdrawalFields;

    // Most recent withdrawal timestamp incremented when withdrawal processed before restaking OR when staking activated
    function test_verifyAndProcessWithdrawal_revert_staleProof() public hasNotRestaked {
        // Set JSON & params
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        _setWithdrawalProofParams();

        // Set timestamp to after withdrawal timestamp
        uint64 timestampOfWithdrawal = Endian.fromLittleEndianUint64(withdrawalToProve.timestampRoot);
        uint256 newTimestamp = timestampOfWithdrawal + 2500;
        cheats.warp(newTimestamp);

        // Activate restaking, setting `mostRecentWithdrawalTimestamp` 
        eigenPodHarness.activateRestaking();

        // Expect revert
        cheats.expectRevert("EigenPod.proofIsForValidTimestamp: beacon chain proof must be at or after mostRecentWithdrawalTimestamp");
        eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );
    }

    function test_verifyAndProcessWithdrawal_revert_statusInactive() public {
        // Set JSON & params
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        _setWithdrawalProofParams();

        // Set status to inactive
        eigenPodHarness.setValidatorStatus(validatorFields[0], IEigenPod.VALIDATOR_STATUS.INACTIVE);

        // Expect revert
        cheats.expectRevert("EigenPod._verifyAndProcessWithdrawal: Validator never proven to have withdrawal credentials pointed to this contract");
        eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );
    }

    function test_verifyAndProcessWithdrawal_withdrawalAlreadyProcessed() public setWithdrawalCredentialsExcess {
        // Set JSON & params
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        _setWithdrawalProofParams();

        // Process withdrawal
        eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );

        // Attempt to process again
        cheats.expectRevert("EigenPod._verifyAndProcessWithdrawal: withdrawal has already been proven for this timestamp");
        eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );
    }

    function test_verifyAndProcessWithdrawal_excess() public setWithdrawalCredentialsExcess {
        // Set JSON & params
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        _setWithdrawalProofParams();

        // Process withdrawal
        eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );

        // Verify storage
        bytes32 validatorPubKeyHash = validatorFields.getPubkeyHash();
        uint64 withdrawalTimestamp = withdrawalToProve.getWithdrawalTimestamp();
        assertTrue(eigenPodHarness.provenWithdrawal(validatorPubKeyHash, withdrawalTimestamp), "Withdrawal not set to proven");
    }

    // regression test for off-by-one error
    function test_verifyAndProcessWithdrawal_atLatestWithdrawalTimestamp() public setWithdrawalCredentialsExcess {
        // Set JSON & params
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        _setWithdrawalProofParams();

        uint64 withdrawalTimestamp = withdrawalToProve.getWithdrawalTimestamp();
        // set the `mostRecentWithdrawalTimestamp` to be equal to the withdrawal timestamp
        eigenPodHarness.setMostRecentWithdrawalTimestamp(withdrawalTimestamp);

        // Process withdrawal
        eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );

        // Verify storage
        bytes32 validatorPubKeyHash = validatorFields.getPubkeyHash();
        assertTrue(eigenPodHarness.provenWithdrawal(validatorPubKeyHash, withdrawalTimestamp), "Withdrawal not set to proven");
    }

    function test_revert_verifyAndProcessWithdrawal_beforeLatestWithdrawalTimestamp() public setWithdrawalCredentialsExcess {
        // Set JSON & params
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        _setWithdrawalProofParams();

        uint64 withdrawalTimestamp = withdrawalToProve.getWithdrawalTimestamp();
        // set the `mostRecentWithdrawalTimestamp` to just after the withdrawal timestamp
        eigenPodHarness.setMostRecentWithdrawalTimestamp(withdrawalTimestamp + 1);

        // Process withdrawal, expect revert
        cheats.expectRevert("EigenPod.proofIsForValidTimestamp: beacon chain proof must be at or after mostRecentWithdrawalTimestamp");
        eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );
    }

    /// @notice Tests processing a full withdrawal > MAX_RESTAKED_GWEI_PER_VALIDATOR
    function test_processFullWithdrawal_excess32ETH() public setWithdrawalCredentialsExcess {
        // Set JSON & params
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        _setWithdrawalProofParams();

        // Get params to check against
        uint64 withdrawalTimestamp = withdrawalToProve.getWithdrawalTimestamp();
        uint40 validatorIndex = uint40(getValidatorIndex());
        uint64 withdrawalAmountGwei = withdrawalFields.getWithdrawalAmountGwei();
        assertGt(withdrawalAmountGwei, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Withdrawal amount should be greater than max restaked balance for this test");

        // Process full withdrawal
        vm.expectEmit(true, true, true, true);
        emit FullWithdrawalRedeemed(validatorIndex, withdrawalTimestamp, podOwner, withdrawalAmountGwei);
        IEigenPod.VerifiedWithdrawal memory vw = eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );

        // Storage checks in _verifyAndProcessWithdrawal
        bytes32 validatorPubKeyHash = validatorFields.getPubkeyHash();
        assertTrue(eigenPodHarness.provenWithdrawal(validatorPubKeyHash, withdrawalTimestamp), "Withdrawal not set to proven");

        // Checks from  _processFullWithdrawal
        assertEq(eigenPod.withdrawableRestakedExecutionLayerGwei(), MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Incorrect withdrawable restaked execution layer gwei");
        // Excess withdrawal amount is diff between restaked balance and total withdrawal amount
        uint64 excessWithdrawalAmount = withdrawalAmountGwei - MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR;
        assertEq(vw.amountToSendGwei, excessWithdrawalAmount, "Amount to send via router is not correct");
        assertEq(vw.sharesDeltaGwei, 0, "Shares delta not correct"); // Shares delta is 0 since restaked balance and amount to withdraw were max
        
        // ValidatorInfo storage update checks
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPodHarness.validatorPubkeyHashToInfo(validatorFields[0]);
        assertEq(uint8(validatorInfo.status), uint8(IEigenPod.VALIDATOR_STATUS.WITHDRAWN), "Validator status should be withdrawn");
        assertEq(validatorInfo.restakedBalanceGwei, 0, "Restaked balance gwei should be 0");
    }

    function test_processFullWithdrawal_lessThan32ETH() public setWithdrawalCredentialsExcess {
        // Set JSON & params
        setJSON("src/test/test-data/fullWithdrawalProof_Latest_28ETH.json");
        _setWithdrawalProofParams();

        // Get params to check against
        uint64 withdrawalTimestamp = withdrawalToProve.getWithdrawalTimestamp();
        uint64 withdrawalAmountGwei = withdrawalFields.getWithdrawalAmountGwei();
        assertLt(withdrawalAmountGwei, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Withdrawal amount should be greater than max restaked balance for this test");

        // Process full withdrawal
        IEigenPod.VerifiedWithdrawal memory vw = eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );

        // Storage checks in _verifyAndProcessWithdrawal
        bytes32 validatorPubKeyHash = validatorFields.getPubkeyHash();
        assertTrue(eigenPodHarness.provenWithdrawal(validatorPubKeyHash, withdrawalTimestamp), "Withdrawal not set to proven");

        // Checks from  _processFullWithdrawal
        assertEq(eigenPod.withdrawableRestakedExecutionLayerGwei(), withdrawalAmountGwei, "Incorrect withdrawable restaked execution layer gwei");
        // Excess withdrawal amount should be 0 since balance is < MAX
        assertEq(vw.amountToSendGwei, 0, "Amount to send via router is not correct");
        int256 expectedSharesDiff = int256(uint256(withdrawalAmountGwei)) - int256(uint256(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR));
        assertEq(vw.sharesDeltaGwei, expectedSharesDiff, "Shares delta not correct"); // Shares delta is 0 since restaked balance and amount to withdraw were max
        
        // ValidatorInfo storage update checks
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPodHarness.validatorPubkeyHashToInfo(validatorFields[0]);
        assertEq(uint8(validatorInfo.status), uint8(IEigenPod.VALIDATOR_STATUS.WITHDRAWN), "Validator status should be withdrawn");
        assertEq(validatorInfo.restakedBalanceGwei, 0, "Restaked balance gwei should be 0");
    }

    function test_processPartialWithdrawal() public setWithdrawalCredentialsExcess {
        // Set JSON & params
        setJSON("./src/test/test-data/partialWithdrawalProof_Latest.json");
        _setWithdrawalProofParams();

        // Get params to check against
        uint64 withdrawalTimestamp = withdrawalToProve.getWithdrawalTimestamp();
        uint40 validatorIndex = uint40(getValidatorIndex());
        uint64 withdrawalAmountGwei = withdrawalFields.getWithdrawalAmountGwei();
        
        // Assert that partial withdrawal code path will be tested
        assertLt(withdrawalToProve.getWithdrawalEpoch(), validatorFields.getWithdrawableEpoch(), "Withdrawal epoch should be less than the withdrawable epoch");

        // Process partial withdrawal
        vm.expectEmit(true, true, true, true);
        emit PartialWithdrawalRedeemed(validatorIndex, withdrawalTimestamp, podOwner, withdrawalAmountGwei);
        IEigenPod.VerifiedWithdrawal memory vw = eigenPodHarness.verifyAndProcessWithdrawal(
            beaconStateRoot,
            withdrawalToProve,
            validatorFieldsProof,
            validatorFields,
            withdrawalFields
        );

        // Storage checks in _verifyAndProcessWithdrawal
        bytes32 validatorPubKeyHash = validatorFields.getPubkeyHash();
        assertTrue(eigenPodHarness.provenWithdrawal(validatorPubKeyHash, withdrawalTimestamp), "Withdrawal not set to proven");

        // Checks from  _processPartialWithdrawal
        assertEq(eigenPod.sumOfPartialWithdrawalsClaimedGwei(), withdrawalAmountGwei, "Incorrect partial withdrawal amount");
        assertEq(vw.amountToSendGwei, withdrawalAmountGwei, "Amount to send via router is not correct");
        assertEq(vw.sharesDeltaGwei, 0, "Shares delta should be 0");

        // Assert validator still has same restaked balance and status
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPodHarness.validatorPubkeyHashToInfo(validatorFields[0]);
        assertEq(uint8(validatorInfo.status), uint8(IEigenPod.VALIDATOR_STATUS.ACTIVE), "Validator status should be active");
        assertEq(validatorInfo.restakedBalanceGwei, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Restaked balance gwei should be max");
    }

    function testFuzz_processFullWithdrawal(bytes32 pubkeyHash, uint64 restakedAmount, uint64 withdrawalAmount) public {
        // Format validatorInfo struct
        IEigenPod.ValidatorInfo memory validatorInfo = IEigenPod.ValidatorInfo({
            validatorIndex: 0,
            restakedBalanceGwei: restakedAmount,
            mostRecentBalanceUpdateTimestamp: 0,
            status: IEigenPod.VALIDATOR_STATUS.ACTIVE
        });

        // Since we're withdrawing using an ACTIVE validator, ensure we have
        // a validator count to decrement
        uint activeValidatorCountBefore = 1 + eigenPodHarness.getActiveValidatorCount();
        eigenPodHarness.setActiveValidatorCount(activeValidatorCountBefore);
        
        // Process full withdrawal.
        IEigenPod.VerifiedWithdrawal memory vw = eigenPodHarness.processFullWithdrawal(0, pubkeyHash, 0, podOwner, withdrawalAmount, validatorInfo);

        // Validate that our activeValidatorCount decreased
        uint activeValidatorCountAfter = eigenPodHarness.getActiveValidatorCount();
        assertEq(activeValidatorCountAfter, activeValidatorCountBefore - 1, "active validator count should decrease when withdrawing active validator");

        // Get expected amounts based on withdrawalAmount
        uint64 amountETHToQueue;
        uint64 amountETHToSend;
        if (withdrawalAmount > MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR){
            amountETHToQueue = MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR;
            amountETHToSend = withdrawalAmount - MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR;
        } else {
            amountETHToQueue = withdrawalAmount;
            amountETHToSend = 0;
        }

        // Check invariant-> amountToQueue + amountToSend = withdrawalAmount
        assertEq(vw.amountToSendGwei + eigenPod.withdrawableRestakedExecutionLayerGwei(), withdrawalAmount, "Amount to queue and send must add up to total withdrawal amount");

        // Check amount to queue and send
        assertEq(vw.amountToSendGwei, amountETHToSend, "Amount to queue is not correct");
        assertEq(eigenPod.withdrawableRestakedExecutionLayerGwei(), amountETHToQueue, "Incorrect withdrawable restaked execution layer gwei");

        // Check shares delta
        int256 expectedSharesDelta = int256(uint256(amountETHToQueue)) - int256(uint256(restakedAmount));
        assertEq(vw.sharesDeltaGwei, expectedSharesDelta, "Shares delta not correct");

        // Storage checks
        IEigenPod.ValidatorInfo memory validatorInfoAfter = eigenPodHarness.validatorPubkeyHashToInfo(pubkeyHash);
        assertEq(uint8(validatorInfoAfter.status), uint8(IEigenPod.VALIDATOR_STATUS.WITHDRAWN), "Validator status should be withdrawn");
        assertEq(validatorInfoAfter.restakedBalanceGwei, 0, "Restaked balance gwei should be 0");
    }

    function testFuzz_processFullWithdrawal_lessMaxRestakedBalance(bytes32 pubkeyHash, uint64 restakedAmount, uint64 withdrawalAmount) public {
        withdrawalAmount = uint64(bound(withdrawalAmount, 0, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR));
        testFuzz_processFullWithdrawal(pubkeyHash, restakedAmount, withdrawalAmount);
    }
    
    function testFuzz_processPartialWithdrawal(        
        uint40 validatorIndex,
        uint64 withdrawalTimestamp,
        address recipient,
        uint64 partialWithdrawalAmountGwei
    ) public {
        IEigenPod.VerifiedWithdrawal memory vw = eigenPodHarness.processPartialWithdrawal(validatorIndex, withdrawalTimestamp, recipient, partialWithdrawalAmountGwei);

        // Checks
        assertEq(eigenPod.sumOfPartialWithdrawalsClaimedGwei(), partialWithdrawalAmountGwei, "Incorrect partial withdrawal amount");
        assertEq(vw.amountToSendGwei, partialWithdrawalAmountGwei, "Amount to send via router is not correct");
        assertEq(vw.sharesDeltaGwei, 0, "Shares delta should be 0");
    }

    function _setWithdrawalProofParams() internal {
        // Set validator index, beacon state root, balance update proof, and validator fields
        beaconStateRoot = getBeaconStateRoot();
        validatorFields = getValidatorFields();
        validatorFieldsProof = abi.encodePacked(getValidatorProof());
        withdrawalToProve = _getWithdrawalProof();  
        withdrawalFields = getWithdrawalFields();
    }

    /// @notice this function just generates a valid proof so that we can test other functionalities of the withdrawal flow
    function _getWithdrawalProof() internal returns (BeaconChainProofs.WithdrawalProof memory) {
        {
            bytes32 blockRoot = getBlockRoot();
            bytes32 slotRoot = getSlotRoot();
            bytes32 timestampRoot = getTimestampRoot();
            bytes32 executionPayloadRoot = getExecutionPayloadRoot();
            bytes memory withdrawalProof = IS_DENEB ? abi.encodePacked(getWithdrawalProofDeneb()) : abi.encodePacked(getWithdrawalProofCapella());
            bytes memory timestampProof = IS_DENEB ? abi.encodePacked(getTimestampProofDeneb()) : abi.encodePacked(getTimestampProofCapella());
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

    ///@notice Effective balance is > 32 ETH
    modifier setWithdrawalCredentialsExcess() {
        // Set JSON and params
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        // Set beacon state root, validatorIndex
        beaconStateRoot = getBeaconStateRoot();
        uint40 validatorIndex = uint40(getValidatorIndex());
        validatorFieldsProof = abi.encodePacked(getWithdrawalCredentialProof()); // Validator fields are proven here
        validatorFields = getValidatorFields();

        // Get an oracle timestamp
        cheats.warp(GOERLI_GENESIS_TIME + 1 days);
        uint64 oracleTimestamp = uint64(block.timestamp);

        eigenPodHarness.verifyWithdrawalCredentials(
            oracleTimestamp,
            beaconStateRoot,
            validatorIndex,
            validatorFieldsProof,
            validatorFields
        );
        _;
    }
}
