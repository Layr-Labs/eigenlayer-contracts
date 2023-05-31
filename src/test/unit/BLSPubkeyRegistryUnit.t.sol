//SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "forge-std/Test.sol";
import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/interfaces/IRegistryCoordinator.sol";
import "../mocks/PublicKeyCompendiumMock.sol";
import "../mocks/RegistryCoordinatorMock.sol";


contract BLSPubkeyRegistryUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    address defaultOperator = address(4545);

    bytes32 internal constant ZERO_PK_HASH = hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";



    BLSPubkeyRegistry public blsPubkeyRegistry;
    BLSPublicKeyCompendiumMock public pkCompendium;
    RegistryCoordinatorMock public registryCoordinator;

    BN254.G1Point internal defaultPubKey;
    uint8 internal defaulQuorumNumber = 0;

    function setUp() external {
        registryCoordinator = new RegistryCoordinatorMock();
        pkCompendium = new BLSPublicKeyCompendiumMock();
        blsPubkeyRegistry = new BLSPubkeyRegistry(registryCoordinator, pkCompendium);

        defaultPubKey = BN254.G1Point(1, 1);
    }

    function testConstructorArgs() public {
        require(blsPubkeyRegistry.registryCoordinator() == registryCoordinator, "registryCoordinator not set correctly");
        require(blsPubkeyRegistry.pubkeyCompendium() == pkCompendium, "pubkeyCompendium not set correctly");
    }

    function testCallRegisterOperatorFromNonCoordinatorAddress(address nonCoordinatorAddress) public {
        cheats.assume(nonCoordinatorAddress != address(registryCoordinator));

        cheats.startPrank(nonCoordinatorAddress);
        cheats.expectRevert(bytes("BLSPubkeyRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        blsPubkeyRegistry.registerOperator(nonCoordinatorAddress, new uint8[](0), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }

    function testCallDeregisterOperatorFromNonCoordinatorAddress(address nonCoordinatorAddress) public {
        cheats.assume(nonCoordinatorAddress != address(registryCoordinator));

        cheats.startPrank(nonCoordinatorAddress);
        cheats.expectRevert(bytes("BLSPubkeyRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        blsPubkeyRegistry.deregisterOperator(nonCoordinatorAddress, new uint8[](0), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }

    function testOperatorDoesNotOwnPubKeyRegister(address operator) public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: operator does not own pubkey"));
        blsPubkeyRegistry.registerOperator(operator, new uint8[](1), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }
    function testOperatorDoesNotOwnPubKeyDeregister(address operator) public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._deregisterOperator: operator does not own pubkey"));
        blsPubkeyRegistry.deregisterOperator(operator, new uint8[](1), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testOperatorRegisterZeroPubkey() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: cannot register zero pubkey"));
        blsPubkeyRegistry.registerOperator(defaultOperator, new uint8[](1), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }
    function testRegisteringWithNoQuorums() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: must register for at least one quorum"));
        blsPubkeyRegistry.registerOperator(defaultOperator, new uint8[](0), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testDeregisteringWithNoQuorums() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._deregisterOperator: must register for at least one quorum"));
        blsPubkeyRegistry.deregisterOperator(defaultOperator, new uint8[](0), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testRegisterOperatorBLSPubkey(address operator) public {
        bytes32 pkHash = BN254.hashG1Point(defaultPubKey);

        cheats.startPrank(operator);
        pkCompendium.registerPublicKey(defaultPubKey);
        cheats.stopPrank();

        //register for one quorum
        uint8[] memory quorumNumbers = new uint8[](1);
        quorumNumbers[0] = defaulQuorumNumber;
        
        cheats.startPrank(address(registryCoordinator));
        bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, defaultPubKey);
        cheats.stopPrank();

        require(registeredpkHash == pkHash, "registeredpkHash not set correctly");
    }

    function testQuorumApkUpdates(uint8[] memory quorumNumbers) public {
        cheats.assume(quorumNumbers.length > 0);
        bytes32 pkHash = BN254.hashG1Point(defaultPubKey);

        BN254.G1Point[] memory quorumApksBefore = new BN254.G1Point[](quorumNumbers.length);
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApksBefore[i] = blsPubkeyRegistry.quorumApk(quorumNumbers[i]);
        }

        cheats.startPrank(defaultOperator);
        pkCompendium.registerPublicKey(defaultPubKey);
        cheats.stopPrank();
        
        cheats.startPrank(address(registryCoordinator));
        blsPubkeyRegistry.registerOperator(defaultOperator, quorumNumbers, defaultPubKey);
        cheats.stopPrank();

        //check quorum apk updates
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            BN254.G1Point memory quorumApkAfter = blsPubkeyRegistry.quorumApk(quorumNumbers[i]);
            require(BN254.hashG1Point(BN254.plus(quorumApkAfter, BN254.negate(quorumApksBefore[i]))) == BN254.hashG1Point(defaultPubKey), "quorum apk not updated correctly");
        }
    }

    // function testRegisterWithNegativeGlobalApk(address operator) external {
    //     testRegisterOperatorBLSPubkey(operator);

    //     BN254.G1Point memory globalApk = blsPubkeyRegistry.globalApk();


    //     BN254.G1Point memory negatedGlobalApk = BN254.negate(globalApk);

    //     //register for one quorum
    //     uint8[] memory quorumNumbers = new uint8[](1);
    //     quorumNumbers[0] = defaulQuorumNumber;
        
    //     cheats.startPrank(address(registryCoordinator));
    //     bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, negatedGlobalApk);
    //     cheats.stopPrank();

    //     BN254.G1Point memory zeroPk = BN254.G1Point(0,0);

    //     require(BN254.hashG1Point(blsPubkeyRegistry.globalApk()) == ZERO_PK_HASH, "globalApk not set correctly");
    // }

    // function testRegisterWithNegativeQuorumApk(address operator) external {
    //     testRegisterOperatorBLSPubkey(defaultOperator);

    //     BN254.G1Point memory quorumApk = blsPubkeyRegistry.quorumApk(defaulQuorumNumber);

    //     BN254.G1Point memory negatedQuorumApk = BN254.negate(quorumApk);

    //     //register for one quorum
    //     uint8[] memory quorumNumbers = new uint8[](1);
    //     quorumNumbers[0] = defaulQuorumNumber;

    //     cheats.startPrank(address(registryCoordinator));
    //     bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, negatedQuorumApk);
    //     cheats.stopPrank();

    //     BN254.G1Point memory zeroPk = BN254.G1Point(0,0);
    //     require(BN254.hashG1Point(blsPubkeyRegistry.quorumApk(defaulQuorumNumber)) == ZERO_PK_HASH, "quorumApk not set correctly");
    // }
    
    // function testQuorumApkUpdatesDeregistration(uint8[] memory quorumNumbers) external {
    //     testQuorumApkUpdates(quorumNumbers);

    //     BN254.G1Point[] memory quorumApksBefore = new BN254.G1Point[](quorumNumbers.length);
    //     for(uint8 i = 0; i < quorumNumbers.length; i++){
    //         quorumApksBefore[i] = blsPubkeyRegistry.quorumApk(quorumNumbers[i]);
    //     }

    //     cheats.startPrank(address(registryCoordinator));
    //     blsPubkeyRegistry.deregisterOperator(defaultOperator, quorumNumbers, defaultPubKey);
    //     cheats.stopPrank();


    //     for(uint8 i = 0; i < quorumNumbers.length; i++){
    //         BN254.G1Point memory quorumApkAfter = blsPubkeyRegistry.quorumApk(quorumNumbers[i]);
    //         require(BN254.hashG1Point(BN254.plus(quorumApksBefore[i], BN254.negate(quorumApkAfter))) == BN254.hashG1Point(defaultPubKey), "quorum apk not updated correctly");
    //     }
    // }

    // function testECADD() public {

    //     BN254.G1Point memory zeroPk = BN254.G1Point(0,0);
    //     BN254.G1Point memory onePK = BN254.G1Point(0,0);

    //     uint256[4] memory input;
    //     BN254.G1Point memory output;
    //     input[0] = zeroPk.X;
    //     input[1] = zeroPk.Y;
    //     input[2] = onePK.X;
    //     input[3] = onePK.Y;
    //     bool success;

    //     // solium-disable-next-line security/no-inline-assembly
    //     assembly {
    //         success := staticcall(sub(gas(), 2000), 6, input, 0x80, output, 0x40)
    //         // Use "invalid" to make gas estimation work
    //         switch success
    //         case 0 {
    //             invalid()
    //         }
    //     }
    //     require(success, "ec-add-failed");
    // }


}