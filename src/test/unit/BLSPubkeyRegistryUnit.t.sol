//SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "forge-std/Test.sol";
import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/interfaces/IRegistryCoordinator.sol";
import "../mocks/PublicKeyCompendiumMock.sol";
import "../mocks/RegistryCoordinatorMock.sol";


contract BLSPubkeyRegistryUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    address operator = address(4545);


    BLSPubkeyRegistry public blsPubkeyRegistry;
    BLSPublicKeyCompendiumMock public pkCompendium;
    RegistryCoordinatorMock public registryCoordinator;

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
        blsPubkeyRegistry.registerOperator(nonCoordinatorAddress, new uint8[](0), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }

    function testOperatorDoesNotOwnPubKey(address operator) public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: operator does not own pubkey"));
        blsPubkeyRegistry.registerOperator(operator, new uint8[](1), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testOperatorRegisterZeroPubkey() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: cannot register zero pubkey"));
        blsPubkeyRegistry.registerOperator(operator, new uint8[](1), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }
    function testRegisteringWithNoQuorums() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSRegistry._registerOperator: must register for at least one quorum"));
        blsPubkeyRegistry.registerOperator(operator, new uint8[](0), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testRegisterOperatorBLSPubkey(address operator) public {
        BN254.G1Point memory pk = BN254.G1Point(1, 1);
        bytes32 pkHash = BN254.hashG1Point(pk);

        cheats.startPrank(operator);
        pkCompendium.registerPublicKey(pk);
        cheats.stopPrank();

        //register for one quorum
        uint8[] memory quorumNumbers = new uint8[](1);
        quorumNumbers[0] = 1;
        
        cheats.startPrank(address(registryCoordinator));
        bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, pk);
        cheats.stopPrank();

        require(registeredpkHash == pkHash, "registeredpkHash not set correctly");
    }

    function testQuorumApkUpdates(uint8[] quorumNumbers) public {
        BN254.G1Point memory pk = BN254.G1Point(1, 1);
        bytes32 pkHash = BN254.hashG1Point(pk);

        BN254.G1Point[] quorumApksBefore = new BN254.G1Point[](quorumNumbers.length);
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApksBefore[i] = blsPubkeyRegistry.quorumApk(quorumNumbers[i]);
        }

        cheats.startPrank(operator);
        pkCompendium.registerPublicKey(pk);
        cheats.stopPrank();
        
        cheats.startPrank(address(registryCoordinator));
        blsPubkeyRegistry.registerOperator(operator, quorumNumbers, pk);
        cheats.stopPrank();

        //check quorum apk updates
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApkAfter = blsPubkeyRegistry.quorumApk(quorumNumbers[i]);
            require(BN254.plus(quorumApkAfter, BN254.negate(quorumApksBefore[i])) == pk, "quorum apk not updated correctly");
        }
    }

    function testRegisterWithNegativeGlobalApk(address operator) external {
        BN254.G1Point memory pk = BN254.G1Point(1, 1);
        testRegisterOperatorBLSPubkey(operator);

        BN254.G1Point memory globalApk = blsPubkeyRegistry.globalApk();


        BN254.G1Point memory negatedGlobalApk = BN254.negate(globalApk);

        //register for one quorum
        uint8[] memory quorumNumbers = new uint8[](1);
        quorumNumbers[0] = 1;
        
        cheats.startPrank(address(registryCoordinator));
        bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, negatedGlobalApk);
        cheats.stopPrank();

        require(blsPubkeyRegistry.globalApk() == BN254.G1Point(0, 0), "globalApk not set correctly");
    }


}