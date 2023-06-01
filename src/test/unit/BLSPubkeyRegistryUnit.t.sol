//SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "forge-std/Test.sol";
import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/interfaces/IRegistryCoordinator.sol";
import "../mocks/PublicKeyCompendiumMock.sol";
import "../mocks/RegistryCoordinatorMock.sol";


contract BLSPubkeyRegistryUnitTests is Test {
    using BN254 for BN254.G1Point;
    Vm cheats = Vm(HEVM_ADDRESS);

    address defaultOperator = address(4545);

    bytes32 internal constant ZERO_PK_HASH = hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";



    BLSPubkeyRegistry public blsPubkeyRegistry;
    BLSPublicKeyCompendiumMock public pkCompendium;
    RegistryCoordinatorMock public registryCoordinator;

    BN254.G1Point internal defaultPubKey =  BN254.G1Point(18260007818883133054078754218619977578772505796600400998181738095793040006897,3432351341799135763167709827653955074218841517684851694584291831827675065899);

    uint8 internal defaulQuorumNumber = 0;

    function setUp() external {
        registryCoordinator = new RegistryCoordinatorMock();
        pkCompendium = new BLSPublicKeyCompendiumMock();
        blsPubkeyRegistry = new BLSPubkeyRegistry(registryCoordinator, pkCompendium);
    }

    function testConstructorArgs() public {
        require(blsPubkeyRegistry.registryCoordinator() == registryCoordinator, "registryCoordinator not set correctly");
        require(blsPubkeyRegistry.pubkeyCompendium() == pkCompendium, "pubkeyCompendium not set correctly");
    }

    function testCallRegisterOperatorFromNonCoordinatorAddress(address nonCoordinatorAddress) public {
        cheats.assume(nonCoordinatorAddress != address(registryCoordinator));

        cheats.startPrank(nonCoordinatorAddress);
        cheats.expectRevert(bytes("BLSPubkeyRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        blsPubkeyRegistry.registerOperator(nonCoordinatorAddress, new bytes(0), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }

    function testCallDeregisterOperatorFromNonCoordinatorAddress(address nonCoordinatorAddress) public {
        cheats.assume(nonCoordinatorAddress != address(registryCoordinator));

        cheats.startPrank(nonCoordinatorAddress);
        cheats.expectRevert(bytes("BLSPubkeyRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        blsPubkeyRegistry.deregisterOperator(nonCoordinatorAddress, new bytes(0), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }

    function testOperatorDoesNotOwnPubKeyRegister(address operator) public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: operator does not own pubkey"));
        blsPubkeyRegistry.registerOperator(operator, new bytes(1), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }
    function testOperatorDoesNotOwnPubKeyDeregister(address operator) public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._deregisterOperator: operator does not own pubkey"));
        blsPubkeyRegistry.deregisterOperator(operator, new bytes(1), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testOperatorRegisterZeroPubkey() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: cannot register zero pubkey"));
        blsPubkeyRegistry.registerOperator(defaultOperator, new bytes(1), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }
    function testRegisteringWithNoQuorums() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: must register for at least one quorum"));
        blsPubkeyRegistry.registerOperator(defaultOperator, new bytes(0), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testDeregisteringWithNoQuorums() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._deregisterOperator: must register for at least one quorum"));
        blsPubkeyRegistry.deregisterOperator(defaultOperator, new bytes(0), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testRegisterOperatorBLSPubkey(address operator, uint256 x, uint256 y) public {
        cheats.assume(x != 0 && y != 0);
        BN254.G1Point memory pubkey = BN254.G1Point(x, y);
        bytes32 pkHash = BN254.hashG1Point(defaultPubKey);

        cheats.startPrank(operator);
        pkCompendium.registerPublicKey(defaultPubKey);
        cheats.stopPrank();

        //register for one quorum
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaulQuorumNumber);
        
        cheats.startPrank(address(registryCoordinator));
        bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, defaultPubKey);
        cheats.stopPrank();

        require(registeredpkHash == pkHash, "registeredpkHash not set correctly");
    }

    function testQuorumApkUpdates(uint8 quorumNumber1, uint8 quorumNumber2) public {
        cheats.assume(quorumNumber1 != quorumNumber2);

        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(quorumNumber1);
        quorumNumbers[1] = bytes1(quorumNumber2);

        bytes32 pkHash = BN254.hashG1Point(defaultPubKey);

        BN254.G1Point[] memory quorumApksBefore = new BN254.G1Point[](quorumNumbers.length);
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApksBefore[i] = blsPubkeyRegistry.quorumApk(uint8(quorumNumbers[i]));
        }

        cheats.startPrank(defaultOperator);
        pkCompendium.registerPublicKey(defaultPubKey);
        cheats.stopPrank();
        
        cheats.startPrank(address(registryCoordinator));
        blsPubkeyRegistry.registerOperator(defaultOperator, quorumNumbers, defaultPubKey);
        cheats.stopPrank();

        //check quorum apk updates
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            BN254.G1Point memory quorumApkAfter = blsPubkeyRegistry.quorumApk(uint8(quorumNumbers[i]));
            bytes32 temp = BN254.hashG1Point(BN254.plus(quorumApkAfter, BN254.negate(quorumApksBefore[i])));
            require(temp == BN254.hashG1Point(defaultPubKey), "quorum apk not updated correctly");
        }
    }

    function testRegisterWithNegativeGlobalApk(address operator) external {
        testRegisterOperatorBLSPubkey(operator, defaultPubKey.X, defaultPubKey.Y);

        (uint256 x, uint256 y)= blsPubkeyRegistry.globalApk();
        BN254.G1Point memory globalApk = BN254.G1Point(x, y);


        BN254.G1Point memory negatedGlobalApk = BN254.negate(globalApk);

        //register for one quorum
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaulQuorumNumber);

        cheats.startPrank(operator);
        pkCompendium.registerPublicKey(negatedGlobalApk);
        cheats.stopPrank();
        
        cheats.startPrank(address(registryCoordinator));
        bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, negatedGlobalApk);
        cheats.stopPrank();

        BN254.G1Point memory zeroPk = BN254.G1Point(0,0);

        (x, y)= blsPubkeyRegistry.globalApk();
        BN254.G1Point memory temp = BN254.G1Point(x, y);

        require(BN254.hashG1Point(temp) == ZERO_PK_HASH, "globalApk not set correctly");
    }

    function testRegisterWithNegativeQuorumApk(address operator) external {
        testRegisterOperatorBLSPubkey(defaultOperator, defaultPubKey.X, defaultPubKey.Y);

        BN254.G1Point memory quorumApk = blsPubkeyRegistry.quorumApk(defaulQuorumNumber);

        BN254.G1Point memory negatedQuorumApk = BN254.negate(quorumApk);

        //register for one quorum
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaulQuorumNumber);

        cheats.startPrank(operator);
        pkCompendium.registerPublicKey(negatedQuorumApk);
        cheats.stopPrank();

        cheats.startPrank(address(registryCoordinator));
        bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, negatedQuorumApk);
        cheats.stopPrank();

        BN254.G1Point memory zeroPk = BN254.G1Point(0,0);
        require(BN254.hashG1Point(blsPubkeyRegistry.quorumApk(defaulQuorumNumber)) == ZERO_PK_HASH, "quorumApk not set correctly");
    }
    
    function testQuorumApkUpdatesDeregistration(uint8 quorumNumber1, uint8 quorumNumber2) external {
        cheats.assume(quorumNumber1 != quorumNumber2);
        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(quorumNumber1);
        quorumNumbers[1] = bytes1(quorumNumber2);

        testQuorumApkUpdates(quorumNumber1, quorumNumber2);

        BN254.G1Point[] memory quorumApksBefore = new BN254.G1Point[](2);
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApksBefore[i] = blsPubkeyRegistry.quorumApk(uint8(quorumNumbers[i]));
        }

        cheats.startPrank(address(registryCoordinator));
        blsPubkeyRegistry.deregisterOperator(defaultOperator, quorumNumbers, defaultPubKey);
        cheats.stopPrank();

        
        BN254.G1Point memory quorumApkAfter;
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApkAfter = blsPubkeyRegistry.quorumApk(uint8(quorumNumbers[i]));
            require(BN254.hashG1Point(quorumApksBefore[i].plus(defaultPubKey.negate())) == BN254.hashG1Point(quorumApkAfter), "quorum apk not updated correctly");
        }
    }
}