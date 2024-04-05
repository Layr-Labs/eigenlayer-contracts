// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodPausingConstants.sol";

import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/utils/ProofParsing.sol";
import "src/test/harnesses/EigenPodManagerWrapper.sol";
import "src/test/mocks/EigenPodMock.sol";
import "src/test/mocks/Dummy.sol";
import "src/test/mocks/ETHDepositMock.sol";
import "src/test/mocks/DelayedWithdrawalRouterMock.sol";
import "src/test/mocks/BeaconChainOracleMock.sol";
import "src/test/mocks/Reenterer.sol";
import "src/test/events/IEigenPodEvents.sol";
import "src/test/events/IEigenPodManagerEvents.sol";

contract EigenPod_PodManager_UnitTests is EigenLayerUnitTestSetup {
    // Contracts Under Test: EigenPodManager & EigenPod
    EigenPod public eigenPod;
    EigenPod public podImplementation;
    IBeacon public eigenPodBeacon;
    EigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    EigenPodManagerWrapper public eigenPodManagerWrapper; // Implementation contract

    // Mocks
    IETHPOSDeposit public ethPOSMock;
    IDelayedWithdrawalRouter public delayedWithdrawalRouterMock;
    BeaconChainOracleMock beaconChainOracle;
    
    // Constants
    uint64 public constant MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32e9;
    uint64 public constant GOERLI_GENESIS_TIME = 1616508000;
    address public initialOwner = address(this);

    // Owner for which proofs are generated; eigenPod above is owned by this address
    bytes internal constant beaconProxyBytecode =
        hex"608060405260405161090e38038061090e83398101604081905261002291610460565b61002e82826000610035565b505061058a565b61003e83610100565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100fb576100f9836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e99190610520565b836102a360201b6100291760201c565b505b505050565b610113816102cf60201b6100551760201c565b6101725760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101e6816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d79190610520565b6102cf60201b6100551760201c565b61024b5760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610169565b806102827fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6102de60201b6100641760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606102c883836040518060600160405280602781526020016108e7602791396102e1565b9392505050565b6001600160a01b03163b151590565b90565b6060600080856001600160a01b0316856040516102fe919061053b565b600060405180830381855af49150503d8060008114610339576040519150601f19603f3d011682016040523d82523d6000602084013e61033e565b606091505b5090925090506103508683838761035a565b9695505050505050565b606083156103c65782516103bf576001600160a01b0385163b6103bf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610169565b50816103d0565b6103d083836103d8565b949350505050565b8151156103e85781518083602001fd5b8060405162461bcd60e51b81526004016101699190610557565b80516001600160a01b038116811461041957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561044f578181015183820152602001610437565b838111156100f95750506000910152565b6000806040838503121561047357600080fd5b61047c83610402565b60208401519092506001600160401b038082111561049957600080fd5b818501915085601f8301126104ad57600080fd5b8151818111156104bf576104bf61041e565b604051601f8201601f19908116603f011681019083821181831017156104e7576104e761041e565b8160405282815288602084870101111561050057600080fd5b610511836020830160208801610434565b80955050505050509250929050565b60006020828403121561053257600080fd5b6102c882610402565b6000825161054d818460208701610434565b9190910192915050565b6020815260008251806020840152610576816040850160208701610434565b601f01601f19169190910160400192915050565b61034e806105996000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610067565b610100565b565b606061004e83836040518060600160405280602781526020016102f260279139610124565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fb9190610249565b905090565b3660008037600080366000845af43d6000803e80801561011f573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161014191906102a2565b600060405180830381855af49150503d806000811461017c576040519150601f19603f3d011682016040523d82523d6000602084013e610181565b606091505b50915091506101928683838761019c565b9695505050505050565b6060831561020d578251610206576001600160a01b0385163b6102065760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610217565b610217838361021f565b949350505050565b81511561022f5781518083602001fd5b8060405162461bcd60e51b81526004016101fd91906102be565b60006020828403121561025b57600080fd5b81516001600160a01b038116811461004e57600080fd5b60005b8381101561028d578181015183820152602001610275565b8381111561029c576000848401525b50505050565b600082516102b4818460208701610272565b9190910192915050565b60208152600082518060208401526102dd816040850160208701610272565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d51e81d3bc5ed20a26aeb05dce7e825c503b2061aa78628027300c8d65b9d89a64736f6c634300080c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564";
    address public constant podOwner = address(42000094993494);

    address public constant podAddress = address(0x49c486E3f4303bc11C02F952Fe5b08D0AB22D443);
    
    function setUp() public override virtual {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy Mocks
        ethPOSMock = new ETHPOSDepositMock();
        delayedWithdrawalRouterMock = new DelayedWithdrawalRouterMock();
        beaconChainOracle = new BeaconChainOracleMock();

        // Deploy proxy contract for EPM
        EmptyContract emptyContract = new EmptyContract();
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Deploy EigenPod Implementation and beacon
        podImplementation = new EigenPod(
            ethPOSMock,
            delayedWithdrawalRouterMock,
            eigenPodManager,
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );

        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));

        // Deploy EigenPodManager implementation
        eigenPodManagerWrapper = new EigenPodManagerWrapper(
            ethPOSMock,
            eigenPodBeacon,
            strategyManagerMock,
            slasherMock,
            delegationManagerMock
        );

        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerWrapper),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                beaconChainOracle,
                initialOwner,
                pauserRegistry,
                0 /*initialPausedStatus*/
            )
        );

        // Below is a hack to get the eigenPod address that proofs prove against
        
        // Deploy Proxy same way as EigenPodManager does
        eigenPod = EigenPod(payable(
            Create2.deploy(
                0,
                bytes32(uint256(uint160(address(podOwner)))),
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
        eigenPod.initialize(address(podOwner));

        // Set storage in EPM
        EigenPodManagerWrapper(address(eigenPodManager)).setPodAddress(podOwner, eigenPod);

        eigenPodManager.setDenebForkTimestamp(type(uint64).max);
    }
}

contract EigenPod_PodManager_UnitTests_EigenPodPausing is EigenPod_PodManager_UnitTests {
    /**
     * 1. verifyBalanceUpdates revert when PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE set
     * 2. verifyAndProcessWithdrawals revert when PAUSED_EIGENPODS_VERIFY_WITHDRAWAL set
     * 3. verifyWithdrawalCredentials revert when PAUSED_EIGENPODS_VERIFY_CREDENTIALS set
     * 4. activateRestaking revert when PAUSED_EIGENPODS_VERIFY_CREDENTIALS set
     */

    /// @notice Index for flag that pauses creation of new EigenPods when set. See EigenPodManager code for details.
    uint8 internal constant PAUSED_NEW_EIGENPODS = 0;
    /// @notice Index for flag that pauses all withdrawal-of-restaked ETH related functionality `function *of the EigenPodManager* when set. See EigenPodManager code for details.
    uint8 internal constant PAUSED_WITHDRAW_RESTAKED_ETH = 1;

    /// @notice Index for flag that pauses the deposit related functions *of the EigenPods* when set. see EigenPod code for details.
    uint8 internal constant PAUSED_EIGENPODS_VERIFY_CREDENTIALS = 2;
    /// @notice Index for flag that pauses the `verifyBalanceUpdate` function *of the EigenPods* when set. see EigenPod code for details.
    uint8 internal constant PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE = 3;
    /// @notice Index for flag that pauses the `verifyBeaconChainFullWithdrawal` function *of the EigenPods* when set. see EigenPod code for details.
    uint8 internal constant PAUSED_EIGENPODS_VERIFY_WITHDRAWAL = 4;

    function test_verifyBalanceUpdates_revert_pausedEigenVerifyBalanceUpdate() public {
        BeaconChainProofs.StateRootProof memory stateRootProofStruct;

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        bytes[] memory proofsArray = new bytes[](1);
        uint40[] memory validatorIndices = new uint40[](1);

        // pause the contract
        cheats.prank(address(pauser));
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE);

        cheats.prank(address(podOwner));
        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        eigenPod.verifyBalanceUpdates(0, validatorIndices, stateRootProofStruct, proofsArray, validatorFieldsArray);
    }

    function test_verifyAndProcessWithdrawals_revert_pausedEigenVerifyWithdrawal() public {
        BeaconChainProofs.StateRootProof memory stateRootProofStruct;
        BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray;

        bytes[] memory validatorFieldsProofArray = new bytes[](1);
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);

        // pause the contract
        cheats.prank(address(pauser));
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_WITHDRAWAL);

        cheats.prank(address(podOwner));
        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        eigenPod.verifyAndProcessWithdrawals(
            0,
            stateRootProofStruct,
            withdrawalProofsArray,
            validatorFieldsProofArray,
            validatorFieldsArray,
            withdrawalFieldsArray
        );
    }

    function test_verifyWithdrawalCredentials_revert_pausedEigenVerifyCredentials() public {
        BeaconChainProofs.StateRootProof memory stateRootProofStruct;

        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        bytes[] memory proofsArray = new bytes[](1);
        uint40[] memory validatorIndices = new uint40[](1);

        // pause the contract
        cheats.prank(address(pauser));
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_CREDENTIALS);

        cheats.prank(address(podOwner));
        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        eigenPod.verifyWithdrawalCredentials(
            0,
            stateRootProofStruct,
            validatorIndices,
            proofsArray,
            validatorFieldsArray
        );
    }

    function test_activateRestaking_revert_pausedEigenVerifyCredentials() public {
        // pause the contract
        cheats.prank(address(pauser));
        eigenPodManager.pause(2 ** PAUSED_EIGENPODS_VERIFY_CREDENTIALS);

        cheats.prank(address(podOwner));
        cheats.expectRevert(bytes("EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"));
        eigenPod.activateRestaking();
    }
}

contract EigenPod_PodManager_UnitTests_EigenPod is EigenPod_PodManager_UnitTests {
    /**
     * @notice Tests function calls from EPM to EigenPod
     * 1. Stake works when pod is deployed
     * 2. Stake when pod is not deployed -> check that ethPOS deposit contract is correct for this and above test
     */

    bytes public constant pubkey = hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";

    function test_stake_podAlreadyDeployed(bytes memory signature, bytes32 depositDataRoot) public {
        uint256 stakeAmount = 32e18;

        uint256 numPods = eigenPodManager.numPods();
        emit log_named_uint("numPods", numPods);

        cheats.startPrank(podOwner);
        cheats.deal(podOwner, stakeAmount);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
    }

    function test_stake_podNotDeployed(bytes memory signature, bytes32 depositDataRoot) public {
        address newPodOwner = address(69696969696);
    
        uint256 stakeAmount = 32e18;

        cheats.startPrank(newPodOwner);
        cheats.deal(newPodOwner, stakeAmount);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
    }
}

contract EigenPod_PodManager_UnitTests_EigenPodManager is EigenPod_PodManager_UnitTests, ProofParsing, IEigenPodEvents {
     /**
     * @notice Tests function calls from EigenPod to EigenPodManager
     * 1. Verify withdrawal credentials and call `recordBeaconChainETHBalanceUpdate` -> assert shares are updated
     * 2. Do a full withdrawal and call `recordBeaconChainETHBalanceUpdate` -> assert shares are updated
     * 3. Do a partial withdrawal and call `recordBeaconChainETHBalanceUpdate` -> assert shares are updated
     * 4. Verify balance updates and call `recordBeaconChainEThBalanceUpdate` -> assert shares are updated
     * 5. Withdraw restaked beacon chain ETH
     */

    using BeaconChainProofs for *;

    // Params to verify withdrawal credentials
    BeaconChainProofs.StateRootProof stateRootProofStruct; 
    uint40[] validatorIndices;
    bytes[] validatorFieldsProofs;
    bytes32[][] validatorFields;
    // BeaconChainProofs.BalanceUpdateProof[] balanceUpdateProof;
    BeaconChainProofs.WithdrawalProof[] withdrawalProofs;
    bytes32[][] withdrawalFields;

    function test_verifyWithdrawalCredentials() public {    
        // Arrange: Set up conditions to verify withdrawal credentials
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _setWithdrawalCredentialParams();

        // Set oracle block root and warp time
        _setOracleBlockRoot();
        cheats.warp(GOERLI_GENESIS_TIME + 1 days);
        uint64 oracleTimestamp = uint64(block.timestamp);

        // Save state for checks 
        int256 initialShares = eigenPodManager.podOwnerShares(podOwner);

        // Act: Verify withdrawal credentials and record the balance update
        cheats.prank(podOwner);
        eigenPod.verifyWithdrawalCredentials(
            oracleTimestamp,
            stateRootProofStruct,
            validatorIndices,
            validatorFieldsProofs,
            validatorFields
        );

        // Assert: Check that the shares are updated correctly
        int256 updatedShares = eigenPodManager.podOwnerShares(podOwner);
        assertTrue(updatedShares != initialShares, "Shares should be updated after verifying withdrawal credentials");
        assertEq(updatedShares, 32e18, "Shares should be 32ETH in wei after verifying withdrawal credentials");
    }

    function test_balanceUpdate_negativeSharesDelta() public {
        // Verify withdrawal credentials
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _verifyWithdrawalCredentials();

        // Set JSON
        setJSON("src/test/test-data/balanceUpdateProof_balance28ETH_302913.json");
        bytes32 validatorPubkeyHash = validatorFields[0].getPubkeyHash();

        // Set proof params, oracle block root, and warp time
        _setBalanceUpdateParams();
        _setOracleBlockRoot();
        cheats.warp(GOERLI_GENESIS_TIME);
        uint64 oracleTimestamp = uint64(block.timestamp);

        // Save state for checks
        int256 initialShares = eigenPodManager.podOwnerShares(podOwner);
        uint64 newValidatorBalance = validatorFields[0].getEffectiveBalanceGwei();

        // Verify balance update
        eigenPod.verifyBalanceUpdates(
            oracleTimestamp,
            validatorIndices,
            stateRootProofStruct,
            validatorFieldsProofs,
            validatorFields
        );

        // Checks
        int256 updatedShares = eigenPodManager.podOwnerShares(podOwner);
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPod.validatorPubkeyHashToInfo(validatorPubkeyHash);
        assertEq(validatorInfo.restakedBalanceGwei, newValidatorBalance, "Restaked balance gwei is incorrect");
        assertLt(updatedShares - initialShares, 0, "Shares delta should be negative");
        int256 expectedSharesDiff = (int256(uint256(newValidatorBalance)) - int256(uint256(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR))) * 1e9;
        assertEq(updatedShares - initialShares, expectedSharesDiff, "Shares delta should be equal to restaked balance");
    }

    function test_balanceUpdate_positiveSharesDelta() public {
        // Verify withdrawal credentials
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913_30ETHBalance.json");
        _verifyWithdrawalCredentials();

        // Set JSON
        setJSON("src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        bytes32 validatorPubkeyHash = validatorFields[0].getPubkeyHash();

        // Set proof params, oracle block root, and warp time
        _setBalanceUpdateParams();
        _setOracleBlockRoot();
        cheats.warp(GOERLI_GENESIS_TIME);
        uint64 oracleTimestamp = uint64(block.timestamp);

        // Save state for checks
        int256 initialShares = eigenPodManager.podOwnerShares(podOwner);
        uint64 newValidatorBalance = validatorFields[0].getEffectiveBalanceGwei();

        // Verify balance update
        eigenPod.verifyBalanceUpdates(
            oracleTimestamp,
            validatorIndices,
            stateRootProofStruct,
            validatorFieldsProofs,
            validatorFields
        );

        // Checks
        int256 updatedShares = eigenPodManager.podOwnerShares(podOwner);
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPod.validatorPubkeyHashToInfo(validatorPubkeyHash);
        assertEq(validatorInfo.restakedBalanceGwei, MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR, "Restaked balance gwei should be max");
        assertGt(updatedShares - initialShares, 0, "Shares delta should be positive");
        assertEq(updatedShares, 32e18, "Shares should be 32ETH");
        assertEq(newValidatorBalance, 32e9, "validator balance should be 32e9 Gwei");
    }

    function test_fullWithdrawal_excess32ETH() public {
        // Verify withdrawal credentials
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913_30ETHBalance.json");
        _verifyWithdrawalCredentials();

        // Set JSON
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        bytes32 validatorPubkeyHash = validatorFields[0].getPubkeyHash();

        // Set proof params, block root
        _setWithdrawalProofParams();        
        _setOracleBlockRoot();

        // Save state for checks; deal EigenPod withdrawal router balance
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(
            withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]
        );
        uint64 leftOverBalanceWEI = uint64(withdrawalAmountGwei - MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR) * 1e9;
        cheats.deal(address(eigenPod), leftOverBalanceWEI);
        int256 initialShares = eigenPodManager.podOwnerShares(podOwner);

        // Withdraw
        eigenPod.verifyAndProcessWithdrawals(
            0,
            stateRootProofStruct,
            withdrawalProofs,
            validatorFieldsProofs,
            validatorFields,
            withdrawalFields
        );

        // Checks
        int256 updatedShares = eigenPodManager.podOwnerShares(podOwner);
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPod.validatorPubkeyHashToInfo(validatorPubkeyHash);
        assertEq(validatorInfo.restakedBalanceGwei, 0, "Restaked balance gwei should be 0");
        assertGt(updatedShares - initialShares, 0, "Shares diff should be positive");
        int256 expectedSharesDiff = (int256(uint256(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR))*1e9) - initialShares;
        assertEq(updatedShares - initialShares, expectedSharesDiff, "Shares delta incorrect");
        assertEq(updatedShares, 32e18, "Shares should be 32e18");
        assertEq(address(delayedWithdrawalRouterMock).balance, leftOverBalanceWEI, "Incorrect amount sent to delayed withdrawal router");
    }

    function test_withdrawRestakedBeaconChainETH() public {
        test_fullWithdrawal_excess32ETH();

        // Deal eigenPod balance - max restaked balance
        cheats.deal(address(eigenPod), 32 ether);

        cheats.startPrank(address(delegationManagerMock));
        vm.expectEmit(true, true, true, true);
        emit RestakedBeaconChainETHWithdrawn(podOwner, 32 ether);
        eigenPodManager.withdrawSharesAsTokens(
            podOwner,
            podOwner,
            uint256(eigenPodManager.podOwnerShares(podOwner))
        );
        cheats.stopPrank();

        // Checks
        assertEq(address(podOwner).balance, 32 ether, "EigenPod balance should be 0");
        assertEq(address(eigenPod).balance, 0, "EigenPod balance should be 0");
        assertEq(eigenPod.withdrawableRestakedExecutionLayerGwei(), 0, "Restaked execution layer gwei should be 0");
    }

    function test_fullWithdrawal_less32ETH() public {
        // Verify withdrawal credentials
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _verifyWithdrawalCredentials();

        // Set JSON
        setJSON("src/test/test-data/fullWithdrawalProof_Latest_28ETH.json");
        bytes32 validatorPubkeyHash = validatorFields[0].getPubkeyHash();

        // Set proof params, block root
        _setWithdrawalProofParams();        
        _setOracleBlockRoot();

        // Save State
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(
            withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]
        );
        int256 initialShares = eigenPodManager.podOwnerShares(podOwner);

        // Withdraw
        eigenPod.verifyAndProcessWithdrawals(
            0,
            stateRootProofStruct,
            withdrawalProofs,
            validatorFieldsProofs,
            validatorFields,
            withdrawalFields
        );

        // Checks
        int256 updatedShares = eigenPodManager.podOwnerShares(podOwner);
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPod.validatorPubkeyHashToInfo(validatorPubkeyHash);
        assertEq(validatorInfo.restakedBalanceGwei, 0, "Restaked balance gwei is incorrect");
        assertLt(updatedShares - initialShares, 0, "Shares delta should be negative");
        int256 expectedSharesDiff = (int256(uint256(withdrawalAmountGwei)) - int256(uint256(MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR))) * 1e9;
        assertEq(updatedShares - initialShares, expectedSharesDiff, "Shares delta incorrect");
    }

    function test_partialWithdrawal() public {
        // Set JSON & params
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _verifyWithdrawalCredentials();
        
        // Set JSON
        setJSON("./src/test/test-data/partialWithdrawalProof_Latest.json");
        bytes32 validatorPubkeyHash = validatorFields[0].getPubkeyHash();

        // Set proof params, block root
        _setWithdrawalProofParams();        
        _setOracleBlockRoot();

        // Assert that partial withdrawal code path will be tested
        assertLt(withdrawalProofs[0].getWithdrawalEpoch(), validatorFields[0].getWithdrawableEpoch(), "Withdrawal epoch should be less than the withdrawable epoch");

        // Save state for checks; deal EigenPod withdrawal router balance
        uint64 withdrawalAmountGwei = Endian.fromLittleEndianUint64(
            withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]
        );
        cheats.deal(address(eigenPod), withdrawalAmountGwei * 1e9); // deal full withdrawal amount since it's a partial withdrawal
        uint64 initialRestakedBalance = (eigenPod.validatorPubkeyHashToInfo(validatorPubkeyHash)).restakedBalanceGwei;
        int256 initialShares = eigenPodManager.podOwnerShares(podOwner);

        // Withdraw
        eigenPod.verifyAndProcessWithdrawals(
            0,
            stateRootProofStruct,
            withdrawalProofs,
            validatorFieldsProofs,
            validatorFields,
            withdrawalFields
        );

        // Checks
        int256 updatedShares = eigenPodManager.podOwnerShares(podOwner);
        IEigenPod.ValidatorInfo memory validatorInfo = eigenPod.validatorPubkeyHashToInfo(validatorPubkeyHash);
        assertEq(validatorInfo.restakedBalanceGwei, initialRestakedBalance, "Restaked balance gwei should be unchanged");
        assertEq(updatedShares - initialShares, 0, "Shares diff should be 0");
        assertEq(address(delayedWithdrawalRouterMock).balance, withdrawalAmountGwei * 1e9, "Incorrect amount sent to delayed withdrawal router");
    }

    // Helper Functions
    function _getStateRootProof() internal returns (BeaconChainProofs.StateRootProof memory) {
        return BeaconChainProofs.StateRootProof(getBeaconStateRoot(), abi.encodePacked(getStateRootProof()));
    }

    function _setOracleBlockRoot() internal {
        bytes32 latestBlockRoot = getLatestBlockRoot();
        //set beaconStateRoot
        beaconChainOracle.setOracleBlockRootAtTimestamp(latestBlockRoot);
    }

    function _verifyWithdrawalCredentials() internal {
        _setWithdrawalCredentialParams();

        // Set oracle block root and warp time
        uint64 oracleTimestamp = 0;
        _setOracleBlockRoot();
        cheats.warp(oracleTimestamp+=1);

        // Act: Verify withdrawal credentials and record the balance update
        cheats.prank(podOwner);
        eigenPod.verifyWithdrawalCredentials(
            oracleTimestamp,
            stateRootProofStruct,
            validatorIndices,
            validatorFieldsProofs,
            validatorFields
        );
    }

    function _setWithdrawalCredentialParams() internal {
        // Reset arrays
        delete validatorIndices;
        delete validatorFields;
        delete validatorFieldsProofs;

        // Set state proof struct
        stateRootProofStruct = _getStateRootProof();

        // Set validator indices
        uint40 validatorIndex = uint40(getValidatorIndex());
        validatorIndices.push(validatorIndex);

        // Set validatorFieldsArray
        validatorFields.push(getValidatorFields());

        // Set validator fields proof
        validatorFieldsProofs.push(abi.encodePacked(getWithdrawalCredentialProof())); // Validator fields are proven here
    }

    function _setBalanceUpdateParams() internal {
        // Reset arrays
        delete validatorIndices;
        delete validatorFields;
        delete validatorFieldsProofs;
        
        // Set state proof struct
        stateRootProofStruct = _getStateRootProof();

        // Set validator indices
        uint40 validatorIndex = uint40(getValidatorIndex());
        validatorIndices.push(validatorIndex);

        // Set validatorFieldsArray
        validatorFields.push(getValidatorFields());

        // Set validator fields proof
        validatorFieldsProofs.push(abi.encodePacked(getBalanceUpdateProof())); // Validator fields are proven here
    }

    function _setWithdrawalProofParams() internal {
        // Reset arrays
        delete validatorFields;
        delete validatorFieldsProofs;
        delete withdrawalFields;
        delete withdrawalProofs;

        // Set state proof struct
        stateRootProofStruct = _getStateRootProof();
    
        // Set validatorFields
        validatorFields.push(getValidatorFields());

        // Set validator fields proof
        validatorFieldsProofs.push(abi.encodePacked(getValidatorProof()));

        // Set withdrawal fields
        withdrawalFields.push(getWithdrawalFields());

        // Set withdrawal proofs
        withdrawalProofs.push(_getWithdrawalProof());
    }

    // function _getBalanceUpdateProof() internal returns (BeaconChainProofs.BalanceUpdateProof memory) {
    //     bytes32 balanceRoot = getBalanceRoot();
    //     BeaconChainProofs.BalanceUpdateProof memory proofs = BeaconChainProofs.BalanceUpdateProof(
    //         abi.encodePacked(getValidatorBalanceProof()),
    //         abi.encodePacked(getWithdrawalCredentialProof()), //technically this is to verify validator pubkey in the validator fields, but the WC proof is effectively the same so we use it here again.
    //         balanceRoot
    //     );
    //     return proofs;
    // }

    /// @notice this function just generates a valid proof so that we can test other functionalities of the withdrawal flow
    function _getWithdrawalProof() internal returns (BeaconChainProofs.WithdrawalProof memory) {
        {
            bytes32 blockRoot = getBlockRoot();
            bytes32 slotRoot = getSlotRoot();
            bytes32 timestampRoot = getTimestampRoot();
            bytes32 executionPayloadRoot = getExecutionPayloadRoot();

            return
                BeaconChainProofs.WithdrawalProof(
                    abi.encodePacked(getWithdrawalProofCapella()),
                    abi.encodePacked(getSlotProof()),
                    abi.encodePacked(getExecutionPayloadProof()),
                    abi.encodePacked(getTimestampProofCapella()),
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
}

